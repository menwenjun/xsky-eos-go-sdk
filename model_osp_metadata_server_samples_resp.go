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

// checks if the OspMetadataServerSamplesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OspMetadataServerSamplesResp{}

// OspMetadataServerSamplesResp struct for OspMetadataServerSamplesResp
type OspMetadataServerSamplesResp struct {
	// osp metadata server samples
	OspMetadataServerSamples []OspMetadataServerStat `json:"osp_metadata_server_samples,omitempty"`
}

// NewOspMetadataServerSamplesResp instantiates a new OspMetadataServerSamplesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOspMetadataServerSamplesResp() *OspMetadataServerSamplesResp {
	this := OspMetadataServerSamplesResp{}
	return &this
}

// NewOspMetadataServerSamplesRespWithDefaults instantiates a new OspMetadataServerSamplesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOspMetadataServerSamplesRespWithDefaults() *OspMetadataServerSamplesResp {
	this := OspMetadataServerSamplesResp{}
	return &this
}

// GetOspMetadataServerSamples returns the OspMetadataServerSamples field value if set, zero value otherwise.
func (o *OspMetadataServerSamplesResp) GetOspMetadataServerSamples() []OspMetadataServerStat {
	if o == nil || IsNil(o.OspMetadataServerSamples) {
		var ret []OspMetadataServerStat
		return ret
	}
	return o.OspMetadataServerSamples
}

// GetOspMetadataServerSamplesOk returns a tuple with the OspMetadataServerSamples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerSamplesResp) GetOspMetadataServerSamplesOk() ([]OspMetadataServerStat, bool) {
	if o == nil || IsNil(o.OspMetadataServerSamples) {
		return nil, false
	}
	return o.OspMetadataServerSamples, true
}

// HasOspMetadataServerSamples returns a boolean if a field has been set.
func (o *OspMetadataServerSamplesResp) HasOspMetadataServerSamples() bool {
	if o != nil && !IsNil(o.OspMetadataServerSamples) {
		return true
	}

	return false
}

// SetOspMetadataServerSamples gets a reference to the given []OspMetadataServerStat and assigns it to the OspMetadataServerSamples field.
func (o *OspMetadataServerSamplesResp) SetOspMetadataServerSamples(v []OspMetadataServerStat) {
	o.OspMetadataServerSamples = v
}

func (o OspMetadataServerSamplesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OspMetadataServerSamplesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OspMetadataServerSamples) {
		toSerialize["osp_metadata_server_samples"] = o.OspMetadataServerSamples
	}
	return toSerialize, nil
}

type NullableOspMetadataServerSamplesResp struct {
	value *OspMetadataServerSamplesResp
	isSet bool
}

func (v NullableOspMetadataServerSamplesResp) Get() *OspMetadataServerSamplesResp {
	return v.value
}

func (v *NullableOspMetadataServerSamplesResp) Set(val *OspMetadataServerSamplesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOspMetadataServerSamplesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOspMetadataServerSamplesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspMetadataServerSamplesResp(val *OspMetadataServerSamplesResp) *NullableOspMetadataServerSamplesResp {
	return &NullableOspMetadataServerSamplesResp{value: val, isSet: true}
}

func (v NullableOspMetadataServerSamplesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspMetadataServerSamplesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


