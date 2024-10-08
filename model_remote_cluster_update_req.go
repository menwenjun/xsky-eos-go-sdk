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

// checks if the RemoteClusterUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteClusterUpdateReq{}

// RemoteClusterUpdateReq struct for RemoteClusterUpdateReq
type RemoteClusterUpdateReq struct {
	RemoteCluster *RemoteClusterUpdateReqRemoteCluster `json:"remote_cluster,omitempty"`
}

// NewRemoteClusterUpdateReq instantiates a new RemoteClusterUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteClusterUpdateReq() *RemoteClusterUpdateReq {
	this := RemoteClusterUpdateReq{}
	return &this
}

// NewRemoteClusterUpdateReqWithDefaults instantiates a new RemoteClusterUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteClusterUpdateReqWithDefaults() *RemoteClusterUpdateReq {
	this := RemoteClusterUpdateReq{}
	return &this
}

// GetRemoteCluster returns the RemoteCluster field value if set, zero value otherwise.
func (o *RemoteClusterUpdateReq) GetRemoteCluster() RemoteClusterUpdateReqRemoteCluster {
	if o == nil || IsNil(o.RemoteCluster) {
		var ret RemoteClusterUpdateReqRemoteCluster
		return ret
	}
	return *o.RemoteCluster
}

// GetRemoteClusterOk returns a tuple with the RemoteCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteClusterUpdateReq) GetRemoteClusterOk() (*RemoteClusterUpdateReqRemoteCluster, bool) {
	if o == nil || IsNil(o.RemoteCluster) {
		return nil, false
	}
	return o.RemoteCluster, true
}

// HasRemoteCluster returns a boolean if a field has been set.
func (o *RemoteClusterUpdateReq) HasRemoteCluster() bool {
	if o != nil && !IsNil(o.RemoteCluster) {
		return true
	}

	return false
}

// SetRemoteCluster gets a reference to the given RemoteClusterUpdateReqRemoteCluster and assigns it to the RemoteCluster field.
func (o *RemoteClusterUpdateReq) SetRemoteCluster(v RemoteClusterUpdateReqRemoteCluster) {
	o.RemoteCluster = &v
}

func (o RemoteClusterUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteClusterUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RemoteCluster) {
		toSerialize["remote_cluster"] = o.RemoteCluster
	}
	return toSerialize, nil
}

type NullableRemoteClusterUpdateReq struct {
	value *RemoteClusterUpdateReq
	isSet bool
}

func (v NullableRemoteClusterUpdateReq) Get() *RemoteClusterUpdateReq {
	return v.value
}

func (v *NullableRemoteClusterUpdateReq) Set(val *RemoteClusterUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteClusterUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteClusterUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteClusterUpdateReq(val *RemoteClusterUpdateReq) *NullableRemoteClusterUpdateReq {
	return &NullableRemoteClusterUpdateReq{value: val, isSet: true}
}

func (v NullableRemoteClusterUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteClusterUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


