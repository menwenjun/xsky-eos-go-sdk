# \NetworkAddressesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkAddress**](NetworkAddressesAPI.md#GetNetworkAddress) | **Get** /network-addresses/{network_address_id} | 
[**ListNetworkAddresses**](NetworkAddressesAPI.md#ListNetworkAddresses) | **Get** /network-addresses/ | 



## GetNetworkAddress

> NetworkAddressResp GetNetworkAddress(ctx, networkAddressId).Execute()





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
	networkAddressId := int64(789) // int64 | network address id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAddressesAPI.GetNetworkAddress(context.Background(), networkAddressId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAddressesAPI.GetNetworkAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkAddress`: NetworkAddressResp
	fmt.Fprintf(os.Stdout, "Response from `NetworkAddressesAPI.GetNetworkAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkAddressId** | **int64** | network address id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkAddressResp**](NetworkAddressResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkAddresses

> NetworkAddressesResp ListNetworkAddresses(ctx).Limit(limit).Offset(offset).NetworkInterfaceId(networkInterfaceId).HostId(hostId).Role(role).VipGroupId(vipGroupId).Execute()





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
	networkInterfaceId := int64(789) // int64 | network interface id (optional)
	hostId := int64(789) // int64 | host id (optional)
	role := "role_example" // string | network address role (optional)
	vipGroupId := int64(789) // int64 | vip group id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkAddressesAPI.ListNetworkAddresses(context.Background()).Limit(limit).Offset(offset).NetworkInterfaceId(networkInterfaceId).HostId(hostId).Role(role).VipGroupId(vipGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkAddressesAPI.ListNetworkAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkAddresses`: NetworkAddressesResp
	fmt.Fprintf(os.Stdout, "Response from `NetworkAddressesAPI.ListNetworkAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **networkInterfaceId** | **int64** | network interface id | 
 **hostId** | **int64** | host id | 
 **role** | **string** | network address role | 
 **vipGroupId** | **int64** | vip group id | 

### Return type

[**NetworkAddressesResp**](NetworkAddressesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

