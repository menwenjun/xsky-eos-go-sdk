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

// checks if the TaskIDResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskIDResp{}

// TaskIDResp struct for TaskIDResp
type TaskIDResp struct {
	TaskId *int64 `json:"task_id,omitempty"`
}

// NewTaskIDResp instantiates a new TaskIDResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskIDResp() *TaskIDResp {
	this := TaskIDResp{}
	return &this
}

// NewTaskIDRespWithDefaults instantiates a new TaskIDResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskIDRespWithDefaults() *TaskIDResp {
	this := TaskIDResp{}
	return &this
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *TaskIDResp) GetTaskId() int64 {
	if o == nil || IsNil(o.TaskId) {
		var ret int64
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskIDResp) GetTaskIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TaskId) {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *TaskIDResp) HasTaskId() bool {
	if o != nil && !IsNil(o.TaskId) {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given int64 and assigns it to the TaskId field.
func (o *TaskIDResp) SetTaskId(v int64) {
	o.TaskId = &v
}

func (o TaskIDResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskIDResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaskId) {
		toSerialize["task_id"] = o.TaskId
	}
	return toSerialize, nil
}

type NullableTaskIDResp struct {
	value *TaskIDResp
	isSet bool
}

func (v NullableTaskIDResp) Get() *TaskIDResp {
	return v.value
}

func (v *NullableTaskIDResp) Set(val *TaskIDResp) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskIDResp) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskIDResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskIDResp(val *TaskIDResp) *NullableTaskIDResp {
	return &NullableTaskIDResp{value: val, isSet: true}
}

func (v NullableTaskIDResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskIDResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


