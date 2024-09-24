# \OsZonePeriodsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOSZonePeriod**](OsZonePeriodsAPI.md#GetOSZonePeriod) | **Get** /os-zone-periods/{period_uuid} | 
[**ListOSZonePeriods**](OsZonePeriodsAPI.md#ListOSZonePeriods) | **Get** /os-zone-periods/ | 



## GetOSZonePeriod

> OSZonePeriodResp GetOSZonePeriod(ctx, periodUuid).Execute()





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
	periodUuid := "periodUuid_example" // string | os zone period uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsZonePeriodsAPI.GetOSZonePeriod(context.Background(), periodUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZonePeriodsAPI.GetOSZonePeriod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSZonePeriod`: OSZonePeriodResp
	fmt.Fprintf(os.Stdout, "Response from `OsZonePeriodsAPI.GetOSZonePeriod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**periodUuid** | **string** | os zone period uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSZonePeriodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSZonePeriodResp**](OSZonePeriodResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOSZonePeriods

> OSZonePeriodsResp ListOSZonePeriods(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsZonePeriodsAPI.ListOSZonePeriods(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZonePeriodsAPI.ListOSZonePeriods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSZonePeriods`: OSZonePeriodsResp
	fmt.Fprintf(os.Stdout, "Response from `OsZonePeriodsAPI.ListOSZonePeriods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSZonePeriodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSZonePeriodsResp**](OSZonePeriodsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

