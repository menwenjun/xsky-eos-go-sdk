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

// checks if the DfsHdfsesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsHdfsesResp{}

// DfsHdfsesResp struct for DfsHdfsesResp
type DfsHdfsesResp struct {
	// dfs hdfs list
	DfsHdfses []DfsHdfsRecord `json:"dfs_hdfses"`
}

type _DfsHdfsesResp DfsHdfsesResp

// NewDfsHdfsesResp instantiates a new DfsHdfsesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsHdfsesResp(dfsHdfses []DfsHdfsRecord) *DfsHdfsesResp {
	this := DfsHdfsesResp{}
	this.DfsHdfses = dfsHdfses
	return &this
}

// NewDfsHdfsesRespWithDefaults instantiates a new DfsHdfsesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsHdfsesRespWithDefaults() *DfsHdfsesResp {
	this := DfsHdfsesResp{}
	return &this
}

// GetDfsHdfses returns the DfsHdfses field value
func (o *DfsHdfsesResp) GetDfsHdfses() []DfsHdfsRecord {
	if o == nil {
		var ret []DfsHdfsRecord
		return ret
	}

	return o.DfsHdfses
}

// GetDfsHdfsesOk returns a tuple with the DfsHdfses field value
// and a boolean to check if the value has been set.
func (o *DfsHdfsesResp) GetDfsHdfsesOk() ([]DfsHdfsRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.DfsHdfses, true
}

// SetDfsHdfses sets field value
func (o *DfsHdfsesResp) SetDfsHdfses(v []DfsHdfsRecord) {
	o.DfsHdfses = v
}

func (o DfsHdfsesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsHdfsesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_hdfses"] = o.DfsHdfses
	return toSerialize, nil
}

func (o *DfsHdfsesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_hdfses",
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

	varDfsHdfsesResp := _DfsHdfsesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsHdfsesResp)

	if err != nil {
		return err
	}

	*o = DfsHdfsesResp(varDfsHdfsesResp)

	return err
}

type NullableDfsHdfsesResp struct {
	value *DfsHdfsesResp
	isSet bool
}

func (v NullableDfsHdfsesResp) Get() *DfsHdfsesResp {
	return v.value
}

func (v *NullableDfsHdfsesResp) Set(val *DfsHdfsesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsHdfsesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsHdfsesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsHdfsesResp(val *DfsHdfsesResp) *NullableDfsHdfsesResp {
	return &NullableDfsHdfsesResp{value: val, isSet: true}
}

func (v NullableDfsHdfsesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsHdfsesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


