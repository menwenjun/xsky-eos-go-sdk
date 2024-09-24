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

// checks if the DfsSMBShareAddACLsReqShare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsSMBShareAddACLsReqShare{}

// DfsSMBShareAddACLsReqShare struct for DfsSMBShareAddACLsReqShare
type DfsSMBShareAddACLsReqShare struct {
	// access control array
	DfsSmbShareAcls []DfsSMBShareACLReq `json:"dfs_smb_share_acls"`
}

type _DfsSMBShareAddACLsReqShare DfsSMBShareAddACLsReqShare

// NewDfsSMBShareAddACLsReqShare instantiates a new DfsSMBShareAddACLsReqShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsSMBShareAddACLsReqShare(dfsSmbShareAcls []DfsSMBShareACLReq) *DfsSMBShareAddACLsReqShare {
	this := DfsSMBShareAddACLsReqShare{}
	this.DfsSmbShareAcls = dfsSmbShareAcls
	return &this
}

// NewDfsSMBShareAddACLsReqShareWithDefaults instantiates a new DfsSMBShareAddACLsReqShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsSMBShareAddACLsReqShareWithDefaults() *DfsSMBShareAddACLsReqShare {
	this := DfsSMBShareAddACLsReqShare{}
	return &this
}

// GetDfsSmbShareAcls returns the DfsSmbShareAcls field value
func (o *DfsSMBShareAddACLsReqShare) GetDfsSmbShareAcls() []DfsSMBShareACLReq {
	if o == nil {
		var ret []DfsSMBShareACLReq
		return ret
	}

	return o.DfsSmbShareAcls
}

// GetDfsSmbShareAclsOk returns a tuple with the DfsSmbShareAcls field value
// and a boolean to check if the value has been set.
func (o *DfsSMBShareAddACLsReqShare) GetDfsSmbShareAclsOk() ([]DfsSMBShareACLReq, bool) {
	if o == nil {
		return nil, false
	}
	return o.DfsSmbShareAcls, true
}

// SetDfsSmbShareAcls sets field value
func (o *DfsSMBShareAddACLsReqShare) SetDfsSmbShareAcls(v []DfsSMBShareACLReq) {
	o.DfsSmbShareAcls = v
}

func (o DfsSMBShareAddACLsReqShare) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsSMBShareAddACLsReqShare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_smb_share_acls"] = o.DfsSmbShareAcls
	return toSerialize, nil
}

func (o *DfsSMBShareAddACLsReqShare) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_smb_share_acls",
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

	varDfsSMBShareAddACLsReqShare := _DfsSMBShareAddACLsReqShare{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsSMBShareAddACLsReqShare)

	if err != nil {
		return err
	}

	*o = DfsSMBShareAddACLsReqShare(varDfsSMBShareAddACLsReqShare)

	return err
}

type NullableDfsSMBShareAddACLsReqShare struct {
	value *DfsSMBShareAddACLsReqShare
	isSet bool
}

func (v NullableDfsSMBShareAddACLsReqShare) Get() *DfsSMBShareAddACLsReqShare {
	return v.value
}

func (v *NullableDfsSMBShareAddACLsReqShare) Set(val *DfsSMBShareAddACLsReqShare) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsSMBShareAddACLsReqShare) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsSMBShareAddACLsReqShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsSMBShareAddACLsReqShare(val *DfsSMBShareAddACLsReqShare) *NullableDfsSMBShareAddACLsReqShare {
	return &NullableDfsSMBShareAddACLsReqShare{value: val, isSet: true}
}

func (v NullableDfsSMBShareAddACLsReqShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsSMBShareAddACLsReqShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


