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

// checks if the Osd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Osd{}

// Osd Osd defines the Ceph Osd +X:model:stats;stats_key=ID;stats_key_type=int64;
type Osd struct {
	ActionStatus *string `json:"action_status,omitempty"`
	BindOsdGroup *OsdGroupNestview `json:"bind_osd_group,omitempty"`
	BindPool *PoolNestview `json:"bind_pool,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	DataDir *string `json:"data_dir,omitempty"`
	Disk *Disk `json:"disk,omitempty"`
	EncryptEnabled *bool `json:"encrypt_enabled,omitempty"`
	ExitCount *int64 `json:"exit_count,omitempty"`
	ExitTime *time.Time `json:"exit_time,omitempty"`
	Host *HostNestview `json:"host,omitempty"`
	Id *int64 `json:"id,omitempty"`
	In *bool `json:"in,omitempty"`
	InitTime *time.Time `json:"init_time,omitempty"`
	LastScrubTime *time.Time `json:"last_scrub_time,omitempty"`
	MetaBytes *int64 `json:"meta_bytes,omitempty"`
	MinAllocSize *int64 `json:"min_alloc_size,omitempty"`
	Name *string `json:"name,omitempty"`
	NumaNode *int64 `json:"numa_node,omitempty"`
	OmapByte *int64 `json:"omap_byte,omitempty"`
	OsdGroup *OsdGroupNestview `json:"osd_group,omitempty"`
	OsdId *int64 `json:"osd_id,omitempty"`
	OspMetadataCluster *OspMetadataCluster `json:"osp_metadata_cluster,omitempty"`
	Partition *Partition `json:"partition,omitempty"`
	Pool *PoolNestview `json:"pool,omitempty"`
	ReadCacheSize *int64 `json:"read_cache_size,omitempty"`
	Role *string `json:"role,omitempty"`
	Status *string `json:"status,omitempty"`
	StatusClass *string `json:"status_class,omitempty"`
	Type *string `json:"type,omitempty"`
	Up *bool `json:"up,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

// NewOsd instantiates a new Osd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsd() *Osd {
	this := Osd{}
	return &this
}

// NewOsdWithDefaults instantiates a new Osd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsdWithDefaults() *Osd {
	this := Osd{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *Osd) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *Osd) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *Osd) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetBindOsdGroup returns the BindOsdGroup field value if set, zero value otherwise.
func (o *Osd) GetBindOsdGroup() OsdGroupNestview {
	if o == nil || IsNil(o.BindOsdGroup) {
		var ret OsdGroupNestview
		return ret
	}
	return *o.BindOsdGroup
}

// GetBindOsdGroupOk returns a tuple with the BindOsdGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetBindOsdGroupOk() (*OsdGroupNestview, bool) {
	if o == nil || IsNil(o.BindOsdGroup) {
		return nil, false
	}
	return o.BindOsdGroup, true
}

// HasBindOsdGroup returns a boolean if a field has been set.
func (o *Osd) HasBindOsdGroup() bool {
	if o != nil && !IsNil(o.BindOsdGroup) {
		return true
	}

	return false
}

// SetBindOsdGroup gets a reference to the given OsdGroupNestview and assigns it to the BindOsdGroup field.
func (o *Osd) SetBindOsdGroup(v OsdGroupNestview) {
	o.BindOsdGroup = &v
}

// GetBindPool returns the BindPool field value if set, zero value otherwise.
func (o *Osd) GetBindPool() PoolNestview {
	if o == nil || IsNil(o.BindPool) {
		var ret PoolNestview
		return ret
	}
	return *o.BindPool
}

// GetBindPoolOk returns a tuple with the BindPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetBindPoolOk() (*PoolNestview, bool) {
	if o == nil || IsNil(o.BindPool) {
		return nil, false
	}
	return o.BindPool, true
}

