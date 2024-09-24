# \MetadataServicesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetadataService**](MetadataServicesAPI.md#GetMetadataService) | **Get** /metadata-services/{metadata_service_id} | 
[**GetMetadataServiceSamples**](MetadataServicesAPI.md#GetMetadataServiceSamples) | **Get** /metadata-services/{metadata_service_id}/samples | 
[**ListMetadataServices**](MetadataServicesAPI.md#ListMetadataServices) | **Get** /metadata-services/ | 



## GetMetadataService

> MetadataServiceResp GetMetadataService(ctx, metadataServiceId).Execute()





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
	metadataServiceId := int64(789) // int64 | metadata service id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataServicesAPI.GetMetadataService(context.Background(), metadataServiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataServicesAPI.GetMetadataService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetadataService`: MetadataServiceResp
	fmt.Fprintf(os.Stdout, "Response from `MetadataServicesAPI.GetMetadataService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataServiceId** | **int64** | metadata service id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetadataServiceResp**](MetadataServiceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadataServiceSamples

> MetadataServiceSamplesResp GetMetadataServiceSamples(ctx, metadataServiceId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	metadataServiceId := int64(789) // int64 | metadata service id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataServicesAPI.GetMetadataServiceSamples(context.Background(), metadataServiceId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataServicesAPI.GetMetadataServiceSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetadataServiceSamples`: MetadataServiceSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `MetadataServicesAPI.GetMetadataServiceSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataServiceId** | **int64** | metadata service id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataServiceSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**MetadataServiceSamplesResp**](MetadataServiceSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetadataServices

> MetadataServicesResp ListMetadataServices(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).DiskIds(diskIds).HostId(hostId).MetadataClusterId(metadataClusterId).Q(q).Sort(sort).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	diskIds := int64(789) // int64 | disk ids (optional)
	hostId := int64(789) // int64 | host id (optional)
	metadataClusterId := int64(789) // int64 | metadata cluster id (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataServicesAPI.ListMetadataServices(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).DiskIds(diskIds).HostId(hostId).MetadataClusterId(metadataClusterId).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataServicesAPI.ListMetadataServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetadataServices`: MetadataServicesResp
	fmt.Fprintf(os.Stdout, "Response from `MetadataServicesAPI.ListMetadataServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetadataServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **diskIds** | **int64** | disk ids | 
 **hostId** | **int64** | host id | 
 **metadataClusterId** | **int64** | metadata cluster id | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**MetadataServicesResp**](MetadataServicesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

