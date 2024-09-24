# \BlockSnapshotsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlockSnapshot**](BlockSnapshotsAPI.md#CreateBlockSnapshot) | **Post** /block-snapshots/ | 
[**DeleteBlockSnapshot**](BlockSnapshotsAPI.md#DeleteBlockSnapshot) | **Delete** /block-snapshots/{snapshot_id} | 
[**GetBlockSnapshot**](BlockSnapshotsAPI.md#GetBlockSnapshot) | **Get** /block-snapshots/{block_snapshot_id} | 
[**ListBlockSnapshots**](BlockSnapshotsAPI.md#ListBlockSnapshots) | **Get** /block-snapshots/ | 
[**UpdateSnapshot**](BlockSnapshotsAPI.md#UpdateSnapshot) | **Patch** /block-snapshots/{snapshot_id} | 



## CreateBlockSnapshot

> SnapshotResp CreateBlockSnapshot(ctx).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/xsky-eos-go-sdk"
)

func main() {
	body := *openapiclient.NewSnapshotCreateReq(*openapiclient.NewSnapshotCreateReqSnapshot(int64(123), "Description_example", "Name_example")) // SnapshotCreateReq | snapshot info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockSnapshotsAPI.CreateBlockSnapshot(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockSnapshotsAPI.CreateBlockSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlockSnapshot`: SnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `BlockSnapshotsAPI.CreateBlockSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlockSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SnapshotCreateReq**](SnapshotCreateReq.md) | snapshot info | 

### Return type

[**SnapshotResp**](SnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlockSnapshot

> SnapshotResp DeleteBlockSnapshot(ctx, snapshotId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/xsky-eos-go-sdk"
)

func main() {
	snapshotId := int64(789) // int64 | snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockSnapshotsAPI.DeleteBlockSnapshot(context.Background(), snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockSnapshotsAPI.DeleteBlockSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlockSnapshot`: SnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `BlockSnapshotsAPI.DeleteBlockSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **int64** | snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlockSnapshotRequest struct via the builder pattern


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


## GetBlockSnapshot

> SnapshotResp GetBlockSnapshot(ctx, blockSnapshotId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/xsky-eos-go-sdk"
)

func main() {
	blockSnapshotId := int64(789) // int64 | the block snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockSnapshotsAPI.GetBlockSnapshot(context.Background(), blockSnapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockSnapshotsAPI.GetBlockSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockSnapshot`: SnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `BlockSnapshotsAPI.GetBlockSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockSnapshotId** | **int64** | the block snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockSnapshotRequest struct via the builder pattern


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


## ListBlockSnapshots

> SnapshotsResp ListBlockSnapshots(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).PoolId(poolId).Uid(uid).BlockVolumeId(blockVolumeId).ConsistentSnapshotId(consistentSnapshotId).Reserved(reserved).Q(q).Sort(sort).All(all).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/xsky-eos-go-sdk"
)

func main() {
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)
	poolId := int64(789) // int64 | pool id (optional)
	uid := "uid_example" // string | snapshot uid (optional)
	blockVolumeId := int64(789) // int64 | volume id (optional)
	consistentSnapshotId := int64(789) // int64 | consistent snapshot id (optional)
	reserved := true // bool | show reserved snapshot or not, default not (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	all := true // bool | show all snapshots (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockSnapshotsAPI.ListBlockSnapshots(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).PoolId(poolId).Uid(uid).BlockVolumeId(blockVolumeId).ConsistentSnapshotId(consistentSnapshotId).Reserved(reserved).Q(q).Sort(sort).All(all).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockSnapshotsAPI.ListBlockSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlockSnapshots`: SnapshotsResp
	fmt.Fprintf(os.Stdout, "Response from `BlockSnapshotsAPI.ListBlockSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlockSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **poolId** | **int64** | pool id | 
 **uid** | **string** | snapshot uid | 
 **blockVolumeId** | **int64** | volume id | 
 **consistentSnapshotId** | **int64** | consistent snapshot id | 
 **reserved** | **bool** | show reserved snapshot or not, default not | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **all** | **bool** | show all snapshots | 

### Return type

[**SnapshotsResp**](SnapshotsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSnapshot

> SnapshotResp UpdateSnapshot(ctx, snapshotId).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/xsky-eos-go-sdk"
)

func main() {
	snapshotId := int64(789) // int64 | snapshot id
	body := *openapiclient.NewSnapshotUpdateReq(*openapiclient.NewSnapshotUpdateReqSnapshot()) // SnapshotUpdateReq | snapshot info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockSnapshotsAPI.UpdateSnapshot(context.Background(), snapshotId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockSnapshotsAPI.UpdateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSnapshot`: SnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `BlockSnapshotsAPI.UpdateSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **int64** | snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SnapshotUpdateReq**](SnapshotUpdateReq.md) | snapshot info | 

### Return type

[**SnapshotResp**](SnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

