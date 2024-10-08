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

// checks if the AdminVIPResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminVIPResp{}

// AdminVIPResp struct for AdminVIPResp
type AdminVIPResp struct {
	AdminVip *AdminVIPResult `json:"admin_vip,omitempty"`
}

// NewAdminVIPResp instantiates a new AdminVIPResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminVIPResp() *AdminVIPResp {
	this := AdminVIPResp{}
	return &this
}

// NewAdminVIPRespWithDefaults instantiates a new AdminVIPResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminVIPRespWithDefaults() *AdminVIPResp {
	this := AdminVIPResp{}
	return &this
}

// GetAdminVip returns the AdminVip field value if set, zero value otherwise.
func (o *AdminVIPResp) GetAdminVip() AdminVIPResult {
	if o == nil || IsNil(o.AdminVip) {
		var ret AdminVIPResult
		return ret
	}
	return *o.AdminVip
}

// GetAdminVipOk returns a tuple with the AdminVip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminVIPResp) GetAdminVipOk() (*AdminVIPResult, bool) {
	if o == nil || IsNil(o.AdminVip) {
		return nil, false
	}
	return o.AdminVip, true
}

// HasAdminVip returns a boolean if a field has been set.
func (o *AdminVIPResp) HasAdminVip() bool {
	if o != nil && !IsNil(o.AdminVip) {
		return true
	}

	return false
}

// SetAdminVip gets a reference to the given AdminVIPResult and assigns it to the AdminVip field.
func (o *AdminVIPResp) SetAdminVip(v AdminVIPResult) {
	o.AdminVip = &v
}

func (o AdminVIPResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminVIPResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminVip) {
		toSerialize["admin_vip"] = o.AdminVip
	}
	return toSerialize, nil
}

type NullableAdminVIPResp struct {
	value *AdminVIPResp
	isSet bool
}

func (v NullableAdminVIPResp) Get() *AdminVIPResp {
	return v.value
}

func (v *NullableAdminVIPResp) Set(val *AdminVIPResp) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminVIPResp) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminVIPResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminVIPResp(val *AdminVIPResp) *NullableAdminVIPResp {
	return &NullableAdminVIPResp{value: val, isSet: true}
}

func (v NullableAdminVIPResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminVIPResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


