/*
XMS API

Testing RbdClientsAPIService

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

func Test_openapi_RbdClientsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RbdClientsAPIService CreateRBDC", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RbdClientsAPI.CreateRBDC(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RbdClientsAPIService DeleteRBDClient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var rbdcId int64

		resp, httpRes, err := apiClient.RbdClientsAPI.DeleteRBDClient(context.Background(), rbdcId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RbdClientsAPIService GetRBDClient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var rbdcId int64

		resp, httpRes, err := apiClient.RbdClientsAPI.GetRBDClient(context.Background(), rbdcId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RbdClientsAPIService ListRBDClients", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RbdClientsAPI.ListRBDClients(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RbdClientsAPIService UpdateRBDClient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var rbdcId int64

		resp, httpRes, err := apiClient.RbdClientsAPI.UpdateRBDClient(context.Background(), rbdcId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RbdClientsAPIService ValidateRBDClientHost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RbdClientsAPI.ValidateRBDClientHost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
