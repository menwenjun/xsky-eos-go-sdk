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

// checks if the AuthSecurityPolicyResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthSecurityPolicyResp{}

// AuthSecurityPolicyResp struct for AuthSecurityPolicyResp
type AuthSecurityPolicyResp struct {
	AuthSecurityPolicy *AuthSecurityPolicy `json:"auth_security_policy,omitempty"`
}

// NewAuthSecurityPolicyResp instantiates a new AuthSecurityPolicyResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthSecurityPolicyResp() *AuthSecurityPolicyResp {
	this := AuthSecurityPolicyResp{}
	return &this
}

// NewAuthSecurityPolicyRespWithDefaults instantiates a new AuthSecurityPolicyResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthSecurityPolicyRespWithDefaults() *AuthSecurityPolicyResp {
	this := AuthSecurityPolicyResp{}
	return &this
}

// GetAuthSecurityPolicy returns the AuthSecurityPolicy field value if set, zero value otherwise.
func (o *AuthSecurityPolicyResp) GetAuthSecurityPolicy() AuthSecurityPolicy {
	if o == nil || IsNil(o.AuthSecurityPolicy) {
		var ret AuthSecurityPolicy
		return ret
	}
	return *o.AuthSecurityPolicy
}

// GetAuthSecurityPolicyOk returns a tuple with the AuthSecurityPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthSecurityPolicyResp) GetAuthSecurityPolicyOk() (*AuthSecurityPolicy, bool) {
	if o == nil || IsNil(o.AuthSecurityPolicy) {
		return nil, false
	}
	return o.AuthSecurityPolicy, true
}

// HasAuthSecurityPolicy returns a boolean if a field has been set.
func (o *AuthSecurityPolicyResp) HasAuthSecurityPolicy() bool {
	if o != nil && !IsNil(o.AuthSecurityPolicy) {
		return true
	}

	return false
}

// SetAuthSecurityPolicy gets a reference to the given AuthSecurityPolicy and assigns it to the AuthSecurityPolicy field.
func (o *AuthSecurityPolicyResp) SetAuthSecurityPolicy(v AuthSecurityPolicy) {
	o.AuthSecurityPolicy = &v
}

func (o AuthSecurityPolicyResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthSecurityPolicyResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthSecurityPolicy) {
		toSerialize["auth_security_policy"] = o.AuthSecurityPolicy
	}
	return toSerialize, nil
}

type NullableAuthSecurityPolicyResp struct {
	value *AuthSecurityPolicyResp
	isSet bool
}

func (v NullableAuthSecurityPolicyResp) Get() *AuthSecurityPolicyResp {
	return v.value
}

func (v *NullableAuthSecurityPolicyResp) Set(val *AuthSecurityPolicyResp) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthSecurityPolicyResp) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthSecurityPolicyResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthSecurityPolicyResp(val *AuthSecurityPolicyResp) *NullableAuthSecurityPolicyResp {
	return &NullableAuthSecurityPolicyResp{value: val, isSet: true}
}

func (v NullableAuthSecurityPolicyResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthSecurityPolicyResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


