# \DpBlockAsyncReplicationPoliciesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpBlockAsyncReplicationPolicy**](DpBlockAsyncReplicationPoliciesAPI.md#CreateDpBlockAsyncReplicationPolicy) | **Post** /dp-block-async-replication-policies/ | 
[**DeleteDpBlockAsyncReplicationPolicy**](DpBlockAsyncReplicationPoliciesAPI.md#DeleteDpBlockAsyncReplicationPolicy) | **Delete** /dp-block-async-replication-policies/{policy_id} | 
[**GetDpBlockAsyncReplicationPolicy**](DpBlockAsyncReplicationPoliciesAPI.md#GetDpBlockAsyncReplicationPolicy) | **Get** /dp-block-async-replication-policies/{policy_id} | 
[**ListDpBlockAsyncReplicationPolicies**](DpBlockAsyncReplicationPoliciesAPI.md#ListDpBlockAsyncReplicationPolicies) | **Get** /dp-block-async-replication-policies/ | 
[**UpdateDpBlockAsyncReplicationPolicy**](DpBlockAsyncReplicationPoliciesAPI.md#UpdateDpBlockAsyncReplicationPolicy) | **Patch** /dp-block-async-replication-policies/{policy_id} | 



## CreateDpBlockAsyncReplicationPolicy

> DpBlockAsyncReplicationPolicyResp CreateDpBlockAsyncReplicationPolicy(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDpBlockAsyncReplicationPolicyCreateReq(*openapiclient.NewDpBlockAsyncReplicationPolicyCreateReqPolicy("CronExpr_example", int64(123), int64(123), "Name_example")) // DpBlockAsyncReplicationPolicyCreateReq | policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationPoliciesAPI.CreateDpBlockAsyncReplicationPolicy(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPoliciesAPI.CreateDpBlockAsyncReplicationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDpBlockAsyncReplicationPolicy`: DpBlockAsyncReplicationPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPoliciesAPI.CreateDpBlockAsyncReplicationPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDpBlockAsyncReplicationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DpBlockAsyncReplicationPolicyCreateReq**](DpBlockAsyncReplicationPolicyCreateReq.md) | policy info | 

### Return type

[**DpBlockAsyncReplicationPolicyResp**](DpBlockAsyncReplicationPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDpBlockAsyncReplicationPolicy

> DpBlockAsyncReplicationPolicyResp DeleteDpBlockAsyncReplicationPolicy(ctx, policyId).Force(force).Execute()





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
	resp, r, err := apiClient.DpBlockAsyncReplicationPoliciesAPI.DeleteDpBlockAsyncReplicationPolicy(context.Background(), policyId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPoliciesAPI.DeleteDpBlockAsyncReplicationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDpBlockAsyncReplicationPolicy`: DpBlockAsyncReplicationPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPoliciesAPI.DeleteDpBlockAsyncReplicationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDpBlockAsyncReplicationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 

### Return type

[**DpBlockAsyncReplicationPolicyResp**](DpBlockAsyncReplicationPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDpBlockAsyncReplicationPolicy

> DpBlockAsyncReplicationPolicyResp GetDpBlockAsyncReplicationPolicy(ctx, policyId).Execute()





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
	resp, r, err := apiClient.DpBlockAsyncReplicationPoliciesAPI.GetDpBlockAsyncReplicationPolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPoliciesAPI.GetDpBlockAsyncReplicationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDpBlockAsyncReplicationPolicy`: DpBlockAsyncReplicationPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPoliciesAPI.GetDpBlockAsyncReplicationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDpBlockAsyncReplicationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpBlockAsyncReplicationPolicyResp**](DpBlockAsyncReplicationPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpBlockAsyncReplicationPolicies

> DpBlockAsyncReplicationPoliciesResp ListDpBlockAsyncReplicationPolicies(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).DpSiteId(dpSiteId).VolumeId(volumeId).Execute()





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
	dpSiteId := int64(789) // int64 | related site id (optional)
	volumeId := int64(789) // int64 | related volume id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationPoliciesAPI.ListDpBlockAsyncReplicationPolicies(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).DpSiteId(dpSiteId).VolumeId(volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPoliciesAPI.ListDpBlockAsyncReplicationPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpBlockAsyncReplicationPolicies`: DpBlockAsyncReplicationPoliciesResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPoliciesAPI.ListDpBlockAsyncReplicationPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpBlockAsyncReplicationPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **dpSiteId** | **int64** | related site id | 
 **volumeId** | **int64** | related volume id | 

### Return type

[**DpBlockAsyncReplicationPoliciesResp**](DpBlockAsyncReplicationPoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDpBlockAsyncReplicationPolicy

> DpBlockAsyncReplicationPolicyResp UpdateDpBlockAsyncReplicationPolicy(ctx, policyId).Body(body).Execute()





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
	body := *openapiclient.NewDpBlockAsyncReplicationPolicyUpdateReq(*openapiclient.NewDpBlockAsyncReplicationPolicyUpdateReqPolicy()) // DpBlockAsyncReplicationPolicyUpdateReq | policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationPoliciesAPI.UpdateDpBlockAsyncReplicationPolicy(context.Background(), policyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPoliciesAPI.UpdateDpBlockAsyncReplicationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDpBlockAsyncReplicationPolicy`: DpBlockAsyncReplicationPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPoliciesAPI.UpdateDpBlockAsyncReplicationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDpBlockAsyncReplicationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DpBlockAsyncReplicationPolicyUpdateReq**](DpBlockAsyncReplicationPolicyUpdateReq.md) | policy info | 

### Return type

[**DpBlockAsyncReplicationPolicyResp**](DpBlockAsyncReplicationPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

