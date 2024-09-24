# \DfsRootfsesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDfsRootfs**](DfsRootfsesAPI.md#CreateDfsRootfs) | **Post** /dfs-rootfses/ | 
[**DeleteDfsRootfs**](DfsRootfsesAPI.md#DeleteDfsRootfs) | **Delete** /dfs-rootfses/{dfs_rootfs_id} | 
[**GetDfsRootfs**](DfsRootfsesAPI.md#GetDfsRootfs) | **Get** /dfs-rootfses/{dfs_rootfs_id} | 
[**GetDfsRootfsSamples**](DfsRootfsesAPI.md#GetDfsRootfsSamples) | **Get** /dfs-rootfses/{dfs_rootfs_id}/samples | 
[**ListDfsRootfses**](DfsRootfsesAPI.md#ListDfsRootfses) | **Get** /dfs-rootfses/ | 
[**SetDfsWormLogPath**](DfsRootfsesAPI.md#SetDfsWormLogPath) | **Patch** /dfs-rootfses/{dfs_rootfs_id}:set-worm-log-path | 
[**SetGCSpeed**](DfsRootfsesAPI.md#SetGCSpeed) | **Post** /dfs-rootfses/{dfs_rootfs_id}:set-gc-speed | 
[**UpdateDfsRootfsActivePool**](DfsRootfsesAPI.md#UpdateDfsRootfsActivePool) | **Patch** /dfs-rootfses/{dfs_rootfs_id}:update-active-pools | 



## CreateDfsRootfs

> DfsRootfsResp CreateDfsRootfs(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsRootfsCreateReq(*openapiclient.NewDfsRootfsCreateReqRootfs(*openapiclient.NewDfsRootfsCreateReqRootfsAuditLog(), int64(123), "Name_example")) // DfsRootfsCreateReq | rootfs info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsRootfsesAPI.CreateDfsRootfs(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsRootfsesAPI.CreateDfsRootfs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsRootfs`: DfsRootfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsRootfsesAPI.CreateDfsRootfs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsRootfsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsRootfsCreateReq**](DfsRootfsCreateReq.md) | rootfs info | 

### Return type

[**DfsRootfsResp**](DfsRootfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsRootfs

> DfsRootfsResp DeleteDfsRootfs(ctx, dfsRootfsId).Force(force).Execute()





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
	dfsRootfsId := int64(789) // int64 | rootfs id
	force := true // bool | force delete (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsRootfsesAPI.DeleteDfsRootfs(context.Background(), dfsRootfsId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsRootfsesAPI.DeleteDfsRootfs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsRootfs`: DfsRootfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsRootfsesAPI.DeleteDfsRootfs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsRootfsId** | **int64** | rootfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsRootfsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete | 

### Return type

[**DfsRootfsResp**](DfsRootfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsRootfs

> DfsRootfsResp GetDfsRootfs(ctx, dfsRootfsId).Execute()





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
	dfsRootfsId := int64(789) // int64 | rootfs id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsRootfsesAPI.GetDfsRootfs(context.Background(), dfsRootfsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsRootfsesAPI.GetDfsRootfs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsRootfs`: DfsRootfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsRootfsesAPI.GetDfsRootfs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsRootfsId** | **int64** | rootfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsRootfsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsRootfsResp**](DfsRootfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsRootfsSamples

> DfsRootfsSamplesResp GetDfsRootfsSamples(ctx, dfsRootfsId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	dfsRootfsId := int64(789) // int64 | dfs rootfs id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsRootfsesAPI.GetDfsRootfsSamples(context.Background(), dfsRootfsId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsRootfsesAPI.GetDfsRootfsSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsRootfsSamples`: DfsRootfsSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsRootfsesAPI.GetDfsRootfsSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsRootfsId** | **int64** | dfs rootfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsRootfsSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**DfsRootfsSamplesResp**](DfsRootfsSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsRootfses

> DfsRootfsesResp ListDfsRootfses(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).PoolId(poolId).FsUserId(fsUserId).FsUserGroupId(fsUserGroupId).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	poolId := int64(789) // int64 | pool id (optional)
	fsUserId := int64(789) // int64 | fs user id (optional)
	fsUserGroupId := int64(789) // int64 | fs user group id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsRootfsesAPI.ListDfsRootfses(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).PoolId(poolId).FsUserId(fsUserId).FsUserGroupId(fsUserGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsRootfsesAPI.ListDfsRootfses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsRootfses`: DfsRootfsesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsRootfsesAPI.ListDfsRootfses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsRootfsesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 
 **poolId** | **int64** | pool id | 
 **fsUserId** | **int64** | fs user id | 
 **fsUserGroupId** | **int64** | fs user group id | 

### Return type

[**DfsRootfsesResp**](DfsRootfsesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDfsWormLogPath

> DfsRootfsResp SetDfsWormLogPath(ctx, dfsRootfsId).Body(body).AllowPathCreate(allowPathCreate).Execute()





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
	dfsRootfsId := int64(789) // int64 | rootfs id
	body := *openapiclient.NewDfsRootfsSetWormLogPathReq(*openapiclient.NewDfsRootfsSetWormLogPathReqRootfs("WormLogPath_example")) // DfsRootfsSetWormLogPathReq | worm log path
	allowPathCreate := true // bool | allow create path when not existed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsRootfsesAPI.SetDfsWormLogPath(context.Background(), dfsRootfsId).Body(body).AllowPathCreate(allowPathCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsRootfsesAPI.SetDfsWormLogPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDfsWormLogPath`: DfsRootfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsRootfsesAPI.SetDfsWormLogPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsRootfsId** | **int64** | rootfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDfsWormLogPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsRootfsSetWormLogPathReq**](DfsRootfsSetWormLogPathReq.md) | worm log path | 
 **allowPathCreate** | **bool** | allow create path when not existed | 

### Return type

[**DfsRootfsResp**](DfsRootfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGCSpeed

> DfsRootfsResp SetGCSpeed(ctx, dfsRootfsId).Body(body).Execute()





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
	dfsRootfsId := int64(789) // int64 | rootfs id
	body := *openapiclient.NewDfsRootfsSetGCSpeedReq(*openapiclient.NewDfsRootfsSetGCSpeedReqRootfs()) // DfsRootfsSetGCSpeedReq | gc speed

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsRootfsesAPI.SetGCSpeed(context.Background(), dfsRootfsId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsRootfsesAPI.SetGCSpeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGCSpeed`: DfsRootfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsRootfsesAPI.SetGCSpeed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsRootfsId** | **int64** | rootfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGCSpeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsRootfsSetGCSpeedReq**](DfsRootfsSetGCSpeedReq.md) | gc speed | 

### Return type

[**DfsRootfsResp**](DfsRootfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsRootfsActivePool

> DfsRootfsResp UpdateDfsRootfsActivePool(ctx, dfsRootfsId).Body(body).Execute()





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
	dfsRootfsId := int64(789) // int64 | rootfs id
	body := *openapiclient.NewDfsRootfsUpdateActivePoolReq(*openapiclient.NewDfsRootfsUpdateActivePoolReqRootfs()) // DfsRootfsUpdateActivePoolReq | active pool ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsRootfsesAPI.UpdateDfsRootfsActivePool(context.Background(), dfsRootfsId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsRootfsesAPI.UpdateDfsRootfsActivePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsRootfsActivePool`: DfsRootfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsRootfsesAPI.UpdateDfsRootfsActivePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsRootfsId** | **int64** | rootfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsRootfsActivePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsRootfsUpdateActivePoolReq**](DfsRootfsUpdateActivePoolReq.md) | active pool ids | 

### Return type

[**DfsRootfsResp**](DfsRootfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

