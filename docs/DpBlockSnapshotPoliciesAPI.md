# \DpBlockSnapshotPoliciesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpBlockSnapshotPolicy**](DpBlockSnapshotPoliciesAPI.md#CreateDpBlockSnapshotPolicy) | **Post** /dp-block-snapshot-policies/ | 
[**DeleteDpBlockSnapshotPolicy**](DpBlockSnapshotPoliciesAPI.md#DeleteDpBlockSnapshotPolicy) | **Delete** /dp-block-snapshot-policies/{policy_id} | 
[**GetDpBlockSnapshotPolicy**](DpBlockSnapshotPoliciesAPI.md#GetDpBlockSnapshotPolicy) | **Get** /dp-block-snapshot-policies/{policy_id} | 
[**ListDpBlockSnapshotPolicies**](DpBlockSnapshotPoliciesAPI.md#ListDpBlockSnapshotPolicies) | **Get** /dp-block-snapshot-policies/ | 
[**UpdateDpBlockSnapshotPolicy**](DpBlockSnapshotPoliciesAPI.md#UpdateDpBlockSnapshotPolicy) | **Patch** /dp-block-snapshot-policies/{policy_id} | 



## CreateDpBlockSnapshotPolicy

> DpBlockSnapshotPolicyResp CreateDpBlockSnapshotPolicy(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDpBlockSnapshotPolicyCreateReq(*openapiclient.NewDpBlockSnapshotPolicyCreateReqPolicy("CronExpr_example", int64(123), "Name_example", int64(123))) // DpBlockSnapshotPolicyCreateReq | policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockSnapshotPoliciesAPI.CreateDpBlockSnapshotPolicy(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockSnapshotPoliciesAPI.CreateDpBlockSnapshotPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDpBlockSnapshotPolicy`: DpBlockSnapshotPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockSnapshotPoliciesAPI.CreateDpBlockSnapshotPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDpBlockSnapshotPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DpBlockSnapshotPolicyCreateReq**](DpBlockSnapshotPolicyCreateReq.md) | policy info | 

### Return type

[**DpBlockSnapshotPolicyResp**](DpBlockSnapshotPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDpBlockSnapshotPolicy

> DpBlockSnapshotPolicyResp DeleteDpBlockSnapshotPolicy(ctx, policyId).Force(force).Execute()





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
	force := true // bool | force delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockSnapshotPoliciesAPI.DeleteDpBlockSnapshotPolicy(context.Background(), policyId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockSnapshotPoliciesAPI.DeleteDpBlockSnapshotPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDpBlockSnapshotPolicy`: DpBlockSnapshotPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockSnapshotPoliciesAPI.DeleteDpBlockSnapshotPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDpBlockSnapshotPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 

### Return type

[**DpBlockSnapshotPolicyResp**](DpBlockSnapshotPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDpBlockSnapshotPolicy

> DpBlockSnapshotPolicyResp GetDpBlockSnapshotPolicy(ctx, policyId).Execute()





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
	resp, r, err := apiClient.DpBlockSnapshotPoliciesAPI.GetDpBlockSnapshotPolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockSnapshotPoliciesAPI.GetDpBlockSnapshotPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDpBlockSnapshotPolicy`: DpBlockSnapshotPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockSnapshotPoliciesAPI.GetDpBlockSnapshotPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDpBlockSnapshotPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpBlockSnapshotPolicyResp**](DpBlockSnapshotPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpBlockSnapshotPolicies

> DpBlockSnapshotPoliciesResp ListDpBlockSnapshotPolicies(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	resp, r, err := apiClient.DpBlockSnapshotPoliciesAPI.ListDpBlockSnapshotPolicies(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockSnapshotPoliciesAPI.ListDpBlockSnapshotPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpBlockSnapshotPolicies`: DpBlockSnapshotPoliciesResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockSnapshotPoliciesAPI.ListDpBlockSnapshotPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpBlockSnapshotPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DpBlockSnapshotPoliciesResp**](DpBlockSnapshotPoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDpBlockSnapshotPolicy

> DpBlockSnapshotPolicyResp UpdateDpBlockSnapshotPolicy(ctx, policyId).Body(body).Execute()





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
	body := *openapiclient.NewDpBlockSnapshotPolicyUpdateReq(*openapiclient.NewDpBlockSnapshotPolicyUpdateReqPolicy()) // DpBlockSnapshotPolicyUpdateReq | policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockSnapshotPoliciesAPI.UpdateDpBlockSnapshotPolicy(context.Background(), policyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockSnapshotPoliciesAPI.UpdateDpBlockSnapshotPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDpBlockSnapshotPolicy`: DpBlockSnapshotPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockSnapshotPoliciesAPI.UpdateDpBlockSnapshotPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDpBlockSnapshotPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DpBlockSnapshotPolicyUpdateReq**](DpBlockSnapshotPolicyUpdateReq.md) | policy info | 

### Return type

[**DpBlockSnapshotPolicyResp**](DpBlockSnapshotPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

