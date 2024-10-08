/*
XMS API

Testing DpBlockReplicationPoliciesAPIService

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

func Test_openapi_DpBlockReplicationPoliciesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DpBlockReplicationPoliciesAPIService CreateDpBlockReplicationPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DpBlockReplicationPoliciesAPI.CreateDpBlockReplicationPolicy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DpBlockReplicationPoliciesAPIService DeleteDpBlockReplicationPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId int64

		httpRes, err := apiClient.DpBlockReplicationPoliciesAPI.DeleteDpBlockReplicationPolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DpBlockReplicationPoliciesAPIService GetDpBlockReplicationPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId int64

		resp, httpRes, err := apiClient.DpBlockReplicationPoliciesAPI.GetDpBlockReplicationPolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DpBlockReplicationPoliciesAPIService ListDpBlockReplicationPolicies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DpBlockReplicationPoliciesAPI.ListDpBlockReplicationPolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DpBlockReplicationPoliciesAPIService UpdateDpBlockReplicationPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId int64

		resp, httpRes, err := apiClient.DpBlockReplicationPoliciesAPI.UpdateDpBlockReplicationPolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
