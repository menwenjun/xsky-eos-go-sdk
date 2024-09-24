# DpBlockBackupJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockSnapshot** | Pointer to [**SnapshotNestview**](SnapshotNestview.md) |  | [optional] 
**BlockVolume** | Pointer to [**VolumeNestview**](VolumeNestview.md) |  | [optional] 
**BlockVolumeSize** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**DiffType** | Pointer to **string** |  | [optional] 
**DpBlockBackupPolicy** | Pointer to [**DpBlockBackupPolicyNestview**](DpBlockBackupPolicyNestview.md) |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**JobSkipReason** | Pointer to **string** |  | [optional] 
**MaxRetryTimes** | Pointer to **int64** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDpBlockBackupJob

`func NewDpBlockBackupJob() *DpBlockBackupJob`

NewDpBlockBackupJob instantiates a new DpBlockBackupJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockBackupJobWithDefaults

`func NewDpBlockBackupJobWithDefaults() *DpBlockBackupJob`

NewDpBlockBackupJobWithDefaults instantiates a new DpBlockBackupJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockSnapshot

`func (o *DpBlockBackupJob) GetBlockSnapshot() SnapshotNestview`

GetBlockSnapshot returns the BlockSnapshot field if non-nil, zero value otherwise.

### GetBlockSnapshotOk

`func (o *DpBlockBackupJob) GetBlockSnapshotOk() (*SnapshotNestview, bool)`

GetBlockSnapshotOk returns a tuple with the BlockSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSnapshot

`func (o *DpBlockBackupJob) SetBlockSnapshot(v SnapshotNestview)`

SetBlockSnapshot sets BlockSnapshot field to given value.

### HasBlockSnapshot

`func (o *DpBlockBackupJob) HasBlockSnapshot() bool`

HasBlockSnapshot returns a boolean if a field has been set.

### GetBlockVolume

`func (o *DpBlockBackupJob) GetBlockVolume() VolumeNestview`

GetBlockVolume returns the BlockVolume field if non-nil, zero value otherwise.

### GetBlockVolumeOk

`func (o *DpBlockBackupJob) GetBlockVolumeOk() (*VolumeNestview, bool)`

GetBlockVolumeOk returns a tuple with the BlockVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolume

`func (o *DpBlockBackupJob) SetBlockVolume(v VolumeNestview)`

SetBlockVolume sets BlockVolume field to given value.

### HasBlockVolume

`func (o *DpBlockBackupJob) HasBlockVolume() bool`

HasBlockVolume returns a boolean if a field has been set.

### GetBlockVolumeSize

`func (o *DpBlockBackupJob) GetBlockVolumeSize() int64`

GetBlockVolumeSize returns the BlockVolumeSize field if non-nil, zero value otherwise.

### GetBlockVolumeSizeOk

`func (o *DpBlockBackupJob) GetBlockVolumeSizeOk() (*int64, bool)`

GetBlockVolumeSizeOk returns a tuple with the BlockVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeSize

`func (o *DpBlockBackupJob) SetBlockVolumeSize(v int64)`

SetBlockVolumeSize sets BlockVolumeSize field to given value.

### HasBlockVolumeSize

`func (o *DpBlockBackupJob) HasBlockVolumeSize() bool`

HasBlockVolumeSize returns a boolean if a field has been set.

### GetCluster

`func (o *DpBlockBackupJob) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpBlockBackupJob) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpBlockBackupJob) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpBlockBackupJob) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDiffType

`func (o *DpBlockBackupJob) GetDiffType() string`

GetDiffType returns the DiffType field if non-nil, zero value otherwise.

### GetDiffTypeOk

`func (o *DpBlockBackupJob) GetDiffTypeOk() (*string, bool)`

GetDiffTypeOk returns a tuple with the DiffType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffType

`func (o *DpBlockBackupJob) SetDiffType(v string)`

SetDiffType sets DiffType field to given value.

### HasDiffType

`func (o *DpBlockBackupJob) HasDiffType() bool`

HasDiffType returns a boolean if a field has been set.

### GetDpBlockBackupPolicy

`func (o *DpBlockBackupJob) GetDpBlockBackupPolicy() DpBlockBackupPolicyNestview`

GetDpBlockBackupPolicy returns the DpBlockBackupPolicy field if non-nil, zero value otherwise.

### GetDpBlockBackupPolicyOk

`func (o *DpBlockBackupJob) GetDpBlockBackupPolicyOk() (*DpBlockBackupPolicyNestview, bool)`

GetDpBlockBackupPolicyOk returns a tuple with the DpBlockBackupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpBlockBackupPolicy

`func (o *DpBlockBackupJob) SetDpBlockBackupPolicy(v DpBlockBackupPolicyNestview)`

SetDpBlockBackupPolicy sets DpBlockBackupPolicy field to given value.

### HasDpBlockBackupPolicy

`func (o *DpBlockBackupJob) HasDpBlockBackupPolicy() bool`

HasDpBlockBackupPolicy returns a boolean if a field has been set.

### GetFinishedAt

`func (o *DpBlockBackupJob) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *DpBlockBackupJob) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *DpBlockBackupJob) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *DpBlockBackupJob) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *DpBlockBackupJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpBlockBackupJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpBlockBackupJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpBlockBackupJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobSkipReason

`func (o *DpBlockBackupJob) GetJobSkipReason() string`

GetJobSkipReason returns the JobSkipReason field if non-nil, zero value otherwise.

### GetJobSkipReasonOk

`func (o *DpBlockBackupJob) GetJobSkipReasonOk() (*string, bool)`

GetJobSkipReasonOk returns a tuple with the JobSkipReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSkipReason

`func (o *DpBlockBackupJob) SetJobSkipReason(v string)`

SetJobSkipReason sets JobSkipReason field to given value.

### HasJobSkipReason

`func (o *DpBlockBackupJob) HasJobSkipReason() bool`

HasJobSkipReason returns a boolean if a field has been set.

### GetMaxRetryTimes

`func (o *DpBlockBackupJob) GetMaxRetryTimes() int64`

GetMaxRetryTimes returns the MaxRetryTimes field if non-nil, zero value otherwise.

### GetMaxRetryTimesOk

`func (o *DpBlockBackupJob) GetMaxRetryTimesOk() (*int64, bool)`

GetMaxRetryTimesOk returns a tuple with the MaxRetryTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetryTimes

`func (o *DpBlockBackupJob) SetMaxRetryTimes(v int64)`

SetMaxRetryTimes sets MaxRetryTimes field to given value.

### HasMaxRetryTimes

`func (o *DpBlockBackupJob) HasMaxRetryTimes() bool`

HasMaxRetryTimes returns a boolean if a field has been set.

### GetProgress

`func (o *DpBlockBackupJob) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DpBlockBackupJob) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DpBlockBackupJob) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *DpBlockBackupJob) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetSize

`func (o *DpBlockBackupJob) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DpBlockBackupJob) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DpBlockBackupJob) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DpBlockBackupJob) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStartedAt

`func (o *DpBlockBackupJob) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DpBlockBackupJob) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DpBlockBackupJob) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DpBlockBackupJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *DpBlockBackupJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpBlockBackupJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpBlockBackupJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpBlockBackupJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DpBlockBackupJob) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DpBlockBackupJob) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DpBlockBackupJob) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DpBlockBackupJob) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


