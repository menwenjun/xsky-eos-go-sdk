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

// checks if the ObjectStorageUsersResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStorageUsersResp{}

// ObjectStorageUsersResp struct for ObjectStorageUsersResp
type ObjectStorageUsersResp struct {
	// object storage users
	OsUsers []ObjectStorageUserRecord `json:"os_users"`
}

type _ObjectStorageUsersResp ObjectStorageUsersResp

// NewObjectStorageUsersResp instantiates a new ObjectStorageUsersResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStorageUsersResp(osUsers []ObjectStorageUserRecord) *ObjectStorageUsersResp {
	this := ObjectStorageUsersResp{}
	this.OsUsers = osUsers
	return &this
}

// NewObjectStorageUsersRespWithDefaults instantiates a new ObjectStorageUsersResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStorageUsersRespWithDefaults() *ObjectStorageUsersResp {
	this := ObjectStorageUsersResp{}
	return &this
}

// GetOsUsers returns the OsUsers field value
func (o *ObjectStorageUsersResp) GetOsUsers() []ObjectStorageUserRecord {
	if o == nil {
		var ret []ObjectStorageUserRecord
		return ret
	}

	return o.OsUsers
}

// GetOsUsersOk returns a tuple with the OsUsers field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageUsersResp) GetOsUsersOk() ([]ObjectStorageUserRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsUsers, true
}

// SetOsUsers sets field value
func (o *ObjectStorageUsersResp) SetOsUsers(v []ObjectStorageUserRecord) {
	o.OsUsers = v
}

func (o ObjectStorageUsersResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStorageUsersResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_users"] = o.OsUsers
	return toSerialize, nil
}

func (o *ObjectStorageUsersResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"os_users",
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

	varObjectStorageUsersResp := _ObjectStorageUsersResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObjectStorageUsersResp)

	if err != nil {
		return err
	}

	*o = ObjectStorageUsersResp(varObjectStorageUsersResp)

	return err
}

type NullableObjectStorageUsersResp struct {
	value *ObjectStorageUsersResp
	isSet bool
}

func (v NullableObjectStorageUsersResp) Get() *ObjectStorageUsersResp {
	return v.value
}

func (v *NullableObjectStorageUsersResp) Set(val *ObjectStorageUsersResp) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStorageUsersResp) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStorageUsersResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStorageUsersResp(val *ObjectStorageUsersResp) *NullableObjectStorageUsersResp {
	return &NullableObjectStorageUsersResp{value: val, isSet: true}
}

func (v NullableObjectStorageUsersResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStorageUsersResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