// HasBindPool returns a boolean if a field has been set.
func (o *Osd) HasBindPool() bool {
	if o != nil && !IsNil(o.BindPool) {
		return true
	}

	return false
}

// SetBindPool gets a reference to the given PoolNestview and assigns it to the BindPool field.
func (o *Osd) SetBindPool(v PoolNestview) {
	o.BindPool = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *Osd) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *Osd) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *Osd) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *Osd) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *Osd) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *Osd) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDataDir returns the DataDir field value if set, zero value otherwise.
func (o *Osd) GetDataDir() string {
	if o == nil || IsNil(o.DataDir) {
		var ret string
		return ret
	}
	return *o.DataDir
}

// GetDataDirOk returns a tuple with the DataDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetDataDirOk() (*string, bool) {
	if o == nil || IsNil(o.DataDir) {
		return nil, false
	}
	return o.DataDir, true
}

// HasDataDir returns a boolean if a field has been set.
func (o *Osd) HasDataDir() bool {
	if o != nil && !IsNil(o.DataDir) {
		return true
	}

	return false
}

// SetDataDir gets a reference to the given string and assigns it to the DataDir field.
func (o *Osd) SetDataDir(v string) {
	o.DataDir = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *Osd) GetDisk() Disk {
	if o == nil || IsNil(o.Disk) {
		var ret Disk
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetDiskOk() (*Disk, bool) {
	if o == nil || IsNil(o.Disk) {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *Osd) HasDisk() bool {
	if o != nil && !IsNil(o.Disk) {
		return true
	}

	return false
}

// SetDisk gets a reference to the given Disk and assigns it to the Disk field.
func (o *Osd) SetDisk(v Disk) {
	o.Disk = &v
}

// GetEncryptEnabled returns the EncryptEnabled field value if set, zero value otherwise.
func (o *Osd) GetEncryptEnabled() bool {
	if o == nil || IsNil(o.EncryptEnabled) {
		var ret bool
		return ret
	}
	return *o.EncryptEnabled
}

// GetEncryptEnabledOk returns a tuple with the EncryptEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetEncryptEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptEnabled) {
		return nil, false
	}
	return o.EncryptEnabled, true
}

// HasEncryptEnabled returns a boolean if a field has been set.
func (o *Osd) HasEncryptEnabled() bool {
	if o != nil && !IsNil(o.EncryptEnabled) {
		return true
	}

	return false
}

// SetEncryptEnabled gets a reference to the given bool and assigns it to the EncryptEnabled field.
func (o *Osd) SetEncryptEnabled(v bool) {
	o.EncryptEnabled = &v
}

// GetExitCount returns the ExitCount field value if set, zero value otherwise.
func (o *Osd) GetExitCount() int64 {
	if o == nil || IsNil(o.ExitCount) {
		var ret int64
		return ret
	}
	return *o.ExitCount
}

// GetExitCountOk returns a tuple with the ExitCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetExitCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ExitCount) {
		return nil, false
	}
	return o.ExitCount, true
}

// HasExitCount returns a boolean if a field has been set.
func (o *Osd) HasExitCount() bool {
	if o != nil && !IsNil(o.ExitCount) {
		return true
	}

	return false
}

// SetExitCount gets a reference to the given int64 and assigns it to the ExitCount field.
func (o *Osd) SetExitCount(v int64) {
	o.ExitCount = &v
}

// GetExitTime returns the ExitTime field value if set, zero value otherwise.
func (o *Osd) GetExitTime() time.Time {
	if o == nil || IsNil(o.ExitTime) {
		var ret time.Time
		return ret
	}
	return *o.ExitTime
}

// GetExitTimeOk returns a tuple with the ExitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetExitTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExitTime) {
		return nil, false
	}
	return o.ExitTime, true
}

// HasExitTime returns a boolean if a field has been set.
func (o *Osd) HasExitTime() bool {
	if o != nil && !IsNil(o.ExitTime) {
		return true
	}

	return false
}

