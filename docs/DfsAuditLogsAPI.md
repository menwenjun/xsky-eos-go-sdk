# \DfsAuditLogsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDfsAuditLogs**](DfsAuditLogsAPI.md#ListDfsAuditLogs) | **Get** /dfs-audit-logs/ | 
[**UpdateDfsAuditLog**](DfsAuditLogsAPI.md#UpdateDfsAuditLog) | **Patch** /dfs-audit-logs/{dfs_audit_log_id} | 



## ListDfsAuditLogs

> DfsAuditLogsResp ListDfsAuditLogs(ctx).ClusterId(clusterId).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsAuditLogsAPI.ListDfsAuditLogs(context.Background()).ClusterId(clusterId).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsAuditLogsAPI.ListDfsAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsAuditLogs`: DfsAuditLogsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsAuditLogsAPI.ListDfsAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DfsAuditLogsResp**](DfsAuditLogsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsAuditLog

> DfsAuditLogResp UpdateDfsAuditLog(ctx, dfsAuditLogId).Body(body).Execute()





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
	dfsAuditLogId := int64(789) // int64 | dfs audit log id
	body := *openapiclient.NewDfsAuditLogUpdateReq(*openapiclient.NewDfsAuditLogUpdateReqAuditLog(int64(123))) // DfsAuditLogUpdateReq | dfs audit log info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsAuditLogsAPI.UpdateDfsAuditLog(context.Background(), dfsAuditLogId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsAuditLogsAPI.UpdateDfsAuditLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsAuditLog`: DfsAuditLogResp
	fmt.Fprintf(os.Stdout, "Response from `DfsAuditLogsAPI.UpdateDfsAuditLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsAuditLogId** | **int64** | dfs audit log id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsAuditLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsAuditLogUpdateReq**](DfsAuditLogUpdateReq.md) | dfs audit log info | 

### Return type

[**DfsAuditLogResp**](DfsAuditLogResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

