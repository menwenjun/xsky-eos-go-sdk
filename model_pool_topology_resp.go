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

// checks if the PoolTopologyResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolTopologyResp{}

// PoolTopologyResp struct for PoolTopologyResp
type PoolTopologyResp struct {
	PoolTopology PoolPlacementNode `json:"pool_topology"`
}

type _PoolTopologyResp PoolTopologyResp

// NewPoolTopologyResp instantiates a new PoolTopologyResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolTopologyResp(poolTopology PoolPlacementNode) *PoolTopologyResp {
	this := PoolTopologyResp{}
	this.PoolTopology = poolTopology
	return &this
}

// NewPoolTopologyRespWithDefaults instantiates a new PoolTopologyResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolTopologyRespWithDefaults() *PoolTopologyResp {
	this := PoolTopologyResp{}
	return &this
}

// GetPoolTopology returns the PoolTopology field value
func (o *PoolTopologyResp) GetPoolTopology() PoolPlacementNode {
	if o == nil {
		var ret PoolPlacementNode
		return ret
	}

	return o.PoolTopology
}

// GetPoolTopologyOk returns a tuple with the PoolTopology field value
// and a boolean to check if the value has been set.
func (o *PoolTopologyResp) GetPoolTopologyOk() (*PoolPlacementNode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolTopology, true
}

// SetPoolTopology sets field value
func (o *PoolTopologyResp) SetPoolTopology(v PoolPlacementNode) {
	o.PoolTopology = v
}

func (o PoolTopologyResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolTopologyResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pool_topology"] = o.PoolTopology
	return toSerialize, nil
}

func (o *PoolTopologyResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pool_topology",
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

	varPoolTopologyResp := _PoolTopologyResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoolTopologyResp)

	if err != nil {
		return err
	}

	*o = PoolTopologyResp(varPoolTopologyResp)

	return err
}

type NullablePoolTopologyResp struct {
	value *PoolTopologyResp
	isSet bool
}

func (v NullablePoolTopologyResp) Get() *PoolTopologyResp {
	return v.value
}

func (v *NullablePoolTopologyResp) Set(val *PoolTopologyResp) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolTopologyResp) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolTopologyResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolTopologyResp(val *PoolTopologyResp) *NullablePoolTopologyResp {
	return &NullablePoolTopologyResp{value: val, isSet: true}
}

func (v NullablePoolTopologyResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolTopologyResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


