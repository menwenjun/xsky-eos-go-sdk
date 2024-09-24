# \OsSamplesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOSSamples**](OsSamplesAPI.md#GetOSSamples) | **Get** /os-samples/ | 
[**GetOSSamplesByBucketName**](OsSamplesAPI.md#GetOSSamplesByBucketName) | **Get** /os-samples/{user_name}/{bucket_name} | 
[**GetOSSamplesByUserName**](OsSamplesAPI.md#GetOSSamplesByUserName) | **Get** /os-samples/{user_name} | 



## GetOSSamples

> OSSampleResp GetOSSamples(ctx).Time(time).ClusterId(clusterId).Execute()





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
	time := "time_example" // string | query time(url encode RFC3339) (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSamplesAPI.GetOSSamples(context.Background()).Time(time).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSamplesAPI.GetOSSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSSamples`: OSSampleResp
	fmt.Fprintf(os.Stdout, "Response from `OsSamplesAPI.GetOSSamples`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOSSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time** | **string** | query time(url encode RFC3339) | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSSampleResp**](OSSampleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSSamplesByBucketName

> OSSampleResp GetOSSamplesByBucketName(ctx, userName, bucketName).Time(time).ClusterId(clusterId).Execute()





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
	userName := "userName_example" // string | os user name (default to "true")
	bucketName := "bucketName_example" // string | os bucket name (default to "true")
	time := "time_example" // string | query time(url encode RFC3339) (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSamplesAPI.GetOSSamplesByBucketName(context.Background(), userName, bucketName).Time(time).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSamplesAPI.GetOSSamplesByBucketName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSSamplesByBucketName`: OSSampleResp
	fmt.Fprintf(os.Stdout, "Response from `OsSamplesAPI.GetOSSamplesByBucketName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userName** | **string** | os user name | [default to &quot;true&quot;]
**bucketName** | **string** | os bucket name | [default to &quot;true&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSSamplesByBucketNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **time** | **string** | query time(url encode RFC3339) | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSSampleResp**](OSSampleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSSamplesByUserName

> OSSampleResp GetOSSamplesByUserName(ctx, userName).Time(time).ClusterId(clusterId).Execute()





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
	userName := "userName_example" // string | os user name
	time := "time_example" // string | query time(url encode RFC3339) (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSamplesAPI.GetOSSamplesByUserName(context.Background(), userName).Time(time).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSamplesAPI.GetOSSamplesByUserName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSSamplesByUserName`: OSSampleResp
	fmt.Fprintf(os.Stdout, "Response from `OsSamplesAPI.GetOSSamplesByUserName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userName** | **string** | os user name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSSamplesByUserNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | **string** | query time(url encode RFC3339) | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSSampleResp**](OSSampleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

