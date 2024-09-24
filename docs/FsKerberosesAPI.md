# \FsKerberosesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFSKerberos**](FsKerberosesAPI.md#CreateFSKerberos) | **Post** /fs-kerberoses/ | 
[**DeleteFSKerberos**](FsKerberosesAPI.md#DeleteFSKerberos) | **Delete** /fs-kerberoses/{fs_kerberos_id} | 
[**GetFSKerberos**](FsKerberosesAPI.md#GetFSKerberos) | **Get** /fs-kerberoses/{fs_kerberos_id} | 
[**ListFSKerberoses**](FsKerberosesAPI.md#ListFSKerberoses) | **Get** /fs-kerberoses/ | 
[**UpdateFSKerberos**](FsKerberosesAPI.md#UpdateFSKerberos) | **Patch** /fs-kerberoses/{fs_kerberos_id} | 
[**VerifyFSKerberos**](FsKerberosesAPI.md#VerifyFSKerberos) | **Post** /fs-kerberoses/{fs_kerberos_id}:verify | 



## CreateFSKerberos

> FSKerberosResp CreateFSKerberos(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewFSKerberosCreateReq(*openapiclient.NewFSKerberosCreateReqInfo("Kdc_example", "Name_example", "Password_example", "Realm_example", "Username_example")) // FSKerberosCreateReq | file storage kerberos info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsKerberosesAPI.CreateFSKerberos(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsKerberosesAPI.CreateFSKerberos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFSKerberos`: FSKerberosResp
	fmt.Fprintf(os.Stdout, "Response from `FsKerberosesAPI.CreateFSKerberos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFSKerberosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FSKerberosCreateReq**](FSKerberosCreateReq.md) | file storage kerberos info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**FSKerberosResp**](FSKerberosResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFSKerberos

> FSKerberosResp DeleteFSKerberos(ctx, fsKerberosId).Execute()





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
	fsKerberosId := int64(789) // int64 | file storage kerberos id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsKerberosesAPI.DeleteFSKerberos(context.Background(), fsKerberosId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsKerberosesAPI.DeleteFSKerberos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFSKerberos`: FSKerberosResp
	fmt.Fprintf(os.Stdout, "Response from `FsKerberosesAPI.DeleteFSKerberos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsKerberosId** | **int64** | file storage kerberos id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFSKerberosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FSKerberosResp**](FSKerberosResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFSKerberos

> FSKerberosResp GetFSKerberos(ctx, fsKerberosId).Execute()





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
	fsKerberosId := int64(789) // int64 | file storage kerberos id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsKerberosesAPI.GetFSKerberos(context.Background(), fsKerberosId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsKerberosesAPI.GetFSKerberos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFSKerberos`: FSKerberosResp
	fmt.Fprintf(os.Stdout, "Response from `FsKerberosesAPI.GetFSKerberos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsKerberosId** | **int64** | file storage kerberos id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFSKerberosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FSKerberosResp**](FSKerberosResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFSKerberoses

> FSKerberosesResp ListFSKerberoses(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).Q(q).Sort(sort).ActionStatus(actionStatus).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	actionStatus := "actionStatus_example" // string | kerberos action status: active, verifying, verifying_error (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsKerberosesAPI.ListFSKerberoses(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).Q(q).Sort(sort).ActionStatus(actionStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsKerberosesAPI.ListFSKerberoses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFSKerberoses`: FSKerberosesResp
	fmt.Fprintf(os.Stdout, "Response from `FsKerberosesAPI.ListFSKerberoses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFSKerberosesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **actionStatus** | **string** | kerberos action status: active, verifying, verifying_error | 

### Return type

[**FSKerberosesResp**](FSKerberosesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFSKerberos

> FSKerberosResp UpdateFSKerberos(ctx, fsKerberosId).Body(body).Execute()





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
	fsKerberosId := int64(789) // int64 | file storage kerberos id
	body := *openapiclient.NewFSKerberosUpdateReq(*openapiclient.NewFSKerberosUpdateReqInfo()) // FSKerberosUpdateReq | file storage kerberos info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsKerberosesAPI.UpdateFSKerberos(context.Background(), fsKerberosId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsKerberosesAPI.UpdateFSKerberos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFSKerberos`: FSKerberosResp
	fmt.Fprintf(os.Stdout, "Response from `FsKerberosesAPI.UpdateFSKerberos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsKerberosId** | **int64** | file storage kerberos id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFSKerberosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FSKerberosUpdateReq**](FSKerberosUpdateReq.md) | file storage kerberos info | 

### Return type

[**FSKerberosResp**](FSKerberosResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyFSKerberos

> FSKerberosResp VerifyFSKerberos(ctx, fsKerberosId).Execute()





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
	fsKerberosId := int64(789) // int64 | file storage kerberos id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsKerberosesAPI.VerifyFSKerberos(context.Background(), fsKerberosId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsKerberosesAPI.VerifyFSKerberos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyFSKerberos`: FSKerberosResp
	fmt.Fprintf(os.Stdout, "Response from `FsKerberosesAPI.VerifyFSKerberos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsKerberosId** | **int64** | file storage kerberos id | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyFSKerberosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FSKerberosResp**](FSKerberosResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

