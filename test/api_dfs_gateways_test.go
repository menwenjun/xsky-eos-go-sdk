/*
XMS API

Testing DfsGatewaysAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_DfsGatewaysAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DfsGatewaysAPIService GetDfsGateway", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dfsGatewayId int64

		resp, httpRes, err := apiClient.DfsGatewaysAPI.GetDfsGateway(context.Background(), dfsGatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsGatewaysAPIService GetDfsGatewaySamples", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dfsGatewayId int64

		resp, httpRes, err := apiClient.DfsGatewaysAPI.GetDfsGatewaySamples(context.Background(), dfsGatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsGatewaysAPIService ListDfsGatewayConnections", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dfsGatewayId int64

		resp, httpRes, err := apiClient.DfsGatewaysAPI.ListDfsGatewayConnections(context.Background(), dfsGatewayId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfsGatewaysAPIService ListDfsGateways", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DfsGatewaysAPI.ListDfsGateways(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
