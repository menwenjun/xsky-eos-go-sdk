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

// checks if the OspMetadataClusterSetStatelessNumReqOspMetadataCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OspMetadataClusterSetStatelessNumReqOspMetadataCluster{}

// OspMetadataClusterSetStatelessNumReqOspMetadataCluster struct for OspMetadataClusterSetStatelessNumReqOspMetadataCluster
type OspMetadataClusterSetStatelessNumReqOspMetadataCluster struct {
	CommitProxyNum *int64 `json:"commit_proxy_num,omitempty"`
	GrvProxyNum *int64 `json:"grv_proxy_num,omitempty"`
	ResolverNum *int64 `json:"resolver_num,omitempty"`
}

// NewOspMetadataClusterSetStatelessNumReqOspMetadataCluster instantiates a new OspMetadataClusterSetStatelessNumReqOspMetadataCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOspMetadataClusterSetStatelessNumReqOspMetadataCluster() *OspMetadataClusterSetStatelessNumReqOspMetadataCluster {
	this := OspMetadataClusterSetStatelessNumReqOspMetadataCluster{}
	return &this
}

// NewOspMetadataClusterSetStatelessNumReqOspMetadataClusterWithDefaults instantiates a new OspMetadataClusterSetStatelessNumReqOspMetadataCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOspMetadataClusterSetStatelessNumReqOspMetadataClusterWithDefaults() *OspMetadataClusterSetStatelessNumReqOspMetadataCluster {
	this := OspMetadataClusterSetStatelessNumReqOspMetadataCluster{}
	return &this
}

// GetCommitProxyNum returns the CommitProxyNum field value if set, zero value otherwise.
func (o *OspMetadataClusterSetStatelessNumReqOspMetadataCluster) GetCommitProxyNum() int64 {
	if o == nil || IsNil(o.CommitProxyNum) {
		var ret int64
		return ret
	}
	return *o.CommitProxyNum
}

// GetCommitProxyNumOk returns a tuple with the CommitProxyNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterSetStatelessNumReqOspMetadataCluster) GetCommitProxyNumOk() (*int64, bool) {
	if o == nil || IsNil(o.CommitProxyNum) {
		return nil, false
	}
	return o.CommitProxyNum, true
}

// HasCommitProxyNum returns a boolean if a field has been set.
func (o *OspMetadataClusterSetStatelessNumReqOspMetadataCluster) HasCommitProxyNum() bool {
	if o != nil && !IsNil(o.CommitProxyNum) {
		return true
	}

	return false
}

// SetCommitProxyNum gets a reference to the given int64 and assigns it to the CommitProxyNum field.
func (o *OspMetadataClusterSetStatelessNumReqOspMetadataCluster) SetCommitProxyNum(v int64) {
	o.CommitProxyNum = &v
}

// GetGrvProxyNum returns the GrvProxyNum field value if set, zero value otherwise.
func (o *OspMetadataClusterSetStatelessNumReqOspMetadataCluster) GetGrvProxyNum() int64 {
	if o == nil || IsNil(o.GrvProxyNum) {
		var ret int64
		return ret
	}
	return *o.GrvProxyNum
}

// GetGrvProxyNumOk returns a tuple with the GrvProxyNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterSetStatelessNumReqOspMetadataCluster) GetGrvProxyNumOk() (*int64, bool) {
	if o == nil || IsNil(o.GrvProxyNum) {
		return nil, false
	}
	return o.GrvProxyNum, true
}

// HasGrvProxyNum returns a boolean if a field has been set.
func (o *OspMetadataClusterSetStatelessNumReqOspMetadataCluster) HasGrvProxyNum() bool {
	if o != nil && !IsNil(o.GrvProxyNum) {
		return true
	}

	return false
}

// SetGrvProxyNum gets a reference to the given int64 and assigns it to the GrvProxyNum field.
func (o *OspMetadataClusterSetStatelessNumReqOspMetadataCluster) SetGrvProxyNum(v int64) {
	o.GrvProxyNum = &v
}

// GetResolverNum returns the ResolverNum field value if set, zero value otherwise.
func (o *OspMetadataClusterSetStatelessNumReqOspMetadataCluster) GetResolverNum() int64 {
	if o == nil || IsNil(o.ResolverNum) {
		var ret int64
		return ret
	}
	return *o.ResolverNum
}

// GetResolverNumOk returns a tuple with the ResolverNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterSetStatelessNumReqOspMetadataCluster) GetResolverNumOk() (*int64, bool) {
	if o == nil || IsNil(o.ResolverNum) {
		return nil, false
	}
	return o.ResolverNum, true
}

// HasResolverNum returns a boolean if a field has been set.
func (o *OspMetadataClusterSetStatelessNumReqOspMetadataCluster) HasResolverNum() bool {
	if o != nil && !IsNil(o.ResolverNum) {
		return true
	}

	return false
}

// SetResolverNum gets a reference to the given int64 and assigns it to the ResolverNum field.
func (o *OspMetadataClusterSetStatelessNumReqOspMetadataCluster) SetResolverNum(v int64) {
	o.ResolverNum = &v
}

func (o OspMetadataClusterSetStatelessNumReqOspMetadataCluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OspMetadataClusterSetStatelessNumReqOspMetadataCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommitProxyNum) {
		toSerialize["commit_proxy_num"] = o.CommitProxyNum
	}
	if !IsNil(o.GrvProxyNum) {
		toSerialize["grv_proxy_num"] = o.GrvProxyNum
	}
	if !IsNil(o.ResolverNum) {
		toSerialize["resolver_num"] = o.ResolverNum
	}
	return toSerialize, nil
}

type NullableOspMetadataClusterSetStatelessNumReqOspMetadataCluster struct {
	value *OspMetadataClusterSetStatelessNumReqOspMetadataCluster
	isSet bool
}

func (v NullableOspMetadataClusterSetStatelessNumReqOspMetadataCluster) Get() *OspMetadataClusterSetStatelessNumReqOspMetadataCluster {
	return v.value
}

func (v *NullableOspMetadataClusterSetStatelessNumReqOspMetadataCluster) Set(val *OspMetadataClusterSetStatelessNumReqOspMetadataCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableOspMetadataClusterSetStatelessNumReqOspMetadataCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableOspMetadataClusterSetStatelessNumReqOspMetadataCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspMetadataClusterSetStatelessNumReqOspMetadataCluster(val *OspMetadataClusterSetStatelessNumReqOspMetadataCluster) *NullableOspMetadataClusterSetStatelessNumReqOspMetadataCluster {
	return &NullableOspMetadataClusterSetStatelessNumReqOspMetadataCluster{value: val, isSet: true}
}

func (v NullableOspMetadataClusterSetStatelessNumReqOspMetadataCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspMetadataClusterSetStatelessNumReqOspMetadataCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


