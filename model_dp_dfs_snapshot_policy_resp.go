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

// checks if the DpDfsSnapshotPolicyResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpDfsSnapshotPolicyResp{}

// DpDfsSnapshotPolicyResp struct for DpDfsSnapshotPolicyResp
type DpDfsSnapshotPolicyResp struct {
	DpDfsSnapshotPolicy DpDfsSnapshotPolicy `json:"dp_dfs_snapshot_policy"`
}

type _DpDfsSnapshotPolicyResp DpDfsSnapshotPolicyResp

// NewDpDfsSnapshotPolicyResp instantiates a new DpDfsSnapshotPolicyResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpDfsSnapshotPolicyResp(dpDfsSnapshotPolicy DpDfsSnapshotPolicy) *DpDfsSnapshotPolicyResp {
	this := DpDfsSnapshotPolicyResp{}
	this.DpDfsSnapshotPolicy = dpDfsSnapshotPolicy
	return &this
}

// NewDpDfsSnapshotPolicyRespWithDefaults instantiates a new DpDfsSnapshotPolicyResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpDfsSnapshotPolicyRespWithDefaults() *DpDfsSnapshotPolicyResp {
	this := DpDfsSnapshotPolicyResp{}
	return &this
}

// GetDpDfsSnapshotPolicy returns the DpDfsSnapshotPolicy field value
func (o *DpDfsSnapshotPolicyResp) GetDpDfsSnapshotPolicy() DpDfsSnapshotPolicy {
	if o == nil {
		var ret DpDfsSnapshotPolicy
		return ret
	}

	return o.DpDfsSnapshotPolicy
}

// GetDpDfsSnapshotPolicyOk returns a tuple with the DpDfsSnapshotPolicy field value
// and a boolean to check if the value has been set.
func (o *DpDfsSnapshotPolicyResp) GetDpDfsSnapshotPolicyOk() (*DpDfsSnapshotPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DpDfsSnapshotPolicy, true
}

// SetDpDfsSnapshotPolicy sets field value
func (o *DpDfsSnapshotPolicyResp) SetDpDfsSnapshotPolicy(v DpDfsSnapshotPolicy) {
	o.DpDfsSnapshotPolicy = v
}

func (o DpDfsSnapshotPolicyResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpDfsSnapshotPolicyResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dp_dfs_snapshot_policy"] = o.DpDfsSnapshotPolicy
	return toSerialize, nil
}

func (o *DpDfsSnapshotPolicyResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dp_dfs_snapshot_policy",
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

	varDpDfsSnapshotPolicyResp := _DpDfsSnapshotPolicyResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDpDfsSnapshotPolicyResp)

	if err != nil {
		return err
	}

	*o = DpDfsSnapshotPolicyResp(varDpDfsSnapshotPolicyResp)

	return err
}

type NullableDpDfsSnapshotPolicyResp struct {
	value *DpDfsSnapshotPolicyResp
	isSet bool
}

func (v NullableDpDfsSnapshotPolicyResp) Get() *DpDfsSnapshotPolicyResp {
	return v.value
}

func (v *NullableDpDfsSnapshotPolicyResp) Set(val *DpDfsSnapshotPolicyResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDpDfsSnapshotPolicyResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDpDfsSnapshotPolicyResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpDfsSnapshotPolicyResp(val *DpDfsSnapshotPolicyResp) *NullableDpDfsSnapshotPolicyResp {
	return &NullableDpDfsSnapshotPolicyResp{value: val, isSet: true}
}

func (v NullableDpDfsSnapshotPolicyResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpDfsSnapshotPolicyResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


