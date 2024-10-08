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

// checks if the ConsistentSnapshotsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsistentSnapshotsResp{}

// ConsistentSnapshotsResp struct for ConsistentSnapshotsResp
type ConsistentSnapshotsResp struct {
	// snapshots
	BlockConsistentSnapshots []ConsistentSnapshot `json:"block_consistent_snapshots"`
}

type _ConsistentSnapshotsResp ConsistentSnapshotsResp

// NewConsistentSnapshotsResp instantiates a new ConsistentSnapshotsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsistentSnapshotsResp(blockConsistentSnapshots []ConsistentSnapshot) *ConsistentSnapshotsResp {
	this := ConsistentSnapshotsResp{}
	this.BlockConsistentSnapshots = blockConsistentSnapshots
	return &this
}

// NewConsistentSnapshotsRespWithDefaults instantiates a new ConsistentSnapshotsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsistentSnapshotsRespWithDefaults() *ConsistentSnapshotsResp {
	this := ConsistentSnapshotsResp{}
	return &this
}

// GetBlockConsistentSnapshots returns the BlockConsistentSnapshots field value
func (o *ConsistentSnapshotsResp) GetBlockConsistentSnapshots() []ConsistentSnapshot {
	if o == nil {
		var ret []ConsistentSnapshot
		return ret
	}

	return o.BlockConsistentSnapshots
}

// GetBlockConsistentSnapshotsOk returns a tuple with the BlockConsistentSnapshots field value
// and a boolean to check if the value has been set.
func (o *ConsistentSnapshotsResp) GetBlockConsistentSnapshotsOk() ([]ConsistentSnapshot, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockConsistentSnapshots, true
}

// SetBlockConsistentSnapshots sets field value
func (o *ConsistentSnapshotsResp) SetBlockConsistentSnapshots(v []ConsistentSnapshot) {
	o.BlockConsistentSnapshots = v
}

func (o ConsistentSnapshotsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsistentSnapshotsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block_consistent_snapshots"] = o.BlockConsistentSnapshots
	return toSerialize, nil
}

func (o *ConsistentSnapshotsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"block_consistent_snapshots",
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

	varConsistentSnapshotsResp := _ConsistentSnapshotsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConsistentSnapshotsResp)

	if err != nil {
		return err
	}

	*o = ConsistentSnapshotsResp(varConsistentSnapshotsResp)

	return err
}

type NullableConsistentSnapshotsResp struct {
	value *ConsistentSnapshotsResp
	isSet bool
}

func (v NullableConsistentSnapshotsResp) Get() *ConsistentSnapshotsResp {
	return v.value
}

func (v *NullableConsistentSnapshotsResp) Set(val *ConsistentSnapshotsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableConsistentSnapshotsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableConsistentSnapshotsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsistentSnapshotsResp(val *ConsistentSnapshotsResp) *NullableConsistentSnapshotsResp {
	return &NullableConsistentSnapshotsResp{value: val, isSet: true}
}

func (v NullableConsistentSnapshotsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsistentSnapshotsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


