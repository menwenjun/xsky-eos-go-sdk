# GcPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | Pointer to **int64** |  | [optional] 
**Enable** | **bool** |  | 
**EndTime** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 

## Methods

### NewGcPolicy

`func NewGcPolicy(enable bool, ) *GcPolicy`

NewGcPolicy instantiates a new GcPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcPolicyWithDefaults

`func NewGcPolicyWithDefaults() *GcPolicy`

NewGcPolicyWithDefaults instantiates a new GcPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidth

`func (o *GcPolicy) GetBandwidth() int64`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *GcPolicy) GetBandwidthOk() (*int64, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *GcPolicy) SetBandwidth(v int64)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *GcPolicy) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetEnable

`func (o *GcPolicy) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GcPolicy) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GcPolicy) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetEndTime

`func (o *GcPolicy) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GcPolicy) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GcPolicy) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GcPolicy) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetLevel

`func (o *GcPolicy) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GcPolicy) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GcPolicy) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *GcPolicy) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetStartTime

`func (o *GcPolicy) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GcPolicy) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GcPolicy) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GcPolicy) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


