/*
XMS API

Testing PartitionsAPIService

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

func Test_openapi_PartitionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PartitionsAPIService GetPartition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var partitionId string

		resp, httpRes, err := apiClient.PartitionsAPI.GetPartition(context.Background(), partitionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartitionsAPIService GetPartitionPredictions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var partitionId int64

		resp, httpRes, err := apiClient.PartitionsAPI.GetPartitionPredictions(context.Background(), partitionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartitionsAPIService ListPartitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PartitionsAPI.ListPartitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PartitionsAPIService UpdatePartitionOspPartitionType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var partitionId int64

		resp, httpRes, err := apiClient.PartitionsAPI.UpdatePartitionOspPartitionType(context.Background(), partitionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
