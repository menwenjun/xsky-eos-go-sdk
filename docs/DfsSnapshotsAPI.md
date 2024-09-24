# \DfsSnapshotsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDfsSnapshot**](DfsSnapshotsAPI.md#CreateDfsSnapshot) | **Post** /dfs-snapshots/ | 
[**DeleteDfsSnapshot**](DfsSnapshotsAPI.md#DeleteDfsSnapshot) | **Delete** /dfs-snapshots/{dfs_snapshot_id} | 
[**GetDfsSnapshot**](DfsSnapshotsAPI.md#GetDfsSnapshot) | **Get** /dfs-snapshots/{dfs_snapshot_id} | 
[**GetDfsSnapshotsOverViewPage**](DfsSnapshotsAPI.md#GetDfsSnapshotsOverViewPage) | **Get** /dfs-snapshots/overview | 
[**ListDfsSnapshots**](DfsSnapshotsAPI.md#ListDfsSnapshots) | **Get** /dfs-snapshots/ | 
[**LockDfsSnapshot**](DfsSnapshotsAPI.md#LockDfsSnapshot) | **Post** /dfs-snapshots/{dfs_snapshot_id}:lock | 
[**RollbackDfsSnapshot**](DfsSnapshotsAPI.md#RollbackDfsSnapshot) | **Post** /dfs-snapshots/{dfs_snapshot_id}:rollback | 
[**UnlockDfsSnapshot**](DfsSnapshotsAPI.md#UnlockDfsSnapshot) | **Post** /dfs-snapshots/{dfs_snapshot_id}:unlock | 
[**UpdateDfsSnapshot**](DfsSnapshotsAPI.md#UpdateDfsSnapshot) | **Patch** /dfs-snapshots/{dfs_snapshot_id} | 



## CreateDfsSnapshot

> DfsSnapshotResp CreateDfsSnapshot(ctx).Body(body).AllowPathCreate(allowPathCreate).Execute()





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
	body := *openapiclient.NewDfsSnapshotCreateReq(*openapiclient.NewDfsSnapshotCreateReqDfsSnapshot("Description_example", "Name_example", "Path_example", "Retention_example", int64(123))) // DfsSnapshotCreateReq | dfs snapshot info
	allowPathCreate := true // bool | allow create path when not existed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSnapshotsAPI.CreateDfsSnapshot(context.Background()).Body(body).AllowPathCreate(allowPathCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSnapshotsAPI.CreateDfsSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsSnapshot`: DfsSnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSnapshotsAPI.CreateDfsSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsSnapshotCreateReq**](DfsSnapshotCreateReq.md) | dfs snapshot info | 
 **allowPathCreate** | **bool** | allow create path when not existed | 

### Return type

[**DfsSnapshotResp**](DfsSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsSnapshot

> DfsSnapshotResp DeleteDfsSnapshot(ctx, dfsSnapshotId).Execute()





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
	dfsSnapshotId := int64(789) // int64 | dfs snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSnapshotsAPI.DeleteDfsSnapshot(context.Background(), dfsSnapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSnapshotsAPI.DeleteDfsSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsSnapshot`: DfsSnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSnapshotsAPI.DeleteDfsSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsSnapshotId** | **int64** | dfs snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsSnapshotResp**](DfsSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsSnapshot

> DfsSnapshotResp GetDfsSnapshot(ctx, dfsSnapshotId).Execute()





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
	dfsSnapshotId := int64(789) // int64 | the dfs snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSnapshotsAPI.GetDfsSnapshot(context.Background(), dfsSnapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSnapshotsAPI.GetDfsSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsSnapshot`: DfsSnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSnapshotsAPI.GetDfsSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsSnapshotId** | **int64** | the dfs snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsSnapshotResp**](DfsSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsSnapshotsOverViewPage

> DfsSnapShotsOverviewPageResp GetDfsSnapshotsOverViewPage(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSnapshotsAPI.GetDfsSnapshotsOverViewPage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSnapshotsAPI.GetDfsSnapshotsOverViewPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsSnapshotsOverViewPage`: DfsSnapShotsOverviewPageResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSnapshotsAPI.GetDfsSnapshotsOverViewPage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsSnapshotsOverViewPageRequest struct via the builder pattern


### Return type

[**DfsSnapShotsOverviewPageResp**](DfsSnapShotsOverviewPageResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsSnapshots

> DfsSnapshotsResp ListDfsSnapshots(ctx).ClusterId(clusterId).DfsPathId(dfsPathId).DpDfsSnapshotId(dpDfsSnapshotId).Path(path).Name(name).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	dfsPathId := int64(789) // int64 | related dfs path id (optional)
	dpDfsSnapshotId := int64(789) // int64 | dp dfs snapshot id (optional)
	path := "path_example" // string | related dfs path (optional)
	name := "name_example" // string | name of dfs snapshot (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSnapshotsAPI.ListDfsSnapshots(context.Background()).ClusterId(clusterId).DfsPathId(dfsPathId).DpDfsSnapshotId(dpDfsSnapshotId).Path(path).Name(name).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSnapshotsAPI.ListDfsSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsSnapshots`: DfsSnapshotsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSnapshotsAPI.ListDfsSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **dfsPathId** | **int64** | related dfs path id | 
 **dpDfsSnapshotId** | **int64** | dp dfs snapshot id | 
 **path** | **string** | related dfs path | 
 **name** | **string** | name of dfs snapshot | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DfsSnapshotsResp**](DfsSnapshotsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockDfsSnapshot

> DfsSnapshotResp LockDfsSnapshot(ctx, dfsSnapshotId).Execute()





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
	dfsSnapshotId := int64(789) // int64 | dfs snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSnapshotsAPI.LockDfsSnapshot(context.Background(), dfsSnapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSnapshotsAPI.LockDfsSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LockDfsSnapshot`: DfsSnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSnapshotsAPI.LockDfsSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsSnapshotId** | **int64** | dfs snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockDfsSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsSnapshotResp**](DfsSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackDfsSnapshot

> DfsSnapshotResp RollbackDfsSnapshot(ctx, dfsSnapshotId).Body(body).Execute()





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
	dfsSnapshotId := int64(789) // int64 | dfs snapshot id
	body := *openapiclient.NewDfsSnapshotRollbackReq(*openapiclient.NewDfsSnapshotRollbackReqDfsSnapshot(false)) // DfsSnapshotRollbackReq | snapshot rollback info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSnapshotsAPI.RollbackDfsSnapshot(context.Background(), dfsSnapshotId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSnapshotsAPI.RollbackDfsSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RollbackDfsSnapshot`: DfsSnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSnapshotsAPI.RollbackDfsSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsSnapshotId** | **int64** | dfs snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackDfsSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsSnapshotRollbackReq**](DfsSnapshotRollbackReq.md) | snapshot rollback info | 

### Return type

[**DfsSnapshotResp**](DfsSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlockDfsSnapshot

> DfsSnapshotResp UnlockDfsSnapshot(ctx, dfsSnapshotId).Execute()





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
	dfsSnapshotId := int64(789) // int64 | dfs snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSnapshotsAPI.UnlockDfsSnapshot(context.Background(), dfsSnapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSnapshotsAPI.UnlockDfsSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnlockDfsSnapshot`: DfsSnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSnapshotsAPI.UnlockDfsSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsSnapshotId** | **int64** | dfs snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockDfsSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsSnapshotResp**](DfsSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsSnapshot

> DfsSnapshotResp UpdateDfsSnapshot(ctx, dfsSnapshotId).Body(body).Execute()





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
	dfsSnapshotId := int64(789) // int64 | dfs snapshot id
	body := *openapiclient.NewDfsSnapshotUpdateReq(*openapiclient.NewDfsSnapshotUpdateReqDfsSnapshot()) // DfsSnapshotUpdateReq | dfs snapshot info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSnapshotsAPI.UpdateDfsSnapshot(context.Background(), dfsSnapshotId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSnapshotsAPI.UpdateDfsSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsSnapshot`: DfsSnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSnapshotsAPI.UpdateDfsSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsSnapshotId** | **int64** | dfs snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsSnapshotUpdateReq**](DfsSnapshotUpdateReq.md) | dfs snapshot info | 

### Return type

[**DfsSnapshotResp**](DfsSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

