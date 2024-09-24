# \DfsGatewayGroupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDfsGateways**](DfsGatewayGroupsAPI.md#AddDfsGateways) | **Post** /dfs-gateway-groups/{dfs_gateway_group_id}:add-gateways | 
[**CreateDfsGatewayGroup**](DfsGatewayGroupsAPI.md#CreateDfsGatewayGroup) | **Post** /dfs-gateway-groups/ | 
[**DeleteDfsGatewayGroup**](DfsGatewayGroupsAPI.md#DeleteDfsGatewayGroup) | **Delete** /dfs-gateway-groups/{dfs_gateway_group_id} | 
[**GetDfsGatewayGroup**](DfsGatewayGroupsAPI.md#GetDfsGatewayGroup) | **Get** /dfs-gateway-groups/{dfs_gateway_group_id} | 
[**ListDfsGatewayGroups**](DfsGatewayGroupsAPI.md#ListDfsGatewayGroups) | **Get** /dfs-gateway-groups/ | 
[**RebuildDfsGateways**](DfsGatewayGroupsAPI.md#RebuildDfsGateways) | **Post** /dfs-gateway-groups/{dfs_gateway_group_id}:rebuild-gateways | 
[**RemoveDfsGateways**](DfsGatewayGroupsAPI.md#RemoveDfsGateways) | **Post** /dfs-gateway-groups/{dfs_gateway_group_id}:remove-gateways | 
[**UpdateDfsGatewayGroup**](DfsGatewayGroupsAPI.md#UpdateDfsGatewayGroup) | **Patch** /dfs-gateway-groups/{dfs_gateway_group_id} | 
[**UpdateDfsGatewayGroupTypes**](DfsGatewayGroupsAPI.md#UpdateDfsGatewayGroupTypes) | **Post** /dfs-gateway-groups/{dfs_gateway_group_id}:update-types | 
[**UpdateDfsGatewayGroupVIPs**](DfsGatewayGroupsAPI.md#UpdateDfsGatewayGroupVIPs) | **Post** /dfs-gateway-groups/{dfs_gateway_group_id}:update-vips | 



## AddDfsGateways

> DfsGatewayGroupResp AddDfsGateways(ctx, dfsGatewayGroupId).Body(body).Execute()





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
	dfsGatewayGroupId := int64(789) // int64 | gateway group id
	body := *openapiclient.NewDfsGatewayGroupAddGatewaysReq(*openapiclient.NewDfsGatewayGroupAddGatewaysReqGatewayGroup([]openapiclient.DfsGatewayReq{*openapiclient.NewDfsGatewayReq(int64(123), int64(123))})) // DfsGatewayGroupAddGatewaysReq | gateways info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayGroupsAPI.AddDfsGateways(context.Background(), dfsGatewayGroupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayGroupsAPI.AddDfsGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDfsGateways`: DfsGatewayGroupResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayGroupsAPI.AddDfsGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayGroupId** | **int64** | gateway group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDfsGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsGatewayGroupAddGatewaysReq**](DfsGatewayGroupAddGatewaysReq.md) | gateways info | 

### Return type

[**DfsGatewayGroupResp**](DfsGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDfsGatewayGroup

> DfsGatewayGroupResp CreateDfsGatewayGroup(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewDfsGatewayGroupCreateReq(*openapiclient.NewDfsGatewayGroupCreateReqGatewayGroup([]openapiclient.DfsGatewayReq{*openapiclient.NewDfsGatewayReq(int64(123), int64(123))}, []string{"DfsVips_example"}, "Name_example", []string{"Types_example"})) // DfsGatewayGroupCreateReq | gateway group info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayGroupsAPI.CreateDfsGatewayGroup(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayGroupsAPI.CreateDfsGatewayGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsGatewayGroup`: DfsGatewayGroupResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayGroupsAPI.CreateDfsGatewayGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsGatewayGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsGatewayGroupCreateReq**](DfsGatewayGroupCreateReq.md) | gateway group info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**DfsGatewayGroupResp**](DfsGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsGatewayGroup

> DfsGatewayGroupResp DeleteDfsGatewayGroup(ctx, dfsGatewayGroupId).Force(force).Execute()





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
	dfsGatewayGroupId := int64(789) // int64 | gateway group id
	force := true // bool | force delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayGroupsAPI.DeleteDfsGatewayGroup(context.Background(), dfsGatewayGroupId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayGroupsAPI.DeleteDfsGatewayGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsGatewayGroup`: DfsGatewayGroupResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayGroupsAPI.DeleteDfsGatewayGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayGroupId** | **int64** | gateway group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsGatewayGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 

### Return type

[**DfsGatewayGroupResp**](DfsGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsGatewayGroup

> DfsGatewayGroupResp GetDfsGatewayGroup(ctx, dfsGatewayGroupId).Execute()





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
	dfsGatewayGroupId := int64(789) // int64 | gateway group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayGroupsAPI.GetDfsGatewayGroup(context.Background(), dfsGatewayGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayGroupsAPI.GetDfsGatewayGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsGatewayGroup`: DfsGatewayGroupResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayGroupsAPI.GetDfsGatewayGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayGroupId** | **int64** | gateway group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsGatewayGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsGatewayGroupResp**](DfsGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsGatewayGroups

> DfsGatewayGroupsResp ListDfsGatewayGroups(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).Type_(type_).Security(security).Q(q).Sort(sort).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	type_ := "type__example" // string | type of dfs gateway group (optional)
	security := "security_example" // string | security of dfs gateway group (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayGroupsAPI.ListDfsGatewayGroups(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).Type_(type_).Security(security).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayGroupsAPI.ListDfsGatewayGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsGatewayGroups`: DfsGatewayGroupsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayGroupsAPI.ListDfsGatewayGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsGatewayGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **type_** | **string** | type of dfs gateway group | 
 **security** | **string** | security of dfs gateway group | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**DfsGatewayGroupsResp**](DfsGatewayGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebuildDfsGateways

> DfsGatewayGroupResp RebuildDfsGateways(ctx, dfsGatewayGroupId).Body(body).Execute()





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
	dfsGatewayGroupId := int64(789) // int64 | gateway group id
	body := *openapiclient.NewDfsGatewayGroupRebuildGatewaysReq(*openapiclient.NewDfsGatewayGroupRebuildGatewaysReqGatewayGroup([]int64{int64(123)})) // DfsGatewayGroupRebuildGatewaysReq | gateways info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayGroupsAPI.RebuildDfsGateways(context.Background(), dfsGatewayGroupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayGroupsAPI.RebuildDfsGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebuildDfsGateways`: DfsGatewayGroupResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayGroupsAPI.RebuildDfsGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayGroupId** | **int64** | gateway group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebuildDfsGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsGatewayGroupRebuildGatewaysReq**](DfsGatewayGroupRebuildGatewaysReq.md) | gateways info | 

### Return type

[**DfsGatewayGroupResp**](DfsGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDfsGateways

> DfsGatewayGroupResp RemoveDfsGateways(ctx, dfsGatewayGroupId).Body(body).Force(force).Execute()





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
	dfsGatewayGroupId := int64(789) // int64 | gateway group id
	body := *openapiclient.NewDfsGatewayGroupRemoveGatewaysReq(*openapiclient.NewDfsGatewayGroupRemoveGatewaysReqGatewayGroup([]int64{int64(123)})) // DfsGatewayGroupRemoveGatewaysReq | gateways info
	force := true // bool | force delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayGroupsAPI.RemoveDfsGateways(context.Background(), dfsGatewayGroupId).Body(body).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayGroupsAPI.RemoveDfsGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveDfsGateways`: DfsGatewayGroupResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayGroupsAPI.RemoveDfsGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayGroupId** | **int64** | gateway group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDfsGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsGatewayGroupRemoveGatewaysReq**](DfsGatewayGroupRemoveGatewaysReq.md) | gateways info | 
 **force** | **bool** | force delete or not | 

### Return type

[**DfsGatewayGroupResp**](DfsGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsGatewayGroup

> DfsGatewayGroupResp UpdateDfsGatewayGroup(ctx, dfsGatewayGroupId).Body(body).Execute()





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
	dfsGatewayGroupId := int64(789) // int64 | gateway group id
	body := *openapiclient.NewDfsGatewayGroupUpdateReq(*openapiclient.NewDfsGatewayGroupUpdateReqGatewayGroup()) // DfsGatewayGroupUpdateReq | gateway group info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayGroupsAPI.UpdateDfsGatewayGroup(context.Background(), dfsGatewayGroupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayGroupsAPI.UpdateDfsGatewayGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsGatewayGroup`: DfsGatewayGroupResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayGroupsAPI.UpdateDfsGatewayGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayGroupId** | **int64** | gateway group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsGatewayGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsGatewayGroupUpdateReq**](DfsGatewayGroupUpdateReq.md) | gateway group info | 

### Return type

[**DfsGatewayGroupResp**](DfsGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsGatewayGroupTypes

> DfsGatewayGroupResp UpdateDfsGatewayGroupTypes(ctx, dfsGatewayGroupId).Body(body).Execute()





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
	dfsGatewayGroupId := int64(789) // int64 | gateway group id
	body := *openapiclient.NewDfsGatewayGroupUpdateTypesReq(*openapiclient.NewDfsGatewayGroupUpdateTypesReqGatewayGroup()) // DfsGatewayGroupUpdateTypesReq | gateway group info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayGroupsAPI.UpdateDfsGatewayGroupTypes(context.Background(), dfsGatewayGroupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayGroupsAPI.UpdateDfsGatewayGroupTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsGatewayGroupTypes`: DfsGatewayGroupResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayGroupsAPI.UpdateDfsGatewayGroupTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayGroupId** | **int64** | gateway group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsGatewayGroupTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsGatewayGroupUpdateTypesReq**](DfsGatewayGroupUpdateTypesReq.md) | gateway group info | 

### Return type

[**DfsGatewayGroupResp**](DfsGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsGatewayGroupVIPs

> DfsGatewayGroupResp UpdateDfsGatewayGroupVIPs(ctx, dfsGatewayGroupId).Body(body).Execute()





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
	dfsGatewayGroupId := int64(789) // int64 | gateway group id
	body := *openapiclient.NewDfsGatewayGroupUpdateVIPsReq(*openapiclient.NewDfsGatewayGroupUpdateVIPsReqGatewayGroup()) // DfsGatewayGroupUpdateVIPsReq | gateway group info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsGatewayGroupsAPI.UpdateDfsGatewayGroupVIPs(context.Background(), dfsGatewayGroupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsGatewayGroupsAPI.UpdateDfsGatewayGroupVIPs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsGatewayGroupVIPs`: DfsGatewayGroupResp
	fmt.Fprintf(os.Stdout, "Response from `DfsGatewayGroupsAPI.UpdateDfsGatewayGroupVIPs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsGatewayGroupId** | **int64** | gateway group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsGatewayGroupVIPsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsGatewayGroupUpdateVIPsReq**](DfsGatewayGroupUpdateVIPsReq.md) | gateway group info | 

### Return type

[**DfsGatewayGroupResp**](DfsGatewayGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

