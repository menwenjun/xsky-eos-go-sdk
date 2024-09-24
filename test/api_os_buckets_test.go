/*
XMS API

Testing OsBucketsAPIService

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

func Test_openapi_OsBucketsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OsBucketsAPIService AddCustomLabels", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.AddCustomLabels(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService AddOSReplicationPaths", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.AddOSReplicationPaths(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService AddOSReplicationZones", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.AddOSReplicationZones(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService BatchGetObjectStorageSamples", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OsBucketsAPI.BatchGetObjectStorageSamples(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService CancelDeleteBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.CancelDeleteBucket(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService CreateBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OsBucketsAPI.CreateBucket(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService CreateObjectStorageBucketNFSClients", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.CreateObjectStorageBucketNFSClients(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService DeleteBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.DeleteBucket(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService DeleteObjectStorageBucketNFSClients", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.DeleteObjectStorageBucketNFSClients(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService DeleteObjectStorageLifecycle", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.DeleteObjectStorageLifecycle(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService GetBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.GetBucket(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService GetOSBucketOriginPullSamples", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.GetOSBucketOriginPullSamples(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService GetObjectStorageBucketNFSClient", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64
		var clientId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.GetObjectStorageBucketNFSClient(context.Background(), bucketId, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService GetObjectStorageBucketSamples", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.GetObjectStorageBucketSamples(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService ListBuckets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OsBucketsAPI.ListBuckets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService ListObjectStorageBucketNFSClients", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.ListObjectStorageBucketNFSClients(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService RemoveCustomLabels", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.RemoveCustomLabels(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService RemoveOSBucketLoggings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.RemoveOSBucketLoggings(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService RemoveOSReplicationPaths", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.RemoveOSReplicationPaths(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService RemoveOSReplicationZones", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.RemoveOSReplicationZones(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService SetAccessLogging", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.SetAccessLogging(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService SetMetadataSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.SetMetadataSearch(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService SetOSBucketPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.SetOSBucketPolicy(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService SetObjectStorageClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.SetObjectStorageClass(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService SetObjectStorageLifecycleRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.SetObjectStorageLifecycleRules(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService SuspendAccessLoggings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.SuspendAccessLoggings(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService SuspendOSReplicationPaths", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.SuspendOSReplicationPaths(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService SwitchOwnerOSZone", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.SwitchOwnerOSZone(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService UnsetAccessLogging", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.UnsetAccessLogging(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService UnsetOSBucketPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.UnsetOSBucketPolicy(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService UnsuspendAccessLogging", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.UnsuspendAccessLogging(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService UnsuspendOSReplicationPaths", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.UnsuspendOSReplicationPaths(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService UpdateBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.UpdateBucket(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService UpdateCustomLabels", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.UpdateCustomLabels(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OsBucketsAPIService UpdateObjectStorageBucketNFSClient", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId int64
		var clientId int64

		resp, httpRes, err := apiClient.OsBucketsAPI.UpdateObjectStorageBucketNFSClient(context.Background(), bucketId, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
