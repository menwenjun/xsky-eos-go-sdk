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

// checks if the S3LoadBalancerGroupCreateReqGroupLoadBalancersElt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3LoadBalancerGroupCreateReqGroupLoadBalancersElt{}

// S3LoadBalancerGroupCreateReqGroupLoadBalancersElt struct for S3LoadBalancerGroupCreateReqGroupLoadBalancersElt
type S3LoadBalancerGroupCreateReqGroupLoadBalancersElt struct {
	// host id of load balancer
	HostId int64 `json:"host_id"`
	// vip will be bounded to interface, exclusive to ip
	InterfaceName *string `json:"interface_name,omitempty"`
	// vip will be bounded to interface of the gateway ip, exclusive to interface name
	Ip *string `json:"ip,omitempty"`
	// vip of load balancer
	Vip string `json:"vip"`
}

type _S3LoadBalancerGroupCreateReqGroupLoadBalancersElt S3LoadBalancerGroupCreateReqGroupLoadBalancersElt

// NewS3LoadBalancerGroupCreateReqGroupLoadBalancersElt instantiates a new S3LoadBalancerGroupCreateReqGroupLoadBalancersElt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3LoadBalancerGroupCreateReqGroupLoadBalancersElt(hostId int64, vip string) *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt {
	this := S3LoadBalancerGroupCreateReqGroupLoadBalancersElt{}
	this.HostId = hostId
	this.Vip = vip
	return &this
}

// NewS3LoadBalancerGroupCreateReqGroupLoadBalancersEltWithDefaults instantiates a new S3LoadBalancerGroupCreateReqGroupLoadBalancersElt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3LoadBalancerGroupCreateReqGroupLoadBalancersEltWithDefaults() *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt {
	this := S3LoadBalancerGroupCreateReqGroupLoadBalancersElt{}
	return &this
}

// GetHostId returns the HostId field value
func (o *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) GetHostId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value
// and a boolean to check if the value has been set.
func (o *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) GetHostIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostId, true
}

// SetHostId sets field value
func (o *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) SetHostId(v int64) {
	o.HostId = v
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) GetInterfaceName() string {
	if o == nil || IsNil(o.InterfaceName) {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) GetInterfaceNameOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceName) {
		return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) HasInterfaceName() bool {
	if o != nil && !IsNil(o.InterfaceName) {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) SetIp(v string) {
	o.Ip = &v
}

// GetVip returns the Vip field value
func (o *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) GetVip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vip
}

// GetVipOk returns a tuple with the Vip field value
// and a boolean to check if the value has been set.
func (o *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) GetVipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vip, true
}

// SetVip sets field value
func (o *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) SetVip(v string) {
	o.Vip = v
}

func (o S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host_id"] = o.HostId
	if !IsNil(o.InterfaceName) {
		toSerialize["interface_name"] = o.InterfaceName
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	toSerialize["vip"] = o.Vip
	return toSerialize, nil
}

func (o *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host_id",
		"vip",
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

	varS3LoadBalancerGroupCreateReqGroupLoadBalancersElt := _S3LoadBalancerGroupCreateReqGroupLoadBalancersElt{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varS3LoadBalancerGroupCreateReqGroupLoadBalancersElt)

	if err != nil {
		return err
	}

	*o = S3LoadBalancerGroupCreateReqGroupLoadBalancersElt(varS3LoadBalancerGroupCreateReqGroupLoadBalancersElt)

	return err
}

type NullableS3LoadBalancerGroupCreateReqGroupLoadBalancersElt struct {
	value *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt
	isSet bool
}

func (v NullableS3LoadBalancerGroupCreateReqGroupLoadBalancersElt) Get() *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt {
	return v.value
}

func (v *NullableS3LoadBalancerGroupCreateReqGroupLoadBalancersElt) Set(val *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) {
	v.value = val
	v.isSet = true
}

func (v NullableS3LoadBalancerGroupCreateReqGroupLoadBalancersElt) IsSet() bool {
	return v.isSet
}

func (v *NullableS3LoadBalancerGroupCreateReqGroupLoadBalancersElt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3LoadBalancerGroupCreateReqGroupLoadBalancersElt(val *S3LoadBalancerGroupCreateReqGroupLoadBalancersElt) *NullableS3LoadBalancerGroupCreateReqGroupLoadBalancersElt {
	return &NullableS3LoadBalancerGroupCreateReqGroupLoadBalancersElt{value: val, isSet: true}
}

func (v NullableS3LoadBalancerGroupCreateReqGroupLoadBalancersElt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3LoadBalancerGroupCreateReqGroupLoadBalancersElt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


