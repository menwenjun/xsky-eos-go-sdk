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

// checks if the UserSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSettings{}

// UserSettings struct for UserSettings
type UserSettings struct {
	User UserSettingsUser `json:"user"`
}

type _UserSettings UserSettings

// NewUserSettings instantiates a new UserSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSettings(user UserSettingsUser) *UserSettings {
	this := UserSettings{}
	this.User = user
	return &this
}

// NewUserSettingsWithDefaults instantiates a new UserSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSettingsWithDefaults() *UserSettings {
	this := UserSettings{}
	return &this
}

// GetUser returns the User field value
func (o *UserSettings) GetUser() UserSettingsUser {
	if o == nil {
		var ret UserSettingsUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserSettings) GetUserOk() (*UserSettingsUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserSettings) SetUser(v UserSettingsUser) {
	o.User = v
}

func (o UserSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *UserSettings) UnmarshalJSON(data []byte) (err error) {
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

	varUserSettings := _UserSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSettings)

	if err != nil {
		return err
	}

	*o = UserSettings(varUserSettings)

	return err
}

type NullableUserSettings struct {
	value *UserSettings
	isSet bool
}

func (v NullableUserSettings) Get() *UserSettings {
	return v.value
}

func (v *NullableUserSettings) Set(val *UserSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSettings(val *UserSettings) *NullableUserSettings {
	return &NullableUserSettings{value: val, isSet: true}
}

func (v NullableUserSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


