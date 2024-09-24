# \NetworkInterfacesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkInterface**](NetworkInterfacesAPI.md#GetNetworkInterface) | **Get** /network-interfaces/{network_interface_id} | 
[**GetNetworkInterfaceSamples**](NetworkInterfacesAPI.md#GetNetworkInterfaceSamples) | **Get** /network-interfaces/{network_interface_id}/samples | 
[**ListNetworkInterfaces**](NetworkInterfacesAPI.md#ListNetworkInterfaces) | **Get** /network-interfaces/ | 



## GetNetworkInterface

> NetworkInterfaceResp GetNetworkInterface(ctx, networkInterfaceId).Execute()





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
	networkInterfaceId := int64(789) // int64 | network interface id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkInterfacesAPI.GetNetworkInterface(context.Background(), networkInterfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfacesAPI.GetNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkInterface`: NetworkInterfaceResp
	fmt.Fprintf(os.Stdout, "Response from `NetworkInterfacesAPI.GetNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInterfaceId** | **int64** | network interface id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkInterfaceResp**](NetworkInterfaceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkInterfaceSamples

> NetworkInterfaceSamplesResp GetNetworkInterfaceSamples(ctx, networkInterfaceId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	networkInterfaceId := int64(789) // int64 | network interface id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkInterfacesAPI.GetNetworkInterfaceSamples(context.Background(), networkInterfaceId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfacesAPI.GetNetworkInterfaceSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkInterfaceSamples`: NetworkInterfaceSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `NetworkInterfacesAPI.GetNetworkInterfaceSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInterfaceId** | **int64** | network interface id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkInterfaceSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**NetworkInterfaceSamplesResp**](NetworkInterfaceSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkInterfaces

> NetworkInterfacesResp ListNetworkInterfaces(ctx).Limit(limit).Offset(offset).MasterNetworkInterfaceId(masterNetworkInterfaceId).HostId(hostId).Role(role).Execute()





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
	masterNetworkInterfaceId := int64(789) // int64 | master network interface id (optional)
	hostId := int64(789) // int64 | host id (optional)
	role := "role_example" // string | network interface role (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkInterfacesAPI.ListNetworkInterfaces(context.Background()).Limit(limit).Offset(offset).MasterNetworkInterfaceId(masterNetworkInterfaceId).HostId(hostId).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfacesAPI.ListNetworkInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkInterfaces`: NetworkInterfacesResp
	fmt.Fprintf(os.Stdout, "Response from `NetworkInterfacesAPI.ListNetworkInterfaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **masterNetworkInterfaceId** | **int64** | master network interface id | 
 **hostId** | **int64** | host id | 
 **role** | **string** | network interface role | 

### Return type

[**NetworkInterfacesResp**](NetworkInterfacesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

