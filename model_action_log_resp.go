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

// checks if the ActionLogResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionLogResp{}

// ActionLogResp struct for ActionLogResp
type ActionLogResp struct {
	ActionLog ActionLog `json:"action_log"`
}

type _ActionLogResp ActionLogResp

// NewActionLogResp instantiates a new ActionLogResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionLogResp(actionLog ActionLog) *ActionLogResp {
	this := ActionLogResp{}
	this.ActionLog = actionLog
	return &this
}

// NewActionLogRespWithDefaults instantiates a new ActionLogResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionLogRespWithDefaults() *ActionLogResp {
	this := ActionLogResp{}
	return &this
}

// GetActionLog returns the ActionLog field value
func (o *ActionLogResp) GetActionLog() ActionLog {
	if o == nil {
		var ret ActionLog
		return ret
	}

	return o.ActionLog
}

// GetActionLogOk returns a tuple with the ActionLog field value
// and a boolean to check if the value has been set.
func (o *ActionLogResp) GetActionLogOk() (*ActionLog, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionLog, true
}

// SetActionLog sets field value
func (o *ActionLogResp) SetActionLog(v ActionLog) {
	o.ActionLog = v
}

func (o ActionLogResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionLogResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action_log"] = o.ActionLog
	return toSerialize, nil
}

func (o *ActionLogResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action_log",
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

	varActionLogResp := _ActionLogResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActionLogResp)

	if err != nil {
		return err
	}

	*o = ActionLogResp(varActionLogResp)

	return err
}

type NullableActionLogResp struct {
	value *ActionLogResp
	isSet bool
}

func (v NullableActionLogResp) Get() *ActionLogResp {
	return v.value
}

func (v *NullableActionLogResp) Set(val *ActionLogResp) {
	v.value = val
	v.isSet = true
}

func (v NullableActionLogResp) IsSet() bool {
	return v.isSet
}

func (v *NullableActionLogResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionLogResp(val *ActionLogResp) *NullableActionLogResp {
	return &NullableActionLogResp{value: val, isSet: true}
}

func (v NullableActionLogResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionLogResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


