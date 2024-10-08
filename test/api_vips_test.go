/*
XMS API

Testing VipsAPIService

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

func Test_openapi_VipsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VipsAPIService CreateVIP", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.VipsAPI.CreateVIP(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VipsAPIService DeleteVIP", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vipId int64

		resp, httpRes, err := apiClient.VipsAPI.DeleteVIP(context.Background(), vipId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VipsAPIService GetVIP", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vipId int64

		resp, httpRes, err := apiClient.VipsAPI.GetVIP(context.Background(), vipId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VipsAPIService ListVIPs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.VipsAPI.ListVIPs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VipsAPIService UpdateVIP", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vipId int64

		resp, httpRes, err := apiClient.VipsAPI.UpdateVIP(context.Background(), vipId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