// SetExitTime gets a reference to the given time.Time and assigns it to the ExitTime field.
func (o *Osd) SetExitTime(v time.Time) {
	o.ExitTime = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *Osd) GetHost() HostNestview {
	if o == nil || IsNil(o.Host) {
		var ret HostNestview
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetHostOk() (*HostNestview, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *Osd) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given HostNestview and assigns it to the Host field.
func (o *Osd) SetHost(v HostNestview) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Osd) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Osd) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Osd) SetId(v int64) {
	o.Id = &v
}

// GetIn returns the In field value if set, zero value otherwise.
func (o *Osd) GetIn() bool {
	if o == nil || IsNil(o.In) {
		var ret bool
		return ret
	}
	return *o.In
}

// GetInOk returns a tuple with the In field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetInOk() (*bool, bool) {
	if o == nil || IsNil(o.In) {
		return nil, false
	}
	return o.In, true
}

// HasIn returns a boolean if a field has been set.
func (o *Osd) HasIn() bool {
	if o != nil && !IsNil(o.In) {
		return true
	}

	return false
}

// SetIn gets a reference to the given bool and assigns it to the In field.
func (o *Osd) SetIn(v bool) {
	o.In = &v
}

// GetInitTime returns the InitTime field value if set, zero value otherwise.
func (o *Osd) GetInitTime() time.Time {
	if o == nil || IsNil(o.InitTime) {
		var ret time.Time
		return ret
	}
	return *o.InitTime
}

// GetInitTimeOk returns a tuple with the InitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetInitTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.InitTime) {
		return nil, false
	}
	return o.InitTime, true
}

// HasInitTime returns a boolean if a field has been set.
func (o *Osd) HasInitTime() bool {
	if o != nil && !IsNil(o.InitTime) {
		return true
	}

	return false
}

// SetInitTime gets a reference to the given time.Time and assigns it to the InitTime field.
func (o *Osd) SetInitTime(v time.Time) {
	o.InitTime = &v
}

// GetLastScrubTime returns the LastScrubTime field value if set, zero value otherwise.
func (o *Osd) GetLastScrubTime() time.Time {
	if o == nil || IsNil(o.LastScrubTime) {
		var ret time.Time
		return ret
	}
	return *o.LastScrubTime
}

// GetLastScrubTimeOk returns a tuple with the LastScrubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetLastScrubTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastScrubTime) {
		return nil, false
	}
	return o.LastScrubTime, true
}

// HasLastScrubTime returns a boolean if a field has been set.
func (o *Osd) HasLastScrubTime() bool {
	if o != nil && !IsNil(o.LastScrubTime) {
		return true
	}

	return false
}

// SetLastScrubTime gets a reference to the given time.Time and assigns it to the LastScrubTime field.
func (o *Osd) SetLastScrubTime(v time.Time) {
	o.LastScrubTime = &v
}

// GetMetaBytes returns the MetaBytes field value if set, zero value otherwise.
func (o *Osd) GetMetaBytes() int64 {
	if o == nil || IsNil(o.MetaBytes) {
		var ret int64
		return ret
	}
	return *o.MetaBytes
}

// GetMetaBytesOk returns a tuple with the MetaBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetMetaBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.MetaBytes) {
		return nil, false
	}
	return o.MetaBytes, true
}

// HasMetaBytes returns a boolean if a field has been set.
func (o *Osd) HasMetaBytes() bool {
	if o != nil && !IsNil(o.MetaBytes) {
		return true
	}

	return false
}

// SetMetaBytes gets a reference to the given int64 and assigns it to the MetaBytes field.
func (o *Osd) SetMetaBytes(v int64) {
	o.MetaBytes = &v
}

// GetMinAllocSize returns the MinAllocSize field value if set, zero value otherwise.
func (o *Osd) GetMinAllocSize() int64 {
	if o == nil || IsNil(o.MinAllocSize) {
		var ret int64
		return ret
	}
	return *o.MinAllocSize
}

