/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the RoleMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMapping{}

// RoleMapping RoleMapping defines role mappings for auth
type RoleMapping struct {
	// time of creating identity platform
	Create *time.Time `json:"create,omitempty"`
	Default *bool `json:"default,omitempty"`
	// id of role mapping
	Id *int64 `json:"id,omitempty"`
	IdentityPlatform *IdentityPlatformNestview `json:"identity_platform,omitempty"`
	// mapped role of sds platform
	Role *string `json:"role,omitempty"`
	// time of updating identity platform
	Update *time.Time `json:"update,omitempty"`
	// roles of external identity platform
	Value *string `json:"value,omitempty"`
}

// NewRoleMapping instantiates a new RoleMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMapping() *RoleMapping {
	this := RoleMapping{}
	return &this
}

// NewRoleMappingWithDefaults instantiates a new RoleMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMappingWithDefaults() *RoleMapping {
	this := RoleMapping{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *RoleMapping) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMapping) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *RoleMapping) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *RoleMapping) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *RoleMapping) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMapping) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *RoleMapping) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *RoleMapping) SetDefault(v bool) {
	o.Default = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleMapping) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMapping) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleMapping) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *RoleMapping) SetId(v int64) {
	o.Id = &v
}

// GetIdentityPlatform returns the IdentityPlatform field value if set, zero value otherwise.
func (o *RoleMapping) GetIdentityPlatform() IdentityPlatformNestview {
	if o == nil || IsNil(o.IdentityPlatform) {
		var ret IdentityPlatformNestview
		return ret
	}
	return *o.IdentityPlatform
}

// GetIdentityPlatformOk returns a tuple with the IdentityPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMapping) GetIdentityPlatformOk() (*IdentityPlatformNestview, bool) {
	if o == nil || IsNil(o.IdentityPlatform) {
		return nil, false
	}
	return o.IdentityPlatform, true
}

// HasIdentityPlatform returns a boolean if a field has been set.
func (o *RoleMapping) HasIdentityPlatform() bool {
	if o != nil && !IsNil(o.IdentityPlatform) {
		return true
	}

	return false
}

// SetIdentityPlatform gets a reference to the given IdentityPlatformNestview and assigns it to the IdentityPlatform field.
func (o *RoleMapping) SetIdentityPlatform(v IdentityPlatformNestview) {
	o.IdentityPlatform = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *RoleMapping) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMapping) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *RoleMapping) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *RoleMapping) SetRole(v string) {
	o.Role = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *RoleMapping) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMapping) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *RoleMapping) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *RoleMapping) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RoleMapping) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMapping) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RoleMapping) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RoleMapping) SetValue(v string) {
	o.Value = &v
}

func (o RoleMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IdentityPlatform) {
		toSerialize["identity_platform"] = o.IdentityPlatform
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableRoleMapping struct {
	value *RoleMapping
	isSet bool
}

func (v NullableRoleMapping) Get() *RoleMapping {
	return v.value
}

func (v *NullableRoleMapping) Set(val *RoleMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMapping(val *RoleMapping) *NullableRoleMapping {
	return &NullableRoleMapping{value: val, isSet: true}
}

func (v NullableRoleMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


