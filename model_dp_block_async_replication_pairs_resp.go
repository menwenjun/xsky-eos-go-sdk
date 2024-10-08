/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DpBlockAsyncReplicationPairsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpBlockAsyncReplicationPairsResp{}

// DpBlockAsyncReplicationPairsResp struct for DpBlockAsyncReplicationPairsResp
type DpBlockAsyncReplicationPairsResp struct {
	DpBlockAsyncReplicationPairs []DpBlockAsyncReplicationPair `json:"dp_block_async_replication_pairs,omitempty"`
}

// NewDpBlockAsyncReplicationPairsResp instantiates a new DpBlockAsyncReplicationPairsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpBlockAsyncReplicationPairsResp() *DpBlockAsyncReplicationPairsResp {
	this := DpBlockAsyncReplicationPairsResp{}
	return &this
}

// NewDpBlockAsyncReplicationPairsRespWithDefaults instantiates a new DpBlockAsyncReplicationPairsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpBlockAsyncReplicationPairsRespWithDefaults() *DpBlockAsyncReplicationPairsResp {
	this := DpBlockAsyncReplicationPairsResp{}
	return &this
}

// GetDpBlockAsyncReplicationPairs returns the DpBlockAsyncReplicationPairs field value if set, zero value otherwise.
func (o *DpBlockAsyncReplicationPairsResp) GetDpBlockAsyncReplicationPairs() []DpBlockAsyncReplicationPair {
	if o == nil || IsNil(o.DpBlockAsyncReplicationPairs) {
		var ret []DpBlockAsyncReplicationPair
		return ret
	}
	return o.DpBlockAsyncReplicationPairs
}

// GetDpBlockAsyncReplicationPairsOk returns a tuple with the DpBlockAsyncReplicationPairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockAsyncReplicationPairsResp) GetDpBlockAsyncReplicationPairsOk() ([]DpBlockAsyncReplicationPair, bool) {
	if o == nil || IsNil(o.DpBlockAsyncReplicationPairs) {
		return nil, false
	}
	return o.DpBlockAsyncReplicationPairs, true
}

// HasDpBlockAsyncReplicationPairs returns a boolean if a field has been set.
func (o *DpBlockAsyncReplicationPairsResp) HasDpBlockAsyncReplicationPairs() bool {
	if o != nil && !IsNil(o.DpBlockAsyncReplicationPairs) {
		return true
	}

	return false
}

// SetDpBlockAsyncReplicationPairs gets a reference to the given []DpBlockAsyncReplicationPair and assigns it to the DpBlockAsyncReplicationPairs field.
func (o *DpBlockAsyncReplicationPairsResp) SetDpBlockAsyncReplicationPairs(v []DpBlockAsyncReplicationPair) {
	o.DpBlockAsyncReplicationPairs = v
}

func (o DpBlockAsyncReplicationPairsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpBlockAsyncReplicationPairsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DpBlockAsyncReplicationPairs) {
		toSerialize["dp_block_async_replication_pairs"] = o.DpBlockAsyncReplicationPairs
	}
	return toSerialize, nil
}

type NullableDpBlockAsyncReplicationPairsResp struct {
	value *DpBlockAsyncReplicationPairsResp
	isSet bool
}

func (v NullableDpBlockAsyncReplicationPairsResp) Get() *DpBlockAsyncReplicationPairsResp {
	return v.value
}

func (v *NullableDpBlockAsyncReplicationPairsResp) Set(val *DpBlockAsyncReplicationPairsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDpBlockAsyncReplicationPairsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDpBlockAsyncReplicationPairsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpBlockAsyncReplicationPairsResp(val *DpBlockAsyncReplicationPairsResp) *NullableDpBlockAsyncReplicationPairsResp {
	return &NullableDpBlockAsyncReplicationPairsResp{value: val, isSet: true}
}

func (v NullableDpBlockAsyncReplicationPairsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpBlockAsyncReplicationPairsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


