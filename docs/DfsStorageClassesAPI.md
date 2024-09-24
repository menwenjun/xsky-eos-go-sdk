# \DfsStorageClassesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDfsStorageClassPools**](DfsStorageClassesAPI.md#AddDfsStorageClassPools) | **Post** /dfs-storage-classes/{dfs_storage_class_id}:add-pools | 
[**CreateDfsStorageClass**](DfsStorageClassesAPI.md#CreateDfsStorageClass) | **Post** /dfs-storage-classes/ | 
[**DeleteDfsStorageClass**](DfsStorageClassesAPI.md#DeleteDfsStorageClass) | **Delete** /dfs-storage-classes/{dfs_storage_class_id} | 
[**GetDfsStorageClass**](DfsStorageClassesAPI.md#GetDfsStorageClass) | **Get** /dfs-storage-classes/{dfs_storage_class_id} | 
[**GetDfsStorageClassPredictions**](DfsStorageClassesAPI.md#GetDfsStorageClassPredictions) | **Get** /dfs-storage-classes/{dfs_storage_class_id}/predictions | 
[**GetDfsStorageClassSamples**](DfsStorageClassesAPI.md#GetDfsStorageClassSamples) | **Get** /dfs-storage-classes/{dfs_storage_class_id}/samples | 
[**ListDfsStorageClasses**](DfsStorageClassesAPI.md#ListDfsStorageClasses) | **Get** /dfs-storage-classes/ | 
[**RemoveDfsStorageClassPools**](DfsStorageClassesAPI.md#RemoveDfsStorageClassPools) | **Post** /dfs-storage-classes/{dfs_storage_class_id}:remove-pools | 
[**UpdateDfsStorageClass**](DfsStorageClassesAPI.md#UpdateDfsStorageClass) | **Patch** /dfs-storage-classes/{dfs_storage_class_id} | 



## AddDfsStorageClassPools

> DfsStorageClassResp AddDfsStorageClassPools(ctx, dfsStorageClassId).Body(body).Execute()





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
	dfsStorageClassId := int64(789) // int64 | dfs storage class id
	body := *openapiclient.NewDfsStorageClassAddPoolReq(*openapiclient.NewDfsStorageClassAddPoolReqDfsStorageClass()) // DfsStorageClassAddPoolReq | pools info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsStorageClassesAPI.AddDfsStorageClassPools(context.Background(), dfsStorageClassId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsStorageClassesAPI.AddDfsStorageClassPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDfsStorageClassPools`: DfsStorageClassResp
	fmt.Fprintf(os.Stdout, "Response from `DfsStorageClassesAPI.AddDfsStorageClassPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsStorageClassId** | **int64** | dfs storage class id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDfsStorageClassPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsStorageClassAddPoolReq**](DfsStorageClassAddPoolReq.md) | pools info | 

### Return type

[**DfsStorageClassResp**](DfsStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDfsStorageClass

> DfsStorageClassResp CreateDfsStorageClass(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsStorageClassCreateReq(*openapiclient.NewDfsStorageClassCreateReqDfsStorageClass("Name_example", []openapiclient.PoolPolicy{*openapiclient.NewPoolPolicy()}, int64(123), int64(123), "WritePolicy_example")) // DfsStorageClassCreateReq | class info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsStorageClassesAPI.CreateDfsStorageClass(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsStorageClassesAPI.CreateDfsStorageClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsStorageClass`: DfsStorageClassResp
	fmt.Fprintf(os.Stdout, "Response from `DfsStorageClassesAPI.CreateDfsStorageClass`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsStorageClassCreateReq**](DfsStorageClassCreateReq.md) | class info | 

### Return type

[**DfsStorageClassResp**](DfsStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsStorageClass

> DfsStorageClassResp DeleteDfsStorageClass(ctx, dfsStorageClassId).Execute()





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
	dfsStorageClassId := int64(789) // int64 | dfs storage class id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsStorageClassesAPI.DeleteDfsStorageClass(context.Background(), dfsStorageClassId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsStorageClassesAPI.DeleteDfsStorageClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsStorageClass`: DfsStorageClassResp
	fmt.Fprintf(os.Stdout, "Response from `DfsStorageClassesAPI.DeleteDfsStorageClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsStorageClassId** | **int64** | dfs storage class id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsStorageClassResp**](DfsStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsStorageClass

> DfsStorageClassResp GetDfsStorageClass(ctx, dfsStorageClassId).Execute()





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
	dfsStorageClassId := int64(789) // int64 | dfs storage class id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsStorageClassesAPI.GetDfsStorageClass(context.Background(), dfsStorageClassId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsStorageClassesAPI.GetDfsStorageClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsStorageClass`: DfsStorageClassResp
	fmt.Fprintf(os.Stdout, "Response from `DfsStorageClassesAPI.GetDfsStorageClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsStorageClassId** | **int64** | dfs storage class id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsStorageClassResp**](DfsStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsStorageClassPredictions

> DfsStorageClassPredictionsResp GetDfsStorageClassPredictions(ctx, dfsStorageClassId).Execute()





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
	dfsStorageClassId := int64(789) // int64 | dfs storage class id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsStorageClassesAPI.GetDfsStorageClassPredictions(context.Background(), dfsStorageClassId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsStorageClassesAPI.GetDfsStorageClassPredictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsStorageClassPredictions`: DfsStorageClassPredictionsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsStorageClassesAPI.GetDfsStorageClassPredictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsStorageClassId** | **int64** | dfs storage class id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsStorageClassPredictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsStorageClassPredictionsResp**](DfsStorageClassPredictionsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsStorageClassSamples

> DfsStorageClassSamplesResp GetDfsStorageClassSamples(ctx, dfsStorageClassId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	dfsStorageClassId := int64(789) // int64 | dfs storage class id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsStorageClassesAPI.GetDfsStorageClassSamples(context.Background(), dfsStorageClassId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsStorageClassesAPI.GetDfsStorageClassSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsStorageClassSamples`: DfsStorageClassSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsStorageClassesAPI.GetDfsStorageClassSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsStorageClassId** | **int64** | dfs storage class id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsStorageClassSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**DfsStorageClassSamplesResp**](DfsStorageClassSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsStorageClasses

> DfsStorageClassesResp ListDfsStorageClasses(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).WritePolicy(writePolicy).Q(q).Sort(sort).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	writePolicy := "writePolicy_example" // string | write policy (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsStorageClassesAPI.ListDfsStorageClasses(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).WritePolicy(writePolicy).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsStorageClassesAPI.ListDfsStorageClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsStorageClasses`: DfsStorageClassesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsStorageClassesAPI.ListDfsStorageClasses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsStorageClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **writePolicy** | **string** | write policy | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DfsStorageClassesResp**](DfsStorageClassesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDfsStorageClassPools

> DfsStorageClassResp RemoveDfsStorageClassPools(ctx, dfsStorageClassId).Body(body).Execute()





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
	dfsStorageClassId := int64(789) // int64 | dfs storage class id
	body := *openapiclient.NewDfsStorageClassRemovePoolReq(*openapiclient.NewDfsStorageClassRemovePoolReqDfsStorageClass([]int64{int64(123)}, false)) // DfsStorageClassRemovePoolReq | pools info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsStorageClassesAPI.RemoveDfsStorageClassPools(context.Background(), dfsStorageClassId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsStorageClassesAPI.RemoveDfsStorageClassPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveDfsStorageClassPools`: DfsStorageClassResp
	fmt.Fprintf(os.Stdout, "Response from `DfsStorageClassesAPI.RemoveDfsStorageClassPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsStorageClassId** | **int64** | dfs storage class id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDfsStorageClassPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsStorageClassRemovePoolReq**](DfsStorageClassRemovePoolReq.md) | pools info | 

### Return type

[**DfsStorageClassResp**](DfsStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsStorageClass

> DfsStorageClassResp UpdateDfsStorageClass(ctx, dfsStorageClassId).Body(body).Execute()





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
	dfsStorageClassId := int64(789) // int64 | dfs storage class id
	body := *openapiclient.NewDfsStorageClassUpdateReq(*openapiclient.NewDfsStorageClassUpdateReqDfsStorageClass()) // DfsStorageClassUpdateReq | class info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsStorageClassesAPI.UpdateDfsStorageClass(context.Background(), dfsStorageClassId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsStorageClassesAPI.UpdateDfsStorageClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsStorageClass`: DfsStorageClassResp
	fmt.Fprintf(os.Stdout, "Response from `DfsStorageClassesAPI.UpdateDfsStorageClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsStorageClassId** | **int64** | dfs storage class id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsStorageClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsStorageClassUpdateReq**](DfsStorageClassUpdateReq.md) | class info | 

### Return type

[**DfsStorageClassResp**](DfsStorageClassResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

