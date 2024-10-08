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

// checks if the VIPInstanceResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VIPInstanceResp{}

// VIPInstanceResp struct for VIPInstanceResp
type VIPInstanceResp struct {
	VipInstance *VIPInstance `json:"vip_instance,omitempty"`
}

// NewVIPInstanceResp instantiates a new VIPInstanceResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVIPInstanceResp() *VIPInstanceResp {
	this := VIPInstanceResp{}
	return &this
}

// NewVIPInstanceRespWithDefaults instantiates a new VIPInstanceResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVIPInstanceRespWithDefaults() *VIPInstanceResp {
	this := VIPInstanceResp{}
	return &this
}

// GetVipInstance returns the VipInstance field value if set, zero value otherwise.
func (o *VIPInstanceResp) GetVipInstance() VIPInstance {
	if o == nil || IsNil(o.VipInstance) {
		var ret VIPInstance
		return ret
	}
	return *o.VipInstance
}

// GetVipInstanceOk returns a tuple with the VipInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VIPInstanceResp) GetVipInstanceOk() (*VIPInstance, bool) {
	if o == nil || IsNil(o.VipInstance) {
		return nil, false
	}
	return o.VipInstance, true
}

// HasVipInstance returns a boolean if a field has been set.
func (o *VIPInstanceResp) HasVipInstance() bool {
	if o != nil && !IsNil(o.VipInstance) {
		return true
	}

	return false
}

// SetVipInstance gets a reference to the given VIPInstance and assigns it to the VipInstance field.
func (o *VIPInstanceResp) SetVipInstance(v VIPInstance) {
	o.VipInstance = &v
}

func (o VIPInstanceResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VIPInstanceResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VipInstance) {
		toSerialize["vip_instance"] = o.VipInstance
	}
	return toSerialize, nil
}

type NullableVIPInstanceResp struct {
	value *VIPInstanceResp
	isSet bool
}

func (v NullableVIPInstanceResp) Get() *VIPInstanceResp {
	return v.value
}

func (v *NullableVIPInstanceResp) Set(val *VIPInstanceResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVIPInstanceResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVIPInstanceResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVIPInstanceResp(val *VIPInstanceResp) *NullableVIPInstanceResp {
	return &NullableVIPInstanceResp{value: val, isSet: true}
}

func (v NullableVIPInstanceResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVIPInstanceResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


