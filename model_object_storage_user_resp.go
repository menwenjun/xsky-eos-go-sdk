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

// checks if the ObjectStorageUserResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStorageUserResp{}

// ObjectStorageUserResp struct for ObjectStorageUserResp
type ObjectStorageUserResp struct {
	OsUser ObjectStorageUserRecord `json:"os_user"`
}

type _ObjectStorageUserResp ObjectStorageUserResp

// NewObjectStorageUserResp instantiates a new ObjectStorageUserResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStorageUserResp(osUser ObjectStorageUserRecord) *ObjectStorageUserResp {
	this := ObjectStorageUserResp{}
	this.OsUser = osUser
	return &this
}

// NewObjectStorageUserRespWithDefaults instantiates a new ObjectStorageUserResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStorageUserRespWithDefaults() *ObjectStorageUserResp {
	this := ObjectStorageUserResp{}
	return &this
}

// GetOsUser returns the OsUser field value
func (o *ObjectStorageUserResp) GetOsUser() ObjectStorageUserRecord {
	if o == nil {
		var ret ObjectStorageUserRecord
		return ret
	}

	return o.OsUser
}

// GetOsUserOk returns a tuple with the OsUser field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageUserResp) GetOsUserOk() (*ObjectStorageUserRecord, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsUser, true
}

// SetOsUser sets field value
func (o *ObjectStorageUserResp) SetOsUser(v ObjectStorageUserRecord) {
	o.OsUser = v
}

func (o ObjectStorageUserResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStorageUserResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_user"] = o.OsUser
	return toSerialize, nil
}

func (o *ObjectStorageUserResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"os_user",
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

	varObjectStorageUserResp := _ObjectStorageUserResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObjectStorageUserResp)

	if err != nil {
		return err
	}

	*o = ObjectStorageUserResp(varObjectStorageUserResp)

	return err
}

type NullableObjectStorageUserResp struct {
	value *ObjectStorageUserResp
	isSet bool
}

func (v NullableObjectStorageUserResp) Get() *ObjectStorageUserResp {
	return v.value
}

func (v *NullableObjectStorageUserResp) Set(val *ObjectStorageUserResp) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStorageUserResp) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStorageUserResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStorageUserResp(val *ObjectStorageUserResp) *NullableObjectStorageUserResp {
	return &NullableObjectStorageUserResp{value: val, isSet: true}
}

func (v NullableObjectStorageUserResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStorageUserResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