// GetMinAllocSizeOk returns a tuple with the MinAllocSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetMinAllocSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MinAllocSize) {
		return nil, false
	}
	return o.MinAllocSize, true
}

// HasMinAllocSize returns a boolean if a field has been set.
func (o *Osd) HasMinAllocSize() bool {
	if o != nil && !IsNil(o.MinAllocSize) {
		return true
	}

	return false
}

// SetMinAllocSize gets a reference to the given int64 and assigns it to the MinAllocSize field.
func (o *Osd) SetMinAllocSize(v int64) {
	o.MinAllocSize = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Osd) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Osd) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Osd) SetName(v string) {
	o.Name = &v
}

// GetNumaNode returns the NumaNode field value if set, zero value otherwise.
func (o *Osd) GetNumaNode() int64 {
	if o == nil || IsNil(o.NumaNode) {
		var ret int64
		return ret
	}
	return *o.NumaNode
}

// GetNumaNodeOk returns a tuple with the NumaNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetNumaNodeOk() (*int64, bool) {
	if o == nil || IsNil(o.NumaNode) {
		return nil, false
	}
	return o.NumaNode, true
}

// HasNumaNode returns a boolean if a field has been set.
func (o *Osd) HasNumaNode() bool {
	if o != nil && !IsNil(o.NumaNode) {
		return true
	}

	return false
}

// SetNumaNode gets a reference to the given int64 and assigns it to the NumaNode field.
func (o *Osd) SetNumaNode(v int64) {
	o.NumaNode = &v
}

// GetOmapByte returns the OmapByte field value if set, zero value otherwise.
func (o *Osd) GetOmapByte() int64 {
	if o == nil || IsNil(o.OmapByte) {
		var ret int64
		return ret
	}
	return *o.OmapByte
}

// GetOmapByteOk returns a tuple with the OmapByte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetOmapByteOk() (*int64, bool) {
	if o == nil || IsNil(o.OmapByte) {
		return nil, false
	}
	return o.OmapByte, true
}

// HasOmapByte returns a boolean if a field has been set.
func (o *Osd) HasOmapByte() bool {
	if o != nil && !IsNil(o.OmapByte) {
		return true
	}

	return false
}

// SetOmapByte gets a reference to the given int64 and assigns it to the OmapByte field.
func (o *Osd) SetOmapByte(v int64) {
	o.OmapByte = &v
}

// GetOsdGroup returns the OsdGroup field value if set, zero value otherwise.
func (o *Osd) GetOsdGroup() OsdGroupNestview {
	if o == nil || IsNil(o.OsdGroup) {
		var ret OsdGroupNestview
		return ret
	}
	return *o.OsdGroup
}

// GetOsdGroupOk returns a tuple with the OsdGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetOsdGroupOk() (*OsdGroupNestview, bool) {
	if o == nil || IsNil(o.OsdGroup) {
		return nil, false
	}
	return o.OsdGroup, true
}

// HasOsdGroup returns a boolean if a field has been set.
func (o *Osd) HasOsdGroup() bool {
	if o != nil && !IsNil(o.OsdGroup) {
		return true
	}

	return false
}

// SetOsdGroup gets a reference to the given OsdGroupNestview and assigns it to the OsdGroup field.
func (o *Osd) SetOsdGroup(v OsdGroupNestview) {
	o.OsdGroup = &v
}

// GetOsdId returns the OsdId field value if set, zero value otherwise.
func (o *Osd) GetOsdId() int64 {
	if o == nil || IsNil(o.OsdId) {
		var ret int64
		return ret
	}
	return *o.OsdId
}

// GetOsdIdOk returns a tuple with the OsdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetOsdIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OsdId) {
		return nil, false
	}
	return o.OsdId, true
}

