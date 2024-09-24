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

// checks if the DfsSMBWindowsACLCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsSMBWindowsACLCreateReq{}

// DfsSMBWindowsACLCreateReq struct for DfsSMBWindowsACLCreateReq
type DfsSMBWindowsACLCreateReq struct {
	DfsSmbWindowsAcl DfsSMBWindowsACLCreateReqWindowsACL `json:"dfs_smb_windows_acl"`
}

type _DfsSMBWindowsACLCreateReq DfsSMBWindowsACLCreateReq

// NewDfsSMBWindowsACLCreateReq instantiates a new DfsSMBWindowsACLCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsSMBWindowsACLCreateReq(dfsSmbWindowsAcl DfsSMBWindowsACLCreateReqWindowsACL) *DfsSMBWindowsACLCreateReq {
	this := DfsSMBWindowsACLCreateReq{}
	this.DfsSmbWindowsAcl = dfsSmbWindowsAcl
	return &this
}

// NewDfsSMBWindowsACLCreateReqWithDefaults instantiates a new DfsSMBWindowsACLCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsSMBWindowsACLCreateReqWithDefaults() *DfsSMBWindowsACLCreateReq {
	this := DfsSMBWindowsACLCreateReq{}
	return &this
}

// GetDfsSmbWindowsAcl returns the DfsSmbWindowsAcl field value
func (o *DfsSMBWindowsACLCreateReq) GetDfsSmbWindowsAcl() DfsSMBWindowsACLCreateReqWindowsACL {
	if o == nil {
		var ret DfsSMBWindowsACLCreateReqWindowsACL
		return ret
	}

	return o.DfsSmbWindowsAcl
}

// GetDfsSmbWindowsAclOk returns a tuple with the DfsSmbWindowsAcl field value
// and a boolean to check if the value has been set.
func (o *DfsSMBWindowsACLCreateReq) GetDfsSmbWindowsAclOk() (*DfsSMBWindowsACLCreateReqWindowsACL, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsSmbWindowsAcl, true
}

// SetDfsSmbWindowsAcl sets field value
func (o *DfsSMBWindowsACLCreateReq) SetDfsSmbWindowsAcl(v DfsSMBWindowsACLCreateReqWindowsACL) {
	o.DfsSmbWindowsAcl = v
}

func (o DfsSMBWindowsACLCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsSMBWindowsACLCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_smb_windows_acl"] = o.DfsSmbWindowsAcl
	return toSerialize, nil
}

func (o *DfsSMBWindowsACLCreateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_smb_windows_acl",
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

	varDfsSMBWindowsACLCreateReq := _DfsSMBWindowsACLCreateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsSMBWindowsACLCreateReq)

	if err != nil {
		return err
	}

	*o = DfsSMBWindowsACLCreateReq(varDfsSMBWindowsACLCreateReq)

	return err
}

type NullableDfsSMBWindowsACLCreateReq struct {
	value *DfsSMBWindowsACLCreateReq
	isSet bool
}

func (v NullableDfsSMBWindowsACLCreateReq) Get() *DfsSMBWindowsACLCreateReq {
	return v.value
}

func (v *NullableDfsSMBWindowsACLCreateReq) Set(val *DfsSMBWindowsACLCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsSMBWindowsACLCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsSMBWindowsACLCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsSMBWindowsACLCreateReq(val *DfsSMBWindowsACLCreateReq) *NullableDfsSMBWindowsACLCreateReq {
	return &NullableDfsSMBWindowsACLCreateReq{value: val, isSet: true}
}

func (v NullableDfsSMBWindowsACLCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsSMBWindowsACLCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


