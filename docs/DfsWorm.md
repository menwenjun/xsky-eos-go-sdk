# DfsWorm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoLockPeriod** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DefaultProtectPeriod** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MaxProtectPeriod** | Pointer to **string** |  | [optional] 
**MinProtectPeriod** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Path** | Pointer to [**DfsPathNestview**](DfsPathNestview.md) |  | [optional] 
**Rootfs** | Pointer to [**DfsRootfsNestview**](DfsRootfsNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsWorm

`func NewDfsWorm() *DfsWorm`

NewDfsWorm instantiates a new DfsWorm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsWormWithDefaults

`func NewDfsWormWithDefaults() *DfsWorm`

NewDfsWormWithDefaults instantiates a new DfsWorm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoLockPeriod

`func (o *DfsWorm) GetAutoLockPeriod() string`

GetAutoLockPeriod returns the AutoLockPeriod field if non-nil, zero value otherwise.

### GetAutoLockPeriodOk

`func (o *DfsWorm) GetAutoLockPeriodOk() (*string, bool)`

GetAutoLockPeriodOk returns a tuple with the AutoLockPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLockPeriod

`func (o *DfsWorm) SetAutoLockPeriod(v string)`

SetAutoLockPeriod sets AutoLockPeriod field to given value.

### HasAutoLockPeriod

`func (o *DfsWorm) HasAutoLockPeriod() bool`

HasAutoLockPeriod returns a boolean if a field has been set.

### GetCluster

`func (o *DfsWorm) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsWorm) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsWorm) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsWorm) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsWorm) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsWorm) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsWorm) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsWorm) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDefaultProtectPeriod

`func (o *DfsWorm) GetDefaultProtectPeriod() string`

GetDefaultProtectPeriod returns the DefaultProtectPeriod field if non-nil, zero value otherwise.

### GetDefaultProtectPeriodOk

`func (o *DfsWorm) GetDefaultProtectPeriodOk() (*string, bool)`

GetDefaultProtectPeriodOk returns a tuple with the DefaultProtectPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProtectPeriod

`func (o *DfsWorm) SetDefaultProtectPeriod(v string)`

SetDefaultProtectPeriod sets DefaultProtectPeriod field to given value.

### HasDefaultProtectPeriod

`func (o *DfsWorm) HasDefaultProtectPeriod() bool`

HasDefaultProtectPeriod returns a boolean if a field has been set.

### GetId

`func (o *DfsWorm) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsWorm) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsWorm) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsWorm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxProtectPeriod

`func (o *DfsWorm) GetMaxProtectPeriod() string`

GetMaxProtectPeriod returns the MaxProtectPeriod field if non-nil, zero value otherwise.

### GetMaxProtectPeriodOk

`func (o *DfsWorm) GetMaxProtectPeriodOk() (*string, bool)`

GetMaxProtectPeriodOk returns a tuple with the MaxProtectPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProtectPeriod

`func (o *DfsWorm) SetMaxProtectPeriod(v string)`

SetMaxProtectPeriod sets MaxProtectPeriod field to given value.

### HasMaxProtectPeriod

`func (o *DfsWorm) HasMaxProtectPeriod() bool`

HasMaxProtectPeriod returns a boolean if a field has been set.

### GetMinProtectPeriod

`func (o *DfsWorm) GetMinProtectPeriod() string`

GetMinProtectPeriod returns the MinProtectPeriod field if non-nil, zero value otherwise.

### GetMinProtectPeriodOk

`func (o *DfsWorm) GetMinProtectPeriodOk() (*string, bool)`

GetMinProtectPeriodOk returns a tuple with the MinProtectPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinProtectPeriod

`func (o *DfsWorm) SetMinProtectPeriod(v string)`

SetMinProtectPeriod sets MinProtectPeriod field to given value.

### HasMinProtectPeriod

`func (o *DfsWorm) HasMinProtectPeriod() bool`

HasMinProtectPeriod returns a boolean if a field has been set.

### GetMode

`func (o *DfsWorm) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DfsWorm) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DfsWorm) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DfsWorm) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPath

`func (o *DfsWorm) GetPath() DfsPathNestview`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsWorm) GetPathOk() (*DfsPathNestview, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsWorm) SetPath(v DfsPathNestview)`

SetPath sets Path field to given value.

### HasPath

`func (o *DfsWorm) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRootfs

`func (o *DfsWorm) GetRootfs() DfsRootfsNestview`

GetRootfs returns the Rootfs field if non-nil, zero value otherwise.

### GetRootfsOk

`func (o *DfsWorm) GetRootfsOk() (*DfsRootfsNestview, bool)`

GetRootfsOk returns a tuple with the Rootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfs

`func (o *DfsWorm) SetRootfs(v DfsRootfsNestview)`

SetRootfs sets Rootfs field to given value.

### HasRootfs

`func (o *DfsWorm) HasRootfs() bool`

HasRootfs returns a boolean if a field has been set.

### GetStatus

`func (o *DfsWorm) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsWorm) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsWorm) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsWorm) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsWorm) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsWorm) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsWorm) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsWorm) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


