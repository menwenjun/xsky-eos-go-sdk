# \FsActiveDirectoriesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFSActiveDirectory**](FsActiveDirectoriesAPI.md#CreateFSActiveDirectory) | **Post** /fs-active-directories/ | 
[**DeleteFSActiveDirectory**](FsActiveDirectoriesAPI.md#DeleteFSActiveDirectory) | **Delete** /fs-active-directories/{fs_active_directory_id} | 
[**GetFSActiveDirectory**](FsActiveDirectoriesAPI.md#GetFSActiveDirectory) | **Get** /fs-active-directories/{fs_active_directory_id} | 
[**ListFSActiveDirectories**](FsActiveDirectoriesAPI.md#ListFSActiveDirectories) | **Get** /fs-active-directories/ | 
[**UpdateFSActiveDirectory**](FsActiveDirectoriesAPI.md#UpdateFSActiveDirectory) | **Patch** /fs-active-directories/{fs_active_directory_id} | 
[**VerifyFSActiveDirectory**](FsActiveDirectoriesAPI.md#VerifyFSActiveDirectory) | **Post** /fs-active-directories/{fs_active_directory_id}:verify | 



## CreateFSActiveDirectory

> FSActiveDirectoryResp CreateFSActiveDirectory(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewFSActiveDirectoryCreateReq(*openapiclient.NewFSActiveDirectoryCreateReqInfo("Name_example", "Password_example", "Realm_example", "Username_example", "Workgroup_example")) // FSActiveDirectoryCreateReq | file storage active directory info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsActiveDirectoriesAPI.CreateFSActiveDirectory(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsActiveDirectoriesAPI.CreateFSActiveDirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFSActiveDirectory`: FSActiveDirectoryResp
	fmt.Fprintf(os.Stdout, "Response from `FsActiveDirectoriesAPI.CreateFSActiveDirectory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFSActiveDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FSActiveDirectoryCreateReq**](FSActiveDirectoryCreateReq.md) | file storage active directory info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**FSActiveDirectoryResp**](FSActiveDirectoryResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFSActiveDirectory

> FSActiveDirectoryResp DeleteFSActiveDirectory(ctx, fsActiveDirectoryId).Execute()





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
	fsActiveDirectoryId := int64(789) // int64 | file storage active directory id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsActiveDirectoriesAPI.DeleteFSActiveDirectory(context.Background(), fsActiveDirectoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsActiveDirectoriesAPI.DeleteFSActiveDirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFSActiveDirectory`: FSActiveDirectoryResp
	fmt.Fprintf(os.Stdout, "Response from `FsActiveDirectoriesAPI.DeleteFSActiveDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsActiveDirectoryId** | **int64** | file storage active directory id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFSActiveDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FSActiveDirectoryResp**](FSActiveDirectoryResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFSActiveDirectory

> FSActiveDirectoryResp GetFSActiveDirectory(ctx, fsActiveDirectoryId).Execute()





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
	fsActiveDirectoryId := int64(789) // int64 | file storage active directory id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsActiveDirectoriesAPI.GetFSActiveDirectory(context.Background(), fsActiveDirectoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsActiveDirectoriesAPI.GetFSActiveDirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFSActiveDirectory`: FSActiveDirectoryResp
	fmt.Fprintf(os.Stdout, "Response from `FsActiveDirectoriesAPI.GetFSActiveDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsActiveDirectoryId** | **int64** | file storage active directory id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFSActiveDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FSActiveDirectoryResp**](FSActiveDirectoryResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFSActiveDirectories

> FSActiveDirectoriesResp ListFSActiveDirectories(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsActiveDirectoriesAPI.ListFSActiveDirectories(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsActiveDirectoriesAPI.ListFSActiveDirectories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFSActiveDirectories`: FSActiveDirectoriesResp
	fmt.Fprintf(os.Stdout, "Response from `FsActiveDirectoriesAPI.ListFSActiveDirectories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFSActiveDirectoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 

### Return type

[**FSActiveDirectoriesResp**](FSActiveDirectoriesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFSActiveDirectory

> FSActiveDirectoryResp UpdateFSActiveDirectory(ctx, fsActiveDirectoryId).Body(body).Execute()





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
	fsActiveDirectoryId := int64(789) // int64 | file storage active directory id
	body := *openapiclient.NewFSActiveDirectoryUpdateReq(*openapiclient.NewFSActiveDirectoryUpdateReqInfo()) // FSActiveDirectoryUpdateReq | file storage active directory info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsActiveDirectoriesAPI.UpdateFSActiveDirectory(context.Background(), fsActiveDirectoryId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsActiveDirectoriesAPI.UpdateFSActiveDirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFSActiveDirectory`: FSActiveDirectoryResp
	fmt.Fprintf(os.Stdout, "Response from `FsActiveDirectoriesAPI.UpdateFSActiveDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsActiveDirectoryId** | **int64** | file storage active directory id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFSActiveDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FSActiveDirectoryUpdateReq**](FSActiveDirectoryUpdateReq.md) | file storage active directory info | 

### Return type

[**FSActiveDirectoryResp**](FSActiveDirectoryResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyFSActiveDirectory

> FSActiveDirectoryResp VerifyFSActiveDirectory(ctx, fsActiveDirectoryId).Execute()





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
	fsActiveDirectoryId := int64(789) // int64 | file storage active directory id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsActiveDirectoriesAPI.VerifyFSActiveDirectory(context.Background(), fsActiveDirectoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsActiveDirectoriesAPI.VerifyFSActiveDirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyFSActiveDirectory`: FSActiveDirectoryResp
	fmt.Fprintf(os.Stdout, "Response from `FsActiveDirectoriesAPI.VerifyFSActiveDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsActiveDirectoryId** | **int64** | file storage active directory id | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyFSActiveDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FSActiveDirectoryResp**](FSActiveDirectoryResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

