# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**AllocatedSize** | Pointer to **int64** |  | [optional] 
**BlockConsistentSnapshot** | Pointer to [**ConsistentSnapshotNestview**](ConsistentSnapshotNestview.md) |  | [optional] 
**BlockVolumeGroupSnapshot** | Pointer to [**VolumeGroupSnapshotNestview**](VolumeGroupSnapshotNestview.md) |  | [optional] 
**ClonedBlockVolumeNum** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Creator** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Passive** | Pointer to **bool** |  | [optional] 
**Pool** | Pointer to [**PoolNestview**](PoolNestview.md) |  | [optional] 
**Protected** | Pointer to **bool** |  | [optional] 
**RemoteCluster** | Pointer to [**RemoteClusterNestview**](RemoteClusterNestview.md) |  | [optional] 
**Reserved** | Pointer to **bool** | Deprecated, used by volume_backup | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**SnapName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Volume** | Pointer to [**VolumeNestview**](VolumeNestview.md) |  | [optional] 

## Methods

### NewSnapshot

`func NewSnapshot() *Snapshot`

NewSnapshot instantiates a new Snapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotWithDefaults

`func NewSnapshotWithDefaults() *Snapshot`

NewSnapshotWithDefaults instantiates a new Snapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *Snapshot) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *Snapshot) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *Snapshot) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *Snapshot) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAllocatedSize

`func (o *Snapshot) GetAllocatedSize() int64`

GetAllocatedSize returns the AllocatedSize field if non-nil, zero value otherwise.

### GetAllocatedSizeOk

`func (o *Snapshot) GetAllocatedSizeOk() (*int64, bool)`

GetAllocatedSizeOk returns a tuple with the AllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSize

`func (o *Snapshot) SetAllocatedSize(v int64)`

SetAllocatedSize sets AllocatedSize field to given value.

### HasAllocatedSize

`func (o *Snapshot) HasAllocatedSize() bool`

HasAllocatedSize returns a boolean if a field has been set.

### GetBlockConsistentSnapshot

`func (o *Snapshot) GetBlockConsistentSnapshot() ConsistentSnapshotNestview`

GetBlockConsistentSnapshot returns the BlockConsistentSnapshot field if non-nil, zero value otherwise.

### GetBlockConsistentSnapshotOk

`func (o *Snapshot) GetBlockConsistentSnapshotOk() (*ConsistentSnapshotNestview, bool)`

GetBlockConsistentSnapshotOk returns a tuple with the BlockConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockConsistentSnapshot

`func (o *Snapshot) SetBlockConsistentSnapshot(v ConsistentSnapshotNestview)`

SetBlockConsistentSnapshot sets BlockConsistentSnapshot field to given value.

### HasBlockConsistentSnapshot

`func (o *Snapshot) HasBlockConsistentSnapshot() bool`

HasBlockConsistentSnapshot returns a boolean if a field has been set.

### GetBlockVolumeGroupSnapshot

`func (o *Snapshot) GetBlockVolumeGroupSnapshot() VolumeGroupSnapshotNestview`

GetBlockVolumeGroupSnapshot returns the BlockVolumeGroupSnapshot field if non-nil, zero value otherwise.

### GetBlockVolumeGroupSnapshotOk

`func (o *Snapshot) GetBlockVolumeGroupSnapshotOk() (*VolumeGroupSnapshotNestview, bool)`

GetBlockVolumeGroupSnapshotOk returns a tuple with the BlockVolumeGroupSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeGroupSnapshot

`func (o *Snapshot) SetBlockVolumeGroupSnapshot(v VolumeGroupSnapshotNestview)`

SetBlockVolumeGroupSnapshot sets BlockVolumeGroupSnapshot field to given value.

### HasBlockVolumeGroupSnapshot

`func (o *Snapshot) HasBlockVolumeGroupSnapshot() bool`

HasBlockVolumeGroupSnapshot returns a boolean if a field has been set.

### GetClonedBlockVolumeNum

`func (o *Snapshot) GetClonedBlockVolumeNum() int64`

