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

// checks if the VolumeReplicationSetReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeReplicationSetReq{}

// VolumeReplicationSetReq struct for VolumeReplicationSetReq
type VolumeReplicationSetReq struct {
	BlockVolume *VolumeReplicationSetReqVolume `json:"block_volume,omitempty"`
}

// NewVolumeReplicationSetReq instantiates a new VolumeReplicationSetReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeReplicationSetReq() *VolumeReplicationSetReq {
	this := VolumeReplicationSetReq{}
	return &this
}

// NewVolumeReplicationSetReqWithDefaults instantiates a new VolumeReplicationSetReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeReplicationSetReqWithDefaults() *VolumeReplicationSetReq {
	this := VolumeReplicationSetReq{}
	return &this
}

// GetBlockVolume returns the BlockVolume field value if set, zero value otherwise.
func (o *VolumeReplicationSetReq) GetBlockVolume() VolumeReplicationSetReqVolume {
	if o == nil || IsNil(o.BlockVolume) {
		var ret VolumeReplicationSetReqVolume
		return ret
	}
	return *o.BlockVolume
}

// GetBlockVolumeOk returns a tuple with the BlockVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeReplicationSetReq) GetBlockVolumeOk() (*VolumeReplicationSetReqVolume, bool) {
	if o == nil || IsNil(o.BlockVolume) {
		return nil, false
	}
	return o.BlockVolume, true
}

// HasBlockVolume returns a boolean if a field has been set.
func (o *VolumeReplicationSetReq) HasBlockVolume() bool {
	if o != nil && !IsNil(o.BlockVolume) {
		return true
	}

	return false
}

// SetBlockVolume gets a reference to the given VolumeReplicationSetReqVolume and assigns it to the BlockVolume field.
func (o *VolumeReplicationSetReq) SetBlockVolume(v VolumeReplicationSetReqVolume) {
	o.BlockVolume = &v
}

func (o VolumeReplicationSetReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeReplicationSetReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockVolume) {
		toSerialize["block_volume"] = o.BlockVolume
	}
	return toSerialize, nil
}

type NullableVolumeReplicationSetReq struct {
	value *VolumeReplicationSetReq
	isSet bool
}

func (v NullableVolumeReplicationSetReq) Get() *VolumeReplicationSetReq {
	return v.value
}

func (v *NullableVolumeReplicationSetReq) Set(val *VolumeReplicationSetReq) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeReplicationSetReq) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeReplicationSetReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeReplicationSetReq(val *VolumeReplicationSetReq) *NullableVolumeReplicationSetReq {
	return &NullableVolumeReplicationSetReq{value: val, isSet: true}
}

func (v NullableVolumeReplicationSetReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeReplicationSetReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


