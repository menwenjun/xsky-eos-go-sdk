/*
XMS API

Testing DfsGatewayZonesAPIService

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

func Test_openapi_DfsGatewayZonesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DfsGatewayZonesAPIService CreateDfsGatewayZone", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DfsGatewayZonesAPI.CreateDfsGatewayZone(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsGatewayZonesAPIService DeleteDfsGatewayZone", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsGatewayZoneId int64

		resp, httpRes, err := apiClient.DfsGatewayZonesAPI.DeleteDfsGatewayZone(context.Background(), dfsGatewayZoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsGatewayZonesAPIService GetDfsGatewayZone", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsGatewayZoneId int64

		resp, httpRes, err := apiClient.DfsGatewayZonesAPI.GetDfsGatewayZone(context.Background(), dfsGatewayZoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsGatewayZonesAPIService GetDfsGatewayZoneSamples", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsGatewayZoneId int64

		resp, httpRes, err := apiClient.DfsGatewayZonesAPI.GetDfsGatewayZoneSamples(context.Background(), dfsGatewayZoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsGatewayZonesAPIService ListDfsGatewayZones", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DfsGatewayZonesAPI.ListDfsGatewayZones(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsGatewayZonesAPIService UpdateDfsGatewayZone", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsGatewayZoneId int64

		resp, httpRes, err := apiClient.DfsGatewayZonesAPI.UpdateDfsGatewayZone(context.Background(), dfsGatewayZoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
