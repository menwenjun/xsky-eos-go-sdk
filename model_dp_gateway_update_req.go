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

// checks if the DpGatewayUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpGatewayUpdateReq{}

// DpGatewayUpdateReq struct for DpGatewayUpdateReq
type DpGatewayUpdateReq struct {
	DpGateway *DpGatewayUpdateReqGateway `json:"dp_gateway,omitempty"`
}

// NewDpGatewayUpdateReq instantiates a new DpGatewayUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpGatewayUpdateReq() *DpGatewayUpdateReq {
	this := DpGatewayUpdateReq{}
	return &this
}

// NewDpGatewayUpdateReqWithDefaults instantiates a new DpGatewayUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpGatewayUpdateReqWithDefaults() *DpGatewayUpdateReq {
	this := DpGatewayUpdateReq{}
	return &this
}

// GetDpGateway returns the DpGateway field value if set, zero value otherwise.
func (o *DpGatewayUpdateReq) GetDpGateway() DpGatewayUpdateReqGateway {
	if o == nil || IsNil(o.DpGateway) {
		var ret DpGatewayUpdateReqGateway
		return ret
	}
	return *o.DpGateway
}

// GetDpGatewayOk returns a tuple with the DpGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpGatewayUpdateReq) GetDpGatewayOk() (*DpGatewayUpdateReqGateway, bool) {
	if o == nil || IsNil(o.DpGateway) {
		return nil, false
	}
	return o.DpGateway, true
}

// HasDpGateway returns a boolean if a field has been set.
func (o *DpGatewayUpdateReq) HasDpGateway() bool {
	if o != nil && !IsNil(o.DpGateway) {
		return true
	}

	return false
}

// SetDpGateway gets a reference to the given DpGatewayUpdateReqGateway and assigns it to the DpGateway field.
func (o *DpGatewayUpdateReq) SetDpGateway(v DpGatewayUpdateReqGateway) {
	o.DpGateway = &v
}

func (o DpGatewayUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpGatewayUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DpGateway) {
		toSerialize["dp_gateway"] = o.DpGateway
	}
	return toSerialize, nil
}

type NullableDpGatewayUpdateReq struct {
	value *DpGatewayUpdateReq
	isSet bool
}

func (v NullableDpGatewayUpdateReq) Get() *DpGatewayUpdateReq {
	return v.value
}

func (v *NullableDpGatewayUpdateReq) Set(val *DpGatewayUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDpGatewayUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDpGatewayUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpGatewayUpdateReq(val *DpGatewayUpdateReq) *NullableDpGatewayUpdateReq {
	return &NullableDpGatewayUpdateReq{value: val, isSet: true}
}

func (v NullableDpGatewayUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpGatewayUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


