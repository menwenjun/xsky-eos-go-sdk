# \ProtectionDomainsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProtectionDomain**](ProtectionDomainsAPI.md#GetProtectionDomain) | **Get** /protection-domains/{protection_domain_id} | 
[**ListProtectionDomains**](ProtectionDomainsAPI.md#ListProtectionDomains) | **Get** /protection-domains/ | 



## GetProtectionDomain

> ProtectionDomainRecordResp GetProtectionDomain(ctx, protectionDomainId).Execute()





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
	protectionDomainId := int64(789) // int64 | protection domain id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectionDomainsAPI.GetProtectionDomain(context.Background(), protectionDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionDomainsAPI.GetProtectionDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectionDomain`: ProtectionDomainRecordResp
	fmt.Fprintf(os.Stdout, "Response from `ProtectionDomainsAPI.GetProtectionDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**protectionDomainId** | **int64** | protection domain id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProtectionDomainRecordResp**](ProtectionDomainRecordResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProtectionDomains

> ProtectionDomainRecordsResp ListProtectionDomains(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProtectionDomainsAPI.ListProtectionDomains(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProtectionDomainsAPI.ListProtectionDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProtectionDomains`: ProtectionDomainRecordsResp
	fmt.Fprintf(os.Stdout, "Response from `ProtectionDomainsAPI.ListProtectionDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProtectionDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**ProtectionDomainRecordsResp**](ProtectionDomainRecordsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

