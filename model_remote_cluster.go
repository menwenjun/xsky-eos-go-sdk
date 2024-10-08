/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the RemoteCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteCluster{}

// RemoteCluster RemoteCluster remote clusters communicated via public API
type RemoteCluster struct {
	AccessToken *string `json:"access_token,omitempty"`
	BlockReplicationNum *int64 `json:"block_replication_num,omitempty"`
	Connected *bool `json:"connected,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	FsId *string `json:"fs_id,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	OsZoneNum *int64 `json:"os_zone_num,omitempty"`
	RemoteClusterId *int64 `json:"remote_cluster_id,omitempty"`
	SnapshotReplicationNum *int64 `json:"snapshot_replication_num,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	Url *string `json:"url,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewRemoteCluster instantiates a new RemoteCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteCluster() *RemoteCluster {
	this := RemoteCluster{}
	return &this
}

// NewRemoteClusterWithDefaults instantiates a new RemoteCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteClusterWithDefaults() *RemoteCluster {
	this := RemoteCluster{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *RemoteCluster) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *RemoteCluster) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *RemoteCluster) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetBlockReplicationNum returns the BlockReplicationNum field value if set, zero value otherwise.
func (o *RemoteCluster) GetBlockReplicationNum() int64 {
	if o == nil || IsNil(o.BlockReplicationNum) {
		var ret int64
		return ret
	}
	return *o.BlockReplicationNum
}

// GetBlockReplicationNumOk returns a tuple with the BlockReplicationNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetBlockReplicationNumOk() (*int64, bool) {
	if o == nil || IsNil(o.BlockReplicationNum) {
		return nil, false
	}
	return o.BlockReplicationNum, true
}

// HasBlockReplicationNum returns a boolean if a field has been set.
func (o *RemoteCluster) HasBlockReplicationNum() bool {
	if o != nil && !IsNil(o.BlockReplicationNum) {
		return true
	}

	return false
}

// SetBlockReplicationNum gets a reference to the given int64 and assigns it to the BlockReplicationNum field.
func (o *RemoteCluster) SetBlockReplicationNum(v int64) {
	o.BlockReplicationNum = &v
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *RemoteCluster) GetConnected() bool {
	if o == nil || IsNil(o.Connected) {
		var ret bool
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Connected) {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *RemoteCluster) HasConnected() bool {
	if o != nil && !IsNil(o.Connected) {
		return true
	}

	return false
}

// SetConnected gets a reference to the given bool and assigns it to the Connected field.
func (o *RemoteCluster) SetConnected(v bool) {
	o.Connected = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *RemoteCluster) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *RemoteCluster) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *RemoteCluster) SetCreate(v time.Time) {
	o.Create = &v
}

// GetFsId returns the FsId field value if set, zero value otherwise.
func (o *RemoteCluster) GetFsId() string {
	if o == nil || IsNil(o.FsId) {
		var ret string
		return ret
	}
	return *o.FsId
}

// GetFsIdOk returns a tuple with the FsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetFsIdOk() (*string, bool) {
	if o == nil || IsNil(o.FsId) {
		return nil, false
	}
	return o.FsId, true
}

// HasFsId returns a boolean if a field has been set.
func (o *RemoteCluster) HasFsId() bool {
	if o != nil && !IsNil(o.FsId) {
		return true
	}

	return false
}

// SetFsId gets a reference to the given string and assigns it to the FsId field.
func (o *RemoteCluster) SetFsId(v string) {
	o.FsId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemoteCluster) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemoteCluster) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *RemoteCluster) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RemoteCluster) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RemoteCluster) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RemoteCluster) SetName(v string) {
	o.Name = &v
}

// GetOsZoneNum returns the OsZoneNum field value if set, zero value otherwise.
func (o *RemoteCluster) GetOsZoneNum() int64 {
	if o == nil || IsNil(o.OsZoneNum) {
		var ret int64
		return ret
	}
	return *o.OsZoneNum
}

// GetOsZoneNumOk returns a tuple with the OsZoneNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetOsZoneNumOk() (*int64, bool) {
	if o == nil || IsNil(o.OsZoneNum) {
		return nil, false
	}
	return o.OsZoneNum, true
}

