# \PoolsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOsdsToPool**](PoolsAPI.md#AddOsdsToPool) | **Put** /pools/{pool_id}/osds | 
[**AddPoolToOutsideBackend**](PoolsAPI.md#AddPoolToOutsideBackend) | **Post** /pools:add-to-outside-backend | 
[**AddPoolsToOspDataBackend**](PoolsAPI.md#AddPoolsToOspDataBackend) | **Post** /pools/:add-pools-to-osp-data-backend | 
[**CalcCapacity**](PoolsAPI.md#CalcCapacity) | **Post** /pools/:calc-capacity | 
[**CheckFull**](PoolsAPI.md#CheckFull) | **Post** /pools/:check-full | 
[**CreatePool**](PoolsAPI.md#CreatePool) | **Post** /pools/ | 
[**DeletePool**](PoolsAPI.md#DeletePool) | **Delete** /pools/{pool_id} | 
[**DisablePoolDeviceTypeCheck**](PoolsAPI.md#DisablePoolDeviceTypeCheck) | **Post** /pools/{pool_id}:disable-device-type-check | 
[**DisablePoolNuma**](PoolsAPI.md#DisablePoolNuma) | **Post** /pools/{pool_id}:disable-numa | 
[**EnablePoolDeviceTypeCheck**](PoolsAPI.md#EnablePoolDeviceTypeCheck) | **Post** /pools/{pool_id}:enable-device-type-check | 
[**EnablePoolNuma**](PoolsAPI.md#EnablePoolNuma) | **Post** /pools/{pool_id}:enable-numa | 
[**GetPool**](PoolsAPI.md#GetPool) | **Get** /pools/{pool_id} | 
[**GetPoolPredictions**](PoolsAPI.md#GetPoolPredictions) | **Get** /pools/{pool_id}/predictions | 
[**GetPoolSamples**](PoolsAPI.md#GetPoolSamples) | **Get** /pools/{pool_id}/samples | 
[**GetPoolTopology**](PoolsAPI.md#GetPoolTopology) | **Get** /pools/{pool_id}/topology | 
[**InitializeEmptyPool**](PoolsAPI.md#InitializeEmptyPool) | **Post** /pools/{pool_id}:initialize | 
[**ListPools**](PoolsAPI.md#ListPools) | **Get** /pools/ | 
[**RemoveFromOutsideBackend**](PoolsAPI.md#RemoveFromOutsideBackend) | **Post** /pools:remove-from-outside-backend | 
[**RemoveOsdsFromPool**](PoolsAPI.md#RemoveOsdsFromPool) | **Delete** /pools/{pool_id}/osds | 
[**RemovePoolsFromOspDataBackend**](PoolsAPI.md#RemovePoolsFromOspDataBackend) | **Post** /pools/:remove-pools-from-osp-data-backend | 
[**ReweightPool**](PoolsAPI.md#ReweightPool) | **Post** /pools/{pool_id}:reweight | 
[**SwitchPoolRole**](PoolsAPI.md#SwitchPoolRole) | **Post** /pools/{pool_id}:switch-role | 
[**UpdateECPoolCrushRule**](PoolsAPI.md#UpdateECPoolCrushRule) | **Post** /pools/{pool_id}:update-ec-crush-rule | 
[**UpdatePool**](PoolsAPI.md#UpdatePool) | **Patch** /pools/{pool_id} | 
[**UpdatePoolGCPolicyPlan**](PoolsAPI.md#UpdatePoolGCPolicyPlan) | **Post** /pools/{pool_id}:update-gc-policy-plan | 



## AddOsdsToPool

> PoolResp AddOsdsToPool(ctx, poolId).Body(body).Execute()





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
	poolId := int64(789) // int64 | pool id
	body := *openapiclient.NewOsdsAddReq([]int64{int64(123)}) // OsdsAddReq | osd infos

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.AddOsdsToPool(context.Background(), poolId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.AddOsdsToPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOsdsToPool`: PoolResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.AddOsdsToPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOsdsToPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OsdsAddReq**](OsdsAddReq.md) | osd infos | 

### Return type

[**PoolResp**](PoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddPoolToOutsideBackend

> OspPoolOutsideResp AddPoolToOutsideBackend(ctx).Body(body).Execute()





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
	body := *openapiclient.NewOspPoolOutsideReq() // OspPoolOutsideReq | pool info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.AddPoolToOutsideBackend(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.AddPoolToOutsideBackend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPoolToOutsideBackend`: OspPoolOutsideResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.AddPoolToOutsideBackend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPoolToOutsideBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OspPoolOutsideReq**](OspPoolOutsideReq.md) | pool info | 

### Return type

[**OspPoolOutsideResp**](OspPoolOutsideResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddPoolsToOspDataBackend

> TaskIDResp AddPoolsToOspDataBackend(ctx).Body(body).Execute()





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
	body := *openapiclient.NewOspDataBackendPoolsOperateReq() // OspDataBackendPoolsOperateReq | pool info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.AddPoolsToOspDataBackend(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.AddPoolsToOspDataBackend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPoolsToOspDataBackend`: TaskIDResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.AddPoolsToOspDataBackend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPoolsToOspDataBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OspDataBackendPoolsOperateReq**](OspDataBackendPoolsOperateReq.md) | pool info | 

### Return type

[**TaskIDResp**](TaskIDResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CalcCapacity

> PoolCapacityResp CalcCapacity(ctx).Body(body).Execute()





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
	body := *openapiclient.NewPoolCapacityReq() // PoolCapacityReq | pool info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.CalcCapacity(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.CalcCapacity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalcCapacity`: PoolCapacityResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.CalcCapacity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCalcCapacityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PoolCapacityReq**](PoolCapacityReq.md) | pool info | 

### Return type

[**PoolCapacityResp**](PoolCapacityResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckFull

> PoolFullCheckResp CheckFull(ctx).Body(body).Execute()





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
	body := *openapiclient.NewPoolFullCheckReq([]int64{int64(123)}) // PoolFullCheckReq | pool ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.CheckFull(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.CheckFull``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckFull`: PoolFullCheckResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.CheckFull`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckFullRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PoolFullCheckReq**](PoolFullCheckReq.md) | pool ids | 

### Return type

[**PoolFullCheckResp**](PoolFullCheckResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePool

> PoolResp CreatePool(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewEmptyPoolCreateReq() // EmptyPoolCreateReq | the empty pool info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.CreatePool(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.CreatePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePool`: PoolResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.CreatePool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EmptyPoolCreateReq**](EmptyPoolCreateReq.md) | the empty pool info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**PoolResp**](PoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePool

> PoolResp DeletePool(ctx, poolId).Force(force).Execute()





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
	poolId := int64(789) // int64 | pool id
	force := true // bool | force delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.DeletePool(context.Background(), poolId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.DeletePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePool`: PoolResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.DeletePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 

### Return type

[**PoolResp**](PoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisablePoolDeviceTypeCheck

> PoolResp DisablePoolDeviceTypeCheck(ctx, poolId).Execute()





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
	poolId := int64(789) // int64 | pool id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.DisablePoolDeviceTypeCheck(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.DisablePoolDeviceTypeCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisablePoolDeviceTypeCheck`: PoolResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.DisablePoolDeviceTypeCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisablePoolDeviceTypeCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoolResp**](PoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisablePoolNuma

> PoolResp DisablePoolNuma(ctx, poolId).Execute()





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
	poolId := int64(789) // int64 | pool id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.DisablePoolNuma(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.DisablePoolNuma``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisablePoolNuma`: PoolResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.DisablePoolNuma`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisablePoolNumaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoolResp**](PoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnablePoolDeviceTypeCheck

> PoolResp EnablePoolDeviceTypeCheck(ctx, poolId).Execute()





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
	poolId := int64(789) // int64 | pool id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.EnablePoolDeviceTypeCheck(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.EnablePoolDeviceTypeCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnablePoolDeviceTypeCheck`: PoolResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.EnablePoolDeviceTypeCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnablePoolDeviceTypeCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoolResp**](PoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnablePoolNuma

> PoolResp EnablePoolNuma(ctx, poolId).Execute()





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
	poolId := int64(789) // int64 | pool id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.EnablePoolNuma(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.EnablePoolNuma``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnablePoolNuma`: PoolResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.EnablePoolNuma`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnablePoolNumaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoolResp**](PoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPool

> PoolResp GetPool(ctx, poolId).Execute()





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
	poolId := int64(789) // int64 | pool id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.GetPool(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.GetPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPool`: PoolResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.GetPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoolResp**](PoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolPredictions

> PoolPredictionsResp GetPoolPredictions(ctx, poolId).Execute()





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
	poolId := int64(789) // int64 | pool id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.GetPoolPredictions(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.GetPoolPredictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoolPredictions`: PoolPredictionsResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.GetPoolPredictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolPredictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoolPredictionsResp**](PoolPredictionsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolSamples

> PoolSamplesResp GetPoolSamples(ctx, poolId).ClientIo(clientIo).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	poolId := int64(789) // int64 | pool id
	clientIo := true // bool | merge client io stat for tier pool (optional)
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.GetPoolSamples(context.Background(), poolId).ClientIo(clientIo).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.GetPoolSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoolSamples`: PoolSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.GetPoolSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientIo** | **bool** | merge client io stat for tier pool | 
 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**PoolSamplesResp**](PoolSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolTopology

> PoolTopologyResp GetPoolTopology(ctx, poolId).Execute()





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
	poolId := int64(789) // int64 | pool id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.GetPoolTopology(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.GetPoolTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoolTopology`: PoolTopologyResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.GetPoolTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoolTopologyResp**](PoolTopologyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeEmptyPool

> PoolResp InitializeEmptyPool(ctx, poolId).Body(body).Execute()





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
	poolId := int64(789) // int64 | pool id
	body := *openapiclient.NewEmptyPoolInitializeReq() // EmptyPoolInitializeReq | the pool info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.InitializeEmptyPool(context.Background(), poolId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.InitializeEmptyPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitializeEmptyPool`: PoolResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.InitializeEmptyPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitializeEmptyPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**EmptyPoolInitializeReq**](EmptyPoolInitializeReq.md) | the pool info | 

### Return type

[**PoolResp**](PoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPools

> PoolsResp ListPools(ctx).Limit(limit).Offset(offset).All(all).ProtectionDomainId(protectionDomainId).ClusterId(clusterId).CompoundOsdOnly(compoundOsdOnly).OsdGroupId(osdGroupId).PoolType(poolType).PoolRole(poolRole).PoolMode(poolMode).Stretched(stretched).WithCompound(withCompound).IsCache(isCache).OsPolicyId(osPolicyId).StorageClassId(storageClassId).StorageClassPoolType(storageClassPoolType).OspDataBackendId(ospDataBackendId).OspClusterId(ospClusterId).Q(q).Sort(sort).PoolIds(poolIds).Execute()





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
	all := true // bool | show all pools (optional)
	protectionDomainId := int64(789) // int64 | protection domain id (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)
	compoundOsdOnly := true // bool | filter pool with only compound osds (optional)
	osdGroupId := int64(789) // int64 | osd group id (optional)
	poolType := "poolType_example" // string | filter pool by type (optional)
	poolRole := "poolRole_example" // string | filter pool by role (optional)
	poolMode := "poolMode_example" // string | filter pool by pool_mode (optional)
	stretched := true // bool | filter stretched pool (optional)
	withCompound := true // bool | with compound pool (optional)
	isCache := true // bool | list cache pool (optional)
	osPolicyId := int64(789) // int64 | filter data pool by object storage policy id (optional)
	storageClassId := int64(789) // int64 | filter data pool by os storage class id (optional)
	storageClassPoolType := "storageClassPoolType_example" // string | storage class pool type(active inactive) to query (optional)
	ospDataBackendId := int64(789) // int64 | osp data backend id (optional)
	ospClusterId := "ospClusterId_example" // string | osp cluster id (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	poolIds := int64(789) // int64 | filter pool by ids (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.ListPools(context.Background()).Limit(limit).Offset(offset).All(all).ProtectionDomainId(protectionDomainId).ClusterId(clusterId).CompoundOsdOnly(compoundOsdOnly).OsdGroupId(osdGroupId).PoolType(poolType).PoolRole(poolRole).PoolMode(poolMode).Stretched(stretched).WithCompound(withCompound).IsCache(isCache).OsPolicyId(osPolicyId).StorageClassId(storageClassId).StorageClassPoolType(storageClassPoolType).OspDataBackendId(ospDataBackendId).OspClusterId(ospClusterId).Q(q).Sort(sort).PoolIds(poolIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.ListPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPools`: PoolsResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.ListPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **all** | **bool** | show all pools | 
 **protectionDomainId** | **int64** | protection domain id | 
 **clusterId** | **string** | cluster id | 
 **compoundOsdOnly** | **bool** | filter pool with only compound osds | 
 **osdGroupId** | **int64** | osd group id | 
 **poolType** | **string** | filter pool by type | 
 **poolRole** | **string** | filter pool by role | 
 **poolMode** | **string** | filter pool by pool_mode | 
 **stretched** | **bool** | filter stretched pool | 
 **withCompound** | **bool** | with compound pool | 
 **isCache** | **bool** | list cache pool | 
 **osPolicyId** | **int64** | filter data pool by object storage policy id | 
 **storageClassId** | **int64** | filter data pool by os storage class id | 
 **storageClassPoolType** | **string** | storage class pool type(active inactive) to query | 
 **ospDataBackendId** | **int64** | osp data backend id | 
 **ospClusterId** | **string** | osp cluster id | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **poolIds** | **int64** | filter pool by ids | 

### Return type

[**PoolsResp**](PoolsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFromOutsideBackend

> OspPoolOutsideResp RemoveFromOutsideBackend(ctx).Body(body).Execute()





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
	body := *openapiclient.NewOspPoolOutsideReq() // OspPoolOutsideReq | pool info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.RemoveFromOutsideBackend(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.RemoveFromOutsideBackend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveFromOutsideBackend`: OspPoolOutsideResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.RemoveFromOutsideBackend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFromOutsideBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OspPoolOutsideReq**](OspPoolOutsideReq.md) | pool info | 

### Return type

[**OspPoolOutsideResp**](OspPoolOutsideResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOsdsFromPool

> PoolResp RemoveOsdsFromPool(ctx, poolId).Body(body).Execute()





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
	poolId := int64(789) // int64 | pool id
	body := *openapiclient.NewOsdsRemoveReq([]int64{int64(123)}) // OsdsRemoveReq | osd infos

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.RemoveOsdsFromPool(context.Background(), poolId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.RemoveOsdsFromPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveOsdsFromPool`: PoolResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.RemoveOsdsFromPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOsdsFromPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OsdsRemoveReq**](OsdsRemoveReq.md) | osd infos | 

### Return type

[**PoolResp**](PoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePoolsFromOspDataBackend

> TaskIDResp RemovePoolsFromOspDataBackend(ctx).Body(body).Execute()





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
	body := *openapiclient.NewOspDataBackendPoolsOperateReq() // OspDataBackendPoolsOperateReq | pool info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.RemovePoolsFromOspDataBackend(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.RemovePoolsFromOspDataBackend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemovePoolsFromOspDataBackend`: TaskIDResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.RemovePoolsFromOspDataBackend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemovePoolsFromOspDataBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OspDataBackendPoolsOperateReq**](OspDataBackendPoolsOperateReq.md) | pool info | 

### Return type

[**TaskIDResp**](TaskIDResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReweightPool

> PoolResp ReweightPool(ctx, poolId).Execute()





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
	poolId := int64(789) // int64 | pool id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.ReweightPool(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.ReweightPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReweightPool`: PoolResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.ReweightPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReweightPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoolResp**](PoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwitchPoolRole

> PoolResp SwitchPoolRole(ctx, poolId).Execute()





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
	poolId := int64(789) // int64 | pool id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.SwitchPoolRole(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.SwitchPoolRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SwitchPoolRole`: PoolResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.SwitchPoolRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchPoolRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoolResp**](PoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateECPoolCrushRule

> PoolResp UpdateECPoolCrushRule(ctx, poolId).Body(body).Execute()





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
	poolId := int64(789) // int64 | pool id
	body := *openapiclient.NewECPoolUpdateCrushRuleReq(*openapiclient.NewECPoolUpdateCrushRuleReqPool(int64(123), int64(123))) // ECPoolUpdateCrushRuleReq | crush rule info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.UpdateECPoolCrushRule(context.Background(), poolId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.UpdateECPoolCrushRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateECPoolCrushRule`: PoolResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.UpdateECPoolCrushRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateECPoolCrushRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ECPoolUpdateCrushRuleReq**](ECPoolUpdateCrushRuleReq.md) | crush rule info | 

### Return type

[**PoolResp**](PoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePool

> PoolResp UpdatePool(ctx, poolId).Body(body).Execute()





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
	poolId := int64(789) // int64 | pool id
	body := *openapiclient.NewPoolUpdateReq(*openapiclient.NewPoolUpdateReqPool()) // PoolUpdateReq | pool info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.UpdatePool(context.Background(), poolId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.UpdatePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePool`: PoolResp
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.UpdatePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PoolUpdateReq**](PoolUpdateReq.md) | pool info | 

### Return type

[**PoolResp**](PoolResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePoolGCPolicyPlan

> UpdatePoolGCPolicyPlan(ctx, poolId).Body(body).Execute()





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
	poolId := int64(789) // int64 | pool id
	body := *openapiclient.NewPoolGCPolicyPlan([]openapiclient.GcPlan{*openapiclient.NewGcPlan([]int64{int64(123)}, []openapiclient.GcPolicy{*openapiclient.NewGcPolicy(false)})}) // PoolGCPolicyPlan | pool gc policy plan

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoolsAPI.UpdatePoolGCPolicyPlan(context.Background(), poolId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.UpdatePoolGCPolicyPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePoolGCPolicyPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PoolGCPolicyPlan**](PoolGCPolicyPlan.md) | pool gc policy plan | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

