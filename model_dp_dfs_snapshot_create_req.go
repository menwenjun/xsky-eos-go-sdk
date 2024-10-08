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

// checks if the DpDfsSnapshotCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpDfsSnapshotCreateReq{}

// DpDfsSnapshotCreateReq struct for DpDfsSnapshotCreateReq
type DpDfsSnapshotCreateReq struct {
	DpDfsSnapshot DpDfsSnapshotCreateReqDpDfsSnapshot `json:"dp_dfs_snapshot"`
}

type _DpDfsSnapshotCreateReq DpDfsSnapshotCreateReq

// NewDpDfsSnapshotCreateReq instantiates a new DpDfsSnapshotCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpDfsSnapshotCreateReq(dpDfsSnapshot DpDfsSnapshotCreateReqDpDfsSnapshot) *DpDfsSnapshotCreateReq {
	this := DpDfsSnapshotCreateReq{}
	this.DpDfsSnapshot = dpDfsSnapshot
	return &this
}

// NewDpDfsSnapshotCreateReqWithDefaults instantiates a new DpDfsSnapshotCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpDfsSnapshotCreateReqWithDefaults() *DpDfsSnapshotCreateReq {
	this := DpDfsSnapshotCreateReq{}
	return &this
}

// GetDpDfsSnapshot returns the DpDfsSnapshot field value
func (o *DpDfsSnapshotCreateReq) GetDpDfsSnapshot() DpDfsSnapshotCreateReqDpDfsSnapshot {
	if o == nil {
		var ret DpDfsSnapshotCreateReqDpDfsSnapshot
		return ret
	}

	return o.DpDfsSnapshot
}

// GetDpDfsSnapshotOk returns a tuple with the DpDfsSnapshot field value
// and a boolean to check if the value has been set.
func (o *DpDfsSnapshotCreateReq) GetDpDfsSnapshotOk() (*DpDfsSnapshotCreateReqDpDfsSnapshot, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DpDfsSnapshot, true
}

// SetDpDfsSnapshot sets field value
func (o *DpDfsSnapshotCreateReq) SetDpDfsSnapshot(v DpDfsSnapshotCreateReqDpDfsSnapshot) {
	o.DpDfsSnapshot = v
}

func (o DpDfsSnapshotCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpDfsSnapshotCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dp_dfs_snapshot"] = o.DpDfsSnapshot
	return toSerialize, nil
}

func (o *DpDfsSnapshotCreateReq) UnmarshalJSON(data []byte) (err error) {
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

	varDpDfsSnapshotCreateReq := _DpDfsSnapshotCreateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDpDfsSnapshotCreateReq)

	if err != nil {
		return err
	}

	*o = DpDfsSnapshotCreateReq(varDpDfsSnapshotCreateReq)

	return err
}

type NullableDpDfsSnapshotCreateReq struct {
	value *DpDfsSnapshotCreateReq
	isSet bool
}

func (v NullableDpDfsSnapshotCreateReq) Get() *DpDfsSnapshotCreateReq {
	return v.value
}

func (v *NullableDpDfsSnapshotCreateReq) Set(val *DpDfsSnapshotCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDpDfsSnapshotCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDpDfsSnapshotCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpDfsSnapshotCreateReq(val *DpDfsSnapshotCreateReq) *NullableDpDfsSnapshotCreateReq {
	return &NullableDpDfsSnapshotCreateReq{value: val, isSet: true}
}

func (v NullableDpDfsSnapshotCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpDfsSnapshotCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


