/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the AuthRSAKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthRSAKey{}

// AuthRSAKey struct for AuthRSAKey
type AuthRSAKey struct {
	Expiration *time.Time `json:"expiration,omitempty"`
	Id *string `json:"id,omitempty"`
	PublicKey *string `json:"public_key,omitempty"`
}

// NewAuthRSAKey instantiates a new AuthRSAKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthRSAKey() *AuthRSAKey {
	this := AuthRSAKey{}
	return &this
}

// NewAuthRSAKeyWithDefaults instantiates a new AuthRSAKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthRSAKeyWithDefaults() *AuthRSAKey {
	this := AuthRSAKey{}
	return &this
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *AuthRSAKey) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration) {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRSAKey) GetExpirationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *AuthRSAKey) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *AuthRSAKey) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthRSAKey) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRSAKey) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthRSAKey) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthRSAKey) SetId(v string) {
	o.Id = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *AuthRSAKey) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRSAKey) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *AuthRSAKey) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *AuthRSAKey) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o AuthRSAKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthRSAKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PublicKey) {
		toSerialize["public_key"] = o.PublicKey
	}
	return toSerialize, nil
}

type NullableAuthRSAKey struct {
	value *AuthRSAKey
	isSet bool
}

func (v NullableAuthRSAKey) Get() *AuthRSAKey {
	return v.value
}

func (v *NullableAuthRSAKey) Set(val *AuthRSAKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthRSAKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthRSAKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthRSAKey(val *AuthRSAKey) *NullableAuthRSAKey {
	return &NullableAuthRSAKey{value: val, isSet: true}
}

func (v NullableAuthRSAKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthRSAKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


