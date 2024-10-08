/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ContactStrategy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactStrategy{}

// ContactStrategy struct for ContactStrategy
type ContactStrategy struct {
	CancelStrategies []AlertStrategy `json:"cancel_strategies,omitempty"`
	DeleteStrategies []AlertStrategy `json:"delete_strategies,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewContactStrategy instantiates a new ContactStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactStrategy() *ContactStrategy {
	this := ContactStrategy{}
	return &this
}

// NewContactStrategyWithDefaults instantiates a new ContactStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactStrategyWithDefaults() *ContactStrategy {
	this := ContactStrategy{}
	return &this
}

// GetCancelStrategies returns the CancelStrategies field value if set, zero value otherwise.
func (o *ContactStrategy) GetCancelStrategies() []AlertStrategy {
	if o == nil || IsNil(o.CancelStrategies) {
		var ret []AlertStrategy
		return ret
	}
	return o.CancelStrategies
}

// GetCancelStrategiesOk returns a tuple with the CancelStrategies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactStrategy) GetCancelStrategiesOk() ([]AlertStrategy, bool) {
	if o == nil || IsNil(o.CancelStrategies) {
		return nil, false
	}
	return o.CancelStrategies, true
}

// HasCancelStrategies returns a boolean if a field has been set.
func (o *ContactStrategy) HasCancelStrategies() bool {
	if o != nil && !IsNil(o.CancelStrategies) {
		return true
	}

	return false
}

// SetCancelStrategies gets a reference to the given []AlertStrategy and assigns it to the CancelStrategies field.
func (o *ContactStrategy) SetCancelStrategies(v []AlertStrategy) {
	o.CancelStrategies = v
}

// GetDeleteStrategies returns the DeleteStrategies field value if set, zero value otherwise.
func (o *ContactStrategy) GetDeleteStrategies() []AlertStrategy {
	if o == nil || IsNil(o.DeleteStrategies) {
		var ret []AlertStrategy
		return ret
	}
	return o.DeleteStrategies
}

// GetDeleteStrategiesOk returns a tuple with the DeleteStrategies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactStrategy) GetDeleteStrategiesOk() ([]AlertStrategy, bool) {
	if o == nil || IsNil(o.DeleteStrategies) {
		return nil, false
	}
	return o.DeleteStrategies, true
}

// HasDeleteStrategies returns a boolean if a field has been set.
func (o *ContactStrategy) HasDeleteStrategies() bool {
	if o != nil && !IsNil(o.DeleteStrategies) {
		return true
	}

	return false
}

// SetDeleteStrategies gets a reference to the given []AlertStrategy and assigns it to the DeleteStrategies field.
func (o *ContactStrategy) SetDeleteStrategies(v []AlertStrategy) {
	o.DeleteStrategies = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContactStrategy) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactStrategy) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContactStrategy) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ContactStrategy) SetType(v string) {
	o.Type = &v
}

func (o ContactStrategy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactStrategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CancelStrategies) {
		toSerialize["cancel_strategies"] = o.CancelStrategies
	}
	if !IsNil(o.DeleteStrategies) {
		toSerialize["delete_strategies"] = o.DeleteStrategies
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableContactStrategy struct {
	value *ContactStrategy
	isSet bool
}

func (v NullableContactStrategy) Get() *ContactStrategy {
	return v.value
}

func (v *NullableContactStrategy) Set(val *ContactStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableContactStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableContactStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactStrategy(val *ContactStrategy) *NullableContactStrategy {
	return &NullableContactStrategy{value: val, isSet: true}
}

func (v NullableContactStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


