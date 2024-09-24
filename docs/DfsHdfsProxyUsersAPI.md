# \DfsHdfsProxyUsersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDfsHdfsProxyUsers**](DfsHdfsProxyUsersAPI.md#ListDfsHdfsProxyUsers) | **Get** /dfs-hdfs-proxy-users/ | 



## ListDfsHdfsProxyUsers

> DfsHdfsProxyUsersResp ListDfsHdfsProxyUsers(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).DfsHdfsId(dfsHdfsId).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	dfsHdfsId := int64(789) // int64 | dfs hdfs id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsProxyUsersAPI.ListDfsHdfsProxyUsers(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).DfsHdfsId(dfsHdfsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsProxyUsersAPI.ListDfsHdfsProxyUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsHdfsProxyUsers`: DfsHdfsProxyUsersResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsProxyUsersAPI.ListDfsHdfsProxyUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsHdfsProxyUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 
 **dfsHdfsId** | **int64** | dfs hdfs id | 

### Return type

[**DfsHdfsProxyUsersResp**](DfsHdfsProxyUsersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

