# \SystemLogsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadSystemLogs**](SystemLogsAPI.md#DownloadSystemLogs) | **Get** /system-logs/archive | 
[**ListSystemLogs**](SystemLogsAPI.md#ListSystemLogs) | **Get** /system-logs/ | 



## DownloadSystemLogs

> *os.File DownloadSystemLogs(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemLogsAPI.DownloadSystemLogs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemLogsAPI.DownloadSystemLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadSystemLogs`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SystemLogsAPI.DownloadSystemLogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadSystemLogsRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystemLogs

> SystemLogsResp ListSystemLogs(ctx).HostId(hostId).Catalog(catalog).Limit(limit).Offset(offset).Execute()





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
	hostId := int64(789) // int64 | The id of host system logs belong to
	catalog := "catalog_example" // string | The name of catalog system logs belong to
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemLogsAPI.ListSystemLogs(context.Background()).HostId(hostId).Catalog(catalog).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemLogsAPI.ListSystemLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSystemLogs`: SystemLogsResp
	fmt.Fprintf(os.Stdout, "Response from `SystemLogsAPI.ListSystemLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSystemLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostId** | **int64** | The id of host system logs belong to | 
 **catalog** | **string** | The name of catalog system logs belong to | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**SystemLogsResp**](SystemLogsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

