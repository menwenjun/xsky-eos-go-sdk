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

// checks if the DfsFileDeleteReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsFileDeleteReq{}

// DfsFileDeleteReq struct for DfsFileDeleteReq
type DfsFileDeleteReq struct {
	DfsFile DfsFileDeleteReqFile `json:"dfs_file"`
}

type _DfsFileDeleteReq DfsFileDeleteReq

// NewDfsFileDeleteReq instantiates a new DfsFileDeleteReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsFileDeleteReq(dfsFile DfsFileDeleteReqFile) *DfsFileDeleteReq {
	this := DfsFileDeleteReq{}
	this.DfsFile = dfsFile
	return &this
}

// NewDfsFileDeleteReqWithDefaults instantiates a new DfsFileDeleteReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsFileDeleteReqWithDefaults() *DfsFileDeleteReq {
	this := DfsFileDeleteReq{}
	return &this
}

// GetDfsFile returns the DfsFile field value
func (o *DfsFileDeleteReq) GetDfsFile() DfsFileDeleteReqFile {
	if o == nil {
		var ret DfsFileDeleteReqFile
		return ret
	}

	return o.DfsFile
}

// GetDfsFileOk returns a tuple with the DfsFile field value
// and a boolean to check if the value has been set.
func (o *DfsFileDeleteReq) GetDfsFileOk() (*DfsFileDeleteReqFile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsFile, true
}

// SetDfsFile sets field value
func (o *DfsFileDeleteReq) SetDfsFile(v DfsFileDeleteReqFile) {
	o.DfsFile = v
}

func (o DfsFileDeleteReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsFileDeleteReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_file"] = o.DfsFile
	return toSerialize, nil
}

func (o *DfsFileDeleteReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_file",
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

	varDfsFileDeleteReq := _DfsFileDeleteReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsFileDeleteReq)

	if err != nil {
		return err
	}

	*o = DfsFileDeleteReq(varDfsFileDeleteReq)

	return err
}

type NullableDfsFileDeleteReq struct {
	value *DfsFileDeleteReq
	isSet bool
}

func (v NullableDfsFileDeleteReq) Get() *DfsFileDeleteReq {
	return v.value
}

func (v *NullableDfsFileDeleteReq) Set(val *DfsFileDeleteReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsFileDeleteReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsFileDeleteReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsFileDeleteReq(val *DfsFileDeleteReq) *NullableDfsFileDeleteReq {
	return &NullableDfsFileDeleteReq{value: val, isSet: true}
}

func (v NullableDfsFileDeleteReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsFileDeleteReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


