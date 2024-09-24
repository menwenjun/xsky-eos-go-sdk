# \TasksAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTask**](TasksAPI.md#GetTask) | **Get** /tasks/{task_id} | 
[**ListTasks**](TasksAPI.md#ListTasks) | **Get** /tasks/ | 



## GetTask

> TaskResp GetTask(ctx, taskId).Execute()





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
	taskId := int64(789) // int64 | the task id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.GetTask(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.GetTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTask`: TaskResp
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int64** | the task id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskResp**](TaskResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTasks

> TasksResp ListTasks(ctx).Status(status).Type_(type_).HostId(hostId).Execute()





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
	status := "status_example" // string | task status (optional)
	type_ := "type__example" // string | task type (optional)
	hostId := int64(789) // int64 | host running the task (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.ListTasks(context.Background()).Status(status).Type_(type_).HostId(hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ListTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTasks`: TasksResp
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ListTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | task status | 
 **type_** | **string** | task type | 
 **hostId** | **int64** | host running the task | 

### Return type

[**TasksResp**](TasksResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

