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

// checks if the NgObjectStorageNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NgObjectStorageNestview{}

// NgObjectStorageNestview struct for NgObjectStorageNestview
type NgObjectStorageNestview struct {
	Id *int64 `json:"id,omitempty"`
}

// NewNgObjectStorageNestview instantiates a new NgObjectStorageNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNgObjectStorageNestview() *NgObjectStorageNestview {
	this := NgObjectStorageNestview{}
	return &this
}

// NewNgObjectStorageNestviewWithDefaults instantiates a new NgObjectStorageNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNgObjectStorageNestviewWithDefaults() *NgObjectStorageNestview {
	this := NgObjectStorageNestview{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NgObjectStorageNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgObjectStorageNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NgObjectStorageNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *NgObjectStorageNestview) SetId(v int64) {
	o.Id = &v
}

func (o NgObjectStorageNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NgObjectStorageNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableNgObjectStorageNestview struct {
	value *NgObjectStorageNestview
	isSet bool
}

func (v NullableNgObjectStorageNestview) Get() *NgObjectStorageNestview {
	return v.value
}

func (v *NullableNgObjectStorageNestview) Set(val *NgObjectStorageNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableNgObjectStorageNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableNgObjectStorageNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNgObjectStorageNestview(val *NgObjectStorageNestview) *NullableNgObjectStorageNestview {
	return &NullableNgObjectStorageNestview{value: val, isSet: true}
}

func (v NullableNgObjectStorageNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNgObjectStorageNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


