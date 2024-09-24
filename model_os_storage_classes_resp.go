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

// checks if the OSStorageClassesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSStorageClassesResp{}

// OSStorageClassesResp struct for OSStorageClassesResp
type OSStorageClassesResp struct {
	OsStorageClasses []OSStorageClass `json:"os_storage_classes,omitempty"`
}

// NewOSStorageClassesResp instantiates a new OSStorageClassesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSStorageClassesResp() *OSStorageClassesResp {
	this := OSStorageClassesResp{}
	return &this
}

// NewOSStorageClassesRespWithDefaults instantiates a new OSStorageClassesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSStorageClassesRespWithDefaults() *OSStorageClassesResp {
	this := OSStorageClassesResp{}
	return &this
}

// GetOsStorageClasses returns the OsStorageClasses field value if set, zero value otherwise.
func (o *OSStorageClassesResp) GetOsStorageClasses() []OSStorageClass {
	if o == nil || IsNil(o.OsStorageClasses) {
		var ret []OSStorageClass
		return ret
	}
	return o.OsStorageClasses
}

// GetOsStorageClassesOk returns a tuple with the OsStorageClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSStorageClassesResp) GetOsStorageClassesOk() ([]OSStorageClass, bool) {
	if o == nil || IsNil(o.OsStorageClasses) {
		return nil, false
	}
	return o.OsStorageClasses, true
}

// HasOsStorageClasses returns a boolean if a field has been set.
func (o *OSStorageClassesResp) HasOsStorageClasses() bool {
	if o != nil && !IsNil(o.OsStorageClasses) {
		return true
	}

	return false
}

// SetOsStorageClasses gets a reference to the given []OSStorageClass and assigns it to the OsStorageClasses field.
func (o *OSStorageClassesResp) SetOsStorageClasses(v []OSStorageClass) {
	o.OsStorageClasses = v
}

func (o OSStorageClassesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSStorageClassesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OsStorageClasses) {
		toSerialize["os_storage_classes"] = o.OsStorageClasses
	}
	return toSerialize, nil
}

type NullableOSStorageClassesResp struct {
	value *OSStorageClassesResp
	isSet bool
}

func (v NullableOSStorageClassesResp) Get() *OSStorageClassesResp {
	return v.value
}

func (v *NullableOSStorageClassesResp) Set(val *OSStorageClassesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOSStorageClassesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOSStorageClassesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSStorageClassesResp(val *OSStorageClassesResp) *NullableOSStorageClassesResp {
	return &NullableOSStorageClassesResp{value: val, isSet: true}
}

func (v NullableOSStorageClassesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSStorageClassesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


