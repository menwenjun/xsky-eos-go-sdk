# VolumeCreateReqVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockSnapshotId** | Pointer to **int64** | id of related block volume snapshot | [optional] 
**CrcCheck** | Pointer to **bool** | crc check or not | [optional] 
**Description** | Pointer to **string** | description of volume | [optional] 
**Flattened** | Pointer to **bool** | flatten or not flatten | [optional] 
**Format** | Pointer to **int64** | volume format: { 128 | 129 (advanced) }, default 128 | [optional] 
**Name** | **string** | name of volume | 
**PerformancePriority** | Pointer to **int64** | performance priority: { 0 | 1 }, default 0 | [optional] 
**PoolId** | **int64** | id of pool belonged to | 
**Qos** | Pointer to [**VolumeQosSpec**](VolumeQosSpec.md) |  | [optional] 
**QosEnabled** | Pointer to **bool** | enable or disable the qos | [optional] 
**RemoteClusterFsId** | Pointer to **string** | replication remote cluster fsid | [optional] 
**ReplicationPool** | Pointer to **string** | replication peer pool | [optional] 
**ReplicationPoolId** | Pointer to **int64** | replication peer pool id | [optional] 
**ReplicationPoolName** | Pointer to **string** | replication peer pool name | [optional] 
**ReplicationVersion** | Pointer to **int64** | replication version | [optional] 
**ReplicationVolume** | Pointer to **string** | replication peer volume | [optional] 
**ReplicationVolumeId** | Pointer to **int64** | replication peer volume id | [optional] 
**ReplicationVolumeName** | Pointer to **string** | replication peer volume name | [optional] 
**Size** | Pointer to **int64** | size of volume | [optional] 
**Sn** | Pointer to **string** | volume sn, used when creating replication volume | [optional] 
**SnapshotReplicationPool** | Pointer to **string** | snapshot replication peer pool | [optional] 
**SnapshotReplicationPoolId** | Pointer to **int64** | snapshot replication peer pool id | [optional] 
**SnapshotReplicationVolume** | Pointer to **string** | snapshot replication peer volume | [optional] 
**SnapshotReplicationVolumeId** | Pointer to **int64** | snapshot replication peer volume id | [optional] 
**Uid** | Pointer to **string** | UID of volume | [optional] 

## Methods

### NewVolumeCreateReqVolume

`func NewVolumeCreateReqVolume(name string, poolId int64, ) *VolumeCreateReqVolume`

NewVolumeCreateReqVolume instantiates a new VolumeCreateReqVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeCreateReqVolumeWithDefaults

`func NewVolumeCreateReqVolumeWithDefaults() *VolumeCreateReqVolume`

NewVolumeCreateReqVolumeWithDefaults instantiates a new VolumeCreateReqVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockSnapshotId

`func (o *VolumeCreateReqVolume) GetBlockSnapshotId() int64`

GetBlockSnapshotId returns the BlockSnapshotId field if non-nil, zero value otherwise.

### GetBlockSnapshotIdOk

`func (o *VolumeCreateReqVolume) GetBlockSnapshotIdOk() (*int64, bool)`

GetBlockSnapshotIdOk returns a tuple with the BlockSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSnapshotId

`func (o *VolumeCreateReqVolume) SetBlockSnapshotId(v int64)`

SetBlockSnapshotId sets BlockSnapshotId field to given value.

### HasBlockSnapshotId

`func (o *VolumeCreateReqVolume) HasBlockSnapshotId() bool`

HasBlockSnapshotId returns a boolean if a field has been set.

### GetCrcCheck

`func (o *VolumeCreateReqVolume) GetCrcCheck() bool`

GetCrcCheck returns the CrcCheck field if non-nil, zero value otherwise.

### GetCrcCheckOk

`func (o *VolumeCreateReqVolume) GetCrcCheckOk() (*bool, bool)`

GetCrcCheckOk returns a tuple with the CrcCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrcCheck

`func (o *VolumeCreateReqVolume) SetCrcCheck(v bool)`

SetCrcCheck sets CrcCheck field to given value.

### HasCrcCheck

`func (o *VolumeCreateReqVolume) HasCrcCheck() bool`

HasCrcCheck returns a boolean if a field has been set.

### GetDescription

