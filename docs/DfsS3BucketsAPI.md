# \DfsS3BucketsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDfsS3Bucket**](DfsS3BucketsAPI.md#CreateDfsS3Bucket) | **Post** /dfs-s3-buckets/ | 
[**DeleteDfsS3Bucket**](DfsS3BucketsAPI.md#DeleteDfsS3Bucket) | **Delete** /dfs-s3-buckets/{bucket_id} | 
[**DeleteDfsS3BucketPolicy**](DfsS3BucketsAPI.md#DeleteDfsS3BucketPolicy) | **Delete** /dfs-s3-buckets/{bucket_id}:delete-bucket-policy | 
[**GetDfsS3Bucket**](DfsS3BucketsAPI.md#GetDfsS3Bucket) | **Get** /dfs-s3-buckets/{bucket_id} | 
[**GetDfsS3BucketSamples**](DfsS3BucketsAPI.md#GetDfsS3BucketSamples) | **Get** /dfs-s3-buckets/{bucket_id}/samples | 
[**ListDfsS3Buckets**](DfsS3BucketsAPI.md#ListDfsS3Buckets) | **Get** /dfs-s3-buckets/ | 
[**SetDfsS3BucketPolicy**](DfsS3BucketsAPI.md#SetDfsS3BucketPolicy) | **Post** /dfs-s3-buckets/{bucket_id}:set-bucket-policy | 
[**UpdateDfsS3Bucket**](DfsS3BucketsAPI.md#UpdateDfsS3Bucket) | **Patch** /dfs-s3-buckets/{bucket_id} | 



## CreateDfsS3Bucket

> DfsS3BucketResp CreateDfsS3Bucket(ctx).Body(body).ClusterId(clusterId).AllowPathCreate(allowPathCreate).Execute()





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
	body := *openapiclient.NewDfsS3BucketCreateReq() // DfsS3BucketCreateReq | bucket info
	clusterId := "clusterId_example" // string | cluster id (optional)
	allowPathCreate := true // bool | allow create path when not existed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsS3BucketsAPI.CreateDfsS3Bucket(context.Background()).Body(body).ClusterId(clusterId).AllowPathCreate(allowPathCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsS3BucketsAPI.CreateDfsS3Bucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsS3Bucket`: DfsS3BucketResp
	fmt.Fprintf(os.Stdout, "Response from `DfsS3BucketsAPI.CreateDfsS3Bucket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsS3BucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsS3BucketCreateReq**](DfsS3BucketCreateReq.md) | bucket info | 
 **clusterId** | **string** | cluster id | 
 **allowPathCreate** | **bool** | allow create path when not existed | 

### Return type

[**DfsS3BucketResp**](DfsS3BucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsS3Bucket

> DfsS3BucketResp DeleteDfsS3Bucket(ctx, bucketId).WithDirectory(withDirectory).Execute()





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
	bucketId := int64(789) // int64 | bucket id
	withDirectory := true // bool | also delete directory (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsS3BucketsAPI.DeleteDfsS3Bucket(context.Background(), bucketId).WithDirectory(withDirectory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsS3BucketsAPI.DeleteDfsS3Bucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsS3Bucket`: DfsS3BucketResp
	fmt.Fprintf(os.Stdout, "Response from `DfsS3BucketsAPI.DeleteDfsS3Bucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsS3BucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDirectory** | **bool** | also delete directory | 

### Return type

[**DfsS3BucketResp**](DfsS3BucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsS3BucketPolicy

> DfsS3BucketResp DeleteDfsS3BucketPolicy(ctx, bucketId).Execute()





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
	bucketId := int64(789) // int64 | bucket id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsS3BucketsAPI.DeleteDfsS3BucketPolicy(context.Background(), bucketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsS3BucketsAPI.DeleteDfsS3BucketPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsS3BucketPolicy`: DfsS3BucketResp
	fmt.Fprintf(os.Stdout, "Response from `DfsS3BucketsAPI.DeleteDfsS3BucketPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsS3BucketPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsS3BucketResp**](DfsS3BucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsS3Bucket

> DfsS3BucketResp GetDfsS3Bucket(ctx, bucketId).Execute()





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
	bucketId := int64(789) // int64 | bucket id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsS3BucketsAPI.GetDfsS3Bucket(context.Background(), bucketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsS3BucketsAPI.GetDfsS3Bucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsS3Bucket`: DfsS3BucketResp
	fmt.Fprintf(os.Stdout, "Response from `DfsS3BucketsAPI.GetDfsS3Bucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsS3BucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsS3BucketResp**](DfsS3BucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsS3BucketSamples

> DfsS3BucketSamplesResp GetDfsS3BucketSamples(ctx, bucketId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	bucketId := int64(789) // int64 | bucket id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsS3BucketsAPI.GetDfsS3BucketSamples(context.Background(), bucketId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsS3BucketsAPI.GetDfsS3BucketSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsS3BucketSamples`: DfsS3BucketSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsS3BucketsAPI.GetDfsS3BucketSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsS3BucketSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**DfsS3BucketSamplesResp**](DfsS3BucketSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsS3Buckets

> DfsS3BucketsResp ListDfsS3Buckets(ctx).ClusterId(clusterId).Limit(limit).Offset(offset).Name(name).OwnerId(ownerId).PathId(pathId).Q(q).Sort(sort).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	name := "name_example" // string | name of dfs s3 buckets (optional)
	ownerId := int64(789) // int64 | The id of user buckets belong to (optional)
	pathId := int64(789) // int64 | The id of bucket path (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsS3BucketsAPI.ListDfsS3Buckets(context.Background()).ClusterId(clusterId).Limit(limit).Offset(offset).Name(name).OwnerId(ownerId).PathId(pathId).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsS3BucketsAPI.ListDfsS3Buckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsS3Buckets`: DfsS3BucketsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsS3BucketsAPI.ListDfsS3Buckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsS3BucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **name** | **string** | name of dfs s3 buckets | 
 **ownerId** | **int64** | The id of user buckets belong to | 
 **pathId** | **int64** | The id of bucket path | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DfsS3BucketsResp**](DfsS3BucketsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDfsS3BucketPolicy

> DfsS3BucketResp SetDfsS3BucketPolicy(ctx, bucketId).Body(body).Execute()





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
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewDfsS3BucketPolicySetReq() // DfsS3BucketPolicySetReq | bucket policy info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsS3BucketsAPI.SetDfsS3BucketPolicy(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsS3BucketsAPI.SetDfsS3BucketPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDfsS3BucketPolicy`: DfsS3BucketResp
	fmt.Fprintf(os.Stdout, "Response from `DfsS3BucketsAPI.SetDfsS3BucketPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDfsS3BucketPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsS3BucketPolicySetReq**](DfsS3BucketPolicySetReq.md) | bucket policy info | 

### Return type

[**DfsS3BucketResp**](DfsS3BucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsS3Bucket

> DfsS3BucketResp UpdateDfsS3Bucket(ctx, bucketId).Body(body).Execute()





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
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewDfsS3BucketUpdateReq() // DfsS3BucketUpdateReq | bucket info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsS3BucketsAPI.UpdateDfsS3Bucket(context.Background(), bucketId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsS3BucketsAPI.UpdateDfsS3Bucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsS3Bucket`: DfsS3BucketResp
	fmt.Fprintf(os.Stdout, "Response from `DfsS3BucketsAPI.UpdateDfsS3Bucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsS3BucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsS3BucketUpdateReq**](DfsS3BucketUpdateReq.md) | bucket info | 

### Return type

[**DfsS3BucketResp**](DfsS3BucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

