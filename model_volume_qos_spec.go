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

// checks if the VolumeQosSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeQosSpec{}

// VolumeQosSpec VolumeQosSpec is a collection of all qos field
type VolumeQosSpec struct {
	BurstTotalBw *int64 `json:"burst_total_bw,omitempty"`
	BurstTotalIops *int64 `json:"burst_total_iops,omitempty"`
	MaxTotalBw *int64 `json:"max_total_bw,omitempty"`
	MaxTotalIops *int64 `json:"max_total_iops,omitempty"`
}

// NewVolumeQosSpec instantiates a new VolumeQosSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeQosSpec() *VolumeQosSpec {
	this := VolumeQosSpec{}
	return &this
}

// NewVolumeQosSpecWithDefaults instantiates a new VolumeQosSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeQosSpecWithDefaults() *VolumeQosSpec {
	this := VolumeQosSpec{}
	return &this
}

// GetBurstTotalBw returns the BurstTotalBw field value if set, zero value otherwise.
func (o *VolumeQosSpec) GetBurstTotalBw() int64 {
	if o == nil || IsNil(o.BurstTotalBw) {
		var ret int64
		return ret
	}
	return *o.BurstTotalBw
}

// GetBurstTotalBwOk returns a tuple with the BurstTotalBw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeQosSpec) GetBurstTotalBwOk() (*int64, bool) {
	if o == nil || IsNil(o.BurstTotalBw) {
		return nil, false
	}
	return o.BurstTotalBw, true
}

// HasBurstTotalBw returns a boolean if a field has been set.
func (o *VolumeQosSpec) HasBurstTotalBw() bool {
	if o != nil && !IsNil(o.BurstTotalBw) {
		return true
	}

	return false
}

// SetBurstTotalBw gets a reference to the given int64 and assigns it to the BurstTotalBw field.
func (o *VolumeQosSpec) SetBurstTotalBw(v int64) {
	o.BurstTotalBw = &v
}

// GetBurstTotalIops returns the BurstTotalIops field value if set, zero value otherwise.
func (o *VolumeQosSpec) GetBurstTotalIops() int64 {
	if o == nil || IsNil(o.BurstTotalIops) {
		var ret int64
		return ret
	}
	return *o.BurstTotalIops
}

// GetBurstTotalIopsOk returns a tuple with the BurstTotalIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeQosSpec) GetBurstTotalIopsOk() (*int64, bool) {
	if o == nil || IsNil(o.BurstTotalIops) {
		return nil, false
	}
	return o.BurstTotalIops, true
}

// HasBurstTotalIops returns a boolean if a field has been set.
func (o *VolumeQosSpec) HasBurstTotalIops() bool {
	if o != nil && !IsNil(o.BurstTotalIops) {
		return true
	}

	return false
}

// SetBurstTotalIops gets a reference to the given int64 and assigns it to the BurstTotalIops field.
func (o *VolumeQosSpec) SetBurstTotalIops(v int64) {
	o.BurstTotalIops = &v
}

// GetMaxTotalBw returns the MaxTotalBw field value if set, zero value otherwise.
func (o *VolumeQosSpec) GetMaxTotalBw() int64 {
	if o == nil || IsNil(o.MaxTotalBw) {
		var ret int64
		return ret
	}
	return *o.MaxTotalBw
}

// GetMaxTotalBwOk returns a tuple with the MaxTotalBw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeQosSpec) GetMaxTotalBwOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxTotalBw) {
		return nil, false
	}
	return o.MaxTotalBw, true
}

// HasMaxTotalBw returns a boolean if a field has been set.
func (o *VolumeQosSpec) HasMaxTotalBw() bool {
	if o != nil && !IsNil(o.MaxTotalBw) {
		return true
	}

	return false
}

// SetMaxTotalBw gets a reference to the given int64 and assigns it to the MaxTotalBw field.
func (o *VolumeQosSpec) SetMaxTotalBw(v int64) {
	o.MaxTotalBw = &v
}

// GetMaxTotalIops returns the MaxTotalIops field value if set, zero value otherwise.
func (o *VolumeQosSpec) GetMaxTotalIops() int64 {
	if o == nil || IsNil(o.MaxTotalIops) {
		var ret int64
		return ret
	}
	return *o.MaxTotalIops
}

// GetMaxTotalIopsOk returns a tuple with the MaxTotalIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeQosSpec) GetMaxTotalIopsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxTotalIops) {
		return nil, false
	}
	return o.MaxTotalIops, true
}

// HasMaxTotalIops returns a boolean if a field has been set.
func (o *VolumeQosSpec) HasMaxTotalIops() bool {
	if o != nil && !IsNil(o.MaxTotalIops) {
		return true
	}

	return false
}

// SetMaxTotalIops gets a reference to the given int64 and assigns it to the MaxTotalIops field.
func (o *VolumeQosSpec) SetMaxTotalIops(v int64) {
	o.MaxTotalIops = &v
}

func (o VolumeQosSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeQosSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BurstTotalBw) {
		toSerialize["burst_total_bw"] = o.BurstTotalBw
	}
	if !IsNil(o.BurstTotalIops) {
		toSerialize["burst_total_iops"] = o.BurstTotalIops
	}
	if !IsNil(o.MaxTotalBw) {
		toSerialize["max_total_bw"] = o.MaxTotalBw
	}
	if !IsNil(o.MaxTotalIops) {
		toSerialize["max_total_iops"] = o.MaxTotalIops
	}
	return toSerialize, nil
}

type NullableVolumeQosSpec struct {
	value *VolumeQosSpec
	isSet bool
}

func (v NullableVolumeQosSpec) Get() *VolumeQosSpec {
	return v.value
}

func (v *NullableVolumeQosSpec) Set(val *VolumeQosSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeQosSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeQosSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeQosSpec(val *VolumeQosSpec) *NullableVolumeQosSpec {
	return &NullableVolumeQosSpec{value: val, isSet: true}
}

func (v NullableVolumeQosSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeQosSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


