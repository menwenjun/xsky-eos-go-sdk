# LifecycleRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Expiration** | Pointer to [**LifecycleExpiration**](LifecycleExpiration.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Transitions** | Pointer to [**[]LifecycleTransition**](LifecycleTransition.md) |  | [optional] 

## Methods

### NewLifecycleRule

`func NewLifecycleRule() *LifecycleRule`

NewLifecycleRule instantiates a new LifecycleRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleRuleWithDefaults

`func NewLifecycleRuleWithDefaults() *LifecycleRule`

NewLifecycleRuleWithDefaults instantiates a new LifecycleRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *LifecycleRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LifecycleRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LifecycleRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *LifecycleRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExpiration

`func (o *LifecycleRule) GetExpiration() LifecycleExpiration`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *LifecycleRule) GetExpirationOk() (*LifecycleExpiration, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *LifecycleRule) SetExpiration(v LifecycleExpiration)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *LifecycleRule) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetName

`func (o *LifecycleRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LifecycleRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LifecycleRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LifecycleRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefix

`func (o *LifecycleRule) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *LifecycleRule) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *LifecycleRule) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *LifecycleRule) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetTransitions

`func (o *LifecycleRule) GetTransitions() []LifecycleTransition`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *LifecycleRule) GetTransitionsOk() (*[]LifecycleTransition, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *LifecycleRule) SetTransitions(v []LifecycleTransition)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *LifecycleRule) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


