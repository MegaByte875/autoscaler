// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package lakeformationiface provides an interface to enable mocking the AWS Lake Formation service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package lakeformationiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/lakeformation"
)

// LakeFormationAPI provides an interface to enable mocking the
// lakeformation.LakeFormation service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Lake Formation.
//    func myFunc(svc lakeformationiface.LakeFormationAPI) bool {
//        // Make svc.AddLFTagsToResource request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := lakeformation.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockLakeFormationClient struct {
//        lakeformationiface.LakeFormationAPI
//    }
//    func (m *mockLakeFormationClient) AddLFTagsToResource(input *lakeformation.AddLFTagsToResourceInput) (*lakeformation.AddLFTagsToResourceOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockLakeFormationClient{}
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
type LakeFormationAPI interface {
	AddLFTagsToResource(*lakeformation.AddLFTagsToResourceInput) (*lakeformation.AddLFTagsToResourceOutput, error)
	AddLFTagsToResourceWithContext(aws.Context, *lakeformation.AddLFTagsToResourceInput, ...request.Option) (*lakeformation.AddLFTagsToResourceOutput, error)
	AddLFTagsToResourceRequest(*lakeformation.AddLFTagsToResourceInput) (*request.Request, *lakeformation.AddLFTagsToResourceOutput)

	BatchGrantPermissions(*lakeformation.BatchGrantPermissionsInput) (*lakeformation.BatchGrantPermissionsOutput, error)
	BatchGrantPermissionsWithContext(aws.Context, *lakeformation.BatchGrantPermissionsInput, ...request.Option) (*lakeformation.BatchGrantPermissionsOutput, error)
	BatchGrantPermissionsRequest(*lakeformation.BatchGrantPermissionsInput) (*request.Request, *lakeformation.BatchGrantPermissionsOutput)

	BatchRevokePermissions(*lakeformation.BatchRevokePermissionsInput) (*lakeformation.BatchRevokePermissionsOutput, error)
	BatchRevokePermissionsWithContext(aws.Context, *lakeformation.BatchRevokePermissionsInput, ...request.Option) (*lakeformation.BatchRevokePermissionsOutput, error)
	BatchRevokePermissionsRequest(*lakeformation.BatchRevokePermissionsInput) (*request.Request, *lakeformation.BatchRevokePermissionsOutput)

	CancelTransaction(*lakeformation.CancelTransactionInput) (*lakeformation.CancelTransactionOutput, error)
	CancelTransactionWithContext(aws.Context, *lakeformation.CancelTransactionInput, ...request.Option) (*lakeformation.CancelTransactionOutput, error)
	CancelTransactionRequest(*lakeformation.CancelTransactionInput) (*request.Request, *lakeformation.CancelTransactionOutput)

	CommitTransaction(*lakeformation.CommitTransactionInput) (*lakeformation.CommitTransactionOutput, error)
	CommitTransactionWithContext(aws.Context, *lakeformation.CommitTransactionInput, ...request.Option) (*lakeformation.CommitTransactionOutput, error)
	CommitTransactionRequest(*lakeformation.CommitTransactionInput) (*request.Request, *lakeformation.CommitTransactionOutput)

	CreateDataCellsFilter(*lakeformation.CreateDataCellsFilterInput) (*lakeformation.CreateDataCellsFilterOutput, error)
	CreateDataCellsFilterWithContext(aws.Context, *lakeformation.CreateDataCellsFilterInput, ...request.Option) (*lakeformation.CreateDataCellsFilterOutput, error)
	CreateDataCellsFilterRequest(*lakeformation.CreateDataCellsFilterInput) (*request.Request, *lakeformation.CreateDataCellsFilterOutput)

	CreateLFTag(*lakeformation.CreateLFTagInput) (*lakeformation.CreateLFTagOutput, error)
	CreateLFTagWithContext(aws.Context, *lakeformation.CreateLFTagInput, ...request.Option) (*lakeformation.CreateLFTagOutput, error)
	CreateLFTagRequest(*lakeformation.CreateLFTagInput) (*request.Request, *lakeformation.CreateLFTagOutput)

	DeleteDataCellsFilter(*lakeformation.DeleteDataCellsFilterInput) (*lakeformation.DeleteDataCellsFilterOutput, error)
	DeleteDataCellsFilterWithContext(aws.Context, *lakeformation.DeleteDataCellsFilterInput, ...request.Option) (*lakeformation.DeleteDataCellsFilterOutput, error)
	DeleteDataCellsFilterRequest(*lakeformation.DeleteDataCellsFilterInput) (*request.Request, *lakeformation.DeleteDataCellsFilterOutput)

	DeleteLFTag(*lakeformation.DeleteLFTagInput) (*lakeformation.DeleteLFTagOutput, error)
	DeleteLFTagWithContext(aws.Context, *lakeformation.DeleteLFTagInput, ...request.Option) (*lakeformation.DeleteLFTagOutput, error)
	DeleteLFTagRequest(*lakeformation.DeleteLFTagInput) (*request.Request, *lakeformation.DeleteLFTagOutput)

	DeleteObjectsOnCancel(*lakeformation.DeleteObjectsOnCancelInput) (*lakeformation.DeleteObjectsOnCancelOutput, error)
	DeleteObjectsOnCancelWithContext(aws.Context, *lakeformation.DeleteObjectsOnCancelInput, ...request.Option) (*lakeformation.DeleteObjectsOnCancelOutput, error)
	DeleteObjectsOnCancelRequest(*lakeformation.DeleteObjectsOnCancelInput) (*request.Request, *lakeformation.DeleteObjectsOnCancelOutput)

	DeregisterResource(*lakeformation.DeregisterResourceInput) (*lakeformation.DeregisterResourceOutput, error)
	DeregisterResourceWithContext(aws.Context, *lakeformation.DeregisterResourceInput, ...request.Option) (*lakeformation.DeregisterResourceOutput, error)
	DeregisterResourceRequest(*lakeformation.DeregisterResourceInput) (*request.Request, *lakeformation.DeregisterResourceOutput)

	DescribeResource(*lakeformation.DescribeResourceInput) (*lakeformation.DescribeResourceOutput, error)
	DescribeResourceWithContext(aws.Context, *lakeformation.DescribeResourceInput, ...request.Option) (*lakeformation.DescribeResourceOutput, error)
	DescribeResourceRequest(*lakeformation.DescribeResourceInput) (*request.Request, *lakeformation.DescribeResourceOutput)

	DescribeTransaction(*lakeformation.DescribeTransactionInput) (*lakeformation.DescribeTransactionOutput, error)
	DescribeTransactionWithContext(aws.Context, *lakeformation.DescribeTransactionInput, ...request.Option) (*lakeformation.DescribeTransactionOutput, error)
	DescribeTransactionRequest(*lakeformation.DescribeTransactionInput) (*request.Request, *lakeformation.DescribeTransactionOutput)

	ExtendTransaction(*lakeformation.ExtendTransactionInput) (*lakeformation.ExtendTransactionOutput, error)
	ExtendTransactionWithContext(aws.Context, *lakeformation.ExtendTransactionInput, ...request.Option) (*lakeformation.ExtendTransactionOutput, error)
	ExtendTransactionRequest(*lakeformation.ExtendTransactionInput) (*request.Request, *lakeformation.ExtendTransactionOutput)

	GetDataLakeSettings(*lakeformation.GetDataLakeSettingsInput) (*lakeformation.GetDataLakeSettingsOutput, error)
	GetDataLakeSettingsWithContext(aws.Context, *lakeformation.GetDataLakeSettingsInput, ...request.Option) (*lakeformation.GetDataLakeSettingsOutput, error)
	GetDataLakeSettingsRequest(*lakeformation.GetDataLakeSettingsInput) (*request.Request, *lakeformation.GetDataLakeSettingsOutput)

	GetEffectivePermissionsForPath(*lakeformation.GetEffectivePermissionsForPathInput) (*lakeformation.GetEffectivePermissionsForPathOutput, error)
	GetEffectivePermissionsForPathWithContext(aws.Context, *lakeformation.GetEffectivePermissionsForPathInput, ...request.Option) (*lakeformation.GetEffectivePermissionsForPathOutput, error)
	GetEffectivePermissionsForPathRequest(*lakeformation.GetEffectivePermissionsForPathInput) (*request.Request, *lakeformation.GetEffectivePermissionsForPathOutput)

	GetEffectivePermissionsForPathPages(*lakeformation.GetEffectivePermissionsForPathInput, func(*lakeformation.GetEffectivePermissionsForPathOutput, bool) bool) error
	GetEffectivePermissionsForPathPagesWithContext(aws.Context, *lakeformation.GetEffectivePermissionsForPathInput, func(*lakeformation.GetEffectivePermissionsForPathOutput, bool) bool, ...request.Option) error

	GetLFTag(*lakeformation.GetLFTagInput) (*lakeformation.GetLFTagOutput, error)
	GetLFTagWithContext(aws.Context, *lakeformation.GetLFTagInput, ...request.Option) (*lakeformation.GetLFTagOutput, error)
	GetLFTagRequest(*lakeformation.GetLFTagInput) (*request.Request, *lakeformation.GetLFTagOutput)

	GetQueryState(*lakeformation.GetQueryStateInput) (*lakeformation.GetQueryStateOutput, error)
	GetQueryStateWithContext(aws.Context, *lakeformation.GetQueryStateInput, ...request.Option) (*lakeformation.GetQueryStateOutput, error)
	GetQueryStateRequest(*lakeformation.GetQueryStateInput) (*request.Request, *lakeformation.GetQueryStateOutput)

	GetQueryStatistics(*lakeformation.GetQueryStatisticsInput) (*lakeformation.GetQueryStatisticsOutput, error)
	GetQueryStatisticsWithContext(aws.Context, *lakeformation.GetQueryStatisticsInput, ...request.Option) (*lakeformation.GetQueryStatisticsOutput, error)
	GetQueryStatisticsRequest(*lakeformation.GetQueryStatisticsInput) (*request.Request, *lakeformation.GetQueryStatisticsOutput)

	GetResourceLFTags(*lakeformation.GetResourceLFTagsInput) (*lakeformation.GetResourceLFTagsOutput, error)
	GetResourceLFTagsWithContext(aws.Context, *lakeformation.GetResourceLFTagsInput, ...request.Option) (*lakeformation.GetResourceLFTagsOutput, error)
	GetResourceLFTagsRequest(*lakeformation.GetResourceLFTagsInput) (*request.Request, *lakeformation.GetResourceLFTagsOutput)

	GetTableObjects(*lakeformation.GetTableObjectsInput) (*lakeformation.GetTableObjectsOutput, error)
	GetTableObjectsWithContext(aws.Context, *lakeformation.GetTableObjectsInput, ...request.Option) (*lakeformation.GetTableObjectsOutput, error)
	GetTableObjectsRequest(*lakeformation.GetTableObjectsInput) (*request.Request, *lakeformation.GetTableObjectsOutput)

	GetTableObjectsPages(*lakeformation.GetTableObjectsInput, func(*lakeformation.GetTableObjectsOutput, bool) bool) error
	GetTableObjectsPagesWithContext(aws.Context, *lakeformation.GetTableObjectsInput, func(*lakeformation.GetTableObjectsOutput, bool) bool, ...request.Option) error

	GetTemporaryGluePartitionCredentials(*lakeformation.GetTemporaryGluePartitionCredentialsInput) (*lakeformation.GetTemporaryGluePartitionCredentialsOutput, error)
	GetTemporaryGluePartitionCredentialsWithContext(aws.Context, *lakeformation.GetTemporaryGluePartitionCredentialsInput, ...request.Option) (*lakeformation.GetTemporaryGluePartitionCredentialsOutput, error)
	GetTemporaryGluePartitionCredentialsRequest(*lakeformation.GetTemporaryGluePartitionCredentialsInput) (*request.Request, *lakeformation.GetTemporaryGluePartitionCredentialsOutput)

	GetTemporaryGlueTableCredentials(*lakeformation.GetTemporaryGlueTableCredentialsInput) (*lakeformation.GetTemporaryGlueTableCredentialsOutput, error)
	GetTemporaryGlueTableCredentialsWithContext(aws.Context, *lakeformation.GetTemporaryGlueTableCredentialsInput, ...request.Option) (*lakeformation.GetTemporaryGlueTableCredentialsOutput, error)
	GetTemporaryGlueTableCredentialsRequest(*lakeformation.GetTemporaryGlueTableCredentialsInput) (*request.Request, *lakeformation.GetTemporaryGlueTableCredentialsOutput)

	GetWorkUnitResults(*lakeformation.GetWorkUnitResultsInput) (*lakeformation.GetWorkUnitResultsOutput, error)
	GetWorkUnitResultsWithContext(aws.Context, *lakeformation.GetWorkUnitResultsInput, ...request.Option) (*lakeformation.GetWorkUnitResultsOutput, error)
	GetWorkUnitResultsRequest(*lakeformation.GetWorkUnitResultsInput) (*request.Request, *lakeformation.GetWorkUnitResultsOutput)

	GetWorkUnits(*lakeformation.GetWorkUnitsInput) (*lakeformation.GetWorkUnitsOutput, error)
	GetWorkUnitsWithContext(aws.Context, *lakeformation.GetWorkUnitsInput, ...request.Option) (*lakeformation.GetWorkUnitsOutput, error)
	GetWorkUnitsRequest(*lakeformation.GetWorkUnitsInput) (*request.Request, *lakeformation.GetWorkUnitsOutput)

	GetWorkUnitsPages(*lakeformation.GetWorkUnitsInput, func(*lakeformation.GetWorkUnitsOutput, bool) bool) error
	GetWorkUnitsPagesWithContext(aws.Context, *lakeformation.GetWorkUnitsInput, func(*lakeformation.GetWorkUnitsOutput, bool) bool, ...request.Option) error

	GrantPermissions(*lakeformation.GrantPermissionsInput) (*lakeformation.GrantPermissionsOutput, error)
	GrantPermissionsWithContext(aws.Context, *lakeformation.GrantPermissionsInput, ...request.Option) (*lakeformation.GrantPermissionsOutput, error)
	GrantPermissionsRequest(*lakeformation.GrantPermissionsInput) (*request.Request, *lakeformation.GrantPermissionsOutput)

	ListDataCellsFilter(*lakeformation.ListDataCellsFilterInput) (*lakeformation.ListDataCellsFilterOutput, error)
	ListDataCellsFilterWithContext(aws.Context, *lakeformation.ListDataCellsFilterInput, ...request.Option) (*lakeformation.ListDataCellsFilterOutput, error)
	ListDataCellsFilterRequest(*lakeformation.ListDataCellsFilterInput) (*request.Request, *lakeformation.ListDataCellsFilterOutput)

	ListDataCellsFilterPages(*lakeformation.ListDataCellsFilterInput, func(*lakeformation.ListDataCellsFilterOutput, bool) bool) error
	ListDataCellsFilterPagesWithContext(aws.Context, *lakeformation.ListDataCellsFilterInput, func(*lakeformation.ListDataCellsFilterOutput, bool) bool, ...request.Option) error

	ListLFTags(*lakeformation.ListLFTagsInput) (*lakeformation.ListLFTagsOutput, error)
	ListLFTagsWithContext(aws.Context, *lakeformation.ListLFTagsInput, ...request.Option) (*lakeformation.ListLFTagsOutput, error)
	ListLFTagsRequest(*lakeformation.ListLFTagsInput) (*request.Request, *lakeformation.ListLFTagsOutput)

	ListLFTagsPages(*lakeformation.ListLFTagsInput, func(*lakeformation.ListLFTagsOutput, bool) bool) error
	ListLFTagsPagesWithContext(aws.Context, *lakeformation.ListLFTagsInput, func(*lakeformation.ListLFTagsOutput, bool) bool, ...request.Option) error

	ListPermissions(*lakeformation.ListPermissionsInput) (*lakeformation.ListPermissionsOutput, error)
	ListPermissionsWithContext(aws.Context, *lakeformation.ListPermissionsInput, ...request.Option) (*lakeformation.ListPermissionsOutput, error)
	ListPermissionsRequest(*lakeformation.ListPermissionsInput) (*request.Request, *lakeformation.ListPermissionsOutput)

	ListPermissionsPages(*lakeformation.ListPermissionsInput, func(*lakeformation.ListPermissionsOutput, bool) bool) error
	ListPermissionsPagesWithContext(aws.Context, *lakeformation.ListPermissionsInput, func(*lakeformation.ListPermissionsOutput, bool) bool, ...request.Option) error

	ListResources(*lakeformation.ListResourcesInput) (*lakeformation.ListResourcesOutput, error)
	ListResourcesWithContext(aws.Context, *lakeformation.ListResourcesInput, ...request.Option) (*lakeformation.ListResourcesOutput, error)
	ListResourcesRequest(*lakeformation.ListResourcesInput) (*request.Request, *lakeformation.ListResourcesOutput)

	ListResourcesPages(*lakeformation.ListResourcesInput, func(*lakeformation.ListResourcesOutput, bool) bool) error
	ListResourcesPagesWithContext(aws.Context, *lakeformation.ListResourcesInput, func(*lakeformation.ListResourcesOutput, bool) bool, ...request.Option) error

	ListTableStorageOptimizers(*lakeformation.ListTableStorageOptimizersInput) (*lakeformation.ListTableStorageOptimizersOutput, error)
	ListTableStorageOptimizersWithContext(aws.Context, *lakeformation.ListTableStorageOptimizersInput, ...request.Option) (*lakeformation.ListTableStorageOptimizersOutput, error)
	ListTableStorageOptimizersRequest(*lakeformation.ListTableStorageOptimizersInput) (*request.Request, *lakeformation.ListTableStorageOptimizersOutput)

	ListTableStorageOptimizersPages(*lakeformation.ListTableStorageOptimizersInput, func(*lakeformation.ListTableStorageOptimizersOutput, bool) bool) error
	ListTableStorageOptimizersPagesWithContext(aws.Context, *lakeformation.ListTableStorageOptimizersInput, func(*lakeformation.ListTableStorageOptimizersOutput, bool) bool, ...request.Option) error

	ListTransactions(*lakeformation.ListTransactionsInput) (*lakeformation.ListTransactionsOutput, error)
	ListTransactionsWithContext(aws.Context, *lakeformation.ListTransactionsInput, ...request.Option) (*lakeformation.ListTransactionsOutput, error)
	ListTransactionsRequest(*lakeformation.ListTransactionsInput) (*request.Request, *lakeformation.ListTransactionsOutput)

	ListTransactionsPages(*lakeformation.ListTransactionsInput, func(*lakeformation.ListTransactionsOutput, bool) bool) error
	ListTransactionsPagesWithContext(aws.Context, *lakeformation.ListTransactionsInput, func(*lakeformation.ListTransactionsOutput, bool) bool, ...request.Option) error

	PutDataLakeSettings(*lakeformation.PutDataLakeSettingsInput) (*lakeformation.PutDataLakeSettingsOutput, error)
	PutDataLakeSettingsWithContext(aws.Context, *lakeformation.PutDataLakeSettingsInput, ...request.Option) (*lakeformation.PutDataLakeSettingsOutput, error)
	PutDataLakeSettingsRequest(*lakeformation.PutDataLakeSettingsInput) (*request.Request, *lakeformation.PutDataLakeSettingsOutput)

	RegisterResource(*lakeformation.RegisterResourceInput) (*lakeformation.RegisterResourceOutput, error)
	RegisterResourceWithContext(aws.Context, *lakeformation.RegisterResourceInput, ...request.Option) (*lakeformation.RegisterResourceOutput, error)
	RegisterResourceRequest(*lakeformation.RegisterResourceInput) (*request.Request, *lakeformation.RegisterResourceOutput)

	RemoveLFTagsFromResource(*lakeformation.RemoveLFTagsFromResourceInput) (*lakeformation.RemoveLFTagsFromResourceOutput, error)
	RemoveLFTagsFromResourceWithContext(aws.Context, *lakeformation.RemoveLFTagsFromResourceInput, ...request.Option) (*lakeformation.RemoveLFTagsFromResourceOutput, error)
	RemoveLFTagsFromResourceRequest(*lakeformation.RemoveLFTagsFromResourceInput) (*request.Request, *lakeformation.RemoveLFTagsFromResourceOutput)

	RevokePermissions(*lakeformation.RevokePermissionsInput) (*lakeformation.RevokePermissionsOutput, error)
	RevokePermissionsWithContext(aws.Context, *lakeformation.RevokePermissionsInput, ...request.Option) (*lakeformation.RevokePermissionsOutput, error)
	RevokePermissionsRequest(*lakeformation.RevokePermissionsInput) (*request.Request, *lakeformation.RevokePermissionsOutput)

	SearchDatabasesByLFTags(*lakeformation.SearchDatabasesByLFTagsInput) (*lakeformation.SearchDatabasesByLFTagsOutput, error)
	SearchDatabasesByLFTagsWithContext(aws.Context, *lakeformation.SearchDatabasesByLFTagsInput, ...request.Option) (*lakeformation.SearchDatabasesByLFTagsOutput, error)
	SearchDatabasesByLFTagsRequest(*lakeformation.SearchDatabasesByLFTagsInput) (*request.Request, *lakeformation.SearchDatabasesByLFTagsOutput)

	SearchDatabasesByLFTagsPages(*lakeformation.SearchDatabasesByLFTagsInput, func(*lakeformation.SearchDatabasesByLFTagsOutput, bool) bool) error
	SearchDatabasesByLFTagsPagesWithContext(aws.Context, *lakeformation.SearchDatabasesByLFTagsInput, func(*lakeformation.SearchDatabasesByLFTagsOutput, bool) bool, ...request.Option) error

	SearchTablesByLFTags(*lakeformation.SearchTablesByLFTagsInput) (*lakeformation.SearchTablesByLFTagsOutput, error)
	SearchTablesByLFTagsWithContext(aws.Context, *lakeformation.SearchTablesByLFTagsInput, ...request.Option) (*lakeformation.SearchTablesByLFTagsOutput, error)
	SearchTablesByLFTagsRequest(*lakeformation.SearchTablesByLFTagsInput) (*request.Request, *lakeformation.SearchTablesByLFTagsOutput)

	SearchTablesByLFTagsPages(*lakeformation.SearchTablesByLFTagsInput, func(*lakeformation.SearchTablesByLFTagsOutput, bool) bool) error
	SearchTablesByLFTagsPagesWithContext(aws.Context, *lakeformation.SearchTablesByLFTagsInput, func(*lakeformation.SearchTablesByLFTagsOutput, bool) bool, ...request.Option) error

	StartQueryPlanning(*lakeformation.StartQueryPlanningInput) (*lakeformation.StartQueryPlanningOutput, error)
	StartQueryPlanningWithContext(aws.Context, *lakeformation.StartQueryPlanningInput, ...request.Option) (*lakeformation.StartQueryPlanningOutput, error)
	StartQueryPlanningRequest(*lakeformation.StartQueryPlanningInput) (*request.Request, *lakeformation.StartQueryPlanningOutput)

	StartTransaction(*lakeformation.StartTransactionInput) (*lakeformation.StartTransactionOutput, error)
	StartTransactionWithContext(aws.Context, *lakeformation.StartTransactionInput, ...request.Option) (*lakeformation.StartTransactionOutput, error)
	StartTransactionRequest(*lakeformation.StartTransactionInput) (*request.Request, *lakeformation.StartTransactionOutput)

	UpdateLFTag(*lakeformation.UpdateLFTagInput) (*lakeformation.UpdateLFTagOutput, error)
	UpdateLFTagWithContext(aws.Context, *lakeformation.UpdateLFTagInput, ...request.Option) (*lakeformation.UpdateLFTagOutput, error)
	UpdateLFTagRequest(*lakeformation.UpdateLFTagInput) (*request.Request, *lakeformation.UpdateLFTagOutput)

	UpdateResource(*lakeformation.UpdateResourceInput) (*lakeformation.UpdateResourceOutput, error)
	UpdateResourceWithContext(aws.Context, *lakeformation.UpdateResourceInput, ...request.Option) (*lakeformation.UpdateResourceOutput, error)
	UpdateResourceRequest(*lakeformation.UpdateResourceInput) (*request.Request, *lakeformation.UpdateResourceOutput)

	UpdateTableObjects(*lakeformation.UpdateTableObjectsInput) (*lakeformation.UpdateTableObjectsOutput, error)
	UpdateTableObjectsWithContext(aws.Context, *lakeformation.UpdateTableObjectsInput, ...request.Option) (*lakeformation.UpdateTableObjectsOutput, error)
	UpdateTableObjectsRequest(*lakeformation.UpdateTableObjectsInput) (*request.Request, *lakeformation.UpdateTableObjectsOutput)

	UpdateTableStorageOptimizer(*lakeformation.UpdateTableStorageOptimizerInput) (*lakeformation.UpdateTableStorageOptimizerOutput, error)
	UpdateTableStorageOptimizerWithContext(aws.Context, *lakeformation.UpdateTableStorageOptimizerInput, ...request.Option) (*lakeformation.UpdateTableStorageOptimizerOutput, error)
	UpdateTableStorageOptimizerRequest(*lakeformation.UpdateTableStorageOptimizerInput) (*request.Request, *lakeformation.UpdateTableStorageOptimizerOutput)
}

var _ LakeFormationAPI = (*lakeformation.LakeFormation)(nil)
