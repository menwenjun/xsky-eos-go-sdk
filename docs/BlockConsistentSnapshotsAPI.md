# \BlockConsistentSnapshotsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlockConsistentSnapshot**](BlockConsistentSnapshotsAPI.md#CreateBlockConsistentSnapshot) | **Post** /block-consistent-snapshots/ | 
[**DeleteBlockConsistentSnapshot**](BlockConsistentSnapshotsAPI.md#DeleteBlockConsistentSnapshot) | **Delete** /block-consistent-snapshots/{consistent_snapshot_id} | 
[**GetBlockConsistentSnapshot**](BlockConsistentSnapshotsAPI.md#GetBlockConsistentSnapshot) | **Get** /block-consistent-snapshots/{block_consistent_snapshot_id} | 
[**ListBlockConsistentSnapshots**](BlockConsistentSnapshotsAPI.md#ListBlockConsistentSnapshots) | **Get** /block-consistent-snapshots/ | 



## CreateBlockConsistentSnapshot

> ConsistentSnapshotResp CreateBlockConsistentSnapshot(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewConsistentSnapshotCreateReq(*openapiclient.NewConsistentSnapshotCreateReqConsistentSnapshot([]int64{int64(123)})) // ConsistentSnapshotCreateReq | consistent snapshot info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockConsistentSnapshotsAPI.CreateBlockConsistentSnapshot(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockConsistentSnapshotsAPI.CreateBlockConsistentSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlockConsistentSnapshot`: ConsistentSnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `BlockConsistentSnapshotsAPI.CreateBlockConsistentSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlockConsistentSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConsistentSnapshotCreateReq**](ConsistentSnapshotCreateReq.md) | consistent snapshot info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**ConsistentSnapshotResp**](ConsistentSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlockConsistentSnapshot

> SnapshotResp DeleteBlockConsistentSnapshot(ctx, consistentSnapshotId).Execute()





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
	consistentSnapshotId := int64(789) // int64 | consistent snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockConsistentSnapshotsAPI.DeleteBlockConsistentSnapshot(context.Background(), consistentSnapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockConsistentSnapshotsAPI.DeleteBlockConsistentSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlockConsistentSnapshot`: SnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `BlockConsistentSnapshotsAPI.DeleteBlockConsistentSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consistentSnapshotId** | **int64** | consistent snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlockConsistentSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SnapshotResp**](SnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockConsistentSnapshot

> SnapshotResp GetBlockConsistentSnapshot(ctx, blockConsistentSnapshotId).Execute()





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
	blockConsistentSnapshotId := int64(789) // int64 | the block consistent snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockConsistentSnapshotsAPI.GetBlockConsistentSnapshot(context.Background(), blockConsistentSnapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockConsistentSnapshotsAPI.GetBlockConsistentSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockConsistentSnapshot`: SnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `BlockConsistentSnapshotsAPI.GetBlockConsistentSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockConsistentSnapshotId** | **int64** | the block consistent snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockConsistentSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SnapshotResp**](SnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlockConsistentSnapshots

> ConsistentSnapshotsResp ListBlockConsistentSnapshots(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockConsistentSnapshotsAPI.ListBlockConsistentSnapshots(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockConsistentSnapshotsAPI.ListBlockConsistentSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlockConsistentSnapshots`: ConsistentSnapshotsResp
	fmt.Fprintf(os.Stdout, "Response from `BlockConsistentSnapshotsAPI.ListBlockConsistentSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlockConsistentSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**ConsistentSnapshotsResp**](ConsistentSnapshotsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

