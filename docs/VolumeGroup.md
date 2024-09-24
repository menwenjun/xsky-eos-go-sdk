# VolumeGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**BlockVolumeGroupSnapshot** | Pointer to [**VolumeGroupSnapshotNestview**](VolumeGroupSnapshotNestview.md) |  | [optional] 
**BlockVolumeGroupSnapshotNum** | Pointer to **int64** |  | [optional] 
**BlockVolumeNum** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Flattened** | Pointer to **bool** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LatestSnapshotTime** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Passive** | Pointer to **bool** |  | [optional] 
**Pool** | Pointer to [**PoolNestview**](PoolNestview.md) |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**SnapshotReplicationRole** | Pointer to **string** |  | [optional] 
**SnapshotReplicationStatus** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewVolumeGroup

`func NewVolumeGroup() *VolumeGroup`

NewVolumeGroup instantiates a new VolumeGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupWithDefaults

`func NewVolumeGroupWithDefaults() *VolumeGroup`

NewVolumeGroupWithDefaults instantiates a new VolumeGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *VolumeGroup) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *VolumeGroup) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *VolumeGroup) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *VolumeGroup) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetBlockVolumeGroupSnapshot

`func (o *VolumeGroup) GetBlockVolumeGroupSnapshot() VolumeGroupSnapshotNestview`

GetBlockVolumeGroupSnapshot returns the BlockVolumeGroupSnapshot field if non-nil, zero value otherwise.

### GetBlockVolumeGroupSnapshotOk

`func (o *VolumeGroup) GetBlockVolumeGroupSnapshotOk() (*VolumeGroupSnapshotNestview, bool)`

GetBlockVolumeGroupSnapshotOk returns a tuple with the BlockVolumeGroupSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeGroupSnapshot

`func (o *VolumeGroup) SetBlockVolumeGroupSnapshot(v VolumeGroupSnapshotNestview)`

SetBlockVolumeGroupSnapshot sets BlockVolumeGroupSnapshot field to given value.

### HasBlockVolumeGroupSnapshot

`func (o *VolumeGroup) HasBlockVolumeGroupSnapshot() bool`

HasBlockVolumeGroupSnapshot returns a boolean if a field has been set.

### GetBlockVolumeGroupSnapshotNum

`func (o *VolumeGroup) GetBlockVolumeGroupSnapshotNum() int64`

GetBlockVolumeGroupSnapshotNum returns the BlockVolumeGroupSnapshotNum field if non-nil, zero value otherwise.

### GetBlockVolumeGroupSnapshotNumOk

`func (o *VolumeGroup) GetBlockVolumeGroupSnapshotNumOk() (*int64, bool)`

GetBlockVolumeGroupSnapshotNumOk returns a tuple with the BlockVolumeGroupSnapshotNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeGroupSnapshotNum

`func (o *VolumeGroup) SetBlockVolumeGroupSnapshotNum(v int64)`

SetBlockVolumeGroupSnapshotNum sets BlockVolumeGroupSnapshotNum field to given value.

### HasBlockVolumeGroupSnapshotNum

`func (o *VolumeGroup) HasBlockVolumeGroupSnapshotNum() bool`

HasBlockVolumeGroupSnapshotNum returns a boolean if a field has been set.

### GetBlockVolumeNum

`func (o *VolumeGroup) GetBlockVolumeNum() int64`

GetBlockVolumeNum returns the BlockVolumeNum field if non-nil, zero value otherwise.

### GetBlockVolumeNumOk

`func (o *VolumeGroup) GetBlockVolumeNumOk() (*int64, bool)`

GetBlockVolumeNumOk returns a tuple with the BlockVolumeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeNum

`func (o *VolumeGroup) SetBlockVolumeNum(v int64)`

SetBlockVolumeNum sets BlockVolumeNum field to given value.

### HasBlockVolumeNum

`func (o *VolumeGroup) HasBlockVolumeNum() bool`

HasBlockVolumeNum returns a boolean if a field has been set.

### GetCluster

