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

// checks if the RegionResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionResp{}

// RegionResp struct for RegionResp
type RegionResp struct {
	Region *Region `json:"region,omitempty"`
}

// NewRegionResp instantiates a new RegionResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionResp() *RegionResp {
	this := RegionResp{}
	return &this
}

// NewRegionRespWithDefaults instantiates a new RegionResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionRespWithDefaults() *RegionResp {
	this := RegionResp{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *RegionResp) GetRegion() Region {
	if o == nil || IsNil(o.Region) {
		var ret Region
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionResp) GetRegionOk() (*Region, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *RegionResp) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given Region and assigns it to the Region field.
func (o *RegionResp) SetRegion(v Region) {
	o.Region = &v
}

func (o RegionResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegionResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return toSerialize, nil
}

type NullableRegionResp struct {
	value *RegionResp
	isSet bool
}

func (v NullableRegionResp) Get() *RegionResp {
	return v.value
}

func (v *NullableRegionResp) Set(val *RegionResp) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionResp) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionResp(val *RegionResp) *NullableRegionResp {
	return &NullableRegionResp{value: val, isSet: true}
}

func (v NullableRegionResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


