# \MappingGroupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVolumes**](MappingGroupsAPI.md#AddVolumes) | **Post** /mapping-groups/{mapping_group_id}/block-volumes | 
[**CreateMappingGroup**](MappingGroupsAPI.md#CreateMappingGroup) | **Post** /mapping-groups/ | 
[**DeleteMappingGroup**](MappingGroupsAPI.md#DeleteMappingGroup) | **Delete** /mapping-groups/{mapping_group_id} | 
[**GetMappingGroup**](MappingGroupsAPI.md#GetMappingGroup) | **Get** /mapping-groups/{mapping_group_id} | 
[**ListMappingGroups**](MappingGroupsAPI.md#ListMappingGroups) | **Get** /mapping-groups/ | 
[**RemoveVolumes**](MappingGroupsAPI.md#RemoveVolumes) | **Delete** /mapping-groups/{mapping_group_id}/block-volumes | 
[**UpdateMappingGroup**](MappingGroupsAPI.md#UpdateMappingGroup) | **Patch** /mapping-groups/{mapping_group_id} | 
[**UpdateMappingGroupClientGroup**](MappingGroupsAPI.md#UpdateMappingGroupClientGroup) | **Patch** /mapping-groups/{mapping_group_id}/client-group | 



## AddVolumes

> MappingGroupResp AddVolumes(ctx, mappingGroupId).BlockVolumeIds(blockVolumeIds).Execute()





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
	mappingGroupId := int64(789) // int64 | mapping group id
	blockVolumeIds := *openapiclient.NewMappingGroupAddVolumesReq([]int64{int64(123)}) // MappingGroupAddVolumesReq | block volume ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingGroupsAPI.AddVolumes(context.Background(), mappingGroupId).BlockVolumeIds(blockVolumeIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingGroupsAPI.AddVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddVolumes`: MappingGroupResp
	fmt.Fprintf(os.Stdout, "Response from `MappingGroupsAPI.AddVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingGroupId** | **int64** | mapping group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blockVolumeIds** | [**MappingGroupAddVolumesReq**](MappingGroupAddVolumesReq.md) | block volume ids | 

### Return type

[**MappingGroupResp**](MappingGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMappingGroup

> MappingGroupResp CreateMappingGroup(ctx).MappingGroup(mappingGroup).Execute()





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
	mappingGroup := *openapiclient.NewMappingGroupCreateReq(*openapiclient.NewMappingGroupReq()) // MappingGroupCreateReq | mapping info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingGroupsAPI.CreateMappingGroup(context.Background()).MappingGroup(mappingGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingGroupsAPI.CreateMappingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMappingGroup`: MappingGroupResp
	fmt.Fprintf(os.Stdout, "Response from `MappingGroupsAPI.CreateMappingGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMappingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappingGroup** | [**MappingGroupCreateReq**](MappingGroupCreateReq.md) | mapping info | 

### Return type

[**MappingGroupResp**](MappingGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMappingGroup

> MappingGroupResp DeleteMappingGroup(ctx, mappingGroupId).Force(force).Execute()





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
	mappingGroupId := int64(789) // int64 | mapping group id
	force := true // bool | force delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingGroupsAPI.DeleteMappingGroup(context.Background(), mappingGroupId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingGroupsAPI.DeleteMappingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMappingGroup`: MappingGroupResp
	fmt.Fprintf(os.Stdout, "Response from `MappingGroupsAPI.DeleteMappingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingGroupId** | **int64** | mapping group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMappingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 

### Return type

[**MappingGroupResp**](MappingGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMappingGroup

> MappingGroupResp GetMappingGroup(ctx, mappingGroupId).Execute()





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
	mappingGroupId := int64(789) // int64 | mapping group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingGroupsAPI.GetMappingGroup(context.Background(), mappingGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingGroupsAPI.GetMappingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMappingGroup`: MappingGroupResp
	fmt.Fprintf(os.Stdout, "Response from `MappingGroupsAPI.GetMappingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingGroupId** | **int64** | mapping group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMappingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MappingGroupResp**](MappingGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappingGroups

> MappingGroupsResp ListMappingGroups(ctx).Limit(limit).Offset(offset).AccessPathId(accessPathId).ClientGroupId(clientGroupId).Execute()





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
	accessPathId := int64(789) // int64 | access path id (optional)
	clientGroupId := int64(789) // int64 | client group id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingGroupsAPI.ListMappingGroups(context.Background()).Limit(limit).Offset(offset).AccessPathId(accessPathId).ClientGroupId(clientGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingGroupsAPI.ListMappingGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMappingGroups`: MappingGroupsResp
	fmt.Fprintf(os.Stdout, "Response from `MappingGroupsAPI.ListMappingGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMappingGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **accessPathId** | **int64** | access path id | 
 **clientGroupId** | **int64** | client group id | 

### Return type

[**MappingGroupsResp**](MappingGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveVolumes

> MappingGroupResp RemoveVolumes(ctx, mappingGroupId).BlockVolumeIds(blockVolumeIds).Force(force).Execute()





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
	mappingGroupId := int64(789) // int64 | mapping group id
	blockVolumeIds := *openapiclient.NewMappingGroupRemoveVolumesReq([]int64{int64(123)}) // MappingGroupRemoveVolumesReq | block volume ids
	force := true // bool | force delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingGroupsAPI.RemoveVolumes(context.Background(), mappingGroupId).BlockVolumeIds(blockVolumeIds).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingGroupsAPI.RemoveVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveVolumes`: MappingGroupResp
	fmt.Fprintf(os.Stdout, "Response from `MappingGroupsAPI.RemoveVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingGroupId** | **int64** | mapping group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blockVolumeIds** | [**MappingGroupRemoveVolumesReq**](MappingGroupRemoveVolumesReq.md) | block volume ids | 
 **force** | **bool** | force delete or not | 

### Return type

[**MappingGroupResp**](MappingGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMappingGroup

> MappingGroupResp UpdateMappingGroup(ctx, mappingGroupId).MappingGroup(mappingGroup).Execute()





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
	mappingGroupId := int64(789) // int64 | mapping group id
	mappingGroup := *openapiclient.NewMappingGroupUpdateReq(*openapiclient.NewMappingGroupUpdateReqMappingGroup()) // MappingGroupUpdateReq | mapping info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingGroupsAPI.UpdateMappingGroup(context.Background(), mappingGroupId).MappingGroup(mappingGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingGroupsAPI.UpdateMappingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMappingGroup`: MappingGroupResp
	fmt.Fprintf(os.Stdout, "Response from `MappingGroupsAPI.UpdateMappingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingGroupId** | **int64** | mapping group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMappingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mappingGroup** | [**MappingGroupUpdateReq**](MappingGroupUpdateReq.md) | mapping info | 

### Return type

[**MappingGroupResp**](MappingGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMappingGroupClientGroup

> MappingGroupResp UpdateMappingGroupClientGroup(ctx, mappingGroupId).ClientGroupId(clientGroupId).Execute()





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
	mappingGroupId := int64(789) // int64 | mapping group id
	clientGroupId := *openapiclient.NewMappingGroupUpdateClientGroupReq(int64(123)) // MappingGroupUpdateClientGroupReq | client group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MappingGroupsAPI.UpdateMappingGroupClientGroup(context.Background(), mappingGroupId).ClientGroupId(clientGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MappingGroupsAPI.UpdateMappingGroupClientGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMappingGroupClientGroup`: MappingGroupResp
	fmt.Fprintf(os.Stdout, "Response from `MappingGroupsAPI.UpdateMappingGroupClientGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingGroupId** | **int64** | mapping group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMappingGroupClientGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientGroupId** | [**MappingGroupUpdateClientGroupReq**](MappingGroupUpdateClientGroupReq.md) | client group id | 

### Return type

[**MappingGroupResp**](MappingGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

