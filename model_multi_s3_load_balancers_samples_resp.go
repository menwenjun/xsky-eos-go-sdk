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

// checks if the MultiS3LoadBalancersSamplesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiS3LoadBalancersSamplesResp{}

// MultiS3LoadBalancersSamplesResp struct for MultiS3LoadBalancersSamplesResp
type MultiS3LoadBalancersSamplesResp struct {
	S3LoadBalancersSamples []S3LoadBalancerSamplesElem `json:"s3_load_balancers_samples,omitempty"`
}

// NewMultiS3LoadBalancersSamplesResp instantiates a new MultiS3LoadBalancersSamplesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiS3LoadBalancersSamplesResp() *MultiS3LoadBalancersSamplesResp {
	this := MultiS3LoadBalancersSamplesResp{}
	return &this
}

// NewMultiS3LoadBalancersSamplesRespWithDefaults instantiates a new MultiS3LoadBalancersSamplesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiS3LoadBalancersSamplesRespWithDefaults() *MultiS3LoadBalancersSamplesResp {
	this := MultiS3LoadBalancersSamplesResp{}
	return &this
}

// GetS3LoadBalancersSamples returns the S3LoadBalancersSamples field value if set, zero value otherwise.
func (o *MultiS3LoadBalancersSamplesResp) GetS3LoadBalancersSamples() []S3LoadBalancerSamplesElem {
	if o == nil || IsNil(o.S3LoadBalancersSamples) {
		var ret []S3LoadBalancerSamplesElem
		return ret
	}
	return o.S3LoadBalancersSamples
}

// GetS3LoadBalancersSamplesOk returns a tuple with the S3LoadBalancersSamples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiS3LoadBalancersSamplesResp) GetS3LoadBalancersSamplesOk() ([]S3LoadBalancerSamplesElem, bool) {
	if o == nil || IsNil(o.S3LoadBalancersSamples) {
		return nil, false
	}
	return o.S3LoadBalancersSamples, true
}

// HasS3LoadBalancersSamples returns a boolean if a field has been set.
func (o *MultiS3LoadBalancersSamplesResp) HasS3LoadBalancersSamples() bool {
	if o != nil && !IsNil(o.S3LoadBalancersSamples) {
		return true
	}

	return false
}

// SetS3LoadBalancersSamples gets a reference to the given []S3LoadBalancerSamplesElem and assigns it to the S3LoadBalancersSamples field.
func (o *MultiS3LoadBalancersSamplesResp) SetS3LoadBalancersSamples(v []S3LoadBalancerSamplesElem) {
	o.S3LoadBalancersSamples = v
}

func (o MultiS3LoadBalancersSamplesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiS3LoadBalancersSamplesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.S3LoadBalancersSamples) {
		toSerialize["s3_load_balancers_samples"] = o.S3LoadBalancersSamples
	}
	return toSerialize, nil
}

type NullableMultiS3LoadBalancersSamplesResp struct {
	value *MultiS3LoadBalancersSamplesResp
	isSet bool
}

func (v NullableMultiS3LoadBalancersSamplesResp) Get() *MultiS3LoadBalancersSamplesResp {
	return v.value
}

func (v *NullableMultiS3LoadBalancersSamplesResp) Set(val *MultiS3LoadBalancersSamplesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiS3LoadBalancersSamplesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiS3LoadBalancersSamplesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiS3LoadBalancersSamplesResp(val *MultiS3LoadBalancersSamplesResp) *NullableMultiS3LoadBalancersSamplesResp {
	return &NullableMultiS3LoadBalancersSamplesResp{value: val, isSet: true}
}

func (v NullableMultiS3LoadBalancersSamplesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiS3LoadBalancersSamplesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


