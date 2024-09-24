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

// checks if the ActionLogCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionLogCreateReq{}

// ActionLogCreateReq struct for ActionLogCreateReq
type ActionLogCreateReq struct {
	ActionLog *ActionLogCreateReqActionLog `json:"action_log,omitempty"`
}

// NewActionLogCreateReq instantiates a new ActionLogCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionLogCreateReq() *ActionLogCreateReq {
	this := ActionLogCreateReq{}
	return &this
}

// NewActionLogCreateReqWithDefaults instantiates a new ActionLogCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionLogCreateReqWithDefaults() *ActionLogCreateReq {
	this := ActionLogCreateReq{}
	return &this
}

// GetActionLog returns the ActionLog field value if set, zero value otherwise.
func (o *ActionLogCreateReq) GetActionLog() ActionLogCreateReqActionLog {
	if o == nil || IsNil(o.ActionLog) {
		var ret ActionLogCreateReqActionLog
		return ret
	}
	return *o.ActionLog
}

// GetActionLogOk returns a tuple with the ActionLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionLogCreateReq) GetActionLogOk() (*ActionLogCreateReqActionLog, bool) {
	if o == nil || IsNil(o.ActionLog) {
		return nil, false
	}
	return o.ActionLog, true
}

// HasActionLog returns a boolean if a field has been set.
func (o *ActionLogCreateReq) HasActionLog() bool {
	if o != nil && !IsNil(o.ActionLog) {
		return true
	}

	return false
}

// SetActionLog gets a reference to the given ActionLogCreateReqActionLog and assigns it to the ActionLog field.
func (o *ActionLogCreateReq) SetActionLog(v ActionLogCreateReqActionLog) {
	o.ActionLog = &v
}

func (o ActionLogCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionLogCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionLog) {
		toSerialize["action_log"] = o.ActionLog
	}
	return toSerialize, nil
}

type NullableActionLogCreateReq struct {
	value *ActionLogCreateReq
	isSet bool
}

func (v NullableActionLogCreateReq) Get() *ActionLogCreateReq {
	return v.value
}

func (v *NullableActionLogCreateReq) Set(val *ActionLogCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableActionLogCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableActionLogCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionLogCreateReq(val *ActionLogCreateReq) *NullableActionLogCreateReq {
	return &NullableActionLogCreateReq{value: val, isSet: true}
}

func (v NullableActionLogCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionLogCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


