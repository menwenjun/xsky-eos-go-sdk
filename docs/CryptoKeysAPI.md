# \CryptoKeysAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCryptoKey**](CryptoKeysAPI.md#CreateCryptoKey) | **Post** /crypto-keys/ | 
[**DownloadCryptoKey**](CryptoKeysAPI.md#DownloadCryptoKey) | **Get** /crypto-keys/{key_id}/key | 
[**GetCryptoKey**](CryptoKeysAPI.md#GetCryptoKey) | **Get** /crypto-keys/{key_id} | 
[**ListCryptoKeys**](CryptoKeysAPI.md#ListCryptoKeys) | **Get** /crypto-keys/ | 
[**UpdateCryptoKey**](CryptoKeysAPI.md#UpdateCryptoKey) | **Patch** /crypto-keys/{key_id} | 



## CreateCryptoKey

> CryptoKeyResp CreateCryptoKey(ctx).Name(name).Key(key).Execute()





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
	name := "name_example" // string | crypto key name
	key := "key_example" // string | crypto key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoKeysAPI.CreateCryptoKey(context.Background()).Name(name).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoKeysAPI.CreateCryptoKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCryptoKey`: CryptoKeyResp
	fmt.Fprintf(os.Stdout, "Response from `CryptoKeysAPI.CreateCryptoKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCryptoKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | crypto key name | 
 **key** | **string** | crypto key | 

### Return type

[**CryptoKeyResp**](CryptoKeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadCryptoKey

> *os.File DownloadCryptoKey(ctx, keyId).Password(password).Execute()



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
	keyId := int64(789) // int64 | crypto key id
	password := "password_example" // string | password (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoKeysAPI.DownloadCryptoKey(context.Background(), keyId).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoKeysAPI.DownloadCryptoKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadCryptoKey`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CryptoKeysAPI.DownloadCryptoKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **int64** | crypto key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadCryptoKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **password** | **string** | password | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCryptoKey

> CryptoKeyResp GetCryptoKey(ctx, keyId).Execute()





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
	keyId := int64(789) // int64 | crypto key id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoKeysAPI.GetCryptoKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoKeysAPI.GetCryptoKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCryptoKey`: CryptoKeyResp
	fmt.Fprintf(os.Stdout, "Response from `CryptoKeysAPI.GetCryptoKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **int64** | crypto key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCryptoKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CryptoKeyResp**](CryptoKeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCryptoKeys

> CryptoKeysResp ListCryptoKeys(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoKeysAPI.ListCryptoKeys(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoKeysAPI.ListCryptoKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCryptoKeys`: CryptoKeysResp
	fmt.Fprintf(os.Stdout, "Response from `CryptoKeysAPI.ListCryptoKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCryptoKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**CryptoKeysResp**](CryptoKeysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCryptoKey

> CryptoKeyResp UpdateCryptoKey(ctx, keyId).Body(body).Execute()





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
	keyId := int64(789) // int64 | crypto key id
	body := *openapiclient.NewCryptoKeyUpdateReq() // CryptoKeyUpdateReq | crypto key info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoKeysAPI.UpdateCryptoKey(context.Background(), keyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoKeysAPI.UpdateCryptoKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCryptoKey`: CryptoKeyResp
	fmt.Fprintf(os.Stdout, "Response from `CryptoKeysAPI.UpdateCryptoKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **int64** | crypto key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCryptoKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CryptoKeyUpdateReq**](CryptoKeyUpdateReq.md) | crypto key info | 

### Return type

[**CryptoKeyResp**](CryptoKeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

