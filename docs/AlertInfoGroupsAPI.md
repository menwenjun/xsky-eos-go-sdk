# \AlertInfoGroupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AckAlertInfoGroup**](AlertInfoGroupsAPI.md#AckAlertInfoGroup) | **Post** /alert-info-groups/{alert_info_group_id}:ack | 
[**CountAlertInfoGroups**](AlertInfoGroupsAPI.md#CountAlertInfoGroups) | **Get** /alert-info-groups/stats | 
[**DeleteAlertInfoGroup**](AlertInfoGroupsAPI.md#DeleteAlertInfoGroup) | **Delete** /alert-info-groups/{alert_info_group_id} | 
[**GetAlertInfoGroup**](AlertInfoGroupsAPI.md#GetAlertInfoGroup) | **Get** /alert-info-groups/{group_id} | 
[**GetAlertInfoGroupsReport**](AlertInfoGroupsAPI.md#GetAlertInfoGroupsReport) | **Get** /alert-info-groups/report | 
[**ListAlertInfoGroups**](AlertInfoGroupsAPI.md#ListAlertInfoGroups) | **Get** /alert-info-groups/ | 
[**ResolveAlertInfoGroup**](AlertInfoGroupsAPI.md#ResolveAlertInfoGroup) | **Post** /alert-info-groups/{alert_info_group_id}:resolve | 



## AckAlertInfoGroup

> AlertInfoGroupResp AckAlertInfoGroup(ctx, alertInfoGroupId).Execute()





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
	alertInfoGroupId := int64(789) // int64 | the id of alert info group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertInfoGroupsAPI.AckAlertInfoGroup(context.Background(), alertInfoGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfoGroupsAPI.AckAlertInfoGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AckAlertInfoGroup`: AlertInfoGroupResp
	fmt.Fprintf(os.Stdout, "Response from `AlertInfoGroupsAPI.AckAlertInfoGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertInfoGroupId** | **int64** | the id of alert info group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAckAlertInfoGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertInfoGroupResp**](AlertInfoGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountAlertInfoGroups

> AlertStatsResp CountAlertInfoGroups(ctx).Acked(acked).Resolved(resolved).ResourceType(resourceType).ResourceId(resourceId).StartTime(startTime).EndTime(endTime).Execute()





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
	acked := true // bool | acked of the most recently alert info in group (optional)
	resolved := true // bool | resolved or not of the most recently alert info in group (optional)
	resourceType := "resourceType_example" // string | resource type of alert info group (optional)
	resourceId := int64(789) // int64 | resource id of alert info group (optional)
	startTime := "startTime_example" // string | start time of create of alert info group (optional)
	endTime := "endTime_example" // string | end time of create of alert info group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertInfoGroupsAPI.CountAlertInfoGroups(context.Background()).Acked(acked).Resolved(resolved).ResourceType(resourceType).ResourceId(resourceId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfoGroupsAPI.CountAlertInfoGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountAlertInfoGroups`: AlertStatsResp
	fmt.Fprintf(os.Stdout, "Response from `AlertInfoGroupsAPI.CountAlertInfoGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountAlertInfoGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acked** | **bool** | acked of the most recently alert info in group | 
 **resolved** | **bool** | resolved or not of the most recently alert info in group | 
 **resourceType** | **string** | resource type of alert info group | 
 **resourceId** | **int64** | resource id of alert info group | 
 **startTime** | **string** | start time of create of alert info group | 
 **endTime** | **string** | end time of create of alert info group | 

### Return type

[**AlertStatsResp**](AlertStatsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertInfoGroup

> AlertInfoGroupResp DeleteAlertInfoGroup(ctx, alertInfoGroupId).Execute()





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
	alertInfoGroupId := int64(789) // int64 | the id of alert info group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertInfoGroupsAPI.DeleteAlertInfoGroup(context.Background(), alertInfoGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfoGroupsAPI.DeleteAlertInfoGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlertInfoGroup`: AlertInfoGroupResp
	fmt.Fprintf(os.Stdout, "Response from `AlertInfoGroupsAPI.DeleteAlertInfoGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertInfoGroupId** | **int64** | the id of alert info group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertInfoGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertInfoGroupResp**](AlertInfoGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertInfoGroup

> AlertInfoGroupResp GetAlertInfoGroup(ctx, groupId).Execute()





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
	groupId := int64(789) // int64 | the id of alert info group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertInfoGroupsAPI.GetAlertInfoGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfoGroupsAPI.GetAlertInfoGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertInfoGroup`: AlertInfoGroupResp
	fmt.Fprintf(os.Stdout, "Response from `AlertInfoGroupsAPI.GetAlertInfoGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | the id of alert info group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertInfoGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertInfoGroupResp**](AlertInfoGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertInfoGroupsReport

> GetAlertInfoGroupsReport(ctx).Level(level).ResourceType(resourceType).ResourceId(resourceId).CreateAfter(createAfter).Acked(acked).Resolved(resolved).Group(group).ExcludeCluster(excludeCluster).IsGlobal(isGlobal).Execute()





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
	level := "level_example" // string | level of alert info group (optional)
	resourceType := "resourceType_example" // string | resource type of alert info group (optional)
	resourceId := int64(789) // int64 | resource ids of alert info group (optional)
	createAfter := "createAfter_example" // string | create_after timestamp of alert info group (optional)
	acked := true // bool | acked of alert info (optional)
	resolved := true // bool | resolved or not of alert info (optional)
	group := "group_example" // string | group of alert info (optional)
	excludeCluster := true // bool | filter to exclude cluster of alert info groups, deprecated, use `is_global` instead. (optional)
	isGlobal := true // bool | filter global alert info groups(exclude alert info groups in any cluster) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertInfoGroupsAPI.GetAlertInfoGroupsReport(context.Background()).Level(level).ResourceType(resourceType).ResourceId(resourceId).CreateAfter(createAfter).Acked(acked).Resolved(resolved).Group(group).ExcludeCluster(excludeCluster).IsGlobal(isGlobal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfoGroupsAPI.GetAlertInfoGroupsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertInfoGroupsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **level** | **string** | level of alert info group | 
 **resourceType** | **string** | resource type of alert info group | 
 **resourceId** | **int64** | resource ids of alert info group | 
 **createAfter** | **string** | create_after timestamp of alert info group | 
 **acked** | **bool** | acked of alert info | 
 **resolved** | **bool** | resolved or not of alert info | 
 **group** | **string** | group of alert info | 
 **excludeCluster** | **bool** | filter to exclude cluster of alert info groups, deprecated, use &#x60;is_global&#x60; instead. | 
 **isGlobal** | **bool** | filter global alert info groups(exclude alert info groups in any cluster) | 

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


## ListAlertInfoGroups

> AlertInfoGroupsResp ListAlertInfoGroups(ctx).QMust(qMust).Q(q).RelatedResource(relatedResource).Sort(sort).Limit(limit).Offset(offset).Level(level).ResourceType(resourceType).ResourceId(resourceId).CreateAfter(createAfter).CreateBefore(createBefore).Acked(acked).Resolved(resolved).Group(group).ExcludeCluster(excludeCluster).IsGlobal(isGlobal).Execute()





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
	qMust := "qMust_example" // string | must query param of search (optional)
	q := "q_example" // string | should query param of search (optional)
	relatedResource := "relatedResource_example" // string | should query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	level := "level_example" // string | level of alert info group (optional)
	resourceType := "resourceType_example" // string | resource type of alert info group (optional)
	resourceId := int64(789) // int64 | resource id of alert info group (optional)
	createAfter := "createAfter_example" // string | create_after start time of alert info group (optional)
	createBefore := "createBefore_example" // string | create_before end time of alert info group (optional)
	acked := true // bool | acked of alert info (optional)
	resolved := true // bool | resolved or not of alert info (optional)
	group := "group_example" // string | group of alert info (optional)
	excludeCluster := true // bool | filter to exclude cluster of alert info groups, deprecated, use `is_global` instead. (optional)
	isGlobal := true // bool | filter global alert info groups(exclude alert info groups in any cluster) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertInfoGroupsAPI.ListAlertInfoGroups(context.Background()).QMust(qMust).Q(q).RelatedResource(relatedResource).Sort(sort).Limit(limit).Offset(offset).Level(level).ResourceType(resourceType).ResourceId(resourceId).CreateAfter(createAfter).CreateBefore(createBefore).Acked(acked).Resolved(resolved).Group(group).ExcludeCluster(excludeCluster).IsGlobal(isGlobal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfoGroupsAPI.ListAlertInfoGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertInfoGroups`: AlertInfoGroupsResp
	fmt.Fprintf(os.Stdout, "Response from `AlertInfoGroupsAPI.ListAlertInfoGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlertInfoGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qMust** | **string** | must query param of search | 
 **q** | **string** | should query param of search | 
 **relatedResource** | **string** | should query param of search | 
 **sort** | **string** | sort param of search | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **level** | **string** | level of alert info group | 
 **resourceType** | **string** | resource type of alert info group | 
 **resourceId** | **int64** | resource id of alert info group | 
 **createAfter** | **string** | create_after start time of alert info group | 
 **createBefore** | **string** | create_before end time of alert info group | 
 **acked** | **bool** | acked of alert info | 
 **resolved** | **bool** | resolved or not of alert info | 
 **group** | **string** | group of alert info | 
 **excludeCluster** | **bool** | filter to exclude cluster of alert info groups, deprecated, use &#x60;is_global&#x60; instead. | 
 **isGlobal** | **bool** | filter global alert info groups(exclude alert info groups in any cluster) | 

### Return type

[**AlertInfoGroupsResp**](AlertInfoGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveAlertInfoGroup

> AlertInfoGroupResp ResolveAlertInfoGroup(ctx, alertInfoGroupId).Execute()





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
	alertInfoGroupId := int64(789) // int64 | the id of alert info group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertInfoGroupsAPI.ResolveAlertInfoGroup(context.Background(), alertInfoGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfoGroupsAPI.ResolveAlertInfoGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveAlertInfoGroup`: AlertInfoGroupResp
	fmt.Fprintf(os.Stdout, "Response from `AlertInfoGroupsAPI.ResolveAlertInfoGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertInfoGroupId** | **int64** | the id of alert info group | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveAlertInfoGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertInfoGroupResp**](AlertInfoGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