// HasOsdId returns a boolean if a field has been set.
func (o *Osd) HasOsdId() bool {
	if o != nil && !IsNil(o.OsdId) {
		return true
	}

	return false
}

// SetOsdId gets a reference to the given int64 and assigns it to the OsdId field.
func (o *Osd) SetOsdId(v int64) {
	o.OsdId = &v
}

// GetOspMetadataCluster returns the OspMetadataCluster field value if set, zero value otherwise.
func (o *Osd) GetOspMetadataCluster() OspMetadataCluster {
	if o == nil || IsNil(o.OspMetadataCluster) {
		var ret OspMetadataCluster
		return ret
	}
	return *o.OspMetadataCluster
}

// GetOspMetadataClusterOk returns a tuple with the OspMetadataCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetOspMetadataClusterOk() (*OspMetadataCluster, bool) {
	if o == nil || IsNil(o.OspMetadataCluster) {
		return nil, false
	}
	return o.OspMetadataCluster, true
}

// HasOspMetadataCluster returns a boolean if a field has been set.
func (o *Osd) HasOspMetadataCluster() bool {
	if o != nil && !IsNil(o.OspMetadataCluster) {
		return true
	}

	return false
}

// SetOspMetadataCluster gets a reference to the given OspMetadataCluster and assigns it to the OspMetadataCluster field.
func (o *Osd) SetOspMetadataCluster(v OspMetadataCluster) {
	o.OspMetadataCluster = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *Osd) GetPartition() Partition {
	if o == nil || IsNil(o.Partition) {
		var ret Partition
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetPartitionOk() (*Partition, bool) {
	if o == nil || IsNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *Osd) HasPartition() bool {
	if o != nil && !IsNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given Partition and assigns it to the Partition field.
func (o *Osd) SetPartition(v Partition) {
	o.Partition = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *Osd) GetPool() PoolNestview {
	if o == nil || IsNil(o.Pool) {
		var ret PoolNestview
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetPoolOk() (*PoolNestview, bool) {
	if o == nil || IsNil(o.Pool) {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *Osd) HasPool() bool {
	if o != nil && !IsNil(o.Pool) {
		return true
	}

	return false
}

// SetPool gets a reference to the given PoolNestview and assigns it to the Pool field.
func (o *Osd) SetPool(v PoolNestview) {
	o.Pool = &v
}

// GetReadCacheSize returns the ReadCacheSize field value if set, zero value otherwise.
func (o *Osd) GetReadCacheSize() int64 {
	if o == nil || IsNil(o.ReadCacheSize) {
		var ret int64
		return ret
	}
	return *o.ReadCacheSize
}

// GetReadCacheSizeOk returns a tuple with the ReadCacheSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetReadCacheSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadCacheSize) {
		return nil, false
	}
	return o.ReadCacheSize, true
}

// HasReadCacheSize returns a boolean if a field has been set.
func (o *Osd) HasReadCacheSize() bool {
	if o != nil && !IsNil(o.ReadCacheSize) {
		return true
	}

	return false
}

// SetReadCacheSize gets a reference to the given int64 and assigns it to the ReadCacheSize field.
func (o *Osd) SetReadCacheSize(v int64) {
	o.ReadCacheSize = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *Osd) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *Osd) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *Osd) SetRole(v string) {
	o.Role = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Osd) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Osd) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Osd) SetStatus(v string) {
	o.Status = &v
}

// GetStatusClass returns the StatusClass field value if set, zero value otherwise.
func (o *Osd) GetStatusClass() string {
	if o == nil || IsNil(o.StatusClass) {
		var ret string
		return ret
	}
	return *o.StatusClass
}

// GetStatusClassOk returns a tuple with the StatusClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetStatusClassOk() (*string, bool) {
	if o == nil || IsNil(o.StatusClass) {
		return nil, false
	}
	return o.StatusClass, true
}

