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

// checks if the DfsFTPShareCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsFTPShareCreateReq{}

// DfsFTPShareCreateReq struct for DfsFTPShareCreateReq
type DfsFTPShareCreateReq struct {
	DfsFtpShare DfsFTPShareCreateReqShare `json:"dfs_ftp_share"`
}

type _DfsFTPShareCreateReq DfsFTPShareCreateReq

// NewDfsFTPShareCreateReq instantiates a new DfsFTPShareCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsFTPShareCreateReq(dfsFtpShare DfsFTPShareCreateReqShare) *DfsFTPShareCreateReq {
	this := DfsFTPShareCreateReq{}
	this.DfsFtpShare = dfsFtpShare
	return &this
}

// NewDfsFTPShareCreateReqWithDefaults instantiates a new DfsFTPShareCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsFTPShareCreateReqWithDefaults() *DfsFTPShareCreateReq {
	this := DfsFTPShareCreateReq{}
	return &this
}

// GetDfsFtpShare returns the DfsFtpShare field value
func (o *DfsFTPShareCreateReq) GetDfsFtpShare() DfsFTPShareCreateReqShare {
	if o == nil {
		var ret DfsFTPShareCreateReqShare
		return ret
	}

	return o.DfsFtpShare
}

// GetDfsFtpShareOk returns a tuple with the DfsFtpShare field value
// and a boolean to check if the value has been set.
func (o *DfsFTPShareCreateReq) GetDfsFtpShareOk() (*DfsFTPShareCreateReqShare, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsFtpShare, true
}

// SetDfsFtpShare sets field value
func (o *DfsFTPShareCreateReq) SetDfsFtpShare(v DfsFTPShareCreateReqShare) {
	o.DfsFtpShare = v
}

func (o DfsFTPShareCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsFTPShareCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_ftp_share"] = o.DfsFtpShare
	return toSerialize, nil
}

func (o *DfsFTPShareCreateReq) UnmarshalJSON(data []byte) (err error) {
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

	varDfsFTPShareCreateReq := _DfsFTPShareCreateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsFTPShareCreateReq)

	if err != nil {
		return err
	}

	*o = DfsFTPShareCreateReq(varDfsFTPShareCreateReq)

	return err
}

type NullableDfsFTPShareCreateReq struct {
	value *DfsFTPShareCreateReq
	isSet bool
}

func (v NullableDfsFTPShareCreateReq) Get() *DfsFTPShareCreateReq {
	return v.value
}

func (v *NullableDfsFTPShareCreateReq) Set(val *DfsFTPShareCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsFTPShareCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsFTPShareCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsFTPShareCreateReq(val *DfsFTPShareCreateReq) *NullableDfsFTPShareCreateReq {
	return &NullableDfsFTPShareCreateReq{value: val, isSet: true}
}

func (v NullableDfsFTPShareCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsFTPShareCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


