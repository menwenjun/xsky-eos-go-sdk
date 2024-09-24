# \BlockVolumeMigrationJobsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelBlockVolumeMigrationJob**](BlockVolumeMigrationJobsAPI.md#CancelBlockVolumeMigrationJob) | **Post** /block-volume-migration-jobs/{block_volume_migration_job_id}:cancel | 
[**GetBlockVolumeMigrationJob**](BlockVolumeMigrationJobsAPI.md#GetBlockVolumeMigrationJob) | **Get** /block-volume-migration-jobs/{block_volume_migration_job_id} | 
[**ListBlockVolumeMigrationJobs**](BlockVolumeMigrationJobsAPI.md#ListBlockVolumeMigrationJobs) | **Get** /block-volume-migration-jobs/ | 
[**UpdateMigration**](BlockVolumeMigrationJobsAPI.md#UpdateMigration) | **Post** /block-volume-migration-jobs/{block_volume_migration_job_id}:update | 



## CancelBlockVolumeMigrationJob

> BlockVolumeMigrationJobResp CancelBlockVolumeMigrationJob(ctx, blockVolumeMigrationJobId).Execute()



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
	blockVolumeMigrationJobId := int64(789) // int64 | block volume migration job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeMigrationJobsAPI.CancelBlockVolumeMigrationJob(context.Background(), blockVolumeMigrationJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeMigrationJobsAPI.CancelBlockVolumeMigrationJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelBlockVolumeMigrationJob`: BlockVolumeMigrationJobResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeMigrationJobsAPI.CancelBlockVolumeMigrationJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeMigrationJobId** | **int64** | block volume migration job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelBlockVolumeMigrationJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlockVolumeMigrationJobResp**](BlockVolumeMigrationJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockVolumeMigrationJob

> BlockVolumeMigrationJobResp GetBlockVolumeMigrationJob(ctx, blockVolumeMigrationJobId).Execute()





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
	blockVolumeMigrationJobId := int64(789) // int64 | volume migration job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeMigrationJobsAPI.GetBlockVolumeMigrationJob(context.Background(), blockVolumeMigrationJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeMigrationJobsAPI.GetBlockVolumeMigrationJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockVolumeMigrationJob`: BlockVolumeMigrationJobResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeMigrationJobsAPI.GetBlockVolumeMigrationJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeMigrationJobId** | **int64** | volume migration job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockVolumeMigrationJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlockVolumeMigrationJobResp**](BlockVolumeMigrationJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlockVolumeMigrationJobs

> BlockVolumeMigrationJobsResp ListBlockVolumeMigrationJobs(ctx).Limit(limit).Offset(offset).Status(status).Q(q).Sort(sort).Execute()





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
	status := "status_example" // string | the status of volume migration job (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeMigrationJobsAPI.ListBlockVolumeMigrationJobs(context.Background()).Limit(limit).Offset(offset).Status(status).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeMigrationJobsAPI.ListBlockVolumeMigrationJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlockVolumeMigrationJobs`: BlockVolumeMigrationJobsResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeMigrationJobsAPI.ListBlockVolumeMigrationJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlockVolumeMigrationJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **status** | **string** | the status of volume migration job | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**BlockVolumeMigrationJobsResp**](BlockVolumeMigrationJobsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMigration

> BlockVolumeMigrationJobResp UpdateMigration(ctx, blockVolumeMigrationJobId).Body(body).Execute()



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
	blockVolumeMigrationJobId := int64(789) // int64 | block volume migration job id
	body := *openapiclient.NewBlockVolumeUpdateMigrationReq() // BlockVolumeUpdateMigrationReq | volume migration udpate info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeMigrationJobsAPI.UpdateMigration(context.Background(), blockVolumeMigrationJobId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeMigrationJobsAPI.UpdateMigration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMigration`: BlockVolumeMigrationJobResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeMigrationJobsAPI.UpdateMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeMigrationJobId** | **int64** | block volume migration job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**BlockVolumeUpdateMigrationReq**](BlockVolumeUpdateMigrationReq.md) | volume migration udpate info | 

### Return type

[**BlockVolumeMigrationJobResp**](BlockVolumeMigrationJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

