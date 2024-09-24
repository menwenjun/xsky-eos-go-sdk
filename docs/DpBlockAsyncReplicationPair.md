# DpBlockAsyncReplicationPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockVolume** | Pointer to [**VolumeNestview**](VolumeNestview.md) |  | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DpBlockAsyncReplicationPolicy** | Pointer to [**DpBlockAsyncReplicationPolicyNestview**](DpBlockAsyncReplicationPolicyNestview.md) |  | [optional] 
**DpVolumeGroupSnapshotReplicationPair** | Pointer to [**DpVolumeGroupSnapshotReplicationPairNestview**](DpVolumeGroupSnapshotReplicationPairNestview.md) |  | [optional] 
**Failovered** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LastSuccessTime** | Pointer to **time.Time** |  | [optional] 
**LocalRole** | Pointer to **string** |  | [optional] 
**PairPoolId** | Pointer to **int64** |  | [optional] 
**PairPoolName** | Pointer to **string** |  | [optional] 
**PairVolumeId** | Pointer to **int64** |  | [optional] 
**PairVolumeName** | Pointer to **string** |  | [optional] 
**Paused** | Pointer to **bool** |  | [optional] 
**PolicyCron** | Pointer to **string** |  | [optional] 
**PreStatus** | Pointer to **string** |  | [optional] 
**RemoteCluster** | Pointer to [**RemoteClusterNestview**](RemoteClusterNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDpBlockAsyncReplicationPair

`func NewDpBlockAsyncReplicationPair() *DpBlockAsyncReplicationPair`

NewDpBlockAsyncReplicationPair instantiates a new DpBlockAsyncReplicationPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockAsyncReplicationPairWithDefaults

`func NewDpBlockAsyncReplicationPairWithDefaults() *DpBlockAsyncReplicationPair`

NewDpBlockAsyncReplicationPairWithDefaults instantiates a new DpBlockAsyncReplicationPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockVolume

`func (o *DpBlockAsyncReplicationPair) GetBlockVolume() VolumeNestview`

GetBlockVolume returns the BlockVolume field if non-nil, zero value otherwise.

### GetBlockVolumeOk

`func (o *DpBlockAsyncReplicationPair) GetBlockVolumeOk() (*VolumeNestview, bool)`

GetBlockVolumeOk returns a tuple with the BlockVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolume

`func (o *DpBlockAsyncReplicationPair) SetBlockVolume(v VolumeNestview)`

SetBlockVolume sets BlockVolume field to given value.

### HasBlockVolume

`func (o *DpBlockAsyncReplicationPair) HasBlockVolume() bool`

HasBlockVolume returns a boolean if a field has been set.

### GetCluster

`func (o *DpBlockAsyncReplicationPair) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpBlockAsyncReplicationPair) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpBlockAsyncReplicationPair) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpBlockAsyncReplicationPair) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnected

`func (o *DpBlockAsyncReplicationPair) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *DpBlockAsyncReplicationPair) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *DpBlockAsyncReplicationPair) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *DpBlockAsyncReplicationPair) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetCreate

`func (o *DpBlockAsyncReplicationPair) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DpBlockAsyncReplicationPair) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DpBlockAsyncReplicationPair) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DpBlockAsyncReplicationPair) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDpBlockAsyncReplicationPolicy

`func (o *DpBlockAsyncReplicationPair) GetDpBlockAsyncReplicationPolicy() DpBlockAsyncReplicationPolicyNestview`

GetDpBlockAsyncReplicationPolicy returns the DpBlockAsyncReplicationPolicy field if non-nil, zero value otherwise.

### GetDpBlockAsyncReplicationPolicyOk

`func (o *DpBlockAsyncReplicationPair) GetDpBlockAsyncReplicationPolicyOk() (*DpBlockAsyncReplicationPolicyNestview, bool)`

GetDpBlockAsyncReplicationPolicyOk returns a tuple with the DpBlockAsyncReplicationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpBlockAsyncReplicationPolicy

`func (o *DpBlockAsyncReplicationPair) SetDpBlockAsyncReplicationPolicy(v DpBlockAsyncReplicationPolicyNestview)`

SetDpBlockAsyncReplicationPolicy sets DpBlockAsyncReplicationPolicy field to given value.

### HasDpBlockAsyncReplicationPolicy

`func (o *DpBlockAsyncReplicationPair) HasDpBlockAsyncReplicationPolicy() bool`

HasDpBlockAsyncReplicationPolicy returns a boolean if a field has been set.

### GetDpVolumeGroupSnapshotReplicationPair

`func (o *DpBlockAsyncReplicationPair) GetDpVolumeGroupSnapshotReplicationPair() DpVolumeGroupSnapshotReplicationPairNestview`

GetDpVolumeGroupSnapshotReplicationPair returns the DpVolumeGroupSnapshotReplicationPair field if non-nil, zero value otherwise.

### GetDpVolumeGroupSnapshotReplicationPairOk

`func (o *DpBlockAsyncReplicationPair) GetDpVolumeGroupSnapshotReplicationPairOk() (*DpVolumeGroupSnapshotReplicationPairNestview, bool)`

GetDpVolumeGroupSnapshotReplicationPairOk returns a tuple with the DpVolumeGroupSnapshotReplicationPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpVolumeGroupSnapshotReplicationPair

`func (o *DpBlockAsyncReplicationPair) SetDpVolumeGroupSnapshotReplicationPair(v DpVolumeGroupSnapshotReplicationPairNestview)`

SetDpVolumeGroupSnapshotReplicationPair sets DpVolumeGroupSnapshotReplicationPair field to given value.

### HasDpVolumeGroupSnapshotReplicationPair

`func (o *DpBlockAsyncReplicationPair) HasDpVolumeGroupSnapshotReplicationPair() bool`

HasDpVolumeGroupSnapshotReplicationPair returns a boolean if a field has been set.

### GetFailovered

`func (o *DpBlockAsyncReplicationPair) GetFailovered() bool`

GetFailovered returns the Failovered field if non-nil, zero value otherwise.

### GetFailoveredOk

`func (o *DpBlockAsyncReplicationPair) GetFailoveredOk() (*bool, bool)`

GetFailoveredOk returns a tuple with the Failovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailovered

`func (o *DpBlockAsyncReplicationPair) SetFailovered(v bool)`

SetFailovered sets Failovered field to given value.

### HasFailovered

`func (o *DpBlockAsyncReplicationPair) HasFailovered() bool`

HasFailovered returns a boolean if a field has been set.

### GetId

`func (o *DpBlockAsyncReplicationPair) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpBlockAsyncReplicationPair) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpBlockAsyncReplicationPair) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpBlockAsyncReplicationPair) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastSuccessTime

`func (o *DpBlockAsyncReplicationPair) GetLastSuccessTime() time.Time`

GetLastSuccessTime returns the LastSuccessTime field if non-nil, zero value otherwise.

### GetLastSuccessTimeOk

`func (o *DpBlockAsyncReplicationPair) GetLastSuccessTimeOk() (*time.Time, bool)`

GetLastSuccessTimeOk returns a tuple with the LastSuccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessTime

`func (o *DpBlockAsyncReplicationPair) SetLastSuccessTime(v time.Time)`

SetLastSuccessTime sets LastSuccessTime field to given value.

### HasLastSuccessTime

`func (o *DpBlockAsyncReplicationPair) HasLastSuccessTime() bool`

HasLastSuccessTime returns a boolean if a field has been set.

### GetLocalRole

`func (o *DpBlockAsyncReplicationPair) GetLocalRole() string`

GetLocalRole returns the LocalRole field if non-nil, zero value otherwise.

### GetLocalRoleOk

`func (o *DpBlockAsyncReplicationPair) GetLocalRoleOk() (*string, bool)`

GetLocalRoleOk returns a tuple with the LocalRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalRole

`func (o *DpBlockAsyncReplicationPair) SetLocalRole(v string)`

SetLocalRole sets LocalRole field to given value.

### HasLocalRole

`func (o *DpBlockAsyncReplicationPair) HasLocalRole() bool`

HasLocalRole returns a boolean if a field has been set.

### GetPairPoolId

`func (o *DpBlockAsyncReplicationPair) GetPairPoolId() int64`

GetPairPoolId returns the PairPoolId field if non-nil, zero value otherwise.

### GetPairPoolIdOk

`func (o *DpBlockAsyncReplicationPair) GetPairPoolIdOk() (*int64, bool)`

GetPairPoolIdOk returns a tuple with the PairPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairPoolId

`func (o *DpBlockAsyncReplicationPair) SetPairPoolId(v int64)`

SetPairPoolId sets PairPoolId field to given value.

### HasPairPoolId

`func (o *DpBlockAsyncReplicationPair) HasPairPoolId() bool`

HasPairPoolId returns a boolean if a field has been set.

### GetPairPoolName

`func (o *DpBlockAsyncReplicationPair) GetPairPoolName() string`

GetPairPoolName returns the PairPoolName field if non-nil, zero value otherwise.

### GetPairPoolNameOk

`func (o *DpBlockAsyncReplicationPair) GetPairPoolNameOk() (*string, bool)`

