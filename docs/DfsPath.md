# DfsPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsPathPerformance** | Pointer to [**DfsPathPerformance**](DfsPathPerformance.md) |  | [optional] 
**DfsPathStat** | Pointer to [**DfsPathStat**](DfsPathStat.md) |  | [optional] 
**DfsStoragePolicyNum** | Pointer to **int64** |  | [optional] 
**DpDfsSnapshotPolicy** | Pointer to [**DpDfsSnapshotPolicy**](DpDfsSnapshotPolicy.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Rootfs** | Pointer to [**DfsRootfsNestview**](DfsRootfsNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StoragePolicyIds** | Pointer to **[]int64** |  | [optional] 
**Stretched** | Pointer to **bool** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsPath

`func NewDfsPath() *DfsPath`

NewDfsPath instantiates a new DfsPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsPathWithDefaults

`func NewDfsPathWithDefaults() *DfsPath`

NewDfsPathWithDefaults instantiates a new DfsPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsPath) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsPath) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsPath) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsPath) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *DfsPath) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsPath) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsPath) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsPath) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsPath) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsPath) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsPath) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsPath) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsPathPerformance

`func (o *DfsPath) GetDfsPathPerformance() DfsPathPerformance`

GetDfsPathPerformance returns the DfsPathPerformance field if non-nil, zero value otherwise.

### GetDfsPathPerformanceOk

`func (o *DfsPath) GetDfsPathPerformanceOk() (*DfsPathPerformance, bool)`

GetDfsPathPerformanceOk returns a tuple with the DfsPathPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPathPerformance

`func (o *DfsPath) SetDfsPathPerformance(v DfsPathPerformance)`

SetDfsPathPerformance sets DfsPathPerformance field to given value.

### HasDfsPathPerformance

`func (o *DfsPath) HasDfsPathPerformance() bool`

HasDfsPathPerformance returns a boolean if a field has been set.

### GetDfsPathStat

`func (o *DfsPath) GetDfsPathStat() DfsPathStat`

GetDfsPathStat returns the DfsPathStat field if non-nil, zero value otherwise.

### GetDfsPathStatOk

`func (o *DfsPath) GetDfsPathStatOk() (*DfsPathStat, bool)`

GetDfsPathStatOk returns a tuple with the DfsPathStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPathStat

`func (o *DfsPath) SetDfsPathStat(v DfsPathStat)`

SetDfsPathStat sets DfsPathStat field to given value.

### HasDfsPathStat

`func (o *DfsPath) HasDfsPathStat() bool`

HasDfsPathStat returns a boolean if a field has been set.

### GetDfsStoragePolicyNum

`func (o *DfsPath) GetDfsStoragePolicyNum() int64`

GetDfsStoragePolicyNum returns the DfsStoragePolicyNum field if non-nil, zero value otherwise.

### GetDfsStoragePolicyNumOk

`func (o *DfsPath) GetDfsStoragePolicyNumOk() (*int64, bool)`

GetDfsStoragePolicyNumOk returns a tuple with the DfsStoragePolicyNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsStoragePolicyNum

`func (o *DfsPath) SetDfsStoragePolicyNum(v int64)`

SetDfsStoragePolicyNum sets DfsStoragePolicyNum field to given value.

### HasDfsStoragePolicyNum

`func (o *DfsPath) HasDfsStoragePolicyNum() bool`

HasDfsStoragePolicyNum returns a boolean if a field has been set.

### GetDpDfsSnapshotPolicy

`func (o *DfsPath) GetDpDfsSnapshotPolicy() DpDfsSnapshotPolicy`

GetDpDfsSnapshotPolicy returns the DpDfsSnapshotPolicy field if non-nil, zero value otherwise.

### GetDpDfsSnapshotPolicyOk

`func (o *DfsPath) GetDpDfsSnapshotPolicyOk() (*DpDfsSnapshotPolicy, bool)`

GetDpDfsSnapshotPolicyOk returns a tuple with the DpDfsSnapshotPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpDfsSnapshotPolicy

`func (o *DfsPath) SetDpDfsSnapshotPolicy(v DpDfsSnapshotPolicy)`

SetDpDfsSnapshotPolicy sets DpDfsSnapshotPolicy field to given value.

### HasDpDfsSnapshotPolicy

`func (o *DfsPath) HasDpDfsSnapshotPolicy() bool`

HasDpDfsSnapshotPolicy returns a boolean if a field has been set.

### GetId

`func (o *DfsPath) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsPath) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsPath) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsPath) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DfsPath) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsPath) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsPath) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsPath) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRootfs

`func (o *DfsPath) GetRootfs() DfsRootfsNestview`

GetRootfs returns the Rootfs field if non-nil, zero value otherwise.

### GetRootfsOk

`func (o *DfsPath) GetRootfsOk() (*DfsRootfsNestview, bool)`

GetRootfsOk returns a tuple with the Rootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfs

`func (o *DfsPath) SetRootfs(v DfsRootfsNestview)`

SetRootfs sets Rootfs field to given value.

### HasRootfs

`func (o *DfsPath) HasRootfs() bool`

HasRootfs returns a boolean if a field has been set.

### GetStatus

`func (o *DfsPath) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsPath) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsPath) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsPath) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStoragePolicyIds

`func (o *DfsPath) GetStoragePolicyIds() []int64`

GetStoragePolicyIds returns the StoragePolicyIds field if non-nil, zero value otherwise.

### GetStoragePolicyIdsOk

`func (o *DfsPath) GetStoragePolicyIdsOk() (*[]int64, bool)`

GetStoragePolicyIdsOk returns a tuple with the StoragePolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicyIds

`func (o *DfsPath) SetStoragePolicyIds(v []int64)`

SetStoragePolicyIds sets StoragePolicyIds field to given value.

### HasStoragePolicyIds

`func (o *DfsPath) HasStoragePolicyIds() bool`

HasStoragePolicyIds returns a boolean if a field has been set.

### GetStretched

`func (o *DfsPath) GetStretched() bool`

GetStretched returns the Stretched field if non-nil, zero value otherwise.

### GetStretchedOk

`func (o *DfsPath) GetStretchedOk() (*bool, bool)`

GetStretchedOk returns a tuple with the Stretched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStretched

`func (o *DfsPath) SetStretched(v bool)`

SetStretched sets Stretched field to given value.

### HasStretched

`func (o *DfsPath) HasStretched() bool`

HasStretched returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsPath) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsPath) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsPath) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsPath) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


