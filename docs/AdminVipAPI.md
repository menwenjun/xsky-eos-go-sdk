# \AdminVipAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateAdminVIP**](AdminVipAPI.md#CreateOrUpdateAdminVIP) | **Put** /admin-vip/ | 
[**FlushAdminVIP**](AdminVipAPI.md#FlushAdminVIP) | **Patch** /admin-vip/ | 
[**GetAdminVIP**](AdminVipAPI.md#GetAdminVIP) | **Get** /admin-vip/ | 
[**GetAdminVIPNetwork**](AdminVipAPI.md#GetAdminVIPNetwork) | **Get** /admin-vip/networks | 



## CreateOrUpdateAdminVIP

> AdminVIPResp CreateOrUpdateAdminVIP(ctx).AdminVip(adminVip).Execute()





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
	adminVip := *openapiclient.NewAdminVIPCreateReq(*openapiclient.NewAdminVIPCreateReqAdminVip(false)) // AdminVIPCreateReq | admin vip info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminVipAPI.CreateOrUpdateAdminVIP(context.Background()).AdminVip(adminVip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminVipAPI.CreateOrUpdateAdminVIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrUpdateAdminVIP`: AdminVIPResp
	fmt.Fprintf(os.Stdout, "Response from `AdminVipAPI.CreateOrUpdateAdminVIP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateAdminVIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminVip** | [**AdminVIPCreateReq**](AdminVIPCreateReq.md) | admin vip info | 

### Return type

[**AdminVIPResp**](AdminVIPResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlushAdminVIP

> AdminVIPResp FlushAdminVIP(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminVipAPI.FlushAdminVIP(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminVipAPI.FlushAdminVIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlushAdminVIP`: AdminVIPResp
	fmt.Fprintf(os.Stdout, "Response from `AdminVipAPI.FlushAdminVIP`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlushAdminVIPRequest struct via the builder pattern


### Return type

[**AdminVIPResp**](AdminVIPResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminVIP

> AdminVIPResp GetAdminVIP(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminVipAPI.GetAdminVIP(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminVipAPI.GetAdminVIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminVIP`: AdminVIPResp
	fmt.Fprintf(os.Stdout, "Response from `AdminVipAPI.GetAdminVIP`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminVIPRequest struct via the builder pattern


### Return type

[**AdminVIPResp**](AdminVIPResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminVIPNetwork

> AdminVIPNetworkResp GetAdminVIPNetwork(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminVipAPI.GetAdminVIPNetwork(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminVipAPI.GetAdminVIPNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminVIPNetwork`: AdminVIPNetworkResp
	fmt.Fprintf(os.Stdout, "Response from `AdminVipAPI.GetAdminVIPNetwork`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminVIPNetworkRequest struct via the builder pattern


### Return type

[**AdminVIPNetworkResp**](AdminVIPNetworkResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

