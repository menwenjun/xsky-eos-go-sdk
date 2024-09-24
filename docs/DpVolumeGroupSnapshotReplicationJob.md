# DpVolumeGroupSnapshotReplicationJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**DiffType** | Pointer to **string** |  | [optional] 
**DpVolumeGroupSnapshotReplicationPair** | Pointer to [**DpVolumeGroupSnapshotReplicationPairNestview**](DpVolumeGroupSnapshotReplicationPairNestview.md) |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MaxRetryTimes** | Pointer to **int64** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Snapshot** | Pointer to [**VolumeGroupSnapshotNestview**](VolumeGroupSnapshotNestview.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**VolumeGroupName** | Pointer to **string** |  | [optional] 

## Methods

### NewDpVolumeGroupSnapshotReplicationJob

`func NewDpVolumeGroupSnapshotReplicationJob() *DpVolumeGroupSnapshotReplicationJob`

NewDpVolumeGroupSnapshotReplicationJob instantiates a new DpVolumeGroupSnapshotReplicationJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpVolumeGroupSnapshotReplicationJobWithDefaults

`func NewDpVolumeGroupSnapshotReplicationJobWithDefaults() *DpVolumeGroupSnapshotReplicationJob`

NewDpVolumeGroupSnapshotReplicationJobWithDefaults instantiates a new DpVolumeGroupSnapshotReplicationJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DpVolumeGroupSnapshotReplicationJob) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpVolumeGroupSnapshotReplicationJob) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpVolumeGroupSnapshotReplicationJob) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpVolumeGroupSnapshotReplicationJob) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDiffType

`func (o *DpVolumeGroupSnapshotReplicationJob) GetDiffType() string`

GetDiffType returns the DiffType field if non-nil, zero value otherwise.

### GetDiffTypeOk

`func (o *DpVolumeGroupSnapshotReplicationJob) GetDiffTypeOk() (*string, bool)`

GetDiffTypeOk returns a tuple with the DiffType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffType

`func (o *DpVolumeGroupSnapshotReplicationJob) SetDiffType(v string)`

SetDiffType sets DiffType field to given value.

### HasDiffType

`func (o *DpVolumeGroupSnapshotReplicationJob) HasDiffType() bool`

HasDiffType returns a boolean if a field has been set.

### GetDpVolumeGroupSnapshotReplicationPair

`func (o *DpVolumeGroupSnapshotReplicationJob) GetDpVolumeGroupSnapshotReplicationPair() DpVolumeGroupSnapshotReplicationPairNestview`

GetDpVolumeGroupSnapshotReplicationPair returns the DpVolumeGroupSnapshotReplicationPair field if non-nil, zero value otherwise.

### GetDpVolumeGroupSnapshotReplicationPairOk

`func (o *DpVolumeGroupSnapshotReplicationJob) GetDpVolumeGroupSnapshotReplicationPairOk() (*DpVolumeGroupSnapshotReplicationPairNestview, bool)`

GetDpVolumeGroupSnapshotReplicationPairOk returns a tuple with the DpVolumeGroupSnapshotReplicationPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpVolumeGroupSnapshotReplicationPair

`func (o *DpVolumeGroupSnapshotReplicationJob) SetDpVolumeGroupSnapshotReplicationPair(v DpVolumeGroupSnapshotReplicationPairNestview)`

SetDpVolumeGroupSnapshotReplicationPair sets DpVolumeGroupSnapshotReplicationPair field to given value.

### HasDpVolumeGroupSnapshotReplicationPair

`func (o *DpVolumeGroupSnapshotReplicationJob) HasDpVolumeGroupSnapshotReplicationPair() bool`

HasDpVolumeGroupSnapshotReplicationPair returns a boolean if a field has been set.

### GetFinishedAt

`func (o *DpVolumeGroupSnapshotReplicationJob) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *DpVolumeGroupSnapshotReplicationJob) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *DpVolumeGroupSnapshotReplicationJob) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *DpVolumeGroupSnapshotReplicationJob) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *DpVolumeGroupSnapshotReplicationJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpVolumeGroupSnapshotReplicationJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpVolumeGroupSnapshotReplicationJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpVolumeGroupSnapshotReplicationJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxRetryTimes

`func (o *DpVolumeGroupSnapshotReplicationJob) GetMaxRetryTimes() int64`

GetMaxRetryTimes returns the MaxRetryTimes field if non-nil, zero value otherwise.

### GetMaxRetryTimesOk

`func (o *DpVolumeGroupSnapshotReplicationJob) GetMaxRetryTimesOk() (*int64, bool)`

GetMaxRetryTimesOk returns a tuple with the MaxRetryTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetryTimes

`func (o *DpVolumeGroupSnapshotReplicationJob) SetMaxRetryTimes(v int64)`

SetMaxRetryTimes sets MaxRetryTimes field to given value.

### HasMaxRetryTimes

`func (o *DpVolumeGroupSnapshotReplicationJob) HasMaxRetryTimes() bool`

HasMaxRetryTimes returns a boolean if a field has been set.

### GetProgress

`func (o *DpVolumeGroupSnapshotReplicationJob) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DpVolumeGroupSnapshotReplicationJob) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DpVolumeGroupSnapshotReplicationJob) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *DpVolumeGroupSnapshotReplicationJob) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetSize

`func (o *DpVolumeGroupSnapshotReplicationJob) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DpVolumeGroupSnapshotReplicationJob) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DpVolumeGroupSnapshotReplicationJob) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DpVolumeGroupSnapshotReplicationJob) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSnapshot

`func (o *DpVolumeGroupSnapshotReplicationJob) GetSnapshot() VolumeGroupSnapshotNestview`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *DpVolumeGroupSnapshotReplicationJob) GetSnapshotOk() (*VolumeGroupSnapshotNestview, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *DpVolumeGroupSnapshotReplicationJob) SetSnapshot(v VolumeGroupSnapshotNestview)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *DpVolumeGroupSnapshotReplicationJob) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetStartedAt

`func (o *DpVolumeGroupSnapshotReplicationJob) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DpVolumeGroupSnapshotReplicationJob) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DpVolumeGroupSnapshotReplicationJob) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DpVolumeGroupSnapshotReplicationJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *DpVolumeGroupSnapshotReplicationJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpVolumeGroupSnapshotReplicationJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpVolumeGroupSnapshotReplicationJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpVolumeGroupSnapshotReplicationJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DpVolumeGroupSnapshotReplicationJob) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DpVolumeGroupSnapshotReplicationJob) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DpVolumeGroupSnapshotReplicationJob) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DpVolumeGroupSnapshotReplicationJob) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVolumeGroupName

`func (o *DpVolumeGroupSnapshotReplicationJob) GetVolumeGroupName() string`

GetVolumeGroupName returns the VolumeGroupName field if non-nil, zero value otherwise.

### GetVolumeGroupNameOk

`func (o *DpVolumeGroupSnapshotReplicationJob) GetVolumeGroupNameOk() (*string, bool)`

GetVolumeGroupNameOk returns a tuple with the VolumeGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupName

`func (o *DpVolumeGroupSnapshotReplicationJob) SetVolumeGroupName(v string)`

SetVolumeGroupName sets VolumeGroupName field to given value.

### HasVolumeGroupName

`func (o *DpVolumeGroupSnapshotReplicationJob) HasVolumeGroupName() bool`

HasVolumeGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


