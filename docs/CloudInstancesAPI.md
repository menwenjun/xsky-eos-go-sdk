# \CloudInstancesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudInstance**](CloudInstancesAPI.md#GetCloudInstance) | **Get** /cloud-instances/{cloud_instance_id} | 
[**GetCloudInstanceSamples**](CloudInstancesAPI.md#GetCloudInstanceSamples) | **Get** /cloud-instances/{cloud_instance_id}/samples | 
[**ListCloudInstances**](CloudInstancesAPI.md#ListCloudInstances) | **Get** /cloud-instances/ | 



## GetCloudInstance

> CloudInstanceResp GetCloudInstance(ctx, cloudInstanceId).Execute()





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
	cloudInstanceId := int64(789) // int64 | cloud instance id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudInstancesAPI.GetCloudInstance(context.Background(), cloudInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInstancesAPI.GetCloudInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudInstance`: CloudInstanceResp
	fmt.Fprintf(os.Stdout, "Response from `CloudInstancesAPI.GetCloudInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudInstanceId** | **int64** | cloud instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudInstanceResp**](CloudInstanceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudInstanceSamples

> CloudInstanceSamplesResp GetCloudInstanceSamples(ctx, cloudInstanceId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	cloudInstanceId := int64(789) // int64 | cloud instance id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudInstancesAPI.GetCloudInstanceSamples(context.Background(), cloudInstanceId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInstancesAPI.GetCloudInstanceSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudInstanceSamples`: CloudInstanceSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `CloudInstancesAPI.GetCloudInstanceSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudInstanceId** | **int64** | cloud instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudInstanceSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**CloudInstanceSamplesResp**](CloudInstanceSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudInstances

> CloudInstancesResp ListCloudInstances(ctx).Limit(limit).Offset(offset).CloudPlatformId(cloudPlatformId).CloudPlatformType(cloudPlatformType).CloudPlatformName(cloudPlatformName).ClusterId(clusterId).Execute()





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
	cloudPlatformId := int64(789) // int64 | cloud platform id (optional)
	cloudPlatformType := "cloudPlatformType_example" // string | cloud platform type (optional)
	cloudPlatformName := "cloudPlatformName_example" // string | cloud platform name (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudInstancesAPI.ListCloudInstances(context.Background()).Limit(limit).Offset(offset).CloudPlatformId(cloudPlatformId).CloudPlatformType(cloudPlatformType).CloudPlatformName(cloudPlatformName).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudInstancesAPI.ListCloudInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCloudInstances`: CloudInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `CloudInstancesAPI.ListCloudInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCloudInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **cloudPlatformId** | **int64** | cloud platform id | 
 **cloudPlatformType** | **string** | cloud platform type | 
 **cloudPlatformName** | **string** | cloud platform name | 
 **clusterId** | **string** | cluster id | 

### Return type

[**CloudInstancesResp**](CloudInstancesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

