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

// checks if the UpdateVolumeStatsReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVolumeStatsReq{}

// UpdateVolumeStatsReq struct for UpdateVolumeStatsReq
type UpdateVolumeStatsReq struct {
	PoolName *string `json:"pool_name,omitempty"`
	VolumeStats *map[string]VolumeStat `json:"volume_stats,omitempty"`
}

// NewUpdateVolumeStatsReq instantiates a new UpdateVolumeStatsReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVolumeStatsReq() *UpdateVolumeStatsReq {
	this := UpdateVolumeStatsReq{}
	return &this
}

// NewUpdateVolumeStatsReqWithDefaults instantiates a new UpdateVolumeStatsReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVolumeStatsReqWithDefaults() *UpdateVolumeStatsReq {
	this := UpdateVolumeStatsReq{}
	return &this
}

// GetPoolName returns the PoolName field value if set, zero value otherwise.
func (o *UpdateVolumeStatsReq) GetPoolName() string {
	if o == nil || IsNil(o.PoolName) {
		var ret string
		return ret
	}
	return *o.PoolName
}

// GetPoolNameOk returns a tuple with the PoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeStatsReq) GetPoolNameOk() (*string, bool) {
	if o == nil || IsNil(o.PoolName) {
		return nil, false
	}
	return o.PoolName, true
}

// HasPoolName returns a boolean if a field has been set.
func (o *UpdateVolumeStatsReq) HasPoolName() bool {
	if o != nil && !IsNil(o.PoolName) {
		return true
	}

	return false
}

// SetPoolName gets a reference to the given string and assigns it to the PoolName field.
func (o *UpdateVolumeStatsReq) SetPoolName(v string) {
	o.PoolName = &v
}

// GetVolumeStats returns the VolumeStats field value if set, zero value otherwise.
func (o *UpdateVolumeStatsReq) GetVolumeStats() map[string]VolumeStat {
	if o == nil || IsNil(o.VolumeStats) {
		var ret map[string]VolumeStat
		return ret
	}
	return *o.VolumeStats
}

// GetVolumeStatsOk returns a tuple with the VolumeStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeStatsReq) GetVolumeStatsOk() (*map[string]VolumeStat, bool) {
	if o == nil || IsNil(o.VolumeStats) {
		return nil, false
	}
	return o.VolumeStats, true
}

// HasVolumeStats returns a boolean if a field has been set.
func (o *UpdateVolumeStatsReq) HasVolumeStats() bool {
	if o != nil && !IsNil(o.VolumeStats) {
		return true
	}

	return false
}

// SetVolumeStats gets a reference to the given map[string]VolumeStat and assigns it to the VolumeStats field.
func (o *UpdateVolumeStatsReq) SetVolumeStats(v map[string]VolumeStat) {
	o.VolumeStats = &v
}

func (o UpdateVolumeStatsReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVolumeStatsReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PoolName) {
		toSerialize["pool_name"] = o.PoolName
	}
	if !IsNil(o.VolumeStats) {
		toSerialize["volume_stats"] = o.VolumeStats
	}
	return toSerialize, nil
}

type NullableUpdateVolumeStatsReq struct {
	value *UpdateVolumeStatsReq
	isSet bool
}

func (v NullableUpdateVolumeStatsReq) Get() *UpdateVolumeStatsReq {
	return v.value
}

func (v *NullableUpdateVolumeStatsReq) Set(val *UpdateVolumeStatsReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVolumeStatsReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVolumeStatsReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVolumeStatsReq(val *UpdateVolumeStatsReq) *NullableUpdateVolumeStatsReq {
	return &NullableUpdateVolumeStatsReq{value: val, isSet: true}
}

func (v NullableUpdateVolumeStatsReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVolumeStatsReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


