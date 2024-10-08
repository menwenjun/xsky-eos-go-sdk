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

// checks if the DfsQuotaUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsQuotaUpdateReq{}

// DfsQuotaUpdateReq struct for DfsQuotaUpdateReq
type DfsQuotaUpdateReq struct {
	DfsQuota *DfsQuotaUpdateReqQuota `json:"dfs_quota,omitempty"`
}

// NewDfsQuotaUpdateReq instantiates a new DfsQuotaUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsQuotaUpdateReq() *DfsQuotaUpdateReq {
	this := DfsQuotaUpdateReq{}
	return &this
}

// NewDfsQuotaUpdateReqWithDefaults instantiates a new DfsQuotaUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsQuotaUpdateReqWithDefaults() *DfsQuotaUpdateReq {
	this := DfsQuotaUpdateReq{}
	return &this
}

// GetDfsQuota returns the DfsQuota field value if set, zero value otherwise.
func (o *DfsQuotaUpdateReq) GetDfsQuota() DfsQuotaUpdateReqQuota {
	if o == nil || IsNil(o.DfsQuota) {
		var ret DfsQuotaUpdateReqQuota
		return ret
	}
	return *o.DfsQuota
}

// GetDfsQuotaOk returns a tuple with the DfsQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuotaUpdateReq) GetDfsQuotaOk() (*DfsQuotaUpdateReqQuota, bool) {
	if o == nil || IsNil(o.DfsQuota) {
		return nil, false
	}
	return o.DfsQuota, true
}

// HasDfsQuota returns a boolean if a field has been set.
func (o *DfsQuotaUpdateReq) HasDfsQuota() bool {
	if o != nil && !IsNil(o.DfsQuota) {
		return true
	}

	return false
}

// SetDfsQuota gets a reference to the given DfsQuotaUpdateReqQuota and assigns it to the DfsQuota field.
func (o *DfsQuotaUpdateReq) SetDfsQuota(v DfsQuotaUpdateReqQuota) {
	o.DfsQuota = &v
}

func (o DfsQuotaUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsQuotaUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DfsQuota) {
		toSerialize["dfs_quota"] = o.DfsQuota
	}
	return toSerialize, nil
}

type NullableDfsQuotaUpdateReq struct {
	value *DfsQuotaUpdateReq
	isSet bool
}

func (v NullableDfsQuotaUpdateReq) Get() *DfsQuotaUpdateReq {
	return v.value
}

func (v *NullableDfsQuotaUpdateReq) Set(val *DfsQuotaUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsQuotaUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsQuotaUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsQuotaUpdateReq(val *DfsQuotaUpdateReq) *NullableDfsQuotaUpdateReq {
	return &NullableDfsQuotaUpdateReq{value: val, isSet: true}
}

func (v NullableDfsQuotaUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsQuotaUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


