# \DfsHdfsConnectionsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDfsHdfsConnections**](DfsHdfsConnectionsAPI.md#ListDfsHdfsConnections) | **Get** /dfs-hdfs-connections/ | 



## ListDfsHdfsConnections

> DfsHdfsConnectionsResp ListDfsHdfsConnections(ctx).Limit(limit).Offset(offset).HdfsId(hdfsId).ClusterId(clusterId).Execute()





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
	hdfsId := int64(789) // int64 | hdfs id (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsHdfsConnectionsAPI.ListDfsHdfsConnections(context.Background()).Limit(limit).Offset(offset).HdfsId(hdfsId).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsHdfsConnectionsAPI.ListDfsHdfsConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsHdfsConnections`: DfsHdfsConnectionsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsHdfsConnectionsAPI.ListDfsHdfsConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsHdfsConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **hdfsId** | **int64** | hdfs id | 
 **clusterId** | **string** | cluster id | 

### Return type

[**DfsHdfsConnectionsResp**](DfsHdfsConnectionsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

