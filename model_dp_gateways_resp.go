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

// checks if the DpGatewaysResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpGatewaysResp{}

// DpGatewaysResp struct for DpGatewaysResp
type DpGatewaysResp struct {
	// dp gateways
	DpGateways []DpGateway `json:"dp_gateways"`
}

type _DpGatewaysResp DpGatewaysResp

// NewDpGatewaysResp instantiates a new DpGatewaysResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpGatewaysResp(dpGateways []DpGateway) *DpGatewaysResp {
	this := DpGatewaysResp{}
	this.DpGateways = dpGateways
	return &this
}

// NewDpGatewaysRespWithDefaults instantiates a new DpGatewaysResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpGatewaysRespWithDefaults() *DpGatewaysResp {
	this := DpGatewaysResp{}
	return &this
}

// GetDpGateways returns the DpGateways field value
func (o *DpGatewaysResp) GetDpGateways() []DpGateway {
	if o == nil {
		var ret []DpGateway
		return ret
	}

	return o.DpGateways
}

// GetDpGatewaysOk returns a tuple with the DpGateways field value
// and a boolean to check if the value has been set.
func (o *DpGatewaysResp) GetDpGatewaysOk() ([]DpGateway, bool) {
	if o == nil {
		return nil, false
	}
	return o.DpGateways, true
}

// SetDpGateways sets field value
func (o *DpGatewaysResp) SetDpGateways(v []DpGateway) {
	o.DpGateways = v
}

func (o DpGatewaysResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpGatewaysResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dp_gateways"] = o.DpGateways
	return toSerialize, nil
}

func (o *DpGatewaysResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dp_gateways",
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

	varDpGatewaysResp := _DpGatewaysResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDpGatewaysResp)

	if err != nil {
		return err
	}

	*o = DpGatewaysResp(varDpGatewaysResp)

	return err
}

type NullableDpGatewaysResp struct {
	value *DpGatewaysResp
	isSet bool
}

func (v NullableDpGatewaysResp) Get() *DpGatewaysResp {
	return v.value
}

func (v *NullableDpGatewaysResp) Set(val *DpGatewaysResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDpGatewaysResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDpGatewaysResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpGatewaysResp(val *DpGatewaysResp) *NullableDpGatewaysResp {
	return &NullableDpGatewaysResp{value: val, isSet: true}
}

func (v NullableDpGatewaysResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpGatewaysResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


