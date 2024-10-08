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

// checks if the DpBlockBackupPoliciesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpBlockBackupPoliciesResp{}

// DpBlockBackupPoliciesResp struct for DpBlockBackupPoliciesResp
type DpBlockBackupPoliciesResp struct {
	DpBlockBackupPolicies []DpBlockBackupPolicy `json:"dp_block_backup_policies,omitempty"`
}

// NewDpBlockBackupPoliciesResp instantiates a new DpBlockBackupPoliciesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpBlockBackupPoliciesResp() *DpBlockBackupPoliciesResp {
	this := DpBlockBackupPoliciesResp{}
	return &this
}

// NewDpBlockBackupPoliciesRespWithDefaults instantiates a new DpBlockBackupPoliciesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpBlockBackupPoliciesRespWithDefaults() *DpBlockBackupPoliciesResp {
	this := DpBlockBackupPoliciesResp{}
	return &this
}

// GetDpBlockBackupPolicies returns the DpBlockBackupPolicies field value if set, zero value otherwise.
func (o *DpBlockBackupPoliciesResp) GetDpBlockBackupPolicies() []DpBlockBackupPolicy {
	if o == nil || IsNil(o.DpBlockBackupPolicies) {
		var ret []DpBlockBackupPolicy
		return ret
	}
	return o.DpBlockBackupPolicies
}

// GetDpBlockBackupPoliciesOk returns a tuple with the DpBlockBackupPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockBackupPoliciesResp) GetDpBlockBackupPoliciesOk() ([]DpBlockBackupPolicy, bool) {
	if o == nil || IsNil(o.DpBlockBackupPolicies) {
		return nil, false
	}
	return o.DpBlockBackupPolicies, true
}

// HasDpBlockBackupPolicies returns a boolean if a field has been set.
func (o *DpBlockBackupPoliciesResp) HasDpBlockBackupPolicies() bool {
	if o != nil && !IsNil(o.DpBlockBackupPolicies) {
		return true
	}

	return false
}

// SetDpBlockBackupPolicies gets a reference to the given []DpBlockBackupPolicy and assigns it to the DpBlockBackupPolicies field.
func (o *DpBlockBackupPoliciesResp) SetDpBlockBackupPolicies(v []DpBlockBackupPolicy) {
	o.DpBlockBackupPolicies = v
}

func (o DpBlockBackupPoliciesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpBlockBackupPoliciesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DpBlockBackupPolicies) {
		toSerialize["dp_block_backup_policies"] = o.DpBlockBackupPolicies
	}
	return toSerialize, nil
}

type NullableDpBlockBackupPoliciesResp struct {
	value *DpBlockBackupPoliciesResp
	isSet bool
}

func (v NullableDpBlockBackupPoliciesResp) Get() *DpBlockBackupPoliciesResp {
	return v.value
}

func (v *NullableDpBlockBackupPoliciesResp) Set(val *DpBlockBackupPoliciesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDpBlockBackupPoliciesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDpBlockBackupPoliciesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpBlockBackupPoliciesResp(val *DpBlockBackupPoliciesResp) *NullableDpBlockBackupPoliciesResp {
	return &NullableDpBlockBackupPoliciesResp{value: val, isSet: true}
}

func (v NullableDpBlockBackupPoliciesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpBlockBackupPoliciesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


