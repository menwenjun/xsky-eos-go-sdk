/*
XMS API

Testing DfsWormsAPIService

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

func Test_openapi_DfsWormsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DfsWormsAPIService CreateDfsWorm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DfsWormsAPI.CreateDfsWorm(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsWormsAPIService DeleteDfsWorm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dfsWormId int64

		resp, httpRes, err := apiClient.DfsWormsAPI.DeleteDfsWorm(context.Background(), dfsWormId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsWormsAPIService GetDfsWorm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dfsWormId int64

		resp, httpRes, err := apiClient.DfsWormsAPI.GetDfsWorm(context.Background(), dfsWormId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsWormsAPIService ListDfsWorms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DfsWormsAPI.ListDfsWorms(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsWormsAPIService UpdateDfsWorm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dfsWormId int64

		resp, httpRes, err := apiClient.DfsWormsAPI.UpdateDfsWorm(context.Background(), dfsWormId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
