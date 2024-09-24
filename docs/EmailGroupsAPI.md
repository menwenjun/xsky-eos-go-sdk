# \EmailGroupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailGroup**](EmailGroupsAPI.md#CreateEmailGroup) | **Post** /email-groups/ | 
[**DeleteEmailGroup**](EmailGroupsAPI.md#DeleteEmailGroup) | **Delete** /email-groups/{group_id} | 
[**GetEmailGroup**](EmailGroupsAPI.md#GetEmailGroup) | **Get** /email-groups/{group_id} | 
[**ListEmailGroups**](EmailGroupsAPI.md#ListEmailGroups) | **Get** /email-groups/ | 
[**UpdateEmailGroup**](EmailGroupsAPI.md#UpdateEmailGroup) | **Put** /email-groups/{group_id} | 



## CreateEmailGroup

> EmailGroupResp CreateEmailGroup(ctx).Body(body).Execute()





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
	body := *openapiclient.NewEmailGroupUpdateReq(*openapiclient.NewEmailGroupUpdateReqEmailGroup([]string{"Emails_example"}, "Name_example")) // EmailGroupUpdateReq | email group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailGroupsAPI.CreateEmailGroup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailGroupsAPI.CreateEmailGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEmailGroup`: EmailGroupResp
	fmt.Fprintf(os.Stdout, "Response from `EmailGroupsAPI.CreateEmailGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EmailGroupUpdateReq**](EmailGroupUpdateReq.md) | email group | 

### Return type

[**EmailGroupResp**](EmailGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmailGroup

> DeleteEmailGroup(ctx, groupId).Execute()





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
	groupId := int64(789) // int64 | the id of email group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailGroupsAPI.DeleteEmailGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailGroupsAPI.DeleteEmailGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | the id of email group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailGroupRequest struct via the builder pattern


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


## GetEmailGroup

> EmailGroupResp GetEmailGroup(ctx, groupId).Execute()





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
	groupId := int64(789) // int64 | the id of email group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailGroupsAPI.GetEmailGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailGroupsAPI.GetEmailGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailGroup`: EmailGroupResp
	fmt.Fprintf(os.Stdout, "Response from `EmailGroupsAPI.GetEmailGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | the id of email group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailGroupResp**](EmailGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailGroups

> EmailGroupsResp ListEmailGroups(ctx).Limit(limit).Offset(offset).AlertGroupId(alertGroupId).Name(name).Execute()





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
	name := "name_example" // string | name of email groups (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailGroupsAPI.ListEmailGroups(context.Background()).Limit(limit).Offset(offset).AlertGroupId(alertGroupId).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailGroupsAPI.ListEmailGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEmailGroups`: EmailGroupsResp
	fmt.Fprintf(os.Stdout, "Response from `EmailGroupsAPI.ListEmailGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEmailGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **alertGroupId** | **int64** | alert group id | 
 **name** | **string** | name of email groups | 

### Return type

[**EmailGroupsResp**](EmailGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailGroup

> EmailGroupResp UpdateEmailGroup(ctx, groupId).Body(body).Execute()





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
	groupId := int64(789) // int64 | the id of email group
	body := *openapiclient.NewEmailGroupUpdateReq(*openapiclient.NewEmailGroupUpdateReqEmailGroup([]string{"Emails_example"}, "Name_example")) // EmailGroupUpdateReq | email group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailGroupsAPI.UpdateEmailGroup(context.Background(), groupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailGroupsAPI.UpdateEmailGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEmailGroup`: EmailGroupResp
	fmt.Fprintf(os.Stdout, "Response from `EmailGroupsAPI.UpdateEmailGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | the id of email group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**EmailGroupUpdateReq**](EmailGroupUpdateReq.md) | email group | 

### Return type

[**EmailGroupResp**](EmailGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

