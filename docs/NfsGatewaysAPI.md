# \NfsGatewaysAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNFSGateway**](NfsGatewaysAPI.md#CreateNFSGateway) | **Post** /nfs-gateways/ | 
[**CreateNFSGatewayBucketMap**](NfsGatewaysAPI.md#CreateNFSGatewayBucketMap) | **Put** /nfs-gateways/{gateway_id}/buckets/{bucket_id} | 
[**DeleteNFSGateway**](NfsGatewaysAPI.md#DeleteNFSGateway) | **Delete** /nfs-gateways/{gateway_id} | 
[**DeleteNFSGatewayBucketMap**](NfsGatewaysAPI.md#DeleteNFSGatewayBucketMap) | **Delete** /nfs-gateways/{gateway_id}/buckets/{bucket_id} | 
[**DoNFSGateway**](NfsGatewaysAPI.md#DoNFSGateway) | **Put** /nfs-gateways/{gateway_id} | 
[**GetNFSGateway**](NfsGatewaysAPI.md#GetNFSGateway) | **Get** /nfs-gateways/{gateway_id} | 
[**GetNFSGatewayBucketMap**](NfsGatewaysAPI.md#GetNFSGatewayBucketMap) | **Get** /nfs-gateways/{gateway_id}/buckets/{bucket_id} | 
[**GetNFSGatewaySamples**](NfsGatewaysAPI.md#GetNFSGatewaySamples) | **Get** /nfs-gateways/{gateway_id}/samples | 
[**ListNFSGatewayBucketMaps**](NfsGatewaysAPI.md#ListNFSGatewayBucketMaps) | **Get** /nfs-gateways/{gateway_id}/buckets | 
[**ListNFSGateways**](NfsGatewaysAPI.md#ListNFSGateways) | **Get** /nfs-gateways/ | 
[**UpdateNFSGateway**](NfsGatewaysAPI.md#UpdateNFSGateway) | **Patch** /nfs-gateways/{gateway_id} | 
[**UpdateNFSGatewayBucketMap**](NfsGatewaysAPI.md#UpdateNFSGatewayBucketMap) | **Patch** /nfs-gateways/{gateway_id}/buckets/{bucket_id} | 
[**UpdateOspExportConfig**](NfsGatewaysAPI.md#UpdateOspExportConfig) | **Post** /nfs-gateways/{gateway_id}:update-osp-export-config | 



## CreateNFSGateway

> NFSGatewayResp CreateNFSGateway(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewNFSGatewayCreateReq(*openapiclient.NewNFSGatewayCreateReqNFSGateway("Description_example", int64(123), "Name_example", int64(123))) // NFSGatewayCreateReq | nfs gateway info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsGatewaysAPI.CreateNFSGateway(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsGatewaysAPI.CreateNFSGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNFSGateway`: NFSGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `NfsGatewaysAPI.CreateNFSGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNFSGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NFSGatewayCreateReq**](NFSGatewayCreateReq.md) | nfs gateway info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**NFSGatewayResp**](NFSGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNFSGatewayBucketMap

> NFSGatewayBucketMapResp CreateNFSGatewayBucketMap(ctx, gatewayId, bucketId).ClusterId(clusterId).Execute()





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
	gatewayId := int64(789) // int64 | nfs gateway id
	bucketId := int64(789) // int64 | bucket id
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsGatewaysAPI.CreateNFSGatewayBucketMap(context.Background(), gatewayId, bucketId).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsGatewaysAPI.CreateNFSGatewayBucketMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNFSGatewayBucketMap`: NFSGatewayBucketMapResp
	fmt.Fprintf(os.Stdout, "Response from `NfsGatewaysAPI.CreateNFSGatewayBucketMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | nfs gateway id | 
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNFSGatewayBucketMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterId** | **string** | cluster id | 

### Return type

[**NFSGatewayBucketMapResp**](NFSGatewayBucketMapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNFSGateway

> NFSGatewayResp DeleteNFSGateway(ctx, gatewayId).Force(force).Execute()





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
	gatewayId := int64(789) // int64 | nfs gateway id
	force := true // bool | force delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsGatewaysAPI.DeleteNFSGateway(context.Background(), gatewayId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsGatewaysAPI.DeleteNFSGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNFSGateway`: NFSGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `NfsGatewaysAPI.DeleteNFSGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | nfs gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNFSGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 

### Return type

[**NFSGatewayResp**](NFSGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNFSGatewayBucketMap

> NFSGatewayBucketMapResp DeleteNFSGatewayBucketMap(ctx, gatewayId, bucketId).Force(force).Execute()





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
	gatewayId := int64(789) // int64 | nfs gateway id
	bucketId := int64(789) // int64 | bucket id
	force := true // bool | force delete or no (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsGatewaysAPI.DeleteNFSGatewayBucketMap(context.Background(), gatewayId, bucketId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsGatewaysAPI.DeleteNFSGatewayBucketMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNFSGatewayBucketMap`: NFSGatewayBucketMapResp
	fmt.Fprintf(os.Stdout, "Response from `NfsGatewaysAPI.DeleteNFSGatewayBucketMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | nfs gateway id | 
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNFSGatewayBucketMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | force delete or no | 

### Return type

[**NFSGatewayBucketMapResp**](NFSGatewayBucketMapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DoNFSGateway

> NFSGatewayResp DoNFSGateway(ctx, gatewayId).Body(body).Force(force).Execute()





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
	gatewayId := int64(789) // int64 | nfs gateway id
	body := *openapiclient.NewNFSGatewayActionReq(*openapiclient.NewNFSGatewayActionReqNFSGateway("Action_example")) // NFSGatewayActionReq | nfs gateway action info
	force := true // bool | force stop or no (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsGatewaysAPI.DoNFSGateway(context.Background(), gatewayId).Body(body).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsGatewaysAPI.DoNFSGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoNFSGateway`: NFSGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `NfsGatewaysAPI.DoNFSGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | nfs gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoNFSGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NFSGatewayActionReq**](NFSGatewayActionReq.md) | nfs gateway action info | 
 **force** | **bool** | force stop or no | 

### Return type

[**NFSGatewayResp**](NFSGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNFSGateway

> NFSGatewayResp GetNFSGateway(ctx, gatewayId).Execute()





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
	gatewayId := int64(789) // int64 | nfs gateway id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsGatewaysAPI.GetNFSGateway(context.Background(), gatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsGatewaysAPI.GetNFSGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNFSGateway`: NFSGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `NfsGatewaysAPI.GetNFSGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | nfs gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNFSGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NFSGatewayResp**](NFSGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNFSGatewayBucketMap

> NFSGatewayBucketMapResp GetNFSGatewayBucketMap(ctx, gatewayId, bucketId).Execute()





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
	gatewayId := int64(789) // int64 | nfs gateway id
	bucketId := int64(789) // int64 | bucket id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsGatewaysAPI.GetNFSGatewayBucketMap(context.Background(), gatewayId, bucketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsGatewaysAPI.GetNFSGatewayBucketMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNFSGatewayBucketMap`: NFSGatewayBucketMapResp
	fmt.Fprintf(os.Stdout, "Response from `NfsGatewaysAPI.GetNFSGatewayBucketMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | nfs gateway id | 
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNFSGatewayBucketMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NFSGatewayBucketMapResp**](NFSGatewayBucketMapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNFSGatewaySamples

> NFSGatewaySamplesResp GetNFSGatewaySamples(ctx, gatewayId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	gatewayId := int64(789) // int64 | gateway id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsGatewaysAPI.GetNFSGatewaySamples(context.Background(), gatewayId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsGatewaysAPI.GetNFSGatewaySamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNFSGatewaySamples`: NFSGatewaySamplesResp
	fmt.Fprintf(os.Stdout, "Response from `NfsGatewaysAPI.GetNFSGatewaySamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNFSGatewaySamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**NFSGatewaySamplesResp**](NFSGatewaySamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNFSGatewayBucketMaps

> NFSGatewayBucketMapsResp ListNFSGatewayBucketMaps(ctx, gatewayId).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()





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
	gatewayId := int64(789) // int64 | nfs gateway id
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsGatewaysAPI.ListNFSGatewayBucketMaps(context.Background(), gatewayId).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsGatewaysAPI.ListNFSGatewayBucketMaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNFSGatewayBucketMaps`: NFSGatewayBucketMapsResp
	fmt.Fprintf(os.Stdout, "Response from `NfsGatewaysAPI.ListNFSGatewayBucketMaps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | nfs gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNFSGatewayBucketMapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 

### Return type

[**NFSGatewayBucketMapsResp**](NFSGatewayBucketMapsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNFSGateways

> NFSGatewaysResp ListNFSGateways(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).OspZoneId(ospZoneId).ClusterId(clusterId).Execute()





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
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	ospZoneId := int64(789) // int64 | osp zone id (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsGatewaysAPI.ListNFSGateways(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).OspZoneId(ospZoneId).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsGatewaysAPI.ListNFSGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNFSGateways`: NFSGatewaysResp
	fmt.Fprintf(os.Stdout, "Response from `NfsGatewaysAPI.ListNFSGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNFSGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **ospZoneId** | **int64** | osp zone id | 
 **clusterId** | **string** | cluster id | 

### Return type

[**NFSGatewaysResp**](NFSGatewaysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNFSGateway

> NFSGatewayResp UpdateNFSGateway(ctx, gatewayId).Body(body).Execute()





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
	gatewayId := int64(789) // int64 | nfs gateway id
	body := *openapiclient.NewNFSGatewayUpdateReq(*openapiclient.NewNFSGatewayUpdateReqNFSGateway()) // NFSGatewayUpdateReq | nfs gateway info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsGatewaysAPI.UpdateNFSGateway(context.Background(), gatewayId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsGatewaysAPI.UpdateNFSGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNFSGateway`: NFSGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `NfsGatewaysAPI.UpdateNFSGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | nfs gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNFSGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NFSGatewayUpdateReq**](NFSGatewayUpdateReq.md) | nfs gateway info | 

### Return type

[**NFSGatewayResp**](NFSGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNFSGatewayBucketMap

> NFSGatewayBucketMapResp UpdateNFSGatewayBucketMap(ctx, gatewayId, bucketId).Body(body).Force(force).Execute()





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
	gatewayId := int64(789) // int64 | nfs gateway id
	bucketId := int64(789) // int64 | bucket id
	body := *openapiclient.NewNFSGatewayBucketMapUpdateReq(*openapiclient.NewNFSGatewayBucketMapUpdateReqNFSGatewayBucketMap()) // NFSGatewayBucketMapUpdateReq | nfs gateway bucket update info
	force := true // bool | force update bucket map (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsGatewaysAPI.UpdateNFSGatewayBucketMap(context.Background(), gatewayId, bucketId).Body(body).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsGatewaysAPI.UpdateNFSGatewayBucketMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNFSGatewayBucketMap`: NFSGatewayBucketMapResp
	fmt.Fprintf(os.Stdout, "Response from `NfsGatewaysAPI.UpdateNFSGatewayBucketMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | nfs gateway id | 
**bucketId** | **int64** | bucket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNFSGatewayBucketMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**NFSGatewayBucketMapUpdateReq**](NFSGatewayBucketMapUpdateReq.md) | nfs gateway bucket update info | 
 **force** | **bool** | force update bucket map | 

### Return type

[**NFSGatewayBucketMapResp**](NFSGatewayBucketMapResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOspExportConfig

> NFSGatewayResp UpdateOspExportConfig(ctx, gatewayId).Execute()





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
	gatewayId := int64(789) // int64 | nfs gateway id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsGatewaysAPI.UpdateOspExportConfig(context.Background(), gatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsGatewaysAPI.UpdateOspExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOspExportConfig`: NFSGatewayResp
	fmt.Fprintf(os.Stdout, "Response from `NfsGatewaysAPI.UpdateOspExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int64** | nfs gateway id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOspExportConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NFSGatewayResp**](NFSGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

