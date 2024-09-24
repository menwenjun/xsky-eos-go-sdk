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

// checks if the OSBuiltinBluRayUpdateReqBuiltinBluRay type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSBuiltinBluRayUpdateReqBuiltinBluRay{}

// OSBuiltinBluRayUpdateReqBuiltinBluRay struct for OSBuiltinBluRayUpdateReqBuiltinBluRay
type OSBuiltinBluRayUpdateReqBuiltinBluRay struct {
	AdminEndpoint *string `json:"admin_endpoint,omitempty"`
	Name *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
	UserName *string `json:"user_name,omitempty"`
}

// NewOSBuiltinBluRayUpdateReqBuiltinBluRay instantiates a new OSBuiltinBluRayUpdateReqBuiltinBluRay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSBuiltinBluRayUpdateReqBuiltinBluRay() *OSBuiltinBluRayUpdateReqBuiltinBluRay {
	this := OSBuiltinBluRayUpdateReqBuiltinBluRay{}
	return &this
}

// NewOSBuiltinBluRayUpdateReqBuiltinBluRayWithDefaults instantiates a new OSBuiltinBluRayUpdateReqBuiltinBluRay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSBuiltinBluRayUpdateReqBuiltinBluRayWithDefaults() *OSBuiltinBluRayUpdateReqBuiltinBluRay {
	this := OSBuiltinBluRayUpdateReqBuiltinBluRay{}
	return &this
}

// GetAdminEndpoint returns the AdminEndpoint field value if set, zero value otherwise.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) GetAdminEndpoint() string {
	if o == nil || IsNil(o.AdminEndpoint) {
		var ret string
		return ret
	}
	return *o.AdminEndpoint
}

// GetAdminEndpointOk returns a tuple with the AdminEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) GetAdminEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.AdminEndpoint) {
		return nil, false
	}
	return o.AdminEndpoint, true
}

// HasAdminEndpoint returns a boolean if a field has been set.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) HasAdminEndpoint() bool {
	if o != nil && !IsNil(o.AdminEndpoint) {
		return true
	}

	return false
}

// SetAdminEndpoint gets a reference to the given string and assigns it to the AdminEndpoint field.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) SetAdminEndpoint(v string) {
	o.AdminEndpoint = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) SetPassword(v string) {
	o.Password = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *OSBuiltinBluRayUpdateReqBuiltinBluRay) SetUserName(v string) {
	o.UserName = &v
}

func (o OSBuiltinBluRayUpdateReqBuiltinBluRay) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSBuiltinBluRayUpdateReqBuiltinBluRay) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminEndpoint) {
		toSerialize["admin_endpoint"] = o.AdminEndpoint
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.UserName) {
		toSerialize["user_name"] = o.UserName
	}
	return toSerialize, nil
}

type NullableOSBuiltinBluRayUpdateReqBuiltinBluRay struct {
	value *OSBuiltinBluRayUpdateReqBuiltinBluRay
	isSet bool
}

func (v NullableOSBuiltinBluRayUpdateReqBuiltinBluRay) Get() *OSBuiltinBluRayUpdateReqBuiltinBluRay {
	return v.value
}

func (v *NullableOSBuiltinBluRayUpdateReqBuiltinBluRay) Set(val *OSBuiltinBluRayUpdateReqBuiltinBluRay) {
	v.value = val
	v.isSet = true
}

func (v NullableOSBuiltinBluRayUpdateReqBuiltinBluRay) IsSet() bool {
	return v.isSet
}

func (v *NullableOSBuiltinBluRayUpdateReqBuiltinBluRay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSBuiltinBluRayUpdateReqBuiltinBluRay(val *OSBuiltinBluRayUpdateReqBuiltinBluRay) *NullableOSBuiltinBluRayUpdateReqBuiltinBluRay {
	return &NullableOSBuiltinBluRayUpdateReqBuiltinBluRay{value: val, isSet: true}
}

func (v NullableOSBuiltinBluRayUpdateReqBuiltinBluRay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSBuiltinBluRayUpdateReqBuiltinBluRay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


