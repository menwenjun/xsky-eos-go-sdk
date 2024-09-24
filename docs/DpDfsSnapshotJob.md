# DpDfsSnapshotJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**DfsPath** | Pointer to [**DfsPath**](DfsPath.md) |  | [optional] 
**DfsSnapshot** | Pointer to [**DfsSnapshotNestview**](DfsSnapshotNestview.md) |  | [optional] 
**DpDfsSnapshot** | Pointer to [**DpDfsSnapshotNestview**](DpDfsSnapshotNestview.md) |  | [optional] 
**DpDfsSnapshotPolicy** | Pointer to [**DpDfsSnapshotPolicy**](DpDfsSnapshotPolicy.md) |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewDpDfsSnapshotJob

`func NewDpDfsSnapshotJob() *DpDfsSnapshotJob`

NewDpDfsSnapshotJob instantiates a new DpDfsSnapshotJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpDfsSnapshotJobWithDefaults

`func NewDpDfsSnapshotJobWithDefaults() *DpDfsSnapshotJob`

NewDpDfsSnapshotJobWithDefaults instantiates a new DpDfsSnapshotJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DpDfsSnapshotJob) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpDfsSnapshotJob) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpDfsSnapshotJob) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpDfsSnapshotJob) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDfsPath

`func (o *DpDfsSnapshotJob) GetDfsPath() DfsPath`

GetDfsPath returns the DfsPath field if non-nil, zero value otherwise.

### GetDfsPathOk

`func (o *DpDfsSnapshotJob) GetDfsPathOk() (*DfsPath, bool)`

GetDfsPathOk returns a tuple with the DfsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPath

`func (o *DpDfsSnapshotJob) SetDfsPath(v DfsPath)`

SetDfsPath sets DfsPath field to given value.

### HasDfsPath

`func (o *DpDfsSnapshotJob) HasDfsPath() bool`

HasDfsPath returns a boolean if a field has been set.

### GetDfsSnapshot

`func (o *DpDfsSnapshotJob) GetDfsSnapshot() DfsSnapshotNestview`

GetDfsSnapshot returns the DfsSnapshot field if non-nil, zero value otherwise.

### GetDfsSnapshotOk

`func (o *DpDfsSnapshotJob) GetDfsSnapshotOk() (*DfsSnapshotNestview, bool)`

GetDfsSnapshotOk returns a tuple with the DfsSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsSnapshot

`func (o *DpDfsSnapshotJob) SetDfsSnapshot(v DfsSnapshotNestview)`

SetDfsSnapshot sets DfsSnapshot field to given value.

### HasDfsSnapshot

`func (o *DpDfsSnapshotJob) HasDfsSnapshot() bool`

HasDfsSnapshot returns a boolean if a field has been set.

### GetDpDfsSnapshot

`func (o *DpDfsSnapshotJob) GetDpDfsSnapshot() DpDfsSnapshotNestview`

GetDpDfsSnapshot returns the DpDfsSnapshot field if non-nil, zero value otherwise.

### GetDpDfsSnapshotOk

`func (o *DpDfsSnapshotJob) GetDpDfsSnapshotOk() (*DpDfsSnapshotNestview, bool)`

GetDpDfsSnapshotOk returns a tuple with the DpDfsSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpDfsSnapshot

`func (o *DpDfsSnapshotJob) SetDpDfsSnapshot(v DpDfsSnapshotNestview)`

SetDpDfsSnapshot sets DpDfsSnapshot field to given value.

### HasDpDfsSnapshot

`func (o *DpDfsSnapshotJob) HasDpDfsSnapshot() bool`

HasDpDfsSnapshot returns a boolean if a field has been set.

### GetDpDfsSnapshotPolicy

`func (o *DpDfsSnapshotJob) GetDpDfsSnapshotPolicy() DpDfsSnapshotPolicy`

GetDpDfsSnapshotPolicy returns the DpDfsSnapshotPolicy field if non-nil, zero value otherwise.

### GetDpDfsSnapshotPolicyOk

`func (o *DpDfsSnapshotJob) GetDpDfsSnapshotPolicyOk() (*DpDfsSnapshotPolicy, bool)`

GetDpDfsSnapshotPolicyOk returns a tuple with the DpDfsSnapshotPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpDfsSnapshotPolicy

`func (o *DpDfsSnapshotJob) SetDpDfsSnapshotPolicy(v DpDfsSnapshotPolicy)`

SetDpDfsSnapshotPolicy sets DpDfsSnapshotPolicy field to given value.

### HasDpDfsSnapshotPolicy

`func (o *DpDfsSnapshotJob) HasDpDfsSnapshotPolicy() bool`

HasDpDfsSnapshotPolicy returns a boolean if a field has been set.

### GetFinishedAt

`func (o *DpDfsSnapshotJob) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *DpDfsSnapshotJob) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *DpDfsSnapshotJob) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *DpDfsSnapshotJob) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *DpDfsSnapshotJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpDfsSnapshotJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpDfsSnapshotJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpDfsSnapshotJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartedAt

`func (o *DpDfsSnapshotJob) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DpDfsSnapshotJob) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DpDfsSnapshotJob) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DpDfsSnapshotJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *DpDfsSnapshotJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpDfsSnapshotJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpDfsSnapshotJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpDfsSnapshotJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


