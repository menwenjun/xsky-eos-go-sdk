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

// checks if the DfsHdfsUpdateIPWhiteListsReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsHdfsUpdateIPWhiteListsReq{}

// DfsHdfsUpdateIPWhiteListsReq struct for DfsHdfsUpdateIPWhiteListsReq
type DfsHdfsUpdateIPWhiteListsReq struct {
	DfsHdfs DfsHdfsUpdateIPWhiteListsReqHdfs `json:"dfs_hdfs"`
}

type _DfsHdfsUpdateIPWhiteListsReq DfsHdfsUpdateIPWhiteListsReq

// NewDfsHdfsUpdateIPWhiteListsReq instantiates a new DfsHdfsUpdateIPWhiteListsReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsHdfsUpdateIPWhiteListsReq(dfsHdfs DfsHdfsUpdateIPWhiteListsReqHdfs) *DfsHdfsUpdateIPWhiteListsReq {
	this := DfsHdfsUpdateIPWhiteListsReq{}
	this.DfsHdfs = dfsHdfs
	return &this
}

// NewDfsHdfsUpdateIPWhiteListsReqWithDefaults instantiates a new DfsHdfsUpdateIPWhiteListsReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsHdfsUpdateIPWhiteListsReqWithDefaults() *DfsHdfsUpdateIPWhiteListsReq {
	this := DfsHdfsUpdateIPWhiteListsReq{}
	return &this
}

// GetDfsHdfs returns the DfsHdfs field value
func (o *DfsHdfsUpdateIPWhiteListsReq) GetDfsHdfs() DfsHdfsUpdateIPWhiteListsReqHdfs {
	if o == nil {
		var ret DfsHdfsUpdateIPWhiteListsReqHdfs
		return ret
	}

	return o.DfsHdfs
}

// GetDfsHdfsOk returns a tuple with the DfsHdfs field value
// and a boolean to check if the value has been set.
func (o *DfsHdfsUpdateIPWhiteListsReq) GetDfsHdfsOk() (*DfsHdfsUpdateIPWhiteListsReqHdfs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsHdfs, true
}

// SetDfsHdfs sets field value
func (o *DfsHdfsUpdateIPWhiteListsReq) SetDfsHdfs(v DfsHdfsUpdateIPWhiteListsReqHdfs) {
	o.DfsHdfs = v
}

func (o DfsHdfsUpdateIPWhiteListsReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsHdfsUpdateIPWhiteListsReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_hdfs"] = o.DfsHdfs
	return toSerialize, nil
}

func (o *DfsHdfsUpdateIPWhiteListsReq) UnmarshalJSON(data []byte) (err error) {
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

	varDfsHdfsUpdateIPWhiteListsReq := _DfsHdfsUpdateIPWhiteListsReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsHdfsUpdateIPWhiteListsReq)

	if err != nil {
		return err
	}

	*o = DfsHdfsUpdateIPWhiteListsReq(varDfsHdfsUpdateIPWhiteListsReq)

	return err
}

type NullableDfsHdfsUpdateIPWhiteListsReq struct {
	value *DfsHdfsUpdateIPWhiteListsReq
	isSet bool
}

func (v NullableDfsHdfsUpdateIPWhiteListsReq) Get() *DfsHdfsUpdateIPWhiteListsReq {
	return v.value
}

func (v *NullableDfsHdfsUpdateIPWhiteListsReq) Set(val *DfsHdfsUpdateIPWhiteListsReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsHdfsUpdateIPWhiteListsReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsHdfsUpdateIPWhiteListsReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsHdfsUpdateIPWhiteListsReq(val *DfsHdfsUpdateIPWhiteListsReq) *NullableDfsHdfsUpdateIPWhiteListsReq {
	return &NullableDfsHdfsUpdateIPWhiteListsReq{value: val, isSet: true}
}

func (v NullableDfsHdfsUpdateIPWhiteListsReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsHdfsUpdateIPWhiteListsReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


