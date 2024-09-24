# DfsAuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LogPath** | Pointer to **string** |  | [optional] 
**Rootfs** | Pointer to [**DfsRootfsNestview**](DfsRootfsNestview.md) |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsAuditLog

`func NewDfsAuditLog() *DfsAuditLog`

NewDfsAuditLog instantiates a new DfsAuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsAuditLogWithDefaults

`func NewDfsAuditLogWithDefaults() *DfsAuditLog`

NewDfsAuditLogWithDefaults instantiates a new DfsAuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *DfsAuditLog) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *DfsAuditLog) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *DfsAuditLog) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *DfsAuditLog) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetCluster

`func (o *DfsAuditLog) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsAuditLog) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsAuditLog) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsAuditLog) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsAuditLog) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsAuditLog) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsAuditLog) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsAuditLog) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *DfsAuditLog) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsAuditLog) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsAuditLog) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsAuditLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogPath

`func (o *DfsAuditLog) GetLogPath() string`

GetLogPath returns the LogPath field if non-nil, zero value otherwise.

### GetLogPathOk

`func (o *DfsAuditLog) GetLogPathOk() (*string, bool)`

GetLogPathOk returns a tuple with the LogPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogPath

`func (o *DfsAuditLog) SetLogPath(v string)`

SetLogPath sets LogPath field to given value.

### HasLogPath

`func (o *DfsAuditLog) HasLogPath() bool`

HasLogPath returns a boolean if a field has been set.

### GetRootfs

`func (o *DfsAuditLog) GetRootfs() DfsRootfsNestview`

GetRootfs returns the Rootfs field if non-nil, zero value otherwise.

### GetRootfsOk

`func (o *DfsAuditLog) GetRootfsOk() (*DfsRootfsNestview, bool)`

GetRootfsOk returns a tuple with the Rootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfs

`func (o *DfsAuditLog) SetRootfs(v DfsRootfsNestview)`

SetRootfs sets Rootfs field to given value.

### HasRootfs

`func (o *DfsAuditLog) HasRootfs() bool`

HasRootfs returns a boolean if a field has been set.

### GetSize

`func (o *DfsAuditLog) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DfsAuditLog) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DfsAuditLog) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DfsAuditLog) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *DfsAuditLog) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsAuditLog) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsAuditLog) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsAuditLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsAuditLog) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsAuditLog) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsAuditLog) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsAuditLog) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


