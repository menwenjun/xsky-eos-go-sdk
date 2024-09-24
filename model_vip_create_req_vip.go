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

// checks if the VIPCreateReqVIP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VIPCreateReqVIP{}

// VIPCreateReqVIP struct for VIPCreateReqVIP
type VIPCreateReqVIP struct {
	Ip string `json:"ip"`
	Mask *int64 `json:"mask,omitempty"`
	NetworkAddressId int64 `json:"network_address_id"`
	VipGroupId int64 `json:"vip_group_id"`
}

type _VIPCreateReqVIP VIPCreateReqVIP

// NewVIPCreateReqVIP instantiates a new VIPCreateReqVIP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVIPCreateReqVIP(ip string, networkAddressId int64, vipGroupId int64) *VIPCreateReqVIP {
	this := VIPCreateReqVIP{}
	this.Ip = ip
	this.NetworkAddressId = networkAddressId
	this.VipGroupId = vipGroupId
	return &this
}

// NewVIPCreateReqVIPWithDefaults instantiates a new VIPCreateReqVIP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVIPCreateReqVIPWithDefaults() *VIPCreateReqVIP {
	this := VIPCreateReqVIP{}
	return &this
}

// GetIp returns the Ip field value
func (o *VIPCreateReqVIP) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *VIPCreateReqVIP) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *VIPCreateReqVIP) SetIp(v string) {
	o.Ip = v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *VIPCreateReqVIP) GetMask() int64 {
	if o == nil || IsNil(o.Mask) {
		var ret int64
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VIPCreateReqVIP) GetMaskOk() (*int64, bool) {
	if o == nil || IsNil(o.Mask) {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *VIPCreateReqVIP) HasMask() bool {
	if o != nil && !IsNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given int64 and assigns it to the Mask field.
func (o *VIPCreateReqVIP) SetMask(v int64) {
	o.Mask = &v
}

// GetNetworkAddressId returns the NetworkAddressId field value
func (o *VIPCreateReqVIP) GetNetworkAddressId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NetworkAddressId
}

// GetNetworkAddressIdOk returns a tuple with the NetworkAddressId field value
// and a boolean to check if the value has been set.
func (o *VIPCreateReqVIP) GetNetworkAddressIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkAddressId, true
}

// SetNetworkAddressId sets field value
func (o *VIPCreateReqVIP) SetNetworkAddressId(v int64) {
	o.NetworkAddressId = v
}

// GetVipGroupId returns the VipGroupId field value
func (o *VIPCreateReqVIP) GetVipGroupId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VipGroupId
}

// GetVipGroupIdOk returns a tuple with the VipGroupId field value
// and a boolean to check if the value has been set.
func (o *VIPCreateReqVIP) GetVipGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VipGroupId, true
}

// SetVipGroupId sets field value
func (o *VIPCreateReqVIP) SetVipGroupId(v int64) {
	o.VipGroupId = v
}

func (o VIPCreateReqVIP) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VIPCreateReqVIP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip"] = o.Ip
	if !IsNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	toSerialize["network_address_id"] = o.NetworkAddressId
	toSerialize["vip_group_id"] = o.VipGroupId
	return toSerialize, nil
}

func (o *VIPCreateReqVIP) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ip",
		"network_address_id",
		"vip_group_id",
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

	varVIPCreateReqVIP := _VIPCreateReqVIP{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVIPCreateReqVIP)

	if err != nil {
		return err
	}

	*o = VIPCreateReqVIP(varVIPCreateReqVIP)

	return err
}

type NullableVIPCreateReqVIP struct {
	value *VIPCreateReqVIP
	isSet bool
}

func (v NullableVIPCreateReqVIP) Get() *VIPCreateReqVIP {
	return v.value
}

func (v *NullableVIPCreateReqVIP) Set(val *VIPCreateReqVIP) {
	v.value = val
	v.isSet = true
}

func (v NullableVIPCreateReqVIP) IsSet() bool {
	return v.isSet
}

func (v *NullableVIPCreateReqVIP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVIPCreateReqVIP(val *VIPCreateReqVIP) *NullableVIPCreateReqVIP {
	return &NullableVIPCreateReqVIP{value: val, isSet: true}
}

func (v NullableVIPCreateReqVIP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVIPCreateReqVIP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


