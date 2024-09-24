/*
XMS API

Testing BlockSnapshotsAPIService

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

func Test_openapi_BlockSnapshotsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BlockSnapshotsAPIService CreateBlockSnapshot", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlockSnapshotsAPI.CreateBlockSnapshot(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockSnapshotsAPIService DeleteBlockSnapshot", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var snapshotId int64

		resp, httpRes, err := apiClient.BlockSnapshotsAPI.DeleteBlockSnapshot(context.Background(), snapshotId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockSnapshotsAPIService GetBlockSnapshot", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var blockSnapshotId int64

		resp, httpRes, err := apiClient.BlockSnapshotsAPI.GetBlockSnapshot(context.Background(), blockSnapshotId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockSnapshotsAPIService ListBlockSnapshots", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlockSnapshotsAPI.ListBlockSnapshots(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockSnapshotsAPIService UpdateSnapshot", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var snapshotId int64

		resp, httpRes, err := apiClient.BlockSnapshotsAPI.UpdateSnapshot(context.Background(), snapshotId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
