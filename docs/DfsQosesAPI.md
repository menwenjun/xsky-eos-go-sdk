# \DfsQosesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDfsQos**](DfsQosesAPI.md#CreateDfsQos) | **Post** /dfs-qoses/ | 
[**DeleteDfsQos**](DfsQosesAPI.md#DeleteDfsQos) | **Delete** /dfs-qoses/{dfs_qos_id} | 
[**GetDfsQos**](DfsQosesAPI.md#GetDfsQos) | **Get** /dfs-qoses/{dfs_qos_id} | 
[**GetDfsQosSamples**](DfsQosesAPI.md#GetDfsQosSamples) | **Get** /dfs-qoses/{dfs_qos_id}/samples | 
[**ListDfsQoses**](DfsQosesAPI.md#ListDfsQoses) | **Get** /dfs-qoses/ | 
[**UpdateDfsQos**](DfsQosesAPI.md#UpdateDfsQos) | **Patch** /dfs-qoses/{dfs_qos_id} | 



## CreateDfsQos

> DfsQosResp CreateDfsQos(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsQosCreateReq(*openapiclient.NewDfsQosCreateReqQos("DfsPath_example", int64(123), int64(123))) // DfsQosCreateReq | qos info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQosesAPI.CreateDfsQos(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQosesAPI.CreateDfsQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsQos`: DfsQosResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQosesAPI.CreateDfsQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsQosCreateReq**](DfsQosCreateReq.md) | qos info | 

### Return type

[**DfsQosResp**](DfsQosResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsQos

> DfsQosResp DeleteDfsQos(ctx, dfsQosId).Execute()





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
	dfsQosId := int64(789) // int64 | dfs qos id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQosesAPI.DeleteDfsQos(context.Background(), dfsQosId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQosesAPI.DeleteDfsQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsQos`: DfsQosResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQosesAPI.DeleteDfsQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsQosId** | **int64** | dfs qos id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsQosResp**](DfsQosResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsQos

> DfsQosResp GetDfsQos(ctx, dfsQosId).Execute()





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
	dfsQosId := int64(789) // int64 | dfs qos id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQosesAPI.GetDfsQos(context.Background(), dfsQosId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQosesAPI.GetDfsQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsQos`: DfsQosResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQosesAPI.GetDfsQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsQosId** | **int64** | dfs qos id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsQosResp**](DfsQosResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsQosSamples

> DfsQosSamplesResp GetDfsQosSamples(ctx, dfsQosId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	dfsQosId := int64(789) // int64 | dfs qos id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQosesAPI.GetDfsQosSamples(context.Background(), dfsQosId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQosesAPI.GetDfsQosSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsQosSamples`: DfsQosSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQosesAPI.GetDfsQosSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsQosId** | **int64** | dfs qos id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsQosSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**DfsQosSamplesResp**](DfsQosSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsQoses

> DfsQosesRecordsResp ListDfsQoses(ctx).PolicyId(policyId).Path(path).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).Execute()





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
	policyId := int64(789) // int64 | policy id
	path := "path_example" // string | dfs qos path (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQosesAPI.ListDfsQoses(context.Background()).PolicyId(policyId).Path(path).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQosesAPI.ListDfsQoses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsQoses`: DfsQosesRecordsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQosesAPI.ListDfsQoses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsQosesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyId** | **int64** | policy id | 
 **path** | **string** | dfs qos path | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 

### Return type

[**DfsQosesRecordsResp**](DfsQosesRecordsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsQos

> DfsQosResp UpdateDfsQos(ctx, dfsQosId).Body(body).ClusterId(clusterId).Execute()





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
	dfsQosId := int64(789) // int64 | qos id
	body := *openapiclient.NewDfsQosUpdateReq(*openapiclient.NewDfsQosUpdateReqQos(int64(123))) // DfsQosUpdateReq | dfs qos info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQosesAPI.UpdateDfsQos(context.Background(), dfsQosId).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQosesAPI.UpdateDfsQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsQos`: DfsQosResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQosesAPI.UpdateDfsQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsQosId** | **int64** | qos id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsQosUpdateReq**](DfsQosUpdateReq.md) | dfs qos info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**DfsQosResp**](DfsQosResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

