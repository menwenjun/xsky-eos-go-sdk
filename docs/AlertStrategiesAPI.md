# \AlertStrategiesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertStrategy**](AlertStrategiesAPI.md#CreateAlertStrategy) | **Post** /alert-strategies/ | 
[**DeleteAlertStrategy**](AlertStrategiesAPI.md#DeleteAlertStrategy) | **Delete** /alert-strategies/{alert_strategy_id} | 
[**GetAlertStrategy**](AlertStrategiesAPI.md#GetAlertStrategy) | **Get** /alert-strategies/{alert_strategy_id} | 
[**ListAlertStrategies**](AlertStrategiesAPI.md#ListAlertStrategies) | **Get** /alert-strategies/ | 
[**UpdateUpdateAlertStrategyAlertContact**](AlertStrategiesAPI.md#UpdateUpdateAlertStrategyAlertContact) | **Patch** /alert-strategies/{alert_strategy_id} | 



## CreateAlertStrategy

> AlertStrategyReq CreateAlertStrategy(ctx).Body(body).Execute()





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
	body := *openapiclient.NewAlertStrategyReq(*openapiclient.NewAlertStrategyReqAlertStrategy()) // AlertStrategyReq | alert strategy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertStrategiesAPI.CreateAlertStrategy(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertStrategiesAPI.CreateAlertStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlertStrategy`: AlertStrategyReq
	fmt.Fprintf(os.Stdout, "Response from `AlertStrategiesAPI.CreateAlertStrategy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AlertStrategyReq**](AlertStrategyReq.md) | alert strategy | 

### Return type

[**AlertStrategyReq**](AlertStrategyReq.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertStrategy

> DeleteAlertStrategy(ctx, alertStrategyId).Execute()





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
	alertStrategyId := int64(789) // int64 | alert strategy id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertStrategiesAPI.DeleteAlertStrategy(context.Background(), alertStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertStrategiesAPI.DeleteAlertStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertStrategyId** | **int64** | alert strategy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertStrategyRequest struct via the builder pattern


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


## GetAlertStrategy

> AlertStrategyResp GetAlertStrategy(ctx, alertStrategyId).Execute()





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
	alertStrategyId := int64(789) // int64 | alert strategy id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertStrategiesAPI.GetAlertStrategy(context.Background(), alertStrategyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertStrategiesAPI.GetAlertStrategy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertStrategy`: AlertStrategyResp
	fmt.Fprintf(os.Stdout, "Response from `AlertStrategiesAPI.GetAlertStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertStrategyId** | **int64** | alert strategy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertStrategyResp**](AlertStrategyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertStrategies

> AlertContactsResp ListAlertStrategies(ctx).Limit(limit).Offset(offset).Name(name).DurationBegin(durationBegin).DurationEnd(durationEnd).DurationLimit(durationLimit).DurationOffset(durationOffset).Q(q).Sort(sort).Execute()





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
	name := "name_example" // string | name of alert contact (optional)
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	durationLimit := int64(789) // int64 | duration limit param (optional)
	durationOffset := int64(789) // int64 | duration offset param (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertStrategiesAPI.ListAlertStrategies(context.Background()).Limit(limit).Offset(offset).Name(name).DurationBegin(durationBegin).DurationEnd(durationEnd).DurationLimit(durationLimit).DurationOffset(durationOffset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertStrategiesAPI.ListAlertStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertStrategies`: AlertContactsResp
	fmt.Fprintf(os.Stdout, "Response from `AlertStrategiesAPI.ListAlertStrategies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlertStrategiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **name** | **string** | name of alert contact | 
 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **durationLimit** | **int64** | duration limit param | 
 **durationOffset** | **int64** | duration offset param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**AlertContactsResp**](AlertContactsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUpdateAlertStrategyAlertContact

> UpdateUpdateAlertStrategyAlertContact(ctx, alertStrategyId).Alert(alert).Execute()





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
	alertStrategyId := int64(789) // int64 | alert strategy id
	alert := *openapiclient.NewAlertStrategyReq(*openapiclient.NewAlertStrategyReqAlertStrategy()) // AlertStrategyReq | strategy update request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertStrategiesAPI.UpdateUpdateAlertStrategyAlertContact(context.Background(), alertStrategyId).Alert(alert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertStrategiesAPI.UpdateUpdateAlertStrategyAlertContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertStrategyId** | **int64** | alert strategy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUpdateAlertStrategyAlertContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alert** | [**AlertStrategyReq**](AlertStrategyReq.md) | strategy update request | 

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

