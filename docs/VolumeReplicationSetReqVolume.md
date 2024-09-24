# VolumeReplicationSetReqVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestPoolId** | Pointer to **int64** | destination pool id | [optional] 
**DestVolumeName** | Pointer to **string** | destination volume name | [optional] 
**DpBlockReplicationPolicyId** | Pointer to **int64** | block replication policy id | [optional] 
**RemoteClusterFsId** | Pointer to **string** | remote cluster | [optional] 
**ReplicationPool** | Pointer to **string** | replication peer pool | [optional] 
**ReplicationPoolId** | Pointer to **int64** | replication peer pool id | [optional] 
**ReplicationPoolName** | Pointer to **string** | replication peer pool name | [optional] 
**ReplicationVersion** | Pointer to **int64** | replication version | [optional] 
**ReplicationVolume** | Pointer to **string** | replication peer volume | [optional] 
**ReplicationVolumeId** | Pointer to **int64** | replication peer volume id | [optional] 
**ReplicationVolumeName** | Pointer to **string** | replication peer volume name | [optional] 

## Methods

### NewVolumeReplicationSetReqVolume

`func NewVolumeReplicationSetReqVolume() *VolumeReplicationSetReqVolume`

NewVolumeReplicationSetReqVolume instantiates a new VolumeReplicationSetReqVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeReplicationSetReqVolumeWithDefaults

`func NewVolumeReplicationSetReqVolumeWithDefaults() *VolumeReplicationSetReqVolume`

NewVolumeReplicationSetReqVolumeWithDefaults instantiates a new VolumeReplicationSetReqVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestPoolId

`func (o *VolumeReplicationSetReqVolume) GetDestPoolId() int64`

GetDestPoolId returns the DestPoolId field if non-nil, zero value otherwise.

### GetDestPoolIdOk

`func (o *VolumeReplicationSetReqVolume) GetDestPoolIdOk() (*int64, bool)`

GetDestPoolIdOk returns a tuple with the DestPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPoolId

`func (o *VolumeReplicationSetReqVolume) SetDestPoolId(v int64)`

SetDestPoolId sets DestPoolId field to given value.

### HasDestPoolId

`func (o *VolumeReplicationSetReqVolume) HasDestPoolId() bool`

HasDestPoolId returns a boolean if a field has been set.

### GetDestVolumeName

`func (o *VolumeReplicationSetReqVolume) GetDestVolumeName() string`

GetDestVolumeName returns the DestVolumeName field if non-nil, zero value otherwise.

### GetDestVolumeNameOk

`func (o *VolumeReplicationSetReqVolume) GetDestVolumeNameOk() (*string, bool)`

GetDestVolumeNameOk returns a tuple with the DestVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestVolumeName

`func (o *VolumeReplicationSetReqVolume) SetDestVolumeName(v string)`

SetDestVolumeName sets DestVolumeName field to given value.

### HasDestVolumeName

`func (o *VolumeReplicationSetReqVolume) HasDestVolumeName() bool`

HasDestVolumeName returns a boolean if a field has been set.

### GetDpBlockReplicationPolicyId

`func (o *VolumeReplicationSetReqVolume) GetDpBlockReplicationPolicyId() int64`

GetDpBlockReplicationPolicyId returns the DpBlockReplicationPolicyId field if non-nil, zero value otherwise.

### GetDpBlockReplicationPolicyIdOk

`func (o *VolumeReplicationSetReqVolume) GetDpBlockReplicationPolicyIdOk() (*int64, bool)`

GetDpBlockReplicationPolicyIdOk returns a tuple with the DpBlockReplicationPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpBlockReplicationPolicyId

`func (o *VolumeReplicationSetReqVolume) SetDpBlockReplicationPolicyId(v int64)`

SetDpBlockReplicationPolicyId sets DpBlockReplicationPolicyId field to given value.

### HasDpBlockReplicationPolicyId

`func (o *VolumeReplicationSetReqVolume) HasDpBlockReplicationPolicyId() bool`

HasDpBlockReplicationPolicyId returns a boolean if a field has been set.

### GetRemoteClusterFsId

`func (o *VolumeReplicationSetReqVolume) GetRemoteClusterFsId() string`

GetRemoteClusterFsId returns the RemoteClusterFsId field if non-nil, zero value otherwise.

### GetRemoteClusterFsIdOk

`func (o *VolumeReplicationSetReqVolume) GetRemoteClusterFsIdOk() (*string, bool)`

GetRemoteClusterFsIdOk returns a tuple with the RemoteClusterFsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClusterFsId

`func (o *VolumeReplicationSetReqVolume) SetRemoteClusterFsId(v string)`

SetRemoteClusterFsId sets RemoteClusterFsId field to given value.

### HasRemoteClusterFsId

`func (o *VolumeReplicationSetReqVolume) HasRemoteClusterFsId() bool`

HasRemoteClusterFsId returns a boolean if a field has been set.

### GetReplicationPool

`func (o *VolumeReplicationSetReqVolume) GetReplicationPool() string`

GetReplicationPool returns the ReplicationPool field if non-nil, zero value otherwise.

### GetReplicationPoolOk

