# \OsZoneLocksAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOSZoneLock**](OsZoneLocksAPI.md#CreateOSZoneLock) | **Post** /os-zone-locks/ | 
[**DeleteOSZoneLock**](OsZoneLocksAPI.md#DeleteOSZoneLock) | **Delete** /os-zone-locks/{lock_uuid} | 
[**GetOSZoneLock**](OsZoneLocksAPI.md#GetOSZoneLock) | **Get** /os-zone-locks/{lock_uuid} | 
[**ListOSZoneLocks**](OsZoneLocksAPI.md#ListOSZoneLocks) | **Get** /os-zone-locks/ | 
[**RefreshOSZoneLock**](OsZoneLocksAPI.md#RefreshOSZoneLock) | **Post** /os-zone-locks/{lock_uuid}:refresh | 



## CreateOSZoneLock

> OSZoneLockResp CreateOSZoneLock(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewOSZoneLockCreateReq() // OSZoneLockCreateReq | os zone lock info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsZoneLocksAPI.CreateOSZoneLock(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZoneLocksAPI.CreateOSZoneLock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOSZoneLock`: OSZoneLockResp
	fmt.Fprintf(os.Stdout, "Response from `OsZoneLocksAPI.CreateOSZoneLock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOSZoneLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OSZoneLockCreateReq**](OSZoneLockCreateReq.md) | os zone lock info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSZoneLockResp**](OSZoneLockResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOSZoneLock

> DeleteOSZoneLock(ctx, lockUuid).Execute()





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
	lockUuid := "lockUuid_example" // string | os zone lock uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OsZoneLocksAPI.DeleteOSZoneLock(context.Background(), lockUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZoneLocksAPI.DeleteOSZoneLock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lockUuid** | **string** | os zone lock uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOSZoneLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSZoneLock

> OSZoneLockResp GetOSZoneLock(ctx, lockUuid).Execute()





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
	lockUuid := "lockUuid_example" // string | os zone lock uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsZoneLocksAPI.GetOSZoneLock(context.Background(), lockUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZoneLocksAPI.GetOSZoneLock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSZoneLock`: OSZoneLockResp
	fmt.Fprintf(os.Stdout, "Response from `OsZoneLocksAPI.GetOSZoneLock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lockUuid** | **string** | os zone lock uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSZoneLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSZoneLockResp**](OSZoneLockResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOSZoneLocks

> OSZoneLocksResp ListOSZoneLocks(ctx).Limit(limit).Offset(offset).All(all).ClusterId(clusterId).Execute()





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
	all := true // bool | show all zone locks (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsZoneLocksAPI.ListOSZoneLocks(context.Background()).Limit(limit).Offset(offset).All(all).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZoneLocksAPI.ListOSZoneLocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSZoneLocks`: OSZoneLocksResp
	fmt.Fprintf(os.Stdout, "Response from `OsZoneLocksAPI.ListOSZoneLocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSZoneLocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **all** | **bool** | show all zone locks | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSZoneLocksResp**](OSZoneLocksResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshOSZoneLock

> OSZoneLockResp RefreshOSZoneLock(ctx, lockUuid).Execute()





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
	lockUuid := "lockUuid_example" // string | os zone lock uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsZoneLocksAPI.RefreshOSZoneLock(context.Background(), lockUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsZoneLocksAPI.RefreshOSZoneLock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshOSZoneLock`: OSZoneLockResp
	fmt.Fprintf(os.Stdout, "Response from `OsZoneLocksAPI.RefreshOSZoneLock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lockUuid** | **string** | os zone lock uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshOSZoneLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSZoneLockResp**](OSZoneLockResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

