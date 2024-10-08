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

// checks if the DfsGatewayGroupRemoveGatewaysReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsGatewayGroupRemoveGatewaysReq{}

// DfsGatewayGroupRemoveGatewaysReq struct for DfsGatewayGroupRemoveGatewaysReq
type DfsGatewayGroupRemoveGatewaysReq struct {
	DfsGatewayGroup DfsGatewayGroupRemoveGatewaysReqGatewayGroup `json:"dfs_gateway_group"`
}

type _DfsGatewayGroupRemoveGatewaysReq DfsGatewayGroupRemoveGatewaysReq

// NewDfsGatewayGroupRemoveGatewaysReq instantiates a new DfsGatewayGroupRemoveGatewaysReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsGatewayGroupRemoveGatewaysReq(dfsGatewayGroup DfsGatewayGroupRemoveGatewaysReqGatewayGroup) *DfsGatewayGroupRemoveGatewaysReq {
	this := DfsGatewayGroupRemoveGatewaysReq{}
	this.DfsGatewayGroup = dfsGatewayGroup
	return &this
}

// NewDfsGatewayGroupRemoveGatewaysReqWithDefaults instantiates a new DfsGatewayGroupRemoveGatewaysReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsGatewayGroupRemoveGatewaysReqWithDefaults() *DfsGatewayGroupRemoveGatewaysReq {
	this := DfsGatewayGroupRemoveGatewaysReq{}
	return &this
}

// GetDfsGatewayGroup returns the DfsGatewayGroup field value
func (o *DfsGatewayGroupRemoveGatewaysReq) GetDfsGatewayGroup() DfsGatewayGroupRemoveGatewaysReqGatewayGroup {
	if o == nil {
		var ret DfsGatewayGroupRemoveGatewaysReqGatewayGroup
		return ret
	}

	return o.DfsGatewayGroup
}

// GetDfsGatewayGroupOk returns a tuple with the DfsGatewayGroup field value
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroupRemoveGatewaysReq) GetDfsGatewayGroupOk() (*DfsGatewayGroupRemoveGatewaysReqGatewayGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsGatewayGroup, true
}

// SetDfsGatewayGroup sets field value
func (o *DfsGatewayGroupRemoveGatewaysReq) SetDfsGatewayGroup(v DfsGatewayGroupRemoveGatewaysReqGatewayGroup) {
	o.DfsGatewayGroup = v
}

func (o DfsGatewayGroupRemoveGatewaysReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsGatewayGroupRemoveGatewaysReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_gateway_group"] = o.DfsGatewayGroup
	return toSerialize, nil
}

func (o *DfsGatewayGroupRemoveGatewaysReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_gateway_group",
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

	varDfsGatewayGroupRemoveGatewaysReq := _DfsGatewayGroupRemoveGatewaysReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsGatewayGroupRemoveGatewaysReq)

	if err != nil {
		return err
	}

	*o = DfsGatewayGroupRemoveGatewaysReq(varDfsGatewayGroupRemoveGatewaysReq)

	return err
}

type NullableDfsGatewayGroupRemoveGatewaysReq struct {
	value *DfsGatewayGroupRemoveGatewaysReq
	isSet bool
}

func (v NullableDfsGatewayGroupRemoveGatewaysReq) Get() *DfsGatewayGroupRemoveGatewaysReq {
	return v.value
}

func (v *NullableDfsGatewayGroupRemoveGatewaysReq) Set(val *DfsGatewayGroupRemoveGatewaysReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsGatewayGroupRemoveGatewaysReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsGatewayGroupRemoveGatewaysReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsGatewayGroupRemoveGatewaysReq(val *DfsGatewayGroupRemoveGatewaysReq) *NullableDfsGatewayGroupRemoveGatewaysReq {
	return &NullableDfsGatewayGroupRemoveGatewaysReq{value: val, isSet: true}
}

func (v NullableDfsGatewayGroupRemoveGatewaysReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsGatewayGroupRemoveGatewaysReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


