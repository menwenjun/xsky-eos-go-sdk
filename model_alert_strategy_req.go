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

// checks if the AlertStrategyReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertStrategyReq{}

// AlertStrategyReq struct for AlertStrategyReq
type AlertStrategyReq struct {
	AlertStrategy AlertStrategyReqAlertStrategy `json:"alert_strategy"`
}

type _AlertStrategyReq AlertStrategyReq

// NewAlertStrategyReq instantiates a new AlertStrategyReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertStrategyReq(alertStrategy AlertStrategyReqAlertStrategy) *AlertStrategyReq {
	this := AlertStrategyReq{}
	this.AlertStrategy = alertStrategy
	return &this
}

// NewAlertStrategyReqWithDefaults instantiates a new AlertStrategyReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertStrategyReqWithDefaults() *AlertStrategyReq {
	this := AlertStrategyReq{}
	return &this
}

// GetAlertStrategy returns the AlertStrategy field value
func (o *AlertStrategyReq) GetAlertStrategy() AlertStrategyReqAlertStrategy {
	if o == nil {
		var ret AlertStrategyReqAlertStrategy
		return ret
	}

	return o.AlertStrategy
}

// GetAlertStrategyOk returns a tuple with the AlertStrategy field value
// and a boolean to check if the value has been set.
func (o *AlertStrategyReq) GetAlertStrategyOk() (*AlertStrategyReqAlertStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertStrategy, true
}

// SetAlertStrategy sets field value
func (o *AlertStrategyReq) SetAlertStrategy(v AlertStrategyReqAlertStrategy) {
	o.AlertStrategy = v
}

func (o AlertStrategyReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertStrategyReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alert_strategy"] = o.AlertStrategy
	return toSerialize, nil
}

func (o *AlertStrategyReq) UnmarshalJSON(data []byte) (err error) {
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

	varAlertStrategyReq := _AlertStrategyReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlertStrategyReq)

	if err != nil {
		return err
	}

	*o = AlertStrategyReq(varAlertStrategyReq)

	return err
}

type NullableAlertStrategyReq struct {
	value *AlertStrategyReq
	isSet bool
}

func (v NullableAlertStrategyReq) Get() *AlertStrategyReq {
	return v.value
}

func (v *NullableAlertStrategyReq) Set(val *AlertStrategyReq) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertStrategyReq) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertStrategyReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertStrategyReq(val *AlertStrategyReq) *NullableAlertStrategyReq {
	return &NullableAlertStrategyReq{value: val, isSet: true}
}

func (v NullableAlertStrategyReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertStrategyReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


