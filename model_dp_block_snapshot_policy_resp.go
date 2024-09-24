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

// checks if the DpBlockSnapshotPolicyResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpBlockSnapshotPolicyResp{}

// DpBlockSnapshotPolicyResp struct for DpBlockSnapshotPolicyResp
type DpBlockSnapshotPolicyResp struct {
	DpBlockSnapshotPolicy *DpBlockSnapshotPolicy `json:"dp_block_snapshot_policy,omitempty"`
}

// NewDpBlockSnapshotPolicyResp instantiates a new DpBlockSnapshotPolicyResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpBlockSnapshotPolicyResp() *DpBlockSnapshotPolicyResp {
	this := DpBlockSnapshotPolicyResp{}
	return &this
}

// NewDpBlockSnapshotPolicyRespWithDefaults instantiates a new DpBlockSnapshotPolicyResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpBlockSnapshotPolicyRespWithDefaults() *DpBlockSnapshotPolicyResp {
	this := DpBlockSnapshotPolicyResp{}
	return &this
}

// GetDpBlockSnapshotPolicy returns the DpBlockSnapshotPolicy field value if set, zero value otherwise.
func (o *DpBlockSnapshotPolicyResp) GetDpBlockSnapshotPolicy() DpBlockSnapshotPolicy {
	if o == nil || IsNil(o.DpBlockSnapshotPolicy) {
		var ret DpBlockSnapshotPolicy
		return ret
	}
	return *o.DpBlockSnapshotPolicy
}

// GetDpBlockSnapshotPolicyOk returns a tuple with the DpBlockSnapshotPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotPolicyResp) GetDpBlockSnapshotPolicyOk() (*DpBlockSnapshotPolicy, bool) {
	if o == nil || IsNil(o.DpBlockSnapshotPolicy) {
		return nil, false
	}
	return o.DpBlockSnapshotPolicy, true
}

// HasDpBlockSnapshotPolicy returns a boolean if a field has been set.
func (o *DpBlockSnapshotPolicyResp) HasDpBlockSnapshotPolicy() bool {
	if o != nil && !IsNil(o.DpBlockSnapshotPolicy) {
		return true
	}

	return false
}

// SetDpBlockSnapshotPolicy gets a reference to the given DpBlockSnapshotPolicy and assigns it to the DpBlockSnapshotPolicy field.
func (o *DpBlockSnapshotPolicyResp) SetDpBlockSnapshotPolicy(v DpBlockSnapshotPolicy) {
	o.DpBlockSnapshotPolicy = &v
}

func (o DpBlockSnapshotPolicyResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpBlockSnapshotPolicyResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DpBlockSnapshotPolicy) {
		toSerialize["dp_block_snapshot_policy"] = o.DpBlockSnapshotPolicy
	}
	return toSerialize, nil
}

type NullableDpBlockSnapshotPolicyResp struct {
	value *DpBlockSnapshotPolicyResp
	isSet bool
}

func (v NullableDpBlockSnapshotPolicyResp) Get() *DpBlockSnapshotPolicyResp {
	return v.value
}

func (v *NullableDpBlockSnapshotPolicyResp) Set(val *DpBlockSnapshotPolicyResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDpBlockSnapshotPolicyResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDpBlockSnapshotPolicyResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpBlockSnapshotPolicyResp(val *DpBlockSnapshotPolicyResp) *NullableDpBlockSnapshotPolicyResp {
	return &NullableDpBlockSnapshotPolicyResp{value: val, isSet: true}
}

func (v NullableDpBlockSnapshotPolicyResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpBlockSnapshotPolicyResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


