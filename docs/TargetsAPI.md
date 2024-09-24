# \TargetsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTargetGatewayIPs**](TargetsAPI.md#AddTargetGatewayIPs) | **Post** /targets/{target_id}:add-gateway-ips | 
[**CreateTarget**](TargetsAPI.md#CreateTarget) | **Post** /targets/ | 
[**DeleteTarget**](TargetsAPI.md#DeleteTarget) | **Delete** /targets/{target_id} | 
[**ListTargets**](TargetsAPI.md#ListTargets) | **Get** /targets/ | 



## AddTargetGatewayIPs

> TargetResp AddTargetGatewayIPs(ctx, targetId).Body(body).Execute()





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
	targetId := int64(789) // int64 | target id
	body := *openapiclient.NewTargetAddGatewayIPsReq() // TargetAddGatewayIPsReq | gateway ips info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetsAPI.AddTargetGatewayIPs(context.Background(), targetId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetsAPI.AddTargetGatewayIPs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTargetGatewayIPs`: TargetResp
	fmt.Fprintf(os.Stdout, "Response from `TargetsAPI.AddTargetGatewayIPs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetId** | **int64** | target id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTargetGatewayIPsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TargetAddGatewayIPsReq**](TargetAddGatewayIPsReq.md) | gateway ips info | 

### Return type

[**TargetResp**](TargetResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTarget

> TargetResp CreateTarget(ctx).Target(target).Execute()





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
	target := *openapiclient.NewTargetCreateReq() // TargetCreateReq | target info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetsAPI.CreateTarget(context.Background()).Target(target).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetsAPI.CreateTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTarget`: TargetResp
	fmt.Fprintf(os.Stdout, "Response from `TargetsAPI.CreateTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **target** | [**TargetCreateReq**](TargetCreateReq.md) | target info | 

### Return type

[**TargetResp**](TargetResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTarget

> TargetResp DeleteTarget(ctx, targetId).Force(force).Execute()





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
	targetId := int64(789) // int64 | target id
	force := true // bool | force delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetsAPI.DeleteTarget(context.Background(), targetId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetsAPI.DeleteTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTarget`: TargetResp
	fmt.Fprintf(os.Stdout, "Response from `TargetsAPI.DeleteTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetId** | **int64** | target id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 

### Return type

[**TargetResp**](TargetResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTargets

> TargetsResp ListTargets(ctx).Limit(limit).Offset(offset).IsFetchLunInfo(isFetchLunInfo).QueryDepth(queryDepth).ClusterId(clusterId).HostId(hostId).AccessPathId(accessPathId).Q(q).Sort(sort).Execute()





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
	isFetchLunInfo := true // bool | whether to fetch lun info from target (optional)
	queryDepth := int64(789) // int64 | query depth (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)
	hostId := int64(789) // int64 | host id (optional)
	accessPathId := int64(789) // int64 | access path id (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetsAPI.ListTargets(context.Background()).Limit(limit).Offset(offset).IsFetchLunInfo(isFetchLunInfo).QueryDepth(queryDepth).ClusterId(clusterId).HostId(hostId).AccessPathId(accessPathId).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetsAPI.ListTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTargets`: TargetsResp
	fmt.Fprintf(os.Stdout, "Response from `TargetsAPI.ListTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **isFetchLunInfo** | **bool** | whether to fetch lun info from target | 
 **queryDepth** | **int64** | query depth | 
 **clusterId** | **string** | cluster id | 
 **hostId** | **int64** | host id | 
 **accessPathId** | **int64** | access path id | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**TargetsResp**](TargetsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

