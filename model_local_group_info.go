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

// checks if the LocalGroupInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalGroupInfo{}

// LocalGroupInfo struct for LocalGroupInfo
type LocalGroupInfo struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewLocalGroupInfo instantiates a new LocalGroupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalGroupInfo() *LocalGroupInfo {
	this := LocalGroupInfo{}
	return &this
}

// NewLocalGroupInfoWithDefaults instantiates a new LocalGroupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalGroupInfoWithDefaults() *LocalGroupInfo {
	this := LocalGroupInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LocalGroupInfo) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalGroupInfo) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LocalGroupInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *LocalGroupInfo) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LocalGroupInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalGroupInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LocalGroupInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LocalGroupInfo) SetName(v string) {
	o.Name = &v
}

func (o LocalGroupInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalGroupInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableLocalGroupInfo struct {
	value *LocalGroupInfo
	isSet bool
}

func (v NullableLocalGroupInfo) Get() *LocalGroupInfo {
	return v.value
}

func (v *NullableLocalGroupInfo) Set(val *LocalGroupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalGroupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalGroupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalGroupInfo(val *LocalGroupInfo) *NullableLocalGroupInfo {
	return &NullableLocalGroupInfo{value: val, isSet: true}
}

func (v NullableLocalGroupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalGroupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


