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

// checks if the AlertWechatConfigResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertWechatConfigResp{}

// AlertWechatConfigResp struct for AlertWechatConfigResp
type AlertWechatConfigResp struct {
	AlertWechatConfig *AlertWechatConfig `json:"alert_wechat_config,omitempty"`
}

// NewAlertWechatConfigResp instantiates a new AlertWechatConfigResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertWechatConfigResp() *AlertWechatConfigResp {
	this := AlertWechatConfigResp{}
	return &this
}

// NewAlertWechatConfigRespWithDefaults instantiates a new AlertWechatConfigResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertWechatConfigRespWithDefaults() *AlertWechatConfigResp {
	this := AlertWechatConfigResp{}
	return &this
}

// GetAlertWechatConfig returns the AlertWechatConfig field value if set, zero value otherwise.
func (o *AlertWechatConfigResp) GetAlertWechatConfig() AlertWechatConfig {
	if o == nil || IsNil(o.AlertWechatConfig) {
		var ret AlertWechatConfig
		return ret
	}
	return *o.AlertWechatConfig
}

// GetAlertWechatConfigOk returns a tuple with the AlertWechatConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertWechatConfigResp) GetAlertWechatConfigOk() (*AlertWechatConfig, bool) {
	if o == nil || IsNil(o.AlertWechatConfig) {
		return nil, false
	}
	return o.AlertWechatConfig, true
}

// HasAlertWechatConfig returns a boolean if a field has been set.
func (o *AlertWechatConfigResp) HasAlertWechatConfig() bool {
	if o != nil && !IsNil(o.AlertWechatConfig) {
		return true
	}

	return false
}

// SetAlertWechatConfig gets a reference to the given AlertWechatConfig and assigns it to the AlertWechatConfig field.
func (o *AlertWechatConfigResp) SetAlertWechatConfig(v AlertWechatConfig) {
	o.AlertWechatConfig = &v
}

func (o AlertWechatConfigResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertWechatConfigResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlertWechatConfig) {
		toSerialize["alert_wechat_config"] = o.AlertWechatConfig
	}
	return toSerialize, nil
}

type NullableAlertWechatConfigResp struct {
	value *AlertWechatConfigResp
	isSet bool
}

func (v NullableAlertWechatConfigResp) Get() *AlertWechatConfigResp {
	return v.value
}

func (v *NullableAlertWechatConfigResp) Set(val *AlertWechatConfigResp) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertWechatConfigResp) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertWechatConfigResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertWechatConfigResp(val *AlertWechatConfigResp) *NullableAlertWechatConfigResp {
	return &NullableAlertWechatConfigResp{value: val, isSet: true}
}

func (v NullableAlertWechatConfigResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertWechatConfigResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


