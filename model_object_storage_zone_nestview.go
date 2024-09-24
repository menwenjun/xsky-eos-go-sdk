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

// checks if the ObjectStorageZoneNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStorageZoneNestview{}

// ObjectStorageZoneNestview struct for ObjectStorageZoneNestview
type ObjectStorageZoneNestview struct {
	Id *int64 `json:"id,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

// NewObjectStorageZoneNestview instantiates a new ObjectStorageZoneNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStorageZoneNestview() *ObjectStorageZoneNestview {
	this := ObjectStorageZoneNestview{}
	return &this
}

// NewObjectStorageZoneNestviewWithDefaults instantiates a new ObjectStorageZoneNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStorageZoneNestviewWithDefaults() *ObjectStorageZoneNestview {
	this := ObjectStorageZoneNestview{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ObjectStorageZoneNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageZoneNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ObjectStorageZoneNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ObjectStorageZoneNestview) SetId(v int64) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ObjectStorageZoneNestview) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageZoneNestview) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ObjectStorageZoneNestview) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ObjectStorageZoneNestview) SetUuid(v string) {
	o.Uuid = &v
}

func (o ObjectStorageZoneNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStorageZoneNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableObjectStorageZoneNestview struct {
	value *ObjectStorageZoneNestview
	isSet bool
}

func (v NullableObjectStorageZoneNestview) Get() *ObjectStorageZoneNestview {
	return v.value
}

func (v *NullableObjectStorageZoneNestview) Set(val *ObjectStorageZoneNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStorageZoneNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStorageZoneNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStorageZoneNestview(val *ObjectStorageZoneNestview) *NullableObjectStorageZoneNestview {
	return &NullableObjectStorageZoneNestview{value: val, isSet: true}
}

func (v NullableObjectStorageZoneNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStorageZoneNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


