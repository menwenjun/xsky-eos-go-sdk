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

// checks if the OSSearchEngineSamplesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSSearchEngineSamplesResp{}

// OSSearchEngineSamplesResp struct for OSSearchEngineSamplesResp
type OSSearchEngineSamplesResp struct {
	// os search engine samples
	OsSearchEngineSamples []OSSearchEngineStat `json:"os_search_engine_samples"`
}

type _OSSearchEngineSamplesResp OSSearchEngineSamplesResp

// NewOSSearchEngineSamplesResp instantiates a new OSSearchEngineSamplesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSSearchEngineSamplesResp(osSearchEngineSamples []OSSearchEngineStat) *OSSearchEngineSamplesResp {
	this := OSSearchEngineSamplesResp{}
	this.OsSearchEngineSamples = osSearchEngineSamples
	return &this
}

// NewOSSearchEngineSamplesRespWithDefaults instantiates a new OSSearchEngineSamplesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSSearchEngineSamplesRespWithDefaults() *OSSearchEngineSamplesResp {
	this := OSSearchEngineSamplesResp{}
	return &this
}

// GetOsSearchEngineSamples returns the OsSearchEngineSamples field value
func (o *OSSearchEngineSamplesResp) GetOsSearchEngineSamples() []OSSearchEngineStat {
	if o == nil {
		var ret []OSSearchEngineStat
		return ret
	}

	return o.OsSearchEngineSamples
}

// GetOsSearchEngineSamplesOk returns a tuple with the OsSearchEngineSamples field value
// and a boolean to check if the value has been set.
func (o *OSSearchEngineSamplesResp) GetOsSearchEngineSamplesOk() ([]OSSearchEngineStat, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsSearchEngineSamples, true
}

// SetOsSearchEngineSamples sets field value
func (o *OSSearchEngineSamplesResp) SetOsSearchEngineSamples(v []OSSearchEngineStat) {
	o.OsSearchEngineSamples = v
}

func (o OSSearchEngineSamplesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSSearchEngineSamplesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_search_engine_samples"] = o.OsSearchEngineSamples
	return toSerialize, nil
}

func (o *OSSearchEngineSamplesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"os_search_engine_samples",
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

	varOSSearchEngineSamplesResp := _OSSearchEngineSamplesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOSSearchEngineSamplesResp)

	if err != nil {
		return err
	}

	*o = OSSearchEngineSamplesResp(varOSSearchEngineSamplesResp)

	return err
}

type NullableOSSearchEngineSamplesResp struct {
	value *OSSearchEngineSamplesResp
	isSet bool
}

func (v NullableOSSearchEngineSamplesResp) Get() *OSSearchEngineSamplesResp {
	return v.value
}

func (v *NullableOSSearchEngineSamplesResp) Set(val *OSSearchEngineSamplesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOSSearchEngineSamplesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOSSearchEngineSamplesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSSearchEngineSamplesResp(val *OSSearchEngineSamplesResp) *NullableOSSearchEngineSamplesResp {
	return &NullableOSSearchEngineSamplesResp{value: val, isSet: true}
}

func (v NullableOSSearchEngineSamplesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSSearchEngineSamplesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


