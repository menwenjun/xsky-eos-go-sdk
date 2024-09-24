# \DpBlockSnapshotRecoveryJobsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpBlockSnapshotRecoveryJob**](DpBlockSnapshotRecoveryJobsAPI.md#CreateDpBlockSnapshotRecoveryJob) | **Post** /dp-block-snapshot-recovery-jobs/ | 
[**DeleteDpBlockSnapshotRecoveryJob**](DpBlockSnapshotRecoveryJobsAPI.md#DeleteDpBlockSnapshotRecoveryJob) | **Delete** /dp-block-snapshot-recovery-jobs/{job_id} | 
[**GetDpBlockSnapshotRecoveryJob**](DpBlockSnapshotRecoveryJobsAPI.md#GetDpBlockSnapshotRecoveryJob) | **Get** /dp-block-snapshot-recovery-jobs/{job_id} | 
[**ListDpBlockSnapshotRecoveryJobs**](DpBlockSnapshotRecoveryJobsAPI.md#ListDpBlockSnapshotRecoveryJobs) | **Get** /dp-block-snapshot-recovery-jobs/ | 



## CreateDpBlockSnapshotRecoveryJob

> DpBlockSnapshotRecoveryJobResp CreateDpBlockSnapshotRecoveryJob(ctx).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := *openapiclient.NewDpBlockSnapshotRecoveryJobCreateReq(*openapiclient.NewDpBlockSnapshotRecoveryJobCreateReqJob(*openapiclient.NewDpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot(time.Now(), int64(123)), *openapiclient.NewDpBlockSnapshotRecoveryJobCreateReqJobBackupVolume("Name_example", "VolumeName_example"), *openapiclient.NewDpBlockSnapshotRecoveryJobCreateReqJobBackupCluster("FsId_example", "Name_example"), *openapiclient.NewDpBlockSnapshotRecoveryJobCreateReqJobVolume("Name_example", int64(123)), int64(123), int64(123), "ResourceType_example")) // DpBlockSnapshotRecoveryJobCreateReq | resource info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockSnapshotRecoveryJobsAPI.CreateDpBlockSnapshotRecoveryJob(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockSnapshotRecoveryJobsAPI.CreateDpBlockSnapshotRecoveryJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDpBlockSnapshotRecoveryJob`: DpBlockSnapshotRecoveryJobResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockSnapshotRecoveryJobsAPI.CreateDpBlockSnapshotRecoveryJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDpBlockSnapshotRecoveryJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DpBlockSnapshotRecoveryJobCreateReq**](DpBlockSnapshotRecoveryJobCreateReq.md) | resource info | 

### Return type

[**DpBlockSnapshotRecoveryJobResp**](DpBlockSnapshotRecoveryJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDpBlockSnapshotRecoveryJob

> DeleteDpBlockSnapshotRecoveryJob(ctx, jobId).Execute()





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
	jobId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DpBlockSnapshotRecoveryJobsAPI.DeleteDpBlockSnapshotRecoveryJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockSnapshotRecoveryJobsAPI.DeleteDpBlockSnapshotRecoveryJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDpBlockSnapshotRecoveryJobRequest struct via the builder pattern


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


## GetDpBlockSnapshotRecoveryJob

> DpBlockSnapshotRecoveryJobResp GetDpBlockSnapshotRecoveryJob(ctx, jobId).Execute()





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
	jobId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockSnapshotRecoveryJobsAPI.GetDpBlockSnapshotRecoveryJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockSnapshotRecoveryJobsAPI.GetDpBlockSnapshotRecoveryJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDpBlockSnapshotRecoveryJob`: DpBlockSnapshotRecoveryJobResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockSnapshotRecoveryJobsAPI.GetDpBlockSnapshotRecoveryJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDpBlockSnapshotRecoveryJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpBlockSnapshotRecoveryJobResp**](DpBlockSnapshotRecoveryJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpBlockSnapshotRecoveryJobs

> DpBlockSnapshotRecoveryJobsResp ListDpBlockSnapshotRecoveryJobs(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	resp, r, err := apiClient.DpBlockSnapshotRecoveryJobsAPI.ListDpBlockSnapshotRecoveryJobs(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockSnapshotRecoveryJobsAPI.ListDpBlockSnapshotRecoveryJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpBlockSnapshotRecoveryJobs`: DpBlockSnapshotRecoveryJobsResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockSnapshotRecoveryJobsAPI.ListDpBlockSnapshotRecoveryJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpBlockSnapshotRecoveryJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DpBlockSnapshotRecoveryJobsResp**](DpBlockSnapshotRecoveryJobsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

