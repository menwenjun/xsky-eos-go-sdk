# VolumeAsyncReplicationProtectionReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestPoolId** | **int64** | destination pool id | 
**DestPoolName** | **string** | destination pool name | 
**DestVolumeName** | **string** | destination volume name | 
**DpBlockAsyncReplicationPolicyId** | **int64** | policy | 
**InitialSyncTime** | Pointer to **time.Time** | initial sync time | [optional] 

## Methods

### NewVolumeAsyncReplicationProtectionReq

`func NewVolumeAsyncReplicationProtectionReq(destPoolId int64, destPoolName string, destVolumeName string, dpBlockAsyncReplicationPolicyId int64, ) *VolumeAsyncReplicationProtectionReq`

NewVolumeAsyncReplicationProtectionReq instantiates a new VolumeAsyncReplicationProtectionReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeAsyncReplicationProtectionReqWithDefaults

`func NewVolumeAsyncReplicationProtectionReqWithDefaults() *VolumeAsyncReplicationProtectionReq`

NewVolumeAsyncReplicationProtectionReqWithDefaults instantiates a new VolumeAsyncReplicationProtectionReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestPoolId

`func (o *VolumeAsyncReplicationProtectionReq) GetDestPoolId() int64`

GetDestPoolId returns the DestPoolId field if non-nil, zero value otherwise.

### GetDestPoolIdOk

`func (o *VolumeAsyncReplicationProtectionReq) GetDestPoolIdOk() (*int64, bool)`

GetDestPoolIdOk returns a tuple with the DestPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPoolId

`func (o *VolumeAsyncReplicationProtectionReq) SetDestPoolId(v int64)`

SetDestPoolId sets DestPoolId field to given value.


### GetDestPoolName

`func (o *VolumeAsyncReplicationProtectionReq) GetDestPoolName() string`

GetDestPoolName returns the DestPoolName field if non-nil, zero value otherwise.

### GetDestPoolNameOk

`func (o *VolumeAsyncReplicationProtectionReq) GetDestPoolNameOk() (*string, bool)`

GetDestPoolNameOk returns a tuple with the DestPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPoolName

`func (o *VolumeAsyncReplicationProtectionReq) SetDestPoolName(v string)`

SetDestPoolName sets DestPoolName field to given value.


### GetDestVolumeName

`func (o *VolumeAsyncReplicationProtectionReq) GetDestVolumeName() string`

GetDestVolumeName returns the DestVolumeName field if non-nil, zero value otherwise.

### GetDestVolumeNameOk

`func (o *VolumeAsyncReplicationProtectionReq) GetDestVolumeNameOk() (*string, bool)`

GetDestVolumeNameOk returns a tuple with the DestVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestVolumeName

`func (o *VolumeAsyncReplicationProtectionReq) SetDestVolumeName(v string)`

SetDestVolumeName sets DestVolumeName field to given value.


### GetDpBlockAsyncReplicationPolicyId

`func (o *VolumeAsyncReplicationProtectionReq) GetDpBlockAsyncReplicationPolicyId() int64`

GetDpBlockAsyncReplicationPolicyId returns the DpBlockAsyncReplicationPolicyId field if non-nil, zero value otherwise.

### GetDpBlockAsyncReplicationPolicyIdOk

`func (o *VolumeAsyncReplicationProtectionReq) GetDpBlockAsyncReplicationPolicyIdOk() (*int64, bool)`

GetDpBlockAsyncReplicationPolicyIdOk returns a tuple with the DpBlockAsyncReplicationPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpBlockAsyncReplicationPolicyId

`func (o *VolumeAsyncReplicationProtectionReq) SetDpBlockAsyncReplicationPolicyId(v int64)`

SetDpBlockAsyncReplicationPolicyId sets DpBlockAsyncReplicationPolicyId field to given value.


### GetInitialSyncTime

`func (o *VolumeAsyncReplicationProtectionReq) GetInitialSyncTime() time.Time`

GetInitialSyncTime returns the InitialSyncTime field if non-nil, zero value otherwise.

### GetInitialSyncTimeOk

`func (o *VolumeAsyncReplicationProtectionReq) GetInitialSyncTimeOk() (*time.Time, bool)`

GetInitialSyncTimeOk returns a tuple with the InitialSyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSyncTime

`func (o *VolumeAsyncReplicationProtectionReq) SetInitialSyncTime(v time.Time)`

SetInitialSyncTime sets InitialSyncTime field to given value.

### HasInitialSyncTime

`func (o *VolumeAsyncReplicationProtectionReq) HasInitialSyncTime() bool`

HasInitialSyncTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


