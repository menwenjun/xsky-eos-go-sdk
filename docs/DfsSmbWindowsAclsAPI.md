# \DfsSmbWindowsAclsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDfsSMBWindowsACL**](DfsSmbWindowsAclsAPI.md#CreateDfsSMBWindowsACL) | **Post** /dfs-smb-windows-acls/ | 
[**DeleteDfsSMBWindowsACL**](DfsSmbWindowsAclsAPI.md#DeleteDfsSMBWindowsACL) | **Delete** /dfs-smb-windows-acls/ | 
[**GetDfsSMBWindowsACL**](DfsSmbWindowsAclsAPI.md#GetDfsSMBWindowsACL) | **Post** /dfs-smb-windows-acls/:get-windows-acl | 
[**SetDfsSMBWindowsACLs**](DfsSmbWindowsAclsAPI.md#SetDfsSMBWindowsACLs) | **Patch** /dfs-smb-windows-acls/:set-acls | 
[**UpdateDfsSMBWindowsACL**](DfsSmbWindowsAclsAPI.md#UpdateDfsSMBWindowsACL) | **Patch** /dfs-smb-windows-acls/ | 
[**UpdateDfsSMBWindowsACLInheritance**](DfsSmbWindowsAclsAPI.md#UpdateDfsSMBWindowsACLInheritance) | **Patch** /dfs-smb-windows-acls/:update-inheritance | 



## CreateDfsSMBWindowsACL

> DfsSMBWindowsACLResp CreateDfsSMBWindowsACL(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsSMBWindowsACLCreateReq(*openapiclient.NewDfsSMBWindowsACLCreateReqWindowsACL("AceType_example", "ApplyTo_example", int64(123), "Path_example", []string{"Permissions_example"}, "PrincipalType_example", "SecurityType_example")) // DfsSMBWindowsACLCreateReq | windows acl info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbWindowsAclsAPI.CreateDfsSMBWindowsACL(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbWindowsAclsAPI.CreateDfsSMBWindowsACL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsSMBWindowsACL`: DfsSMBWindowsACLResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbWindowsAclsAPI.CreateDfsSMBWindowsACL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsSMBWindowsACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsSMBWindowsACLCreateReq**](DfsSMBWindowsACLCreateReq.md) | windows acl info | 

### Return type

[**DfsSMBWindowsACLResp**](DfsSMBWindowsACLResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsSMBWindowsACL

> DfsSMBWindowsACLResp DeleteDfsSMBWindowsACL(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsSMBWindowsACLDeleteReq(*openapiclient.NewDfsSMBWindowsACLDeleteReqWindowsACL("AceType_example", "ApplyTo_example", int64(123), "Path_example", []string{"Permissions_example"}, "Principal_example")) // DfsSMBWindowsACLDeleteReq | windows acl info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbWindowsAclsAPI.DeleteDfsSMBWindowsACL(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbWindowsAclsAPI.DeleteDfsSMBWindowsACL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsSMBWindowsACL`: DfsSMBWindowsACLResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbWindowsAclsAPI.DeleteDfsSMBWindowsACL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsSMBWindowsACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsSMBWindowsACLDeleteReq**](DfsSMBWindowsACLDeleteReq.md) | windows acl info | 

### Return type

[**DfsSMBWindowsACLResp**](DfsSMBWindowsACLResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsSMBWindowsACL

> DfsSMBWindowsACLResp GetDfsSMBWindowsACL(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsSMBWindowsACLGetReq(*openapiclient.NewDfsSMBWindowsACLGetReqWindowsACL(int64(123), "Path_example")) // DfsSMBWindowsACLGetReq | windows acl info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbWindowsAclsAPI.GetDfsSMBWindowsACL(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbWindowsAclsAPI.GetDfsSMBWindowsACL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsSMBWindowsACL`: DfsSMBWindowsACLResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbWindowsAclsAPI.GetDfsSMBWindowsACL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsSMBWindowsACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsSMBWindowsACLGetReq**](DfsSMBWindowsACLGetReq.md) | windows acl info | 

### Return type

[**DfsSMBWindowsACLResp**](DfsSMBWindowsACLResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDfsSMBWindowsACLs

> DfsSMBWindowsACLResp SetDfsSMBWindowsACLs(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsSMBWindowsACLSetReq(int64(123), []openapiclient.DfsSMBWindowsACLSet{*openapiclient.NewDfsSMBWindowsACLSet("AceType_example", "ActionType_example", "ApplyTo_example", []string{"Permissions_example"}, "PrincipalType_example")}, "Path_example") // DfsSMBWindowsACLSetReq | windows acls info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbWindowsAclsAPI.SetDfsSMBWindowsACLs(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbWindowsAclsAPI.SetDfsSMBWindowsACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDfsSMBWindowsACLs`: DfsSMBWindowsACLResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbWindowsAclsAPI.SetDfsSMBWindowsACLs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDfsSMBWindowsACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsSMBWindowsACLSetReq**](DfsSMBWindowsACLSetReq.md) | windows acls info | 

### Return type

[**DfsSMBWindowsACLResp**](DfsSMBWindowsACLResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsSMBWindowsACL

> DfsSMBWindowsACLResp UpdateDfsSMBWindowsACL(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsSMBWindowsACLUpdateReq(*openapiclient.NewDfsSMBWindowsACLUpdateReqWindowsACL("AceType_example", "ApplyTo_example", int64(123), "Path_example", []string{"Permissions_example"}, "Principal_example")) // DfsSMBWindowsACLUpdateReq | windows acl info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbWindowsAclsAPI.UpdateDfsSMBWindowsACL(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbWindowsAclsAPI.UpdateDfsSMBWindowsACL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsSMBWindowsACL`: DfsSMBWindowsACLResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbWindowsAclsAPI.UpdateDfsSMBWindowsACL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsSMBWindowsACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsSMBWindowsACLUpdateReq**](DfsSMBWindowsACLUpdateReq.md) | windows acl info | 

### Return type

[**DfsSMBWindowsACLResp**](DfsSMBWindowsACLResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsSMBWindowsACLInheritance

> DfsSMBWindowsACLResp UpdateDfsSMBWindowsACLInheritance(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsSMBWindowsACLInheritanceUpdateReq(*openapiclient.NewDfsSMBWindowsACLInheritanceUpdateReqWindowsACL("ActionType_example", int64(123), "Path_example")) // DfsSMBWindowsACLInheritanceUpdateReq | windows acl info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbWindowsAclsAPI.UpdateDfsSMBWindowsACLInheritance(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbWindowsAclsAPI.UpdateDfsSMBWindowsACLInheritance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsSMBWindowsACLInheritance`: DfsSMBWindowsACLResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbWindowsAclsAPI.UpdateDfsSMBWindowsACLInheritance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsSMBWindowsACLInheritanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsSMBWindowsACLInheritanceUpdateReq**](DfsSMBWindowsACLInheritanceUpdateReq.md) | windows acl info | 

### Return type

[**DfsSMBWindowsACLResp**](DfsSMBWindowsACLResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

