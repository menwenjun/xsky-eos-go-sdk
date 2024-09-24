# \DnsZonesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDNSZone**](DnsZonesAPI.md#GetDNSZone) | **Get** /dns-zones/{zone_id} | 
[**ListDNSZones**](DnsZonesAPI.md#ListDNSZones) | **Get** /dns-zones/ | 
[**UpdateDNSZone**](DnsZonesAPI.md#UpdateDNSZone) | **Patch** /dns-zones/{zone_id} | 



## GetDNSZone

> DNSZoneResp GetDNSZone(ctx, zoneId).Execute()





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
	zoneId := int64(789) // int64 | dns zone id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsZonesAPI.GetDNSZone(context.Background(), zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsZonesAPI.GetDNSZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDNSZone`: DNSZoneResp
	fmt.Fprintf(os.Stdout, "Response from `DnsZonesAPI.GetDNSZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int64** | dns zone id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDNSZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DNSZoneResp**](DNSZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDNSZones

> DNSZonesResp ListDNSZones(ctx).Limit(limit).Offset(offset).DnsGatewayGroupId(dnsGatewayGroupId).Execute()





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
	dnsGatewayGroupId := int64(789) // int64 | dns gateway group id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsZonesAPI.ListDNSZones(context.Background()).Limit(limit).Offset(offset).DnsGatewayGroupId(dnsGatewayGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsZonesAPI.ListDNSZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDNSZones`: DNSZonesResp
	fmt.Fprintf(os.Stdout, "Response from `DnsZonesAPI.ListDNSZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDNSZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **dnsGatewayGroupId** | **int64** | dns gateway group id | 

### Return type

[**DNSZonesResp**](DNSZonesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDNSZone

> DNSZoneResp UpdateDNSZone(ctx, zoneId).DnsZone(dnsZone).Execute()





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
	zoneId := int64(789) // int64 | dns zone id
	dnsZone := *openapiclient.NewDNSZoneUpdateReq() // DNSZoneUpdateReq | DNS zone info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsZonesAPI.UpdateDNSZone(context.Background(), zoneId).DnsZone(dnsZone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsZonesAPI.UpdateDNSZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDNSZone`: DNSZoneResp
	fmt.Fprintf(os.Stdout, "Response from `DnsZonesAPI.UpdateDNSZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int64** | dns zone id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDNSZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsZone** | [**DNSZoneUpdateReq**](DNSZoneUpdateReq.md) | DNS zone info | 

### Return type

[**DNSZoneResp**](DNSZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

