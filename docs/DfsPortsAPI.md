# \DfsPortsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDfsPorts**](DfsPortsAPI.md#ListDfsPorts) | **Get** /dfs-ports/ | 



## ListDfsPorts

> DfsPortsResp ListDfsPorts(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).DfsGatewayZoneId(dfsGatewayZoneId).Share(share).Execute()





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
	dfsGatewayZoneId := "dfsGatewayZoneId_example" // string | dfs gateway zone id (optional)
	share := "share_example" // string | share (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsPortsAPI.ListDfsPorts(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).DfsGatewayZoneId(dfsGatewayZoneId).Share(share).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsPortsAPI.ListDfsPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsPorts`: DfsPortsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsPortsAPI.ListDfsPorts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **dfsGatewayZoneId** | **string** | dfs gateway zone id | 
 **share** | **string** | share | 

### Return type

[**DfsPortsResp**](DfsPortsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

