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

// checks if the CephHostReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CephHostReq{}

// CephHostReq struct for CephHostReq
type CephHostReq struct {
	// gateway ips for s3
	GatewayIps []string `json:"gateway_ips,omitempty"`
	// boot node host id
	HostId int64 `json:"host_id"`
	// cluster private ip for internal data access
	PrivateIp *string `json:"private_ip,omitempty"`
	// public ip for outside data access
	PublicIp *string `json:"public_ip,omitempty"`
	// host roles
	Roles []string `json:"roles,omitempty"`
	// storage server, storage client or storage witness
	Type *string `json:"type,omitempty"`
}

type _CephHostReq CephHostReq

// NewCephHostReq instantiates a new CephHostReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCephHostReq(hostId int64) *CephHostReq {
	this := CephHostReq{}
	this.HostId = hostId
	return &this
}

// NewCephHostReqWithDefaults instantiates a new CephHostReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCephHostReqWithDefaults() *CephHostReq {
	this := CephHostReq{}
	return &this
}

// GetGatewayIps returns the GatewayIps field value if set, zero value otherwise.
func (o *CephHostReq) GetGatewayIps() []string {
	if o == nil || IsNil(o.GatewayIps) {
		var ret []string
		return ret
	}
	return o.GatewayIps
}

// GetGatewayIpsOk returns a tuple with the GatewayIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CephHostReq) GetGatewayIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.GatewayIps) {
		return nil, false
	}
	return o.GatewayIps, true
}

// HasGatewayIps returns a boolean if a field has been set.
func (o *CephHostReq) HasGatewayIps() bool {
	if o != nil && !IsNil(o.GatewayIps) {
		return true
	}

	return false
}

// SetGatewayIps gets a reference to the given []string and assigns it to the GatewayIps field.
func (o *CephHostReq) SetGatewayIps(v []string) {
	o.GatewayIps = v
}

// GetHostId returns the HostId field value
func (o *CephHostReq) GetHostId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value
// and a boolean to check if the value has been set.
func (o *CephHostReq) GetHostIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostId, true
}

// SetHostId sets field value
func (o *CephHostReq) SetHostId(v int64) {
	o.HostId = v
}

// GetPrivateIp returns the PrivateIp field value if set, zero value otherwise.
func (o *CephHostReq) GetPrivateIp() string {
	if o == nil || IsNil(o.PrivateIp) {
		var ret string
		return ret
	}
	return *o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CephHostReq) GetPrivateIpOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateIp) {
		return nil, false
	}
	return o.PrivateIp, true
}

// HasPrivateIp returns a boolean if a field has been set.
func (o *CephHostReq) HasPrivateIp() bool {
	if o != nil && !IsNil(o.PrivateIp) {
		return true
	}

	return false
}

// SetPrivateIp gets a reference to the given string and assigns it to the PrivateIp field.
func (o *CephHostReq) SetPrivateIp(v string) {
	o.PrivateIp = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *CephHostReq) GetPublicIp() string {
	if o == nil || IsNil(o.PublicIp) {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CephHostReq) GetPublicIpOk() (*string, bool) {
	if o == nil || IsNil(o.PublicIp) {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *CephHostReq) HasPublicIp() bool {
	if o != nil && !IsNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *CephHostReq) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *CephHostReq) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CephHostReq) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CephHostReq) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *CephHostReq) SetRoles(v []string) {
	o.Roles = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CephHostReq) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CephHostReq) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CephHostReq) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CephHostReq) SetType(v string) {
	o.Type = &v
}

func (o CephHostReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CephHostReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayIps) {
		toSerialize["gateway_ips"] = o.GatewayIps
	}
	toSerialize["host_id"] = o.HostId
	if !IsNil(o.PrivateIp) {
		toSerialize["private_ip"] = o.PrivateIp
	}
	if !IsNil(o.PublicIp) {
		toSerialize["public_ip"] = o.PublicIp
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *CephHostReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host_id",
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

	varCephHostReq := _CephHostReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCephHostReq)

	if err != nil {
		return err
	}

	*o = CephHostReq(varCephHostReq)

	return err
}

type NullableCephHostReq struct {
	value *CephHostReq
	isSet bool
}

func (v NullableCephHostReq) Get() *CephHostReq {
	return v.value
}

func (v *NullableCephHostReq) Set(val *CephHostReq) {
	v.value = val
	v.isSet = true
}

func (v NullableCephHostReq) IsSet() bool {
	return v.isSet
}

func (v *NullableCephHostReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCephHostReq(val *CephHostReq) *NullableCephHostReq {
	return &NullableCephHostReq{value: val, isSet: true}
}

func (v NullableCephHostReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCephHostReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


