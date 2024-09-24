# \DisksAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePartitions**](DisksAPI.md#CreatePartitions) | **Put** /disks/{disk_id}/partitions | 
[**DeletePartitions**](DisksAPI.md#DeletePartitions) | **Delete** /disks/{disk_id}/partitions | 
[**GetDisk**](DisksAPI.md#GetDisk) | **Get** /disks/{disk_id} | 
[**GetDiskSamples**](DisksAPI.md#GetDiskSamples) | **Get** /disks/{disk_id}/samples | 
[**ListDisks**](DisksAPI.md#ListDisks) | **Get** /disks/ | 
[**RebuildDisk**](DisksAPI.md#RebuildDisk) | **Post** /disks/{disk_id}:rebuild | 
[**UpdateDisk**](DisksAPI.md#UpdateDisk) | **Patch** /disks/{disk_id} | 



## CreatePartitions

> DiskResp CreatePartitions(ctx, diskId).Body(body).Num(num).Execute()





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
	diskId := int64(789) // int64 | disk id
	body := *openapiclient.NewPartitionsCreateReq() // PartitionsCreateReq | partitions info
	num := int64(789) // int64 | num of cache partitions to create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisksAPI.CreatePartitions(context.Background(), diskId).Body(body).Num(num).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisksAPI.CreatePartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePartitions`: DiskResp
	fmt.Fprintf(os.Stdout, "Response from `DisksAPI.CreatePartitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**diskId** | **int64** | disk id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PartitionsCreateReq**](PartitionsCreateReq.md) | partitions info | 
 **num** | **int64** | num of cache partitions to create | 

### Return type

[**DiskResp**](DiskResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePartitions

> DiskResp DeletePartitions(ctx, diskId).Execute()





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
	diskId := int64(789) // int64 | disk id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisksAPI.DeletePartitions(context.Background(), diskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisksAPI.DeletePartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePartitions`: DiskResp
	fmt.Fprintf(os.Stdout, "Response from `DisksAPI.DeletePartitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**diskId** | **int64** | disk id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DiskResp**](DiskResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDisk

> DiskResp GetDisk(ctx, diskId).Execute()





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
	diskId := int64(789) // int64 | disk id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisksAPI.GetDisk(context.Background(), diskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisksAPI.GetDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDisk`: DiskResp
	fmt.Fprintf(os.Stdout, "Response from `DisksAPI.GetDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**diskId** | **int64** | disk id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DiskResp**](DiskResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiskSamples

> DiskSamplesResp GetDiskSamples(ctx, diskId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	diskId := int64(789) // int64 | disk id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisksAPI.GetDiskSamples(context.Background(), diskId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisksAPI.GetDiskSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiskSamples`: DiskSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `DisksAPI.GetDiskSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**diskId** | **int64** | disk id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiskSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**DiskSamplesResp**](DiskSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDisks

> DisksResp ListDisks(ctx).Limit(limit).Offset(offset).HostId(hostId).ClusterId(clusterId).RebuildDiskId(rebuildDiskId).IsCache(isCache).DiskType(diskType).Usage(usage).Used(used).Q(q).Sort(sort).Device(device).Status(status).Ids(ids).HostIds(hostIds).OrderBy(orderBy).Execute()





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
	hostId := int64(789) // int64 | host id (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)
	rebuildDiskId := int64(789) // int64 | disk id to order disk by size and type for rebuilding disk (optional)
	isCache := true // bool | filter cache disk, deprecated, use `usage=partition` instead (optional)
	diskType := "diskType_example" // string | filter disk type (optional)
	usage := "usage_example" // string | filter disk usage (optional)
	used := true // bool | filter used (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	device := "device_example" // string | device name of disk (optional)
	status := "status_example" // string | disk status (optional)
	ids := "ids_example" // string | disk comma separate ids (optional)
	hostIds := "hostIds_example" // string | disk.host comma separate ids (optional)
	orderBy := "orderBy_example" // string | sort param without search, split with comma, e.g. -host.name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisksAPI.ListDisks(context.Background()).Limit(limit).Offset(offset).HostId(hostId).ClusterId(clusterId).RebuildDiskId(rebuildDiskId).IsCache(isCache).DiskType(diskType).Usage(usage).Used(used).Q(q).Sort(sort).Device(device).Status(status).Ids(ids).HostIds(hostIds).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisksAPI.ListDisks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDisks`: DisksResp
	fmt.Fprintf(os.Stdout, "Response from `DisksAPI.ListDisks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDisksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **hostId** | **int64** | host id | 
 **clusterId** | **string** | cluster id | 
 **rebuildDiskId** | **int64** | disk id to order disk by size and type for rebuilding disk | 
 **isCache** | **bool** | filter cache disk, deprecated, use &#x60;usage&#x3D;partition&#x60; instead | 
 **diskType** | **string** | filter disk type | 
 **usage** | **string** | filter disk usage | 
 **used** | **bool** | filter used | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **device** | **string** | device name of disk | 
 **status** | **string** | disk status | 
 **ids** | **string** | disk comma separate ids | 
 **hostIds** | **string** | disk.host comma separate ids | 
 **orderBy** | **string** | sort param without search, split with comma, e.g. -host.name | 

### Return type

[**DisksResp**](DisksResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebuildDisk

> DiskResp RebuildDisk(ctx, diskId).Body(body).Execute()





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
	diskId := int64(789) // int64 | disk id
	body := *openapiclient.NewDiskRebuildReq(*openapiclient.NewDiskRebuildReqDisk()) // DiskRebuildReq | new disk info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisksAPI.RebuildDisk(context.Background(), diskId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisksAPI.RebuildDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebuildDisk`: DiskResp
	fmt.Fprintf(os.Stdout, "Response from `DisksAPI.RebuildDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**diskId** | **int64** | disk id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebuildDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DiskRebuildReq**](DiskRebuildReq.md) | new disk info | 

### Return type

[**DiskResp**](DiskResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDisk

> DiskResp UpdateDisk(ctx, diskId).Disk(disk).Execute()





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
	diskId := int64(789) // int64 | disk id
	disk := *openapiclient.NewDiskUpdateReq(*openapiclient.NewDiskUpdateReqDisk()) // DiskUpdateReq | disk info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisksAPI.UpdateDisk(context.Background(), diskId).Disk(disk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisksAPI.UpdateDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDisk`: DiskResp
	fmt.Fprintf(os.Stdout, "Response from `DisksAPI.UpdateDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**diskId** | **int64** | disk id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disk** | [**DiskUpdateReq**](DiskUpdateReq.md) | disk info | 

### Return type

[**DiskResp**](DiskResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

