# \DpVolumeGroupSnapshotReplicationJobsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDpVolumeGroupSnapshotReplicationJob**](DpVolumeGroupSnapshotReplicationJobsAPI.md#DeleteDpVolumeGroupSnapshotReplicationJob) | **Delete** /dp-volume-group-snapshot-replication-jobs/{job_id} | 
[**GetDpVolumeGroupSnapshotReplicationJob**](DpVolumeGroupSnapshotReplicationJobsAPI.md#GetDpVolumeGroupSnapshotReplicationJob) | **Get** /dp-volume-group-snapshot-replication-jobs/{job_id} | 
[**ListDpVolumeGroupSnapshotReplicationJob**](DpVolumeGroupSnapshotReplicationJobsAPI.md#ListDpVolumeGroupSnapshotReplicationJob) | **Get** /dp-volume-group-snapshot-replication-jobs/ | 



## DeleteDpVolumeGroupSnapshotReplicationJob

> DeleteDpVolumeGroupSnapshotReplicationJob(ctx, jobId).Execute()





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
	jobId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DpVolumeGroupSnapshotReplicationJobsAPI.DeleteDpVolumeGroupSnapshotReplicationJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationJobsAPI.DeleteDpVolumeGroupSnapshotReplicationJob``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDpVolumeGroupSnapshotReplicationJobRequest struct via the builder pattern


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


## GetDpVolumeGroupSnapshotReplicationJob

> DpVolumeGroupSnapshotReplicationJobResp GetDpVolumeGroupSnapshotReplicationJob(ctx, jobId).Execute()





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
	jobId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpVolumeGroupSnapshotReplicationJobsAPI.GetDpVolumeGroupSnapshotReplicationJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationJobsAPI.GetDpVolumeGroupSnapshotReplicationJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDpVolumeGroupSnapshotReplicationJob`: DpVolumeGroupSnapshotReplicationJobResp
	fmt.Fprintf(os.Stdout, "Response from `DpVolumeGroupSnapshotReplicationJobsAPI.GetDpVolumeGroupSnapshotReplicationJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDpVolumeGroupSnapshotReplicationJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpVolumeGroupSnapshotReplicationJobResp**](DpVolumeGroupSnapshotReplicationJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpVolumeGroupSnapshotReplicationJob

> DpVolumeGroupSnapshotReplicationJobsResp ListDpVolumeGroupSnapshotReplicationJob(ctx).DpVolumeGroupSnapshotReplicationPairId(dpVolumeGroupSnapshotReplicationPairId).Execute()





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
	dpVolumeGroupSnapshotReplicationPairId := int64(789) // int64 | related pair id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpVolumeGroupSnapshotReplicationJobsAPI.ListDpVolumeGroupSnapshotReplicationJob(context.Background()).DpVolumeGroupSnapshotReplicationPairId(dpVolumeGroupSnapshotReplicationPairId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationJobsAPI.ListDpVolumeGroupSnapshotReplicationJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpVolumeGroupSnapshotReplicationJob`: DpVolumeGroupSnapshotReplicationJobsResp
	fmt.Fprintf(os.Stdout, "Response from `DpVolumeGroupSnapshotReplicationJobsAPI.ListDpVolumeGroupSnapshotReplicationJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpVolumeGroupSnapshotReplicationJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dpVolumeGroupSnapshotReplicationPairId** | **int64** | related pair id | 

### Return type

[**DpVolumeGroupSnapshotReplicationJobsResp**](DpVolumeGroupSnapshotReplicationJobsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

