# DpBlockSnapshotJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockSnapshot** | Pointer to [**SnapshotNestview**](SnapshotNestview.md) |  | [optional] 
**BlockVolume** | Pointer to [**VolumeNestview**](VolumeNestview.md) |  | [optional] 
**BlockVolumeSize** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**DiffType** | Pointer to **string** |  | [optional] 
**DpBlockSnapshotPolicy** | Pointer to [**DpBlockSnapshotPolicyNestview**](DpBlockSnapshotPolicyNestview.md) |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDpBlockSnapshotJob

`func NewDpBlockSnapshotJob() *DpBlockSnapshotJob`

NewDpBlockSnapshotJob instantiates a new DpBlockSnapshotJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockSnapshotJobWithDefaults

`func NewDpBlockSnapshotJobWithDefaults() *DpBlockSnapshotJob`

NewDpBlockSnapshotJobWithDefaults instantiates a new DpBlockSnapshotJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockSnapshot

`func (o *DpBlockSnapshotJob) GetBlockSnapshot() SnapshotNestview`

GetBlockSnapshot returns the BlockSnapshot field if non-nil, zero value otherwise.

### GetBlockSnapshotOk

`func (o *DpBlockSnapshotJob) GetBlockSnapshotOk() (*SnapshotNestview, bool)`

GetBlockSnapshotOk returns a tuple with the BlockSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSnapshot

`func (o *DpBlockSnapshotJob) SetBlockSnapshot(v SnapshotNestview)`

SetBlockSnapshot sets BlockSnapshot field to given value.

### HasBlockSnapshot

`func (o *DpBlockSnapshotJob) HasBlockSnapshot() bool`

HasBlockSnapshot returns a boolean if a field has been set.

### GetBlockVolume

`func (o *DpBlockSnapshotJob) GetBlockVolume() VolumeNestview`

GetBlockVolume returns the BlockVolume field if non-nil, zero value otherwise.

### GetBlockVolumeOk

`func (o *DpBlockSnapshotJob) GetBlockVolumeOk() (*VolumeNestview, bool)`

GetBlockVolumeOk returns a tuple with the BlockVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolume

`func (o *DpBlockSnapshotJob) SetBlockVolume(v VolumeNestview)`

SetBlockVolume sets BlockVolume field to given value.

### HasBlockVolume

`func (o *DpBlockSnapshotJob) HasBlockVolume() bool`

HasBlockVolume returns a boolean if a field has been set.

### GetBlockVolumeSize

`func (o *DpBlockSnapshotJob) GetBlockVolumeSize() int64`

GetBlockVolumeSize returns the BlockVolumeSize field if non-nil, zero value otherwise.

### GetBlockVolumeSizeOk

`func (o *DpBlockSnapshotJob) GetBlockVolumeSizeOk() (*int64, bool)`

GetBlockVolumeSizeOk returns a tuple with the BlockVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeSize

`func (o *DpBlockSnapshotJob) SetBlockVolumeSize(v int64)`

SetBlockVolumeSize sets BlockVolumeSize field to given value.

### HasBlockVolumeSize

`func (o *DpBlockSnapshotJob) HasBlockVolumeSize() bool`

HasBlockVolumeSize returns a boolean if a field has been set.

### GetCluster

`func (o *DpBlockSnapshotJob) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpBlockSnapshotJob) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpBlockSnapshotJob) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpBlockSnapshotJob) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDiffType

`func (o *DpBlockSnapshotJob) GetDiffType() string`

GetDiffType returns the DiffType field if non-nil, zero value otherwise.

### GetDiffTypeOk

`func (o *DpBlockSnapshotJob) GetDiffTypeOk() (*string, bool)`

GetDiffTypeOk returns a tuple with the DiffType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffType

`func (o *DpBlockSnapshotJob) SetDiffType(v string)`

SetDiffType sets DiffType field to given value.

### HasDiffType

`func (o *DpBlockSnapshotJob) HasDiffType() bool`

HasDiffType returns a boolean if a field has been set.

### GetDpBlockSnapshotPolicy

`func (o *DpBlockSnapshotJob) GetDpBlockSnapshotPolicy() DpBlockSnapshotPolicyNestview`

GetDpBlockSnapshotPolicy returns the DpBlockSnapshotPolicy field if non-nil, zero value otherwise.

### GetDpBlockSnapshotPolicyOk

`func (o *DpBlockSnapshotJob) GetDpBlockSnapshotPolicyOk() (*DpBlockSnapshotPolicyNestview, bool)`

GetDpBlockSnapshotPolicyOk returns a tuple with the DpBlockSnapshotPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpBlockSnapshotPolicy

`func (o *DpBlockSnapshotJob) SetDpBlockSnapshotPolicy(v DpBlockSnapshotPolicyNestview)`

SetDpBlockSnapshotPolicy sets DpBlockSnapshotPolicy field to given value.

### HasDpBlockSnapshotPolicy

`func (o *DpBlockSnapshotJob) HasDpBlockSnapshotPolicy() bool`

HasDpBlockSnapshotPolicy returns a boolean if a field has been set.

### GetFinishedAt

`func (o *DpBlockSnapshotJob) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *DpBlockSnapshotJob) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *DpBlockSnapshotJob) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *DpBlockSnapshotJob) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *DpBlockSnapshotJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpBlockSnapshotJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpBlockSnapshotJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpBlockSnapshotJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProgress

`func (o *DpBlockSnapshotJob) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DpBlockSnapshotJob) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DpBlockSnapshotJob) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *DpBlockSnapshotJob) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetSize

`func (o *DpBlockSnapshotJob) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DpBlockSnapshotJob) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DpBlockSnapshotJob) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DpBlockSnapshotJob) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStartedAt

`func (o *DpBlockSnapshotJob) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DpBlockSnapshotJob) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DpBlockSnapshotJob) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DpBlockSnapshotJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *DpBlockSnapshotJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpBlockSnapshotJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpBlockSnapshotJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpBlockSnapshotJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DpBlockSnapshotJob) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DpBlockSnapshotJob) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DpBlockSnapshotJob) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DpBlockSnapshotJob) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


