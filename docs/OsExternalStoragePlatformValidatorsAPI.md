# \OsExternalStoragePlatformValidatorsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOSExStoragePlatformValidator**](OsExternalStoragePlatformValidatorsAPI.md#CreateOSExStoragePlatformValidator) | **Post** /os-external-storage-platform-validators/ | 
[**GetOSExStoragePlatformValidator**](OsExternalStoragePlatformValidatorsAPI.md#GetOSExStoragePlatformValidator) | **Get** /os-external-storage-platform-validators/{validator_id} | 



## CreateOSExStoragePlatformValidator

> OSExStoragePlatformValidatorResp CreateOSExStoragePlatformValidator(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewOSExStoragePlatformValidatorCreateReq() // OSExStoragePlatformValidatorCreateReq | external storage platform validator info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsExternalStoragePlatformValidatorsAPI.CreateOSExStoragePlatformValidator(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsExternalStoragePlatformValidatorsAPI.CreateOSExStoragePlatformValidator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOSExStoragePlatformValidator`: OSExStoragePlatformValidatorResp
	fmt.Fprintf(os.Stdout, "Response from `OsExternalStoragePlatformValidatorsAPI.CreateOSExStoragePlatformValidator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOSExStoragePlatformValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OSExStoragePlatformValidatorCreateReq**](OSExStoragePlatformValidatorCreateReq.md) | external storage platform validator info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSExStoragePlatformValidatorResp**](OSExStoragePlatformValidatorResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSExStoragePlatformValidator

> OSExStoragePlatformValidatorResp GetOSExStoragePlatformValidator(ctx, validatorId).Execute()





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
	validatorId := int64(789) // int64 | external storage platform validator id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsExternalStoragePlatformValidatorsAPI.GetOSExStoragePlatformValidator(context.Background(), validatorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsExternalStoragePlatformValidatorsAPI.GetOSExStoragePlatformValidator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSExStoragePlatformValidator`: OSExStoragePlatformValidatorResp
	fmt.Fprintf(os.Stdout, "Response from `OsExternalStoragePlatformValidatorsAPI.GetOSExStoragePlatformValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**validatorId** | **int64** | external storage platform validator id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSExStoragePlatformValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSExStoragePlatformValidatorResp**](OSExStoragePlatformValidatorResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

