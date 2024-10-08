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

// checks if the SettingResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingResp{}

// SettingResp struct for SettingResp
type SettingResp struct {
	Settings map[string]map[string]interface{} `json:"settings,omitempty"`
}

// NewSettingResp instantiates a new SettingResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingResp() *SettingResp {
	this := SettingResp{}
	return &this
}

// NewSettingRespWithDefaults instantiates a new SettingResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingRespWithDefaults() *SettingResp {
	this := SettingResp{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SettingResp) GetSettings() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Settings) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingResp) GetSettingsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Settings) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SettingResp) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]map[string]interface{} and assigns it to the Settings field.
func (o *SettingResp) SetSettings(v map[string]map[string]interface{}) {
	o.Settings = v
}

func (o SettingResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableSettingResp struct {
	value *SettingResp
	isSet bool
}

func (v NullableSettingResp) Get() *SettingResp {
	return v.value
}

func (v *NullableSettingResp) Set(val *SettingResp) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingResp) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingResp(val *SettingResp) *NullableSettingResp {
	return &NullableSettingResp{value: val, isSet: true}
}

func (v NullableSettingResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


