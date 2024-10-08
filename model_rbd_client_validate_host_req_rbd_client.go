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

// checks if the RBDClientValidateHostReqRBDClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RBDClientValidateHostReqRBDClient{}

// RBDClientValidateHostReqRBDClient struct for RBDClientValidateHostReqRBDClient
type RBDClientValidateHostReqRBDClient struct {
	// admin ip
	AdminIp string `json:"admin_ip"`
	// rbdc monitor listen port
	Port *int64 `json:"port,omitempty"`
	// storage server or storage client
	Token *string `json:"token,omitempty"`
}

type _RBDClientValidateHostReqRBDClient RBDClientValidateHostReqRBDClient

// NewRBDClientValidateHostReqRBDClient instantiates a new RBDClientValidateHostReqRBDClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRBDClientValidateHostReqRBDClient(adminIp string) *RBDClientValidateHostReqRBDClient {
	this := RBDClientValidateHostReqRBDClient{}
	this.AdminIp = adminIp
	return &this
}

// NewRBDClientValidateHostReqRBDClientWithDefaults instantiates a new RBDClientValidateHostReqRBDClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRBDClientValidateHostReqRBDClientWithDefaults() *RBDClientValidateHostReqRBDClient {
	this := RBDClientValidateHostReqRBDClient{}
	return &this
}

// GetAdminIp returns the AdminIp field value
func (o *RBDClientValidateHostReqRBDClient) GetAdminIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdminIp
}

// GetAdminIpOk returns a tuple with the AdminIp field value
// and a boolean to check if the value has been set.
func (o *RBDClientValidateHostReqRBDClient) GetAdminIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdminIp, true
}

// SetAdminIp sets field value
func (o *RBDClientValidateHostReqRBDClient) SetAdminIp(v string) {
	o.AdminIp = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *RBDClientValidateHostReqRBDClient) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClientValidateHostReqRBDClient) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *RBDClientValidateHostReqRBDClient) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *RBDClientValidateHostReqRBDClient) SetPort(v int64) {
	o.Port = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RBDClientValidateHostReqRBDClient) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClientValidateHostReqRBDClient) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RBDClientValidateHostReqRBDClient) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RBDClientValidateHostReqRBDClient) SetToken(v string) {
	o.Token = &v
}

func (o RBDClientValidateHostReqRBDClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RBDClientValidateHostReqRBDClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["admin_ip"] = o.AdminIp
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

func (o *RBDClientValidateHostReqRBDClient) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"admin_ip",
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

	varRBDClientValidateHostReqRBDClient := _RBDClientValidateHostReqRBDClient{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRBDClientValidateHostReqRBDClient)

	if err != nil {
		return err
	}

	*o = RBDClientValidateHostReqRBDClient(varRBDClientValidateHostReqRBDClient)

	return err
}

type NullableRBDClientValidateHostReqRBDClient struct {
	value *RBDClientValidateHostReqRBDClient
	isSet bool
}

func (v NullableRBDClientValidateHostReqRBDClient) Get() *RBDClientValidateHostReqRBDClient {
	return v.value
}

func (v *NullableRBDClientValidateHostReqRBDClient) Set(val *RBDClientValidateHostReqRBDClient) {
	v.value = val
	v.isSet = true
}

func (v NullableRBDClientValidateHostReqRBDClient) IsSet() bool {
	return v.isSet
}

func (v *NullableRBDClientValidateHostReqRBDClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRBDClientValidateHostReqRBDClient(val *RBDClientValidateHostReqRBDClient) *NullableRBDClientValidateHostReqRBDClient {
	return &NullableRBDClientValidateHostReqRBDClient{value: val, isSet: true}
}

func (v NullableRBDClientValidateHostReqRBDClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRBDClientValidateHostReqRBDClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


