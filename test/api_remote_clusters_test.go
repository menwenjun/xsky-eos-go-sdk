/*
XMS API

Testing RemoteClustersAPIService

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

func Test_openapi_RemoteClustersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RemoteClustersAPIService CreateRemoteCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RemoteClustersAPI.CreateRemoteCluster(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemoteClustersAPIService DeleteRemoteCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int64

		resp, httpRes, err := apiClient.RemoteClustersAPI.DeleteRemoteCluster(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemoteClustersAPIService GetRemoteCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int64

		resp, httpRes, err := apiClient.RemoteClustersAPI.GetRemoteCluster(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemoteClustersAPIService ListRemoteClusters", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RemoteClustersAPI.ListRemoteClusters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemoteClustersAPIService UpdateRemoteCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int64

		resp, httpRes, err := apiClient.RemoteClustersAPI.UpdateRemoteCluster(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
