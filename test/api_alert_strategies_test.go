/*
XMS API

Testing AlertStrategiesAPIService

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

func Test_openapi_AlertStrategiesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AlertStrategiesAPIService CreateAlertStrategy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AlertStrategiesAPI.CreateAlertStrategy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertStrategiesAPIService DeleteAlertStrategy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var alertStrategyId int64

		httpRes, err := apiClient.AlertStrategiesAPI.DeleteAlertStrategy(context.Background(), alertStrategyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertStrategiesAPIService GetAlertStrategy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var alertStrategyId int64

		resp, httpRes, err := apiClient.AlertStrategiesAPI.GetAlertStrategy(context.Background(), alertStrategyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertStrategiesAPIService ListAlertStrategies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AlertStrategiesAPI.ListAlertStrategies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlertStrategiesAPIService UpdateUpdateAlertStrategyAlertContact", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var alertStrategyId int64

		httpRes, err := apiClient.AlertStrategiesAPI.UpdateUpdateAlertStrategyAlertContact(context.Background(), alertStrategyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