GetClonedBlockVolumeNum returns the ClonedBlockVolumeNum field if non-nil, zero value otherwise.

### GetClonedBlockVolumeNumOk

`func (o *Snapshot) GetClonedBlockVolumeNumOk() (*int64, bool)`

GetClonedBlockVolumeNumOk returns a tuple with the ClonedBlockVolumeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedBlockVolumeNum

`func (o *Snapshot) SetClonedBlockVolumeNum(v int64)`

SetClonedBlockVolumeNum sets ClonedBlockVolumeNum field to given value.

### HasClonedBlockVolumeNum

`func (o *Snapshot) HasClonedBlockVolumeNum() bool`

HasClonedBlockVolumeNum returns a boolean if a field has been set.

### GetCluster

`func (o *Snapshot) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Snapshot) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Snapshot) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Snapshot) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *Snapshot) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Snapshot) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Snapshot) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Snapshot) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCreator

`func (o *Snapshot) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Snapshot) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Snapshot) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Snapshot) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *Snapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Snapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Snapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Snapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHidden

`func (o *Snapshot) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Snapshot) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Snapshot) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Snapshot) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *Snapshot) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Snapshot) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Snapshot) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Snapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Snapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Snapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Snapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Snapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassive

`func (o *Snapshot) GetPassive() bool`

GetPassive returns the Passive field if non-nil, zero value otherwise.

### GetPassiveOk

`func (o *Snapshot) GetPassiveOk() (*bool, bool)`

GetPassiveOk returns a tuple with the Passive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassive

`func (o *Snapshot) SetPassive(v bool)`

SetPassive sets Passive field to given value.

### HasPassive

`func (o *Snapshot) HasPassive() bool`

HasPassive returns a boolean if a field has been set.

### GetPool

`func (o *Snapshot) GetPool() PoolNestview`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *Snapshot) GetPoolOk() (*PoolNestview, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *Snapshot) SetPool(v PoolNestview)`

SetPool sets Pool field to given value.

### HasPool

`func (o *Snapshot) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetProtected

`func (o *Snapshot) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *Snapshot) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *Snapshot) SetProtected(v bool)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *Snapshot) HasProtected() bool`

HasProtected returns a boolean if a field has been set.

### GetRemoteCluster

`func (o *Snapshot) GetRemoteCluster() RemoteClusterNestview`

GetRemoteCluster returns the RemoteCluster field if non-nil, zero value otherwise.

### GetRemoteClusterOk

`func (o *Snapshot) GetRemoteClusterOk() (*RemoteClusterNestview, bool)`

GetRemoteClusterOk returns a tuple with the RemoteCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCluster

`func (o *Snapshot) SetRemoteCluster(v RemoteClusterNestview)`

SetRemoteCluster sets RemoteCluster field to given value.

### HasRemoteCluster

`func (o *Snapshot) HasRemoteCluster() bool`

HasRemoteCluster returns a boolean if a field has been set.

### GetReserved

`func (o *Snapshot) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *Snapshot) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *Snapshot) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *Snapshot) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetSize

`func (o *Snapshot) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Snapshot) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Snapshot) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *Snapshot) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSnapName

`func (o *Snapshot) GetSnapName() string`

GetSnapName returns the SnapName field if non-nil, zero value otherwise.

### GetSnapNameOk

`func (o *Snapshot) GetSnapNameOk() (*string, bool)`

GetSnapNameOk returns a tuple with the SnapName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapName

`func (o *Snapshot) SetSnapName(v string)`

SetSnapName sets SnapName field to given value.

### HasSnapName

`func (o *Snapshot) HasSnapName() bool`

HasSnapName returns a boolean if a field has been set.

### GetStatus

`func (o *Snapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Snapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Snapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Snapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUid

`func (o *Snapshot) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Snapshot) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Snapshot) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Snapshot) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUpdate

`func (o *Snapshot) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Snapshot) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Snapshot) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Snapshot) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVolume

`func (o *Snapshot) GetVolume() VolumeNestview`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Snapshot) GetVolumeOk() (*VolumeNestview, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Snapshot) SetVolume(v VolumeNestview)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *Snapshot) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


