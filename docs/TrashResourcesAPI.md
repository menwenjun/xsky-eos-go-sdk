# \TrashResourcesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTrashResource**](TrashResourcesAPI.md#DeleteTrashResource) | **Delete** /trash-resources/{trash_resource_id} | 
[**GetTrashResource**](TrashResourcesAPI.md#GetTrashResource) | **Get** /trash-resources/{trash_resource_id} | 
[**ListTrashResources**](TrashResourcesAPI.md#ListTrashResources) | **Get** /trash-resources/ | 
[**RestoreTrashResource**](TrashResourcesAPI.md#RestoreTrashResource) | **Post** /trash-resources/{trash_resource_id}:restore | 



## DeleteTrashResource

> TrashResourceResp DeleteTrashResource(ctx, trashResourceId).Execute()





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
	trashResourceId := int64(789) // int64 | trash resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrashResourcesAPI.DeleteTrashResource(context.Background(), trashResourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrashResourcesAPI.DeleteTrashResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTrashResource`: TrashResourceResp
	fmt.Fprintf(os.Stdout, "Response from `TrashResourcesAPI.DeleteTrashResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trashResourceId** | **int64** | trash resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrashResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrashResourceResp**](TrashResourceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrashResource

> TrashResourceResp GetTrashResource(ctx, trashResourceId).Execute()





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
	trashResourceId := int64(789) // int64 | trash resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrashResourcesAPI.GetTrashResource(context.Background(), trashResourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrashResourcesAPI.GetTrashResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrashResource`: TrashResourceResp
	fmt.Fprintf(os.Stdout, "Response from `TrashResourcesAPI.GetTrashResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trashResourceId** | **int64** | trash resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrashResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrashResourceResp**](TrashResourceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrashResources

> TrashResourcesResp ListTrashResources(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).Type_(type_).Q(q).Sort(sort).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	type_ := "type__example" // string | the type of trash resources (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrashResourcesAPI.ListTrashResources(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).Type_(type_).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrashResourcesAPI.ListTrashResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTrashResources`: TrashResourcesResp
	fmt.Fprintf(os.Stdout, "Response from `TrashResourcesAPI.ListTrashResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTrashResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **type_** | **string** | the type of trash resources | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**TrashResourcesResp**](TrashResourcesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreTrashResource

> TrashResourceResp RestoreTrashResource(ctx, trashResourceId).Execute()





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
	trashResourceId := int64(789) // int64 | trash resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrashResourcesAPI.RestoreTrashResource(context.Background(), trashResourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrashResourcesAPI.RestoreTrashResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreTrashResource`: TrashResourceResp
	fmt.Fprintf(os.Stdout, "Response from `TrashResourcesAPI.RestoreTrashResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trashResourceId** | **int64** | trash resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreTrashResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrashResourceResp**](TrashResourceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

