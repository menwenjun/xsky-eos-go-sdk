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

// checks if the EmptyPoolInitializeReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmptyPoolInitializeReq{}

// EmptyPoolInitializeReq struct for EmptyPoolInitializeReq
type EmptyPoolInitializeReq struct {
	Pool *EmptyPoolInitializeReqPool `json:"pool,omitempty"`
}

// NewEmptyPoolInitializeReq instantiates a new EmptyPoolInitializeReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmptyPoolInitializeReq() *EmptyPoolInitializeReq {
	this := EmptyPoolInitializeReq{}
	return &this
}

// NewEmptyPoolInitializeReqWithDefaults instantiates a new EmptyPoolInitializeReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmptyPoolInitializeReqWithDefaults() *EmptyPoolInitializeReq {
	this := EmptyPoolInitializeReq{}
	return &this
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *EmptyPoolInitializeReq) GetPool() EmptyPoolInitializeReqPool {
	if o == nil || IsNil(o.Pool) {
		var ret EmptyPoolInitializeReqPool
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmptyPoolInitializeReq) GetPoolOk() (*EmptyPoolInitializeReqPool, bool) {
	if o == nil || IsNil(o.Pool) {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *EmptyPoolInitializeReq) HasPool() bool {
	if o != nil && !IsNil(o.Pool) {
		return true
	}

	return false
}

// SetPool gets a reference to the given EmptyPoolInitializeReqPool and assigns it to the Pool field.
func (o *EmptyPoolInitializeReq) SetPool(v EmptyPoolInitializeReqPool) {
	o.Pool = &v
}

func (o EmptyPoolInitializeReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmptyPoolInitializeReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pool) {
		toSerialize["pool"] = o.Pool
	}
	return toSerialize, nil
}

type NullableEmptyPoolInitializeReq struct {
	value *EmptyPoolInitializeReq
	isSet bool
}

func (v NullableEmptyPoolInitializeReq) Get() *EmptyPoolInitializeReq {
	return v.value
}

func (v *NullableEmptyPoolInitializeReq) Set(val *EmptyPoolInitializeReq) {
	v.value = val
	v.isSet = true
}

func (v NullableEmptyPoolInitializeReq) IsSet() bool {
	return v.isSet
}

func (v *NullableEmptyPoolInitializeReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmptyPoolInitializeReq(val *EmptyPoolInitializeReq) *NullableEmptyPoolInitializeReq {
	return &NullableEmptyPoolInitializeReq{value: val, isSet: true}
}

func (v NullableEmptyPoolInitializeReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmptyPoolInitializeReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


