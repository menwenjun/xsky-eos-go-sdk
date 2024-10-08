/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AuthLoginReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthLoginReq{}

// AuthLoginReq struct for AuthLoginReq
type AuthLoginReq struct {
	Auth AuthLoginReqAuth `json:"auth"`
}

type _AuthLoginReq AuthLoginReq

// NewAuthLoginReq instantiates a new AuthLoginReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthLoginReq(auth AuthLoginReqAuth) *AuthLoginReq {
	this := AuthLoginReq{}
	this.Auth = auth
	return &this
}

// NewAuthLoginReqWithDefaults instantiates a new AuthLoginReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthLoginReqWithDefaults() *AuthLoginReq {
	this := AuthLoginReq{}
	return &this
}

// GetAuth returns the Auth field value
func (o *AuthLoginReq) GetAuth() AuthLoginReqAuth {
	if o == nil {
		var ret AuthLoginReqAuth
		return ret
	}

	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *AuthLoginReq) GetAuthOk() (*AuthLoginReqAuth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value
func (o *AuthLoginReq) SetAuth(v AuthLoginReqAuth) {
	o.Auth = v
}

func (o AuthLoginReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthLoginReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auth"] = o.Auth
	return toSerialize, nil
}

func (o *AuthLoginReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auth",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAuthLoginReq := _AuthLoginReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthLoginReq)

	if err != nil {
		return err
	}

	*o = AuthLoginReq(varAuthLoginReq)

	return err
}

type NullableAuthLoginReq struct {
	value *AuthLoginReq
	isSet bool
}

func (v NullableAuthLoginReq) Get() *AuthLoginReq {
	return v.value
}

func (v *NullableAuthLoginReq) Set(val *AuthLoginReq) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthLoginReq) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthLoginReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthLoginReq(val *AuthLoginReq) *NullableAuthLoginReq {
	return &NullableAuthLoginReq{value: val, isSet: true}
}

func (v NullableAuthLoginReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthLoginReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


