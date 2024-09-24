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

// checks if the AlertDingDingConfigUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertDingDingConfigUpdateReq{}

// AlertDingDingConfigUpdateReq struct for AlertDingDingConfigUpdateReq
type AlertDingDingConfigUpdateReq struct {
	AlertDingdingConfig *AlertDingDingConfigUpdateReqAlertDingDingConfig `json:"alert_dingding_config,omitempty"`
}

// NewAlertDingDingConfigUpdateReq instantiates a new AlertDingDingConfigUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertDingDingConfigUpdateReq() *AlertDingDingConfigUpdateReq {
	this := AlertDingDingConfigUpdateReq{}
	return &this
}

// NewAlertDingDingConfigUpdateReqWithDefaults instantiates a new AlertDingDingConfigUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertDingDingConfigUpdateReqWithDefaults() *AlertDingDingConfigUpdateReq {
	this := AlertDingDingConfigUpdateReq{}
	return &this
}

// GetAlertDingdingConfig returns the AlertDingdingConfig field value if set, zero value otherwise.
func (o *AlertDingDingConfigUpdateReq) GetAlertDingdingConfig() AlertDingDingConfigUpdateReqAlertDingDingConfig {
	if o == nil || IsNil(o.AlertDingdingConfig) {
		var ret AlertDingDingConfigUpdateReqAlertDingDingConfig
		return ret
	}
	return *o.AlertDingdingConfig
}

// GetAlertDingdingConfigOk returns a tuple with the AlertDingdingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDingDingConfigUpdateReq) GetAlertDingdingConfigOk() (*AlertDingDingConfigUpdateReqAlertDingDingConfig, bool) {
	if o == nil || IsNil(o.AlertDingdingConfig) {
		return nil, false
	}
	return o.AlertDingdingConfig, true
}

// HasAlertDingdingConfig returns a boolean if a field has been set.
func (o *AlertDingDingConfigUpdateReq) HasAlertDingdingConfig() bool {
	if o != nil && !IsNil(o.AlertDingdingConfig) {
		return true
	}

	return false
}

// SetAlertDingdingConfig gets a reference to the given AlertDingDingConfigUpdateReqAlertDingDingConfig and assigns it to the AlertDingdingConfig field.
func (o *AlertDingDingConfigUpdateReq) SetAlertDingdingConfig(v AlertDingDingConfigUpdateReqAlertDingDingConfig) {
	o.AlertDingdingConfig = &v
}

func (o AlertDingDingConfigUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertDingDingConfigUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlertDingdingConfig) {
		toSerialize["alert_dingding_config"] = o.AlertDingdingConfig
	}
	return toSerialize, nil
}

type NullableAlertDingDingConfigUpdateReq struct {
	value *AlertDingDingConfigUpdateReq
	isSet bool
}

func (v NullableAlertDingDingConfigUpdateReq) Get() *AlertDingDingConfigUpdateReq {
	return v.value
}

func (v *NullableAlertDingDingConfigUpdateReq) Set(val *AlertDingDingConfigUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertDingDingConfigUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertDingDingConfigUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertDingDingConfigUpdateReq(val *AlertDingDingConfigUpdateReq) *NullableAlertDingDingConfigUpdateReq {
	return &NullableAlertDingDingConfigUpdateReq{value: val, isSet: true}
}

func (v NullableAlertDingDingConfigUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertDingDingConfigUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


