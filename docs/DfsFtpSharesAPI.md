# \DfsFtpSharesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDfsFTPShareACLs**](DfsFtpSharesAPI.md#AddDfsFTPShareACLs) | **Post** /dfs-ftp-shares/{dfs_ftp_share_id}:add-acls | 
[**CreateDfsFTPShare**](DfsFtpSharesAPI.md#CreateDfsFTPShare) | **Post** /dfs-ftp-shares/ | 
[**DeleteDfsFTPShare**](DfsFtpSharesAPI.md#DeleteDfsFTPShare) | **Delete** /dfs-ftp-shares/{dfs_ftp_share_id} | 
[**GetDfsFTPShare**](DfsFtpSharesAPI.md#GetDfsFTPShare) | **Get** /dfs-ftp-shares/{dfs_ftp_share_id} | 
[**ListDfsFTPShares**](DfsFtpSharesAPI.md#ListDfsFTPShares) | **Get** /dfs-ftp-shares/ | 
[**RemoveDfsFTPShareACLs**](DfsFtpSharesAPI.md#RemoveDfsFTPShareACLs) | **Post** /dfs-ftp-shares/{dfs_ftp_share_id}:remove-acls | 
[**UpdateDfsFTPShare**](DfsFtpSharesAPI.md#UpdateDfsFTPShare) | **Patch** /dfs-ftp-shares/{dfs_ftp_share_id} | 
[**UpdateDfsFTPShareACLs**](DfsFtpSharesAPI.md#UpdateDfsFTPShareACLs) | **Post** /dfs-ftp-shares/{dfs_ftp_share_id}:update-acls | 



## AddDfsFTPShareACLs

> DfsFTPShareResp AddDfsFTPShareACLs(ctx, dfsFtpShareId).Body(body).Execute()





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
	dfsFtpShareId := int64(789) // int64 | dfs ftp share id
	body := *openapiclient.NewDfsFTPShareAddACLsReq(*openapiclient.NewDfsFTPShareAddACLsReqShare([]openapiclient.DfsFTPShareACLReq{*openapiclient.NewDfsFTPShareACLReq()})) // DfsFTPShareAddACLsReq | ftp share acls info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsFtpSharesAPI.AddDfsFTPShareACLs(context.Background(), dfsFtpShareId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFtpSharesAPI.AddDfsFTPShareACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDfsFTPShareACLs`: DfsFTPShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsFtpSharesAPI.AddDfsFTPShareACLs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsFtpShareId** | **int64** | dfs ftp share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDfsFTPShareACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsFTPShareAddACLsReq**](DfsFTPShareAddACLsReq.md) | ftp share acls info | 

### Return type

[**DfsFTPShareResp**](DfsFTPShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDfsFTPShare

> DfsFTPShareResp CreateDfsFTPShare(ctx).Body(body).AllowPathCreate(allowPathCreate).Execute()





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
	body := *openapiclient.NewDfsFTPShareCreateReq(*openapiclient.NewDfsFTPShareCreateReqShare(int64(123), "Name_example", "Path_example")) // DfsFTPShareCreateReq | share info
	allowPathCreate := true // bool | allow create path when not existed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsFtpSharesAPI.CreateDfsFTPShare(context.Background()).Body(body).AllowPathCreate(allowPathCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFtpSharesAPI.CreateDfsFTPShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsFTPShare`: DfsFTPShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsFtpSharesAPI.CreateDfsFTPShare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsFTPShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsFTPShareCreateReq**](DfsFTPShareCreateReq.md) | share info | 
 **allowPathCreate** | **bool** | allow create path when not existed | 

### Return type

[**DfsFTPShareResp**](DfsFTPShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsFTPShare

> DfsFTPShareResp DeleteDfsFTPShare(ctx, dfsFtpShareId).Force(force).WithDirectory(withDirectory).Execute()





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
	dfsFtpShareId := int64(789) // int64 | share id
	force := true // bool | force delete or not (optional)
	withDirectory := true // bool | also delete directory (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsFtpSharesAPI.DeleteDfsFTPShare(context.Background(), dfsFtpShareId).Force(force).WithDirectory(withDirectory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFtpSharesAPI.DeleteDfsFTPShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsFTPShare`: DfsFTPShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsFtpSharesAPI.DeleteDfsFTPShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsFtpShareId** | **int64** | share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsFTPShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 
 **withDirectory** | **bool** | also delete directory | 

### Return type

[**DfsFTPShareResp**](DfsFTPShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsFTPShare

> DfsFTPShareResp GetDfsFTPShare(ctx, dfsFtpShareId).Execute()





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
	dfsFtpShareId := int64(789) // int64 | share id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsFtpSharesAPI.GetDfsFTPShare(context.Background(), dfsFtpShareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFtpSharesAPI.GetDfsFTPShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsFTPShare`: DfsFTPShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsFtpSharesAPI.GetDfsFTPShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsFtpShareId** | **int64** | share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsFTPShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsFTPShareResp**](DfsFTPShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsFTPShares

> DfsFTPSharesResp ListDfsFTPShares(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).DfsRootfsId(dfsRootfsId).Path(path).DfsGatewayGroupId(dfsGatewayGroupId).Q(q).Sort(sort).Execute()





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
	path := "path_example" // string | related dfs path (optional)
	dfsGatewayGroupId := int64(789) // int64 | dfs gateway group id (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsFtpSharesAPI.ListDfsFTPShares(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).DfsRootfsId(dfsRootfsId).Path(path).DfsGatewayGroupId(dfsGatewayGroupId).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFtpSharesAPI.ListDfsFTPShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsFTPShares`: DfsFTPSharesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsFtpSharesAPI.ListDfsFTPShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsFTPSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **dfsRootfsId** | **int64** | dfs rootfs id | 
 **path** | **string** | related dfs path | 
 **dfsGatewayGroupId** | **int64** | dfs gateway group id | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DfsFTPSharesResp**](DfsFTPSharesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDfsFTPShareACLs

> DfsFTPShareResp RemoveDfsFTPShareACLs(ctx, dfsFtpShareId).Body(body).Execute()





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
	dfsFtpShareId := int64(789) // int64 | dfs ftp share id
	body := *openapiclient.NewDfsFTPShareRemoveACLsReq(*openapiclient.NewDfsFTPShareRemoveACLsReqShare([]int64{int64(123)})) // DfsFTPShareRemoveACLsReq | share acls info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsFtpSharesAPI.RemoveDfsFTPShareACLs(context.Background(), dfsFtpShareId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFtpSharesAPI.RemoveDfsFTPShareACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveDfsFTPShareACLs`: DfsFTPShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsFtpSharesAPI.RemoveDfsFTPShareACLs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsFtpShareId** | **int64** | dfs ftp share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDfsFTPShareACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsFTPShareRemoveACLsReq**](DfsFTPShareRemoveACLsReq.md) | share acls info | 

### Return type

[**DfsFTPShareResp**](DfsFTPShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsFTPShare

> DfsFTPShareResp UpdateDfsFTPShare(ctx, dfsFtpShareId).Body(body).Execute()





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
	dfsFtpShareId := int64(789) // int64 | share id
	body := *openapiclient.NewDfsFTPShareUpdateReq(*openapiclient.NewDfsFTPShareUpdateReqShare()) // DfsFTPShareUpdateReq | share info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsFtpSharesAPI.UpdateDfsFTPShare(context.Background(), dfsFtpShareId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFtpSharesAPI.UpdateDfsFTPShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsFTPShare`: DfsFTPShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsFtpSharesAPI.UpdateDfsFTPShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsFtpShareId** | **int64** | share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsFTPShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsFTPShareUpdateReq**](DfsFTPShareUpdateReq.md) | share info | 

### Return type

[**DfsFTPShareResp**](DfsFTPShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsFTPShareACLs

> DfsFTPShareResp UpdateDfsFTPShareACLs(ctx, dfsFtpShareId).Body(body).Execute()





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
	dfsFtpShareId := int64(789) // int64 | ftp share id
	body := *openapiclient.NewDfsFTPShareUpdateACLsReq(*openapiclient.NewDfsFTPShareUpdateACLsReqShare()) // DfsFTPShareUpdateACLsReq | ftp share acls info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsFtpSharesAPI.UpdateDfsFTPShareACLs(context.Background(), dfsFtpShareId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFtpSharesAPI.UpdateDfsFTPShareACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsFTPShareACLs`: DfsFTPShareResp
	fmt.Fprintf(os.Stdout, "Response from `DfsFtpSharesAPI.UpdateDfsFTPShareACLs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsFtpShareId** | **int64** | ftp share id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsFTPShareACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsFTPShareUpdateACLsReq**](DfsFTPShareUpdateACLsReq.md) | ftp share acls info | 

### Return type

[**DfsFTPShareResp**](DfsFTPShareResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

