# \CloudVolumeAttachmentsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCloudVolumeAttachments**](CloudVolumeAttachmentsAPI.md#ListCloudVolumeAttachments) | **Get** /cloud-volume-attachments/ | 



## ListCloudVolumeAttachments

> CloudVolumeAttachmentsResp ListCloudVolumeAttachments(ctx).Limit(limit).Offset(offset).CloudInstanceId(cloudInstanceId).Execute()





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
	cloudInstanceId := int64(789) // int64 | cloud instance id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudVolumeAttachmentsAPI.ListCloudVolumeAttachments(context.Background()).Limit(limit).Offset(offset).CloudInstanceId(cloudInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudVolumeAttachmentsAPI.ListCloudVolumeAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCloudVolumeAttachments`: CloudVolumeAttachmentsResp
	fmt.Fprintf(os.Stdout, "Response from `CloudVolumeAttachmentsAPI.ListCloudVolumeAttachments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCloudVolumeAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **cloudInstanceId** | **int64** | cloud instance id | 

### Return type

[**CloudVolumeAttachmentsResp**](CloudVolumeAttachmentsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

