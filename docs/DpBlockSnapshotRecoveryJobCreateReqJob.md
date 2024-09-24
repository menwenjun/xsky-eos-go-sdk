# DpBlockSnapshotRecoveryJobCreateReqJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupBlockSnapshot** | [**DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot**](DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot.md) |  | 
**BackupBlockVolume** | [**DpBlockSnapshotRecoveryJobCreateReqJobBackupVolume**](DpBlockSnapshotRecoveryJobCreateReqJobBackupVolume.md) |  | 
**BackupCluster** | [**DpBlockSnapshotRecoveryJobCreateReqJobBackupCluster**](DpBlockSnapshotRecoveryJobCreateReqJobBackupCluster.md) |  | 
**BlockVolume** | [**DpBlockSnapshotRecoveryJobCreateReqJobVolume**](DpBlockSnapshotRecoveryJobCreateReqJobVolume.md) |  | 
**DpGatewayId** | **int64** |  | 
**DpSiteId** | **int64** |  | 
**ResourceType** | **string** |  | 

## Methods

### NewDpBlockSnapshotRecoveryJobCreateReqJob

`func NewDpBlockSnapshotRecoveryJobCreateReqJob(backupBlockSnapshot DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot, backupBlockVolume DpBlockSnapshotRecoveryJobCreateReqJobBackupVolume, backupCluster DpBlockSnapshotRecoveryJobCreateReqJobBackupCluster, blockVolume DpBlockSnapshotRecoveryJobCreateReqJobVolume, dpGatewayId int64, dpSiteId int64, resourceType string, ) *DpBlockSnapshotRecoveryJobCreateReqJob`

NewDpBlockSnapshotRecoveryJobCreateReqJob instantiates a new DpBlockSnapshotRecoveryJobCreateReqJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockSnapshotRecoveryJobCreateReqJobWithDefaults

`func NewDpBlockSnapshotRecoveryJobCreateReqJobWithDefaults() *DpBlockSnapshotRecoveryJobCreateReqJob`

NewDpBlockSnapshotRecoveryJobCreateReqJobWithDefaults instantiates a new DpBlockSnapshotRecoveryJobCreateReqJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupBlockSnapshot

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) GetBackupBlockSnapshot() DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot`

GetBackupBlockSnapshot returns the BackupBlockSnapshot field if non-nil, zero value otherwise.

### GetBackupBlockSnapshotOk

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) GetBackupBlockSnapshotOk() (*DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot, bool)`

GetBackupBlockSnapshotOk returns a tuple with the BackupBlockSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupBlockSnapshot

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) SetBackupBlockSnapshot(v DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot)`

SetBackupBlockSnapshot sets BackupBlockSnapshot field to given value.


### GetBackupBlockVolume

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) GetBackupBlockVolume() DpBlockSnapshotRecoveryJobCreateReqJobBackupVolume`

GetBackupBlockVolume returns the BackupBlockVolume field if non-nil, zero value otherwise.

### GetBackupBlockVolumeOk

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) GetBackupBlockVolumeOk() (*DpBlockSnapshotRecoveryJobCreateReqJobBackupVolume, bool)`

GetBackupBlockVolumeOk returns a tuple with the BackupBlockVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupBlockVolume

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) SetBackupBlockVolume(v DpBlockSnapshotRecoveryJobCreateReqJobBackupVolume)`

SetBackupBlockVolume sets BackupBlockVolume field to given value.


### GetBackupCluster

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) GetBackupCluster() DpBlockSnapshotRecoveryJobCreateReqJobBackupCluster`

GetBackupCluster returns the BackupCluster field if non-nil, zero value otherwise.

### GetBackupClusterOk

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) GetBackupClusterOk() (*DpBlockSnapshotRecoveryJobCreateReqJobBackupCluster, bool)`

GetBackupClusterOk returns a tuple with the BackupCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCluster

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) SetBackupCluster(v DpBlockSnapshotRecoveryJobCreateReqJobBackupCluster)`

SetBackupCluster sets BackupCluster field to given value.


### GetBlockVolume

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) GetBlockVolume() DpBlockSnapshotRecoveryJobCreateReqJobVolume`

GetBlockVolume returns the BlockVolume field if non-nil, zero value otherwise.

### GetBlockVolumeOk

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) GetBlockVolumeOk() (*DpBlockSnapshotRecoveryJobCreateReqJobVolume, bool)`

GetBlockVolumeOk returns a tuple with the BlockVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolume

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) SetBlockVolume(v DpBlockSnapshotRecoveryJobCreateReqJobVolume)`

SetBlockVolume sets BlockVolume field to given value.


### GetDpGatewayId

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) GetDpGatewayId() int64`

GetDpGatewayId returns the DpGatewayId field if non-nil, zero value otherwise.

### GetDpGatewayIdOk

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) GetDpGatewayIdOk() (*int64, bool)`

GetDpGatewayIdOk returns a tuple with the DpGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpGatewayId

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) SetDpGatewayId(v int64)`

SetDpGatewayId sets DpGatewayId field to given value.


### GetDpSiteId

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) GetDpSiteId() int64`

GetDpSiteId returns the DpSiteId field if non-nil, zero value otherwise.

### GetDpSiteIdOk

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) GetDpSiteIdOk() (*int64, bool)`

GetDpSiteIdOk returns a tuple with the DpSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpSiteId

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) SetDpSiteId(v int64)`

SetDpSiteId sets DpSiteId field to given value.


### GetResourceType

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DpBlockSnapshotRecoveryJobCreateReqJob) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


