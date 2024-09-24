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

// checks if the DfsSMBShareResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsSMBShareResp{}

// DfsSMBShareResp struct for DfsSMBShareResp
type DfsSMBShareResp struct {
	DfsSmbShare DfsSMBShareRecord `json:"dfs_smb_share"`
}

type _DfsSMBShareResp DfsSMBShareResp

// NewDfsSMBShareResp instantiates a new DfsSMBShareResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsSMBShareResp(dfsSmbShare DfsSMBShareRecord) *DfsSMBShareResp {
	this := DfsSMBShareResp{}
	this.DfsSmbShare = dfsSmbShare
	return &this
}

// NewDfsSMBShareRespWithDefaults instantiates a new DfsSMBShareResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsSMBShareRespWithDefaults() *DfsSMBShareResp {
	this := DfsSMBShareResp{}
	return &this
}

// GetDfsSmbShare returns the DfsSmbShare field value
func (o *DfsSMBShareResp) GetDfsSmbShare() DfsSMBShareRecord {
	if o == nil {
		var ret DfsSMBShareRecord
		return ret
	}

	return o.DfsSmbShare
}

// GetDfsSmbShareOk returns a tuple with the DfsSmbShare field value
// and a boolean to check if the value has been set.
func (o *DfsSMBShareResp) GetDfsSmbShareOk() (*DfsSMBShareRecord, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsSmbShare, true
}

// SetDfsSmbShare sets field value
func (o *DfsSMBShareResp) SetDfsSmbShare(v DfsSMBShareRecord) {
	o.DfsSmbShare = v
}

func (o DfsSMBShareResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsSMBShareResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_smb_share"] = o.DfsSmbShare
	return toSerialize, nil
}

func (o *DfsSMBShareResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_smb_share",
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

	varDfsSMBShareResp := _DfsSMBShareResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsSMBShareResp)

	if err != nil {
		return err
	}

	*o = DfsSMBShareResp(varDfsSMBShareResp)

	return err
}

type NullableDfsSMBShareResp struct {
	value *DfsSMBShareResp
	isSet bool
}

func (v NullableDfsSMBShareResp) Get() *DfsSMBShareResp {
	return v.value
}

func (v *NullableDfsSMBShareResp) Set(val *DfsSMBShareResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsSMBShareResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsSMBShareResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsSMBShareResp(val *DfsSMBShareResp) *NullableDfsSMBShareResp {
	return &NullableDfsSMBShareResp{value: val, isSet: true}
}

func (v NullableDfsSMBShareResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsSMBShareResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


