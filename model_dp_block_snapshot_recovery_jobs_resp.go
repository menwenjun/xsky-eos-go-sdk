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

// checks if the DpBlockSnapshotRecoveryJobsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpBlockSnapshotRecoveryJobsResp{}

// DpBlockSnapshotRecoveryJobsResp struct for DpBlockSnapshotRecoveryJobsResp
type DpBlockSnapshotRecoveryJobsResp struct {
	DpBlockSnapshotRecoveryJobs []DpBlockSnapshotRecoveryJob `json:"dp_block_snapshot_recovery_jobs,omitempty"`
}

// NewDpBlockSnapshotRecoveryJobsResp instantiates a new DpBlockSnapshotRecoveryJobsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpBlockSnapshotRecoveryJobsResp() *DpBlockSnapshotRecoveryJobsResp {
	this := DpBlockSnapshotRecoveryJobsResp{}
	return &this
}

// NewDpBlockSnapshotRecoveryJobsRespWithDefaults instantiates a new DpBlockSnapshotRecoveryJobsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpBlockSnapshotRecoveryJobsRespWithDefaults() *DpBlockSnapshotRecoveryJobsResp {
	this := DpBlockSnapshotRecoveryJobsResp{}
	return &this
}

// GetDpBlockSnapshotRecoveryJobs returns the DpBlockSnapshotRecoveryJobs field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJobsResp) GetDpBlockSnapshotRecoveryJobs() []DpBlockSnapshotRecoveryJob {
	if o == nil || IsNil(o.DpBlockSnapshotRecoveryJobs) {
		var ret []DpBlockSnapshotRecoveryJob
		return ret
	}
	return o.DpBlockSnapshotRecoveryJobs
}

// GetDpBlockSnapshotRecoveryJobsOk returns a tuple with the DpBlockSnapshotRecoveryJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJobsResp) GetDpBlockSnapshotRecoveryJobsOk() ([]DpBlockSnapshotRecoveryJob, bool) {
	if o == nil || IsNil(o.DpBlockSnapshotRecoveryJobs) {
		return nil, false
	}
	return o.DpBlockSnapshotRecoveryJobs, true
}

// HasDpBlockSnapshotRecoveryJobs returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJobsResp) HasDpBlockSnapshotRecoveryJobs() bool {
	if o != nil && !IsNil(o.DpBlockSnapshotRecoveryJobs) {
		return true
	}

	return false
}

// SetDpBlockSnapshotRecoveryJobs gets a reference to the given []DpBlockSnapshotRecoveryJob and assigns it to the DpBlockSnapshotRecoveryJobs field.
func (o *DpBlockSnapshotRecoveryJobsResp) SetDpBlockSnapshotRecoveryJobs(v []DpBlockSnapshotRecoveryJob) {
	o.DpBlockSnapshotRecoveryJobs = v
}

func (o DpBlockSnapshotRecoveryJobsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpBlockSnapshotRecoveryJobsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DpBlockSnapshotRecoveryJobs) {
		toSerialize["dp_block_snapshot_recovery_jobs"] = o.DpBlockSnapshotRecoveryJobs
	}
	return toSerialize, nil
}

type NullableDpBlockSnapshotRecoveryJobsResp struct {
	value *DpBlockSnapshotRecoveryJobsResp
	isSet bool
}

func (v NullableDpBlockSnapshotRecoveryJobsResp) Get() *DpBlockSnapshotRecoveryJobsResp {
	return v.value
}

func (v *NullableDpBlockSnapshotRecoveryJobsResp) Set(val *DpBlockSnapshotRecoveryJobsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDpBlockSnapshotRecoveryJobsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDpBlockSnapshotRecoveryJobsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpBlockSnapshotRecoveryJobsResp(val *DpBlockSnapshotRecoveryJobsResp) *NullableDpBlockSnapshotRecoveryJobsResp {
	return &NullableDpBlockSnapshotRecoveryJobsResp{value: val, isSet: true}
}

func (v NullableDpBlockSnapshotRecoveryJobsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpBlockSnapshotRecoveryJobsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


