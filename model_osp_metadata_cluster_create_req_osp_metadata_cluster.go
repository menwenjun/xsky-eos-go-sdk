/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OspMetadataClusterCreateReqOspMetadataCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OspMetadataClusterCreateReqOspMetadataCluster{}

// OspMetadataClusterCreateReqOspMetadataCluster struct for OspMetadataClusterCreateReqOspMetadataCluster
type OspMetadataClusterCreateReqOspMetadataCluster struct {
	ClusterId *int64 `json:"cluster_id,omitempty"`
	CommitProxyNum int64 `json:"commit_proxy_num"`
	CoordinatorHostIds []int64 `json:"coordinator_host_ids"`
	CoordinatorNum int64 `json:"coordinator_num"`
	DeployMode *string `json:"deploy_mode,omitempty"`
	GrvProxyNum int64 `json:"grv_proxy_num"`
	LogNum int64 `json:"log_num"`
	Name string `json:"name"`
	OsdIds []OspMetadataClusterService `json:"osd_ids,omitempty"`
	PartitionIds []OspMetadataClusterService `json:"partition_ids,omitempty"`
	ProtectionType string `json:"protection_type"`
	ReplicationNum int64 `json:"replication_num"`
	ResolverNum int64 `json:"resolver_num"`
	StatelessServerNum int64 `json:"stateless_server_num"`
	Type string `json:"type"`
}

type _OspMetadataClusterCreateReqOspMetadataCluster OspMetadataClusterCreateReqOspMetadataCluster

// NewOspMetadataClusterCreateReqOspMetadataCluster instantiates a new OspMetadataClusterCreateReqOspMetadataCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOspMetadataClusterCreateReqOspMetadataCluster(commitProxyNum int64, coordinatorHostIds []int64, coordinatorNum int64, grvProxyNum int64, logNum int64, name string, protectionType string, replicationNum int64, resolverNum int64, statelessServerNum int64, type_ string) *OspMetadataClusterCreateReqOspMetadataCluster {
	this := OspMetadataClusterCreateReqOspMetadataCluster{}
	this.CommitProxyNum = commitProxyNum
	this.CoordinatorHostIds = coordinatorHostIds
	this.CoordinatorNum = coordinatorNum
	this.GrvProxyNum = grvProxyNum
	this.LogNum = logNum
	this.Name = name
	this.ProtectionType = protectionType
	this.ReplicationNum = replicationNum
	this.ResolverNum = resolverNum
	this.StatelessServerNum = statelessServerNum
	this.Type = type_
	return &this
}

// NewOspMetadataClusterCreateReqOspMetadataClusterWithDefaults instantiates a new OspMetadataClusterCreateReqOspMetadataCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOspMetadataClusterCreateReqOspMetadataClusterWithDefaults() *OspMetadataClusterCreateReqOspMetadataCluster {
	this := OspMetadataClusterCreateReqOspMetadataCluster{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetClusterId() int64 {
	if o == nil || IsNil(o.ClusterId) {
		var ret int64
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetClusterIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given int64 and assigns it to the ClusterId field.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetClusterId(v int64) {
	o.ClusterId = &v
}

// GetCommitProxyNum returns the CommitProxyNum field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetCommitProxyNum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CommitProxyNum
}

// GetCommitProxyNumOk returns a tuple with the CommitProxyNum field value
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetCommitProxyNumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitProxyNum, true
}

// SetCommitProxyNum sets field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetCommitProxyNum(v int64) {
	o.CommitProxyNum = v
}

// GetCoordinatorHostIds returns the CoordinatorHostIds field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetCoordinatorHostIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.CoordinatorHostIds
}

// GetCoordinatorHostIdsOk returns a tuple with the CoordinatorHostIds field value
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetCoordinatorHostIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoordinatorHostIds, true
}

// SetCoordinatorHostIds sets field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetCoordinatorHostIds(v []int64) {
	o.CoordinatorHostIds = v
}

