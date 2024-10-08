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

// checks if the DpBlockSnapshotRecoveryJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpBlockSnapshotRecoveryJob{}

// DpBlockSnapshotRecoveryJob DpBlockSnapshotRecoveryJob recovery a block resource
type DpBlockSnapshotRecoveryJob struct {
	BackupBlockSnapshot *DpBackupBlockSnapshot `json:"backup_block_snapshot,omitempty"`
	BackupBlockVolume *DpBackupBlockVolume `json:"backup_block_volume,omitempty"`
	BackupCluster *DpBackupCluster `json:"backup_cluster,omitempty"`
	BlockVolume *VolumeNestview `json:"block_volume,omitempty"`
	Cluster *Cluster `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	DpGateway *DpGatewayNestview `json:"dp_gateway,omitempty"`
	DpSite *DpSiteNestview `json:"dp_site,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	Id *int64 `json:"id,omitempty"`
	MaxRetryTimes *int64 `json:"max_retry_times,omitempty"`
	Progress *float64 `json:"progress,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
	StartedAt *time.Time `json:"started_at,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewDpBlockSnapshotRecoveryJob instantiates a new DpBlockSnapshotRecoveryJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpBlockSnapshotRecoveryJob() *DpBlockSnapshotRecoveryJob {
	this := DpBlockSnapshotRecoveryJob{}
	return &this
}

// NewDpBlockSnapshotRecoveryJobWithDefaults instantiates a new DpBlockSnapshotRecoveryJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpBlockSnapshotRecoveryJobWithDefaults() *DpBlockSnapshotRecoveryJob {
	this := DpBlockSnapshotRecoveryJob{}
	return &this
}

// GetBackupBlockSnapshot returns the BackupBlockSnapshot field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetBackupBlockSnapshot() DpBackupBlockSnapshot {
	if o == nil || IsNil(o.BackupBlockSnapshot) {
		var ret DpBackupBlockSnapshot
		return ret
	}
	return *o.BackupBlockSnapshot
}

// GetBackupBlockSnapshotOk returns a tuple with the BackupBlockSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetBackupBlockSnapshotOk() (*DpBackupBlockSnapshot, bool) {
	if o == nil || IsNil(o.BackupBlockSnapshot) {
		return nil, false
	}
	return o.BackupBlockSnapshot, true
}

// HasBackupBlockSnapshot returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasBackupBlockSnapshot() bool {
	if o != nil && !IsNil(o.BackupBlockSnapshot) {
		return true
	}

	return false
}

// SetBackupBlockSnapshot gets a reference to the given DpBackupBlockSnapshot and assigns it to the BackupBlockSnapshot field.
func (o *DpBlockSnapshotRecoveryJob) SetBackupBlockSnapshot(v DpBackupBlockSnapshot) {
	o.BackupBlockSnapshot = &v
}

// GetBackupBlockVolume returns the BackupBlockVolume field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetBackupBlockVolume() DpBackupBlockVolume {
	if o == nil || IsNil(o.BackupBlockVolume) {
		var ret DpBackupBlockVolume
		return ret
	}
	return *o.BackupBlockVolume
}

// GetBackupBlockVolumeOk returns a tuple with the BackupBlockVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetBackupBlockVolumeOk() (*DpBackupBlockVolume, bool) {
	if o == nil || IsNil(o.BackupBlockVolume) {
		return nil, false
	}
	return o.BackupBlockVolume, true
}

// HasBackupBlockVolume returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasBackupBlockVolume() bool {
	if o != nil && !IsNil(o.BackupBlockVolume) {
		return true
	}

	return false
}

// SetBackupBlockVolume gets a reference to the given DpBackupBlockVolume and assigns it to the BackupBlockVolume field.
func (o *DpBlockSnapshotRecoveryJob) SetBackupBlockVolume(v DpBackupBlockVolume) {
	o.BackupBlockVolume = &v
}

// GetBackupCluster returns the BackupCluster field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetBackupCluster() DpBackupCluster {
	if o == nil || IsNil(o.BackupCluster) {
		var ret DpBackupCluster
		return ret
	}
	return *o.BackupCluster
}

// GetBackupClusterOk returns a tuple with the BackupCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetBackupClusterOk() (*DpBackupCluster, bool) {
	if o == nil || IsNil(o.BackupCluster) {
		return nil, false
	}
	return o.BackupCluster, true
}

// HasBackupCluster returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasBackupCluster() bool {
	if o != nil && !IsNil(o.BackupCluster) {
		return true
	}

	return false
}

// SetBackupCluster gets a reference to the given DpBackupCluster and assigns it to the BackupCluster field.
func (o *DpBlockSnapshotRecoveryJob) SetBackupCluster(v DpBackupCluster) {
	o.BackupCluster = &v
}

// GetBlockVolume returns the BlockVolume field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetBlockVolume() VolumeNestview {
	if o == nil || IsNil(o.BlockVolume) {
		var ret VolumeNestview
		return ret
	}
	return *o.BlockVolume
}

// GetBlockVolumeOk returns a tuple with the BlockVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetBlockVolumeOk() (*VolumeNestview, bool) {
	if o == nil || IsNil(o.BlockVolume) {
		return nil, false
	}
	return o.BlockVolume, true
}

// HasBlockVolume returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasBlockVolume() bool {
	if o != nil && !IsNil(o.BlockVolume) {
		return true
	}

	return false
}

