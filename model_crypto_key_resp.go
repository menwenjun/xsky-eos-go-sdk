/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CryptoKeyResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CryptoKeyResp{}

// CryptoKeyResp struct for CryptoKeyResp
type CryptoKeyResp struct {
	CryptoKey CryptoKey `json:"crypto_key"`
}

type _CryptoKeyResp CryptoKeyResp

// NewCryptoKeyResp instantiates a new CryptoKeyResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptoKeyResp(cryptoKey CryptoKey) *CryptoKeyResp {
	this := CryptoKeyResp{}
	this.CryptoKey = cryptoKey
	return &this
}

// NewCryptoKeyRespWithDefaults instantiates a new CryptoKeyResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptoKeyRespWithDefaults() *CryptoKeyResp {
	this := CryptoKeyResp{}
	return &this
}

// GetCryptoKey returns the CryptoKey field value
func (o *CryptoKeyResp) GetCryptoKey() CryptoKey {
	if o == nil {
		var ret CryptoKey
		return ret
	}

	return o.CryptoKey
}

// GetCryptoKeyOk returns a tuple with the CryptoKey field value
// and a boolean to check if the value has been set.
func (o *CryptoKeyResp) GetCryptoKeyOk() (*CryptoKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CryptoKey, true
}

// SetCryptoKey sets field value
func (o *CryptoKeyResp) SetCryptoKey(v CryptoKey) {
	o.CryptoKey = v
}

func (o CryptoKeyResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CryptoKeyResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["crypto_key"] = o.CryptoKey
	return toSerialize, nil
}

func (o *CryptoKeyResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"crypto_key",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCryptoKeyResp := _CryptoKeyResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCryptoKeyResp)

	if err != nil {
		return err
	}

	*o = CryptoKeyResp(varCryptoKeyResp)

	return err
}

type NullableCryptoKeyResp struct {
	value *CryptoKeyResp
	isSet bool
}

func (v NullableCryptoKeyResp) Get() *CryptoKeyResp {
	return v.value
}

func (v *NullableCryptoKeyResp) Set(val *CryptoKeyResp) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptoKeyResp) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptoKeyResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptoKeyResp(val *CryptoKeyResp) *NullableCryptoKeyResp {
	return &NullableCryptoKeyResp{value: val, isSet: true}
}

func (v NullableCryptoKeyResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptoKeyResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


