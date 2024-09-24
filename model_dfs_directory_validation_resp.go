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

// checks if the DfsDirectoryValidationResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsDirectoryValidationResp{}

// DfsDirectoryValidationResp struct for DfsDirectoryValidationResp
type DfsDirectoryValidationResp struct {
	DfsDirectory DfsDirectoryValidationRespDirectory `json:"dfs_directory"`
}

type _DfsDirectoryValidationResp DfsDirectoryValidationResp

// NewDfsDirectoryValidationResp instantiates a new DfsDirectoryValidationResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsDirectoryValidationResp(dfsDirectory DfsDirectoryValidationRespDirectory) *DfsDirectoryValidationResp {
	this := DfsDirectoryValidationResp{}
	this.DfsDirectory = dfsDirectory
	return &this
}

// NewDfsDirectoryValidationRespWithDefaults instantiates a new DfsDirectoryValidationResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsDirectoryValidationRespWithDefaults() *DfsDirectoryValidationResp {
	this := DfsDirectoryValidationResp{}
	return &this
}

// GetDfsDirectory returns the DfsDirectory field value
func (o *DfsDirectoryValidationResp) GetDfsDirectory() DfsDirectoryValidationRespDirectory {
	if o == nil {
		var ret DfsDirectoryValidationRespDirectory
		return ret
	}

	return o.DfsDirectory
}

// GetDfsDirectoryOk returns a tuple with the DfsDirectory field value
// and a boolean to check if the value has been set.
func (o *DfsDirectoryValidationResp) GetDfsDirectoryOk() (*DfsDirectoryValidationRespDirectory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsDirectory, true
}

// SetDfsDirectory sets field value
func (o *DfsDirectoryValidationResp) SetDfsDirectory(v DfsDirectoryValidationRespDirectory) {
	o.DfsDirectory = v
}

func (o DfsDirectoryValidationResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsDirectoryValidationResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_directory"] = o.DfsDirectory
	return toSerialize, nil
}

func (o *DfsDirectoryValidationResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_directory",
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

	varDfsDirectoryValidationResp := _DfsDirectoryValidationResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsDirectoryValidationResp)

	if err != nil {
		return err
	}

	*o = DfsDirectoryValidationResp(varDfsDirectoryValidationResp)

	return err
}

type NullableDfsDirectoryValidationResp struct {
	value *DfsDirectoryValidationResp
	isSet bool
}

func (v NullableDfsDirectoryValidationResp) Get() *DfsDirectoryValidationResp {
	return v.value
}

func (v *NullableDfsDirectoryValidationResp) Set(val *DfsDirectoryValidationResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsDirectoryValidationResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsDirectoryValidationResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsDirectoryValidationResp(val *DfsDirectoryValidationResp) *NullableDfsDirectoryValidationResp {
	return &NullableDfsDirectoryValidationResp{value: val, isSet: true}
}

func (v NullableDfsDirectoryValidationResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsDirectoryValidationResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


