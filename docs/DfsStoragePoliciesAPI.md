# \DfsStoragePoliciesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDfsStoragePolicy**](DfsStoragePoliciesAPI.md#CreateDfsStoragePolicy) | **Post** /dfs-storage-policies/ | 
[**DeleteDfsStoragePolicy**](DfsStoragePoliciesAPI.md#DeleteDfsStoragePolicy) | **Delete** /dfs-storage-policies/{dfs_storage_policy_id} | 
[**GetDfsStoragePolicy**](DfsStoragePoliciesAPI.md#GetDfsStoragePolicy) | **Get** /dfs-storage-policies/{dfs_storage_policy_id} | 
[**ListDfsStoragePolicies**](DfsStoragePoliciesAPI.md#ListDfsStoragePolicies) | **Get** /dfs-storage-policies/ | 
[**UnlinkDfsStoragePolicyAndDfsPath**](DfsStoragePoliciesAPI.md#UnlinkDfsStoragePolicyAndDfsPath) | **Patch** /dfs-storage-policies/{dfs_storage_policy_id}:unlink-dfs-path | 
[**UpdateDfsStoragePolicy**](DfsStoragePoliciesAPI.md#UpdateDfsStoragePolicy) | **Patch** /dfs-storage-policies/{dfs_storage_policy_id} | 



## CreateDfsStoragePolicy

> DfsStoragePolicyResp CreateDfsStoragePolicy(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsStoragePolicyCreateReq(*openapiclient.NewDfsStoragePolicyCreateReqDfsStoragePolicy("Name_example", int64(123), int64(123))) // DfsStoragePolicyCreateReq | storagepolicy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsStoragePoliciesAPI.CreateDfsStoragePolicy(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsStoragePoliciesAPI.CreateDfsStoragePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsStoragePolicy`: DfsStoragePolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DfsStoragePoliciesAPI.CreateDfsStoragePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsStoragePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsStoragePolicyCreateReq**](DfsStoragePolicyCreateReq.md) | storagepolicy info | 

### Return type

[**DfsStoragePolicyResp**](DfsStoragePolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsStoragePolicy

> DfsStoragePolicyResp DeleteDfsStoragePolicy(ctx, dfsStoragePolicyId).Execute()





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
	dfsStoragePolicyId := int64(789) // int64 | dfs storage policy id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsStoragePoliciesAPI.DeleteDfsStoragePolicy(context.Background(), dfsStoragePolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsStoragePoliciesAPI.DeleteDfsStoragePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsStoragePolicy`: DfsStoragePolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DfsStoragePoliciesAPI.DeleteDfsStoragePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsStoragePolicyId** | **int64** | dfs storage policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsStoragePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsStoragePolicyResp**](DfsStoragePolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsStoragePolicy

> DfsStoragePolicyResp GetDfsStoragePolicy(ctx, dfsStoragePolicyId).Execute()





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
	dfsStoragePolicyId := int64(789) // int64 | dfs storage policy id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsStoragePoliciesAPI.GetDfsStoragePolicy(context.Background(), dfsStoragePolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsStoragePoliciesAPI.GetDfsStoragePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsStoragePolicy`: DfsStoragePolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DfsStoragePoliciesAPI.GetDfsStoragePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsStoragePolicyId** | **int64** | dfs storage policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsStoragePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsStoragePolicyResp**](DfsStoragePolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsStoragePolicies

> DfsStoragePoliciesResp ListDfsStoragePolicies(ctx).Limit(limit).Offset(offset).DfsRootfsId(dfsRootfsId).DfsPathId(dfsPathId).DfsStorageClassId(dfsStorageClassId).ClusterId(clusterId).Execute()





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
	dfsRootfsId := int64(789) // int64 | dfs rootfs id (optional)
	dfsPathId := int64(789) // int64 | dfs path id (optional)
	dfsStorageClassId := int64(789) // int64 | dfs class id (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsStoragePoliciesAPI.ListDfsStoragePolicies(context.Background()).Limit(limit).Offset(offset).DfsRootfsId(dfsRootfsId).DfsPathId(dfsPathId).DfsStorageClassId(dfsStorageClassId).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsStoragePoliciesAPI.ListDfsStoragePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsStoragePolicies`: DfsStoragePoliciesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsStoragePoliciesAPI.ListDfsStoragePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsStoragePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **dfsRootfsId** | **int64** | dfs rootfs id | 
 **dfsPathId** | **int64** | dfs path id | 
 **dfsStorageClassId** | **int64** | dfs class id | 
 **clusterId** | **string** | cluster id | 

### Return type

[**DfsStoragePoliciesResp**](DfsStoragePoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkDfsStoragePolicyAndDfsPath

> DfsStoragePolicyResp UnlinkDfsStoragePolicyAndDfsPath(ctx, dfsStoragePolicyId).Body(body).Execute()





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
	dfsStoragePolicyId := int64(789) // int64 | dfs storage policy id
	body := *openapiclient.NewDfsStoragePolicyUnlinkDfsPathReq(*openapiclient.NewDfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicy(int64(123), []string{"UnlinkPaths_example"})) // DfsStoragePolicyUnlinkDfsPathReq | policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsStoragePoliciesAPI.UnlinkDfsStoragePolicyAndDfsPath(context.Background(), dfsStoragePolicyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsStoragePoliciesAPI.UnlinkDfsStoragePolicyAndDfsPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnlinkDfsStoragePolicyAndDfsPath`: DfsStoragePolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DfsStoragePoliciesAPI.UnlinkDfsStoragePolicyAndDfsPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsStoragePolicyId** | **int64** | dfs storage policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkDfsStoragePolicyAndDfsPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsStoragePolicyUnlinkDfsPathReq**](DfsStoragePolicyUnlinkDfsPathReq.md) | policy info | 

### Return type

[**DfsStoragePolicyResp**](DfsStoragePolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsStoragePolicy

> DfsStoragePolicyResp UpdateDfsStoragePolicy(ctx, dfsStoragePolicyId).Body(body).Execute()





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
	dfsStoragePolicyId := int64(789) // int64 | dfs storage policy id
	body := *openapiclient.NewDfsStoragePolicyUpdateReq(*openapiclient.NewDfsStoragePolicyUpdateReqDfsStoragePolicy(int64(123), int64(123))) // DfsStoragePolicyUpdateReq | storagepolicy info (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsStoragePoliciesAPI.UpdateDfsStoragePolicy(context.Background(), dfsStoragePolicyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsStoragePoliciesAPI.UpdateDfsStoragePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsStoragePolicy`: DfsStoragePolicyResp
	fmt.Fprintf(os.Stdout, "Response from `DfsStoragePoliciesAPI.UpdateDfsStoragePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsStoragePolicyId** | **int64** | dfs storage policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsStoragePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsStoragePolicyUpdateReq**](DfsStoragePolicyUpdateReq.md) | storagepolicy info | 

### Return type

[**DfsStoragePolicyResp**](DfsStoragePolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

