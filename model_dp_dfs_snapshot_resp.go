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

// checks if the DpDfsSnapshotResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpDfsSnapshotResp{}

// DpDfsSnapshotResp struct for DpDfsSnapshotResp
type DpDfsSnapshotResp struct {
	DpDfsSnapshot DpDfsSnapshot `json:"dp_dfs_snapshot"`
}

type _DpDfsSnapshotResp DpDfsSnapshotResp

// NewDpDfsSnapshotResp instantiates a new DpDfsSnapshotResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpDfsSnapshotResp(dpDfsSnapshot DpDfsSnapshot) *DpDfsSnapshotResp {
	this := DpDfsSnapshotResp{}
	this.DpDfsSnapshot = dpDfsSnapshot
	return &this
}

// NewDpDfsSnapshotRespWithDefaults instantiates a new DpDfsSnapshotResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpDfsSnapshotRespWithDefaults() *DpDfsSnapshotResp {
	this := DpDfsSnapshotResp{}
	return &this
}

// GetDpDfsSnapshot returns the DpDfsSnapshot field value
func (o *DpDfsSnapshotResp) GetDpDfsSnapshot() DpDfsSnapshot {
	if o == nil {
		var ret DpDfsSnapshot
		return ret
	}

	return o.DpDfsSnapshot
}

// GetDpDfsSnapshotOk returns a tuple with the DpDfsSnapshot field value
// and a boolean to check if the value has been set.
func (o *DpDfsSnapshotResp) GetDpDfsSnapshotOk() (*DpDfsSnapshot, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DpDfsSnapshot, true
}

// SetDpDfsSnapshot sets field value
func (o *DpDfsSnapshotResp) SetDpDfsSnapshot(v DpDfsSnapshot) {
	o.DpDfsSnapshot = v
}

func (o DpDfsSnapshotResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpDfsSnapshotResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dp_dfs_snapshot"] = o.DpDfsSnapshot
	return toSerialize, nil
}

func (o *DpDfsSnapshotResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dp_dfs_snapshot",
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

	varDpDfsSnapshotResp := _DpDfsSnapshotResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDpDfsSnapshotResp)

	if err != nil {
		return err
	}

	*o = DpDfsSnapshotResp(varDpDfsSnapshotResp)

	return err
}

type NullableDpDfsSnapshotResp struct {
	value *DpDfsSnapshotResp
	isSet bool
}

func (v NullableDpDfsSnapshotResp) Get() *DpDfsSnapshotResp {
	return v.value
}

func (v *NullableDpDfsSnapshotResp) Set(val *DpDfsSnapshotResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDpDfsSnapshotResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDpDfsSnapshotResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpDfsSnapshotResp(val *DpDfsSnapshotResp) *NullableDpDfsSnapshotResp {
	return &NullableDpDfsSnapshotResp{value: val, isSet: true}
}

func (v NullableDpDfsSnapshotResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpDfsSnapshotResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


