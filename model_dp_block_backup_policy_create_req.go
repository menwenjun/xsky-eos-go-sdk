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

// checks if the DpBlockBackupPolicyCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpBlockBackupPolicyCreateReq{}

// DpBlockBackupPolicyCreateReq struct for DpBlockBackupPolicyCreateReq
type DpBlockBackupPolicyCreateReq struct {
	DpBlockBackupPolicy DpBlockBackupPolicyCreateReqPolicy `json:"dp_block_backup_policy"`
}

type _DpBlockBackupPolicyCreateReq DpBlockBackupPolicyCreateReq

// NewDpBlockBackupPolicyCreateReq instantiates a new DpBlockBackupPolicyCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpBlockBackupPolicyCreateReq(dpBlockBackupPolicy DpBlockBackupPolicyCreateReqPolicy) *DpBlockBackupPolicyCreateReq {
	this := DpBlockBackupPolicyCreateReq{}
	this.DpBlockBackupPolicy = dpBlockBackupPolicy
	return &this
}

// NewDpBlockBackupPolicyCreateReqWithDefaults instantiates a new DpBlockBackupPolicyCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpBlockBackupPolicyCreateReqWithDefaults() *DpBlockBackupPolicyCreateReq {
	this := DpBlockBackupPolicyCreateReq{}
	return &this
}

// GetDpBlockBackupPolicy returns the DpBlockBackupPolicy field value
func (o *DpBlockBackupPolicyCreateReq) GetDpBlockBackupPolicy() DpBlockBackupPolicyCreateReqPolicy {
	if o == nil {
		var ret DpBlockBackupPolicyCreateReqPolicy
		return ret
	}

	return o.DpBlockBackupPolicy
}

// GetDpBlockBackupPolicyOk returns a tuple with the DpBlockBackupPolicy field value
// and a boolean to check if the value has been set.
func (o *DpBlockBackupPolicyCreateReq) GetDpBlockBackupPolicyOk() (*DpBlockBackupPolicyCreateReqPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DpBlockBackupPolicy, true
}

// SetDpBlockBackupPolicy sets field value
func (o *DpBlockBackupPolicyCreateReq) SetDpBlockBackupPolicy(v DpBlockBackupPolicyCreateReqPolicy) {
	o.DpBlockBackupPolicy = v
}

func (o DpBlockBackupPolicyCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpBlockBackupPolicyCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dp_block_backup_policy"] = o.DpBlockBackupPolicy
	return toSerialize, nil
}

func (o *DpBlockBackupPolicyCreateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dp_block_backup_policy",
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

	varDpBlockBackupPolicyCreateReq := _DpBlockBackupPolicyCreateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDpBlockBackupPolicyCreateReq)

	if err != nil {
		return err
	}

	*o = DpBlockBackupPolicyCreateReq(varDpBlockBackupPolicyCreateReq)

	return err
}

type NullableDpBlockBackupPolicyCreateReq struct {
	value *DpBlockBackupPolicyCreateReq
	isSet bool
}

func (v NullableDpBlockBackupPolicyCreateReq) Get() *DpBlockBackupPolicyCreateReq {
	return v.value
}

func (v *NullableDpBlockBackupPolicyCreateReq) Set(val *DpBlockBackupPolicyCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDpBlockBackupPolicyCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDpBlockBackupPolicyCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpBlockBackupPolicyCreateReq(val *DpBlockBackupPolicyCreateReq) *NullableDpBlockBackupPolicyCreateReq {
	return &NullableDpBlockBackupPolicyCreateReq{value: val, isSet: true}
}

func (v NullableDpBlockBackupPolicyCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpBlockBackupPolicyCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


