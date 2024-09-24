# \HostMaintainValidatorsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHostMaintainValidator**](HostMaintainValidatorsAPI.md#CreateHostMaintainValidator) | **Post** /host-maintain-validators/ | 
[**GetHostMaintainValidator**](HostMaintainValidatorsAPI.md#GetHostMaintainValidator) | **Get** /host-maintain-validators/{host_maintain_validator_id} | 



## CreateHostMaintainValidator

> HostMaintainValidatorResp CreateHostMaintainValidator(ctx).Body(body).Execute()





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
	body := *openapiclient.NewHostMaintainValidatorCreateReq() // HostMaintainValidatorCreateReq | host maintain validator info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostMaintainValidatorsAPI.CreateHostMaintainValidator(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostMaintainValidatorsAPI.CreateHostMaintainValidator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHostMaintainValidator`: HostMaintainValidatorResp
	fmt.Fprintf(os.Stdout, "Response from `HostMaintainValidatorsAPI.CreateHostMaintainValidator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostMaintainValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HostMaintainValidatorCreateReq**](HostMaintainValidatorCreateReq.md) | host maintain validator info | 

### Return type

[**HostMaintainValidatorResp**](HostMaintainValidatorResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostMaintainValidator

> HostMaintainValidatorResp GetHostMaintainValidator(ctx, hostMaintainValidatorId).Execute()





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
	hostMaintainValidatorId := int64(789) // int64 | host maintain validator id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostMaintainValidatorsAPI.GetHostMaintainValidator(context.Background(), hostMaintainValidatorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostMaintainValidatorsAPI.GetHostMaintainValidator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostMaintainValidator`: HostMaintainValidatorResp
	fmt.Fprintf(os.Stdout, "Response from `HostMaintainValidatorsAPI.GetHostMaintainValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostMaintainValidatorId** | **int64** | host maintain validator id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostMaintainValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostMaintainValidatorResp**](HostMaintainValidatorResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

