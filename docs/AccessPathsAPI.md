# \AccessPathsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessPath**](AccessPathsAPI.md#CreateAccessPath) | **Post** /access-paths/ | 
[**DeleteAccessPath**](AccessPathsAPI.md#DeleteAccessPath) | **Delete** /access-paths/{access_path_id} | 
[**GetAccessPath**](AccessPathsAPI.md#GetAccessPath) | **Get** /access-paths/{access_path_id} | 
[**ListAccessPaths**](AccessPathsAPI.md#ListAccessPaths) | **Get** /access-paths/ | 
[**UpdateAccessPath**](AccessPathsAPI.md#UpdateAccessPath) | **Patch** /access-paths/{access_path_id} | 



## CreateAccessPath

> AccessPathResp CreateAccessPath(ctx).AccessPath(accessPath).ClusterId(clusterId).Execute()





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
	accessPath := *openapiclient.NewAccessPathCreateReq() // AccessPathCreateReq | access path info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessPathsAPI.CreateAccessPath(context.Background()).AccessPath(accessPath).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessPathsAPI.CreateAccessPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccessPath`: AccessPathResp
	fmt.Fprintf(os.Stdout, "Response from `AccessPathsAPI.CreateAccessPath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessPath** | [**AccessPathCreateReq**](AccessPathCreateReq.md) | access path info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**AccessPathResp**](AccessPathResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessPath

> AccessPathResp DeleteAccessPath(ctx, accessPathId).Execute()





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
	accessPathId := int64(789) // int64 | access path id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessPathsAPI.DeleteAccessPath(context.Background(), accessPathId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessPathsAPI.DeleteAccessPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccessPath`: AccessPathResp
	fmt.Fprintf(os.Stdout, "Response from `AccessPathsAPI.DeleteAccessPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessPathId** | **int64** | access path id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessPathResp**](AccessPathResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessPath

> AccessPathResp GetAccessPath(ctx, accessPathId).Execute()





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
	accessPathId := int64(789) // int64 | access path id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessPathsAPI.GetAccessPath(context.Background(), accessPathId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessPathsAPI.GetAccessPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessPath`: AccessPathResp
	fmt.Fprintf(os.Stdout, "Response from `AccessPathsAPI.GetAccessPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessPathId** | **int64** | access path id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessPathResp**](AccessPathResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessPaths

> AccessPathsResp ListAccessPaths(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).Q(q).Uid(uid).ClientGroupId(clientGroupId).All(all).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	q := "q_example" // string | query param of search (optional)
	uid := "uid_example" // string | access path uid (optional)
	clientGroupId := int64(789) // int64 | related client group id (optional)
	all := true // bool | show all access targets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessPathsAPI.ListAccessPaths(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).Q(q).Uid(uid).ClientGroupId(clientGroupId).All(all).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessPathsAPI.ListAccessPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccessPaths`: AccessPathsResp
	fmt.Fprintf(os.Stdout, "Response from `AccessPathsAPI.ListAccessPaths`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **q** | **string** | query param of search | 
 **uid** | **string** | access path uid | 
 **clientGroupId** | **int64** | related client group id | 
 **all** | **bool** | show all access targets | 

### Return type

[**AccessPathsResp**](AccessPathsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessPath

> AccessPathResp UpdateAccessPath(ctx, accessPathId).AccessPath(accessPath).Execute()





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
	accessPathId := int64(789) // int64 | access path id
	accessPath := *openapiclient.NewAccessPathUpdateReq(*openapiclient.NewAccessPathUpdateReqAccessPath()) // AccessPathUpdateReq | access path info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessPathsAPI.UpdateAccessPath(context.Background(), accessPathId).AccessPath(accessPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessPathsAPI.UpdateAccessPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccessPath`: AccessPathResp
	fmt.Fprintf(os.Stdout, "Response from `AccessPathsAPI.UpdateAccessPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessPathId** | **int64** | access path id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessPath** | [**AccessPathUpdateReq**](AccessPathUpdateReq.md) | access path info | 

### Return type

[**AccessPathResp**](AccessPathResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

