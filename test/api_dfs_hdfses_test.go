/*
XMS API

Testing DfsHdfsesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/menwenjun/xsky-eos-go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_DfsHdfsesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DfsHdfsesAPIService AddDfsHdfsACLs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsHdfsId int64

		resp, httpRes, err := apiClient.DfsHdfsesAPI.AddDfsHdfsACLs(context.Background(), dfsHdfsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsHdfsesAPIService AddDfsHdfsIPWhiteLists", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsHdfsId int64

		resp, httpRes, err := apiClient.DfsHdfsesAPI.AddDfsHdfsIPWhiteLists(context.Background(), dfsHdfsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsHdfsesAPIService AddDfsHdfsProxyUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsHdfsId int64

		resp, httpRes, err := apiClient.DfsHdfsesAPI.AddDfsHdfsProxyUsers(context.Background(), dfsHdfsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsHdfsesAPIService CreateDfsHdfs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DfsHdfsesAPI.CreateDfsHdfs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsHdfsesAPIService DeleteDfsHdfs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsHdfsId int64

		resp, httpRes, err := apiClient.DfsHdfsesAPI.DeleteDfsHdfs(context.Background(), dfsHdfsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsHdfsesAPIService GetDfsHdfs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsHdfsId int64

		resp, httpRes, err := apiClient.DfsHdfsesAPI.GetDfsHdfs(context.Background(), dfsHdfsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsHdfsesAPIService ListDfsHdfses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DfsHdfsesAPI.ListDfsHdfses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsHdfsesAPIService RemoveDfsHdfsACLs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsHdfsId int64

		resp, httpRes, err := apiClient.DfsHdfsesAPI.RemoveDfsHdfsACLs(context.Background(), dfsHdfsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsHdfsesAPIService RemoveDfsHdfsIPWhiteLists", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsHdfsId int64

		resp, httpRes, err := apiClient.DfsHdfsesAPI.RemoveDfsHdfsIPWhiteLists(context.Background(), dfsHdfsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsHdfsesAPIService RemoveDfsHdfsProxyUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsHdfsId int64

		resp, httpRes, err := apiClient.DfsHdfsesAPI.RemoveDfsHdfsProxyUsers(context.Background(), dfsHdfsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsHdfsesAPIService UpdateDfsHdfs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsHdfsId int64

		resp, httpRes, err := apiClient.DfsHdfsesAPI.UpdateDfsHdfs(context.Background(), dfsHdfsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsHdfsesAPIService UpdateDfsHdfsACLs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsHdfsId int64

		resp, httpRes, err := apiClient.DfsHdfsesAPI.UpdateDfsHdfsACLs(context.Background(), dfsHdfsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsHdfsesAPIService UpdateDfsHdfsIPWhiteLists", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsHdfsId int64

		resp, httpRes, err := apiClient.DfsHdfsesAPI.UpdateDfsHdfsIPWhiteLists(context.Background(), dfsHdfsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsHdfsesAPIService UpdateDfsHdfsProxyUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsHdfsId int64

		resp, httpRes, err := apiClient.DfsHdfsesAPI.UpdateDfsHdfsProxyUsers(context.Background(), dfsHdfsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
