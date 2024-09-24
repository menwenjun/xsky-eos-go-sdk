# \DnsLoadBalancePoliciesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDNSLoadBalancePolicies**](DnsLoadBalancePoliciesAPI.md#ListDNSLoadBalancePolicies) | **Get** /dns-load-balance-policies/ | 



## ListDNSLoadBalancePolicies

> DNSLoadBalancePoliciesResp ListDNSLoadBalancePolicies(ctx).Limit(limit).Offset(offset).ResourceType(resourceType).Execute()





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
	resourceType := "resourceType_example" // string | resource type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsLoadBalancePoliciesAPI.ListDNSLoadBalancePolicies(context.Background()).Limit(limit).Offset(offset).ResourceType(resourceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsLoadBalancePoliciesAPI.ListDNSLoadBalancePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDNSLoadBalancePolicies`: DNSLoadBalancePoliciesResp
	fmt.Fprintf(os.Stdout, "Response from `DnsLoadBalancePoliciesAPI.ListDNSLoadBalancePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDNSLoadBalancePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **resourceType** | **string** | resource type | 

### Return type

[**DNSLoadBalancePoliciesResp**](DNSLoadBalancePoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

