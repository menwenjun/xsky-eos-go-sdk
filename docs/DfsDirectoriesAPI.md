# \DfsDirectoriesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckDirectoryResources**](DfsDirectoriesAPI.md#CheckDirectoryResources) | **Post** /dfs-directories/:check-resource | 
[**CreateDfsCSIDirectory**](DfsDirectoriesAPI.md#CreateDfsCSIDirectory) | **Post** /dfs-directories/:mkcsidir | 
[**CreateDfsDirectories**](DfsDirectoriesAPI.md#CreateDfsDirectories) | **Post** /dfs-directories/:create | 
[**CreateDfsDirectory**](DfsDirectoriesAPI.md#CreateDfsDirectory) | **Post** /dfs-directories/:mkdir | 
[**DeleteDfsDirectories**](DfsDirectoriesAPI.md#DeleteDfsDirectories) | **Post** /dfs-directories/:delete | 
[**DeleteDfsDirectory**](DfsDirectoriesAPI.md#DeleteDfsDirectory) | **Post** /dfs-directories/:rmdir | 
[**DirectoryValidator**](DfsDirectoriesAPI.md#DirectoryValidator) | **Get** /dfs-directories/directory-validator | 
[**UpdatePolicyOnDfsDirectory**](DfsDirectoriesAPI.md#UpdatePolicyOnDfsDirectory) | **Patch** /dfs-directories/:update-storage-policy | 



## CheckDirectoryResources

> DfsDirectoriesCheckResourcesResp CheckDirectoryResources(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsDirectoriesCheckResourcesReq(*openapiclient.NewDfsDirectoriesCheckResourcesReqDirectories(int64(123), []string{"Directories_example"})) // DfsDirectoriesCheckResourcesReq | directory info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsDirectoriesAPI.CheckDirectoryResources(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsDirectoriesAPI.CheckDirectoryResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckDirectoryResources`: DfsDirectoriesCheckResourcesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsDirectoriesAPI.CheckDirectoryResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckDirectoryResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsDirectoriesCheckResourcesReq**](DfsDirectoriesCheckResourcesReq.md) | directory info | 

### Return type

[**DfsDirectoriesCheckResourcesResp**](DfsDirectoriesCheckResourcesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDfsCSIDirectory

> DfsCSIDirectoryResp CreateDfsCSIDirectory(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsCSIDirectoryReq(*openapiclient.NewDfsCSIDirectoryReqDirectory(int64(123), "Path_example")) // DfsCSIDirectoryReq | directory info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsDirectoriesAPI.CreateDfsCSIDirectory(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsDirectoriesAPI.CreateDfsCSIDirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsCSIDirectory`: DfsCSIDirectoryResp
	fmt.Fprintf(os.Stdout, "Response from `DfsDirectoriesAPI.CreateDfsCSIDirectory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsCSIDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsCSIDirectoryReq**](DfsCSIDirectoryReq.md) | directory info | 

### Return type

[**DfsCSIDirectoryResp**](DfsCSIDirectoryResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDfsDirectories

> DfsDirectoriesResp CreateDfsDirectories(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsDirectoriesReq(*openapiclient.NewDfsDirectoriesReqDirectories(int64(123), []string{"Directories_example"})) // DfsDirectoriesReq | directories info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsDirectoriesAPI.CreateDfsDirectories(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsDirectoriesAPI.CreateDfsDirectories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsDirectories`: DfsDirectoriesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsDirectoriesAPI.CreateDfsDirectories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsDirectoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsDirectoriesReq**](DfsDirectoriesReq.md) | directories info | 

### Return type

[**DfsDirectoriesResp**](DfsDirectoriesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDfsDirectory

> DfsDirectoryResp CreateDfsDirectory(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsDirectoryReq(*openapiclient.NewDfsDirectoryReqDirectory(int64(123), "Path_example")) // DfsDirectoryReq | directory info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsDirectoriesAPI.CreateDfsDirectory(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsDirectoriesAPI.CreateDfsDirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsDirectory`: DfsDirectoryResp
	fmt.Fprintf(os.Stdout, "Response from `DfsDirectoriesAPI.CreateDfsDirectory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsDirectoryReq**](DfsDirectoryReq.md) | directory info | 

### Return type

[**DfsDirectoryResp**](DfsDirectoryResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsDirectories

> DfsDirectoriesResp DeleteDfsDirectories(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsDirectoriesReq(*openapiclient.NewDfsDirectoriesReqDirectories(int64(123), []string{"Directories_example"})) // DfsDirectoriesReq | directories info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsDirectoriesAPI.DeleteDfsDirectories(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsDirectoriesAPI.DeleteDfsDirectories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsDirectories`: DfsDirectoriesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsDirectoriesAPI.DeleteDfsDirectories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsDirectoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsDirectoriesReq**](DfsDirectoriesReq.md) | directories info | 

### Return type

[**DfsDirectoriesResp**](DfsDirectoriesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsDirectory

> DfsDirectoryResp DeleteDfsDirectory(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsDirectoryReq(*openapiclient.NewDfsDirectoryReqDirectory(int64(123), "Path_example")) // DfsDirectoryReq | directory info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsDirectoriesAPI.DeleteDfsDirectory(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsDirectoriesAPI.DeleteDfsDirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsDirectory`: DfsDirectoryResp
	fmt.Fprintf(os.Stdout, "Response from `DfsDirectoriesAPI.DeleteDfsDirectory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsDirectoryReq**](DfsDirectoryReq.md) | directory info | 

### Return type

[**DfsDirectoryResp**](DfsDirectoryResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryValidator

> DfsDirectoryValidationResp DirectoryValidator(ctx).RootfsId(rootfsId).Path(path).Execute()





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
	rootfsId := int64(789) // int64 | dfs rootfs id
	path := "path_example" // string | dfs quota path

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsDirectoriesAPI.DirectoryValidator(context.Background()).RootfsId(rootfsId).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsDirectoriesAPI.DirectoryValidator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DirectoryValidator`: DfsDirectoryValidationResp
	fmt.Fprintf(os.Stdout, "Response from `DfsDirectoriesAPI.DirectoryValidator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rootfsId** | **int64** | dfs rootfs id | 
 **path** | **string** | dfs quota path | 

### Return type

[**DfsDirectoryValidationResp**](DfsDirectoryValidationResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicyOnDfsDirectory

> DfsDirectoryResp UpdatePolicyOnDfsDirectory(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsDirectoryReq(*openapiclient.NewDfsDirectoryReqDirectory(int64(123), "Path_example")) // DfsDirectoryReq | policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsDirectoriesAPI.UpdatePolicyOnDfsDirectory(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsDirectoriesAPI.UpdatePolicyOnDfsDirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePolicyOnDfsDirectory`: DfsDirectoryResp
	fmt.Fprintf(os.Stdout, "Response from `DfsDirectoriesAPI.UpdatePolicyOnDfsDirectory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyOnDfsDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsDirectoryReq**](DfsDirectoryReq.md) | policy info | 

### Return type

[**DfsDirectoryResp**](DfsDirectoryResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

