/*
XMS API

Testing ClusterServicesAPIService

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

func Test_openapi_ClusterServicesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ClusterServicesAPIService GetClusterService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId int64

		resp, httpRes, err := apiClient.ClusterServicesAPI.GetClusterService(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterServicesAPIService ListClusterServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ClusterServicesAPI.ListClusterServices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterServicesAPIService RebuildClusterService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId int64

		resp, httpRes, err := apiClient.ClusterServicesAPI.RebuildClusterService(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterServicesAPIService ReindexDocs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId int64

		resp, httpRes, err := apiClient.ClusterServicesAPI.ReindexDocs(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
