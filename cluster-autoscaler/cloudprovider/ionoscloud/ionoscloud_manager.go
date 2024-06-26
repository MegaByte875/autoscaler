/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ionoscloud

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	ionos "k8s.io/autoscaler/cluster-autoscaler/cloudprovider/ionoscloud/ionos-cloud-sdk-go"
	"k8s.io/klog/v2"
)

const (
	envKeyClusterID         = "IONOS_CLUSTER_ID"
	envKeyToken             = ionos.IonosTokenEnvVar
	envKeyEndpoint          = ionos.IonosApiUrlEnvVar
	envKeyInsecure          = "IONOS_INSECURE"
	envKeyPollTimeout       = "IONOS_POLL_TIMEOUT"
	envKeyPollInterval      = "IONOS_POLL_INTERVAL"
	envKeyTokensPath        = "IONOS_TOKENS_PATH"
	envKeyAdditionalHeaders = "IONOS_ADDITIONAL_HEADERS"
	defaultTimeout          = 15 * time.Minute
	defaultInterval         = 30 * time.Second
)

// IonosCloudManager handles IonosCloud communication and data caching of node groups.
type IonosCloudManager interface {
	// GetNodeGroupTargetSize gets node group target size.
	GetNodeGroupTargetSize(nodeGroup cloudprovider.NodeGroup) (int, error)
	// GetNodeGroupSize gets node group size.
	GetNodeGroupSize(nodeGroup cloudprovider.NodeGroup) (int, error)
	// SetNodeGroupSize sets the node group size.
	SetNodeGroupSize(nodeGroup cloudprovider.NodeGroup, size int) error
	// DeleteNode deletes a single node.
	DeleteNode(nodeGroup cloudprovider.NodeGroup, nodeId string) error
	// GetInstancesForNodeGroup returns the instances for the given node group.
	GetInstancesForNodeGroup(nodeGroup cloudprovider.NodeGroup) ([]cloudprovider.Instance, error)
	// GetNodeGroupForNode returns the node group that the node belongs to.
	GetNodeGroupForNode(node *apiv1.Node) cloudprovider.NodeGroup
	// TryLockNodeGroup tries to acquire a lock for a node group.
	TryLockNodeGroup(nodeGroup cloudprovider.NodeGroup) bool
	// UnlockNodeGroup releases a node group lock.
	UnlockNodeGroup(nodeGroup cloudprovider.NodeGroup)
	// GetNodeGroups returns the list of managed node groups.
	GetNodeGroups() []cloudprovider.NodeGroup
}

// Config holds information necessary to construct IonosCloud API clients.
type Config struct {
	// ClusterID is the ID of the cluster to autoscale.
	ClusterID string
	// Token is an IonosCloud API access token.
	Token string
	// Endpoint overrides the default API URL.
	Endpoint string
	// Insecure configures the IonosCloud API client to use insecure connection.
	Insecure bool
	// PollTimeout is the timeout for polling a node pool after an update.
	PollTimeout time.Duration
	// PollInterval is the interval in which a node pool is polled after an update.
	PollInterval time.Duration
	// TokensPath points to a directory that contains token files
	TokensPath string
	// AdditionalHeaders are additional HTTP headers to append to each IonosCloud API request.
	AdditionalHeaders map[string]string
}

// LoadConfigFromEnv loads the IonosCloud client config from env.
func LoadConfigFromEnv() (*Config, error) {
	config := &Config{
		Token:        os.Getenv(envKeyToken),
		Endpoint:     os.Getenv(envKeyEndpoint),
		TokensPath:   os.Getenv(envKeyTokensPath),
		PollInterval: defaultInterval,
		PollTimeout:  defaultTimeout,
	}

	if config.ClusterID = os.Getenv(envKeyClusterID); config.ClusterID == "" {
		return nil, fmt.Errorf("missing value for %s", envKeyClusterID)
	}

	if config.Token == "" && config.TokensPath == "" {
		return nil, fmt.Errorf("missing value for either %s or %s", envKeyToken, envKeyTokensPath)
	}

	var err error
	if insecure := os.Getenv(envKeyInsecure); insecure != "" {
		config.Insecure, err = strconv.ParseBool(insecure)
		if err != nil {
			return nil, fmt.Errorf("invalid value for %s: %s", envKeyInsecure, insecure)
		}
	}
	if interval := os.Getenv(envKeyPollInterval); interval != "" {
		config.PollInterval, err = time.ParseDuration(interval)
		if err != nil {
			return nil, fmt.Errorf("invalid value for %s: %s", envKeyPollInterval, interval)
		}
	}
	if timeout := os.Getenv(envKeyPollTimeout); timeout != "" {
		config.PollTimeout, err = time.ParseDuration(timeout)
		if err != nil {
			return nil, fmt.Errorf("invalid value for %s: %s", envKeyPollTimeout, timeout)
		}
	}

	if rawHeaders := os.Getenv(envKeyAdditionalHeaders); rawHeaders != "" {
		config.AdditionalHeaders = make(map[string]string)
		for _, item := range strings.Split(rawHeaders, ";") {
			item = strings.TrimSpace(item)
			if item == "" {
				// ignore empty items and trailing semicolons
				continue
			}
			k, v, ok := strings.Cut(item, ":")
			if !ok {
				return nil, fmt.Errorf("missing colon in additional headers: %q", rawHeaders)
			}
			config.AdditionalHeaders[k] = v
		}
	}

	return config, nil
}

