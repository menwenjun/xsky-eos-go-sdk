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

// checks if the Snapshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Snapshot{}

// Snapshot Snapshot is the corresponding rbd snapshot
type Snapshot struct {
	ActionStatus *string `json:"action_status,omitempty"`
	AllocatedSize *int64 `json:"allocated_size,omitempty"`
	BlockConsistentSnapshot *ConsistentSnapshotNestview `json:"block_consistent_snapshot,omitempty"`
	BlockVolumeGroupSnapshot *VolumeGroupSnapshotNestview `json:"block_volume_group_snapshot,omitempty"`
	ClonedBlockVolumeNum *int64 `json:"cloned_block_volume_num,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Creator *string `json:"creator,omitempty"`
	Description *string `json:"description,omitempty"`
	Hidden *bool `json:"hidden,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Passive *bool `json:"passive,omitempty"`
	Pool *PoolNestview `json:"pool,omitempty"`
	Protected *bool `json:"protected,omitempty"`
	RemoteCluster *RemoteClusterNestview `json:"remote_cluster,omitempty"`
	// Deprecated, used by volume_backup
	Reserved *bool `json:"reserved,omitempty"`
	Size *int64 `json:"size,omitempty"`
	SnapName *string `json:"snap_name,omitempty"`
	Status *string `json:"status,omitempty"`
	Uid *string `json:"uid,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	Volume *VolumeNestview `json:"volume,omitempty"`
}

// NewSnapshot instantiates a new Snapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshot() *Snapshot {
	this := Snapshot{}
	return &this
}

// NewSnapshotWithDefaults instantiates a new Snapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotWithDefaults() *Snapshot {
	this := Snapshot{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *Snapshot) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *Snapshot) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *Snapshot) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetAllocatedSize returns the AllocatedSize field value if set, zero value otherwise.
func (o *Snapshot) GetAllocatedSize() int64 {
	if o == nil || IsNil(o.AllocatedSize) {
		var ret int64
		return ret
	}
	return *o.AllocatedSize
}

// GetAllocatedSizeOk returns a tuple with the AllocatedSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetAllocatedSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.AllocatedSize) {
		return nil, false
	}
	return o.AllocatedSize, true
}

// HasAllocatedSize returns a boolean if a field has been set.
func (o *Snapshot) HasAllocatedSize() bool {
	if o != nil && !IsNil(o.AllocatedSize) {
		return true
	}

	return false
}

// SetAllocatedSize gets a reference to the given int64 and assigns it to the AllocatedSize field.
func (o *Snapshot) SetAllocatedSize(v int64) {
	o.AllocatedSize = &v
}

// GetBlockConsistentSnapshot returns the BlockConsistentSnapshot field value if set, zero value otherwise.
func (o *Snapshot) GetBlockConsistentSnapshot() ConsistentSnapshotNestview {
	if o == nil || IsNil(o.BlockConsistentSnapshot) {
		var ret ConsistentSnapshotNestview
		return ret
	}
	return *o.BlockConsistentSnapshot
}

// GetBlockConsistentSnapshotOk returns a tuple with the BlockConsistentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetBlockConsistentSnapshotOk() (*ConsistentSnapshotNestview, bool) {
	if o == nil || IsNil(o.BlockConsistentSnapshot) {
		return nil, false
	}
	return o.BlockConsistentSnapshot, true
}

// HasBlockConsistentSnapshot returns a boolean if a field has been set.
func (o *Snapshot) HasBlockConsistentSnapshot() bool {
	if o != nil && !IsNil(o.BlockConsistentSnapshot) {
		return true
	}

	return false
}

// SetBlockConsistentSnapshot gets a reference to the given ConsistentSnapshotNestview and assigns it to the BlockConsistentSnapshot field.
func (o *Snapshot) SetBlockConsistentSnapshot(v ConsistentSnapshotNestview) {
	o.BlockConsistentSnapshot = &v
}

// GetBlockVolumeGroupSnapshot returns the BlockVolumeGroupSnapshot field value if set, zero value otherwise.
func (o *Snapshot) GetBlockVolumeGroupSnapshot() VolumeGroupSnapshotNestview {
	if o == nil || IsNil(o.BlockVolumeGroupSnapshot) {
		var ret VolumeGroupSnapshotNestview
		return ret
	}
	return *o.BlockVolumeGroupSnapshot
}

// GetBlockVolumeGroupSnapshotOk returns a tuple with the BlockVolumeGroupSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetBlockVolumeGroupSnapshotOk() (*VolumeGroupSnapshotNestview, bool) {
	if o == nil || IsNil(o.BlockVolumeGroupSnapshot) {
		return nil, false
	}
	return o.BlockVolumeGroupSnapshot, true
}

// HasBlockVolumeGroupSnapshot returns a boolean if a field has been set.
func (o *Snapshot) HasBlockVolumeGroupSnapshot() bool {
	if o != nil && !IsNil(o.BlockVolumeGroupSnapshot) {
		return true
	}

	return false
}

// SetBlockVolumeGroupSnapshot gets a reference to the given VolumeGroupSnapshotNestview and assigns it to the BlockVolumeGroupSnapshot field.
func (o *Snapshot) SetBlockVolumeGroupSnapshot(v VolumeGroupSnapshotNestview) {
	o.BlockVolumeGroupSnapshot = &v
}

// GetClonedBlockVolumeNum returns the ClonedBlockVolumeNum field value if set, zero value otherwise.
func (o *Snapshot) GetClonedBlockVolumeNum() int64 {
	if o == nil || IsNil(o.ClonedBlockVolumeNum) {
		var ret int64
		return ret
	}
	return *o.ClonedBlockVolumeNum
}

// GetClonedBlockVolumeNumOk returns a tuple with the ClonedBlockVolumeNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetClonedBlockVolumeNumOk() (*int64, bool) {
	if o == nil || IsNil(o.ClonedBlockVolumeNum) {
		return nil, false
	}
	return o.ClonedBlockVolumeNum, true
}

// HasClonedBlockVolumeNum returns a boolean if a field has been set.
func (o *Snapshot) HasClonedBlockVolumeNum() bool {
	if o != nil && !IsNil(o.ClonedBlockVolumeNum) {
		return true
	}

	return false
}

// SetClonedBlockVolumeNum gets a reference to the given int64 and assigns it to the ClonedBlockVolumeNum field.
func (o *Snapshot) SetClonedBlockVolumeNum(v int64) {
	o.ClonedBlockVolumeNum = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *Snapshot) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *Snapshot) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *Snapshot) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *Snapshot) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *Snapshot) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *Snapshot) SetCreate(v time.Time) {
	o.Create = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *Snapshot) GetCreator() string {
	if o == nil || IsNil(o.Creator) {
		var ret string
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetCreatorOk() (*string, bool) {
	if o == nil || IsNil(o.Creator) {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *Snapshot) HasCreator() bool {
	if o != nil && !IsNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given string and assigns it to the Creator field.
func (o *Snapshot) SetCreator(v string) {
	o.Creator = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Snapshot) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Snapshot) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Snapshot) SetDescription(v string) {
	o.Description = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *Snapshot) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *Snapshot) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *Snapshot) SetHidden(v bool) {
	o.Hidden = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Snapshot) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Snapshot) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Snapshot) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Snapshot) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Snapshot) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Snapshot) SetName(v string) {
	o.Name = &v
}

// GetPassive returns the Passive field value if set, zero value otherwise.
func (o *Snapshot) GetPassive() bool {
	if o == nil || IsNil(o.Passive) {
		var ret bool
		return ret
	}
	return *o.Passive
}

// GetPassiveOk returns a tuple with the Passive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetPassiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Passive) {
		return nil, false
	}
	return o.Passive, true
}

// HasPassive returns a boolean if a field has been set.
func (o *Snapshot) HasPassive() bool {
	if o != nil && !IsNil(o.Passive) {
		return true
	}

	return false
}

// SetPassive gets a reference to the given bool and assigns it to the Passive field.
func (o *Snapshot) SetPassive(v bool) {
	o.Passive = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *Snapshot) GetPool() PoolNestview {
	if o == nil || IsNil(o.Pool) {
		var ret PoolNestview
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetPoolOk() (*PoolNestview, bool) {
	if o == nil || IsNil(o.Pool) {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *Snapshot) HasPool() bool {
	if o != nil && !IsNil(o.Pool) {
		return true
	}

	return false
}

// SetPool gets a reference to the given PoolNestview and assigns it to the Pool field.
func (o *Snapshot) SetPool(v PoolNestview) {
	o.Pool = &v
}

// GetProtected returns the Protected field value if set, zero value otherwise.
func (o *Snapshot) GetProtected() bool {
	if o == nil || IsNil(o.Protected) {
		var ret bool
		return ret
	}
	return *o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Protected) {
		return nil, false
	}
	return o.Protected, true
}

// HasProtected returns a boolean if a field has been set.
func (o *Snapshot) HasProtected() bool {
	if o != nil && !IsNil(o.Protected) {
		return true
	}

	return false
}

// SetProtected gets a reference to the given bool and assigns it to the Protected field.
func (o *Snapshot) SetProtected(v bool) {
	o.Protected = &v
}

// GetRemoteCluster returns the RemoteCluster field value if set, zero value otherwise.
func (o *Snapshot) GetRemoteCluster() RemoteClusterNestview {
	if o == nil || IsNil(o.RemoteCluster) {
		var ret RemoteClusterNestview
		return ret
	}
	return *o.RemoteCluster
}

// GetRemoteClusterOk returns a tuple with the RemoteCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetRemoteClusterOk() (*RemoteClusterNestview, bool) {
	if o == nil || IsNil(o.RemoteCluster) {
		return nil, false
	}
	return o.RemoteCluster, true
}

// HasRemoteCluster returns a boolean if a field has been set.
func (o *Snapshot) HasRemoteCluster() bool {
	if o != nil && !IsNil(o.RemoteCluster) {
		return true
	}

	return false
}

// SetRemoteCluster gets a reference to the given RemoteClusterNestview and assigns it to the RemoteCluster field.
func (o *Snapshot) SetRemoteCluster(v RemoteClusterNestview) {
	o.RemoteCluster = &v
}

// GetReserved returns the Reserved field value if set, zero value otherwise.
func (o *Snapshot) GetReserved() bool {
	if o == nil || IsNil(o.Reserved) {
		var ret bool
		return ret
	}
	return *o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetReservedOk() (*bool, bool) {
	if o == nil || IsNil(o.Reserved) {
		return nil, false
	}
	return o.Reserved, true
}

// HasReserved returns a boolean if a field has been set.
func (o *Snapshot) HasReserved() bool {
	if o != nil && !IsNil(o.Reserved) {
		return true
	}

	return false
}

// SetReserved gets a reference to the given bool and assigns it to the Reserved field.
func (o *Snapshot) SetReserved(v bool) {
	o.Reserved = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Snapshot) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Snapshot) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *Snapshot) SetSize(v int64) {
	o.Size = &v
}

// GetSnapName returns the SnapName field value if set, zero value otherwise.
func (o *Snapshot) GetSnapName() string {
	if o == nil || IsNil(o.SnapName) {
		var ret string
		return ret
	}
	return *o.SnapName
}

// GetSnapNameOk returns a tuple with the SnapName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetSnapNameOk() (*string, bool) {
	if o == nil || IsNil(o.SnapName) {
		return nil, false
	}
	return o.SnapName, true
}

// HasSnapName returns a boolean if a field has been set.
func (o *Snapshot) HasSnapName() bool {
	if o != nil && !IsNil(o.SnapName) {
		return true
	}

	return false
}

// SetSnapName gets a reference to the given string and assigns it to the SnapName field.
func (o *Snapshot) SetSnapName(v string) {
	o.SnapName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Snapshot) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Snapshot) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Snapshot) SetStatus(v string) {
	o.Status = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *Snapshot) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *Snapshot) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *Snapshot) SetUid(v string) {
	o.Uid = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *Snapshot) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *Snapshot) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *Snapshot) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *Snapshot) GetVolume() VolumeNestview {
	if o == nil || IsNil(o.Volume) {
		var ret VolumeNestview
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetVolumeOk() (*VolumeNestview, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *Snapshot) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given VolumeNestview and assigns it to the Volume field.
func (o *Snapshot) SetVolume(v VolumeNestview) {
	o.Volume = &v
}

func (o Snapshot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Snapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionStatus) {
		toSerialize["action_status"] = o.ActionStatus
	}
	if !IsNil(o.AllocatedSize) {
		toSerialize["allocated_size"] = o.AllocatedSize
	}
	if !IsNil(o.BlockConsistentSnapshot) {
		toSerialize["block_consistent_snapshot"] = o.BlockConsistentSnapshot
	}
	if !IsNil(o.BlockVolumeGroupSnapshot) {
		toSerialize["block_volume_group_snapshot"] = o.BlockVolumeGroupSnapshot
	}
	if !IsNil(o.ClonedBlockVolumeNum) {
		toSerialize["cloned_block_volume_num"] = o.ClonedBlockVolumeNum
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Creator) {
		toSerialize["creator"] = o.Creator
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Passive) {
		toSerialize["passive"] = o.Passive
	}
	if !IsNil(o.Pool) {
		toSerialize["pool"] = o.Pool
	}
	if !IsNil(o.Protected) {
		toSerialize["protected"] = o.Protected
	}
	if !IsNil(o.RemoteCluster) {
		toSerialize["remote_cluster"] = o.RemoteCluster
	}
	if !IsNil(o.Reserved) {
		toSerialize["reserved"] = o.Reserved
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.SnapName) {
		toSerialize["snap_name"] = o.SnapName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

type NullableSnapshot struct {
	value *Snapshot
	isSet bool
}

func (v NullableSnapshot) Get() *Snapshot {
	return v.value
}

func (v *NullableSnapshot) Set(val *Snapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshot(val *Snapshot) *NullableSnapshot {
	return &NullableSnapshot{value: val, isSet: true}
}

func (v NullableSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


