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

// checks if the Gateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Gateway{}

// Gateway Gateway defines osp gateway infos
type Gateway struct {
	ActionStatus *string `json:"action_status,omitempty"`
	GatewayIp *string `json:"gateway_ip,omitempty"`
	GatewayName *string `json:"gateway_name,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	PlatformIp *string `json:"platform_ip,omitempty"`
	Role *string `json:"role,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewGateway instantiates a new Gateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGateway() *Gateway {
	this := Gateway{}
	return &this
}

// NewGatewayWithDefaults instantiates a new Gateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayWithDefaults() *Gateway {
	this := Gateway{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *Gateway) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *Gateway) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *Gateway) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetGatewayIp returns the GatewayIp field value if set, zero value otherwise.
func (o *Gateway) GetGatewayIp() string {
	if o == nil || IsNil(o.GatewayIp) {
		var ret string
		return ret
	}
	return *o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetGatewayIpOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayIp) {
		return nil, false
	}
	return o.GatewayIp, true
}

// HasGatewayIp returns a boolean if a field has been set.
func (o *Gateway) HasGatewayIp() bool {
	if o != nil && !IsNil(o.GatewayIp) {
		return true
	}

	return false
}

// SetGatewayIp gets a reference to the given string and assigns it to the GatewayIp field.
func (o *Gateway) SetGatewayIp(v string) {
	o.GatewayIp = &v
}

// GetGatewayName returns the GatewayName field value if set, zero value otherwise.
func (o *Gateway) GetGatewayName() string {
	if o == nil || IsNil(o.GatewayName) {
		var ret string
		return ret
	}
	return *o.GatewayName
}

// GetGatewayNameOk returns a tuple with the GatewayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetGatewayNameOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayName) {
		return nil, false
	}
	return o.GatewayName, true
}

// HasGatewayName returns a boolean if a field has been set.
func (o *Gateway) HasGatewayName() bool {
	if o != nil && !IsNil(o.GatewayName) {
		return true
	}

	return false
}

// SetGatewayName gets a reference to the given string and assigns it to the GatewayName field.
func (o *Gateway) SetGatewayName(v string) {
	o.GatewayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Gateway) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Gateway) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Gateway) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Gateway) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Gateway) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Gateway) SetName(v string) {
	o.Name = &v
}

// GetPlatformIp returns the PlatformIp field value if set, zero value otherwise.
func (o *Gateway) GetPlatformIp() string {
	if o == nil || IsNil(o.PlatformIp) {
		var ret string
		return ret
	}
	return *o.PlatformIp
}

// GetPlatformIpOk returns a tuple with the PlatformIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetPlatformIpOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformIp) {
		return nil, false
	}
	return o.PlatformIp, true
}

// HasPlatformIp returns a boolean if a field has been set.
func (o *Gateway) HasPlatformIp() bool {
	if o != nil && !IsNil(o.PlatformIp) {
		return true
	}

	return false
}

// SetPlatformIp gets a reference to the given string and assigns it to the PlatformIp field.
func (o *Gateway) SetPlatformIp(v string) {
	o.PlatformIp = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *Gateway) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *Gateway) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *Gateway) SetRole(v string) {
	o.Role = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Gateway) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Gateway) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Gateway) SetStatus(v string) {
	o.Status = &v
}

func (o Gateway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Gateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionStatus) {
		toSerialize["action_status"] = o.ActionStatus
	}
	if !IsNil(o.GatewayIp) {
		toSerialize["gateway_ip"] = o.GatewayIp
	}
	if !IsNil(o.GatewayName) {
		toSerialize["gateway_name"] = o.GatewayName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PlatformIp) {
		toSerialize["platform_ip"] = o.PlatformIp
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableGateway struct {
	value *Gateway
	isSet bool
}

func (v NullableGateway) Get() *Gateway {
	return v.value
}

func (v *NullableGateway) Set(val *Gateway) {
	v.value = val
	v.isSet = true
}

func (v NullableGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGateway(val *Gateway) *NullableGateway {
	return &NullableGateway{value: val, isSet: true}
}

func (v NullableGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