// GetCoordinatorNum returns the CoordinatorNum field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetCoordinatorNum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CoordinatorNum
}

// GetCoordinatorNumOk returns a tuple with the CoordinatorNum field value
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetCoordinatorNumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CoordinatorNum, true
}

// SetCoordinatorNum sets field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetCoordinatorNum(v int64) {
	o.CoordinatorNum = v
}

// GetDeployMode returns the DeployMode field value if set, zero value otherwise.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetDeployMode() string {
	if o == nil || IsNil(o.DeployMode) {
		var ret string
		return ret
	}
	return *o.DeployMode
}

// GetDeployModeOk returns a tuple with the DeployMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetDeployModeOk() (*string, bool) {
	if o == nil || IsNil(o.DeployMode) {
		return nil, false
	}
	return o.DeployMode, true
}

// HasDeployMode returns a boolean if a field has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) HasDeployMode() bool {
	if o != nil && !IsNil(o.DeployMode) {
		return true
	}

	return false
}

// SetDeployMode gets a reference to the given string and assigns it to the DeployMode field.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetDeployMode(v string) {
	o.DeployMode = &v
}

// GetGrvProxyNum returns the GrvProxyNum field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetGrvProxyNum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GrvProxyNum
}

// GetGrvProxyNumOk returns a tuple with the GrvProxyNum field value
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetGrvProxyNumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrvProxyNum, true
}

// SetGrvProxyNum sets field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetGrvProxyNum(v int64) {
	o.GrvProxyNum = v
}

// GetLogNum returns the LogNum field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetLogNum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LogNum
}

// GetLogNumOk returns a tuple with the LogNum field value
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetLogNumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogNum, true
}

// SetLogNum sets field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetLogNum(v int64) {
	o.LogNum = v
}

// GetName returns the Name field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetName(v string) {
	o.Name = v
}

// GetOsdIds returns the OsdIds field value if set, zero value otherwise.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetOsdIds() []OspMetadataClusterService {
	if o == nil || IsNil(o.OsdIds) {
		var ret []OspMetadataClusterService
		return ret
	}
	return o.OsdIds
}

// GetOsdIdsOk returns a tuple with the OsdIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetOsdIdsOk() ([]OspMetadataClusterService, bool) {
	if o == nil || IsNil(o.OsdIds) {
		return nil, false
	}
	return o.OsdIds, true
}

// HasOsdIds returns a boolean if a field has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) HasOsdIds() bool {
	if o != nil && !IsNil(o.OsdIds) {
		return true
	}

	return false
}

// SetOsdIds gets a reference to the given []OspMetadataClusterService and assigns it to the OsdIds field.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetOsdIds(v []OspMetadataClusterService) {
	o.OsdIds = v
}

// GetPartitionIds returns the PartitionIds field value if set, zero value otherwise.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetPartitionIds() []OspMetadataClusterService {
	if o == nil || IsNil(o.PartitionIds) {
		var ret []OspMetadataClusterService
		return ret
	}
	return o.PartitionIds
}

// GetPartitionIdsOk returns a tuple with the PartitionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetPartitionIdsOk() ([]OspMetadataClusterService, bool) {
	if o == nil || IsNil(o.PartitionIds) {
		return nil, false
	}
	return o.PartitionIds, true
}

// HasPartitionIds returns a boolean if a field has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) HasPartitionIds() bool {
	if o != nil && !IsNil(o.PartitionIds) {
		return true
	}

	return false
}

// SetPartitionIds gets a reference to the given []OspMetadataClusterService and assigns it to the PartitionIds field.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetPartitionIds(v []OspMetadataClusterService) {
	o.PartitionIds = v
}

// GetProtectionType returns the ProtectionType field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetProtectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProtectionType
}

// GetProtectionTypeOk returns a tuple with the ProtectionType field value
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetProtectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProtectionType, true
}

// SetProtectionType sets field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetProtectionType(v string) {
	o.ProtectionType = v
}

