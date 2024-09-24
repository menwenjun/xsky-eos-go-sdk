# \AlertsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountAlerts**](AlertsAPI.md#CountAlerts) | **Get** /alerts/stats | 
[**DeleteAlert**](AlertsAPI.md#DeleteAlert) | **Delete** /alerts/{alert_id} | 
[**DeleteAlerts**](AlertsAPI.md#DeleteAlerts) | **Delete** /alerts/ | 
[**DoAlert**](AlertsAPI.md#DoAlert) | **Put** /alerts/{alert_id} | 
[**DoAlerts**](AlertsAPI.md#DoAlerts) | **Patch** /alerts/ | 
[**GetAlert**](AlertsAPI.md#GetAlert) | **Get** /alerts/{alert_id} | 
[**ListAlerts**](AlertsAPI.md#ListAlerts) | **Get** /alerts/ | 
[**ResolveAlert**](AlertsAPI.md#ResolveAlert) | **Post** /alerts/{alert_id}:resolve | 
[**ResolveAlerts**](AlertsAPI.md#ResolveAlerts) | **Post** /alerts/:resolve | 



## CountAlerts

> AlertStatsResp CountAlerts(ctx).Acked(acked).Resolved(resolved).ResourceType(resourceType).ResourceId(resourceId).Execute()





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
	acked := true // bool | acked of alert (optional)
	resolved := true // bool | resolved or not of alert (optional)
	resourceType := "resourceType_example" // string | resource type of alert (optional)
	resourceId := int64(789) // int64 | resource id of alert (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CountAlerts(context.Background()).Acked(acked).Resolved(resolved).ResourceType(resourceType).ResourceId(resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CountAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountAlerts`: AlertStatsResp
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CountAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acked** | **bool** | acked of alert | 
 **resolved** | **bool** | resolved or not of alert | 
 **resourceType** | **string** | resource type of alert | 
 **resourceId** | **int64** | resource id of alert | 

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


## DeleteAlert

> DeleteAlert(ctx, alertId).Execute()





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
	alertId := int64(789) // int64 | alert id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertsAPI.DeleteAlert(context.Background(), alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **int64** | alert id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertRequest struct via the builder pattern


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


## DeleteAlerts

> DeleteAlerts(ctx).Before(before).Execute()





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
	before := "before_example" // string | create time of last alert (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertsAPI.DeleteAlerts(context.Background()).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **string** | create time of last alert | 

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


## DoAlert

> AlertResp DoAlert(ctx, alertId).Alert(alert).Execute()





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
	alertId := int64(789) // int64 | the id of alert
	alert := *openapiclient.NewAlertActionReq(*openapiclient.NewAlertActionReqAlert("Action_example")) // AlertActionReq | the alert action info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.DoAlert(context.Background(), alertId).Alert(alert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DoAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoAlert`: AlertResp
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.DoAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **int64** | the id of alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alert** | [**AlertActionReq**](AlertActionReq.md) | the alert action info | 

### Return type

[**AlertResp**](AlertResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DoAlerts

> DoAlerts(ctx).Alert(alert).Execute()





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
	alert := *openapiclient.NewAlertsActionReq(*openapiclient.NewAlertsActionReqAlerts("Action_example")) // AlertsActionReq | the alerts action info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertsAPI.DoAlerts(context.Background()).Alert(alert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DoAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDoAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alert** | [**AlertsActionReq**](AlertsActionReq.md) | the alerts action info | 

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


## GetAlert

> AlertResp GetAlert(ctx, alertId).Execute()





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
	alertId := int64(789) // int64 | alert id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlert(context.Background(), alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlert`: AlertResp
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **int64** | alert id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertResp**](AlertResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlerts

> AlertsResp ListAlerts(ctx).Limit(limit).Offset(offset).ResourceType(resourceType).ResourceId(resourceId).AlertType(alertType).Acked(acked).Resolved(resolved).ResolveType(resolveType).Level(level).DurationBegin(durationBegin).DurationEnd(durationEnd).DurationLimit(durationLimit).DurationOffset(durationOffset).Q(q).RelatedResource(relatedResource).Sort(sort).Execute()





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
	resourceType := "resourceType_example" // string | resource type of alert (optional)
	resourceId := int64(789) // int64 | resource id of alert (optional)
	alertType := "alertType_example" // string | type of alert (optional)
	acked := true // bool | acked of alert (optional)
	resolved := true // bool | resolved or not of alert (optional)
	resolveType := "resolveType_example" // string | resolve type of alert (optional)
	level := "level_example" // string | level of alert (optional)
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	durationLimit := int64(789) // int64 | duration limit param (optional)
	durationOffset := int64(789) // int64 | duration offset param (optional)
	q := "q_example" // string | query param of search (optional)
	relatedResource := "relatedResource_example" // string | related resource info of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ListAlerts(context.Background()).Limit(limit).Offset(offset).ResourceType(resourceType).ResourceId(resourceId).AlertType(alertType).Acked(acked).Resolved(resolved).ResolveType(resolveType).Level(level).DurationBegin(durationBegin).DurationEnd(durationEnd).DurationLimit(durationLimit).DurationOffset(durationOffset).Q(q).RelatedResource(relatedResource).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlerts`: AlertsResp
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **resourceType** | **string** | resource type of alert | 
 **resourceId** | **int64** | resource id of alert | 
 **alertType** | **string** | type of alert | 
 **acked** | **bool** | acked of alert | 
 **resolved** | **bool** | resolved or not of alert | 
 **resolveType** | **string** | resolve type of alert | 
 **level** | **string** | level of alert | 
 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **durationLimit** | **int64** | duration limit param | 
 **durationOffset** | **int64** | duration offset param | 
 **q** | **string** | query param of search | 
 **relatedResource** | **string** | related resource info of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**AlertsResp**](AlertsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveAlert

> AlertResp ResolveAlert(ctx, alertId).Execute()





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
	alertId := int64(789) // int64 | the id of alert

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ResolveAlert(context.Background(), alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ResolveAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveAlert`: AlertResp
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ResolveAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **int64** | the id of alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertResp**](AlertResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveAlerts

> ResolveAlerts(ctx).Alert(alert).Execute()





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
	alert := *openapiclient.NewAlertsResolveReq(*openapiclient.NewAlertsResolveReqAlerts()) // AlertsResolveReq | the alerts resolve info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertsAPI.ResolveAlerts(context.Background()).Alert(alert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ResolveAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResolveAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alert** | [**AlertsResolveReq**](AlertsResolveReq.md) | the alerts resolve info | 

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

