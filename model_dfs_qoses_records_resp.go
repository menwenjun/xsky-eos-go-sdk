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

// checks if the DfsQosesRecordsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsQosesRecordsResp{}

// DfsQosesRecordsResp struct for DfsQosesRecordsResp
type DfsQosesRecordsResp struct {
	// dfs qoses records
	DfsQoses []DfsQosRecord `json:"dfs_qoses"`
}

type _DfsQosesRecordsResp DfsQosesRecordsResp

// NewDfsQosesRecordsResp instantiates a new DfsQosesRecordsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsQosesRecordsResp(dfsQoses []DfsQosRecord) *DfsQosesRecordsResp {
	this := DfsQosesRecordsResp{}
	this.DfsQoses = dfsQoses
	return &this
}

// NewDfsQosesRecordsRespWithDefaults instantiates a new DfsQosesRecordsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsQosesRecordsRespWithDefaults() *DfsQosesRecordsResp {
	this := DfsQosesRecordsResp{}
	return &this
}

// GetDfsQoses returns the DfsQoses field value
func (o *DfsQosesRecordsResp) GetDfsQoses() []DfsQosRecord {
	if o == nil {
		var ret []DfsQosRecord
		return ret
	}

	return o.DfsQoses
}

// GetDfsQosesOk returns a tuple with the DfsQoses field value
// and a boolean to check if the value has been set.
func (o *DfsQosesRecordsResp) GetDfsQosesOk() ([]DfsQosRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.DfsQoses, true
}

// SetDfsQoses sets field value
func (o *DfsQosesRecordsResp) SetDfsQoses(v []DfsQosRecord) {
	o.DfsQoses = v
}

func (o DfsQosesRecordsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsQosesRecordsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_qoses"] = o.DfsQoses
	return toSerialize, nil
}

func (o *DfsQosesRecordsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_qoses",
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

	varDfsQosesRecordsResp := _DfsQosesRecordsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsQosesRecordsResp)

	if err != nil {
		return err
	}

	*o = DfsQosesRecordsResp(varDfsQosesRecordsResp)

	return err
}

type NullableDfsQosesRecordsResp struct {
	value *DfsQosesRecordsResp
	isSet bool
}

func (v NullableDfsQosesRecordsResp) Get() *DfsQosesRecordsResp {
	return v.value
}

func (v *NullableDfsQosesRecordsResp) Set(val *DfsQosesRecordsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsQosesRecordsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsQosesRecordsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsQosesRecordsResp(val *DfsQosesRecordsResp) *NullableDfsQosesRecordsResp {
	return &NullableDfsQosesRecordsResp{value: val, isSet: true}
}

func (v NullableDfsQosesRecordsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsQosesRecordsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


