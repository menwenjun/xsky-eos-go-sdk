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

// checks if the MetadataClusterSetTransLocaterReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataClusterSetTransLocaterReq{}

// MetadataClusterSetTransLocaterReq struct for MetadataClusterSetTransLocaterReq
type MetadataClusterSetTransLocaterReq struct {
	MetadataCluster *MetadataClusterSetTransLocaterReqMetadataCluster `json:"metadata_cluster,omitempty"`
}

// NewMetadataClusterSetTransLocaterReq instantiates a new MetadataClusterSetTransLocaterReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataClusterSetTransLocaterReq() *MetadataClusterSetTransLocaterReq {
	this := MetadataClusterSetTransLocaterReq{}
	return &this
}

// NewMetadataClusterSetTransLocaterReqWithDefaults instantiates a new MetadataClusterSetTransLocaterReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataClusterSetTransLocaterReqWithDefaults() *MetadataClusterSetTransLocaterReq {
	this := MetadataClusterSetTransLocaterReq{}
	return &this
}

// GetMetadataCluster returns the MetadataCluster field value if set, zero value otherwise.
func (o *MetadataClusterSetTransLocaterReq) GetMetadataCluster() MetadataClusterSetTransLocaterReqMetadataCluster {
	if o == nil || IsNil(o.MetadataCluster) {
		var ret MetadataClusterSetTransLocaterReqMetadataCluster
		return ret
	}
	return *o.MetadataCluster
}

// GetMetadataClusterOk returns a tuple with the MetadataCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataClusterSetTransLocaterReq) GetMetadataClusterOk() (*MetadataClusterSetTransLocaterReqMetadataCluster, bool) {
	if o == nil || IsNil(o.MetadataCluster) {
		return nil, false
	}
	return o.MetadataCluster, true
}

// HasMetadataCluster returns a boolean if a field has been set.
func (o *MetadataClusterSetTransLocaterReq) HasMetadataCluster() bool {
	if o != nil && !IsNil(o.MetadataCluster) {
		return true
	}

	return false
}

// SetMetadataCluster gets a reference to the given MetadataClusterSetTransLocaterReqMetadataCluster and assigns it to the MetadataCluster field.
func (o *MetadataClusterSetTransLocaterReq) SetMetadataCluster(v MetadataClusterSetTransLocaterReqMetadataCluster) {
	o.MetadataCluster = &v
}

func (o MetadataClusterSetTransLocaterReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataClusterSetTransLocaterReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetadataCluster) {
		toSerialize["metadata_cluster"] = o.MetadataCluster
	}
	return toSerialize, nil
}

type NullableMetadataClusterSetTransLocaterReq struct {
	value *MetadataClusterSetTransLocaterReq
	isSet bool
}

func (v NullableMetadataClusterSetTransLocaterReq) Get() *MetadataClusterSetTransLocaterReq {
	return v.value
}

func (v *NullableMetadataClusterSetTransLocaterReq) Set(val *MetadataClusterSetTransLocaterReq) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataClusterSetTransLocaterReq) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataClusterSetTransLocaterReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataClusterSetTransLocaterReq(val *MetadataClusterSetTransLocaterReq) *NullableMetadataClusterSetTransLocaterReq {
	return &NullableMetadataClusterSetTransLocaterReq{value: val, isSet: true}
}

func (v NullableMetadataClusterSetTransLocaterReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataClusterSetTransLocaterReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


