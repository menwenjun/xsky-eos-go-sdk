# \OsZonesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObjectStorageZone**](OsZonesAPI.md#CreateObjectStorageZone) | **Post** /os-zones/ | 
[**DeleteObjectStorageZone**](OsZonesAPI.md#DeleteObjectStorageZone) | **Delete** /os-zones/{zone_uuid} | 
[**GetObjectStorageZone**](OsZonesAPI.md#GetObjectStorageZone) | **Get** /os-zones/{zone_uuid} | 
[**GetObjectStorageZoneSamples**](OsZonesAPI.md#GetObjectStorageZoneSamples) | **Get** /os-zones/{zone_uuid}/samples | 
[**ListObjectStorageZones**](OsZonesAPI.md#ListObjectStorageZones) | **Get** /os-zones/ | 
[**UpdateOSZonesClockDiff**](OsZonesAPI.md#UpdateOSZonesClockDiff) | **Post** /os-zones/{zone_uuid} | 



## CreateObjectStorageZone

> ObjectStorageZoneResp CreateObjectStorageZone(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewObjectStorageZoneCreateReq() // ObjectStorageZoneCreateReq | zone info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsZonesAPI.CreateObjectStorageZone(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZonesAPI.CreateObjectStorageZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateObjectStorageZone`: ObjectStorageZoneResp
	fmt.Fprintf(os.Stdout, "Response from `OsZonesAPI.CreateObjectStorageZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectStorageZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ObjectStorageZoneCreateReq**](ObjectStorageZoneCreateReq.md) | zone info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**ObjectStorageZoneResp**](ObjectStorageZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObjectStorageZone

> ObjectStorageZoneResp DeleteObjectStorageZone(ctx, zoneUuid).Force(force).Execute()





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
	zoneUuid := "zoneUuid_example" // string | os zone uuid
	force := true // bool | delete os zone forcefully or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsZonesAPI.DeleteObjectStorageZone(context.Background(), zoneUuid).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZonesAPI.DeleteObjectStorageZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteObjectStorageZone`: ObjectStorageZoneResp
	fmt.Fprintf(os.Stdout, "Response from `OsZonesAPI.DeleteObjectStorageZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneUuid** | **string** | os zone uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectStorageZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | delete os zone forcefully or not | 

### Return type

[**ObjectStorageZoneResp**](ObjectStorageZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageZone

> ObjectStorageZoneRecordResp GetObjectStorageZone(ctx, zoneUuid).Execute()





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
	zoneUuid := "zoneUuid_example" // string | object storage zone uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsZonesAPI.GetObjectStorageZone(context.Background(), zoneUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZonesAPI.GetObjectStorageZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageZone`: ObjectStorageZoneRecordResp
	fmt.Fprintf(os.Stdout, "Response from `OsZonesAPI.GetObjectStorageZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneUuid** | **string** | object storage zone uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectStorageZoneRecordResp**](ObjectStorageZoneRecordResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageZoneSamples

> ObjectStorageZoneSamplesResp GetObjectStorageZoneSamples(ctx, zoneUuid).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	zoneUuid := "zoneUuid_example" // string | object storage zone uuid (default to "true")
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsZonesAPI.GetObjectStorageZoneSamples(context.Background(), zoneUuid).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZonesAPI.GetObjectStorageZoneSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageZoneSamples`: ObjectStorageZoneSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `OsZonesAPI.GetObjectStorageZoneSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneUuid** | **string** | object storage zone uuid | [default to &quot;true&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageZoneSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**ObjectStorageZoneSamplesResp**](ObjectStorageZoneSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObjectStorageZones

> ObjectStorageZonesRecordResp ListObjectStorageZones(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).Local(local).Master(master).Name(name).ClusterId(clusterId).Execute()





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
	local := true // bool | local or non-local zones (optional)
	master := true // bool | master or non-master zones (optional)
	name := "name_example" // string | name of zone (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsZonesAPI.ListObjectStorageZones(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).Local(local).Master(master).Name(name).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZonesAPI.ListObjectStorageZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListObjectStorageZones`: ObjectStorageZonesRecordResp
	fmt.Fprintf(os.Stdout, "Response from `OsZonesAPI.ListObjectStorageZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListObjectStorageZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **local** | **bool** | local or non-local zones | 
 **master** | **bool** | master or non-master zones | 
 **name** | **string** | name of zone | 
 **clusterId** | **string** | cluster id | 

### Return type

[**ObjectStorageZonesRecordResp**](ObjectStorageZonesRecordResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOSZonesClockDiff

> ObjectStorageZoneResp UpdateOSZonesClockDiff(ctx, zoneUuid).Body(body).Execute()





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
	zoneUuid := "zoneUuid_example" // string | os zone uuid
	body := *openapiclient.NewOSZonePairsUpdateReq() // OSZonePairsUpdateReq | zone pairs info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsZonesAPI.UpdateOSZonesClockDiff(context.Background(), zoneUuid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZonesAPI.UpdateOSZonesClockDiff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOSZonesClockDiff`: ObjectStorageZoneResp
	fmt.Fprintf(os.Stdout, "Response from `OsZonesAPI.UpdateOSZonesClockDiff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneUuid** | **string** | os zone uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOSZonesClockDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSZonePairsUpdateReq**](OSZonePairsUpdateReq.md) | zone pairs info | 

### Return type

[**ObjectStorageZoneResp**](ObjectStorageZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

