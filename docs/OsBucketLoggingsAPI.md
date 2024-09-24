# \OsBucketLoggingsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOSBucketLoggings**](OsBucketLoggingsAPI.md#ListOSBucketLoggings) | **Get** /os-bucket-loggings/ | 



## ListOSBucketLoggings

> OSBucketLoggingsResp ListOSBucketLoggings(ctx).Limit(limit).Offset(offset).SourceBucketName(sourceBucketName).TargetBucketName(targetBucketName).Q(q).Sort(sort).ClusterId(clusterId).Execute()





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
	sourceBucketName := "sourceBucketName_example" // string | source bucket name (optional)
	targetBucketName := "targetBucketName_example" // string | target bucket name (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsBucketLoggingsAPI.ListOSBucketLoggings(context.Background()).Limit(limit).Offset(offset).SourceBucketName(sourceBucketName).TargetBucketName(targetBucketName).Q(q).Sort(sort).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsBucketLoggingsAPI.ListOSBucketLoggings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSBucketLoggings`: OSBucketLoggingsResp
	fmt.Fprintf(os.Stdout, "Response from `OsBucketLoggingsAPI.ListOSBucketLoggings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSBucketLoggingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **sourceBucketName** | **string** | source bucket name | 
 **targetBucketName** | **string** | target bucket name | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSBucketLoggingsResp**](OSBucketLoggingsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

