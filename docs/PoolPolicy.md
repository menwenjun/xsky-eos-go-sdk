# PoolPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | is active pool | [optional] 
**FromDefault** | Pointer to **bool** | judge the source of the add pool | [optional] 
**PoolId** | Pointer to **int64** | pool id | [optional] 
**Threshold** | Pointer to **int64** | threshold for auto active pool | [optional] 

## Methods

### NewPoolPolicy

`func NewPoolPolicy() *PoolPolicy`

NewPoolPolicy instantiates a new PoolPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolPolicyWithDefaults

`func NewPoolPolicyWithDefaults() *PoolPolicy`

NewPoolPolicyWithDefaults instantiates a new PoolPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *PoolPolicy) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PoolPolicy) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PoolPolicy) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PoolPolicy) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFromDefault

`func (o *PoolPolicy) GetFromDefault() bool`

GetFromDefault returns the FromDefault field if non-nil, zero value otherwise.

### GetFromDefaultOk

`func (o *PoolPolicy) GetFromDefaultOk() (*bool, bool)`

GetFromDefaultOk returns a tuple with the FromDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDefault

`func (o *PoolPolicy) SetFromDefault(v bool)`

SetFromDefault sets FromDefault field to given value.

### HasFromDefault

`func (o *PoolPolicy) HasFromDefault() bool`

HasFromDefault returns a boolean if a field has been set.

### GetPoolId

`func (o *PoolPolicy) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *PoolPolicy) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *PoolPolicy) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *PoolPolicy) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetThreshold

`func (o *PoolPolicy) GetThreshold() int64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *PoolPolicy) GetThresholdOk() (*int64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *PoolPolicy) SetThreshold(v int64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *PoolPolicy) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


