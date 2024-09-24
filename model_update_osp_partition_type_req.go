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

// checks if the UpdateOspPartitionTypeReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOspPartitionTypeReq{}

// UpdateOspPartitionTypeReq struct for UpdateOspPartitionTypeReq
type UpdateOspPartitionTypeReq struct {
	// partition type
	Type *string `json:"type,omitempty"`
}

// NewUpdateOspPartitionTypeReq instantiates a new UpdateOspPartitionTypeReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOspPartitionTypeReq() *UpdateOspPartitionTypeReq {
	this := UpdateOspPartitionTypeReq{}
	return &this
}

// NewUpdateOspPartitionTypeReqWithDefaults instantiates a new UpdateOspPartitionTypeReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOspPartitionTypeReqWithDefaults() *UpdateOspPartitionTypeReq {
	this := UpdateOspPartitionTypeReq{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateOspPartitionTypeReq) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOspPartitionTypeReq) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateOspPartitionTypeReq) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateOspPartitionTypeReq) SetType(v string) {
	o.Type = &v
}

func (o UpdateOspPartitionTypeReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOspPartitionTypeReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUpdateOspPartitionTypeReq struct {
	value *UpdateOspPartitionTypeReq
	isSet bool
}

func (v NullableUpdateOspPartitionTypeReq) Get() *UpdateOspPartitionTypeReq {
	return v.value
}

func (v *NullableUpdateOspPartitionTypeReq) Set(val *UpdateOspPartitionTypeReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOspPartitionTypeReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOspPartitionTypeReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOspPartitionTypeReq(val *UpdateOspPartitionTypeReq) *NullableUpdateOspPartitionTypeReq {
	return &NullableUpdateOspPartitionTypeReq{value: val, isSet: true}
}

func (v NullableUpdateOspPartitionTypeReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOspPartitionTypeReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


