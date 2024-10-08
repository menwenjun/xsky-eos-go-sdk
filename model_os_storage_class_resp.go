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

// checks if the OSStorageClassResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSStorageClassResp{}

// OSStorageClassResp struct for OSStorageClassResp
type OSStorageClassResp struct {
	OsStorageClass *OSStorageClass `json:"os_storage_class,omitempty"`
}

// NewOSStorageClassResp instantiates a new OSStorageClassResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSStorageClassResp() *OSStorageClassResp {
	this := OSStorageClassResp{}
	return &this
}

// NewOSStorageClassRespWithDefaults instantiates a new OSStorageClassResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSStorageClassRespWithDefaults() *OSStorageClassResp {
	this := OSStorageClassResp{}
	return &this
}

// GetOsStorageClass returns the OsStorageClass field value if set, zero value otherwise.
func (o *OSStorageClassResp) GetOsStorageClass() OSStorageClass {
	if o == nil || IsNil(o.OsStorageClass) {
		var ret OSStorageClass
		return ret
	}
	return *o.OsStorageClass
}

// GetOsStorageClassOk returns a tuple with the OsStorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSStorageClassResp) GetOsStorageClassOk() (*OSStorageClass, bool) {
	if o == nil || IsNil(o.OsStorageClass) {
		return nil, false
	}
	return o.OsStorageClass, true
}

// HasOsStorageClass returns a boolean if a field has been set.
func (o *OSStorageClassResp) HasOsStorageClass() bool {
	if o != nil && !IsNil(o.OsStorageClass) {
		return true
	}

	return false
}

// SetOsStorageClass gets a reference to the given OSStorageClass and assigns it to the OsStorageClass field.
func (o *OSStorageClassResp) SetOsStorageClass(v OSStorageClass) {
	o.OsStorageClass = &v
}

func (o OSStorageClassResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSStorageClassResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OsStorageClass) {
		toSerialize["os_storage_class"] = o.OsStorageClass
	}
	return toSerialize, nil
}

type NullableOSStorageClassResp struct {
	value *OSStorageClassResp
	isSet bool
}

func (v NullableOSStorageClassResp) Get() *OSStorageClassResp {
	return v.value
}

func (v *NullableOSStorageClassResp) Set(val *OSStorageClassResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOSStorageClassResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOSStorageClassResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSStorageClassResp(val *OSStorageClassResp) *NullableOSStorageClassResp {
	return &NullableOSStorageClassResp{value: val, isSet: true}
}

func (v NullableOSStorageClassResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSStorageClassResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


