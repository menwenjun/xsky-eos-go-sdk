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

// checks if the DfsSMBWindowsACLResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsSMBWindowsACLResp{}

// DfsSMBWindowsACLResp struct for DfsSMBWindowsACLResp
type DfsSMBWindowsACLResp struct {
	DfsSmbWindowsAcl DfsSMBWindowsACL `json:"dfs_smb_windows_acl"`
}

type _DfsSMBWindowsACLResp DfsSMBWindowsACLResp

// NewDfsSMBWindowsACLResp instantiates a new DfsSMBWindowsACLResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsSMBWindowsACLResp(dfsSmbWindowsAcl DfsSMBWindowsACL) *DfsSMBWindowsACLResp {
	this := DfsSMBWindowsACLResp{}
	this.DfsSmbWindowsAcl = dfsSmbWindowsAcl
	return &this
}

// NewDfsSMBWindowsACLRespWithDefaults instantiates a new DfsSMBWindowsACLResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsSMBWindowsACLRespWithDefaults() *DfsSMBWindowsACLResp {
	this := DfsSMBWindowsACLResp{}
	return &this
}

// GetDfsSmbWindowsAcl returns the DfsSmbWindowsAcl field value
func (o *DfsSMBWindowsACLResp) GetDfsSmbWindowsAcl() DfsSMBWindowsACL {
	if o == nil {
		var ret DfsSMBWindowsACL
		return ret
	}

	return o.DfsSmbWindowsAcl
}

// GetDfsSmbWindowsAclOk returns a tuple with the DfsSmbWindowsAcl field value
// and a boolean to check if the value has been set.
func (o *DfsSMBWindowsACLResp) GetDfsSmbWindowsAclOk() (*DfsSMBWindowsACL, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsSmbWindowsAcl, true
}

// SetDfsSmbWindowsAcl sets field value
func (o *DfsSMBWindowsACLResp) SetDfsSmbWindowsAcl(v DfsSMBWindowsACL) {
	o.DfsSmbWindowsAcl = v
}

func (o DfsSMBWindowsACLResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsSMBWindowsACLResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_smb_windows_acl"] = o.DfsSmbWindowsAcl
	return toSerialize, nil
}

func (o *DfsSMBWindowsACLResp) UnmarshalJSON(data []byte) (err error) {
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

	varDfsSMBWindowsACLResp := _DfsSMBWindowsACLResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsSMBWindowsACLResp)

	if err != nil {
		return err
	}

	*o = DfsSMBWindowsACLResp(varDfsSMBWindowsACLResp)

	return err
}

type NullableDfsSMBWindowsACLResp struct {
	value *DfsSMBWindowsACLResp
	isSet bool
}

func (v NullableDfsSMBWindowsACLResp) Get() *DfsSMBWindowsACLResp {
	return v.value
}

func (v *NullableDfsSMBWindowsACLResp) Set(val *DfsSMBWindowsACLResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsSMBWindowsACLResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsSMBWindowsACLResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsSMBWindowsACLResp(val *DfsSMBWindowsACLResp) *NullableDfsSMBWindowsACLResp {
	return &NullableDfsSMBWindowsACLResp{value: val, isSet: true}
}

func (v NullableDfsSMBWindowsACLResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsSMBWindowsACLResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


