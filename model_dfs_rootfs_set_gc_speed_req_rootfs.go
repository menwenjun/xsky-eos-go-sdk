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

// checks if the DfsRootfsSetGCSpeedReqRootfs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsRootfsSetGCSpeedReqRootfs{}

// DfsRootfsSetGCSpeedReqRootfs struct for DfsRootfsSetGCSpeedReqRootfs
type DfsRootfsSetGCSpeedReqRootfs struct {
	GcSpeed *string `json:"gc_speed,omitempty"`
}

// NewDfsRootfsSetGCSpeedReqRootfs instantiates a new DfsRootfsSetGCSpeedReqRootfs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsRootfsSetGCSpeedReqRootfs() *DfsRootfsSetGCSpeedReqRootfs {
	this := DfsRootfsSetGCSpeedReqRootfs{}
	return &this
}

// NewDfsRootfsSetGCSpeedReqRootfsWithDefaults instantiates a new DfsRootfsSetGCSpeedReqRootfs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsRootfsSetGCSpeedReqRootfsWithDefaults() *DfsRootfsSetGCSpeedReqRootfs {
	this := DfsRootfsSetGCSpeedReqRootfs{}
	return &this
}

// GetGcSpeed returns the GcSpeed field value if set, zero value otherwise.
func (o *DfsRootfsSetGCSpeedReqRootfs) GetGcSpeed() string {
	if o == nil || IsNil(o.GcSpeed) {
		var ret string
		return ret
	}
	return *o.GcSpeed
}

// GetGcSpeedOk returns a tuple with the GcSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsRootfsSetGCSpeedReqRootfs) GetGcSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.GcSpeed) {
		return nil, false
	}
	return o.GcSpeed, true
}

// HasGcSpeed returns a boolean if a field has been set.
func (o *DfsRootfsSetGCSpeedReqRootfs) HasGcSpeed() bool {
	if o != nil && !IsNil(o.GcSpeed) {
		return true
	}

	return false
}

// SetGcSpeed gets a reference to the given string and assigns it to the GcSpeed field.
func (o *DfsRootfsSetGCSpeedReqRootfs) SetGcSpeed(v string) {
	o.GcSpeed = &v
}

func (o DfsRootfsSetGCSpeedReqRootfs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsRootfsSetGCSpeedReqRootfs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GcSpeed) {
		toSerialize["gc_speed"] = o.GcSpeed
	}
	return toSerialize, nil
}

type NullableDfsRootfsSetGCSpeedReqRootfs struct {
	value *DfsRootfsSetGCSpeedReqRootfs
	isSet bool
}

func (v NullableDfsRootfsSetGCSpeedReqRootfs) Get() *DfsRootfsSetGCSpeedReqRootfs {
	return v.value
}

func (v *NullableDfsRootfsSetGCSpeedReqRootfs) Set(val *DfsRootfsSetGCSpeedReqRootfs) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsRootfsSetGCSpeedReqRootfs) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsRootfsSetGCSpeedReqRootfs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsRootfsSetGCSpeedReqRootfs(val *DfsRootfsSetGCSpeedReqRootfs) *NullableDfsRootfsSetGCSpeedReqRootfs {
	return &NullableDfsRootfsSetGCSpeedReqRootfs{value: val, isSet: true}
}

func (v NullableDfsRootfsSetGCSpeedReqRootfs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsRootfsSetGCSpeedReqRootfs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


