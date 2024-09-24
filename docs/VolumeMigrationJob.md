# VolumeMigrationJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DestPool** | Pointer to [**PoolNestview**](PoolNestview.md) |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**RemainingTime** | Pointer to **int64** |  | [optional] 
**SourcePool** | Pointer to [**PoolNestview**](PoolNestview.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Stripes** | Pointer to **int64** |  | [optional] 
**Task** | Pointer to [**Task**](Task.md) |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Volume** | Pointer to [**VolumeNestview**](VolumeNestview.md) |  | [optional] 

## Methods

### NewVolumeMigrationJob

`func NewVolumeMigrationJob() *VolumeMigrationJob`

NewVolumeMigrationJob instantiates a new VolumeMigrationJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeMigrationJobWithDefaults

`func NewVolumeMigrationJobWithDefaults() *VolumeMigrationJob`

NewVolumeMigrationJobWithDefaults instantiates a new VolumeMigrationJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *VolumeMigrationJob) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VolumeMigrationJob) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VolumeMigrationJob) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VolumeMigrationJob) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *VolumeMigrationJob) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *VolumeMigrationJob) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *VolumeMigrationJob) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *VolumeMigrationJob) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDestPool

`func (o *VolumeMigrationJob) GetDestPool() PoolNestview`

GetDestPool returns the DestPool field if non-nil, zero value otherwise.

### GetDestPoolOk

`func (o *VolumeMigrationJob) GetDestPoolOk() (*PoolNestview, bool)`

GetDestPoolOk returns a tuple with the DestPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPool

`func (o *VolumeMigrationJob) SetDestPool(v PoolNestview)`

SetDestPool sets DestPool field to given value.

### HasDestPool

`func (o *VolumeMigrationJob) HasDestPool() bool`

HasDestPool returns a boolean if a field has been set.

### GetFinishedAt

`func (o *VolumeMigrationJob) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *VolumeMigrationJob) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *VolumeMigrationJob) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *VolumeMigrationJob) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetHost

`func (o *VolumeMigrationJob) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VolumeMigrationJob) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VolumeMigrationJob) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *VolumeMigrationJob) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *VolumeMigrationJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeMigrationJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeMigrationJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeMigrationJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemainingTime

`func (o *VolumeMigrationJob) GetRemainingTime() int64`

GetRemainingTime returns the RemainingTime field if non-nil, zero value otherwise.

### GetRemainingTimeOk

`func (o *VolumeMigrationJob) GetRemainingTimeOk() (*int64, bool)`

GetRemainingTimeOk returns a tuple with the RemainingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingTime

`func (o *VolumeMigrationJob) SetRemainingTime(v int64)`

SetRemainingTime sets RemainingTime field to given value.

### HasRemainingTime

`func (o *VolumeMigrationJob) HasRemainingTime() bool`

HasRemainingTime returns a boolean if a field has been set.

### GetSourcePool

`func (o *VolumeMigrationJob) GetSourcePool() PoolNestview`

GetSourcePool returns the SourcePool field if non-nil, zero value otherwise.

### GetSourcePoolOk

`func (o *VolumeMigrationJob) GetSourcePoolOk() (*PoolNestview, bool)`

GetSourcePoolOk returns a tuple with the SourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePool

`func (o *VolumeMigrationJob) SetSourcePool(v PoolNestview)`

SetSourcePool sets SourcePool field to given value.

### HasSourcePool

`func (o *VolumeMigrationJob) HasSourcePool() bool`

HasSourcePool returns a boolean if a field has been set.

### GetStartedAt

`func (o *VolumeMigrationJob) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *VolumeMigrationJob) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *VolumeMigrationJob) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *VolumeMigrationJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *VolumeMigrationJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumeMigrationJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumeMigrationJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumeMigrationJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStripes

`func (o *VolumeMigrationJob) GetStripes() int64`

GetStripes returns the Stripes field if non-nil, zero value otherwise.

### GetStripesOk

`func (o *VolumeMigrationJob) GetStripesOk() (*int64, bool)`

GetStripesOk returns a tuple with the Stripes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripes

`func (o *VolumeMigrationJob) SetStripes(v int64)`

SetStripes sets Stripes field to given value.

### HasStripes

`func (o *VolumeMigrationJob) HasStripes() bool`

HasStripes returns a boolean if a field has been set.

### GetTask

`func (o *VolumeMigrationJob) GetTask() Task`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *VolumeMigrationJob) GetTaskOk() (*Task, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *VolumeMigrationJob) SetTask(v Task)`

SetTask sets Task field to given value.

### HasTask

`func (o *VolumeMigrationJob) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetUpdate

`func (o *VolumeMigrationJob) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *VolumeMigrationJob) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *VolumeMigrationJob) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *VolumeMigrationJob) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVolume

`func (o *VolumeMigrationJob) GetVolume() VolumeNestview`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *VolumeMigrationJob) GetVolumeOk() (*VolumeNestview, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *VolumeMigrationJob) SetVolume(v VolumeNestview)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *VolumeMigrationJob) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


