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

// checks if the DpVolumeGroupSnapshotReplicationJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpVolumeGroupSnapshotReplicationJob{}

// DpVolumeGroupSnapshotReplicationJob DpVolumeGroupSnapshotReplicationJob is one execution of a dp volume group snapshot protection +X:model:generate;related_depth=2;check_get;with_detailed;
type DpVolumeGroupSnapshotReplicationJob struct {
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	DiffType *string `json:"diff_type,omitempty"`
	DpVolumeGroupSnapshotReplicationPair *DpVolumeGroupSnapshotReplicationPairNestview `json:"dp_volume_group_snapshot_replication_pair,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	Id *int64 `json:"id,omitempty"`
	MaxRetryTimes *int64 `json:"max_retry_times,omitempty"`
	Progress *float64 `json:"progress,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Snapshot *VolumeGroupSnapshotNestview `json:"snapshot,omitempty"`
	StartedAt *time.Time `json:"started_at,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	VolumeGroupName *string `json:"volume_group_name,omitempty"`
}

// NewDpVolumeGroupSnapshotReplicationJob instantiates a new DpVolumeGroupSnapshotReplicationJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpVolumeGroupSnapshotReplicationJob() *DpVolumeGroupSnapshotReplicationJob {
	this := DpVolumeGroupSnapshotReplicationJob{}
	return &this
}

// NewDpVolumeGroupSnapshotReplicationJobWithDefaults instantiates a new DpVolumeGroupSnapshotReplicationJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpVolumeGroupSnapshotReplicationJobWithDefaults() *DpVolumeGroupSnapshotReplicationJob {
	this := DpVolumeGroupSnapshotReplicationJob{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationJob) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *DpVolumeGroupSnapshotReplicationJob) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetDiffType returns the DiffType field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationJob) GetDiffType() string {
	if o == nil || IsNil(o.DiffType) {
		var ret string
		return ret
	}
	return *o.DiffType
}

// GetDiffTypeOk returns a tuple with the DiffType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) GetDiffTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DiffType) {
		return nil, false
	}
	return o.DiffType, true
}

// HasDiffType returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) HasDiffType() bool {
	if o != nil && !IsNil(o.DiffType) {
		return true
	}

	return false
}

// SetDiffType gets a reference to the given string and assigns it to the DiffType field.
func (o *DpVolumeGroupSnapshotReplicationJob) SetDiffType(v string) {
	o.DiffType = &v
}

// GetDpVolumeGroupSnapshotReplicationPair returns the DpVolumeGroupSnapshotReplicationPair field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationJob) GetDpVolumeGroupSnapshotReplicationPair() DpVolumeGroupSnapshotReplicationPairNestview {
	if o == nil || IsNil(o.DpVolumeGroupSnapshotReplicationPair) {
		var ret DpVolumeGroupSnapshotReplicationPairNestview
		return ret
	}
	return *o.DpVolumeGroupSnapshotReplicationPair
}

// GetDpVolumeGroupSnapshotReplicationPairOk returns a tuple with the DpVolumeGroupSnapshotReplicationPair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) GetDpVolumeGroupSnapshotReplicationPairOk() (*DpVolumeGroupSnapshotReplicationPairNestview, bool) {
	if o == nil || IsNil(o.DpVolumeGroupSnapshotReplicationPair) {
		return nil, false
	}
	return o.DpVolumeGroupSnapshotReplicationPair, true
}

// HasDpVolumeGroupSnapshotReplicationPair returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) HasDpVolumeGroupSnapshotReplicationPair() bool {
	if o != nil && !IsNil(o.DpVolumeGroupSnapshotReplicationPair) {
		return true
	}

	return false
}

// SetDpVolumeGroupSnapshotReplicationPair gets a reference to the given DpVolumeGroupSnapshotReplicationPairNestview and assigns it to the DpVolumeGroupSnapshotReplicationPair field.
func (o *DpVolumeGroupSnapshotReplicationJob) SetDpVolumeGroupSnapshotReplicationPair(v DpVolumeGroupSnapshotReplicationPairNestview) {
	o.DpVolumeGroupSnapshotReplicationPair = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationJob) GetFinishedAt() time.Time {
	if o == nil || IsNil(o.FinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FinishedAt) {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) HasFinishedAt() bool {
	if o != nil && !IsNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *DpVolumeGroupSnapshotReplicationJob) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationJob) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DpVolumeGroupSnapshotReplicationJob) SetId(v int64) {
	o.Id = &v
}

// GetMaxRetryTimes returns the MaxRetryTimes field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationJob) GetMaxRetryTimes() int64 {
	if o == nil || IsNil(o.MaxRetryTimes) {
		var ret int64
		return ret
	}
	return *o.MaxRetryTimes
}

// GetMaxRetryTimesOk returns a tuple with the MaxRetryTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) GetMaxRetryTimesOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxRetryTimes) {
		return nil, false
	}
	return o.MaxRetryTimes, true
}

// HasMaxRetryTimes returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) HasMaxRetryTimes() bool {
	if o != nil && !IsNil(o.MaxRetryTimes) {
		return true
	}

	return false
}

// SetMaxRetryTimes gets a reference to the given int64 and assigns it to the MaxRetryTimes field.
func (o *DpVolumeGroupSnapshotReplicationJob) SetMaxRetryTimes(v int64) {
	o.MaxRetryTimes = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationJob) GetProgress() float64 {
	if o == nil || IsNil(o.Progress) {
		var ret float64
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) GetProgressOk() (*float64, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given float64 and assigns it to the Progress field.
func (o *DpVolumeGroupSnapshotReplicationJob) SetProgress(v float64) {
	o.Progress = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationJob) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *DpVolumeGroupSnapshotReplicationJob) SetSize(v int64) {
	o.Size = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationJob) GetSnapshot() VolumeGroupSnapshotNestview {
	if o == nil || IsNil(o.Snapshot) {
		var ret VolumeGroupSnapshotNestview
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) GetSnapshotOk() (*VolumeGroupSnapshotNestview, bool) {
	if o == nil || IsNil(o.Snapshot) {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) HasSnapshot() bool {
	if o != nil && !IsNil(o.Snapshot) {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given VolumeGroupSnapshotNestview and assigns it to the Snapshot field.
func (o *DpVolumeGroupSnapshotReplicationJob) SetSnapshot(v VolumeGroupSnapshotNestview) {
	o.Snapshot = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationJob) GetStartedAt() time.Time {
	if o == nil || IsNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *DpVolumeGroupSnapshotReplicationJob) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationJob) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DpVolumeGroupSnapshotReplicationJob) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationJob) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DpVolumeGroupSnapshotReplicationJob) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVolumeGroupName returns the VolumeGroupName field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationJob) GetVolumeGroupName() string {
	if o == nil || IsNil(o.VolumeGroupName) {
		var ret string
		return ret
	}
	return *o.VolumeGroupName
}

// GetVolumeGroupNameOk returns a tuple with the VolumeGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) GetVolumeGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeGroupName) {
		return nil, false
	}
	return o.VolumeGroupName, true
}

// HasVolumeGroupName returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationJob) HasVolumeGroupName() bool {
	if o != nil && !IsNil(o.VolumeGroupName) {
		return true
	}

	return false
}

// SetVolumeGroupName gets a reference to the given string and assigns it to the VolumeGroupName field.
func (o *DpVolumeGroupSnapshotReplicationJob) SetVolumeGroupName(v string) {
	o.VolumeGroupName = &v
}

func (o DpVolumeGroupSnapshotReplicationJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpVolumeGroupSnapshotReplicationJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.DiffType) {
		toSerialize["diff_type"] = o.DiffType
	}
	if !IsNil(o.DpVolumeGroupSnapshotReplicationPair) {
		toSerialize["dp_volume_group_snapshot_replication_pair"] = o.DpVolumeGroupSnapshotReplicationPair
	}
	if !IsNil(o.FinishedAt) {
		toSerialize["finished_at"] = o.FinishedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MaxRetryTimes) {
		toSerialize["max_retry_times"] = o.MaxRetryTimes
	}
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Snapshot) {
		toSerialize["snapshot"] = o.Snapshot
	}
	if !IsNil(o.StartedAt) {
		toSerialize["started_at"] = o.StartedAt
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.VolumeGroupName) {
		toSerialize["volume_group_name"] = o.VolumeGroupName
	}
	return toSerialize, nil
}

type NullableDpVolumeGroupSnapshotReplicationJob struct {
	value *DpVolumeGroupSnapshotReplicationJob
	isSet bool
}

func (v NullableDpVolumeGroupSnapshotReplicationJob) Get() *DpVolumeGroupSnapshotReplicationJob {
	return v.value
}

func (v *NullableDpVolumeGroupSnapshotReplicationJob) Set(val *DpVolumeGroupSnapshotReplicationJob) {
	v.value = val
	v.isSet = true
}

func (v NullableDpVolumeGroupSnapshotReplicationJob) IsSet() bool {
	return v.isSet
}

func (v *NullableDpVolumeGroupSnapshotReplicationJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpVolumeGroupSnapshotReplicationJob(val *DpVolumeGroupSnapshotReplicationJob) *NullableDpVolumeGroupSnapshotReplicationJob {
	return &NullableDpVolumeGroupSnapshotReplicationJob{value: val, isSet: true}
}

func (v NullableDpVolumeGroupSnapshotReplicationJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpVolumeGroupSnapshotReplicationJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


