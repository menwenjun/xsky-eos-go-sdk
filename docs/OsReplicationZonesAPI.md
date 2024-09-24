# \OsReplicationZonesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOSReplicationZone**](OsReplicationZonesAPI.md#CreateOSReplicationZone) | **Post** /os-replication-zones/ | 
[**DeleteOSReplicationZone**](OsReplicationZonesAPI.md#DeleteOSReplicationZone) | **Delete** /os-replication-zones/{zone_uuid} | 
[**GetOSReplicationZone**](OsReplicationZonesAPI.md#GetOSReplicationZone) | **Get** /os-replication-zones/{zone_uuid} | 
[**GetOSReplicationZoneSamples**](OsReplicationZonesAPI.md#GetOSReplicationZoneSamples) | **Get** /os-replication-zones/{zone_uuid}/samples | 
[**ListOSReplicationZones**](OsReplicationZonesAPI.md#ListOSReplicationZones) | **Get** /os-replication-zones/ | 
[**UpdateOSReplicationZone**](OsReplicationZonesAPI.md#UpdateOSReplicationZone) | **Patch** /os-replication-zones/{zone_uuid} | 



## CreateOSReplicationZone

> OSReplicationZoneResp CreateOSReplicationZone(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewOSReplicationZoneCreateReq(*openapiclient.NewOSReplicationZoneCreateReqOSReplicationZone(int64(123), "Uuid_example")) // OSReplicationZoneCreateReq | os replication zone info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsReplicationZonesAPI.CreateOSReplicationZone(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsReplicationZonesAPI.CreateOSReplicationZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOSReplicationZone`: OSReplicationZoneResp
	fmt.Fprintf(os.Stdout, "Response from `OsReplicationZonesAPI.CreateOSReplicationZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOSReplicationZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OSReplicationZoneCreateReq**](OSReplicationZoneCreateReq.md) | os replication zone info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSReplicationZoneResp**](OSReplicationZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOSReplicationZone

> OSReplicationZoneResp DeleteOSReplicationZone(ctx, zoneUuid).Execute()





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
	zoneUuid := "zoneUuid_example" // string | os replication zone uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsReplicationZonesAPI.DeleteOSReplicationZone(context.Background(), zoneUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsReplicationZonesAPI.DeleteOSReplicationZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOSReplicationZone`: OSReplicationZoneResp
	fmt.Fprintf(os.Stdout, "Response from `OsReplicationZonesAPI.DeleteOSReplicationZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneUuid** | **string** | os replication zone uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOSReplicationZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSReplicationZoneResp**](OSReplicationZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSReplicationZone

> OSReplicationZoneRecordResp GetOSReplicationZone(ctx, zoneUuid).Execute()





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
	zoneUuid := "zoneUuid_example" // string | os replication zone uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsReplicationZonesAPI.GetOSReplicationZone(context.Background(), zoneUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsReplicationZonesAPI.GetOSReplicationZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSReplicationZone`: OSReplicationZoneRecordResp
	fmt.Fprintf(os.Stdout, "Response from `OsReplicationZonesAPI.GetOSReplicationZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneUuid** | **string** | os replication zone uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSReplicationZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSReplicationZoneRecordResp**](OSReplicationZoneRecordResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSReplicationZoneSamples

> OSReplicationZoneSamplesResp GetOSReplicationZoneSamples(ctx, zoneUuid).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	zoneUuid := "zoneUuid_example" // string | os replication zone uuid (default to "true")
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsReplicationZonesAPI.GetOSReplicationZoneSamples(context.Background(), zoneUuid).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsReplicationZonesAPI.GetOSReplicationZoneSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSReplicationZoneSamples`: OSReplicationZoneSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `OsReplicationZonesAPI.GetOSReplicationZoneSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneUuid** | **string** | os replication zone uuid | [default to &quot;true&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSReplicationZoneSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**OSReplicationZoneSamplesResp**](OSReplicationZoneSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOSReplicationZones

> OSReplicationZoneRecordsResp ListOSReplicationZones(ctx).Limit(limit).Offset(offset).Marker(marker).ReplicationUuid(replicationUuid).OsZoneUuid(osZoneUuid).ClusterId(clusterId).Execute()





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
	marker := "marker_example" // string | paging param (optional)
	replicationUuid := "replicationUuid_example" // string | os replication uuid (optional)
	osZoneUuid := "osZoneUuid_example" // string | os zone uuid (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsReplicationZonesAPI.ListOSReplicationZones(context.Background()).Limit(limit).Offset(offset).Marker(marker).ReplicationUuid(replicationUuid).OsZoneUuid(osZoneUuid).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsReplicationZonesAPI.ListOSReplicationZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSReplicationZones`: OSReplicationZoneRecordsResp
	fmt.Fprintf(os.Stdout, "Response from `OsReplicationZonesAPI.ListOSReplicationZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSReplicationZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **marker** | **string** | paging param | 
 **replicationUuid** | **string** | os replication uuid | 
 **osZoneUuid** | **string** | os zone uuid | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSReplicationZoneRecordsResp**](OSReplicationZoneRecordsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOSReplicationZone

> OSReplicationZoneResp UpdateOSReplicationZone(ctx, zoneUuid).Execute()





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
	zoneUuid := "zoneUuid_example" // string | os replication zone uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsReplicationZonesAPI.UpdateOSReplicationZone(context.Background(), zoneUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsReplicationZonesAPI.UpdateOSReplicationZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOSReplicationZone`: OSReplicationZoneResp
	fmt.Fprintf(os.Stdout, "Response from `OsReplicationZonesAPI.UpdateOSReplicationZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneUuid** | **string** | os replication zone uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOSReplicationZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSReplicationZoneResp**](OSReplicationZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

