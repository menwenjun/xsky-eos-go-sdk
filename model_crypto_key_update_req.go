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

// checks if the CryptoKeyUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CryptoKeyUpdateReq{}

// CryptoKeyUpdateReq struct for CryptoKeyUpdateReq
type CryptoKeyUpdateReq struct {
	CryptoKey *CryptoKeyUpdateReqKey `json:"crypto_key,omitempty"`
}

// NewCryptoKeyUpdateReq instantiates a new CryptoKeyUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptoKeyUpdateReq() *CryptoKeyUpdateReq {
	this := CryptoKeyUpdateReq{}
	return &this
}

// NewCryptoKeyUpdateReqWithDefaults instantiates a new CryptoKeyUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptoKeyUpdateReqWithDefaults() *CryptoKeyUpdateReq {
	this := CryptoKeyUpdateReq{}
	return &this
}

// GetCryptoKey returns the CryptoKey field value if set, zero value otherwise.
func (o *CryptoKeyUpdateReq) GetCryptoKey() CryptoKeyUpdateReqKey {
	if o == nil || IsNil(o.CryptoKey) {
		var ret CryptoKeyUpdateReqKey
		return ret
	}
	return *o.CryptoKey
}

// GetCryptoKeyOk returns a tuple with the CryptoKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CryptoKeyUpdateReq) GetCryptoKeyOk() (*CryptoKeyUpdateReqKey, bool) {
	if o == nil || IsNil(o.CryptoKey) {
		return nil, false
	}
	return o.CryptoKey, true
}

// HasCryptoKey returns a boolean if a field has been set.
func (o *CryptoKeyUpdateReq) HasCryptoKey() bool {
	if o != nil && !IsNil(o.CryptoKey) {
		return true
	}

	return false
}

// SetCryptoKey gets a reference to the given CryptoKeyUpdateReqKey and assigns it to the CryptoKey field.
func (o *CryptoKeyUpdateReq) SetCryptoKey(v CryptoKeyUpdateReqKey) {
	o.CryptoKey = &v
}

func (o CryptoKeyUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CryptoKeyUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CryptoKey) {
		toSerialize["crypto_key"] = o.CryptoKey
	}
	return toSerialize, nil
}

type NullableCryptoKeyUpdateReq struct {
	value *CryptoKeyUpdateReq
	isSet bool
}

func (v NullableCryptoKeyUpdateReq) Get() *CryptoKeyUpdateReq {
	return v.value
}

func (v *NullableCryptoKeyUpdateReq) Set(val *CryptoKeyUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptoKeyUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptoKeyUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptoKeyUpdateReq(val *CryptoKeyUpdateReq) *NullableCryptoKeyUpdateReq {
	return &NullableCryptoKeyUpdateReq{value: val, isSet: true}
}

func (v NullableCryptoKeyUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptoKeyUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


