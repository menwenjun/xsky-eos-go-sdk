# DpSiteConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockAsyncReplication** | Pointer to [**AsyncReplicationConfig**](AsyncReplicationConfig.md) |  | [optional] 
**BlockSnapshotBackup** | Pointer to [**SnapshotBackupConfig**](SnapshotBackupConfig.md) |  | [optional] 
**BlockVolumeReplication** | Pointer to [**BlockReplicationConfig**](BlockReplicationConfig.md) |  | [optional] 

## Methods

### NewDpSiteConfig

`func NewDpSiteConfig() *DpSiteConfig`

NewDpSiteConfig instantiates a new DpSiteConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpSiteConfigWithDefaults

`func NewDpSiteConfigWithDefaults() *DpSiteConfig`

NewDpSiteConfigWithDefaults instantiates a new DpSiteConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockAsyncReplication

`func (o *DpSiteConfig) GetBlockAsyncReplication() AsyncReplicationConfig`

GetBlockAsyncReplication returns the BlockAsyncReplication field if non-nil, zero value otherwise.

### GetBlockAsyncReplicationOk

`func (o *DpSiteConfig) GetBlockAsyncReplicationOk() (*AsyncReplicationConfig, bool)`

GetBlockAsyncReplicationOk returns a tuple with the BlockAsyncReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAsyncReplication

`func (o *DpSiteConfig) SetBlockAsyncReplication(v AsyncReplicationConfig)`

SetBlockAsyncReplication sets BlockAsyncReplication field to given value.

### HasBlockAsyncReplication

`func (o *DpSiteConfig) HasBlockAsyncReplication() bool`

HasBlockAsyncReplication returns a boolean if a field has been set.

### GetBlockSnapshotBackup

`func (o *DpSiteConfig) GetBlockSnapshotBackup() SnapshotBackupConfig`

GetBlockSnapshotBackup returns the BlockSnapshotBackup field if non-nil, zero value otherwise.

### GetBlockSnapshotBackupOk

`func (o *DpSiteConfig) GetBlockSnapshotBackupOk() (*SnapshotBackupConfig, bool)`

GetBlockSnapshotBackupOk returns a tuple with the BlockSnapshotBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSnapshotBackup

`func (o *DpSiteConfig) SetBlockSnapshotBackup(v SnapshotBackupConfig)`

SetBlockSnapshotBackup sets BlockSnapshotBackup field to given value.

### HasBlockSnapshotBackup

`func (o *DpSiteConfig) HasBlockSnapshotBackup() bool`

HasBlockSnapshotBackup returns a boolean if a field has been set.

### GetBlockVolumeReplication

`func (o *DpSiteConfig) GetBlockVolumeReplication() BlockReplicationConfig`

GetBlockVolumeReplication returns the BlockVolumeReplication field if non-nil, zero value otherwise.

### GetBlockVolumeReplicationOk

`func (o *DpSiteConfig) GetBlockVolumeReplicationOk() (*BlockReplicationConfig, bool)`

GetBlockVolumeReplicationOk returns a tuple with the BlockVolumeReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeReplication

`func (o *DpSiteConfig) SetBlockVolumeReplication(v BlockReplicationConfig)`

SetBlockVolumeReplication sets BlockVolumeReplication field to given value.

### HasBlockVolumeReplication

`func (o *DpSiteConfig) HasBlockVolumeReplication() bool`

HasBlockVolumeReplication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


