# \DfsNfsSharesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDfsNFSShareACLs**](DfsNfsSharesAPI.md#AddDfsNFSShareACLs) | **Post** /dfs-nfs-shares/{dfs_nfs_share_id}:add-acls | 
[**CreateDfsNFSShare**](DfsNfsSharesAPI.md#CreateDfsNFSShare) | **Post** /dfs-nfs-shares/ | 
[**DeleteDfsNFSShare**](DfsNfsSharesAPI.md#DeleteDfsNFSShare) | **Delete** /dfs-nfs-shares/{dfs_nfs_share_id} | 
[**GetDfsNFSShare**](DfsNfsSharesAPI.md#GetDfsNFSShare) | **Get** /dfs-nfs-shares/{dfs_nfs_share_id} | 
[**ListDfsNFSShares**](DfsNfsSharesAPI.md#ListDfsNFSShares) | **Get** /dfs-nfs-shares/ | 
[**RemoveDfsNFSShareACLs**](DfsNfsSharesAPI.md#RemoveDfsNFSShareACLs) | **Post** /dfs-nfs-shares/{dfs_nfs_share_id}:remove-acls | 
[**SetDfsNFSShareACLs**](DfsNfsSharesAPI.md#SetDfsNFSShareACLs) | **Post** /dfs-nfs-shares/{dfs_nfs_share_id}:set-acls | 
[**UpdateDfsNFSShare**](DfsNfsSharesAPI.md#UpdateDfsNFSShare) | **Patch** /dfs-nfs-shares/{dfs_nfs_share_id} | 
[**UpdateDfsNFSShareACLs**](DfsNfsSharesAPI.md#UpdateDfsNFSShareACLs) | **Post** /dfs-nfs-shares/{dfs_nfs_share_id}:update-acls | 



## AddDfsNFSShareACLs

> DfsNFSShareResp AddDfsNFSShareACLs(ctx, dfsNfsShareId).Body(body).Execute()





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
	dfsNfsShareId := int64(789) // int64 | dfs nfs shares id
	body := *openapiclient.NewDfsNFSShareAddACLsReq(*openapiclient.NewDfsNFSShareAddACLsReqShare([]openapiclient.DfsNFSShareACLReq{*openapiclient.NewDfsNFSShareACLReq()})) // DfsNFSShareAddACLsReq | share acls info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsNfsSharesAPI.AddDfsNFSShareACLs(context.Background(), dfsNfsShareId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsNfsSharesAPI.AddDfsNFSShareACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDfsNFSShareACLs`: DfsNFSShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsNfsSharesAPI.AddDfsNFSShareACLs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsNfsShareId** | **int64** | dfs nfs shares id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDfsNFSShareACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsNFSShareAddACLsReq**](DfsNFSShareAddACLsReq.md) | share acls info | 

### Return type

[**DfsNFSShareResp**](DfsNFSShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDfsNFSShare

> DfsNFSShareResp CreateDfsNFSShare(ctx).Body(body).AllowPathCreate(allowPathCreate).Execute()





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
	body := *openapiclient.NewDfsNFSShareCreateReq(*openapiclient.NewDfsNFSShareCreateReqShare(int64(123), "Name_example", "Path_example")) // DfsNFSShareCreateReq | share info
	allowPathCreate := true // bool | allow create path when not existed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsNfsSharesAPI.CreateDfsNFSShare(context.Background()).Body(body).AllowPathCreate(allowPathCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsNfsSharesAPI.CreateDfsNFSShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsNFSShare`: DfsNFSShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsNfsSharesAPI.CreateDfsNFSShare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsNFSShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsNFSShareCreateReq**](DfsNFSShareCreateReq.md) | share info | 
 **allowPathCreate** | **bool** | allow create path when not existed | 

### Return type

[**DfsNFSShareResp**](DfsNFSShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsNFSShare

> DfsNFSShareResp DeleteDfsNFSShare(ctx, dfsNfsShareId).Force(force).WithDirectory(withDirectory).Execute()





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
	dfsNfsShareId := int64(789) // int64 | share id
	force := true // bool | force delete or not (optional)
	withDirectory := true // bool | also delete directory (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsNfsSharesAPI.DeleteDfsNFSShare(context.Background(), dfsNfsShareId).Force(force).WithDirectory(withDirectory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsNfsSharesAPI.DeleteDfsNFSShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsNFSShare`: DfsNFSShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsNfsSharesAPI.DeleteDfsNFSShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsNfsShareId** | **int64** | share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsNFSShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 
 **withDirectory** | **bool** | also delete directory | 

### Return type

[**DfsNFSShareResp**](DfsNFSShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsNFSShare

> DfsNFSShareResp GetDfsNFSShare(ctx, dfsNfsShareId).Execute()





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
	dfsNfsShareId := int64(789) // int64 | share id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsNfsSharesAPI.GetDfsNFSShare(context.Background(), dfsNfsShareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsNfsSharesAPI.GetDfsNFSShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsNFSShare`: DfsNFSShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsNfsSharesAPI.GetDfsNFSShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsNfsShareId** | **int64** | share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsNFSShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsNFSShareResp**](DfsNFSShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsNFSShares

> DfsNFSSharesResp ListDfsNFSShares(ctx).Limit(limit).Offset(offset).DfsRootfsId(dfsRootfsId).Path(path).DfsGatewayGroupId(dfsGatewayGroupId).NfsVersions(nfsVersions).Q(q).Sort(sort).ClusterId(clusterId).Execute()





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
	dfsRootfsId := int64(789) // int64 | dfs rootfs id (optional)
	path := "path_example" // string | related dfs path (optional)
	dfsGatewayGroupId := int64(789) // int64 | dfs gateway group id (optional)
	nfsVersions := "nfsVersions_example" // string | nfs share version (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsNfsSharesAPI.ListDfsNFSShares(context.Background()).Limit(limit).Offset(offset).DfsRootfsId(dfsRootfsId).Path(path).DfsGatewayGroupId(dfsGatewayGroupId).NfsVersions(nfsVersions).Q(q).Sort(sort).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsNfsSharesAPI.ListDfsNFSShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsNFSShares`: DfsNFSSharesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsNfsSharesAPI.ListDfsNFSShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsNFSSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **dfsRootfsId** | **int64** | dfs rootfs id | 
 **path** | **string** | related dfs path | 
 **dfsGatewayGroupId** | **int64** | dfs gateway group id | 
 **nfsVersions** | **string** | nfs share version | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 

### Return type

[**DfsNFSSharesResp**](DfsNFSSharesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDfsNFSShareACLs

> DfsNFSShareResp RemoveDfsNFSShareACLs(ctx, dfsNfsShareId).Body(body).Execute()





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
	dfsNfsShareId := int64(789) // int64 | dfs nfs shares id
	body := *openapiclient.NewDfsNFSShareRemoveACLsReq(*openapiclient.NewDfsNFSShareRemoveACLsReqShare([]int64{int64(123)})) // DfsNFSShareRemoveACLsReq | share acls info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsNfsSharesAPI.RemoveDfsNFSShareACLs(context.Background(), dfsNfsShareId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsNfsSharesAPI.RemoveDfsNFSShareACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveDfsNFSShareACLs`: DfsNFSShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsNfsSharesAPI.RemoveDfsNFSShareACLs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsNfsShareId** | **int64** | dfs nfs shares id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDfsNFSShareACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsNFSShareRemoveACLsReq**](DfsNFSShareRemoveACLsReq.md) | share acls info | 

### Return type

[**DfsNFSShareResp**](DfsNFSShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDfsNFSShareACLs

> DfsNFSShareResp SetDfsNFSShareACLs(ctx, dfsNfsShareId).Body(body).Execute()





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
	dfsNfsShareId := int64(789) // int64 | share id
	body := *openapiclient.NewDfsNFSShareSetACLsReq(*openapiclient.NewDfsNFSShareSetACLsReqShare([]openapiclient.DfsNFSShareACLReq{*openapiclient.NewDfsNFSShareACLReq()})) // DfsNFSShareSetACLsReq | share info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsNfsSharesAPI.SetDfsNFSShareACLs(context.Background(), dfsNfsShareId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsNfsSharesAPI.SetDfsNFSShareACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDfsNFSShareACLs`: DfsNFSShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsNfsSharesAPI.SetDfsNFSShareACLs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsNfsShareId** | **int64** | share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDfsNFSShareACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsNFSShareSetACLsReq**](DfsNFSShareSetACLsReq.md) | share info | 

### Return type

[**DfsNFSShareResp**](DfsNFSShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsNFSShare

> DfsNFSShareResp UpdateDfsNFSShare(ctx, dfsNfsShareId).Body(body).Execute()





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
	dfsNfsShareId := int64(789) // int64 | share id
	body := *openapiclient.NewDfsNFSShareUpdateReq(*openapiclient.NewDfsNFSShareUpdateReqShare()) // DfsNFSShareUpdateReq | share info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsNfsSharesAPI.UpdateDfsNFSShare(context.Background(), dfsNfsShareId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsNfsSharesAPI.UpdateDfsNFSShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsNFSShare`: DfsNFSShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsNfsSharesAPI.UpdateDfsNFSShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsNfsShareId** | **int64** | share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsNFSShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsNFSShareUpdateReq**](DfsNFSShareUpdateReq.md) | share info | 

### Return type

[**DfsNFSShareResp**](DfsNFSShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsNFSShareACLs

> DfsNFSShareResp UpdateDfsNFSShareACLs(ctx, dfsNfsShareId).Body(body).Execute()





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
	dfsNfsShareId := int64(789) // int64 | share id
	body := *openapiclient.NewDfsNFSShareUpdateACLsReq(*openapiclient.NewDfsNFSShareUpdateACLsReqShare()) // DfsNFSShareUpdateACLsReq | share info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsNfsSharesAPI.UpdateDfsNFSShareACLs(context.Background(), dfsNfsShareId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsNfsSharesAPI.UpdateDfsNFSShareACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsNFSShareACLs`: DfsNFSShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsNfsSharesAPI.UpdateDfsNFSShareACLs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsNfsShareId** | **int64** | share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsNFSShareACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsNFSShareUpdateACLsReq**](DfsNFSShareUpdateACLsReq.md) | share info | 

### Return type

[**DfsNFSShareResp**](DfsNFSShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

