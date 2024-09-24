# LifecycleExpiration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **int64** |  | [optional] 
**StartAt** | Pointer to **string** |  | [optional] 
**StopAt** | Pointer to **string** |  | [optional] 
**DeleteType** | Pointer to **string** |  | [optional] 

## Methods

### NewLifecycleExpiration

`func NewLifecycleExpiration() *LifecycleExpiration`

NewLifecycleExpiration instantiates a new LifecycleExpiration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleExpirationWithDefaults

`func NewLifecycleExpirationWithDefaults() *LifecycleExpiration`

NewLifecycleExpirationWithDefaults instantiates a new LifecycleExpiration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *LifecycleExpiration) GetDays() int64`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *LifecycleExpiration) GetDaysOk() (*int64, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *LifecycleExpiration) SetDays(v int64)`

SetDays sets Days field to given value.

### HasDays

`func (o *LifecycleExpiration) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetStartAt

`func (o *LifecycleExpiration) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *LifecycleExpiration) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *LifecycleExpiration) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *LifecycleExpiration) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetStopAt

`func (o *LifecycleExpiration) GetStopAt() string`

GetStopAt returns the StopAt field if non-nil, zero value otherwise.

### GetStopAtOk

`func (o *LifecycleExpiration) GetStopAtOk() (*string, bool)`

GetStopAtOk returns a tuple with the StopAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopAt

`func (o *LifecycleExpiration) SetStopAt(v string)`

SetStopAt sets StopAt field to given value.

### HasStopAt

`func (o *LifecycleExpiration) HasStopAt() bool`

HasStopAt returns a boolean if a field has been set.

### GetDeleteType

`func (o *LifecycleExpiration) GetDeleteType() string`

GetDeleteType returns the DeleteType field if non-nil, zero value otherwise.

### GetDeleteTypeOk

`func (o *LifecycleExpiration) GetDeleteTypeOk() (*string, bool)`

GetDeleteTypeOk returns a tuple with the DeleteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteType

`func (o *LifecycleExpiration) SetDeleteType(v string)`

SetDeleteType sets DeleteType field to given value.

### HasDeleteType

`func (o *LifecycleExpiration) HasDeleteType() bool`

HasDeleteType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


