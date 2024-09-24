# \AlertWechatsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlertWechatConfig**](AlertWechatsAPI.md#GetAlertWechatConfig) | **Get** /alert-wechats/config | 
[**SendWechatMessage**](AlertWechatsAPI.md#SendWechatMessage) | **Post** /alert-wechats/ | 
[**UpdateAlertWechatConfig**](AlertWechatsAPI.md#UpdateAlertWechatConfig) | **Patch** /alert-wechats/config | 



## GetAlertWechatConfig

> AlertWechatConfigResp GetAlertWechatConfig(ctx).Execute()





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
	resp, r, err := apiClient.AlertWechatsAPI.GetAlertWechatConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertWechatsAPI.GetAlertWechatConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertWechatConfig`: AlertWechatConfigResp
	fmt.Fprintf(os.Stdout, "Response from `AlertWechatsAPI.GetAlertWechatConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertWechatConfigRequest struct via the builder pattern


### Return type

[**AlertWechatConfigResp**](AlertWechatConfigResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendWechatMessage

> AlertWechatSendResp SendWechatMessage(ctx).Body(body).Execute()





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
	body := *openapiclient.NewAlertWechatSendReq(*openapiclient.NewAlertWechatSendReqAlertWechatSend()) // AlertWechatSendReq | wechat info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertWechatsAPI.SendWechatMessage(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertWechatsAPI.SendWechatMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendWechatMessage`: AlertWechatSendResp
	fmt.Fprintf(os.Stdout, "Response from `AlertWechatsAPI.SendWechatMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendWechatMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AlertWechatSendReq**](AlertWechatSendReq.md) | wechat info | 

### Return type

[**AlertWechatSendResp**](AlertWechatSendResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertWechatConfig

> AlertWechatConfigResp UpdateAlertWechatConfig(ctx).Body(body).Execute()





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
	body := *openapiclient.NewAlertWechatConfigUpdateReq(*openapiclient.NewAlertWechatConfigUpdateReqAlertWechatConfig()) // AlertWechatConfigUpdateReq | alert wechat config info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertWechatsAPI.UpdateAlertWechatConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertWechatsAPI.UpdateAlertWechatConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlertWechatConfig`: AlertWechatConfigResp
	fmt.Fprintf(os.Stdout, "Response from `AlertWechatsAPI.UpdateAlertWechatConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertWechatConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AlertWechatConfigUpdateReq**](AlertWechatConfigUpdateReq.md) | alert wechat config info | 

### Return type

[**AlertWechatConfigResp**](AlertWechatConfigResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

