/*
XMS API

Testing FsUsersAPIService

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

func Test_openapi_FsUsersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FsUsersAPIService CreateFSUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FsUsersAPI.CreateFSUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FsUsersAPIService DeleteFSUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fsUserId int64

		httpRes, err := apiClient.FsUsersAPI.DeleteFSUser(context.Background(), fsUserId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FsUsersAPIService GetFSUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fsUserId int64

		resp, httpRes, err := apiClient.FsUsersAPI.GetFSUser(context.Background(), fsUserId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FsUsersAPIService ListFSUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FsUsersAPI.ListFSUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FsUsersAPIService UpdateFSUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fsUserId int64

		resp, httpRes, err := apiClient.FsUsersAPI.UpdateFSUser(context.Background(), fsUserId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FsUsersAPIService UpdateFSUserSecondaryGroups", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fsUserId int64

		resp, httpRes, err := apiClient.FsUsersAPI.UpdateFSUserSecondaryGroups(context.Background(), fsUserId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FsUsersAPIService VerifyFSUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FsUsersAPI.VerifyFSUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
