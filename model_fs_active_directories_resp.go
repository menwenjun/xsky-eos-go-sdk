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

// checks if the FSActiveDirectoriesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FSActiveDirectoriesResp{}

// FSActiveDirectoriesResp struct for FSActiveDirectoriesResp
type FSActiveDirectoriesResp struct {
	// file storage active directories
	FsActiveDirectories []FSActiveDirectory `json:"fs_active_directories"`
}

type _FSActiveDirectoriesResp FSActiveDirectoriesResp

// NewFSActiveDirectoriesResp instantiates a new FSActiveDirectoriesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFSActiveDirectoriesResp(fsActiveDirectories []FSActiveDirectory) *FSActiveDirectoriesResp {
	this := FSActiveDirectoriesResp{}
	this.FsActiveDirectories = fsActiveDirectories
	return &this
}

// NewFSActiveDirectoriesRespWithDefaults instantiates a new FSActiveDirectoriesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFSActiveDirectoriesRespWithDefaults() *FSActiveDirectoriesResp {
	this := FSActiveDirectoriesResp{}
	return &this
}

// GetFsActiveDirectories returns the FsActiveDirectories field value
func (o *FSActiveDirectoriesResp) GetFsActiveDirectories() []FSActiveDirectory {
	if o == nil {
		var ret []FSActiveDirectory
		return ret
	}

	return o.FsActiveDirectories
}

// GetFsActiveDirectoriesOk returns a tuple with the FsActiveDirectories field value
// and a boolean to check if the value has been set.
func (o *FSActiveDirectoriesResp) GetFsActiveDirectoriesOk() ([]FSActiveDirectory, bool) {
	if o == nil {
		return nil, false
	}
	return o.FsActiveDirectories, true
}

// SetFsActiveDirectories sets field value
func (o *FSActiveDirectoriesResp) SetFsActiveDirectories(v []FSActiveDirectory) {
	o.FsActiveDirectories = v
}

func (o FSActiveDirectoriesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FSActiveDirectoriesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fs_active_directories"] = o.FsActiveDirectories
	return toSerialize, nil
}

func (o *FSActiveDirectoriesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fs_active_directories",
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

	varFSActiveDirectoriesResp := _FSActiveDirectoriesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFSActiveDirectoriesResp)

	if err != nil {
		return err
	}

	*o = FSActiveDirectoriesResp(varFSActiveDirectoriesResp)

	return err
}

type NullableFSActiveDirectoriesResp struct {
	value *FSActiveDirectoriesResp
	isSet bool
}

func (v NullableFSActiveDirectoriesResp) Get() *FSActiveDirectoriesResp {
	return v.value
}

func (v *NullableFSActiveDirectoriesResp) Set(val *FSActiveDirectoriesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableFSActiveDirectoriesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableFSActiveDirectoriesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFSActiveDirectoriesResp(val *FSActiveDirectoriesResp) *NullableFSActiveDirectoriesResp {
	return &NullableFSActiveDirectoriesResp{value: val, isSet: true}
}

func (v NullableFSActiveDirectoriesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFSActiveDirectoriesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


