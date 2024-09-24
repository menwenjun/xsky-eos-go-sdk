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

// checks if the DpVolumeGroupSnapshotReplicationPairUpdateReqPair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpVolumeGroupSnapshotReplicationPairUpdateReqPair{}

// DpVolumeGroupSnapshotReplicationPairUpdateReqPair struct for DpVolumeGroupSnapshotReplicationPairUpdateReqPair
type DpVolumeGroupSnapshotReplicationPairUpdateReqPair struct {
	// policy cron
	MasterPolicyCron *string `json:"master_policy_cron,omitempty"`
	// master volume group name
	MasterVolumeGroupName *string `json:"master_volume_group_name,omitempty"`
	// slave volume group id
	SlaveVolumeGroupId *int64 `json:"slave_volume_group_id,omitempty"`
	// pair status
	Status *string `json:"status,omitempty"`
}

// NewDpVolumeGroupSnapshotReplicationPairUpdateReqPair instantiates a new DpVolumeGroupSnapshotReplicationPairUpdateReqPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpVolumeGroupSnapshotReplicationPairUpdateReqPair() *DpVolumeGroupSnapshotReplicationPairUpdateReqPair {
	this := DpVolumeGroupSnapshotReplicationPairUpdateReqPair{}
	return &this
}

// NewDpVolumeGroupSnapshotReplicationPairUpdateReqPairWithDefaults instantiates a new DpVolumeGroupSnapshotReplicationPairUpdateReqPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpVolumeGroupSnapshotReplicationPairUpdateReqPairWithDefaults() *DpVolumeGroupSnapshotReplicationPairUpdateReqPair {
	this := DpVolumeGroupSnapshotReplicationPairUpdateReqPair{}
	return &this
}

// GetMasterPolicyCron returns the MasterPolicyCron field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) GetMasterPolicyCron() string {
	if o == nil || IsNil(o.MasterPolicyCron) {
		var ret string
		return ret
	}
	return *o.MasterPolicyCron
}

// GetMasterPolicyCronOk returns a tuple with the MasterPolicyCron field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) GetMasterPolicyCronOk() (*string, bool) {
	if o == nil || IsNil(o.MasterPolicyCron) {
		return nil, false
	}
	return o.MasterPolicyCron, true
}

// HasMasterPolicyCron returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) HasMasterPolicyCron() bool {
	if o != nil && !IsNil(o.MasterPolicyCron) {
		return true
	}

	return false
}

// SetMasterPolicyCron gets a reference to the given string and assigns it to the MasterPolicyCron field.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) SetMasterPolicyCron(v string) {
	o.MasterPolicyCron = &v
}

// GetMasterVolumeGroupName returns the MasterVolumeGroupName field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) GetMasterVolumeGroupName() string {
	if o == nil || IsNil(o.MasterVolumeGroupName) {
		var ret string
		return ret
	}
	return *o.MasterVolumeGroupName
}

// GetMasterVolumeGroupNameOk returns a tuple with the MasterVolumeGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) GetMasterVolumeGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.MasterVolumeGroupName) {
		return nil, false
	}
	return o.MasterVolumeGroupName, true
}

// HasMasterVolumeGroupName returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) HasMasterVolumeGroupName() bool {
	if o != nil && !IsNil(o.MasterVolumeGroupName) {
		return true
	}

	return false
}

// SetMasterVolumeGroupName gets a reference to the given string and assigns it to the MasterVolumeGroupName field.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) SetMasterVolumeGroupName(v string) {
	o.MasterVolumeGroupName = &v
}

// GetSlaveVolumeGroupId returns the SlaveVolumeGroupId field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) GetSlaveVolumeGroupId() int64 {
	if o == nil || IsNil(o.SlaveVolumeGroupId) {
		var ret int64
		return ret
	}
	return *o.SlaveVolumeGroupId
}

// GetSlaveVolumeGroupIdOk returns a tuple with the SlaveVolumeGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) GetSlaveVolumeGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SlaveVolumeGroupId) {
		return nil, false
	}
	return o.SlaveVolumeGroupId, true
}

// HasSlaveVolumeGroupId returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) HasSlaveVolumeGroupId() bool {
	if o != nil && !IsNil(o.SlaveVolumeGroupId) {
		return true
	}

	return false
}

// SetSlaveVolumeGroupId gets a reference to the given int64 and assigns it to the SlaveVolumeGroupId field.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) SetSlaveVolumeGroupId(v int64) {
	o.SlaveVolumeGroupId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) SetStatus(v string) {
	o.Status = &v
}

func (o DpVolumeGroupSnapshotReplicationPairUpdateReqPair) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpVolumeGroupSnapshotReplicationPairUpdateReqPair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MasterPolicyCron) {
		toSerialize["master_policy_cron"] = o.MasterPolicyCron
	}
	if !IsNil(o.MasterVolumeGroupName) {
		toSerialize["master_volume_group_name"] = o.MasterVolumeGroupName
	}
	if !IsNil(o.SlaveVolumeGroupId) {
		toSerialize["slave_volume_group_id"] = o.SlaveVolumeGroupId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableDpVolumeGroupSnapshotReplicationPairUpdateReqPair struct {
	value *DpVolumeGroupSnapshotReplicationPairUpdateReqPair
	isSet bool
}

func (v NullableDpVolumeGroupSnapshotReplicationPairUpdateReqPair) Get() *DpVolumeGroupSnapshotReplicationPairUpdateReqPair {
	return v.value
}

func (v *NullableDpVolumeGroupSnapshotReplicationPairUpdateReqPair) Set(val *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) {
	v.value = val
	v.isSet = true
}

func (v NullableDpVolumeGroupSnapshotReplicationPairUpdateReqPair) IsSet() bool {
	return v.isSet
}

func (v *NullableDpVolumeGroupSnapshotReplicationPairUpdateReqPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpVolumeGroupSnapshotReplicationPairUpdateReqPair(val *DpVolumeGroupSnapshotReplicationPairUpdateReqPair) *NullableDpVolumeGroupSnapshotReplicationPairUpdateReqPair {
	return &NullableDpVolumeGroupSnapshotReplicationPairUpdateReqPair{value: val, isSet: true}
}

func (v NullableDpVolumeGroupSnapshotReplicationPairUpdateReqPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpVolumeGroupSnapshotReplicationPairUpdateReqPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


