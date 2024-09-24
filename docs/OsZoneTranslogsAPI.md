# \OsZoneTranslogsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOSZoneTranslog**](OsZoneTranslogsAPI.md#GetOSZoneTranslog) | **Get** /os-zone-translogs/{translog_uuid} | 
[**ListOSZoneTranslogs**](OsZoneTranslogsAPI.md#ListOSZoneTranslogs) | **Get** /os-zone-translogs/ | 



## GetOSZoneTranslog

> OSZoneTranslogResp GetOSZoneTranslog(ctx, translogUuid).Execute()





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
	translogUuid := "translogUuid_example" // string | translog uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsZoneTranslogsAPI.GetOSZoneTranslog(context.Background(), translogUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZoneTranslogsAPI.GetOSZoneTranslog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSZoneTranslog`: OSZoneTranslogResp
	fmt.Fprintf(os.Stdout, "Response from `OsZoneTranslogsAPI.GetOSZoneTranslog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**translogUuid** | **string** | translog uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSZoneTranslogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSZoneTranslogResp**](OSZoneTranslogResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOSZoneTranslogs

> OSZoneTranslogsResp ListOSZoneTranslogs(ctx).Limit(limit).Offset(offset).EpochUuid(epochUuid).ClusterId(clusterId).Execute()





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
	epochUuid := "epochUuid_example" // string | paging param (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsZoneTranslogsAPI.ListOSZoneTranslogs(context.Background()).Limit(limit).Offset(offset).EpochUuid(epochUuid).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZoneTranslogsAPI.ListOSZoneTranslogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSZoneTranslogs`: OSZoneTranslogsResp
	fmt.Fprintf(os.Stdout, "Response from `OsZoneTranslogsAPI.ListOSZoneTranslogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSZoneTranslogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **epochUuid** | **string** | paging param | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSZoneTranslogsResp**](OSZoneTranslogsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

