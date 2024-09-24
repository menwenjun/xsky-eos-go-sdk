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

// checks if the AccessTokenResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenResp{}

// AccessTokenResp struct for AccessTokenResp
type AccessTokenResp struct {
	AccessToken *AccessToken `json:"access_token,omitempty"`
}

// NewAccessTokenResp instantiates a new AccessTokenResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenResp() *AccessTokenResp {
	this := AccessTokenResp{}
	return &this
}

// NewAccessTokenRespWithDefaults instantiates a new AccessTokenResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenRespWithDefaults() *AccessTokenResp {
	this := AccessTokenResp{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *AccessTokenResp) GetAccessToken() AccessToken {
	if o == nil || IsNil(o.AccessToken) {
		var ret AccessToken
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenResp) GetAccessTokenOk() (*AccessToken, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AccessTokenResp) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given AccessToken and assigns it to the AccessToken field.
func (o *AccessTokenResp) SetAccessToken(v AccessToken) {
	o.AccessToken = &v
}

func (o AccessTokenResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokenResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessToken) {
		toSerialize["access_token"] = o.AccessToken
	}
	return toSerialize, nil
}

type NullableAccessTokenResp struct {
	value *AccessTokenResp
	isSet bool
}

func (v NullableAccessTokenResp) Get() *AccessTokenResp {
	return v.value
}

func (v *NullableAccessTokenResp) Set(val *AccessTokenResp) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenResp) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenResp(val *AccessTokenResp) *NullableAccessTokenResp {
	return &NullableAccessTokenResp{value: val, isSet: true}
}

func (v NullableAccessTokenResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


