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

// checks if the VolumeGroupUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeGroupUpdateReq{}

// VolumeGroupUpdateReq struct for VolumeGroupUpdateReq
type VolumeGroupUpdateReq struct {
	BlockVolumeGroup VolumeGroupUpdateReqVolumeGroup `json:"block_volume_group"`
}

type _VolumeGroupUpdateReq VolumeGroupUpdateReq

// NewVolumeGroupUpdateReq instantiates a new VolumeGroupUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeGroupUpdateReq(blockVolumeGroup VolumeGroupUpdateReqVolumeGroup) *VolumeGroupUpdateReq {
	this := VolumeGroupUpdateReq{}
	this.BlockVolumeGroup = blockVolumeGroup
	return &this
}

// NewVolumeGroupUpdateReqWithDefaults instantiates a new VolumeGroupUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeGroupUpdateReqWithDefaults() *VolumeGroupUpdateReq {
	this := VolumeGroupUpdateReq{}
	return &this
}

// GetBlockVolumeGroup returns the BlockVolumeGroup field value
func (o *VolumeGroupUpdateReq) GetBlockVolumeGroup() VolumeGroupUpdateReqVolumeGroup {
	if o == nil {
		var ret VolumeGroupUpdateReqVolumeGroup
		return ret
	}

	return o.BlockVolumeGroup
}

// GetBlockVolumeGroupOk returns a tuple with the BlockVolumeGroup field value
// and a boolean to check if the value has been set.
func (o *VolumeGroupUpdateReq) GetBlockVolumeGroupOk() (*VolumeGroupUpdateReqVolumeGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockVolumeGroup, true
}

// SetBlockVolumeGroup sets field value
func (o *VolumeGroupUpdateReq) SetBlockVolumeGroup(v VolumeGroupUpdateReqVolumeGroup) {
	o.BlockVolumeGroup = v
}

func (o VolumeGroupUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeGroupUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block_volume_group"] = o.BlockVolumeGroup
	return toSerialize, nil
}

func (o *VolumeGroupUpdateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"block_volume_group",
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

	varVolumeGroupUpdateReq := _VolumeGroupUpdateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVolumeGroupUpdateReq)

	if err != nil {
		return err
	}

	*o = VolumeGroupUpdateReq(varVolumeGroupUpdateReq)

	return err
}

type NullableVolumeGroupUpdateReq struct {
	value *VolumeGroupUpdateReq
	isSet bool
}

func (v NullableVolumeGroupUpdateReq) Get() *VolumeGroupUpdateReq {
	return v.value
}

func (v *NullableVolumeGroupUpdateReq) Set(val *VolumeGroupUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeGroupUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeGroupUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeGroupUpdateReq(val *VolumeGroupUpdateReq) *NullableVolumeGroupUpdateReq {
	return &NullableVolumeGroupUpdateReq{value: val, isSet: true}
}

func (v NullableVolumeGroupUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeGroupUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


