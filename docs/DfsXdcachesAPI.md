# \DfsXdcachesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableDfsXDCache**](DfsXdcachesAPI.md#DisableDfsXDCache) | **Post** /dfs-xdcaches/:disable | 
[**EnableDfsXDCache**](DfsXdcachesAPI.md#EnableDfsXDCache) | **Post** /dfs-xdcaches/:enable | 
[**ListDfsXDCaches**](DfsXdcachesAPI.md#ListDfsXDCaches) | **Get** /dfs-xdcaches/ | 
[**ViewDfsXDCache**](DfsXdcachesAPI.md#ViewDfsXDCache) | **Get** /dfs-xdcaches/:view-xdcache | 



## DisableDfsXDCache

> DfsXDCacheResp DisableDfsXDCache(ctx).Body(body).AllowPathCreate(allowPathCreate).Execute()





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
	body := *openapiclient.NewDfsXDCacheDisableReq(*openapiclient.NewXDCache("Path_example", int64(123))) // DfsXDCacheDisableReq | xdcache info
	allowPathCreate := true // bool | allow create path when not existed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsXdcachesAPI.DisableDfsXDCache(context.Background()).Body(body).AllowPathCreate(allowPathCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsXdcachesAPI.DisableDfsXDCache``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableDfsXDCache`: DfsXDCacheResp
	fmt.Fprintf(os.Stdout, "Response from `DfsXdcachesAPI.DisableDfsXDCache`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableDfsXDCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsXDCacheDisableReq**](DfsXDCacheDisableReq.md) | xdcache info | 
 **allowPathCreate** | **bool** | allow create path when not existed | 

### Return type

[**DfsXDCacheResp**](DfsXDCacheResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableDfsXDCache

> DfsXDCacheResp EnableDfsXDCache(ctx).Body(body).AllowPathCreate(allowPathCreate).Execute()





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
	body := *openapiclient.NewDfsXDCacheEnableReq(*openapiclient.NewXDCache("Path_example", int64(123))) // DfsXDCacheEnableReq | xdcache info
	allowPathCreate := true // bool | allow create path when not existed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsXdcachesAPI.EnableDfsXDCache(context.Background()).Body(body).AllowPathCreate(allowPathCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsXdcachesAPI.EnableDfsXDCache``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableDfsXDCache`: DfsXDCacheResp
	fmt.Fprintf(os.Stdout, "Response from `DfsXdcachesAPI.EnableDfsXDCache`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableDfsXDCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsXDCacheEnableReq**](DfsXDCacheEnableReq.md) | xdcache info | 
 **allowPathCreate** | **bool** | allow create path when not existed | 

### Return type

[**DfsXDCacheResp**](DfsXDCacheResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsXDCaches

> DfsXDCachesResp ListDfsXDCaches(ctx).ClusterId(clusterId).Path(path).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	path := "path_example" // string | related dfs path (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsXdcachesAPI.ListDfsXDCaches(context.Background()).ClusterId(clusterId).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsXdcachesAPI.ListDfsXDCaches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsXDCaches`: DfsXDCachesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsXdcachesAPI.ListDfsXDCaches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsXDCachesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **path** | **string** | related dfs path | 

### Return type

[**DfsXDCachesResp**](DfsXDCachesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewDfsXDCache

> DfsXDCacheViewResp ViewDfsXDCache(ctx).Path(path).ClusterId(clusterId).Execute()





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
	path := "path_example" // string | related dfs path
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsXdcachesAPI.ViewDfsXDCache(context.Background()).Path(path).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsXdcachesAPI.ViewDfsXDCache``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ViewDfsXDCache`: DfsXDCacheViewResp
	fmt.Fprintf(os.Stdout, "Response from `DfsXdcachesAPI.ViewDfsXDCache`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiViewDfsXDCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | related dfs path | 
 **clusterId** | **string** | cluster id | 

### Return type

[**DfsXDCacheViewResp**](DfsXDCacheViewResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

