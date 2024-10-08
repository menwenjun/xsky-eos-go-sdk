/*
XMS API

Testing OspGatewaysAPIService

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

func Test_openapi_OspGatewaysAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OspGatewaysAPIService CreateOspGateway", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OspGatewaysAPI.CreateOspGateway(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OspGatewaysAPIService DeleteOspGateway", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var ospGatewayId string

		resp, httpRes, err := apiClient.OspGatewaysAPI.DeleteOspGateway(context.Background(), ospGatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OspGatewaysAPIService GetOspGateway", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var ospGatewayId int64

		resp, httpRes, err := apiClient.OspGatewaysAPI.GetOspGateway(context.Background(), ospGatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OspGatewaysAPIService GetOspGatewaySamples", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var ospGatewayId int64

		resp, httpRes, err := apiClient.OspGatewaysAPI.GetOspGatewaySamples(context.Background(), ospGatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OspGatewaysAPIService GetOspGatewaysStatSumByZone", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OspGatewaysAPI.GetOspGatewaysStatSumByZone(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OspGatewaysAPIService ListOspGateways", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OspGatewaysAPI.ListOspGateways(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OspGatewaysAPIService RestartOspGateway", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var ospGatewayId int64

		resp, httpRes, err := apiClient.OspGatewaysAPI.RestartOspGateway(context.Background(), ospGatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OspGatewaysAPIService RestartOspGateways", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.OspGatewaysAPI.RestartOspGateways(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OspGatewaysAPIService StartOspGateway", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var ospGatewayId int64

		resp, httpRes, err := apiClient.OspGatewaysAPI.StartOspGateway(context.Background(), ospGatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OspGatewaysAPIService StopOspGateway", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var ospGatewayId string

		resp, httpRes, err := apiClient.OspGatewaysAPI.StopOspGateway(context.Background(), ospGatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
