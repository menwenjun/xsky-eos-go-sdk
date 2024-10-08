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

// checks if the DfsCSIDirectoryRespDirectory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsCSIDirectoryRespDirectory{}

// DfsCSIDirectoryRespDirectory struct for DfsCSIDirectoryRespDirectory
type DfsCSIDirectoryRespDirectory struct {
	DfsRootfs DfsRootfs `json:"dfs_rootfs"`
	// directory path
	Path string `json:"path"`
}

type _DfsCSIDirectoryRespDirectory DfsCSIDirectoryRespDirectory

// NewDfsCSIDirectoryRespDirectory instantiates a new DfsCSIDirectoryRespDirectory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsCSIDirectoryRespDirectory(dfsRootfs DfsRootfs, path string) *DfsCSIDirectoryRespDirectory {
	this := DfsCSIDirectoryRespDirectory{}
	this.DfsRootfs = dfsRootfs
	this.Path = path
	return &this
}

// NewDfsCSIDirectoryRespDirectoryWithDefaults instantiates a new DfsCSIDirectoryRespDirectory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsCSIDirectoryRespDirectoryWithDefaults() *DfsCSIDirectoryRespDirectory {
	this := DfsCSIDirectoryRespDirectory{}
	return &this
}

// GetDfsRootfs returns the DfsRootfs field value
func (o *DfsCSIDirectoryRespDirectory) GetDfsRootfs() DfsRootfs {
	if o == nil {
		var ret DfsRootfs
		return ret
	}

	return o.DfsRootfs
}

// GetDfsRootfsOk returns a tuple with the DfsRootfs field value
// and a boolean to check if the value has been set.
func (o *DfsCSIDirectoryRespDirectory) GetDfsRootfsOk() (*DfsRootfs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsRootfs, true
}

// SetDfsRootfs sets field value
func (o *DfsCSIDirectoryRespDirectory) SetDfsRootfs(v DfsRootfs) {
	o.DfsRootfs = v
}

// GetPath returns the Path field value
func (o *DfsCSIDirectoryRespDirectory) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DfsCSIDirectoryRespDirectory) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DfsCSIDirectoryRespDirectory) SetPath(v string) {
	o.Path = v
}

func (o DfsCSIDirectoryRespDirectory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsCSIDirectoryRespDirectory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_rootfs"] = o.DfsRootfs
	toSerialize["path"] = o.Path
	return toSerialize, nil
}

func (o *DfsCSIDirectoryRespDirectory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_rootfs",
		"path",
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

	varDfsCSIDirectoryRespDirectory := _DfsCSIDirectoryRespDirectory{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsCSIDirectoryRespDirectory)

	if err != nil {
		return err
	}

	*o = DfsCSIDirectoryRespDirectory(varDfsCSIDirectoryRespDirectory)

	return err
}

type NullableDfsCSIDirectoryRespDirectory struct {
	value *DfsCSIDirectoryRespDirectory
	isSet bool
}

func (v NullableDfsCSIDirectoryRespDirectory) Get() *DfsCSIDirectoryRespDirectory {
	return v.value
}

func (v *NullableDfsCSIDirectoryRespDirectory) Set(val *DfsCSIDirectoryRespDirectory) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsCSIDirectoryRespDirectory) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsCSIDirectoryRespDirectory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsCSIDirectoryRespDirectory(val *DfsCSIDirectoryRespDirectory) *NullableDfsCSIDirectoryRespDirectory {
	return &NullableDfsCSIDirectoryRespDirectory{value: val, isSet: true}
}

func (v NullableDfsCSIDirectoryRespDirectory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsCSIDirectoryRespDirectory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


