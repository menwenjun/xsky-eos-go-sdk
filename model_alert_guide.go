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

// checks if the AlertGuide type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertGuide{}

// AlertGuide struct for AlertGuide
type AlertGuide struct {
	Status *string `json:"status,omitempty"`
}

// NewAlertGuide instantiates a new AlertGuide object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertGuide() *AlertGuide {
	this := AlertGuide{}
	return &this
}

// NewAlertGuideWithDefaults instantiates a new AlertGuide object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertGuideWithDefaults() *AlertGuide {
	this := AlertGuide{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlertGuide) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGuide) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlertGuide) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlertGuide) SetStatus(v string) {
	o.Status = &v
}

func (o AlertGuide) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertGuide) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAlertGuide struct {
	value *AlertGuide
	isSet bool
}

func (v NullableAlertGuide) Get() *AlertGuide {
	return v.value
}

func (v *NullableAlertGuide) Set(val *AlertGuide) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertGuide) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertGuide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertGuide(val *AlertGuide) *NullableAlertGuide {
	return &NullableAlertGuide{value: val, isSet: true}
}

func (v NullableAlertGuide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertGuide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


