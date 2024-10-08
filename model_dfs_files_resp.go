/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DfsFilesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsFilesResp{}

// DfsFilesResp struct for DfsFilesResp
type DfsFilesResp struct {
	// dfs file list
	DfsFiles []DfsFile `json:"dfs_files,omitempty"`
	// for ls directory, set to true when reach end
	Eof *bool `json:"eof,omitempty"`
}

// NewDfsFilesResp instantiates a new DfsFilesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsFilesResp() *DfsFilesResp {
	this := DfsFilesResp{}
	return &this
}

// NewDfsFilesRespWithDefaults instantiates a new DfsFilesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsFilesRespWithDefaults() *DfsFilesResp {
	this := DfsFilesResp{}
	return &this
}

// GetDfsFiles returns the DfsFiles field value if set, zero value otherwise.
func (o *DfsFilesResp) GetDfsFiles() []DfsFile {
	if o == nil || IsNil(o.DfsFiles) {
		var ret []DfsFile
		return ret
	}
	return o.DfsFiles
}

// GetDfsFilesOk returns a tuple with the DfsFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFilesResp) GetDfsFilesOk() ([]DfsFile, bool) {
	if o == nil || IsNil(o.DfsFiles) {
		return nil, false
	}
	return o.DfsFiles, true
}

// HasDfsFiles returns a boolean if a field has been set.
func (o *DfsFilesResp) HasDfsFiles() bool {
	if o != nil && !IsNil(o.DfsFiles) {
		return true
	}

	return false
}

// SetDfsFiles gets a reference to the given []DfsFile and assigns it to the DfsFiles field.
func (o *DfsFilesResp) SetDfsFiles(v []DfsFile) {
	o.DfsFiles = v
}

// GetEof returns the Eof field value if set, zero value otherwise.
func (o *DfsFilesResp) GetEof() bool {
	if o == nil || IsNil(o.Eof) {
		var ret bool
		return ret
	}
	return *o.Eof
}

// GetEofOk returns a tuple with the Eof field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFilesResp) GetEofOk() (*bool, bool) {
	if o == nil || IsNil(o.Eof) {
		return nil, false
	}
	return o.Eof, true
}

// HasEof returns a boolean if a field has been set.
func (o *DfsFilesResp) HasEof() bool {
	if o != nil && !IsNil(o.Eof) {
		return true
	}

	return false
}

// SetEof gets a reference to the given bool and assigns it to the Eof field.
func (o *DfsFilesResp) SetEof(v bool) {
	o.Eof = &v
}

func (o DfsFilesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsFilesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DfsFiles) {
		toSerialize["dfs_files"] = o.DfsFiles
	}
	if !IsNil(o.Eof) {
		toSerialize["eof"] = o.Eof
	}
	return toSerialize, nil
}

type NullableDfsFilesResp struct {
	value *DfsFilesResp
	isSet bool
}

func (v NullableDfsFilesResp) Get() *DfsFilesResp {
	return v.value
}

func (v *NullableDfsFilesResp) Set(val *DfsFilesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsFilesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsFilesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsFilesResp(val *DfsFilesResp) *NullableDfsFilesResp {
	return &NullableDfsFilesResp{value: val, isSet: true}
}

func (v NullableDfsFilesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsFilesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