`func (o *VolumeReplicationSetReqVolume) GetReplicationPoolOk() (*string, bool)`

GetReplicationPoolOk returns a tuple with the ReplicationPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPool

`func (o *VolumeReplicationSetReqVolume) SetReplicationPool(v string)`

SetReplicationPool sets ReplicationPool field to given value.

### HasReplicationPool

`func (o *VolumeReplicationSetReqVolume) HasReplicationPool() bool`

HasReplicationPool returns a boolean if a field has been set.

### GetReplicationPoolId

`func (o *VolumeReplicationSetReqVolume) GetReplicationPoolId() int64`

GetReplicationPoolId returns the ReplicationPoolId field if non-nil, zero value otherwise.

### GetReplicationPoolIdOk

`func (o *VolumeReplicationSetReqVolume) GetReplicationPoolIdOk() (*int64, bool)`

GetReplicationPoolIdOk returns a tuple with the ReplicationPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPoolId

`func (o *VolumeReplicationSetReqVolume) SetReplicationPoolId(v int64)`

SetReplicationPoolId sets ReplicationPoolId field to given value.

### HasReplicationPoolId

`func (o *VolumeReplicationSetReqVolume) HasReplicationPoolId() bool`

HasReplicationPoolId returns a boolean if a field has been set.

### GetReplicationPoolName

`func (o *VolumeReplicationSetReqVolume) GetReplicationPoolName() string`

GetReplicationPoolName returns the ReplicationPoolName field if non-nil, zero value otherwise.

### GetReplicationPoolNameOk

`func (o *VolumeReplicationSetReqVolume) GetReplicationPoolNameOk() (*string, bool)`

GetReplicationPoolNameOk returns a tuple with the ReplicationPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPoolName

`func (o *VolumeReplicationSetReqVolume) SetReplicationPoolName(v string)`

SetReplicationPoolName sets ReplicationPoolName field to given value.

### HasReplicationPoolName

`func (o *VolumeReplicationSetReqVolume) HasReplicationPoolName() bool`

HasReplicationPoolName returns a boolean if a field has been set.

### GetReplicationVersion

`func (o *VolumeReplicationSetReqVolume) GetReplicationVersion() int64`

GetReplicationVersion returns the ReplicationVersion field if non-nil, zero value otherwise.

### GetReplicationVersionOk

`func (o *VolumeReplicationSetReqVolume) GetReplicationVersionOk() (*int64, bool)`

GetReplicationVersionOk returns a tuple with the ReplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVersion

`func (o *VolumeReplicationSetReqVolume) SetReplicationVersion(v int64)`

SetReplicationVersion sets ReplicationVersion field to given value.

### HasReplicationVersion

`func (o *VolumeReplicationSetReqVolume) HasReplicationVersion() bool`

HasReplicationVersion returns a boolean if a field has been set.

### GetReplicationVolume

`func (o *VolumeReplicationSetReqVolume) GetReplicationVolume() string`

GetReplicationVolume returns the ReplicationVolume field if non-nil, zero value otherwise.

### GetReplicationVolumeOk

`func (o *VolumeReplicationSetReqVolume) GetReplicationVolumeOk() (*string, bool)`

GetReplicationVolumeOk returns a tuple with the ReplicationVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVolume

`func (o *VolumeReplicationSetReqVolume) SetReplicationVolume(v string)`

SetReplicationVolume sets ReplicationVolume field to given value.

### HasReplicationVolume

`func (o *VolumeReplicationSetReqVolume) HasReplicationVolume() bool`

HasReplicationVolume returns a boolean if a field has been set.

### GetReplicationVolumeId

`func (o *VolumeReplicationSetReqVolume) GetReplicationVolumeId() int64`

GetReplicationVolumeId returns the ReplicationVolumeId field if non-nil, zero value otherwise.

### GetReplicationVolumeIdOk

`func (o *VolumeReplicationSetReqVolume) GetReplicationVolumeIdOk() (*int64, bool)`

GetReplicationVolumeIdOk returns a tuple with the ReplicationVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVolumeId

`func (o *VolumeReplicationSetReqVolume) SetReplicationVolumeId(v int64)`

SetReplicationVolumeId sets ReplicationVolumeId field to given value.

### HasReplicationVolumeId

`func (o *VolumeReplicationSetReqVolume) HasReplicationVolumeId() bool`

HasReplicationVolumeId returns a boolean if a field has been set.

### GetReplicationVolumeName

`func (o *VolumeReplicationSetReqVolume) GetReplicationVolumeName() string`

GetReplicationVolumeName returns the ReplicationVolumeName field if non-nil, zero value otherwise.

### GetReplicationVolumeNameOk

`func (o *VolumeReplicationSetReqVolume) GetReplicationVolumeNameOk() (*string, bool)`

GetReplicationVolumeNameOk returns a tuple with the ReplicationVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVolumeName

`func (o *VolumeReplicationSetReqVolume) SetReplicationVolumeName(v string)`

SetReplicationVolumeName sets ReplicationVolumeName field to given value.

### HasReplicationVolumeName

`func (o *VolumeReplicationSetReqVolume) HasReplicationVolumeName() bool`

HasReplicationVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


