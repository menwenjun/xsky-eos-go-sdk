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

// checks if the DNSGatewayGroupRemoveGatewayReqGatewaysElt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DNSGatewayGroupRemoveGatewayReqGatewaysElt{}

// DNSGatewayGroupRemoveGatewayReqGatewaysElt struct for DNSGatewayGroupRemoveGatewayReqGatewaysElt
type DNSGatewayGroupRemoveGatewayReqGatewaysElt struct {
	// dns gateway id
	Id *int64 `json:"id,omitempty"`
}

// NewDNSGatewayGroupRemoveGatewayReqGatewaysElt instantiates a new DNSGatewayGroupRemoveGatewayReqGatewaysElt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNSGatewayGroupRemoveGatewayReqGatewaysElt() *DNSGatewayGroupRemoveGatewayReqGatewaysElt {
	this := DNSGatewayGroupRemoveGatewayReqGatewaysElt{}
	return &this
}

// NewDNSGatewayGroupRemoveGatewayReqGatewaysEltWithDefaults instantiates a new DNSGatewayGroupRemoveGatewayReqGatewaysElt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNSGatewayGroupRemoveGatewayReqGatewaysEltWithDefaults() *DNSGatewayGroupRemoveGatewayReqGatewaysElt {
	this := DNSGatewayGroupRemoveGatewayReqGatewaysElt{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DNSGatewayGroupRemoveGatewayReqGatewaysElt) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSGatewayGroupRemoveGatewayReqGatewaysElt) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DNSGatewayGroupRemoveGatewayReqGatewaysElt) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DNSGatewayGroupRemoveGatewayReqGatewaysElt) SetId(v int64) {
	o.Id = &v
}

func (o DNSGatewayGroupRemoveGatewayReqGatewaysElt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DNSGatewayGroupRemoveGatewayReqGatewaysElt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableDNSGatewayGroupRemoveGatewayReqGatewaysElt struct {
	value *DNSGatewayGroupRemoveGatewayReqGatewaysElt
	isSet bool
}

func (v NullableDNSGatewayGroupRemoveGatewayReqGatewaysElt) Get() *DNSGatewayGroupRemoveGatewayReqGatewaysElt {
	return v.value
}

func (v *NullableDNSGatewayGroupRemoveGatewayReqGatewaysElt) Set(val *DNSGatewayGroupRemoveGatewayReqGatewaysElt) {
	v.value = val
	v.isSet = true
}

func (v NullableDNSGatewayGroupRemoveGatewayReqGatewaysElt) IsSet() bool {
	return v.isSet
}

func (v *NullableDNSGatewayGroupRemoveGatewayReqGatewaysElt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNSGatewayGroupRemoveGatewayReqGatewaysElt(val *DNSGatewayGroupRemoveGatewayReqGatewaysElt) *NullableDNSGatewayGroupRemoveGatewayReqGatewaysElt {
	return &NullableDNSGatewayGroupRemoveGatewayReqGatewaysElt{value: val, isSet: true}
}

func (v NullableDNSGatewayGroupRemoveGatewayReqGatewaysElt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNSGatewayGroupRemoveGatewayReqGatewaysElt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


