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

// checks if the WechatContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WechatContact{}

// WechatContact struct for WechatContact
type WechatContact struct {
	ContactId *string `json:"contact_id,omitempty"`
	WechatContactType *string `json:"wechat_contact_type,omitempty"`
}

// NewWechatContact instantiates a new WechatContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWechatContact() *WechatContact {
	this := WechatContact{}
	return &this
}

// NewWechatContactWithDefaults instantiates a new WechatContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWechatContactWithDefaults() *WechatContact {
	this := WechatContact{}
	return &this
}

// GetContactId returns the ContactId field value if set, zero value otherwise.
func (o *WechatContact) GetContactId() string {
	if o == nil || IsNil(o.ContactId) {
		var ret string
		return ret
	}
	return *o.ContactId
}

// GetContactIdOk returns a tuple with the ContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WechatContact) GetContactIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContactId) {
		return nil, false
	}
	return o.ContactId, true
}

// HasContactId returns a boolean if a field has been set.
func (o *WechatContact) HasContactId() bool {
	if o != nil && !IsNil(o.ContactId) {
		return true
	}

	return false
}

// SetContactId gets a reference to the given string and assigns it to the ContactId field.
func (o *WechatContact) SetContactId(v string) {
	o.ContactId = &v
}

// GetWechatContactType returns the WechatContactType field value if set, zero value otherwise.
func (o *WechatContact) GetWechatContactType() string {
	if o == nil || IsNil(o.WechatContactType) {
		var ret string
		return ret
	}
	return *o.WechatContactType
}

// GetWechatContactTypeOk returns a tuple with the WechatContactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WechatContact) GetWechatContactTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WechatContactType) {
		return nil, false
	}
	return o.WechatContactType, true
}

// HasWechatContactType returns a boolean if a field has been set.
func (o *WechatContact) HasWechatContactType() bool {
	if o != nil && !IsNil(o.WechatContactType) {
		return true
	}

	return false
}

// SetWechatContactType gets a reference to the given string and assigns it to the WechatContactType field.
func (o *WechatContact) SetWechatContactType(v string) {
	o.WechatContactType = &v
}

func (o WechatContact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WechatContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactId) {
		toSerialize["contact_id"] = o.ContactId
	}
	if !IsNil(o.WechatContactType) {
		toSerialize["wechat_contact_type"] = o.WechatContactType
	}
	return toSerialize, nil
}

type NullableWechatContact struct {
	value *WechatContact
	isSet bool
}

func (v NullableWechatContact) Get() *WechatContact {
	return v.value
}

func (v *NullableWechatContact) Set(val *WechatContact) {
	v.value = val
	v.isSet = true
}

func (v NullableWechatContact) IsSet() bool {
	return v.isSet
}

func (v *NullableWechatContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWechatContact(val *WechatContact) *NullableWechatContact {
	return &NullableWechatContact{value: val, isSet: true}
}

func (v NullableWechatContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWechatContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


