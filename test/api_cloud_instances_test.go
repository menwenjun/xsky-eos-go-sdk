/*
XMS API

Testing CloudInstancesAPIService

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

func Test_openapi_CloudInstancesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CloudInstancesAPIService GetCloudInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudInstanceId int64

		resp, httpRes, err := apiClient.CloudInstancesAPI.GetCloudInstance(context.Background(), cloudInstanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudInstancesAPIService GetCloudInstanceSamples", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudInstanceId int64

		resp, httpRes, err := apiClient.CloudInstancesAPI.GetCloudInstanceSamples(context.Background(), cloudInstanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudInstancesAPIService ListCloudInstances", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CloudInstancesAPI.ListCloudInstances(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
