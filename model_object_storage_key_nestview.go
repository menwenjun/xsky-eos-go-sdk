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

// checks if the ObjectStorageKeyNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStorageKeyNestview{}

// ObjectStorageKeyNestview struct for ObjectStorageKeyNestview
type ObjectStorageKeyNestview struct {
	Id *int64 `json:"id,omitempty"`
}

// NewObjectStorageKeyNestview instantiates a new ObjectStorageKeyNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStorageKeyNestview() *ObjectStorageKeyNestview {
	this := ObjectStorageKeyNestview{}
	return &this
}

// NewObjectStorageKeyNestviewWithDefaults instantiates a new ObjectStorageKeyNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStorageKeyNestviewWithDefaults() *ObjectStorageKeyNestview {
	this := ObjectStorageKeyNestview{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ObjectStorageKeyNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageKeyNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ObjectStorageKeyNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ObjectStorageKeyNestview) SetId(v int64) {
	o.Id = &v
}

func (o ObjectStorageKeyNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStorageKeyNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableObjectStorageKeyNestview struct {
	value *ObjectStorageKeyNestview
	isSet bool
}

func (v NullableObjectStorageKeyNestview) Get() *ObjectStorageKeyNestview {
	return v.value
}

func (v *NullableObjectStorageKeyNestview) Set(val *ObjectStorageKeyNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStorageKeyNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStorageKeyNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStorageKeyNestview(val *ObjectStorageKeyNestview) *NullableObjectStorageKeyNestview {
	return &NullableObjectStorageKeyNestview{value: val, isSet: true}
}

func (v NullableObjectStorageKeyNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStorageKeyNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


