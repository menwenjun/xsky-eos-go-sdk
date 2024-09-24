# \VipsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVIP**](VipsAPI.md#CreateVIP) | **Post** /vips/ | 
[**DeleteVIP**](VipsAPI.md#DeleteVIP) | **Delete** /vips/{vip_id} | 
[**GetVIP**](VipsAPI.md#GetVIP) | **Get** /vips/{vip_id} | 
[**ListVIPs**](VipsAPI.md#ListVIPs) | **Get** /vips/ | 
[**UpdateVIP**](VipsAPI.md#UpdateVIP) | **Patch** /vips/{vip_id} | 



## CreateVIP

> VIPResp CreateVIP(ctx).Vip(vip).Execute()





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
	vip := *openapiclient.NewVIPCreateReq(*openapiclient.NewVIPCreateReqVIP("Ip_example", int64(123), int64(123))) // VIPCreateReq | vip info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipsAPI.CreateVIP(context.Background()).Vip(vip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipsAPI.CreateVIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVIP`: VIPResp
	fmt.Fprintf(os.Stdout, "Response from `VipsAPI.CreateVIP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vip** | [**VIPCreateReq**](VIPCreateReq.md) | vip info | 

### Return type

[**VIPResp**](VIPResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVIP

> VIPResp DeleteVIP(ctx, vipId).Execute()





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
	vipId := int64(789) // int64 | vip id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipsAPI.DeleteVIP(context.Background(), vipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipsAPI.DeleteVIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVIP`: VIPResp
	fmt.Fprintf(os.Stdout, "Response from `VipsAPI.DeleteVIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vipId** | **int64** | vip id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VIPResp**](VIPResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVIP

> VIPResp GetVIP(ctx, vipId).Execute()





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
	vipId := int64(789) // int64 | vip id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipsAPI.GetVIP(context.Background(), vipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipsAPI.GetVIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVIP`: VIPResp
	fmt.Fprintf(os.Stdout, "Response from `VipsAPI.GetVIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vipId** | **int64** | vip id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VIPResp**](VIPResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVIPs

> VIPsResp ListVIPs(ctx).Limit(limit).Offset(offset).VipGroupId(vipGroupId).Execute()





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
	vipGroupId := int64(789) // int64 | vip_group id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipsAPI.ListVIPs(context.Background()).Limit(limit).Offset(offset).VipGroupId(vipGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipsAPI.ListVIPs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVIPs`: VIPsResp
	fmt.Fprintf(os.Stdout, "Response from `VipsAPI.ListVIPs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVIPsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **vipGroupId** | **int64** | vip_group id | 

### Return type

[**VIPsResp**](VIPsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVIP

> VIPResp UpdateVIP(ctx, vipId).Execute()





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
	vipId := int64(789) // int64 | vip id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipsAPI.UpdateVIP(context.Background(), vipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipsAPI.UpdateVIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVIP`: VIPResp
	fmt.Fprintf(os.Stdout, "Response from `VipsAPI.UpdateVIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vipId** | **int64** | vip id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VIPResp**](VIPResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

