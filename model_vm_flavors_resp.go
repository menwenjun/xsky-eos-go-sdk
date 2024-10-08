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

// checks if the VMFlavorsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMFlavorsResp{}

// VMFlavorsResp struct for VMFlavorsResp
type VMFlavorsResp struct {
	VmFlavors []VMFlavor `json:"vm_flavors,omitempty"`
}

// NewVMFlavorsResp instantiates a new VMFlavorsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMFlavorsResp() *VMFlavorsResp {
	this := VMFlavorsResp{}
	return &this
}

// NewVMFlavorsRespWithDefaults instantiates a new VMFlavorsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMFlavorsRespWithDefaults() *VMFlavorsResp {
	this := VMFlavorsResp{}
	return &this
}

// GetVmFlavors returns the VmFlavors field value if set, zero value otherwise.
func (o *VMFlavorsResp) GetVmFlavors() []VMFlavor {
	if o == nil || IsNil(o.VmFlavors) {
		var ret []VMFlavor
		return ret
	}
	return o.VmFlavors
}

// GetVmFlavorsOk returns a tuple with the VmFlavors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMFlavorsResp) GetVmFlavorsOk() ([]VMFlavor, bool) {
	if o == nil || IsNil(o.VmFlavors) {
		return nil, false
	}
	return o.VmFlavors, true
}

// HasVmFlavors returns a boolean if a field has been set.
func (o *VMFlavorsResp) HasVmFlavors() bool {
	if o != nil && !IsNil(o.VmFlavors) {
		return true
	}

	return false
}

// SetVmFlavors gets a reference to the given []VMFlavor and assigns it to the VmFlavors field.
func (o *VMFlavorsResp) SetVmFlavors(v []VMFlavor) {
	o.VmFlavors = v
}

func (o VMFlavorsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMFlavorsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VmFlavors) {
		toSerialize["vm_flavors"] = o.VmFlavors
	}
	return toSerialize, nil
}

type NullableVMFlavorsResp struct {
	value *VMFlavorsResp
	isSet bool
}

func (v NullableVMFlavorsResp) Get() *VMFlavorsResp {
	return v.value
}

func (v *NullableVMFlavorsResp) Set(val *VMFlavorsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVMFlavorsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVMFlavorsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMFlavorsResp(val *VMFlavorsResp) *NullableVMFlavorsResp {
	return &NullableVMFlavorsResp{value: val, isSet: true}
}

func (v NullableVMFlavorsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMFlavorsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


