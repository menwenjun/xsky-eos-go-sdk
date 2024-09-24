# \PlacementNodesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePlacementNode**](PlacementNodesAPI.md#CreatePlacementNode) | **Post** /placement-nodes/ | 
[**DeletePlacementNode**](PlacementNodesAPI.md#DeletePlacementNode) | **Delete** /placement-nodes/{placement_node_id} | 
[**GetPlacementNode**](PlacementNodesAPI.md#GetPlacementNode) | **Get** /placement-nodes/{placement_node_id} | 
[**GetPlacementNodeTopology**](PlacementNodesAPI.md#GetPlacementNodeTopology) | **Get** /placement-nodes/{placement_node_id}/topology | 
[**GetTopologyFromOsds**](PlacementNodesAPI.md#GetTopologyFromOsds) | **Post** /placement-nodes/:topology-from-osd | 
[**ListPlacementNodes**](PlacementNodesAPI.md#ListPlacementNodes) | **Get** /placement-nodes/ | 
[**UpdatePlacementNode**](PlacementNodesAPI.md#UpdatePlacementNode) | **Patch** /placement-nodes/{placement_node_id} | 



## CreatePlacementNode

> PlacementNodeResp CreatePlacementNode(ctx).Body(body).Execute()





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
	body := *openapiclient.NewPlacementNodeCreateReq() // PlacementNodeCreateReq | placement node info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlacementNodesAPI.CreatePlacementNode(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlacementNodesAPI.CreatePlacementNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlacementNode`: PlacementNodeResp
	fmt.Fprintf(os.Stdout, "Response from `PlacementNodesAPI.CreatePlacementNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlacementNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PlacementNodeCreateReq**](PlacementNodeCreateReq.md) | placement node info | 

### Return type

[**PlacementNodeResp**](PlacementNodeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlacementNode

> DeletePlacementNode(ctx, placementNodeId).Execute()





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
	placementNodeId := int64(789) // int64 | placement node id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlacementNodesAPI.DeletePlacementNode(context.Background(), placementNodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlacementNodesAPI.DeletePlacementNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placementNodeId** | **int64** | placement node id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlacementNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetPlacementNode

> PlacementNodeResp GetPlacementNode(ctx, placementNodeId).Execute()





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
	placementNodeId := int64(789) // int64 | placement node id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlacementNodesAPI.GetPlacementNode(context.Background(), placementNodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlacementNodesAPI.GetPlacementNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlacementNode`: PlacementNodeResp
	fmt.Fprintf(os.Stdout, "Response from `PlacementNodesAPI.GetPlacementNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placementNodeId** | **int64** | placement node id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlacementNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlacementNodeResp**](PlacementNodeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlacementNodeTopology

> PlacementNodeTopologyResp GetPlacementNodeTopology(ctx, placementNodeId).ClusterId(clusterId).OsdType(osdType).OsdRole(osdRole).WithCompound(withCompound).WithHybrid(withHybrid).Execute()





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
	placementNodeId := int64(789) // int64 | placement node id
	clusterId := "clusterId_example" // string | cluster id (optional)
	osdType := "osdType_example" // string | osd type (optional)
	osdRole := "osdRole_example" // string | osd role (optional)
	withCompound := true // bool | with compound osd (optional)
	withHybrid := true // bool | with hybrid osd (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlacementNodesAPI.GetPlacementNodeTopology(context.Background(), placementNodeId).ClusterId(clusterId).OsdType(osdType).OsdRole(osdRole).WithCompound(withCompound).WithHybrid(withHybrid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlacementNodesAPI.GetPlacementNodeTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlacementNodeTopology`: PlacementNodeTopologyResp
	fmt.Fprintf(os.Stdout, "Response from `PlacementNodesAPI.GetPlacementNodeTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placementNodeId** | **int64** | placement node id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlacementNodeTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterId** | **string** | cluster id | 
 **osdType** | **string** | osd type | 
 **osdRole** | **string** | osd role | 
 **withCompound** | **bool** | with compound osd | 
 **withHybrid** | **bool** | with hybrid osd | 

### Return type

[**PlacementNodeTopologyResp**](PlacementNodeTopologyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopologyFromOsds

> PlacementNodeTopologyResp GetTopologyFromOsds(ctx).Body(body).Execute()





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
	body := *openapiclient.NewTopologyFromOsdReq([]int64{int64(123)}) // TopologyFromOsdReq | osds

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlacementNodesAPI.GetTopologyFromOsds(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlacementNodesAPI.GetTopologyFromOsds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopologyFromOsds`: PlacementNodeTopologyResp
	fmt.Fprintf(os.Stdout, "Response from `PlacementNodesAPI.GetTopologyFromOsds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologyFromOsdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TopologyFromOsdReq**](TopologyFromOsdReq.md) | osds | 

### Return type

[**PlacementNodeTopologyResp**](PlacementNodeTopologyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlacementNodes

> PlacementNodesResp ListPlacementNodes(ctx).Limit(limit).Offset(offset).Type_(type_).ParentId(parentId).Execute()





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
	type_ := "type__example" // string | filter placement nodes by type (optional)
	parentId := int64(789) // int64 | filter placement nodes by parent (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlacementNodesAPI.ListPlacementNodes(context.Background()).Limit(limit).Offset(offset).Type_(type_).ParentId(parentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlacementNodesAPI.ListPlacementNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPlacementNodes`: PlacementNodesResp
	fmt.Fprintf(os.Stdout, "Response from `PlacementNodesAPI.ListPlacementNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPlacementNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **type_** | **string** | filter placement nodes by type | 
 **parentId** | **int64** | filter placement nodes by parent | 

### Return type

[**PlacementNodesResp**](PlacementNodesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlacementNode

> PlacementNodeResp UpdatePlacementNode(ctx, placementNodeId).Body(body).Execute()





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
	placementNodeId := int64(789) // int64 | the placement node id
	body := *openapiclient.NewPlacementNodeUpdateReq() // PlacementNodeUpdateReq | the placement node info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlacementNodesAPI.UpdatePlacementNode(context.Background(), placementNodeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlacementNodesAPI.UpdatePlacementNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePlacementNode`: PlacementNodeResp
	fmt.Fprintf(os.Stdout, "Response from `PlacementNodesAPI.UpdatePlacementNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placementNodeId** | **int64** | the placement node id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlacementNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PlacementNodeUpdateReq**](PlacementNodeUpdateReq.md) | the placement node info | 

### Return type

[**PlacementNodeResp**](PlacementNodeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

