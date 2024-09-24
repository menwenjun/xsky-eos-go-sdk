# \DfsS3KeysAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDfsS3Key**](DfsS3KeysAPI.md#CreateDfsS3Key) | **Post** /dfs-s3-keys/ | 
[**DeleteDfsS3Key**](DfsS3KeysAPI.md#DeleteDfsS3Key) | **Delete** /dfs-s3-keys/{key_id} | 
[**GetDfsS3Key**](DfsS3KeysAPI.md#GetDfsS3Key) | **Get** /dfs-s3-keys/{key_id} | 
[**ListDfsS3Keys**](DfsS3KeysAPI.md#ListDfsS3Keys) | **Get** /dfs-s3-keys/ | 
[**UpdateDfsS3Key**](DfsS3KeysAPI.md#UpdateDfsS3Key) | **Patch** /dfs-s3-keys/{key_id} | 



## CreateDfsS3Key

> DfsS3KeyResp CreateDfsS3Key(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewDfsS3KeyCreateReq() // DfsS3KeyCreateReq | key info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsS3KeysAPI.CreateDfsS3Key(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsS3KeysAPI.CreateDfsS3Key``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsS3Key`: DfsS3KeyResp
	fmt.Fprintf(os.Stdout, "Response from `DfsS3KeysAPI.CreateDfsS3Key`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsS3KeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsS3KeyCreateReq**](DfsS3KeyCreateReq.md) | key info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**DfsS3KeyResp**](DfsS3KeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsS3Key

> DfsS3KeyResp DeleteDfsS3Key(ctx, keyId).Execute()





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
	keyId := int64(789) // int64 | key id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsS3KeysAPI.DeleteDfsS3Key(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsS3KeysAPI.DeleteDfsS3Key``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsS3Key`: DfsS3KeyResp
	fmt.Fprintf(os.Stdout, "Response from `DfsS3KeysAPI.DeleteDfsS3Key`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **int64** | key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsS3KeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsS3KeyResp**](DfsS3KeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsS3Key

> DfsS3KeyResp GetDfsS3Key(ctx, keyId).Execute()





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
	keyId := int64(789) // int64 | user id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsS3KeysAPI.GetDfsS3Key(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsS3KeysAPI.GetDfsS3Key``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsS3Key`: DfsS3KeyResp
	fmt.Fprintf(os.Stdout, "Response from `DfsS3KeysAPI.GetDfsS3Key`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **int64** | user id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsS3KeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsS3KeyResp**](DfsS3KeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsS3Keys

> DfsS3KeysResp ListDfsS3Keys(ctx).ClusterId(clusterId).UserId(userId).Limit(limit).Offset(offset).Execute()





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
	userId := int64(789) // int64 | dfs s3 user id (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsS3KeysAPI.ListDfsS3Keys(context.Background()).ClusterId(clusterId).UserId(userId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsS3KeysAPI.ListDfsS3Keys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsS3Keys`: DfsS3KeysResp
	fmt.Fprintf(os.Stdout, "Response from `DfsS3KeysAPI.ListDfsS3Keys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsS3KeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **userId** | **int64** | dfs s3 user id | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**DfsS3KeysResp**](DfsS3KeysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsS3Key

> DfsS3KeyResp UpdateDfsS3Key(ctx, keyId).Body(body).Execute()





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
	keyId := int64(789) // int64 | key id
	body := *openapiclient.NewDfsS3KeyUpdateReq() // DfsS3KeyUpdateReq | key info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsS3KeysAPI.UpdateDfsS3Key(context.Background(), keyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsS3KeysAPI.UpdateDfsS3Key``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsS3Key`: DfsS3KeyResp
	fmt.Fprintf(os.Stdout, "Response from `DfsS3KeysAPI.UpdateDfsS3Key`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **int64** | key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsS3KeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsS3KeyUpdateReq**](DfsS3KeyUpdateReq.md) | key info | 

### Return type

[**DfsS3KeyResp**](DfsS3KeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

