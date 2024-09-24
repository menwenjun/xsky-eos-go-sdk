# \VipInstancesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVIPInstance**](VipInstancesAPI.md#GetVIPInstance) | **Get** /vip-instances/{vip_instance_id} | 
[**ListVIPInstances**](VipInstancesAPI.md#ListVIPInstances) | **Get** /vip-instances/ | 



## GetVIPInstance

> VIPInstanceResp GetVIPInstance(ctx, vipInstanceId).Execute()





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
	vipInstanceId := int64(789) // int64 | vip group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipInstancesAPI.GetVIPInstance(context.Background(), vipInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipInstancesAPI.GetVIPInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVIPInstance`: VIPInstanceResp
	fmt.Fprintf(os.Stdout, "Response from `VipInstancesAPI.GetVIPInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vipInstanceId** | **int64** | vip group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVIPInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VIPInstanceResp**](VIPInstanceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVIPInstances

> VIPInstancesResp ListVIPInstances(ctx).Limit(limit).Offset(offset).VipId(vipId).Execute()





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
	vipId := int64(789) // int64 | vip id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipInstancesAPI.ListVIPInstances(context.Background()).Limit(limit).Offset(offset).VipId(vipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipInstancesAPI.ListVIPInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVIPInstances`: VIPInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `VipInstancesAPI.ListVIPInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVIPInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **vipId** | **int64** | vip id | 

### Return type

[**VIPInstancesResp**](VIPInstancesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

