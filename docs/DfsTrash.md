# DfsTrash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**ExpiredTime** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Path** | Pointer to [**DfsPathNestview**](DfsPathNestview.md) |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**ProgressInfo** | Pointer to [**ProgressInfo**](ProgressInfo.md) |  | [optional] 
**Rootfs** | Pointer to [**DfsRootfsNestview**](DfsRootfsNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalFiles** | Pointer to **int64** |  | [optional] 
**TotalKbyte** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsTrash

`func NewDfsTrash() *DfsTrash`

NewDfsTrash instantiates a new DfsTrash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsTrashWithDefaults

`func NewDfsTrashWithDefaults() *DfsTrash`

NewDfsTrashWithDefaults instantiates a new DfsTrash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsTrash) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsTrash) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsTrash) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsTrash) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *DfsTrash) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsTrash) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsTrash) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsTrash) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsTrash) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsTrash) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsTrash) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsTrash) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetExpiredTime

`func (o *DfsTrash) GetExpiredTime() int64`

GetExpiredTime returns the ExpiredTime field if non-nil, zero value otherwise.

### GetExpiredTimeOk

`func (o *DfsTrash) GetExpiredTimeOk() (*int64, bool)`

GetExpiredTimeOk returns a tuple with the ExpiredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTime

`func (o *DfsTrash) SetExpiredTime(v int64)`

SetExpiredTime sets ExpiredTime field to given value.

### HasExpiredTime

`func (o *DfsTrash) HasExpiredTime() bool`

HasExpiredTime returns a boolean if a field has been set.

### GetId

`func (o *DfsTrash) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsTrash) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsTrash) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsTrash) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *DfsTrash) GetPath() DfsPathNestview`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsTrash) GetPathOk() (*DfsPathNestview, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsTrash) SetPath(v DfsPathNestview)`

SetPath sets Path field to given value.

### HasPath

`func (o *DfsTrash) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProgress

`func (o *DfsTrash) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DfsTrash) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DfsTrash) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *DfsTrash) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProgressInfo

`func (o *DfsTrash) GetProgressInfo() ProgressInfo`

GetProgressInfo returns the ProgressInfo field if non-nil, zero value otherwise.

### GetProgressInfoOk

`func (o *DfsTrash) GetProgressInfoOk() (*ProgressInfo, bool)`

GetProgressInfoOk returns a tuple with the ProgressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressInfo

`func (o *DfsTrash) SetProgressInfo(v ProgressInfo)`

SetProgressInfo sets ProgressInfo field to given value.

### HasProgressInfo

`func (o *DfsTrash) HasProgressInfo() bool`

HasProgressInfo returns a boolean if a field has been set.

### GetRootfs

`func (o *DfsTrash) GetRootfs() DfsRootfsNestview`

GetRootfs returns the Rootfs field if non-nil, zero value otherwise.

### GetRootfsOk

`func (o *DfsTrash) GetRootfsOk() (*DfsRootfsNestview, bool)`

GetRootfsOk returns a tuple with the Rootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfs

`func (o *DfsTrash) SetRootfs(v DfsRootfsNestview)`

SetRootfs sets Rootfs field to given value.

### HasRootfs

`func (o *DfsTrash) HasRootfs() bool`

HasRootfs returns a boolean if a field has been set.

### GetStatus

`func (o *DfsTrash) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsTrash) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsTrash) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsTrash) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalFiles

`func (o *DfsTrash) GetTotalFiles() int64`

GetTotalFiles returns the TotalFiles field if non-nil, zero value otherwise.

### GetTotalFilesOk

`func (o *DfsTrash) GetTotalFilesOk() (*int64, bool)`

GetTotalFilesOk returns a tuple with the TotalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFiles

`func (o *DfsTrash) SetTotalFiles(v int64)`

SetTotalFiles sets TotalFiles field to given value.

### HasTotalFiles

`func (o *DfsTrash) HasTotalFiles() bool`

HasTotalFiles returns a boolean if a field has been set.

### GetTotalKbyte

`func (o *DfsTrash) GetTotalKbyte() int64`

GetTotalKbyte returns the TotalKbyte field if non-nil, zero value otherwise.

### GetTotalKbyteOk

`func (o *DfsTrash) GetTotalKbyteOk() (*int64, bool)`

GetTotalKbyteOk returns a tuple with the TotalKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalKbyte

`func (o *DfsTrash) SetTotalKbyte(v int64)`

SetTotalKbyte sets TotalKbyte field to given value.

### HasTotalKbyte

`func (o *DfsTrash) HasTotalKbyte() bool`

HasTotalKbyte returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsTrash) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsTrash) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsTrash) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsTrash) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


