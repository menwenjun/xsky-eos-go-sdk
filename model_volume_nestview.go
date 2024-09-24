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

// checks if the VolumeNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeNestview{}

// VolumeNestview struct for VolumeNestview
type VolumeNestview struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	OriginalName *string `json:"original_name,omitempty"`
	Uid *string `json:"uid,omitempty"`
}

// NewVolumeNestview instantiates a new VolumeNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeNestview() *VolumeNestview {
	this := VolumeNestview{}
	return &this
}

// NewVolumeNestviewWithDefaults instantiates a new VolumeNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeNestviewWithDefaults() *VolumeNestview {
	this := VolumeNestview{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VolumeNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VolumeNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *VolumeNestview) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VolumeNestview) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeNestview) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VolumeNestview) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VolumeNestview) SetName(v string) {
	o.Name = &v
}

// GetOriginalName returns the OriginalName field value if set, zero value otherwise.
func (o *VolumeNestview) GetOriginalName() string {
	if o == nil || IsNil(o.OriginalName) {
		var ret string
		return ret
	}
	return *o.OriginalName
}

// GetOriginalNameOk returns a tuple with the OriginalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeNestview) GetOriginalNameOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalName) {
		return nil, false
	}
	return o.OriginalName, true
}

// HasOriginalName returns a boolean if a field has been set.
func (o *VolumeNestview) HasOriginalName() bool {
	if o != nil && !IsNil(o.OriginalName) {
		return true
	}

	return false
}

// SetOriginalName gets a reference to the given string and assigns it to the OriginalName field.
func (o *VolumeNestview) SetOriginalName(v string) {
	o.OriginalName = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *VolumeNestview) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeNestview) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *VolumeNestview) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *VolumeNestview) SetUid(v string) {
	o.Uid = &v
}

func (o VolumeNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OriginalName) {
		toSerialize["original_name"] = o.OriginalName
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableVolumeNestview struct {
	value *VolumeNestview
	isSet bool
}

func (v NullableVolumeNestview) Get() *VolumeNestview {
	return v.value
}

func (v *NullableVolumeNestview) Set(val *VolumeNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeNestview(val *VolumeNestview) *NullableVolumeNestview {
	return &NullableVolumeNestview{value: val, isSet: true}
}

func (v NullableVolumeNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


