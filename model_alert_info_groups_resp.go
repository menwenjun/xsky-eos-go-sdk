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

// checks if the AlertInfoGroupsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertInfoGroupsResp{}

// AlertInfoGroupsResp struct for AlertInfoGroupsResp
type AlertInfoGroupsResp struct {
	// alert info groups
	AlertInfoGroups []AlertInfoGroupRecord `json:"alert_info_groups"`
}

type _AlertInfoGroupsResp AlertInfoGroupsResp

// NewAlertInfoGroupsResp instantiates a new AlertInfoGroupsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertInfoGroupsResp(alertInfoGroups []AlertInfoGroupRecord) *AlertInfoGroupsResp {
	this := AlertInfoGroupsResp{}
	this.AlertInfoGroups = alertInfoGroups
	return &this
}

// NewAlertInfoGroupsRespWithDefaults instantiates a new AlertInfoGroupsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertInfoGroupsRespWithDefaults() *AlertInfoGroupsResp {
	this := AlertInfoGroupsResp{}
	return &this
}

// GetAlertInfoGroups returns the AlertInfoGroups field value
func (o *AlertInfoGroupsResp) GetAlertInfoGroups() []AlertInfoGroupRecord {
	if o == nil {
		var ret []AlertInfoGroupRecord
		return ret
	}

	return o.AlertInfoGroups
}

// GetAlertInfoGroupsOk returns a tuple with the AlertInfoGroups field value
// and a boolean to check if the value has been set.
func (o *AlertInfoGroupsResp) GetAlertInfoGroupsOk() ([]AlertInfoGroupRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertInfoGroups, true
}

// SetAlertInfoGroups sets field value
func (o *AlertInfoGroupsResp) SetAlertInfoGroups(v []AlertInfoGroupRecord) {
	o.AlertInfoGroups = v
}

func (o AlertInfoGroupsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertInfoGroupsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alert_info_groups"] = o.AlertInfoGroups
	return toSerialize, nil
}

func (o *AlertInfoGroupsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alert_info_groups",
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

	varAlertInfoGroupsResp := _AlertInfoGroupsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlertInfoGroupsResp)

	if err != nil {
		return err
	}

	*o = AlertInfoGroupsResp(varAlertInfoGroupsResp)

	return err
}

type NullableAlertInfoGroupsResp struct {
	value *AlertInfoGroupsResp
	isSet bool
}

func (v NullableAlertInfoGroupsResp) Get() *AlertInfoGroupsResp {
	return v.value
}

func (v *NullableAlertInfoGroupsResp) Set(val *AlertInfoGroupsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertInfoGroupsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertInfoGroupsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertInfoGroupsResp(val *AlertInfoGroupsResp) *NullableAlertInfoGroupsResp {
	return &NullableAlertInfoGroupsResp{value: val, isSet: true}
}

func (v NullableAlertInfoGroupsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertInfoGroupsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