GetPairPoolNameOk returns a tuple with the PairPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairPoolName

`func (o *DpBlockAsyncReplicationPair) SetPairPoolName(v string)`

SetPairPoolName sets PairPoolName field to given value.

### HasPairPoolName

`func (o *DpBlockAsyncReplicationPair) HasPairPoolName() bool`

HasPairPoolName returns a boolean if a field has been set.

### GetPairVolumeId

`func (o *DpBlockAsyncReplicationPair) GetPairVolumeId() int64`

GetPairVolumeId returns the PairVolumeId field if non-nil, zero value otherwise.

### GetPairVolumeIdOk

`func (o *DpBlockAsyncReplicationPair) GetPairVolumeIdOk() (*int64, bool)`

GetPairVolumeIdOk returns a tuple with the PairVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairVolumeId

`func (o *DpBlockAsyncReplicationPair) SetPairVolumeId(v int64)`

SetPairVolumeId sets PairVolumeId field to given value.

### HasPairVolumeId

`func (o *DpBlockAsyncReplicationPair) HasPairVolumeId() bool`

HasPairVolumeId returns a boolean if a field has been set.

### GetPairVolumeName

`func (o *DpBlockAsyncReplicationPair) GetPairVolumeName() string`

GetPairVolumeName returns the PairVolumeName field if non-nil, zero value otherwise.

### GetPairVolumeNameOk

`func (o *DpBlockAsyncReplicationPair) GetPairVolumeNameOk() (*string, bool)`

GetPairVolumeNameOk returns a tuple with the PairVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairVolumeName

`func (o *DpBlockAsyncReplicationPair) SetPairVolumeName(v string)`

SetPairVolumeName sets PairVolumeName field to given value.

### HasPairVolumeName

`func (o *DpBlockAsyncReplicationPair) HasPairVolumeName() bool`

HasPairVolumeName returns a boolean if a field has been set.

### GetPaused

`func (o *DpBlockAsyncReplicationPair) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *DpBlockAsyncReplicationPair) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *DpBlockAsyncReplicationPair) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *DpBlockAsyncReplicationPair) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetPolicyCron

`func (o *DpBlockAsyncReplicationPair) GetPolicyCron() string`

GetPolicyCron returns the PolicyCron field if non-nil, zero value otherwise.

### GetPolicyCronOk

`func (o *DpBlockAsyncReplicationPair) GetPolicyCronOk() (*string, bool)`

GetPolicyCronOk returns a tuple with the PolicyCron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCron

`func (o *DpBlockAsyncReplicationPair) SetPolicyCron(v string)`

SetPolicyCron sets PolicyCron field to given value.

### HasPolicyCron

`func (o *DpBlockAsyncReplicationPair) HasPolicyCron() bool`

HasPolicyCron returns a boolean if a field has been set.

### GetPreStatus

`func (o *DpBlockAsyncReplicationPair) GetPreStatus() string`

GetPreStatus returns the PreStatus field if non-nil, zero value otherwise.

### GetPreStatusOk

`func (o *DpBlockAsyncReplicationPair) GetPreStatusOk() (*string, bool)`

GetPreStatusOk returns a tuple with the PreStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreStatus

`func (o *DpBlockAsyncReplicationPair) SetPreStatus(v string)`

SetPreStatus sets PreStatus field to given value.

### HasPreStatus

`func (o *DpBlockAsyncReplicationPair) HasPreStatus() bool`

HasPreStatus returns a boolean if a field has been set.

### GetRemoteCluster

`func (o *DpBlockAsyncReplicationPair) GetRemoteCluster() RemoteClusterNestview`

GetRemoteCluster returns the RemoteCluster field if non-nil, zero value otherwise.

### GetRemoteClusterOk

`func (o *DpBlockAsyncReplicationPair) GetRemoteClusterOk() (*RemoteClusterNestview, bool)`

GetRemoteClusterOk returns a tuple with the RemoteCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCluster

`func (o *DpBlockAsyncReplicationPair) SetRemoteCluster(v RemoteClusterNestview)`

SetRemoteCluster sets RemoteCluster field to given value.

### HasRemoteCluster

`func (o *DpBlockAsyncReplicationPair) HasRemoteCluster() bool`

HasRemoteCluster returns a boolean if a field has been set.

### GetStatus

`func (o *DpBlockAsyncReplicationPair) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpBlockAsyncReplicationPair) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpBlockAsyncReplicationPair) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpBlockAsyncReplicationPair) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DpBlockAsyncReplicationPair) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DpBlockAsyncReplicationPair) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DpBlockAsyncReplicationPair) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DpBlockAsyncReplicationPair) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


