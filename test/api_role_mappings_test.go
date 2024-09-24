/*
XMS API

Testing RoleMappingsAPIService

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

func Test_openapi_RoleMappingsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RoleMappingsAPIService CreateRoleMapping", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RoleMappingsAPI.CreateRoleMapping(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleMappingsAPIService DeleteRoleMapping", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var roleMappingId int64

		httpRes, err := apiClient.RoleMappingsAPI.DeleteRoleMapping(context.Background(), roleMappingId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleMappingsAPIService GetRoleMapping", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var roleMappingId int64

		resp, httpRes, err := apiClient.RoleMappingsAPI.GetRoleMapping(context.Background(), roleMappingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleMappingsAPIService ListRoleMappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RoleMappingsAPI.ListRoleMappings(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleMappingsAPIService UpdateRoleMapping", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var roleMappingId int64

		resp, httpRes, err := apiClient.RoleMappingsAPI.UpdateRoleMapping(context.Background(), roleMappingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
