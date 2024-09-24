# \DomainUserValidatorsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomainUserValidator**](DomainUserValidatorsAPI.md#CreateDomainUserValidator) | **Post** /domain-user-validators/ | 
[**GetDomainUserValidator**](DomainUserValidatorsAPI.md#GetDomainUserValidator) | **Get** /domain-user-validators/{domain_user_validator_id} | 



## CreateDomainUserValidator

> DomainUserValidatorResp CreateDomainUserValidator(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDomainUserValidatorCreateReq() // DomainUserValidatorCreateReq | domain user validator info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainUserValidatorsAPI.CreateDomainUserValidator(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainUserValidatorsAPI.CreateDomainUserValidator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDomainUserValidator`: DomainUserValidatorResp
	fmt.Fprintf(os.Stdout, "Response from `DomainUserValidatorsAPI.CreateDomainUserValidator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainUserValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DomainUserValidatorCreateReq**](DomainUserValidatorCreateReq.md) | domain user validator info | 

### Return type

[**DomainUserValidatorResp**](DomainUserValidatorResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainUserValidator

> DomainUserValidatorResp GetDomainUserValidator(ctx, domainUserValidatorId).Execute()





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
	domainUserValidatorId := int64(789) // int64 | domain user validator id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainUserValidatorsAPI.GetDomainUserValidator(context.Background(), domainUserValidatorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainUserValidatorsAPI.GetDomainUserValidator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainUserValidator`: DomainUserValidatorResp
	fmt.Fprintf(os.Stdout, "Response from `DomainUserValidatorsAPI.GetDomainUserValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainUserValidatorId** | **int64** | domain user validator id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainUserValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainUserValidatorResp**](DomainUserValidatorResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

