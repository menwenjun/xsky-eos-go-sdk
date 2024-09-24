# \DpVolumeGroupSnapshotReplicationPairsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncFailoverDpVolumeGroupSnapshotReplicationPair**](DpVolumeGroupSnapshotReplicationPairsAPI.md#AsyncFailoverDpVolumeGroupSnapshotReplicationPair) | **Post** /dp-volume-group-snapshot-replication-pairs/{pair_id}:async-failover | 
[**CreateDpVolumeGroupSnapshotReplicationPair**](DpVolumeGroupSnapshotReplicationPairsAPI.md#CreateDpVolumeGroupSnapshotReplicationPair) | **Post** /dp-volume-group-snapshot-replication-pairs/ | 
[**DeleteDpVolumeGroupSnapshotReplicationPair**](DpVolumeGroupSnapshotReplicationPairsAPI.md#DeleteDpVolumeGroupSnapshotReplicationPair) | **Delete** /dp-volume-group-snapshot-replication-pairs/{pair_id} | 
[**GetDpVolumeGroupSnapshotReplicationPair**](DpVolumeGroupSnapshotReplicationPairsAPI.md#GetDpVolumeGroupSnapshotReplicationPair) | **Get** /dp-volume-group-snapshot-replication-pairs/{pair_id} | 
[**ListDpVolumeGroupSnapshotReplicationPair**](DpVolumeGroupSnapshotReplicationPairsAPI.md#ListDpVolumeGroupSnapshotReplicationPair) | **Get** /dp-volume-group-snapshot-replication-pairs/ | 
[**PauseDpVolumeGroupSnapshotReplicationPair**](DpVolumeGroupSnapshotReplicationPairsAPI.md#PauseDpVolumeGroupSnapshotReplicationPair) | **Post** /dp-volume-group-snapshot-replication-pairs/{pair_id}:pause | 
[**ResumeDpVolumeGroupSnapshotReplicationPair**](DpVolumeGroupSnapshotReplicationPairsAPI.md#ResumeDpVolumeGroupSnapshotReplicationPair) | **Post** /dp-volume-group-snapshot-replication-pairs/{pair_id}:resume | 
[**UpdateDpVolumeGroupSnapshotReplicationPair**](DpVolumeGroupSnapshotReplicationPairsAPI.md#UpdateDpVolumeGroupSnapshotReplicationPair) | **Patch** /dp-volume-group-snapshot-replication-pairs/{pair_id} | 



## AsyncFailoverDpVolumeGroupSnapshotReplicationPair

> DpVolumeGroupSnapshotReplicationPairsResp AsyncFailoverDpVolumeGroupSnapshotReplicationPair(ctx, pairId).Execute()





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
	pairId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpVolumeGroupSnapshotReplicationPairsAPI.AsyncFailoverDpVolumeGroupSnapshotReplicationPair(context.Background(), pairId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationPairsAPI.AsyncFailoverDpVolumeGroupSnapshotReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AsyncFailoverDpVolumeGroupSnapshotReplicationPair`: DpVolumeGroupSnapshotReplicationPairsResp
	fmt.Fprintf(os.Stdout, "Response from `DpVolumeGroupSnapshotReplicationPairsAPI.AsyncFailoverDpVolumeGroupSnapshotReplicationPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAsyncFailoverDpVolumeGroupSnapshotReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpVolumeGroupSnapshotReplicationPairsResp**](DpVolumeGroupSnapshotReplicationPairsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDpVolumeGroupSnapshotReplicationPair

> DpVolumeGroupSnapshotReplicationPairResp CreateDpVolumeGroupSnapshotReplicationPair(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDpVolumeGroupSnapshotReplicationPairCreateReq(*openapiclient.NewDpVolumeGroupSnapshotReplicationPairCreateReqGroupPair("ChainName_example", "MasterClusterFsId_example", "MasterGateway_example", int64(123), int64(123), "MasterVolumeGroupName_example", "PolicyCron_example", "SlaveGateway_example")) // DpVolumeGroupSnapshotReplicationPairCreateReq | pair info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpVolumeGroupSnapshotReplicationPairsAPI.CreateDpVolumeGroupSnapshotReplicationPair(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationPairsAPI.CreateDpVolumeGroupSnapshotReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDpVolumeGroupSnapshotReplicationPair`: DpVolumeGroupSnapshotReplicationPairResp
	fmt.Fprintf(os.Stdout, "Response from `DpVolumeGroupSnapshotReplicationPairsAPI.CreateDpVolumeGroupSnapshotReplicationPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDpVolumeGroupSnapshotReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DpVolumeGroupSnapshotReplicationPairCreateReq**](DpVolumeGroupSnapshotReplicationPairCreateReq.md) | pair info | 

### Return type

[**DpVolumeGroupSnapshotReplicationPairResp**](DpVolumeGroupSnapshotReplicationPairResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDpVolumeGroupSnapshotReplicationPair

> DeleteDpVolumeGroupSnapshotReplicationPair(ctx, pairId).Force(force).Execute()





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
	pairId := int64(789) // int64 | resource id
	force := true // bool | force delete volume group pair or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DpVolumeGroupSnapshotReplicationPairsAPI.DeleteDpVolumeGroupSnapshotReplicationPair(context.Background(), pairId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationPairsAPI.DeleteDpVolumeGroupSnapshotReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDpVolumeGroupSnapshotReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete volume group pair or not | 

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


## GetDpVolumeGroupSnapshotReplicationPair

> DpVolumeGroupSnapshotReplicationPairResp GetDpVolumeGroupSnapshotReplicationPair(ctx, pairId).Execute()





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
	pairId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpVolumeGroupSnapshotReplicationPairsAPI.GetDpVolumeGroupSnapshotReplicationPair(context.Background(), pairId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationPairsAPI.GetDpVolumeGroupSnapshotReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDpVolumeGroupSnapshotReplicationPair`: DpVolumeGroupSnapshotReplicationPairResp
	fmt.Fprintf(os.Stdout, "Response from `DpVolumeGroupSnapshotReplicationPairsAPI.GetDpVolumeGroupSnapshotReplicationPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDpVolumeGroupSnapshotReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpVolumeGroupSnapshotReplicationPairResp**](DpVolumeGroupSnapshotReplicationPairResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpVolumeGroupSnapshotReplicationPair

> DpVolumeGroupSnapshotReplicationPairsResp ListDpVolumeGroupSnapshotReplicationPair(ctx).VolumeGroupId(volumeGroupId).DpVolumeGroupSnapshotReplicationPolicyId(dpVolumeGroupSnapshotReplicationPolicyId).Execute()





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
	volumeGroupId := int64(789) // int64 | related volume group id (optional)
	dpVolumeGroupSnapshotReplicationPolicyId := int64(789) // int64 | related policy id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpVolumeGroupSnapshotReplicationPairsAPI.ListDpVolumeGroupSnapshotReplicationPair(context.Background()).VolumeGroupId(volumeGroupId).DpVolumeGroupSnapshotReplicationPolicyId(dpVolumeGroupSnapshotReplicationPolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationPairsAPI.ListDpVolumeGroupSnapshotReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpVolumeGroupSnapshotReplicationPair`: DpVolumeGroupSnapshotReplicationPairsResp
	fmt.Fprintf(os.Stdout, "Response from `DpVolumeGroupSnapshotReplicationPairsAPI.ListDpVolumeGroupSnapshotReplicationPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpVolumeGroupSnapshotReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeGroupId** | **int64** | related volume group id | 
 **dpVolumeGroupSnapshotReplicationPolicyId** | **int64** | related policy id | 

### Return type

[**DpVolumeGroupSnapshotReplicationPairsResp**](DpVolumeGroupSnapshotReplicationPairsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseDpVolumeGroupSnapshotReplicationPair

> DpVolumeGroupSnapshotReplicationPairsResp PauseDpVolumeGroupSnapshotReplicationPair(ctx, pairId).Execute()





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
	pairId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpVolumeGroupSnapshotReplicationPairsAPI.PauseDpVolumeGroupSnapshotReplicationPair(context.Background(), pairId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationPairsAPI.PauseDpVolumeGroupSnapshotReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseDpVolumeGroupSnapshotReplicationPair`: DpVolumeGroupSnapshotReplicationPairsResp
	fmt.Fprintf(os.Stdout, "Response from `DpVolumeGroupSnapshotReplicationPairsAPI.PauseDpVolumeGroupSnapshotReplicationPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseDpVolumeGroupSnapshotReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpVolumeGroupSnapshotReplicationPairsResp**](DpVolumeGroupSnapshotReplicationPairsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeDpVolumeGroupSnapshotReplicationPair

> DpVolumeGroupSnapshotReplicationPairResp ResumeDpVolumeGroupSnapshotReplicationPair(ctx, pairId).Execute()





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
	pairId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpVolumeGroupSnapshotReplicationPairsAPI.ResumeDpVolumeGroupSnapshotReplicationPair(context.Background(), pairId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationPairsAPI.ResumeDpVolumeGroupSnapshotReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumeDpVolumeGroupSnapshotReplicationPair`: DpVolumeGroupSnapshotReplicationPairResp
	fmt.Fprintf(os.Stdout, "Response from `DpVolumeGroupSnapshotReplicationPairsAPI.ResumeDpVolumeGroupSnapshotReplicationPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeDpVolumeGroupSnapshotReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpVolumeGroupSnapshotReplicationPairResp**](DpVolumeGroupSnapshotReplicationPairResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDpVolumeGroupSnapshotReplicationPair

> DpVolumeGroupSnapshotReplicationPairResp UpdateDpVolumeGroupSnapshotReplicationPair(ctx, pairId).Body(body).Execute()





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
	pairId := int64(789) // int64 | resource id
	body := *openapiclient.NewDpVolumeGroupSnapshotReplicationPairUpdateReq(*openapiclient.NewDpVolumeGroupSnapshotReplicationPairUpdateReqPair()) // DpVolumeGroupSnapshotReplicationPairUpdateReq | pair info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpVolumeGroupSnapshotReplicationPairsAPI.UpdateDpVolumeGroupSnapshotReplicationPair(context.Background(), pairId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpVolumeGroupSnapshotReplicationPairsAPI.UpdateDpVolumeGroupSnapshotReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDpVolumeGroupSnapshotReplicationPair`: DpVolumeGroupSnapshotReplicationPairResp
	fmt.Fprintf(os.Stdout, "Response from `DpVolumeGroupSnapshotReplicationPairsAPI.UpdateDpVolumeGroupSnapshotReplicationPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDpVolumeGroupSnapshotReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DpVolumeGroupSnapshotReplicationPairUpdateReq**](DpVolumeGroupSnapshotReplicationPairUpdateReq.md) | pair info | 

### Return type

[**DpVolumeGroupSnapshotReplicationPairResp**](DpVolumeGroupSnapshotReplicationPairResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

