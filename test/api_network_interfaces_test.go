/*
XMS API

Testing NetworkInterfacesAPIService

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

func Test_openapi_NetworkInterfacesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NetworkInterfacesAPIService GetNetworkInterface", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var networkInterfaceId int64

		resp, httpRes, err := apiClient.NetworkInterfacesAPI.GetNetworkInterface(context.Background(), networkInterfaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkInterfacesAPIService GetNetworkInterfaceSamples", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var networkInterfaceId int64

		resp, httpRes, err := apiClient.NetworkInterfacesAPI.GetNetworkInterfaceSamples(context.Background(), networkInterfaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkInterfacesAPIService ListNetworkInterfaces", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.NetworkInterfacesAPI.ListNetworkInterfaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
