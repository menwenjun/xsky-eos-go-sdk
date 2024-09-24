# \DpDfsSnapshotPoliciesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpDfsSnapshotPolicy**](DpDfsSnapshotPoliciesAPI.md#CreateDpDfsSnapshotPolicy) | **Post** /dp-dfs-snapshot-policies/ | 
[**DeleteDPDfsSnapshotPolicy**](DpDfsSnapshotPoliciesAPI.md#DeleteDPDfsSnapshotPolicy) | **Delete** /dp-dfs-snapshot-policies/{policy_id} | 
[**GetDpDfsSnapshotPolicy**](DpDfsSnapshotPoliciesAPI.md#GetDpDfsSnapshotPolicy) | **Get** /dp-dfs-snapshot-policies/{policy_id} | 
[**ListDpDfsSnapshotPolicies**](DpDfsSnapshotPoliciesAPI.md#ListDpDfsSnapshotPolicies) | **Get** /dp-dfs-snapshot-policies/ | 
[**UpdateDpDfsSnapshotPolicy**](DpDfsSnapshotPoliciesAPI.md#UpdateDpDfsSnapshotPolicy) | **Patch** /dp-dfs-snapshot-policies/{policy_id} | 



## CreateDpDfsSnapshotPolicy

> DpDfsSnapshotPolicyResp CreateDpDfsSnapshotPolicy(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDpDfsSnapshotPolicyCreateReq(*openapiclient.NewDpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy("CronExpr_example", "Name_example", "RetainTime_example")) // DpDfsSnapshotPolicyCreateReq | dp dfs snapshot policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpDfsSnapshotPoliciesAPI.CreateDpDfsSnapshotPolicy(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpDfsSnapshotPoliciesAPI.CreateDpDfsSnapshotPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDpDfsSnapshotPolicy`: DpDfsSnapshotPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpDfsSnapshotPoliciesAPI.CreateDpDfsSnapshotPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDpDfsSnapshotPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DpDfsSnapshotPolicyCreateReq**](DpDfsSnapshotPolicyCreateReq.md) | dp dfs snapshot policy info | 

### Return type

[**DpDfsSnapshotPolicyResp**](DpDfsSnapshotPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDPDfsSnapshotPolicy

> DeleteDPDfsSnapshotPolicy(ctx, policyId).Force(force).Execute()





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
	policyId := int64(789) // int64 | dp dfs snapshot policy id
	force := true // bool | force delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DpDfsSnapshotPoliciesAPI.DeleteDPDfsSnapshotPolicy(context.Background(), policyId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpDfsSnapshotPoliciesAPI.DeleteDPDfsSnapshotPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | dp dfs snapshot policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDPDfsSnapshotPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 

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


## GetDpDfsSnapshotPolicy

> DpDfsSnapshotPolicyResp GetDpDfsSnapshotPolicy(ctx, policyId).Execute()





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
	policyId := int64(789) // int64 | the dp dfs snapshot policy id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpDfsSnapshotPoliciesAPI.GetDpDfsSnapshotPolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpDfsSnapshotPoliciesAPI.GetDpDfsSnapshotPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDpDfsSnapshotPolicy`: DpDfsSnapshotPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpDfsSnapshotPoliciesAPI.GetDpDfsSnapshotPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | the dp dfs snapshot policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDpDfsSnapshotPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpDfsSnapshotPolicyResp**](DpDfsSnapshotPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpDfsSnapshotPolicies

> DpDfsSnapshotPoliciesResp ListDpDfsSnapshotPolicies(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).DfsPathName(dfsPathName).Execute()





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
	dfsPathName := "dfsPathName_example" // string | show dp dfs snapshot policies of specific dfs path (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpDfsSnapshotPoliciesAPI.ListDpDfsSnapshotPolicies(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).DfsPathName(dfsPathName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpDfsSnapshotPoliciesAPI.ListDpDfsSnapshotPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpDfsSnapshotPolicies`: DpDfsSnapshotPoliciesResp
	fmt.Fprintf(os.Stdout, "Response from `DpDfsSnapshotPoliciesAPI.ListDpDfsSnapshotPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpDfsSnapshotPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **dfsPathName** | **string** | show dp dfs snapshot policies of specific dfs path | 

### Return type

[**DpDfsSnapshotPoliciesResp**](DpDfsSnapshotPoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDpDfsSnapshotPolicy

> DpDfsSnapshotPolicyResp UpdateDpDfsSnapshotPolicy(ctx, policyId).Body(body).Execute()





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
	policyId := int64(789) // int64 | dp dfs snapshot policy id
	body := *openapiclient.NewDpDfsSnapshotPolicyUpdateReq(*openapiclient.NewDpDfsSnapshotPolicyUpdateReqDpDfsSnapshotPolicy()) // DpDfsSnapshotPolicyUpdateReq | dp dfs snapshot policy update info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpDfsSnapshotPoliciesAPI.UpdateDpDfsSnapshotPolicy(context.Background(), policyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpDfsSnapshotPoliciesAPI.UpdateDpDfsSnapshotPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDpDfsSnapshotPolicy`: DpDfsSnapshotPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpDfsSnapshotPoliciesAPI.UpdateDpDfsSnapshotPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | dp dfs snapshot policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDpDfsSnapshotPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DpDfsSnapshotPolicyUpdateReq**](DpDfsSnapshotPolicyUpdateReq.md) | dp dfs snapshot policy update info | 

### Return type

[**DpDfsSnapshotPolicyResp**](DpDfsSnapshotPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

