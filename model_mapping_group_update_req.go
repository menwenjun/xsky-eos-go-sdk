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

// checks if the MappingGroupUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MappingGroupUpdateReq{}

// MappingGroupUpdateReq struct for MappingGroupUpdateReq
type MappingGroupUpdateReq struct {
	MappingGroup MappingGroupUpdateReqMappingGroup `json:"mapping_group"`
}

type _MappingGroupUpdateReq MappingGroupUpdateReq

// NewMappingGroupUpdateReq instantiates a new MappingGroupUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMappingGroupUpdateReq(mappingGroup MappingGroupUpdateReqMappingGroup) *MappingGroupUpdateReq {
	this := MappingGroupUpdateReq{}
	this.MappingGroup = mappingGroup
	return &this
}

// NewMappingGroupUpdateReqWithDefaults instantiates a new MappingGroupUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMappingGroupUpdateReqWithDefaults() *MappingGroupUpdateReq {
	this := MappingGroupUpdateReq{}
	return &this
}

// GetMappingGroup returns the MappingGroup field value
func (o *MappingGroupUpdateReq) GetMappingGroup() MappingGroupUpdateReqMappingGroup {
	if o == nil {
		var ret MappingGroupUpdateReqMappingGroup
		return ret
	}

	return o.MappingGroup
}

// GetMappingGroupOk returns a tuple with the MappingGroup field value
// and a boolean to check if the value has been set.
func (o *MappingGroupUpdateReq) GetMappingGroupOk() (*MappingGroupUpdateReqMappingGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MappingGroup, true
}

// SetMappingGroup sets field value
func (o *MappingGroupUpdateReq) SetMappingGroup(v MappingGroupUpdateReqMappingGroup) {
	o.MappingGroup = v
}

func (o MappingGroupUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MappingGroupUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mapping_group"] = o.MappingGroup
	return toSerialize, nil
}

func (o *MappingGroupUpdateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mapping_group",
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

	varMappingGroupUpdateReq := _MappingGroupUpdateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMappingGroupUpdateReq)

	if err != nil {
		return err
	}

	*o = MappingGroupUpdateReq(varMappingGroupUpdateReq)

	return err
}

type NullableMappingGroupUpdateReq struct {
	value *MappingGroupUpdateReq
	isSet bool
}

func (v NullableMappingGroupUpdateReq) Get() *MappingGroupUpdateReq {
	return v.value
}

func (v *NullableMappingGroupUpdateReq) Set(val *MappingGroupUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableMappingGroupUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableMappingGroupUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMappingGroupUpdateReq(val *MappingGroupUpdateReq) *NullableMappingGroupUpdateReq {
	return &NullableMappingGroupUpdateReq{value: val, isSet: true}
}

func (v NullableMappingGroupUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMappingGroupUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


