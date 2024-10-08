/*
XMS API

Testing BlockConsistentSnapshotsAPIService

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

func Test_openapi_BlockConsistentSnapshotsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BlockConsistentSnapshotsAPIService CreateBlockConsistentSnapshot", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlockConsistentSnapshotsAPI.CreateBlockConsistentSnapshot(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConsistentSnapshotsAPIService DeleteBlockConsistentSnapshot", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var consistentSnapshotId int64

		resp, httpRes, err := apiClient.BlockConsistentSnapshotsAPI.DeleteBlockConsistentSnapshot(context.Background(), consistentSnapshotId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConsistentSnapshotsAPIService GetBlockConsistentSnapshot", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var blockConsistentSnapshotId int64

		resp, httpRes, err := apiClient.BlockConsistentSnapshotsAPI.GetBlockConsistentSnapshot(context.Background(), blockConsistentSnapshotId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConsistentSnapshotsAPIService ListBlockConsistentSnapshots", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlockConsistentSnapshotsAPI.ListBlockConsistentSnapshots(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
