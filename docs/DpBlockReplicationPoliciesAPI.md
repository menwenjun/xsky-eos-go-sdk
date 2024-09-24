# \DpBlockReplicationPoliciesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpBlockReplicationPolicy**](DpBlockReplicationPoliciesAPI.md#CreateDpBlockReplicationPolicy) | **Post** /dp-block-replication-policies/ | 
[**DeleteDpBlockReplicationPolicy**](DpBlockReplicationPoliciesAPI.md#DeleteDpBlockReplicationPolicy) | **Delete** /dp-block-replication-policies/{policy_id} | 
[**GetDpBlockReplicationPolicy**](DpBlockReplicationPoliciesAPI.md#GetDpBlockReplicationPolicy) | **Get** /dp-block-replication-policies/{policy_id} | 
[**ListDpBlockReplicationPolicies**](DpBlockReplicationPoliciesAPI.md#ListDpBlockReplicationPolicies) | **Get** /dp-block-replication-policies/ | 
[**UpdateDpBlockReplicationPolicy**](DpBlockReplicationPoliciesAPI.md#UpdateDpBlockReplicationPolicy) | **Patch** /dp-block-replication-policies/{policy_id} | 



## CreateDpBlockReplicationPolicy

> DpBlockReplicationPolicyResp CreateDpBlockReplicationPolicy(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDpBlockReplicationPolicyCreateReq(*openapiclient.NewDpBlockReplicationPolicyCreateReqPolicy(int64(123), "Name_example", int64(123))) // DpBlockReplicationPolicyCreateReq | policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockReplicationPoliciesAPI.CreateDpBlockReplicationPolicy(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockReplicationPoliciesAPI.CreateDpBlockReplicationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDpBlockReplicationPolicy`: DpBlockReplicationPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockReplicationPoliciesAPI.CreateDpBlockReplicationPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDpBlockReplicationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DpBlockReplicationPolicyCreateReq**](DpBlockReplicationPolicyCreateReq.md) | policy info | 

### Return type

[**DpBlockReplicationPolicyResp**](DpBlockReplicationPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDpBlockReplicationPolicy

> DeleteDpBlockReplicationPolicy(ctx, policyId).Execute()





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
	policyId := int64(789) // int64 | policy id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DpBlockReplicationPoliciesAPI.DeleteDpBlockReplicationPolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockReplicationPoliciesAPI.DeleteDpBlockReplicationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDpBlockReplicationPolicyRequest struct via the builder pattern


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


## GetDpBlockReplicationPolicy

> DpBlockReplicationPolicyResp GetDpBlockReplicationPolicy(ctx, policyId).Execute()





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
	policyId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockReplicationPoliciesAPI.GetDpBlockReplicationPolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockReplicationPoliciesAPI.GetDpBlockReplicationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDpBlockReplicationPolicy`: DpBlockReplicationPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockReplicationPoliciesAPI.GetDpBlockReplicationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDpBlockReplicationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpBlockReplicationPolicyResp**](DpBlockReplicationPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpBlockReplicationPolicies

> DpBlockReplicationPoliciesResp ListDpBlockReplicationPolicies(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	resp, r, err := apiClient.DpBlockReplicationPoliciesAPI.ListDpBlockReplicationPolicies(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockReplicationPoliciesAPI.ListDpBlockReplicationPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpBlockReplicationPolicies`: DpBlockReplicationPoliciesResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockReplicationPoliciesAPI.ListDpBlockReplicationPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpBlockReplicationPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DpBlockReplicationPoliciesResp**](DpBlockReplicationPoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDpBlockReplicationPolicy

> DpBlockReplicationPolicyResp UpdateDpBlockReplicationPolicy(ctx, policyId).Body(body).Execute()





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
	policyId := int64(789) // int64 | resource id
	body := *openapiclient.NewDpBlockReplicationPolicyUpdateReq(*openapiclient.NewDpBlockReplicationPolicyUpdateReqPolicy()) // DpBlockReplicationPolicyUpdateReq | policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockReplicationPoliciesAPI.UpdateDpBlockReplicationPolicy(context.Background(), policyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockReplicationPoliciesAPI.UpdateDpBlockReplicationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDpBlockReplicationPolicy`: DpBlockReplicationPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockReplicationPoliciesAPI.UpdateDpBlockReplicationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDpBlockReplicationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DpBlockReplicationPolicyUpdateReq**](DpBlockReplicationPolicyUpdateReq.md) | policy info | 

### Return type

[**DpBlockReplicationPolicyResp**](DpBlockReplicationPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

