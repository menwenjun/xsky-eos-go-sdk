/*
XMS API

Testing DfsStoragePoliciesAPIService

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

func Test_openapi_DfsStoragePoliciesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DfsStoragePoliciesAPIService CreateDfsStoragePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DfsStoragePoliciesAPI.CreateDfsStoragePolicy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsStoragePoliciesAPIService DeleteDfsStoragePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsStoragePolicyId int64

		resp, httpRes, err := apiClient.DfsStoragePoliciesAPI.DeleteDfsStoragePolicy(context.Background(), dfsStoragePolicyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsStoragePoliciesAPIService GetDfsStoragePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsStoragePolicyId int64

		resp, httpRes, err := apiClient.DfsStoragePoliciesAPI.GetDfsStoragePolicy(context.Background(), dfsStoragePolicyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsStoragePoliciesAPIService ListDfsStoragePolicies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DfsStoragePoliciesAPI.ListDfsStoragePolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsStoragePoliciesAPIService UnlinkDfsStoragePolicyAndDfsPath", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsStoragePolicyId int64

		resp, httpRes, err := apiClient.DfsStoragePoliciesAPI.UnlinkDfsStoragePolicyAndDfsPath(context.Background(), dfsStoragePolicyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsStoragePoliciesAPIService UpdateDfsStoragePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsStoragePolicyId int64

		resp, httpRes, err := apiClient.DfsStoragePoliciesAPI.UpdateDfsStoragePolicy(context.Background(), dfsStoragePolicyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
