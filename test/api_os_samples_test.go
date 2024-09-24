/*
XMS API

Testing OsSamplesAPIService

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

func Test_openapi_OsSamplesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OsSamplesAPIService GetOSSamples", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OsSamplesAPI.GetOSSamples(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsSamplesAPIService GetOSSamplesByBucketName", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userName string
		var bucketName string

		resp, httpRes, err := apiClient.OsSamplesAPI.GetOSSamplesByBucketName(context.Background(), userName, bucketName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsSamplesAPIService GetOSSamplesByUserName", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userName string

		resp, httpRes, err := apiClient.OsSamplesAPI.GetOSSamplesByUserName(context.Background(), userName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
