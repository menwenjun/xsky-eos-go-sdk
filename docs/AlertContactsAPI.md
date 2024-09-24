# \AlertContactsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAlertContact**](AlertContactsAPI.md#CheckAlertContact) | **Get** /alert-contacts/{alert_contact_id}/check | 
[**CreateAlertContact**](AlertContactsAPI.md#CreateAlertContact) | **Post** /alert-contacts/ | 
[**DeleteAlertContact**](AlertContactsAPI.md#DeleteAlertContact) | **Delete** /alert-contacts/{alert_contact_id} | 
[**GetAlertContact**](AlertContactsAPI.md#GetAlertContact) | **Get** /alert-contacts/{alert_contact_id} | 
[**ListAlertContacts**](AlertContactsAPI.md#ListAlertContacts) | **Get** /alert-contacts/ | 
[**UpdateAlertContact**](AlertContactsAPI.md#UpdateAlertContact) | **Patch** /alert-contacts/{alert_contact_id} | 



## CheckAlertContact

> AlertContactStrategyResp CheckAlertContact(ctx, alertContactId).Execute()





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
	alertContactId := int64(789) // int64 | alert contact id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertContactsAPI.CheckAlertContact(context.Background(), alertContactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertContactsAPI.CheckAlertContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckAlertContact`: AlertContactStrategyResp
	fmt.Fprintf(os.Stdout, "Response from `AlertContactsAPI.CheckAlertContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertContactId** | **int64** | alert contact id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckAlertContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertContactStrategyResp**](AlertContactStrategyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAlertContact

> AlertContactReq CreateAlertContact(ctx).Body(body).Execute()





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
	body := *openapiclient.NewAlertContactReq(*openapiclient.NewAlertContactReqAlertContact()) // AlertContactReq | alert contact

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertContactsAPI.CreateAlertContact(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertContactsAPI.CreateAlertContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlertContact`: AlertContactReq
	fmt.Fprintf(os.Stdout, "Response from `AlertContactsAPI.CreateAlertContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AlertContactReq**](AlertContactReq.md) | alert contact | 

### Return type

[**AlertContactReq**](AlertContactReq.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertContact

> DeleteAlertContact(ctx, alertContactId).Execute()





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
	alertContactId := int64(789) // int64 | alert contact id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertContactsAPI.DeleteAlertContact(context.Background(), alertContactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertContactsAPI.DeleteAlertContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertContactId** | **int64** | alert contact id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertContactRequest struct via the builder pattern


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


## GetAlertContact

> AlertContactResp GetAlertContact(ctx, alertContactId).Execute()





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
	alertContactId := int64(789) // int64 | alert contact id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertContactsAPI.GetAlertContact(context.Background(), alertContactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertContactsAPI.GetAlertContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertContact`: AlertContactResp
	fmt.Fprintf(os.Stdout, "Response from `AlertContactsAPI.GetAlertContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertContactId** | **int64** | alert contact id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertContactResp**](AlertContactResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertContacts

> AlertContactsResp ListAlertContacts(ctx).Limit(limit).Offset(offset).Name(name).DurationBegin(durationBegin).DurationEnd(durationEnd).DurationLimit(durationLimit).DurationOffset(durationOffset).Q(q).Sort(sort).Execute()





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
	name := "name_example" // string | name of alert contact (optional)
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	durationLimit := int64(789) // int64 | duration limit param (optional)
	durationOffset := int64(789) // int64 | duration offset param (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertContactsAPI.ListAlertContacts(context.Background()).Limit(limit).Offset(offset).Name(name).DurationBegin(durationBegin).DurationEnd(durationEnd).DurationLimit(durationLimit).DurationOffset(durationOffset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertContactsAPI.ListAlertContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertContacts`: AlertContactsResp
	fmt.Fprintf(os.Stdout, "Response from `AlertContactsAPI.ListAlertContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlertContactsRequest struct via the builder pattern


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


## UpdateAlertContact

> UpdateAlertContact(ctx, alertContactId).Alert(alert).Execute()





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
	alertContactId := int64(789) // int64 | alert contact id
	alert := *openapiclient.NewAlertsActionReq(*openapiclient.NewAlertsActionReqAlerts("Action_example")) // AlertsActionReq | the alerts action info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertContactsAPI.UpdateAlertContact(context.Background(), alertContactId).Alert(alert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertContactsAPI.UpdateAlertContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertContactId** | **int64** | alert contact id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertContactRequest struct via the builder pattern


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

