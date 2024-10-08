/*
XMS API

Testing DfsGatewayServicesAPIService

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

func Test_openapi_DfsGatewayServicesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DfsGatewayServicesAPIService BatchGetDfsGatewayServiceSamples", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DfsGatewayServicesAPI.BatchGetDfsGatewayServiceSamples(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsGatewayServicesAPIService GetDfsGatewayService", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsGatewayServiceId int64

		resp, httpRes, err := apiClient.DfsGatewayServicesAPI.GetDfsGatewayService(context.Background(), dfsGatewayServiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsGatewayServicesAPIService GetDfsGatewayServiceSamples", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsGatewayServiceId int64

		resp, httpRes, err := apiClient.DfsGatewayServicesAPI.GetDfsGatewayServiceSamples(context.Background(), dfsGatewayServiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsGatewayServicesAPIService ListDfsGatewayServices", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DfsGatewayServicesAPI.ListDfsGatewayServices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
