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

// checks if the VolumeGroupResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeGroupResp{}

// VolumeGroupResp struct for VolumeGroupResp
type VolumeGroupResp struct {
	BlockVolumeGroup *VolumeGroup `json:"block_volume_group,omitempty"`
}

// NewVolumeGroupResp instantiates a new VolumeGroupResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeGroupResp() *VolumeGroupResp {
	this := VolumeGroupResp{}
	return &this
}

// NewVolumeGroupRespWithDefaults instantiates a new VolumeGroupResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeGroupRespWithDefaults() *VolumeGroupResp {
	this := VolumeGroupResp{}
	return &this
}

// GetBlockVolumeGroup returns the BlockVolumeGroup field value if set, zero value otherwise.
func (o *VolumeGroupResp) GetBlockVolumeGroup() VolumeGroup {
	if o == nil || IsNil(o.BlockVolumeGroup) {
		var ret VolumeGroup
		return ret
	}
	return *o.BlockVolumeGroup
}

// GetBlockVolumeGroupOk returns a tuple with the BlockVolumeGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeGroupResp) GetBlockVolumeGroupOk() (*VolumeGroup, bool) {
	if o == nil || IsNil(o.BlockVolumeGroup) {
		return nil, false
	}
	return o.BlockVolumeGroup, true
}

// HasBlockVolumeGroup returns a boolean if a field has been set.
func (o *VolumeGroupResp) HasBlockVolumeGroup() bool {
	if o != nil && !IsNil(o.BlockVolumeGroup) {
		return true
	}

	return false
}

// SetBlockVolumeGroup gets a reference to the given VolumeGroup and assigns it to the BlockVolumeGroup field.
func (o *VolumeGroupResp) SetBlockVolumeGroup(v VolumeGroup) {
	o.BlockVolumeGroup = &v
}

func (o VolumeGroupResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeGroupResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockVolumeGroup) {
		toSerialize["block_volume_group"] = o.BlockVolumeGroup
	}
	return toSerialize, nil
}

type NullableVolumeGroupResp struct {
	value *VolumeGroupResp
	isSet bool
}

func (v NullableVolumeGroupResp) Get() *VolumeGroupResp {
	return v.value
}

func (v *NullableVolumeGroupResp) Set(val *VolumeGroupResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeGroupResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeGroupResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeGroupResp(val *VolumeGroupResp) *NullableVolumeGroupResp {
	return &NullableVolumeGroupResp{value: val, isSet: true}
}

func (v NullableVolumeGroupResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeGroupResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


