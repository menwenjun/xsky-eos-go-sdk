# \DfsPathPerformancesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDfsPathPerformance**](DfsPathPerformancesAPI.md#CreateDfsPathPerformance) | **Post** /dfs-path-performances/ | 
[**DeleteDfsPathPerformance**](DfsPathPerformancesAPI.md#DeleteDfsPathPerformance) | **Delete** /dfs-path-performances/{dfs_path_performance_id} | 
[**GetDfsPathPerformance**](DfsPathPerformancesAPI.md#GetDfsPathPerformance) | **Get** /dfs-path-performances/{dfs_path_performance_id} | 
[**GetDfsPathPerformanceSamples**](DfsPathPerformancesAPI.md#GetDfsPathPerformanceSamples) | **Get** /dfs-path-performances/{dfs_path_performance_id}/samples | 
[**ListDfsPathPerformances**](DfsPathPerformancesAPI.md#ListDfsPathPerformances) | **Get** /dfs-path-performances/ | 



## CreateDfsPathPerformance

> DfsPathPerformanceResp CreateDfsPathPerformance(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsPathPerformanceCreateReq(*openapiclient.NewDfsPathPerformanceCreateReqPathPerformance(int64(123), "Path_example")) // DfsPathPerformanceCreateReq | path performance info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsPathPerformancesAPI.CreateDfsPathPerformance(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsPathPerformancesAPI.CreateDfsPathPerformance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsPathPerformance`: DfsPathPerformanceResp
	fmt.Fprintf(os.Stdout, "Response from `DfsPathPerformancesAPI.CreateDfsPathPerformance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsPathPerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsPathPerformanceCreateReq**](DfsPathPerformanceCreateReq.md) | path performance info | 

### Return type

[**DfsPathPerformanceResp**](DfsPathPerformanceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsPathPerformance

> DfsPathPerformanceResp DeleteDfsPathPerformance(ctx, dfsPathPerformanceId).Execute()





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
	dfsPathPerformanceId := int64(789) // int64 | dfs path performance id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsPathPerformancesAPI.DeleteDfsPathPerformance(context.Background(), dfsPathPerformanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsPathPerformancesAPI.DeleteDfsPathPerformance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsPathPerformance`: DfsPathPerformanceResp
	fmt.Fprintf(os.Stdout, "Response from `DfsPathPerformancesAPI.DeleteDfsPathPerformance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsPathPerformanceId** | **int64** | dfs path performance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsPathPerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsPathPerformanceResp**](DfsPathPerformanceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsPathPerformance

> DfsPathPerformanceResp GetDfsPathPerformance(ctx, dfsPathPerformanceId).Execute()





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
	dfsPathPerformanceId := int64(789) // int64 | path performance id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsPathPerformancesAPI.GetDfsPathPerformance(context.Background(), dfsPathPerformanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsPathPerformancesAPI.GetDfsPathPerformance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsPathPerformance`: DfsPathPerformanceResp
	fmt.Fprintf(os.Stdout, "Response from `DfsPathPerformancesAPI.GetDfsPathPerformance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsPathPerformanceId** | **int64** | path performance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsPathPerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsPathPerformanceResp**](DfsPathPerformanceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsPathPerformanceSamples

> DfsPathPerformanceSamplesResp GetDfsPathPerformanceSamples(ctx, dfsPathPerformanceId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	dfsPathPerformanceId := int64(789) // int64 | dfs path performance id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsPathPerformancesAPI.GetDfsPathPerformanceSamples(context.Background(), dfsPathPerformanceId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsPathPerformancesAPI.GetDfsPathPerformanceSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsPathPerformanceSamples`: DfsPathPerformanceSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsPathPerformancesAPI.GetDfsPathPerformanceSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsPathPerformanceId** | **int64** | dfs path performance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsPathPerformanceSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**DfsPathPerformanceSamplesResp**](DfsPathPerformanceSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsPathPerformances

> DfsPathPerformancesResp ListDfsPathPerformances(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).Path(path).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	path := "path_example" // string | path (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsPathPerformancesAPI.ListDfsPathPerformances(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsPathPerformancesAPI.ListDfsPathPerformances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsPathPerformances`: DfsPathPerformancesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsPathPerformancesAPI.ListDfsPathPerformances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsPathPerformancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **path** | **string** | path | 

### Return type

[**DfsPathPerformancesResp**](DfsPathPerformancesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

