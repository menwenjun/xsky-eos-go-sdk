# DfsSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Creator** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DpDfsSnapshotPolicy** | Pointer to [**DpDfsSnapshotPolicyNestview**](DpDfsSnapshotPolicyNestview.md) |  | [optional] 
**ExpireTime** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Lock** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Path** | Pointer to [**DfsPathNestview**](DfsPathNestview.md) |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**Rootfs** | Pointer to [**DfsRootfsNestview**](DfsRootfsNestview.md) |  | [optional] 
**SnapshotName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsSnapshot

`func NewDfsSnapshot() *DfsSnapshot`

NewDfsSnapshot instantiates a new DfsSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSnapshotWithDefaults

`func NewDfsSnapshotWithDefaults() *DfsSnapshot`

NewDfsSnapshotWithDefaults instantiates a new DfsSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsSnapshot) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsSnapshot) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsSnapshot) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsSnapshot) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *DfsSnapshot) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsSnapshot) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsSnapshot) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsSnapshot) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsSnapshot) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsSnapshot) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsSnapshot) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsSnapshot) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCreator

`func (o *DfsSnapshot) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *DfsSnapshot) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *DfsSnapshot) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *DfsSnapshot) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *DfsSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsSnapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDpDfsSnapshotPolicy

`func (o *DfsSnapshot) GetDpDfsSnapshotPolicy() DpDfsSnapshotPolicyNestview`

GetDpDfsSnapshotPolicy returns the DpDfsSnapshotPolicy field if non-nil, zero value otherwise.

### GetDpDfsSnapshotPolicyOk

`func (o *DfsSnapshot) GetDpDfsSnapshotPolicyOk() (*DpDfsSnapshotPolicyNestview, bool)`

GetDpDfsSnapshotPolicyOk returns a tuple with the DpDfsSnapshotPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpDfsSnapshotPolicy

`func (o *DfsSnapshot) SetDpDfsSnapshotPolicy(v DpDfsSnapshotPolicyNestview)`

SetDpDfsSnapshotPolicy sets DpDfsSnapshotPolicy field to given value.

### HasDpDfsSnapshotPolicy

`func (o *DfsSnapshot) HasDpDfsSnapshotPolicy() bool`

HasDpDfsSnapshotPolicy returns a boolean if a field has been set.

### GetExpireTime

`func (o *DfsSnapshot) GetExpireTime() time.Time`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *DfsSnapshot) GetExpireTimeOk() (*time.Time, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *DfsSnapshot) SetExpireTime(v time.Time)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *DfsSnapshot) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetId

`func (o *DfsSnapshot) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsSnapshot) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsSnapshot) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLock

`func (o *DfsSnapshot) GetLock() bool`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *DfsSnapshot) GetLockOk() (*bool, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *DfsSnapshot) SetLock(v bool)`

SetLock sets Lock field to given value.

### HasLock

`func (o *DfsSnapshot) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetName

`func (o *DfsSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsSnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsSnapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *DfsSnapshot) GetPath() DfsPathNestview`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsSnapshot) GetPathOk() (*DfsPathNestview, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsSnapshot) SetPath(v DfsPathNestview)`

SetPath sets Path field to given value.

### HasPath

`func (o *DfsSnapshot) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProgress

`func (o *DfsSnapshot) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DfsSnapshot) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DfsSnapshot) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *DfsSnapshot) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetRootfs

`func (o *DfsSnapshot) GetRootfs() DfsRootfsNestview`

GetRootfs returns the Rootfs field if non-nil, zero value otherwise.

### GetRootfsOk

`func (o *DfsSnapshot) GetRootfsOk() (*DfsRootfsNestview, bool)`

GetRootfsOk returns a tuple with the Rootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfs

`func (o *DfsSnapshot) SetRootfs(v DfsRootfsNestview)`

SetRootfs sets Rootfs field to given value.

### HasRootfs

`func (o *DfsSnapshot) HasRootfs() bool`

HasRootfs returns a boolean if a field has been set.

### GetSnapshotName

`func (o *DfsSnapshot) GetSnapshotName() string`

GetSnapshotName returns the SnapshotName field if non-nil, zero value otherwise.

### GetSnapshotNameOk

`func (o *DfsSnapshot) GetSnapshotNameOk() (*string, bool)`

GetSnapshotNameOk returns a tuple with the SnapshotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotName

`func (o *DfsSnapshot) SetSnapshotName(v string)`

SetSnapshotName sets SnapshotName field to given value.

### HasSnapshotName

`func (o *DfsSnapshot) HasSnapshotName() bool`

HasSnapshotName returns a boolean if a field has been set.

### GetStatus

`func (o *DfsSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsSnapshot) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsSnapshot) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsSnapshot) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsSnapshot) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


