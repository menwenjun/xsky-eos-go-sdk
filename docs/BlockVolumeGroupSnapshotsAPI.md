# \BlockVolumeGroupSnapshotsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlockVolumeGroupSnapshot**](BlockVolumeGroupSnapshotsAPI.md#CreateBlockVolumeGroupSnapshot) | **Post** /block-volume-group-snapshots/ | 
[**DeleteBlockVolumeGroupSnapshot**](BlockVolumeGroupSnapshotsAPI.md#DeleteBlockVolumeGroupSnapshot) | **Delete** /block-volume-group-snapshots/{block_volume_group_snapshot_id} | 
[**GetBlockVolumeGroupSnapshot**](BlockVolumeGroupSnapshotsAPI.md#GetBlockVolumeGroupSnapshot) | **Get** /block-volume-group-snapshots/{block_volume_group_snapshot_id} | 
[**ListBlockVolumeGroupSnapshots**](BlockVolumeGroupSnapshotsAPI.md#ListBlockVolumeGroupSnapshots) | **Get** /block-volume-group-snapshots/ | 
[**UpdateBlockVolumeGroupSnapshot**](BlockVolumeGroupSnapshotsAPI.md#UpdateBlockVolumeGroupSnapshot) | **Patch** /block-volume-group-snapshots/{block_volume_group_snapshot_id} | 



## CreateBlockVolumeGroupSnapshot

> VolumeGroupSnapshotResp CreateBlockVolumeGroupSnapshot(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewVolumeGroupSnapshotCreateReq(*openapiclient.NewVolumeGroupSnapshotCreateReqVolumeGroupSnapshot(int64(123), "Description_example", "Name_example")) // VolumeGroupSnapshotCreateReq | volume group snapshot info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeGroupSnapshotsAPI.CreateBlockVolumeGroupSnapshot(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeGroupSnapshotsAPI.CreateBlockVolumeGroupSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlockVolumeGroupSnapshot`: VolumeGroupSnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeGroupSnapshotsAPI.CreateBlockVolumeGroupSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlockVolumeGroupSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VolumeGroupSnapshotCreateReq**](VolumeGroupSnapshotCreateReq.md) | volume group snapshot info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**VolumeGroupSnapshotResp**](VolumeGroupSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlockVolumeGroupSnapshot

> VolumeGroupSnapshotResp DeleteBlockVolumeGroupSnapshot(ctx, blockVolumeGroupSnapshotId).Execute()





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
	blockVolumeGroupSnapshotId := int64(789) // int64 | block volume group snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeGroupSnapshotsAPI.DeleteBlockVolumeGroupSnapshot(context.Background(), blockVolumeGroupSnapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeGroupSnapshotsAPI.DeleteBlockVolumeGroupSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlockVolumeGroupSnapshot`: VolumeGroupSnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeGroupSnapshotsAPI.DeleteBlockVolumeGroupSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeGroupSnapshotId** | **int64** | block volume group snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlockVolumeGroupSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeGroupSnapshotResp**](VolumeGroupSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockVolumeGroupSnapshot

> VolumeGroupSnapshotResp GetBlockVolumeGroupSnapshot(ctx, blockVolumeGroupSnapshotId).Execute()





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
	blockVolumeGroupSnapshotId := int64(789) // int64 | the block volume group snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeGroupSnapshotsAPI.GetBlockVolumeGroupSnapshot(context.Background(), blockVolumeGroupSnapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeGroupSnapshotsAPI.GetBlockVolumeGroupSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockVolumeGroupSnapshot`: VolumeGroupSnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeGroupSnapshotsAPI.GetBlockVolumeGroupSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeGroupSnapshotId** | **int64** | the block volume group snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockVolumeGroupSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeGroupSnapshotResp**](VolumeGroupSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlockVolumeGroupSnapshots

> VolumeGroupSnapshotsResp ListBlockVolumeGroupSnapshots(ctx).ClusterId(clusterId).BlockVolumeGroupId(blockVolumeGroupId).Name(name).Passive(passive).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	blockVolumeGroupId := int64(789) // int64 | related volume group id (optional)
	name := "name_example" // string | name of volume group snapshot (optional)
	passive := true // bool | passive or not (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeGroupSnapshotsAPI.ListBlockVolumeGroupSnapshots(context.Background()).ClusterId(clusterId).BlockVolumeGroupId(blockVolumeGroupId).Name(name).Passive(passive).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeGroupSnapshotsAPI.ListBlockVolumeGroupSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlockVolumeGroupSnapshots`: VolumeGroupSnapshotsResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeGroupSnapshotsAPI.ListBlockVolumeGroupSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlockVolumeGroupSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **blockVolumeGroupId** | **int64** | related volume group id | 
 **name** | **string** | name of volume group snapshot | 
 **passive** | **bool** | passive or not | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**VolumeGroupSnapshotsResp**](VolumeGroupSnapshotsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlockVolumeGroupSnapshot

> VolumeGroupSnapshotResp UpdateBlockVolumeGroupSnapshot(ctx, blockVolumeGroupSnapshotId).Body(body).Execute()





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
	blockVolumeGroupSnapshotId := int64(789) // int64 | the block volume group snapshot id
	body := *openapiclient.NewVolumeGroupSnapshotUpdateReq(*openapiclient.NewVolumeGroupSnapshotUpdateReqVolumeGroupSnapshot()) // VolumeGroupSnapshotUpdateReq | volume group snapshot info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeGroupSnapshotsAPI.UpdateBlockVolumeGroupSnapshot(context.Background(), blockVolumeGroupSnapshotId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeGroupSnapshotsAPI.UpdateBlockVolumeGroupSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBlockVolumeGroupSnapshot`: VolumeGroupSnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeGroupSnapshotsAPI.UpdateBlockVolumeGroupSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeGroupSnapshotId** | **int64** | the block volume group snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlockVolumeGroupSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VolumeGroupSnapshotUpdateReq**](VolumeGroupSnapshotUpdateReq.md) | volume group snapshot info | 

### Return type

[**VolumeGroupSnapshotResp**](VolumeGroupSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

