# \AlertDingdingsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlertDingDingConfig**](AlertDingdingsAPI.md#GetAlertDingDingConfig) | **Get** /alert-dingdings/config | 
[**SendDingDingMessage**](AlertDingdingsAPI.md#SendDingDingMessage) | **Post** /alert-dingdings/ | 
[**UpdateAlertDingDingConfig**](AlertDingdingsAPI.md#UpdateAlertDingDingConfig) | **Patch** /alert-dingdings/config | 



## GetAlertDingDingConfig

> AlertDingDingConfigResp GetAlertDingDingConfig(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertDingdingsAPI.GetAlertDingDingConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertDingdingsAPI.GetAlertDingDingConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertDingDingConfig`: AlertDingDingConfigResp
	fmt.Fprintf(os.Stdout, "Response from `AlertDingdingsAPI.GetAlertDingDingConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertDingDingConfigRequest struct via the builder pattern


### Return type

[**AlertDingDingConfigResp**](AlertDingDingConfigResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendDingDingMessage

> AlertDingDingSendResp SendDingDingMessage(ctx).Body(body).Execute()





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
	body := *openapiclient.NewAlertDingDingSendReq("Body_example") // AlertDingDingSendReq | dingding info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertDingdingsAPI.SendDingDingMessage(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertDingdingsAPI.SendDingDingMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendDingDingMessage`: AlertDingDingSendResp
	fmt.Fprintf(os.Stdout, "Response from `AlertDingdingsAPI.SendDingDingMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendDingDingMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AlertDingDingSendReq**](AlertDingDingSendReq.md) | dingding info | 

### Return type

[**AlertDingDingSendResp**](AlertDingDingSendResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertDingDingConfig

> AlertDingDingConfigResp UpdateAlertDingDingConfig(ctx).Body(body).Execute()





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
	body := *openapiclient.NewAlertDingDingConfigUpdateReq() // AlertDingDingConfigUpdateReq | alert dingding config info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertDingdingsAPI.UpdateAlertDingDingConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertDingdingsAPI.UpdateAlertDingDingConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlertDingDingConfig`: AlertDingDingConfigResp
	fmt.Fprintf(os.Stdout, "Response from `AlertDingdingsAPI.UpdateAlertDingDingConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertDingDingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AlertDingDingConfigUpdateReq**](AlertDingDingConfigUpdateReq.md) | alert dingding config info | 

### Return type

[**AlertDingDingConfigResp**](AlertDingDingConfigResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

