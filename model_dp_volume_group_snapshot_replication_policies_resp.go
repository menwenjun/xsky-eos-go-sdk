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

// checks if the DpVolumeGroupSnapshotReplicationPoliciesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpVolumeGroupSnapshotReplicationPoliciesResp{}

// DpVolumeGroupSnapshotReplicationPoliciesResp struct for DpVolumeGroupSnapshotReplicationPoliciesResp
type DpVolumeGroupSnapshotReplicationPoliciesResp struct {
	DpVolumeGroupSnapshotReplicationPolicies []DpVolumeGroupSnapshotReplicationPolicy `json:"dp_volume_group_snapshot_replication_policies,omitempty"`
}

// NewDpVolumeGroupSnapshotReplicationPoliciesResp instantiates a new DpVolumeGroupSnapshotReplicationPoliciesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpVolumeGroupSnapshotReplicationPoliciesResp() *DpVolumeGroupSnapshotReplicationPoliciesResp {
	this := DpVolumeGroupSnapshotReplicationPoliciesResp{}
	return &this
}

// NewDpVolumeGroupSnapshotReplicationPoliciesRespWithDefaults instantiates a new DpVolumeGroupSnapshotReplicationPoliciesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpVolumeGroupSnapshotReplicationPoliciesRespWithDefaults() *DpVolumeGroupSnapshotReplicationPoliciesResp {
	this := DpVolumeGroupSnapshotReplicationPoliciesResp{}
	return &this
}

// GetDpVolumeGroupSnapshotReplicationPolicies returns the DpVolumeGroupSnapshotReplicationPolicies field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationPoliciesResp) GetDpVolumeGroupSnapshotReplicationPolicies() []DpVolumeGroupSnapshotReplicationPolicy {
	if o == nil || IsNil(o.DpVolumeGroupSnapshotReplicationPolicies) {
		var ret []DpVolumeGroupSnapshotReplicationPolicy
		return ret
	}
	return o.DpVolumeGroupSnapshotReplicationPolicies
}

// GetDpVolumeGroupSnapshotReplicationPoliciesOk returns a tuple with the DpVolumeGroupSnapshotReplicationPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationPoliciesResp) GetDpVolumeGroupSnapshotReplicationPoliciesOk() ([]DpVolumeGroupSnapshotReplicationPolicy, bool) {
	if o == nil || IsNil(o.DpVolumeGroupSnapshotReplicationPolicies) {
		return nil, false
	}
	return o.DpVolumeGroupSnapshotReplicationPolicies, true
}

// HasDpVolumeGroupSnapshotReplicationPolicies returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationPoliciesResp) HasDpVolumeGroupSnapshotReplicationPolicies() bool {
	if o != nil && !IsNil(o.DpVolumeGroupSnapshotReplicationPolicies) {
		return true
	}

	return false
}

// SetDpVolumeGroupSnapshotReplicationPolicies gets a reference to the given []DpVolumeGroupSnapshotReplicationPolicy and assigns it to the DpVolumeGroupSnapshotReplicationPolicies field.
func (o *DpVolumeGroupSnapshotReplicationPoliciesResp) SetDpVolumeGroupSnapshotReplicationPolicies(v []DpVolumeGroupSnapshotReplicationPolicy) {
	o.DpVolumeGroupSnapshotReplicationPolicies = v
}

func (o DpVolumeGroupSnapshotReplicationPoliciesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpVolumeGroupSnapshotReplicationPoliciesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DpVolumeGroupSnapshotReplicationPolicies) {
		toSerialize["dp_volume_group_snapshot_replication_policies"] = o.DpVolumeGroupSnapshotReplicationPolicies
	}
	return toSerialize, nil
}

type NullableDpVolumeGroupSnapshotReplicationPoliciesResp struct {
	value *DpVolumeGroupSnapshotReplicationPoliciesResp
	isSet bool
}

func (v NullableDpVolumeGroupSnapshotReplicationPoliciesResp) Get() *DpVolumeGroupSnapshotReplicationPoliciesResp {
	return v.value
}

func (v *NullableDpVolumeGroupSnapshotReplicationPoliciesResp) Set(val *DpVolumeGroupSnapshotReplicationPoliciesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDpVolumeGroupSnapshotReplicationPoliciesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDpVolumeGroupSnapshotReplicationPoliciesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpVolumeGroupSnapshotReplicationPoliciesResp(val *DpVolumeGroupSnapshotReplicationPoliciesResp) *NullableDpVolumeGroupSnapshotReplicationPoliciesResp {
	return &NullableDpVolumeGroupSnapshotReplicationPoliciesResp{value: val, isSet: true}
}

func (v NullableDpVolumeGroupSnapshotReplicationPoliciesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpVolumeGroupSnapshotReplicationPoliciesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