// GetReplicationNum returns the ReplicationNum field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetReplicationNum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ReplicationNum
}

// GetReplicationNumOk returns a tuple with the ReplicationNum field value
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetReplicationNumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplicationNum, true
}

// SetReplicationNum sets field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetReplicationNum(v int64) {
	o.ReplicationNum = v
}

// GetResolverNum returns the ResolverNum field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetResolverNum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ResolverNum
}

// GetResolverNumOk returns a tuple with the ResolverNum field value
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetResolverNumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResolverNum, true
}

// SetResolverNum sets field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetResolverNum(v int64) {
	o.ResolverNum = v
}

// GetStatelessServerNum returns the StatelessServerNum field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetStatelessServerNum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StatelessServerNum
}

// GetStatelessServerNumOk returns a tuple with the StatelessServerNum field value
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetStatelessServerNumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatelessServerNum, true
}

// SetStatelessServerNum sets field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetStatelessServerNum(v int64) {
	o.StatelessServerNum = v
}

// GetType returns the Type field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetType(v string) {
	o.Type = v
}

func (o OspMetadataClusterCreateReqOspMetadataCluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OspMetadataClusterCreateReqOspMetadataCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterId) {
		toSerialize["cluster_id"] = o.ClusterId
	}
	toSerialize["commit_proxy_num"] = o.CommitProxyNum
	toSerialize["coordinator_host_ids"] = o.CoordinatorHostIds
	toSerialize["coordinator_num"] = o.CoordinatorNum
	if !IsNil(o.DeployMode) {
		toSerialize["deploy_mode"] = o.DeployMode
	}
	toSerialize["grv_proxy_num"] = o.GrvProxyNum
	toSerialize["log_num"] = o.LogNum
	toSerialize["name"] = o.Name
	if !IsNil(o.OsdIds) {
		toSerialize["osd_ids"] = o.OsdIds
	}
	if !IsNil(o.PartitionIds) {
		toSerialize["partition_ids"] = o.PartitionIds
	}
	toSerialize["protection_type"] = o.ProtectionType
	toSerialize["replication_num"] = o.ReplicationNum
	toSerialize["resolver_num"] = o.ResolverNum
	toSerialize["stateless_server_num"] = o.StatelessServerNum
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *OspMetadataClusterCreateReqOspMetadataCluster) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"commit_proxy_num",
		"coordinator_host_ids",
		"coordinator_num",
		"grv_proxy_num",
		"log_num",
		"name",
		"protection_type",
		"replication_num",
		"resolver_num",
		"stateless_server_num",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOspMetadataClusterCreateReqOspMetadataCluster := _OspMetadataClusterCreateReqOspMetadataCluster{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOspMetadataClusterCreateReqOspMetadataCluster)

	if err != nil {
		return err
	}

	*o = OspMetadataClusterCreateReqOspMetadataCluster(varOspMetadataClusterCreateReqOspMetadataCluster)

	return err
}

type NullableOspMetadataClusterCreateReqOspMetadataCluster struct {
	value *OspMetadataClusterCreateReqOspMetadataCluster
	isSet bool
}

func (v NullableOspMetadataClusterCreateReqOspMetadataCluster) Get() *OspMetadataClusterCreateReqOspMetadataCluster {
	return v.value
}

func (v *NullableOspMetadataClusterCreateReqOspMetadataCluster) Set(val *OspMetadataClusterCreateReqOspMetadataCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableOspMetadataClusterCreateReqOspMetadataCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableOspMetadataClusterCreateReqOspMetadataCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspMetadataClusterCreateReqOspMetadataCluster(val *OspMetadataClusterCreateReqOspMetadataCluster) *NullableOspMetadataClusterCreateReqOspMetadataCluster {
	return &NullableOspMetadataClusterCreateReqOspMetadataCluster{value: val, isSet: true}
}

func (v NullableOspMetadataClusterCreateReqOspMetadataCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspMetadataClusterCreateReqOspMetadataCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


