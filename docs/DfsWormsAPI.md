# \DfsWormsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDfsWorm**](DfsWormsAPI.md#CreateDfsWorm) | **Post** /dfs-worms/ | 
[**DeleteDfsWorm**](DfsWormsAPI.md#DeleteDfsWorm) | **Delete** /dfs-worms/{dfs_worm_id} | 
[**GetDfsWorm**](DfsWormsAPI.md#GetDfsWorm) | **Get** /dfs-worms/{dfs_worm_id} | 
[**ListDfsWorms**](DfsWormsAPI.md#ListDfsWorms) | **Get** /dfs-worms/ | 
[**UpdateDfsWorm**](DfsWormsAPI.md#UpdateDfsWorm) | **Patch** /dfs-worms/{dfs_worm_id} | 



## CreateDfsWorm

> DfsWormResp CreateDfsWorm(ctx).Body(body).AllowPathCreate(allowPathCreate).Execute()





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
	body := *openapiclient.NewDfsWormCreateReq(*openapiclient.NewDfsWormCreateReqWorm("AutoLockPeriod_example", "DefaultProtectPeriod_example", "MaxProtectPeriod_example", "MinProtectPeriod_example", "Path_example", int64(123))) // DfsWormCreateReq | worm info
	allowPathCreate := true // bool | allow create path when not existed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsWormsAPI.CreateDfsWorm(context.Background()).Body(body).AllowPathCreate(allowPathCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsWormsAPI.CreateDfsWorm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsWorm`: DfsWormResp
	fmt.Fprintf(os.Stdout, "Response from `DfsWormsAPI.CreateDfsWorm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsWormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsWormCreateReq**](DfsWormCreateReq.md) | worm info | 
 **allowPathCreate** | **bool** | allow create path when not existed | 

### Return type

[**DfsWormResp**](DfsWormResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsWorm

> DfsWormResp DeleteDfsWorm(ctx, dfsWormId).Execute()





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
	dfsWormId := int64(789) // int64 | worm id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsWormsAPI.DeleteDfsWorm(context.Background(), dfsWormId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsWormsAPI.DeleteDfsWorm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsWorm`: DfsWormResp
	fmt.Fprintf(os.Stdout, "Response from `DfsWormsAPI.DeleteDfsWorm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsWormId** | **int64** | worm id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsWormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsWormResp**](DfsWormResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsWorm

> DfsWormResp GetDfsWorm(ctx, dfsWormId).Execute()





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
	dfsWormId := int64(789) // int64 | the dfs worm id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsWormsAPI.GetDfsWorm(context.Background(), dfsWormId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsWormsAPI.GetDfsWorm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsWorm`: DfsWormResp
	fmt.Fprintf(os.Stdout, "Response from `DfsWormsAPI.GetDfsWorm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsWormId** | **int64** | the dfs worm id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsWormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsWormResp**](DfsWormResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsWorms

> DfsWormsResp ListDfsWorms(ctx).ClusterId(clusterId).DfsPathId(dfsPathId).Path(path).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	path := "path_example" // string | related dfs path (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsWormsAPI.ListDfsWorms(context.Background()).ClusterId(clusterId).DfsPathId(dfsPathId).Path(path).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsWormsAPI.ListDfsWorms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsWorms`: DfsWormsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsWormsAPI.ListDfsWorms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsWormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **dfsPathId** | **int64** | related dfs path id | 
 **path** | **string** | related dfs path | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DfsWormsResp**](DfsWormsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsWorm

> DfsWormResp UpdateDfsWorm(ctx, dfsWormId).Body(body).Execute()





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
	dfsWormId := int64(789) // int64 | dfs worm id
	body := *openapiclient.NewDfsWormUpdateReq(*openapiclient.NewDfsWormUpdateReqWorm()) // DfsWormUpdateReq | dfs worm info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsWormsAPI.UpdateDfsWorm(context.Background(), dfsWormId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsWormsAPI.UpdateDfsWorm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsWorm`: DfsWormResp
	fmt.Fprintf(os.Stdout, "Response from `DfsWormsAPI.UpdateDfsWorm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsWormId** | **int64** | dfs worm id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsWormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsWormUpdateReq**](DfsWormUpdateReq.md) | dfs worm info | 

### Return type

[**DfsWormResp**](DfsWormResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

