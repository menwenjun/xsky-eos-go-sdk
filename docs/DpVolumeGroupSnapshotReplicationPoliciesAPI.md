# \DpVolumeGroupSnapshotReplicationPoliciesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpVolumeGroupSnapshotReplicationPolicy**](DpVolumeGroupSnapshotReplicationPoliciesAPI.md#CreateDpVolumeGroupSnapshotReplicationPolicy) | **Post** /dp-volume-group-snapshot-replication-policies/ | 
[**DeleteDpVolumeGroupSnapshotReplicationPolicy**](DpVolumeGroupSnapshotReplicationPoliciesAPI.md#DeleteDpVolumeGroupSnapshotReplicationPolicy) | **Delete** /dp-volume-group-snapshot-replication-policies/{policy_id} | 
[**GetDpVolumeGroupSnapshotReplicationPolicy**](DpVolumeGroupSnapshotReplicationPoliciesAPI.md#GetDpVolumeGroupSnapshotReplicationPolicy) | **Get** /dp-volume-group-snapshot-replication-policies/{policy_id} | 
[**ListDpVolumeGroupSnapshotReplicationPolicies**](DpVolumeGroupSnapshotReplicationPoliciesAPI.md#ListDpVolumeGroupSnapshotReplicationPolicies) | **Get** /dp-volume-group-snapshot-replication-policies/ | 
[**UpdateDpVolumeGroupSnapshotReplicationPolicy**](DpVolumeGroupSnapshotReplicationPoliciesAPI.md#UpdateDpVolumeGroupSnapshotReplicationPolicy) | **Patch** /dp-volume-group-snapshot-replication-policies/{policy_id} | 



## CreateDpVolumeGroupSnapshotReplicationPolicy

> DpVolumeGroupSnapshotReplicationPolicyResp CreateDpVolumeGroupSnapshotReplicationPolicy(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDpVolumeGroupSnapshotReplicationPolicyCreateReq(*openapiclient.NewDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy("CronExpr_example", int64(123), int64(123), "Name_example")) // DpVolumeGroupSnapshotReplicationPolicyCreateReq | policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpVolumeGroupSnapshotReplicationPoliciesAPI.CreateDpVolumeGroupSnapshotReplicationPolicy(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationPoliciesAPI.CreateDpVolumeGroupSnapshotReplicationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDpVolumeGroupSnapshotReplicationPolicy`: DpVolumeGroupSnapshotReplicationPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpVolumeGroupSnapshotReplicationPoliciesAPI.CreateDpVolumeGroupSnapshotReplicationPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDpVolumeGroupSnapshotReplicationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DpVolumeGroupSnapshotReplicationPolicyCreateReq**](DpVolumeGroupSnapshotReplicationPolicyCreateReq.md) | policy info | 

### Return type

[**DpVolumeGroupSnapshotReplicationPolicyResp**](DpVolumeGroupSnapshotReplicationPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDpVolumeGroupSnapshotReplicationPolicy

> DpVolumeGroupSnapshotReplicationPolicyResp DeleteDpVolumeGroupSnapshotReplicationPolicy(ctx, policyId).Force(force).Execute()





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
	resp, r, err := apiClient.DpVolumeGroupSnapshotReplicationPoliciesAPI.DeleteDpVolumeGroupSnapshotReplicationPolicy(context.Background(), policyId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationPoliciesAPI.DeleteDpVolumeGroupSnapshotReplicationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDpVolumeGroupSnapshotReplicationPolicy`: DpVolumeGroupSnapshotReplicationPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpVolumeGroupSnapshotReplicationPoliciesAPI.DeleteDpVolumeGroupSnapshotReplicationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDpVolumeGroupSnapshotReplicationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 

### Return type

[**DpVolumeGroupSnapshotReplicationPolicyResp**](DpVolumeGroupSnapshotReplicationPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDpVolumeGroupSnapshotReplicationPolicy

> DpVolumeGroupSnapshotReplicationPolicyResp GetDpVolumeGroupSnapshotReplicationPolicy(ctx, policyId).Execute()





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
	resp, r, err := apiClient.DpVolumeGroupSnapshotReplicationPoliciesAPI.GetDpVolumeGroupSnapshotReplicationPolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationPoliciesAPI.GetDpVolumeGroupSnapshotReplicationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDpVolumeGroupSnapshotReplicationPolicy`: DpVolumeGroupSnapshotReplicationPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpVolumeGroupSnapshotReplicationPoliciesAPI.GetDpVolumeGroupSnapshotReplicationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDpVolumeGroupSnapshotReplicationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpVolumeGroupSnapshotReplicationPolicyResp**](DpVolumeGroupSnapshotReplicationPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpVolumeGroupSnapshotReplicationPolicies

> DpVolumeGroupSnapshotReplicationPoliciesResp ListDpVolumeGroupSnapshotReplicationPolicies(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).DpSiteId(dpSiteId).VolumeGroupId(volumeGroupId).Execute()





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
	volumeGroupId := int64(789) // int64 | related volume group id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpVolumeGroupSnapshotReplicationPoliciesAPI.ListDpVolumeGroupSnapshotReplicationPolicies(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).DpSiteId(dpSiteId).VolumeGroupId(volumeGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationPoliciesAPI.ListDpVolumeGroupSnapshotReplicationPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpVolumeGroupSnapshotReplicationPolicies`: DpVolumeGroupSnapshotReplicationPoliciesResp
	fmt.Fprintf(os.Stdout, "Response from `DpVolumeGroupSnapshotReplicationPoliciesAPI.ListDpVolumeGroupSnapshotReplicationPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpVolumeGroupSnapshotReplicationPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **dpSiteId** | **int64** | related site id | 
 **volumeGroupId** | **int64** | related volume group id | 

### Return type

[**DpVolumeGroupSnapshotReplicationPoliciesResp**](DpVolumeGroupSnapshotReplicationPoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDpVolumeGroupSnapshotReplicationPolicy

> DpVolumeGroupSnapshotReplicationPolicyResp UpdateDpVolumeGroupSnapshotReplicationPolicy(ctx, policyId).Body(body).Execute()





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
	body := *openapiclient.NewDpVolumeGroupSnapshotReplicationPolicyUpdateReq(*openapiclient.NewDpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy()) // DpVolumeGroupSnapshotReplicationPolicyUpdateReq | policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpVolumeGroupSnapshotReplicationPoliciesAPI.UpdateDpVolumeGroupSnapshotReplicationPolicy(context.Background(), policyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationPoliciesAPI.UpdateDpVolumeGroupSnapshotReplicationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDpVolumeGroupSnapshotReplicationPolicy`: DpVolumeGroupSnapshotReplicationPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpVolumeGroupSnapshotReplicationPoliciesAPI.UpdateDpVolumeGroupSnapshotReplicationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDpVolumeGroupSnapshotReplicationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DpVolumeGroupSnapshotReplicationPolicyUpdateReq**](DpVolumeGroupSnapshotReplicationPolicyUpdateReq.md) | policy info | 

### Return type

[**DpVolumeGroupSnapshotReplicationPolicyResp**](DpVolumeGroupSnapshotReplicationPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

