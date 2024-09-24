# \CloudPlatformsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudPlatform**](CloudPlatformsAPI.md#CreateCloudPlatform) | **Post** /cloud-platforms/ | 
[**DeleteCloudPlatform**](CloudPlatformsAPI.md#DeleteCloudPlatform) | **Delete** /cloud-platforms/{cloud_platform_id} | 
[**GetCloudPlatform**](CloudPlatformsAPI.md#GetCloudPlatform) | **Get** /cloud-platforms/{cloud_platform_id} | 
[**ListCloudPlatforms**](CloudPlatformsAPI.md#ListCloudPlatforms) | **Get** /cloud-platforms/ | 
[**SyncCloudPlatform**](CloudPlatformsAPI.md#SyncCloudPlatform) | **Post** /cloud-platforms/{cloud_platform_id}:sync | 
[**UpdateCloudPlatform**](CloudPlatformsAPI.md#UpdateCloudPlatform) | **Patch** /cloud-platforms/{cloud_platform_id} | 



## CreateCloudPlatform

> CloudPlatformResp CreateCloudPlatform(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewCloudPlatformCreateReq() // CloudPlatformCreateReq | cloud platform info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudPlatformsAPI.CreateCloudPlatform(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudPlatformsAPI.CreateCloudPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCloudPlatform`: CloudPlatformResp
	fmt.Fprintf(os.Stdout, "Response from `CloudPlatformsAPI.CreateCloudPlatform`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CloudPlatformCreateReq**](CloudPlatformCreateReq.md) | cloud platform info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**CloudPlatformResp**](CloudPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCloudPlatform

> DeleteCloudPlatform(ctx, cloudPlatformId).Execute()





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
	cloudPlatformId := int64(789) // int64 | cloud platform id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudPlatformsAPI.DeleteCloudPlatform(context.Background(), cloudPlatformId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudPlatformsAPI.DeleteCloudPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudPlatformId** | **int64** | cloud platform id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudPlatform

> CloudPlatformResp GetCloudPlatform(ctx, cloudPlatformId).Execute()





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
	cloudPlatformId := int64(789) // int64 | cloud platform id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudPlatformsAPI.GetCloudPlatform(context.Background(), cloudPlatformId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudPlatformsAPI.GetCloudPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudPlatform`: CloudPlatformResp
	fmt.Fprintf(os.Stdout, "Response from `CloudPlatformsAPI.GetCloudPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudPlatformId** | **int64** | cloud platform id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudPlatformResp**](CloudPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudPlatforms

> CloudPlatformsResp ListCloudPlatforms(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()





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
	resp, r, err := apiClient.CloudPlatformsAPI.ListCloudPlatforms(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudPlatformsAPI.ListCloudPlatforms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCloudPlatforms`: CloudPlatformsResp
	fmt.Fprintf(os.Stdout, "Response from `CloudPlatformsAPI.ListCloudPlatforms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCloudPlatformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 

### Return type

[**CloudPlatformsResp**](CloudPlatformsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncCloudPlatform

> CloudPlatformResp SyncCloudPlatform(ctx, cloudPlatformId).Execute()





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
	cloudPlatformId := int64(789) // int64 | cloud platform id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudPlatformsAPI.SyncCloudPlatform(context.Background(), cloudPlatformId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudPlatformsAPI.SyncCloudPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncCloudPlatform`: CloudPlatformResp
	fmt.Fprintf(os.Stdout, "Response from `CloudPlatformsAPI.SyncCloudPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudPlatformId** | **int64** | cloud platform id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncCloudPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudPlatformResp**](CloudPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudPlatform

> CloudPlatformResp UpdateCloudPlatform(ctx, cloudPlatformId).Body(body).Execute()





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
	cloudPlatformId := int64(789) // int64 | cloud platform id
	body := *openapiclient.NewCloudPlatformUpdateReq() // CloudPlatformUpdateReq | cloud platform info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudPlatformsAPI.UpdateCloudPlatform(context.Background(), cloudPlatformId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudPlatformsAPI.UpdateCloudPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCloudPlatform`: CloudPlatformResp
	fmt.Fprintf(os.Stdout, "Response from `CloudPlatformsAPI.UpdateCloudPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudPlatformId** | **int64** | cloud platform id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CloudPlatformUpdateReq**](CloudPlatformUpdateReq.md) | cloud platform info | 

### Return type

[**CloudPlatformResp**](CloudPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

