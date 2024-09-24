# \HostValidatorsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHostValidator**](HostValidatorsAPI.md#CreateHostValidator) | **Post** /host-validators/ | 
[**GetHostValidator**](HostValidatorsAPI.md#GetHostValidator) | **Get** /host-validators/{host_validator_id} | 



## CreateHostValidator

> HostValidatorResp CreateHostValidator(ctx).Body(body).Execute()





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
	body := *openapiclient.NewHostValidatorCreateReq() // HostValidatorCreateReq | host validator info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostValidatorsAPI.CreateHostValidator(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostValidatorsAPI.CreateHostValidator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHostValidator`: HostValidatorResp
	fmt.Fprintf(os.Stdout, "Response from `HostValidatorsAPI.CreateHostValidator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HostValidatorCreateReq**](HostValidatorCreateReq.md) | host validator info | 

### Return type

[**HostValidatorResp**](HostValidatorResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostValidator

> HostValidatorResp GetHostValidator(ctx, hostValidatorId).Execute()





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
	hostValidatorId := int64(789) // int64 | host validator id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostValidatorsAPI.GetHostValidator(context.Background(), hostValidatorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostValidatorsAPI.GetHostValidator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostValidator`: HostValidatorResp
	fmt.Fprintf(os.Stdout, "Response from `HostValidatorsAPI.GetHostValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostValidatorId** | **int64** | host validator id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostValidatorResp**](HostValidatorResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

