/*
XMS API

Testing EmailGroupsAPIService

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

func Test_openapi_EmailGroupsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EmailGroupsAPIService CreateEmailGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EmailGroupsAPI.CreateEmailGroup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailGroupsAPIService DeleteEmailGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId int64

		httpRes, err := apiClient.EmailGroupsAPI.DeleteEmailGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailGroupsAPIService GetEmailGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId int64

		resp, httpRes, err := apiClient.EmailGroupsAPI.GetEmailGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailGroupsAPIService ListEmailGroups", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EmailGroupsAPI.ListEmailGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailGroupsAPIService UpdateEmailGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId int64

		resp, httpRes, err := apiClient.EmailGroupsAPI.UpdateEmailGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
