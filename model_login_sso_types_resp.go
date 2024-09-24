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

// checks if the LoginSsoTypesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginSsoTypesResp{}

// LoginSsoTypesResp struct for LoginSsoTypesResp
type LoginSsoTypesResp struct {
	// SSOTypes use to Unmarshal config
	SsoTypes []SsoType `json:"sso_types,omitempty"`
}

// NewLoginSsoTypesResp instantiates a new LoginSsoTypesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginSsoTypesResp() *LoginSsoTypesResp {
	this := LoginSsoTypesResp{}
	return &this
}

// NewLoginSsoTypesRespWithDefaults instantiates a new LoginSsoTypesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginSsoTypesRespWithDefaults() *LoginSsoTypesResp {
	this := LoginSsoTypesResp{}
	return &this
}

// GetSsoTypes returns the SsoTypes field value if set, zero value otherwise.
func (o *LoginSsoTypesResp) GetSsoTypes() []SsoType {
	if o == nil || IsNil(o.SsoTypes) {
		var ret []SsoType
		return ret
	}
	return o.SsoTypes
}

// GetSsoTypesOk returns a tuple with the SsoTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginSsoTypesResp) GetSsoTypesOk() ([]SsoType, bool) {
	if o == nil || IsNil(o.SsoTypes) {
		return nil, false
	}
	return o.SsoTypes, true
}

// HasSsoTypes returns a boolean if a field has been set.
func (o *LoginSsoTypesResp) HasSsoTypes() bool {
	if o != nil && !IsNil(o.SsoTypes) {
		return true
	}

	return false
}

// SetSsoTypes gets a reference to the given []SsoType and assigns it to the SsoTypes field.
func (o *LoginSsoTypesResp) SetSsoTypes(v []SsoType) {
	o.SsoTypes = v
}

func (o LoginSsoTypesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginSsoTypesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SsoTypes) {
		toSerialize["sso_types"] = o.SsoTypes
	}
	return toSerialize, nil
}

type NullableLoginSsoTypesResp struct {
	value *LoginSsoTypesResp
	isSet bool
}

func (v NullableLoginSsoTypesResp) Get() *LoginSsoTypesResp {
	return v.value
}

func (v *NullableLoginSsoTypesResp) Set(val *LoginSsoTypesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginSsoTypesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginSsoTypesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginSsoTypesResp(val *LoginSsoTypesResp) *NullableLoginSsoTypesResp {
	return &NullableLoginSsoTypesResp{value: val, isSet: true}
}

func (v NullableLoginSsoTypesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginSsoTypesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


