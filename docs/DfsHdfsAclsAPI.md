# \DfsHdfsAclsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDfsHdfsACLs**](DfsHdfsAclsAPI.md#ListDfsHdfsACLs) | **Get** /dfs-hdfs-acls/ | 



## ListDfsHdfsACLs

> DfsHdfsACLsResp ListDfsHdfsACLs(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).DfsHdfsId(dfsHdfsId).Q(q).Sort(sort).Execute()





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
	dfsHdfsId := int64(789) // int64 | dfs hdfs id (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsAclsAPI.ListDfsHdfsACLs(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).DfsHdfsId(dfsHdfsId).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsAclsAPI.ListDfsHdfsACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsHdfsACLs`: DfsHdfsACLsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsAclsAPI.ListDfsHdfsACLs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsHdfsACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **dfsHdfsId** | **int64** | dfs hdfs id | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DfsHdfsACLsResp**](DfsHdfsACLsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