`func (o *VolumeCreateReqVolume) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VolumeCreateReqVolume) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VolumeCreateReqVolume) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VolumeCreateReqVolume) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFlattened

`func (o *VolumeCreateReqVolume) GetFlattened() bool`

GetFlattened returns the Flattened field if non-nil, zero value otherwise.

### GetFlattenedOk

`func (o *VolumeCreateReqVolume) GetFlattenedOk() (*bool, bool)`

GetFlattenedOk returns a tuple with the Flattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattened

`func (o *VolumeCreateReqVolume) SetFlattened(v bool)`

SetFlattened sets Flattened field to given value.

### HasFlattened

`func (o *VolumeCreateReqVolume) HasFlattened() bool`

HasFlattened returns a boolean if a field has been set.

### GetFormat

`func (o *VolumeCreateReqVolume) GetFormat() int64`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *VolumeCreateReqVolume) GetFormatOk() (*int64, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *VolumeCreateReqVolume) SetFormat(v int64)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *VolumeCreateReqVolume) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *VolumeCreateReqVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeCreateReqVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeCreateReqVolume) SetName(v string)`

SetName sets Name field to given value.


### GetPerformancePriority

`func (o *VolumeCreateReqVolume) GetPerformancePriority() int64`

GetPerformancePriority returns the PerformancePriority field if non-nil, zero value otherwise.

### GetPerformancePriorityOk

`func (o *VolumeCreateReqVolume) GetPerformancePriorityOk() (*int64, bool)`

GetPerformancePriorityOk returns a tuple with the PerformancePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformancePriority

`func (o *VolumeCreateReqVolume) SetPerformancePriority(v int64)`

SetPerformancePriority sets PerformancePriority field to given value.

### HasPerformancePriority

`func (o *VolumeCreateReqVolume) HasPerformancePriority() bool`

HasPerformancePriority returns a boolean if a field has been set.

### GetPoolId

`func (o *VolumeCreateReqVolume) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *VolumeCreateReqVolume) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *VolumeCreateReqVolume) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.


### GetQos

`func (o *VolumeCreateReqVolume) GetQos() VolumeQosSpec`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *VolumeCreateReqVolume) GetQosOk() (*VolumeQosSpec, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *VolumeCreateReqVolume) SetQos(v VolumeQosSpec)`

SetQos sets Qos field to given value.

### HasQos

