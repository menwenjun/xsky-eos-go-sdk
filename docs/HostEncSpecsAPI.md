# \HostEncSpecsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHostEncSpec**](HostEncSpecsAPI.md#CreateHostEncSpec) | **Post** /host-enc-specs/ | 
[**DeleteHostEncSpec**](HostEncSpecsAPI.md#DeleteHostEncSpec) | **Delete** /host-enc-specs/{spec_id} | 
[**GetHostEncSpec**](HostEncSpecsAPI.md#GetHostEncSpec) | **Get** /host-enc-specs/{spec_id} | 
[**ListHostEncSpecs**](HostEncSpecsAPI.md#ListHostEncSpecs) | **Get** /host-enc-specs/ | 



## CreateHostEncSpec

> HostEncSpecResp CreateHostEncSpec(ctx).Body(body).Execute()





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
	body := *openapiclient.NewHostEncSpecCreateReq() // HostEncSpecCreateReq | host enclosure spec info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostEncSpecsAPI.CreateHostEncSpec(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostEncSpecsAPI.CreateHostEncSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHostEncSpec`: HostEncSpecResp
	fmt.Fprintf(os.Stdout, "Response from `HostEncSpecsAPI.CreateHostEncSpec`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostEncSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HostEncSpecCreateReq**](HostEncSpecCreateReq.md) | host enclosure spec info | 

### Return type

[**HostEncSpecResp**](HostEncSpecResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHostEncSpec

> DeleteHostEncSpec(ctx, specId).Execute()





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
	specId := int64(789) // int64 | host enclosure spec id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostEncSpecsAPI.DeleteHostEncSpec(context.Background(), specId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostEncSpecsAPI.DeleteHostEncSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**specId** | **int64** | host enclosure spec id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHostEncSpecRequest struct via the builder pattern


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


## GetHostEncSpec

> HostEncSpecResp GetHostEncSpec(ctx, specId).Execute()





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
	specId := int64(789) // int64 | host enclosure spec id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostEncSpecsAPI.GetHostEncSpec(context.Background(), specId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostEncSpecsAPI.GetHostEncSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostEncSpec`: HostEncSpecResp
	fmt.Fprintf(os.Stdout, "Response from `HostEncSpecsAPI.GetHostEncSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**specId** | **int64** | host enclosure spec id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostEncSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostEncSpecResp**](HostEncSpecResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHostEncSpecs

> HostEncSpecsResp ListHostEncSpecs(ctx).Limit(limit).Offset(offset).Model(model).Vendor(vendor).Execute()





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
	model := "model_example" // string | host model (optional)
	vendor := "vendor_example" // string | host vendor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostEncSpecsAPI.ListHostEncSpecs(context.Background()).Limit(limit).Offset(offset).Model(model).Vendor(vendor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostEncSpecsAPI.ListHostEncSpecs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHostEncSpecs`: HostEncSpecsResp
	fmt.Fprintf(os.Stdout, "Response from `HostEncSpecsAPI.ListHostEncSpecs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHostEncSpecsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **model** | **string** | host model | 
 **vendor** | **string** | host vendor | 

### Return type

[**HostEncSpecsResp**](HostEncSpecsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

