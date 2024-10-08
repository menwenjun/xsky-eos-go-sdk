/*
XMS API

Testing OsGatewaysAPIService

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

func Test_openapi_OsGatewaysAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OsGatewaysAPIService CreateGateway", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OsGatewaysAPI.CreateGateway(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsGatewaysAPIService DeleteGateway", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var gatewayId int64

		resp, httpRes, err := apiClient.OsGatewaysAPI.DeleteGateway(context.Background(), gatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsGatewaysAPIService GetGateway", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var gatewayId int64

		resp, httpRes, err := apiClient.OsGatewaysAPI.GetGateway(context.Background(), gatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsGatewaysAPIService GetGatewaySamples", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var gatewayId int64

		resp, httpRes, err := apiClient.OsGatewaysAPI.GetGatewaySamples(context.Background(), gatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsGatewaysAPIService ListGateways", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OsGatewaysAPI.ListGateways(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsGatewaysAPIService UpdateGateway", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var gatewayId int64

		resp, httpRes, err := apiClient.OsGatewaysAPI.UpdateGateway(context.Background(), gatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
