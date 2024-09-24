# \DpBlockBackupPoliciesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpBlockBackupPolicy**](DpBlockBackupPoliciesAPI.md#CreateDpBlockBackupPolicy) | **Post** /dp-block-backup-policies/ | 
[**DeleteDpBlockBackupPolicy**](DpBlockBackupPoliciesAPI.md#DeleteDpBlockBackupPolicy) | **Delete** /dp-block-backup-policies/{policy_id} | 
[**GetDpBlockBackupPolicy**](DpBlockBackupPoliciesAPI.md#GetDpBlockBackupPolicy) | **Get** /dp-block-backup-policies/{policy_id} | 
[**ListDpBlockBackupPolicies**](DpBlockBackupPoliciesAPI.md#ListDpBlockBackupPolicies) | **Get** /dp-block-backup-policies/ | 
[**UpdateDpBlockBackupPolicy**](DpBlockBackupPoliciesAPI.md#UpdateDpBlockBackupPolicy) | **Patch** /dp-block-backup-policies/{policy_id} | 



## CreateDpBlockBackupPolicy

> DpBlockBackupPolicyResp CreateDpBlockBackupPolicy(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDpBlockBackupPolicyCreateReq(*openapiclient.NewDpBlockBackupPolicyCreateReqPolicy(int64(123), int64(123), "FullCronExpr_example", "Name_example", int64(123))) // DpBlockBackupPolicyCreateReq | policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockBackupPoliciesAPI.CreateDpBlockBackupPolicy(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockBackupPoliciesAPI.CreateDpBlockBackupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDpBlockBackupPolicy`: DpBlockBackupPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockBackupPoliciesAPI.CreateDpBlockBackupPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDpBlockBackupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DpBlockBackupPolicyCreateReq**](DpBlockBackupPolicyCreateReq.md) | policy info | 

### Return type

[**DpBlockBackupPolicyResp**](DpBlockBackupPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDpBlockBackupPolicy

> DpBlockBackupPolicyResp DeleteDpBlockBackupPolicy(ctx, policyId).Force(force).Execute()





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
	policyId := int64(789) // int64 | resource id
	force := true // bool | force delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockBackupPoliciesAPI.DeleteDpBlockBackupPolicy(context.Background(), policyId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockBackupPoliciesAPI.DeleteDpBlockBackupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDpBlockBackupPolicy`: DpBlockBackupPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockBackupPoliciesAPI.DeleteDpBlockBackupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDpBlockBackupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 

### Return type

[**DpBlockBackupPolicyResp**](DpBlockBackupPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDpBlockBackupPolicy

> DpBlockBackupPolicyResp GetDpBlockBackupPolicy(ctx, policyId).Execute()





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
	policyId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockBackupPoliciesAPI.GetDpBlockBackupPolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockBackupPoliciesAPI.GetDpBlockBackupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDpBlockBackupPolicy`: DpBlockBackupPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockBackupPoliciesAPI.GetDpBlockBackupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDpBlockBackupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpBlockBackupPolicyResp**](DpBlockBackupPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpBlockBackupPolicies

> DpBlockBackupPoliciesResp ListDpBlockBackupPolicies(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).BlockVolumeId(blockVolumeId).DpSiteId(dpSiteId).Execute()





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
	blockVolumeId := int64(789) // int64 | show dp block backup policies of specific block volume (optional)
	dpSiteId := int64(789) // int64 | related site id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockBackupPoliciesAPI.ListDpBlockBackupPolicies(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).BlockVolumeId(blockVolumeId).DpSiteId(dpSiteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockBackupPoliciesAPI.ListDpBlockBackupPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpBlockBackupPolicies`: DpBlockBackupPoliciesResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockBackupPoliciesAPI.ListDpBlockBackupPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpBlockBackupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **blockVolumeId** | **int64** | show dp block backup policies of specific block volume | 
 **dpSiteId** | **int64** | related site id | 

### Return type

[**DpBlockBackupPoliciesResp**](DpBlockBackupPoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDpBlockBackupPolicy

> DpBlockBackupPolicyResp UpdateDpBlockBackupPolicy(ctx, policyId).Body(body).Execute()





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
	policyId := int64(789) // int64 | resource id
	body := *openapiclient.NewDpBlockBackupPolicyUpdateReq(*openapiclient.NewDpBlockBackupPolicyUpdateReqPolicy()) // DpBlockBackupPolicyUpdateReq | policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockBackupPoliciesAPI.UpdateDpBlockBackupPolicy(context.Background(), policyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockBackupPoliciesAPI.UpdateDpBlockBackupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDpBlockBackupPolicy`: DpBlockBackupPolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockBackupPoliciesAPI.UpdateDpBlockBackupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDpBlockBackupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DpBlockBackupPolicyUpdateReq**](DpBlockBackupPolicyUpdateReq.md) | policy info | 

### Return type

[**DpBlockBackupPolicyResp**](DpBlockBackupPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

