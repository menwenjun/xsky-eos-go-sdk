# \DfsHdfsesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDfsHdfsACLs**](DfsHdfsesAPI.md#AddDfsHdfsACLs) | **Post** /dfs-hdfses/{dfs_hdfs_id}:add-acls | 
[**AddDfsHdfsIPWhiteLists**](DfsHdfsesAPI.md#AddDfsHdfsIPWhiteLists) | **Post** /dfs-hdfses/{dfs_hdfs_id}:add-ip-white-lists | 
[**AddDfsHdfsProxyUsers**](DfsHdfsesAPI.md#AddDfsHdfsProxyUsers) | **Post** /dfs-hdfses/{dfs_hdfs_id}:add-proxy-users | 
[**CreateDfsHdfs**](DfsHdfsesAPI.md#CreateDfsHdfs) | **Post** /dfs-hdfses/ | 
[**DeleteDfsHdfs**](DfsHdfsesAPI.md#DeleteDfsHdfs) | **Delete** /dfs-hdfses/{dfs_hdfs_id} | 
[**GetDfsHdfs**](DfsHdfsesAPI.md#GetDfsHdfs) | **Get** /dfs-hdfses/{dfs_hdfs_id} | 
[**ListDfsHdfses**](DfsHdfsesAPI.md#ListDfsHdfses) | **Get** /dfs-hdfses/ | 
[**RemoveDfsHdfsACLs**](DfsHdfsesAPI.md#RemoveDfsHdfsACLs) | **Post** /dfs-hdfses/{dfs_hdfs_id}:remove-acls | 
[**RemoveDfsHdfsIPWhiteLists**](DfsHdfsesAPI.md#RemoveDfsHdfsIPWhiteLists) | **Post** /dfs-hdfses/{dfs_hdfs_id}:remove-ip-white-lists | 
[**RemoveDfsHdfsProxyUsers**](DfsHdfsesAPI.md#RemoveDfsHdfsProxyUsers) | **Post** /dfs-hdfses/{dfs_hdfs_id}:remove-proxy-users | 
[**UpdateDfsHdfs**](DfsHdfsesAPI.md#UpdateDfsHdfs) | **Patch** /dfs-hdfses/{dfs_hdfs_id} | 
[**UpdateDfsHdfsACLs**](DfsHdfsesAPI.md#UpdateDfsHdfsACLs) | **Patch** /dfs-hdfses/{dfs_hdfs_id}:update-acls | 
[**UpdateDfsHdfsIPWhiteLists**](DfsHdfsesAPI.md#UpdateDfsHdfsIPWhiteLists) | **Patch** /dfs-hdfses/{dfs_hdfs_id}:update-ip-white-lists | 
[**UpdateDfsHdfsProxyUsers**](DfsHdfsesAPI.md#UpdateDfsHdfsProxyUsers) | **Patch** /dfs-hdfses/{dfs_hdfs_id}:update-proxy-users | 



## AddDfsHdfsACLs

> DfsHdfsResp AddDfsHdfsACLs(ctx, dfsHdfsId).Body(body).Execute()





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
	dfsHdfsId := int64(789) // int64 | hdfs id
	body := *openapiclient.NewDfsHdfsAddACLsReq(*openapiclient.NewDfsHdfsAddACLsReqHdfs([]openapiclient.DfsHdfsACLReq{*openapiclient.NewDfsHdfsACLReq()})) // DfsHdfsAddACLsReq | dfs hdfs info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsesAPI.AddDfsHdfsACLs(context.Background(), dfsHdfsId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsesAPI.AddDfsHdfsACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDfsHdfsACLs`: DfsHdfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsesAPI.AddDfsHdfsACLs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsHdfsId** | **int64** | hdfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDfsHdfsACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsHdfsAddACLsReq**](DfsHdfsAddACLsReq.md) | dfs hdfs info | 

### Return type

[**DfsHdfsResp**](DfsHdfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddDfsHdfsIPWhiteLists

> DfsHdfsResp AddDfsHdfsIPWhiteLists(ctx, dfsHdfsId).Body(body).Execute()





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
	dfsHdfsId := int64(789) // int64 | hdfs id
	body := *openapiclient.NewDfsHdfsAddIPWhiteListReq(*openapiclient.NewDfsHdfsAddIPWhiteListReqHdfs([]openapiclient.DfsHdfsIPWhiteList{*openapiclient.NewDfsHdfsIPWhiteList()})) // DfsHdfsAddIPWhiteListReq | dfs hdfs info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsesAPI.AddDfsHdfsIPWhiteLists(context.Background(), dfsHdfsId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsesAPI.AddDfsHdfsIPWhiteLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDfsHdfsIPWhiteLists`: DfsHdfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsesAPI.AddDfsHdfsIPWhiteLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsHdfsId** | **int64** | hdfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDfsHdfsIPWhiteListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsHdfsAddIPWhiteListReq**](DfsHdfsAddIPWhiteListReq.md) | dfs hdfs info | 

### Return type

[**DfsHdfsResp**](DfsHdfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddDfsHdfsProxyUsers

> DfsHdfsResp AddDfsHdfsProxyUsers(ctx, dfsHdfsId).Body(body).Execute()





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
	dfsHdfsId := int64(789) // int64 | hdfs id
	body := *openapiclient.NewDfsHdfsAddProxyUsersReq(*openapiclient.NewDfsHdfsAddProxyUsersReqHdfs([]openapiclient.DfsHdfsProxyUserReq{*openapiclient.NewDfsHdfsProxyUserReq()})) // DfsHdfsAddProxyUsersReq | dfs hdfs info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsesAPI.AddDfsHdfsProxyUsers(context.Background(), dfsHdfsId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsesAPI.AddDfsHdfsProxyUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDfsHdfsProxyUsers`: DfsHdfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsesAPI.AddDfsHdfsProxyUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsHdfsId** | **int64** | hdfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDfsHdfsProxyUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsHdfsAddProxyUsersReq**](DfsHdfsAddProxyUsersReq.md) | dfs hdfs info | 

### Return type

[**DfsHdfsResp**](DfsHdfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDfsHdfs

> DfsHdfsResp CreateDfsHdfs(ctx).Body(body).AllowPathCreate(allowPathCreate).Execute()





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
	body := *openapiclient.NewDfsHdfsCreateReq() // DfsHdfsCreateReq | hdfs info
	allowPathCreate := true // bool | allow create path when not existed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsesAPI.CreateDfsHdfs(context.Background()).Body(body).AllowPathCreate(allowPathCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsesAPI.CreateDfsHdfs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsHdfs`: DfsHdfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsesAPI.CreateDfsHdfs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsHdfsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsHdfsCreateReq**](DfsHdfsCreateReq.md) | hdfs info | 
 **allowPathCreate** | **bool** | allow create path when not existed | 

### Return type

[**DfsHdfsResp**](DfsHdfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsHdfs

> DfsHdfsResp DeleteDfsHdfs(ctx, dfsHdfsId).Force(force).WithDirectory(withDirectory).Execute()





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
	dfsHdfsId := int64(789) // int64 | dfs hdfs id
	force := true // bool | force delete or not (optional)
	withDirectory := true // bool | also delete directory (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsesAPI.DeleteDfsHdfs(context.Background(), dfsHdfsId).Force(force).WithDirectory(withDirectory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsesAPI.DeleteDfsHdfs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsHdfs`: DfsHdfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsesAPI.DeleteDfsHdfs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsHdfsId** | **int64** | dfs hdfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsHdfsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 
 **withDirectory** | **bool** | also delete directory | 

### Return type

[**DfsHdfsResp**](DfsHdfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsHdfs

> DfsHdfsResp GetDfsHdfs(ctx, dfsHdfsId).Execute()





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
	dfsHdfsId := int64(789) // int64 | dfs hdfs id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsesAPI.GetDfsHdfs(context.Background(), dfsHdfsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsesAPI.GetDfsHdfs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsHdfs`: DfsHdfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsesAPI.GetDfsHdfs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsHdfsId** | **int64** | dfs hdfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsHdfsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsHdfsResp**](DfsHdfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsHdfses

> DfsHdfsesResp ListDfsHdfses(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).Path(path).Auth(auth).DfsGatewayZoneId(dfsGatewayZoneId).Q(q).Sort(sort).Execute()





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
	path := "path_example" // string | related dfs path (optional)
	auth := "auth_example" // string | authType for hdfs: simple, kerberos (optional)
	dfsGatewayZoneId := int64(789) // int64 | dfs gateway zone id (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsesAPI.ListDfsHdfses(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).Path(path).Auth(auth).DfsGatewayZoneId(dfsGatewayZoneId).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsesAPI.ListDfsHdfses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsHdfses`: DfsHdfsesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsesAPI.ListDfsHdfses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsHdfsesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **path** | **string** | related dfs path | 
 **auth** | **string** | authType for hdfs: simple, kerberos | 
 **dfsGatewayZoneId** | **int64** | dfs gateway zone id | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DfsHdfsesResp**](DfsHdfsesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDfsHdfsACLs

> DfsHdfsResp RemoveDfsHdfsACLs(ctx, dfsHdfsId).Body(body).Execute()





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
	dfsHdfsId := int64(789) // int64 | dfs hdfs id
	body := *openapiclient.NewDfsHdfsRemoveACLsReq(*openapiclient.NewDfsHdfsRemoveACLsReqHdfs([]int64{int64(123)})) // DfsHdfsRemoveACLsReq | hdfs acls info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsesAPI.RemoveDfsHdfsACLs(context.Background(), dfsHdfsId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsesAPI.RemoveDfsHdfsACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveDfsHdfsACLs`: DfsHdfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsesAPI.RemoveDfsHdfsACLs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsHdfsId** | **int64** | dfs hdfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDfsHdfsACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsHdfsRemoveACLsReq**](DfsHdfsRemoveACLsReq.md) | hdfs acls info | 

### Return type

[**DfsHdfsResp**](DfsHdfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDfsHdfsIPWhiteLists

> DfsHdfsResp RemoveDfsHdfsIPWhiteLists(ctx, dfsHdfsId).Body(body).Execute()





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
	dfsHdfsId := int64(789) // int64 | dfs hdfs id
	body := *openapiclient.NewDfsHdfsRemoveIPWhiteListReq(*openapiclient.NewDfsHdfsRemoveIPWhiteListReqHdfs([]int64{int64(123)})) // DfsHdfsRemoveIPWhiteListReq | hdfs ip white list info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsesAPI.RemoveDfsHdfsIPWhiteLists(context.Background(), dfsHdfsId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsesAPI.RemoveDfsHdfsIPWhiteLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveDfsHdfsIPWhiteLists`: DfsHdfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsesAPI.RemoveDfsHdfsIPWhiteLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsHdfsId** | **int64** | dfs hdfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDfsHdfsIPWhiteListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsHdfsRemoveIPWhiteListReq**](DfsHdfsRemoveIPWhiteListReq.md) | hdfs ip white list info | 

### Return type

[**DfsHdfsResp**](DfsHdfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDfsHdfsProxyUsers

> DfsHdfsResp RemoveDfsHdfsProxyUsers(ctx, dfsHdfsId).Body(body).Execute()





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
	dfsHdfsId := int64(789) // int64 | dfs hdfs id
	body := *openapiclient.NewDfsHdfsRemoveProxyUsersReq(*openapiclient.NewDfsHdfsRemoveProxyUsersReqHdfs([]int64{int64(123)})) // DfsHdfsRemoveProxyUsersReq | hdfs proxy users info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsesAPI.RemoveDfsHdfsProxyUsers(context.Background(), dfsHdfsId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsesAPI.RemoveDfsHdfsProxyUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveDfsHdfsProxyUsers`: DfsHdfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsesAPI.RemoveDfsHdfsProxyUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsHdfsId** | **int64** | dfs hdfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDfsHdfsProxyUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsHdfsRemoveProxyUsersReq**](DfsHdfsRemoveProxyUsersReq.md) | hdfs proxy users info | 

### Return type

[**DfsHdfsResp**](DfsHdfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsHdfs

> DfsHdfsResp UpdateDfsHdfs(ctx, dfsHdfsId).Body(body).Execute()





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
	dfsHdfsId := int64(789) // int64 | hdfs id
	body := *openapiclient.NewDfsHdfsUpdateReq() // DfsHdfsUpdateReq | dfs hdfs info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsesAPI.UpdateDfsHdfs(context.Background(), dfsHdfsId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsesAPI.UpdateDfsHdfs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsHdfs`: DfsHdfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsesAPI.UpdateDfsHdfs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsHdfsId** | **int64** | hdfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsHdfsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsHdfsUpdateReq**](DfsHdfsUpdateReq.md) | dfs hdfs info | 

### Return type

[**DfsHdfsResp**](DfsHdfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsHdfsACLs

> DfsHdfsResp UpdateDfsHdfsACLs(ctx, dfsHdfsId).Body(body).Execute()





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
	dfsHdfsId := int64(789) // int64 | hdfs id
	body := *openapiclient.NewDfsHdfsUpdateACLsReq(*openapiclient.NewDfsHdfsUpdateACLsReqHdfs()) // DfsHdfsUpdateACLsReq | hdfs acls info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsesAPI.UpdateDfsHdfsACLs(context.Background(), dfsHdfsId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsesAPI.UpdateDfsHdfsACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsHdfsACLs`: DfsHdfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsesAPI.UpdateDfsHdfsACLs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsHdfsId** | **int64** | hdfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsHdfsACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsHdfsUpdateACLsReq**](DfsHdfsUpdateACLsReq.md) | hdfs acls info | 

### Return type

[**DfsHdfsResp**](DfsHdfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsHdfsIPWhiteLists

> DfsHdfsResp UpdateDfsHdfsIPWhiteLists(ctx, dfsHdfsId).Body(body).Execute()





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
	dfsHdfsId := int64(789) // int64 | hdfs id
	body := *openapiclient.NewDfsHdfsUpdateIPWhiteListsReq(*openapiclient.NewDfsHdfsUpdateIPWhiteListsReqHdfs()) // DfsHdfsUpdateIPWhiteListsReq | hdfs ip white lists info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsesAPI.UpdateDfsHdfsIPWhiteLists(context.Background(), dfsHdfsId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsesAPI.UpdateDfsHdfsIPWhiteLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsHdfsIPWhiteLists`: DfsHdfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsesAPI.UpdateDfsHdfsIPWhiteLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsHdfsId** | **int64** | hdfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsHdfsIPWhiteListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsHdfsUpdateIPWhiteListsReq**](DfsHdfsUpdateIPWhiteListsReq.md) | hdfs ip white lists info | 

### Return type

[**DfsHdfsResp**](DfsHdfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsHdfsProxyUsers

> DfsHdfsResp UpdateDfsHdfsProxyUsers(ctx, dfsHdfsId).Body(body).Execute()





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
	dfsHdfsId := int64(789) // int64 | hdfs id
	body := *openapiclient.NewDfsHdfsUpdateProxyUsersReq(*openapiclient.NewDfsHdfsUpdateProxyUsersReqHdfs()) // DfsHdfsUpdateProxyUsersReq | hdfs proxy user info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsesAPI.UpdateDfsHdfsProxyUsers(context.Background(), dfsHdfsId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsesAPI.UpdateDfsHdfsProxyUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsHdfsProxyUsers`: DfsHdfsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsesAPI.UpdateDfsHdfsProxyUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsHdfsId** | **int64** | hdfs id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsHdfsProxyUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsHdfsUpdateProxyUsersReq**](DfsHdfsUpdateProxyUsersReq.md) | hdfs proxy user info | 

### Return type

[**DfsHdfsResp**](DfsHdfsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

