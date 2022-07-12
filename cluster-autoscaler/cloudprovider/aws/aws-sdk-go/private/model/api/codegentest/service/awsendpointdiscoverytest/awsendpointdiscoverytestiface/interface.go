// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package awsendpointdiscoverytestiface provides an interface to enable mocking the AwsEndpointDiscoveryTest service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package awsendpointdiscoverytestiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/model/api/codegentest/service/awsendpointdiscoverytest"
)

// AwsEndpointDiscoveryTestAPI provides an interface to enable mocking the
// awsendpointdiscoverytest.AwsEndpointDiscoveryTest service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AwsEndpointDiscoveryTest.
//    func myFunc(svc awsendpointdiscoverytestiface.AwsEndpointDiscoveryTestAPI) bool {
//        // Make svc.DescribeEndpoints request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := awsendpointdiscoverytest.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAwsEndpointDiscoveryTestClient struct {
//        awsendpointdiscoverytestiface.AwsEndpointDiscoveryTestAPI
//    }
//    func (m *mockAwsEndpointDiscoveryTestClient) DescribeEndpoints(input *awsendpointdiscoverytest.DescribeEndpointsInput) (*awsendpointdiscoverytest.DescribeEndpointsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAwsEndpointDiscoveryTestClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type AwsEndpointDiscoveryTestAPI interface {
	DescribeEndpoints(*awsendpointdiscoverytest.DescribeEndpointsInput) (*awsendpointdiscoverytest.DescribeEndpointsOutput, error)
	DescribeEndpointsWithContext(aws.Context, *awsendpointdiscoverytest.DescribeEndpointsInput, ...request.Option) (*awsendpointdiscoverytest.DescribeEndpointsOutput, error)
	DescribeEndpointsRequest(*awsendpointdiscoverytest.DescribeEndpointsInput) (*request.Request, *awsendpointdiscoverytest.DescribeEndpointsOutput)

	TestDiscoveryIdentifiersRequired(*awsendpointdiscoverytest.TestDiscoveryIdentifiersRequiredInput) (*awsendpointdiscoverytest.TestDiscoveryIdentifiersRequiredOutput, error)
	TestDiscoveryIdentifiersRequiredWithContext(aws.Context, *awsendpointdiscoverytest.TestDiscoveryIdentifiersRequiredInput, ...request.Option) (*awsendpointdiscoverytest.TestDiscoveryIdentifiersRequiredOutput, error)
	TestDiscoveryIdentifiersRequiredRequest(*awsendpointdiscoverytest.TestDiscoveryIdentifiersRequiredInput) (*request.Request, *awsendpointdiscoverytest.TestDiscoveryIdentifiersRequiredOutput)

	TestDiscoveryOptional(*awsendpointdiscoverytest.TestDiscoveryOptionalInput) (*awsendpointdiscoverytest.TestDiscoveryOptionalOutput, error)
	TestDiscoveryOptionalWithContext(aws.Context, *awsendpointdiscoverytest.TestDiscoveryOptionalInput, ...request.Option) (*awsendpointdiscoverytest.TestDiscoveryOptionalOutput, error)
	TestDiscoveryOptionalRequest(*awsendpointdiscoverytest.TestDiscoveryOptionalInput) (*request.Request, *awsendpointdiscoverytest.TestDiscoveryOptionalOutput)

	TestDiscoveryRequired(*awsendpointdiscoverytest.TestDiscoveryRequiredInput) (*awsendpointdiscoverytest.TestDiscoveryRequiredOutput, error)
	TestDiscoveryRequiredWithContext(aws.Context, *awsendpointdiscoverytest.TestDiscoveryRequiredInput, ...request.Option) (*awsendpointdiscoverytest.TestDiscoveryRequiredOutput, error)
	TestDiscoveryRequiredRequest(*awsendpointdiscoverytest.TestDiscoveryRequiredInput) (*request.Request, *awsendpointdiscoverytest.TestDiscoveryRequiredOutput)
}

var _ AwsEndpointDiscoveryTestAPI = (*awsendpointdiscoverytest.AwsEndpointDiscoveryTest)(nil)
