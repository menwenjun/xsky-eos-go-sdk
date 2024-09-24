# \OsBuiltinBluRaysAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOSBuiltinBluRay**](OsBuiltinBluRaysAPI.md#CreateOSBuiltinBluRay) | **Post** /os-builtin-blu-rays/ | 
[**DeleteOSBuiltinBluRay**](OsBuiltinBluRaysAPI.md#DeleteOSBuiltinBluRay) | **Delete** /os-builtin-blu-rays/{builtin_blu_ray_id} | 
[**GetOSBuiltinBluRay**](OsBuiltinBluRaysAPI.md#GetOSBuiltinBluRay) | **Get** /os-builtin-blu-rays/{builtin_blu_ray_id} | 
[**ListOSBuiltinBluRays**](OsBuiltinBluRaysAPI.md#ListOSBuiltinBluRays) | **Get** /os-builtin-blu-rays/ | 
[**UpdateOSBuiltinBluRay**](OsBuiltinBluRaysAPI.md#UpdateOSBuiltinBluRay) | **Patch** /os-builtin-blu-rays/{builtin_blu_ray_id} | 



## CreateOSBuiltinBluRay

> OSBuiltinBluRayResp CreateOSBuiltinBluRay(ctx).Body(body).Execute()





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
	body := *openapiclient.NewOSBuiltinBluRayCreateReq() // OSBuiltinBluRayCreateReq | builtin blu ray info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBuiltinBluRaysAPI.CreateOSBuiltinBluRay(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBuiltinBluRaysAPI.CreateOSBuiltinBluRay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOSBuiltinBluRay`: OSBuiltinBluRayResp
	fmt.Fprintf(os.Stdout, "Response from `OsBuiltinBluRaysAPI.CreateOSBuiltinBluRay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOSBuiltinBluRayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OSBuiltinBluRayCreateReq**](OSBuiltinBluRayCreateReq.md) | builtin blu ray info | 

### Return type

[**OSBuiltinBluRayResp**](OSBuiltinBluRayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOSBuiltinBluRay

> DeleteOSBuiltinBluRay(ctx, builtinBluRayId).Execute()





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
	builtinBluRayId := int64(789) // int64 | builtin blu ray id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OsBuiltinBluRaysAPI.DeleteOSBuiltinBluRay(context.Background(), builtinBluRayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBuiltinBluRaysAPI.DeleteOSBuiltinBluRay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**builtinBluRayId** | **int64** | builtin blu ray id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOSBuiltinBluRayRequest struct via the builder pattern


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


## GetOSBuiltinBluRay

> OSBuiltinBluRayResp GetOSBuiltinBluRay(ctx, builtinBluRayId).Execute()





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
	builtinBluRayId := int64(789) // int64 | builtin blu ray id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBuiltinBluRaysAPI.GetOSBuiltinBluRay(context.Background(), builtinBluRayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBuiltinBluRaysAPI.GetOSBuiltinBluRay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSBuiltinBluRay`: OSBuiltinBluRayResp
	fmt.Fprintf(os.Stdout, "Response from `OsBuiltinBluRaysAPI.GetOSBuiltinBluRay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**builtinBluRayId** | **int64** | builtin blu ray id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSBuiltinBluRayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSBuiltinBluRayResp**](OSBuiltinBluRayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOSBuiltinBluRays

> OSBuiltinBluRaysResp ListOSBuiltinBluRays(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBuiltinBluRaysAPI.ListOSBuiltinBluRays(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBuiltinBluRaysAPI.ListOSBuiltinBluRays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSBuiltinBluRays`: OSBuiltinBluRaysResp
	fmt.Fprintf(os.Stdout, "Response from `OsBuiltinBluRaysAPI.ListOSBuiltinBluRays`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSBuiltinBluRaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**OSBuiltinBluRaysResp**](OSBuiltinBluRaysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOSBuiltinBluRay

> OSBuiltinBluRayResp UpdateOSBuiltinBluRay(ctx, builtinBluRayId).Body(body).Execute()





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
	builtinBluRayId := int64(789) // int64 | builtin blu ray id
	body := *openapiclient.NewOSBuiltinBluRayUpdateReq() // OSBuiltinBluRayUpdateReq | builtin blu ray info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBuiltinBluRaysAPI.UpdateOSBuiltinBluRay(context.Background(), builtinBluRayId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBuiltinBluRaysAPI.UpdateOSBuiltinBluRay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOSBuiltinBluRay`: OSBuiltinBluRayResp
	fmt.Fprintf(os.Stdout, "Response from `OsBuiltinBluRaysAPI.UpdateOSBuiltinBluRay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**builtinBluRayId** | **int64** | builtin blu ray id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOSBuiltinBluRayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSBuiltinBluRayUpdateReq**](OSBuiltinBluRayUpdateReq.md) | builtin blu ray info | 

### Return type

[**OSBuiltinBluRayResp**](OSBuiltinBluRayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

