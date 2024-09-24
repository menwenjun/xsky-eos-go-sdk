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

// checks if the UserUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUpdateReq{}

// UserUpdateReq struct for UserUpdateReq
type UserUpdateReq struct {
	User UserUpdateReqUser `json:"user"`
}

type _UserUpdateReq UserUpdateReq

// NewUserUpdateReq instantiates a new UserUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUpdateReq(user UserUpdateReqUser) *UserUpdateReq {
	this := UserUpdateReq{}
	this.User = user
	return &this
}

// NewUserUpdateReqWithDefaults instantiates a new UserUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUpdateReqWithDefaults() *UserUpdateReq {
	this := UserUpdateReq{}
	return &this
}

// GetUser returns the User field value
func (o *UserUpdateReq) GetUser() UserUpdateReqUser {
	if o == nil {
		var ret UserUpdateReqUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserUpdateReq) GetUserOk() (*UserUpdateReqUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserUpdateReq) SetUser(v UserUpdateReqUser) {
	o.User = v
}

func (o UserUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *UserUpdateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user",
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

	varUserUpdateReq := _UserUpdateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserUpdateReq)

	if err != nil {
		return err
	}

	*o = UserUpdateReq(varUserUpdateReq)

	return err
}

type NullableUserUpdateReq struct {
	value *UserUpdateReq
	isSet bool
}

func (v NullableUserUpdateReq) Get() *UserUpdateReq {
	return v.value
}

func (v *NullableUserUpdateReq) Set(val *UserUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUpdateReq(val *UserUpdateReq) *NullableUserUpdateReq {
	return &NullableUserUpdateReq{value: val, isSet: true}
}

func (v NullableUserUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


