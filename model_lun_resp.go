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

// checks if the LunResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LunResp{}

// LunResp struct for LunResp
type LunResp struct {
	Lun Lun `json:"lun"`
}

type _LunResp LunResp

// NewLunResp instantiates a new LunResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLunResp(lun Lun) *LunResp {
	this := LunResp{}
	this.Lun = lun
	return &this
}

// NewLunRespWithDefaults instantiates a new LunResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLunRespWithDefaults() *LunResp {
	this := LunResp{}
	return &this
}

// GetLun returns the Lun field value
func (o *LunResp) GetLun() Lun {
	if o == nil {
		var ret Lun
		return ret
	}

	return o.Lun
}

// GetLunOk returns a tuple with the Lun field value
// and a boolean to check if the value has been set.
func (o *LunResp) GetLunOk() (*Lun, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lun, true
}

// SetLun sets field value
func (o *LunResp) SetLun(v Lun) {
	o.Lun = v
}

func (o LunResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LunResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lun"] = o.Lun
	return toSerialize, nil
}

func (o *LunResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lun",
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

	varLunResp := _LunResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLunResp)

	if err != nil {
		return err
	}

	*o = LunResp(varLunResp)

	return err
}

type NullableLunResp struct {
	value *LunResp
	isSet bool
}

func (v NullableLunResp) Get() *LunResp {
	return v.value
}

func (v *NullableLunResp) Set(val *LunResp) {
	v.value = val
	v.isSet = true
}

func (v NullableLunResp) IsSet() bool {
	return v.isSet
}

func (v *NullableLunResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLunResp(val *LunResp) *NullableLunResp {
	return &NullableLunResp{value: val, isSet: true}
}

func (v NullableLunResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLunResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


