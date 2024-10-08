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

// checks if the OspGatewaySamplesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OspGatewaySamplesResp{}

// OspGatewaySamplesResp struct for OspGatewaySamplesResp
type OspGatewaySamplesResp struct {
	// osp gateway samples
	OspGatewaySamples []OspGatewayStat `json:"osp_gateway_samples"`
}

type _OspGatewaySamplesResp OspGatewaySamplesResp

// NewOspGatewaySamplesResp instantiates a new OspGatewaySamplesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOspGatewaySamplesResp(ospGatewaySamples []OspGatewayStat) *OspGatewaySamplesResp {
	this := OspGatewaySamplesResp{}
	this.OspGatewaySamples = ospGatewaySamples
	return &this
}

// NewOspGatewaySamplesRespWithDefaults instantiates a new OspGatewaySamplesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOspGatewaySamplesRespWithDefaults() *OspGatewaySamplesResp {
	this := OspGatewaySamplesResp{}
	return &this
}

// GetOspGatewaySamples returns the OspGatewaySamples field value
func (o *OspGatewaySamplesResp) GetOspGatewaySamples() []OspGatewayStat {
	if o == nil {
		var ret []OspGatewayStat
		return ret
	}

	return o.OspGatewaySamples
}

// GetOspGatewaySamplesOk returns a tuple with the OspGatewaySamples field value
// and a boolean to check if the value has been set.
func (o *OspGatewaySamplesResp) GetOspGatewaySamplesOk() ([]OspGatewayStat, bool) {
	if o == nil {
		return nil, false
	}
	return o.OspGatewaySamples, true
}

// SetOspGatewaySamples sets field value
func (o *OspGatewaySamplesResp) SetOspGatewaySamples(v []OspGatewayStat) {
	o.OspGatewaySamples = v
}

func (o OspGatewaySamplesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OspGatewaySamplesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["osp_gateway_samples"] = o.OspGatewaySamples
	return toSerialize, nil
}

func (o *OspGatewaySamplesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"osp_gateway_samples",
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

	varOspGatewaySamplesResp := _OspGatewaySamplesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOspGatewaySamplesResp)

	if err != nil {
		return err
	}

	*o = OspGatewaySamplesResp(varOspGatewaySamplesResp)

	return err
}

type NullableOspGatewaySamplesResp struct {
	value *OspGatewaySamplesResp
	isSet bool
}

func (v NullableOspGatewaySamplesResp) Get() *OspGatewaySamplesResp {
	return v.value
}

func (v *NullableOspGatewaySamplesResp) Set(val *OspGatewaySamplesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOspGatewaySamplesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOspGatewaySamplesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspGatewaySamplesResp(val *OspGatewaySamplesResp) *NullableOspGatewaySamplesResp {
	return &NullableOspGatewaySamplesResp{value: val, isSet: true}
}

func (v NullableOspGatewaySamplesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspGatewaySamplesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


