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

// checks if the OSExternalStorageClassesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSExternalStorageClassesResp{}

// OSExternalStorageClassesResp struct for OSExternalStorageClassesResp
type OSExternalStorageClassesResp struct {
	OsExternalStorageClasses []OSExternalStorageClass `json:"os_external_storage_classes,omitempty"`
}

// NewOSExternalStorageClassesResp instantiates a new OSExternalStorageClassesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSExternalStorageClassesResp() *OSExternalStorageClassesResp {
	this := OSExternalStorageClassesResp{}
	return &this
}

// NewOSExternalStorageClassesRespWithDefaults instantiates a new OSExternalStorageClassesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSExternalStorageClassesRespWithDefaults() *OSExternalStorageClassesResp {
	this := OSExternalStorageClassesResp{}
	return &this
}

// GetOsExternalStorageClasses returns the OsExternalStorageClasses field value if set, zero value otherwise.
func (o *OSExternalStorageClassesResp) GetOsExternalStorageClasses() []OSExternalStorageClass {
	if o == nil || IsNil(o.OsExternalStorageClasses) {
		var ret []OSExternalStorageClass
		return ret
	}
	return o.OsExternalStorageClasses
}

// GetOsExternalStorageClassesOk returns a tuple with the OsExternalStorageClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassesResp) GetOsExternalStorageClassesOk() ([]OSExternalStorageClass, bool) {
	if o == nil || IsNil(o.OsExternalStorageClasses) {
		return nil, false
	}
	return o.OsExternalStorageClasses, true
}

// HasOsExternalStorageClasses returns a boolean if a field has been set.
func (o *OSExternalStorageClassesResp) HasOsExternalStorageClasses() bool {
	if o != nil && !IsNil(o.OsExternalStorageClasses) {
		return true
	}

	return false
}

// SetOsExternalStorageClasses gets a reference to the given []OSExternalStorageClass and assigns it to the OsExternalStorageClasses field.
func (o *OSExternalStorageClassesResp) SetOsExternalStorageClasses(v []OSExternalStorageClass) {
	o.OsExternalStorageClasses = v
}

func (o OSExternalStorageClassesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSExternalStorageClassesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OsExternalStorageClasses) {
		toSerialize["os_external_storage_classes"] = o.OsExternalStorageClasses
	}
	return toSerialize, nil
}

type NullableOSExternalStorageClassesResp struct {
	value *OSExternalStorageClassesResp
	isSet bool
}

func (v NullableOSExternalStorageClassesResp) Get() *OSExternalStorageClassesResp {
	return v.value
}

func (v *NullableOSExternalStorageClassesResp) Set(val *OSExternalStorageClassesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOSExternalStorageClassesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOSExternalStorageClassesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSExternalStorageClassesResp(val *OSExternalStorageClassesResp) *NullableOSExternalStorageClassesResp {
	return &NullableOSExternalStorageClassesResp{value: val, isSet: true}
}

func (v NullableOSExternalStorageClassesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSExternalStorageClassesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


