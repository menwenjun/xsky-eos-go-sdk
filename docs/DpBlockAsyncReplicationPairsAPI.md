# \DpBlockAsyncReplicationPairsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncFailoverDpBlockAsyncReplicationPair**](DpBlockAsyncReplicationPairsAPI.md#AsyncFailoverDpBlockAsyncReplicationPair) | **Post** /dp-block-async-replication-pairs/{pair_id}:async-failover | 
[**CreateDpBlockAsyncReplicationPair**](DpBlockAsyncReplicationPairsAPI.md#CreateDpBlockAsyncReplicationPair) | **Post** /dp-block-async-replication-pairs/ | 
[**DeleteDpBlockAsyncReplicationPair**](DpBlockAsyncReplicationPairsAPI.md#DeleteDpBlockAsyncReplicationPair) | **Delete** /dp-block-async-replication-pairs/{pair_id} | 
[**FailbackDpBlockAsyncReplicationPair**](DpBlockAsyncReplicationPairsAPI.md#FailbackDpBlockAsyncReplicationPair) | **Post** /dp-block-async-replication-pairs/{pair_id}:failback | 
[**GetDpBlockAsyncReplicationPair**](DpBlockAsyncReplicationPairsAPI.md#GetDpBlockAsyncReplicationPair) | **Get** /dp-block-async-replication-pairs/{pair_id} | 
[**ListDpBlockAsyncReplicationPair**](DpBlockAsyncReplicationPairsAPI.md#ListDpBlockAsyncReplicationPair) | **Get** /dp-block-async-replication-pairs/ | 
[**PauseDpBlockAsyncReplicationPair**](DpBlockAsyncReplicationPairsAPI.md#PauseDpBlockAsyncReplicationPair) | **Post** /dp-block-async-replication-pairs/{pair_id}:pause | 
[**ResumeDpBlockAsyncReplicationPair**](DpBlockAsyncReplicationPairsAPI.md#ResumeDpBlockAsyncReplicationPair) | **Post** /dp-block-async-replication-pairs/{pair_id}:resume | 
[**RollbackDpBlockAsyncReplicationPair**](DpBlockAsyncReplicationPairsAPI.md#RollbackDpBlockAsyncReplicationPair) | **Post** /dp-block-async-replication-pairs/{pair_id}:rollback | 
[**SyncDpBlockAsyncReplicationPair**](DpBlockAsyncReplicationPairsAPI.md#SyncDpBlockAsyncReplicationPair) | **Post** /dp-block-async-replication-pairs/{pair_id}:sync | 
[**SyncFailoverDpBlockAsyncReplicationPair**](DpBlockAsyncReplicationPairsAPI.md#SyncFailoverDpBlockAsyncReplicationPair) | **Post** /dp-block-async-replication-pairs/{pair_id}:sync-failover | 
[**UpdateDpBlockAsyncReplicationPair**](DpBlockAsyncReplicationPairsAPI.md#UpdateDpBlockAsyncReplicationPair) | **Patch** /dp-block-async-replication-pairs/{pair_id} | 



## AsyncFailoverDpBlockAsyncReplicationPair

> DpBlockAsyncReplicationPairResp AsyncFailoverDpBlockAsyncReplicationPair(ctx, pairId).Execute()





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
	pairId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationPairsAPI.AsyncFailoverDpBlockAsyncReplicationPair(context.Background(), pairId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPairsAPI.AsyncFailoverDpBlockAsyncReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AsyncFailoverDpBlockAsyncReplicationPair`: DpBlockAsyncReplicationPairResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPairsAPI.AsyncFailoverDpBlockAsyncReplicationPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAsyncFailoverDpBlockAsyncReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpBlockAsyncReplicationPairResp**](DpBlockAsyncReplicationPairResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDpBlockAsyncReplicationPair

> DpBlockAsyncReplicationPairResp CreateDpBlockAsyncReplicationPair(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDpBlockAsyncReplicationPairCreateReq(*openapiclient.NewDpBlockAsyncReplicationPairCreateReqPair("ChainName_example", int64(123), "MasterClusterFsId_example", "MasterGateway_example", int64(123), int64(123), "MasterPoolName_example", int64(123), "MasterVolumeName_example", int64(123), "PolicyCron_example", "SlaveGateway_example", int64(123), "SlavePoolName_example", "SlaveVolumeName_example")) // DpBlockAsyncReplicationPairCreateReq | pair info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationPairsAPI.CreateDpBlockAsyncReplicationPair(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPairsAPI.CreateDpBlockAsyncReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDpBlockAsyncReplicationPair`: DpBlockAsyncReplicationPairResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPairsAPI.CreateDpBlockAsyncReplicationPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDpBlockAsyncReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DpBlockAsyncReplicationPairCreateReq**](DpBlockAsyncReplicationPairCreateReq.md) | pair info | 

### Return type

[**DpBlockAsyncReplicationPairResp**](DpBlockAsyncReplicationPairResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDpBlockAsyncReplicationPair

> DpBlockAsyncReplicationPairResp DeleteDpBlockAsyncReplicationPair(ctx, pairId).ReserveVolume(reserveVolume).Execute()





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
	pairId := int64(789) // int64 | resource id
	reserveVolume := true // bool | reserve replicated volume or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationPairsAPI.DeleteDpBlockAsyncReplicationPair(context.Background(), pairId).ReserveVolume(reserveVolume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPairsAPI.DeleteDpBlockAsyncReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDpBlockAsyncReplicationPair`: DpBlockAsyncReplicationPairResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPairsAPI.DeleteDpBlockAsyncReplicationPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDpBlockAsyncReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reserveVolume** | **bool** | reserve replicated volume or not | 

### Return type

[**DpBlockAsyncReplicationPairResp**](DpBlockAsyncReplicationPairResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FailbackDpBlockAsyncReplicationPair

> DpBlockAsyncReplicationPairResp FailbackDpBlockAsyncReplicationPair(ctx, pairId).Execute()





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
	pairId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationPairsAPI.FailbackDpBlockAsyncReplicationPair(context.Background(), pairId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPairsAPI.FailbackDpBlockAsyncReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FailbackDpBlockAsyncReplicationPair`: DpBlockAsyncReplicationPairResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPairsAPI.FailbackDpBlockAsyncReplicationPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFailbackDpBlockAsyncReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpBlockAsyncReplicationPairResp**](DpBlockAsyncReplicationPairResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDpBlockAsyncReplicationPair

> DpBlockAsyncReplicationPairResp GetDpBlockAsyncReplicationPair(ctx, pairId).Execute()





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
	pairId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationPairsAPI.GetDpBlockAsyncReplicationPair(context.Background(), pairId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPairsAPI.GetDpBlockAsyncReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDpBlockAsyncReplicationPair`: DpBlockAsyncReplicationPairResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPairsAPI.GetDpBlockAsyncReplicationPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDpBlockAsyncReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpBlockAsyncReplicationPairResp**](DpBlockAsyncReplicationPairResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpBlockAsyncReplicationPair

> DpBlockAsyncReplicationPairsResp ListDpBlockAsyncReplicationPair(ctx).BlockVolumeId(blockVolumeId).DpBlockAsyncReplicationPolicyId(dpBlockAsyncReplicationPolicyId).Execute()





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
	blockVolumeId := int64(789) // int64 | show volume snapshot replication pairs of specific block volume (optional)
	dpBlockAsyncReplicationPolicyId := int64(789) // int64 | show volume snapshot replication pairs of specific block async replication policy (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationPairsAPI.ListDpBlockAsyncReplicationPair(context.Background()).BlockVolumeId(blockVolumeId).DpBlockAsyncReplicationPolicyId(dpBlockAsyncReplicationPolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPairsAPI.ListDpBlockAsyncReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpBlockAsyncReplicationPair`: DpBlockAsyncReplicationPairsResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPairsAPI.ListDpBlockAsyncReplicationPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpBlockAsyncReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockVolumeId** | **int64** | show volume snapshot replication pairs of specific block volume | 
 **dpBlockAsyncReplicationPolicyId** | **int64** | show volume snapshot replication pairs of specific block async replication policy | 

### Return type

[**DpBlockAsyncReplicationPairsResp**](DpBlockAsyncReplicationPairsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseDpBlockAsyncReplicationPair

> DpBlockAsyncReplicationPairResp PauseDpBlockAsyncReplicationPair(ctx, pairId).Execute()





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
	pairId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationPairsAPI.PauseDpBlockAsyncReplicationPair(context.Background(), pairId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPairsAPI.PauseDpBlockAsyncReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseDpBlockAsyncReplicationPair`: DpBlockAsyncReplicationPairResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPairsAPI.PauseDpBlockAsyncReplicationPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseDpBlockAsyncReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpBlockAsyncReplicationPairResp**](DpBlockAsyncReplicationPairResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeDpBlockAsyncReplicationPair

> DpBlockAsyncReplicationPairResp ResumeDpBlockAsyncReplicationPair(ctx, pairId).Execute()





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
	pairId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationPairsAPI.ResumeDpBlockAsyncReplicationPair(context.Background(), pairId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPairsAPI.ResumeDpBlockAsyncReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumeDpBlockAsyncReplicationPair`: DpBlockAsyncReplicationPairResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPairsAPI.ResumeDpBlockAsyncReplicationPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeDpBlockAsyncReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpBlockAsyncReplicationPairResp**](DpBlockAsyncReplicationPairResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackDpBlockAsyncReplicationPair

> DpBlockAsyncReplicationPairResp RollbackDpBlockAsyncReplicationPair(ctx, pairId).Execute()





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
	pairId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationPairsAPI.RollbackDpBlockAsyncReplicationPair(context.Background(), pairId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPairsAPI.RollbackDpBlockAsyncReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RollbackDpBlockAsyncReplicationPair`: DpBlockAsyncReplicationPairResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPairsAPI.RollbackDpBlockAsyncReplicationPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackDpBlockAsyncReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpBlockAsyncReplicationPairResp**](DpBlockAsyncReplicationPairResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncDpBlockAsyncReplicationPair

> DpBlockAsyncReplicationPairResp SyncDpBlockAsyncReplicationPair(ctx, pairId).Execute()





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
	pairId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationPairsAPI.SyncDpBlockAsyncReplicationPair(context.Background(), pairId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPairsAPI.SyncDpBlockAsyncReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncDpBlockAsyncReplicationPair`: DpBlockAsyncReplicationPairResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPairsAPI.SyncDpBlockAsyncReplicationPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncDpBlockAsyncReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpBlockAsyncReplicationPairResp**](DpBlockAsyncReplicationPairResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncFailoverDpBlockAsyncReplicationPair

> DpBlockAsyncReplicationPairResp SyncFailoverDpBlockAsyncReplicationPair(ctx, pairId).Execute()





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
	pairId := int64(789) // int64 | resource id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationPairsAPI.SyncFailoverDpBlockAsyncReplicationPair(context.Background(), pairId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPairsAPI.SyncFailoverDpBlockAsyncReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncFailoverDpBlockAsyncReplicationPair`: DpBlockAsyncReplicationPairResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPairsAPI.SyncFailoverDpBlockAsyncReplicationPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncFailoverDpBlockAsyncReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpBlockAsyncReplicationPairResp**](DpBlockAsyncReplicationPairResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDpBlockAsyncReplicationPair

> DpBlockAsyncReplicationPairResp UpdateDpBlockAsyncReplicationPair(ctx, pairId).Body(body).Execute()





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
	pairId := int64(789) // int64 | resource id
	body := *openapiclient.NewDpBlockAsyncReplicationPairUpdateReq(*openapiclient.NewDpBlockAsyncReplicationPairUpdateReqPair()) // DpBlockAsyncReplicationPairUpdateReq | pair info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpBlockAsyncReplicationPairsAPI.UpdateDpBlockAsyncReplicationPair(context.Background(), pairId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpBlockAsyncReplicationPairsAPI.UpdateDpBlockAsyncReplicationPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDpBlockAsyncReplicationPair`: DpBlockAsyncReplicationPairResp
	fmt.Fprintf(os.Stdout, "Response from `DpBlockAsyncReplicationPairsAPI.UpdateDpBlockAsyncReplicationPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pairId** | **int64** | resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDpBlockAsyncReplicationPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DpBlockAsyncReplicationPairUpdateReq**](DpBlockAsyncReplicationPairUpdateReq.md) | pair info | 

### Return type

[**DpBlockAsyncReplicationPairResp**](DpBlockAsyncReplicationPairResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

