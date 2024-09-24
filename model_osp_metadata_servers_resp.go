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

// checks if the OspMetadataServersResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OspMetadataServersResp{}

// OspMetadataServersResp struct for OspMetadataServersResp
type OspMetadataServersResp struct {
	OspMetadataServers []OspMetadataServerRecord `json:"osp_metadata_servers,omitempty"`
}

// NewOspMetadataServersResp instantiates a new OspMetadataServersResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOspMetadataServersResp() *OspMetadataServersResp {
	this := OspMetadataServersResp{}
	return &this
}

// NewOspMetadataServersRespWithDefaults instantiates a new OspMetadataServersResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOspMetadataServersRespWithDefaults() *OspMetadataServersResp {
	this := OspMetadataServersResp{}
	return &this
}

// GetOspMetadataServers returns the OspMetadataServers field value if set, zero value otherwise.
func (o *OspMetadataServersResp) GetOspMetadataServers() []OspMetadataServerRecord {
	if o == nil || IsNil(o.OspMetadataServers) {
		var ret []OspMetadataServerRecord
		return ret
	}
	return o.OspMetadataServers
}

// GetOspMetadataServersOk returns a tuple with the OspMetadataServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServersResp) GetOspMetadataServersOk() ([]OspMetadataServerRecord, bool) {
	if o == nil || IsNil(o.OspMetadataServers) {
		return nil, false
	}
	return o.OspMetadataServers, true
}

// HasOspMetadataServers returns a boolean if a field has been set.
func (o *OspMetadataServersResp) HasOspMetadataServers() bool {
	if o != nil && !IsNil(o.OspMetadataServers) {
		return true
	}

	return false
}

// SetOspMetadataServers gets a reference to the given []OspMetadataServerRecord and assigns it to the OspMetadataServers field.
func (o *OspMetadataServersResp) SetOspMetadataServers(v []OspMetadataServerRecord) {
	o.OspMetadataServers = v
}

func (o OspMetadataServersResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OspMetadataServersResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OspMetadataServers) {
		toSerialize["osp_metadata_servers"] = o.OspMetadataServers
	}
	return toSerialize, nil
}

type NullableOspMetadataServersResp struct {
	value *OspMetadataServersResp
	isSet bool
}

func (v NullableOspMetadataServersResp) Get() *OspMetadataServersResp {
	return v.value
}

func (v *NullableOspMetadataServersResp) Set(val *OspMetadataServersResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOspMetadataServersResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOspMetadataServersResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspMetadataServersResp(val *OspMetadataServersResp) *NullableOspMetadataServersResp {
	return &NullableOspMetadataServersResp{value: val, isSet: true}
}

func (v NullableOspMetadataServersResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspMetadataServersResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


