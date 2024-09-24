# \OsUsersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObjectStorageUser**](OsUsersAPI.md#CreateObjectStorageUser) | **Post** /os-users/ | 
[**DeleteObjectStorageUser**](OsUsersAPI.md#DeleteObjectStorageUser) | **Delete** /os-users/{user_id} | 
[**GetObjectStorageUser**](OsUsersAPI.md#GetObjectStorageUser) | **Get** /os-users/{user_id} | 
[**GetObjectStorageUserSamples**](OsUsersAPI.md#GetObjectStorageUserSamples) | **Get** /os-users/{user_id}/samples | 
[**ListObjectStorageUsers**](OsUsersAPI.md#ListObjectStorageUsers) | **Get** /os-users/ | 
[**UpdateObjectStorageUser**](OsUsersAPI.md#UpdateObjectStorageUser) | **Patch** /os-users/{user_id} | 



## CreateObjectStorageUser

> ObjectStorageUserResp CreateObjectStorageUser(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewObjectStorageUserCreateReq() // ObjectStorageUserCreateReq | user info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsUsersAPI.CreateObjectStorageUser(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsUsersAPI.CreateObjectStorageUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateObjectStorageUser`: ObjectStorageUserResp
	fmt.Fprintf(os.Stdout, "Response from `OsUsersAPI.CreateObjectStorageUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectStorageUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ObjectStorageUserCreateReq**](ObjectStorageUserCreateReq.md) | user info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**ObjectStorageUserResp**](ObjectStorageUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObjectStorageUser

> ObjectStorageUserResp DeleteObjectStorageUser(ctx, userId).Execute()





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
	userId := int64(789) // int64 | user id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsUsersAPI.DeleteObjectStorageUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsUsersAPI.DeleteObjectStorageUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteObjectStorageUser`: ObjectStorageUserResp
	fmt.Fprintf(os.Stdout, "Response from `OsUsersAPI.DeleteObjectStorageUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | user id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectStorageUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStorageUserResp**](ObjectStorageUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageUser

> ObjectStorageUserResp GetObjectStorageUser(ctx, userId).Execute()





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
	userId := int64(789) // int64 | user id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsUsersAPI.GetObjectStorageUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsUsersAPI.GetObjectStorageUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageUser`: ObjectStorageUserResp
	fmt.Fprintf(os.Stdout, "Response from `OsUsersAPI.GetObjectStorageUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | user id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStorageUserResp**](ObjectStorageUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageUserSamples

> ObjectStorageUserSamplesResp GetObjectStorageUserSamples(ctx, userId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	userId := int64(789) // int64 | user id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsUsersAPI.GetObjectStorageUserSamples(context.Background(), userId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsUsersAPI.GetObjectStorageUserSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageUserSamples`: ObjectStorageUserSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `OsUsersAPI.GetObjectStorageUserSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | user id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageUserSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**ObjectStorageUserSamplesResp**](ObjectStorageUserSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObjectStorageUsers

> ObjectStorageUsersResp ListObjectStorageUsers(ctx).ClusterId(clusterId).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsUsersAPI.ListObjectStorageUsers(context.Background()).ClusterId(clusterId).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsUsersAPI.ListObjectStorageUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListObjectStorageUsers`: ObjectStorageUsersResp
	fmt.Fprintf(os.Stdout, "Response from `OsUsersAPI.ListObjectStorageUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListObjectStorageUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**ObjectStorageUsersResp**](ObjectStorageUsersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateObjectStorageUser

> ObjectStorageUserResp UpdateObjectStorageUser(ctx, userId).Body(body).Execute()





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
	userId := int64(789) // int64 | user id
	body := *openapiclient.NewObjectStorageUserUpdateReq() // ObjectStorageUserUpdateReq | user info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsUsersAPI.UpdateObjectStorageUser(context.Background(), userId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsUsersAPI.UpdateObjectStorageUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateObjectStorageUser`: ObjectStorageUserResp
	fmt.Fprintf(os.Stdout, "Response from `OsUsersAPI.UpdateObjectStorageUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | user id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectStorageUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ObjectStorageUserUpdateReq**](ObjectStorageUserUpdateReq.md) | user info | 

### Return type

[**ObjectStorageUserResp**](ObjectStorageUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

