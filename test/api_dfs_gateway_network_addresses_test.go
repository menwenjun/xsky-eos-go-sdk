/*
XMS API

Testing DfsGatewayNetworkAddressesAPIService

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

func Test_openapi_DfsGatewayNetworkAddressesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DfsGatewayNetworkAddressesAPIService GetDfsGatewayNetworkAddress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dfsGatewayNetworkAddressId int64

		resp, httpRes, err := apiClient.DfsGatewayNetworkAddressesAPI.GetDfsGatewayNetworkAddress(context.Background(), dfsGatewayNetworkAddressId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsGatewayNetworkAddressesAPIService ListDfsGatewayNetworkAddresses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DfsGatewayNetworkAddressesAPI.ListDfsGatewayNetworkAddresses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
