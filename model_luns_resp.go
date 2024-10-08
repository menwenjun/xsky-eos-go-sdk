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

// checks if the LunsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LunsResp{}

// LunsResp struct for LunsResp
type LunsResp struct {
	// luns
	Luns []Lun `json:"luns"`
}

type _LunsResp LunsResp

// NewLunsResp instantiates a new LunsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLunsResp(luns []Lun) *LunsResp {
	this := LunsResp{}
	this.Luns = luns
	return &this
}

// NewLunsRespWithDefaults instantiates a new LunsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLunsRespWithDefaults() *LunsResp {
	this := LunsResp{}
	return &this
}

// GetLuns returns the Luns field value
func (o *LunsResp) GetLuns() []Lun {
	if o == nil {
		var ret []Lun
		return ret
	}

	return o.Luns
}

// GetLunsOk returns a tuple with the Luns field value
// and a boolean to check if the value has been set.
func (o *LunsResp) GetLunsOk() ([]Lun, bool) {
	if o == nil {
		return nil, false
	}
	return o.Luns, true
}

// SetLuns sets field value
func (o *LunsResp) SetLuns(v []Lun) {
	o.Luns = v
}

func (o LunsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LunsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["luns"] = o.Luns
	return toSerialize, nil
}

func (o *LunsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"luns",
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

	varLunsResp := _LunsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLunsResp)

	if err != nil {
		return err
	}

	*o = LunsResp(varLunsResp)

	return err
}

type NullableLunsResp struct {
	value *LunsResp
	isSet bool
}

func (v NullableLunsResp) Get() *LunsResp {
	return v.value
}

func (v *NullableLunsResp) Set(val *LunsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableLunsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableLunsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLunsResp(val *LunsResp) *NullableLunsResp {
	return &NullableLunsResp{value: val, isSet: true}
}

func (v NullableLunsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLunsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


