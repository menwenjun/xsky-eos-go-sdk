# \DfsS3BucketConnectionsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDfsS3BucketConnections**](DfsS3BucketConnectionsAPI.md#ListDfsS3BucketConnections) | **Get** /dfs-s3-bucket-connections/ | 



## ListDfsS3BucketConnections

> DfsS3BucketConnectionsResp ListDfsS3BucketConnections(ctx).Limit(limit).Offset(offset).BucketId(bucketId).Execute()





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
	bucketId := int64(789) // int64 | bucket id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsS3BucketConnectionsAPI.ListDfsS3BucketConnections(context.Background()).Limit(limit).Offset(offset).BucketId(bucketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsS3BucketConnectionsAPI.ListDfsS3BucketConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsS3BucketConnections`: DfsS3BucketConnectionsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsS3BucketConnectionsAPI.ListDfsS3BucketConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsS3BucketConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **bucketId** | **int64** | bucket id | 

### Return type

[**DfsS3BucketConnectionsResp**](DfsS3BucketConnectionsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

