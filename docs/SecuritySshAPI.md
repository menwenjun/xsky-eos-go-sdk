# \SecuritySshAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSSHValidation**](SecuritySshAPI.md#CreateSSHValidation) | **Post** /security-ssh/validation | 
[**ListSSHConfigs**](SecuritySshAPI.md#ListSSHConfigs) | **Get** /security-ssh/ | 
[**UpdateSSHConfig**](SecuritySshAPI.md#UpdateSSHConfig) | **Patch** /security-ssh/{config_id} | 



## CreateSSHValidation

> SecuritySSHValidationResp CreateSSHValidation(ctx).Body(body).Execute()





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
	body := *openapiclient.NewSecuritySSHValidationReq() // SecuritySSHValidationReq | ssh validation info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecuritySshAPI.CreateSSHValidation(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecuritySshAPI.CreateSSHValidation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSSHValidation`: SecuritySSHValidationResp
	fmt.Fprintf(os.Stdout, "Response from `SecuritySshAPI.CreateSSHValidation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHValidationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SecuritySSHValidationReq**](SecuritySSHValidationReq.md) | ssh validation info | 

### Return type

[**SecuritySSHValidationResp**](SecuritySSHValidationResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSSHConfigs

> SecuritySSHConfigResp ListSSHConfigs(ctx).Execute()





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
	resp, r, err := apiClient.SecuritySshAPI.ListSSHConfigs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecuritySshAPI.ListSSHConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSSHConfigs`: SecuritySSHConfigResp
	fmt.Fprintf(os.Stdout, "Response from `SecuritySshAPI.ListSSHConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSSHConfigsRequest struct via the builder pattern


### Return type

[**SecuritySSHConfigResp**](SecuritySSHConfigResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSSHConfig

> SecuritySSHConfigResp UpdateSSHConfig(ctx, configId).Body(body).Execute()





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
	configId := int64(789) // int64 | security ssh config id
	body := *openapiclient.NewSecuritySSHConfigReq() // SecuritySSHConfigReq | ssh validation info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecuritySshAPI.UpdateSSHConfig(context.Background(), configId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecuritySshAPI.UpdateSSHConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSSHConfig`: SecuritySSHConfigResp
	fmt.Fprintf(os.Stdout, "Response from `SecuritySshAPI.UpdateSSHConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **int64** | security ssh config id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSSHConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SecuritySSHConfigReq**](SecuritySSHConfigReq.md) | ssh validation info | 

### Return type

[**SecuritySSHConfigResp**](SecuritySSHConfigResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

