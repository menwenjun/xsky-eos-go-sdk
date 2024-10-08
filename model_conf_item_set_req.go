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

// checks if the ConfItemSetReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfItemSetReq{}

// ConfItemSetReq struct for ConfItemSetReq
type ConfItemSetReq struct {
	Conf ConfItemSetReqConfItem `json:"conf"`
}

type _ConfItemSetReq ConfItemSetReq

// NewConfItemSetReq instantiates a new ConfItemSetReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfItemSetReq(conf ConfItemSetReqConfItem) *ConfItemSetReq {
	this := ConfItemSetReq{}
	this.Conf = conf
	return &this
}

// NewConfItemSetReqWithDefaults instantiates a new ConfItemSetReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfItemSetReqWithDefaults() *ConfItemSetReq {
	this := ConfItemSetReq{}
	return &this
}

// GetConf returns the Conf field value
func (o *ConfItemSetReq) GetConf() ConfItemSetReqConfItem {
	if o == nil {
		var ret ConfItemSetReqConfItem
		return ret
	}

	return o.Conf
}

// GetConfOk returns a tuple with the Conf field value
// and a boolean to check if the value has been set.
func (o *ConfItemSetReq) GetConfOk() (*ConfItemSetReqConfItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conf, true
}

// SetConf sets field value
func (o *ConfItemSetReq) SetConf(v ConfItemSetReqConfItem) {
	o.Conf = v
}

func (o ConfItemSetReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfItemSetReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conf"] = o.Conf
	return toSerialize, nil
}

func (o *ConfItemSetReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"conf",
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

	varConfItemSetReq := _ConfItemSetReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConfItemSetReq)

	if err != nil {
		return err
	}

	*o = ConfItemSetReq(varConfItemSetReq)

	return err
}

type NullableConfItemSetReq struct {
	value *ConfItemSetReq
	isSet bool
}

func (v NullableConfItemSetReq) Get() *ConfItemSetReq {
	return v.value
}

func (v *NullableConfItemSetReq) Set(val *ConfItemSetReq) {
	v.value = val
	v.isSet = true
}

func (v NullableConfItemSetReq) IsSet() bool {
	return v.isSet
}

func (v *NullableConfItemSetReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfItemSetReq(val *ConfItemSetReq) *NullableConfItemSetReq {
	return &NullableConfItemSetReq{value: val, isSet: true}
}

func (v NullableConfItemSetReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfItemSetReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


