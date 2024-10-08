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

// checks if the DfsPortsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsPortsResp{}

// DfsPortsResp struct for DfsPortsResp
type DfsPortsResp struct {
	DfsPorts *map[string][]int64 `json:"dfs_ports,omitempty"`
}

// NewDfsPortsResp instantiates a new DfsPortsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsPortsResp() *DfsPortsResp {
	this := DfsPortsResp{}
	return &this
}

// NewDfsPortsRespWithDefaults instantiates a new DfsPortsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsPortsRespWithDefaults() *DfsPortsResp {
	this := DfsPortsResp{}
	return &this
}

// GetDfsPorts returns the DfsPorts field value if set, zero value otherwise.
func (o *DfsPortsResp) GetDfsPorts() map[string][]int64 {
	if o == nil || IsNil(o.DfsPorts) {
		var ret map[string][]int64
		return ret
	}
	return *o.DfsPorts
}

// GetDfsPortsOk returns a tuple with the DfsPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPortsResp) GetDfsPortsOk() (*map[string][]int64, bool) {
	if o == nil || IsNil(o.DfsPorts) {
		return nil, false
	}
	return o.DfsPorts, true
}

// HasDfsPorts returns a boolean if a field has been set.
func (o *DfsPortsResp) HasDfsPorts() bool {
	if o != nil && !IsNil(o.DfsPorts) {
		return true
	}

	return false
}

// SetDfsPorts gets a reference to the given map[string][]int64 and assigns it to the DfsPorts field.
func (o *DfsPortsResp) SetDfsPorts(v map[string][]int64) {
	o.DfsPorts = &v
}

func (o DfsPortsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsPortsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DfsPorts) {
		toSerialize["dfs_ports"] = o.DfsPorts
	}
	return toSerialize, nil
}

type NullableDfsPortsResp struct {
	value *DfsPortsResp
	isSet bool
}

func (v NullableDfsPortsResp) Get() *DfsPortsResp {
	return v.value
}

func (v *NullableDfsPortsResp) Set(val *DfsPortsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsPortsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsPortsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsPortsResp(val *DfsPortsResp) *NullableDfsPortsResp {
	return &NullableDfsPortsResp{value: val, isSet: true}
}

func (v NullableDfsPortsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsPortsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


