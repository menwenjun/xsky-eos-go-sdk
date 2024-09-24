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

// checks if the DfsSMBWindowsACLDeleteReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsSMBWindowsACLDeleteReq{}

// DfsSMBWindowsACLDeleteReq struct for DfsSMBWindowsACLDeleteReq
type DfsSMBWindowsACLDeleteReq struct {
	DfsSmbWindowsAcl DfsSMBWindowsACLDeleteReqWindowsACL `json:"dfs_smb_windows_acl"`
}

type _DfsSMBWindowsACLDeleteReq DfsSMBWindowsACLDeleteReq

// NewDfsSMBWindowsACLDeleteReq instantiates a new DfsSMBWindowsACLDeleteReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsSMBWindowsACLDeleteReq(dfsSmbWindowsAcl DfsSMBWindowsACLDeleteReqWindowsACL) *DfsSMBWindowsACLDeleteReq {
	this := DfsSMBWindowsACLDeleteReq{}
	this.DfsSmbWindowsAcl = dfsSmbWindowsAcl
	return &this
}

// NewDfsSMBWindowsACLDeleteReqWithDefaults instantiates a new DfsSMBWindowsACLDeleteReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsSMBWindowsACLDeleteReqWithDefaults() *DfsSMBWindowsACLDeleteReq {
	this := DfsSMBWindowsACLDeleteReq{}
	return &this
}

// GetDfsSmbWindowsAcl returns the DfsSmbWindowsAcl field value
func (o *DfsSMBWindowsACLDeleteReq) GetDfsSmbWindowsAcl() DfsSMBWindowsACLDeleteReqWindowsACL {
	if o == nil {
		var ret DfsSMBWindowsACLDeleteReqWindowsACL
		return ret
	}

	return o.DfsSmbWindowsAcl
}

// GetDfsSmbWindowsAclOk returns a tuple with the DfsSmbWindowsAcl field value
// and a boolean to check if the value has been set.
func (o *DfsSMBWindowsACLDeleteReq) GetDfsSmbWindowsAclOk() (*DfsSMBWindowsACLDeleteReqWindowsACL, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsSmbWindowsAcl, true
}

// SetDfsSmbWindowsAcl sets field value
func (o *DfsSMBWindowsACLDeleteReq) SetDfsSmbWindowsAcl(v DfsSMBWindowsACLDeleteReqWindowsACL) {
	o.DfsSmbWindowsAcl = v
}

func (o DfsSMBWindowsACLDeleteReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsSMBWindowsACLDeleteReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_smb_windows_acl"] = o.DfsSmbWindowsAcl
	return toSerialize, nil
}

func (o *DfsSMBWindowsACLDeleteReq) UnmarshalJSON(data []byte) (err error) {
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

	varDfsSMBWindowsACLDeleteReq := _DfsSMBWindowsACLDeleteReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsSMBWindowsACLDeleteReq)

	if err != nil {
		return err
	}

	*o = DfsSMBWindowsACLDeleteReq(varDfsSMBWindowsACLDeleteReq)

	return err
}

type NullableDfsSMBWindowsACLDeleteReq struct {
	value *DfsSMBWindowsACLDeleteReq
	isSet bool
}

func (v NullableDfsSMBWindowsACLDeleteReq) Get() *DfsSMBWindowsACLDeleteReq {
	return v.value
}

func (v *NullableDfsSMBWindowsACLDeleteReq) Set(val *DfsSMBWindowsACLDeleteReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsSMBWindowsACLDeleteReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsSMBWindowsACLDeleteReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsSMBWindowsACLDeleteReq(val *DfsSMBWindowsACLDeleteReq) *NullableDfsSMBWindowsACLDeleteReq {
	return &NullableDfsSMBWindowsACLDeleteReq{value: val, isSet: true}
}

func (v NullableDfsSMBWindowsACLDeleteReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsSMBWindowsACLDeleteReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


