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

// checks if the OsdGroupSamplesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsdGroupSamplesResp{}

// OsdGroupSamplesResp struct for OsdGroupSamplesResp
type OsdGroupSamplesResp struct {
	OsdGroupSamples []OsdGroupStat `json:"osd_group_samples"`
}

type _OsdGroupSamplesResp OsdGroupSamplesResp

// NewOsdGroupSamplesResp instantiates a new OsdGroupSamplesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsdGroupSamplesResp(osdGroupSamples []OsdGroupStat) *OsdGroupSamplesResp {
	this := OsdGroupSamplesResp{}
	this.OsdGroupSamples = osdGroupSamples
	return &this
}

// NewOsdGroupSamplesRespWithDefaults instantiates a new OsdGroupSamplesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsdGroupSamplesRespWithDefaults() *OsdGroupSamplesResp {
	this := OsdGroupSamplesResp{}
	return &this
}

// GetOsdGroupSamples returns the OsdGroupSamples field value
func (o *OsdGroupSamplesResp) GetOsdGroupSamples() []OsdGroupStat {
	if o == nil {
		var ret []OsdGroupStat
		return ret
	}

	return o.OsdGroupSamples
}

// GetOsdGroupSamplesOk returns a tuple with the OsdGroupSamples field value
// and a boolean to check if the value has been set.
func (o *OsdGroupSamplesResp) GetOsdGroupSamplesOk() ([]OsdGroupStat, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsdGroupSamples, true
}

// SetOsdGroupSamples sets field value
func (o *OsdGroupSamplesResp) SetOsdGroupSamples(v []OsdGroupStat) {
	o.OsdGroupSamples = v
}

func (o OsdGroupSamplesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsdGroupSamplesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["osd_group_samples"] = o.OsdGroupSamples
	return toSerialize, nil
}

func (o *OsdGroupSamplesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"osd_group_samples",
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

	varOsdGroupSamplesResp := _OsdGroupSamplesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOsdGroupSamplesResp)

	if err != nil {
		return err
	}

	*o = OsdGroupSamplesResp(varOsdGroupSamplesResp)

	return err
}

type NullableOsdGroupSamplesResp struct {
	value *OsdGroupSamplesResp
	isSet bool
}

func (v NullableOsdGroupSamplesResp) Get() *OsdGroupSamplesResp {
	return v.value
}

func (v *NullableOsdGroupSamplesResp) Set(val *OsdGroupSamplesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOsdGroupSamplesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOsdGroupSamplesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsdGroupSamplesResp(val *OsdGroupSamplesResp) *NullableOsdGroupSamplesResp {
	return &NullableOsdGroupSamplesResp{value: val, isSet: true}
}

func (v NullableOsdGroupSamplesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsdGroupSamplesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