`func (o *VolumeGroup) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VolumeGroup) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VolumeGroup) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VolumeGroup) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *VolumeGroup) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *VolumeGroup) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *VolumeGroup) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *VolumeGroup) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *VolumeGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VolumeGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VolumeGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VolumeGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFlattened

`func (o *VolumeGroup) GetFlattened() bool`

GetFlattened returns the Flattened field if non-nil, zero value otherwise.

### GetFlattenedOk

`func (o *VolumeGroup) GetFlattenedOk() (*bool, bool)`

GetFlattenedOk returns a tuple with the Flattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattened

`func (o *VolumeGroup) SetFlattened(v bool)`

SetFlattened sets Flattened field to given value.

### HasFlattened

`func (o *VolumeGroup) HasFlattened() bool`

HasFlattened returns a boolean if a field has been set.

### GetGroupName

`func (o *VolumeGroup) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *VolumeGroup) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *VolumeGroup) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *VolumeGroup) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetId

`func (o *VolumeGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatestSnapshotTime

`func (o *VolumeGroup) GetLatestSnapshotTime() time.Time`

GetLatestSnapshotTime returns the LatestSnapshotTime field if non-nil, zero value otherwise.

### GetLatestSnapshotTimeOk

`func (o *VolumeGroup) GetLatestSnapshotTimeOk() (*time.Time, bool)`

GetLatestSnapshotTimeOk returns a tuple with the LatestSnapshotTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSnapshotTime

`func (o *VolumeGroup) SetLatestSnapshotTime(v time.Time)`

SetLatestSnapshotTime sets LatestSnapshotTime field to given value.

### HasLatestSnapshotTime

`func (o *VolumeGroup) HasLatestSnapshotTime() bool`

HasLatestSnapshotTime returns a boolean if a field has been set.

### GetName

`func (o *VolumeGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassive

`func (o *VolumeGroup) GetPassive() bool`

GetPassive returns the Passive field if non-nil, zero value otherwise.

### GetPassiveOk

`func (o *VolumeGroup) GetPassiveOk() (*bool, bool)`

GetPassiveOk returns a tuple with the Passive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassive

`func (o *VolumeGroup) SetPassive(v bool)`

SetPassive sets Passive field to given value.

### HasPassive

`func (o *VolumeGroup) HasPassive() bool`

HasPassive returns a boolean if a field has been set.

### GetPool

`func (o *VolumeGroup) GetPool() PoolNestview`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *VolumeGroup) GetPoolOk() (*PoolNestview, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *VolumeGroup) SetPool(v PoolNestview)`

SetPool sets Pool field to given value.

### HasPool

`func (o *VolumeGroup) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetProgress

`func (o *VolumeGroup) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *VolumeGroup) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *VolumeGroup) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *VolumeGroup) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetSnapshotReplicationRole

`func (o *VolumeGroup) GetSnapshotReplicationRole() string`

GetSnapshotReplicationRole returns the SnapshotReplicationRole field if non-nil, zero value otherwise.

### GetSnapshotReplicationRoleOk

`func (o *VolumeGroup) GetSnapshotReplicationRoleOk() (*string, bool)`

GetSnapshotReplicationRoleOk returns a tuple with the SnapshotReplicationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotReplicationRole

`func (o *VolumeGroup) SetSnapshotReplicationRole(v string)`

SetSnapshotReplicationRole sets SnapshotReplicationRole field to given value.

### HasSnapshotReplicationRole

`func (o *VolumeGroup) HasSnapshotReplicationRole() bool`

HasSnapshotReplicationRole returns a boolean if a field has been set.

### GetSnapshotReplicationStatus

`func (o *VolumeGroup) GetSnapshotReplicationStatus() string`

GetSnapshotReplicationStatus returns the SnapshotReplicationStatus field if non-nil, zero value otherwise.

### GetSnapshotReplicationStatusOk

`func (o *VolumeGroup) GetSnapshotReplicationStatusOk() (*string, bool)`

GetSnapshotReplicationStatusOk returns a tuple with the SnapshotReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotReplicationStatus

`func (o *VolumeGroup) SetSnapshotReplicationStatus(v string)`

SetSnapshotReplicationStatus sets SnapshotReplicationStatus field to given value.

### HasSnapshotReplicationStatus

`func (o *VolumeGroup) HasSnapshotReplicationStatus() bool`

HasSnapshotReplicationStatus returns a boolean if a field has been set.

### GetStatus

`func (o *VolumeGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumeGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumeGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumeGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *VolumeGroup) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *VolumeGroup) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *VolumeGroup) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *VolumeGroup) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


