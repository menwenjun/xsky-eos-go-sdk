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

// checks if the HostNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostNestview{}

// HostNestview struct for HostNestview
type HostNestview struct {
	AdminIp string `json:"admin_ip"`
	GatewayIps *string `json:"gateway_ips,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	NumaNodes *string `json:"numa_nodes,omitempty"`
	Up *bool `json:"up,omitempty"`
}

type _HostNestview HostNestview

// NewHostNestview instantiates a new HostNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostNestview(adminIp string) *HostNestview {
	this := HostNestview{}
	this.AdminIp = adminIp
	return &this
}

// NewHostNestviewWithDefaults instantiates a new HostNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostNestviewWithDefaults() *HostNestview {
	this := HostNestview{}
	return &this
}

// GetAdminIp returns the AdminIp field value
func (o *HostNestview) GetAdminIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdminIp
}

// GetAdminIpOk returns a tuple with the AdminIp field value
// and a boolean to check if the value has been set.
func (o *HostNestview) GetAdminIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdminIp, true
}

// SetAdminIp sets field value
func (o *HostNestview) SetAdminIp(v string) {
	o.AdminIp = v
}

// GetGatewayIps returns the GatewayIps field value if set, zero value otherwise.
func (o *HostNestview) GetGatewayIps() string {
	if o == nil || IsNil(o.GatewayIps) {
		var ret string
		return ret
	}
	return *o.GatewayIps
}

// GetGatewayIpsOk returns a tuple with the GatewayIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostNestview) GetGatewayIpsOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayIps) {
		return nil, false
	}
	return o.GatewayIps, true
}

// HasGatewayIps returns a boolean if a field has been set.
func (o *HostNestview) HasGatewayIps() bool {
	if o != nil && !IsNil(o.GatewayIps) {
		return true
	}

	return false
}

// SetGatewayIps gets a reference to the given string and assigns it to the GatewayIps field.
func (o *HostNestview) SetGatewayIps(v string) {
	o.GatewayIps = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HostNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HostNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *HostNestview) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HostNestview) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostNestview) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HostNestview) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HostNestview) SetName(v string) {
	o.Name = &v
}

// GetNumaNodes returns the NumaNodes field value if set, zero value otherwise.
func (o *HostNestview) GetNumaNodes() string {
	if o == nil || IsNil(o.NumaNodes) {
		var ret string
		return ret
	}
	return *o.NumaNodes
}

// GetNumaNodesOk returns a tuple with the NumaNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostNestview) GetNumaNodesOk() (*string, bool) {
	if o == nil || IsNil(o.NumaNodes) {
		return nil, false
	}
	return o.NumaNodes, true
}

// HasNumaNodes returns a boolean if a field has been set.
func (o *HostNestview) HasNumaNodes() bool {
	if o != nil && !IsNil(o.NumaNodes) {
		return true
	}

	return false
}

// SetNumaNodes gets a reference to the given string and assigns it to the NumaNodes field.
func (o *HostNestview) SetNumaNodes(v string) {
	o.NumaNodes = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *HostNestview) GetUp() bool {
	if o == nil || IsNil(o.Up) {
		var ret bool
		return ret
	}
	return *o.Up
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostNestview) GetUpOk() (*bool, bool) {
	if o == nil || IsNil(o.Up) {
		return nil, false
	}
	return o.Up, true
}

// HasUp returns a boolean if a field has been set.
func (o *HostNestview) HasUp() bool {
	if o != nil && !IsNil(o.Up) {
		return true
	}

	return false
}

// SetUp gets a reference to the given bool and assigns it to the Up field.
func (o *HostNestview) SetUp(v bool) {
	o.Up = &v
}

func (o HostNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["admin_ip"] = o.AdminIp
	if !IsNil(o.GatewayIps) {
		toSerialize["gateway_ips"] = o.GatewayIps
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NumaNodes) {
		toSerialize["numa_nodes"] = o.NumaNodes
	}
	if !IsNil(o.Up) {
		toSerialize["up"] = o.Up
	}
	return toSerialize, nil
}

func (o *HostNestview) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"admin_ip",
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

	varHostNestview := _HostNestview{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHostNestview)

	if err != nil {
		return err
	}

	*o = HostNestview(varHostNestview)

	return err
}

type NullableHostNestview struct {
	value *HostNestview
	isSet bool
}

func (v NullableHostNestview) Get() *HostNestview {
	return v.value
}

func (v *NullableHostNestview) Set(val *HostNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableHostNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableHostNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostNestview(val *HostNestview) *NullableHostNestview {
	return &NullableHostNestview{value: val, isSet: true}
}

func (v NullableHostNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


