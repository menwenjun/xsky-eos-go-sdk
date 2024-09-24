# \AlertInfosAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AckAlertInfo**](AlertInfosAPI.md#AckAlertInfo) | **Post** /alert-infos/{alert_info_id}:ack | 
[**CountAlertInfos**](AlertInfosAPI.md#CountAlertInfos) | **Get** /alert-infos/stats | 
[**CountBatchDeleteAlertInfos**](AlertInfosAPI.md#CountBatchDeleteAlertInfos) | **Get** /alert-infos/batch-delete | 
[**DeleteAlertInfo**](AlertInfosAPI.md#DeleteAlertInfo) | **Delete** /alert-infos/{alert_info_id} | 
[**DeleteAlertInfos**](AlertInfosAPI.md#DeleteAlertInfos) | **Delete** /alert-infos/ | 
[**GetAlertInfo**](AlertInfosAPI.md#GetAlertInfo) | **Get** /alert-infos/{alert_info_id} | 
[**GetAlertInfosReport**](AlertInfosAPI.md#GetAlertInfosReport) | **Get** /alert-infos/report | 
[**ListAlertInfos**](AlertInfosAPI.md#ListAlertInfos) | **Get** /alert-infos/ | 
[**ResolveAlertInfo**](AlertInfosAPI.md#ResolveAlertInfo) | **Post** /alert-infos/{alert_info_id}:resolve | 



## AckAlertInfo

> AckAlertInfo(ctx, alertInfoId).Execute()





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
	alertInfoId := int64(789) // int64 | the id of alert info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertInfosAPI.AckAlertInfo(context.Background(), alertInfoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfosAPI.AckAlertInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertInfoId** | **int64** | the id of alert info | 

### Other Parameters

Other parameters are passed through a pointer to a apiAckAlertInfoRequest struct via the builder pattern


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


## CountAlertInfos

> AlertStatsResp CountAlertInfos(ctx).Acked(acked).Resolved(resolved).ResourceType(resourceType).ResourceId(resourceId).Execute()





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
	acked := true // bool | acked of alert info (optional)
	resolved := true // bool | resolved or not of alert info (optional)
	resourceType := "resourceType_example" // string | resource type of alert info (optional)
	resourceId := int64(789) // int64 | resource id of alert info (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertInfosAPI.CountAlertInfos(context.Background()).Acked(acked).Resolved(resolved).ResourceType(resourceType).ResourceId(resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfosAPI.CountAlertInfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountAlertInfos`: AlertStatsResp
	fmt.Fprintf(os.Stdout, "Response from `AlertInfosAPI.CountAlertInfos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountAlertInfosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acked** | **bool** | acked of alert info | 
 **resolved** | **bool** | resolved or not of alert info | 
 **resourceType** | **string** | resource type of alert info | 
 **resourceId** | **int64** | resource id of alert info | 

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


## CountBatchDeleteAlertInfos

> AlertInfoBatchDeleteResp CountBatchDeleteAlertInfos(ctx).CreateAfter(createAfter).CreateBefore(createBefore).ClusterId(clusterId).Execute()





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
	createAfter := "createAfter_example" // string | create_after timestamp of alert info (optional)
	createBefore := "createBefore_example" // string | create_before end time of alert info (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertInfosAPI.CountBatchDeleteAlertInfos(context.Background()).CreateAfter(createAfter).CreateBefore(createBefore).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfosAPI.CountBatchDeleteAlertInfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountBatchDeleteAlertInfos`: AlertInfoBatchDeleteResp
	fmt.Fprintf(os.Stdout, "Response from `AlertInfosAPI.CountBatchDeleteAlertInfos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountBatchDeleteAlertInfosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAfter** | **string** | create_after timestamp of alert info | 
 **createBefore** | **string** | create_before end time of alert info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**AlertInfoBatchDeleteResp**](AlertInfoBatchDeleteResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertInfo

> AlertInfoResp DeleteAlertInfo(ctx, alertInfoId).Execute()





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
	alertInfoId := int64(789) // int64 | the id of alert info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertInfosAPI.DeleteAlertInfo(context.Background(), alertInfoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfosAPI.DeleteAlertInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlertInfo`: AlertInfoResp
	fmt.Fprintf(os.Stdout, "Response from `AlertInfosAPI.DeleteAlertInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertInfoId** | **int64** | the id of alert info | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertInfoResp**](AlertInfoResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertInfos

> AlertInfoBatchDeleteResp DeleteAlertInfos(ctx).CreateAfter(createAfter).CreateBefore(createBefore).ClusterId(clusterId).Execute()





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
	createAfter := "createAfter_example" // string | create_after start time of alert info (optional)
	createBefore := "createBefore_example" // string | create_before end time of alert info (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertInfosAPI.DeleteAlertInfos(context.Background()).CreateAfter(createAfter).CreateBefore(createBefore).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfosAPI.DeleteAlertInfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlertInfos`: AlertInfoBatchDeleteResp
	fmt.Fprintf(os.Stdout, "Response from `AlertInfosAPI.DeleteAlertInfos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertInfosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAfter** | **string** | create_after start time of alert info | 
 **createBefore** | **string** | create_before end time of alert info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**AlertInfoBatchDeleteResp**](AlertInfoBatchDeleteResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertInfo

> AlertInfoResp GetAlertInfo(ctx, alertInfoId).Execute()





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
	alertInfoId := int64(789) // int64 | alert info id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertInfosAPI.GetAlertInfo(context.Background(), alertInfoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfosAPI.GetAlertInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertInfo`: AlertInfoResp
	fmt.Fprintf(os.Stdout, "Response from `AlertInfosAPI.GetAlertInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertInfoId** | **int64** | alert info id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertInfoResp**](AlertInfoResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertInfosReport

> GetAlertInfosReport(ctx).Level(level).ResourceType(resourceType).ResourceId(resourceId).CreateAfter(createAfter).Acked(acked).Resolved(resolved).Group(group).ExcludeCluster(excludeCluster).IsGlobal(isGlobal).Execute()





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
	level := "level_example" // string | level of alert info (optional)
	resourceType := "resourceType_example" // string | resource type of alert info (optional)
	resourceId := int64(789) // int64 | resource id of alert info (optional)
	createAfter := "createAfter_example" // string | create_after timestamp of alert info (optional)
	acked := true // bool | acked of alert info (optional)
	resolved := true // bool | resolved or not of alert info (optional)
	group := "group_example" // string | group of alert info (optional)
	excludeCluster := true // bool | filter to exclude cluster of alert info, deprecated, use `is_global` instead. (optional)
	isGlobal := true // bool | filter global alert info groups(exclude alert info in any cluster) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertInfosAPI.GetAlertInfosReport(context.Background()).Level(level).ResourceType(resourceType).ResourceId(resourceId).CreateAfter(createAfter).Acked(acked).Resolved(resolved).Group(group).ExcludeCluster(excludeCluster).IsGlobal(isGlobal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfosAPI.GetAlertInfosReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertInfosReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **level** | **string** | level of alert info | 
 **resourceType** | **string** | resource type of alert info | 
 **resourceId** | **int64** | resource id of alert info | 
 **createAfter** | **string** | create_after timestamp of alert info | 
 **acked** | **bool** | acked of alert info | 
 **resolved** | **bool** | resolved or not of alert info | 
 **group** | **string** | group of alert info | 
 **excludeCluster** | **bool** | filter to exclude cluster of alert info, deprecated, use &#x60;is_global&#x60; instead. | 
 **isGlobal** | **bool** | filter global alert info groups(exclude alert info in any cluster) | 

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


## ListAlertInfos

> AlertInfosResp ListAlertInfos(ctx).QMust(qMust).Q(q).RelatedResource(relatedResource).Sort(sort).Limit(limit).Offset(offset).Level(level).ResourceType(resourceType).ResourceId(resourceId).CreateAfter(createAfter).Acked(acked).Resolved(resolved).Group(group).ExcludeCluster(excludeCluster).IsGlobal(isGlobal).Execute()





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
	level := "level_example" // string | level of alert info (optional)
	resourceType := "resourceType_example" // string | resource type of alert info (optional)
	resourceId := int64(789) // int64 | resource id of alert info (optional)
	createAfter := "createAfter_example" // string | create_after timestamp of alert info (optional)
	acked := true // bool | acked of alert info (optional)
	resolved := true // bool | resolved or not of alert info (optional)
	group := "group_example" // string | group of alert info (optional)
	excludeCluster := true // bool | filter to exclude cluster of alert info, deprecated, use `is_global` instead (optional)
	isGlobal := true // bool | filter global alert info(exclude alert info in any cluster) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertInfosAPI.ListAlertInfos(context.Background()).QMust(qMust).Q(q).RelatedResource(relatedResource).Sort(sort).Limit(limit).Offset(offset).Level(level).ResourceType(resourceType).ResourceId(resourceId).CreateAfter(createAfter).Acked(acked).Resolved(resolved).Group(group).ExcludeCluster(excludeCluster).IsGlobal(isGlobal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfosAPI.ListAlertInfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertInfos`: AlertInfosResp
	fmt.Fprintf(os.Stdout, "Response from `AlertInfosAPI.ListAlertInfos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlertInfosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qMust** | **string** | must query param of search | 
 **q** | **string** | should query param of search | 
 **relatedResource** | **string** | should query param of search | 
 **sort** | **string** | sort param of search | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **level** | **string** | level of alert info | 
 **resourceType** | **string** | resource type of alert info | 
 **resourceId** | **int64** | resource id of alert info | 
 **createAfter** | **string** | create_after timestamp of alert info | 
 **acked** | **bool** | acked of alert info | 
 **resolved** | **bool** | resolved or not of alert info | 
 **group** | **string** | group of alert info | 
 **excludeCluster** | **bool** | filter to exclude cluster of alert info, deprecated, use &#x60;is_global&#x60; instead | 
 **isGlobal** | **bool** | filter global alert info(exclude alert info in any cluster) | 

### Return type

[**AlertInfosResp**](AlertInfosResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveAlertInfo

> ResolveAlertInfo(ctx, alertInfoId).Execute()





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
	alertInfoId := int64(789) // int64 | the id of alert info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertInfosAPI.ResolveAlertInfo(context.Background(), alertInfoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfosAPI.ResolveAlertInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertInfoId** | **int64** | the id of alert info | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveAlertInfoRequest struct via the builder pattern


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

