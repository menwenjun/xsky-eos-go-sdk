# \DfsSmbSharesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDfsSMBShareACLs**](DfsSmbSharesAPI.md#AddDfsSMBShareACLs) | **Post** /dfs-smb-shares/{dfs_smb_share_id}:add-acls | 
[**CreateDfsSMBShare**](DfsSmbSharesAPI.md#CreateDfsSMBShare) | **Post** /dfs-smb-shares/ | 
[**DeleteDfsSMBShare**](DfsSmbSharesAPI.md#DeleteDfsSMBShare) | **Delete** /dfs-smb-shares/{dfs_smb_share_id} | 
[**GetDfsSMBShare**](DfsSmbSharesAPI.md#GetDfsSMBShare) | **Get** /dfs-smb-shares/{dfs_smb_share_id} | 
[**ListDfsSMBShares**](DfsSmbSharesAPI.md#ListDfsSMBShares) | **Get** /dfs-smb-shares/ | 
[**RemoveDfsSMBShareACLs**](DfsSmbSharesAPI.md#RemoveDfsSMBShareACLs) | **Post** /dfs-smb-shares/{dfs_smb_share_id}:remove-acls | 
[**UpdateDfsSMBShare**](DfsSmbSharesAPI.md#UpdateDfsSMBShare) | **Patch** /dfs-smb-shares/{dfs_smb_share_id} | 
[**UpdateDfsSMBShareACLs**](DfsSmbSharesAPI.md#UpdateDfsSMBShareACLs) | **Post** /dfs-smb-shares/{dfs_smb_share_id}:update-acls | 



## AddDfsSMBShareACLs

> DfsSMBShareResp AddDfsSMBShareACLs(ctx, dfsSmbShareId).Body(body).Execute()





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
	dfsSmbShareId := int64(789) // int64 | dfs smb share id
	body := *openapiclient.NewDfsSMBShareAddACLsReq(*openapiclient.NewDfsSMBShareAddACLsReqShare([]openapiclient.DfsSMBShareACLReq{*openapiclient.NewDfsSMBShareACLReq()})) // DfsSMBShareAddACLsReq | share acls info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbSharesAPI.AddDfsSMBShareACLs(context.Background(), dfsSmbShareId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbSharesAPI.AddDfsSMBShareACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDfsSMBShareACLs`: DfsSMBShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbSharesAPI.AddDfsSMBShareACLs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsSmbShareId** | **int64** | dfs smb share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDfsSMBShareACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsSMBShareAddACLsReq**](DfsSMBShareAddACLsReq.md) | share acls info | 

### Return type

[**DfsSMBShareResp**](DfsSMBShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDfsSMBShare

> DfsSMBShareResp CreateDfsSMBShare(ctx).Body(body).AllowPathCreate(allowPathCreate).Execute()





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
	body := *openapiclient.NewDfsSMBShareCreateReq(*openapiclient.NewDfsSMBShareCreateReqShare(int64(123), "Name_example", "Path_example")) // DfsSMBShareCreateReq | share info
	allowPathCreate := true // bool | allow create path when not existed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbSharesAPI.CreateDfsSMBShare(context.Background()).Body(body).AllowPathCreate(allowPathCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbSharesAPI.CreateDfsSMBShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsSMBShare`: DfsSMBShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbSharesAPI.CreateDfsSMBShare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsSMBShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsSMBShareCreateReq**](DfsSMBShareCreateReq.md) | share info | 
 **allowPathCreate** | **bool** | allow create path when not existed | 

### Return type

[**DfsSMBShareResp**](DfsSMBShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsSMBShare

> DfsSMBShareResp DeleteDfsSMBShare(ctx, dfsSmbShareId).Force(force).WithDirectory(withDirectory).Execute()





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
	dfsSmbShareId := int64(789) // int64 | share id
	force := true // bool | force delete or not (optional)
	withDirectory := true // bool | also delete directory (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbSharesAPI.DeleteDfsSMBShare(context.Background(), dfsSmbShareId).Force(force).WithDirectory(withDirectory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbSharesAPI.DeleteDfsSMBShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsSMBShare`: DfsSMBShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbSharesAPI.DeleteDfsSMBShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsSmbShareId** | **int64** | share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsSMBShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 
 **withDirectory** | **bool** | also delete directory | 

### Return type

[**DfsSMBShareResp**](DfsSMBShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsSMBShare

> DfsSMBShareResp GetDfsSMBShare(ctx, dfsSmbShareId).Execute()





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
	dfsSmbShareId := int64(789) // int64 | share id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbSharesAPI.GetDfsSMBShare(context.Background(), dfsSmbShareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbSharesAPI.GetDfsSMBShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsSMBShare`: DfsSMBShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbSharesAPI.GetDfsSMBShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsSmbShareId** | **int64** | share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsSMBShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsSMBShareResp**](DfsSMBShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsSMBShares

> DfsSMBSharesResp ListDfsSMBShares(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).DfsRootfsId(dfsRootfsId).DfsGatewayGroupId(dfsGatewayGroupId).Q(q).Path(path).Sort(sort).Execute()





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
	dfsRootfsId := int64(789) // int64 | dfs rootfs id (optional)
	dfsGatewayGroupId := int64(789) // int64 | dfs gateway group id (optional)
	q := "q_example" // string | query param of search (optional)
	path := "path_example" // string | related dfs path (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbSharesAPI.ListDfsSMBShares(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).DfsRootfsId(dfsRootfsId).DfsGatewayGroupId(dfsGatewayGroupId).Q(q).Path(path).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbSharesAPI.ListDfsSMBShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsSMBShares`: DfsSMBSharesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbSharesAPI.ListDfsSMBShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsSMBSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **dfsRootfsId** | **int64** | dfs rootfs id | 
 **dfsGatewayGroupId** | **int64** | dfs gateway group id | 
 **q** | **string** | query param of search | 
 **path** | **string** | related dfs path | 
 **sort** | **string** | sort param of search | 

### Return type

[**DfsSMBSharesResp**](DfsSMBSharesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDfsSMBShareACLs

> DfsSMBShareResp RemoveDfsSMBShareACLs(ctx, dfsSmbShareId).Body(body).Execute()





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
	dfsSmbShareId := int64(789) // int64 | dfs smb share id
	body := *openapiclient.NewDfsSMBShareRemoveACLsReq(*openapiclient.NewDfsSMBShareRemoveACLsReqShare([]int64{int64(123)})) // DfsSMBShareRemoveACLsReq | share acls info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbSharesAPI.RemoveDfsSMBShareACLs(context.Background(), dfsSmbShareId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbSharesAPI.RemoveDfsSMBShareACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveDfsSMBShareACLs`: DfsSMBShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbSharesAPI.RemoveDfsSMBShareACLs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsSmbShareId** | **int64** | dfs smb share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDfsSMBShareACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsSMBShareRemoveACLsReq**](DfsSMBShareRemoveACLsReq.md) | share acls info | 

### Return type

[**DfsSMBShareResp**](DfsSMBShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsSMBShare

> DfsSMBShareResp UpdateDfsSMBShare(ctx, dfsSmbShareId).Body(body).Execute()





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
	dfsSmbShareId := int64(789) // int64 | share id
	body := *openapiclient.NewDfsSMBShareUpdateReq(*openapiclient.NewDfsSMBShareUpdateReqShare()) // DfsSMBShareUpdateReq | share info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbSharesAPI.UpdateDfsSMBShare(context.Background(), dfsSmbShareId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbSharesAPI.UpdateDfsSMBShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsSMBShare`: DfsSMBShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbSharesAPI.UpdateDfsSMBShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsSmbShareId** | **int64** | share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsSMBShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsSMBShareUpdateReq**](DfsSMBShareUpdateReq.md) | share info | 

### Return type

[**DfsSMBShareResp**](DfsSMBShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsSMBShareACLs

> DfsSMBShareResp UpdateDfsSMBShareACLs(ctx, dfsSmbShareId).Body(body).Execute()





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
	dfsSmbShareId := int64(789) // int64 | share id
	body := *openapiclient.NewDfsSMBShareUpdateACLsReq(*openapiclient.NewDfsSMBShareUpdateACLsReqShare()) // DfsSMBShareUpdateACLsReq | share acls info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbSharesAPI.UpdateDfsSMBShareACLs(context.Background(), dfsSmbShareId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbSharesAPI.UpdateDfsSMBShareACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsSMBShareACLs`: DfsSMBShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbSharesAPI.UpdateDfsSMBShareACLs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsSmbShareId** | **int64** | share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsSMBShareACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsSMBShareUpdateACLsReq**](DfsSMBShareUpdateACLsReq.md) | share acls info | 

### Return type

[**DfsSMBShareResp**](DfsSMBShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

