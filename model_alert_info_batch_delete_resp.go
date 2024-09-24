/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the AlertInfoBatchDeleteResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertInfoBatchDeleteResp{}

// AlertInfoBatchDeleteResp struct for AlertInfoBatchDeleteResp
type AlertInfoBatchDeleteResp struct {
	CreateAfter *time.Time `json:"create_after,omitempty"`
	CreateBefore *time.Time `json:"create_before,omitempty"`
	Total *int64 `json:"total,omitempty"`
	UnresolvedCount *int64 `json:"unresolved_count,omitempty"`
}

// NewAlertInfoBatchDeleteResp instantiates a new AlertInfoBatchDeleteResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertInfoBatchDeleteResp() *AlertInfoBatchDeleteResp {
	this := AlertInfoBatchDeleteResp{}
	return &this
}

// NewAlertInfoBatchDeleteRespWithDefaults instantiates a new AlertInfoBatchDeleteResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertInfoBatchDeleteRespWithDefaults() *AlertInfoBatchDeleteResp {
	this := AlertInfoBatchDeleteResp{}
	return &this
}

// GetCreateAfter returns the CreateAfter field value if set, zero value otherwise.
func (o *AlertInfoBatchDeleteResp) GetCreateAfter() time.Time {
	if o == nil || IsNil(o.CreateAfter) {
		var ret time.Time
		return ret
	}
	return *o.CreateAfter
}

// GetCreateAfterOk returns a tuple with the CreateAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertInfoBatchDeleteResp) GetCreateAfterOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateAfter) {
		return nil, false
	}
	return o.CreateAfter, true
}

// HasCreateAfter returns a boolean if a field has been set.
func (o *AlertInfoBatchDeleteResp) HasCreateAfter() bool {
	if o != nil && !IsNil(o.CreateAfter) {
		return true
	}

	return false
}

// SetCreateAfter gets a reference to the given time.Time and assigns it to the CreateAfter field.
func (o *AlertInfoBatchDeleteResp) SetCreateAfter(v time.Time) {
	o.CreateAfter = &v
}

// GetCreateBefore returns the CreateBefore field value if set, zero value otherwise.
func (o *AlertInfoBatchDeleteResp) GetCreateBefore() time.Time {
	if o == nil || IsNil(o.CreateBefore) {
		var ret time.Time
		return ret
	}
	return *o.CreateBefore
}

// GetCreateBeforeOk returns a tuple with the CreateBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertInfoBatchDeleteResp) GetCreateBeforeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateBefore) {
		return nil, false
	}
	return o.CreateBefore, true
}

// HasCreateBefore returns a boolean if a field has been set.
func (o *AlertInfoBatchDeleteResp) HasCreateBefore() bool {
	if o != nil && !IsNil(o.CreateBefore) {
		return true
	}

	return false
}

// SetCreateBefore gets a reference to the given time.Time and assigns it to the CreateBefore field.
func (o *AlertInfoBatchDeleteResp) SetCreateBefore(v time.Time) {
	o.CreateBefore = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *AlertInfoBatchDeleteResp) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertInfoBatchDeleteResp) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *AlertInfoBatchDeleteResp) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *AlertInfoBatchDeleteResp) SetTotal(v int64) {
	o.Total = &v
}

// GetUnresolvedCount returns the UnresolvedCount field value if set, zero value otherwise.
func (o *AlertInfoBatchDeleteResp) GetUnresolvedCount() int64 {
	if o == nil || IsNil(o.UnresolvedCount) {
		var ret int64
		return ret
	}
	return *o.UnresolvedCount
}

// GetUnresolvedCountOk returns a tuple with the UnresolvedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertInfoBatchDeleteResp) GetUnresolvedCountOk() (*int64, bool) {
	if o == nil || IsNil(o.UnresolvedCount) {
		return nil, false
	}
	return o.UnresolvedCount, true
}

// HasUnresolvedCount returns a boolean if a field has been set.
func (o *AlertInfoBatchDeleteResp) HasUnresolvedCount() bool {
	if o != nil && !IsNil(o.UnresolvedCount) {
		return true
	}

	return false
}

// SetUnresolvedCount gets a reference to the given int64 and assigns it to the UnresolvedCount field.
func (o *AlertInfoBatchDeleteResp) SetUnresolvedCount(v int64) {
	o.UnresolvedCount = &v
}

func (o AlertInfoBatchDeleteResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertInfoBatchDeleteResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateAfter) {
		toSerialize["create_after"] = o.CreateAfter
	}
	if !IsNil(o.CreateBefore) {
		toSerialize["create_before"] = o.CreateBefore
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.UnresolvedCount) {
		toSerialize["unresolved_count"] = o.UnresolvedCount
	}
	return toSerialize, nil
}

type NullableAlertInfoBatchDeleteResp struct {
	value *AlertInfoBatchDeleteResp
	isSet bool
}

func (v NullableAlertInfoBatchDeleteResp) Get() *AlertInfoBatchDeleteResp {
	return v.value
}

func (v *NullableAlertInfoBatchDeleteResp) Set(val *AlertInfoBatchDeleteResp) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertInfoBatchDeleteResp) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertInfoBatchDeleteResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertInfoBatchDeleteResp(val *AlertInfoBatchDeleteResp) *NullableAlertInfoBatchDeleteResp {
	return &NullableAlertInfoBatchDeleteResp{value: val, isSet: true}
}

func (v NullableAlertInfoBatchDeleteResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertInfoBatchDeleteResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


