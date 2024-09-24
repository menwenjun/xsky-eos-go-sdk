# \RemoteClustersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRemoteCluster**](RemoteClustersAPI.md#CreateRemoteCluster) | **Post** /remote-clusters/ | 
[**DeleteRemoteCluster**](RemoteClustersAPI.md#DeleteRemoteCluster) | **Delete** /remote-clusters/{cluster_id} | 
[**GetRemoteCluster**](RemoteClustersAPI.md#GetRemoteCluster) | **Get** /remote-clusters/{cluster_id} | 
[**ListRemoteClusters**](RemoteClustersAPI.md#ListRemoteClusters) | **Get** /remote-clusters/ | 
[**UpdateRemoteCluster**](RemoteClustersAPI.md#UpdateRemoteCluster) | **Patch** /remote-clusters/{cluster_id} | 



## CreateRemoteCluster

> RemoteClusterResp CreateRemoteCluster(ctx).Body(body).Execute()





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
	body := *openapiclient.NewRemoteClusterCreateReq() // RemoteClusterCreateReq | remote cluster info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteClustersAPI.CreateRemoteCluster(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteClustersAPI.CreateRemoteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRemoteCluster`: RemoteClusterResp
	fmt.Fprintf(os.Stdout, "Response from `RemoteClustersAPI.CreateRemoteCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemoteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RemoteClusterCreateReq**](RemoteClusterCreateReq.md) | remote cluster info | 

### Return type

[**RemoteClusterResp**](RemoteClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRemoteCluster

> RemoteClusterResp DeleteRemoteCluster(ctx, clusterId).Execute()





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
	clusterId := int64(789) // int64 | remote cluster id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteClustersAPI.DeleteRemoteCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteClustersAPI.DeleteRemoteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRemoteCluster`: RemoteClusterResp
	fmt.Fprintf(os.Stdout, "Response from `RemoteClustersAPI.DeleteRemoteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int64** | remote cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemoteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoteClusterResp**](RemoteClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteCluster

> RemoteClusterResp GetRemoteCluster(ctx, clusterId).Execute()





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
	clusterId := int64(789) // int64 | remote cluster id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteClustersAPI.GetRemoteCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteClustersAPI.GetRemoteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteCluster`: RemoteClusterResp
	fmt.Fprintf(os.Stdout, "Response from `RemoteClustersAPI.GetRemoteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int64** | remote cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoteClusterResp**](RemoteClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRemoteClusters

> RemoteClustersResp ListRemoteClusters(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).FsId(fsId).Execute()





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
	fsId := "fsId_example" // string | fsid of cluster (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteClustersAPI.ListRemoteClusters(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).FsId(fsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteClustersAPI.ListRemoteClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRemoteClusters`: RemoteClustersResp
	fmt.Fprintf(os.Stdout, "Response from `RemoteClustersAPI.ListRemoteClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRemoteClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **fsId** | **string** | fsid of cluster | 

### Return type

[**RemoteClustersResp**](RemoteClustersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRemoteCluster

> RemoteClusterResp UpdateRemoteCluster(ctx, clusterId).Body(body).Execute()





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
	clusterId := int64(789) // int64 | remote cluster id
	body := *openapiclient.NewRemoteClusterUpdateReq() // RemoteClusterUpdateReq | remote cluster info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteClustersAPI.UpdateRemoteCluster(context.Background(), clusterId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteClustersAPI.UpdateRemoteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRemoteCluster`: RemoteClusterResp
	fmt.Fprintf(os.Stdout, "Response from `RemoteClustersAPI.UpdateRemoteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int64** | remote cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RemoteClusterUpdateReq**](RemoteClusterUpdateReq.md) | remote cluster info | 

### Return type

[**RemoteClusterResp**](RemoteClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

