/*
XMS API

Testing DisksAPIService

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

func Test_openapi_DisksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DisksAPIService CreatePartitions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var diskId int64

		resp, httpRes, err := apiClient.DisksAPI.CreatePartitions(context.Background(), diskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DisksAPIService DeletePartitions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var diskId int64

		resp, httpRes, err := apiClient.DisksAPI.DeletePartitions(context.Background(), diskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DisksAPIService GetDisk", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var diskId int64

		resp, httpRes, err := apiClient.DisksAPI.GetDisk(context.Background(), diskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DisksAPIService GetDiskSamples", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var diskId int64

		resp, httpRes, err := apiClient.DisksAPI.GetDiskSamples(context.Background(), diskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DisksAPIService ListDisks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DisksAPI.ListDisks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DisksAPIService RebuildDisk", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var diskId int64

		resp, httpRes, err := apiClient.DisksAPI.RebuildDisk(context.Background(), diskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DisksAPIService UpdateDisk", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var diskId int64

		resp, httpRes, err := apiClient.DisksAPI.UpdateDisk(context.Background(), diskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
