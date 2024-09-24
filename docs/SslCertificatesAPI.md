# \SslCertificatesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSSLCertificate**](SslCertificatesAPI.md#CreateSSLCertificate) | **Post** /ssl-certificates/ | 
[**DeleteSSLCertificate**](SslCertificatesAPI.md#DeleteSSLCertificate) | **Delete** /ssl-certificates/{certificate_id} | 
[**GetSSLCertificate**](SslCertificatesAPI.md#GetSSLCertificate) | **Get** /ssl-certificates/{certificate_id} | 
[**ListSSLCertificates**](SslCertificatesAPI.md#ListSSLCertificates) | **Get** /ssl-certificates/ | 
[**UpdateSSLCertificate**](SslCertificatesAPI.md#UpdateSSLCertificate) | **Patch** /ssl-certificates/{certificate_id} | 



## CreateSSLCertificate

> SSLCertificateResp CreateSSLCertificate(ctx).Body(body).Execute()





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
	body := *openapiclient.NewSSLCertificateCreateReq(*openapiclient.NewSSLCertificateCreateReqCertificate("Certificate_example", "Name_example", "PrivateKey_example")) // SSLCertificateCreateReq | ssl certificate info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SslCertificatesAPI.CreateSSLCertificate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SslCertificatesAPI.CreateSSLCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSSLCertificate`: SSLCertificateResp
	fmt.Fprintf(os.Stdout, "Response from `SslCertificatesAPI.CreateSSLCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSLCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SSLCertificateCreateReq**](SSLCertificateCreateReq.md) | ssl certificate info | 

### Return type

[**SSLCertificateResp**](SSLCertificateResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSSLCertificate

> SSLCertificateResp DeleteSSLCertificate(ctx, certificateId).Execute()





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
	certificateId := int64(789) // int64 | certificate id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SslCertificatesAPI.DeleteSSLCertificate(context.Background(), certificateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SslCertificatesAPI.DeleteSSLCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSSLCertificate`: SSLCertificateResp
	fmt.Fprintf(os.Stdout, "Response from `SslCertificatesAPI.DeleteSSLCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **int64** | certificate id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSLCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SSLCertificateResp**](SSLCertificateResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSLCertificate

> SSLCertificateResp GetSSLCertificate(ctx, certificateId).Limit(limit).Offset(offset).Execute()





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
	certificateId := int64(789) // int64 | certificate id
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SslCertificatesAPI.GetSSLCertificate(context.Background(), certificateId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SslCertificatesAPI.GetSSLCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSSLCertificate`: SSLCertificateResp
	fmt.Fprintf(os.Stdout, "Response from `SslCertificatesAPI.GetSSLCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **int64** | certificate id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSLCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**SSLCertificateResp**](SSLCertificateResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSSLCertificates

> SSLCertificatesResp ListSSLCertificates(ctx).Limit(limit).Offset(offset).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SslCertificatesAPI.ListSSLCertificates(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SslCertificatesAPI.ListSSLCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSSLCertificates`: SSLCertificatesResp
	fmt.Fprintf(os.Stdout, "Response from `SslCertificatesAPI.ListSSLCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSSLCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**SSLCertificatesResp**](SSLCertificatesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSSLCertificate

> SSLCertificateResp UpdateSSLCertificate(ctx, certificateId).Body(body).Force(force).Execute()





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
	certificateId := int64(789) // int64 | certificate id
	body := *openapiclient.NewSSLCertificateUpdateReq(*openapiclient.NewSSLCertificateUpdateReqCertificate()) // SSLCertificateUpdateReq | ssl certificate info
	force := true // bool | force switch certification type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SslCertificatesAPI.UpdateSSLCertificate(context.Background(), certificateId).Body(body).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SslCertificatesAPI.UpdateSSLCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSSLCertificate`: SSLCertificateResp
	fmt.Fprintf(os.Stdout, "Response from `SslCertificatesAPI.UpdateSSLCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **int64** | certificate id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSSLCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SSLCertificateUpdateReq**](SSLCertificateUpdateReq.md) | ssl certificate info | 
 **force** | **bool** | force switch certification type | 

### Return type

[**SSLCertificateResp**](SSLCertificateResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

