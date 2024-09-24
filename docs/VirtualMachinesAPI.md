# \VirtualMachinesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVirtualMachine**](VirtualMachinesAPI.md#GetVirtualMachine) | **Get** /virtual-machines/{virtual_machine_id} | 
[**ListVirtualMachines**](VirtualMachinesAPI.md#ListVirtualMachines) | **Get** /virtual-machines/ | 
[**UpdateVirtualMachine**](VirtualMachinesAPI.md#UpdateVirtualMachine) | **Patch** /virtual-machines/{virtual_machine_id} | 



## GetVirtualMachine

> VirtualMachineResp GetVirtualMachine(ctx, virtualMachineId).Execute()





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
	virtualMachineId := int64(789) // int64 | virtual machine id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.GetVirtualMachine(context.Background(), virtualMachineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachine`: VirtualMachineResp
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | **int64** | virtual machine id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualMachineResp**](VirtualMachineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualMachines

> VirtualMachinesResp ListVirtualMachines(ctx).Limit(limit).Offset(offset).Execute()





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
	resp, r, err := apiClient.VirtualMachinesAPI.ListVirtualMachines(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.ListVirtualMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVirtualMachines`: VirtualMachinesResp
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.ListVirtualMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**VirtualMachinesResp**](VirtualMachinesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualMachine

> VirtualMachineResp UpdateVirtualMachine(ctx, virtualMachineId).CpuNum(cpuNum).Execute()





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
	virtualMachineId := int64(789) // int64 | virtual machine id
	cpuNum := int64(789) // int64 | new cpu number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.UpdateVirtualMachine(context.Background(), virtualMachineId).CpuNum(cpuNum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.UpdateVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVirtualMachine`: VirtualMachineResp
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.UpdateVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | **int64** | virtual machine id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cpuNum** | **int64** | new cpu number | 

### Return type

[**VirtualMachineResp**](VirtualMachineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

