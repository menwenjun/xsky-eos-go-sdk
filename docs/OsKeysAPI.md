# \OsKeysAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKey**](OsKeysAPI.md#CreateKey) | **Post** /os-keys/ | 
[**DeleteKey**](OsKeysAPI.md#DeleteKey) | **Delete** /os-keys/{key_id} | 
[**GetKey**](OsKeysAPI.md#GetKey) | **Get** /os-keys/{key_id} | 
[**ListKeys**](OsKeysAPI.md#ListKeys) | **Get** /os-keys/ | 
[**UpdateKey**](OsKeysAPI.md#UpdateKey) | **Patch** /os-keys/{key_id} | 



## CreateKey

> ObjectStorageKeyResp CreateKey(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewObjectStorageKeyCreateReq() // ObjectStorageKeyCreateReq | key info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsKeysAPI.CreateKey(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsKeysAPI.CreateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKey`: ObjectStorageKeyResp
	fmt.Fprintf(os.Stdout, "Response from `OsKeysAPI.CreateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ObjectStorageKeyCreateReq**](ObjectStorageKeyCreateReq.md) | key info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**ObjectStorageKeyResp**](ObjectStorageKeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKey

> ObjectStorageKeyResp DeleteKey(ctx, keyId).Execute()





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
	keyId := int64(789) // int64 | key id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsKeysAPI.DeleteKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsKeysAPI.DeleteKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKey`: ObjectStorageKeyResp
	fmt.Fprintf(os.Stdout, "Response from `OsKeysAPI.DeleteKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **int64** | key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStorageKeyResp**](ObjectStorageKeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKey

> ObjectStorageKeyResp GetKey(ctx, keyId).Execute()





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
	keyId := int64(789) // int64 | user id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsKeysAPI.GetKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsKeysAPI.GetKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKey`: ObjectStorageKeyResp
	fmt.Fprintf(os.Stdout, "Response from `OsKeysAPI.GetKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **int64** | user id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStorageKeyResp**](ObjectStorageKeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeys

> ObjectStorageKeysResp ListKeys(ctx).ClusterId(clusterId).Limit(limit).Offset(offset).UserId(userId).Execute()





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
	userId := int64(789) // int64 | object storage user id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsKeysAPI.ListKeys(context.Background()).ClusterId(clusterId).Limit(limit).Offset(offset).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsKeysAPI.ListKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKeys`: ObjectStorageKeysResp
	fmt.Fprintf(os.Stdout, "Response from `OsKeysAPI.ListKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **userId** | **int64** | object storage user id | 

### Return type

[**ObjectStorageKeysResp**](ObjectStorageKeysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKey

> ObjectStorageKeyResp UpdateKey(ctx, keyId).Body(body).Execute()





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
	keyId := int64(789) // int64 | key id
	body := *openapiclient.NewObjectStorageKeyUpdateReq() // ObjectStorageKeyUpdateReq | key info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsKeysAPI.UpdateKey(context.Background(), keyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsKeysAPI.UpdateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKey`: ObjectStorageKeyResp
	fmt.Fprintf(os.Stdout, "Response from `OsKeysAPI.UpdateKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **int64** | key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ObjectStorageKeyUpdateReq**](ObjectStorageKeyUpdateReq.md) | key info | 

### Return type

[**ObjectStorageKeyResp**](ObjectStorageKeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

