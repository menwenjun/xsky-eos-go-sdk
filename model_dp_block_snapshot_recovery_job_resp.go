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

// checks if the DpBlockSnapshotRecoveryJobResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpBlockSnapshotRecoveryJobResp{}

// DpBlockSnapshotRecoveryJobResp struct for DpBlockSnapshotRecoveryJobResp
type DpBlockSnapshotRecoveryJobResp struct {
	DpBlockSnapshotRecoveryJob *DpBlockSnapshotRecoveryJob `json:"dp_block_snapshot_recovery_job,omitempty"`
}

// NewDpBlockSnapshotRecoveryJobResp instantiates a new DpBlockSnapshotRecoveryJobResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpBlockSnapshotRecoveryJobResp() *DpBlockSnapshotRecoveryJobResp {
	this := DpBlockSnapshotRecoveryJobResp{}
	return &this
}

// NewDpBlockSnapshotRecoveryJobRespWithDefaults instantiates a new DpBlockSnapshotRecoveryJobResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpBlockSnapshotRecoveryJobRespWithDefaults() *DpBlockSnapshotRecoveryJobResp {
	this := DpBlockSnapshotRecoveryJobResp{}
	return &this
}

// GetDpBlockSnapshotRecoveryJob returns the DpBlockSnapshotRecoveryJob field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJobResp) GetDpBlockSnapshotRecoveryJob() DpBlockSnapshotRecoveryJob {
	if o == nil || IsNil(o.DpBlockSnapshotRecoveryJob) {
		var ret DpBlockSnapshotRecoveryJob
		return ret
	}
	return *o.DpBlockSnapshotRecoveryJob
}

// GetDpBlockSnapshotRecoveryJobOk returns a tuple with the DpBlockSnapshotRecoveryJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJobResp) GetDpBlockSnapshotRecoveryJobOk() (*DpBlockSnapshotRecoveryJob, bool) {
	if o == nil || IsNil(o.DpBlockSnapshotRecoveryJob) {
		return nil, false
	}
	return o.DpBlockSnapshotRecoveryJob, true
}

// HasDpBlockSnapshotRecoveryJob returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJobResp) HasDpBlockSnapshotRecoveryJob() bool {
	if o != nil && !IsNil(o.DpBlockSnapshotRecoveryJob) {
		return true
	}

	return false
}

// SetDpBlockSnapshotRecoveryJob gets a reference to the given DpBlockSnapshotRecoveryJob and assigns it to the DpBlockSnapshotRecoveryJob field.
func (o *DpBlockSnapshotRecoveryJobResp) SetDpBlockSnapshotRecoveryJob(v DpBlockSnapshotRecoveryJob) {
	o.DpBlockSnapshotRecoveryJob = &v
}

func (o DpBlockSnapshotRecoveryJobResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpBlockSnapshotRecoveryJobResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DpBlockSnapshotRecoveryJob) {
		toSerialize["dp_block_snapshot_recovery_job"] = o.DpBlockSnapshotRecoveryJob
	}
	return toSerialize, nil
}

type NullableDpBlockSnapshotRecoveryJobResp struct {
	value *DpBlockSnapshotRecoveryJobResp
	isSet bool
}

func (v NullableDpBlockSnapshotRecoveryJobResp) Get() *DpBlockSnapshotRecoveryJobResp {
	return v.value
}

func (v *NullableDpBlockSnapshotRecoveryJobResp) Set(val *DpBlockSnapshotRecoveryJobResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDpBlockSnapshotRecoveryJobResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDpBlockSnapshotRecoveryJobResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpBlockSnapshotRecoveryJobResp(val *DpBlockSnapshotRecoveryJobResp) *NullableDpBlockSnapshotRecoveryJobResp {
	return &NullableDpBlockSnapshotRecoveryJobResp{value: val, isSet: true}
}

func (v NullableDpBlockSnapshotRecoveryJobResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpBlockSnapshotRecoveryJobResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


