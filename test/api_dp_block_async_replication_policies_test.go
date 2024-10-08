/*
XMS API

Testing DpBlockAsyncReplicationPoliciesAPIService

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

func Test_openapi_DpBlockAsyncReplicationPoliciesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DpBlockAsyncReplicationPoliciesAPIService CreateDpBlockAsyncReplicationPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DpBlockAsyncReplicationPoliciesAPI.CreateDpBlockAsyncReplicationPolicy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DpBlockAsyncReplicationPoliciesAPIService DeleteDpBlockAsyncReplicationPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId int64

		resp, httpRes, err := apiClient.DpBlockAsyncReplicationPoliciesAPI.DeleteDpBlockAsyncReplicationPolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DpBlockAsyncReplicationPoliciesAPIService GetDpBlockAsyncReplicationPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId int64

		resp, httpRes, err := apiClient.DpBlockAsyncReplicationPoliciesAPI.GetDpBlockAsyncReplicationPolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DpBlockAsyncReplicationPoliciesAPIService ListDpBlockAsyncReplicationPolicies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DpBlockAsyncReplicationPoliciesAPI.ListDpBlockAsyncReplicationPolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DpBlockAsyncReplicationPoliciesAPIService UpdateDpBlockAsyncReplicationPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId int64

		resp, httpRes, err := apiClient.DpBlockAsyncReplicationPoliciesAPI.UpdateDpBlockAsyncReplicationPolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
