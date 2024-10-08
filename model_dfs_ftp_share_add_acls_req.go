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

// checks if the DfsFTPShareAddACLsReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsFTPShareAddACLsReq{}

// DfsFTPShareAddACLsReq struct for DfsFTPShareAddACLsReq
type DfsFTPShareAddACLsReq struct {
	DfsFtpShare DfsFTPShareAddACLsReqShare `json:"dfs_ftp_share"`
}

type _DfsFTPShareAddACLsReq DfsFTPShareAddACLsReq

// NewDfsFTPShareAddACLsReq instantiates a new DfsFTPShareAddACLsReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsFTPShareAddACLsReq(dfsFtpShare DfsFTPShareAddACLsReqShare) *DfsFTPShareAddACLsReq {
	this := DfsFTPShareAddACLsReq{}
	this.DfsFtpShare = dfsFtpShare
	return &this
}

// NewDfsFTPShareAddACLsReqWithDefaults instantiates a new DfsFTPShareAddACLsReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsFTPShareAddACLsReqWithDefaults() *DfsFTPShareAddACLsReq {
	this := DfsFTPShareAddACLsReq{}
	return &this
}

// GetDfsFtpShare returns the DfsFtpShare field value
func (o *DfsFTPShareAddACLsReq) GetDfsFtpShare() DfsFTPShareAddACLsReqShare {
	if o == nil {
		var ret DfsFTPShareAddACLsReqShare
		return ret
	}

	return o.DfsFtpShare
}

// GetDfsFtpShareOk returns a tuple with the DfsFtpShare field value
// and a boolean to check if the value has been set.
func (o *DfsFTPShareAddACLsReq) GetDfsFtpShareOk() (*DfsFTPShareAddACLsReqShare, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsFtpShare, true
}

// SetDfsFtpShare sets field value
func (o *DfsFTPShareAddACLsReq) SetDfsFtpShare(v DfsFTPShareAddACLsReqShare) {
	o.DfsFtpShare = v
}

func (o DfsFTPShareAddACLsReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsFTPShareAddACLsReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_ftp_share"] = o.DfsFtpShare
	return toSerialize, nil
}

func (o *DfsFTPShareAddACLsReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_ftp_share",
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

	varDfsFTPShareAddACLsReq := _DfsFTPShareAddACLsReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsFTPShareAddACLsReq)

	if err != nil {
		return err
	}

	*o = DfsFTPShareAddACLsReq(varDfsFTPShareAddACLsReq)

	return err
}

type NullableDfsFTPShareAddACLsReq struct {
	value *DfsFTPShareAddACLsReq
	isSet bool
}

func (v NullableDfsFTPShareAddACLsReq) Get() *DfsFTPShareAddACLsReq {
	return v.value
}

func (v *NullableDfsFTPShareAddACLsReq) Set(val *DfsFTPShareAddACLsReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsFTPShareAddACLsReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsFTPShareAddACLsReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsFTPShareAddACLsReq(val *DfsFTPShareAddACLsReq) *NullableDfsFTPShareAddACLsReq {
	return &NullableDfsFTPShareAddACLsReq{value: val, isSet: true}
}

func (v NullableDfsFTPShareAddACLsReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsFTPShareAddACLsReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


