# \DpBlockAsyncReplicationJobsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDpBlockAsyncReplicationJob**](DpBlockAsyncReplicationJobsAPI.md#DeleteDpBlockAsyncReplicationJob) | **Delete** /dp-block-async-replication-jobs/{job_id} | 
[**GetDpBlockAsyncReplicationJob**](DpBlockAsyncReplicationJobsAPI.md#GetDpBlockAsyncReplicationJob) | **Get** /dp-block-async-replication-jobs/{job_id} | 
[**ListDpBlockAsyncReplicationJob**](DpBlockAsyncReplicationJobsAPI.md#ListDpBlockAsyncReplicationJob) | **Get** /dp-block-async-replication-jobs/ | 



## DeleteDpBlockAsyncReplicationJob

> DeleteDpBlockAsyncReplicationJob(ctx, jobId).Execute()





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
	r, err := apiClient.DpBlockAsyncReplicationJobsAPI.DeleteDpBlockAsyncReplicationJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationJobsAPI.DeleteDpBlockAsyncReplicationJob``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDpBlockAsyncReplicationJobRequest struct via the builder pattern


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


## GetDpBlockAsyncReplicationJob

> DpBlockAsyncReplicationJobResp GetDpBlockAsyncReplicationJob(ctx, jobId).Execute()





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
	resp, r, err := apiClient.DpBlockAsyncReplicationJobsAPI.GetDpBlockAsyncReplicationJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationJobsAPI.GetDpBlockAsyncReplicationJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDpBlockAsyncReplicationJob`: DpBlockAsyncReplicationJobResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationJobsAPI.GetDpBlockAsyncReplicationJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDpBlockAsyncReplicationJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpBlockAsyncReplicationJobResp**](DpBlockAsyncReplicationJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpBlockAsyncReplicationJob

> DpBlockAsyncReplicationJobsResp ListDpBlockAsyncReplicationJob(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationJobsAPI.ListDpBlockAsyncReplicationJob(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationJobsAPI.ListDpBlockAsyncReplicationJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpBlockAsyncReplicationJob`: DpBlockAsyncReplicationJobsResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationJobsAPI.ListDpBlockAsyncReplicationJob`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDpBlockAsyncReplicationJobRequest struct via the builder pattern


### Return type

[**DpBlockAsyncReplicationJobsResp**](DpBlockAsyncReplicationJobsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

