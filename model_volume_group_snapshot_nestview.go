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

// checks if the VolumeGroupSnapshotNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeGroupSnapshotNestview{}

// VolumeGroupSnapshotNestview struct for VolumeGroupSnapshotNestview
type VolumeGroupSnapshotNestview struct {
	GroupSnapshotName *string `json:"group_snapshot_name,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewVolumeGroupSnapshotNestview instantiates a new VolumeGroupSnapshotNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeGroupSnapshotNestview() *VolumeGroupSnapshotNestview {
	this := VolumeGroupSnapshotNestview{}
	return &this
}

// NewVolumeGroupSnapshotNestviewWithDefaults instantiates a new VolumeGroupSnapshotNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeGroupSnapshotNestviewWithDefaults() *VolumeGroupSnapshotNestview {
	this := VolumeGroupSnapshotNestview{}
	return &this
}

// GetGroupSnapshotName returns the GroupSnapshotName field value if set, zero value otherwise.
func (o *VolumeGroupSnapshotNestview) GetGroupSnapshotName() string {
	if o == nil || IsNil(o.GroupSnapshotName) {
		var ret string
		return ret
	}
	return *o.GroupSnapshotName
}

// GetGroupSnapshotNameOk returns a tuple with the GroupSnapshotName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeGroupSnapshotNestview) GetGroupSnapshotNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupSnapshotName) {
		return nil, false
	}
	return o.GroupSnapshotName, true
}

// HasGroupSnapshotName returns a boolean if a field has been set.
func (o *VolumeGroupSnapshotNestview) HasGroupSnapshotName() bool {
	if o != nil && !IsNil(o.GroupSnapshotName) {
		return true
	}

	return false
}

// SetGroupSnapshotName gets a reference to the given string and assigns it to the GroupSnapshotName field.
func (o *VolumeGroupSnapshotNestview) SetGroupSnapshotName(v string) {
	o.GroupSnapshotName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VolumeGroupSnapshotNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeGroupSnapshotNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VolumeGroupSnapshotNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *VolumeGroupSnapshotNestview) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VolumeGroupSnapshotNestview) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeGroupSnapshotNestview) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VolumeGroupSnapshotNestview) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VolumeGroupSnapshotNestview) SetName(v string) {
	o.Name = &v
}

func (o VolumeGroupSnapshotNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeGroupSnapshotNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupSnapshotName) {
		toSerialize["group_snapshot_name"] = o.GroupSnapshotName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableVolumeGroupSnapshotNestview struct {
	value *VolumeGroupSnapshotNestview
	isSet bool
}

func (v NullableVolumeGroupSnapshotNestview) Get() *VolumeGroupSnapshotNestview {
	return v.value
}

func (v *NullableVolumeGroupSnapshotNestview) Set(val *VolumeGroupSnapshotNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeGroupSnapshotNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeGroupSnapshotNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeGroupSnapshotNestview(val *VolumeGroupSnapshotNestview) *NullableVolumeGroupSnapshotNestview {
	return &NullableVolumeGroupSnapshotNestview{value: val, isSet: true}
}

func (v NullableVolumeGroupSnapshotNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeGroupSnapshotNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


