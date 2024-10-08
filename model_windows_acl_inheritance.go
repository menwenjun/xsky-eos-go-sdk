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

// checks if the WindowsACLInheritance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WindowsACLInheritance{}

// WindowsACLInheritance WindowsACLInheritance windows acl inheritance info
type WindowsACLInheritance struct {
	ApplyTo *string `json:"apply_to,omitempty"`
	InheritedFrom *string `json:"inherited_from,omitempty"`
	MessageInfo *string `json:"message_info,omitempty"`
}

// NewWindowsACLInheritance instantiates a new WindowsACLInheritance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWindowsACLInheritance() *WindowsACLInheritance {
	this := WindowsACLInheritance{}
	return &this
}

// NewWindowsACLInheritanceWithDefaults instantiates a new WindowsACLInheritance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWindowsACLInheritanceWithDefaults() *WindowsACLInheritance {
	this := WindowsACLInheritance{}
	return &this
}

// GetApplyTo returns the ApplyTo field value if set, zero value otherwise.
func (o *WindowsACLInheritance) GetApplyTo() string {
	if o == nil || IsNil(o.ApplyTo) {
		var ret string
		return ret
	}
	return *o.ApplyTo
}

// GetApplyToOk returns a tuple with the ApplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsACLInheritance) GetApplyToOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyTo) {
		return nil, false
	}
	return o.ApplyTo, true
}

// HasApplyTo returns a boolean if a field has been set.
func (o *WindowsACLInheritance) HasApplyTo() bool {
	if o != nil && !IsNil(o.ApplyTo) {
		return true
	}

	return false
}

// SetApplyTo gets a reference to the given string and assigns it to the ApplyTo field.
func (o *WindowsACLInheritance) SetApplyTo(v string) {
	o.ApplyTo = &v
}

// GetInheritedFrom returns the InheritedFrom field value if set, zero value otherwise.
func (o *WindowsACLInheritance) GetInheritedFrom() string {
	if o == nil || IsNil(o.InheritedFrom) {
		var ret string
		return ret
	}
	return *o.InheritedFrom
}

// GetInheritedFromOk returns a tuple with the InheritedFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsACLInheritance) GetInheritedFromOk() (*string, bool) {
	if o == nil || IsNil(o.InheritedFrom) {
		return nil, false
	}
	return o.InheritedFrom, true
}

// HasInheritedFrom returns a boolean if a field has been set.
func (o *WindowsACLInheritance) HasInheritedFrom() bool {
	if o != nil && !IsNil(o.InheritedFrom) {
		return true
	}

	return false
}

// SetInheritedFrom gets a reference to the given string and assigns it to the InheritedFrom field.
func (o *WindowsACLInheritance) SetInheritedFrom(v string) {
	o.InheritedFrom = &v
}

// GetMessageInfo returns the MessageInfo field value if set, zero value otherwise.
func (o *WindowsACLInheritance) GetMessageInfo() string {
	if o == nil || IsNil(o.MessageInfo) {
		var ret string
		return ret
	}
	return *o.MessageInfo
}

// GetMessageInfoOk returns a tuple with the MessageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsACLInheritance) GetMessageInfoOk() (*string, bool) {
	if o == nil || IsNil(o.MessageInfo) {
		return nil, false
	}
	return o.MessageInfo, true
}

// HasMessageInfo returns a boolean if a field has been set.
func (o *WindowsACLInheritance) HasMessageInfo() bool {
	if o != nil && !IsNil(o.MessageInfo) {
		return true
	}

	return false
}

// SetMessageInfo gets a reference to the given string and assigns it to the MessageInfo field.
func (o *WindowsACLInheritance) SetMessageInfo(v string) {
	o.MessageInfo = &v
}

func (o WindowsACLInheritance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WindowsACLInheritance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyTo) {
		toSerialize["apply_to"] = o.ApplyTo
	}
	if !IsNil(o.InheritedFrom) {
		toSerialize["inherited_from"] = o.InheritedFrom
	}
	if !IsNil(o.MessageInfo) {
		toSerialize["message_info"] = o.MessageInfo
	}
	return toSerialize, nil
}

type NullableWindowsACLInheritance struct {
	value *WindowsACLInheritance
	isSet bool
}

func (v NullableWindowsACLInheritance) Get() *WindowsACLInheritance {
	return v.value
}

func (v *NullableWindowsACLInheritance) Set(val *WindowsACLInheritance) {
	v.value = val
	v.isSet = true
}

func (v NullableWindowsACLInheritance) IsSet() bool {
	return v.isSet
}

func (v *NullableWindowsACLInheritance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWindowsACLInheritance(val *WindowsACLInheritance) *NullableWindowsACLInheritance {
	return &NullableWindowsACLInheritance{value: val, isSet: true}
}

func (v NullableWindowsACLInheritance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWindowsACLInheritance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


