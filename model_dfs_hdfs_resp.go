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

// checks if the DfsHdfsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsHdfsResp{}

// DfsHdfsResp struct for DfsHdfsResp
type DfsHdfsResp struct {
	DfsHdfs DfsHdfsRecord `json:"dfs_hdfs"`
}

type _DfsHdfsResp DfsHdfsResp

// NewDfsHdfsResp instantiates a new DfsHdfsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsHdfsResp(dfsHdfs DfsHdfsRecord) *DfsHdfsResp {
	this := DfsHdfsResp{}
	this.DfsHdfs = dfsHdfs
	return &this
}

// NewDfsHdfsRespWithDefaults instantiates a new DfsHdfsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsHdfsRespWithDefaults() *DfsHdfsResp {
	this := DfsHdfsResp{}
	return &this
}

// GetDfsHdfs returns the DfsHdfs field value
func (o *DfsHdfsResp) GetDfsHdfs() DfsHdfsRecord {
	if o == nil {
		var ret DfsHdfsRecord
		return ret
	}

	return o.DfsHdfs
}

// GetDfsHdfsOk returns a tuple with the DfsHdfs field value
// and a boolean to check if the value has been set.
func (o *DfsHdfsResp) GetDfsHdfsOk() (*DfsHdfsRecord, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsHdfs, true
}

// SetDfsHdfs sets field value
func (o *DfsHdfsResp) SetDfsHdfs(v DfsHdfsRecord) {
	o.DfsHdfs = v
}

func (o DfsHdfsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsHdfsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_hdfs"] = o.DfsHdfs
	return toSerialize, nil
}

func (o *DfsHdfsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_hdfs",
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

	varDfsHdfsResp := _DfsHdfsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsHdfsResp)

	if err != nil {
		return err
	}

	*o = DfsHdfsResp(varDfsHdfsResp)

	return err
}

type NullableDfsHdfsResp struct {
	value *DfsHdfsResp
	isSet bool
}

func (v NullableDfsHdfsResp) Get() *DfsHdfsResp {
	return v.value
}

func (v *NullableDfsHdfsResp) Set(val *DfsHdfsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsHdfsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsHdfsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsHdfsResp(val *DfsHdfsResp) *NullableDfsHdfsResp {
	return &NullableDfsHdfsResp{value: val, isSet: true}
}

func (v NullableDfsHdfsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsHdfsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


