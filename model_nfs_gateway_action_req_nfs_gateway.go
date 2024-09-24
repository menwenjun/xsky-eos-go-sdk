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

// checks if the NFSGatewayActionReqNFSGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NFSGatewayActionReqNFSGateway{}

// NFSGatewayActionReqNFSGateway struct for NFSGatewayActionReqNFSGateway
type NFSGatewayActionReqNFSGateway struct {
	Action string `json:"action"`
}

type _NFSGatewayActionReqNFSGateway NFSGatewayActionReqNFSGateway

// NewNFSGatewayActionReqNFSGateway instantiates a new NFSGatewayActionReqNFSGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFSGatewayActionReqNFSGateway(action string) *NFSGatewayActionReqNFSGateway {
	this := NFSGatewayActionReqNFSGateway{}
	this.Action = action
	return &this
}

// NewNFSGatewayActionReqNFSGatewayWithDefaults instantiates a new NFSGatewayActionReqNFSGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFSGatewayActionReqNFSGatewayWithDefaults() *NFSGatewayActionReqNFSGateway {
	this := NFSGatewayActionReqNFSGateway{}
	return &this
}

// GetAction returns the Action field value
func (o *NFSGatewayActionReqNFSGateway) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *NFSGatewayActionReqNFSGateway) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *NFSGatewayActionReqNFSGateway) SetAction(v string) {
	o.Action = v
}

func (o NFSGatewayActionReqNFSGateway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NFSGatewayActionReqNFSGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

func (o *NFSGatewayActionReqNFSGateway) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
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

	varNFSGatewayActionReqNFSGateway := _NFSGatewayActionReqNFSGateway{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNFSGatewayActionReqNFSGateway)

	if err != nil {
		return err
	}

	*o = NFSGatewayActionReqNFSGateway(varNFSGatewayActionReqNFSGateway)

	return err
}

type NullableNFSGatewayActionReqNFSGateway struct {
	value *NFSGatewayActionReqNFSGateway
	isSet bool
}

func (v NullableNFSGatewayActionReqNFSGateway) Get() *NFSGatewayActionReqNFSGateway {
	return v.value
}

func (v *NullableNFSGatewayActionReqNFSGateway) Set(val *NFSGatewayActionReqNFSGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableNFSGatewayActionReqNFSGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableNFSGatewayActionReqNFSGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFSGatewayActionReqNFSGateway(val *NFSGatewayActionReqNFSGateway) *NullableNFSGatewayActionReqNFSGateway {
	return &NullableNFSGatewayActionReqNFSGateway{value: val, isSet: true}
}

func (v NullableNFSGatewayActionReqNFSGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFSGatewayActionReqNFSGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


