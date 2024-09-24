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

// checks if the SnapshotUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotUpdateReq{}

// SnapshotUpdateReq struct for SnapshotUpdateReq
type SnapshotUpdateReq struct {
	BlockSnapshot SnapshotUpdateReqSnapshot `json:"block_snapshot"`
}

type _SnapshotUpdateReq SnapshotUpdateReq

// NewSnapshotUpdateReq instantiates a new SnapshotUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotUpdateReq(blockSnapshot SnapshotUpdateReqSnapshot) *SnapshotUpdateReq {
	this := SnapshotUpdateReq{}
	this.BlockSnapshot = blockSnapshot
	return &this
}

// NewSnapshotUpdateReqWithDefaults instantiates a new SnapshotUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotUpdateReqWithDefaults() *SnapshotUpdateReq {
	this := SnapshotUpdateReq{}
	return &this
}

// GetBlockSnapshot returns the BlockSnapshot field value
func (o *SnapshotUpdateReq) GetBlockSnapshot() SnapshotUpdateReqSnapshot {
	if o == nil {
		var ret SnapshotUpdateReqSnapshot
		return ret
	}

	return o.BlockSnapshot
}

// GetBlockSnapshotOk returns a tuple with the BlockSnapshot field value
// and a boolean to check if the value has been set.
func (o *SnapshotUpdateReq) GetBlockSnapshotOk() (*SnapshotUpdateReqSnapshot, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockSnapshot, true
}

// SetBlockSnapshot sets field value
func (o *SnapshotUpdateReq) SetBlockSnapshot(v SnapshotUpdateReqSnapshot) {
	o.BlockSnapshot = v
}

func (o SnapshotUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block_snapshot"] = o.BlockSnapshot
	return toSerialize, nil
}

func (o *SnapshotUpdateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"block_snapshot",
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

	varSnapshotUpdateReq := _SnapshotUpdateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSnapshotUpdateReq)

	if err != nil {
		return err
	}

	*o = SnapshotUpdateReq(varSnapshotUpdateReq)

	return err
}

type NullableSnapshotUpdateReq struct {
	value *SnapshotUpdateReq
	isSet bool
}

func (v NullableSnapshotUpdateReq) Get() *SnapshotUpdateReq {
	return v.value
}

func (v *NullableSnapshotUpdateReq) Set(val *SnapshotUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotUpdateReq(val *SnapshotUpdateReq) *NullableSnapshotUpdateReq {
	return &NullableSnapshotUpdateReq{value: val, isSet: true}
}

func (v NullableSnapshotUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


