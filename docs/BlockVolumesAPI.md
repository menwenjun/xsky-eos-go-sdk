# \BlockVolumesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchGetBlockVolumeSamples**](BlockVolumesAPI.md#BatchGetBlockVolumeSamples) | **Get** /block-volumes/samples | 
[**CreateBlockVolume**](BlockVolumesAPI.md#CreateBlockVolume) | **Post** /block-volumes/ | 
[**DeleteBlockVolume**](BlockVolumesAPI.md#DeleteBlockVolume) | **Delete** /block-volumes/{block_volume_id} | 
[**GetBlockVolume**](BlockVolumesAPI.md#GetBlockVolume) | **Get** /block-volumes/{block_volume_id} | 
[**GetBlockVolumeSamples**](BlockVolumesAPI.md#GetBlockVolumeSamples) | **Get** /block-volumes/{block_volume_id}/samples | 
[**ListBlockVolumes**](BlockVolumesAPI.md#ListBlockVolumes) | **Get** /block-volumes/ | 
[**MigrateBlockVolume**](BlockVolumesAPI.md#MigrateBlockVolume) | **Post** /block-volumes/{block_volume_id}:migrate | 
[**RebuildBlockVolumeReplication**](BlockVolumesAPI.md#RebuildBlockVolumeReplication) | **Post** /block-volumes/{block_volume_id}:rebuild-replication | 
[**SetAsyncReplicationProtection**](BlockVolumesAPI.md#SetAsyncReplicationProtection) | **Post** /block-volumes/{block_volume_id}:set-async-replication-protection | 
[**SetBackupProtection**](BlockVolumesAPI.md#SetBackupProtection) | **Post** /block-volumes/{block_volume_id}:set-backup-protection | 
[**SetBlockVolumeCrc**](BlockVolumesAPI.md#SetBlockVolumeCrc) | **Post** /block-volumes/{block_volume_id}:set-crc | 
[**SetBlockVolumeReplication**](BlockVolumesAPI.md#SetBlockVolumeReplication) | **Post** /block-volumes/{block_volume_id}:set-replication | 
[**SetSnapshotProtection**](BlockVolumesAPI.md#SetSnapshotProtection) | **Post** /block-volumes/{block_volume_id}:set-snapshot-protection | 
[**SuspendBlockVolumeReplication**](BlockVolumesAPI.md#SuspendBlockVolumeReplication) | **Post** /block-volumes/{block_volume_id}:suspend-replication | 
[**UnsetAsyncReplicationProtection**](BlockVolumesAPI.md#UnsetAsyncReplicationProtection) | **Post** /block-volumes/{block_volume_id}:unset-async-replication-protection | 
[**UnsetBackupProtection**](BlockVolumesAPI.md#UnsetBackupProtection) | **Post** /block-volumes/{block_volume_id}:unset-backup-protection | 
[**UnsetBlockVolumeCrc**](BlockVolumesAPI.md#UnsetBlockVolumeCrc) | **Post** /block-volumes/{block_volume_id}:unset-crc | 
[**UnsetBlockVolumeReplication**](BlockVolumesAPI.md#UnsetBlockVolumeReplication) | **Post** /block-volumes/{block_volume_id}:unset-replication | 
[**UnsetSnapshotProtection**](BlockVolumesAPI.md#UnsetSnapshotProtection) | **Post** /block-volumes/{block_volume_id}:unset-snapshot-protection | 
[**UpdateBlockVolume**](BlockVolumesAPI.md#UpdateBlockVolume) | **Patch** /block-volumes/{block_volume_id} | 
[**UpdateBlockVolumeVolumeName**](BlockVolumesAPI.md#UpdateBlockVolumeVolumeName) | **Post** /block-volumes/{block_volume_id}:update-volume-name | 
[**UpdateVolumeStat**](BlockVolumesAPI.md#UpdateVolumeStat) | **Post** /block-volumes/:update-stat | 
[**UpdateVolumeStats**](BlockVolumesAPI.md#UpdateVolumeStats) | **Post** /block-volumes/:update-stats | 



## BatchGetBlockVolumeSamples

> MultiVolumesSamplesResp BatchGetBlockVolumeSamples(ctx).Ids(ids).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	ids := "ids_example" // string | volume ids; example: id1,id2,id3
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.BatchGetBlockVolumeSamples(context.Background()).Ids(ids).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.BatchGetBlockVolumeSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchGetBlockVolumeSamples`: MultiVolumesSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.BatchGetBlockVolumeSamples`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchGetBlockVolumeSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | volume ids; example: id1,id2,id3 | 
 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**MultiVolumesSamplesResp**](MultiVolumesSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBlockVolume

> VolumeResp CreateBlockVolume(ctx).Body(body).Execute()





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
	body := *openapiclient.NewVolumeCreateReq(*openapiclient.NewVolumeCreateReqVolume("Name_example", int64(123))) // VolumeCreateReq | volume info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.CreateBlockVolume(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.CreateBlockVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlockVolume`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.CreateBlockVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlockVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VolumeCreateReq**](VolumeCreateReq.md) | volume info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlockVolume

> VolumeResp DeleteBlockVolume(ctx, blockVolumeId).BypassTrash(bypassTrash).Execute()





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
	blockVolumeId := int64(789) // int64 | volume id
	bypassTrash := true // bool | bypass trash or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.DeleteBlockVolume(context.Background(), blockVolumeId).BypassTrash(bypassTrash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.DeleteBlockVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlockVolume`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.DeleteBlockVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlockVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bypassTrash** | **bool** | bypass trash or not | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockVolume

> VolumeResp GetBlockVolume(ctx, blockVolumeId).Execute()





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
	blockVolumeId := int64(789) // int64 | block volume id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.GetBlockVolume(context.Background(), blockVolumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.GetBlockVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockVolume`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.GetBlockVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockVolumeSamples

> VolumeSamplesResp GetBlockVolumeSamples(ctx, blockVolumeId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	blockVolumeId := int64(789) // int64 | block volume id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.GetBlockVolumeSamples(context.Background(), blockVolumeId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.GetBlockVolumeSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockVolumeSamples`: VolumeSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.GetBlockVolumeSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockVolumeSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**VolumeSamplesResp**](VolumeSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlockVolumes

> VolumesResp ListBlockVolumes(ctx).Limit(limit).Offset(offset).PoolId(poolId).PoolName(poolName).ClusterId(clusterId).BlockSnapshotId(blockSnapshotId).Name(name).NamePrefix(namePrefix).VolumeName(volumeName).Uid(uid).ClientGroupId(clientGroupId).MappingGroupId(mappingGroupId).ExcludeMappingGroupId(excludeMappingGroupId).AccessPathId(accessPathId).Passive(passive).Recycled(recycled).Status(status).WithNotUsed(withNotUsed).Q(q).Sort(sort).All(all).DpBlockBackupPolicyId(dpBlockBackupPolicyId).DpBlockAsyncReplicationPolicyId(dpBlockAsyncReplicationPolicyId).CouldHaveIo(couldHaveIo).RbdClientId(rbdClientId).Execute()





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
	poolId := int64(789) // int64 | The id of the pool volumes belong to (optional)
	poolName := "poolName_example" // string | The pool_name of the pool volumes belong to (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)
	blockSnapshotId := int64(789) // int64 | related snapshot id (optional)
	name := "name_example" // string | name of volume (optional)
	namePrefix := "namePrefix_example" // string | name prefix of volume (optional)
	volumeName := "volumeName_example" // string | volume_name of volume (optional)
	uid := "uid_example" // string | uid of volume (optional)
	clientGroupId := int64(789) // int64 | related client group id (optional)
	mappingGroupId := int64(789) // int64 | related mapping group id (optional)
	excludeMappingGroupId := int64(789) // int64 | exclude mapping group id, with access path id (optional)
	accessPathId := int64(789) // int64 | related access path id (optional)
	passive := true // bool | passive or not (optional)
	recycled := true // bool | recycled or not (optional)
	status := "status_example" // string | filter with status (optional)
	withNotUsed := true // bool | list with not used volumes, can be used with access path id (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	all := true // bool | show all volumes (optional)
	dpBlockBackupPolicyId := int64(789) // int64 | show volumes of specific dp block backup policy (optional)
	dpBlockAsyncReplicationPolicyId := int64(789) // int64 | show volumes of specific dp block async replication policy (optional)
	couldHaveIo := true // bool | show volumes without volume that cannot have io (optional)
	rbdClientId := int64(789) // int64 | rbd client id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.ListBlockVolumes(context.Background()).Limit(limit).Offset(offset).PoolId(poolId).PoolName(poolName).ClusterId(clusterId).BlockSnapshotId(blockSnapshotId).Name(name).NamePrefix(namePrefix).VolumeName(volumeName).Uid(uid).ClientGroupId(clientGroupId).MappingGroupId(mappingGroupId).ExcludeMappingGroupId(excludeMappingGroupId).AccessPathId(accessPathId).Passive(passive).Recycled(recycled).Status(status).WithNotUsed(withNotUsed).Q(q).Sort(sort).All(all).DpBlockBackupPolicyId(dpBlockBackupPolicyId).DpBlockAsyncReplicationPolicyId(dpBlockAsyncReplicationPolicyId).CouldHaveIo(couldHaveIo).RbdClientId(rbdClientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.ListBlockVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlockVolumes`: VolumesResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.ListBlockVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlockVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **poolId** | **int64** | The id of the pool volumes belong to | 
 **poolName** | **string** | The pool_name of the pool volumes belong to | 
 **clusterId** | **string** | cluster id | 
 **blockSnapshotId** | **int64** | related snapshot id | 
 **name** | **string** | name of volume | 
 **namePrefix** | **string** | name prefix of volume | 
 **volumeName** | **string** | volume_name of volume | 
 **uid** | **string** | uid of volume | 
 **clientGroupId** | **int64** | related client group id | 
 **mappingGroupId** | **int64** | related mapping group id | 
 **excludeMappingGroupId** | **int64** | exclude mapping group id, with access path id | 
 **accessPathId** | **int64** | related access path id | 
 **passive** | **bool** | passive or not | 
 **recycled** | **bool** | recycled or not | 
 **status** | **string** | filter with status | 
 **withNotUsed** | **bool** | list with not used volumes, can be used with access path id | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **all** | **bool** | show all volumes | 
 **dpBlockBackupPolicyId** | **int64** | show volumes of specific dp block backup policy | 
 **dpBlockAsyncReplicationPolicyId** | **int64** | show volumes of specific dp block async replication policy | 
 **couldHaveIo** | **bool** | show volumes without volume that cannot have io | 
 **rbdClientId** | **int64** | rbd client id | 

### Return type

[**VolumesResp**](VolumesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrateBlockVolume

> VolumeResp MigrateBlockVolume(ctx, blockVolumeId).Body(body).Execute()



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
	blockVolumeId := int64(789) // int64 | the block volume id
	body := *openapiclient.NewVolumeMigrateReq() // VolumeMigrateReq | volume migration info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.MigrateBlockVolume(context.Background(), blockVolumeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.MigrateBlockVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrateBlockVolume`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.MigrateBlockVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | the block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrateBlockVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VolumeMigrateReq**](VolumeMigrateReq.md) | volume migration info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebuildBlockVolumeReplication

> VolumeResp RebuildBlockVolumeReplication(ctx, blockVolumeId).Force(force).Execute()





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
	blockVolumeId := int64(789) // int64 | block volume id
	force := true // bool | force rebuild or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.RebuildBlockVolumeReplication(context.Background(), blockVolumeId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.RebuildBlockVolumeReplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebuildBlockVolumeReplication`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.RebuildBlockVolumeReplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebuildBlockVolumeReplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force rebuild or not | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAsyncReplicationProtection

> VolumeResp SetAsyncReplicationProtection(ctx, blockVolumeId).Body(body).Execute()





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
	blockVolumeId := int64(789) // int64 | the block volume id
	body := *openapiclient.NewVolumeAsyncReplicationProtectionReq(int64(123), "DestPoolName_example", "DestVolumeName_example", int64(123)) // VolumeAsyncReplicationProtectionReq | request info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.SetAsyncReplicationProtection(context.Background(), blockVolumeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.SetAsyncReplicationProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAsyncReplicationProtection`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.SetAsyncReplicationProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | the block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAsyncReplicationProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VolumeAsyncReplicationProtectionReq**](VolumeAsyncReplicationProtectionReq.md) | request info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBackupProtection

> VolumeResp SetBackupProtection(ctx, blockVolumeId).Body(body).Execute()





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
	blockVolumeId := int64(789) // int64 | the block volume id
	body := *openapiclient.NewVolumeBackupProtectionReq() // VolumeBackupProtectionReq | request info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.SetBackupProtection(context.Background(), blockVolumeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.SetBackupProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetBackupProtection`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.SetBackupProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | the block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetBackupProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VolumeBackupProtectionReq**](VolumeBackupProtectionReq.md) | request info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBlockVolumeCrc

> VolumeResp SetBlockVolumeCrc(ctx, blockVolumeId).Body(body).Execute()





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
	blockVolumeId := int64(789) // int64 | block volume id
	body := *openapiclient.NewVolumeCrcSetReq() // VolumeCrcSetReq | volume crc info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.SetBlockVolumeCrc(context.Background(), blockVolumeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.SetBlockVolumeCrc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetBlockVolumeCrc`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.SetBlockVolumeCrc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetBlockVolumeCrcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VolumeCrcSetReq**](VolumeCrcSetReq.md) | volume crc info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBlockVolumeReplication

> VolumeResp SetBlockVolumeReplication(ctx, blockVolumeId).Body(body).Execute()





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
	blockVolumeId := int64(789) // int64 | block volume id
	body := *openapiclient.NewVolumeReplicationSetReq() // VolumeReplicationSetReq | volume replication info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.SetBlockVolumeReplication(context.Background(), blockVolumeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.SetBlockVolumeReplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetBlockVolumeReplication`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.SetBlockVolumeReplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetBlockVolumeReplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VolumeReplicationSetReq**](VolumeReplicationSetReq.md) | volume replication info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSnapshotProtection

> VolumeResp SetSnapshotProtection(ctx, blockVolumeId).Body(body).Execute()





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
	blockVolumeId := int64(789) // int64 | the block volume id
	body := *openapiclient.NewVolumeSnapshotProtectionReq() // VolumeSnapshotProtectionReq | request info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.SetSnapshotProtection(context.Background(), blockVolumeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.SetSnapshotProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSnapshotProtection`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.SetSnapshotProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | the block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSnapshotProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VolumeSnapshotProtectionReq**](VolumeSnapshotProtectionReq.md) | request info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendBlockVolumeReplication

> VolumeResp SuspendBlockVolumeReplication(ctx, blockVolumeId).Execute()





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
	blockVolumeId := int64(789) // int64 | block volume id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.SuspendBlockVolumeReplication(context.Background(), blockVolumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.SuspendBlockVolumeReplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuspendBlockVolumeReplication`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.SuspendBlockVolumeReplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendBlockVolumeReplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetAsyncReplicationProtection

> VolumeResp UnsetAsyncReplicationProtection(ctx, blockVolumeId).Force(force).ReserveVolume(reserveVolume).Execute()





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
	blockVolumeId := int64(789) // int64 | the block volume id
	force := true // bool | force unset or not (optional)
	reserveVolume := true // bool | reserve replicated volume or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.UnsetAsyncReplicationProtection(context.Background(), blockVolumeId).Force(force).ReserveVolume(reserveVolume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.UnsetAsyncReplicationProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsetAsyncReplicationProtection`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.UnsetAsyncReplicationProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | the block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetAsyncReplicationProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force unset or not | 
 **reserveVolume** | **bool** | reserve replicated volume or not | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetBackupProtection

> VolumeResp UnsetBackupProtection(ctx, blockVolumeId).Force(force).Execute()





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
	blockVolumeId := int64(789) // int64 | the block volume id
	force := true // bool | force unset or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.UnsetBackupProtection(context.Background(), blockVolumeId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.UnsetBackupProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsetBackupProtection`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.UnsetBackupProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | the block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetBackupProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force unset or not | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetBlockVolumeCrc

> VolumeResp UnsetBlockVolumeCrc(ctx, blockVolumeId).Body(body).Execute()





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
	blockVolumeId := int64(789) // int64 | block volume id
	body := *openapiclient.NewVolumeCrcSetReq() // VolumeCrcSetReq | volume crc info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.UnsetBlockVolumeCrc(context.Background(), blockVolumeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.UnsetBlockVolumeCrc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsetBlockVolumeCrc`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.UnsetBlockVolumeCrc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetBlockVolumeCrcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VolumeCrcSetReq**](VolumeCrcSetReq.md) | volume crc info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetBlockVolumeReplication

> VolumeResp UnsetBlockVolumeReplication(ctx, blockVolumeId).Execute()





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
	blockVolumeId := int64(789) // int64 | block volume id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.UnsetBlockVolumeReplication(context.Background(), blockVolumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.UnsetBlockVolumeReplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsetBlockVolumeReplication`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.UnsetBlockVolumeReplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetBlockVolumeReplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetSnapshotProtection

> VolumeResp UnsetSnapshotProtection(ctx, blockVolumeId).Force(force).Execute()





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
	blockVolumeId := int64(789) // int64 | the block volume id
	force := true // bool | force unset or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.UnsetSnapshotProtection(context.Background(), blockVolumeId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.UnsetSnapshotProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsetSnapshotProtection`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.UnsetSnapshotProtection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | the block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetSnapshotProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force unset or not | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlockVolume

> VolumeResp UpdateBlockVolume(ctx, blockVolumeId).Body(body).Execute()





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
	blockVolumeId := int64(789) // int64 | the block volume id
	body := *openapiclient.NewVolumeUpdateReq(*openapiclient.NewVolumeUpdateReqVolume()) // VolumeUpdateReq | volume info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.UpdateBlockVolume(context.Background(), blockVolumeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.UpdateBlockVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBlockVolume`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.UpdateBlockVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | the block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlockVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VolumeUpdateReq**](VolumeUpdateReq.md) | volume info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlockVolumeVolumeName

> VolumeResp UpdateBlockVolumeVolumeName(ctx, blockVolumeId).Body(body).Execute()





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
	blockVolumeId := int64(789) // int64 | the block volume id
	body := *openapiclient.NewVolumeUpdateVolumeNameReq(*openapiclient.NewVolumeUpdateVolumeNameReqVolume("VolumeName_example")) // VolumeUpdateVolumeNameReq | volume volume_name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockVolumesAPI.UpdateBlockVolumeVolumeName(context.Background(), blockVolumeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.UpdateBlockVolumeVolumeName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBlockVolumeVolumeName`: VolumeResp
	fmt.Fprintf(os.Stdout, "Response from `BlockVolumesAPI.UpdateBlockVolumeVolumeName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockVolumeId** | **int64** | the block volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlockVolumeVolumeNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VolumeUpdateVolumeNameReq**](VolumeUpdateVolumeNameReq.md) | volume volume_name | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVolumeStat

> UpdateVolumeStat(ctx).Body(body).Execute()





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
	body := *openapiclient.NewUpdateVolumeStatReq() // UpdateVolumeStatReq | volume stat

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlockVolumesAPI.UpdateVolumeStat(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.UpdateVolumeStat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeStatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateVolumeStatReq**](UpdateVolumeStatReq.md) | volume stat | 

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


## UpdateVolumeStats

> UpdateVolumeStats(ctx).Body(body).Execute()





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
	body := *openapiclient.NewUpdateVolumeStatsReq() // UpdateVolumeStatsReq | volume stats

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlockVolumesAPI.UpdateVolumeStats(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockVolumesAPI.UpdateVolumeStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateVolumeStatsReq**](UpdateVolumeStatsReq.md) | volume stats | 

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

