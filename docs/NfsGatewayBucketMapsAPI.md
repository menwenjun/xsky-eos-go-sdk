# \NfsGatewayBucketMapsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNFSGatewayS3BucketMaps**](NfsGatewayBucketMapsAPI.md#ListNFSGatewayS3BucketMaps) | **Get** /nfs-gateway-bucket-maps/ | 



## ListNFSGatewayS3BucketMaps

> NFSGatewayBucketMapsResp ListNFSGatewayS3BucketMaps(ctx).Limit(limit).Offset(offset).NfsGatewayId(nfsGatewayId).OsBucketId(osBucketId).ClusterId(clusterId).Execute()





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
	nfsGatewayId := int64(789) // int64 | nfs gateway id (optional)
	osBucketId := int64(789) // int64 | s3 bucket id (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsGatewayBucketMapsAPI.ListNFSGatewayS3BucketMaps(context.Background()).Limit(limit).Offset(offset).NfsGatewayId(nfsGatewayId).OsBucketId(osBucketId).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsGatewayBucketMapsAPI.ListNFSGatewayS3BucketMaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNFSGatewayS3BucketMaps`: NFSGatewayBucketMapsResp
	fmt.Fprintf(os.Stdout, "Response from `NfsGatewayBucketMapsAPI.ListNFSGatewayS3BucketMaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNFSGatewayS3BucketMapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **nfsGatewayId** | **int64** | nfs gateway id | 
 **osBucketId** | **int64** | s3 bucket id | 
 **clusterId** | **string** | cluster id | 

### Return type

[**NFSGatewayBucketMapsResp**](NFSGatewayBucketMapsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

