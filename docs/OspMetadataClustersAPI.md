# \OspMetadataClustersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPartitions**](OspMetadataClustersAPI.md#AddPartitions) | **Post** /osp-metadata-clusters/{id}:add-partitions | 
[**CreateOspMetadataCluster**](OspMetadataClustersAPI.md#CreateOspMetadataCluster) | **Post** /osp-metadata-clusters/ | 
[**DeleteOspMetadataCluster**](OspMetadataClustersAPI.md#DeleteOspMetadataCluster) | **Delete** /osp-metadata-clusters/{id} | 
[**GetOspMetadataCluster**](OspMetadataClustersAPI.md#GetOspMetadataCluster) | **Get** /osp-metadata-clusters/{id} | 
[**GetOspMetadataClusterSamples**](OspMetadataClustersAPI.md#GetOspMetadataClusterSamples) | **Get** /osp-metadata-clusters/{id}/samples | 
[**ListOspMetadataClusters**](OspMetadataClustersAPI.md#ListOspMetadataClusters) | **Get** /osp-metadata-clusters/ | 
[**RemovePartitions**](OspMetadataClustersAPI.md#RemovePartitions) | **Post** /osp-metadata-clusters/{id}:remove-partitions | 
[**SetCoordinator**](OspMetadataClustersAPI.md#SetCoordinator) | **Post** /osp-metadata-clusters/{id}:set-coordinator | 
[**SetDeployMode**](OspMetadataClustersAPI.md#SetDeployMode) | **Post** /osp-metadata-clusters/{id}:set-deploy-mode | 
[**SetLogServices**](OspMetadataClustersAPI.md#SetLogServices) | **Post** /osp-metadata-clusters/{id}:set-log-services | 
[**SetMetadataServices**](OspMetadataClustersAPI.md#SetMetadataServices) | **Post** /osp-metadata-clusters/{id}:set-metadata-services | 
[**SetMinAvailableSpaceRatio**](OspMetadataClustersAPI.md#SetMinAvailableSpaceRatio) | **Post** /osp-metadata-clusters/{id}:set-min-available-space-ratio | 
[**SetName**](OspMetadataClustersAPI.md#SetName) | **Post** /osp-metadata-clusters/{id}:set-name | 
[**SetPrimaryDc**](OspMetadataClustersAPI.md#SetPrimaryDc) | **Post** /osp-metadata-clusters/{id}:set-primary-dc | 
[**SetSecurityPolicy**](OspMetadataClustersAPI.md#SetSecurityPolicy) | **Post** /osp-metadata-clusters/{id}:set-security-policy | 
[**SetServices**](OspMetadataClustersAPI.md#SetServices) | **Post** /osp-metadata-clusters/{id}:set-services | 
[**SetSpaceUsageFactor**](OspMetadataClustersAPI.md#SetSpaceUsageFactor) | **Post** /osp-metadata-clusters/{id}:set-space-usage-factor | 
[**SetStatelessNum**](OspMetadataClustersAPI.md#SetStatelessNum) | **Post** /osp-metadata-clusters/{id}:set-stateless-num | 
[**SetStatelessServer**](OspMetadataClustersAPI.md#SetStatelessServer) | **Post** /osp-metadata-clusters/{id}:set-stateless-server | 
[**UpgradeOspMetadataCluster**](OspMetadataClustersAPI.md#UpgradeOspMetadataCluster) | **Post** /osp-metadata-clusters/{id}:upgrade-service | 



## AddPartitions

> OspMetadataClusterResp AddPartitions(ctx, id).Body(body).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id
	body := *openapiclient.NewOspMetadataClusterUpdateReq() // OspMetadataClusterUpdateReq | osp metadata cluster info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.AddPartitions(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.AddPartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPartitions`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.AddPartitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OspMetadataClusterUpdateReq**](OspMetadataClusterUpdateReq.md) | osp metadata cluster info | 

### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOspMetadataCluster

> OspMetadataClusterResp CreateOspMetadataCluster(ctx).Body(body).Execute()





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
	body := *openapiclient.NewOspMetadataClusterCreateReq() // OspMetadataClusterCreateReq | osp metadata cluster info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.CreateOspMetadataCluster(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.CreateOspMetadataCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOspMetadataCluster`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.CreateOspMetadataCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOspMetadataClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OspMetadataClusterCreateReq**](OspMetadataClusterCreateReq.md) | osp metadata cluster info | 

### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOspMetadataCluster

> OspMetadataClusterResp DeleteOspMetadataCluster(ctx, id).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.DeleteOspMetadataCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.DeleteOspMetadataCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOspMetadataCluster`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.DeleteOspMetadataCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOspMetadataClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOspMetadataCluster

> OspMetadataClusterResp GetOspMetadataCluster(ctx, id).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.GetOspMetadataCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.GetOspMetadataCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOspMetadataCluster`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.GetOspMetadataCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOspMetadataClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOspMetadataClusterSamples

> OspMetadataClusterSamplesResp GetOspMetadataClusterSamples(ctx, id).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.GetOspMetadataClusterSamples(context.Background(), id).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.GetOspMetadataClusterSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOspMetadataClusterSamples`: OspMetadataClusterSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.GetOspMetadataClusterSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOspMetadataClusterSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**OspMetadataClusterSamplesResp**](OspMetadataClusterSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOspMetadataClusters

> OspMetadataClustersResp ListOspMetadataClusters(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).ObjectStorageId(objectStorageId).Execute()





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
	clusterId := int64(789) // int64 | The id of osp metadata cluster's cluster (optional)
	objectStorageId := int64(789) // int64 | The id of osp metadata cluster's object storage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.ListOspMetadataClusters(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).ObjectStorageId(objectStorageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.ListOspMetadataClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOspMetadataClusters`: OspMetadataClustersResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.ListOspMetadataClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOspMetadataClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **int64** | The id of osp metadata cluster&#39;s cluster | 
 **objectStorageId** | **int64** | The id of osp metadata cluster&#39;s object storage | 

### Return type

[**OspMetadataClustersResp**](OspMetadataClustersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePartitions

> OspMetadataClusterResp RemovePartitions(ctx, id).Body(body).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id
	body := *openapiclient.NewOspMetadataClusterUpdateReq() // OspMetadataClusterUpdateReq | osp metadata cluster info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.RemovePartitions(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.RemovePartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemovePartitions`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.RemovePartitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OspMetadataClusterUpdateReq**](OspMetadataClusterUpdateReq.md) | osp metadata cluster info | 

### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCoordinator

> OspMetadataClusterResp SetCoordinator(ctx, id).Body(body).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id
	body := *openapiclient.NewOspMetadataClusterUpdateReq() // OspMetadataClusterUpdateReq | osp metadata cluster info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.SetCoordinator(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.SetCoordinator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetCoordinator`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.SetCoordinator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCoordinatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OspMetadataClusterUpdateReq**](OspMetadataClusterUpdateReq.md) | osp metadata cluster info | 

### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDeployMode

> OspMetadataClusterResp SetDeployMode(ctx, id).Body(body).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id
	body := *openapiclient.NewOspMetadataClusterUpdateReq() // OspMetadataClusterUpdateReq | osp metadata cluster info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.SetDeployMode(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.SetDeployMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDeployMode`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.SetDeployMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDeployModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OspMetadataClusterUpdateReq**](OspMetadataClusterUpdateReq.md) | osp metadata cluster info | 

### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLogServices

> OspMetadataClusterResp SetLogServices(ctx, id).Body(body).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id
	body := *openapiclient.NewOspMetadataClusterUpdateReq() // OspMetadataClusterUpdateReq | osp metadata cluster info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.SetLogServices(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.SetLogServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetLogServices`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.SetLogServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLogServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OspMetadataClusterUpdateReq**](OspMetadataClusterUpdateReq.md) | osp metadata cluster info | 

### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMetadataServices

> OspMetadataClusterResp SetMetadataServices(ctx, id).Body(body).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id
	body := *openapiclient.NewOspMetadataClusterUpdateReq() // OspMetadataClusterUpdateReq | osp metadata cluster info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.SetMetadataServices(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.SetMetadataServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetMetadataServices`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.SetMetadataServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMetadataServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OspMetadataClusterUpdateReq**](OspMetadataClusterUpdateReq.md) | osp metadata cluster info | 

### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMinAvailableSpaceRatio

> OspMetadataClusterResp SetMinAvailableSpaceRatio(ctx, id).Body(body).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id
	body := *openapiclient.NewOspMetadataClusterSetMinAvailableSpaceRatioReq() // OspMetadataClusterSetMinAvailableSpaceRatioReq | osp metadata cluster min available space ratio info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.SetMinAvailableSpaceRatio(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.SetMinAvailableSpaceRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetMinAvailableSpaceRatio`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.SetMinAvailableSpaceRatio`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMinAvailableSpaceRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OspMetadataClusterSetMinAvailableSpaceRatioReq**](OspMetadataClusterSetMinAvailableSpaceRatioReq.md) | osp metadata cluster min available space ratio info | 

### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetName

> OspMetadataClusterResp SetName(ctx, id).Body(body).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id
	body := *openapiclient.NewOspMetadataClusterSetNameReq() // OspMetadataClusterSetNameReq | osp metadata cluster info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.SetName(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.SetName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetName`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.SetName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OspMetadataClusterSetNameReq**](OspMetadataClusterSetNameReq.md) | osp metadata cluster info | 

### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPrimaryDc

> OspMetadataClusterResp SetPrimaryDc(ctx, id).Body(body).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id
	body := *openapiclient.NewOspMetadataClusterSetPrimaryDcReq() // OspMetadataClusterSetPrimaryDcReq | osp metadata cluster primary data center info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.SetPrimaryDc(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.SetPrimaryDc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPrimaryDc`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.SetPrimaryDc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPrimaryDcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OspMetadataClusterSetPrimaryDcReq**](OspMetadataClusterSetPrimaryDcReq.md) | osp metadata cluster primary data center info | 

### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSecurityPolicy

> OspMetadataClusterResp SetSecurityPolicy(ctx, id).Body(body).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id
	body := *openapiclient.NewOspMetadataClusterUpdateReq() // OspMetadataClusterUpdateReq | osp metadata cluster info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.SetSecurityPolicy(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.SetSecurityPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSecurityPolicy`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.SetSecurityPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSecurityPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OspMetadataClusterUpdateReq**](OspMetadataClusterUpdateReq.md) | osp metadata cluster info | 

### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetServices

> OspMetadataClusterResp SetServices(ctx, id).Body(body).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id
	body := *openapiclient.NewOspMetadataClusterUpdateReq() // OspMetadataClusterUpdateReq | osp metadata cluster info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.SetServices(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.SetServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetServices`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.SetServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OspMetadataClusterUpdateReq**](OspMetadataClusterUpdateReq.md) | osp metadata cluster info | 

### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSpaceUsageFactor

> OspMetadataClusterResp SetSpaceUsageFactor(ctx, id).Body(body).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id
	body := *openapiclient.NewOspMetadataClusterSetSpaceUsageFactorReq() // OspMetadataClusterSetSpaceUsageFactorReq | osp metadata cluster space usage factor info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.SetSpaceUsageFactor(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.SetSpaceUsageFactor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSpaceUsageFactor`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.SetSpaceUsageFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSpaceUsageFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OspMetadataClusterSetSpaceUsageFactorReq**](OspMetadataClusterSetSpaceUsageFactorReq.md) | osp metadata cluster space usage factor info | 

### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetStatelessNum

> OspMetadataClusterResp SetStatelessNum(ctx, id).Body(body).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id
	body := *openapiclient.NewOspMetadataClusterSetStatelessNumReq() // OspMetadataClusterSetStatelessNumReq | osp metadata cluster transaction roles info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.SetStatelessNum(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.SetStatelessNum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetStatelessNum`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.SetStatelessNum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetStatelessNumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OspMetadataClusterSetStatelessNumReq**](OspMetadataClusterSetStatelessNumReq.md) | osp metadata cluster transaction roles info | 

### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetStatelessServer

> OspMetadataClusterResp SetStatelessServer(ctx, id).Body(body).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id
	body := *openapiclient.NewOspMetadataClusterSetStatelessServerReq() // OspMetadataClusterSetStatelessServerReq | osp metadata cluster info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.SetStatelessServer(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.SetStatelessServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetStatelessServer`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.SetStatelessServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetStatelessServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OspMetadataClusterSetStatelessServerReq**](OspMetadataClusterSetStatelessServerReq.md) | osp metadata cluster info | 

### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeOspMetadataCluster

> OspMetadataClusterResp UpgradeOspMetadataCluster(ctx, id).Execute()





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
	id := int64(789) // int64 | osp metadata cluster id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataClustersAPI.UpgradeOspMetadataCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataClustersAPI.UpgradeOspMetadataCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeOspMetadataCluster`: OspMetadataClusterResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataClustersAPI.UpgradeOspMetadataCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeOspMetadataClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OspMetadataClusterResp**](OspMetadataClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

