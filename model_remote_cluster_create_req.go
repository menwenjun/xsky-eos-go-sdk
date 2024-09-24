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

// checks if the RemoteClusterCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteClusterCreateReq{}

// RemoteClusterCreateReq struct for RemoteClusterCreateReq
type RemoteClusterCreateReq struct {
	RemoteCluster *RemoteClusterCreateReqRemoteCluster `json:"remote_cluster,omitempty"`
}

// NewRemoteClusterCreateReq instantiates a new RemoteClusterCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteClusterCreateReq() *RemoteClusterCreateReq {
	this := RemoteClusterCreateReq{}
	return &this
}

// NewRemoteClusterCreateReqWithDefaults instantiates a new RemoteClusterCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteClusterCreateReqWithDefaults() *RemoteClusterCreateReq {
	this := RemoteClusterCreateReq{}
	return &this
}

// GetRemoteCluster returns the RemoteCluster field value if set, zero value otherwise.
func (o *RemoteClusterCreateReq) GetRemoteCluster() RemoteClusterCreateReqRemoteCluster {
	if o == nil || IsNil(o.RemoteCluster) {
		var ret RemoteClusterCreateReqRemoteCluster
		return ret
	}
	return *o.RemoteCluster
}

// GetRemoteClusterOk returns a tuple with the RemoteCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteClusterCreateReq) GetRemoteClusterOk() (*RemoteClusterCreateReqRemoteCluster, bool) {
	if o == nil || IsNil(o.RemoteCluster) {
		return nil, false
	}
	return o.RemoteCluster, true
}

// HasRemoteCluster returns a boolean if a field has been set.
func (o *RemoteClusterCreateReq) HasRemoteCluster() bool {
	if o != nil && !IsNil(o.RemoteCluster) {
		return true
	}

	return false
}

// SetRemoteCluster gets a reference to the given RemoteClusterCreateReqRemoteCluster and assigns it to the RemoteCluster field.
func (o *RemoteClusterCreateReq) SetRemoteCluster(v RemoteClusterCreateReqRemoteCluster) {
	o.RemoteCluster = &v
}

func (o RemoteClusterCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteClusterCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RemoteCluster) {
		toSerialize["remote_cluster"] = o.RemoteCluster
	}
	return toSerialize, nil
}

type NullableRemoteClusterCreateReq struct {
	value *RemoteClusterCreateReq
	isSet bool
}

func (v NullableRemoteClusterCreateReq) Get() *RemoteClusterCreateReq {
	return v.value
}

func (v *NullableRemoteClusterCreateReq) Set(val *RemoteClusterCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteClusterCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteClusterCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteClusterCreateReq(val *RemoteClusterCreateReq) *NullableRemoteClusterCreateReq {
	return &NullableRemoteClusterCreateReq{value: val, isSet: true}
}

func (v NullableRemoteClusterCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteClusterCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


