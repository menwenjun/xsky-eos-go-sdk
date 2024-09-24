# \DpDfsSnapshotJobsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDpDfsSnapshotJobs**](DpDfsSnapshotJobsAPI.md#ListDpDfsSnapshotJobs) | **Get** /dp-dfs-snapshot-jobs/ | 



## ListDpDfsSnapshotJobs

> DpDfsSnapshotJobsResp ListDpDfsSnapshotJobs(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).DpDfsSnapshotPolicyId(dpDfsSnapshotPolicyId).Execute()





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
	dpDfsSnapshotPolicyId := int64(789) // int64 | dp dfs snapshot policy (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpDfsSnapshotJobsAPI.ListDpDfsSnapshotJobs(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).DpDfsSnapshotPolicyId(dpDfsSnapshotPolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpDfsSnapshotJobsAPI.ListDpDfsSnapshotJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpDfsSnapshotJobs`: DpDfsSnapshotJobsResp
	fmt.Fprintf(os.Stdout, "Response from `DpDfsSnapshotJobsAPI.ListDpDfsSnapshotJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpDfsSnapshotJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **dpDfsSnapshotPolicyId** | **int64** | dp dfs snapshot policy | 

### Return type

[**DpDfsSnapshotJobsResp**](DpDfsSnapshotJobsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

