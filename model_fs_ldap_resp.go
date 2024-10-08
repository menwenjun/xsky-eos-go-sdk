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

// checks if the FSLdapResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FSLdapResp{}

// FSLdapResp struct for FSLdapResp
type FSLdapResp struct {
	FsLdap FSLdap `json:"fs_ldap"`
}

type _FSLdapResp FSLdapResp

// NewFSLdapResp instantiates a new FSLdapResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFSLdapResp(fsLdap FSLdap) *FSLdapResp {
	this := FSLdapResp{}
	this.FsLdap = fsLdap
	return &this
}

// NewFSLdapRespWithDefaults instantiates a new FSLdapResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFSLdapRespWithDefaults() *FSLdapResp {
	this := FSLdapResp{}
	return &this
}

// GetFsLdap returns the FsLdap field value
func (o *FSLdapResp) GetFsLdap() FSLdap {
	if o == nil {
		var ret FSLdap
		return ret
	}

	return o.FsLdap
}

// GetFsLdapOk returns a tuple with the FsLdap field value
// and a boolean to check if the value has been set.
func (o *FSLdapResp) GetFsLdapOk() (*FSLdap, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FsLdap, true
}

// SetFsLdap sets field value
func (o *FSLdapResp) SetFsLdap(v FSLdap) {
	o.FsLdap = v
}

func (o FSLdapResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FSLdapResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fs_ldap"] = o.FsLdap
	return toSerialize, nil
}

func (o *FSLdapResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fs_ldap",
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

	varFSLdapResp := _FSLdapResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFSLdapResp)

	if err != nil {
		return err
	}

	*o = FSLdapResp(varFSLdapResp)

	return err
}

type NullableFSLdapResp struct {
	value *FSLdapResp
	isSet bool
}

func (v NullableFSLdapResp) Get() *FSLdapResp {
	return v.value
}

func (v *NullableFSLdapResp) Set(val *FSLdapResp) {
	v.value = val
	v.isSet = true
}

func (v NullableFSLdapResp) IsSet() bool {
	return v.isSet
}

func (v *NullableFSLdapResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFSLdapResp(val *FSLdapResp) *NullableFSLdapResp {
	return &NullableFSLdapResp{value: val, isSet: true}
}

func (v NullableFSLdapResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFSLdapResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


