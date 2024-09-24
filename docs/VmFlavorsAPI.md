# \VmFlavorsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVMFlavor**](VmFlavorsAPI.md#GetVMFlavor) | **Get** /vm-flavors/{vm_flavor_id} | 
[**ListVMFlavors**](VmFlavorsAPI.md#ListVMFlavors) | **Get** /vm-flavors/ | 



## GetVMFlavor

> VMFlavorResp GetVMFlavor(ctx, vmFlavorId).Execute()





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
	vmFlavorId := int64(789) // int64 | vm flavor id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VmFlavorsAPI.GetVMFlavor(context.Background(), vmFlavorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VmFlavorsAPI.GetVMFlavor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMFlavor`: VMFlavorResp
	fmt.Fprintf(os.Stdout, "Response from `VmFlavorsAPI.GetVMFlavor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmFlavorId** | **int64** | vm flavor id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMFlavorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VMFlavorResp**](VMFlavorResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVMFlavors

> VMFlavorsResp ListVMFlavors(ctx).Limit(limit).Offset(offset).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VmFlavorsAPI.ListVMFlavors(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VmFlavorsAPI.ListVMFlavors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVMFlavors`: VMFlavorsResp
	fmt.Fprintf(os.Stdout, "Response from `VmFlavorsAPI.ListVMFlavors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVMFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**VMFlavorsResp**](VMFlavorsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

