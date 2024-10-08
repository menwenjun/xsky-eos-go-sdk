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

// checks if the DfsGatewayZoneResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsGatewayZoneResp{}

// DfsGatewayZoneResp struct for DfsGatewayZoneResp
type DfsGatewayZoneResp struct {
	DfsGatewayZone DfsGatewayZoneRecord `json:"dfs_gateway_zone"`
}

type _DfsGatewayZoneResp DfsGatewayZoneResp

// NewDfsGatewayZoneResp instantiates a new DfsGatewayZoneResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsGatewayZoneResp(dfsGatewayZone DfsGatewayZoneRecord) *DfsGatewayZoneResp {
	this := DfsGatewayZoneResp{}
	this.DfsGatewayZone = dfsGatewayZone
	return &this
}

// NewDfsGatewayZoneRespWithDefaults instantiates a new DfsGatewayZoneResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsGatewayZoneRespWithDefaults() *DfsGatewayZoneResp {
	this := DfsGatewayZoneResp{}
	return &this
}

// GetDfsGatewayZone returns the DfsGatewayZone field value
func (o *DfsGatewayZoneResp) GetDfsGatewayZone() DfsGatewayZoneRecord {
	if o == nil {
		var ret DfsGatewayZoneRecord
		return ret
	}

	return o.DfsGatewayZone
}

// GetDfsGatewayZoneOk returns a tuple with the DfsGatewayZone field value
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneResp) GetDfsGatewayZoneOk() (*DfsGatewayZoneRecord, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsGatewayZone, true
}

// SetDfsGatewayZone sets field value
func (o *DfsGatewayZoneResp) SetDfsGatewayZone(v DfsGatewayZoneRecord) {
	o.DfsGatewayZone = v
}

func (o DfsGatewayZoneResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsGatewayZoneResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_gateway_zone"] = o.DfsGatewayZone
	return toSerialize, nil
}

func (o *DfsGatewayZoneResp) UnmarshalJSON(data []byte) (err error) {
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

	varDfsGatewayZoneResp := _DfsGatewayZoneResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsGatewayZoneResp)

	if err != nil {
		return err
	}

	*o = DfsGatewayZoneResp(varDfsGatewayZoneResp)

	return err
}

type NullableDfsGatewayZoneResp struct {
	value *DfsGatewayZoneResp
	isSet bool
}

func (v NullableDfsGatewayZoneResp) Get() *DfsGatewayZoneResp {
	return v.value
}

func (v *NullableDfsGatewayZoneResp) Set(val *DfsGatewayZoneResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsGatewayZoneResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsGatewayZoneResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsGatewayZoneResp(val *DfsGatewayZoneResp) *NullableDfsGatewayZoneResp {
	return &NullableDfsGatewayZoneResp{value: val, isSet: true}
}

func (v NullableDfsGatewayZoneResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsGatewayZoneResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


