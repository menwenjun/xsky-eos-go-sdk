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

// checks if the AccessTokenUpdateReqAccessToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenUpdateReqAccessToken{}

// AccessTokenUpdateReqAccessToken struct for AccessTokenUpdateReqAccessToken
type AccessTokenUpdateReqAccessToken struct {
	// name of access token
	Name string `json:"name"`
}

type _AccessTokenUpdateReqAccessToken AccessTokenUpdateReqAccessToken

// NewAccessTokenUpdateReqAccessToken instantiates a new AccessTokenUpdateReqAccessToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenUpdateReqAccessToken(name string) *AccessTokenUpdateReqAccessToken {
	this := AccessTokenUpdateReqAccessToken{}
	this.Name = name
	return &this
}

// NewAccessTokenUpdateReqAccessTokenWithDefaults instantiates a new AccessTokenUpdateReqAccessToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenUpdateReqAccessTokenWithDefaults() *AccessTokenUpdateReqAccessToken {
	this := AccessTokenUpdateReqAccessToken{}
	return &this
}

// GetName returns the Name field value
func (o *AccessTokenUpdateReqAccessToken) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccessTokenUpdateReqAccessToken) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccessTokenUpdateReqAccessToken) SetName(v string) {
	o.Name = v
}

func (o AccessTokenUpdateReqAccessToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokenUpdateReqAccessToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *AccessTokenUpdateReqAccessToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varAccessTokenUpdateReqAccessToken := _AccessTokenUpdateReqAccessToken{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessTokenUpdateReqAccessToken)

	if err != nil {
		return err
	}

	*o = AccessTokenUpdateReqAccessToken(varAccessTokenUpdateReqAccessToken)

	return err
}

type NullableAccessTokenUpdateReqAccessToken struct {
	value *AccessTokenUpdateReqAccessToken
	isSet bool
}

func (v NullableAccessTokenUpdateReqAccessToken) Get() *AccessTokenUpdateReqAccessToken {
	return v.value
}

func (v *NullableAccessTokenUpdateReqAccessToken) Set(val *AccessTokenUpdateReqAccessToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenUpdateReqAccessToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenUpdateReqAccessToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenUpdateReqAccessToken(val *AccessTokenUpdateReqAccessToken) *NullableAccessTokenUpdateReqAccessToken {
	return &NullableAccessTokenUpdateReqAccessToken{value: val, isSet: true}
}

func (v NullableAccessTokenUpdateReqAccessToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenUpdateReqAccessToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


