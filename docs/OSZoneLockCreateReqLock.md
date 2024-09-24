# OSZoneLockCreateReqLock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | key of lock, for example name of os bucket | 
**ResourceType** | **string** | resource type of lock, including os_bucket | 
**Value** | **string** | value of lock, for example action of os bucket | 

## Methods

### NewOSZoneLockCreateReqLock

`func NewOSZoneLockCreateReqLock(key string, resourceType string, value string, ) *OSZoneLockCreateReqLock`

NewOSZoneLockCreateReqLock instantiates a new OSZoneLockCreateReqLock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSZoneLockCreateReqLockWithDefaults

`func NewOSZoneLockCreateReqLockWithDefaults() *OSZoneLockCreateReqLock`

NewOSZoneLockCreateReqLockWithDefaults instantiates a new OSZoneLockCreateReqLock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *OSZoneLockCreateReqLock) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OSZoneLockCreateReqLock) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OSZoneLockCreateReqLock) SetKey(v string)`

SetKey sets Key field to given value.


### GetResourceType

`func (o *OSZoneLockCreateReqLock) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *OSZoneLockCreateReqLock) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *OSZoneLockCreateReqLock) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetValue

`func (o *OSZoneLockCreateReqLock) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OSZoneLockCreateReqLock) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OSZoneLockCreateReqLock) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


