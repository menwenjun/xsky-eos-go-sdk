# \PartitionsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPartition**](PartitionsAPI.md#GetPartition) | **Get** /partitions/{partition_id} | 
[**GetPartitionPredictions**](PartitionsAPI.md#GetPartitionPredictions) | **Get** /partitions/{partition_id}/predictions | 
[**ListPartitions**](PartitionsAPI.md#ListPartitions) | **Get** /partitions/ | 
[**UpdatePartitionOspPartitionType**](PartitionsAPI.md#UpdatePartitionOspPartitionType) | **Post** /partitions/{partition_id}:update-osp-partition-type | 



## GetPartition

> PartitionResp GetPartition(ctx, partitionId).Execute()





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
	partitionId := "partitionId_example" // string | partition id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PartitionsAPI.GetPartition(context.Background(), partitionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartitionsAPI.GetPartition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPartition`: PartitionResp
	fmt.Fprintf(os.Stdout, "Response from `PartitionsAPI.GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionId** | **string** | partition id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PartitionResp**](PartitionResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartitionPredictions

> PartitionPredictionsResp GetPartitionPredictions(ctx, partitionId).Execute()





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
	partitionId := int64(789) // int64 | partition id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PartitionsAPI.GetPartitionPredictions(context.Background(), partitionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartitionsAPI.GetPartitionPredictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPartitionPredictions`: PartitionPredictionsResp
	fmt.Fprintf(os.Stdout, "Response from `PartitionsAPI.GetPartitionPredictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionId** | **int64** | partition id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartitionPredictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PartitionPredictionsResp**](PartitionPredictionsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPartitions

> PartitionsResp ListPartitions(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).HostId(hostId).ClusterId(clusterId).Type_(type_).DiskId(diskId).Execute()





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
	hostId := int64(789) // int64 | host id (optional)
	clusterId := int64(789) // int64 | cluster id (optional)
	type_ := "type__example" // string | partition type (optional)
	diskId := int64(789) // int64 | disk id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PartitionsAPI.ListPartitions(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).HostId(hostId).ClusterId(clusterId).Type_(type_).DiskId(diskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartitionsAPI.ListPartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPartitions`: PartitionsResp
	fmt.Fprintf(os.Stdout, "Response from `PartitionsAPI.ListPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **hostId** | **int64** | host id | 
 **clusterId** | **int64** | cluster id | 
 **type_** | **string** | partition type | 
 **diskId** | **int64** | disk id | 

### Return type

[**PartitionsResp**](PartitionsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartitionOspPartitionType

> OsdResp UpdatePartitionOspPartitionType(ctx, partitionId).Body(body).Execute()





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
	partitionId := int64(789) // int64 | partition id
	body := *openapiclient.NewUpdateOspPartitionTypeReq() // UpdateOspPartitionTypeReq | osp partition type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PartitionsAPI.UpdatePartitionOspPartitionType(context.Background(), partitionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartitionsAPI.UpdatePartitionOspPartitionType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePartitionOspPartitionType`: OsdResp
	fmt.Fprintf(os.Stdout, "Response from `PartitionsAPI.UpdatePartitionOspPartitionType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionId** | **int64** | partition id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePartitionOspPartitionTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateOspPartitionTypeReq**](UpdateOspPartitionTypeReq.md) | osp partition type | 

### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

