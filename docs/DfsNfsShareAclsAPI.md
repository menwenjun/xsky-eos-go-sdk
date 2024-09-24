# \DfsNfsShareAclsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDfsNFSShareACLs**](DfsNfsShareAclsAPI.md#ListDfsNFSShareACLs) | **Get** /dfs-nfs-share-acls/ | 



## ListDfsNFSShareACLs

> DfsNFSShareACLsResp ListDfsNFSShareACLs(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).DfsNfsShareId(dfsNfsShareId).Execute()





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
	dfsNfsShareId := int64(789) // int64 | dfs nfs shares id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsNfsShareAclsAPI.ListDfsNFSShareACLs(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).DfsNfsShareId(dfsNfsShareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsNfsShareAclsAPI.ListDfsNFSShareACLs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsNFSShareACLs`: DfsNFSShareACLsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsNfsShareAclsAPI.ListDfsNFSShareACLs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsNFSShareACLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 
 **dfsNfsShareId** | **int64** | dfs nfs shares id | 

### Return type

[**DfsNFSShareACLsResp**](DfsNFSShareACLsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

