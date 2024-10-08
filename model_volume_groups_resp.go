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

// checks if the VolumeGroupsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeGroupsResp{}

// VolumeGroupsResp struct for VolumeGroupsResp
type VolumeGroupsResp struct {
	BlockVolumeGroups []VolumeGroup `json:"block_volume_groups,omitempty"`
}

// NewVolumeGroupsResp instantiates a new VolumeGroupsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeGroupsResp() *VolumeGroupsResp {
	this := VolumeGroupsResp{}
	return &this
}

// NewVolumeGroupsRespWithDefaults instantiates a new VolumeGroupsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeGroupsRespWithDefaults() *VolumeGroupsResp {
	this := VolumeGroupsResp{}
	return &this
}

// GetBlockVolumeGroups returns the BlockVolumeGroups field value if set, zero value otherwise.
func (o *VolumeGroupsResp) GetBlockVolumeGroups() []VolumeGroup {
	if o == nil || IsNil(o.BlockVolumeGroups) {
		var ret []VolumeGroup
		return ret
	}
	return o.BlockVolumeGroups
}

// GetBlockVolumeGroupsOk returns a tuple with the BlockVolumeGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeGroupsResp) GetBlockVolumeGroupsOk() ([]VolumeGroup, bool) {
	if o == nil || IsNil(o.BlockVolumeGroups) {
		return nil, false
	}
	return o.BlockVolumeGroups, true
}

// HasBlockVolumeGroups returns a boolean if a field has been set.
func (o *VolumeGroupsResp) HasBlockVolumeGroups() bool {
	if o != nil && !IsNil(o.BlockVolumeGroups) {
		return true
	}

	return false
}

// SetBlockVolumeGroups gets a reference to the given []VolumeGroup and assigns it to the BlockVolumeGroups field.
func (o *VolumeGroupsResp) SetBlockVolumeGroups(v []VolumeGroup) {
	o.BlockVolumeGroups = v
}

func (o VolumeGroupsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeGroupsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockVolumeGroups) {
		toSerialize["block_volume_groups"] = o.BlockVolumeGroups
	}
	return toSerialize, nil
}

type NullableVolumeGroupsResp struct {
	value *VolumeGroupsResp
	isSet bool
}

func (v NullableVolumeGroupsResp) Get() *VolumeGroupsResp {
	return v.value
}

func (v *NullableVolumeGroupsResp) Set(val *VolumeGroupsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeGroupsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeGroupsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeGroupsResp(val *VolumeGroupsResp) *NullableVolumeGroupsResp {
	return &NullableVolumeGroupsResp{value: val, isSet: true}
}

func (v NullableVolumeGroupsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeGroupsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


