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

// checks if the ClusterAddHostsReqCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterAddHostsReqCluster{}

// ClusterAddHostsReqCluster struct for ClusterAddHostsReqCluster
type ClusterAddHostsReqCluster struct {
	// auto select osp monitor role if true
	AutoSelectOspMonitorRole *bool `json:"auto_select_osp_monitor_role,omitempty"`
	// ceph host list
	CephHosts []CephHostReq `json:"ceph_hosts,omitempty"`
	// osp host list
	OspHosts []OspHostReq `json:"osp_hosts,omitempty"`
}

// NewClusterAddHostsReqCluster instantiates a new ClusterAddHostsReqCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterAddHostsReqCluster() *ClusterAddHostsReqCluster {
	this := ClusterAddHostsReqCluster{}
	return &this
}

// NewClusterAddHostsReqClusterWithDefaults instantiates a new ClusterAddHostsReqCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterAddHostsReqClusterWithDefaults() *ClusterAddHostsReqCluster {
	this := ClusterAddHostsReqCluster{}
	return &this
}

// GetAutoSelectOspMonitorRole returns the AutoSelectOspMonitorRole field value if set, zero value otherwise.
func (o *ClusterAddHostsReqCluster) GetAutoSelectOspMonitorRole() bool {
	if o == nil || IsNil(o.AutoSelectOspMonitorRole) {
		var ret bool
		return ret
	}
	return *o.AutoSelectOspMonitorRole
}

// GetAutoSelectOspMonitorRoleOk returns a tuple with the AutoSelectOspMonitorRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAddHostsReqCluster) GetAutoSelectOspMonitorRoleOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoSelectOspMonitorRole) {
		return nil, false
	}
	return o.AutoSelectOspMonitorRole, true
}

// HasAutoSelectOspMonitorRole returns a boolean if a field has been set.
func (o *ClusterAddHostsReqCluster) HasAutoSelectOspMonitorRole() bool {
	if o != nil && !IsNil(o.AutoSelectOspMonitorRole) {
		return true
	}

	return false
}

// SetAutoSelectOspMonitorRole gets a reference to the given bool and assigns it to the AutoSelectOspMonitorRole field.
func (o *ClusterAddHostsReqCluster) SetAutoSelectOspMonitorRole(v bool) {
	o.AutoSelectOspMonitorRole = &v
}

// GetCephHosts returns the CephHosts field value if set, zero value otherwise.
func (o *ClusterAddHostsReqCluster) GetCephHosts() []CephHostReq {
	if o == nil || IsNil(o.CephHosts) {
		var ret []CephHostReq
		return ret
	}
	return o.CephHosts
}

// GetCephHostsOk returns a tuple with the CephHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAddHostsReqCluster) GetCephHostsOk() ([]CephHostReq, bool) {
	if o == nil || IsNil(o.CephHosts) {
		return nil, false
	}
	return o.CephHosts, true
}

// HasCephHosts returns a boolean if a field has been set.
func (o *ClusterAddHostsReqCluster) HasCephHosts() bool {
	if o != nil && !IsNil(o.CephHosts) {
		return true
	}

	return false
}

// SetCephHosts gets a reference to the given []CephHostReq and assigns it to the CephHosts field.
func (o *ClusterAddHostsReqCluster) SetCephHosts(v []CephHostReq) {
	o.CephHosts = v
}

// GetOspHosts returns the OspHosts field value if set, zero value otherwise.
func (o *ClusterAddHostsReqCluster) GetOspHosts() []OspHostReq {
	if o == nil || IsNil(o.OspHosts) {
		var ret []OspHostReq
		return ret
	}
	return o.OspHosts
}

// GetOspHostsOk returns a tuple with the OspHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAddHostsReqCluster) GetOspHostsOk() ([]OspHostReq, bool) {
	if o == nil || IsNil(o.OspHosts) {
		return nil, false
	}
	return o.OspHosts, true
}

// HasOspHosts returns a boolean if a field has been set.
func (o *ClusterAddHostsReqCluster) HasOspHosts() bool {
	if o != nil && !IsNil(o.OspHosts) {
		return true
	}

	return false
}

// SetOspHosts gets a reference to the given []OspHostReq and assigns it to the OspHosts field.
func (o *ClusterAddHostsReqCluster) SetOspHosts(v []OspHostReq) {
	o.OspHosts = v
}

func (o ClusterAddHostsReqCluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterAddHostsReqCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoSelectOspMonitorRole) {
		toSerialize["auto_select_osp_monitor_role"] = o.AutoSelectOspMonitorRole
	}
	if !IsNil(o.CephHosts) {
		toSerialize["ceph_hosts"] = o.CephHosts
	}
	if !IsNil(o.OspHosts) {
		toSerialize["osp_hosts"] = o.OspHosts
	}
	return toSerialize, nil
}

type NullableClusterAddHostsReqCluster struct {
	value *ClusterAddHostsReqCluster
	isSet bool
}

func (v NullableClusterAddHostsReqCluster) Get() *ClusterAddHostsReqCluster {
	return v.value
}

func (v *NullableClusterAddHostsReqCluster) Set(val *ClusterAddHostsReqCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterAddHostsReqCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterAddHostsReqCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterAddHostsReqCluster(val *ClusterAddHostsReqCluster) *NullableClusterAddHostsReqCluster {
	return &NullableClusterAddHostsReqCluster{value: val, isSet: true}
}

func (v NullableClusterAddHostsReqCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterAddHostsReqCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


