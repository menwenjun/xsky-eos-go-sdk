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

// checks if the DfsGatewayZonesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsGatewayZonesResp{}

// DfsGatewayZonesResp struct for DfsGatewayZonesResp
type DfsGatewayZonesResp struct {
	// dfs gateway zones
	DfsGatewayZones []DfsGatewayZoneRecord `json:"dfs_gateway_zones"`
}

type _DfsGatewayZonesResp DfsGatewayZonesResp

// NewDfsGatewayZonesResp instantiates a new DfsGatewayZonesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsGatewayZonesResp(dfsGatewayZones []DfsGatewayZoneRecord) *DfsGatewayZonesResp {
	this := DfsGatewayZonesResp{}
	this.DfsGatewayZones = dfsGatewayZones
	return &this
}

// NewDfsGatewayZonesRespWithDefaults instantiates a new DfsGatewayZonesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsGatewayZonesRespWithDefaults() *DfsGatewayZonesResp {
	this := DfsGatewayZonesResp{}
	return &this
}

// GetDfsGatewayZones returns the DfsGatewayZones field value
func (o *DfsGatewayZonesResp) GetDfsGatewayZones() []DfsGatewayZoneRecord {
	if o == nil {
		var ret []DfsGatewayZoneRecord
		return ret
	}

	return o.DfsGatewayZones
}

// GetDfsGatewayZonesOk returns a tuple with the DfsGatewayZones field value
// and a boolean to check if the value has been set.
func (o *DfsGatewayZonesResp) GetDfsGatewayZonesOk() ([]DfsGatewayZoneRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.DfsGatewayZones, true
}

// SetDfsGatewayZones sets field value
func (o *DfsGatewayZonesResp) SetDfsGatewayZones(v []DfsGatewayZoneRecord) {
	o.DfsGatewayZones = v
}

func (o DfsGatewayZonesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsGatewayZonesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_gateway_zones"] = o.DfsGatewayZones
	return toSerialize, nil
}

func (o *DfsGatewayZonesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_gateway_zones",
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

	varDfsGatewayZonesResp := _DfsGatewayZonesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsGatewayZonesResp)

	if err != nil {
		return err
	}

	*o = DfsGatewayZonesResp(varDfsGatewayZonesResp)

	return err
}

type NullableDfsGatewayZonesResp struct {
	value *DfsGatewayZonesResp
	isSet bool
}

func (v NullableDfsGatewayZonesResp) Get() *DfsGatewayZonesResp {
	return v.value
}

func (v *NullableDfsGatewayZonesResp) Set(val *DfsGatewayZonesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsGatewayZonesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsGatewayZonesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsGatewayZonesResp(val *DfsGatewayZonesResp) *NullableDfsGatewayZonesResp {
	return &NullableDfsGatewayZonesResp{value: val, isSet: true}
}

func (v NullableDfsGatewayZonesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsGatewayZonesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


