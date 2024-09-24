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

// checks if the NestedTrash type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedTrash{}

// NestedTrash struct for NestedTrash
type NestedTrash struct {
	// trash id
	Id *int64 `json:"id,omitempty"`
	// trash path
	Path *string `json:"path,omitempty"`
}

// NewNestedTrash instantiates a new NestedTrash object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedTrash() *NestedTrash {
	this := NestedTrash{}
	return &this
}

// NewNestedTrashWithDefaults instantiates a new NestedTrash object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedTrashWithDefaults() *NestedTrash {
	this := NestedTrash{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NestedTrash) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedTrash) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NestedTrash) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *NestedTrash) SetId(v int64) {
	o.Id = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *NestedTrash) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedTrash) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *NestedTrash) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *NestedTrash) SetPath(v string) {
	o.Path = &v
}

func (o NestedTrash) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedTrash) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableNestedTrash struct {
	value *NestedTrash
	isSet bool
}

func (v NullableNestedTrash) Get() *NestedTrash {
	return v.value
}

func (v *NullableNestedTrash) Set(val *NestedTrash) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedTrash) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedTrash) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedTrash(val *NestedTrash) *NullableNestedTrash {
	return &NullableNestedTrash{value: val, isSet: true}
}

func (v NullableNestedTrash) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedTrash) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