`func (o *VolumeCreateReqVolume) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetQosEnabled

`func (o *VolumeCreateReqVolume) GetQosEnabled() bool`

GetQosEnabled returns the QosEnabled field if non-nil, zero value otherwise.

### GetQosEnabledOk

`func (o *VolumeCreateReqVolume) GetQosEnabledOk() (*bool, bool)`

GetQosEnabledOk returns a tuple with the QosEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosEnabled

`func (o *VolumeCreateReqVolume) SetQosEnabled(v bool)`

SetQosEnabled sets QosEnabled field to given value.

### HasQosEnabled

`func (o *VolumeCreateReqVolume) HasQosEnabled() bool`

HasQosEnabled returns a boolean if a field has been set.

### GetRemoteClusterFsId

`func (o *VolumeCreateReqVolume) GetRemoteClusterFsId() string`

GetRemoteClusterFsId returns the RemoteClusterFsId field if non-nil, zero value otherwise.

### GetRemoteClusterFsIdOk

`func (o *VolumeCreateReqVolume) GetRemoteClusterFsIdOk() (*string, bool)`

GetRemoteClusterFsIdOk returns a tuple with the RemoteClusterFsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClusterFsId

`func (o *VolumeCreateReqVolume) SetRemoteClusterFsId(v string)`

SetRemoteClusterFsId sets RemoteClusterFsId field to given value.

### HasRemoteClusterFsId

`func (o *VolumeCreateReqVolume) HasRemoteClusterFsId() bool`

HasRemoteClusterFsId returns a boolean if a field has been set.

### GetReplicationPool

`func (o *VolumeCreateReqVolume) GetReplicationPool() string`

GetReplicationPool returns the ReplicationPool field if non-nil, zero value otherwise.

### GetReplicationPoolOk

`func (o *VolumeCreateReqVolume) GetReplicationPoolOk() (*string, bool)`

GetReplicationPoolOk returns a tuple with the ReplicationPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPool

`func (o *VolumeCreateReqVolume) SetReplicationPool(v string)`

SetReplicationPool sets ReplicationPool field to given value.

### HasReplicationPool

`func (o *VolumeCreateReqVolume) HasReplicationPool() bool`

HasReplicationPool returns a boolean if a field has been set.

### GetReplicationPoolId

`func (o *VolumeCreateReqVolume) GetReplicationPoolId() int64`

GetReplicationPoolId returns the ReplicationPoolId field if non-nil, zero value otherwise.

### GetReplicationPoolIdOk

`func (o *VolumeCreateReqVolume) GetReplicationPoolIdOk() (*int64, bool)`

GetReplicationPoolIdOk returns a tuple with the ReplicationPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPoolId

`func (o *VolumeCreateReqVolume) SetReplicationPoolId(v int64)`

SetReplicationPoolId sets ReplicationPoolId field to given value.

### HasReplicationPoolId

`func (o *VolumeCreateReqVolume) HasReplicationPoolId() bool`

HasReplicationPoolId returns a boolean if a field has been set.

### GetReplicationPoolName

`func (o *VolumeCreateReqVolume) GetReplicationPoolName() string`

GetReplicationPoolName returns the ReplicationPoolName field if non-nil, zero value otherwise.

### GetReplicationPoolNameOk

`func (o *VolumeCreateReqVolume) GetReplicationPoolNameOk() (*string, bool)`

GetReplicationPoolNameOk returns a tuple with the ReplicationPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPoolName

`func (o *VolumeCreateReqVolume) SetReplicationPoolName(v string)`

SetReplicationPoolName sets ReplicationPoolName field to given value.

### HasReplicationPoolName

`func (o *VolumeCreateReqVolume) HasReplicationPoolName() bool`

HasReplicationPoolName returns a boolean if a field has been set.

### GetReplicationVersion

`func (o *VolumeCreateReqVolume) GetReplicationVersion() int64`

GetReplicationVersion returns the ReplicationVersion field if non-nil, zero value otherwise.

### GetReplicationVersionOk

`func (o *VolumeCreateReqVolume) GetReplicationVersionOk() (*int64, bool)`

GetReplicationVersionOk returns a tuple with the ReplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVersion

`func (o *VolumeCreateReqVolume) SetReplicationVersion(v int64)`

SetReplicationVersion sets ReplicationVersion field to given value.

### HasReplicationVersion

`func (o *VolumeCreateReqVolume) HasReplicationVersion() bool`

HasReplicationVersion returns a boolean if a field has been set.

### GetReplicationVolume

`func (o *VolumeCreateReqVolume) GetReplicationVolume() string`

GetReplicationVolume returns the ReplicationVolume field if non-nil, zero value otherwise.

### GetReplicationVolumeOk

`func (o *VolumeCreateReqVolume) GetReplicationVolumeOk() (*string, bool)`

GetReplicationVolumeOk returns a tuple with the ReplicationVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVolume

`func (o *VolumeCreateReqVolume) SetReplicationVolume(v string)`

SetReplicationVolume sets ReplicationVolume field to given value.

### HasReplicationVolume

`func (o *VolumeCreateReqVolume) HasReplicationVolume() bool`

HasReplicationVolume returns a boolean if a field has been set.

### GetReplicationVolumeId

`func (o *VolumeCreateReqVolume) GetReplicationVolumeId() int64`

GetReplicationVolumeId returns the ReplicationVolumeId field if non-nil, zero value otherwise.

### GetReplicationVolumeIdOk

`func (o *VolumeCreateReqVolume) GetReplicationVolumeIdOk() (*int64, bool)`

GetReplicationVolumeIdOk returns a tuple with the ReplicationVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVolumeId

`func (o *VolumeCreateReqVolume) SetReplicationVolumeId(v int64)`

SetReplicationVolumeId sets ReplicationVolumeId field to given value.

### HasReplicationVolumeId

`func (o *VolumeCreateReqVolume) HasReplicationVolumeId() bool`

HasReplicationVolumeId returns a boolean if a field has been set.

### GetReplicationVolumeName

`func (o *VolumeCreateReqVolume) GetReplicationVolumeName() string`

GetReplicationVolumeName returns the ReplicationVolumeName field if non-nil, zero value otherwise.

### GetReplicationVolumeNameOk

`func (o *VolumeCreateReqVolume) GetReplicationVolumeNameOk() (*string, bool)`

GetReplicationVolumeNameOk returns a tuple with the ReplicationVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVolumeName

`func (o *VolumeCreateReqVolume) SetReplicationVolumeName(v string)`

SetReplicationVolumeName sets ReplicationVolumeName field to given value.

### HasReplicationVolumeName

`func (o *VolumeCreateReqVolume) HasReplicationVolumeName() bool`

HasReplicationVolumeName returns a boolean if a field has been set.

### GetSize

`func (o *VolumeCreateReqVolume) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumeCreateReqVolume) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumeCreateReqVolume) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *VolumeCreateReqVolume) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSn

