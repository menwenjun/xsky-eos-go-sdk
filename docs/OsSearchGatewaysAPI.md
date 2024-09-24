# \OsSearchGatewaysAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOSSearchGatewaySamples**](OsSearchGatewaysAPI.md#GetOSSearchGatewaySamples) | **Get** /os-search-gateways/{gateway_id}/samples | 
[**GetOSSearchGateways**](OsSearchGatewaysAPI.md#GetOSSearchGateways) | **Get** /os-search-gateways/{gateway_id} | 
[**ListOSSearchGateways**](OsSearchGatewaysAPI.md#ListOSSearchGateways) | **Get** /os-search-gateways/ | 



## GetOSSearchGatewaySamples

> OSSearchGatewaySamplesResp GetOSSearchGatewaySamples(ctx, gatewayId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	gatewayId := int64(789) // int64 | os search gateway id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSearchGatewaysAPI.GetOSSearchGatewaySamples(context.Background(), gatewayId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSearchGatewaysAPI.GetOSSearchGatewaySamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSSearchGatewaySamples`: OSSearchGatewaySamplesResp
	fmt.Fprintf(os.Stdout, "Response from `OsSearchGatewaysAPI.GetOSSearchGatewaySamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | os search gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSSearchGatewaySamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**OSSearchGatewaySamplesResp**](OSSearchGatewaySamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSSearchGateways

> OSSearchGatewayResp GetOSSearchGateways(ctx, gatewayId).Execute()





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
	gatewayId := int64(789) // int64 | os search gateway id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSearchGatewaysAPI.GetOSSearchGateways(context.Background(), gatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSearchGatewaysAPI.GetOSSearchGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSSearchGateways`: OSSearchGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `OsSearchGatewaysAPI.GetOSSearchGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | os search gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSSearchGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSSearchGatewayResp**](OSSearchGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOSSearchGateways

> OSSearchGatewaysResp ListOSSearchGateways(ctx).OsSearchEngineId(osSearchEngineId).ClusterId(clusterId).Q(q).Sort(sort).Execute()





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
	osSearchEngineId := int64(789) // int64 | os search engine id (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSearchGatewaysAPI.ListOSSearchGateways(context.Background()).OsSearchEngineId(osSearchEngineId).ClusterId(clusterId).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSearchGatewaysAPI.ListOSSearchGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSSearchGateways`: OSSearchGatewaysResp
	fmt.Fprintf(os.Stdout, "Response from `OsSearchGatewaysAPI.ListOSSearchGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSSearchGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **osSearchEngineId** | **int64** | os search engine id | 
 **clusterId** | **string** | cluster id | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**OSSearchGatewaysResp**](OSSearchGatewaysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

