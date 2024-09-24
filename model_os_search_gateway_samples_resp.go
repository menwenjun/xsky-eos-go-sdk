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

// checks if the OSSearchGatewaySamplesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSSearchGatewaySamplesResp{}

// OSSearchGatewaySamplesResp struct for OSSearchGatewaySamplesResp
type OSSearchGatewaySamplesResp struct {
	OsSearchGatewaySamples []OSSearchGatewayStat `json:"os_search_gateway_samples"`
}

type _OSSearchGatewaySamplesResp OSSearchGatewaySamplesResp

// NewOSSearchGatewaySamplesResp instantiates a new OSSearchGatewaySamplesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSSearchGatewaySamplesResp(osSearchGatewaySamples []OSSearchGatewayStat) *OSSearchGatewaySamplesResp {
	this := OSSearchGatewaySamplesResp{}
	this.OsSearchGatewaySamples = osSearchGatewaySamples
	return &this
}

// NewOSSearchGatewaySamplesRespWithDefaults instantiates a new OSSearchGatewaySamplesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSSearchGatewaySamplesRespWithDefaults() *OSSearchGatewaySamplesResp {
	this := OSSearchGatewaySamplesResp{}
	return &this
}

// GetOsSearchGatewaySamples returns the OsSearchGatewaySamples field value
func (o *OSSearchGatewaySamplesResp) GetOsSearchGatewaySamples() []OSSearchGatewayStat {
	if o == nil {
		var ret []OSSearchGatewayStat
		return ret
	}

	return o.OsSearchGatewaySamples
}

// GetOsSearchGatewaySamplesOk returns a tuple with the OsSearchGatewaySamples field value
// and a boolean to check if the value has been set.
func (o *OSSearchGatewaySamplesResp) GetOsSearchGatewaySamplesOk() ([]OSSearchGatewayStat, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsSearchGatewaySamples, true
}

// SetOsSearchGatewaySamples sets field value
func (o *OSSearchGatewaySamplesResp) SetOsSearchGatewaySamples(v []OSSearchGatewayStat) {
	o.OsSearchGatewaySamples = v
}

func (o OSSearchGatewaySamplesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSSearchGatewaySamplesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_search_gateway_samples"] = o.OsSearchGatewaySamples
	return toSerialize, nil
}

func (o *OSSearchGatewaySamplesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"os_search_gateway_samples",
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

	varOSSearchGatewaySamplesResp := _OSSearchGatewaySamplesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOSSearchGatewaySamplesResp)

	if err != nil {
		return err
	}

	*o = OSSearchGatewaySamplesResp(varOSSearchGatewaySamplesResp)

	return err
}

type NullableOSSearchGatewaySamplesResp struct {
	value *OSSearchGatewaySamplesResp
	isSet bool
}

func (v NullableOSSearchGatewaySamplesResp) Get() *OSSearchGatewaySamplesResp {
	return v.value
}

func (v *NullableOSSearchGatewaySamplesResp) Set(val *OSSearchGatewaySamplesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOSSearchGatewaySamplesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOSSearchGatewaySamplesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSSearchGatewaySamplesResp(val *OSSearchGatewaySamplesResp) *NullableOSSearchGatewaySamplesResp {
	return &NullableOSSearchGatewaySamplesResp{value: val, isSet: true}
}

func (v NullableOSSearchGatewaySamplesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSSearchGatewaySamplesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


