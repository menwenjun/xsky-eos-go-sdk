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

// checks if the BootNodeReqBootNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BootNodeReqBootNode{}

// BootNodeReqBootNode struct for BootNodeReqBootNode
type BootNodeReqBootNode struct {
	// admin network: 10.0.0.0/24
	AdminNetwork *string `json:"admin_network,omitempty"`
	// fs id
	FsId *string `json:"fs_id,omitempty"`
	// gateway networks, multiple networks splited by comma: 10.0.3.0/24,10.0.4.0/24
	GatewayNetwork *string `json:"gateway_network,omitempty"`
	// private network : 10.0.2.0/24
	PrivateNetwork string `json:"private_network"`
	// public network: 10.0.1.0/24
	PublicNetwork string `json:"public_network"`
	// region name
	RegionName *string `json:"region_name,omitempty"`
}

type _BootNodeReqBootNode BootNodeReqBootNode

// NewBootNodeReqBootNode instantiates a new BootNodeReqBootNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootNodeReqBootNode(privateNetwork string, publicNetwork string) *BootNodeReqBootNode {
	this := BootNodeReqBootNode{}
	this.PrivateNetwork = privateNetwork
	this.PublicNetwork = publicNetwork
	return &this
}

// NewBootNodeReqBootNodeWithDefaults instantiates a new BootNodeReqBootNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootNodeReqBootNodeWithDefaults() *BootNodeReqBootNode {
	this := BootNodeReqBootNode{}
	return &this
}

// GetAdminNetwork returns the AdminNetwork field value if set, zero value otherwise.
func (o *BootNodeReqBootNode) GetAdminNetwork() string {
	if o == nil || IsNil(o.AdminNetwork) {
		var ret string
		return ret
	}
	return *o.AdminNetwork
}

// GetAdminNetworkOk returns a tuple with the AdminNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootNodeReqBootNode) GetAdminNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.AdminNetwork) {
		return nil, false
	}
	return o.AdminNetwork, true
}

// HasAdminNetwork returns a boolean if a field has been set.
func (o *BootNodeReqBootNode) HasAdminNetwork() bool {
	if o != nil && !IsNil(o.AdminNetwork) {
		return true
	}

	return false
}

// SetAdminNetwork gets a reference to the given string and assigns it to the AdminNetwork field.
func (o *BootNodeReqBootNode) SetAdminNetwork(v string) {
	o.AdminNetwork = &v
}

// GetFsId returns the FsId field value if set, zero value otherwise.
func (o *BootNodeReqBootNode) GetFsId() string {
	if o == nil || IsNil(o.FsId) {
		var ret string
		return ret
	}
	return *o.FsId
}

// GetFsIdOk returns a tuple with the FsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootNodeReqBootNode) GetFsIdOk() (*string, bool) {
	if o == nil || IsNil(o.FsId) {
		return nil, false
	}
	return o.FsId, true
}

// HasFsId returns a boolean if a field has been set.
func (o *BootNodeReqBootNode) HasFsId() bool {
	if o != nil && !IsNil(o.FsId) {
		return true
	}

	return false
}

// SetFsId gets a reference to the given string and assigns it to the FsId field.
func (o *BootNodeReqBootNode) SetFsId(v string) {
	o.FsId = &v
}

// GetGatewayNetwork returns the GatewayNetwork field value if set, zero value otherwise.
func (o *BootNodeReqBootNode) GetGatewayNetwork() string {
	if o == nil || IsNil(o.GatewayNetwork) {
		var ret string
		return ret
	}
	return *o.GatewayNetwork
}

// GetGatewayNetworkOk returns a tuple with the GatewayNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootNodeReqBootNode) GetGatewayNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayNetwork) {
		return nil, false
	}
	return o.GatewayNetwork, true
}

// HasGatewayNetwork returns a boolean if a field has been set.
func (o *BootNodeReqBootNode) HasGatewayNetwork() bool {
	if o != nil && !IsNil(o.GatewayNetwork) {
		return true
	}

	return false
}

// SetGatewayNetwork gets a reference to the given string and assigns it to the GatewayNetwork field.
func (o *BootNodeReqBootNode) SetGatewayNetwork(v string) {
	o.GatewayNetwork = &v
}

// GetPrivateNetwork returns the PrivateNetwork field value
func (o *BootNodeReqBootNode) GetPrivateNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateNetwork
}

// GetPrivateNetworkOk returns a tuple with the PrivateNetwork field value
// and a boolean to check if the value has been set.
func (o *BootNodeReqBootNode) GetPrivateNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateNetwork, true
}

// SetPrivateNetwork sets field value
func (o *BootNodeReqBootNode) SetPrivateNetwork(v string) {
	o.PrivateNetwork = v
}

// GetPublicNetwork returns the PublicNetwork field value
func (o *BootNodeReqBootNode) GetPublicNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicNetwork
}

// GetPublicNetworkOk returns a tuple with the PublicNetwork field value
// and a boolean to check if the value has been set.
func (o *BootNodeReqBootNode) GetPublicNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicNetwork, true
}

// SetPublicNetwork sets field value
func (o *BootNodeReqBootNode) SetPublicNetwork(v string) {
	o.PublicNetwork = v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *BootNodeReqBootNode) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootNodeReqBootNode) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *BootNodeReqBootNode) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *BootNodeReqBootNode) SetRegionName(v string) {
	o.RegionName = &v
}

func (o BootNodeReqBootNode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BootNodeReqBootNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminNetwork) {
		toSerialize["admin_network"] = o.AdminNetwork
	}
	if !IsNil(o.FsId) {
		toSerialize["fs_id"] = o.FsId
	}
	if !IsNil(o.GatewayNetwork) {
		toSerialize["gateway_network"] = o.GatewayNetwork
	}
	toSerialize["private_network"] = o.PrivateNetwork
	toSerialize["public_network"] = o.PublicNetwork
	if !IsNil(o.RegionName) {
		toSerialize["region_name"] = o.RegionName
	}
	return toSerialize, nil
}

func (o *BootNodeReqBootNode) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varBootNodeReqBootNode := _BootNodeReqBootNode{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBootNodeReqBootNode)

	if err != nil {
		return err
	}

	*o = BootNodeReqBootNode(varBootNodeReqBootNode)

	return err
}

type NullableBootNodeReqBootNode struct {
	value *BootNodeReqBootNode
	isSet bool
}

func (v NullableBootNodeReqBootNode) Get() *BootNodeReqBootNode {
	return v.value
}

func (v *NullableBootNodeReqBootNode) Set(val *BootNodeReqBootNode) {
	v.value = val
	v.isSet = true
}

func (v NullableBootNodeReqBootNode) IsSet() bool {
	return v.isSet
}

func (v *NullableBootNodeReqBootNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootNodeReqBootNode(val *BootNodeReqBootNode) *NullableBootNodeReqBootNode {
	return &NullableBootNodeReqBootNode{value: val, isSet: true}
}

func (v NullableBootNodeReqBootNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootNodeReqBootNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


