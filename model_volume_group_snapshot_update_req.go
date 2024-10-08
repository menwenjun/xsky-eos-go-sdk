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

// checks if the VolumeGroupSnapshotUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeGroupSnapshotUpdateReq{}

// VolumeGroupSnapshotUpdateReq struct for VolumeGroupSnapshotUpdateReq
type VolumeGroupSnapshotUpdateReq struct {
	BlockVolumeGroupSnapshot VolumeGroupSnapshotUpdateReqVolumeGroupSnapshot `json:"block_volume_group_snapshot"`
}

type _VolumeGroupSnapshotUpdateReq VolumeGroupSnapshotUpdateReq

// NewVolumeGroupSnapshotUpdateReq instantiates a new VolumeGroupSnapshotUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeGroupSnapshotUpdateReq(blockVolumeGroupSnapshot VolumeGroupSnapshotUpdateReqVolumeGroupSnapshot) *VolumeGroupSnapshotUpdateReq {
	this := VolumeGroupSnapshotUpdateReq{}
	this.BlockVolumeGroupSnapshot = blockVolumeGroupSnapshot
	return &this
}

// NewVolumeGroupSnapshotUpdateReqWithDefaults instantiates a new VolumeGroupSnapshotUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeGroupSnapshotUpdateReqWithDefaults() *VolumeGroupSnapshotUpdateReq {
	this := VolumeGroupSnapshotUpdateReq{}
	return &this
}

// GetBlockVolumeGroupSnapshot returns the BlockVolumeGroupSnapshot field value
func (o *VolumeGroupSnapshotUpdateReq) GetBlockVolumeGroupSnapshot() VolumeGroupSnapshotUpdateReqVolumeGroupSnapshot {
	if o == nil {
		var ret VolumeGroupSnapshotUpdateReqVolumeGroupSnapshot
		return ret
	}

	return o.BlockVolumeGroupSnapshot
}

// GetBlockVolumeGroupSnapshotOk returns a tuple with the BlockVolumeGroupSnapshot field value
// and a boolean to check if the value has been set.
func (o *VolumeGroupSnapshotUpdateReq) GetBlockVolumeGroupSnapshotOk() (*VolumeGroupSnapshotUpdateReqVolumeGroupSnapshot, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockVolumeGroupSnapshot, true
}

// SetBlockVolumeGroupSnapshot sets field value
func (o *VolumeGroupSnapshotUpdateReq) SetBlockVolumeGroupSnapshot(v VolumeGroupSnapshotUpdateReqVolumeGroupSnapshot) {
	o.BlockVolumeGroupSnapshot = v
}

func (o VolumeGroupSnapshotUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeGroupSnapshotUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block_volume_group_snapshot"] = o.BlockVolumeGroupSnapshot
	return toSerialize, nil
}

func (o *VolumeGroupSnapshotUpdateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"block_volume_group_snapshot",
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

	varVolumeGroupSnapshotUpdateReq := _VolumeGroupSnapshotUpdateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVolumeGroupSnapshotUpdateReq)

	if err != nil {
		return err
	}

	*o = VolumeGroupSnapshotUpdateReq(varVolumeGroupSnapshotUpdateReq)

	return err
}

type NullableVolumeGroupSnapshotUpdateReq struct {
	value *VolumeGroupSnapshotUpdateReq
	isSet bool
}

func (v NullableVolumeGroupSnapshotUpdateReq) Get() *VolumeGroupSnapshotUpdateReq {
	return v.value
}

func (v *NullableVolumeGroupSnapshotUpdateReq) Set(val *VolumeGroupSnapshotUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeGroupSnapshotUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeGroupSnapshotUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeGroupSnapshotUpdateReq(val *VolumeGroupSnapshotUpdateReq) *NullableVolumeGroupSnapshotUpdateReq {
	return &NullableVolumeGroupSnapshotUpdateReq{value: val, isSet: true}
}

func (v NullableVolumeGroupSnapshotUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeGroupSnapshotUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


