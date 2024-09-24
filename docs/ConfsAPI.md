# \ConfsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfItem**](ConfsAPI.md#DeleteConfItem) | **Delete** /confs/{type}/{key} | 
[**GetConfItem**](ConfsAPI.md#GetConfItem) | **Get** /confs/{type}/{key} | 
[**ListConfItems**](ConfsAPI.md#ListConfItems) | **Get** /confs/{type} | 
[**ListConfItemsQuery**](ConfsAPI.md#ListConfItemsQuery) | **Get** /confs/ | 
[**SetConfItem**](ConfsAPI.md#SetConfItem) | **Put** /confs/ | 
[**SetConfItems**](ConfsAPI.md#SetConfItems) | **Put** /confs/{type} | 



## DeleteConfItem

> DeleteConfItem(ctx, type_, key).Execute()





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
	type_ := "type__example" // string | filter the type of conf item
	key := "key_example" // string | filter the key of conf item

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfsAPI.DeleteConfItem(context.Background(), type_, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfsAPI.DeleteConfItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | filter the type of conf item | 
**key** | **string** | filter the key of conf item | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfItemRequest struct via the builder pattern


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


## GetConfItem

> ConfItemResp GetConfItem(ctx, type_, key).ClusterId(clusterId).Execute()





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
	type_ := "type__example" // string | filter the type of conf
	key := "key_example" // string | filter the key of conf item
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfsAPI.GetConfItem(context.Background(), type_, key).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfsAPI.GetConfItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfItem`: ConfItemResp
	fmt.Fprintf(os.Stdout, "Response from `ConfsAPI.GetConfItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | filter the type of conf | 
**key** | **string** | filter the key of conf item | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterId** | **string** | cluster id | 

### Return type

[**ConfItemResp**](ConfItemResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfItems

> ConfItemsResp ListConfItems(ctx, type_).Limit(limit).Offset(offset).Execute()





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
	type_ := "type__example" // string | filter the type of conf
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfsAPI.ListConfItems(context.Background(), type_).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfsAPI.ListConfItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConfItems`: ConfItemsResp
	fmt.Fprintf(os.Stdout, "Response from `ConfsAPI.ListConfItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | filter the type of conf | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConfItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**ConfItemsResp**](ConfItemsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfItemsQuery

> ConfItemsResp ListConfItemsQuery(ctx).Limit(limit).Offset(offset).Type_(type_).Key(key).Execute()





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
	type_ := "type__example" // string | filter the type of conf (optional)
	key := "key_example" // string | filter the key of conf (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfsAPI.ListConfItemsQuery(context.Background()).Limit(limit).Offset(offset).Type_(type_).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfsAPI.ListConfItemsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConfItemsQuery`: ConfItemsResp
	fmt.Fprintf(os.Stdout, "Response from `ConfsAPI.ListConfItemsQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConfItemsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **type_** | **string** | filter the type of conf | 
 **key** | **string** | filter the key of conf | 

### Return type

[**ConfItemsResp**](ConfItemsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfItem

> ConfItem SetConfItem(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewConfItemSetReq(*openapiclient.NewConfItemSetReqConfItem("Key_example", "Type_example", "Value_example")) // ConfItemSetReq | conf item
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfsAPI.SetConfItem(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfsAPI.SetConfItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetConfItem`: ConfItem
	fmt.Fprintf(os.Stdout, "Response from `ConfsAPI.SetConfItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetConfItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfItemSetReq**](ConfItemSetReq.md) | conf item | 
 **clusterId** | **string** | cluster id | 

### Return type

[**ConfItem**](ConfItem.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfItems

> SetConfItems(ctx, type_).Body(body).ClusterId(clusterId).Execute()





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
	type_ := "type__example" // string | filter the type of conf
	body := *openapiclient.NewConfItemsSetReq(map[string]string{"key": "Inner_example"}) // ConfItemsSetReq | conf items
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfsAPI.SetConfItems(context.Background(), type_).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfsAPI.SetConfItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | filter the type of conf | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetConfItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConfItemsSetReq**](ConfItemsSetReq.md) | conf items | 
 **clusterId** | **string** | cluster id | 

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

