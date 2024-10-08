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

// checks if the VolumeCrcSetReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeCrcSetReq{}

// VolumeCrcSetReq struct for VolumeCrcSetReq
type VolumeCrcSetReq struct {
	BlockVolume *VolumeCrcSetReqVolume `json:"block_volume,omitempty"`
}

// NewVolumeCrcSetReq instantiates a new VolumeCrcSetReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeCrcSetReq() *VolumeCrcSetReq {
	this := VolumeCrcSetReq{}
	return &this
}

// NewVolumeCrcSetReqWithDefaults instantiates a new VolumeCrcSetReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeCrcSetReqWithDefaults() *VolumeCrcSetReq {
	this := VolumeCrcSetReq{}
	return &this
}

// GetBlockVolume returns the BlockVolume field value if set, zero value otherwise.
func (o *VolumeCrcSetReq) GetBlockVolume() VolumeCrcSetReqVolume {
	if o == nil || IsNil(o.BlockVolume) {
		var ret VolumeCrcSetReqVolume
		return ret
	}
	return *o.BlockVolume
}

// GetBlockVolumeOk returns a tuple with the BlockVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeCrcSetReq) GetBlockVolumeOk() (*VolumeCrcSetReqVolume, bool) {
	if o == nil || IsNil(o.BlockVolume) {
		return nil, false
	}
	return o.BlockVolume, true
}

// HasBlockVolume returns a boolean if a field has been set.
func (o *VolumeCrcSetReq) HasBlockVolume() bool {
	if o != nil && !IsNil(o.BlockVolume) {
		return true
	}

	return false
}

// SetBlockVolume gets a reference to the given VolumeCrcSetReqVolume and assigns it to the BlockVolume field.
func (o *VolumeCrcSetReq) SetBlockVolume(v VolumeCrcSetReqVolume) {
	o.BlockVolume = &v
}

func (o VolumeCrcSetReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeCrcSetReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockVolume) {
		toSerialize["block_volume"] = o.BlockVolume
	}
	return toSerialize, nil
}

type NullableVolumeCrcSetReq struct {
	value *VolumeCrcSetReq
	isSet bool
}

func (v NullableVolumeCrcSetReq) Get() *VolumeCrcSetReq {
	return v.value
}

func (v *NullableVolumeCrcSetReq) Set(val *VolumeCrcSetReq) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeCrcSetReq) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeCrcSetReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeCrcSetReq(val *VolumeCrcSetReq) *NullableVolumeCrcSetReq {
	return &NullableVolumeCrcSetReq{value: val, isSet: true}
}

func (v NullableVolumeCrcSetReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeCrcSetReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


