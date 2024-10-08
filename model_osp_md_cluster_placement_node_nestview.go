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

// checks if the OspMdClusterPlacementNodeNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OspMdClusterPlacementNodeNestview{}

// OspMdClusterPlacementNodeNestview struct for OspMdClusterPlacementNodeNestview
type OspMdClusterPlacementNodeNestview struct {
	Id *int64 `json:"id,omitempty"`
}

// NewOspMdClusterPlacementNodeNestview instantiates a new OspMdClusterPlacementNodeNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOspMdClusterPlacementNodeNestview() *OspMdClusterPlacementNodeNestview {
	this := OspMdClusterPlacementNodeNestview{}
	return &this
}

// NewOspMdClusterPlacementNodeNestviewWithDefaults instantiates a new OspMdClusterPlacementNodeNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOspMdClusterPlacementNodeNestviewWithDefaults() *OspMdClusterPlacementNodeNestview {
	this := OspMdClusterPlacementNodeNestview{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OspMdClusterPlacementNodeNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMdClusterPlacementNodeNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OspMdClusterPlacementNodeNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OspMdClusterPlacementNodeNestview) SetId(v int64) {
	o.Id = &v
}

func (o OspMdClusterPlacementNodeNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OspMdClusterPlacementNodeNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableOspMdClusterPlacementNodeNestview struct {
	value *OspMdClusterPlacementNodeNestview
	isSet bool
}

func (v NullableOspMdClusterPlacementNodeNestview) Get() *OspMdClusterPlacementNodeNestview {
	return v.value
}

func (v *NullableOspMdClusterPlacementNodeNestview) Set(val *OspMdClusterPlacementNodeNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableOspMdClusterPlacementNodeNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableOspMdClusterPlacementNodeNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspMdClusterPlacementNodeNestview(val *OspMdClusterPlacementNodeNestview) *NullableOspMdClusterPlacementNodeNestview {
	return &NullableOspMdClusterPlacementNodeNestview{value: val, isSet: true}
}

func (v NullableOspMdClusterPlacementNodeNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspMdClusterPlacementNodeNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


