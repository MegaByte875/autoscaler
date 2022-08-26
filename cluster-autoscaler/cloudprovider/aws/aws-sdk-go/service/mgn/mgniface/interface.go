// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mgniface provides an interface to enable mocking the Application Migration Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mgniface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/mgn"
)

// MgnAPI provides an interface to enable mocking the
// mgn.Mgn service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Application Migration Service.
//	func myFunc(svc mgniface.MgnAPI) bool {
//	    // Make svc.ChangeServerLifeCycleState request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := mgn.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockMgnClient struct {
//	    mgniface.MgnAPI
//	}
//	func (m *mockMgnClient) ChangeServerLifeCycleState(input *mgn.ChangeServerLifeCycleStateInput) (*mgn.ChangeServerLifeCycleStateOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockMgnClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type MgnAPI interface {
	ChangeServerLifeCycleState(*mgn.ChangeServerLifeCycleStateInput) (*mgn.ChangeServerLifeCycleStateOutput, error)
	ChangeServerLifeCycleStateWithContext(aws.Context, *mgn.ChangeServerLifeCycleStateInput, ...request.Option) (*mgn.ChangeServerLifeCycleStateOutput, error)
	ChangeServerLifeCycleStateRequest(*mgn.ChangeServerLifeCycleStateInput) (*request.Request, *mgn.ChangeServerLifeCycleStateOutput)

	CreateReplicationConfigurationTemplate(*mgn.CreateReplicationConfigurationTemplateInput) (*mgn.CreateReplicationConfigurationTemplateOutput, error)
	CreateReplicationConfigurationTemplateWithContext(aws.Context, *mgn.CreateReplicationConfigurationTemplateInput, ...request.Option) (*mgn.CreateReplicationConfigurationTemplateOutput, error)
	CreateReplicationConfigurationTemplateRequest(*mgn.CreateReplicationConfigurationTemplateInput) (*request.Request, *mgn.CreateReplicationConfigurationTemplateOutput)

	DeleteJob(*mgn.DeleteJobInput) (*mgn.DeleteJobOutput, error)
	DeleteJobWithContext(aws.Context, *mgn.DeleteJobInput, ...request.Option) (*mgn.DeleteJobOutput, error)
	DeleteJobRequest(*mgn.DeleteJobInput) (*request.Request, *mgn.DeleteJobOutput)

	DeleteReplicationConfigurationTemplate(*mgn.DeleteReplicationConfigurationTemplateInput) (*mgn.DeleteReplicationConfigurationTemplateOutput, error)
	DeleteReplicationConfigurationTemplateWithContext(aws.Context, *mgn.DeleteReplicationConfigurationTemplateInput, ...request.Option) (*mgn.DeleteReplicationConfigurationTemplateOutput, error)
	DeleteReplicationConfigurationTemplateRequest(*mgn.DeleteReplicationConfigurationTemplateInput) (*request.Request, *mgn.DeleteReplicationConfigurationTemplateOutput)

	DeleteSourceServer(*mgn.DeleteSourceServerInput) (*mgn.DeleteSourceServerOutput, error)
	DeleteSourceServerWithContext(aws.Context, *mgn.DeleteSourceServerInput, ...request.Option) (*mgn.DeleteSourceServerOutput, error)
	DeleteSourceServerRequest(*mgn.DeleteSourceServerInput) (*request.Request, *mgn.DeleteSourceServerOutput)

	DeleteVcenterClient(*mgn.DeleteVcenterClientInput) (*mgn.DeleteVcenterClientOutput, error)
	DeleteVcenterClientWithContext(aws.Context, *mgn.DeleteVcenterClientInput, ...request.Option) (*mgn.DeleteVcenterClientOutput, error)
	DeleteVcenterClientRequest(*mgn.DeleteVcenterClientInput) (*request.Request, *mgn.DeleteVcenterClientOutput)

	DescribeJobLogItems(*mgn.DescribeJobLogItemsInput) (*mgn.DescribeJobLogItemsOutput, error)
	DescribeJobLogItemsWithContext(aws.Context, *mgn.DescribeJobLogItemsInput, ...request.Option) (*mgn.DescribeJobLogItemsOutput, error)
	DescribeJobLogItemsRequest(*mgn.DescribeJobLogItemsInput) (*request.Request, *mgn.DescribeJobLogItemsOutput)

	DescribeJobLogItemsPages(*mgn.DescribeJobLogItemsInput, func(*mgn.DescribeJobLogItemsOutput, bool) bool) error
	DescribeJobLogItemsPagesWithContext(aws.Context, *mgn.DescribeJobLogItemsInput, func(*mgn.DescribeJobLogItemsOutput, bool) bool, ...request.Option) error

	DescribeJobs(*mgn.DescribeJobsInput) (*mgn.DescribeJobsOutput, error)
	DescribeJobsWithContext(aws.Context, *mgn.DescribeJobsInput, ...request.Option) (*mgn.DescribeJobsOutput, error)
	DescribeJobsRequest(*mgn.DescribeJobsInput) (*request.Request, *mgn.DescribeJobsOutput)

	DescribeJobsPages(*mgn.DescribeJobsInput, func(*mgn.DescribeJobsOutput, bool) bool) error
	DescribeJobsPagesWithContext(aws.Context, *mgn.DescribeJobsInput, func(*mgn.DescribeJobsOutput, bool) bool, ...request.Option) error

	DescribeReplicationConfigurationTemplates(*mgn.DescribeReplicationConfigurationTemplatesInput) (*mgn.DescribeReplicationConfigurationTemplatesOutput, error)
	DescribeReplicationConfigurationTemplatesWithContext(aws.Context, *mgn.DescribeReplicationConfigurationTemplatesInput, ...request.Option) (*mgn.DescribeReplicationConfigurationTemplatesOutput, error)
	DescribeReplicationConfigurationTemplatesRequest(*mgn.DescribeReplicationConfigurationTemplatesInput) (*request.Request, *mgn.DescribeReplicationConfigurationTemplatesOutput)

	DescribeReplicationConfigurationTemplatesPages(*mgn.DescribeReplicationConfigurationTemplatesInput, func(*mgn.DescribeReplicationConfigurationTemplatesOutput, bool) bool) error
	DescribeReplicationConfigurationTemplatesPagesWithContext(aws.Context, *mgn.DescribeReplicationConfigurationTemplatesInput, func(*mgn.DescribeReplicationConfigurationTemplatesOutput, bool) bool, ...request.Option) error

	DescribeSourceServers(*mgn.DescribeSourceServersInput) (*mgn.DescribeSourceServersOutput, error)
	DescribeSourceServersWithContext(aws.Context, *mgn.DescribeSourceServersInput, ...request.Option) (*mgn.DescribeSourceServersOutput, error)
	DescribeSourceServersRequest(*mgn.DescribeSourceServersInput) (*request.Request, *mgn.DescribeSourceServersOutput)

	DescribeSourceServersPages(*mgn.DescribeSourceServersInput, func(*mgn.DescribeSourceServersOutput, bool) bool) error
	DescribeSourceServersPagesWithContext(aws.Context, *mgn.DescribeSourceServersInput, func(*mgn.DescribeSourceServersOutput, bool) bool, ...request.Option) error

	DescribeVcenterClients(*mgn.DescribeVcenterClientsInput) (*mgn.DescribeVcenterClientsOutput, error)
	DescribeVcenterClientsWithContext(aws.Context, *mgn.DescribeVcenterClientsInput, ...request.Option) (*mgn.DescribeVcenterClientsOutput, error)
	DescribeVcenterClientsRequest(*mgn.DescribeVcenterClientsInput) (*request.Request, *mgn.DescribeVcenterClientsOutput)

	DescribeVcenterClientsPages(*mgn.DescribeVcenterClientsInput, func(*mgn.DescribeVcenterClientsOutput, bool) bool) error
	DescribeVcenterClientsPagesWithContext(aws.Context, *mgn.DescribeVcenterClientsInput, func(*mgn.DescribeVcenterClientsOutput, bool) bool, ...request.Option) error

	DisconnectFromService(*mgn.DisconnectFromServiceInput) (*mgn.DisconnectFromServiceOutput, error)
	DisconnectFromServiceWithContext(aws.Context, *mgn.DisconnectFromServiceInput, ...request.Option) (*mgn.DisconnectFromServiceOutput, error)
	DisconnectFromServiceRequest(*mgn.DisconnectFromServiceInput) (*request.Request, *mgn.DisconnectFromServiceOutput)

	FinalizeCutover(*mgn.FinalizeCutoverInput) (*mgn.FinalizeCutoverOutput, error)
	FinalizeCutoverWithContext(aws.Context, *mgn.FinalizeCutoverInput, ...request.Option) (*mgn.FinalizeCutoverOutput, error)
	FinalizeCutoverRequest(*mgn.FinalizeCutoverInput) (*request.Request, *mgn.FinalizeCutoverOutput)

	GetLaunchConfiguration(*mgn.GetLaunchConfigurationInput) (*mgn.GetLaunchConfigurationOutput, error)
	GetLaunchConfigurationWithContext(aws.Context, *mgn.GetLaunchConfigurationInput, ...request.Option) (*mgn.GetLaunchConfigurationOutput, error)
	GetLaunchConfigurationRequest(*mgn.GetLaunchConfigurationInput) (*request.Request, *mgn.GetLaunchConfigurationOutput)

	GetReplicationConfiguration(*mgn.GetReplicationConfigurationInput) (*mgn.GetReplicationConfigurationOutput, error)
	GetReplicationConfigurationWithContext(aws.Context, *mgn.GetReplicationConfigurationInput, ...request.Option) (*mgn.GetReplicationConfigurationOutput, error)
	GetReplicationConfigurationRequest(*mgn.GetReplicationConfigurationInput) (*request.Request, *mgn.GetReplicationConfigurationOutput)

	InitializeService(*mgn.InitializeServiceInput) (*mgn.InitializeServiceOutput, error)
	InitializeServiceWithContext(aws.Context, *mgn.InitializeServiceInput, ...request.Option) (*mgn.InitializeServiceOutput, error)
	InitializeServiceRequest(*mgn.InitializeServiceInput) (*request.Request, *mgn.InitializeServiceOutput)

	ListTagsForResource(*mgn.ListTagsForResourceInput) (*mgn.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *mgn.ListTagsForResourceInput, ...request.Option) (*mgn.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*mgn.ListTagsForResourceInput) (*request.Request, *mgn.ListTagsForResourceOutput)

	MarkAsArchived(*mgn.MarkAsArchivedInput) (*mgn.MarkAsArchivedOutput, error)
	MarkAsArchivedWithContext(aws.Context, *mgn.MarkAsArchivedInput, ...request.Option) (*mgn.MarkAsArchivedOutput, error)
	MarkAsArchivedRequest(*mgn.MarkAsArchivedInput) (*request.Request, *mgn.MarkAsArchivedOutput)

	RetryDataReplication(*mgn.RetryDataReplicationInput) (*mgn.RetryDataReplicationOutput, error)
	RetryDataReplicationWithContext(aws.Context, *mgn.RetryDataReplicationInput, ...request.Option) (*mgn.RetryDataReplicationOutput, error)
	RetryDataReplicationRequest(*mgn.RetryDataReplicationInput) (*request.Request, *mgn.RetryDataReplicationOutput)

	StartCutover(*mgn.StartCutoverInput) (*mgn.StartCutoverOutput, error)
	StartCutoverWithContext(aws.Context, *mgn.StartCutoverInput, ...request.Option) (*mgn.StartCutoverOutput, error)
	StartCutoverRequest(*mgn.StartCutoverInput) (*request.Request, *mgn.StartCutoverOutput)

	StartReplication(*mgn.StartReplicationInput) (*mgn.StartReplicationOutput, error)
	StartReplicationWithContext(aws.Context, *mgn.StartReplicationInput, ...request.Option) (*mgn.StartReplicationOutput, error)
	StartReplicationRequest(*mgn.StartReplicationInput) (*request.Request, *mgn.StartReplicationOutput)

	StartTest(*mgn.StartTestInput) (*mgn.StartTestOutput, error)
	StartTestWithContext(aws.Context, *mgn.StartTestInput, ...request.Option) (*mgn.StartTestOutput, error)
	StartTestRequest(*mgn.StartTestInput) (*request.Request, *mgn.StartTestOutput)

	TagResource(*mgn.TagResourceInput) (*mgn.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *mgn.TagResourceInput, ...request.Option) (*mgn.TagResourceOutput, error)
	TagResourceRequest(*mgn.TagResourceInput) (*request.Request, *mgn.TagResourceOutput)

	TerminateTargetInstances(*mgn.TerminateTargetInstancesInput) (*mgn.TerminateTargetInstancesOutput, error)
	TerminateTargetInstancesWithContext(aws.Context, *mgn.TerminateTargetInstancesInput, ...request.Option) (*mgn.TerminateTargetInstancesOutput, error)
	TerminateTargetInstancesRequest(*mgn.TerminateTargetInstancesInput) (*request.Request, *mgn.TerminateTargetInstancesOutput)

	UntagResource(*mgn.UntagResourceInput) (*mgn.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *mgn.UntagResourceInput, ...request.Option) (*mgn.UntagResourceOutput, error)
	UntagResourceRequest(*mgn.UntagResourceInput) (*request.Request, *mgn.UntagResourceOutput)

	UpdateLaunchConfiguration(*mgn.UpdateLaunchConfigurationInput) (*mgn.UpdateLaunchConfigurationOutput, error)
	UpdateLaunchConfigurationWithContext(aws.Context, *mgn.UpdateLaunchConfigurationInput, ...request.Option) (*mgn.UpdateLaunchConfigurationOutput, error)
	UpdateLaunchConfigurationRequest(*mgn.UpdateLaunchConfigurationInput) (*request.Request, *mgn.UpdateLaunchConfigurationOutput)

	UpdateReplicationConfiguration(*mgn.UpdateReplicationConfigurationInput) (*mgn.UpdateReplicationConfigurationOutput, error)
	UpdateReplicationConfigurationWithContext(aws.Context, *mgn.UpdateReplicationConfigurationInput, ...request.Option) (*mgn.UpdateReplicationConfigurationOutput, error)
	UpdateReplicationConfigurationRequest(*mgn.UpdateReplicationConfigurationInput) (*request.Request, *mgn.UpdateReplicationConfigurationOutput)

	UpdateReplicationConfigurationTemplate(*mgn.UpdateReplicationConfigurationTemplateInput) (*mgn.UpdateReplicationConfigurationTemplateOutput, error)
	UpdateReplicationConfigurationTemplateWithContext(aws.Context, *mgn.UpdateReplicationConfigurationTemplateInput, ...request.Option) (*mgn.UpdateReplicationConfigurationTemplateOutput, error)
	UpdateReplicationConfigurationTemplateRequest(*mgn.UpdateReplicationConfigurationTemplateInput) (*request.Request, *mgn.UpdateReplicationConfigurationTemplateOutput)

	UpdateSourceServerReplicationType(*mgn.UpdateSourceServerReplicationTypeInput) (*mgn.UpdateSourceServerReplicationTypeOutput, error)
	UpdateSourceServerReplicationTypeWithContext(aws.Context, *mgn.UpdateSourceServerReplicationTypeInput, ...request.Option) (*mgn.UpdateSourceServerReplicationTypeOutput, error)
	UpdateSourceServerReplicationTypeRequest(*mgn.UpdateSourceServerReplicationTypeInput) (*request.Request, *mgn.UpdateSourceServerReplicationTypeOutput)
}

var _ MgnAPI = (*mgn.Mgn)(nil)
