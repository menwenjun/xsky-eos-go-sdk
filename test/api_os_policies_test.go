/*
XMS API

Testing OsPoliciesAPIService

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

func Test_openapi_OsPoliciesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OsPoliciesAPIService CreatePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OsPoliciesAPI.CreatePolicy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsPoliciesAPIService DeletePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId int64

		resp, httpRes, err := apiClient.OsPoliciesAPI.DeletePolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsPoliciesAPIService GetPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId int64

		resp, httpRes, err := apiClient.OsPoliciesAPI.GetPolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsPoliciesAPIService ListPolicies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OsPoliciesAPI.ListPolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsPoliciesAPIService UpdatePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyId int64

		resp, httpRes, err := apiClient.OsPoliciesAPI.UpdatePolicy(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
