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

// checks if the DisksResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisksResp{}

// DisksResp struct for DisksResp
type DisksResp struct {
	// disks
	Disks []DiskRecord `json:"disks"`
}

type _DisksResp DisksResp

// NewDisksResp instantiates a new DisksResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisksResp(disks []DiskRecord) *DisksResp {
	this := DisksResp{}
	this.Disks = disks
	return &this
}

// NewDisksRespWithDefaults instantiates a new DisksResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisksRespWithDefaults() *DisksResp {
	this := DisksResp{}
	return &this
}

// GetDisks returns the Disks field value
func (o *DisksResp) GetDisks() []DiskRecord {
	if o == nil {
		var ret []DiskRecord
		return ret
	}

	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value
// and a boolean to check if the value has been set.
func (o *DisksResp) GetDisksOk() ([]DiskRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disks, true
}

// SetDisks sets field value
func (o *DisksResp) SetDisks(v []DiskRecord) {
	o.Disks = v
}

func (o DisksResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisksResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["disks"] = o.Disks
	return toSerialize, nil
}

func (o *DisksResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"disks",
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

	varDisksResp := _DisksResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDisksResp)

	if err != nil {
		return err
	}

	*o = DisksResp(varDisksResp)

	return err
}

type NullableDisksResp struct {
	value *DisksResp
	isSet bool
}

func (v NullableDisksResp) Get() *DisksResp {
	return v.value
}

func (v *NullableDisksResp) Set(val *DisksResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDisksResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDisksResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisksResp(val *DisksResp) *NullableDisksResp {
	return &NullableDisksResp{value: val, isSet: true}
}

func (v NullableDisksResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisksResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