// HasOsZoneNum returns a boolean if a field has been set.
func (o *RemoteCluster) HasOsZoneNum() bool {
	if o != nil && !IsNil(o.OsZoneNum) {
		return true
	}

	return false
}

// SetOsZoneNum gets a reference to the given int64 and assigns it to the OsZoneNum field.
func (o *RemoteCluster) SetOsZoneNum(v int64) {
	o.OsZoneNum = &v
}

// GetRemoteClusterId returns the RemoteClusterId field value if set, zero value otherwise.
func (o *RemoteCluster) GetRemoteClusterId() int64 {
	if o == nil || IsNil(o.RemoteClusterId) {
		var ret int64
		return ret
	}
	return *o.RemoteClusterId
}

// GetRemoteClusterIdOk returns a tuple with the RemoteClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetRemoteClusterIdOk() (*int64, bool) {
	if o == nil || IsNil(o.RemoteClusterId) {
		return nil, false
	}
	return o.RemoteClusterId, true
}

// HasRemoteClusterId returns a boolean if a field has been set.
func (o *RemoteCluster) HasRemoteClusterId() bool {
	if o != nil && !IsNil(o.RemoteClusterId) {
		return true
	}

	return false
}

// SetRemoteClusterId gets a reference to the given int64 and assigns it to the RemoteClusterId field.
func (o *RemoteCluster) SetRemoteClusterId(v int64) {
	o.RemoteClusterId = &v
}

// GetSnapshotReplicationNum returns the SnapshotReplicationNum field value if set, zero value otherwise.
func (o *RemoteCluster) GetSnapshotReplicationNum() int64 {
	if o == nil || IsNil(o.SnapshotReplicationNum) {
		var ret int64
		return ret
	}
	return *o.SnapshotReplicationNum
}

// GetSnapshotReplicationNumOk returns a tuple with the SnapshotReplicationNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetSnapshotReplicationNumOk() (*int64, bool) {
	if o == nil || IsNil(o.SnapshotReplicationNum) {
		return nil, false
	}
	return o.SnapshotReplicationNum, true
}

// HasSnapshotReplicationNum returns a boolean if a field has been set.
func (o *RemoteCluster) HasSnapshotReplicationNum() bool {
	if o != nil && !IsNil(o.SnapshotReplicationNum) {
		return true
	}

	return false
}

// SetSnapshotReplicationNum gets a reference to the given int64 and assigns it to the SnapshotReplicationNum field.
func (o *RemoteCluster) SetSnapshotReplicationNum(v int64) {
	o.SnapshotReplicationNum = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RemoteCluster) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RemoteCluster) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RemoteCluster) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *RemoteCluster) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *RemoteCluster) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *RemoteCluster) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RemoteCluster) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RemoteCluster) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RemoteCluster) SetUrl(v string) {
	o.Url = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RemoteCluster) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RemoteCluster) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *RemoteCluster) SetVersion(v string) {
	o.Version = &v
}

func (o RemoteCluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessToken) {
		toSerialize["access_token"] = o.AccessToken
	}
	if !IsNil(o.BlockReplicationNum) {
		toSerialize["block_replication_num"] = o.BlockReplicationNum
	}
	if !IsNil(o.Connected) {
		toSerialize["connected"] = o.Connected
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.FsId) {
		toSerialize["fs_id"] = o.FsId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OsZoneNum) {
		toSerialize["os_zone_num"] = o.OsZoneNum
	}
	if !IsNil(o.RemoteClusterId) {
		toSerialize["remote_cluster_id"] = o.RemoteClusterId
	}
	if !IsNil(o.SnapshotReplicationNum) {
		toSerialize["snapshot_replication_num"] = o.SnapshotReplicationNum
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableRemoteCluster struct {
	value *RemoteCluster
	isSet bool
}

func (v NullableRemoteCluster) Get() *RemoteCluster {
	return v.value
}

func (v *NullableRemoteCluster) Set(val *RemoteCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteCluster(val *RemoteCluster) *NullableRemoteCluster {
	return &NullableRemoteCluster{value: val, isSet: true}
}

func (v NullableRemoteCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


