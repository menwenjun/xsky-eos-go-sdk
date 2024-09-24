# DfsWormCreateReqWorm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoLockPeriod** | **string** | automatic locking period | 
**DefaultProtectPeriod** | **string** | default protect period | 
**MaxProtectPeriod** | **string** | maximum protect period | 
**MinProtectPeriod** | **string** | minimum protect period | 
**Mode** | Pointer to **string** | worm mode, value is enterprise or compliance | [optional] 
**Path** | **string** | worm root path | 
**RootfsId** | **int64** | id of dfs rootfs | 

## Methods

### NewDfsWormCreateReqWorm

`func NewDfsWormCreateReqWorm(autoLockPeriod string, defaultProtectPeriod string, maxProtectPeriod string, minProtectPeriod string, path string, rootfsId int64, ) *DfsWormCreateReqWorm`

NewDfsWormCreateReqWorm instantiates a new DfsWormCreateReqWorm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsWormCreateReqWormWithDefaults

`func NewDfsWormCreateReqWormWithDefaults() *DfsWormCreateReqWorm`

NewDfsWormCreateReqWormWithDefaults instantiates a new DfsWormCreateReqWorm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoLockPeriod

`func (o *DfsWormCreateReqWorm) GetAutoLockPeriod() string`

GetAutoLockPeriod returns the AutoLockPeriod field if non-nil, zero value otherwise.

### GetAutoLockPeriodOk

`func (o *DfsWormCreateReqWorm) GetAutoLockPeriodOk() (*string, bool)`

GetAutoLockPeriodOk returns a tuple with the AutoLockPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLockPeriod

`func (o *DfsWormCreateReqWorm) SetAutoLockPeriod(v string)`

SetAutoLockPeriod sets AutoLockPeriod field to given value.


### GetDefaultProtectPeriod

`func (o *DfsWormCreateReqWorm) GetDefaultProtectPeriod() string`

GetDefaultProtectPeriod returns the DefaultProtectPeriod field if non-nil, zero value otherwise.

### GetDefaultProtectPeriodOk

`func (o *DfsWormCreateReqWorm) GetDefaultProtectPeriodOk() (*string, bool)`

GetDefaultProtectPeriodOk returns a tuple with the DefaultProtectPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProtectPeriod

`func (o *DfsWormCreateReqWorm) SetDefaultProtectPeriod(v string)`

SetDefaultProtectPeriod sets DefaultProtectPeriod field to given value.


### GetMaxProtectPeriod

`func (o *DfsWormCreateReqWorm) GetMaxProtectPeriod() string`

GetMaxProtectPeriod returns the MaxProtectPeriod field if non-nil, zero value otherwise.

### GetMaxProtectPeriodOk

`func (o *DfsWormCreateReqWorm) GetMaxProtectPeriodOk() (*string, bool)`

GetMaxProtectPeriodOk returns a tuple with the MaxProtectPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProtectPeriod

`func (o *DfsWormCreateReqWorm) SetMaxProtectPeriod(v string)`

SetMaxProtectPeriod sets MaxProtectPeriod field to given value.


### GetMinProtectPeriod

`func (o *DfsWormCreateReqWorm) GetMinProtectPeriod() string`

GetMinProtectPeriod returns the MinProtectPeriod field if non-nil, zero value otherwise.

### GetMinProtectPeriodOk

`func (o *DfsWormCreateReqWorm) GetMinProtectPeriodOk() (*string, bool)`

GetMinProtectPeriodOk returns a tuple with the MinProtectPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinProtectPeriod

`func (o *DfsWormCreateReqWorm) SetMinProtectPeriod(v string)`

SetMinProtectPeriod sets MinProtectPeriod field to given value.


### GetMode

`func (o *DfsWormCreateReqWorm) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DfsWormCreateReqWorm) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DfsWormCreateReqWorm) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DfsWormCreateReqWorm) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPath

`func (o *DfsWormCreateReqWorm) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsWormCreateReqWorm) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsWormCreateReqWorm) SetPath(v string)`

SetPath sets Path field to given value.


### GetRootfsId

`func (o *DfsWormCreateReqWorm) GetRootfsId() int64`

GetRootfsId returns the RootfsId field if non-nil, zero value otherwise.

### GetRootfsIdOk

`func (o *DfsWormCreateReqWorm) GetRootfsIdOk() (*int64, bool)`

GetRootfsIdOk returns a tuple with the RootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfsId

`func (o *DfsWormCreateReqWorm) SetRootfsId(v int64)`

SetRootfsId sets RootfsId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


