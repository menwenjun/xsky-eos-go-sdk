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

// checks if the OspMetadataClustersResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OspMetadataClustersResp{}

// OspMetadataClustersResp struct for OspMetadataClustersResp
type OspMetadataClustersResp struct {
	OspMetadataClusters []OspMetadataClusterRecord `json:"osp_metadata_clusters,omitempty"`
}

// NewOspMetadataClustersResp instantiates a new OspMetadataClustersResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOspMetadataClustersResp() *OspMetadataClustersResp {
	this := OspMetadataClustersResp{}
	return &this
}

// NewOspMetadataClustersRespWithDefaults instantiates a new OspMetadataClustersResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOspMetadataClustersRespWithDefaults() *OspMetadataClustersResp {
	this := OspMetadataClustersResp{}
	return &this
}

// GetOspMetadataClusters returns the OspMetadataClusters field value if set, zero value otherwise.
func (o *OspMetadataClustersResp) GetOspMetadataClusters() []OspMetadataClusterRecord {
	if o == nil || IsNil(o.OspMetadataClusters) {
		var ret []OspMetadataClusterRecord
		return ret
	}
	return o.OspMetadataClusters
}

// GetOspMetadataClustersOk returns a tuple with the OspMetadataClusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClustersResp) GetOspMetadataClustersOk() ([]OspMetadataClusterRecord, bool) {
	if o == nil || IsNil(o.OspMetadataClusters) {
		return nil, false
	}
	return o.OspMetadataClusters, true
}

// HasOspMetadataClusters returns a boolean if a field has been set.
func (o *OspMetadataClustersResp) HasOspMetadataClusters() bool {
	if o != nil && !IsNil(o.OspMetadataClusters) {
		return true
	}

	return false
}

// SetOspMetadataClusters gets a reference to the given []OspMetadataClusterRecord and assigns it to the OspMetadataClusters field.
func (o *OspMetadataClustersResp) SetOspMetadataClusters(v []OspMetadataClusterRecord) {
	o.OspMetadataClusters = v
}

func (o OspMetadataClustersResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OspMetadataClustersResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OspMetadataClusters) {
		toSerialize["osp_metadata_clusters"] = o.OspMetadataClusters
	}
	return toSerialize, nil
}

type NullableOspMetadataClustersResp struct {
	value *OspMetadataClustersResp
	isSet bool
}

func (v NullableOspMetadataClustersResp) Get() *OspMetadataClustersResp {
	return v.value
}

func (v *NullableOspMetadataClustersResp) Set(val *OspMetadataClustersResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOspMetadataClustersResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOspMetadataClustersResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspMetadataClustersResp(val *OspMetadataClustersResp) *NullableOspMetadataClustersResp {
	return &NullableOspMetadataClustersResp{value: val, isSet: true}
}

func (v NullableOspMetadataClustersResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspMetadataClustersResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


