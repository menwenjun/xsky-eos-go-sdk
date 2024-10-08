/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TargetResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetResp{}

// TargetResp struct for TargetResp
type TargetResp struct {
	Target Target `json:"target"`
}

type _TargetResp TargetResp

// NewTargetResp instantiates a new TargetResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetResp(target Target) *TargetResp {
	this := TargetResp{}
	this.Target = target
	return &this
}

// NewTargetRespWithDefaults instantiates a new TargetResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetRespWithDefaults() *TargetResp {
	this := TargetResp{}
	return &this
}

// GetTarget returns the Target field value
func (o *TargetResp) GetTarget() Target {
	if o == nil {
		var ret Target
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *TargetResp) GetTargetOk() (*Target, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *TargetResp) SetTarget(v Target) {
	o.Target = v
}

func (o TargetResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["target"] = o.Target
	return toSerialize, nil
}

func (o *TargetResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"target",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTargetResp := _TargetResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTargetResp)

	if err != nil {
		return err
	}

	*o = TargetResp(varTargetResp)

	return err
}

type NullableTargetResp struct {
	value *TargetResp
	isSet bool
}

func (v NullableTargetResp) Get() *TargetResp {
	return v.value
}

func (v *NullableTargetResp) Set(val *TargetResp) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetResp) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetResp(val *TargetResp) *NullableTargetResp {
	return &NullableTargetResp{value: val, isSet: true}
}

func (v NullableTargetResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


