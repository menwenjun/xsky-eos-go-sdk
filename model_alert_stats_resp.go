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

// checks if the AlertStatsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertStatsResp{}

// AlertStatsResp struct for AlertStatsResp
type AlertStatsResp struct {
	AlertStats AlertStats `json:"alert_stats"`
}

type _AlertStatsResp AlertStatsResp

// NewAlertStatsResp instantiates a new AlertStatsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertStatsResp(alertStats AlertStats) *AlertStatsResp {
	this := AlertStatsResp{}
	this.AlertStats = alertStats
	return &this
}

// NewAlertStatsRespWithDefaults instantiates a new AlertStatsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertStatsRespWithDefaults() *AlertStatsResp {
	this := AlertStatsResp{}
	return &this
}

// GetAlertStats returns the AlertStats field value
func (o *AlertStatsResp) GetAlertStats() AlertStats {
	if o == nil {
		var ret AlertStats
		return ret
	}

	return o.AlertStats
}

// GetAlertStatsOk returns a tuple with the AlertStats field value
// and a boolean to check if the value has been set.
func (o *AlertStatsResp) GetAlertStatsOk() (*AlertStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertStats, true
}

// SetAlertStats sets field value
func (o *AlertStatsResp) SetAlertStats(v AlertStats) {
	o.AlertStats = v
}

func (o AlertStatsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertStatsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alert_stats"] = o.AlertStats
	return toSerialize, nil
}

func (o *AlertStatsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alert_stats",
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

	varAlertStatsResp := _AlertStatsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlertStatsResp)

	if err != nil {
		return err
	}

	*o = AlertStatsResp(varAlertStatsResp)

	return err
}

type NullableAlertStatsResp struct {
	value *AlertStatsResp
	isSet bool
}

func (v NullableAlertStatsResp) Get() *AlertStatsResp {
	return v.value
}

func (v *NullableAlertStatsResp) Set(val *AlertStatsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertStatsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertStatsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertStatsResp(val *AlertStatsResp) *NullableAlertStatsResp {
	return &NullableAlertStatsResp{value: val, isSet: true}
}

func (v NullableAlertStatsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertStatsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


