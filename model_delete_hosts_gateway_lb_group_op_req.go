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

// checks if the DeleteHostsGatewayLbGroupOpReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteHostsGatewayLbGroupOpReq{}

// DeleteHostsGatewayLbGroupOpReq struct for DeleteHostsGatewayLbGroupOpReq
type DeleteHostsGatewayLbGroupOpReq struct {
	OspZone *DeleteHostsGatewayLbGroupOpReqOspZone `json:"osp_zone,omitempty"`
}

// NewDeleteHostsGatewayLbGroupOpReq instantiates a new DeleteHostsGatewayLbGroupOpReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteHostsGatewayLbGroupOpReq() *DeleteHostsGatewayLbGroupOpReq {
	this := DeleteHostsGatewayLbGroupOpReq{}
	return &this
}

// NewDeleteHostsGatewayLbGroupOpReqWithDefaults instantiates a new DeleteHostsGatewayLbGroupOpReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteHostsGatewayLbGroupOpReqWithDefaults() *DeleteHostsGatewayLbGroupOpReq {
	this := DeleteHostsGatewayLbGroupOpReq{}
	return &this
}

// GetOspZone returns the OspZone field value if set, zero value otherwise.
func (o *DeleteHostsGatewayLbGroupOpReq) GetOspZone() DeleteHostsGatewayLbGroupOpReqOspZone {
	if o == nil || IsNil(o.OspZone) {
		var ret DeleteHostsGatewayLbGroupOpReqOspZone
		return ret
	}
	return *o.OspZone
}

// GetOspZoneOk returns a tuple with the OspZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteHostsGatewayLbGroupOpReq) GetOspZoneOk() (*DeleteHostsGatewayLbGroupOpReqOspZone, bool) {
	if o == nil || IsNil(o.OspZone) {
		return nil, false
	}
	return o.OspZone, true
}

// HasOspZone returns a boolean if a field has been set.
func (o *DeleteHostsGatewayLbGroupOpReq) HasOspZone() bool {
	if o != nil && !IsNil(o.OspZone) {
		return true
	}

	return false
}

// SetOspZone gets a reference to the given DeleteHostsGatewayLbGroupOpReqOspZone and assigns it to the OspZone field.
func (o *DeleteHostsGatewayLbGroupOpReq) SetOspZone(v DeleteHostsGatewayLbGroupOpReqOspZone) {
	o.OspZone = &v
}

func (o DeleteHostsGatewayLbGroupOpReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteHostsGatewayLbGroupOpReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OspZone) {
		toSerialize["osp_zone"] = o.OspZone
	}
	return toSerialize, nil
}

type NullableDeleteHostsGatewayLbGroupOpReq struct {
	value *DeleteHostsGatewayLbGroupOpReq
	isSet bool
}

func (v NullableDeleteHostsGatewayLbGroupOpReq) Get() *DeleteHostsGatewayLbGroupOpReq {
	return v.value
}

func (v *NullableDeleteHostsGatewayLbGroupOpReq) Set(val *DeleteHostsGatewayLbGroupOpReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteHostsGatewayLbGroupOpReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteHostsGatewayLbGroupOpReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteHostsGatewayLbGroupOpReq(val *DeleteHostsGatewayLbGroupOpReq) *NullableDeleteHostsGatewayLbGroupOpReq {
	return &NullableDeleteHostsGatewayLbGroupOpReq{value: val, isSet: true}
}

func (v NullableDeleteHostsGatewayLbGroupOpReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteHostsGatewayLbGroupOpReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


