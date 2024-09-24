# \OspGatewaysAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOspGateway**](OspGatewaysAPI.md#CreateOspGateway) | **Post** /osp-gateways/ | 
[**DeleteOspGateway**](OspGatewaysAPI.md#DeleteOspGateway) | **Delete** /osp-gateways/{osp_gateway_id} | 
[**GetOspGateway**](OspGatewaysAPI.md#GetOspGateway) | **Get** /osp-gateways/{osp_gateway_id} | 
[**GetOspGatewaySamples**](OspGatewaysAPI.md#GetOspGatewaySamples) | **Get** /osp-gateways/{osp_gateway_id}/samples | 
[**GetOspGatewaysStatSumByZone**](OspGatewaysAPI.md#GetOspGatewaysStatSumByZone) | **Get** /osp-gateways/osp_zone_stats | 
[**ListOspGateways**](OspGatewaysAPI.md#ListOspGateways) | **Get** /osp-gateways/ | 
[**RestartOspGateway**](OspGatewaysAPI.md#RestartOspGateway) | **Post** /osp-gateways/{osp_gateway_id}:restart | 
[**RestartOspGateways**](OspGatewaysAPI.md#RestartOspGateways) | **Post** /osp-gateways/:restart | 
[**StartOspGateway**](OspGatewaysAPI.md#StartOspGateway) | **Post** /osp-gateways/{osp_gateway_id}:start | 
[**StopOspGateway**](OspGatewaysAPI.md#StopOspGateway) | **Post** /osp-gateways/{osp_gateway_id}:stop | 



## CreateOspGateway

> OspGatewayResp CreateOspGateway(ctx).Body(body).Execute()





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
	body := *openapiclient.NewOspGatewayCreateReq() // OspGatewayCreateReq | gateway info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspGatewaysAPI.CreateOspGateway(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspGatewaysAPI.CreateOspGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOspGateway`: OspGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `OspGatewaysAPI.CreateOspGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOspGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OspGatewayCreateReq**](OspGatewayCreateReq.md) | gateway info | 

### Return type

[**OspGatewayResp**](OspGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOspGateway

> OspGatewayResp DeleteOspGateway(ctx, ospGatewayId).Execute()





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
	ospGatewayId := "ospGatewayId_example" // string | osp gateway id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspGatewaysAPI.DeleteOspGateway(context.Background(), ospGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspGatewaysAPI.DeleteOspGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOspGateway`: OspGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `OspGatewaysAPI.DeleteOspGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ospGatewayId** | **string** | osp gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOspGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OspGatewayResp**](OspGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOspGateway

> OspGatewayResp GetOspGateway(ctx, ospGatewayId).Execute()





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
	ospGatewayId := int64(789) // int64 | osp gateway id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspGatewaysAPI.GetOspGateway(context.Background(), ospGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspGatewaysAPI.GetOspGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOspGateway`: OspGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `OspGatewaysAPI.GetOspGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ospGatewayId** | **int64** | osp gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOspGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OspGatewayResp**](OspGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOspGatewaySamples

> OspGatewaySamplesResp GetOspGatewaySamples(ctx, ospGatewayId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	ospGatewayId := int64(789) // int64 | osp gateway id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspGatewaysAPI.GetOspGatewaySamples(context.Background(), ospGatewayId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspGatewaysAPI.GetOspGatewaySamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOspGatewaySamples`: OspGatewaySamplesResp
	fmt.Fprintf(os.Stdout, "Response from `OspGatewaysAPI.GetOspGatewaySamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ospGatewayId** | **int64** | osp gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOspGatewaySamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**OspGatewaySamplesResp**](OspGatewaySamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOspGatewaysStatSumByZone

> OspZoneStatResp GetOspGatewaysStatSumByZone(ctx).OspZoneId(ospZoneId).ClusterId(clusterId).Execute()





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
	ospZoneId := int64(789) // int64 | osp zone id
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspGatewaysAPI.GetOspGatewaysStatSumByZone(context.Background()).OspZoneId(ospZoneId).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspGatewaysAPI.GetOspGatewaysStatSumByZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOspGatewaysStatSumByZone`: OspZoneStatResp
	fmt.Fprintf(os.Stdout, "Response from `OspGatewaysAPI.GetOspGatewaysStatSumByZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOspGatewaysStatSumByZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ospZoneId** | **int64** | osp zone id | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OspZoneStatResp**](OspZoneStatResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOspGateways

> OspGatewaysResp ListOspGateways(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).OspZoneId(ospZoneId).GatewayName(gatewayName).Role(role).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	ospZoneId := int64(789) // int64 | osp zone id (optional)
	gatewayName := "gatewayName_example" // string | gateway name (optional)
	role := "role_example" // string | filter by gateway role (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspGatewaysAPI.ListOspGateways(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).OspZoneId(ospZoneId).GatewayName(gatewayName).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspGatewaysAPI.ListOspGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOspGateways`: OspGatewaysResp
	fmt.Fprintf(os.Stdout, "Response from `OspGatewaysAPI.ListOspGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOspGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **ospZoneId** | **int64** | osp zone id | 
 **gatewayName** | **string** | gateway name | 
 **role** | **string** | filter by gateway role | 

### Return type

[**OspGatewaysResp**](OspGatewaysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartOspGateway

> OspGatewayResp RestartOspGateway(ctx, ospGatewayId).Execute()





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
	ospGatewayId := int64(789) // int64 | osp gateway id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspGatewaysAPI.RestartOspGateway(context.Background(), ospGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspGatewaysAPI.RestartOspGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestartOspGateway`: OspGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `OspGatewaysAPI.RestartOspGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ospGatewayId** | **int64** | osp gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartOspGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OspGatewayResp**](OspGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartOspGateways

> RestartOspGateways(ctx).ClusterId(clusterId).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OspGatewaysAPI.RestartOspGateways(context.Background()).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspGatewaysAPI.RestartOspGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestartOspGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 

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


## StartOspGateway

> OspGatewayResp StartOspGateway(ctx, ospGatewayId).Execute()





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
	ospGatewayId := int64(789) // int64 | osp gateway id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspGatewaysAPI.StartOspGateway(context.Background(), ospGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspGatewaysAPI.StartOspGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartOspGateway`: OspGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `OspGatewaysAPI.StartOspGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ospGatewayId** | **int64** | osp gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartOspGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OspGatewayResp**](OspGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopOspGateway

> OspGatewayResp StopOspGateway(ctx, ospGatewayId).Execute()





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
	ospGatewayId := "ospGatewayId_example" // string | osp gateway id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspGatewaysAPI.StopOspGateway(context.Background(), ospGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspGatewaysAPI.StopOspGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopOspGateway`: OspGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `OspGatewaysAPI.StopOspGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ospGatewayId** | **string** | osp gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopOspGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OspGatewayResp**](OspGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

