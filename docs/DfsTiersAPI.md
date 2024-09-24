# \DfsTiersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDfsTierPools**](DfsTiersAPI.md#AddDfsTierPools) | **Post** /dfs-tiers/{dfs_tier_id}:add-pools | 
[**GetDfsTier**](DfsTiersAPI.md#GetDfsTier) | **Get** /dfs-tiers/{dfs_tier_id} | 
[**ListDfsTiers**](DfsTiersAPI.md#ListDfsTiers) | **Get** /dfs-tiers/ | 
[**RemoveDfsTierPools**](DfsTiersAPI.md#RemoveDfsTierPools) | **Post** /dfs-tiers/{dfs_tier_id}:remove-pools | 



## AddDfsTierPools

> DfsTierResp AddDfsTierPools(ctx, dfsTierId).Body(body).Execute()





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
	dfsTierId := int64(789) // int64 | dfs tier id
	body := *openapiclient.NewDfsTierAddPoolReq(*openapiclient.NewDfsTierAddPoolReqTier()) // DfsTierAddPoolReq | pools info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsTiersAPI.AddDfsTierPools(context.Background(), dfsTierId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsTiersAPI.AddDfsTierPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDfsTierPools`: DfsTierResp
	fmt.Fprintf(os.Stdout, "Response from `DfsTiersAPI.AddDfsTierPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsTierId** | **int64** | dfs tier id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDfsTierPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsTierAddPoolReq**](DfsTierAddPoolReq.md) | pools info | 

### Return type

[**DfsTierResp**](DfsTierResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsTier

> DfsTierResp GetDfsTier(ctx, dfsTierId).Execute()





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
	dfsTierId := int64(789) // int64 | dfs tier id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsTiersAPI.GetDfsTier(context.Background(), dfsTierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsTiersAPI.GetDfsTier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsTier`: DfsTierResp
	fmt.Fprintf(os.Stdout, "Response from `DfsTiersAPI.GetDfsTier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsTierId** | **int64** | dfs tier id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsTierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsTierResp**](DfsTierResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsTiers

> DfsTiersResp ListDfsTiers(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).DfsRootfsId(dfsRootfsId).Q(q).Sort(sort).Execute()





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
	dfsRootfsId := int64(789) // int64 | dfs rootfs id (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsTiersAPI.ListDfsTiers(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).DfsRootfsId(dfsRootfsId).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsTiersAPI.ListDfsTiers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsTiers`: DfsTiersResp
	fmt.Fprintf(os.Stdout, "Response from `DfsTiersAPI.ListDfsTiers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsTiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **dfsRootfsId** | **int64** | dfs rootfs id | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DfsTiersResp**](DfsTiersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDfsTierPools

> DfsTierResp RemoveDfsTierPools(ctx, dfsTierId).Body(body).Execute()





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
	dfsTierId := int64(789) // int64 | dfs tier id
	body := *openapiclient.NewDfsTierRemovePoolReq(*openapiclient.NewDfsTierRemovePoolReqTier([]int64{int64(123)}, false)) // DfsTierRemovePoolReq | pools info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsTiersAPI.RemoveDfsTierPools(context.Background(), dfsTierId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsTiersAPI.RemoveDfsTierPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveDfsTierPools`: DfsTierResp
	fmt.Fprintf(os.Stdout, "Response from `DfsTiersAPI.RemoveDfsTierPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsTierId** | **int64** | dfs tier id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDfsTierPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsTierRemovePoolReq**](DfsTierRemovePoolReq.md) | pools info | 

### Return type

[**DfsTierResp**](DfsTierResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

