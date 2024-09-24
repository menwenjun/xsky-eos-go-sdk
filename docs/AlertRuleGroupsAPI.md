# \AlertRuleGroupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlertRuleGroupDetail**](AlertRuleGroupsAPI.md#GetAlertRuleGroupDetail) | **Get** /alert-rule-groups/{id}/detail | 
[**ListAlertRuleGroups**](AlertRuleGroupsAPI.md#ListAlertRuleGroups) | **Get** /alert-rule-groups/ | 
[**UpdateAlertRuleGroupDetail**](AlertRuleGroupsAPI.md#UpdateAlertRuleGroupDetail) | **Patch** /alert-rule-groups/{id}/detail | 



## GetAlertRuleGroupDetail

> AlertRuleGroupDetailResp GetAlertRuleGroupDetail(ctx, id).Execute()





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
	id := int64(789) // int64 | alert rule group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRuleGroupsAPI.GetAlertRuleGroupDetail(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRuleGroupsAPI.GetAlertRuleGroupDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertRuleGroupDetail`: AlertRuleGroupDetailResp
	fmt.Fprintf(os.Stdout, "Response from `AlertRuleGroupsAPI.GetAlertRuleGroupDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | alert rule group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRuleGroupDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertRuleGroupDetailResp**](AlertRuleGroupDetailResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertRuleGroups

> AlertRuleGroupsResp ListAlertRuleGroups(ctx).Limit(limit).Offset(offset).ResourceType(resourceType).Name(name).Level(level).Enabled(enabled).Execute()





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
	resourceType := "resourceType_example" // string | resource type of alert rule group (optional)
	name := "name_example" // string | name of alert rule group (optional)
	level := "level_example" // string | level of alert rule group (optional)
	enabled := true // bool | enabled or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRuleGroupsAPI.ListAlertRuleGroups(context.Background()).Limit(limit).Offset(offset).ResourceType(resourceType).Name(name).Level(level).Enabled(enabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRuleGroupsAPI.ListAlertRuleGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertRuleGroups`: AlertRuleGroupsResp
	fmt.Fprintf(os.Stdout, "Response from `AlertRuleGroupsAPI.ListAlertRuleGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlertRuleGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **resourceType** | **string** | resource type of alert rule group | 
 **name** | **string** | name of alert rule group | 
 **level** | **string** | level of alert rule group | 
 **enabled** | **bool** | enabled or not | 

### Return type

[**AlertRuleGroupsResp**](AlertRuleGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertRuleGroupDetail

> AlertRuleGroupDetailResp UpdateAlertRuleGroupDetail(ctx, id).Body(body).Restore(restore).Execute()





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
	id := int64(789) // int64 | alert rule group id
	body := *openapiclient.NewAlertRuleGroupDetailUpdateReq() // AlertRuleGroupDetailUpdateReq | alert rule group detail
	restore := true // bool | restore default setup (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRuleGroupsAPI.UpdateAlertRuleGroupDetail(context.Background(), id).Body(body).Restore(restore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRuleGroupsAPI.UpdateAlertRuleGroupDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlertRuleGroupDetail`: AlertRuleGroupDetailResp
	fmt.Fprintf(os.Stdout, "Response from `AlertRuleGroupsAPI.UpdateAlertRuleGroupDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | alert rule group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertRuleGroupDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AlertRuleGroupDetailUpdateReq**](AlertRuleGroupDetailUpdateReq.md) | alert rule group detail | 
 **restore** | **bool** | restore default setup | 

### Return type

[**AlertRuleGroupDetailResp**](AlertRuleGroupDetailResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

