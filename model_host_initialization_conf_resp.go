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

// checks if the HostInitializationConfResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostInitializationConfResp{}

// HostInitializationConfResp struct for HostInitializationConfResp
type HostInitializationConfResp struct {
	HostInitializationConf *HostInitializationConf `json:"host_initialization_conf,omitempty"`
}

// NewHostInitializationConfResp instantiates a new HostInitializationConfResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostInitializationConfResp() *HostInitializationConfResp {
	this := HostInitializationConfResp{}
	return &this
}

// NewHostInitializationConfRespWithDefaults instantiates a new HostInitializationConfResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostInitializationConfRespWithDefaults() *HostInitializationConfResp {
	this := HostInitializationConfResp{}
	return &this
}

// GetHostInitializationConf returns the HostInitializationConf field value if set, zero value otherwise.
func (o *HostInitializationConfResp) GetHostInitializationConf() HostInitializationConf {
	if o == nil || IsNil(o.HostInitializationConf) {
		var ret HostInitializationConf
		return ret
	}
	return *o.HostInitializationConf
}

// GetHostInitializationConfOk returns a tuple with the HostInitializationConf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitializationConfResp) GetHostInitializationConfOk() (*HostInitializationConf, bool) {
	if o == nil || IsNil(o.HostInitializationConf) {
		return nil, false
	}
	return o.HostInitializationConf, true
}

// HasHostInitializationConf returns a boolean if a field has been set.
func (o *HostInitializationConfResp) HasHostInitializationConf() bool {
	if o != nil && !IsNil(o.HostInitializationConf) {
		return true
	}

	return false
}

// SetHostInitializationConf gets a reference to the given HostInitializationConf and assigns it to the HostInitializationConf field.
func (o *HostInitializationConfResp) SetHostInitializationConf(v HostInitializationConf) {
	o.HostInitializationConf = &v
}

func (o HostInitializationConfResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostInitializationConfResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HostInitializationConf) {
		toSerialize["host_initialization_conf"] = o.HostInitializationConf
	}
	return toSerialize, nil
}

type NullableHostInitializationConfResp struct {
	value *HostInitializationConfResp
	isSet bool
}

func (v NullableHostInitializationConfResp) Get() *HostInitializationConfResp {
	return v.value
}

func (v *NullableHostInitializationConfResp) Set(val *HostInitializationConfResp) {
	v.value = val
	v.isSet = true
}

func (v NullableHostInitializationConfResp) IsSet() bool {
	return v.isSet
}

func (v *NullableHostInitializationConfResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostInitializationConfResp(val *HostInitializationConfResp) *NullableHostInitializationConfResp {
	return &NullableHostInitializationConfResp{value: val, isSet: true}
}

func (v NullableHostInitializationConfResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostInitializationConfResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


