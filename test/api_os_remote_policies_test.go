/*
XMS API

Testing OsRemotePoliciesAPIService

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

func Test_openapi_OsRemotePoliciesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OsRemotePoliciesAPIService GetOSRemotePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyUuid string

		resp, httpRes, err := apiClient.OsRemotePoliciesAPI.GetOSRemotePolicy(context.Background(), policyUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsRemotePoliciesAPIService ListOSRemotePolicies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OsRemotePoliciesAPI.ListOSRemotePolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
