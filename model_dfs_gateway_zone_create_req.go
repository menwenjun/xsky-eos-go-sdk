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

// checks if the DfsGatewayZoneCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsGatewayZoneCreateReq{}

// DfsGatewayZoneCreateReq struct for DfsGatewayZoneCreateReq
type DfsGatewayZoneCreateReq struct {
	DfsGatewayZone DfsGatewayZoneCreateReqGatewayZone `json:"dfs_gateway_zone"`
}

type _DfsGatewayZoneCreateReq DfsGatewayZoneCreateReq

// NewDfsGatewayZoneCreateReq instantiates a new DfsGatewayZoneCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsGatewayZoneCreateReq(dfsGatewayZone DfsGatewayZoneCreateReqGatewayZone) *DfsGatewayZoneCreateReq {
	this := DfsGatewayZoneCreateReq{}
	this.DfsGatewayZone = dfsGatewayZone
	return &this
}

// NewDfsGatewayZoneCreateReqWithDefaults instantiates a new DfsGatewayZoneCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsGatewayZoneCreateReqWithDefaults() *DfsGatewayZoneCreateReq {
	this := DfsGatewayZoneCreateReq{}
	return &this
}

// GetDfsGatewayZone returns the DfsGatewayZone field value
func (o *DfsGatewayZoneCreateReq) GetDfsGatewayZone() DfsGatewayZoneCreateReqGatewayZone {
	if o == nil {
		var ret DfsGatewayZoneCreateReqGatewayZone
		return ret
	}

	return o.DfsGatewayZone
}

// GetDfsGatewayZoneOk returns a tuple with the DfsGatewayZone field value
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneCreateReq) GetDfsGatewayZoneOk() (*DfsGatewayZoneCreateReqGatewayZone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsGatewayZone, true
}

// SetDfsGatewayZone sets field value
func (o *DfsGatewayZoneCreateReq) SetDfsGatewayZone(v DfsGatewayZoneCreateReqGatewayZone) {
	o.DfsGatewayZone = v
}

func (o DfsGatewayZoneCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsGatewayZoneCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_gateway_zone"] = o.DfsGatewayZone
	return toSerialize, nil
}

func (o *DfsGatewayZoneCreateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_gateway_zone",
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

	varDfsGatewayZoneCreateReq := _DfsGatewayZoneCreateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsGatewayZoneCreateReq)

	if err != nil {
		return err
	}

	*o = DfsGatewayZoneCreateReq(varDfsGatewayZoneCreateReq)

	return err
}

type NullableDfsGatewayZoneCreateReq struct {
	value *DfsGatewayZoneCreateReq
	isSet bool
}

func (v NullableDfsGatewayZoneCreateReq) Get() *DfsGatewayZoneCreateReq {
	return v.value
}

func (v *NullableDfsGatewayZoneCreateReq) Set(val *DfsGatewayZoneCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsGatewayZoneCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsGatewayZoneCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsGatewayZoneCreateReq(val *DfsGatewayZoneCreateReq) *NullableDfsGatewayZoneCreateReq {
	return &NullableDfsGatewayZoneCreateReq{value: val, isSet: true}
}

func (v NullableDfsGatewayZoneCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsGatewayZoneCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


