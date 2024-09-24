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

// checks if the AlertWechatSendResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertWechatSendResp{}

// AlertWechatSendResp struct for AlertWechatSendResp
type AlertWechatSendResp struct {
	Code *int64 `json:"Code,omitempty"`
}

// NewAlertWechatSendResp instantiates a new AlertWechatSendResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertWechatSendResp() *AlertWechatSendResp {
	this := AlertWechatSendResp{}
	return &this
}

// NewAlertWechatSendRespWithDefaults instantiates a new AlertWechatSendResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertWechatSendRespWithDefaults() *AlertWechatSendResp {
	this := AlertWechatSendResp{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AlertWechatSendResp) GetCode() int64 {
	if o == nil || IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertWechatSendResp) GetCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AlertWechatSendResp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *AlertWechatSendResp) SetCode(v int64) {
	o.Code = &v
}

func (o AlertWechatSendResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertWechatSendResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["Code"] = o.Code
	}
	return toSerialize, nil
}

type NullableAlertWechatSendResp struct {
	value *AlertWechatSendResp
	isSet bool
}

func (v NullableAlertWechatSendResp) Get() *AlertWechatSendResp {
	return v.value
}

func (v *NullableAlertWechatSendResp) Set(val *AlertWechatSendResp) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertWechatSendResp) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertWechatSendResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertWechatSendResp(val *AlertWechatSendResp) *NullableAlertWechatSendResp {
	return &NullableAlertWechatSendResp{value: val, isSet: true}
}

func (v NullableAlertWechatSendResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertWechatSendResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


