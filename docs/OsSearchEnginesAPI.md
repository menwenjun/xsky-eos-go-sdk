# \OsSearchEnginesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOSSearchGateways**](OsSearchEnginesAPI.md#AddOSSearchGateways) | **Post** /os-search-engines/{os_search_engine_id}:add-os-search-gateways | 
[**ChangeOSSearchEngine**](OsSearchEnginesAPI.md#ChangeOSSearchEngine) | **Patch** /os-search-engines/{os_search_engine_id} | 
[**CreateOSSearchEngine**](OsSearchEnginesAPI.md#CreateOSSearchEngine) | **Post** /os-search-engines/ | 
[**DeleteOSSearchEngine**](OsSearchEnginesAPI.md#DeleteOSSearchEngine) | **Delete** /os-search-engines/{os_search_engine_id} | 
[**GetOSSearchEngine**](OsSearchEnginesAPI.md#GetOSSearchEngine) | **Get** /os-search-engines/{os_search_engine_id} | 
[**GetOSSearchEngineSamples**](OsSearchEnginesAPI.md#GetOSSearchEngineSamples) | **Get** /os-search-engines/{os_search_engine_id}/samples | 
[**ListOSSearchEngines**](OsSearchEnginesAPI.md#ListOSSearchEngines) | **Get** /os-search-engines/ | 
[**RemoveOSSearchGateways**](OsSearchEnginesAPI.md#RemoveOSSearchGateways) | **Post** /os-search-engines/{os_search_engine_id}:remove-os-search-gateways | 
[**StartOSSearchEngine**](OsSearchEnginesAPI.md#StartOSSearchEngine) | **Post** /os-search-engines/{os_search_engine_id}:start | 
[**StopOSSearchEngine**](OsSearchEnginesAPI.md#StopOSSearchEngine) | **Post** /os-search-engines/{os_search_engine_id}:stop | 



## AddOSSearchGateways

> OSSearchEngineResp AddOSSearchGateways(ctx, osSearchEngineId).Body(body).Execute()





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
	osSearchEngineId := int64(789) // int64 | OS search engine id
	body := *openapiclient.NewOSSearchEngineAddGatewaysReq(*openapiclient.NewOSSearchEngineAddGatewaysReqSearchEngine([]openapiclient.OSSearchGatewayReq{*openapiclient.NewOSSearchGatewayReq(int64(123), int64(123))})) // OSSearchEngineAddGatewaysReq | os search gateways info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSearchEnginesAPI.AddOSSearchGateways(context.Background(), osSearchEngineId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSearchEnginesAPI.AddOSSearchGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOSSearchGateways`: OSSearchEngineResp
	fmt.Fprintf(os.Stdout, "Response from `OsSearchEnginesAPI.AddOSSearchGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osSearchEngineId** | **int64** | OS search engine id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOSSearchGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSSearchEngineAddGatewaysReq**](OSSearchEngineAddGatewaysReq.md) | os search gateways info | 

### Return type

[**OSSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeOSSearchEngine

> OSSearchEngineResp ChangeOSSearchEngine(ctx, osSearchEngineId).Body(body).Execute()





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
	osSearchEngineId := int64(789) // int64 | OS search engine id
	body := *openapiclient.NewOSSearchEngineUpdateReq(*openapiclient.NewOSSearchEngineUpdateReqSearchEngine()) // OSSearchEngineUpdateReq | os search gateways info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSearchEnginesAPI.ChangeOSSearchEngine(context.Background(), osSearchEngineId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSearchEnginesAPI.ChangeOSSearchEngine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeOSSearchEngine`: OSSearchEngineResp
	fmt.Fprintf(os.Stdout, "Response from `OsSearchEnginesAPI.ChangeOSSearchEngine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osSearchEngineId** | **int64** | OS search engine id | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeOSSearchEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSSearchEngineUpdateReq**](OSSearchEngineUpdateReq.md) | os search gateways info | 

### Return type

[**OSSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOSSearchEngine

> OSSearchEngineResp CreateOSSearchEngine(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewOSSearchEngineCreateReq(*openapiclient.NewOSSearchEngineCreateReqSearchEngine(int64(123), int64(123), []openapiclient.OSSearchGatewayReq{*openapiclient.NewOSSearchGatewayReq(int64(123), int64(123))})) // OSSearchEngineCreateReq | OS search engine info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSearchEnginesAPI.CreateOSSearchEngine(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSearchEnginesAPI.CreateOSSearchEngine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOSSearchEngine`: OSSearchEngineResp
	fmt.Fprintf(os.Stdout, "Response from `OsSearchEnginesAPI.CreateOSSearchEngine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOSSearchEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OSSearchEngineCreateReq**](OSSearchEngineCreateReq.md) | OS search engine info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOSSearchEngine

> OSSearchEngineResp DeleteOSSearchEngine(ctx, osSearchEngineId).Execute()





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
	osSearchEngineId := int64(789) // int64 | OS search engine id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSearchEnginesAPI.DeleteOSSearchEngine(context.Background(), osSearchEngineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSearchEnginesAPI.DeleteOSSearchEngine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOSSearchEngine`: OSSearchEngineResp
	fmt.Fprintf(os.Stdout, "Response from `OsSearchEnginesAPI.DeleteOSSearchEngine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osSearchEngineId** | **int64** | OS search engine id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOSSearchEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSSearchEngine

> OSSearchEngineResp GetOSSearchEngine(ctx, osSearchEngineId).Execute()





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
	osSearchEngineId := int64(789) // int64 | OS search engine id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSearchEnginesAPI.GetOSSearchEngine(context.Background(), osSearchEngineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSearchEnginesAPI.GetOSSearchEngine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSSearchEngine`: OSSearchEngineResp
	fmt.Fprintf(os.Stdout, "Response from `OsSearchEnginesAPI.GetOSSearchEngine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osSearchEngineId** | **int64** | OS search engine id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSSearchEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOSSearchEngineSamples

> OSSearchEngineSamplesResp GetOSSearchEngineSamples(ctx, osSearchEngineId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	osSearchEngineId := int64(789) // int64 | OS search engine id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSearchEnginesAPI.GetOSSearchEngineSamples(context.Background(), osSearchEngineId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSearchEnginesAPI.GetOSSearchEngineSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSSearchEngineSamples`: OSSearchEngineSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `OsSearchEnginesAPI.GetOSSearchEngineSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osSearchEngineId** | **int64** | OS search engine id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSSearchEngineSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**OSSearchEngineSamplesResp**](OSSearchEngineSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOSSearchEngines

> OSSearchEnginesResp ListOSSearchEngines(ctx).ClusterId(clusterId).Limit(limit).Offset(offset).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSearchEnginesAPI.ListOSSearchEngines(context.Background()).ClusterId(clusterId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSearchEnginesAPI.ListOSSearchEngines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSSearchEngines`: OSSearchEnginesResp
	fmt.Fprintf(os.Stdout, "Response from `OsSearchEnginesAPI.ListOSSearchEngines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSSearchEnginesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**OSSearchEnginesResp**](OSSearchEnginesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOSSearchGateways

> OSSearchEngineResp RemoveOSSearchGateways(ctx, osSearchEngineId).Body(body).Execute()





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
	osSearchEngineId := int64(789) // int64 | OS search engine id
	body := *openapiclient.NewOSSearchEngineRemoveGatewaysReq(*openapiclient.NewOSSearchEngineRemoveGatewaysReqSearchEngine([]int64{int64(123)})) // OSSearchEngineRemoveGatewaysReq | os search gateways info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSearchEnginesAPI.RemoveOSSearchGateways(context.Background(), osSearchEngineId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSearchEnginesAPI.RemoveOSSearchGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveOSSearchGateways`: OSSearchEngineResp
	fmt.Fprintf(os.Stdout, "Response from `OsSearchEnginesAPI.RemoveOSSearchGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osSearchEngineId** | **int64** | OS search engine id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOSSearchGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OSSearchEngineRemoveGatewaysReq**](OSSearchEngineRemoveGatewaysReq.md) | os search gateways info | 

### Return type

[**OSSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartOSSearchEngine

> OSSearchEngineResp StartOSSearchEngine(ctx, osSearchEngineId).Execute()





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
	osSearchEngineId := int64(789) // int64 | OS search engine id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSearchEnginesAPI.StartOSSearchEngine(context.Background(), osSearchEngineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSearchEnginesAPI.StartOSSearchEngine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartOSSearchEngine`: OSSearchEngineResp
	fmt.Fprintf(os.Stdout, "Response from `OsSearchEnginesAPI.StartOSSearchEngine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osSearchEngineId** | **int64** | OS search engine id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartOSSearchEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopOSSearchEngine

> OSSearchEngineResp StopOSSearchEngine(ctx, osSearchEngineId).Execute()





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
	osSearchEngineId := int64(789) // int64 | OS search engine id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsSearchEnginesAPI.StopOSSearchEngine(context.Background(), osSearchEngineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsSearchEnginesAPI.StopOSSearchEngine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopOSSearchEngine`: OSSearchEngineResp
	fmt.Fprintf(os.Stdout, "Response from `OsSearchEnginesAPI.StopOSSearchEngine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osSearchEngineId** | **int64** | OS search engine id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopOSSearchEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

