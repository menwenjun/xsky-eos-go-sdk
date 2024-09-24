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

// checks if the AlertRuleGroupDetailUpdateReqGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertRuleGroupDetailUpdateReqGroup{}

// AlertRuleGroupDetailUpdateReqGroup struct for AlertRuleGroupDetailUpdateReqGroup
type AlertRuleGroupDetailUpdateReqGroup struct {
	// auto resolved or not
	AutoResolve *bool `json:"auto_resolve,omitempty"`
	// enable the alert rule group or not
	Enabled *bool `json:"enabled,omitempty"`
	// trigger time in seconds
	TriggerTimeSeconds *int64 `json:"trigger_time_seconds,omitempty"`
}

// NewAlertRuleGroupDetailUpdateReqGroup instantiates a new AlertRuleGroupDetailUpdateReqGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertRuleGroupDetailUpdateReqGroup() *AlertRuleGroupDetailUpdateReqGroup {
	this := AlertRuleGroupDetailUpdateReqGroup{}
	return &this
}

// NewAlertRuleGroupDetailUpdateReqGroupWithDefaults instantiates a new AlertRuleGroupDetailUpdateReqGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertRuleGroupDetailUpdateReqGroupWithDefaults() *AlertRuleGroupDetailUpdateReqGroup {
	this := AlertRuleGroupDetailUpdateReqGroup{}
	return &this
}

// GetAutoResolve returns the AutoResolve field value if set, zero value otherwise.
func (o *AlertRuleGroupDetailUpdateReqGroup) GetAutoResolve() bool {
	if o == nil || IsNil(o.AutoResolve) {
		var ret bool
		return ret
	}
	return *o.AutoResolve
}

// GetAutoResolveOk returns a tuple with the AutoResolve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRuleGroupDetailUpdateReqGroup) GetAutoResolveOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoResolve) {
		return nil, false
	}
	return o.AutoResolve, true
}

// HasAutoResolve returns a boolean if a field has been set.
func (o *AlertRuleGroupDetailUpdateReqGroup) HasAutoResolve() bool {
	if o != nil && !IsNil(o.AutoResolve) {
		return true
	}

	return false
}

// SetAutoResolve gets a reference to the given bool and assigns it to the AutoResolve field.
func (o *AlertRuleGroupDetailUpdateReqGroup) SetAutoResolve(v bool) {
	o.AutoResolve = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AlertRuleGroupDetailUpdateReqGroup) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRuleGroupDetailUpdateReqGroup) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AlertRuleGroupDetailUpdateReqGroup) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AlertRuleGroupDetailUpdateReqGroup) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTriggerTimeSeconds returns the TriggerTimeSeconds field value if set, zero value otherwise.
func (o *AlertRuleGroupDetailUpdateReqGroup) GetTriggerTimeSeconds() int64 {
	if o == nil || IsNil(o.TriggerTimeSeconds) {
		var ret int64
		return ret
	}
	return *o.TriggerTimeSeconds
}

// GetTriggerTimeSecondsOk returns a tuple with the TriggerTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRuleGroupDetailUpdateReqGroup) GetTriggerTimeSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.TriggerTimeSeconds) {
		return nil, false
	}
	return o.TriggerTimeSeconds, true
}

// HasTriggerTimeSeconds returns a boolean if a field has been set.
func (o *AlertRuleGroupDetailUpdateReqGroup) HasTriggerTimeSeconds() bool {
	if o != nil && !IsNil(o.TriggerTimeSeconds) {
		return true
	}

	return false
}

// SetTriggerTimeSeconds gets a reference to the given int64 and assigns it to the TriggerTimeSeconds field.
func (o *AlertRuleGroupDetailUpdateReqGroup) SetTriggerTimeSeconds(v int64) {
	o.TriggerTimeSeconds = &v
}

func (o AlertRuleGroupDetailUpdateReqGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertRuleGroupDetailUpdateReqGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoResolve) {
		toSerialize["auto_resolve"] = o.AutoResolve
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.TriggerTimeSeconds) {
		toSerialize["trigger_time_seconds"] = o.TriggerTimeSeconds
	}
	return toSerialize, nil
}

type NullableAlertRuleGroupDetailUpdateReqGroup struct {
	value *AlertRuleGroupDetailUpdateReqGroup
	isSet bool
}

func (v NullableAlertRuleGroupDetailUpdateReqGroup) Get() *AlertRuleGroupDetailUpdateReqGroup {
	return v.value
}

func (v *NullableAlertRuleGroupDetailUpdateReqGroup) Set(val *AlertRuleGroupDetailUpdateReqGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertRuleGroupDetailUpdateReqGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertRuleGroupDetailUpdateReqGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertRuleGroupDetailUpdateReqGroup(val *AlertRuleGroupDetailUpdateReqGroup) *NullableAlertRuleGroupDetailUpdateReqGroup {
	return &NullableAlertRuleGroupDetailUpdateReqGroup{value: val, isSet: true}
}

func (v NullableAlertRuleGroupDetailUpdateReqGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertRuleGroupDetailUpdateReqGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


