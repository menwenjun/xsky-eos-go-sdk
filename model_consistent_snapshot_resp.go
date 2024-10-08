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

// checks if the ConsistentSnapshotResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsistentSnapshotResp{}

// ConsistentSnapshotResp struct for ConsistentSnapshotResp
type ConsistentSnapshotResp struct {
	BlockConsistentSnapshot ConsistentSnapshot `json:"block_consistent_snapshot"`
}

type _ConsistentSnapshotResp ConsistentSnapshotResp

// NewConsistentSnapshotResp instantiates a new ConsistentSnapshotResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsistentSnapshotResp(blockConsistentSnapshot ConsistentSnapshot) *ConsistentSnapshotResp {
	this := ConsistentSnapshotResp{}
	this.BlockConsistentSnapshot = blockConsistentSnapshot
	return &this
}

// NewConsistentSnapshotRespWithDefaults instantiates a new ConsistentSnapshotResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsistentSnapshotRespWithDefaults() *ConsistentSnapshotResp {
	this := ConsistentSnapshotResp{}
	return &this
}

// GetBlockConsistentSnapshot returns the BlockConsistentSnapshot field value
func (o *ConsistentSnapshotResp) GetBlockConsistentSnapshot() ConsistentSnapshot {
	if o == nil {
		var ret ConsistentSnapshot
		return ret
	}

	return o.BlockConsistentSnapshot
}

// GetBlockConsistentSnapshotOk returns a tuple with the BlockConsistentSnapshot field value
// and a boolean to check if the value has been set.
func (o *ConsistentSnapshotResp) GetBlockConsistentSnapshotOk() (*ConsistentSnapshot, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockConsistentSnapshot, true
}

// SetBlockConsistentSnapshot sets field value
func (o *ConsistentSnapshotResp) SetBlockConsistentSnapshot(v ConsistentSnapshot) {
	o.BlockConsistentSnapshot = v
}

func (o ConsistentSnapshotResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsistentSnapshotResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block_consistent_snapshot"] = o.BlockConsistentSnapshot
	return toSerialize, nil
}

func (o *ConsistentSnapshotResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"block_consistent_snapshot",
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

	varConsistentSnapshotResp := _ConsistentSnapshotResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConsistentSnapshotResp)

	if err != nil {
		return err
	}

	*o = ConsistentSnapshotResp(varConsistentSnapshotResp)

	return err
}

type NullableConsistentSnapshotResp struct {
	value *ConsistentSnapshotResp
	isSet bool
}

func (v NullableConsistentSnapshotResp) Get() *ConsistentSnapshotResp {
	return v.value
}

func (v *NullableConsistentSnapshotResp) Set(val *ConsistentSnapshotResp) {
	v.value = val
	v.isSet = true
}

func (v NullableConsistentSnapshotResp) IsSet() bool {
	return v.isSet
}

func (v *NullableConsistentSnapshotResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsistentSnapshotResp(val *ConsistentSnapshotResp) *NullableConsistentSnapshotResp {
	return &NullableConsistentSnapshotResp{value: val, isSet: true}
}

func (v NullableConsistentSnapshotResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsistentSnapshotResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


