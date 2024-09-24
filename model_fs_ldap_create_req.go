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

// checks if the FSLdapCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FSLdapCreateReq{}

// FSLdapCreateReq struct for FSLdapCreateReq
type FSLdapCreateReq struct {
	FsLdap FSLdapCreateReqInfo `json:"fs_ldap"`
}

type _FSLdapCreateReq FSLdapCreateReq

// NewFSLdapCreateReq instantiates a new FSLdapCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFSLdapCreateReq(fsLdap FSLdapCreateReqInfo) *FSLdapCreateReq {
	this := FSLdapCreateReq{}
	this.FsLdap = fsLdap
	return &this
}

// NewFSLdapCreateReqWithDefaults instantiates a new FSLdapCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFSLdapCreateReqWithDefaults() *FSLdapCreateReq {
	this := FSLdapCreateReq{}
	return &this
}

// GetFsLdap returns the FsLdap field value
func (o *FSLdapCreateReq) GetFsLdap() FSLdapCreateReqInfo {
	if o == nil {
		var ret FSLdapCreateReqInfo
		return ret
	}

	return o.FsLdap
}

// GetFsLdapOk returns a tuple with the FsLdap field value
// and a boolean to check if the value has been set.
func (o *FSLdapCreateReq) GetFsLdapOk() (*FSLdapCreateReqInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FsLdap, true
}

// SetFsLdap sets field value
func (o *FSLdapCreateReq) SetFsLdap(v FSLdapCreateReqInfo) {
	o.FsLdap = v
}

func (o FSLdapCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FSLdapCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fs_ldap"] = o.FsLdap
	return toSerialize, nil
}

func (o *FSLdapCreateReq) UnmarshalJSON(data []byte) (err error) {
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

	varFSLdapCreateReq := _FSLdapCreateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFSLdapCreateReq)

	if err != nil {
		return err
	}

	*o = FSLdapCreateReq(varFSLdapCreateReq)

	return err
}

type NullableFSLdapCreateReq struct {
	value *FSLdapCreateReq
	isSet bool
}

func (v NullableFSLdapCreateReq) Get() *FSLdapCreateReq {
	return v.value
}

func (v *NullableFSLdapCreateReq) Set(val *FSLdapCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableFSLdapCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableFSLdapCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFSLdapCreateReq(val *FSLdapCreateReq) *NullableFSLdapCreateReq {
	return &NullableFSLdapCreateReq{value: val, isSet: true}
}

func (v NullableFSLdapCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFSLdapCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


