/*
XMS API

Testing ProtectionDomainsAPIService

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

func Test_openapi_ProtectionDomainsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProtectionDomainsAPIService GetProtectionDomain", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var protectionDomainId int64

		resp, httpRes, err := apiClient.ProtectionDomainsAPI.GetProtectionDomain(context.Background(), protectionDomainId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtectionDomainsAPIService ListProtectionDomains", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProtectionDomainsAPI.ListProtectionDomains(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
