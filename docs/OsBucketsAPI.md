# \OsBucketsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCustomLabels**](OsBucketsAPI.md#AddCustomLabels) | **Post** /os-buckets/{bucket_id}:add-custom-labels | 
[**AddOSReplicationPaths**](OsBucketsAPI.md#AddOSReplicationPaths) | **Post** /os-buckets/{bucket_id}:add-os-replication-paths | 
[**AddOSReplicationZones**](OsBucketsAPI.md#AddOSReplicationZones) | **Post** /os-buckets/{bucket_id}:add-os-replication-zones | 
[**BatchGetObjectStorageSamples**](OsBucketsAPI.md#BatchGetObjectStorageSamples) | **Get** /os-buckets/samples | 
[**CancelDeleteBucket**](OsBucketsAPI.md#CancelDeleteBucket) | **Post** /os-buckets/{bucket_id}:cancel | 
[**CreateBucket**](OsBucketsAPI.md#CreateBucket) | **Post** /os-buckets/ | 
[**CreateObjectStorageBucketNFSClients**](OsBucketsAPI.md#CreateObjectStorageBucketNFSClients) | **Post** /os-buckets/{bucket_id}/nfs-clients | 
[**DeleteBucket**](OsBucketsAPI.md#DeleteBucket) | **Delete** /os-buckets/{bucket_id} | 
[**DeleteObjectStorageBucketNFSClients**](OsBucketsAPI.md#DeleteObjectStorageBucketNFSClients) | **Delete** /os-buckets/{bucket_id}/nfs-clients | 
[**DeleteObjectStorageLifecycle**](OsBucketsAPI.md#DeleteObjectStorageLifecycle) | **Delete** /os-buckets/{bucket_id}/lifecycle | 
[**GetBucket**](OsBucketsAPI.md#GetBucket) | **Get** /os-buckets/{bucket_id} | 
[**GetOSBucketOriginPullSamples**](OsBucketsAPI.md#GetOSBucketOriginPullSamples) | **Get** /os-buckets/{bucket_id}/origin_pull_samples | 
[**GetObjectStorageBucketNFSClient**](OsBucketsAPI.md#GetObjectStorageBucketNFSClient) | **Get** /os-buckets/{bucket_id}/nfs-clients/{client_id} | 
[**GetObjectStorageBucketSamples**](OsBucketsAPI.md#GetObjectStorageBucketSamples) | **Get** /os-buckets/{bucket_id}/samples | 
[**ListBuckets**](OsBucketsAPI.md#ListBuckets) | **Get** /os-buckets/ | 
[**ListObjectStorageBucketNFSClients**](OsBucketsAPI.md#ListObjectStorageBucketNFSClients) | **Get** /os-buckets/{bucket_id}/nfs-clients | 
[**RemoveCustomLabels**](OsBucketsAPI.md#RemoveCustomLabels) | **Post** /os-buckets/{bucket_id}:remove-custom-labels | 
[**RemoveOSBucketLoggings**](OsBucketsAPI.md#RemoveOSBucketLoggings) | **Post** /os-buckets/{bucket_id}:remove-os-bucket-loggings | 
[**RemoveOSReplicationPaths**](OsBucketsAPI.md#RemoveOSReplicationPaths) | **Post** /os-buckets/{bucket_id}:remove-os-replication-paths | 
[**RemoveOSReplicationZones**](OsBucketsAPI.md#RemoveOSReplicationZones) | **Post** /os-buckets/{bucket_id}:remove-os-replication-zones | 
[**SetAccessLogging**](OsBucketsAPI.md#SetAccessLogging) | **Post** /os-buckets/{bucket_id}:set-access-logging | 
[**SetMetadataSearch**](OsBucketsAPI.md#SetMetadataSearch) | **Post** /os-buckets/{bucket_id}:set-metadata-search | 
[**SetOSBucketPolicy**](OsBucketsAPI.md#SetOSBucketPolicy) | **Post** /os-buckets/{bucket_id}:set-bucket-policy | 
[**SetObjectStorageClass**](OsBucketsAPI.md#SetObjectStorageClass) | **Post** /os-buckets/{bucket_id}:set-object-storage-class | 
[**SetObjectStorageLifecycleRules**](OsBucketsAPI.md#SetObjectStorageLifecycleRules) | **Put** /os-buckets/{bucket_id}/lifecycle | 
[**SuspendAccessLoggings**](OsBucketsAPI.md#SuspendAccessLoggings) | **Post** /os-buckets/{bucket_id}:suspend-access-logging | 
[**SuspendOSReplicationPaths**](OsBucketsAPI.md#SuspendOSReplicationPaths) | **Post** /os-buckets/{bucket_id}:suspend-os-replication-paths | 
[**SwitchOwnerOSZone**](OsBucketsAPI.md#SwitchOwnerOSZone) | **Post** /os-buckets/{bucket_id}:switch-owner-os-zone | 
[**UnsetAccessLogging**](OsBucketsAPI.md#UnsetAccessLogging) | **Post** /os-buckets/{bucket_id}:unset-access-logging | 
[**UnsetOSBucketPolicy**](OsBucketsAPI.md#UnsetOSBucketPolicy) | **Post** /os-buckets/{bucket_id}:unset-bucket-policy | 
[**UnsuspendAccessLogging**](OsBucketsAPI.md#UnsuspendAccessLogging) | **Post** /os-buckets/{bucket_id}:unsuspend-access-logging | 
[**UnsuspendOSReplicationPaths**](OsBucketsAPI.md#UnsuspendOSReplicationPaths) | **Post** /os-buckets/{bucket_id}:unsuspend-os-replication-paths | 
[**UpdateBucket**](OsBucketsAPI.md#UpdateBucket) | **Patch** /os-buckets/{bucket_id} | 
[**UpdateCustomLabels**](OsBucketsAPI.md#UpdateCustomLabels) | **Post** /os-buckets/{bucket_id}:update-custom-labels | 
[**UpdateObjectStorageBucketNFSClient**](OsBucketsAPI.md#UpdateObjectStorageBucketNFSClient) | **Patch** /os-buckets/{bucket_id}/nfs-clients/{client_id} | 



## AddCustomLabels

> ObjectStorageBucketResp AddCustomLabels(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewOSBucketCustomLabelsAddReq(*openapiclient.NewOSBucketCustomLabelsAddReqBucket([]openapiclient.OSBucketCustomLabelsAddReqBucketLabelsElt{*openapiclient.NewOSBucketCustomLabelsAddReqBucketLabelsElt("Category_example", "Name_example", "Type_example")})) // OSBucketCustomLabelsAddReq | bucket custom labels info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.AddCustomLabels(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.AddCustomLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddCustomLabels`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.AddCustomLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCustomLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBucketCustomLabelsAddReq**](OSBucketCustomLabelsAddReq.md) | bucket custom labels info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddOSReplicationPaths

> ObjectStorageBucketResp AddOSReplicationPaths(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewOSBucketAddReplicationPathsReq(*openapiclient.NewOSBucketAddReplicationPathsReqBucket()) // OSBucketAddReplicationPathsReq | replication paths info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.AddOSReplicationPaths(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.AddOSReplicationPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOSReplicationPaths`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.AddOSReplicationPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOSReplicationPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBucketAddReplicationPathsReq**](OSBucketAddReplicationPathsReq.md) | replication paths info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddOSReplicationZones

> ObjectStorageBucketResp AddOSReplicationZones(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewOSBucketAddReplicationZonesReq(*openapiclient.NewOSBucketAddReplicationZonesReqBucket()) // OSBucketAddReplicationZonesReq | replication zones info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.AddOSReplicationZones(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.AddOSReplicationZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOSReplicationZones`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.AddOSReplicationZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOSReplicationZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBucketAddReplicationZonesReq**](OSBucketAddReplicationZonesReq.md) | replication zones info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchGetObjectStorageSamples

> MultiObjectStorageBucketsSamplesResp BatchGetObjectStorageSamples(ctx).Ids(ids).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	ids := "ids_example" // string | bucket ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.BatchGetObjectStorageSamples(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.BatchGetObjectStorageSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchGetObjectStorageSamples`: MultiObjectStorageBucketsSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.BatchGetObjectStorageSamples`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchGetObjectStorageSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | bucket ids | 

### Return type

[**MultiObjectStorageBucketsSamplesResp**](MultiObjectStorageBucketsSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelDeleteBucket

> ObjectStorageBucketResp CancelDeleteBucket(ctx, bucketId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.CancelDeleteBucket(context.Background(), bucketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.CancelDeleteBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelDeleteBucket`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.CancelDeleteBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelDeleteBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBucket

> ObjectStorageBucketResp CreateBucket(ctx).Body(body).ClusterId(clusterId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := *openapiclient.NewObjectStorageBucketCreateReq() // ObjectStorageBucketCreateReq | bucket info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.CreateBucket(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.CreateBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBucket`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.CreateBucket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ObjectStorageBucketCreateReq**](ObjectStorageBucketCreateReq.md) | bucket info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateObjectStorageBucketNFSClients

> RawObjectStorageBucketResp CreateObjectStorageBucketNFSClients(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewObjectStorageBucketNFSClientsCreateReq([]openapiclient.ObjectStorageBucketNFSClientsCreateReqClientsElt{*openapiclient.NewObjectStorageBucketNFSClientsCreateReqClientsElt("Client_example", "Permission_example")}) // ObjectStorageBucketNFSClientsCreateReq | nfs client info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.CreateObjectStorageBucketNFSClients(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.CreateObjectStorageBucketNFSClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateObjectStorageBucketNFSClients`: RawObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.CreateObjectStorageBucketNFSClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectStorageBucketNFSClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ObjectStorageBucketNFSClientsCreateReq**](ObjectStorageBucketNFSClientsCreateReq.md) | nfs client info | 

### Return type

[**RawObjectStorageBucketResp**](RawObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBucket

> ObjectStorageBucketResp DeleteBucket(ctx, bucketId).Force(force).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	force := true // bool | force delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.DeleteBucket(context.Background(), bucketId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.DeleteBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBucket`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.DeleteBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObjectStorageBucketNFSClients

> RawObjectStorageBucketResp DeleteObjectStorageBucketNFSClients(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewObjectStorageBucketNFSClientsDeleteReq([]int64{int64(123)}) // ObjectStorageBucketNFSClientsDeleteReq | nfs client info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.DeleteObjectStorageBucketNFSClients(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.DeleteObjectStorageBucketNFSClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteObjectStorageBucketNFSClients`: RawObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.DeleteObjectStorageBucketNFSClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectStorageBucketNFSClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ObjectStorageBucketNFSClientsDeleteReq**](ObjectStorageBucketNFSClientsDeleteReq.md) | nfs client info | 

### Return type

[**RawObjectStorageBucketResp**](RawObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObjectStorageLifecycle

> ObjectStorageBucketResp DeleteObjectStorageLifecycle(ctx, bucketId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.DeleteObjectStorageLifecycle(context.Background(), bucketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.DeleteObjectStorageLifecycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteObjectStorageLifecycle`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.DeleteObjectStorageLifecycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectStorageLifecycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBucket

> ObjectStorageBucketResp GetBucket(ctx, bucketId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.GetBucket(context.Background(), bucketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.GetBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBucket`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.GetBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSBucketOriginPullSamples

> OSBucketOriginPullSamplesResp GetOSBucketOriginPullSamples(ctx, bucketId).OriginMode(originMode).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	originMode := "originMode_example" // string | origin mode (optional)
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.GetOSBucketOriginPullSamples(context.Background(), bucketId).OriginMode(originMode).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.GetOSBucketOriginPullSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSBucketOriginPullSamples`: OSBucketOriginPullSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.GetOSBucketOriginPullSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSBucketOriginPullSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **originMode** | **string** | origin mode | 
 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**OSBucketOriginPullSamplesResp**](OSBucketOriginPullSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageBucketNFSClient

> ObjectStorageBucketNFSClientResp GetObjectStorageBucketNFSClient(ctx, bucketId, clientId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	clientId := int64(789) // int64 | nfs client id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.GetObjectStorageBucketNFSClient(context.Background(), bucketId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.GetObjectStorageBucketNFSClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageBucketNFSClient`: ObjectStorageBucketNFSClientResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.GetObjectStorageBucketNFSClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 
**clientId** | **int64** | nfs client id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageBucketNFSClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ObjectStorageBucketNFSClientResp**](ObjectStorageBucketNFSClientResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageBucketSamples

> ObjectStorageBucketSamplesResp GetObjectStorageBucketSamples(ctx, bucketId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.GetObjectStorageBucketSamples(context.Background(), bucketId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.GetObjectStorageBucketSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageBucketSamples`: ObjectStorageBucketSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.GetObjectStorageBucketSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageBucketSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**ObjectStorageBucketSamplesResp**](ObjectStorageBucketSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuckets

> ObjectStorageBucketsResp ListBuckets(ctx).ClusterId(clusterId).Limit(limit).Offset(offset).Name(name).OsPolicyId(osPolicyId).OsUserId(osUserId).ReplicationUuid(replicationUuid).Virtual(virtual).Q(q).Sort(sort).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterId := "clusterId_example" // string | cluster id (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	name := "name_example" // string | name of object storage buckets (optional)
	osPolicyId := int64(789) // int64 | The id of policy object storage buckets belong to (optional)
	osUserId := int64(789) // int64 | The id of user object storage buckets belong to (optional)
	replicationUuid := "replicationUuid_example" // string | The uuid of replication os buckets belong to (optional)
	virtual := true // bool | Virtual bucket or not (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.ListBuckets(context.Background()).ClusterId(clusterId).Limit(limit).Offset(offset).Name(name).OsPolicyId(osPolicyId).OsUserId(osUserId).ReplicationUuid(replicationUuid).Virtual(virtual).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.ListBuckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBuckets`: ObjectStorageBucketsResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.ListBuckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **name** | **string** | name of object storage buckets | 
 **osPolicyId** | **int64** | The id of policy object storage buckets belong to | 
 **osUserId** | **int64** | The id of user object storage buckets belong to | 
 **replicationUuid** | **string** | The uuid of replication os buckets belong to | 
 **virtual** | **bool** | Virtual bucket or not | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**ObjectStorageBucketsResp**](ObjectStorageBucketsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObjectStorageBucketNFSClients

> ObjectStorageBucketNFSClientsResp ListObjectStorageBucketNFSClients(ctx, bucketId).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.ListObjectStorageBucketNFSClients(context.Background(), bucketId).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.ListObjectStorageBucketNFSClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListObjectStorageBucketNFSClients`: ObjectStorageBucketNFSClientsResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.ListObjectStorageBucketNFSClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListObjectStorageBucketNFSClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 

### Return type

[**ObjectStorageBucketNFSClientsResp**](ObjectStorageBucketNFSClientsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCustomLabels

> ObjectStorageBucketResp RemoveCustomLabels(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | object storage bucket id
	body := *openapiclient.NewOSBucketCustomLabelsRemoveReq(*openapiclient.NewOSBucketCustomLabelsRemoveReqBucket([]int64{int64(123)})) // OSBucketCustomLabelsRemoveReq | custom labels info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.RemoveCustomLabels(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.RemoveCustomLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveCustomLabels`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.RemoveCustomLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | object storage bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCustomLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBucketCustomLabelsRemoveReq**](OSBucketCustomLabelsRemoveReq.md) | custom labels info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOSBucketLoggings

> ObjectStorageBucketResp RemoveOSBucketLoggings(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewOSBucketRemoveLoggingsReq(*openapiclient.NewOSBucketRemoveLoggingsReqBucket([]int64{int64(123)})) // OSBucketRemoveLoggingsReq | os bucket loggings info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.RemoveOSBucketLoggings(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.RemoveOSBucketLoggings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveOSBucketLoggings`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.RemoveOSBucketLoggings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOSBucketLoggingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBucketRemoveLoggingsReq**](OSBucketRemoveLoggingsReq.md) | os bucket loggings info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOSReplicationPaths

> ObjectStorageBucketResp RemoveOSReplicationPaths(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewOSBucketRemoveReplicationPathsReq(*openapiclient.NewOSBucketRemoveReplicationPathsReqBucket([]string{"OsReplicationPathUuids_example"})) // OSBucketRemoveReplicationPathsReq | replication paths info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.RemoveOSReplicationPaths(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.RemoveOSReplicationPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveOSReplicationPaths`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.RemoveOSReplicationPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOSReplicationPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBucketRemoveReplicationPathsReq**](OSBucketRemoveReplicationPathsReq.md) | replication paths info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOSReplicationZones

> ObjectStorageBucketResp RemoveOSReplicationZones(ctx, bucketId).Body(body).Force(force).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewOSBucketRemoveReplicationZonesReq(*openapiclient.NewOSBucketRemoveReplicationZonesReqBucket([]string{"OsReplicationZoneUuids_example"})) // OSBucketRemoveReplicationZonesReq | replication zones info
	force := true // bool | force delete os replication zone even if target zone is dead (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.RemoveOSReplicationZones(context.Background(), bucketId).Body(body).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.RemoveOSReplicationZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveOSReplicationZones`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.RemoveOSReplicationZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOSReplicationZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBucketRemoveReplicationZonesReq**](OSBucketRemoveReplicationZonesReq.md) | replication zones info | 
 **force** | **bool** | force delete os replication zone even if target zone is dead | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAccessLogging

> ObjectStorageBucketResp SetAccessLogging(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewOSBucketSetAccessLoggingReq(*openapiclient.NewOSBucketSetAccessLoggingReqBucket()) // OSBucketSetAccessLoggingReq | access logging info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.SetAccessLogging(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.SetAccessLogging``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAccessLogging`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.SetAccessLogging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAccessLoggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBucketSetAccessLoggingReq**](OSBucketSetAccessLoggingReq.md) | access logging info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMetadataSearch

> ObjectStorageBucketsResp SetMetadataSearch(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | object storage bucket id
	body := *openapiclient.NewOSBucketMetadataSearchSetReq(*openapiclient.NewOSBucketMetadataSearchSetReqBucket(false)) // OSBucketMetadataSearchSetReq | bucket metadata search info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.SetMetadataSearch(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.SetMetadataSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetMetadataSearch`: ObjectStorageBucketsResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.SetMetadataSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | object storage bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMetadataSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBucketMetadataSearchSetReq**](OSBucketMetadataSearchSetReq.md) | bucket metadata search info | 

### Return type

[**ObjectStorageBucketsResp**](ObjectStorageBucketsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOSBucketPolicy

> ObjectStorageBucketResp SetOSBucketPolicy(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewOSBucketPolicySetReq() // OSBucketPolicySetReq | bucket policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.SetOSBucketPolicy(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.SetOSBucketPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOSBucketPolicy`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.SetOSBucketPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetOSBucketPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBucketPolicySetReq**](OSBucketPolicySetReq.md) | bucket policy info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetObjectStorageClass

> ObjectStorageBucketResp SetObjectStorageClass(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewOSBucketSetObjectStorageClassReq(*openapiclient.NewOSBucketSetObjectStorageClassReqBucket(*openapiclient.NewOSBucketObjectStorageClass())) // OSBucketSetObjectStorageClassReq | object storage class info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.SetObjectStorageClass(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.SetObjectStorageClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetObjectStorageClass`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.SetObjectStorageClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetObjectStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBucketSetObjectStorageClassReq**](OSBucketSetObjectStorageClassReq.md) | object storage class info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetObjectStorageLifecycleRules

> ObjectStorageBucketResp SetObjectStorageLifecycleRules(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewObjectStorageLifecycleSetReq() // ObjectStorageLifecycleSetReq | lifecyce rules info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.SetObjectStorageLifecycleRules(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.SetObjectStorageLifecycleRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetObjectStorageLifecycleRules`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.SetObjectStorageLifecycleRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetObjectStorageLifecycleRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ObjectStorageLifecycleSetReq**](ObjectStorageLifecycleSetReq.md) | lifecyce rules info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendAccessLoggings

> ObjectStorageBucketResp SuspendAccessLoggings(ctx, bucketId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.SuspendAccessLoggings(context.Background(), bucketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.SuspendAccessLoggings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuspendAccessLoggings`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.SuspendAccessLoggings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendAccessLoggingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendOSReplicationPaths

> ObjectStorageBucketResp SuspendOSReplicationPaths(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewOSBucketUpdateReplicationPathsReq(*openapiclient.NewOSBucketUpdateReplicationPathsReqBucket([]string{"OsReplicationPathUuids_example"})) // OSBucketUpdateReplicationPathsReq | replication paths info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.SuspendOSReplicationPaths(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.SuspendOSReplicationPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuspendOSReplicationPaths`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.SuspendOSReplicationPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendOSReplicationPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBucketUpdateReplicationPathsReq**](OSBucketUpdateReplicationPathsReq.md) | replication paths info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwitchOwnerOSZone

> ObjectStorageBucketResp SwitchOwnerOSZone(ctx, bucketId).Body(body).Force(force).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewOSBucketSwitchOwnerOSZoneReq(*openapiclient.NewOSBucketSwitchOwnerOSZoneReqBucket()) // OSBucketSwitchOwnerOSZoneReq | new owner os zone info
	force := true // bool | force to switch even if old owner zone is dead (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.SwitchOwnerOSZone(context.Background(), bucketId).Body(body).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.SwitchOwnerOSZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SwitchOwnerOSZone`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.SwitchOwnerOSZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchOwnerOSZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBucketSwitchOwnerOSZoneReq**](OSBucketSwitchOwnerOSZoneReq.md) | new owner os zone info | 
 **force** | **bool** | force to switch even if old owner zone is dead | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetAccessLogging

> ObjectStorageBucketResp UnsetAccessLogging(ctx, bucketId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.UnsetAccessLogging(context.Background(), bucketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.UnsetAccessLogging``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsetAccessLogging`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.UnsetAccessLogging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetAccessLoggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetOSBucketPolicy

> ObjectStorageBucketResp UnsetOSBucketPolicy(ctx, bucketId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.UnsetOSBucketPolicy(context.Background(), bucketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.UnsetOSBucketPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsetOSBucketPolicy`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.UnsetOSBucketPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetOSBucketPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsuspendAccessLogging

> ObjectStorageBucketResp UnsuspendAccessLogging(ctx, bucketId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.UnsuspendAccessLogging(context.Background(), bucketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.UnsuspendAccessLogging``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsuspendAccessLogging`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.UnsuspendAccessLogging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsuspendAccessLoggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsuspendOSReplicationPaths

> ObjectStorageBucketResp UnsuspendOSReplicationPaths(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewOSBucketUpdateReplicationPathsReq(*openapiclient.NewOSBucketUpdateReplicationPathsReqBucket([]string{"OsReplicationPathUuids_example"})) // OSBucketUpdateReplicationPathsReq | replication paths info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.UnsuspendOSReplicationPaths(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.UnsuspendOSReplicationPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsuspendOSReplicationPaths`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.UnsuspendOSReplicationPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsuspendOSReplicationPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBucketUpdateReplicationPathsReq**](OSBucketUpdateReplicationPathsReq.md) | replication paths info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBucket

> ObjectStorageBucketResp UpdateBucket(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewObjectStorageBucketUpdateReq() // ObjectStorageBucketUpdateReq | bucket info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.UpdateBucket(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.UpdateBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBucket`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.UpdateBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ObjectStorageBucketUpdateReq**](ObjectStorageBucketUpdateReq.md) | bucket info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomLabels

> ObjectStorageBucketResp UpdateCustomLabels(ctx, bucketId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | object storage bucket id
	body := *openapiclient.NewOSBucketCustomLabelsUpdateReq(*openapiclient.NewOSBucketCustomLabelsUpdateReqBucket([]openapiclient.OSBucketCustomLabelsUpdateReqBucketLabelsElt{*openapiclient.NewOSBucketCustomLabelsUpdateReqBucketLabelsElt(int64(123), "Type_example")})) // OSBucketCustomLabelsUpdateReq | custom labels info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.UpdateCustomLabels(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.UpdateCustomLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomLabels`: ObjectStorageBucketResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.UpdateCustomLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | object storage bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBucketCustomLabelsUpdateReq**](OSBucketCustomLabelsUpdateReq.md) | custom labels info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateObjectStorageBucketNFSClient

> ObjectStorageBucketNFSClientResp UpdateObjectStorageBucketNFSClient(ctx, bucketId, clientId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bucketId := int64(789) // int64 | bucket id
	clientId := int64(789) // int64 | nfs client id
	body := *openapiclient.NewObjectStorageBucketNFSClientUpdateReq(*openapiclient.NewObjectStorageBucketNFSClientUpdateReqClient()) // ObjectStorageBucketNFSClientUpdateReq | nfs client info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketsAPI.UpdateObjectStorageBucketNFSClient(context.Background(), bucketId, clientId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketsAPI.UpdateObjectStorageBucketNFSClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateObjectStorageBucketNFSClient`: ObjectStorageBucketNFSClientResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketsAPI.UpdateObjectStorageBucketNFSClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 
**clientId** | **int64** | nfs client id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectStorageBucketNFSClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ObjectStorageBucketNFSClientUpdateReq**](ObjectStorageBucketNFSClientUpdateReq.md) | nfs client info | 

### Return type

[**ObjectStorageBucketNFSClientResp**](ObjectStorageBucketNFSClientResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

