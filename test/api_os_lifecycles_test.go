/*
XMS API

Testing OsLifecyclesAPIService

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

func Test_openapi_OsLifecyclesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OsLifecyclesAPIService GetLifecycle", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var lifecycleId int64

		resp, httpRes, err := apiClient.OsLifecyclesAPI.GetLifecycle(context.Background(), lifecycleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsLifecyclesAPIService ListLifecycles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OsLifecyclesAPI.ListLifecycles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
