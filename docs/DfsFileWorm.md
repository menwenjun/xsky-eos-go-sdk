# DfsFileWorm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoLockTime** | Pointer to **time.Time** | automatic lock time of file | [optional] 
**DfsWorm** | Pointer to [**DfsWorm**](DfsWorm.md) |  | [optional] 
**ExpireTime** | Pointer to **time.Time** | worm expire time of file | [optional] 
**State** | Pointer to **string** | worm state of file | [optional] 

## Methods

### NewDfsFileWorm

`func NewDfsFileWorm() *DfsFileWorm`

NewDfsFileWorm instantiates a new DfsFileWorm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsFileWormWithDefaults

`func NewDfsFileWormWithDefaults() *DfsFileWorm`

NewDfsFileWormWithDefaults instantiates a new DfsFileWorm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoLockTime

`func (o *DfsFileWorm) GetAutoLockTime() time.Time`

GetAutoLockTime returns the AutoLockTime field if non-nil, zero value otherwise.

### GetAutoLockTimeOk

`func (o *DfsFileWorm) GetAutoLockTimeOk() (*time.Time, bool)`

GetAutoLockTimeOk returns a tuple with the AutoLockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLockTime

`func (o *DfsFileWorm) SetAutoLockTime(v time.Time)`

SetAutoLockTime sets AutoLockTime field to given value.

### HasAutoLockTime

`func (o *DfsFileWorm) HasAutoLockTime() bool`

HasAutoLockTime returns a boolean if a field has been set.

### GetDfsWorm

`func (o *DfsFileWorm) GetDfsWorm() DfsWorm`

GetDfsWorm returns the DfsWorm field if non-nil, zero value otherwise.

### GetDfsWormOk

`func (o *DfsFileWorm) GetDfsWormOk() (*DfsWorm, bool)`

GetDfsWormOk returns a tuple with the DfsWorm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsWorm

`func (o *DfsFileWorm) SetDfsWorm(v DfsWorm)`

SetDfsWorm sets DfsWorm field to given value.

### HasDfsWorm

`func (o *DfsFileWorm) HasDfsWorm() bool`

HasDfsWorm returns a boolean if a field has been set.

### GetExpireTime

`func (o *DfsFileWorm) GetExpireTime() time.Time`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *DfsFileWorm) GetExpireTimeOk() (*time.Time, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *DfsFileWorm) SetExpireTime(v time.Time)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *DfsFileWorm) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetState

`func (o *DfsFileWorm) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DfsFileWorm) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DfsFileWorm) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DfsFileWorm) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


