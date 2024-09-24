# \CloudVolumesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudVolume**](CloudVolumesAPI.md#GetCloudVolume) | **Get** /cloud-volumes/{cloud_volume_id} | 
[**ListCloudVolumes**](CloudVolumesAPI.md#ListCloudVolumes) | **Get** /cloud-volumes/ | 



## GetCloudVolume

> CloudVolumeResp GetCloudVolume(ctx, cloudVolumeId).Execute()





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
	cloudVolumeId := int64(789) // int64 | cloud volume id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudVolumesAPI.GetCloudVolume(context.Background(), cloudVolumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudVolumesAPI.GetCloudVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudVolume`: CloudVolumeResp
	fmt.Fprintf(os.Stdout, "Response from `CloudVolumesAPI.GetCloudVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudVolumeId** | **int64** | cloud volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudVolumeResp**](CloudVolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudVolumes

> CloudVolumesResp ListCloudVolumes(ctx).Limit(limit).Offset(offset).CloudInstanceId(cloudInstanceId).Execute()





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
	cloudInstanceId := int64(789) // int64 | cloud instance id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudVolumesAPI.ListCloudVolumes(context.Background()).Limit(limit).Offset(offset).CloudInstanceId(cloudInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudVolumesAPI.ListCloudVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCloudVolumes`: CloudVolumesResp
	fmt.Fprintf(os.Stdout, "Response from `CloudVolumesAPI.ListCloudVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCloudVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **cloudInstanceId** | **int64** | cloud instance id | 

### Return type

[**CloudVolumesResp**](CloudVolumesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

