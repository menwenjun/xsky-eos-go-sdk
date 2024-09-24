# \ActionLogsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateActionLog**](ActionLogsAPI.md#CreateActionLog) | **Post** /action-logs/ | 
[**GetActionLog**](ActionLogsAPI.md#GetActionLog) | **Get** /action-logs/{action_log_id} | 
[**GetActionLogUser**](ActionLogsAPI.md#GetActionLogUser) | **Get** /action-logs/:user | 
[**ListActionLogs**](ActionLogsAPI.md#ListActionLogs) | **Get** /action-logs/ | 
[**UpdateActionLog**](ActionLogsAPI.md#UpdateActionLog) | **Patch** /action-logs/{action_log_id} | 



## CreateActionLog

> ActionLogResp CreateActionLog(ctx).Body(body).Execute()





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
	body := *openapiclient.NewActionLogCreateReq() // ActionLogCreateReq | action log info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionLogsAPI.CreateActionLog(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionLogsAPI.CreateActionLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateActionLog`: ActionLogResp
	fmt.Fprintf(os.Stdout, "Response from `ActionLogsAPI.CreateActionLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateActionLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ActionLogCreateReq**](ActionLogCreateReq.md) | action log info | 

### Return type

[**ActionLogResp**](ActionLogResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActionLog

> ActionLogResp GetActionLog(ctx, actionLogId).Execute()





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
	actionLogId := int64(789) // int64 | action log id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionLogsAPI.GetActionLog(context.Background(), actionLogId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionLogsAPI.GetActionLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActionLog`: ActionLogResp
	fmt.Fprintf(os.Stdout, "Response from `ActionLogsAPI.GetActionLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionLogId** | **int64** | action log id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActionLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionLogResp**](ActionLogResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActionLogUser

> ActionLogUserResp GetActionLogUser(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionLogsAPI.GetActionLogUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionLogsAPI.GetActionLogUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActionLogUser`: ActionLogUserResp
	fmt.Fprintf(os.Stdout, "Response from `ActionLogsAPI.GetActionLogUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActionLogUserRequest struct via the builder pattern


### Return type

[**ActionLogUserResp**](ActionLogUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActionLogs

> ActionLogsResp ListActionLogs(ctx).Limit(limit).Offset(offset).Action(action).ResourceType(resourceType).Status(status).UserId(userId).ParentId(parentId).CreateBegin(createBegin).CreateEnd(createEnd).Q(q).RelatedResource(relatedResource).Sort(sort).ClusterId(clusterId).Execute()





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
	action := "action_example" // string | action of action logs (optional)
	resourceType := "resourceType_example" // string | resource type of action logs (optional)
	status := "status_example" // string | status of action logs (optional)
	userId := int64(789) // int64 | user id of action logs (optional)
	parentId := int64(789) // int64 | parent action log id of action logs (optional)
	createBegin := "createBegin_example" // string | create begin timestamp (optional)
	createEnd := "createEnd_example" // string | create end timestamp (optional)
	q := "q_example" // string | query param of search (optional)
	relatedResource := "relatedResource_example" // string | related resource info of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionLogsAPI.ListActionLogs(context.Background()).Limit(limit).Offset(offset).Action(action).ResourceType(resourceType).Status(status).UserId(userId).ParentId(parentId).CreateBegin(createBegin).CreateEnd(createEnd).Q(q).RelatedResource(relatedResource).Sort(sort).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionLogsAPI.ListActionLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListActionLogs`: ActionLogsResp
	fmt.Fprintf(os.Stdout, "Response from `ActionLogsAPI.ListActionLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListActionLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **action** | **string** | action of action logs | 
 **resourceType** | **string** | resource type of action logs | 
 **status** | **string** | status of action logs | 
 **userId** | **int64** | user id of action logs | 
 **parentId** | **int64** | parent action log id of action logs | 
 **createBegin** | **string** | create begin timestamp | 
 **createEnd** | **string** | create end timestamp | 
 **q** | **string** | query param of search | 
 **relatedResource** | **string** | related resource info of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 

### Return type

[**ActionLogsResp**](ActionLogsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActionLog

> ActionLogResp UpdateActionLog(ctx, actionLogId).Body(body).Execute()





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
	actionLogId := int64(789) // int64 | action log id
	body := *openapiclient.NewActionLogUpdateReq() // ActionLogUpdateReq | action log info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionLogsAPI.UpdateActionLog(context.Background(), actionLogId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionLogsAPI.UpdateActionLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateActionLog`: ActionLogResp
	fmt.Fprintf(os.Stdout, "Response from `ActionLogsAPI.UpdateActionLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionLogId** | **int64** | action log id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActionLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ActionLogUpdateReq**](ActionLogUpdateReq.md) | action log info | 

### Return type

[**ActionLogResp**](ActionLogResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

