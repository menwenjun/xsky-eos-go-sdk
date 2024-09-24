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

// checks if the DfsHdfsRemoveIPWhiteListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsHdfsRemoveIPWhiteListReq{}

// DfsHdfsRemoveIPWhiteListReq struct for DfsHdfsRemoveIPWhiteListReq
type DfsHdfsRemoveIPWhiteListReq struct {
	DfsHdfs DfsHdfsRemoveIPWhiteListReqHdfs `json:"dfs_hdfs"`
}

type _DfsHdfsRemoveIPWhiteListReq DfsHdfsRemoveIPWhiteListReq

// NewDfsHdfsRemoveIPWhiteListReq instantiates a new DfsHdfsRemoveIPWhiteListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsHdfsRemoveIPWhiteListReq(dfsHdfs DfsHdfsRemoveIPWhiteListReqHdfs) *DfsHdfsRemoveIPWhiteListReq {
	this := DfsHdfsRemoveIPWhiteListReq{}
	this.DfsHdfs = dfsHdfs
	return &this
}

// NewDfsHdfsRemoveIPWhiteListReqWithDefaults instantiates a new DfsHdfsRemoveIPWhiteListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsHdfsRemoveIPWhiteListReqWithDefaults() *DfsHdfsRemoveIPWhiteListReq {
	this := DfsHdfsRemoveIPWhiteListReq{}
	return &this
}

// GetDfsHdfs returns the DfsHdfs field value
func (o *DfsHdfsRemoveIPWhiteListReq) GetDfsHdfs() DfsHdfsRemoveIPWhiteListReqHdfs {
	if o == nil {
		var ret DfsHdfsRemoveIPWhiteListReqHdfs
		return ret
	}

	return o.DfsHdfs
}

// GetDfsHdfsOk returns a tuple with the DfsHdfs field value
// and a boolean to check if the value has been set.
func (o *DfsHdfsRemoveIPWhiteListReq) GetDfsHdfsOk() (*DfsHdfsRemoveIPWhiteListReqHdfs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsHdfs, true
}

// SetDfsHdfs sets field value
func (o *DfsHdfsRemoveIPWhiteListReq) SetDfsHdfs(v DfsHdfsRemoveIPWhiteListReqHdfs) {
	o.DfsHdfs = v
}

func (o DfsHdfsRemoveIPWhiteListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsHdfsRemoveIPWhiteListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_hdfs"] = o.DfsHdfs
	return toSerialize, nil
}

func (o *DfsHdfsRemoveIPWhiteListReq) UnmarshalJSON(data []byte) (err error) {
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

	varDfsHdfsRemoveIPWhiteListReq := _DfsHdfsRemoveIPWhiteListReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsHdfsRemoveIPWhiteListReq)

	if err != nil {
		return err
	}

	*o = DfsHdfsRemoveIPWhiteListReq(varDfsHdfsRemoveIPWhiteListReq)

	return err
}

type NullableDfsHdfsRemoveIPWhiteListReq struct {
	value *DfsHdfsRemoveIPWhiteListReq
	isSet bool
}

func (v NullableDfsHdfsRemoveIPWhiteListReq) Get() *DfsHdfsRemoveIPWhiteListReq {
	return v.value
}

func (v *NullableDfsHdfsRemoveIPWhiteListReq) Set(val *DfsHdfsRemoveIPWhiteListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsHdfsRemoveIPWhiteListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsHdfsRemoveIPWhiteListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsHdfsRemoveIPWhiteListReq(val *DfsHdfsRemoveIPWhiteListReq) *NullableDfsHdfsRemoveIPWhiteListReq {
	return &NullableDfsHdfsRemoveIPWhiteListReq{value: val, isSet: true}
}

func (v NullableDfsHdfsRemoveIPWhiteListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsHdfsRemoveIPWhiteListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


