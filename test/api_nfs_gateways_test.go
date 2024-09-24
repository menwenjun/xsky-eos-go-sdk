/*
XMS API

Testing NfsGatewaysAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_NfsGatewaysAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NfsGatewaysAPIService CreateNFSGateway", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NfsGatewaysAPI.CreateNFSGateway(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NfsGatewaysAPIService CreateNFSGatewayBucketMap", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gatewayId int64
		var bucketId int64

		resp, httpRes, err := apiClient.NfsGatewaysAPI.CreateNFSGatewayBucketMap(context.Background(), gatewayId, bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NfsGatewaysAPIService DeleteNFSGateway", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gatewayId int64

		resp, httpRes, err := apiClient.NfsGatewaysAPI.DeleteNFSGateway(context.Background(), gatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NfsGatewaysAPIService DeleteNFSGatewayBucketMap", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gatewayId int64
		var bucketId int64

		resp, httpRes, err := apiClient.NfsGatewaysAPI.DeleteNFSGatewayBucketMap(context.Background(), gatewayId, bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NfsGatewaysAPIService DoNFSGateway", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gatewayId int64

		resp, httpRes, err := apiClient.NfsGatewaysAPI.DoNFSGateway(context.Background(), gatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NfsGatewaysAPIService GetNFSGateway", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gatewayId int64

		resp, httpRes, err := apiClient.NfsGatewaysAPI.GetNFSGateway(context.Background(), gatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NfsGatewaysAPIService GetNFSGatewayBucketMap", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gatewayId int64
		var bucketId int64

		resp, httpRes, err := apiClient.NfsGatewaysAPI.GetNFSGatewayBucketMap(context.Background(), gatewayId, bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NfsGatewaysAPIService GetNFSGatewaySamples", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gatewayId int64

		resp, httpRes, err := apiClient.NfsGatewaysAPI.GetNFSGatewaySamples(context.Background(), gatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NfsGatewaysAPIService ListNFSGatewayBucketMaps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gatewayId int64

		resp, httpRes, err := apiClient.NfsGatewaysAPI.ListNFSGatewayBucketMaps(context.Background(), gatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NfsGatewaysAPIService ListNFSGateways", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NfsGatewaysAPI.ListNFSGateways(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NfsGatewaysAPIService UpdateNFSGateway", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gatewayId int64

		resp, httpRes, err := apiClient.NfsGatewaysAPI.UpdateNFSGateway(context.Background(), gatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NfsGatewaysAPIService UpdateNFSGatewayBucketMap", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gatewayId int64
		var bucketId int64

		resp, httpRes, err := apiClient.NfsGatewaysAPI.UpdateNFSGatewayBucketMap(context.Background(), gatewayId, bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NfsGatewaysAPIService UpdateOspExportConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gatewayId int64

		resp, httpRes, err := apiClient.NfsGatewaysAPI.UpdateOspExportConfig(context.Background(), gatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
