# \DfsGatewaysAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDfsGateway**](DfsGatewaysAPI.md#GetDfsGateway) | **Get** /dfs-gateways/{dfs_gateway_id} | 
[**GetDfsGatewaySamples**](DfsGatewaysAPI.md#GetDfsGatewaySamples) | **Get** /dfs-gateways/{dfs_gateway_id}/samples | 
[**ListDfsGatewayConnections**](DfsGatewaysAPI.md#ListDfsGatewayConnections) | **Get** /dfs-gateways/{dfs_gateway_id}/connections | 
[**ListDfsGateways**](DfsGatewaysAPI.md#ListDfsGateways) | **Get** /dfs-gateways/ | 



## GetDfsGateway

> DfsGatewayResp GetDfsGateway(ctx, dfsGatewayId).Execute()





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
	dfsGatewayId := int64(789) // int64 | gateway id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewaysAPI.GetDfsGateway(context.Background(), dfsGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewaysAPI.GetDfsGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsGateway`: DfsGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewaysAPI.GetDfsGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayId** | **int64** | gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsGatewayResp**](DfsGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsGatewaySamples

> DfsGatewaySamplesResp GetDfsGatewaySamples(ctx, dfsGatewayId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	dfsGatewayId := int64(789) // int64 | dfs gateway id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewaysAPI.GetDfsGatewaySamples(context.Background(), dfsGatewayId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewaysAPI.GetDfsGatewaySamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsGatewaySamples`: DfsGatewaySamplesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewaysAPI.GetDfsGatewaySamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayId** | **int64** | dfs gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsGatewaySamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**DfsGatewaySamplesResp**](DfsGatewaySamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsGatewayConnections

> DfsGatewayConnectionsResp ListDfsGatewayConnections(ctx, dfsGatewayId).Types(types).Execute()





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
	dfsGatewayId := int64(789) // int64 | dfs gateway id
	types := "types_example" // string | share types, value is smb|nfs|ftp|s3|hdfs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewaysAPI.ListDfsGatewayConnections(context.Background(), dfsGatewayId).Types(types).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewaysAPI.ListDfsGatewayConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsGatewayConnections`: DfsGatewayConnectionsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewaysAPI.ListDfsGatewayConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayId** | **int64** | dfs gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDfsGatewayConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **types** | **string** | share types, value is smb|nfs|ftp|s3|hdfs | 

### Return type

[**DfsGatewayConnectionsResp**](DfsGatewayConnectionsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsGateways

> DfsGatewaysResp ListDfsGateways(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).DfsGatewayGroupId(dfsGatewayGroupId).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	dfsGatewayGroupId := int64(789) // int64 | dfs gateway group id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewaysAPI.ListDfsGateways(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).DfsGatewayGroupId(dfsGatewayGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewaysAPI.ListDfsGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsGateways`: DfsGatewaysResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewaysAPI.ListDfsGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 
 **dfsGatewayGroupId** | **int64** | dfs gateway group id | 

### Return type

[**DfsGatewaysResp**](DfsGatewaysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

