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

// checks if the UpdateAlertRuleResourceBlacklistReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAlertRuleResourceBlacklistReq{}

// UpdateAlertRuleResourceBlacklistReq struct for UpdateAlertRuleResourceBlacklistReq
type UpdateAlertRuleResourceBlacklistReq struct {
	AlertRuleResourceBlacklist UpdateAlertRuleResourceBlacklistReqAlertRuleResourceBlacklist `json:"alert_rule_resource_blacklist"`
}

type _UpdateAlertRuleResourceBlacklistReq UpdateAlertRuleResourceBlacklistReq

// NewUpdateAlertRuleResourceBlacklistReq instantiates a new UpdateAlertRuleResourceBlacklistReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAlertRuleResourceBlacklistReq(alertRuleResourceBlacklist UpdateAlertRuleResourceBlacklistReqAlertRuleResourceBlacklist) *UpdateAlertRuleResourceBlacklistReq {
	this := UpdateAlertRuleResourceBlacklistReq{}
	this.AlertRuleResourceBlacklist = alertRuleResourceBlacklist
	return &this
}

// NewUpdateAlertRuleResourceBlacklistReqWithDefaults instantiates a new UpdateAlertRuleResourceBlacklistReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAlertRuleResourceBlacklistReqWithDefaults() *UpdateAlertRuleResourceBlacklistReq {
	this := UpdateAlertRuleResourceBlacklistReq{}
	return &this
}

// GetAlertRuleResourceBlacklist returns the AlertRuleResourceBlacklist field value
func (o *UpdateAlertRuleResourceBlacklistReq) GetAlertRuleResourceBlacklist() UpdateAlertRuleResourceBlacklistReqAlertRuleResourceBlacklist {
	if o == nil {
		var ret UpdateAlertRuleResourceBlacklistReqAlertRuleResourceBlacklist
		return ret
	}

	return o.AlertRuleResourceBlacklist
}

// GetAlertRuleResourceBlacklistOk returns a tuple with the AlertRuleResourceBlacklist field value
// and a boolean to check if the value has been set.
func (o *UpdateAlertRuleResourceBlacklistReq) GetAlertRuleResourceBlacklistOk() (*UpdateAlertRuleResourceBlacklistReqAlertRuleResourceBlacklist, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertRuleResourceBlacklist, true
}

// SetAlertRuleResourceBlacklist sets field value
func (o *UpdateAlertRuleResourceBlacklistReq) SetAlertRuleResourceBlacklist(v UpdateAlertRuleResourceBlacklistReqAlertRuleResourceBlacklist) {
	o.AlertRuleResourceBlacklist = v
}

func (o UpdateAlertRuleResourceBlacklistReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAlertRuleResourceBlacklistReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alert_rule_resource_blacklist"] = o.AlertRuleResourceBlacklist
	return toSerialize, nil
}

func (o *UpdateAlertRuleResourceBlacklistReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alert_rule_resource_blacklist",
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

	varUpdateAlertRuleResourceBlacklistReq := _UpdateAlertRuleResourceBlacklistReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateAlertRuleResourceBlacklistReq)

	if err != nil {
		return err
	}

	*o = UpdateAlertRuleResourceBlacklistReq(varUpdateAlertRuleResourceBlacklistReq)

	return err
}

type NullableUpdateAlertRuleResourceBlacklistReq struct {
	value *UpdateAlertRuleResourceBlacklistReq
	isSet bool
}

func (v NullableUpdateAlertRuleResourceBlacklistReq) Get() *UpdateAlertRuleResourceBlacklistReq {
	return v.value
}

func (v *NullableUpdateAlertRuleResourceBlacklistReq) Set(val *UpdateAlertRuleResourceBlacklistReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAlertRuleResourceBlacklistReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAlertRuleResourceBlacklistReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAlertRuleResourceBlacklistReq(val *UpdateAlertRuleResourceBlacklistReq) *NullableUpdateAlertRuleResourceBlacklistReq {
	return &NullableUpdateAlertRuleResourceBlacklistReq{value: val, isSet: true}
}

func (v NullableUpdateAlertRuleResourceBlacklistReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAlertRuleResourceBlacklistReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


