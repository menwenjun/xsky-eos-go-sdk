# DpVolumeGroupSnapshotReplicationPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DpVolumeGroupSnapshotReplicationPolicy** | Pointer to [**DpVolumeGroupSnapshotReplicationPolicyNestview**](DpVolumeGroupSnapshotReplicationPolicyNestview.md) |  | [optional] 
**Failovered** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LastSuccessTime** | Pointer to **time.Time** |  | [optional] 
**LocalRole** | Pointer to **string** |  | [optional] 
**PairVolumeGroupId** | Pointer to **int64** |  | [optional] 
**PairVolumeGroupName** | Pointer to **string** |  | [optional] 
**Paused** | Pointer to **bool** |  | [optional] 
**PolicyCron** | Pointer to **string** |  | [optional] 
**PreStatus** | Pointer to **string** |  | [optional] 
**RemoteCluster** | Pointer to [**RemoteClusterNestview**](RemoteClusterNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**VolumeGroup** | Pointer to [**VolumeGroupNestview**](VolumeGroupNestview.md) |  | [optional] 

## Methods

### NewDpVolumeGroupSnapshotReplicationPair

`func NewDpVolumeGroupSnapshotReplicationPair() *DpVolumeGroupSnapshotReplicationPair`

NewDpVolumeGroupSnapshotReplicationPair instantiates a new DpVolumeGroupSnapshotReplicationPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpVolumeGroupSnapshotReplicationPairWithDefaults

`func NewDpVolumeGroupSnapshotReplicationPairWithDefaults() *DpVolumeGroupSnapshotReplicationPair`

NewDpVolumeGroupSnapshotReplicationPairWithDefaults instantiates a new DpVolumeGroupSnapshotReplicationPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DpVolumeGroupSnapshotReplicationPair) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpVolumeGroupSnapshotReplicationPair) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpVolumeGroupSnapshotReplicationPair) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnected

`func (o *DpVolumeGroupSnapshotReplicationPair) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *DpVolumeGroupSnapshotReplicationPair) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *DpVolumeGroupSnapshotReplicationPair) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetCreate

`func (o *DpVolumeGroupSnapshotReplicationPair) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DpVolumeGroupSnapshotReplicationPair) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DpVolumeGroupSnapshotReplicationPair) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDpVolumeGroupSnapshotReplicationPolicy

`func (o *DpVolumeGroupSnapshotReplicationPair) GetDpVolumeGroupSnapshotReplicationPolicy() DpVolumeGroupSnapshotReplicationPolicyNestview`

GetDpVolumeGroupSnapshotReplicationPolicy returns the DpVolumeGroupSnapshotReplicationPolicy field if non-nil, zero value otherwise.

### GetDpVolumeGroupSnapshotReplicationPolicyOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetDpVolumeGroupSnapshotReplicationPolicyOk() (*DpVolumeGroupSnapshotReplicationPolicyNestview, bool)`

GetDpVolumeGroupSnapshotReplicationPolicyOk returns a tuple with the DpVolumeGroupSnapshotReplicationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpVolumeGroupSnapshotReplicationPolicy

`func (o *DpVolumeGroupSnapshotReplicationPair) SetDpVolumeGroupSnapshotReplicationPolicy(v DpVolumeGroupSnapshotReplicationPolicyNestview)`

SetDpVolumeGroupSnapshotReplicationPolicy sets DpVolumeGroupSnapshotReplicationPolicy field to given value.

### HasDpVolumeGroupSnapshotReplicationPolicy

`func (o *DpVolumeGroupSnapshotReplicationPair) HasDpVolumeGroupSnapshotReplicationPolicy() bool`

HasDpVolumeGroupSnapshotReplicationPolicy returns a boolean if a field has been set.

### GetFailovered

`func (o *DpVolumeGroupSnapshotReplicationPair) GetFailovered() bool`

GetFailovered returns the Failovered field if non-nil, zero value otherwise.

### GetFailoveredOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetFailoveredOk() (*bool, bool)`

GetFailoveredOk returns a tuple with the Failovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailovered

`func (o *DpVolumeGroupSnapshotReplicationPair) SetFailovered(v bool)`

SetFailovered sets Failovered field to given value.

### HasFailovered

`func (o *DpVolumeGroupSnapshotReplicationPair) HasFailovered() bool`

HasFailovered returns a boolean if a field has been set.

### GetId

`func (o *DpVolumeGroupSnapshotReplicationPair) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpVolumeGroupSnapshotReplicationPair) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpVolumeGroupSnapshotReplicationPair) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastSuccessTime

`func (o *DpVolumeGroupSnapshotReplicationPair) GetLastSuccessTime() time.Time`

GetLastSuccessTime returns the LastSuccessTime field if non-nil, zero value otherwise.

### GetLastSuccessTimeOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetLastSuccessTimeOk() (*time.Time, bool)`

GetLastSuccessTimeOk returns a tuple with the LastSuccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessTime

`func (o *DpVolumeGroupSnapshotReplicationPair) SetLastSuccessTime(v time.Time)`

SetLastSuccessTime sets LastSuccessTime field to given value.

### HasLastSuccessTime

`func (o *DpVolumeGroupSnapshotReplicationPair) HasLastSuccessTime() bool`

HasLastSuccessTime returns a boolean if a field has been set.

### GetLocalRole

`func (o *DpVolumeGroupSnapshotReplicationPair) GetLocalRole() string`

GetLocalRole returns the LocalRole field if non-nil, zero value otherwise.

### GetLocalRoleOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetLocalRoleOk() (*string, bool)`