type ionosCloudManagerImpl struct {
	cache  *IonosCache
	client *AutoscalingClient
}

// CreateIonosCloudManager initializes a new IonosCloudManager.
func CreateIonosCloudManager(nodeGroupsConfig []string, userAgent string) (IonosCloudManager, error) {
	klog.V(4).Info("Creating IonosCloud manager")
	config, err := LoadConfigFromEnv()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	manager := newManager(NewAutoscalingClient(config, userAgent))

	if err := manager.initExplicitNodeGroups(nodeGroupsConfig); err != nil {
		return nil, fmt.Errorf("failed to load pre-configured node groups: %w", err)
	}

	return manager, nil
}

func newManager(client *AutoscalingClient) *ionosCloudManagerImpl {
	return &ionosCloudManagerImpl{
		cache:  NewIonosCache(),
		client: client,
	}
}

// initExplicitNodeGroups adds a list of pre-configured node groups to the cache.
// The node groups are parsed from a list of strings in the format of <min>:<max>:<id>.
func (manager *ionosCloudManagerImpl) initExplicitNodeGroups(nodeGroupsConfig []string) error {
	if len(nodeGroupsConfig) == 0 {
		return errors.New("missing value for --nodes flag")
	}

	for _, config := range nodeGroupsConfig {
		parts := strings.Split(config, ":")
		if len(parts) != 3 {
			return fmt.Errorf("invalid autoscaling group config: %s", config)
		}
		min, err := strconv.ParseUint(parts[0], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid value for min: %s", parts[0])
		}
		max, err := strconv.ParseUint(parts[1], 10, 32)
		if err != nil || max == 0 {
			return fmt.Errorf("invalid value for max: %s", parts[1])
		}
		if _, err := uuid.Parse(parts[2]); err != nil {
			return fmt.Errorf("invalid value for id: %s", parts[2])
		}

		np := &nodePool{
			id:      parts[2],
			min:     int(min),
			max:     int(max),
			manager: manager,
		}
		fetchedNodePool, err := manager.client.GetNodePool(np.Id())
		if err != nil {
			return fmt.Errorf("failed to fetch configured node pool %s: %w", np.Id(), err)
		}
		manager.cache.AddNodeGroup(np)
		manager.cache.SetNodeGroupTargetSize(np.Id(), int(*fetchedNodePool.Properties.NodeCount))

		if err := manager.refreshInstancesForNodeGroup(np.Id()); err != nil {
			if !errors.Is(err, errMissingNodeID) {
				return err
			}
			klog.V(4).Infof("node pool %s has nodes with missing node ID", np.Id())
		}
	}
	return nil
}

// GetNodeGroupSize gets node group size.
func (manager *ionosCloudManagerImpl) GetNodeGroupSize(nodeGroup cloudprovider.NodeGroup) (int, error) {
	size, found := manager.cache.GetNodeGroupSize(nodeGroup.Id())
	if !found {
		nodes, err := manager.client.ListNodes(nodeGroup.Id())
		if err != nil {
			return 0, err
		}
		size = len(nodes)
		manager.cache.SetNodeGroupSize(nodeGroup.Id(), size)
	}
	return size, nil
}

// GetNodeGroupSize gets node group target size.
func (manager *ionosCloudManagerImpl) GetNodeGroupTargetSize(nodeGroup cloudprovider.NodeGroup) (int, error) {
	size, found := manager.cache.GetNodeGroupTargetSize(nodeGroup.Id())
	if !found {
		fetchedNodePool, err := manager.client.GetNodePool(nodeGroup.Id())
		if err != nil {
			return 0, err
		}
		size = int(*fetchedNodePool.Properties.NodeCount)
		manager.cache.SetNodeGroupTargetSize(nodeGroup.Id(), size)
	}
	return size, nil
}

