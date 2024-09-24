# \FcPortsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearFCPortErrCode**](FcPortsAPI.md#ClearFCPortErrCode) | **Post** /fc-ports/{fc_port_id}:clear-err-code | 
[**GetFCPort**](FcPortsAPI.md#GetFCPort) | **Get** /fc-ports/{fc_port_id} | 
[**ListFCPorts**](FcPortsAPI.md#ListFCPorts) | **Get** /fc-ports/ | 
[**ResetFCPort**](FcPortsAPI.md#ResetFCPort) | **Post** /fc-ports/{fc_port_id}:reset | 
[**SetFCPort**](FcPortsAPI.md#SetFCPort) | **Patch** /fc-ports/{fc_port_id} | 



## ClearFCPortErrCode

> FCPortResp ClearFCPortErrCode(ctx, fcPortId).Execute()





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
	fcPortId := int64(789) // int64 | fc port id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FcPortsAPI.ClearFCPortErrCode(context.Background(), fcPortId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FcPortsAPI.ClearFCPortErrCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearFCPortErrCode`: FCPortResp
	fmt.Fprintf(os.Stdout, "Response from `FcPortsAPI.ClearFCPortErrCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fcPortId** | **int64** | fc port id | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearFCPortErrCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FCPortResp**](FCPortResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFCPort

> FCPortResp GetFCPort(ctx, fcPortId).Execute()





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
	fcPortId := int64(789) // int64 | fc port id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FcPortsAPI.GetFCPort(context.Background(), fcPortId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FcPortsAPI.GetFCPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFCPort`: FCPortResp
	fmt.Fprintf(os.Stdout, "Response from `FcPortsAPI.GetFCPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fcPortId** | **int64** | fc port id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFCPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FCPortResp**](FCPortResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFCPorts

> FCPortsResp ListFCPorts(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).Q(q).Sort(sort).Execute()





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
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FcPortsAPI.ListFCPorts(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FcPortsAPI.ListFCPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFCPorts`: FCPortsResp
	fmt.Fprintf(os.Stdout, "Response from `FcPortsAPI.ListFCPorts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFCPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**FCPortsResp**](FCPortsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetFCPort

> FCPortResp ResetFCPort(ctx, fcPortId).Execute()





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
	fcPortId := int64(789) // int64 | fc port id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FcPortsAPI.ResetFCPort(context.Background(), fcPortId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FcPortsAPI.ResetFCPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetFCPort`: FCPortResp
	fmt.Fprintf(os.Stdout, "Response from `FcPortsAPI.ResetFCPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fcPortId** | **int64** | fc port id | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetFCPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FCPortResp**](FCPortResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFCPort

> FCPortResp SetFCPort(ctx, fcPortId).Execute()





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
	fcPortId := int64(789) // int64 | fc port id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FcPortsAPI.SetFCPort(context.Background(), fcPortId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FcPortsAPI.SetFCPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFCPort`: FCPortResp
	fmt.Fprintf(os.Stdout, "Response from `FcPortsAPI.SetFCPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fcPortId** | **int64** | fc port id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFCPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FCPortResp**](FCPortResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

