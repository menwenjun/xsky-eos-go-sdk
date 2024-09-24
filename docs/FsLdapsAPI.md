# \FsLdapsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFSLdap**](FsLdapsAPI.md#CreateFSLdap) | **Post** /fs-ldaps/ | 
[**DeleteFSLdap**](FsLdapsAPI.md#DeleteFSLdap) | **Delete** /fs-ldaps/{fs_ldap_id} | 
[**GetFSLdap**](FsLdapsAPI.md#GetFSLdap) | **Get** /fs-ldaps/{fs_ldap_id} | 
[**ListFSLdaps**](FsLdapsAPI.md#ListFSLdaps) | **Get** /fs-ldaps/ | 
[**UpdateFSLdap**](FsLdapsAPI.md#UpdateFSLdap) | **Patch** /fs-ldaps/{fs_ldap_id} | 
[**VerifyFSLdap**](FsLdapsAPI.md#VerifyFSLdap) | **Post** /fs-ldaps/{fs_ldap_id}:verify | 



## CreateFSLdap

> FSLdapResp CreateFSLdap(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewFSLdapCreateReq(*openapiclient.NewFSLdapCreateReqInfo("Ip_example", "Name_example", int64(123), "Suffix_example")) // FSLdapCreateReq | file storage ldap info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsLdapsAPI.CreateFSLdap(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsLdapsAPI.CreateFSLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFSLdap`: FSLdapResp
	fmt.Fprintf(os.Stdout, "Response from `FsLdapsAPI.CreateFSLdap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFSLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FSLdapCreateReq**](FSLdapCreateReq.md) | file storage ldap info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**FSLdapResp**](FSLdapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFSLdap

> DeleteFSLdap(ctx, fsLdapId).Execute()





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
	fsLdapId := int64(789) // int64 | file storage ldap id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FsLdapsAPI.DeleteFSLdap(context.Background(), fsLdapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsLdapsAPI.DeleteFSLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsLdapId** | **int64** | file storage ldap id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFSLdapRequest struct via the builder pattern


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


## GetFSLdap

> FSLdapResp GetFSLdap(ctx, fsLdapId).Execute()





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
	fsLdapId := int64(789) // int64 | file storage ldap id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsLdapsAPI.GetFSLdap(context.Background(), fsLdapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsLdapsAPI.GetFSLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFSLdap`: FSLdapResp
	fmt.Fprintf(os.Stdout, "Response from `FsLdapsAPI.GetFSLdap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsLdapId** | **int64** | file storage ldap id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFSLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FSLdapResp**](FSLdapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFSLdaps

> FSLdapsResp ListFSLdaps(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsLdapsAPI.ListFSLdaps(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsLdapsAPI.ListFSLdaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFSLdaps`: FSLdapsResp
	fmt.Fprintf(os.Stdout, "Response from `FsLdapsAPI.ListFSLdaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFSLdapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 

### Return type

[**FSLdapsResp**](FSLdapsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFSLdap

> FSLdapResp UpdateFSLdap(ctx, fsLdapId).Body(body).Execute()





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
	fsLdapId := int64(789) // int64 | file storage ldap id
	body := *openapiclient.NewFSLdapUpdateReq(*openapiclient.NewFSLdapUpdateReqInfo("Ip_example", "Name_example", int64(123), "Suffix_example")) // FSLdapUpdateReq | file storage ldap info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsLdapsAPI.UpdateFSLdap(context.Background(), fsLdapId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsLdapsAPI.UpdateFSLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFSLdap`: FSLdapResp
	fmt.Fprintf(os.Stdout, "Response from `FsLdapsAPI.UpdateFSLdap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsLdapId** | **int64** | file storage ldap id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFSLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FSLdapUpdateReq**](FSLdapUpdateReq.md) | file storage ldap info | 

### Return type

[**FSLdapResp**](FSLdapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyFSLdap

> FSLdapResp VerifyFSLdap(ctx, fsLdapId).Execute()





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
	fsLdapId := int64(789) // int64 | file storage ldap id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FsLdapsAPI.VerifyFSLdap(context.Background(), fsLdapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FsLdapsAPI.VerifyFSLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyFSLdap`: FSLdapResp
	fmt.Fprintf(os.Stdout, "Response from `FsLdapsAPI.VerifyFSLdap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsLdapId** | **int64** | file storage ldap id | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyFSLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FSLdapResp**](FSLdapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

