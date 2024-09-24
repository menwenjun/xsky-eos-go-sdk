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

// checks if the ObjectStoragePolicyUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStoragePolicyUpdateReq{}

// ObjectStoragePolicyUpdateReq struct for ObjectStoragePolicyUpdateReq
type ObjectStoragePolicyUpdateReq struct {
	OsPolicy *ObjectStoragePolicyUpdateReqPolicy `json:"os_policy,omitempty"`
}

// NewObjectStoragePolicyUpdateReq instantiates a new ObjectStoragePolicyUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStoragePolicyUpdateReq() *ObjectStoragePolicyUpdateReq {
	this := ObjectStoragePolicyUpdateReq{}
	return &this
}

// NewObjectStoragePolicyUpdateReqWithDefaults instantiates a new ObjectStoragePolicyUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStoragePolicyUpdateReqWithDefaults() *ObjectStoragePolicyUpdateReq {
	this := ObjectStoragePolicyUpdateReq{}
	return &this
}

// GetOsPolicy returns the OsPolicy field value if set, zero value otherwise.
func (o *ObjectStoragePolicyUpdateReq) GetOsPolicy() ObjectStoragePolicyUpdateReqPolicy {
	if o == nil || IsNil(o.OsPolicy) {
		var ret ObjectStoragePolicyUpdateReqPolicy
		return ret
	}
	return *o.OsPolicy
}

// GetOsPolicyOk returns a tuple with the OsPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePolicyUpdateReq) GetOsPolicyOk() (*ObjectStoragePolicyUpdateReqPolicy, bool) {
	if o == nil || IsNil(o.OsPolicy) {
		return nil, false
	}
	return o.OsPolicy, true
}

// HasOsPolicy returns a boolean if a field has been set.
func (o *ObjectStoragePolicyUpdateReq) HasOsPolicy() bool {
	if o != nil && !IsNil(o.OsPolicy) {
		return true
	}

	return false
}

// SetOsPolicy gets a reference to the given ObjectStoragePolicyUpdateReqPolicy and assigns it to the OsPolicy field.
func (o *ObjectStoragePolicyUpdateReq) SetOsPolicy(v ObjectStoragePolicyUpdateReqPolicy) {
	o.OsPolicy = &v
}

func (o ObjectStoragePolicyUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStoragePolicyUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OsPolicy) {
		toSerialize["os_policy"] = o.OsPolicy
	}
	return toSerialize, nil
}

type NullableObjectStoragePolicyUpdateReq struct {
	value *ObjectStoragePolicyUpdateReq
	isSet bool
}

func (v NullableObjectStoragePolicyUpdateReq) Get() *ObjectStoragePolicyUpdateReq {
	return v.value
}

func (v *NullableObjectStoragePolicyUpdateReq) Set(val *ObjectStoragePolicyUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStoragePolicyUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStoragePolicyUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStoragePolicyUpdateReq(val *ObjectStoragePolicyUpdateReq) *NullableObjectStoragePolicyUpdateReq {
	return &NullableObjectStoragePolicyUpdateReq{value: val, isSet: true}
}

func (v NullableObjectStoragePolicyUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStoragePolicyUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


