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

// checks if the OSReplicationZoneRecordResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSReplicationZoneRecordResp{}

// OSReplicationZoneRecordResp struct for OSReplicationZoneRecordResp
type OSReplicationZoneRecordResp struct {
	OsReplicationZone *OSReplicationZoneRecord `json:"os_replication_zone,omitempty"`
}

// NewOSReplicationZoneRecordResp instantiates a new OSReplicationZoneRecordResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSReplicationZoneRecordResp() *OSReplicationZoneRecordResp {
	this := OSReplicationZoneRecordResp{}
	return &this
}

// NewOSReplicationZoneRecordRespWithDefaults instantiates a new OSReplicationZoneRecordResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSReplicationZoneRecordRespWithDefaults() *OSReplicationZoneRecordResp {
	this := OSReplicationZoneRecordResp{}
	return &this
}

// GetOsReplicationZone returns the OsReplicationZone field value if set, zero value otherwise.
func (o *OSReplicationZoneRecordResp) GetOsReplicationZone() OSReplicationZoneRecord {
	if o == nil || IsNil(o.OsReplicationZone) {
		var ret OSReplicationZoneRecord
		return ret
	}
	return *o.OsReplicationZone
}

// GetOsReplicationZoneOk returns a tuple with the OsReplicationZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationZoneRecordResp) GetOsReplicationZoneOk() (*OSReplicationZoneRecord, bool) {
	if o == nil || IsNil(o.OsReplicationZone) {
		return nil, false
	}
	return o.OsReplicationZone, true
}

// HasOsReplicationZone returns a boolean if a field has been set.
func (o *OSReplicationZoneRecordResp) HasOsReplicationZone() bool {
	if o != nil && !IsNil(o.OsReplicationZone) {
		return true
	}

	return false
}

// SetOsReplicationZone gets a reference to the given OSReplicationZoneRecord and assigns it to the OsReplicationZone field.
func (o *OSReplicationZoneRecordResp) SetOsReplicationZone(v OSReplicationZoneRecord) {
	o.OsReplicationZone = &v
}

func (o OSReplicationZoneRecordResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSReplicationZoneRecordResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OsReplicationZone) {
		toSerialize["os_replication_zone"] = o.OsReplicationZone
	}
	return toSerialize, nil
}

type NullableOSReplicationZoneRecordResp struct {
	value *OSReplicationZoneRecordResp
	isSet bool
}

func (v NullableOSReplicationZoneRecordResp) Get() *OSReplicationZoneRecordResp {
	return v.value
}

func (v *NullableOSReplicationZoneRecordResp) Set(val *OSReplicationZoneRecordResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOSReplicationZoneRecordResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOSReplicationZoneRecordResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSReplicationZoneRecordResp(val *OSReplicationZoneRecordResp) *NullableOSReplicationZoneRecordResp {
	return &NullableOSReplicationZoneRecordResp{value: val, isSet: true}
}

func (v NullableOSReplicationZoneRecordResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSReplicationZoneRecordResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


