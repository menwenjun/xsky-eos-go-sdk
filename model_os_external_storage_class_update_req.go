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

// checks if the OSExternalStorageClassUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSExternalStorageClassUpdateReq{}

// OSExternalStorageClassUpdateReq struct for OSExternalStorageClassUpdateReq
type OSExternalStorageClassUpdateReq struct {
	OsExternalStorageClass *OSExternalStorageClassUpdateReqExternalStorageClass `json:"os_external_storage_class,omitempty"`
}

// NewOSExternalStorageClassUpdateReq instantiates a new OSExternalStorageClassUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSExternalStorageClassUpdateReq() *OSExternalStorageClassUpdateReq {
	this := OSExternalStorageClassUpdateReq{}
	return &this
}

// NewOSExternalStorageClassUpdateReqWithDefaults instantiates a new OSExternalStorageClassUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSExternalStorageClassUpdateReqWithDefaults() *OSExternalStorageClassUpdateReq {
	this := OSExternalStorageClassUpdateReq{}
	return &this
}

// GetOsExternalStorageClass returns the OsExternalStorageClass field value if set, zero value otherwise.
func (o *OSExternalStorageClassUpdateReq) GetOsExternalStorageClass() OSExternalStorageClassUpdateReqExternalStorageClass {
	if o == nil || IsNil(o.OsExternalStorageClass) {
		var ret OSExternalStorageClassUpdateReqExternalStorageClass
		return ret
	}
	return *o.OsExternalStorageClass
}

// GetOsExternalStorageClassOk returns a tuple with the OsExternalStorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassUpdateReq) GetOsExternalStorageClassOk() (*OSExternalStorageClassUpdateReqExternalStorageClass, bool) {
	if o == nil || IsNil(o.OsExternalStorageClass) {
		return nil, false
	}
	return o.OsExternalStorageClass, true
}

// HasOsExternalStorageClass returns a boolean if a field has been set.
func (o *OSExternalStorageClassUpdateReq) HasOsExternalStorageClass() bool {
	if o != nil && !IsNil(o.OsExternalStorageClass) {
		return true
	}

	return false
}

// SetOsExternalStorageClass gets a reference to the given OSExternalStorageClassUpdateReqExternalStorageClass and assigns it to the OsExternalStorageClass field.
func (o *OSExternalStorageClassUpdateReq) SetOsExternalStorageClass(v OSExternalStorageClassUpdateReqExternalStorageClass) {
	o.OsExternalStorageClass = &v
}

func (o OSExternalStorageClassUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSExternalStorageClassUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OsExternalStorageClass) {
		toSerialize["os_external_storage_class"] = o.OsExternalStorageClass
	}
	return toSerialize, nil
}

type NullableOSExternalStorageClassUpdateReq struct {
	value *OSExternalStorageClassUpdateReq
	isSet bool
}

func (v NullableOSExternalStorageClassUpdateReq) Get() *OSExternalStorageClassUpdateReq {
	return v.value
}

func (v *NullableOSExternalStorageClassUpdateReq) Set(val *OSExternalStorageClassUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableOSExternalStorageClassUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableOSExternalStorageClassUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSExternalStorageClassUpdateReq(val *OSExternalStorageClassUpdateReq) *NullableOSExternalStorageClassUpdateReq {
	return &NullableOSExternalStorageClassUpdateReq{value: val, isSet: true}
}

func (v NullableOSExternalStorageClassUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSExternalStorageClassUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


