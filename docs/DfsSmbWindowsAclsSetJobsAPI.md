# \DfsSmbWindowsAclsSetJobsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDfsSMBWindowsACLsSetJob**](DfsSmbWindowsAclsSetJobsAPI.md#DeleteDfsSMBWindowsACLsSetJob) | **Delete** /dfs-smb-windows-acls-set-jobs/{acls_set_job_id} | 
[**ListDfsSmbWindowsACLsSetJobs**](DfsSmbWindowsAclsSetJobsAPI.md#ListDfsSmbWindowsACLsSetJobs) | **Get** /dfs-smb-windows-acls-set-jobs/ | 
[**RetryDfsSMBWindowsACLsSetJob**](DfsSmbWindowsAclsSetJobsAPI.md#RetryDfsSMBWindowsACLsSetJob) | **Post** /dfs-smb-windows-acls-set-jobs/{acls_set_job_id}:retry | 



## DeleteDfsSMBWindowsACLsSetJob

> DfsSMBWindowsACLsSetJobResp DeleteDfsSMBWindowsACLsSetJob(ctx, aclsSetJobId).Execute()





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
	aclsSetJobId := int64(789) // int64 | dfs smb windows acl set job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbWindowsAclsSetJobsAPI.DeleteDfsSMBWindowsACLsSetJob(context.Background(), aclsSetJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbWindowsAclsSetJobsAPI.DeleteDfsSMBWindowsACLsSetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsSMBWindowsACLsSetJob`: DfsSMBWindowsACLsSetJobResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbWindowsAclsSetJobsAPI.DeleteDfsSMBWindowsACLsSetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aclsSetJobId** | **int64** | dfs smb windows acl set job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsSMBWindowsACLsSetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsSMBWindowsACLsSetJobResp**](DfsSMBWindowsACLsSetJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsSmbWindowsACLsSetJobs

> DfsSMBWindowsACLsSetJobsResp ListDfsSmbWindowsACLsSetJobs(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).Status(status).Path(path).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	status := "status_example" // string | status (optional)
	path := "path_example" // string | path (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbWindowsAclsSetJobsAPI.ListDfsSmbWindowsACLsSetJobs(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).Status(status).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbWindowsAclsSetJobsAPI.ListDfsSmbWindowsACLsSetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsSmbWindowsACLsSetJobs`: DfsSMBWindowsACLsSetJobsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbWindowsAclsSetJobsAPI.ListDfsSmbWindowsACLsSetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsSmbWindowsACLsSetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **status** | **string** | status | 
 **path** | **string** | path | 

### Return type

[**DfsSMBWindowsACLsSetJobsResp**](DfsSMBWindowsACLsSetJobsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryDfsSMBWindowsACLsSetJob

> DfsSMBWindowsACLsSetJobResp RetryDfsSMBWindowsACLsSetJob(ctx, aclsSetJobId).Resume(resume).Execute()





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
	aclsSetJobId := int64(789) // int64 | dfs smb windows acls set job id
	resume := true // bool | retry with resume or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbWindowsAclsSetJobsAPI.RetryDfsSMBWindowsACLsSetJob(context.Background(), aclsSetJobId).Resume(resume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbWindowsAclsSetJobsAPI.RetryDfsSMBWindowsACLsSetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryDfsSMBWindowsACLsSetJob`: DfsSMBWindowsACLsSetJobResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbWindowsAclsSetJobsAPI.RetryDfsSMBWindowsACLsSetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aclsSetJobId** | **int64** | dfs smb windows acls set job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryDfsSMBWindowsACLsSetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resume** | **bool** | retry with resume or not | 

### Return type

[**DfsSMBWindowsACLsSetJobResp**](DfsSMBWindowsACLsSetJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

