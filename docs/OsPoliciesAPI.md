# \OsPoliciesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicy**](OsPoliciesAPI.md#CreatePolicy) | **Post** /os-policies/ | 
[**DeletePolicy**](OsPoliciesAPI.md#DeletePolicy) | **Delete** /os-policies/{policy_id} | 
[**GetPolicy**](OsPoliciesAPI.md#GetPolicy) | **Get** /os-policies/{policy_id} | 
[**ListPolicies**](OsPoliciesAPI.md#ListPolicies) | **Get** /os-policies/ | 
[**UpdatePolicy**](OsPoliciesAPI.md#UpdatePolicy) | **Patch** /os-policies/{policy_id} | 



## CreatePolicy

> ObjectStoragePolicyResp CreatePolicy(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewObjectStoragePolicyCreateReq() // ObjectStoragePolicyCreateReq | policy info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsPoliciesAPI.CreatePolicy(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsPoliciesAPI.CreatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePolicy`: ObjectStoragePolicyResp
	fmt.Fprintf(os.Stdout, "Response from `OsPoliciesAPI.CreatePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ObjectStoragePolicyCreateReq**](ObjectStoragePolicyCreateReq.md) | policy info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**ObjectStoragePolicyResp**](ObjectStoragePolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicy

> ObjectStoragePolicyResp DeletePolicy(ctx, policyId).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsPoliciesAPI.DeletePolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsPoliciesAPI.DeletePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePolicy`: ObjectStoragePolicyResp
	fmt.Fprintf(os.Stdout, "Response from `OsPoliciesAPI.DeletePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStoragePolicyResp**](ObjectStoragePolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicy

> ObjectStoragePolicyResp GetPolicy(ctx, policyId).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsPoliciesAPI.GetPolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsPoliciesAPI.GetPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicy`: ObjectStoragePolicyResp
	fmt.Fprintf(os.Stdout, "Response from `OsPoliciesAPI.GetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStoragePolicyResp**](ObjectStoragePolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicies

> ObjectStoragePoliciesResp ListPolicies(ctx).ClusterId(clusterId).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsPoliciesAPI.ListPolicies(context.Background()).ClusterId(clusterId).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsPoliciesAPI.ListPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPolicies`: ObjectStoragePoliciesResp
	fmt.Fprintf(os.Stdout, "Response from `OsPoliciesAPI.ListPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**ObjectStoragePoliciesResp**](ObjectStoragePoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicy

> ObjectStoragePolicyResp UpdatePolicy(ctx, policyId).Body(body).Execute()





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
	body := *openapiclient.NewObjectStoragePolicyUpdateReq() // ObjectStoragePolicyUpdateReq | policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsPoliciesAPI.UpdatePolicy(context.Background(), policyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsPoliciesAPI.UpdatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePolicy`: ObjectStoragePolicyResp
	fmt.Fprintf(os.Stdout, "Response from `OsPoliciesAPI.UpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ObjectStoragePolicyUpdateReq**](ObjectStoragePolicyUpdateReq.md) | policy info | 

### Return type

[**ObjectStoragePolicyResp**](ObjectStoragePolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

