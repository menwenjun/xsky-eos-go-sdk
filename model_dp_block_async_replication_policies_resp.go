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

// checks if the DpBlockAsyncReplicationPoliciesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpBlockAsyncReplicationPoliciesResp{}

// DpBlockAsyncReplicationPoliciesResp struct for DpBlockAsyncReplicationPoliciesResp
type DpBlockAsyncReplicationPoliciesResp struct {
	DpBlockAsyncReplicationPolicies []DpBlockAsyncReplicationPolicy `json:"dp_block_async_replication_policies,omitempty"`
}

// NewDpBlockAsyncReplicationPoliciesResp instantiates a new DpBlockAsyncReplicationPoliciesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpBlockAsyncReplicationPoliciesResp() *DpBlockAsyncReplicationPoliciesResp {
	this := DpBlockAsyncReplicationPoliciesResp{}
	return &this
}

// NewDpBlockAsyncReplicationPoliciesRespWithDefaults instantiates a new DpBlockAsyncReplicationPoliciesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpBlockAsyncReplicationPoliciesRespWithDefaults() *DpBlockAsyncReplicationPoliciesResp {
	this := DpBlockAsyncReplicationPoliciesResp{}
	return &this
}

// GetDpBlockAsyncReplicationPolicies returns the DpBlockAsyncReplicationPolicies field value if set, zero value otherwise.
func (o *DpBlockAsyncReplicationPoliciesResp) GetDpBlockAsyncReplicationPolicies() []DpBlockAsyncReplicationPolicy {
	if o == nil || IsNil(o.DpBlockAsyncReplicationPolicies) {
		var ret []DpBlockAsyncReplicationPolicy
		return ret
	}
	return o.DpBlockAsyncReplicationPolicies
}

// GetDpBlockAsyncReplicationPoliciesOk returns a tuple with the DpBlockAsyncReplicationPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockAsyncReplicationPoliciesResp) GetDpBlockAsyncReplicationPoliciesOk() ([]DpBlockAsyncReplicationPolicy, bool) {
	if o == nil || IsNil(o.DpBlockAsyncReplicationPolicies) {
		return nil, false
	}
	return o.DpBlockAsyncReplicationPolicies, true
}

// HasDpBlockAsyncReplicationPolicies returns a boolean if a field has been set.
func (o *DpBlockAsyncReplicationPoliciesResp) HasDpBlockAsyncReplicationPolicies() bool {
	if o != nil && !IsNil(o.DpBlockAsyncReplicationPolicies) {
		return true
	}

	return false
}

// SetDpBlockAsyncReplicationPolicies gets a reference to the given []DpBlockAsyncReplicationPolicy and assigns it to the DpBlockAsyncReplicationPolicies field.
func (o *DpBlockAsyncReplicationPoliciesResp) SetDpBlockAsyncReplicationPolicies(v []DpBlockAsyncReplicationPolicy) {
	o.DpBlockAsyncReplicationPolicies = v
}

func (o DpBlockAsyncReplicationPoliciesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpBlockAsyncReplicationPoliciesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DpBlockAsyncReplicationPolicies) {
		toSerialize["dp_block_async_replication_policies"] = o.DpBlockAsyncReplicationPolicies
	}
	return toSerialize, nil
}

type NullableDpBlockAsyncReplicationPoliciesResp struct {
	value *DpBlockAsyncReplicationPoliciesResp
	isSet bool
}

func (v NullableDpBlockAsyncReplicationPoliciesResp) Get() *DpBlockAsyncReplicationPoliciesResp {
	return v.value
}

func (v *NullableDpBlockAsyncReplicationPoliciesResp) Set(val *DpBlockAsyncReplicationPoliciesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDpBlockAsyncReplicationPoliciesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDpBlockAsyncReplicationPoliciesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpBlockAsyncReplicationPoliciesResp(val *DpBlockAsyncReplicationPoliciesResp) *NullableDpBlockAsyncReplicationPoliciesResp {
	return &NullableDpBlockAsyncReplicationPoliciesResp{value: val, isSet: true}
}

func (v NullableDpBlockAsyncReplicationPoliciesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpBlockAsyncReplicationPoliciesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


