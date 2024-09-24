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

// checks if the VolumeSamplesElem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeSamplesElem{}

// VolumeSamplesElem struct for VolumeSamplesElem
type VolumeSamplesElem struct {
	Id *int64 `json:"id,omitempty"`
	Paging *Paging `json:"paging,omitempty"`
	Samples []VolumeStat `json:"samples,omitempty"`
}

// NewVolumeSamplesElem instantiates a new VolumeSamplesElem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeSamplesElem() *VolumeSamplesElem {
	this := VolumeSamplesElem{}
	return &this
}

// NewVolumeSamplesElemWithDefaults instantiates a new VolumeSamplesElem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeSamplesElemWithDefaults() *VolumeSamplesElem {
	this := VolumeSamplesElem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VolumeSamplesElem) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeSamplesElem) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VolumeSamplesElem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *VolumeSamplesElem) SetId(v int64) {
	o.Id = &v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *VolumeSamplesElem) GetPaging() Paging {
	if o == nil || IsNil(o.Paging) {
		var ret Paging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeSamplesElem) GetPagingOk() (*Paging, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *VolumeSamplesElem) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given Paging and assigns it to the Paging field.
func (o *VolumeSamplesElem) SetPaging(v Paging) {
	o.Paging = &v
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *VolumeSamplesElem) GetSamples() []VolumeStat {
	if o == nil || IsNil(o.Samples) {
		var ret []VolumeStat
		return ret
	}
	return o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeSamplesElem) GetSamplesOk() ([]VolumeStat, bool) {
	if o == nil || IsNil(o.Samples) {
		return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *VolumeSamplesElem) HasSamples() bool {
	if o != nil && !IsNil(o.Samples) {
		return true
	}

	return false
}

// SetSamples gets a reference to the given []VolumeStat and assigns it to the Samples field.
func (o *VolumeSamplesElem) SetSamples(v []VolumeStat) {
	o.Samples = v
}

func (o VolumeSamplesElem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeSamplesElem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	if !IsNil(o.Samples) {
		toSerialize["samples"] = o.Samples
	}
	return toSerialize, nil
}

type NullableVolumeSamplesElem struct {
	value *VolumeSamplesElem
	isSet bool
}

func (v NullableVolumeSamplesElem) Get() *VolumeSamplesElem {
	return v.value
}

func (v *NullableVolumeSamplesElem) Set(val *VolumeSamplesElem) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeSamplesElem) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeSamplesElem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeSamplesElem(val *VolumeSamplesElem) *NullableVolumeSamplesElem {
	return &NullableVolumeSamplesElem{value: val, isSet: true}
}

func (v NullableVolumeSamplesElem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeSamplesElem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


