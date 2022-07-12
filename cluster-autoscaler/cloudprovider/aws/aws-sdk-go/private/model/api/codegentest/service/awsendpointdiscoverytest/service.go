// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package awsendpointdiscoverytest

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/client"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/client/metadata"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/crr"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/signer/v4"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// AwsEndpointDiscoveryTest provides the API operation methods for making requests to
// AwsEndpointDiscoveryTest. See this package's package overview docs
// for details on the service.
//
// AwsEndpointDiscoveryTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type AwsEndpointDiscoveryTest struct {
	*client.Client
	endpointCache *crr.EndpointCache
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "AwsEndpointDiscoveryTest"        // Name of service.
	EndpointsID = "awsendpointdiscoverytestservice" // ID to lookup a service endpoint with.
	ServiceID   = "AwsEndpointDiscoveryTest"        // ServiceID is a unique identifier of a specific service.
)

// New creates a new instance of the AwsEndpointDiscoveryTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     mySession := session.Must(session.NewSession())
//
//     // Create a AwsEndpointDiscoveryTest client from just a session.
//     svc := awsendpointdiscoverytest.New(mySession)
//
//     // Create a AwsEndpointDiscoveryTest client with additional configuration
//     svc := awsendpointdiscoverytest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *AwsEndpointDiscoveryTest {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "awsendpointdiscoverytestservice"
	}
	return newClient(*c.Config, c.Handlers, c.PartitionID, c.Endpoint, c.SigningRegion, c.SigningName, c.ResolvedRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, partitionID, endpoint, signingRegion, signingName, resolvedRegion string) *AwsEndpointDiscoveryTest {
	svc := &AwsEndpointDiscoveryTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:    ServiceName,
				ServiceID:      ServiceID,
				SigningName:    signingName,
				SigningRegion:  signingRegion,
				PartitionID:    partitionID,
				Endpoint:       endpoint,
				APIVersion:     "2018-08-31",
				ResolvedRegion: resolvedRegion,
				JSONVersion:    "1.1",
				TargetPrefix:   "AwsEndpointDiscoveryTestService",
			},
			handlers,
		),
	}
	svc.endpointCache = crr.NewEndpointCache(10)

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a AwsEndpointDiscoveryTest operation and runs any
// custom request initialization.
func (c *AwsEndpointDiscoveryTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