// SetBlockVolume gets a reference to the given VolumeNestview and assigns it to the BlockVolume field.
func (o *DpBlockSnapshotRecoveryJob) SetBlockVolume(v VolumeNestview) {
	o.BlockVolume = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetCluster() Cluster {
	if o == nil || IsNil(o.Cluster) {
		var ret Cluster
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetClusterOk() (*Cluster, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given Cluster and assigns it to the Cluster field.
func (o *DpBlockSnapshotRecoveryJob) SetCluster(v Cluster) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DpBlockSnapshotRecoveryJob) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDpGateway returns the DpGateway field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetDpGateway() DpGatewayNestview {
	if o == nil || IsNil(o.DpGateway) {
		var ret DpGatewayNestview
		return ret
	}
	return *o.DpGateway
}

// GetDpGatewayOk returns a tuple with the DpGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetDpGatewayOk() (*DpGatewayNestview, bool) {
	if o == nil || IsNil(o.DpGateway) {
		return nil, false
	}
	return o.DpGateway, true
}

// HasDpGateway returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasDpGateway() bool {
	if o != nil && !IsNil(o.DpGateway) {
		return true
	}

	return false
}

// SetDpGateway gets a reference to the given DpGatewayNestview and assigns it to the DpGateway field.
func (o *DpBlockSnapshotRecoveryJob) SetDpGateway(v DpGatewayNestview) {
	o.DpGateway = &v
}

// GetDpSite returns the DpSite field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetDpSite() DpSiteNestview {
	if o == nil || IsNil(o.DpSite) {
		var ret DpSiteNestview
		return ret
	}
	return *o.DpSite
}

// GetDpSiteOk returns a tuple with the DpSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetDpSiteOk() (*DpSiteNestview, bool) {
	if o == nil || IsNil(o.DpSite) {
		return nil, false
	}
	return o.DpSite, true
}

// HasDpSite returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasDpSite() bool {
	if o != nil && !IsNil(o.DpSite) {
		return true
	}

	return false
}

// SetDpSite gets a reference to the given DpSiteNestview and assigns it to the DpSite field.
func (o *DpBlockSnapshotRecoveryJob) SetDpSite(v DpSiteNestview) {
	o.DpSite = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetFinishedAt() time.Time {
	if o == nil || IsNil(o.FinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FinishedAt) {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasFinishedAt() bool {
	if o != nil && !IsNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *DpBlockSnapshotRecoveryJob) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DpBlockSnapshotRecoveryJob) SetId(v int64) {
	o.Id = &v
}

// GetMaxRetryTimes returns the MaxRetryTimes field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetMaxRetryTimes() int64 {
	if o == nil || IsNil(o.MaxRetryTimes) {
		var ret int64
		return ret
	}
	return *o.MaxRetryTimes
}

// GetMaxRetryTimesOk returns a tuple with the MaxRetryTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetMaxRetryTimesOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxRetryTimes) {
		return nil, false
	}
	return o.MaxRetryTimes, true
}

// HasMaxRetryTimes returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasMaxRetryTimes() bool {
	if o != nil && !IsNil(o.MaxRetryTimes) {
		return true
	}

	return false
}

// SetMaxRetryTimes gets a reference to the given int64 and assigns it to the MaxRetryTimes field.
func (o *DpBlockSnapshotRecoveryJob) SetMaxRetryTimes(v int64) {
	o.MaxRetryTimes = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetProgress() float64 {
	if o == nil || IsNil(o.Progress) {
		var ret float64
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetProgressOk() (*float64, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given float64 and assigns it to the Progress field.
func (o *DpBlockSnapshotRecoveryJob) SetProgress(v float64) {
	o.Progress = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *DpBlockSnapshotRecoveryJob) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetStartedAt() time.Time {
	if o == nil || IsNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *DpBlockSnapshotRecoveryJob) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DpBlockSnapshotRecoveryJob) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DpBlockSnapshotRecoveryJob) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DpBlockSnapshotRecoveryJob) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockSnapshotRecoveryJob) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DpBlockSnapshotRecoveryJob) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DpBlockSnapshotRecoveryJob) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o DpBlockSnapshotRecoveryJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpBlockSnapshotRecoveryJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackupBlockSnapshot) {
		toSerialize["backup_block_snapshot"] = o.BackupBlockSnapshot
	}
	if !IsNil(o.BackupBlockVolume) {
		toSerialize["backup_block_volume"] = o.BackupBlockVolume
	}
	if !IsNil(o.BackupCluster) {
		toSerialize["backup_cluster"] = o.BackupCluster
	}
	if !IsNil(o.BlockVolume) {
		toSerialize["block_volume"] = o.BlockVolume
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.DpGateway) {
		toSerialize["dp_gateway"] = o.DpGateway
	}
	if !IsNil(o.DpSite) {
		toSerialize["dp_site"] = o.DpSite
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
	if !IsNil(o.ResourceType) {
		toSerialize["resource_type"] = o.ResourceType
	}
	if !IsNil(o.StartedAt) {
		toSerialize["started_at"] = o.StartedAt
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableDpBlockSnapshotRecoveryJob struct {
	value *DpBlockSnapshotRecoveryJob
	isSet bool
}

func (v NullableDpBlockSnapshotRecoveryJob) Get() *DpBlockSnapshotRecoveryJob {
	return v.value
}

func (v *NullableDpBlockSnapshotRecoveryJob) Set(val *DpBlockSnapshotRecoveryJob) {
	v.value = val
	v.isSet = true
}

func (v NullableDpBlockSnapshotRecoveryJob) IsSet() bool {
	return v.isSet
}

func (v *NullableDpBlockSnapshotRecoveryJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpBlockSnapshotRecoveryJob(val *DpBlockSnapshotRecoveryJob) *NullableDpBlockSnapshotRecoveryJob {
	return &NullableDpBlockSnapshotRecoveryJob{value: val, isSet: true}
}

func (v NullableDpBlockSnapshotRecoveryJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpBlockSnapshotRecoveryJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


