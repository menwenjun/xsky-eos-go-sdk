# \DfsGatewayZonesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDfsGatewayZone**](DfsGatewayZonesAPI.md#CreateDfsGatewayZone) | **Post** /dfs-gateway-zones/ | 
[**DeleteDfsGatewayZone**](DfsGatewayZonesAPI.md#DeleteDfsGatewayZone) | **Delete** /dfs-gateway-zones/{dfs_gateway_zone_id} | 
[**GetDfsGatewayZone**](DfsGatewayZonesAPI.md#GetDfsGatewayZone) | **Get** /dfs-gateway-zones/{dfs_gateway_zone_id} | 
[**GetDfsGatewayZoneSamples**](DfsGatewayZonesAPI.md#GetDfsGatewayZoneSamples) | **Get** /dfs-gateway-zones/{dfs_gateway_zone_id}/samples | 
[**ListDfsGatewayZones**](DfsGatewayZonesAPI.md#ListDfsGatewayZones) | **Get** /dfs-gateway-zones/ | 
[**UpdateDfsGatewayZone**](DfsGatewayZonesAPI.md#UpdateDfsGatewayZone) | **Patch** /dfs-gateway-zones/{dfs_gateway_zone_id} | 



## CreateDfsGatewayZone

> DfsGatewayZoneResp CreateDfsGatewayZone(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewDfsGatewayZoneCreateReq(*openapiclient.NewDfsGatewayZoneCreateReqGatewayZone(int64(123), []openapiclient.DfsGatewayReq{*openapiclient.NewDfsGatewayReq(int64(123), int64(123))}, []string{"DfsVips_example"}, "Name_example")) // DfsGatewayZoneCreateReq | gateway zone info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayZonesAPI.CreateDfsGatewayZone(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayZonesAPI.CreateDfsGatewayZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsGatewayZone`: DfsGatewayZoneResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayZonesAPI.CreateDfsGatewayZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsGatewayZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsGatewayZoneCreateReq**](DfsGatewayZoneCreateReq.md) | gateway zone info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**DfsGatewayZoneResp**](DfsGatewayZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsGatewayZone

> DfsGatewayZoneResp DeleteDfsGatewayZone(ctx, dfsGatewayZoneId).Force(force).Execute()





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
	dfsGatewayZoneId := int64(789) // int64 | gateway zone id
	force := true // bool | force delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayZonesAPI.DeleteDfsGatewayZone(context.Background(), dfsGatewayZoneId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayZonesAPI.DeleteDfsGatewayZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsGatewayZone`: DfsGatewayZoneResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayZonesAPI.DeleteDfsGatewayZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayZoneId** | **int64** | gateway zone id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsGatewayZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 

### Return type

[**DfsGatewayZoneResp**](DfsGatewayZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsGatewayZone

> DfsGatewayZoneResp GetDfsGatewayZone(ctx, dfsGatewayZoneId).Execute()





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
	dfsGatewayZoneId := int64(789) // int64 | gateway zone id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayZonesAPI.GetDfsGatewayZone(context.Background(), dfsGatewayZoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayZonesAPI.GetDfsGatewayZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsGatewayZone`: DfsGatewayZoneResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayZonesAPI.GetDfsGatewayZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayZoneId** | **int64** | gateway zone id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsGatewayZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsGatewayZoneResp**](DfsGatewayZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsGatewayZoneSamples

> DfsGatewayZoneSamplesResp GetDfsGatewayZoneSamples(ctx, dfsGatewayZoneId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	dfsGatewayZoneId := int64(789) // int64 | dfs gateway zone id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayZonesAPI.GetDfsGatewayZoneSamples(context.Background(), dfsGatewayZoneId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayZonesAPI.GetDfsGatewayZoneSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsGatewayZoneSamples`: DfsGatewayZoneSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayZonesAPI.GetDfsGatewayZoneSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayZoneId** | **int64** | dfs gateway zone id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsGatewayZoneSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**DfsGatewayZoneSamplesResp**](DfsGatewayZoneSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsGatewayZones

> DfsGatewayZonesResp ListDfsGatewayZones(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayZonesAPI.ListDfsGatewayZones(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayZonesAPI.ListDfsGatewayZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsGatewayZones`: DfsGatewayZonesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayZonesAPI.ListDfsGatewayZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsGatewayZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 

### Return type

[**DfsGatewayZonesResp**](DfsGatewayZonesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsGatewayZone

> DfsGatewayZoneResp UpdateDfsGatewayZone(ctx, dfsGatewayZoneId).Body(body).Execute()





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
	dfsGatewayZoneId := int64(789) // int64 | gateway zone id
	body := *openapiclient.NewDfsGatewayZoneUpdateReq(*openapiclient.NewDfsGatewayZoneUpdateReqGatewayZone()) // DfsGatewayZoneUpdateReq | gateway zone info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayZonesAPI.UpdateDfsGatewayZone(context.Background(), dfsGatewayZoneId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayZonesAPI.UpdateDfsGatewayZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsGatewayZone`: DfsGatewayZoneResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayZonesAPI.UpdateDfsGatewayZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayZoneId** | **int64** | gateway zone id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsGatewayZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsGatewayZoneUpdateReq**](DfsGatewayZoneUpdateReq.md) | gateway zone info | 

### Return type

[**DfsGatewayZoneResp**](DfsGatewayZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

