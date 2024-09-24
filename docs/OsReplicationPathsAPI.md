# \OsReplicationPathsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOSReplicationPath**](OsReplicationPathsAPI.md#GetOSReplicationPath) | **Get** /os-replication-paths/{path_uuid} | 
[**ListOSReplicationPaths**](OsReplicationPathsAPI.md#ListOSReplicationPaths) | **Get** /os-replication-paths/ | 



## GetOSReplicationPath

> OSReplicationPathResp GetOSReplicationPath(ctx, pathUuid).Execute()





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
	pathUuid := "pathUuid_example" // string | replication path uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsReplicationPathsAPI.GetOSReplicationPath(context.Background(), pathUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsReplicationPathsAPI.GetOSReplicationPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSReplicationPath`: OSReplicationPathResp
	fmt.Fprintf(os.Stdout, "Response from `OsReplicationPathsAPI.GetOSReplicationPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathUuid** | **string** | replication path uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSReplicationPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSReplicationPathResp**](OSReplicationPathResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOSReplicationPaths

> OSReplicationPathsResp ListOSReplicationPaths(ctx).Limit(limit).Offset(offset).Marker(marker).ReplicationUuid(replicationUuid).AllowUnknownZone(allowUnknownZone).ClusterId(clusterId).Execute()





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
	allowUnknownZone := true // bool | allow has nil zone result(default true) (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsReplicationPathsAPI.ListOSReplicationPaths(context.Background()).Limit(limit).Offset(offset).Marker(marker).ReplicationUuid(replicationUuid).AllowUnknownZone(allowUnknownZone).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsReplicationPathsAPI.ListOSReplicationPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSReplicationPaths`: OSReplicationPathsResp
	fmt.Fprintf(os.Stdout, "Response from `OsReplicationPathsAPI.ListOSReplicationPaths`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSReplicationPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **marker** | **string** | paging param | 
 **replicationUuid** | **string** | os replication uuid | 
 **allowUnknownZone** | **bool** | allow has nil zone result(default true) | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSReplicationPathsResp**](OSReplicationPathsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

