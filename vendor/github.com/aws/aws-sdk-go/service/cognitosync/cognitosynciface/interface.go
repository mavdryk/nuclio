// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cognitosynciface provides an interface to enable mocking the Amazon Cognito Sync service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cognitosynciface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cognitosync"
)

// CognitoSyncAPI provides an interface to enable mocking the
// cognitosync.CognitoSync service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Cognito Sync.
//    func myFunc(svc cognitosynciface.CognitoSyncAPI) bool {
//        // Make svc.BulkPublish request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cognitosync.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCognitoSyncClient struct {
//        cognitosynciface.CognitoSyncAPI
//    }
//    func (m *mockCognitoSyncClient) BulkPublish(input *cognitosync.BulkPublishInput) (*cognitosync.BulkPublishOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCognitoSyncClient{}
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
type CognitoSyncAPI interface {
	BulkPublish(*cognitosync.BulkPublishInput) (*cognitosync.BulkPublishOutput, error)
	BulkPublishWithContext(aws.Context, *cognitosync.BulkPublishInput, ...request.Option) (*cognitosync.BulkPublishOutput, error)
	BulkPublishRequest(*cognitosync.BulkPublishInput) (*request.Request, *cognitosync.BulkPublishOutput)

	DeleteDataset(*cognitosync.DeleteDatasetInput) (*cognitosync.DeleteDatasetOutput, error)
	DeleteDatasetWithContext(aws.Context, *cognitosync.DeleteDatasetInput, ...request.Option) (*cognitosync.DeleteDatasetOutput, error)
	DeleteDatasetRequest(*cognitosync.DeleteDatasetInput) (*request.Request, *cognitosync.DeleteDatasetOutput)

	DescribeDataset(*cognitosync.DescribeDatasetInput) (*cognitosync.DescribeDatasetOutput, error)
	DescribeDatasetWithContext(aws.Context, *cognitosync.DescribeDatasetInput, ...request.Option) (*cognitosync.DescribeDatasetOutput, error)
	DescribeDatasetRequest(*cognitosync.DescribeDatasetInput) (*request.Request, *cognitosync.DescribeDatasetOutput)

	DescribeIdentityPoolUsage(*cognitosync.DescribeIdentityPoolUsageInput) (*cognitosync.DescribeIdentityPoolUsageOutput, error)
	DescribeIdentityPoolUsageWithContext(aws.Context, *cognitosync.DescribeIdentityPoolUsageInput, ...request.Option) (*cognitosync.DescribeIdentityPoolUsageOutput, error)
	DescribeIdentityPoolUsageRequest(*cognitosync.DescribeIdentityPoolUsageInput) (*request.Request, *cognitosync.DescribeIdentityPoolUsageOutput)

	DescribeIdentityUsage(*cognitosync.DescribeIdentityUsageInput) (*cognitosync.DescribeIdentityUsageOutput, error)
	DescribeIdentityUsageWithContext(aws.Context, *cognitosync.DescribeIdentityUsageInput, ...request.Option) (*cognitosync.DescribeIdentityUsageOutput, error)
	DescribeIdentityUsageRequest(*cognitosync.DescribeIdentityUsageInput) (*request.Request, *cognitosync.DescribeIdentityUsageOutput)

	GetBulkPublishDetails(*cognitosync.GetBulkPublishDetailsInput) (*cognitosync.GetBulkPublishDetailsOutput, error)
	GetBulkPublishDetailsWithContext(aws.Context, *cognitosync.GetBulkPublishDetailsInput, ...request.Option) (*cognitosync.GetBulkPublishDetailsOutput, error)
	GetBulkPublishDetailsRequest(*cognitosync.GetBulkPublishDetailsInput) (*request.Request, *cognitosync.GetBulkPublishDetailsOutput)

	GetCognitoEvents(*cognitosync.GetCognitoEventsInput) (*cognitosync.GetCognitoEventsOutput, error)
	GetCognitoEventsWithContext(aws.Context, *cognitosync.GetCognitoEventsInput, ...request.Option) (*cognitosync.GetCognitoEventsOutput, error)
	GetCognitoEventsRequest(*cognitosync.GetCognitoEventsInput) (*request.Request, *cognitosync.GetCognitoEventsOutput)

	GetIdentityPoolConfiguration(*cognitosync.GetIdentityPoolConfigurationInput) (*cognitosync.GetIdentityPoolConfigurationOutput, error)
	GetIdentityPoolConfigurationWithContext(aws.Context, *cognitosync.GetIdentityPoolConfigurationInput, ...request.Option) (*cognitosync.GetIdentityPoolConfigurationOutput, error)
	GetIdentityPoolConfigurationRequest(*cognitosync.GetIdentityPoolConfigurationInput) (*request.Request, *cognitosync.GetIdentityPoolConfigurationOutput)

	ListDatasets(*cognitosync.ListDatasetsInput) (*cognitosync.ListDatasetsOutput, error)
	ListDatasetsWithContext(aws.Context, *cognitosync.ListDatasetsInput, ...request.Option) (*cognitosync.ListDatasetsOutput, error)
	ListDatasetsRequest(*cognitosync.ListDatasetsInput) (*request.Request, *cognitosync.ListDatasetsOutput)

	ListIdentityPoolUsage(*cognitosync.ListIdentityPoolUsageInput) (*cognitosync.ListIdentityPoolUsageOutput, error)
	ListIdentityPoolUsageWithContext(aws.Context, *cognitosync.ListIdentityPoolUsageInput, ...request.Option) (*cognitosync.ListIdentityPoolUsageOutput, error)
	ListIdentityPoolUsageRequest(*cognitosync.ListIdentityPoolUsageInput) (*request.Request, *cognitosync.ListIdentityPoolUsageOutput)

	ListRecords(*cognitosync.ListRecordsInput) (*cognitosync.ListRecordsOutput, error)
	ListRecordsWithContext(aws.Context, *cognitosync.ListRecordsInput, ...request.Option) (*cognitosync.ListRecordsOutput, error)
	ListRecordsRequest(*cognitosync.ListRecordsInput) (*request.Request, *cognitosync.ListRecordsOutput)

	RegisterDevice(*cognitosync.RegisterDeviceInput) (*cognitosync.RegisterDeviceOutput, error)
	RegisterDeviceWithContext(aws.Context, *cognitosync.RegisterDeviceInput, ...request.Option) (*cognitosync.RegisterDeviceOutput, error)
	RegisterDeviceRequest(*cognitosync.RegisterDeviceInput) (*request.Request, *cognitosync.RegisterDeviceOutput)

	SetCognitoEvents(*cognitosync.SetCognitoEventsInput) (*cognitosync.SetCognitoEventsOutput, error)
	SetCognitoEventsWithContext(aws.Context, *cognitosync.SetCognitoEventsInput, ...request.Option) (*cognitosync.SetCognitoEventsOutput, error)
	SetCognitoEventsRequest(*cognitosync.SetCognitoEventsInput) (*request.Request, *cognitosync.SetCognitoEventsOutput)

	SetIdentityPoolConfiguration(*cognitosync.SetIdentityPoolConfigurationInput) (*cognitosync.SetIdentityPoolConfigurationOutput, error)
	SetIdentityPoolConfigurationWithContext(aws.Context, *cognitosync.SetIdentityPoolConfigurationInput, ...request.Option) (*cognitosync.SetIdentityPoolConfigurationOutput, error)
	SetIdentityPoolConfigurationRequest(*cognitosync.SetIdentityPoolConfigurationInput) (*request.Request, *cognitosync.SetIdentityPoolConfigurationOutput)

	SubscribeToDataset(*cognitosync.SubscribeToDatasetInput) (*cognitosync.SubscribeToDatasetOutput, error)
	SubscribeToDatasetWithContext(aws.Context, *cognitosync.SubscribeToDatasetInput, ...request.Option) (*cognitosync.SubscribeToDatasetOutput, error)
	SubscribeToDatasetRequest(*cognitosync.SubscribeToDatasetInput) (*request.Request, *cognitosync.SubscribeToDatasetOutput)

	UnsubscribeFromDataset(*cognitosync.UnsubscribeFromDatasetInput) (*cognitosync.UnsubscribeFromDatasetOutput, error)
	UnsubscribeFromDatasetWithContext(aws.Context, *cognitosync.UnsubscribeFromDatasetInput, ...request.Option) (*cognitosync.UnsubscribeFromDatasetOutput, error)
	UnsubscribeFromDatasetRequest(*cognitosync.UnsubscribeFromDatasetInput) (*request.Request, *cognitosync.UnsubscribeFromDatasetOutput)

	UpdateRecords(*cognitosync.UpdateRecordsInput) (*cognitosync.UpdateRecordsOutput, error)
	UpdateRecordsWithContext(aws.Context, *cognitosync.UpdateRecordsInput, ...request.Option) (*cognitosync.UpdateRecordsOutput, error)
	UpdateRecordsRequest(*cognitosync.UpdateRecordsInput) (*request.Request, *cognitosync.UpdateRecordsOutput)
}

var _ CognitoSyncAPI = (*cognitosync.CognitoSync)(nil)
