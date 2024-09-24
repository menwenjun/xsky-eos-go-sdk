# DpBlockSnapshotRecoveryJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupBlockSnapshot** | Pointer to [**DpBackupBlockSnapshot**](DpBackupBlockSnapshot.md) |  | [optional] 
**BackupBlockVolume** | Pointer to [**DpBackupBlockVolume**](DpBackupBlockVolume.md) |  | [optional] 
**BackupCluster** | Pointer to [**DpBackupCluster**](DpBackupCluster.md) |  | [optional] 
**BlockVolume** | Pointer to [**VolumeNestview**](VolumeNestview.md) |  | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DpGateway** | Pointer to [**DpGatewayNestview**](DpGatewayNestview.md) |  | [optional] 
**DpSite** | Pointer to [**DpSiteNestview**](DpSiteNestview.md) |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MaxRetryTimes** | Pointer to **int64** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDpBlockSnapshotRecoveryJob

`func NewDpBlockSnapshotRecoveryJob() *DpBlockSnapshotRecoveryJob`

NewDpBlockSnapshotRecoveryJob instantiates a new DpBlockSnapshotRecoveryJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockSnapshotRecoveryJobWithDefaults

`func NewDpBlockSnapshotRecoveryJobWithDefaults() *DpBlockSnapshotRecoveryJob`

NewDpBlockSnapshotRecoveryJobWithDefaults instantiates a new DpBlockSnapshotRecoveryJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupBlockSnapshot

`func (o *DpBlockSnapshotRecoveryJob) GetBackupBlockSnapshot() DpBackupBlockSnapshot`

GetBackupBlockSnapshot returns the BackupBlockSnapshot field if non-nil, zero value otherwise.

### GetBackupBlockSnapshotOk

`func (o *DpBlockSnapshotRecoveryJob) GetBackupBlockSnapshotOk() (*DpBackupBlockSnapshot, bool)`

GetBackupBlockSnapshotOk returns a tuple with the BackupBlockSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupBlockSnapshot

`func (o *DpBlockSnapshotRecoveryJob) SetBackupBlockSnapshot(v DpBackupBlockSnapshot)`

SetBackupBlockSnapshot sets BackupBlockSnapshot field to given value.

### HasBackupBlockSnapshot

`func (o *DpBlockSnapshotRecoveryJob) HasBackupBlockSnapshot() bool`

HasBackupBlockSnapshot returns a boolean if a field has been set.

### GetBackupBlockVolume

`func (o *DpBlockSnapshotRecoveryJob) GetBackupBlockVolume() DpBackupBlockVolume`

GetBackupBlockVolume returns the BackupBlockVolume field if non-nil, zero value otherwise.

### GetBackupBlockVolumeOk

`func (o *DpBlockSnapshotRecoveryJob) GetBackupBlockVolumeOk() (*DpBackupBlockVolume, bool)`

GetBackupBlockVolumeOk returns a tuple with the BackupBlockVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupBlockVolume

`func (o *DpBlockSnapshotRecoveryJob) SetBackupBlockVolume(v DpBackupBlockVolume)`

SetBackupBlockVolume sets BackupBlockVolume field to given value.

### HasBackupBlockVolume

`func (o *DpBlockSnapshotRecoveryJob) HasBackupBlockVolume() bool`

HasBackupBlockVolume returns a boolean if a field has been set.

### GetBackupCluster

`func (o *DpBlockSnapshotRecoveryJob) GetBackupCluster() DpBackupCluster`

GetBackupCluster returns the BackupCluster field if non-nil, zero value otherwise.

### GetBackupClusterOk

`func (o *DpBlockSnapshotRecoveryJob) GetBackupClusterOk() (*DpBackupCluster, bool)`

GetBackupClusterOk returns a tuple with the BackupCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCluster

`func (o *DpBlockSnapshotRecoveryJob) SetBackupCluster(v DpBackupCluster)`

SetBackupCluster sets BackupCluster field to given value.

### HasBackupCluster

`func (o *DpBlockSnapshotRecoveryJob) HasBackupCluster() bool`

HasBackupCluster returns a boolean if a field has been set.

### GetBlockVolume

`func (o *DpBlockSnapshotRecoveryJob) GetBlockVolume() VolumeNestview`

GetBlockVolume returns the BlockVolume field if non-nil, zero value otherwise.

### GetBlockVolumeOk

`func (o *DpBlockSnapshotRecoveryJob) GetBlockVolumeOk() (*VolumeNestview, bool)`

GetBlockVolumeOk returns a tuple with the BlockVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolume

