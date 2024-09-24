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

// checks if the AlertRuleSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertRuleSchema{}

// AlertRuleSchema AlertRuleSchema defines alert rule schema
type AlertRuleSchema struct {
	Events []string `json:"events,omitempty"`
	Types []string `json:"types,omitempty"`
}

// NewAlertRuleSchema instantiates a new AlertRuleSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertRuleSchema() *AlertRuleSchema {
	this := AlertRuleSchema{}
	return &this
}

// NewAlertRuleSchemaWithDefaults instantiates a new AlertRuleSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertRuleSchemaWithDefaults() *AlertRuleSchema {
	this := AlertRuleSchema{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *AlertRuleSchema) GetEvents() []string {
	if o == nil || IsNil(o.Events) {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRuleSchema) GetEventsOk() ([]string, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *AlertRuleSchema) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *AlertRuleSchema) SetEvents(v []string) {
	o.Events = v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *AlertRuleSchema) GetTypes() []string {
	if o == nil || IsNil(o.Types) {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRuleSchema) GetTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *AlertRuleSchema) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *AlertRuleSchema) SetTypes(v []string) {
	o.Types = v
}

func (o AlertRuleSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertRuleSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	return toSerialize, nil
}

type NullableAlertRuleSchema struct {
	value *AlertRuleSchema
	isSet bool
}

func (v NullableAlertRuleSchema) Get() *AlertRuleSchema {
	return v.value
}

func (v *NullableAlertRuleSchema) Set(val *AlertRuleSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertRuleSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertRuleSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertRuleSchema(val *AlertRuleSchema) *NullableAlertRuleSchema {
	return &NullableAlertRuleSchema{value: val, isSet: true}
}

func (v NullableAlertRuleSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertRuleSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


