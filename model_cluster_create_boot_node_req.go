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

// checks if the ClusterCreateBootNodeReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterCreateBootNodeReq{}

// ClusterCreateBootNodeReq struct for ClusterCreateBootNodeReq
type ClusterCreateBootNodeReq struct {
	// gateway ips for s3
	GatewayIps []string `json:"gateway_ips,omitempty"`
	// gateway network
	GatewayNetwork *string `json:"gateway_network,omitempty"`
	// boot node host id
	HostId int64 `json:"host_id"`
	// cluster private ip for internal data access
	PrivateIp *string `json:"private_ip,omitempty"`
	// private network
	PrivateNetwork string `json:"private_network"`
	// public ip for outside data access
	PublicIp *string `json:"public_ip,omitempty"`
	// public network
	PublicNetwork string `json:"public_network"`
	// host roles: admin,monitor,block_storage_gateway,file_storage_gateway,s3_gateway,nfs_gateway
	Roles []string `json:"roles,omitempty"`
}

type _ClusterCreateBootNodeReq ClusterCreateBootNodeReq

// NewClusterCreateBootNodeReq instantiates a new ClusterCreateBootNodeReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterCreateBootNodeReq(hostId int64, privateNetwork string, publicNetwork string) *ClusterCreateBootNodeReq {
	this := ClusterCreateBootNodeReq{}
	this.HostId = hostId
	this.PrivateNetwork = privateNetwork
	this.PublicNetwork = publicNetwork
	return &this
}

// NewClusterCreateBootNodeReqWithDefaults instantiates a new ClusterCreateBootNodeReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterCreateBootNodeReqWithDefaults() *ClusterCreateBootNodeReq {
	this := ClusterCreateBootNodeReq{}
	return &this
}

// GetGatewayIps returns the GatewayIps field value if set, zero value otherwise.
func (o *ClusterCreateBootNodeReq) GetGatewayIps() []string {
	if o == nil || IsNil(o.GatewayIps) {
		var ret []string
		return ret
	}
	return o.GatewayIps
}

// GetGatewayIpsOk returns a tuple with the GatewayIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreateBootNodeReq) GetGatewayIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.GatewayIps) {
		return nil, false
	}
	return o.GatewayIps, true
}

// HasGatewayIps returns a boolean if a field has been set.
func (o *ClusterCreateBootNodeReq) HasGatewayIps() bool {
	if o != nil && !IsNil(o.GatewayIps) {
		return true
	}

	return false
}

// SetGatewayIps gets a reference to the given []string and assigns it to the GatewayIps field.
func (o *ClusterCreateBootNodeReq) SetGatewayIps(v []string) {
	o.GatewayIps = v
}

// GetGatewayNetwork returns the GatewayNetwork field value if set, zero value otherwise.
func (o *ClusterCreateBootNodeReq) GetGatewayNetwork() string {
	if o == nil || IsNil(o.GatewayNetwork) {
		var ret string
		return ret
	}
	return *o.GatewayNetwork
}

// GetGatewayNetworkOk returns a tuple with the GatewayNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreateBootNodeReq) GetGatewayNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayNetwork) {
		return nil, false
	}
	return o.GatewayNetwork, true
}

// HasGatewayNetwork returns a boolean if a field has been set.
func (o *ClusterCreateBootNodeReq) HasGatewayNetwork() bool {
	if o != nil && !IsNil(o.GatewayNetwork) {
		return true
	}

	return false
}

// SetGatewayNetwork gets a reference to the given string and assigns it to the GatewayNetwork field.
func (o *ClusterCreateBootNodeReq) SetGatewayNetwork(v string) {
	o.GatewayNetwork = &v
}

// GetHostId returns the HostId field value
func (o *ClusterCreateBootNodeReq) GetHostId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value
// and a boolean to check if the value has been set.
func (o *ClusterCreateBootNodeReq) GetHostIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostId, true
}

// SetHostId sets field value
func (o *ClusterCreateBootNodeReq) SetHostId(v int64) {
	o.HostId = v
}

