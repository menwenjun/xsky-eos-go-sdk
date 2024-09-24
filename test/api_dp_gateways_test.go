/*
XMS API

Testing DpGatewaysAPIService

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

func Test_openapi_DpGatewaysAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DpGatewaysAPIService CreateDpGateway", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DpGatewaysAPI.CreateDpGateway(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DpGatewaysAPIService DeleteDpGateway", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var gatewayId int64

		resp, httpRes, err := apiClient.DpGatewaysAPI.DeleteDpGateway(context.Background(), gatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DpGatewaysAPIService GetDpGateway", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var gatewayId int64

		resp, httpRes, err := apiClient.DpGatewaysAPI.GetDpGateway(context.Background(), gatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DpGatewaysAPIService ListDpGateways", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DpGatewaysAPI.ListDpGateways(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DpGatewaysAPIService UpdateDpGateway", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var gatewayId int64

		resp, httpRes, err := apiClient.DpGatewaysAPI.UpdateDpGateway(context.Background(), gatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
