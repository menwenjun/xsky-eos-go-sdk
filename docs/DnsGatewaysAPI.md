# \DnsGatewaysAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDNSGateway**](DnsGatewaysAPI.md#GetDNSGateway) | **Get** /dns-gateways/{gateway_id} | 
[**ListDNSGateways**](DnsGatewaysAPI.md#ListDNSGateways) | **Get** /dns-gateways/ | 



## GetDNSGateway

> DNSGatewayResp GetDNSGateway(ctx, gatewayId).Execute()





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
	gatewayId := int64(789) // int64 | dns gateway id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsGatewaysAPI.GetDNSGateway(context.Background(), gatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsGatewaysAPI.GetDNSGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDNSGateway`: DNSGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `DnsGatewaysAPI.GetDNSGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | dns gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDNSGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DNSGatewayResp**](DNSGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDNSGateways

> DNSGatewaysResp ListDNSGateways(ctx).Limit(limit).Offset(offset).DnsGatewayGroupId(dnsGatewayGroupId).Execute()





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
	dnsGatewayGroupId := int64(789) // int64 | dns_gateway_group id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsGatewaysAPI.ListDNSGateways(context.Background()).Limit(limit).Offset(offset).DnsGatewayGroupId(dnsGatewayGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsGatewaysAPI.ListDNSGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDNSGateways`: DNSGatewaysResp
	fmt.Fprintf(os.Stdout, "Response from `DnsGatewaysAPI.ListDNSGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDNSGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **dnsGatewayGroupId** | **int64** | dns_gateway_group id | 

### Return type

[**DNSGatewaysResp**](DNSGatewaysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

