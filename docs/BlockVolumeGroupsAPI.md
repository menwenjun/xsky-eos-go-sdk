# \BlockVolumeGroupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlockVolumeGroup**](BlockVolumeGroupsAPI.md#CreateBlockVolumeGroup) | **Post** /block-volume-groups/ | 
[**DeleteBlockVolumeGroup**](BlockVolumeGroupsAPI.md#DeleteBlockVolumeGroup) | **Delete** /block-volume-groups/{block_volume_group_id} | 
[**GetBlockVolumeGroup**](BlockVolumeGroupsAPI.md#GetBlockVolumeGroup) | **Get** /block-volume-groups/{block_volume_group_id} | 
[**ListBlockVolumeGroups**](BlockVolumeGroupsAPI.md#ListBlockVolumeGroups) | **Get** /block-volume-groups/ | 
[**RollbackVolumeGroup**](BlockVolumeGroupsAPI.md#RollbackVolumeGroup) | **Post** /block-volume-groups/{block_volume_group_id}:rollback | 
[**SetVolumeGroupSnapshotReplicationProtection**](BlockVolumeGroupsAPI.md#SetVolumeGroupSnapshotReplicationProtection) | **Post** /block-volume-groups/{block_volume_group_id}:set-snapshot-replication-protection | 
[**UnsetVolumeGroupSnapshotReplicationProtection**](BlockVolumeGroupsAPI.md#UnsetVolumeGroupSnapshotReplicationProtection) | **Post** /block-volume-groups/{block_volume_group_id}:unset-snapshot-replication-protection | 
[**UpdateBlockVolumeGroup**](BlockVolumeGroupsAPI.md#UpdateBlockVolumeGroup) | **Patch** /block-volume-groups/{block_volume_group_id} | 



## CreateBlockVolumeGroup

> VolumeGroupResp CreateBlockVolumeGroup(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewVolumeGroupCreateReq(*openapiclient.NewVolumeGroupCreateReqVolumeGroup("Name_example")) // VolumeGroupCreateReq | volume group info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeGroupsAPI.CreateBlockVolumeGroup(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeGroupsAPI.CreateBlockVolumeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlockVolumeGroup`: VolumeGroupResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeGroupsAPI.CreateBlockVolumeGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlockVolumeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VolumeGroupCreateReq**](VolumeGroupCreateReq.md) | volume group info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**VolumeGroupResp**](VolumeGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlockVolumeGroup

> VolumeGroupResp DeleteBlockVolumeGroup(ctx, blockVolumeGroupId).Execute()





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
	blockVolumeGroupId := int64(789) // int64 | block volume group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeGroupsAPI.DeleteBlockVolumeGroup(context.Background(), blockVolumeGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeGroupsAPI.DeleteBlockVolumeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlockVolumeGroup`: VolumeGroupResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeGroupsAPI.DeleteBlockVolumeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeGroupId** | **int64** | block volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlockVolumeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeGroupResp**](VolumeGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockVolumeGroup

> VolumeGroupResp GetBlockVolumeGroup(ctx, blockVolumeGroupId).Execute()





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
	blockVolumeGroupId := int64(789) // int64 | block volume group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeGroupsAPI.GetBlockVolumeGroup(context.Background(), blockVolumeGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeGroupsAPI.GetBlockVolumeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockVolumeGroup`: VolumeGroupResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeGroupsAPI.GetBlockVolumeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeGroupId** | **int64** | block volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockVolumeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeGroupResp**](VolumeGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlockVolumeGroups

> VolumeGroupsResp ListBlockVolumeGroups(ctx).ClusterId(clusterId).Passive(passive).Name(name).Limit(limit).Offset(offset).Q(q).Sort(sort).DpVolumeGroupSnapshotReplicationPolicyId(dpVolumeGroupSnapshotReplicationPolicyId).Execute()





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
	passive := true // bool | passive or not (optional)
	name := "name_example" // string | name of volume group (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	dpVolumeGroupSnapshotReplicationPolicyId := int64(789) // int64 | show volume groups of specific dp block async replication policy (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeGroupsAPI.ListBlockVolumeGroups(context.Background()).ClusterId(clusterId).Passive(passive).Name(name).Limit(limit).Offset(offset).Q(q).Sort(sort).DpVolumeGroupSnapshotReplicationPolicyId(dpVolumeGroupSnapshotReplicationPolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeGroupsAPI.ListBlockVolumeGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlockVolumeGroups`: VolumeGroupsResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeGroupsAPI.ListBlockVolumeGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlockVolumeGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 
 **passive** | **bool** | passive or not | 
 **name** | **string** | name of volume group | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **dpVolumeGroupSnapshotReplicationPolicyId** | **int64** | show volume groups of specific dp block async replication policy | 

### Return type

[**VolumeGroupsResp**](VolumeGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackVolumeGroup

> VolumeGroupResp RollbackVolumeGroup(ctx, blockVolumeGroupId).Body(body).Execute()





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
	blockVolumeGroupId := int64(789) // int64 | block volume group id
	body := *openapiclient.NewVolumeGroupRollbackReq(*openapiclient.NewVolumeGroupRollbackReqVolumeGroup()) // VolumeGroupRollbackReq | rollback info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeGroupsAPI.RollbackVolumeGroup(context.Background(), blockVolumeGroupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeGroupsAPI.RollbackVolumeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RollbackVolumeGroup`: VolumeGroupResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeGroupsAPI.RollbackVolumeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeGroupId** | **int64** | block volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackVolumeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VolumeGroupRollbackReq**](VolumeGroupRollbackReq.md) | rollback info | 

### Return type

[**VolumeGroupResp**](VolumeGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVolumeGroupSnapshotReplicationProtection

> VolumeGroupResp SetVolumeGroupSnapshotReplicationProtection(ctx, blockVolumeGroupId).Body(body).Execute()





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
	blockVolumeGroupId := int64(789) // int64 | the block volume group id
	body := *openapiclient.NewVolumeGroupSnapshotReplicationProtectionReq("DestVolumeGroupName_example", int64(123), []openapiclient.VolumePairInfo{*openapiclient.NewVolumePairInfo(int64(123), "DestPoolName_example", "DestVolumeName_example", int64(123))}) // VolumeGroupSnapshotReplicationProtectionReq | request info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeGroupsAPI.SetVolumeGroupSnapshotReplicationProtection(context.Background(), blockVolumeGroupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeGroupsAPI.SetVolumeGroupSnapshotReplicationProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVolumeGroupSnapshotReplicationProtection`: VolumeGroupResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeGroupsAPI.SetVolumeGroupSnapshotReplicationProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeGroupId** | **int64** | the block volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVolumeGroupSnapshotReplicationProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VolumeGroupSnapshotReplicationProtectionReq**](VolumeGroupSnapshotReplicationProtectionReq.md) | request info | 

### Return type

[**VolumeGroupResp**](VolumeGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetVolumeGroupSnapshotReplicationProtection

> VolumeGroupResp UnsetVolumeGroupSnapshotReplicationProtection(ctx, blockVolumeGroupId).Force(force).ReserveVolumeGroup(reserveVolumeGroup).Execute()





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
	blockVolumeGroupId := int64(789) // int64 | the volume group id
	force := true // bool | force unset or not (optional)
	reserveVolumeGroup := true // bool | reserve replicated volume group or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeGroupsAPI.UnsetVolumeGroupSnapshotReplicationProtection(context.Background(), blockVolumeGroupId).Force(force).ReserveVolumeGroup(reserveVolumeGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeGroupsAPI.UnsetVolumeGroupSnapshotReplicationProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsetVolumeGroupSnapshotReplicationProtection`: VolumeGroupResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeGroupsAPI.UnsetVolumeGroupSnapshotReplicationProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeGroupId** | **int64** | the volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetVolumeGroupSnapshotReplicationProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force unset or not | 
 **reserveVolumeGroup** | **bool** | reserve replicated volume group or not | 

### Return type

[**VolumeGroupResp**](VolumeGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlockVolumeGroup

> VolumeGroupResp UpdateBlockVolumeGroup(ctx, blockVolumeGroupId).Body(body).Execute()





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
	blockVolumeGroupId := int64(789) // int64 | block volume group id
	body := *openapiclient.NewVolumeGroupUpdateReq(*openapiclient.NewVolumeGroupUpdateReqVolumeGroup("Name_example")) // VolumeGroupUpdateReq | volume group info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumeGroupsAPI.UpdateBlockVolumeGroup(context.Background(), blockVolumeGroupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumeGroupsAPI.UpdateBlockVolumeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBlockVolumeGroup`: VolumeGroupResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumeGroupsAPI.UpdateBlockVolumeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeGroupId** | **int64** | block volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlockVolumeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VolumeGroupUpdateReq**](VolumeGroupUpdateReq.md) | volume group info | 

### Return type

[**VolumeGroupResp**](VolumeGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

