# \DpBlockBackupJobsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDpBlockBackupJobs**](DpBlockBackupJobsAPI.md#ListDpBlockBackupJobs) | **Get** /dp-block-backup-jobs/ | 



## ListDpBlockBackupJobs

> DpBlockBackupJobsResp ListDpBlockBackupJobs(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).BlockVolumeId(blockVolumeId).BlockSnapshotId(blockSnapshotId).DpBlockBackupPolicyId(dpBlockBackupPolicyId).Execute()





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
	blockVolumeId := int64(789) // int64 | block volume (optional)
	blockSnapshotId := int64(789) // int64 | object storage bucket (optional)
	dpBlockBackupPolicyId := int64(789) // int64 | dp block backup policy (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockBackupJobsAPI.ListDpBlockBackupJobs(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).BlockVolumeId(blockVolumeId).BlockSnapshotId(blockSnapshotId).DpBlockBackupPolicyId(dpBlockBackupPolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockBackupJobsAPI.ListDpBlockBackupJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpBlockBackupJobs`: DpBlockBackupJobsResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockBackupJobsAPI.ListDpBlockBackupJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpBlockBackupJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **blockVolumeId** | **int64** | block volume | 
 **blockSnapshotId** | **int64** | object storage bucket | 
 **dpBlockBackupPolicyId** | **int64** | dp block backup policy | 

### Return type

[**DpBlockBackupJobsResp**](DpBlockBackupJobsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

