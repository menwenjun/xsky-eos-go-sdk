# \DfsGatewayServicesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchGetDfsGatewayServiceSamples**](DfsGatewayServicesAPI.md#BatchGetDfsGatewayServiceSamples) | **Get** /dfs-gateway-services/samples | 
[**GetDfsGatewayService**](DfsGatewayServicesAPI.md#GetDfsGatewayService) | **Get** /dfs-gateway-services/{dfs_gateway_service_id} | 
[**GetDfsGatewayServiceSamples**](DfsGatewayServicesAPI.md#GetDfsGatewayServiceSamples) | **Get** /dfs-gateway-services/{dfs_gateway_service_id}/samples | 
[**ListDfsGatewayServices**](DfsGatewayServicesAPI.md#ListDfsGatewayServices) | **Get** /dfs-gateway-services/ | 



## BatchGetDfsGatewayServiceSamples

> MultiDfsGatewayServicesSamplesResp BatchGetDfsGatewayServiceSamples(ctx).ClusterId(clusterId).Ids(ids).DfsGatewayId(dfsGatewayId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	ids := "ids_example" // string | dfs gateway service id, format 1,2,3 (optional)
	dfsGatewayId := int64(789) // int64 | dfs gateway service id (optional)
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayServicesAPI.BatchGetDfsGatewayServiceSamples(context.Background()).ClusterId(clusterId).Ids(ids).DfsGatewayId(dfsGatewayId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayServicesAPI.BatchGetDfsGatewayServiceSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchGetDfsGatewayServiceSamples`: MultiDfsGatewayServicesSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayServicesAPI.BatchGetDfsGatewayServiceSamples`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchGetDfsGatewayServiceSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **ids** | **string** | dfs gateway service id, format 1,2,3 | 
 **dfsGatewayId** | **int64** | dfs gateway service id | 
 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**MultiDfsGatewayServicesSamplesResp**](MultiDfsGatewayServicesSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsGatewayService

> DfsGatewayServiceResp GetDfsGatewayService(ctx, dfsGatewayServiceId).Execute()





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
	dfsGatewayServiceId := int64(789) // int64 | gateway service id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayServicesAPI.GetDfsGatewayService(context.Background(), dfsGatewayServiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayServicesAPI.GetDfsGatewayService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsGatewayService`: DfsGatewayServiceResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayServicesAPI.GetDfsGatewayService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayServiceId** | **int64** | gateway service id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsGatewayServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsGatewayServiceResp**](DfsGatewayServiceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsGatewayServiceSamples

> DfsGatewayServiceSamplesResp GetDfsGatewayServiceSamples(ctx, dfsGatewayServiceId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	dfsGatewayServiceId := int64(789) // int64 | dfs gateway service id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayServicesAPI.GetDfsGatewayServiceSamples(context.Background(), dfsGatewayServiceId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayServicesAPI.GetDfsGatewayServiceSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsGatewayServiceSamples`: DfsGatewayServiceSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayServicesAPI.GetDfsGatewayServiceSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayServiceId** | **int64** | dfs gateway service id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsGatewayServiceSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**DfsGatewayServiceSamplesResp**](DfsGatewayServiceSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsGatewayServices

> DfsGatewaysResp ListDfsGatewayServices(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).HostId(hostId).DfsGatewayId(dfsGatewayId).DfsGatewayGroupId(dfsGatewayGroupId).Execute()





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
	hostId := int64(789) // int64 | host id (optional)
	dfsGatewayId := int64(789) // int64 | dfs gateway id (optional)
	dfsGatewayGroupId := int64(789) // int64 | dfs gateway group id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayServicesAPI.ListDfsGatewayServices(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).HostId(hostId).DfsGatewayId(dfsGatewayId).DfsGatewayGroupId(dfsGatewayGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayServicesAPI.ListDfsGatewayServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsGatewayServices`: DfsGatewaysResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayServicesAPI.ListDfsGatewayServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsGatewayServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 
 **hostId** | **int64** | host id | 
 **dfsGatewayId** | **int64** | dfs gateway id | 
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

