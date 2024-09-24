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

// checks if the NFSGatewayCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NFSGatewayCreateReq{}

// NFSGatewayCreateReq struct for NFSGatewayCreateReq
type NFSGatewayCreateReq struct {
	NfsGateway NFSGatewayCreateReqNFSGateway `json:"nfs_gateway"`
}

type _NFSGatewayCreateReq NFSGatewayCreateReq

// NewNFSGatewayCreateReq instantiates a new NFSGatewayCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFSGatewayCreateReq(nfsGateway NFSGatewayCreateReqNFSGateway) *NFSGatewayCreateReq {
	this := NFSGatewayCreateReq{}
	this.NfsGateway = nfsGateway
	return &this
}

// NewNFSGatewayCreateReqWithDefaults instantiates a new NFSGatewayCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFSGatewayCreateReqWithDefaults() *NFSGatewayCreateReq {
	this := NFSGatewayCreateReq{}
	return &this
}

// GetNfsGateway returns the NfsGateway field value
func (o *NFSGatewayCreateReq) GetNfsGateway() NFSGatewayCreateReqNFSGateway {
	if o == nil {
		var ret NFSGatewayCreateReqNFSGateway
		return ret
	}

	return o.NfsGateway
}

// GetNfsGatewayOk returns a tuple with the NfsGateway field value
// and a boolean to check if the value has been set.
func (o *NFSGatewayCreateReq) GetNfsGatewayOk() (*NFSGatewayCreateReqNFSGateway, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfsGateway, true
}

// SetNfsGateway sets field value
func (o *NFSGatewayCreateReq) SetNfsGateway(v NFSGatewayCreateReqNFSGateway) {
	o.NfsGateway = v
}

func (o NFSGatewayCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NFSGatewayCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nfs_gateway"] = o.NfsGateway
	return toSerialize, nil
}

func (o *NFSGatewayCreateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nfs_gateway",
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

	varNFSGatewayCreateReq := _NFSGatewayCreateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNFSGatewayCreateReq)

	if err != nil {
		return err
	}

	*o = NFSGatewayCreateReq(varNFSGatewayCreateReq)

	return err
}

type NullableNFSGatewayCreateReq struct {
	value *NFSGatewayCreateReq
	isSet bool
}

func (v NullableNFSGatewayCreateReq) Get() *NFSGatewayCreateReq {
	return v.value
}

func (v *NullableNFSGatewayCreateReq) Set(val *NFSGatewayCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableNFSGatewayCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableNFSGatewayCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFSGatewayCreateReq(val *NFSGatewayCreateReq) *NullableNFSGatewayCreateReq {
	return &NullableNFSGatewayCreateReq{value: val, isSet: true}
}

func (v NullableNFSGatewayCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFSGatewayCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


