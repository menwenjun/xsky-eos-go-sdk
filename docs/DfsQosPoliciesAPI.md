# \DfsQosPoliciesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDfsQosPolicy**](DfsQosPoliciesAPI.md#CreateDfsQosPolicy) | **Post** /dfs-qos-policies/ | 
[**DeleteDfsQosPolicy**](DfsQosPoliciesAPI.md#DeleteDfsQosPolicy) | **Delete** /dfs-qos-policies/{policy_id} | 
[**GetDfsQosPolicy**](DfsQosPoliciesAPI.md#GetDfsQosPolicy) | **Get** /dfs-qos-policies/{policy_id} | 
[**ListDfsQosPolicies**](DfsQosPoliciesAPI.md#ListDfsQosPolicies) | **Get** /dfs-qos-policies/ | 
[**UpdateDfsQosPolicy**](DfsQosPoliciesAPI.md#UpdateDfsQosPolicy) | **Patch** /dfs-qos-policies/{policy_id} | 



## CreateDfsQosPolicy

> DfsQosPolicyResp CreateDfsQosPolicy(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewDfsQosPolicyCreateReq(*openapiclient.NewDfsQosPolicyCreateReqQosPolicy("Name_example")) // DfsQosPolicyCreateReq | qos policy info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQosPoliciesAPI.CreateDfsQosPolicy(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQosPoliciesAPI.CreateDfsQosPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsQosPolicy`: DfsQosPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQosPoliciesAPI.CreateDfsQosPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsQosPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsQosPolicyCreateReq**](DfsQosPolicyCreateReq.md) | qos policy info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**DfsQosPolicyResp**](DfsQosPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsQosPolicy

> DfsQosPolicyResp DeleteDfsQosPolicy(ctx, policyId).ClusterId(clusterId).Execute()





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
	policyId := int64(789) // int64 | policy id
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQosPoliciesAPI.DeleteDfsQosPolicy(context.Background(), policyId).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQosPoliciesAPI.DeleteDfsQosPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsQosPolicy`: DfsQosPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQosPoliciesAPI.DeleteDfsQosPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsQosPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterId** | **string** | cluster id | 

### Return type

[**DfsQosPolicyResp**](DfsQosPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsQosPolicy

> DfsQosPolicyResp GetDfsQosPolicy(ctx, policyId).Execute()





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
	policyId := int64(789) // int64 | dfs qos policy id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQosPoliciesAPI.GetDfsQosPolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQosPoliciesAPI.GetDfsQosPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsQosPolicy`: DfsQosPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQosPoliciesAPI.GetDfsQosPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | dfs qos policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsQosPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsQosPolicyResp**](DfsQosPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsQosPolicies

> DfsQosPoliciesResp ListDfsQosPolicies(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQosPoliciesAPI.ListDfsQosPolicies(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQosPoliciesAPI.ListDfsQosPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsQosPolicies`: DfsQosPoliciesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQosPoliciesAPI.ListDfsQosPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsQosPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DfsQosPoliciesResp**](DfsQosPoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsQosPolicy

> DfsQosPolicyResp UpdateDfsQosPolicy(ctx, policyId).Body(body).Execute()





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
	policyId := int64(789) // int64 | policy id
	body := *openapiclient.NewDfsQosPolicyUpdateReq(*openapiclient.NewDfsQosPolicyUpdateReqQosPolicy("Name_example")) // DfsQosPolicyUpdateReq | dfs qos policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQosPoliciesAPI.UpdateDfsQosPolicy(context.Background(), policyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQosPoliciesAPI.UpdateDfsQosPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsQosPolicy`: DfsQosPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQosPoliciesAPI.UpdateDfsQosPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsQosPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsQosPolicyUpdateReq**](DfsQosPolicyUpdateReq.md) | dfs qos policy info | 

### Return type

[**DfsQosPolicyResp**](DfsQosPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

