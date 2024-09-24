# \DpGatewaysAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpGateway**](DpGatewaysAPI.md#CreateDpGateway) | **Post** /dp-gateways/ | 
[**DeleteDpGateway**](DpGatewaysAPI.md#DeleteDpGateway) | **Delete** /dp-gateways/{gateway_id} | 
[**GetDpGateway**](DpGatewaysAPI.md#GetDpGateway) | **Get** /dp-gateways/{gateway_id} | 
[**ListDpGateways**](DpGatewaysAPI.md#ListDpGateways) | **Get** /dp-gateways/ | 
[**UpdateDpGateway**](DpGatewaysAPI.md#UpdateDpGateway) | **Patch** /dp-gateways/{gateway_id} | 



## CreateDpGateway

> DpGatewayResp CreateDpGateway(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDpGatewayCreateReq() // DpGatewayCreateReq | dp gateway info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpGatewaysAPI.CreateDpGateway(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpGatewaysAPI.CreateDpGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDpGateway`: DpGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `DpGatewaysAPI.CreateDpGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDpGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DpGatewayCreateReq**](DpGatewayCreateReq.md) | dp gateway info | 

### Return type

[**DpGatewayResp**](DpGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDpGateway

> DpGatewayResp DeleteDpGateway(ctx, gatewayId).Execute()





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
	gatewayId := int64(789) // int64 | dp gateway id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpGatewaysAPI.DeleteDpGateway(context.Background(), gatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpGatewaysAPI.DeleteDpGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDpGateway`: DpGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `DpGatewaysAPI.DeleteDpGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | dp gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDpGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpGatewayResp**](DpGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDpGateway

> DpGatewayResp GetDpGateway(ctx, gatewayId).Execute()





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
	gatewayId := int64(789) // int64 | dp gateway id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpGatewaysAPI.GetDpGateway(context.Background(), gatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpGatewaysAPI.GetDpGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDpGateway`: DpGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `DpGatewaysAPI.GetDpGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | dp gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDpGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpGatewayResp**](DpGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpGateways

> DpGatewaysResp ListDpGateways(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	resp, r, err := apiClient.DpGatewaysAPI.ListDpGateways(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpGatewaysAPI.ListDpGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpGateways`: DpGatewaysResp
	fmt.Fprintf(os.Stdout, "Response from `DpGatewaysAPI.ListDpGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DpGatewaysResp**](DpGatewaysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDpGateway

> DpGatewayResp UpdateDpGateway(ctx, gatewayId).Body(body).Execute()





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
	gatewayId := int64(789) // int64 | dp gateway id
	body := *openapiclient.NewDpGatewayUpdateReq() // DpGatewayUpdateReq | dp gateway info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpGatewaysAPI.UpdateDpGateway(context.Background(), gatewayId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpGatewaysAPI.UpdateDpGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDpGateway`: DpGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `DpGatewaysAPI.UpdateDpGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | dp gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDpGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DpGatewayUpdateReq**](DpGatewayUpdateReq.md) | dp gateway info | 

### Return type

[**DpGatewayResp**](DpGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

