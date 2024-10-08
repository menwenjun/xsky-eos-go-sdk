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

// checks if the DfsTiersResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsTiersResp{}

// DfsTiersResp struct for DfsTiersResp
type DfsTiersResp struct {
	// dfs tiers
	DfsTiers []DfsTier `json:"dfs_tiers"`
}

type _DfsTiersResp DfsTiersResp

// NewDfsTiersResp instantiates a new DfsTiersResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsTiersResp(dfsTiers []DfsTier) *DfsTiersResp {
	this := DfsTiersResp{}
	this.DfsTiers = dfsTiers
	return &this
}

// NewDfsTiersRespWithDefaults instantiates a new DfsTiersResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsTiersRespWithDefaults() *DfsTiersResp {
	this := DfsTiersResp{}
	return &this
}

// GetDfsTiers returns the DfsTiers field value
func (o *DfsTiersResp) GetDfsTiers() []DfsTier {
	if o == nil {
		var ret []DfsTier
		return ret
	}

	return o.DfsTiers
}

// GetDfsTiersOk returns a tuple with the DfsTiers field value
// and a boolean to check if the value has been set.
func (o *DfsTiersResp) GetDfsTiersOk() ([]DfsTier, bool) {
	if o == nil {
		return nil, false
	}
	return o.DfsTiers, true
}

// SetDfsTiers sets field value
func (o *DfsTiersResp) SetDfsTiers(v []DfsTier) {
	o.DfsTiers = v
}

func (o DfsTiersResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsTiersResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_tiers"] = o.DfsTiers
	return toSerialize, nil
}

func (o *DfsTiersResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_tiers",
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

	varDfsTiersResp := _DfsTiersResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsTiersResp)

	if err != nil {
		return err
	}

	*o = DfsTiersResp(varDfsTiersResp)

	return err
}

type NullableDfsTiersResp struct {
	value *DfsTiersResp
	isSet bool
}

func (v NullableDfsTiersResp) Get() *DfsTiersResp {
	return v.value
}

func (v *NullableDfsTiersResp) Set(val *DfsTiersResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsTiersResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsTiersResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsTiersResp(val *DfsTiersResp) *NullableDfsTiersResp {
	return &NullableDfsTiersResp{value: val, isSet: true}
}

func (v NullableDfsTiersResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsTiersResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


