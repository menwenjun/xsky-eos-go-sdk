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

// checks if the IdentityPlatformUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityPlatformUpdateReq{}

// IdentityPlatformUpdateReq struct for IdentityPlatformUpdateReq
type IdentityPlatformUpdateReq struct {
	IdentityPlatform *IdentityPlatformUpdateReqPlatform `json:"identity_platform,omitempty"`
}

// NewIdentityPlatformUpdateReq instantiates a new IdentityPlatformUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityPlatformUpdateReq() *IdentityPlatformUpdateReq {
	this := IdentityPlatformUpdateReq{}
	return &this
}

// NewIdentityPlatformUpdateReqWithDefaults instantiates a new IdentityPlatformUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityPlatformUpdateReqWithDefaults() *IdentityPlatformUpdateReq {
	this := IdentityPlatformUpdateReq{}
	return &this
}

// GetIdentityPlatform returns the IdentityPlatform field value if set, zero value otherwise.
func (o *IdentityPlatformUpdateReq) GetIdentityPlatform() IdentityPlatformUpdateReqPlatform {
	if o == nil || IsNil(o.IdentityPlatform) {
		var ret IdentityPlatformUpdateReqPlatform
		return ret
	}
	return *o.IdentityPlatform
}

// GetIdentityPlatformOk returns a tuple with the IdentityPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityPlatformUpdateReq) GetIdentityPlatformOk() (*IdentityPlatformUpdateReqPlatform, bool) {
	if o == nil || IsNil(o.IdentityPlatform) {
		return nil, false
	}
	return o.IdentityPlatform, true
}

// HasIdentityPlatform returns a boolean if a field has been set.
func (o *IdentityPlatformUpdateReq) HasIdentityPlatform() bool {
	if o != nil && !IsNil(o.IdentityPlatform) {
		return true
	}

	return false
}

// SetIdentityPlatform gets a reference to the given IdentityPlatformUpdateReqPlatform and assigns it to the IdentityPlatform field.
func (o *IdentityPlatformUpdateReq) SetIdentityPlatform(v IdentityPlatformUpdateReqPlatform) {
	o.IdentityPlatform = &v
}

func (o IdentityPlatformUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityPlatformUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdentityPlatform) {
		toSerialize["identity_platform"] = o.IdentityPlatform
	}
	return toSerialize, nil
}

type NullableIdentityPlatformUpdateReq struct {
	value *IdentityPlatformUpdateReq
	isSet bool
}

func (v NullableIdentityPlatformUpdateReq) Get() *IdentityPlatformUpdateReq {
	return v.value
}

func (v *NullableIdentityPlatformUpdateReq) Set(val *IdentityPlatformUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityPlatformUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityPlatformUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityPlatformUpdateReq(val *IdentityPlatformUpdateReq) *NullableIdentityPlatformUpdateReq {
	return &NullableIdentityPlatformUpdateReq{value: val, isSet: true}
}

func (v NullableIdentityPlatformUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityPlatformUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


