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

// checks if the CryptoKeyUpdateReqKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CryptoKeyUpdateReqKey{}

// CryptoKeyUpdateReqKey struct for CryptoKeyUpdateReqKey
type CryptoKeyUpdateReqKey struct {
	// enable or disable the key
	Enabled *bool `json:"enabled,omitempty"`
}

// NewCryptoKeyUpdateReqKey instantiates a new CryptoKeyUpdateReqKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptoKeyUpdateReqKey() *CryptoKeyUpdateReqKey {
	this := CryptoKeyUpdateReqKey{}
	return &this
}

// NewCryptoKeyUpdateReqKeyWithDefaults instantiates a new CryptoKeyUpdateReqKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptoKeyUpdateReqKeyWithDefaults() *CryptoKeyUpdateReqKey {
	this := CryptoKeyUpdateReqKey{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CryptoKeyUpdateReqKey) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CryptoKeyUpdateReqKey) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CryptoKeyUpdateReqKey) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CryptoKeyUpdateReqKey) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CryptoKeyUpdateReqKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CryptoKeyUpdateReqKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableCryptoKeyUpdateReqKey struct {
	value *CryptoKeyUpdateReqKey
	isSet bool
}

func (v NullableCryptoKeyUpdateReqKey) Get() *CryptoKeyUpdateReqKey {
	return v.value
}

func (v *NullableCryptoKeyUpdateReqKey) Set(val *CryptoKeyUpdateReqKey) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptoKeyUpdateReqKey) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptoKeyUpdateReqKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptoKeyUpdateReqKey(val *CryptoKeyUpdateReqKey) *NullableCryptoKeyUpdateReqKey {
	return &NullableCryptoKeyUpdateReqKey{value: val, isSet: true}
}

func (v NullableCryptoKeyUpdateReqKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptoKeyUpdateReqKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


