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

// checks if the DfsRootfsesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsRootfsesResp{}

// DfsRootfsesResp struct for DfsRootfsesResp
type DfsRootfsesResp struct {
	// dfs rootfs records
	DfsRootfses []DfsRootfsRecord `json:"dfs_rootfses"`
}

type _DfsRootfsesResp DfsRootfsesResp

// NewDfsRootfsesResp instantiates a new DfsRootfsesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsRootfsesResp(dfsRootfses []DfsRootfsRecord) *DfsRootfsesResp {
	this := DfsRootfsesResp{}
	this.DfsRootfses = dfsRootfses
	return &this
}

// NewDfsRootfsesRespWithDefaults instantiates a new DfsRootfsesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsRootfsesRespWithDefaults() *DfsRootfsesResp {
	this := DfsRootfsesResp{}
	return &this
}

// GetDfsRootfses returns the DfsRootfses field value
func (o *DfsRootfsesResp) GetDfsRootfses() []DfsRootfsRecord {
	if o == nil {
		var ret []DfsRootfsRecord
		return ret
	}

	return o.DfsRootfses
}

// GetDfsRootfsesOk returns a tuple with the DfsRootfses field value
// and a boolean to check if the value has been set.
func (o *DfsRootfsesResp) GetDfsRootfsesOk() ([]DfsRootfsRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.DfsRootfses, true
}

// SetDfsRootfses sets field value
func (o *DfsRootfsesResp) SetDfsRootfses(v []DfsRootfsRecord) {
	o.DfsRootfses = v
}

func (o DfsRootfsesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsRootfsesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_rootfses"] = o.DfsRootfses
	return toSerialize, nil
}

func (o *DfsRootfsesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_rootfses",
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

	varDfsRootfsesResp := _DfsRootfsesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsRootfsesResp)

	if err != nil {
		return err
	}

	*o = DfsRootfsesResp(varDfsRootfsesResp)

	return err
}

type NullableDfsRootfsesResp struct {
	value *DfsRootfsesResp
	isSet bool
}

func (v NullableDfsRootfsesResp) Get() *DfsRootfsesResp {
	return v.value
}

func (v *NullableDfsRootfsesResp) Set(val *DfsRootfsesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsRootfsesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsRootfsesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsRootfsesResp(val *DfsRootfsesResp) *NullableDfsRootfsesResp {
	return &NullableDfsRootfsesResp{value: val, isSet: true}
}

func (v NullableDfsRootfsesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsRootfsesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