// HasStatusClass returns a boolean if a field has been set.
func (o *Osd) HasStatusClass() bool {
	if o != nil && !IsNil(o.StatusClass) {
		return true
	}

	return false
}

// SetStatusClass gets a reference to the given string and assigns it to the StatusClass field.
func (o *Osd) SetStatusClass(v string) {
	o.StatusClass = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Osd) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Osd) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Osd) SetType(v string) {
	o.Type = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *Osd) GetUp() bool {
	if o == nil || IsNil(o.Up) {
		var ret bool
		return ret
	}
	return *o.Up
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetUpOk() (*bool, bool) {
	if o == nil || IsNil(o.Up) {
		return nil, false
	}
	return o.Up, true
}

// HasUp returns a boolean if a field has been set.
func (o *Osd) HasUp() bool {
	if o != nil && !IsNil(o.Up) {
		return true
	}

	return false
}

// SetUp gets a reference to the given bool and assigns it to the Up field.
func (o *Osd) SetUp(v bool) {
	o.Up = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *Osd) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *Osd) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *Osd) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Osd) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osd) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Osd) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Osd) SetUuid(v string) {
	o.Uuid = &v
}

func (o Osd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Osd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionStatus) {
		toSerialize["action_status"] = o.ActionStatus
	}
	if !IsNil(o.BindOsdGroup) {
		toSerialize["bind_osd_group"] = o.BindOsdGroup
	}
	if !IsNil(o.BindPool) {
		toSerialize["bind_pool"] = o.BindPool
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.DataDir) {
		toSerialize["data_dir"] = o.DataDir
	}
	if !IsNil(o.Disk) {
		toSerialize["disk"] = o.Disk
	}
	if !IsNil(o.EncryptEnabled) {
		toSerialize["encrypt_enabled"] = o.EncryptEnabled
	}
	if !IsNil(o.ExitCount) {
		toSerialize["exit_count"] = o.ExitCount
	}
	if !IsNil(o.ExitTime) {
		toSerialize["exit_time"] = o.ExitTime
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.In) {
		toSerialize["in"] = o.In
	}
	if !IsNil(o.InitTime) {
		toSerialize["init_time"] = o.InitTime
	}
	if !IsNil(o.LastScrubTime) {
		toSerialize["last_scrub_time"] = o.LastScrubTime
	}
	if !IsNil(o.MetaBytes) {
		toSerialize["meta_bytes"] = o.MetaBytes
	}
	if !IsNil(o.MinAllocSize) {
		toSerialize["min_alloc_size"] = o.MinAllocSize
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NumaNode) {
		toSerialize["numa_node"] = o.NumaNode
	}
	if !IsNil(o.OmapByte) {
		toSerialize["omap_byte"] = o.OmapByte
	}
	if !IsNil(o.OsdGroup) {
		toSerialize["osd_group"] = o.OsdGroup
	}
	if !IsNil(o.OsdId) {
		toSerialize["osd_id"] = o.OsdId
	}
	if !IsNil(o.OspMetadataCluster) {
		toSerialize["osp_metadata_cluster"] = o.OspMetadataCluster
	}
	if !IsNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	if !IsNil(o.Pool) {
		toSerialize["pool"] = o.Pool
	}
	if !IsNil(o.ReadCacheSize) {
		toSerialize["read_cache_size"] = o.ReadCacheSize
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusClass) {
		toSerialize["status_class"] = o.StatusClass
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Up) {
		toSerialize["up"] = o.Up
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableOsd struct {
	value *Osd
	isSet bool
}

func (v NullableOsd) Get() *Osd {
	return v.value
}

func (v *NullableOsd) Set(val *Osd) {
	v.value = val
	v.isSet = true
}

func (v NullableOsd) IsSet() bool {
	return v.isSet
}

func (v *NullableOsd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsd(val *Osd) *NullableOsd {
	return &NullableOsd{value: val, isSet: true}
}

func (v NullableOsd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


