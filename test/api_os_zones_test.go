/*
XMS API

Testing OsZonesAPIService

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

func Test_openapi_OsZonesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OsZonesAPIService CreateObjectStorageZone", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OsZonesAPI.CreateObjectStorageZone(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsZonesAPIService DeleteObjectStorageZone", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var zoneUuid string

		resp, httpRes, err := apiClient.OsZonesAPI.DeleteObjectStorageZone(context.Background(), zoneUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsZonesAPIService GetObjectStorageZone", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var zoneUuid string

		resp, httpRes, err := apiClient.OsZonesAPI.GetObjectStorageZone(context.Background(), zoneUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsZonesAPIService GetObjectStorageZoneSamples", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var zoneUuid string

		resp, httpRes, err := apiClient.OsZonesAPI.GetObjectStorageZoneSamples(context.Background(), zoneUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsZonesAPIService ListObjectStorageZones", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OsZonesAPI.ListObjectStorageZones(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsZonesAPIService UpdateOSZonesClockDiff", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var zoneUuid string

		resp, httpRes, err := apiClient.OsZonesAPI.UpdateOSZonesClockDiff(context.Background(), zoneUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
