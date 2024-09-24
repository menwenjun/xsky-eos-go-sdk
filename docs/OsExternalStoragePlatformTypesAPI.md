# \OsExternalStoragePlatformTypesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOSExternalStoragePlatformType**](OsExternalStoragePlatformTypesAPI.md#CreateOSExternalStoragePlatformType) | **Post** /os-external-storage-platform-types/ | 
[**GetOSExternalStoragePlatformType**](OsExternalStoragePlatformTypesAPI.md#GetOSExternalStoragePlatformType) | **Get** /os-external-storage-platform-types/{external_storage_platform_type_id} | 
[**ListOSExternalStoragePlatformTypes**](OsExternalStoragePlatformTypesAPI.md#ListOSExternalStoragePlatformTypes) | **Get** /os-external-storage-platform-types/ | 



## CreateOSExternalStoragePlatformType

> OSExternalStoragePlatformTypeResp CreateOSExternalStoragePlatformType(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewOSExternalStoragePlatformTypeCreateReq() // OSExternalStoragePlatformTypeCreateReq | external storage platform type info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsExternalStoragePlatformTypesAPI.CreateOSExternalStoragePlatformType(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsExternalStoragePlatformTypesAPI.CreateOSExternalStoragePlatformType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOSExternalStoragePlatformType`: OSExternalStoragePlatformTypeResp
	fmt.Fprintf(os.Stdout, "Response from `OsExternalStoragePlatformTypesAPI.CreateOSExternalStoragePlatformType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOSExternalStoragePlatformTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OSExternalStoragePlatformTypeCreateReq**](OSExternalStoragePlatformTypeCreateReq.md) | external storage platform type info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSExternalStoragePlatformTypeResp**](OSExternalStoragePlatformTypeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSExternalStoragePlatformType

> OSExternalStoragePlatformTypeResp GetOSExternalStoragePlatformType(ctx, externalStoragePlatformTypeId).Execute()





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
	externalStoragePlatformTypeId := int64(789) // int64 | external storage platform type id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsExternalStoragePlatformTypesAPI.GetOSExternalStoragePlatformType(context.Background(), externalStoragePlatformTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsExternalStoragePlatformTypesAPI.GetOSExternalStoragePlatformType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSExternalStoragePlatformType`: OSExternalStoragePlatformTypeResp
	fmt.Fprintf(os.Stdout, "Response from `OsExternalStoragePlatformTypesAPI.GetOSExternalStoragePlatformType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalStoragePlatformTypeId** | **int64** | external storage platform type id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSExternalStoragePlatformTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSExternalStoragePlatformTypeResp**](OSExternalStoragePlatformTypeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOSExternalStoragePlatformTypes

> OSExternalStoragePlatformTypesResp ListOSExternalStoragePlatformTypes(ctx).Limit(limit).Offset(offset).Platform(platform).PlatformType(platformType).ClusterId(clusterId).Execute()





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
	platform := "platform_example" // string | platform (optional)
	platformType := "platformType_example" // string | platform type (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsExternalStoragePlatformTypesAPI.ListOSExternalStoragePlatformTypes(context.Background()).Limit(limit).Offset(offset).Platform(platform).PlatformType(platformType).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsExternalStoragePlatformTypesAPI.ListOSExternalStoragePlatformTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSExternalStoragePlatformTypes`: OSExternalStoragePlatformTypesResp
	fmt.Fprintf(os.Stdout, "Response from `OsExternalStoragePlatformTypesAPI.ListOSExternalStoragePlatformTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSExternalStoragePlatformTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **platform** | **string** | platform | 
 **platformType** | **string** | platform type | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSExternalStoragePlatformTypesResp**](OSExternalStoragePlatformTypesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

