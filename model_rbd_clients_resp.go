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

// checks if the RBDClientsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RBDClientsResp{}

// RBDClientsResp struct for RBDClientsResp
type RBDClientsResp struct {
	RbdClients []RBDClient `json:"rbd_clients"`
}

type _RBDClientsResp RBDClientsResp

// NewRBDClientsResp instantiates a new RBDClientsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRBDClientsResp(rbdClients []RBDClient) *RBDClientsResp {
	this := RBDClientsResp{}
	this.RbdClients = rbdClients
	return &this
}

// NewRBDClientsRespWithDefaults instantiates a new RBDClientsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRBDClientsRespWithDefaults() *RBDClientsResp {
	this := RBDClientsResp{}
	return &this
}

// GetRbdClients returns the RbdClients field value
func (o *RBDClientsResp) GetRbdClients() []RBDClient {
	if o == nil {
		var ret []RBDClient
		return ret
	}

	return o.RbdClients
}

// GetRbdClientsOk returns a tuple with the RbdClients field value
// and a boolean to check if the value has been set.
func (o *RBDClientsResp) GetRbdClientsOk() ([]RBDClient, bool) {
	if o == nil {
		return nil, false
	}
	return o.RbdClients, true
}

// SetRbdClients sets field value
func (o *RBDClientsResp) SetRbdClients(v []RBDClient) {
	o.RbdClients = v
}

func (o RBDClientsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RBDClientsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rbd_clients"] = o.RbdClients
	return toSerialize, nil
}

func (o *RBDClientsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rbd_clients",
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

	varRBDClientsResp := _RBDClientsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRBDClientsResp)

	if err != nil {
		return err
	}

	*o = RBDClientsResp(varRBDClientsResp)

	return err
}

type NullableRBDClientsResp struct {
	value *RBDClientsResp
	isSet bool
}

func (v NullableRBDClientsResp) Get() *RBDClientsResp {
	return v.value
}

func (v *NullableRBDClientsResp) Set(val *RBDClientsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableRBDClientsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableRBDClientsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRBDClientsResp(val *RBDClientsResp) *NullableRBDClientsResp {
	return &NullableRBDClientsResp{value: val, isSet: true}
}

func (v NullableRBDClientsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRBDClientsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