// GetPrivateIp returns the PrivateIp field value if set, zero value otherwise.
func (o *ClusterCreateBootNodeReq) GetPrivateIp() string {
	if o == nil || IsNil(o.PrivateIp) {
		var ret string
		return ret
	}
	return *o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreateBootNodeReq) GetPrivateIpOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateIp) {
		return nil, false
	}
	return o.PrivateIp, true
}

// HasPrivateIp returns a boolean if a field has been set.
func (o *ClusterCreateBootNodeReq) HasPrivateIp() bool {
	if o != nil && !IsNil(o.PrivateIp) {
		return true
	}

	return false
}

// SetPrivateIp gets a reference to the given string and assigns it to the PrivateIp field.
func (o *ClusterCreateBootNodeReq) SetPrivateIp(v string) {
	o.PrivateIp = &v
}

// GetPrivateNetwork returns the PrivateNetwork field value
func (o *ClusterCreateBootNodeReq) GetPrivateNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateNetwork
}

// GetPrivateNetworkOk returns a tuple with the PrivateNetwork field value
// and a boolean to check if the value has been set.
func (o *ClusterCreateBootNodeReq) GetPrivateNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateNetwork, true
}

// SetPrivateNetwork sets field value
func (o *ClusterCreateBootNodeReq) SetPrivateNetwork(v string) {
	o.PrivateNetwork = v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *ClusterCreateBootNodeReq) GetPublicIp() string {
	if o == nil || IsNil(o.PublicIp) {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreateBootNodeReq) GetPublicIpOk() (*string, bool) {
	if o == nil || IsNil(o.PublicIp) {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *ClusterCreateBootNodeReq) HasPublicIp() bool {
	if o != nil && !IsNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *ClusterCreateBootNodeReq) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetPublicNetwork returns the PublicNetwork field value
func (o *ClusterCreateBootNodeReq) GetPublicNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicNetwork
}

// GetPublicNetworkOk returns a tuple with the PublicNetwork field value
// and a boolean to check if the value has been set.
func (o *ClusterCreateBootNodeReq) GetPublicNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicNetwork, true
}

// SetPublicNetwork sets field value
func (o *ClusterCreateBootNodeReq) SetPublicNetwork(v string) {
	o.PublicNetwork = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ClusterCreateBootNodeReq) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreateBootNodeReq) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ClusterCreateBootNodeReq) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *ClusterCreateBootNodeReq) SetRoles(v []string) {
	o.Roles = v
}

func (o ClusterCreateBootNodeReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterCreateBootNodeReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayIps) {
		toSerialize["gateway_ips"] = o.GatewayIps
	}
	if !IsNil(o.GatewayNetwork) {
		toSerialize["gateway_network"] = o.GatewayNetwork
	}
	toSerialize["host_id"] = o.HostId
	if !IsNil(o.PrivateIp) {
		toSerialize["private_ip"] = o.PrivateIp
	}
	toSerialize["private_network"] = o.PrivateNetwork
	if !IsNil(o.PublicIp) {
		toSerialize["public_ip"] = o.PublicIp
	}
	toSerialize["public_network"] = o.PublicNetwork
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

func (o *ClusterCreateBootNodeReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host_id",
		"private_network",
		"public_network",
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

	varClusterCreateBootNodeReq := _ClusterCreateBootNodeReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClusterCreateBootNodeReq)

	if err != nil {
		return err
	}

	*o = ClusterCreateBootNodeReq(varClusterCreateBootNodeReq)

	return err
}

type NullableClusterCreateBootNodeReq struct {
	value *ClusterCreateBootNodeReq
	isSet bool
}

func (v NullableClusterCreateBootNodeReq) Get() *ClusterCreateBootNodeReq {
	return v.value
}

func (v *NullableClusterCreateBootNodeReq) Set(val *ClusterCreateBootNodeReq) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterCreateBootNodeReq) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterCreateBootNodeReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterCreateBootNodeReq(val *ClusterCreateBootNodeReq) *NullableClusterCreateBootNodeReq {
	return &NullableClusterCreateBootNodeReq{value: val, isSet: true}
}

func (v NullableClusterCreateBootNodeReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterCreateBootNodeReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


