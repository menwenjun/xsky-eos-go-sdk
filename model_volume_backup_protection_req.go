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

// checks if the VolumeBackupProtectionReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeBackupProtectionReq{}

// VolumeBackupProtectionReq struct for VolumeBackupProtectionReq
type VolumeBackupProtectionReq struct {
	BlockVolume *VolumeBackupProtectionReqVolume `json:"block_volume,omitempty"`
}

// NewVolumeBackupProtectionReq instantiates a new VolumeBackupProtectionReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeBackupProtectionReq() *VolumeBackupProtectionReq {
	this := VolumeBackupProtectionReq{}
	return &this
}

// NewVolumeBackupProtectionReqWithDefaults instantiates a new VolumeBackupProtectionReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeBackupProtectionReqWithDefaults() *VolumeBackupProtectionReq {
	this := VolumeBackupProtectionReq{}
	return &this
}

// GetBlockVolume returns the BlockVolume field value if set, zero value otherwise.
func (o *VolumeBackupProtectionReq) GetBlockVolume() VolumeBackupProtectionReqVolume {
	if o == nil || IsNil(o.BlockVolume) {
		var ret VolumeBackupProtectionReqVolume
		return ret
	}
	return *o.BlockVolume
}

// GetBlockVolumeOk returns a tuple with the BlockVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeBackupProtectionReq) GetBlockVolumeOk() (*VolumeBackupProtectionReqVolume, bool) {
	if o == nil || IsNil(o.BlockVolume) {
		return nil, false
	}
	return o.BlockVolume, true
}

// HasBlockVolume returns a boolean if a field has been set.
func (o *VolumeBackupProtectionReq) HasBlockVolume() bool {
	if o != nil && !IsNil(o.BlockVolume) {
		return true
	}

	return false
}

// SetBlockVolume gets a reference to the given VolumeBackupProtectionReqVolume and assigns it to the BlockVolume field.
func (o *VolumeBackupProtectionReq) SetBlockVolume(v VolumeBackupProtectionReqVolume) {
	o.BlockVolume = &v
}

func (o VolumeBackupProtectionReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeBackupProtectionReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockVolume) {
		toSerialize["block_volume"] = o.BlockVolume
	}
	return toSerialize, nil
}

type NullableVolumeBackupProtectionReq struct {
	value *VolumeBackupProtectionReq
	isSet bool
}

func (v NullableVolumeBackupProtectionReq) Get() *VolumeBackupProtectionReq {
	return v.value
}

func (v *NullableVolumeBackupProtectionReq) Set(val *VolumeBackupProtectionReq) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeBackupProtectionReq) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeBackupProtectionReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeBackupProtectionReq(val *VolumeBackupProtectionReq) *NullableVolumeBackupProtectionReq {
	return &NullableVolumeBackupProtectionReq{value: val, isSet: true}
}

func (v NullableVolumeBackupProtectionReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeBackupProtectionReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