GetLocalRoleOk returns a tuple with the LocalRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalRole

`func (o *DpVolumeGroupSnapshotReplicationPair) SetLocalRole(v string)`

SetLocalRole sets LocalRole field to given value.

### HasLocalRole

`func (o *DpVolumeGroupSnapshotReplicationPair) HasLocalRole() bool`

HasLocalRole returns a boolean if a field has been set.

### GetPairVolumeGroupId

`func (o *DpVolumeGroupSnapshotReplicationPair) GetPairVolumeGroupId() int64`

GetPairVolumeGroupId returns the PairVolumeGroupId field if non-nil, zero value otherwise.

### GetPairVolumeGroupIdOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetPairVolumeGroupIdOk() (*int64, bool)`

GetPairVolumeGroupIdOk returns a tuple with the PairVolumeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairVolumeGroupId

`func (o *DpVolumeGroupSnapshotReplicationPair) SetPairVolumeGroupId(v int64)`

SetPairVolumeGroupId sets PairVolumeGroupId field to given value.

### HasPairVolumeGroupId

`func (o *DpVolumeGroupSnapshotReplicationPair) HasPairVolumeGroupId() bool`

HasPairVolumeGroupId returns a boolean if a field has been set.

### GetPairVolumeGroupName

`func (o *DpVolumeGroupSnapshotReplicationPair) GetPairVolumeGroupName() string`

GetPairVolumeGroupName returns the PairVolumeGroupName field if non-nil, zero value otherwise.

### GetPairVolumeGroupNameOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetPairVolumeGroupNameOk() (*string, bool)`

GetPairVolumeGroupNameOk returns a tuple with the PairVolumeGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairVolumeGroupName

`func (o *DpVolumeGroupSnapshotReplicationPair) SetPairVolumeGroupName(v string)`

SetPairVolumeGroupName sets PairVolumeGroupName field to given value.

### HasPairVolumeGroupName

`func (o *DpVolumeGroupSnapshotReplicationPair) HasPairVolumeGroupName() bool`

HasPairVolumeGroupName returns a boolean if a field has been set.

### GetPaused

`func (o *DpVolumeGroupSnapshotReplicationPair) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *DpVolumeGroupSnapshotReplicationPair) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *DpVolumeGroupSnapshotReplicationPair) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetPolicyCron

`func (o *DpVolumeGroupSnapshotReplicationPair) GetPolicyCron() string`

GetPolicyCron returns the PolicyCron field if non-nil, zero value otherwise.

### GetPolicyCronOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetPolicyCronOk() (*string, bool)`

GetPolicyCronOk returns a tuple with the PolicyCron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCron

`func (o *DpVolumeGroupSnapshotReplicationPair) SetPolicyCron(v string)`

SetPolicyCron sets PolicyCron field to given value.

### HasPolicyCron

`func (o *DpVolumeGroupSnapshotReplicationPair) HasPolicyCron() bool`

HasPolicyCron returns a boolean if a field has been set.

### GetPreStatus

`func (o *DpVolumeGroupSnapshotReplicationPair) GetPreStatus() string`

GetPreStatus returns the PreStatus field if non-nil, zero value otherwise.

### GetPreStatusOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetPreStatusOk() (*string, bool)`

GetPreStatusOk returns a tuple with the PreStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreStatus

`func (o *DpVolumeGroupSnapshotReplicationPair) SetPreStatus(v string)`

SetPreStatus sets PreStatus field to given value.

### HasPreStatus

`func (o *DpVolumeGroupSnapshotReplicationPair) HasPreStatus() bool`

HasPreStatus returns a boolean if a field has been set.

### GetRemoteCluster

`func (o *DpVolumeGroupSnapshotReplicationPair) GetRemoteCluster() RemoteClusterNestview`

GetRemoteCluster returns the RemoteCluster field if non-nil, zero value otherwise.

### GetRemoteClusterOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetRemoteClusterOk() (*RemoteClusterNestview, bool)`

GetRemoteClusterOk returns a tuple with the RemoteCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCluster

`func (o *DpVolumeGroupSnapshotReplicationPair) SetRemoteCluster(v RemoteClusterNestview)`

SetRemoteCluster sets RemoteCluster field to given value.

### HasRemoteCluster

`func (o *DpVolumeGroupSnapshotReplicationPair) HasRemoteCluster() bool`

HasRemoteCluster returns a boolean if a field has been set.

### GetStatus

`func (o *DpVolumeGroupSnapshotReplicationPair) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpVolumeGroupSnapshotReplicationPair) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpVolumeGroupSnapshotReplicationPair) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DpVolumeGroupSnapshotReplicationPair) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DpVolumeGroupSnapshotReplicationPair) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DpVolumeGroupSnapshotReplicationPair) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVolumeGroup

`func (o *DpVolumeGroupSnapshotReplicationPair) GetVolumeGroup() VolumeGroupNestview`

GetVolumeGroup returns the VolumeGroup field if non-nil, zero value otherwise.

### GetVolumeGroupOk

`func (o *DpVolumeGroupSnapshotReplicationPair) GetVolumeGroupOk() (*VolumeGroupNestview, bool)`

GetVolumeGroupOk returns a tuple with the VolumeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroup

`func (o *DpVolumeGroupSnapshotReplicationPair) SetVolumeGroup(v VolumeGroupNestview)`

SetVolumeGroup sets VolumeGroup field to given value.

### HasVolumeGroup

`func (o *DpVolumeGroupSnapshotReplicationPair) HasVolumeGroup() bool`

HasVolumeGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


