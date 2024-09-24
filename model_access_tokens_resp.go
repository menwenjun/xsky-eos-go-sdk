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

// checks if the AccessTokensResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokensResp{}

// AccessTokensResp struct for AccessTokensResp
type AccessTokensResp struct {
	// access tokens
	AccessTokens []AccessToken `json:"access_tokens"`
}

type _AccessTokensResp AccessTokensResp

// NewAccessTokensResp instantiates a new AccessTokensResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokensResp(accessTokens []AccessToken) *AccessTokensResp {
	this := AccessTokensResp{}
	this.AccessTokens = accessTokens
	return &this
}

// NewAccessTokensRespWithDefaults instantiates a new AccessTokensResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokensRespWithDefaults() *AccessTokensResp {
	this := AccessTokensResp{}
	return &this
}

// GetAccessTokens returns the AccessTokens field value
func (o *AccessTokensResp) GetAccessTokens() []AccessToken {
	if o == nil {
		var ret []AccessToken
		return ret
	}

	return o.AccessTokens
}

// GetAccessTokensOk returns a tuple with the AccessTokens field value
// and a boolean to check if the value has been set.
func (o *AccessTokensResp) GetAccessTokensOk() ([]AccessToken, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessTokens, true
}

// SetAccessTokens sets field value
func (o *AccessTokensResp) SetAccessTokens(v []AccessToken) {
	o.AccessTokens = v
}

func (o AccessTokensResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokensResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_tokens"] = o.AccessTokens
	return toSerialize, nil
}

func (o *AccessTokensResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_tokens",
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

	varAccessTokensResp := _AccessTokensResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessTokensResp)

	if err != nil {
		return err
	}

	*o = AccessTokensResp(varAccessTokensResp)

	return err
}

type NullableAccessTokensResp struct {
	value *AccessTokensResp
	isSet bool
}

func (v NullableAccessTokensResp) Get() *AccessTokensResp {
	return v.value
}

func (v *NullableAccessTokensResp) Set(val *AccessTokensResp) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokensResp) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokensResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokensResp(val *AccessTokensResp) *NullableAccessTokensResp {
	return &NullableAccessTokensResp{value: val, isSet: true}
}

func (v NullableAccessTokensResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokensResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


