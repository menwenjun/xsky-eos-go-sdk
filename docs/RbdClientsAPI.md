# \RbdClientsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRBDC**](RbdClientsAPI.md#CreateRBDC) | **Post** /rbd-clients/ | 
[**DeleteRBDClient**](RbdClientsAPI.md#DeleteRBDClient) | **Delete** /rbd-clients/{rbdc_id} | 
[**GetRBDClient**](RbdClientsAPI.md#GetRBDClient) | **Get** /rbd-clients/{rbdc_id} | 
[**ListRBDClients**](RbdClientsAPI.md#ListRBDClients) | **Get** /rbd-clients/ | 
[**UpdateRBDClient**](RbdClientsAPI.md#UpdateRBDClient) | **Patch** /rbd-clients/{rbdc_id} | 
[**ValidateRBDClientHost**](RbdClientsAPI.md#ValidateRBDClientHost) | **Post** /rbd-clients/:validate-host | 



## CreateRBDC

> RBDClientResp CreateRBDC(ctx).Body(body).Execute()





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
	body := *openapiclient.NewRBDClientCreateReq() // RBDClientCreateReq | rbdc info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbdClientsAPI.CreateRBDC(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdClientsAPI.CreateRBDC``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRBDC`: RBDClientResp
	fmt.Fprintf(os.Stdout, "Response from `RbdClientsAPI.CreateRBDC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRBDCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RBDClientCreateReq**](RBDClientCreateReq.md) | rbdc info | 

### Return type

[**RBDClientResp**](RBDClientResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRBDClient

> RBDClientResp DeleteRBDClient(ctx, rbdcId).Force(force).SkipUninstall(skipUninstall).Execute()





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
	rbdcId := int64(789) // int64 | rbdc id
	force := true // bool | force delete or not (optional)
	skipUninstall := true // bool | skip uninstallation for managed rbdc (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbdClientsAPI.DeleteRBDClient(context.Background(), rbdcId).Force(force).SkipUninstall(skipUninstall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdClientsAPI.DeleteRBDClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRBDClient`: RBDClientResp
	fmt.Fprintf(os.Stdout, "Response from `RbdClientsAPI.DeleteRBDClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rbdcId** | **int64** | rbdc id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRBDClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 
 **skipUninstall** | **bool** | skip uninstallation for managed rbdc | 

### Return type

[**RBDClientResp**](RBDClientResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRBDClient

> RBDClientResp GetRBDClient(ctx, rbdcId).Execute()





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
	rbdcId := int64(789) // int64 | the rbd client id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbdClientsAPI.GetRBDClient(context.Background(), rbdcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdClientsAPI.GetRBDClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRBDClient`: RBDClientResp
	fmt.Fprintf(os.Stdout, "Response from `RbdClientsAPI.GetRBDClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rbdcId** | **int64** | the rbd client id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRBDClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RBDClientResp**](RBDClientResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRBDClients

> RBDClientsResp ListRBDClients(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).Type_(type_).Q(q).Sort(sort).Execute()





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
	type_ := "type__example" // string | filter by rbdc type (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbdClientsAPI.ListRBDClients(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).Type_(type_).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdClientsAPI.ListRBDClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRBDClients`: RBDClientsResp
	fmt.Fprintf(os.Stdout, "Response from `RbdClientsAPI.ListRBDClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRBDClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **type_** | **string** | filter by rbdc type | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**RBDClientsResp**](RBDClientsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRBDClient

> RBDClientResp UpdateRBDClient(ctx, rbdcId).Body(body).Execute()





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
	rbdcId := int64(789) // int64 | rbdc id
	body := *openapiclient.NewRBDClientUpdateReq(*openapiclient.NewRBDClientUpdateReqRBDClient()) // RBDClientUpdateReq | rbdc info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbdClientsAPI.UpdateRBDClient(context.Background(), rbdcId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdClientsAPI.UpdateRBDClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRBDClient`: RBDClientResp
	fmt.Fprintf(os.Stdout, "Response from `RbdClientsAPI.UpdateRBDClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rbdcId** | **int64** | rbdc id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRBDClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RBDClientUpdateReq**](RBDClientUpdateReq.md) | rbdc info | 

### Return type

[**RBDClientResp**](RBDClientResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateRBDClientHost

> RBDClientValidateHostResp ValidateRBDClientHost(ctx).Body(body).Execute()





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
	body := *openapiclient.NewRBDClientValidateHostReq() // RBDClientValidateHostReq | validate rbdc info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbdClientsAPI.ValidateRBDClientHost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdClientsAPI.ValidateRBDClientHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateRBDClientHost`: RBDClientValidateHostResp
	fmt.Fprintf(os.Stdout, "Response from `RbdClientsAPI.ValidateRBDClientHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateRBDClientHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RBDClientValidateHostReq**](RBDClientValidateHostReq.md) | validate rbdc info | 

### Return type

[**RBDClientValidateHostResp**](RBDClientValidateHostResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

