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

// checks if the OSZonePairsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSZonePairsResp{}

// OSZonePairsResp struct for OSZonePairsResp
type OSZonePairsResp struct {
	OsZonePairs []OSZonePair `json:"os_zone_pairs"`
}

type _OSZonePairsResp OSZonePairsResp

// NewOSZonePairsResp instantiates a new OSZonePairsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSZonePairsResp(osZonePairs []OSZonePair) *OSZonePairsResp {
	this := OSZonePairsResp{}
	this.OsZonePairs = osZonePairs
	return &this
}

// NewOSZonePairsRespWithDefaults instantiates a new OSZonePairsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSZonePairsRespWithDefaults() *OSZonePairsResp {
	this := OSZonePairsResp{}
	return &this
}

// GetOsZonePairs returns the OsZonePairs field value
func (o *OSZonePairsResp) GetOsZonePairs() []OSZonePair {
	if o == nil {
		var ret []OSZonePair
		return ret
	}

	return o.OsZonePairs
}

// GetOsZonePairsOk returns a tuple with the OsZonePairs field value
// and a boolean to check if the value has been set.
func (o *OSZonePairsResp) GetOsZonePairsOk() ([]OSZonePair, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsZonePairs, true
}

// SetOsZonePairs sets field value
func (o *OSZonePairsResp) SetOsZonePairs(v []OSZonePair) {
	o.OsZonePairs = v
}

func (o OSZonePairsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSZonePairsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_zone_pairs"] = o.OsZonePairs
	return toSerialize, nil
}

func (o *OSZonePairsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"os_zone_pairs",
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

	varOSZonePairsResp := _OSZonePairsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOSZonePairsResp)

	if err != nil {
		return err
	}

	*o = OSZonePairsResp(varOSZonePairsResp)

	return err
}

type NullableOSZonePairsResp struct {
	value *OSZonePairsResp
	isSet bool
}

func (v NullableOSZonePairsResp) Get() *OSZonePairsResp {
	return v.value
}

func (v *NullableOSZonePairsResp) Set(val *OSZonePairsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOSZonePairsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOSZonePairsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSZonePairsResp(val *OSZonePairsResp) *NullableOSZonePairsResp {
	return &NullableOSZonePairsResp{value: val, isSet: true}
}

func (v NullableOSZonePairsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSZonePairsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


