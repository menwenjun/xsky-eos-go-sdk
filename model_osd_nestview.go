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

// checks if the OsdNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsdNestview{}

// OsdNestview struct for OsdNestview
type OsdNestview struct {
	EncryptEnabled *bool `json:"encrypt_enabled,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	NumaNode *int64 `json:"numa_node,omitempty"`
}

// NewOsdNestview instantiates a new OsdNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsdNestview() *OsdNestview {
	this := OsdNestview{}
	return &this
}

// NewOsdNestviewWithDefaults instantiates a new OsdNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsdNestviewWithDefaults() *OsdNestview {
	this := OsdNestview{}
	return &this
}

// GetEncryptEnabled returns the EncryptEnabled field value if set, zero value otherwise.
func (o *OsdNestview) GetEncryptEnabled() bool {
	if o == nil || IsNil(o.EncryptEnabled) {
		var ret bool
		return ret
	}
	return *o.EncryptEnabled
}

// GetEncryptEnabledOk returns a tuple with the EncryptEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsdNestview) GetEncryptEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptEnabled) {
		return nil, false
	}
	return o.EncryptEnabled, true
}

// HasEncryptEnabled returns a boolean if a field has been set.
func (o *OsdNestview) HasEncryptEnabled() bool {
	if o != nil && !IsNil(o.EncryptEnabled) {
		return true
	}

	return false
}

// SetEncryptEnabled gets a reference to the given bool and assigns it to the EncryptEnabled field.
func (o *OsdNestview) SetEncryptEnabled(v bool) {
	o.EncryptEnabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OsdNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsdNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OsdNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OsdNestview) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsdNestview) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsdNestview) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsdNestview) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsdNestview) SetName(v string) {
	o.Name = &v
}

// GetNumaNode returns the NumaNode field value if set, zero value otherwise.
func (o *OsdNestview) GetNumaNode() int64 {
	if o == nil || IsNil(o.NumaNode) {
		var ret int64
		return ret
	}
	return *o.NumaNode
}

// GetNumaNodeOk returns a tuple with the NumaNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsdNestview) GetNumaNodeOk() (*int64, bool) {
	if o == nil || IsNil(o.NumaNode) {
		return nil, false
	}
	return o.NumaNode, true
}

// HasNumaNode returns a boolean if a field has been set.
func (o *OsdNestview) HasNumaNode() bool {
	if o != nil && !IsNil(o.NumaNode) {
		return true
	}

	return false
}

// SetNumaNode gets a reference to the given int64 and assigns it to the NumaNode field.
func (o *OsdNestview) SetNumaNode(v int64) {
	o.NumaNode = &v
}

func (o OsdNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsdNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EncryptEnabled) {
		toSerialize["encrypt_enabled"] = o.EncryptEnabled
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NumaNode) {
		toSerialize["numa_node"] = o.NumaNode
	}
	return toSerialize, nil
}

type NullableOsdNestview struct {
	value *OsdNestview
	isSet bool
}

func (v NullableOsdNestview) Get() *OsdNestview {
	return v.value
}

func (v *NullableOsdNestview) Set(val *OsdNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableOsdNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableOsdNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsdNestview(val *OsdNestview) *NullableOsdNestview {
	return &NullableOsdNestview{value: val, isSet: true}
}

func (v NullableOsdNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsdNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


