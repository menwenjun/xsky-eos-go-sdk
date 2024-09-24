/*
XMS API

Testing HostMaintainValidatorsAPIService

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

func Test_openapi_HostMaintainValidatorsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HostMaintainValidatorsAPIService CreateHostMaintainValidator", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.HostMaintainValidatorsAPI.CreateHostMaintainValidator(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HostMaintainValidatorsAPIService GetHostMaintainValidator", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hostMaintainValidatorId int64

		resp, httpRes, err := apiClient.HostMaintainValidatorsAPI.GetHostMaintainValidator(context.Background(), hostMaintainValidatorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
