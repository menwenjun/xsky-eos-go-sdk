# \MetadataClustersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMetadataServicesToCluster**](MetadataClustersAPI.md#AddMetadataServicesToCluster) | **Post** /metadata-clusters/{metadata_cluster_id}:add-metadata-services | 
[**CreateMetadataCluster**](MetadataClustersAPI.md#CreateMetadataCluster) | **Post** /metadata-clusters/ | 
[**DeleteMetadataCluster**](MetadataClustersAPI.md#DeleteMetadataCluster) | **Delete** /metadata-clusters/{metadata_cluster_id} | 
[**GetMetadataCluster**](MetadataClustersAPI.md#GetMetadataCluster) | **Get** /metadata-clusters/{metadata_cluster_id} | 
[**GetMetadataClusterPredictions**](MetadataClustersAPI.md#GetMetadataClusterPredictions) | **Get** /metadata-clusters/{metadata_cluster_id}/predictions | 
[**GetMetadataClusterSamples**](MetadataClustersAPI.md#GetMetadataClusterSamples) | **Get** /metadata-clusters/{metadata_cluster_id}/samples | 
[**ListMetadataClusters**](MetadataClustersAPI.md#ListMetadataClusters) | **Get** /metadata-clusters/ | 
[**RemoveMetadataServicesFromCluster**](MetadataClustersAPI.md#RemoveMetadataServicesFromCluster) | **Post** /metadata-clusters/{metadata_cluster_id}:remove-metadata-services | 
[**SetMetadataClusterToStretched**](MetadataClustersAPI.md#SetMetadataClusterToStretched) | **Patch** /metadata-clusters/{metadata_cluster_id}:set-stretched | 
[**SetMetadataClusterTransLocator**](MetadataClustersAPI.md#SetMetadataClusterTransLocator) | **Post** /metadata-clusters/{metadata_cluster_id}:set-trans-locator | 
[**UpdateMetadataCluster**](MetadataClustersAPI.md#UpdateMetadataCluster) | **Patch** /metadata-clusters/{metadata_cluster_id} | 
[**UpdatePrimaryDc**](MetadataClustersAPI.md#UpdatePrimaryDc) | **Patch** /metadata-clusters/{metadata_cluster_id}:update-primary-dc | 



## AddMetadataServicesToCluster

> MetadataClusterResp AddMetadataServicesToCluster(ctx, metadataClusterId).Body(body).Execute()





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
	metadataClusterId := int64(789) // int64 | metadata cluster id
	body := *openapiclient.NewMetadataClusterAddServicesReq() // MetadataClusterAddServicesReq | metadata services

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataClustersAPI.AddMetadataServicesToCluster(context.Background(), metadataClusterId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataClustersAPI.AddMetadataServicesToCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddMetadataServicesToCluster`: MetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `MetadataClustersAPI.AddMetadataServicesToCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataClusterId** | **int64** | metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMetadataServicesToClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MetadataClusterAddServicesReq**](MetadataClusterAddServicesReq.md) | metadata services | 

### Return type

[**MetadataClusterResp**](MetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMetadataCluster

> MetadataClusterResp CreateMetadataCluster(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewMetadataClusterCreateReq() // MetadataClusterCreateReq | metadata cluster info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataClustersAPI.CreateMetadataCluster(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataClustersAPI.CreateMetadataCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMetadataCluster`: MetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `MetadataClustersAPI.CreateMetadataCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetadataClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MetadataClusterCreateReq**](MetadataClusterCreateReq.md) | metadata cluster info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**MetadataClusterResp**](MetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMetadataCluster

> MetadataClusterResp DeleteMetadataCluster(ctx, metadataClusterId).Execute()





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
	metadataClusterId := int64(789) // int64 | metadata cluster id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataClustersAPI.DeleteMetadataCluster(context.Background(), metadataClusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataClustersAPI.DeleteMetadataCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMetadataCluster`: MetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `MetadataClustersAPI.DeleteMetadataCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataClusterId** | **int64** | metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetadataClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetadataClusterResp**](MetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadataCluster

> MetadataClusterResp GetMetadataCluster(ctx, metadataClusterId).Execute()





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
	metadataClusterId := int64(789) // int64 | metadata cluster id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataClustersAPI.GetMetadataCluster(context.Background(), metadataClusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataClustersAPI.GetMetadataCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetadataCluster`: MetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `MetadataClustersAPI.GetMetadataCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataClusterId** | **int64** | metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetadataClusterResp**](MetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadataClusterPredictions

> MetadataClusterPredictionsResp GetMetadataClusterPredictions(ctx, metadataClusterId).Execute()





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
	metadataClusterId := int64(789) // int64 | metadata cluster id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataClustersAPI.GetMetadataClusterPredictions(context.Background(), metadataClusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataClustersAPI.GetMetadataClusterPredictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetadataClusterPredictions`: MetadataClusterPredictionsResp
	fmt.Fprintf(os.Stdout, "Response from `MetadataClustersAPI.GetMetadataClusterPredictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataClusterId** | **int64** | metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataClusterPredictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetadataClusterPredictionsResp**](MetadataClusterPredictionsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadataClusterSamples

> MetadataClusterSamplesResp GetMetadataClusterSamples(ctx, metadataClusterId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	metadataClusterId := int64(789) // int64 | metadata cluster id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataClustersAPI.GetMetadataClusterSamples(context.Background(), metadataClusterId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataClustersAPI.GetMetadataClusterSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetadataClusterSamples`: MetadataClusterSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `MetadataClustersAPI.GetMetadataClusterSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataClusterId** | **int64** | metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataClusterSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**MetadataClusterSamplesResp**](MetadataClusterSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetadataClusters

> MetadataClustersResp ListMetadataClusters(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).HostId(hostId).MetadataClusterId(metadataClusterId).Q(q).Sort(sort).Execute()





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
	hostId := int64(789) // int64 | host id (optional)
	metadataClusterId := int64(789) // int64 | metadata cluster id (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataClustersAPI.ListMetadataClusters(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).HostId(hostId).MetadataClusterId(metadataClusterId).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataClustersAPI.ListMetadataClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetadataClusters`: MetadataClustersResp
	fmt.Fprintf(os.Stdout, "Response from `MetadataClustersAPI.ListMetadataClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetadataClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **hostId** | **int64** | host id | 
 **metadataClusterId** | **int64** | metadata cluster id | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**MetadataClustersResp**](MetadataClustersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMetadataServicesFromCluster

> MetadataClusterResp RemoveMetadataServicesFromCluster(ctx, metadataClusterId).Body(body).Execute()





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
	metadataClusterId := int64(789) // int64 | metadata cluster id
	body := *openapiclient.NewMetadataClusterRemoveServicesReq() // MetadataClusterRemoveServicesReq | metadata services

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataClustersAPI.RemoveMetadataServicesFromCluster(context.Background(), metadataClusterId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataClustersAPI.RemoveMetadataServicesFromCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveMetadataServicesFromCluster`: MetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `MetadataClustersAPI.RemoveMetadataServicesFromCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataClusterId** | **int64** | metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMetadataServicesFromClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MetadataClusterRemoveServicesReq**](MetadataClusterRemoveServicesReq.md) | metadata services | 

### Return type

[**MetadataClusterResp**](MetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMetadataClusterToStretched

> MetadataClusterResp SetMetadataClusterToStretched(ctx, metadataClusterId).Body(body).Execute()





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
	metadataClusterId := int64(789) // int64 | metadata cluster id
	body := *openapiclient.NewMetadataClusterSetStretchedReq() // MetadataClusterSetStretchedReq | cluster stretched info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataClustersAPI.SetMetadataClusterToStretched(context.Background(), metadataClusterId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataClustersAPI.SetMetadataClusterToStretched``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetMetadataClusterToStretched`: MetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `MetadataClustersAPI.SetMetadataClusterToStretched`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataClusterId** | **int64** | metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMetadataClusterToStretchedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MetadataClusterSetStretchedReq**](MetadataClusterSetStretchedReq.md) | cluster stretched info | 

### Return type

[**MetadataClusterResp**](MetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMetadataClusterTransLocator

> MetadataClusterResp SetMetadataClusterTransLocator(ctx, metadataClusterId).Body(body).Execute()





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
	metadataClusterId := int64(789) // int64 | metadata cluster id
	body := *openapiclient.NewMetadataClusterSetTransLocaterReq() // MetadataClusterSetTransLocaterReq | transfer locator

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataClustersAPI.SetMetadataClusterTransLocator(context.Background(), metadataClusterId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataClustersAPI.SetMetadataClusterTransLocator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetMetadataClusterTransLocator`: MetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `MetadataClustersAPI.SetMetadataClusterTransLocator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataClusterId** | **int64** | metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMetadataClusterTransLocatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MetadataClusterSetTransLocaterReq**](MetadataClusterSetTransLocaterReq.md) | transfer locator | 

### Return type

[**MetadataClusterResp**](MetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetadataCluster

> MetadataClusterResp UpdateMetadataCluster(ctx, metadataClusterId).Body(body).Execute()





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
	metadataClusterId := int64(789) // int64 | metadata cluster id
	body := *openapiclient.NewMetadataClusterUpdateReq() // MetadataClusterUpdateReq | metadata cluster info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataClustersAPI.UpdateMetadataCluster(context.Background(), metadataClusterId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataClustersAPI.UpdateMetadataCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMetadataCluster`: MetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `MetadataClustersAPI.UpdateMetadataCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataClusterId** | **int64** | metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetadataClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MetadataClusterUpdateReq**](MetadataClusterUpdateReq.md) | metadata cluster info | 

### Return type

[**MetadataClusterResp**](MetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrimaryDc

> MetadataClusterResp UpdatePrimaryDc(ctx, metadataClusterId).Body(body).Execute()





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
	metadataClusterId := int64(789) // int64 | metadata cluster id
	body := *openapiclient.NewMetadataClusterUpdatePrimaryDcReq() // MetadataClusterUpdatePrimaryDcReq | metadata cluster info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataClustersAPI.UpdatePrimaryDc(context.Background(), metadataClusterId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataClustersAPI.UpdatePrimaryDc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrimaryDc`: MetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `MetadataClustersAPI.UpdatePrimaryDc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataClusterId** | **int64** | metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrimaryDcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MetadataClusterUpdatePrimaryDcReq**](MetadataClusterUpdatePrimaryDcReq.md) | metadata cluster info | 

### Return type

[**MetadataClusterResp**](MetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

