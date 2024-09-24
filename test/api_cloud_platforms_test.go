/*
XMS API

Testing CloudPlatformsAPIService

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

func Test_openapi_CloudPlatformsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CloudPlatformsAPIService CreateCloudPlatform", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CloudPlatformsAPI.CreateCloudPlatform(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudPlatformsAPIService DeleteCloudPlatform", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudPlatformId int64

		httpRes, err := apiClient.CloudPlatformsAPI.DeleteCloudPlatform(context.Background(), cloudPlatformId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudPlatformsAPIService GetCloudPlatform", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudPlatformId int64

		resp, httpRes, err := apiClient.CloudPlatformsAPI.GetCloudPlatform(context.Background(), cloudPlatformId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudPlatformsAPIService ListCloudPlatforms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CloudPlatformsAPI.ListCloudPlatforms(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudPlatformsAPIService SyncCloudPlatform", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudPlatformId int64

		resp, httpRes, err := apiClient.CloudPlatformsAPI.SyncCloudPlatform(context.Background(), cloudPlatformId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudPlatformsAPIService UpdateCloudPlatform", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudPlatformId int64

		resp, httpRes, err := apiClient.CloudPlatformsAPI.UpdateCloudPlatform(context.Background(), cloudPlatformId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
