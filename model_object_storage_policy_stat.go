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

// checks if the ObjectStoragePolicyStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStoragePolicyStat{}

// ObjectStoragePolicyStat ObjectStoragePolicyStat define the object storage policy statistics of a user
type ObjectStoragePolicyStat struct {
	StorageClasses *map[string]OSStorageClassStat `json:"storage_classes,omitempty"`
}

// NewObjectStoragePolicyStat instantiates a new ObjectStoragePolicyStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStoragePolicyStat() *ObjectStoragePolicyStat {
	this := ObjectStoragePolicyStat{}
	return &this
}

// NewObjectStoragePolicyStatWithDefaults instantiates a new ObjectStoragePolicyStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStoragePolicyStatWithDefaults() *ObjectStoragePolicyStat {
	this := ObjectStoragePolicyStat{}
	return &this
}

// GetStorageClasses returns the StorageClasses field value if set, zero value otherwise.
func (o *ObjectStoragePolicyStat) GetStorageClasses() map[string]OSStorageClassStat {
	if o == nil || IsNil(o.StorageClasses) {
		var ret map[string]OSStorageClassStat
		return ret
	}
	return *o.StorageClasses
}

// GetStorageClassesOk returns a tuple with the StorageClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePolicyStat) GetStorageClassesOk() (*map[string]OSStorageClassStat, bool) {
	if o == nil || IsNil(o.StorageClasses) {
		return nil, false
	}
	return o.StorageClasses, true
}

// HasStorageClasses returns a boolean if a field has been set.
func (o *ObjectStoragePolicyStat) HasStorageClasses() bool {
	if o != nil && !IsNil(o.StorageClasses) {
		return true
	}

	return false
}

// SetStorageClasses gets a reference to the given map[string]OSStorageClassStat and assigns it to the StorageClasses field.
func (o *ObjectStoragePolicyStat) SetStorageClasses(v map[string]OSStorageClassStat) {
	o.StorageClasses = &v
}

func (o ObjectStoragePolicyStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStoragePolicyStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StorageClasses) {
		toSerialize["storage_classes"] = o.StorageClasses
	}
	return toSerialize, nil
}

type NullableObjectStoragePolicyStat struct {
	value *ObjectStoragePolicyStat
	isSet bool
}

func (v NullableObjectStoragePolicyStat) Get() *ObjectStoragePolicyStat {
	return v.value
}

func (v *NullableObjectStoragePolicyStat) Set(val *ObjectStoragePolicyStat) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStoragePolicyStat) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStoragePolicyStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStoragePolicyStat(val *ObjectStoragePolicyStat) *NullableObjectStoragePolicyStat {
	return &NullableObjectStoragePolicyStat{value: val, isSet: true}
}

func (v NullableObjectStoragePolicyStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStoragePolicyStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


