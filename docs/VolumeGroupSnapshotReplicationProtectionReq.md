# VolumeGroupSnapshotReplicationProtectionReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestVolumeGroupName** | **string** | destination volume name | 
**DpVolumeGroupSnapshotReplicationPolicyId** | **int64** | policy | 
**InitialSyncTime** | Pointer to **time.Time** | initial sync time | [optional] 
**VolumePairs** | [**[]VolumePairInfo**](VolumePairInfo.md) | volume pairs | 

## Methods

### NewVolumeGroupSnapshotReplicationProtectionReq

`func NewVolumeGroupSnapshotReplicationProtectionReq(destVolumeGroupName string, dpVolumeGroupSnapshotReplicationPolicyId int64, volumePairs []VolumePairInfo, ) *VolumeGroupSnapshotReplicationProtectionReq`

NewVolumeGroupSnapshotReplicationProtectionReq instantiates a new VolumeGroupSnapshotReplicationProtectionReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupSnapshotReplicationProtectionReqWithDefaults

`func NewVolumeGroupSnapshotReplicationProtectionReqWithDefaults() *VolumeGroupSnapshotReplicationProtectionReq`

NewVolumeGroupSnapshotReplicationProtectionReqWithDefaults instantiates a new VolumeGroupSnapshotReplicationProtectionReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestVolumeGroupName

`func (o *VolumeGroupSnapshotReplicationProtectionReq) GetDestVolumeGroupName() string`

GetDestVolumeGroupName returns the DestVolumeGroupName field if non-nil, zero value otherwise.

### GetDestVolumeGroupNameOk

`func (o *VolumeGroupSnapshotReplicationProtectionReq) GetDestVolumeGroupNameOk() (*string, bool)`

GetDestVolumeGroupNameOk returns a tuple with the DestVolumeGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestVolumeGroupName

`func (o *VolumeGroupSnapshotReplicationProtectionReq) SetDestVolumeGroupName(v string)`

SetDestVolumeGroupName sets DestVolumeGroupName field to given value.


### GetDpVolumeGroupSnapshotReplicationPolicyId

`func (o *VolumeGroupSnapshotReplicationProtectionReq) GetDpVolumeGroupSnapshotReplicationPolicyId() int64`

GetDpVolumeGroupSnapshotReplicationPolicyId returns the DpVolumeGroupSnapshotReplicationPolicyId field if non-nil, zero value otherwise.

### GetDpVolumeGroupSnapshotReplicationPolicyIdOk

`func (o *VolumeGroupSnapshotReplicationProtectionReq) GetDpVolumeGroupSnapshotReplicationPolicyIdOk() (*int64, bool)`

GetDpVolumeGroupSnapshotReplicationPolicyIdOk returns a tuple with the DpVolumeGroupSnapshotReplicationPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpVolumeGroupSnapshotReplicationPolicyId

`func (o *VolumeGroupSnapshotReplicationProtectionReq) SetDpVolumeGroupSnapshotReplicationPolicyId(v int64)`

SetDpVolumeGroupSnapshotReplicationPolicyId sets DpVolumeGroupSnapshotReplicationPolicyId field to given value.


### GetInitialSyncTime

`func (o *VolumeGroupSnapshotReplicationProtectionReq) GetInitialSyncTime() time.Time`

GetInitialSyncTime returns the InitialSyncTime field if non-nil, zero value otherwise.

### GetInitialSyncTimeOk

`func (o *VolumeGroupSnapshotReplicationProtectionReq) GetInitialSyncTimeOk() (*time.Time, bool)`

GetInitialSyncTimeOk returns a tuple with the InitialSyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSyncTime

`func (o *VolumeGroupSnapshotReplicationProtectionReq) SetInitialSyncTime(v time.Time)`

SetInitialSyncTime sets InitialSyncTime field to given value.

### HasInitialSyncTime

`func (o *VolumeGroupSnapshotReplicationProtectionReq) HasInitialSyncTime() bool`

HasInitialSyncTime returns a boolean if a field has been set.

### GetVolumePairs

`func (o *VolumeGroupSnapshotReplicationProtectionReq) GetVolumePairs() []VolumePairInfo`

GetVolumePairs returns the VolumePairs field if non-nil, zero value otherwise.

### GetVolumePairsOk

`func (o *VolumeGroupSnapshotReplicationProtectionReq) GetVolumePairsOk() (*[]VolumePairInfo, bool)`

GetVolumePairsOk returns a tuple with the VolumePairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePairs

`func (o *VolumeGroupSnapshotReplicationProtectionReq) SetVolumePairs(v []VolumePairInfo)`

SetVolumePairs sets VolumePairs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


