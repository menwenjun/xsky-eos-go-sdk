# DpBlockAsyncReplicationJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**DiffType** | Pointer to **string** |  | [optional] 
**DpBlockAsyncReplicationPair** | Pointer to [**DpBlockAsyncReplicationPairNestview**](DpBlockAsyncReplicationPairNestview.md) |  | [optional] 
**DpVolumeGroupSnapshotReplicationJob** | Pointer to [**DpVolumeGroupSnapshotReplicationJobNestview**](DpVolumeGroupSnapshotReplicationJobNestview.md) |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MaxRetryTimes** | Pointer to **int64** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Snapshot** | Pointer to [**SnapshotNestview**](SnapshotNestview.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**VolumeName** | Pointer to **string** |  | [optional] 

## Methods

### NewDpBlockAsyncReplicationJob

`func NewDpBlockAsyncReplicationJob() *DpBlockAsyncReplicationJob`

NewDpBlockAsyncReplicationJob instantiates a new DpBlockAsyncReplicationJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockAsyncReplicationJobWithDefaults

`func NewDpBlockAsyncReplicationJobWithDefaults() *DpBlockAsyncReplicationJob`

NewDpBlockAsyncReplicationJobWithDefaults instantiates a new DpBlockAsyncReplicationJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DpBlockAsyncReplicationJob) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpBlockAsyncReplicationJob) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpBlockAsyncReplicationJob) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpBlockAsyncReplicationJob) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDiffType

`func (o *DpBlockAsyncReplicationJob) GetDiffType() string`

GetDiffType returns the DiffType field if non-nil, zero value otherwise.

### GetDiffTypeOk

`func (o *DpBlockAsyncReplicationJob) GetDiffTypeOk() (*string, bool)`

GetDiffTypeOk returns a tuple with the DiffType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffType

`func (o *DpBlockAsyncReplicationJob) SetDiffType(v string)`

SetDiffType sets DiffType field to given value.

### HasDiffType

`func (o *DpBlockAsyncReplicationJob) HasDiffType() bool`

HasDiffType returns a boolean if a field has been set.

### GetDpBlockAsyncReplicationPair

`func (o *DpBlockAsyncReplicationJob) GetDpBlockAsyncReplicationPair() DpBlockAsyncReplicationPairNestview`

GetDpBlockAsyncReplicationPair returns the DpBlockAsyncReplicationPair field if non-nil, zero value otherwise.

### GetDpBlockAsyncReplicationPairOk

`func (o *DpBlockAsyncReplicationJob) GetDpBlockAsyncReplicationPairOk() (*DpBlockAsyncReplicationPairNestview, bool)`

GetDpBlockAsyncReplicationPairOk returns a tuple with the DpBlockAsyncReplicationPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpBlockAsyncReplicationPair

`func (o *DpBlockAsyncReplicationJob) SetDpBlockAsyncReplicationPair(v DpBlockAsyncReplicationPairNestview)`

SetDpBlockAsyncReplicationPair sets DpBlockAsyncReplicationPair field to given value.

### HasDpBlockAsyncReplicationPair

`func (o *DpBlockAsyncReplicationJob) HasDpBlockAsyncReplicationPair() bool`

HasDpBlockAsyncReplicationPair returns a boolean if a field has been set.

### GetDpVolumeGroupSnapshotReplicationJob

`func (o *DpBlockAsyncReplicationJob) GetDpVolumeGroupSnapshotReplicationJob() DpVolumeGroupSnapshotReplicationJobNestview`

GetDpVolumeGroupSnapshotReplicationJob returns the DpVolumeGroupSnapshotReplicationJob field if non-nil, zero value otherwise.

### GetDpVolumeGroupSnapshotReplicationJobOk

`func (o *DpBlockAsyncReplicationJob) GetDpVolumeGroupSnapshotReplicationJobOk() (*DpVolumeGroupSnapshotReplicationJobNestview, bool)`

GetDpVolumeGroupSnapshotReplicationJobOk returns a tuple with the DpVolumeGroupSnapshotReplicationJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpVolumeGroupSnapshotReplicationJob

`func (o *DpBlockAsyncReplicationJob) SetDpVolumeGroupSnapshotReplicationJob(v DpVolumeGroupSnapshotReplicationJobNestview)`

SetDpVolumeGroupSnapshotReplicationJob sets DpVolumeGroupSnapshotReplicationJob field to given value.

### HasDpVolumeGroupSnapshotReplicationJob

`func (o *DpBlockAsyncReplicationJob) HasDpVolumeGroupSnapshotReplicationJob() bool`

HasDpVolumeGroupSnapshotReplicationJob returns a boolean if a field has been set.

### GetFinishedAt

`func (o *DpBlockAsyncReplicationJob) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *DpBlockAsyncReplicationJob) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *DpBlockAsyncReplicationJob) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *DpBlockAsyncReplicationJob) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *DpBlockAsyncReplicationJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpBlockAsyncReplicationJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpBlockAsyncReplicationJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpBlockAsyncReplicationJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxRetryTimes

`func (o *DpBlockAsyncReplicationJob) GetMaxRetryTimes() int64`

GetMaxRetryTimes returns the MaxRetryTimes field if non-nil, zero value otherwise.

### GetMaxRetryTimesOk

`func (o *DpBlockAsyncReplicationJob) GetMaxRetryTimesOk() (*int64, bool)`

GetMaxRetryTimesOk returns a tuple with the MaxRetryTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetryTimes

`func (o *DpBlockAsyncReplicationJob) SetMaxRetryTimes(v int64)`

SetMaxRetryTimes sets MaxRetryTimes field to given value.

### HasMaxRetryTimes

`func (o *DpBlockAsyncReplicationJob) HasMaxRetryTimes() bool`

HasMaxRetryTimes returns a boolean if a field has been set.

### GetProgress

`func (o *DpBlockAsyncReplicationJob) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DpBlockAsyncReplicationJob) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DpBlockAsyncReplicationJob) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *DpBlockAsyncReplicationJob) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetSize

`func (o *DpBlockAsyncReplicationJob) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DpBlockAsyncReplicationJob) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DpBlockAsyncReplicationJob) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DpBlockAsyncReplicationJob) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSnapshot

`func (o *DpBlockAsyncReplicationJob) GetSnapshot() SnapshotNestview`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *DpBlockAsyncReplicationJob) GetSnapshotOk() (*SnapshotNestview, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *DpBlockAsyncReplicationJob) SetSnapshot(v SnapshotNestview)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *DpBlockAsyncReplicationJob) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetStartedAt

`func (o *DpBlockAsyncReplicationJob) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DpBlockAsyncReplicationJob) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DpBlockAsyncReplicationJob) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DpBlockAsyncReplicationJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *DpBlockAsyncReplicationJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpBlockAsyncReplicationJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpBlockAsyncReplicationJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpBlockAsyncReplicationJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DpBlockAsyncReplicationJob) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DpBlockAsyncReplicationJob) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DpBlockAsyncReplicationJob) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DpBlockAsyncReplicationJob) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVolumeName

`func (o *DpBlockAsyncReplicationJob) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *DpBlockAsyncReplicationJob) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *DpBlockAsyncReplicationJob) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *DpBlockAsyncReplicationJob) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


