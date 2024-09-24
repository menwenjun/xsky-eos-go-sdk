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

// checks if the OspGatewayResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OspGatewayResp{}

// OspGatewayResp struct for OspGatewayResp
type OspGatewayResp struct {
	OspGateway *OspGatewayRecord `json:"osp_gateway,omitempty"`
}

// NewOspGatewayResp instantiates a new OspGatewayResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOspGatewayResp() *OspGatewayResp {
	this := OspGatewayResp{}
	return &this
}

// NewOspGatewayRespWithDefaults instantiates a new OspGatewayResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOspGatewayRespWithDefaults() *OspGatewayResp {
	this := OspGatewayResp{}
	return &this
}

// GetOspGateway returns the OspGateway field value if set, zero value otherwise.
func (o *OspGatewayResp) GetOspGateway() OspGatewayRecord {
	if o == nil || IsNil(o.OspGateway) {
		var ret OspGatewayRecord
		return ret
	}
	return *o.OspGateway
}

// GetOspGatewayOk returns a tuple with the OspGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayResp) GetOspGatewayOk() (*OspGatewayRecord, bool) {
	if o == nil || IsNil(o.OspGateway) {
		return nil, false
	}
	return o.OspGateway, true
}

// HasOspGateway returns a boolean if a field has been set.
func (o *OspGatewayResp) HasOspGateway() bool {
	if o != nil && !IsNil(o.OspGateway) {
		return true
	}

	return false
}

// SetOspGateway gets a reference to the given OspGatewayRecord and assigns it to the OspGateway field.
func (o *OspGatewayResp) SetOspGateway(v OspGatewayRecord) {
	o.OspGateway = &v
}

func (o OspGatewayResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OspGatewayResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OspGateway) {
		toSerialize["osp_gateway"] = o.OspGateway
	}
	return toSerialize, nil
}

type NullableOspGatewayResp struct {
	value *OspGatewayResp
	isSet bool
}

func (v NullableOspGatewayResp) Get() *OspGatewayResp {
	return v.value
}

func (v *NullableOspGatewayResp) Set(val *OspGatewayResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOspGatewayResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOspGatewayResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspGatewayResp(val *OspGatewayResp) *NullableOspGatewayResp {
	return &NullableOspGatewayResp{value: val, isSet: true}
}

func (v NullableOspGatewayResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspGatewayResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


