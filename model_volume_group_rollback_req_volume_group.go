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

// checks if the VolumeGroupRollbackReqVolumeGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeGroupRollbackReqVolumeGroup{}

// VolumeGroupRollbackReqVolumeGroup struct for VolumeGroupRollbackReqVolumeGroup
type VolumeGroupRollbackReqVolumeGroup struct {
	BlockVolumeGroupSnapshotId *int64 `json:"block_volume_group_snapshot_id,omitempty"`
}

// NewVolumeGroupRollbackReqVolumeGroup instantiates a new VolumeGroupRollbackReqVolumeGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeGroupRollbackReqVolumeGroup() *VolumeGroupRollbackReqVolumeGroup {
	this := VolumeGroupRollbackReqVolumeGroup{}
	return &this
}

// NewVolumeGroupRollbackReqVolumeGroupWithDefaults instantiates a new VolumeGroupRollbackReqVolumeGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeGroupRollbackReqVolumeGroupWithDefaults() *VolumeGroupRollbackReqVolumeGroup {
	this := VolumeGroupRollbackReqVolumeGroup{}
	return &this
}

// GetBlockVolumeGroupSnapshotId returns the BlockVolumeGroupSnapshotId field value if set, zero value otherwise.
func (o *VolumeGroupRollbackReqVolumeGroup) GetBlockVolumeGroupSnapshotId() int64 {
	if o == nil || IsNil(o.BlockVolumeGroupSnapshotId) {
		var ret int64
		return ret
	}
	return *o.BlockVolumeGroupSnapshotId
}

// GetBlockVolumeGroupSnapshotIdOk returns a tuple with the BlockVolumeGroupSnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeGroupRollbackReqVolumeGroup) GetBlockVolumeGroupSnapshotIdOk() (*int64, bool) {
	if o == nil || IsNil(o.BlockVolumeGroupSnapshotId) {
		return nil, false
	}
	return o.BlockVolumeGroupSnapshotId, true
}

// HasBlockVolumeGroupSnapshotId returns a boolean if a field has been set.
func (o *VolumeGroupRollbackReqVolumeGroup) HasBlockVolumeGroupSnapshotId() bool {
	if o != nil && !IsNil(o.BlockVolumeGroupSnapshotId) {
		return true
	}

	return false
}

// SetBlockVolumeGroupSnapshotId gets a reference to the given int64 and assigns it to the BlockVolumeGroupSnapshotId field.
func (o *VolumeGroupRollbackReqVolumeGroup) SetBlockVolumeGroupSnapshotId(v int64) {
	o.BlockVolumeGroupSnapshotId = &v
}

func (o VolumeGroupRollbackReqVolumeGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeGroupRollbackReqVolumeGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockVolumeGroupSnapshotId) {
		toSerialize["block_volume_group_snapshot_id"] = o.BlockVolumeGroupSnapshotId
	}
	return toSerialize, nil
}

type NullableVolumeGroupRollbackReqVolumeGroup struct {
	value *VolumeGroupRollbackReqVolumeGroup
	isSet bool
}

func (v NullableVolumeGroupRollbackReqVolumeGroup) Get() *VolumeGroupRollbackReqVolumeGroup {
	return v.value
}

func (v *NullableVolumeGroupRollbackReqVolumeGroup) Set(val *VolumeGroupRollbackReqVolumeGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeGroupRollbackReqVolumeGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeGroupRollbackReqVolumeGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeGroupRollbackReqVolumeGroup(val *VolumeGroupRollbackReqVolumeGroup) *NullableVolumeGroupRollbackReqVolumeGroup {
	return &NullableVolumeGroupRollbackReqVolumeGroup{value: val, isSet: true}
}

func (v NullableVolumeGroupRollbackReqVolumeGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeGroupRollbackReqVolumeGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


