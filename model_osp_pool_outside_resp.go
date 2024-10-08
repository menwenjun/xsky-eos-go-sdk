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

// checks if the OspPoolOutsideResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OspPoolOutsideResp{}

// OspPoolOutsideResp struct for OspPoolOutsideResp
type OspPoolOutsideResp struct {
	PoolIds []int64 `json:"pool_ids,omitempty"`
}

// NewOspPoolOutsideResp instantiates a new OspPoolOutsideResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOspPoolOutsideResp() *OspPoolOutsideResp {
	this := OspPoolOutsideResp{}
	return &this
}

// NewOspPoolOutsideRespWithDefaults instantiates a new OspPoolOutsideResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOspPoolOutsideRespWithDefaults() *OspPoolOutsideResp {
	this := OspPoolOutsideResp{}
	return &this
}

// GetPoolIds returns the PoolIds field value if set, zero value otherwise.
func (o *OspPoolOutsideResp) GetPoolIds() []int64 {
	if o == nil || IsNil(o.PoolIds) {
		var ret []int64
		return ret
	}
	return o.PoolIds
}

// GetPoolIdsOk returns a tuple with the PoolIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspPoolOutsideResp) GetPoolIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.PoolIds) {
		return nil, false
	}
	return o.PoolIds, true
}

// HasPoolIds returns a boolean if a field has been set.
func (o *OspPoolOutsideResp) HasPoolIds() bool {
	if o != nil && !IsNil(o.PoolIds) {
		return true
	}

	return false
}

// SetPoolIds gets a reference to the given []int64 and assigns it to the PoolIds field.
func (o *OspPoolOutsideResp) SetPoolIds(v []int64) {
	o.PoolIds = v
}

func (o OspPoolOutsideResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OspPoolOutsideResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PoolIds) {
		toSerialize["pool_ids"] = o.PoolIds
	}
	return toSerialize, nil
}

type NullableOspPoolOutsideResp struct {
	value *OspPoolOutsideResp
	isSet bool
}

func (v NullableOspPoolOutsideResp) Get() *OspPoolOutsideResp {
	return v.value
}

func (v *NullableOspPoolOutsideResp) Set(val *OspPoolOutsideResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOspPoolOutsideResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOspPoolOutsideResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspPoolOutsideResp(val *OspPoolOutsideResp) *NullableOspPoolOutsideResp {
	return &NullableOspPoolOutsideResp{value: val, isSet: true}
}

func (v NullableOspPoolOutsideResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspPoolOutsideResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


