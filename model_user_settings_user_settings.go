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

// checks if the UserSettingsUserSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSettingsUserSettings{}

// UserSettingsUserSettings struct for UserSettingsUserSettings
type UserSettingsUserSettings struct {
	// dashboard setting
	Dashboard *string `json:"dashboard,omitempty"`
}

// NewUserSettingsUserSettings instantiates a new UserSettingsUserSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSettingsUserSettings() *UserSettingsUserSettings {
	this := UserSettingsUserSettings{}
	return &this
}

// NewUserSettingsUserSettingsWithDefaults instantiates a new UserSettingsUserSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSettingsUserSettingsWithDefaults() *UserSettingsUserSettings {
	this := UserSettingsUserSettings{}
	return &this
}

// GetDashboard returns the Dashboard field value if set, zero value otherwise.
func (o *UserSettingsUserSettings) GetDashboard() string {
	if o == nil || IsNil(o.Dashboard) {
		var ret string
		return ret
	}
	return *o.Dashboard
}

// GetDashboardOk returns a tuple with the Dashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsUserSettings) GetDashboardOk() (*string, bool) {
	if o == nil || IsNil(o.Dashboard) {
		return nil, false
	}
	return o.Dashboard, true
}

// HasDashboard returns a boolean if a field has been set.
func (o *UserSettingsUserSettings) HasDashboard() bool {
	if o != nil && !IsNil(o.Dashboard) {
		return true
	}

	return false
}

// SetDashboard gets a reference to the given string and assigns it to the Dashboard field.
func (o *UserSettingsUserSettings) SetDashboard(v string) {
	o.Dashboard = &v
}

func (o UserSettingsUserSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSettingsUserSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dashboard) {
		toSerialize["dashboard"] = o.Dashboard
	}
	return toSerialize, nil
}

type NullableUserSettingsUserSettings struct {
	value *UserSettingsUserSettings
	isSet bool
}

func (v NullableUserSettingsUserSettings) Get() *UserSettingsUserSettings {
	return v.value
}

func (v *NullableUserSettingsUserSettings) Set(val *UserSettingsUserSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSettingsUserSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSettingsUserSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSettingsUserSettings(val *UserSettingsUserSettings) *NullableUserSettingsUserSettings {
	return &NullableUserSettingsUserSettings{value: val, isSet: true}
}

func (v NullableUserSettingsUserSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSettingsUserSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


