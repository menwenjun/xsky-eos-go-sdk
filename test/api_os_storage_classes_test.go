/*
XMS API

Testing OsStorageClassesAPIService

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

func Test_openapi_OsStorageClassesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OsStorageClassesAPIService CreateOSStorageClass", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OsStorageClassesAPI.CreateOSStorageClass(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsStorageClassesAPIService DeleteOSStorageClass", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var storageClassId int64

		resp, httpRes, err := apiClient.OsStorageClassesAPI.DeleteOSStorageClass(context.Background(), storageClassId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsStorageClassesAPIService GetOSStorageClass", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var storageClassId string

		resp, httpRes, err := apiClient.OsStorageClassesAPI.GetOSStorageClass(context.Background(), storageClassId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsStorageClassesAPIService ListOSStorageClasses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OsStorageClassesAPI.ListOSStorageClasses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsStorageClassesAPIService UpdateOSStorageClass", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var storageClassId int64

		resp, httpRes, err := apiClient.OsStorageClassesAPI.UpdateOSStorageClass(context.Background(), storageClassId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
