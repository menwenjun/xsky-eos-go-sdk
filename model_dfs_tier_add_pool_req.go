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

// checks if the DfsTierAddPoolReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsTierAddPoolReq{}

// DfsTierAddPoolReq struct for DfsTierAddPoolReq
type DfsTierAddPoolReq struct {
	DfsTier DfsTierAddPoolReqTier `json:"dfs_tier"`
}

type _DfsTierAddPoolReq DfsTierAddPoolReq

// NewDfsTierAddPoolReq instantiates a new DfsTierAddPoolReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsTierAddPoolReq(dfsTier DfsTierAddPoolReqTier) *DfsTierAddPoolReq {
	this := DfsTierAddPoolReq{}
	this.DfsTier = dfsTier
	return &this
}

// NewDfsTierAddPoolReqWithDefaults instantiates a new DfsTierAddPoolReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsTierAddPoolReqWithDefaults() *DfsTierAddPoolReq {
	this := DfsTierAddPoolReq{}
	return &this
}

// GetDfsTier returns the DfsTier field value
func (o *DfsTierAddPoolReq) GetDfsTier() DfsTierAddPoolReqTier {
	if o == nil {
		var ret DfsTierAddPoolReqTier
		return ret
	}

	return o.DfsTier
}

// GetDfsTierOk returns a tuple with the DfsTier field value
// and a boolean to check if the value has been set.
func (o *DfsTierAddPoolReq) GetDfsTierOk() (*DfsTierAddPoolReqTier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsTier, true
}

// SetDfsTier sets field value
func (o *DfsTierAddPoolReq) SetDfsTier(v DfsTierAddPoolReqTier) {
	o.DfsTier = v
}

func (o DfsTierAddPoolReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsTierAddPoolReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_tier"] = o.DfsTier
	return toSerialize, nil
}

func (o *DfsTierAddPoolReq) UnmarshalJSON(data []byte) (err error) {
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

	varDfsTierAddPoolReq := _DfsTierAddPoolReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsTierAddPoolReq)

	if err != nil {
		return err
	}

	*o = DfsTierAddPoolReq(varDfsTierAddPoolReq)

	return err
}

type NullableDfsTierAddPoolReq struct {
	value *DfsTierAddPoolReq
	isSet bool
}

func (v NullableDfsTierAddPoolReq) Get() *DfsTierAddPoolReq {
	return v.value
}

func (v *NullableDfsTierAddPoolReq) Set(val *DfsTierAddPoolReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsTierAddPoolReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsTierAddPoolReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsTierAddPoolReq(val *DfsTierAddPoolReq) *NullableDfsTierAddPoolReq {
	return &NullableDfsTierAddPoolReq{value: val, isSet: true}
}

func (v NullableDfsTierAddPoolReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsTierAddPoolReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


