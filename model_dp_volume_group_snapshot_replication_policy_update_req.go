/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DpVolumeGroupSnapshotReplicationPolicyUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpVolumeGroupSnapshotReplicationPolicyUpdateReq{}

// DpVolumeGroupSnapshotReplicationPolicyUpdateReq struct for DpVolumeGroupSnapshotReplicationPolicyUpdateReq
type DpVolumeGroupSnapshotReplicationPolicyUpdateReq struct {
	DpVolumeGroupSnapshotReplicationPolicy DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy `json:"dp_volume_group_snapshot_replication_policy"`
}

type _DpVolumeGroupSnapshotReplicationPolicyUpdateReq DpVolumeGroupSnapshotReplicationPolicyUpdateReq

// NewDpVolumeGroupSnapshotReplicationPolicyUpdateReq instantiates a new DpVolumeGroupSnapshotReplicationPolicyUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpVolumeGroupSnapshotReplicationPolicyUpdateReq(dpVolumeGroupSnapshotReplicationPolicy DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy) *DpVolumeGroupSnapshotReplicationPolicyUpdateReq {
	this := DpVolumeGroupSnapshotReplicationPolicyUpdateReq{}
	this.DpVolumeGroupSnapshotReplicationPolicy = dpVolumeGroupSnapshotReplicationPolicy
	return &this
}

// NewDpVolumeGroupSnapshotReplicationPolicyUpdateReqWithDefaults instantiates a new DpVolumeGroupSnapshotReplicationPolicyUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpVolumeGroupSnapshotReplicationPolicyUpdateReqWithDefaults() *DpVolumeGroupSnapshotReplicationPolicyUpdateReq {
	this := DpVolumeGroupSnapshotReplicationPolicyUpdateReq{}
	return &this
}

// GetDpVolumeGroupSnapshotReplicationPolicy returns the DpVolumeGroupSnapshotReplicationPolicy field value
func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReq) GetDpVolumeGroupSnapshotReplicationPolicy() DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy {
	if o == nil {
		var ret DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy
		return ret
	}

	return o.DpVolumeGroupSnapshotReplicationPolicy
}

// GetDpVolumeGroupSnapshotReplicationPolicyOk returns a tuple with the DpVolumeGroupSnapshotReplicationPolicy field value
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReq) GetDpVolumeGroupSnapshotReplicationPolicyOk() (*DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DpVolumeGroupSnapshotReplicationPolicy, true
}

// SetDpVolumeGroupSnapshotReplicationPolicy sets field value
func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReq) SetDpVolumeGroupSnapshotReplicationPolicy(v DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy) {
	o.DpVolumeGroupSnapshotReplicationPolicy = v
}

func (o DpVolumeGroupSnapshotReplicationPolicyUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpVolumeGroupSnapshotReplicationPolicyUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dp_volume_group_snapshot_replication_policy"] = o.DpVolumeGroupSnapshotReplicationPolicy
	return toSerialize, nil
}

func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dp_volume_group_snapshot_replication_policy",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDpVolumeGroupSnapshotReplicationPolicyUpdateReq := _DpVolumeGroupSnapshotReplicationPolicyUpdateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDpVolumeGroupSnapshotReplicationPolicyUpdateReq)

	if err != nil {
		return err
	}

	*o = DpVolumeGroupSnapshotReplicationPolicyUpdateReq(varDpVolumeGroupSnapshotReplicationPolicyUpdateReq)

	return err
}

type NullableDpVolumeGroupSnapshotReplicationPolicyUpdateReq struct {
	value *DpVolumeGroupSnapshotReplicationPolicyUpdateReq
	isSet bool
}

func (v NullableDpVolumeGroupSnapshotReplicationPolicyUpdateReq) Get() *DpVolumeGroupSnapshotReplicationPolicyUpdateReq {
	return v.value
}

func (v *NullableDpVolumeGroupSnapshotReplicationPolicyUpdateReq) Set(val *DpVolumeGroupSnapshotReplicationPolicyUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDpVolumeGroupSnapshotReplicationPolicyUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDpVolumeGroupSnapshotReplicationPolicyUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpVolumeGroupSnapshotReplicationPolicyUpdateReq(val *DpVolumeGroupSnapshotReplicationPolicyUpdateReq) *NullableDpVolumeGroupSnapshotReplicationPolicyUpdateReq {
	return &NullableDpVolumeGroupSnapshotReplicationPolicyUpdateReq{value: val, isSet: true}
}

func (v NullableDpVolumeGroupSnapshotReplicationPolicyUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpVolumeGroupSnapshotReplicationPolicyUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