`func (o *VolumeCreateReqVolume) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *VolumeCreateReqVolume) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *VolumeCreateReqVolume) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *VolumeCreateReqVolume) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetSnapshotReplicationPool

`func (o *VolumeCreateReqVolume) GetSnapshotReplicationPool() string`

GetSnapshotReplicationPool returns the SnapshotReplicationPool field if non-nil, zero value otherwise.

### GetSnapshotReplicationPoolOk

`func (o *VolumeCreateReqVolume) GetSnapshotReplicationPoolOk() (*string, bool)`

GetSnapshotReplicationPoolOk returns a tuple with the SnapshotReplicationPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotReplicationPool

`func (o *VolumeCreateReqVolume) SetSnapshotReplicationPool(v string)`

SetSnapshotReplicationPool sets SnapshotReplicationPool field to given value.

### HasSnapshotReplicationPool

`func (o *VolumeCreateReqVolume) HasSnapshotReplicationPool() bool`

HasSnapshotReplicationPool returns a boolean if a field has been set.

### GetSnapshotReplicationPoolId

`func (o *VolumeCreateReqVolume) GetSnapshotReplicationPoolId() int64`

GetSnapshotReplicationPoolId returns the SnapshotReplicationPoolId field if non-nil, zero value otherwise.

### GetSnapshotReplicationPoolIdOk

`func (o *VolumeCreateReqVolume) GetSnapshotReplicationPoolIdOk() (*int64, bool)`

GetSnapshotReplicationPoolIdOk returns a tuple with the SnapshotReplicationPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotReplicationPoolId

`func (o *VolumeCreateReqVolume) SetSnapshotReplicationPoolId(v int64)`

SetSnapshotReplicationPoolId sets SnapshotReplicationPoolId field to given value.

### HasSnapshotReplicationPoolId

`func (o *VolumeCreateReqVolume) HasSnapshotReplicationPoolId() bool`

HasSnapshotReplicationPoolId returns a boolean if a field has been set.

### GetSnapshotReplicationVolume

`func (o *VolumeCreateReqVolume) GetSnapshotReplicationVolume() string`

GetSnapshotReplicationVolume returns the SnapshotReplicationVolume field if non-nil, zero value otherwise.

### GetSnapshotReplicationVolumeOk

`func (o *VolumeCreateReqVolume) GetSnapshotReplicationVolumeOk() (*string, bool)`

GetSnapshotReplicationVolumeOk returns a tuple with the SnapshotReplicationVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotReplicationVolume

`func (o *VolumeCreateReqVolume) SetSnapshotReplicationVolume(v string)`

SetSnapshotReplicationVolume sets SnapshotReplicationVolume field to given value.

### HasSnapshotReplicationVolume

`func (o *VolumeCreateReqVolume) HasSnapshotReplicationVolume() bool`

HasSnapshotReplicationVolume returns a boolean if a field has been set.

### GetSnapshotReplicationVolumeId

`func (o *VolumeCreateReqVolume) GetSnapshotReplicationVolumeId() int64`

GetSnapshotReplicationVolumeId returns the SnapshotReplicationVolumeId field if non-nil, zero value otherwise.

### GetSnapshotReplicationVolumeIdOk

`func (o *VolumeCreateReqVolume) GetSnapshotReplicationVolumeIdOk() (*int64, bool)`

GetSnapshotReplicationVolumeIdOk returns a tuple with the SnapshotReplicationVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotReplicationVolumeId

`func (o *VolumeCreateReqVolume) SetSnapshotReplicationVolumeId(v int64)`

SetSnapshotReplicationVolumeId sets SnapshotReplicationVolumeId field to given value.

### HasSnapshotReplicationVolumeId

`func (o *VolumeCreateReqVolume) HasSnapshotReplicationVolumeId() bool`

HasSnapshotReplicationVolumeId returns a boolean if a field has been set.

### GetUid

`func (o *VolumeCreateReqVolume) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *VolumeCreateReqVolume) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *VolumeCreateReqVolume) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *VolumeCreateReqVolume) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