`func (o *DpBlockSnapshotRecoveryJob) SetBlockVolume(v VolumeNestview)`

SetBlockVolume sets BlockVolume field to given value.

### HasBlockVolume

`func (o *DpBlockSnapshotRecoveryJob) HasBlockVolume() bool`

HasBlockVolume returns a boolean if a field has been set.

### GetCluster

`func (o *DpBlockSnapshotRecoveryJob) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpBlockSnapshotRecoveryJob) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpBlockSnapshotRecoveryJob) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpBlockSnapshotRecoveryJob) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DpBlockSnapshotRecoveryJob) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DpBlockSnapshotRecoveryJob) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DpBlockSnapshotRecoveryJob) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DpBlockSnapshotRecoveryJob) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDpGateway

`func (o *DpBlockSnapshotRecoveryJob) GetDpGateway() DpGatewayNestview`

GetDpGateway returns the DpGateway field if non-nil, zero value otherwise.

### GetDpGatewayOk

`func (o *DpBlockSnapshotRecoveryJob) GetDpGatewayOk() (*DpGatewayNestview, bool)`

GetDpGatewayOk returns a tuple with the DpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpGateway

`func (o *DpBlockSnapshotRecoveryJob) SetDpGateway(v DpGatewayNestview)`

SetDpGateway sets DpGateway field to given value.

### HasDpGateway

`func (o *DpBlockSnapshotRecoveryJob) HasDpGateway() bool`

HasDpGateway returns a boolean if a field has been set.

### GetDpSite

`func (o *DpBlockSnapshotRecoveryJob) GetDpSite() DpSiteNestview`

GetDpSite returns the DpSite field if non-nil, zero value otherwise.

### GetDpSiteOk

`func (o *DpBlockSnapshotRecoveryJob) GetDpSiteOk() (*DpSiteNestview, bool)`

GetDpSiteOk returns a tuple with the DpSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpSite

`func (o *DpBlockSnapshotRecoveryJob) SetDpSite(v DpSiteNestview)`

SetDpSite sets DpSite field to given value.

### HasDpSite

`func (o *DpBlockSnapshotRecoveryJob) HasDpSite() bool`

HasDpSite returns a boolean if a field has been set.

### GetFinishedAt

`func (o *DpBlockSnapshotRecoveryJob) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *DpBlockSnapshotRecoveryJob) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *DpBlockSnapshotRecoveryJob) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *DpBlockSnapshotRecoveryJob) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *DpBlockSnapshotRecoveryJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpBlockSnapshotRecoveryJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpBlockSnapshotRecoveryJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpBlockSnapshotRecoveryJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxRetryTimes

`func (o *DpBlockSnapshotRecoveryJob) GetMaxRetryTimes() int64`

GetMaxRetryTimes returns the MaxRetryTimes field if non-nil, zero value otherwise.

### GetMaxRetryTimesOk

`func (o *DpBlockSnapshotRecoveryJob) GetMaxRetryTimesOk() (*int64, bool)`

GetMaxRetryTimesOk returns a tuple with the MaxRetryTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetryTimes

`func (o *DpBlockSnapshotRecoveryJob) SetMaxRetryTimes(v int64)`

SetMaxRetryTimes sets MaxRetryTimes field to given value.

### HasMaxRetryTimes

`func (o *DpBlockSnapshotRecoveryJob) HasMaxRetryTimes() bool`

HasMaxRetryTimes returns a boolean if a field has been set.

### GetProgress

`func (o *DpBlockSnapshotRecoveryJob) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DpBlockSnapshotRecoveryJob) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DpBlockSnapshotRecoveryJob) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *DpBlockSnapshotRecoveryJob) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetResourceType

`func (o *DpBlockSnapshotRecoveryJob) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DpBlockSnapshotRecoveryJob) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DpBlockSnapshotRecoveryJob) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *DpBlockSnapshotRecoveryJob) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStartedAt

`func (o *DpBlockSnapshotRecoveryJob) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DpBlockSnapshotRecoveryJob) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DpBlockSnapshotRecoveryJob) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DpBlockSnapshotRecoveryJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *DpBlockSnapshotRecoveryJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpBlockSnapshotRecoveryJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpBlockSnapshotRecoveryJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpBlockSnapshotRecoveryJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DpBlockSnapshotRecoveryJob) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DpBlockSnapshotRecoveryJob) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DpBlockSnapshotRecoveryJob) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DpBlockSnapshotRecoveryJob) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DpBlockSnapshotRecoveryJob) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DpBlockSnapshotRecoveryJob) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DpBlockSnapshotRecoveryJob) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DpBlockSnapshotRecoveryJob) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


