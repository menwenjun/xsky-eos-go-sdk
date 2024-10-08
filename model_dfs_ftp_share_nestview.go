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

// checks if the DfsFTPShareNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsFTPShareNestview{}

// DfsFTPShareNestview struct for DfsFTPShareNestview
type DfsFTPShareNestview struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewDfsFTPShareNestview instantiates a new DfsFTPShareNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsFTPShareNestview() *DfsFTPShareNestview {
	this := DfsFTPShareNestview{}
	return &this
}

// NewDfsFTPShareNestviewWithDefaults instantiates a new DfsFTPShareNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsFTPShareNestviewWithDefaults() *DfsFTPShareNestview {
	this := DfsFTPShareNestview{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsFTPShareNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsFTPShareNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsFTPShareNestview) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DfsFTPShareNestview) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareNestview) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DfsFTPShareNestview) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DfsFTPShareNestview) SetName(v string) {
	o.Name = &v
}

func (o DfsFTPShareNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsFTPShareNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableDfsFTPShareNestview struct {
	value *DfsFTPShareNestview
	isSet bool
}

func (v NullableDfsFTPShareNestview) Get() *DfsFTPShareNestview {
	return v.value
}

func (v *NullableDfsFTPShareNestview) Set(val *DfsFTPShareNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsFTPShareNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsFTPShareNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsFTPShareNestview(val *DfsFTPShareNestview) *NullableDfsFTPShareNestview {
	return &NullableDfsFTPShareNestview{value: val, isSet: true}
}

func (v NullableDfsFTPShareNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsFTPShareNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


