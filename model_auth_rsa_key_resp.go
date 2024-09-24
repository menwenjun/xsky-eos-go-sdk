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

// checks if the AuthRSAKeyResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthRSAKeyResp{}

// AuthRSAKeyResp struct for AuthRSAKeyResp
type AuthRSAKeyResp struct {
	RsaKey *AuthRSAKey `json:"rsa_key,omitempty"`
}

// NewAuthRSAKeyResp instantiates a new AuthRSAKeyResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthRSAKeyResp() *AuthRSAKeyResp {
	this := AuthRSAKeyResp{}
	return &this
}

// NewAuthRSAKeyRespWithDefaults instantiates a new AuthRSAKeyResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthRSAKeyRespWithDefaults() *AuthRSAKeyResp {
	this := AuthRSAKeyResp{}
	return &this
}

// GetRsaKey returns the RsaKey field value if set, zero value otherwise.
func (o *AuthRSAKeyResp) GetRsaKey() AuthRSAKey {
	if o == nil || IsNil(o.RsaKey) {
		var ret AuthRSAKey
		return ret
	}
	return *o.RsaKey
}

// GetRsaKeyOk returns a tuple with the RsaKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRSAKeyResp) GetRsaKeyOk() (*AuthRSAKey, bool) {
	if o == nil || IsNil(o.RsaKey) {
		return nil, false
	}
	return o.RsaKey, true
}

// HasRsaKey returns a boolean if a field has been set.
func (o *AuthRSAKeyResp) HasRsaKey() bool {
	if o != nil && !IsNil(o.RsaKey) {
		return true
	}

	return false
}

// SetRsaKey gets a reference to the given AuthRSAKey and assigns it to the RsaKey field.
func (o *AuthRSAKeyResp) SetRsaKey(v AuthRSAKey) {
	o.RsaKey = &v
}

func (o AuthRSAKeyResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthRSAKeyResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RsaKey) {
		toSerialize["rsa_key"] = o.RsaKey
	}
	return toSerialize, nil
}

type NullableAuthRSAKeyResp struct {
	value *AuthRSAKeyResp
	isSet bool
}

func (v NullableAuthRSAKeyResp) Get() *AuthRSAKeyResp {
	return v.value
}

func (v *NullableAuthRSAKeyResp) Set(val *AuthRSAKeyResp) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthRSAKeyResp) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthRSAKeyResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthRSAKeyResp(val *AuthRSAKeyResp) *NullableAuthRSAKeyResp {
	return &NullableAuthRSAKeyResp{value: val, isSet: true}
}

func (v NullableAuthRSAKeyResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthRSAKeyResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