// SetNodeGroupSize sets the node group size.
func (manager *ionosCloudManagerImpl) SetNodeGroupSize(nodeGroup cloudprovider.NodeGroup, size int) error {
	klog.V(1).Infof("Setting node group size of %s to %d", nodeGroup.Id(), size)
	if err := manager.client.ResizeNodePool(nodeGroup.Id(), int32(size)); err != nil {
		return fmt.Errorf("node group resize failed: %w", err)
	}
	manager.cache.InvalidateNodeGroupTargetSize(nodeGroup.Id())
	if err := manager.client.WaitForNodePoolResize(nodeGroup.Id(), size); err != nil {
		return fmt.Errorf("wait for node group resize failed: %w", err)
	}
	if err := manager.refreshInstancesForNodeGroup(nodeGroup.Id()); err != nil {
		return fmt.Errorf("cache refresh after resize failed: %w", err)
	}
	klog.V(1).Infof("Successfully increased node group size")
	return nil
}

// DeleteNode deletes a single node.
func (manager *ionosCloudManagerImpl) DeleteNode(nodeGroup cloudprovider.NodeGroup, nodeID string) error {
	klog.V(1).Infof("Deleting node %s from node group %s", nodeID, nodeGroup.Id())
	size, err := manager.GetNodeGroupSize(nodeGroup)
	if err != nil {
		return err
	}
	manager.cache.InvalidateNodeGroupTargetSize(nodeGroup.Id())
	if err := manager.client.DeleteNode(nodeGroup.Id(), nodeID); err != nil {
		return fmt.Errorf("delete node %s failed: %w", nodeID, err)
	}
	targetSize := size - 1
	if err := manager.client.WaitForNodePoolResize(nodeGroup.Id(), targetSize); err != nil {
		return err
	}
	manager.cache.SetNodeGroupSize(nodeGroup.Id(), targetSize)
	manager.cache.RemoveInstanceFromCache(nodeID)
	klog.V(1).Infof("Successfully deleted node %s from node group %s", nodeID, nodeGroup.Id())
	return nil
}

// GetInstancesForNodeGroup returns the instances for the given node group.
func (manager *ionosCloudManagerImpl) GetInstancesForNodeGroup(nodeGroup cloudprovider.NodeGroup) ([]cloudprovider.Instance, error) {
	return manager.fetchInstancesForNodeGroup(nodeGroup.Id())
}

func (manager *ionosCloudManagerImpl) GetNodeGroupForNode(node *apiv1.Node) cloudprovider.NodeGroup {
	nodeID := convertToNodeID(node.Spec.ProviderID)
	return manager.cache.GetNodeGroupForNode(nodeID)
}

func (manager *ionosCloudManagerImpl) refreshInstancesForNodeGroup(id string) error {
	instances, err := manager.fetchInstancesForNodeGroup(id)
	if err != nil {
		return err
	}
	manager.cache.SetInstancesCacheForNodeGroup(id, instances)
	manager.cache.SetNodeGroupSize(id, len(instances))
	return nil
}

func (manager *ionosCloudManagerImpl) fetchInstancesForNodeGroup(id string) ([]cloudprovider.Instance, error) {
	klog.V(4).Infof("Refreshing instances for node group: %s", id)
	kubernetesNodes, err := manager.client.ListNodes(id)
	if err != nil {
		return nil, fmt.Errorf("failed to list nodes for node group %s: %w", id, err)
	}

	instances, err := convertToInstances(kubernetesNodes)
	if err != nil {
		return nil, fmt.Errorf("failed to convert instances for node group %s: %w", id, err)
	}

	return instances, nil
}

func (manager *ionosCloudManagerImpl) GetNodeGroups() []cloudprovider.NodeGroup {
	return manager.cache.GetNodeGroups()
}

// TryLockNodeGroup tries to acquire a lock for a node group.
func (manager *ionosCloudManagerImpl) TryLockNodeGroup(nodeGroup cloudprovider.NodeGroup) bool {
	return manager.cache.TryLockNodeGroup(nodeGroup)
}

// UnlockNodeGroup releases a node group lock.
func (manager *ionosCloudManagerImpl) UnlockNodeGroup(nodeGroup cloudprovider.NodeGroup) {
	manager.cache.UnlockNodeGroup(nodeGroup)
}
