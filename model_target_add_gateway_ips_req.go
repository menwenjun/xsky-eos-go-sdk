/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TargetAddGatewayIPsReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetAddGatewayIPsReq{}

// TargetAddGatewayIPsReq struct for TargetAddGatewayIPsReq
type TargetAddGatewayIPsReq struct {
	Target *TargetAddGatewayIPsReqTarget `json:"target,omitempty"`
}

// NewTargetAddGatewayIPsReq instantiates a new TargetAddGatewayIPsReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetAddGatewayIPsReq() *TargetAddGatewayIPsReq {
	this := TargetAddGatewayIPsReq{}
	return &this
}

// NewTargetAddGatewayIPsReqWithDefaults instantiates a new TargetAddGatewayIPsReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetAddGatewayIPsReqWithDefaults() *TargetAddGatewayIPsReq {
	this := TargetAddGatewayIPsReq{}
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *TargetAddGatewayIPsReq) GetTarget() TargetAddGatewayIPsReqTarget {
	if o == nil || IsNil(o.Target) {
		var ret TargetAddGatewayIPsReqTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetAddGatewayIPsReq) GetTargetOk() (*TargetAddGatewayIPsReqTarget, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *TargetAddGatewayIPsReq) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given TargetAddGatewayIPsReqTarget and assigns it to the Target field.
func (o *TargetAddGatewayIPsReq) SetTarget(v TargetAddGatewayIPsReqTarget) {
	o.Target = &v
}

func (o TargetAddGatewayIPsReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetAddGatewayIPsReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

type NullableTargetAddGatewayIPsReq struct {
	value *TargetAddGatewayIPsReq
	isSet bool
}

func (v NullableTargetAddGatewayIPsReq) Get() *TargetAddGatewayIPsReq {
	return v.value
}

func (v *NullableTargetAddGatewayIPsReq) Set(val *TargetAddGatewayIPsReq) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetAddGatewayIPsReq) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetAddGatewayIPsReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetAddGatewayIPsReq(val *TargetAddGatewayIPsReq) *NullableTargetAddGatewayIPsReq {
	return &NullableTargetAddGatewayIPsReq{value: val, isSet: true}
}

func (v NullableTargetAddGatewayIPsReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetAddGatewayIPsReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


