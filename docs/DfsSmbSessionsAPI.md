# \DfsSmbSessionsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDfsSMBSessions**](DfsSmbSessionsAPI.md#ListDfsSMBSessions) | **Get** /dfs-smb-sessions/ | 



## ListDfsSMBSessions

> DfsSMBSessionsResp ListDfsSMBSessions(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).DfsSmbShareId(dfsSmbShareId).Q(q).Sort(sort).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	dfsSmbShareId := int64(789) // int64 | dfs smb share id (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsSmbSessionsAPI.ListDfsSMBSessions(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).DfsSmbShareId(dfsSmbShareId).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsSmbSessionsAPI.ListDfsSMBSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsSMBSessions`: DfsSMBSessionsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsSmbSessionsAPI.ListDfsSMBSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsSMBSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **dfsSmbShareId** | **int64** | dfs smb share id | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DfsSMBSessionsResp**](DfsSMBSessionsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

