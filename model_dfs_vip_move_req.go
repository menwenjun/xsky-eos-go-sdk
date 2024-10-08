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

// checks if the DfsVIPMoveReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsVIPMoveReq{}

// DfsVIPMoveReq struct for DfsVIPMoveReq
type DfsVIPMoveReq struct {
	DfsVip DfsVIPMoveReqVIP `json:"dfs_vip"`
}

type _DfsVIPMoveReq DfsVIPMoveReq

// NewDfsVIPMoveReq instantiates a new DfsVIPMoveReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsVIPMoveReq(dfsVip DfsVIPMoveReqVIP) *DfsVIPMoveReq {
	this := DfsVIPMoveReq{}
	this.DfsVip = dfsVip
	return &this
}

// NewDfsVIPMoveReqWithDefaults instantiates a new DfsVIPMoveReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsVIPMoveReqWithDefaults() *DfsVIPMoveReq {
	this := DfsVIPMoveReq{}
	return &this
}

// GetDfsVip returns the DfsVip field value
func (o *DfsVIPMoveReq) GetDfsVip() DfsVIPMoveReqVIP {
	if o == nil {
		var ret DfsVIPMoveReqVIP
		return ret
	}

	return o.DfsVip
}

// GetDfsVipOk returns a tuple with the DfsVip field value
// and a boolean to check if the value has been set.
func (o *DfsVIPMoveReq) GetDfsVipOk() (*DfsVIPMoveReqVIP, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsVip, true
}

// SetDfsVip sets field value
func (o *DfsVIPMoveReq) SetDfsVip(v DfsVIPMoveReqVIP) {
	o.DfsVip = v
}

func (o DfsVIPMoveReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsVIPMoveReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_vip"] = o.DfsVip
	return toSerialize, nil
}

func (o *DfsVIPMoveReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_vip",
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

	varDfsVIPMoveReq := _DfsVIPMoveReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsVIPMoveReq)

	if err != nil {
		return err
	}

	*o = DfsVIPMoveReq(varDfsVIPMoveReq)

	return err
}

type NullableDfsVIPMoveReq struct {
	value *DfsVIPMoveReq
	isSet bool
}

func (v NullableDfsVIPMoveReq) Get() *DfsVIPMoveReq {
	return v.value
}

func (v *NullableDfsVIPMoveReq) Set(val *DfsVIPMoveReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsVIPMoveReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsVIPMoveReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsVIPMoveReq(val *DfsVIPMoveReq) *NullableDfsVIPMoveReq {
	return &NullableDfsVIPMoveReq{value: val, isSet: true}
}

func (v NullableDfsVIPMoveReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsVIPMoveReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


