/*
XMS API

Testing SearchCapabilitiesAPIService

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

func Test_openapi_SearchCapabilitiesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SearchCapabilitiesAPIService SearchCapabilites", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SearchCapabilitiesAPI.SearchCapabilites(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
