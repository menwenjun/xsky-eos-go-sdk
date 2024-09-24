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

// checks if the DpGatewayNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpGatewayNestview{}

// DpGatewayNestview struct for DpGatewayNestview
type DpGatewayNestview struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewDpGatewayNestview instantiates a new DpGatewayNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpGatewayNestview() *DpGatewayNestview {
	this := DpGatewayNestview{}
	return &this
}

// NewDpGatewayNestviewWithDefaults instantiates a new DpGatewayNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpGatewayNestviewWithDefaults() *DpGatewayNestview {
	this := DpGatewayNestview{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DpGatewayNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpGatewayNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DpGatewayNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DpGatewayNestview) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DpGatewayNestview) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpGatewayNestview) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DpGatewayNestview) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DpGatewayNestview) SetName(v string) {
	o.Name = &v
}

func (o DpGatewayNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpGatewayNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableDpGatewayNestview struct {
	value *DpGatewayNestview
	isSet bool
}

func (v NullableDpGatewayNestview) Get() *DpGatewayNestview {
	return v.value
}

func (v *NullableDpGatewayNestview) Set(val *DpGatewayNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableDpGatewayNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableDpGatewayNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpGatewayNestview(val *DpGatewayNestview) *NullableDpGatewayNestview {
	return &NullableDpGatewayNestview{value: val, isSet: true}
}

func (v NullableDpGatewayNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpGatewayNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


