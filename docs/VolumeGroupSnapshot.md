# VolumeGroupSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**BlockSnapshotNum** | Pointer to **int64** |  | [optional] 
**BlockVolumeGroup** | Pointer to [**VolumeGroupNestview**](VolumeGroupNestview.md) |  | [optional] 
**ClonedBlockVolumeGroupNum** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Creator** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GroupSnapshotName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Passive** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewVolumeGroupSnapshot

`func NewVolumeGroupSnapshot() *VolumeGroupSnapshot`

NewVolumeGroupSnapshot instantiates a new VolumeGroupSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupSnapshotWithDefaults

`func NewVolumeGroupSnapshotWithDefaults() *VolumeGroupSnapshot`

NewVolumeGroupSnapshotWithDefaults instantiates a new VolumeGroupSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *VolumeGroupSnapshot) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *VolumeGroupSnapshot) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *VolumeGroupSnapshot) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *VolumeGroupSnapshot) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetBlockSnapshotNum

`func (o *VolumeGroupSnapshot) GetBlockSnapshotNum() int64`

GetBlockSnapshotNum returns the BlockSnapshotNum field if non-nil, zero value otherwise.

### GetBlockSnapshotNumOk

`func (o *VolumeGroupSnapshot) GetBlockSnapshotNumOk() (*int64, bool)`

GetBlockSnapshotNumOk returns a tuple with the BlockSnapshotNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSnapshotNum

`func (o *VolumeGroupSnapshot) SetBlockSnapshotNum(v int64)`

SetBlockSnapshotNum sets BlockSnapshotNum field to given value.

### HasBlockSnapshotNum

`func (o *VolumeGroupSnapshot) HasBlockSnapshotNum() bool`

HasBlockSnapshotNum returns a boolean if a field has been set.

### GetBlockVolumeGroup

`func (o *VolumeGroupSnapshot) GetBlockVolumeGroup() VolumeGroupNestview`

GetBlockVolumeGroup returns the BlockVolumeGroup field if non-nil, zero value otherwise.

### GetBlockVolumeGroupOk

`func (o *VolumeGroupSnapshot) GetBlockVolumeGroupOk() (*VolumeGroupNestview, bool)`

GetBlockVolumeGroupOk returns a tuple with the BlockVolumeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeGroup

`func (o *VolumeGroupSnapshot) SetBlockVolumeGroup(v VolumeGroupNestview)`

SetBlockVolumeGroup sets BlockVolumeGroup field to given value.

### HasBlockVolumeGroup

`func (o *VolumeGroupSnapshot) HasBlockVolumeGroup() bool`

HasBlockVolumeGroup returns a boolean if a field has been set.

### GetClonedBlockVolumeGroupNum

`func (o *VolumeGroupSnapshot) GetClonedBlockVolumeGroupNum() int64`

GetClonedBlockVolumeGroupNum returns the ClonedBlockVolumeGroupNum field if non-nil, zero value otherwise.

### GetClonedBlockVolumeGroupNumOk

`func (o *VolumeGroupSnapshot) GetClonedBlockVolumeGroupNumOk() (*int64, bool)`

GetClonedBlockVolumeGroupNumOk returns a tuple with the ClonedBlockVolumeGroupNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedBlockVolumeGroupNum

`func (o *VolumeGroupSnapshot) SetClonedBlockVolumeGroupNum(v int64)`

SetClonedBlockVolumeGroupNum sets ClonedBlockVolumeGroupNum field to given value.

### HasClonedBlockVolumeGroupNum

`func (o *VolumeGroupSnapshot) HasClonedBlockVolumeGroupNum() bool`

HasClonedBlockVolumeGroupNum returns a boolean if a field has been set.

### GetCluster

`func (o *VolumeGroupSnapshot) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VolumeGroupSnapshot) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VolumeGroupSnapshot) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VolumeGroupSnapshot) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *VolumeGroupSnapshot) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *VolumeGroupSnapshot) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *VolumeGroupSnapshot) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *VolumeGroupSnapshot) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCreator

`func (o *VolumeGroupSnapshot) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *VolumeGroupSnapshot) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *VolumeGroupSnapshot) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *VolumeGroupSnapshot) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *VolumeGroupSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VolumeGroupSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VolumeGroupSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VolumeGroupSnapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupSnapshotName

`func (o *VolumeGroupSnapshot) GetGroupSnapshotName() string`

GetGroupSnapshotName returns the GroupSnapshotName field if non-nil, zero value otherwise.

### GetGroupSnapshotNameOk

`func (o *VolumeGroupSnapshot) GetGroupSnapshotNameOk() (*string, bool)`

GetGroupSnapshotNameOk returns a tuple with the GroupSnapshotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSnapshotName

`func (o *VolumeGroupSnapshot) SetGroupSnapshotName(v string)`

SetGroupSnapshotName sets GroupSnapshotName field to given value.

### HasGroupSnapshotName

`func (o *VolumeGroupSnapshot) HasGroupSnapshotName() bool`

HasGroupSnapshotName returns a boolean if a field has been set.

### GetId

`func (o *VolumeGroupSnapshot) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeGroupSnapshot) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeGroupSnapshot) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeGroupSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VolumeGroupSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeGroupSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeGroupSnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeGroupSnapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassive

`func (o *VolumeGroupSnapshot) GetPassive() bool`

GetPassive returns the Passive field if non-nil, zero value otherwise.

### GetPassiveOk

`func (o *VolumeGroupSnapshot) GetPassiveOk() (*bool, bool)`

GetPassiveOk returns a tuple with the Passive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassive

`func (o *VolumeGroupSnapshot) SetPassive(v bool)`

SetPassive sets Passive field to given value.

### HasPassive

`func (o *VolumeGroupSnapshot) HasPassive() bool`

HasPassive returns a boolean if a field has been set.

### GetStatus

`func (o *VolumeGroupSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumeGroupSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumeGroupSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumeGroupSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *VolumeGroupSnapshot) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *VolumeGroupSnapshot) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *VolumeGroupSnapshot) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *VolumeGroupSnapshot) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


