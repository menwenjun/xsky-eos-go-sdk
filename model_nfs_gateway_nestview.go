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

// checks if the NFSGatewayNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NFSGatewayNestview{}

// NFSGatewayNestview struct for NFSGatewayNestview
type NFSGatewayNestview struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewNFSGatewayNestview instantiates a new NFSGatewayNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFSGatewayNestview() *NFSGatewayNestview {
	this := NFSGatewayNestview{}
	return &this
}

// NewNFSGatewayNestviewWithDefaults instantiates a new NFSGatewayNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFSGatewayNestviewWithDefaults() *NFSGatewayNestview {
	this := NFSGatewayNestview{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NFSGatewayNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGatewayNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NFSGatewayNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *NFSGatewayNestview) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NFSGatewayNestview) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGatewayNestview) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NFSGatewayNestview) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NFSGatewayNestview) SetName(v string) {
	o.Name = &v
}

func (o NFSGatewayNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NFSGatewayNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableNFSGatewayNestview struct {
	value *NFSGatewayNestview
	isSet bool
}

func (v NullableNFSGatewayNestview) Get() *NFSGatewayNestview {
	return v.value
}

func (v *NullableNFSGatewayNestview) Set(val *NFSGatewayNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableNFSGatewayNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableNFSGatewayNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFSGatewayNestview(val *NFSGatewayNestview) *NullableNFSGatewayNestview {
	return &NullableNFSGatewayNestview{value: val, isSet: true}
}

func (v NullableNFSGatewayNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFSGatewayNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


