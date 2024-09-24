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

// checks if the DfsTierResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsTierResp{}

// DfsTierResp struct for DfsTierResp
type DfsTierResp struct {
	DfsTier DfsTier `json:"dfs_tier"`
}

type _DfsTierResp DfsTierResp

// NewDfsTierResp instantiates a new DfsTierResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsTierResp(dfsTier DfsTier) *DfsTierResp {
	this := DfsTierResp{}
	this.DfsTier = dfsTier
	return &this
}

// NewDfsTierRespWithDefaults instantiates a new DfsTierResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsTierRespWithDefaults() *DfsTierResp {
	this := DfsTierResp{}
	return &this
}

// GetDfsTier returns the DfsTier field value
func (o *DfsTierResp) GetDfsTier() DfsTier {
	if o == nil {
		var ret DfsTier
		return ret
	}

	return o.DfsTier
}

// GetDfsTierOk returns a tuple with the DfsTier field value
// and a boolean to check if the value has been set.
func (o *DfsTierResp) GetDfsTierOk() (*DfsTier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsTier, true
}

// SetDfsTier sets field value
func (o *DfsTierResp) SetDfsTier(v DfsTier) {
	o.DfsTier = v
}

func (o DfsTierResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsTierResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_tier"] = o.DfsTier
	return toSerialize, nil
}

func (o *DfsTierResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_tier",
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

	varDfsTierResp := _DfsTierResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsTierResp)

	if err != nil {
		return err
	}

	*o = DfsTierResp(varDfsTierResp)

	return err
}

type NullableDfsTierResp struct {
	value *DfsTierResp
	isSet bool
}

func (v NullableDfsTierResp) Get() *DfsTierResp {
	return v.value
}

func (v *NullableDfsTierResp) Set(val *DfsTierResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsTierResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsTierResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsTierResp(val *DfsTierResp) *NullableDfsTierResp {
	return &NullableDfsTierResp{value: val, isSet: true}
}

func (v NullableDfsTierResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsTierResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


