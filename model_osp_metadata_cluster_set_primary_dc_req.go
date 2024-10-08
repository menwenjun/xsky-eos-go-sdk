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

// checks if the OspMetadataClusterSetPrimaryDcReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OspMetadataClusterSetPrimaryDcReq{}

// OspMetadataClusterSetPrimaryDcReq struct for OspMetadataClusterSetPrimaryDcReq
type OspMetadataClusterSetPrimaryDcReq struct {
	OspMetadataCluster *OspMetadataClusterSetPrimaryDcReqOspMetadataCluster `json:"osp_metadata_cluster,omitempty"`
}

// NewOspMetadataClusterSetPrimaryDcReq instantiates a new OspMetadataClusterSetPrimaryDcReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOspMetadataClusterSetPrimaryDcReq() *OspMetadataClusterSetPrimaryDcReq {
	this := OspMetadataClusterSetPrimaryDcReq{}
	return &this
}

// NewOspMetadataClusterSetPrimaryDcReqWithDefaults instantiates a new OspMetadataClusterSetPrimaryDcReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOspMetadataClusterSetPrimaryDcReqWithDefaults() *OspMetadataClusterSetPrimaryDcReq {
	this := OspMetadataClusterSetPrimaryDcReq{}
	return &this
}

// GetOspMetadataCluster returns the OspMetadataCluster field value if set, zero value otherwise.
func (o *OspMetadataClusterSetPrimaryDcReq) GetOspMetadataCluster() OspMetadataClusterSetPrimaryDcReqOspMetadataCluster {
	if o == nil || IsNil(o.OspMetadataCluster) {
		var ret OspMetadataClusterSetPrimaryDcReqOspMetadataCluster
		return ret
	}
	return *o.OspMetadataCluster
}

// GetOspMetadataClusterOk returns a tuple with the OspMetadataCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterSetPrimaryDcReq) GetOspMetadataClusterOk() (*OspMetadataClusterSetPrimaryDcReqOspMetadataCluster, bool) {
	if o == nil || IsNil(o.OspMetadataCluster) {
		return nil, false
	}
	return o.OspMetadataCluster, true
}

// HasOspMetadataCluster returns a boolean if a field has been set.
func (o *OspMetadataClusterSetPrimaryDcReq) HasOspMetadataCluster() bool {
	if o != nil && !IsNil(o.OspMetadataCluster) {
		return true
	}

	return false
}

// SetOspMetadataCluster gets a reference to the given OspMetadataClusterSetPrimaryDcReqOspMetadataCluster and assigns it to the OspMetadataCluster field.
func (o *OspMetadataClusterSetPrimaryDcReq) SetOspMetadataCluster(v OspMetadataClusterSetPrimaryDcReqOspMetadataCluster) {
	o.OspMetadataCluster = &v
}

func (o OspMetadataClusterSetPrimaryDcReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OspMetadataClusterSetPrimaryDcReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OspMetadataCluster) {
		toSerialize["osp_metadata_cluster"] = o.OspMetadataCluster
	}
	return toSerialize, nil
}

type NullableOspMetadataClusterSetPrimaryDcReq struct {
	value *OspMetadataClusterSetPrimaryDcReq
	isSet bool
}

func (v NullableOspMetadataClusterSetPrimaryDcReq) Get() *OspMetadataClusterSetPrimaryDcReq {
	return v.value
}

func (v *NullableOspMetadataClusterSetPrimaryDcReq) Set(val *OspMetadataClusterSetPrimaryDcReq) {
	v.value = val
	v.isSet = true
}

func (v NullableOspMetadataClusterSetPrimaryDcReq) IsSet() bool {
	return v.isSet
}

func (v *NullableOspMetadataClusterSetPrimaryDcReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspMetadataClusterSetPrimaryDcReq(val *OspMetadataClusterSetPrimaryDcReq) *NullableOspMetadataClusterSetPrimaryDcReq {
	return &NullableOspMetadataClusterSetPrimaryDcReq{value: val, isSet: true}
}

func (v NullableOspMetadataClusterSetPrimaryDcReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspMetadataClusterSetPrimaryDcReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


