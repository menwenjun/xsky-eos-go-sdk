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

// checks if the AlertStrategyResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertStrategyResp{}

// AlertStrategyResp struct for AlertStrategyResp
type AlertStrategyResp struct {
	AlertStrategy AlertStrategy `json:"alert_strategy"`
}

type _AlertStrategyResp AlertStrategyResp

// NewAlertStrategyResp instantiates a new AlertStrategyResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertStrategyResp(alertStrategy AlertStrategy) *AlertStrategyResp {
	this := AlertStrategyResp{}
	this.AlertStrategy = alertStrategy
	return &this
}

// NewAlertStrategyRespWithDefaults instantiates a new AlertStrategyResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertStrategyRespWithDefaults() *AlertStrategyResp {
	this := AlertStrategyResp{}
	return &this
}

// GetAlertStrategy returns the AlertStrategy field value
func (o *AlertStrategyResp) GetAlertStrategy() AlertStrategy {
	if o == nil {
		var ret AlertStrategy
		return ret
	}

	return o.AlertStrategy
}

// GetAlertStrategyOk returns a tuple with the AlertStrategy field value
// and a boolean to check if the value has been set.
func (o *AlertStrategyResp) GetAlertStrategyOk() (*AlertStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertStrategy, true
}

// SetAlertStrategy sets field value
func (o *AlertStrategyResp) SetAlertStrategy(v AlertStrategy) {
	o.AlertStrategy = v
}

func (o AlertStrategyResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertStrategyResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alert_strategy"] = o.AlertStrategy
	return toSerialize, nil
}

func (o *AlertStrategyResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alert_strategy",
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

	varAlertStrategyResp := _AlertStrategyResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlertStrategyResp)

	if err != nil {
		return err
	}

	*o = AlertStrategyResp(varAlertStrategyResp)

	return err
}

type NullableAlertStrategyResp struct {
	value *AlertStrategyResp
	isSet bool
}

func (v NullableAlertStrategyResp) Get() *AlertStrategyResp {
	return v.value
}

func (v *NullableAlertStrategyResp) Set(val *AlertStrategyResp) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertStrategyResp) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertStrategyResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertStrategyResp(val *AlertStrategyResp) *NullableAlertStrategyResp {
	return &NullableAlertStrategyResp{value: val, isSet: true}
}

func (v NullableAlertStrategyResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertStrategyResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


