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

// checks if the AlertInfoNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertInfoNestview{}

// AlertInfoNestview struct for AlertInfoNestview
type AlertInfoNestview struct {
	AlarmId *string `json:"alarm_id,omitempty"`
	Id *int64 `json:"id,omitempty"`
}

// NewAlertInfoNestview instantiates a new AlertInfoNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertInfoNestview() *AlertInfoNestview {
	this := AlertInfoNestview{}
	return &this
}

// NewAlertInfoNestviewWithDefaults instantiates a new AlertInfoNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertInfoNestviewWithDefaults() *AlertInfoNestview {
	this := AlertInfoNestview{}
	return &this
}

// GetAlarmId returns the AlarmId field value if set, zero value otherwise.
func (o *AlertInfoNestview) GetAlarmId() string {
	if o == nil || IsNil(o.AlarmId) {
		var ret string
		return ret
	}
	return *o.AlarmId
}

// GetAlarmIdOk returns a tuple with the AlarmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertInfoNestview) GetAlarmIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlarmId) {
		return nil, false
	}
	return o.AlarmId, true
}

// HasAlarmId returns a boolean if a field has been set.
func (o *AlertInfoNestview) HasAlarmId() bool {
	if o != nil && !IsNil(o.AlarmId) {
		return true
	}

	return false
}

// SetAlarmId gets a reference to the given string and assigns it to the AlarmId field.
func (o *AlertInfoNestview) SetAlarmId(v string) {
	o.AlarmId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertInfoNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertInfoNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertInfoNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AlertInfoNestview) SetId(v int64) {
	o.Id = &v
}

func (o AlertInfoNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertInfoNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlarmId) {
		toSerialize["alarm_id"] = o.AlarmId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAlertInfoNestview struct {
	value *AlertInfoNestview
	isSet bool
}

func (v NullableAlertInfoNestview) Get() *AlertInfoNestview {
	return v.value
}

func (v *NullableAlertInfoNestview) Set(val *AlertInfoNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertInfoNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertInfoNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertInfoNestview(val *AlertInfoNestview) *NullableAlertInfoNestview {
	return &NullableAlertInfoNestview{value: val, isSet: true}
}

func (v NullableAlertInfoNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertInfoNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


