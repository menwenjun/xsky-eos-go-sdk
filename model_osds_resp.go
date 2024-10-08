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

// checks if the OsdsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsdsResp{}

// OsdsResp struct for OsdsResp
type OsdsResp struct {
	// osds
	Osds []OsdRecord `json:"osds"`
}

type _OsdsResp OsdsResp

// NewOsdsResp instantiates a new OsdsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsdsResp(osds []OsdRecord) *OsdsResp {
	this := OsdsResp{}
	this.Osds = osds
	return &this
}

// NewOsdsRespWithDefaults instantiates a new OsdsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsdsRespWithDefaults() *OsdsResp {
	this := OsdsResp{}
	return &this
}

// GetOsds returns the Osds field value
func (o *OsdsResp) GetOsds() []OsdRecord {
	if o == nil {
		var ret []OsdRecord
		return ret
	}

	return o.Osds
}

// GetOsdsOk returns a tuple with the Osds field value
// and a boolean to check if the value has been set.
func (o *OsdsResp) GetOsdsOk() ([]OsdRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.Osds, true
}

// SetOsds sets field value
func (o *OsdsResp) SetOsds(v []OsdRecord) {
	o.Osds = v
}

func (o OsdsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsdsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["osds"] = o.Osds
	return toSerialize, nil
}

func (o *OsdsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"osds",
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

	varOsdsResp := _OsdsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOsdsResp)

	if err != nil {
		return err
	}

	*o = OsdsResp(varOsdsResp)

	return err
}

type NullableOsdsResp struct {
	value *OsdsResp
	isSet bool
}

func (v NullableOsdsResp) Get() *OsdsResp {
	return v.value
}

func (v *NullableOsdsResp) Set(val *OsdsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOsdsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOsdsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsdsResp(val *OsdsResp) *NullableOsdsResp {
	return &NullableOsdsResp{value: val, isSet: true}
}

func (v NullableOsdsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsdsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


