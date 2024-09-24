# \OsStorageClassesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOSStorageClass**](OsStorageClassesAPI.md#CreateOSStorageClass) | **Post** /os-storage-classes/ | 
[**DeleteOSStorageClass**](OsStorageClassesAPI.md#DeleteOSStorageClass) | **Delete** /os-storage-classes/{storage_class_id} | 
[**GetOSStorageClass**](OsStorageClassesAPI.md#GetOSStorageClass) | **Get** /os-storage-classes/{storage_class_id} | 
[**ListOSStorageClasses**](OsStorageClassesAPI.md#ListOSStorageClasses) | **Get** /os-storage-classes/ | 
[**UpdateOSStorageClass**](OsStorageClassesAPI.md#UpdateOSStorageClass) | **Patch** /os-storage-classes/{storage_class_id} | 



## CreateOSStorageClass

> OSStorageClassResp CreateOSStorageClass(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewOSStorageClassCreateReq() // OSStorageClassCreateReq | storage class info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsStorageClassesAPI.CreateOSStorageClass(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsStorageClassesAPI.CreateOSStorageClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOSStorageClass`: OSStorageClassResp
	fmt.Fprintf(os.Stdout, "Response from `OsStorageClassesAPI.CreateOSStorageClass`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOSStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OSStorageClassCreateReq**](OSStorageClassCreateReq.md) | storage class info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSStorageClassResp**](OSStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOSStorageClass

> OSStorageClassResp DeleteOSStorageClass(ctx, storageClassId).Execute()





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
	storageClassId := int64(789) // int64 | storage class id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsStorageClassesAPI.DeleteOSStorageClass(context.Background(), storageClassId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsStorageClassesAPI.DeleteOSStorageClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOSStorageClass`: OSStorageClassResp
	fmt.Fprintf(os.Stdout, "Response from `OsStorageClassesAPI.DeleteOSStorageClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageClassId** | **int64** | storage class id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOSStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSStorageClassResp**](OSStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSStorageClass

> OSStorageClassResp GetOSStorageClass(ctx, storageClassId).Execute()





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
	storageClassId := "storageClassId_example" // string | storage class id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsStorageClassesAPI.GetOSStorageClass(context.Background(), storageClassId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsStorageClassesAPI.GetOSStorageClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSStorageClass`: OSStorageClassResp
	fmt.Fprintf(os.Stdout, "Response from `OsStorageClassesAPI.GetOSStorageClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageClassId** | **string** | storage class id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSStorageClassResp**](OSStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOSStorageClasses

> OSStorageClassesResp ListOSStorageClasses(ctx).Limit(limit).Offset(offset).OsPolicyId(osPolicyId).ClusterId(clusterId).Execute()





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
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	osPolicyId := int64(789) // int64 | os policy id (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsStorageClassesAPI.ListOSStorageClasses(context.Background()).Limit(limit).Offset(offset).OsPolicyId(osPolicyId).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsStorageClassesAPI.ListOSStorageClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSStorageClasses`: OSStorageClassesResp
	fmt.Fprintf(os.Stdout, "Response from `OsStorageClassesAPI.ListOSStorageClasses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSStorageClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **osPolicyId** | **int64** | os policy id | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSStorageClassesResp**](OSStorageClassesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOSStorageClass

> OSStorageClassResp UpdateOSStorageClass(ctx, storageClassId).Body(body).Execute()





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
	storageClassId := int64(789) // int64 | storage class id
	body := *openapiclient.NewOSStorageClassUpdateReq() // OSStorageClassUpdateReq | storage class info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsStorageClassesAPI.UpdateOSStorageClass(context.Background(), storageClassId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsStorageClassesAPI.UpdateOSStorageClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOSStorageClass`: OSStorageClassResp
	fmt.Fprintf(os.Stdout, "Response from `OsStorageClassesAPI.UpdateOSStorageClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageClassId** | **int64** | storage class id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOSStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSStorageClassUpdateReq**](OSStorageClassUpdateReq.md) | storage class info | 

### Return type

[**OSStorageClassResp**](OSStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

