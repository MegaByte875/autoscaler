// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package clouddirectoryiface provides an interface to enable mocking the Amazon CloudDirectory service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package clouddirectoryiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/clouddirectory"
)

// CloudDirectoryAPI provides an interface to enable mocking the
// clouddirectory.CloudDirectory service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CloudDirectory.
//    func myFunc(svc clouddirectoryiface.CloudDirectoryAPI) bool {
//        // Make svc.AddFacetToObject request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := clouddirectory.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudDirectoryClient struct {
//        clouddirectoryiface.CloudDirectoryAPI
//    }
//    func (m *mockCloudDirectoryClient) AddFacetToObject(input *clouddirectory.AddFacetToObjectInput) (*clouddirectory.AddFacetToObjectOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudDirectoryClient{}
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
type CloudDirectoryAPI interface {
	AddFacetToObject(*clouddirectory.AddFacetToObjectInput) (*clouddirectory.AddFacetToObjectOutput, error)
	AddFacetToObjectWithContext(aws.Context, *clouddirectory.AddFacetToObjectInput, ...request.Option) (*clouddirectory.AddFacetToObjectOutput, error)
	AddFacetToObjectRequest(*clouddirectory.AddFacetToObjectInput) (*request.Request, *clouddirectory.AddFacetToObjectOutput)

	ApplySchema(*clouddirectory.ApplySchemaInput) (*clouddirectory.ApplySchemaOutput, error)
	ApplySchemaWithContext(aws.Context, *clouddirectory.ApplySchemaInput, ...request.Option) (*clouddirectory.ApplySchemaOutput, error)
	ApplySchemaRequest(*clouddirectory.ApplySchemaInput) (*request.Request, *clouddirectory.ApplySchemaOutput)

	AttachObject(*clouddirectory.AttachObjectInput) (*clouddirectory.AttachObjectOutput, error)
	AttachObjectWithContext(aws.Context, *clouddirectory.AttachObjectInput, ...request.Option) (*clouddirectory.AttachObjectOutput, error)
	AttachObjectRequest(*clouddirectory.AttachObjectInput) (*request.Request, *clouddirectory.AttachObjectOutput)

	AttachPolicy(*clouddirectory.AttachPolicyInput) (*clouddirectory.AttachPolicyOutput, error)
	AttachPolicyWithContext(aws.Context, *clouddirectory.AttachPolicyInput, ...request.Option) (*clouddirectory.AttachPolicyOutput, error)
	AttachPolicyRequest(*clouddirectory.AttachPolicyInput) (*request.Request, *clouddirectory.AttachPolicyOutput)

	AttachToIndex(*clouddirectory.AttachToIndexInput) (*clouddirectory.AttachToIndexOutput, error)
	AttachToIndexWithContext(aws.Context, *clouddirectory.AttachToIndexInput, ...request.Option) (*clouddirectory.AttachToIndexOutput, error)
	AttachToIndexRequest(*clouddirectory.AttachToIndexInput) (*request.Request, *clouddirectory.AttachToIndexOutput)

	AttachTypedLink(*clouddirectory.AttachTypedLinkInput) (*clouddirectory.AttachTypedLinkOutput, error)
	AttachTypedLinkWithContext(aws.Context, *clouddirectory.AttachTypedLinkInput, ...request.Option) (*clouddirectory.AttachTypedLinkOutput, error)
	AttachTypedLinkRequest(*clouddirectory.AttachTypedLinkInput) (*request.Request, *clouddirectory.AttachTypedLinkOutput)

	BatchRead(*clouddirectory.BatchReadInput) (*clouddirectory.BatchReadOutput, error)
	BatchReadWithContext(aws.Context, *clouddirectory.BatchReadInput, ...request.Option) (*clouddirectory.BatchReadOutput, error)
	BatchReadRequest(*clouddirectory.BatchReadInput) (*request.Request, *clouddirectory.BatchReadOutput)

	BatchWrite(*clouddirectory.BatchWriteInput) (*clouddirectory.BatchWriteOutput, error)
	BatchWriteWithContext(aws.Context, *clouddirectory.BatchWriteInput, ...request.Option) (*clouddirectory.BatchWriteOutput, error)
	BatchWriteRequest(*clouddirectory.BatchWriteInput) (*request.Request, *clouddirectory.BatchWriteOutput)

	CreateDirectory(*clouddirectory.CreateDirectoryInput) (*clouddirectory.CreateDirectoryOutput, error)
	CreateDirectoryWithContext(aws.Context, *clouddirectory.CreateDirectoryInput, ...request.Option) (*clouddirectory.CreateDirectoryOutput, error)
	CreateDirectoryRequest(*clouddirectory.CreateDirectoryInput) (*request.Request, *clouddirectory.CreateDirectoryOutput)

	CreateFacet(*clouddirectory.CreateFacetInput) (*clouddirectory.CreateFacetOutput, error)
	CreateFacetWithContext(aws.Context, *clouddirectory.CreateFacetInput, ...request.Option) (*clouddirectory.CreateFacetOutput, error)
	CreateFacetRequest(*clouddirectory.CreateFacetInput) (*request.Request, *clouddirectory.CreateFacetOutput)

	CreateIndex(*clouddirectory.CreateIndexInput) (*clouddirectory.CreateIndexOutput, error)
	CreateIndexWithContext(aws.Context, *clouddirectory.CreateIndexInput, ...request.Option) (*clouddirectory.CreateIndexOutput, error)
	CreateIndexRequest(*clouddirectory.CreateIndexInput) (*request.Request, *clouddirectory.CreateIndexOutput)

	CreateObject(*clouddirectory.CreateObjectInput) (*clouddirectory.CreateObjectOutput, error)
	CreateObjectWithContext(aws.Context, *clouddirectory.CreateObjectInput, ...request.Option) (*clouddirectory.CreateObjectOutput, error)
	CreateObjectRequest(*clouddirectory.CreateObjectInput) (*request.Request, *clouddirectory.CreateObjectOutput)

	CreateSchema(*clouddirectory.CreateSchemaInput) (*clouddirectory.CreateSchemaOutput, error)
	CreateSchemaWithContext(aws.Context, *clouddirectory.CreateSchemaInput, ...request.Option) (*clouddirectory.CreateSchemaOutput, error)
	CreateSchemaRequest(*clouddirectory.CreateSchemaInput) (*request.Request, *clouddirectory.CreateSchemaOutput)

	CreateTypedLinkFacet(*clouddirectory.CreateTypedLinkFacetInput) (*clouddirectory.CreateTypedLinkFacetOutput, error)
	CreateTypedLinkFacetWithContext(aws.Context, *clouddirectory.CreateTypedLinkFacetInput, ...request.Option) (*clouddirectory.CreateTypedLinkFacetOutput, error)
	CreateTypedLinkFacetRequest(*clouddirectory.CreateTypedLinkFacetInput) (*request.Request, *clouddirectory.CreateTypedLinkFacetOutput)

	DeleteDirectory(*clouddirectory.DeleteDirectoryInput) (*clouddirectory.DeleteDirectoryOutput, error)
	DeleteDirectoryWithContext(aws.Context, *clouddirectory.DeleteDirectoryInput, ...request.Option) (*clouddirectory.DeleteDirectoryOutput, error)
	DeleteDirectoryRequest(*clouddirectory.DeleteDirectoryInput) (*request.Request, *clouddirectory.DeleteDirectoryOutput)

	DeleteFacet(*clouddirectory.DeleteFacetInput) (*clouddirectory.DeleteFacetOutput, error)
	DeleteFacetWithContext(aws.Context, *clouddirectory.DeleteFacetInput, ...request.Option) (*clouddirectory.DeleteFacetOutput, error)
	DeleteFacetRequest(*clouddirectory.DeleteFacetInput) (*request.Request, *clouddirectory.DeleteFacetOutput)

	DeleteObject(*clouddirectory.DeleteObjectInput) (*clouddirectory.DeleteObjectOutput, error)
	DeleteObjectWithContext(aws.Context, *clouddirectory.DeleteObjectInput, ...request.Option) (*clouddirectory.DeleteObjectOutput, error)
	DeleteObjectRequest(*clouddirectory.DeleteObjectInput) (*request.Request, *clouddirectory.DeleteObjectOutput)

	DeleteSchema(*clouddirectory.DeleteSchemaInput) (*clouddirectory.DeleteSchemaOutput, error)
	DeleteSchemaWithContext(aws.Context, *clouddirectory.DeleteSchemaInput, ...request.Option) (*clouddirectory.DeleteSchemaOutput, error)
	DeleteSchemaRequest(*clouddirectory.DeleteSchemaInput) (*request.Request, *clouddirectory.DeleteSchemaOutput)

	DeleteTypedLinkFacet(*clouddirectory.DeleteTypedLinkFacetInput) (*clouddirectory.DeleteTypedLinkFacetOutput, error)
	DeleteTypedLinkFacetWithContext(aws.Context, *clouddirectory.DeleteTypedLinkFacetInput, ...request.Option) (*clouddirectory.DeleteTypedLinkFacetOutput, error)
	DeleteTypedLinkFacetRequest(*clouddirectory.DeleteTypedLinkFacetInput) (*request.Request, *clouddirectory.DeleteTypedLinkFacetOutput)

	DetachFromIndex(*clouddirectory.DetachFromIndexInput) (*clouddirectory.DetachFromIndexOutput, error)
	DetachFromIndexWithContext(aws.Context, *clouddirectory.DetachFromIndexInput, ...request.Option) (*clouddirectory.DetachFromIndexOutput, error)
	DetachFromIndexRequest(*clouddirectory.DetachFromIndexInput) (*request.Request, *clouddirectory.DetachFromIndexOutput)

	DetachObject(*clouddirectory.DetachObjectInput) (*clouddirectory.DetachObjectOutput, error)
	DetachObjectWithContext(aws.Context, *clouddirectory.DetachObjectInput, ...request.Option) (*clouddirectory.DetachObjectOutput, error)
	DetachObjectRequest(*clouddirectory.DetachObjectInput) (*request.Request, *clouddirectory.DetachObjectOutput)

	DetachPolicy(*clouddirectory.DetachPolicyInput) (*clouddirectory.DetachPolicyOutput, error)
	DetachPolicyWithContext(aws.Context, *clouddirectory.DetachPolicyInput, ...request.Option) (*clouddirectory.DetachPolicyOutput, error)
	DetachPolicyRequest(*clouddirectory.DetachPolicyInput) (*request.Request, *clouddirectory.DetachPolicyOutput)

	DetachTypedLink(*clouddirectory.DetachTypedLinkInput) (*clouddirectory.DetachTypedLinkOutput, error)
	DetachTypedLinkWithContext(aws.Context, *clouddirectory.DetachTypedLinkInput, ...request.Option) (*clouddirectory.DetachTypedLinkOutput, error)
	DetachTypedLinkRequest(*clouddirectory.DetachTypedLinkInput) (*request.Request, *clouddirectory.DetachTypedLinkOutput)

	DisableDirectory(*clouddirectory.DisableDirectoryInput) (*clouddirectory.DisableDirectoryOutput, error)
	DisableDirectoryWithContext(aws.Context, *clouddirectory.DisableDirectoryInput, ...request.Option) (*clouddirectory.DisableDirectoryOutput, error)
	DisableDirectoryRequest(*clouddirectory.DisableDirectoryInput) (*request.Request, *clouddirectory.DisableDirectoryOutput)

	EnableDirectory(*clouddirectory.EnableDirectoryInput) (*clouddirectory.EnableDirectoryOutput, error)
	EnableDirectoryWithContext(aws.Context, *clouddirectory.EnableDirectoryInput, ...request.Option) (*clouddirectory.EnableDirectoryOutput, error)
	EnableDirectoryRequest(*clouddirectory.EnableDirectoryInput) (*request.Request, *clouddirectory.EnableDirectoryOutput)

	GetAppliedSchemaVersion(*clouddirectory.GetAppliedSchemaVersionInput) (*clouddirectory.GetAppliedSchemaVersionOutput, error)
	GetAppliedSchemaVersionWithContext(aws.Context, *clouddirectory.GetAppliedSchemaVersionInput, ...request.Option) (*clouddirectory.GetAppliedSchemaVersionOutput, error)
	GetAppliedSchemaVersionRequest(*clouddirectory.GetAppliedSchemaVersionInput) (*request.Request, *clouddirectory.GetAppliedSchemaVersionOutput)

	GetDirectory(*clouddirectory.GetDirectoryInput) (*clouddirectory.GetDirectoryOutput, error)
	GetDirectoryWithContext(aws.Context, *clouddirectory.GetDirectoryInput, ...request.Option) (*clouddirectory.GetDirectoryOutput, error)
	GetDirectoryRequest(*clouddirectory.GetDirectoryInput) (*request.Request, *clouddirectory.GetDirectoryOutput)

	GetFacet(*clouddirectory.GetFacetInput) (*clouddirectory.GetFacetOutput, error)
	GetFacetWithContext(aws.Context, *clouddirectory.GetFacetInput, ...request.Option) (*clouddirectory.GetFacetOutput, error)
	GetFacetRequest(*clouddirectory.GetFacetInput) (*request.Request, *clouddirectory.GetFacetOutput)

	GetLinkAttributes(*clouddirectory.GetLinkAttributesInput) (*clouddirectory.GetLinkAttributesOutput, error)
	GetLinkAttributesWithContext(aws.Context, *clouddirectory.GetLinkAttributesInput, ...request.Option) (*clouddirectory.GetLinkAttributesOutput, error)
	GetLinkAttributesRequest(*clouddirectory.GetLinkAttributesInput) (*request.Request, *clouddirectory.GetLinkAttributesOutput)

	GetObjectAttributes(*clouddirectory.GetObjectAttributesInput) (*clouddirectory.GetObjectAttributesOutput, error)
	GetObjectAttributesWithContext(aws.Context, *clouddirectory.GetObjectAttributesInput, ...request.Option) (*clouddirectory.GetObjectAttributesOutput, error)
	GetObjectAttributesRequest(*clouddirectory.GetObjectAttributesInput) (*request.Request, *clouddirectory.GetObjectAttributesOutput)

	GetObjectInformation(*clouddirectory.GetObjectInformationInput) (*clouddirectory.GetObjectInformationOutput, error)
	GetObjectInformationWithContext(aws.Context, *clouddirectory.GetObjectInformationInput, ...request.Option) (*clouddirectory.GetObjectInformationOutput, error)
	GetObjectInformationRequest(*clouddirectory.GetObjectInformationInput) (*request.Request, *clouddirectory.GetObjectInformationOutput)

	GetSchemaAsJson(*clouddirectory.GetSchemaAsJsonInput) (*clouddirectory.GetSchemaAsJsonOutput, error)
	GetSchemaAsJsonWithContext(aws.Context, *clouddirectory.GetSchemaAsJsonInput, ...request.Option) (*clouddirectory.GetSchemaAsJsonOutput, error)
	GetSchemaAsJsonRequest(*clouddirectory.GetSchemaAsJsonInput) (*request.Request, *clouddirectory.GetSchemaAsJsonOutput)

	GetTypedLinkFacetInformation(*clouddirectory.GetTypedLinkFacetInformationInput) (*clouddirectory.GetTypedLinkFacetInformationOutput, error)
	GetTypedLinkFacetInformationWithContext(aws.Context, *clouddirectory.GetTypedLinkFacetInformationInput, ...request.Option) (*clouddirectory.GetTypedLinkFacetInformationOutput, error)
	GetTypedLinkFacetInformationRequest(*clouddirectory.GetTypedLinkFacetInformationInput) (*request.Request, *clouddirectory.GetTypedLinkFacetInformationOutput)

	ListAppliedSchemaArns(*clouddirectory.ListAppliedSchemaArnsInput) (*clouddirectory.ListAppliedSchemaArnsOutput, error)
	ListAppliedSchemaArnsWithContext(aws.Context, *clouddirectory.ListAppliedSchemaArnsInput, ...request.Option) (*clouddirectory.ListAppliedSchemaArnsOutput, error)
	ListAppliedSchemaArnsRequest(*clouddirectory.ListAppliedSchemaArnsInput) (*request.Request, *clouddirectory.ListAppliedSchemaArnsOutput)

	ListAppliedSchemaArnsPages(*clouddirectory.ListAppliedSchemaArnsInput, func(*clouddirectory.ListAppliedSchemaArnsOutput, bool) bool) error
	ListAppliedSchemaArnsPagesWithContext(aws.Context, *clouddirectory.ListAppliedSchemaArnsInput, func(*clouddirectory.ListAppliedSchemaArnsOutput, bool) bool, ...request.Option) error

	ListAttachedIndices(*clouddirectory.ListAttachedIndicesInput) (*clouddirectory.ListAttachedIndicesOutput, error)
	ListAttachedIndicesWithContext(aws.Context, *clouddirectory.ListAttachedIndicesInput, ...request.Option) (*clouddirectory.ListAttachedIndicesOutput, error)
	ListAttachedIndicesRequest(*clouddirectory.ListAttachedIndicesInput) (*request.Request, *clouddirectory.ListAttachedIndicesOutput)

	ListAttachedIndicesPages(*clouddirectory.ListAttachedIndicesInput, func(*clouddirectory.ListAttachedIndicesOutput, bool) bool) error
	ListAttachedIndicesPagesWithContext(aws.Context, *clouddirectory.ListAttachedIndicesInput, func(*clouddirectory.ListAttachedIndicesOutput, bool) bool, ...request.Option) error

	ListDevelopmentSchemaArns(*clouddirectory.ListDevelopmentSchemaArnsInput) (*clouddirectory.ListDevelopmentSchemaArnsOutput, error)
	ListDevelopmentSchemaArnsWithContext(aws.Context, *clouddirectory.ListDevelopmentSchemaArnsInput, ...request.Option) (*clouddirectory.ListDevelopmentSchemaArnsOutput, error)
	ListDevelopmentSchemaArnsRequest(*clouddirectory.ListDevelopmentSchemaArnsInput) (*request.Request, *clouddirectory.ListDevelopmentSchemaArnsOutput)

	ListDevelopmentSchemaArnsPages(*clouddirectory.ListDevelopmentSchemaArnsInput, func(*clouddirectory.ListDevelopmentSchemaArnsOutput, bool) bool) error
	ListDevelopmentSchemaArnsPagesWithContext(aws.Context, *clouddirectory.ListDevelopmentSchemaArnsInput, func(*clouddirectory.ListDevelopmentSchemaArnsOutput, bool) bool, ...request.Option) error

	ListDirectories(*clouddirectory.ListDirectoriesInput) (*clouddirectory.ListDirectoriesOutput, error)
	ListDirectoriesWithContext(aws.Context, *clouddirectory.ListDirectoriesInput, ...request.Option) (*clouddirectory.ListDirectoriesOutput, error)
	ListDirectoriesRequest(*clouddirectory.ListDirectoriesInput) (*request.Request, *clouddirectory.ListDirectoriesOutput)

	ListDirectoriesPages(*clouddirectory.ListDirectoriesInput, func(*clouddirectory.ListDirectoriesOutput, bool) bool) error
	ListDirectoriesPagesWithContext(aws.Context, *clouddirectory.ListDirectoriesInput, func(*clouddirectory.ListDirectoriesOutput, bool) bool, ...request.Option) error

	ListFacetAttributes(*clouddirectory.ListFacetAttributesInput) (*clouddirectory.ListFacetAttributesOutput, error)
	ListFacetAttributesWithContext(aws.Context, *clouddirectory.ListFacetAttributesInput, ...request.Option) (*clouddirectory.ListFacetAttributesOutput, error)
	ListFacetAttributesRequest(*clouddirectory.ListFacetAttributesInput) (*request.Request, *clouddirectory.ListFacetAttributesOutput)

	ListFacetAttributesPages(*clouddirectory.ListFacetAttributesInput, func(*clouddirectory.ListFacetAttributesOutput, bool) bool) error
	ListFacetAttributesPagesWithContext(aws.Context, *clouddirectory.ListFacetAttributesInput, func(*clouddirectory.ListFacetAttributesOutput, bool) bool, ...request.Option) error

	ListFacetNames(*clouddirectory.ListFacetNamesInput) (*clouddirectory.ListFacetNamesOutput, error)
	ListFacetNamesWithContext(aws.Context, *clouddirectory.ListFacetNamesInput, ...request.Option) (*clouddirectory.ListFacetNamesOutput, error)
	ListFacetNamesRequest(*clouddirectory.ListFacetNamesInput) (*request.Request, *clouddirectory.ListFacetNamesOutput)

	ListFacetNamesPages(*clouddirectory.ListFacetNamesInput, func(*clouddirectory.ListFacetNamesOutput, bool) bool) error
	ListFacetNamesPagesWithContext(aws.Context, *clouddirectory.ListFacetNamesInput, func(*clouddirectory.ListFacetNamesOutput, bool) bool, ...request.Option) error

	ListIncomingTypedLinks(*clouddirectory.ListIncomingTypedLinksInput) (*clouddirectory.ListIncomingTypedLinksOutput, error)
	ListIncomingTypedLinksWithContext(aws.Context, *clouddirectory.ListIncomingTypedLinksInput, ...request.Option) (*clouddirectory.ListIncomingTypedLinksOutput, error)
	ListIncomingTypedLinksRequest(*clouddirectory.ListIncomingTypedLinksInput) (*request.Request, *clouddirectory.ListIncomingTypedLinksOutput)

	ListIndex(*clouddirectory.ListIndexInput) (*clouddirectory.ListIndexOutput, error)
	ListIndexWithContext(aws.Context, *clouddirectory.ListIndexInput, ...request.Option) (*clouddirectory.ListIndexOutput, error)
	ListIndexRequest(*clouddirectory.ListIndexInput) (*request.Request, *clouddirectory.ListIndexOutput)

	ListIndexPages(*clouddirectory.ListIndexInput, func(*clouddirectory.ListIndexOutput, bool) bool) error
	ListIndexPagesWithContext(aws.Context, *clouddirectory.ListIndexInput, func(*clouddirectory.ListIndexOutput, bool) bool, ...request.Option) error

	ListManagedSchemaArns(*clouddirectory.ListManagedSchemaArnsInput) (*clouddirectory.ListManagedSchemaArnsOutput, error)
	ListManagedSchemaArnsWithContext(aws.Context, *clouddirectory.ListManagedSchemaArnsInput, ...request.Option) (*clouddirectory.ListManagedSchemaArnsOutput, error)
	ListManagedSchemaArnsRequest(*clouddirectory.ListManagedSchemaArnsInput) (*request.Request, *clouddirectory.ListManagedSchemaArnsOutput)

	ListManagedSchemaArnsPages(*clouddirectory.ListManagedSchemaArnsInput, func(*clouddirectory.ListManagedSchemaArnsOutput, bool) bool) error
	ListManagedSchemaArnsPagesWithContext(aws.Context, *clouddirectory.ListManagedSchemaArnsInput, func(*clouddirectory.ListManagedSchemaArnsOutput, bool) bool, ...request.Option) error

	ListObjectAttributes(*clouddirectory.ListObjectAttributesInput) (*clouddirectory.ListObjectAttributesOutput, error)
	ListObjectAttributesWithContext(aws.Context, *clouddirectory.ListObjectAttributesInput, ...request.Option) (*clouddirectory.ListObjectAttributesOutput, error)
	ListObjectAttributesRequest(*clouddirectory.ListObjectAttributesInput) (*request.Request, *clouddirectory.ListObjectAttributesOutput)

	ListObjectAttributesPages(*clouddirectory.ListObjectAttributesInput, func(*clouddirectory.ListObjectAttributesOutput, bool) bool) error
	ListObjectAttributesPagesWithContext(aws.Context, *clouddirectory.ListObjectAttributesInput, func(*clouddirectory.ListObjectAttributesOutput, bool) bool, ...request.Option) error

	ListObjectChildren(*clouddirectory.ListObjectChildrenInput) (*clouddirectory.ListObjectChildrenOutput, error)
	ListObjectChildrenWithContext(aws.Context, *clouddirectory.ListObjectChildrenInput, ...request.Option) (*clouddirectory.ListObjectChildrenOutput, error)
	ListObjectChildrenRequest(*clouddirectory.ListObjectChildrenInput) (*request.Request, *clouddirectory.ListObjectChildrenOutput)

	ListObjectChildrenPages(*clouddirectory.ListObjectChildrenInput, func(*clouddirectory.ListObjectChildrenOutput, bool) bool) error
	ListObjectChildrenPagesWithContext(aws.Context, *clouddirectory.ListObjectChildrenInput, func(*clouddirectory.ListObjectChildrenOutput, bool) bool, ...request.Option) error

	ListObjectParentPaths(*clouddirectory.ListObjectParentPathsInput) (*clouddirectory.ListObjectParentPathsOutput, error)
	ListObjectParentPathsWithContext(aws.Context, *clouddirectory.ListObjectParentPathsInput, ...request.Option) (*clouddirectory.ListObjectParentPathsOutput, error)
	ListObjectParentPathsRequest(*clouddirectory.ListObjectParentPathsInput) (*request.Request, *clouddirectory.ListObjectParentPathsOutput)

	ListObjectParentPathsPages(*clouddirectory.ListObjectParentPathsInput, func(*clouddirectory.ListObjectParentPathsOutput, bool) bool) error
	ListObjectParentPathsPagesWithContext(aws.Context, *clouddirectory.ListObjectParentPathsInput, func(*clouddirectory.ListObjectParentPathsOutput, bool) bool, ...request.Option) error

	ListObjectParents(*clouddirectory.ListObjectParentsInput) (*clouddirectory.ListObjectParentsOutput, error)
	ListObjectParentsWithContext(aws.Context, *clouddirectory.ListObjectParentsInput, ...request.Option) (*clouddirectory.ListObjectParentsOutput, error)
	ListObjectParentsRequest(*clouddirectory.ListObjectParentsInput) (*request.Request, *clouddirectory.ListObjectParentsOutput)

	ListObjectParentsPages(*clouddirectory.ListObjectParentsInput, func(*clouddirectory.ListObjectParentsOutput, bool) bool) error
	ListObjectParentsPagesWithContext(aws.Context, *clouddirectory.ListObjectParentsInput, func(*clouddirectory.ListObjectParentsOutput, bool) bool, ...request.Option) error

	ListObjectPolicies(*clouddirectory.ListObjectPoliciesInput) (*clouddirectory.ListObjectPoliciesOutput, error)
	ListObjectPoliciesWithContext(aws.Context, *clouddirectory.ListObjectPoliciesInput, ...request.Option) (*clouddirectory.ListObjectPoliciesOutput, error)
	ListObjectPoliciesRequest(*clouddirectory.ListObjectPoliciesInput) (*request.Request, *clouddirectory.ListObjectPoliciesOutput)

	ListObjectPoliciesPages(*clouddirectory.ListObjectPoliciesInput, func(*clouddirectory.ListObjectPoliciesOutput, bool) bool) error
	ListObjectPoliciesPagesWithContext(aws.Context, *clouddirectory.ListObjectPoliciesInput, func(*clouddirectory.ListObjectPoliciesOutput, bool) bool, ...request.Option) error

	ListOutgoingTypedLinks(*clouddirectory.ListOutgoingTypedLinksInput) (*clouddirectory.ListOutgoingTypedLinksOutput, error)
	ListOutgoingTypedLinksWithContext(aws.Context, *clouddirectory.ListOutgoingTypedLinksInput, ...request.Option) (*clouddirectory.ListOutgoingTypedLinksOutput, error)
	ListOutgoingTypedLinksRequest(*clouddirectory.ListOutgoingTypedLinksInput) (*request.Request, *clouddirectory.ListOutgoingTypedLinksOutput)

	ListPolicyAttachments(*clouddirectory.ListPolicyAttachmentsInput) (*clouddirectory.ListPolicyAttachmentsOutput, error)
	ListPolicyAttachmentsWithContext(aws.Context, *clouddirectory.ListPolicyAttachmentsInput, ...request.Option) (*clouddirectory.ListPolicyAttachmentsOutput, error)
	ListPolicyAttachmentsRequest(*clouddirectory.ListPolicyAttachmentsInput) (*request.Request, *clouddirectory.ListPolicyAttachmentsOutput)

	ListPolicyAttachmentsPages(*clouddirectory.ListPolicyAttachmentsInput, func(*clouddirectory.ListPolicyAttachmentsOutput, bool) bool) error
	ListPolicyAttachmentsPagesWithContext(aws.Context, *clouddirectory.ListPolicyAttachmentsInput, func(*clouddirectory.ListPolicyAttachmentsOutput, bool) bool, ...request.Option) error

	ListPublishedSchemaArns(*clouddirectory.ListPublishedSchemaArnsInput) (*clouddirectory.ListPublishedSchemaArnsOutput, error)
	ListPublishedSchemaArnsWithContext(aws.Context, *clouddirectory.ListPublishedSchemaArnsInput, ...request.Option) (*clouddirectory.ListPublishedSchemaArnsOutput, error)
	ListPublishedSchemaArnsRequest(*clouddirectory.ListPublishedSchemaArnsInput) (*request.Request, *clouddirectory.ListPublishedSchemaArnsOutput)

	ListPublishedSchemaArnsPages(*clouddirectory.ListPublishedSchemaArnsInput, func(*clouddirectory.ListPublishedSchemaArnsOutput, bool) bool) error
	ListPublishedSchemaArnsPagesWithContext(aws.Context, *clouddirectory.ListPublishedSchemaArnsInput, func(*clouddirectory.ListPublishedSchemaArnsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*clouddirectory.ListTagsForResourceInput) (*clouddirectory.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *clouddirectory.ListTagsForResourceInput, ...request.Option) (*clouddirectory.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*clouddirectory.ListTagsForResourceInput) (*request.Request, *clouddirectory.ListTagsForResourceOutput)

	ListTagsForResourcePages(*clouddirectory.ListTagsForResourceInput, func(*clouddirectory.ListTagsForResourceOutput, bool) bool) error
	ListTagsForResourcePagesWithContext(aws.Context, *clouddirectory.ListTagsForResourceInput, func(*clouddirectory.ListTagsForResourceOutput, bool) bool, ...request.Option) error

	ListTypedLinkFacetAttributes(*clouddirectory.ListTypedLinkFacetAttributesInput) (*clouddirectory.ListTypedLinkFacetAttributesOutput, error)
	ListTypedLinkFacetAttributesWithContext(aws.Context, *clouddirectory.ListTypedLinkFacetAttributesInput, ...request.Option) (*clouddirectory.ListTypedLinkFacetAttributesOutput, error)
	ListTypedLinkFacetAttributesRequest(*clouddirectory.ListTypedLinkFacetAttributesInput) (*request.Request, *clouddirectory.ListTypedLinkFacetAttributesOutput)

	ListTypedLinkFacetAttributesPages(*clouddirectory.ListTypedLinkFacetAttributesInput, func(*clouddirectory.ListTypedLinkFacetAttributesOutput, bool) bool) error
	ListTypedLinkFacetAttributesPagesWithContext(aws.Context, *clouddirectory.ListTypedLinkFacetAttributesInput, func(*clouddirectory.ListTypedLinkFacetAttributesOutput, bool) bool, ...request.Option) error

	ListTypedLinkFacetNames(*clouddirectory.ListTypedLinkFacetNamesInput) (*clouddirectory.ListTypedLinkFacetNamesOutput, error)
	ListTypedLinkFacetNamesWithContext(aws.Context, *clouddirectory.ListTypedLinkFacetNamesInput, ...request.Option) (*clouddirectory.ListTypedLinkFacetNamesOutput, error)
	ListTypedLinkFacetNamesRequest(*clouddirectory.ListTypedLinkFacetNamesInput) (*request.Request, *clouddirectory.ListTypedLinkFacetNamesOutput)

	ListTypedLinkFacetNamesPages(*clouddirectory.ListTypedLinkFacetNamesInput, func(*clouddirectory.ListTypedLinkFacetNamesOutput, bool) bool) error
	ListTypedLinkFacetNamesPagesWithContext(aws.Context, *clouddirectory.ListTypedLinkFacetNamesInput, func(*clouddirectory.ListTypedLinkFacetNamesOutput, bool) bool, ...request.Option) error

	LookupPolicy(*clouddirectory.LookupPolicyInput) (*clouddirectory.LookupPolicyOutput, error)
	LookupPolicyWithContext(aws.Context, *clouddirectory.LookupPolicyInput, ...request.Option) (*clouddirectory.LookupPolicyOutput, error)
	LookupPolicyRequest(*clouddirectory.LookupPolicyInput) (*request.Request, *clouddirectory.LookupPolicyOutput)

	LookupPolicyPages(*clouddirectory.LookupPolicyInput, func(*clouddirectory.LookupPolicyOutput, bool) bool) error
	LookupPolicyPagesWithContext(aws.Context, *clouddirectory.LookupPolicyInput, func(*clouddirectory.LookupPolicyOutput, bool) bool, ...request.Option) error

	PublishSchema(*clouddirectory.PublishSchemaInput) (*clouddirectory.PublishSchemaOutput, error)
	PublishSchemaWithContext(aws.Context, *clouddirectory.PublishSchemaInput, ...request.Option) (*clouddirectory.PublishSchemaOutput, error)
	PublishSchemaRequest(*clouddirectory.PublishSchemaInput) (*request.Request, *clouddirectory.PublishSchemaOutput)

	PutSchemaFromJson(*clouddirectory.PutSchemaFromJsonInput) (*clouddirectory.PutSchemaFromJsonOutput, error)
	PutSchemaFromJsonWithContext(aws.Context, *clouddirectory.PutSchemaFromJsonInput, ...request.Option) (*clouddirectory.PutSchemaFromJsonOutput, error)
	PutSchemaFromJsonRequest(*clouddirectory.PutSchemaFromJsonInput) (*request.Request, *clouddirectory.PutSchemaFromJsonOutput)

	RemoveFacetFromObject(*clouddirectory.RemoveFacetFromObjectInput) (*clouddirectory.RemoveFacetFromObjectOutput, error)
	RemoveFacetFromObjectWithContext(aws.Context, *clouddirectory.RemoveFacetFromObjectInput, ...request.Option) (*clouddirectory.RemoveFacetFromObjectOutput, error)
	RemoveFacetFromObjectRequest(*clouddirectory.RemoveFacetFromObjectInput) (*request.Request, *clouddirectory.RemoveFacetFromObjectOutput)

	TagResource(*clouddirectory.TagResourceInput) (*clouddirectory.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *clouddirectory.TagResourceInput, ...request.Option) (*clouddirectory.TagResourceOutput, error)
	TagResourceRequest(*clouddirectory.TagResourceInput) (*request.Request, *clouddirectory.TagResourceOutput)

	UntagResource(*clouddirectory.UntagResourceInput) (*clouddirectory.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *clouddirectory.UntagResourceInput, ...request.Option) (*clouddirectory.UntagResourceOutput, error)
	UntagResourceRequest(*clouddirectory.UntagResourceInput) (*request.Request, *clouddirectory.UntagResourceOutput)

	UpdateFacet(*clouddirectory.UpdateFacetInput) (*clouddirectory.UpdateFacetOutput, error)
	UpdateFacetWithContext(aws.Context, *clouddirectory.UpdateFacetInput, ...request.Option) (*clouddirectory.UpdateFacetOutput, error)
	UpdateFacetRequest(*clouddirectory.UpdateFacetInput) (*request.Request, *clouddirectory.UpdateFacetOutput)

	UpdateLinkAttributes(*clouddirectory.UpdateLinkAttributesInput) (*clouddirectory.UpdateLinkAttributesOutput, error)
	UpdateLinkAttributesWithContext(aws.Context, *clouddirectory.UpdateLinkAttributesInput, ...request.Option) (*clouddirectory.UpdateLinkAttributesOutput, error)
	UpdateLinkAttributesRequest(*clouddirectory.UpdateLinkAttributesInput) (*request.Request, *clouddirectory.UpdateLinkAttributesOutput)

	UpdateObjectAttributes(*clouddirectory.UpdateObjectAttributesInput) (*clouddirectory.UpdateObjectAttributesOutput, error)
	UpdateObjectAttributesWithContext(aws.Context, *clouddirectory.UpdateObjectAttributesInput, ...request.Option) (*clouddirectory.UpdateObjectAttributesOutput, error)
	UpdateObjectAttributesRequest(*clouddirectory.UpdateObjectAttributesInput) (*request.Request, *clouddirectory.UpdateObjectAttributesOutput)

	UpdateSchema(*clouddirectory.UpdateSchemaInput) (*clouddirectory.UpdateSchemaOutput, error)
	UpdateSchemaWithContext(aws.Context, *clouddirectory.UpdateSchemaInput, ...request.Option) (*clouddirectory.UpdateSchemaOutput, error)
	UpdateSchemaRequest(*clouddirectory.UpdateSchemaInput) (*request.Request, *clouddirectory.UpdateSchemaOutput)

	UpdateTypedLinkFacet(*clouddirectory.UpdateTypedLinkFacetInput) (*clouddirectory.UpdateTypedLinkFacetOutput, error)
	UpdateTypedLinkFacetWithContext(aws.Context, *clouddirectory.UpdateTypedLinkFacetInput, ...request.Option) (*clouddirectory.UpdateTypedLinkFacetOutput, error)
	UpdateTypedLinkFacetRequest(*clouddirectory.UpdateTypedLinkFacetInput) (*request.Request, *clouddirectory.UpdateTypedLinkFacetOutput)

	UpgradeAppliedSchema(*clouddirectory.UpgradeAppliedSchemaInput) (*clouddirectory.UpgradeAppliedSchemaOutput, error)
	UpgradeAppliedSchemaWithContext(aws.Context, *clouddirectory.UpgradeAppliedSchemaInput, ...request.Option) (*clouddirectory.UpgradeAppliedSchemaOutput, error)
	UpgradeAppliedSchemaRequest(*clouddirectory.UpgradeAppliedSchemaInput) (*request.Request, *clouddirectory.UpgradeAppliedSchemaOutput)

	UpgradePublishedSchema(*clouddirectory.UpgradePublishedSchemaInput) (*clouddirectory.UpgradePublishedSchemaOutput, error)
	UpgradePublishedSchemaWithContext(aws.Context, *clouddirectory.UpgradePublishedSchemaInput, ...request.Option) (*clouddirectory.UpgradePublishedSchemaOutput, error)
	UpgradePublishedSchemaRequest(*clouddirectory.UpgradePublishedSchemaInput) (*request.Request, *clouddirectory.UpgradePublishedSchemaOutput)
}

var _ CloudDirectoryAPI = (*clouddirectory.CloudDirectory)(nil)
