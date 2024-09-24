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

// checks if the ClusterServicesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterServicesResp{}

// ClusterServicesResp struct for ClusterServicesResp
type ClusterServicesResp struct {
	ClusterServices []ClusterService `json:"cluster_services,omitempty"`
}

// NewClusterServicesResp instantiates a new ClusterServicesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterServicesResp() *ClusterServicesResp {
	this := ClusterServicesResp{}
	return &this
}

// NewClusterServicesRespWithDefaults instantiates a new ClusterServicesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterServicesRespWithDefaults() *ClusterServicesResp {
	this := ClusterServicesResp{}
	return &this
}

// GetClusterServices returns the ClusterServices field value if set, zero value otherwise.
func (o *ClusterServicesResp) GetClusterServices() []ClusterService {
	if o == nil || IsNil(o.ClusterServices) {
		var ret []ClusterService
		return ret
	}
	return o.ClusterServices
}

// GetClusterServicesOk returns a tuple with the ClusterServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterServicesResp) GetClusterServicesOk() ([]ClusterService, bool) {
	if o == nil || IsNil(o.ClusterServices) {
		return nil, false
	}
	return o.ClusterServices, true
}

// HasClusterServices returns a boolean if a field has been set.
func (o *ClusterServicesResp) HasClusterServices() bool {
	if o != nil && !IsNil(o.ClusterServices) {
		return true
	}

	return false
}

// SetClusterServices gets a reference to the given []ClusterService and assigns it to the ClusterServices field.
func (o *ClusterServicesResp) SetClusterServices(v []ClusterService) {
	o.ClusterServices = v
}

func (o ClusterServicesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterServicesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterServices) {
		toSerialize["cluster_services"] = o.ClusterServices
	}
	return toSerialize, nil
}

type NullableClusterServicesResp struct {
	value *ClusterServicesResp
	isSet bool
}

func (v NullableClusterServicesResp) Get() *ClusterServicesResp {
	return v.value
}

func (v *NullableClusterServicesResp) Set(val *ClusterServicesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterServicesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterServicesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterServicesResp(val *ClusterServicesResp) *NullableClusterServicesResp {
	return &NullableClusterServicesResp{value: val, isSet: true}
}

func (v NullableClusterServicesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterServicesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


