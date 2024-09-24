/*
XMS API

Testing AlertGroupsAPIService

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

func Test_openapi_AlertGroupsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AlertGroupsAPIService CreateAlertGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AlertGroupsAPI.CreateAlertGroup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertGroupsAPIService DeleteAlertGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId int64

		httpRes, err := apiClient.AlertGroupsAPI.DeleteAlertGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertGroupsAPIService GetAlertGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId int64

		resp, httpRes, err := apiClient.AlertGroupsAPI.GetAlertGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertGroupsAPIService ListAlertGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AlertGroupsAPI.ListAlertGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertGroupsAPIService UpdateAlertGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId int64

		resp, httpRes, err := apiClient.AlertGroupsAPI.UpdateAlertGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
