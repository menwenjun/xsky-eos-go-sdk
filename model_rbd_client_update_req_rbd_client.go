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

// checks if the RBDClientUpdateReqRBDClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RBDClientUpdateReqRBDClient{}

// RBDClientUpdateReqRBDClient struct for RBDClientUpdateReqRBDClient
type RBDClientUpdateReqRBDClient struct {
	Port *int64 `json:"port,omitempty"`
	Token *string `json:"token,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewRBDClientUpdateReqRBDClient instantiates a new RBDClientUpdateReqRBDClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRBDClientUpdateReqRBDClient() *RBDClientUpdateReqRBDClient {
	this := RBDClientUpdateReqRBDClient{}
	return &this
}

// NewRBDClientUpdateReqRBDClientWithDefaults instantiates a new RBDClientUpdateReqRBDClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRBDClientUpdateReqRBDClientWithDefaults() *RBDClientUpdateReqRBDClient {
	this := RBDClientUpdateReqRBDClient{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *RBDClientUpdateReqRBDClient) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClientUpdateReqRBDClient) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *RBDClientUpdateReqRBDClient) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *RBDClientUpdateReqRBDClient) SetPort(v int64) {
	o.Port = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RBDClientUpdateReqRBDClient) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClientUpdateReqRBDClient) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RBDClientUpdateReqRBDClient) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RBDClientUpdateReqRBDClient) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RBDClientUpdateReqRBDClient) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClientUpdateReqRBDClient) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RBDClientUpdateReqRBDClient) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RBDClientUpdateReqRBDClient) SetType(v string) {
	o.Type = &v
}

func (o RBDClientUpdateReqRBDClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RBDClientUpdateReqRBDClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableRBDClientUpdateReqRBDClient struct {
	value *RBDClientUpdateReqRBDClient
	isSet bool
}

func (v NullableRBDClientUpdateReqRBDClient) Get() *RBDClientUpdateReqRBDClient {
	return v.value
}

func (v *NullableRBDClientUpdateReqRBDClient) Set(val *RBDClientUpdateReqRBDClient) {
	v.value = val
	v.isSet = true
}

func (v NullableRBDClientUpdateReqRBDClient) IsSet() bool {
	return v.isSet
}

func (v *NullableRBDClientUpdateReqRBDClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRBDClientUpdateReqRBDClient(val *RBDClientUpdateReqRBDClient) *NullableRBDClientUpdateReqRBDClient {
	return &NullableRBDClientUpdateReqRBDClient{value: val, isSet: true}
}

func (v NullableRBDClientUpdateReqRBDClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRBDClientUpdateReqRBDClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


