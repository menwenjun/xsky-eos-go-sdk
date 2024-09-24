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

// checks if the ServicesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicesResp{}

// ServicesResp struct for ServicesResp
type ServicesResp struct {
	// services
	Services []Service `json:"services"`
}

type _ServicesResp ServicesResp

// NewServicesResp instantiates a new ServicesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesResp(services []Service) *ServicesResp {
	this := ServicesResp{}
	this.Services = services
	return &this
}

// NewServicesRespWithDefaults instantiates a new ServicesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesRespWithDefaults() *ServicesResp {
	this := ServicesResp{}
	return &this
}

// GetServices returns the Services field value
func (o *ServicesResp) GetServices() []Service {
	if o == nil {
		var ret []Service
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *ServicesResp) GetServicesOk() ([]Service, bool) {
	if o == nil {
		return nil, false
	}
	return o.Services, true
}

// SetServices sets field value
func (o *ServicesResp) SetServices(v []Service) {
	o.Services = v
}

func (o ServicesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["services"] = o.Services
	return toSerialize, nil
}

func (o *ServicesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"services",
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

	varServicesResp := _ServicesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServicesResp)

	if err != nil {
		return err
	}

	*o = ServicesResp(varServicesResp)

	return err
}

type NullableServicesResp struct {
	value *ServicesResp
	isSet bool
}

func (v NullableServicesResp) Get() *ServicesResp {
	return v.value
}

func (v *NullableServicesResp) Set(val *ServicesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesResp(val *ServicesResp) *NullableServicesResp {
	return &NullableServicesResp{value: val, isSet: true}
}

func (v NullableServicesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


