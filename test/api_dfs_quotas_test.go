/*
XMS API

Testing DfsQuotasAPIService

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

func Test_openapi_DfsQuotasAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DfsQuotasAPIService CreateDfsQuota", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DfsQuotasAPI.CreateDfsQuota(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsQuotasAPIService DeleteDfsQuota", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dfsQuotaId int64

		resp, httpRes, err := apiClient.DfsQuotasAPI.DeleteDfsQuota(context.Background(), dfsQuotaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsQuotasAPIService DfsQuotaOverview", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DfsQuotasAPI.DfsQuotaOverview(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsQuotasAPIService GetDfsQuota", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dfsQuotaId int64

		resp, httpRes, err := apiClient.DfsQuotasAPI.GetDfsQuota(context.Background(), dfsQuotaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsQuotasAPIService GetDfsQuotaPredictions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dfsQuotaId int64

		resp, httpRes, err := apiClient.DfsQuotasAPI.GetDfsQuotaPredictions(context.Background(), dfsQuotaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsQuotasAPIService GetDfsQuotaSamples", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dfsQuotaId int64

		resp, httpRes, err := apiClient.DfsQuotasAPI.GetDfsQuotaSamples(context.Background(), dfsQuotaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsQuotasAPIService ListDfsQuotas", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DfsQuotasAPI.ListDfsQuotas(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsQuotasAPIService PathValidator", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DfsQuotasAPI.PathValidator(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsQuotasAPIService UpdateDfsQuota", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dfsQuotaId int64

		resp, httpRes, err := apiClient.DfsQuotasAPI.UpdateDfsQuota(context.Background(), dfsQuotaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
