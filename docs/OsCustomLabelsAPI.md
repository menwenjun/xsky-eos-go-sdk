# \OsCustomLabelsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOSCustomLabel**](OsCustomLabelsAPI.md#GetOSCustomLabel) | **Get** /os-custom-labels/{os_custom_label_id} | 
[**ListOSCustomLabels**](OsCustomLabelsAPI.md#ListOSCustomLabels) | **Get** /os-custom-labels/ | 



## GetOSCustomLabel

> OSCustomLabelResp GetOSCustomLabel(ctx, osCustomLabelId).Execute()





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
	osCustomLabelId := int64(789) // int64 | os custom label id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsCustomLabelsAPI.GetOSCustomLabel(context.Background(), osCustomLabelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsCustomLabelsAPI.GetOSCustomLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSCustomLabel`: OSCustomLabelResp
	fmt.Fprintf(os.Stdout, "Response from `OsCustomLabelsAPI.GetOSCustomLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osCustomLabelId** | **int64** | os custom label id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSCustomLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSCustomLabelResp**](OSCustomLabelResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOSCustomLabels

> OSCustomLabelsResp ListOSCustomLabels(ctx).BucketId(bucketId).Category(category).ClusterId(clusterId).Limit(limit).Offset(offset).Execute()





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
	bucketId := int64(789) // int64 | bucket id (optional)
	category := "category_example" // string | label category: meta, default, tagging (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsCustomLabelsAPI.ListOSCustomLabels(context.Background()).BucketId(bucketId).Category(category).ClusterId(clusterId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsCustomLabelsAPI.ListOSCustomLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSCustomLabels`: OSCustomLabelsResp
	fmt.Fprintf(os.Stdout, "Response from `OsCustomLabelsAPI.ListOSCustomLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSCustomLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketId** | **int64** | bucket id | 
 **category** | **string** | label category: meta, default, tagging | 
 **clusterId** | **string** | cluster id | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**OSCustomLabelsResp**](OSCustomLabelsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

