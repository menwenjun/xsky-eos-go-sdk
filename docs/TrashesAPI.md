# \TrashesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CleanupTrash**](TrashesAPI.md#CleanupTrash) | **Post** /trashes/{trash_id}:cleanup | 
[**GetTrash**](TrashesAPI.md#GetTrash) | **Get** /trashes/{trash_id} | 
[**ListTrashes**](TrashesAPI.md#ListTrashes) | **Get** /trashes/ | 
[**UpdateTrash**](TrashesAPI.md#UpdateTrash) | **Patch** /trashes/{trash_id} | 



## CleanupTrash

> TrashResp CleanupTrash(ctx, trashId).Execute()





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
	trashId := int64(789) // int64 | trash id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrashesAPI.CleanupTrash(context.Background(), trashId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrashesAPI.CleanupTrash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CleanupTrash`: TrashResp
	fmt.Fprintf(os.Stdout, "Response from `TrashesAPI.CleanupTrash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trashId** | **int64** | trash id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCleanupTrashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrashResp**](TrashResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrash

> TrashResp GetTrash(ctx, trashId).Execute()





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
	trashId := int64(789) // int64 | trash id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrashesAPI.GetTrash(context.Background(), trashId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrashesAPI.GetTrash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrash`: TrashResp
	fmt.Fprintf(os.Stdout, "Response from `TrashesAPI.GetTrash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trashId** | **int64** | trash id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrashResp**](TrashResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrashes

> TrashesResp ListTrashes(ctx).Limit(limit).Offset(offset).Type_(type_).Q(q).Sort(sort).Execute()





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
	type_ := "type__example" // string | the type of trash (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrashesAPI.ListTrashes(context.Background()).Limit(limit).Offset(offset).Type_(type_).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrashesAPI.ListTrashes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTrashes`: TrashesResp
	fmt.Fprintf(os.Stdout, "Response from `TrashesAPI.ListTrashes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTrashesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **type_** | **string** | the type of trash | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**TrashesResp**](TrashesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrash

> TrashResp UpdateTrash(ctx, trashId).Body(body).Execute()





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
	trashId := int64(789) // int64 | trash id
	body := *openapiclient.NewTrashUpdateReq(*openapiclient.NewTrashUpdateReqTrash()) // TrashUpdateReq | trash info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrashesAPI.UpdateTrash(context.Background(), trashId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrashesAPI.UpdateTrash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTrash`: TrashResp
	fmt.Fprintf(os.Stdout, "Response from `TrashesAPI.UpdateTrash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trashId** | **int64** | trash id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TrashUpdateReq**](TrashUpdateReq.md) | trash info | 

### Return type

[**TrashResp**](TrashResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

