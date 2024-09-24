# \DfsGatewayNetworkAddressesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDfsGatewayNetworkAddress**](DfsGatewayNetworkAddressesAPI.md#GetDfsGatewayNetworkAddress) | **Get** /dfs-gateway-network-addresses/{dfs_gateway_network_address_id} | 
[**ListDfsGatewayNetworkAddresses**](DfsGatewayNetworkAddressesAPI.md#ListDfsGatewayNetworkAddresses) | **Get** /dfs-gateway-network-addresses/ | 



## GetDfsGatewayNetworkAddress

> DfsGatewayNetworkAddressResp GetDfsGatewayNetworkAddress(ctx, dfsGatewayNetworkAddressId).Execute()





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
	dfsGatewayNetworkAddressId := int64(789) // int64 | gateway network address id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayNetworkAddressesAPI.GetDfsGatewayNetworkAddress(context.Background(), dfsGatewayNetworkAddressId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayNetworkAddressesAPI.GetDfsGatewayNetworkAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsGatewayNetworkAddress`: DfsGatewayNetworkAddressResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayNetworkAddressesAPI.GetDfsGatewayNetworkAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayNetworkAddressId** | **int64** | gateway network address id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsGatewayNetworkAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsGatewayNetworkAddressResp**](DfsGatewayNetworkAddressResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsGatewayNetworkAddresses

> DfsGatewayNetworkAddressesResp ListDfsGatewayNetworkAddresses(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).DfsGatewayGroupId(dfsGatewayGroupId).DfsGatewayZoneId(dfsGatewayZoneId).Execute()





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
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)
	dfsGatewayGroupId := int64(789) // int64 | dfs gateway group id (optional)
	dfsGatewayZoneId := int64(789) // int64 | dfs gateway zone id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayNetworkAddressesAPI.ListDfsGatewayNetworkAddresses(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).DfsGatewayGroupId(dfsGatewayGroupId).DfsGatewayZoneId(dfsGatewayZoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayNetworkAddressesAPI.ListDfsGatewayNetworkAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsGatewayNetworkAddresses`: DfsGatewayNetworkAddressesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayNetworkAddressesAPI.ListDfsGatewayNetworkAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsGatewayNetworkAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 
 **dfsGatewayGroupId** | **int64** | dfs gateway group id | 
 **dfsGatewayZoneId** | **int64** | dfs gateway zone id | 

### Return type

[**DfsGatewayNetworkAddressesResp**](DfsGatewayNetworkAddressesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

