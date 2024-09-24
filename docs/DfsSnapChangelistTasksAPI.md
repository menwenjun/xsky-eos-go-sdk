# \DfsSnapChangelistTasksAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDfsSnapChangelistTask**](DfsSnapChangelistTasksAPI.md#CreateDfsSnapChangelistTask) | **Post** /dfs-snap-changelist-tasks/ | 
[**DeleteDfsSnapChangelistTask**](DfsSnapChangelistTasksAPI.md#DeleteDfsSnapChangelistTask) | **Delete** /dfs-snap-changelist-tasks/{dfs_snap_changelist_task_id} | 
[**GetDfsSnapChangelistTask**](DfsSnapChangelistTasksAPI.md#GetDfsSnapChangelistTask) | **Get** /dfs-snap-changelist-tasks/{dfs_snap_changelist_task_id} | 
[**ListDfsSnapChangelistTasks**](DfsSnapChangelistTasksAPI.md#ListDfsSnapChangelistTasks) | **Get** /dfs-snap-changelist-tasks/ | 



## CreateDfsSnapChangelistTask

> DfsSnapChangelistTaskResp CreateDfsSnapChangelistTask(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsSnapChangelistTaskCreateReq(*openapiclient.NewDfsSnapChangelistTaskCreateReqSnapChangelistTask(int64(123), "Path_example")) // DfsSnapChangelistTaskCreateReq | task info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSnapChangelistTasksAPI.CreateDfsSnapChangelistTask(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSnapChangelistTasksAPI.CreateDfsSnapChangelistTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsSnapChangelistTask`: DfsSnapChangelistTaskResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSnapChangelistTasksAPI.CreateDfsSnapChangelistTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsSnapChangelistTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsSnapChangelistTaskCreateReq**](DfsSnapChangelistTaskCreateReq.md) | task info | 

### Return type

[**DfsSnapChangelistTaskResp**](DfsSnapChangelistTaskResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsSnapChangelistTask

> DeleteDfsSnapChangelistTask(ctx, dfsSnapChangelistTaskId).Execute()





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
	dfsSnapChangelistTaskId := int64(789) // int64 | task id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DfsSnapChangelistTasksAPI.DeleteDfsSnapChangelistTask(context.Background(), dfsSnapChangelistTaskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSnapChangelistTasksAPI.DeleteDfsSnapChangelistTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsSnapChangelistTaskId** | **int64** | task id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsSnapChangelistTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsSnapChangelistTask

> DfsSnapChangelistTaskResp GetDfsSnapChangelistTask(ctx, dfsSnapChangelistTaskId).Execute()





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
	dfsSnapChangelistTaskId := int64(789) // int64 | the dfs snap change list task id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSnapChangelistTasksAPI.GetDfsSnapChangelistTask(context.Background(), dfsSnapChangelistTaskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSnapChangelistTasksAPI.GetDfsSnapChangelistTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsSnapChangelistTask`: DfsSnapChangelistTaskResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSnapChangelistTasksAPI.GetDfsSnapChangelistTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsSnapChangelistTaskId** | **int64** | the dfs snap change list task id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsSnapChangelistTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsSnapChangelistTaskResp**](DfsSnapChangelistTaskResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsSnapChangelistTasks

> DfsSnapChangelistTasksResp ListDfsSnapChangelistTasks(ctx).ClusterId(clusterId).Path(path).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	path := "path_example" // string | related dfs path (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSnapChangelistTasksAPI.ListDfsSnapChangelistTasks(context.Background()).ClusterId(clusterId).Path(path).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSnapChangelistTasksAPI.ListDfsSnapChangelistTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsSnapChangelistTasks`: DfsSnapChangelistTasksResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSnapChangelistTasksAPI.ListDfsSnapChangelistTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsSnapChangelistTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **path** | **string** | related dfs path | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DfsSnapChangelistTasksResp**](DfsSnapChangelistTasksResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

