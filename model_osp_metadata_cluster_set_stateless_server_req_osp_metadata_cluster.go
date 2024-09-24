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

// checks if the OspMetadataClusterSetStatelessServerReqOspMetadataCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OspMetadataClusterSetStatelessServerReqOspMetadataCluster{}

// OspMetadataClusterSetStatelessServerReqOspMetadataCluster struct for OspMetadataClusterSetStatelessServerReqOspMetadataCluster
type OspMetadataClusterSetStatelessServerReqOspMetadataCluster struct {
	HostIds []int64 `json:"host_ids,omitempty"`
	StatelessServerNum *int64 `json:"stateless_server_num,omitempty"`
}

// NewOspMetadataClusterSetStatelessServerReqOspMetadataCluster instantiates a new OspMetadataClusterSetStatelessServerReqOspMetadataCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOspMetadataClusterSetStatelessServerReqOspMetadataCluster() *OspMetadataClusterSetStatelessServerReqOspMetadataCluster {
	this := OspMetadataClusterSetStatelessServerReqOspMetadataCluster{}
	return &this
}

// NewOspMetadataClusterSetStatelessServerReqOspMetadataClusterWithDefaults instantiates a new OspMetadataClusterSetStatelessServerReqOspMetadataCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOspMetadataClusterSetStatelessServerReqOspMetadataClusterWithDefaults() *OspMetadataClusterSetStatelessServerReqOspMetadataCluster {
	this := OspMetadataClusterSetStatelessServerReqOspMetadataCluster{}
	return &this
}

// GetHostIds returns the HostIds field value if set, zero value otherwise.
func (o *OspMetadataClusterSetStatelessServerReqOspMetadataCluster) GetHostIds() []int64 {
	if o == nil || IsNil(o.HostIds) {
		var ret []int64
		return ret
	}
	return o.HostIds
}

// GetHostIdsOk returns a tuple with the HostIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterSetStatelessServerReqOspMetadataCluster) GetHostIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.HostIds) {
		return nil, false
	}
	return o.HostIds, true
}

// HasHostIds returns a boolean if a field has been set.
func (o *OspMetadataClusterSetStatelessServerReqOspMetadataCluster) HasHostIds() bool {
	if o != nil && !IsNil(o.HostIds) {
		return true
	}

	return false
}

// SetHostIds gets a reference to the given []int64 and assigns it to the HostIds field.
func (o *OspMetadataClusterSetStatelessServerReqOspMetadataCluster) SetHostIds(v []int64) {
	o.HostIds = v
}

// GetStatelessServerNum returns the StatelessServerNum field value if set, zero value otherwise.
func (o *OspMetadataClusterSetStatelessServerReqOspMetadataCluster) GetStatelessServerNum() int64 {
	if o == nil || IsNil(o.StatelessServerNum) {
		var ret int64
		return ret
	}
	return *o.StatelessServerNum
}

// GetStatelessServerNumOk returns a tuple with the StatelessServerNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterSetStatelessServerReqOspMetadataCluster) GetStatelessServerNumOk() (*int64, bool) {
	if o == nil || IsNil(o.StatelessServerNum) {
		return nil, false
	}
	return o.StatelessServerNum, true
}

// HasStatelessServerNum returns a boolean if a field has been set.
func (o *OspMetadataClusterSetStatelessServerReqOspMetadataCluster) HasStatelessServerNum() bool {
	if o != nil && !IsNil(o.StatelessServerNum) {
		return true
	}

	return false
}

// SetStatelessServerNum gets a reference to the given int64 and assigns it to the StatelessServerNum field.
func (o *OspMetadataClusterSetStatelessServerReqOspMetadataCluster) SetStatelessServerNum(v int64) {
	o.StatelessServerNum = &v
}

func (o OspMetadataClusterSetStatelessServerReqOspMetadataCluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OspMetadataClusterSetStatelessServerReqOspMetadataCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HostIds) {
		toSerialize["host_ids"] = o.HostIds
	}
	if !IsNil(o.StatelessServerNum) {
		toSerialize["stateless_server_num"] = o.StatelessServerNum
	}
	return toSerialize, nil
}

type NullableOspMetadataClusterSetStatelessServerReqOspMetadataCluster struct {
	value *OspMetadataClusterSetStatelessServerReqOspMetadataCluster
	isSet bool
}

func (v NullableOspMetadataClusterSetStatelessServerReqOspMetadataCluster) Get() *OspMetadataClusterSetStatelessServerReqOspMetadataCluster {
	return v.value
}

func (v *NullableOspMetadataClusterSetStatelessServerReqOspMetadataCluster) Set(val *OspMetadataClusterSetStatelessServerReqOspMetadataCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableOspMetadataClusterSetStatelessServerReqOspMetadataCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableOspMetadataClusterSetStatelessServerReqOspMetadataCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspMetadataClusterSetStatelessServerReqOspMetadataCluster(val *OspMetadataClusterSetStatelessServerReqOspMetadataCluster) *NullableOspMetadataClusterSetStatelessServerReqOspMetadataCluster {
	return &NullableOspMetadataClusterSetStatelessServerReqOspMetadataCluster{value: val, isSet: true}
}

func (v NullableOspMetadataClusterSetStatelessServerReqOspMetadataCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspMetadataClusterSetStatelessServerReqOspMetadataCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


