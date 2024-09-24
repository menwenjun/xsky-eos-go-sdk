# \AlertRulesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertRuleResourceBlacklist**](AlertRulesAPI.md#CreateAlertRuleResourceBlacklist) | **Post** /alert-rules/{rule_id}/blacklist | 
[**DeleteAlertRule**](AlertRulesAPI.md#DeleteAlertRule) | **Delete** /alert-rules/{rule_id} | 
[**DeleteAlertRuleResourceBlacklist**](AlertRulesAPI.md#DeleteAlertRuleResourceBlacklist) | **Delete** /alert-rules/{rule_id}/blacklist | 
[**GetAlertRule**](AlertRulesAPI.md#GetAlertRule) | **Get** /alert-rules/{rule_id} | 
[**GetAlertRuleSchema**](AlertRulesAPI.md#GetAlertRuleSchema) | **Get** /alert-rules/schema | 
[**GetAlertRulesReport**](AlertRulesAPI.md#GetAlertRulesReport) | **Get** /alert-rules/report | 
[**ListAlertRuleResourceBlacklist**](AlertRulesAPI.md#ListAlertRuleResourceBlacklist) | **Get** /alert-rules/{rule_id}/blacklist | 
[**ListAlertRules**](AlertRulesAPI.md#ListAlertRules) | **Get** /alert-rules/ | 
[**UpdateAlertRule**](AlertRulesAPI.md#UpdateAlertRule) | **Patch** /alert-rules/{rule_id} | 



## CreateAlertRuleResourceBlacklist

> AlertRuleResourceBlacklistResp CreateAlertRuleResourceBlacklist(ctx, ruleId).Body(body).Execute()





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
	ruleId := int64(789) // int64 | the id of alert rule
	body := *openapiclient.NewUpdateAlertRuleResourceBlacklistReq(*openapiclient.NewUpdateAlertRuleResourceBlacklistReqAlertRuleResourceBlacklist()) // UpdateAlertRuleResourceBlacklistReq | resource blacklist

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRulesAPI.CreateAlertRuleResourceBlacklist(context.Background(), ruleId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.CreateAlertRuleResourceBlacklist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlertRuleResourceBlacklist`: AlertRuleResourceBlacklistResp
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.CreateAlertRuleResourceBlacklist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int64** | the id of alert rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertRuleResourceBlacklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateAlertRuleResourceBlacklistReq**](UpdateAlertRuleResourceBlacklistReq.md) | resource blacklist | 

### Return type

[**AlertRuleResourceBlacklistResp**](AlertRuleResourceBlacklistResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertRule

> DeleteAlertRule(ctx, ruleId).Execute()





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
	ruleId := int64(789) // int64 | the id of alert rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertRulesAPI.DeleteAlertRule(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.DeleteAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int64** | the id of alert rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertRuleRequest struct via the builder pattern


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


## DeleteAlertRuleResourceBlacklist

> DeleteAlertRuleResourceBlacklist(ctx, ruleId).Body(body).Execute()





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
	ruleId := int64(789) // int64 | the id of alert rule
	body := *openapiclient.NewUpdateAlertRuleResourceBlacklistReq(*openapiclient.NewUpdateAlertRuleResourceBlacklistReqAlertRuleResourceBlacklist()) // UpdateAlertRuleResourceBlacklistReq | resource blacklist

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertRulesAPI.DeleteAlertRuleResourceBlacklist(context.Background(), ruleId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.DeleteAlertRuleResourceBlacklist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int64** | the id of alert rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertRuleResourceBlacklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateAlertRuleResourceBlacklistReq**](UpdateAlertRuleResourceBlacklistReq.md) | resource blacklist | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertRule

> AlertRuleResp GetAlertRule(ctx, ruleId).Execute()





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
	ruleId := int64(789) // int64 | the id of alert rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRulesAPI.GetAlertRule(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.GetAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertRule`: AlertRuleResp
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.GetAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int64** | the id of alert rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertRuleResp**](AlertRuleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertRuleSchema

> AlertRuleSchemaResp GetAlertRuleSchema(ctx).Execute()





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
	resp, r, err := apiClient.AlertRulesAPI.GetAlertRuleSchema(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.GetAlertRuleSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertRuleSchema`: AlertRuleSchemaResp
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.GetAlertRuleSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRuleSchemaRequest struct via the builder pattern


### Return type

[**AlertRuleSchemaResp**](AlertRuleSchemaResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertRulesReport

> GetAlertRulesReport(ctx).ResourceType(resourceType).GroupName(groupName).Level(level).Enabled(enabled).BasicType(basicType).Execute()





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
	resourceType := "resourceType_example" // string | resource type of alert rule (optional)
	groupName := "groupName_example" // string | group name of alert rule (optional)
	level := "level_example" // string | level of alert rule (optional)
	enabled := true // bool | enabled or not (optional)
	basicType := "basicType_example" // string | basic type of alert rule (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertRulesAPI.GetAlertRulesReport(context.Background()).ResourceType(resourceType).GroupName(groupName).Level(level).Enabled(enabled).BasicType(basicType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.GetAlertRulesReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRulesReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceType** | **string** | resource type of alert rule | 
 **groupName** | **string** | group name of alert rule | 
 **level** | **string** | level of alert rule | 
 **enabled** | **bool** | enabled or not | 
 **basicType** | **string** | basic type of alert rule | 

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


## ListAlertRuleResourceBlacklist

> AlertRuleResourceBlacklistResp ListAlertRuleResourceBlacklist(ctx, ruleId).Blacklist(blacklist).Limit(limit).Offset(offset).Execute()





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
	ruleId := int64(789) // int64 | the id of alert rule
	blacklist := "blacklist_example" // string | filter resource in blacklist or not
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRulesAPI.ListAlertRuleResourceBlacklist(context.Background(), ruleId).Blacklist(blacklist).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.ListAlertRuleResourceBlacklist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertRuleResourceBlacklist`: AlertRuleResourceBlacklistResp
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.ListAlertRuleResourceBlacklist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int64** | the id of alert rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertRuleResourceBlacklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blacklist** | **string** | filter resource in blacklist or not | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**AlertRuleResourceBlacklistResp**](AlertRuleResourceBlacklistResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertRules

> AlertRulesResp ListAlertRules(ctx).Limit(limit).Offset(offset).AlertGroupId(alertGroupId).ResourceType(resourceType).GroupName(groupName).Level(level).Enabled(enabled).Q(q).Sort(sort).BasicType(basicType).Execute()





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
	alertGroupId := int64(789) // int64 | alert group id (optional)
	resourceType := "resourceType_example" // string | resource type of alert rule (optional)
	groupName := "groupName_example" // string | group name of alert rule (optional)
	level := "level_example" // string | level of alert rule (optional)
	enabled := true // bool | enabled or not (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	basicType := "basicType_example" // string | basic type of alert rule (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRulesAPI.ListAlertRules(context.Background()).Limit(limit).Offset(offset).AlertGroupId(alertGroupId).ResourceType(resourceType).GroupName(groupName).Level(level).Enabled(enabled).Q(q).Sort(sort).BasicType(basicType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.ListAlertRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertRules`: AlertRulesResp
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.ListAlertRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlertRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **alertGroupId** | **int64** | alert group id | 
 **resourceType** | **string** | resource type of alert rule | 
 **groupName** | **string** | group name of alert rule | 
 **level** | **string** | level of alert rule | 
 **enabled** | **bool** | enabled or not | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **basicType** | **string** | basic type of alert rule | 

### Return type

[**AlertRulesResp**](AlertRulesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertRule

> AlertRuleResp UpdateAlertRule(ctx, ruleId).Body(body).Execute()





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
	ruleId := int64(789) // int64 | the id of alert rule
	body := *openapiclient.NewAlertRuleUpdateReq(*openapiclient.NewAlertRuleUpdateReqRule()) // AlertRuleUpdateReq | alert rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRulesAPI.UpdateAlertRule(context.Background(), ruleId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.UpdateAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlertRule`: AlertRuleResp
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.UpdateAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int64** | the id of alert rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AlertRuleUpdateReq**](AlertRuleUpdateReq.md) | alert rule | 

### Return type

[**AlertRuleResp**](AlertRuleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

