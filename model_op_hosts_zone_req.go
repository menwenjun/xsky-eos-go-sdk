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

// checks if the OpHostsZoneReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpHostsZoneReq{}

// OpHostsZoneReq struct for OpHostsZoneReq
type OpHostsZoneReq struct {
	HostsZoneInfo *OpHostsZoneReqHostsZoneInfo `json:"hosts_zone_info,omitempty"`
}

// NewOpHostsZoneReq instantiates a new OpHostsZoneReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpHostsZoneReq() *OpHostsZoneReq {
	this := OpHostsZoneReq{}
	return &this
}

// NewOpHostsZoneReqWithDefaults instantiates a new OpHostsZoneReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpHostsZoneReqWithDefaults() *OpHostsZoneReq {
	this := OpHostsZoneReq{}
	return &this
}

// GetHostsZoneInfo returns the HostsZoneInfo field value if set, zero value otherwise.
func (o *OpHostsZoneReq) GetHostsZoneInfo() OpHostsZoneReqHostsZoneInfo {
	if o == nil || IsNil(o.HostsZoneInfo) {
		var ret OpHostsZoneReqHostsZoneInfo
		return ret
	}
	return *o.HostsZoneInfo
}

// GetHostsZoneInfoOk returns a tuple with the HostsZoneInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpHostsZoneReq) GetHostsZoneInfoOk() (*OpHostsZoneReqHostsZoneInfo, bool) {
	if o == nil || IsNil(o.HostsZoneInfo) {
		return nil, false
	}
	return o.HostsZoneInfo, true
}

// HasHostsZoneInfo returns a boolean if a field has been set.
func (o *OpHostsZoneReq) HasHostsZoneInfo() bool {
	if o != nil && !IsNil(o.HostsZoneInfo) {
		return true
	}

	return false
}

// SetHostsZoneInfo gets a reference to the given OpHostsZoneReqHostsZoneInfo and assigns it to the HostsZoneInfo field.
func (o *OpHostsZoneReq) SetHostsZoneInfo(v OpHostsZoneReqHostsZoneInfo) {
	o.HostsZoneInfo = &v
}

func (o OpHostsZoneReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpHostsZoneReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HostsZoneInfo) {
		toSerialize["hosts_zone_info"] = o.HostsZoneInfo
	}
	return toSerialize, nil
}

type NullableOpHostsZoneReq struct {
	value *OpHostsZoneReq
	isSet bool
}

func (v NullableOpHostsZoneReq) Get() *OpHostsZoneReq {
	return v.value
}

func (v *NullableOpHostsZoneReq) Set(val *OpHostsZoneReq) {
	v.value = val
	v.isSet = true
}

func (v NullableOpHostsZoneReq) IsSet() bool {
	return v.isSet
}

func (v *NullableOpHostsZoneReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpHostsZoneReq(val *OpHostsZoneReq) *NullableOpHostsZoneReq {
	return &NullableOpHostsZoneReq{value: val, isSet: true}
}

func (v NullableOpHostsZoneReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpHostsZoneReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


