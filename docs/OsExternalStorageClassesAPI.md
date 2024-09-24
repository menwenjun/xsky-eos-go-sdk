# \OsExternalStorageClassesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalStorageClass**](OsExternalStorageClassesAPI.md#CreateExternalStorageClass) | **Post** /os-external-storage-classes/ | 
[**DeleteOSExternalStorageClass**](OsExternalStorageClassesAPI.md#DeleteOSExternalStorageClass) | **Delete** /os-external-storage-classes/{external_storage_class_id} | 
[**GetOSExternalStorageClass**](OsExternalStorageClassesAPI.md#GetOSExternalStorageClass) | **Get** /os-external-storage-classes/{external_storage_class_id} | 
[**ListOSExternalStorageClasses**](OsExternalStorageClassesAPI.md#ListOSExternalStorageClasses) | **Get** /os-external-storage-classes/ | 
[**UpdateOSExternalStorageClass**](OsExternalStorageClassesAPI.md#UpdateOSExternalStorageClass) | **Patch** /os-external-storage-classes/{external_storage_class_id} | 



## CreateExternalStorageClass

> OSExternalStorageClassResp CreateExternalStorageClass(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewOSExternalStorageClassCreateReq() // OSExternalStorageClassCreateReq | external storage class info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsExternalStorageClassesAPI.CreateExternalStorageClass(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsExternalStorageClassesAPI.CreateExternalStorageClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalStorageClass`: OSExternalStorageClassResp
	fmt.Fprintf(os.Stdout, "Response from `OsExternalStorageClassesAPI.CreateExternalStorageClass`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OSExternalStorageClassCreateReq**](OSExternalStorageClassCreateReq.md) | external storage class info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSExternalStorageClassResp**](OSExternalStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOSExternalStorageClass

> OSExternalStorageClassResp DeleteOSExternalStorageClass(ctx, externalStorageClassId).ClusterId(clusterId).Execute()





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
	externalStorageClassId := int64(789) // int64 | external storage class id
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsExternalStorageClassesAPI.DeleteOSExternalStorageClass(context.Background(), externalStorageClassId).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsExternalStorageClassesAPI.DeleteOSExternalStorageClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOSExternalStorageClass`: OSExternalStorageClassResp
	fmt.Fprintf(os.Stdout, "Response from `OsExternalStorageClassesAPI.DeleteOSExternalStorageClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalStorageClassId** | **int64** | external storage class id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOSExternalStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterId** | **string** | cluster id | 

### Return type

[**OSExternalStorageClassResp**](OSExternalStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSExternalStorageClass

> OSExternalStorageClassResp GetOSExternalStorageClass(ctx, externalStorageClassId).Execute()





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
	externalStorageClassId := int64(789) // int64 | external storage class id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsExternalStorageClassesAPI.GetOSExternalStorageClass(context.Background(), externalStorageClassId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsExternalStorageClassesAPI.GetOSExternalStorageClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSExternalStorageClass`: OSExternalStorageClassResp
	fmt.Fprintf(os.Stdout, "Response from `OsExternalStorageClassesAPI.GetOSExternalStorageClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalStorageClassId** | **int64** | external storage class id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSExternalStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSExternalStorageClassResp**](OSExternalStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOSExternalStorageClasses

> OSExternalStorageClassesResp ListOSExternalStorageClasses(ctx).Limit(limit).Offset(offset).OsPolicyId(osPolicyId).GeneralStatus(generalStatus).ClusterId(clusterId).Execute()





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
	generalStatus := true // bool | query records with active or error status (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsExternalStorageClassesAPI.ListOSExternalStorageClasses(context.Background()).Limit(limit).Offset(offset).OsPolicyId(osPolicyId).GeneralStatus(generalStatus).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsExternalStorageClassesAPI.ListOSExternalStorageClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSExternalStorageClasses`: OSExternalStorageClassesResp
	fmt.Fprintf(os.Stdout, "Response from `OsExternalStorageClassesAPI.ListOSExternalStorageClasses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSExternalStorageClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **osPolicyId** | **int64** | os policy id | 
 **generalStatus** | **bool** | query records with active or error status | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSExternalStorageClassesResp**](OSExternalStorageClassesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOSExternalStorageClass

> OSExternalStorageClassResp UpdateOSExternalStorageClass(ctx, externalStorageClassId).Body(body).ClusterId(clusterId).Execute()





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
	externalStorageClassId := int64(789) // int64 | external storage class id
	body := *openapiclient.NewOSExternalStorageClassUpdateReq() // OSExternalStorageClassUpdateReq | external storage class info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsExternalStorageClassesAPI.UpdateOSExternalStorageClass(context.Background(), externalStorageClassId).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsExternalStorageClassesAPI.UpdateOSExternalStorageClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOSExternalStorageClass`: OSExternalStorageClassResp
	fmt.Fprintf(os.Stdout, "Response from `OsExternalStorageClassesAPI.UpdateOSExternalStorageClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalStorageClassId** | **int64** | external storage class id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOSExternalStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSExternalStorageClassUpdateReq**](OSExternalStorageClassUpdateReq.md) | external storage class info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSExternalStorageClassResp**](OSExternalStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

