# \DpDfsSnapshotsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpDfsSnapshot**](DpDfsSnapshotsAPI.md#CreateDpDfsSnapshot) | **Post** /dp-dfs-snapshots/ | 
[**DeleteDPDfsSnapshot**](DpDfsSnapshotsAPI.md#DeleteDPDfsSnapshot) | **Delete** /dp-dfs-snapshots/{dp_dfs_snapshot_id} | 
[**GetDpDfsSnapshot**](DpDfsSnapshotsAPI.md#GetDpDfsSnapshot) | **Get** /dp-dfs-snapshots/{dp_dfs_snapshot_id} | 
[**ListDpDfsSnapshots**](DpDfsSnapshotsAPI.md#ListDpDfsSnapshots) | **Get** /dp-dfs-snapshots/ | 
[**ListDpDfsSnapshotsByDfsPathName**](DpDfsSnapshotsAPI.md#ListDpDfsSnapshotsByDfsPathName) | **Post** /dp-dfs-snapshots/_search | 



## CreateDpDfsSnapshot

> DpDfsSnapshotResp CreateDpDfsSnapshot(ctx).Body(body).AllowPathCreate(allowPathCreate).Execute()





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
	body := *openapiclient.NewDpDfsSnapshotCreateReq(*openapiclient.NewDpDfsSnapshotCreateReqDpDfsSnapshot(int64(123), "Path_example", int64(123), false)) // DpDfsSnapshotCreateReq | dp dfs snapshot info
	allowPathCreate := true // bool | allow create path when not existed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpDfsSnapshotsAPI.CreateDpDfsSnapshot(context.Background()).Body(body).AllowPathCreate(allowPathCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpDfsSnapshotsAPI.CreateDpDfsSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDpDfsSnapshot`: DpDfsSnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `DpDfsSnapshotsAPI.CreateDpDfsSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDpDfsSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DpDfsSnapshotCreateReq**](DpDfsSnapshotCreateReq.md) | dp dfs snapshot info | 
 **allowPathCreate** | **bool** | allow create path when not existed | 

### Return type

[**DpDfsSnapshotResp**](DpDfsSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDPDfsSnapshot

> DpDfsSnapshotResp DeleteDPDfsSnapshot(ctx, dpDfsSnapshotId).Execute()





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
	dpDfsSnapshotId := int64(789) // int64 | dp dfs snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpDfsSnapshotsAPI.DeleteDPDfsSnapshot(context.Background(), dpDfsSnapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpDfsSnapshotsAPI.DeleteDPDfsSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDPDfsSnapshot`: DpDfsSnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `DpDfsSnapshotsAPI.DeleteDPDfsSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dpDfsSnapshotId** | **int64** | dp dfs snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDPDfsSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpDfsSnapshotResp**](DpDfsSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDpDfsSnapshot

> DpDfsSnapshotResp GetDpDfsSnapshot(ctx, dpDfsSnapshotId).Execute()





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
	dpDfsSnapshotId := int64(789) // int64 | the dp dfs snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpDfsSnapshotsAPI.GetDpDfsSnapshot(context.Background(), dpDfsSnapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpDfsSnapshotsAPI.GetDpDfsSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDpDfsSnapshot`: DpDfsSnapshotResp
	fmt.Fprintf(os.Stdout, "Response from `DpDfsSnapshotsAPI.GetDpDfsSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dpDfsSnapshotId** | **int64** | the dp dfs snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDpDfsSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpDfsSnapshotResp**](DpDfsSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpDfsSnapshots

> DpDfsSnapshotsResp ListDpDfsSnapshots(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).PolicyId(policyId).PathId(pathId).Execute()





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
	policyId := int64(789) // int64 | related dp dfs snapshot policy id (optional)
	pathId := int64(789) // int64 | related dfs path id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpDfsSnapshotsAPI.ListDpDfsSnapshots(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).PolicyId(policyId).PathId(pathId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpDfsSnapshotsAPI.ListDpDfsSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpDfsSnapshots`: DpDfsSnapshotsResp
	fmt.Fprintf(os.Stdout, "Response from `DpDfsSnapshotsAPI.ListDpDfsSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpDfsSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **policyId** | **int64** | related dp dfs snapshot policy id | 
 **pathId** | **int64** | related dfs path id | 

### Return type

[**DpDfsSnapshotsResp**](DpDfsSnapshotsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpDfsSnapshotsByDfsPathName

> DpDfsSnapshotsResp ListDpDfsSnapshotsByDfsPathName(ctx).Execute()





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
	resp, r, err := apiClient.DpDfsSnapshotsAPI.ListDpDfsSnapshotsByDfsPathName(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpDfsSnapshotsAPI.ListDpDfsSnapshotsByDfsPathName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpDfsSnapshotsByDfsPathName`: DpDfsSnapshotsResp
	fmt.Fprintf(os.Stdout, "Response from `DpDfsSnapshotsAPI.ListDpDfsSnapshotsByDfsPathName`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDpDfsSnapshotsByDfsPathNameRequest struct via the builder pattern


### Return type

[**DpDfsSnapshotsResp**](DpDfsSnapshotsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

