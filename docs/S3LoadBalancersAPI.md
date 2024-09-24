# \S3LoadBalancersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchGetS3LoadBalancerSamples**](S3LoadBalancersAPI.md#BatchGetS3LoadBalancerSamples) | **Get** /s3-load-balancers/samples | 
[**GetS3LoadBalancer**](S3LoadBalancersAPI.md#GetS3LoadBalancer) | **Get** /s3-load-balancers/{load_balancer_id} | 
[**GetS3LoadBalancerSamples**](S3LoadBalancersAPI.md#GetS3LoadBalancerSamples) | **Get** /s3-load-balancers/{s3_load_balancer_id}/samples | 
[**ListS3LoadBalancers**](S3LoadBalancersAPI.md#ListS3LoadBalancers) | **Get** /s3-load-balancers/ | 



## BatchGetS3LoadBalancerSamples

> MultiS3LoadBalancersSamplesResp BatchGetS3LoadBalancerSamples(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3LoadBalancersAPI.BatchGetS3LoadBalancerSamples(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3LoadBalancersAPI.BatchGetS3LoadBalancerSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchGetS3LoadBalancerSamples`: MultiS3LoadBalancersSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `S3LoadBalancersAPI.BatchGetS3LoadBalancerSamples`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBatchGetS3LoadBalancerSamplesRequest struct via the builder pattern


### Return type

[**MultiS3LoadBalancersSamplesResp**](MultiS3LoadBalancersSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetS3LoadBalancer

> S3LoadBalancerResp GetS3LoadBalancer(ctx, loadBalancerId).Execute()





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
	loadBalancerId := int64(789) // int64 | s3 load balancer id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3LoadBalancersAPI.GetS3LoadBalancer(context.Background(), loadBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3LoadBalancersAPI.GetS3LoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetS3LoadBalancer`: S3LoadBalancerResp
	fmt.Fprintf(os.Stdout, "Response from `S3LoadBalancersAPI.GetS3LoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **int64** | s3 load balancer id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetS3LoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**S3LoadBalancerResp**](S3LoadBalancerResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetS3LoadBalancerSamples

> S3LoadBalancerSamplesResp GetS3LoadBalancerSamples(ctx, s3LoadBalancerId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	s3LoadBalancerId := int64(789) // int64 | s3 load balancer id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3LoadBalancersAPI.GetS3LoadBalancerSamples(context.Background(), s3LoadBalancerId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3LoadBalancersAPI.GetS3LoadBalancerSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetS3LoadBalancerSamples`: S3LoadBalancerSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `S3LoadBalancersAPI.GetS3LoadBalancerSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**s3LoadBalancerId** | **int64** | s3 load balancer id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetS3LoadBalancerSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**S3LoadBalancerSamplesResp**](S3LoadBalancerSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListS3LoadBalancers

> S3LoadBalancersResp ListS3LoadBalancers(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).GroupId(groupId).OspZoneId(ospZoneId).HostId(hostId).Execute()





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
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)
	groupId := int64(789) // int64 | s3 load balancer group id (optional)
	ospZoneId := int64(789) // int64 | osp zone id (optional)
	hostId := int64(789) // int64 | host id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3LoadBalancersAPI.ListS3LoadBalancers(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).GroupId(groupId).OspZoneId(ospZoneId).HostId(hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3LoadBalancersAPI.ListS3LoadBalancers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListS3LoadBalancers`: S3LoadBalancersResp
	fmt.Fprintf(os.Stdout, "Response from `S3LoadBalancersAPI.ListS3LoadBalancers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListS3LoadBalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 
 **groupId** | **int64** | s3 load balancer group id | 
 **ospZoneId** | **int64** | osp zone id | 
 **hostId** | **int64** | host id | 

### Return type

[**S3LoadBalancersResp**](S3LoadBalancersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

