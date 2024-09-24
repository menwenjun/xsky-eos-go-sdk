# \OsOriginPullRulesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOSOriginPullRule**](OsOriginPullRulesAPI.md#CreateOSOriginPullRule) | **Post** /os-origin-pull-rules/ | 
[**DeleteOSOriginPullRule**](OsOriginPullRulesAPI.md#DeleteOSOriginPullRule) | **Delete** /os-origin-pull-rules/{rule_id} | 
[**GetOSOriginPullRule**](OsOriginPullRulesAPI.md#GetOSOriginPullRule) | **Get** /os-origin-pull-rules/{rule_id} | 
[**ListOSOriginPullRules**](OsOriginPullRulesAPI.md#ListOSOriginPullRules) | **Get** /os-origin-pull-rules/ | 
[**UpdateOSOriginPullRule**](OsOriginPullRulesAPI.md#UpdateOSOriginPullRule) | **Patch** /os-origin-pull-rules/{rule_id} | 



## CreateOSOriginPullRule

> OSOriginPullRuleResp CreateOSOriginPullRule(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewOSOriginPullRuleCreateReq() // OSOriginPullRuleCreateReq | origin pull rule info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsOriginPullRulesAPI.CreateOSOriginPullRule(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsOriginPullRulesAPI.CreateOSOriginPullRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOSOriginPullRule`: OSOriginPullRuleResp
	fmt.Fprintf(os.Stdout, "Response from `OsOriginPullRulesAPI.CreateOSOriginPullRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOSOriginPullRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OSOriginPullRuleCreateReq**](OSOriginPullRuleCreateReq.md) | origin pull rule info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSOriginPullRuleResp**](OSOriginPullRuleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOSOriginPullRule

> OSOriginPullRuleResp DeleteOSOriginPullRule(ctx, ruleId).Execute()





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
	ruleId := int64(789) // int64 | origin pull rule id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsOriginPullRulesAPI.DeleteOSOriginPullRule(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsOriginPullRulesAPI.DeleteOSOriginPullRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOSOriginPullRule`: OSOriginPullRuleResp
	fmt.Fprintf(os.Stdout, "Response from `OsOriginPullRulesAPI.DeleteOSOriginPullRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int64** | origin pull rule id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOSOriginPullRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSOriginPullRuleResp**](OSOriginPullRuleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSOriginPullRule

> OSOriginPullRuleResp GetOSOriginPullRule(ctx, ruleId).Execute()





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
	ruleId := int64(789) // int64 | rule id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsOriginPullRulesAPI.GetOSOriginPullRule(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsOriginPullRulesAPI.GetOSOriginPullRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSOriginPullRule`: OSOriginPullRuleResp
	fmt.Fprintf(os.Stdout, "Response from `OsOriginPullRulesAPI.GetOSOriginPullRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int64** | rule id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSOriginPullRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSOriginPullRuleResp**](OSOriginPullRuleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOSOriginPullRules

> OSOriginPullRulesResp ListOSOriginPullRules(ctx).ClusterId(clusterId).Limit(limit).Offset(offset).BucketId(bucketId).S3LbGroupId(s3LbGroupId).Execute()





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
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	bucketId := int64(789) // int64 | bucket id (optional)
	s3LbGroupId := int64(789) // int64 | s3 lb group id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsOriginPullRulesAPI.ListOSOriginPullRules(context.Background()).ClusterId(clusterId).Limit(limit).Offset(offset).BucketId(bucketId).S3LbGroupId(s3LbGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsOriginPullRulesAPI.ListOSOriginPullRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSOriginPullRules`: OSOriginPullRulesResp
	fmt.Fprintf(os.Stdout, "Response from `OsOriginPullRulesAPI.ListOSOriginPullRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSOriginPullRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **bucketId** | **int64** | bucket id | 
 **s3LbGroupId** | **int64** | s3 lb group id | 

### Return type

[**OSOriginPullRulesResp**](OSOriginPullRulesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOSOriginPullRule

> OSOriginPullRuleResp UpdateOSOriginPullRule(ctx, ruleId).Body(body).Execute()





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
	ruleId := int64(789) // int64 | origin pull rule id
	body := *openapiclient.NewOSOriginPullRuleUpdateReq() // OSOriginPullRuleUpdateReq | origin pull rule info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsOriginPullRulesAPI.UpdateOSOriginPullRule(context.Background(), ruleId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsOriginPullRulesAPI.UpdateOSOriginPullRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOSOriginPullRule`: OSOriginPullRuleResp
	fmt.Fprintf(os.Stdout, "Response from `OsOriginPullRulesAPI.UpdateOSOriginPullRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int64** | origin pull rule id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOSOriginPullRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSOriginPullRuleUpdateReq**](OSOriginPullRuleUpdateReq.md) | origin pull rule info | 

### Return type

[**OSOriginPullRuleResp**](OSOriginPullRuleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

