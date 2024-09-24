/*
XMS API

Testing OsZoneTranslogsAPIService

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

func Test_openapi_OsZoneTranslogsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OsZoneTranslogsAPIService GetOSZoneTranslog", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var translogUuid string

		resp, httpRes, err := apiClient.OsZoneTranslogsAPI.GetOSZoneTranslog(context.Background(), translogUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsZoneTranslogsAPIService ListOSZoneTranslogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OsZoneTranslogsAPI.ListOSZoneTranslogs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
