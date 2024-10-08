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

// checks if the ServiceResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceResp{}

// ServiceResp struct for ServiceResp
type ServiceResp struct {
	Service Service `json:"service"`
}

type _ServiceResp ServiceResp

// NewServiceResp instantiates a new ServiceResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceResp(service Service) *ServiceResp {
	this := ServiceResp{}
	this.Service = service
	return &this
}

// NewServiceRespWithDefaults instantiates a new ServiceResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceRespWithDefaults() *ServiceResp {
	this := ServiceResp{}
	return &this
}

// GetService returns the Service field value
func (o *ServiceResp) GetService() Service {
	if o == nil {
		var ret Service
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ServiceResp) GetServiceOk() (*Service, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *ServiceResp) SetService(v Service) {
	o.Service = v
}

func (o ServiceResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["service"] = o.Service
	return toSerialize, nil
}

func (o *ServiceResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"service",
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

	varServiceResp := _ServiceResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceResp)

	if err != nil {
		return err
	}

	*o = ServiceResp(varServiceResp)

	return err
}

type NullableServiceResp struct {
	value *ServiceResp
	isSet bool
}

func (v NullableServiceResp) Get() *ServiceResp {
	return v.value
}

func (v *NullableServiceResp) Set(val *ServiceResp) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceResp) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceResp(val *ServiceResp) *NullableServiceResp {
	return &NullableServiceResp{value: val, isSet: true}
}

func (v NullableServiceResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


