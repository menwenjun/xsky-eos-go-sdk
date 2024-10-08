/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RegionUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionUpdateReq{}

// RegionUpdateReq struct for RegionUpdateReq
type RegionUpdateReq struct {
	Region RegionUpdateReqRegion `json:"region"`
}

type _RegionUpdateReq RegionUpdateReq

// NewRegionUpdateReq instantiates a new RegionUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionUpdateReq(region RegionUpdateReqRegion) *RegionUpdateReq {
	this := RegionUpdateReq{}
	this.Region = region
	return &this
}

// NewRegionUpdateReqWithDefaults instantiates a new RegionUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionUpdateReqWithDefaults() *RegionUpdateReq {
	this := RegionUpdateReq{}
	return &this
}

// GetRegion returns the Region field value
func (o *RegionUpdateReq) GetRegion() RegionUpdateReqRegion {
	if o == nil {
		var ret RegionUpdateReqRegion
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *RegionUpdateReq) GetRegionOk() (*RegionUpdateReqRegion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *RegionUpdateReq) SetRegion(v RegionUpdateReqRegion) {
	o.Region = v
}

func (o RegionUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegionUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["region"] = o.Region
	return toSerialize, nil
}

func (o *RegionUpdateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"region",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRegionUpdateReq := _RegionUpdateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegionUpdateReq)

	if err != nil {
		return err
	}

	*o = RegionUpdateReq(varRegionUpdateReq)

	return err
}

type NullableRegionUpdateReq struct {
	value *RegionUpdateReq
	isSet bool
}

func (v NullableRegionUpdateReq) Get() *RegionUpdateReq {
	return v.value
}

func (v *NullableRegionUpdateReq) Set(val *RegionUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionUpdateReq(val *RegionUpdateReq) *NullableRegionUpdateReq {
	return &NullableRegionUpdateReq{value: val, isSet: true}
}

func (v NullableRegionUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


