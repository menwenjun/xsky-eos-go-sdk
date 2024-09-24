# \AlertGroupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertGroup**](AlertGroupsAPI.md#CreateAlertGroup) | **Post** /alert-groups/ | 
[**DeleteAlertGroup**](AlertGroupsAPI.md#DeleteAlertGroup) | **Delete** /alert-groups/{group_id} | 
[**GetAlertGroup**](AlertGroupsAPI.md#GetAlertGroup) | **Get** /alert-groups/{group_id} | 
[**ListAlertGroups**](AlertGroupsAPI.md#ListAlertGroups) | **Get** /alert-groups/ | 
[**UpdateAlertGroup**](AlertGroupsAPI.md#UpdateAlertGroup) | **Patch** /alert-groups/{group_id} | 



## CreateAlertGroup

> AlertGroupResp CreateAlertGroup(ctx).Body(body).Execute()





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
	body := *openapiclient.NewAlertGroupCreateReq(*openapiclient.NewAlertGroupCreateReqAlertGroup([]int64{int64(123)}, []int64{int64(123)}, "Name_example")) // AlertGroupCreateReq | alert group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertGroupsAPI.CreateAlertGroup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsAPI.CreateAlertGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlertGroup`: AlertGroupResp
	fmt.Fprintf(os.Stdout, "Response from `AlertGroupsAPI.CreateAlertGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AlertGroupCreateReq**](AlertGroupCreateReq.md) | alert group | 

### Return type

[**AlertGroupResp**](AlertGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertGroup

> DeleteAlertGroup(ctx, groupId).Execute()





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
	groupId := int64(789) // int64 | the id of alert group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertGroupsAPI.DeleteAlertGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsAPI.DeleteAlertGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | the id of alert group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertGroupRequest struct via the builder pattern


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


## GetAlertGroup

> AlertGroupResp GetAlertGroup(ctx, groupId).Execute()





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
	groupId := int64(789) // int64 | the id of alert group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertGroupsAPI.GetAlertGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsAPI.GetAlertGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertGroup`: AlertGroupResp
	fmt.Fprintf(os.Stdout, "Response from `AlertGroupsAPI.GetAlertGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | the id of alert group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertGroupResp**](AlertGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertGroups

> AlertGroupsResp ListAlertGroups(ctx).Limit(limit).Offset(offset).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertGroupsAPI.ListAlertGroups(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsAPI.ListAlertGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertGroups`: AlertGroupsResp
	fmt.Fprintf(os.Stdout, "Response from `AlertGroupsAPI.ListAlertGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlertGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**AlertGroupsResp**](AlertGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertGroup

> AlertGroupResp UpdateAlertGroup(ctx, groupId).Body(body).Execute()





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
	groupId := int64(789) // int64 | the id of alert group
	body := *openapiclient.NewAlertGroupUpdateReq(*openapiclient.NewAlertGroupUpdateReqAlertGroup()) // AlertGroupUpdateReq | alert group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertGroupsAPI.UpdateAlertGroup(context.Background(), groupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsAPI.UpdateAlertGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlertGroup`: AlertGroupResp
	fmt.Fprintf(os.Stdout, "Response from `AlertGroupsAPI.UpdateAlertGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | the id of alert group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AlertGroupUpdateReq**](AlertGroupUpdateReq.md) | alert group | 

### Return type

[**AlertGroupResp**](AlertGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

