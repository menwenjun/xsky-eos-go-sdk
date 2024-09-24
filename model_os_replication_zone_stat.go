/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the OSReplicationZoneStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSReplicationZoneStat{}

// OSReplicationZoneStat OSReplicationZoneStat contains sync stat info of os replication zone
type OSReplicationZoneStat struct {
	BandwidthKbyte *float64 `json:"bandwidth_kbyte,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	ObjectKbytesPs *float64 `json:"object_kbytes_ps,omitempty"`
	ObjectsPm *float64 `json:"objects_pm,omitempty"`
	TotalBytes *int64 `json:"total_bytes,omitempty"`
	TotalObjectBytes *int64 `json:"total_object_bytes,omitempty"`
	TotalObjects *int64 `json:"total_objects,omitempty"`
}

// NewOSReplicationZoneStat instantiates a new OSReplicationZoneStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSReplicationZoneStat() *OSReplicationZoneStat {
	this := OSReplicationZoneStat{}
	return &this
}

// NewOSReplicationZoneStatWithDefaults instantiates a new OSReplicationZoneStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSReplicationZoneStatWithDefaults() *OSReplicationZoneStat {
	this := OSReplicationZoneStat{}
	return &this
}

// GetBandwidthKbyte returns the BandwidthKbyte field value if set, zero value otherwise.
func (o *OSReplicationZoneStat) GetBandwidthKbyte() float64 {
	if o == nil || IsNil(o.BandwidthKbyte) {
		var ret float64
		return ret
	}
	return *o.BandwidthKbyte
}

// GetBandwidthKbyteOk returns a tuple with the BandwidthKbyte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationZoneStat) GetBandwidthKbyteOk() (*float64, bool) {
	if o == nil || IsNil(o.BandwidthKbyte) {
		return nil, false
	}
	return o.BandwidthKbyte, true
}

// HasBandwidthKbyte returns a boolean if a field has been set.
func (o *OSReplicationZoneStat) HasBandwidthKbyte() bool {
	if o != nil && !IsNil(o.BandwidthKbyte) {
		return true
	}

	return false
}

// SetBandwidthKbyte gets a reference to the given float64 and assigns it to the BandwidthKbyte field.
func (o *OSReplicationZoneStat) SetBandwidthKbyte(v float64) {
	o.BandwidthKbyte = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *OSReplicationZoneStat) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationZoneStat) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *OSReplicationZoneStat) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *OSReplicationZoneStat) SetCreate(v time.Time) {
	o.Create = &v
}

// GetObjectKbytesPs returns the ObjectKbytesPs field value if set, zero value otherwise.
func (o *OSReplicationZoneStat) GetObjectKbytesPs() float64 {
	if o == nil || IsNil(o.ObjectKbytesPs) {
		var ret float64
		return ret
	}
	return *o.ObjectKbytesPs
}

// GetObjectKbytesPsOk returns a tuple with the ObjectKbytesPs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationZoneStat) GetObjectKbytesPsOk() (*float64, bool) {
	if o == nil || IsNil(o.ObjectKbytesPs) {
		return nil, false
	}
	return o.ObjectKbytesPs, true
}

// HasObjectKbytesPs returns a boolean if a field has been set.
func (o *OSReplicationZoneStat) HasObjectKbytesPs() bool {
	if o != nil && !IsNil(o.ObjectKbytesPs) {
		return true
	}

	return false
}

// SetObjectKbytesPs gets a reference to the given float64 and assigns it to the ObjectKbytesPs field.
func (o *OSReplicationZoneStat) SetObjectKbytesPs(v float64) {
	o.ObjectKbytesPs = &v
}

// GetObjectsPm returns the ObjectsPm field value if set, zero value otherwise.
func (o *OSReplicationZoneStat) GetObjectsPm() float64 {
	if o == nil || IsNil(o.ObjectsPm) {
		var ret float64
		return ret
	}
	return *o.ObjectsPm
}

// GetObjectsPmOk returns a tuple with the ObjectsPm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationZoneStat) GetObjectsPmOk() (*float64, bool) {
	if o == nil || IsNil(o.ObjectsPm) {
		return nil, false
	}
	return o.ObjectsPm, true
}

// HasObjectsPm returns a boolean if a field has been set.
func (o *OSReplicationZoneStat) HasObjectsPm() bool {
	if o != nil && !IsNil(o.ObjectsPm) {
		return true
	}

	return false
}

// SetObjectsPm gets a reference to the given float64 and assigns it to the ObjectsPm field.
func (o *OSReplicationZoneStat) SetObjectsPm(v float64) {
	o.ObjectsPm = &v
}

// GetTotalBytes returns the TotalBytes field value if set, zero value otherwise.
func (o *OSReplicationZoneStat) GetTotalBytes() int64 {
	if o == nil || IsNil(o.TotalBytes) {
		var ret int64
		return ret
	}
	return *o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationZoneStat) GetTotalBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalBytes) {
		return nil, false
	}
	return o.TotalBytes, true
}

// HasTotalBytes returns a boolean if a field has been set.
func (o *OSReplicationZoneStat) HasTotalBytes() bool {
	if o != nil && !IsNil(o.TotalBytes) {
		return true
	}

	return false
}

// SetTotalBytes gets a reference to the given int64 and assigns it to the TotalBytes field.
func (o *OSReplicationZoneStat) SetTotalBytes(v int64) {
	o.TotalBytes = &v
}

// GetTotalObjectBytes returns the TotalObjectBytes field value if set, zero value otherwise.
func (o *OSReplicationZoneStat) GetTotalObjectBytes() int64 {
	if o == nil || IsNil(o.TotalObjectBytes) {
		var ret int64
		return ret
	}
	return *o.TotalObjectBytes
}

// GetTotalObjectBytesOk returns a tuple with the TotalObjectBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationZoneStat) GetTotalObjectBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalObjectBytes) {
		return nil, false
	}
	return o.TotalObjectBytes, true
}

// HasTotalObjectBytes returns a boolean if a field has been set.
func (o *OSReplicationZoneStat) HasTotalObjectBytes() bool {
	if o != nil && !IsNil(o.TotalObjectBytes) {
		return true
	}

	return false
}

// SetTotalObjectBytes gets a reference to the given int64 and assigns it to the TotalObjectBytes field.
func (o *OSReplicationZoneStat) SetTotalObjectBytes(v int64) {
	o.TotalObjectBytes = &v
}

// GetTotalObjects returns the TotalObjects field value if set, zero value otherwise.
func (o *OSReplicationZoneStat) GetTotalObjects() int64 {
	if o == nil || IsNil(o.TotalObjects) {
		var ret int64
		return ret
	}
	return *o.TotalObjects
}

// GetTotalObjectsOk returns a tuple with the TotalObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationZoneStat) GetTotalObjectsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalObjects) {
		return nil, false
	}
	return o.TotalObjects, true
}

// HasTotalObjects returns a boolean if a field has been set.
func (o *OSReplicationZoneStat) HasTotalObjects() bool {
	if o != nil && !IsNil(o.TotalObjects) {
		return true
	}

	return false
}

// SetTotalObjects gets a reference to the given int64 and assigns it to the TotalObjects field.
func (o *OSReplicationZoneStat) SetTotalObjects(v int64) {
	o.TotalObjects = &v
}

func (o OSReplicationZoneStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSReplicationZoneStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BandwidthKbyte) {
		toSerialize["bandwidth_kbyte"] = o.BandwidthKbyte
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.ObjectKbytesPs) {
		toSerialize["object_kbytes_ps"] = o.ObjectKbytesPs
	}
	if !IsNil(o.ObjectsPm) {
		toSerialize["objects_pm"] = o.ObjectsPm
	}
	if !IsNil(o.TotalBytes) {
		toSerialize["total_bytes"] = o.TotalBytes
	}
	if !IsNil(o.TotalObjectBytes) {
		toSerialize["total_object_bytes"] = o.TotalObjectBytes
	}
	if !IsNil(o.TotalObjects) {
		toSerialize["total_objects"] = o.TotalObjects
	}
	return toSerialize, nil
}

type NullableOSReplicationZoneStat struct {
	value *OSReplicationZoneStat
	isSet bool
}

func (v NullableOSReplicationZoneStat) Get() *OSReplicationZoneStat {
	return v.value
}

func (v *NullableOSReplicationZoneStat) Set(val *OSReplicationZoneStat) {
	v.value = val
	v.isSet = true
}

func (v NullableOSReplicationZoneStat) IsSet() bool {
	return v.isSet
}

func (v *NullableOSReplicationZoneStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSReplicationZoneStat(val *OSReplicationZoneStat) *NullableOSReplicationZoneStat {
	return &NullableOSReplicationZoneStat{value: val, isSet: true}
}

func (v NullableOSReplicationZoneStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSReplicationZoneStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


