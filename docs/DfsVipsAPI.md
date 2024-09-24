# \DfsVipsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDfsVIP**](DfsVipsAPI.md#GetDfsVIP) | **Get** /dfs-vips/{dfs_vip_id} | 
[**ListDfsVIPs**](DfsVipsAPI.md#ListDfsVIPs) | **Get** /dfs-vips/ | 
[**MoveDfsVIP**](DfsVipsAPI.md#MoveDfsVIP) | **Post** /dfs-vips/{dfs_vip_id} | 



## GetDfsVIP

> DfsVIPResp GetDfsVIP(ctx, dfsVipId).Execute()





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
	dfsVipId := int64(789) // int64 | vip id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsVipsAPI.GetDfsVIP(context.Background(), dfsVipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsVipsAPI.GetDfsVIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsVIP`: DfsVIPResp
	fmt.Fprintf(os.Stdout, "Response from `DfsVipsAPI.GetDfsVIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsVipId** | **int64** | vip id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsVIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsVIPResp**](DfsVIPResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsVIPs

> DfsVIPsResp ListDfsVIPs(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).DfsGatewayGroupId(dfsGatewayGroupId).DfsGatewayZoneId(dfsGatewayZoneId).PrimaryGatewayId(primaryGatewayId).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	dfsGatewayGroupId := int64(789) // int64 | dfs vip group id (optional)
	dfsGatewayZoneId := int64(789) // int64 | dfs gateway zone id (optional)
	primaryGatewayId := int64(789) // int64 | primary gateway id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsVipsAPI.ListDfsVIPs(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).DfsGatewayGroupId(dfsGatewayGroupId).DfsGatewayZoneId(dfsGatewayZoneId).PrimaryGatewayId(primaryGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsVipsAPI.ListDfsVIPs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsVIPs`: DfsVIPsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsVipsAPI.ListDfsVIPs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsVIPsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **dfsGatewayGroupId** | **int64** | dfs vip group id | 
 **dfsGatewayZoneId** | **int64** | dfs gateway zone id | 
 **primaryGatewayId** | **int64** | primary gateway id | 

### Return type

[**DfsVIPsResp**](DfsVIPsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveDfsVIP

> DfsVIPResp MoveDfsVIP(ctx, dfsVipId).Body(body).Execute()





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
	dfsVipId := int64(789) // int64 | vip id
	body := *openapiclient.NewDfsVIPMoveReq(*openapiclient.NewDfsVIPMoveReqVIP(int64(123))) // DfsVIPMoveReq | moving vip info (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsVipsAPI.MoveDfsVIP(context.Background(), dfsVipId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsVipsAPI.MoveDfsVIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveDfsVIP`: DfsVIPResp
	fmt.Fprintf(os.Stdout, "Response from `DfsVipsAPI.MoveDfsVIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsVipId** | **int64** | vip id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveDfsVIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsVIPMoveReq**](DfsVIPMoveReq.md) | moving vip info | 

### Return type

[**DfsVIPResp**](DfsVIPResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

