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

// checks if the ObjectStorageGatewayCreateReqGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStorageGatewayCreateReqGateway{}

// ObjectStorageGatewayCreateReqGateway struct for ObjectStorageGatewayCreateReqGateway
type ObjectStorageGatewayCreateReqGateway struct {
	Description *string `json:"description,omitempty"`
	GatewayIp *string `json:"gateway_ip,omitempty"`
	HostId int64 `json:"host_id"`
	Name string `json:"name"`
	Port int64 `json:"port"`
	Role *string `json:"role,omitempty"`
}

type _ObjectStorageGatewayCreateReqGateway ObjectStorageGatewayCreateReqGateway

// NewObjectStorageGatewayCreateReqGateway instantiates a new ObjectStorageGatewayCreateReqGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStorageGatewayCreateReqGateway(hostId int64, name string, port int64) *ObjectStorageGatewayCreateReqGateway {
	this := ObjectStorageGatewayCreateReqGateway{}
	this.HostId = hostId
	this.Name = name
	this.Port = port
	return &this
}

// NewObjectStorageGatewayCreateReqGatewayWithDefaults instantiates a new ObjectStorageGatewayCreateReqGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStorageGatewayCreateReqGatewayWithDefaults() *ObjectStorageGatewayCreateReqGateway {
	this := ObjectStorageGatewayCreateReqGateway{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ObjectStorageGatewayCreateReqGateway) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageGatewayCreateReqGateway) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ObjectStorageGatewayCreateReqGateway) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ObjectStorageGatewayCreateReqGateway) SetDescription(v string) {
	o.Description = &v
}

// GetGatewayIp returns the GatewayIp field value if set, zero value otherwise.
func (o *ObjectStorageGatewayCreateReqGateway) GetGatewayIp() string {
	if o == nil || IsNil(o.GatewayIp) {
		var ret string
		return ret
	}
	return *o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageGatewayCreateReqGateway) GetGatewayIpOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayIp) {
		return nil, false
	}
	return o.GatewayIp, true
}

// HasGatewayIp returns a boolean if a field has been set.
func (o *ObjectStorageGatewayCreateReqGateway) HasGatewayIp() bool {
	if o != nil && !IsNil(o.GatewayIp) {
		return true
	}

	return false
}

// SetGatewayIp gets a reference to the given string and assigns it to the GatewayIp field.
func (o *ObjectStorageGatewayCreateReqGateway) SetGatewayIp(v string) {
	o.GatewayIp = &v
}

// GetHostId returns the HostId field value
func (o *ObjectStorageGatewayCreateReqGateway) GetHostId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageGatewayCreateReqGateway) GetHostIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostId, true
}

// SetHostId sets field value
func (o *ObjectStorageGatewayCreateReqGateway) SetHostId(v int64) {
	o.HostId = v
}

// GetName returns the Name field value
func (o *ObjectStorageGatewayCreateReqGateway) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageGatewayCreateReqGateway) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ObjectStorageGatewayCreateReqGateway) SetName(v string) {
	o.Name = v
}

// GetPort returns the Port field value
func (o *ObjectStorageGatewayCreateReqGateway) GetPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageGatewayCreateReqGateway) GetPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ObjectStorageGatewayCreateReqGateway) SetPort(v int64) {
	o.Port = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ObjectStorageGatewayCreateReqGateway) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageGatewayCreateReqGateway) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ObjectStorageGatewayCreateReqGateway) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ObjectStorageGatewayCreateReqGateway) SetRole(v string) {
	o.Role = &v
}

func (o ObjectStorageGatewayCreateReqGateway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStorageGatewayCreateReqGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.GatewayIp) {
		toSerialize["gateway_ip"] = o.GatewayIp
	}
	toSerialize["host_id"] = o.HostId
	toSerialize["name"] = o.Name
	toSerialize["port"] = o.Port
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

func (o *ObjectStorageGatewayCreateReqGateway) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host_id",
		"name",
		"port",
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

	varObjectStorageGatewayCreateReqGateway := _ObjectStorageGatewayCreateReqGateway{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObjectStorageGatewayCreateReqGateway)

	if err != nil {
		return err
	}

	*o = ObjectStorageGatewayCreateReqGateway(varObjectStorageGatewayCreateReqGateway)

	return err
}

type NullableObjectStorageGatewayCreateReqGateway struct {
	value *ObjectStorageGatewayCreateReqGateway
	isSet bool
}

func (v NullableObjectStorageGatewayCreateReqGateway) Get() *ObjectStorageGatewayCreateReqGateway {
	return v.value
}

func (v *NullableObjectStorageGatewayCreateReqGateway) Set(val *ObjectStorageGatewayCreateReqGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStorageGatewayCreateReqGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStorageGatewayCreateReqGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStorageGatewayCreateReqGateway(val *ObjectStorageGatewayCreateReqGateway) *NullableObjectStorageGatewayCreateReqGateway {
	return &NullableObjectStorageGatewayCreateReqGateway{value: val, isSet: true}
}

func (v NullableObjectStorageGatewayCreateReqGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStorageGatewayCreateReqGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


