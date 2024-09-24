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

// checks if the DfsQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsQuota{}

// DfsQuota DfsQuota defines model of dfs quota
type DfsQuota struct {
	ActionStatus *string `json:"action_status,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	DfsPath *DfsPathNestview `json:"dfs_path,omitempty"`
	DfsRootfs *DfsRootfsNestview `json:"dfs_rootfs,omitempty"`
	DomainUser *DomainUserNestview `json:"domain_user,omitempty"`
	DomainUserGroup *DomainUserGroupNestview `json:"domain_user_group,omitempty"`
	FilesExceedTime *time.Time `json:"files_exceed_time,omitempty"`
	FilesGraceTime *int64 `json:"files_grace_time,omitempty"`
	FilesHardQuota *int64 `json:"files_hard_quota,omitempty"`
	FilesRatio *float32 `json:"files_ratio,omitempty"`
	FilesSoftQuota *int64 `json:"files_soft_quota,omitempty"`
	FilesSuggestQuota *int64 `json:"files_suggest_quota,omitempty"`
	FsUser *FSUserNestview `json:"fs_user,omitempty"`
	FsUserGroup *FSUserGroupNestview `json:"fs_user_group,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Inode *int64 `json:"inode,omitempty"`
	Shared *bool `json:"shared,omitempty"`
	SizeExceedTime *time.Time `json:"size_exceed_time,omitempty"`
	SizeGraceTime *int64 `json:"size_grace_time,omitempty"`
	SizeHardQuota *int64 `json:"size_hard_quota,omitempty"`
	SizeRatio *float32 `json:"size_ratio,omitempty"`
	SizeSoftQuota *int64 `json:"size_soft_quota,omitempty"`
	SizeSuggestQuota *int64 `json:"size_suggest_quota,omitempty"`
	SourceType *string `json:"source_type,omitempty"`
	SourceUuid *string `json:"source_uuid,omitempty"`
	Status *string `json:"status,omitempty"`
	Type *string `json:"type,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewDfsQuota instantiates a new DfsQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsQuota() *DfsQuota {
	this := DfsQuota{}
	return &this
}

// NewDfsQuotaWithDefaults instantiates a new DfsQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsQuotaWithDefaults() *DfsQuota {
	this := DfsQuota{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *DfsQuota) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *DfsQuota) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *DfsQuota) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DfsQuota) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DfsQuota) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *DfsQuota) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DfsQuota) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DfsQuota) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DfsQuota) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDfsPath returns the DfsPath field value if set, zero value otherwise.
func (o *DfsQuota) GetDfsPath() DfsPathNestview {
	if o == nil || IsNil(o.DfsPath) {
		var ret DfsPathNestview
		return ret
	}
	return *o.DfsPath
}

// GetDfsPathOk returns a tuple with the DfsPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetDfsPathOk() (*DfsPathNestview, bool) {
	if o == nil || IsNil(o.DfsPath) {
		return nil, false
	}
	return o.DfsPath, true
}

// HasDfsPath returns a boolean if a field has been set.
func (o *DfsQuota) HasDfsPath() bool {
	if o != nil && !IsNil(o.DfsPath) {
		return true
	}

	return false
}

// SetDfsPath gets a reference to the given DfsPathNestview and assigns it to the DfsPath field.
func (o *DfsQuota) SetDfsPath(v DfsPathNestview) {
	o.DfsPath = &v
}

// GetDfsRootfs returns the DfsRootfs field value if set, zero value otherwise.
func (o *DfsQuota) GetDfsRootfs() DfsRootfsNestview {
	if o == nil || IsNil(o.DfsRootfs) {
		var ret DfsRootfsNestview
		return ret
	}
	return *o.DfsRootfs
}

// GetDfsRootfsOk returns a tuple with the DfsRootfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetDfsRootfsOk() (*DfsRootfsNestview, bool) {
	if o == nil || IsNil(o.DfsRootfs) {
		return nil, false
	}
	return o.DfsRootfs, true
}

// HasDfsRootfs returns a boolean if a field has been set.
func (o *DfsQuota) HasDfsRootfs() bool {
	if o != nil && !IsNil(o.DfsRootfs) {
		return true
	}

	return false
}

// SetDfsRootfs gets a reference to the given DfsRootfsNestview and assigns it to the DfsRootfs field.
func (o *DfsQuota) SetDfsRootfs(v DfsRootfsNestview) {
	o.DfsRootfs = &v
}

// GetDomainUser returns the DomainUser field value if set, zero value otherwise.
func (o *DfsQuota) GetDomainUser() DomainUserNestview {
	if o == nil || IsNil(o.DomainUser) {
		var ret DomainUserNestview
		return ret
	}
	return *o.DomainUser
}

// GetDomainUserOk returns a tuple with the DomainUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetDomainUserOk() (*DomainUserNestview, bool) {
	if o == nil || IsNil(o.DomainUser) {
		return nil, false
	}
	return o.DomainUser, true
}

// HasDomainUser returns a boolean if a field has been set.
func (o *DfsQuota) HasDomainUser() bool {
	if o != nil && !IsNil(o.DomainUser) {
		return true
	}

	return false
}

// SetDomainUser gets a reference to the given DomainUserNestview and assigns it to the DomainUser field.
func (o *DfsQuota) SetDomainUser(v DomainUserNestview) {
	o.DomainUser = &v
}

// GetDomainUserGroup returns the DomainUserGroup field value if set, zero value otherwise.
func (o *DfsQuota) GetDomainUserGroup() DomainUserGroupNestview {
	if o == nil || IsNil(o.DomainUserGroup) {
		var ret DomainUserGroupNestview
		return ret
	}
	return *o.DomainUserGroup
}

// GetDomainUserGroupOk returns a tuple with the DomainUserGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetDomainUserGroupOk() (*DomainUserGroupNestview, bool) {
	if o == nil || IsNil(o.DomainUserGroup) {
		return nil, false
	}
	return o.DomainUserGroup, true
}

// HasDomainUserGroup returns a boolean if a field has been set.
func (o *DfsQuota) HasDomainUserGroup() bool {
	if o != nil && !IsNil(o.DomainUserGroup) {
		return true
	}

	return false
}

// SetDomainUserGroup gets a reference to the given DomainUserGroupNestview and assigns it to the DomainUserGroup field.
func (o *DfsQuota) SetDomainUserGroup(v DomainUserGroupNestview) {
	o.DomainUserGroup = &v
}

// GetFilesExceedTime returns the FilesExceedTime field value if set, zero value otherwise.
func (o *DfsQuota) GetFilesExceedTime() time.Time {
	if o == nil || IsNil(o.FilesExceedTime) {
		var ret time.Time
		return ret
	}
	return *o.FilesExceedTime
}

// GetFilesExceedTimeOk returns a tuple with the FilesExceedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetFilesExceedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FilesExceedTime) {
		return nil, false
	}
	return o.FilesExceedTime, true
}

// HasFilesExceedTime returns a boolean if a field has been set.
func (o *DfsQuota) HasFilesExceedTime() bool {
	if o != nil && !IsNil(o.FilesExceedTime) {
		return true
	}

	return false
}

// SetFilesExceedTime gets a reference to the given time.Time and assigns it to the FilesExceedTime field.
func (o *DfsQuota) SetFilesExceedTime(v time.Time) {
	o.FilesExceedTime = &v
}

// GetFilesGraceTime returns the FilesGraceTime field value if set, zero value otherwise.
func (o *DfsQuota) GetFilesGraceTime() int64 {
	if o == nil || IsNil(o.FilesGraceTime) {
		var ret int64
		return ret
	}
	return *o.FilesGraceTime
}

// GetFilesGraceTimeOk returns a tuple with the FilesGraceTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetFilesGraceTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.FilesGraceTime) {
		return nil, false
	}
	return o.FilesGraceTime, true
}

// HasFilesGraceTime returns a boolean if a field has been set.
func (o *DfsQuota) HasFilesGraceTime() bool {
	if o != nil && !IsNil(o.FilesGraceTime) {
		return true
	}

	return false
}

// SetFilesGraceTime gets a reference to the given int64 and assigns it to the FilesGraceTime field.
func (o *DfsQuota) SetFilesGraceTime(v int64) {
	o.FilesGraceTime = &v
}

// GetFilesHardQuota returns the FilesHardQuota field value if set, zero value otherwise.
func (o *DfsQuota) GetFilesHardQuota() int64 {
	if o == nil || IsNil(o.FilesHardQuota) {
		var ret int64
		return ret
	}
	return *o.FilesHardQuota
}

// GetFilesHardQuotaOk returns a tuple with the FilesHardQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetFilesHardQuotaOk() (*int64, bool) {
	if o == nil || IsNil(o.FilesHardQuota) {
		return nil, false
	}
	return o.FilesHardQuota, true
}

// HasFilesHardQuota returns a boolean if a field has been set.
func (o *DfsQuota) HasFilesHardQuota() bool {
	if o != nil && !IsNil(o.FilesHardQuota) {
		return true
	}

	return false
}

// SetFilesHardQuota gets a reference to the given int64 and assigns it to the FilesHardQuota field.
func (o *DfsQuota) SetFilesHardQuota(v int64) {
	o.FilesHardQuota = &v
}

// GetFilesRatio returns the FilesRatio field value if set, zero value otherwise.
func (o *DfsQuota) GetFilesRatio() float32 {
	if o == nil || IsNil(o.FilesRatio) {
		var ret float32
		return ret
	}
	return *o.FilesRatio
}

// GetFilesRatioOk returns a tuple with the FilesRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetFilesRatioOk() (*float32, bool) {
	if o == nil || IsNil(o.FilesRatio) {
		return nil, false
	}
	return o.FilesRatio, true
}

// HasFilesRatio returns a boolean if a field has been set.
func (o *DfsQuota) HasFilesRatio() bool {
	if o != nil && !IsNil(o.FilesRatio) {
		return true
	}

	return false
}

// SetFilesRatio gets a reference to the given float32 and assigns it to the FilesRatio field.
func (o *DfsQuota) SetFilesRatio(v float32) {
	o.FilesRatio = &v
}

// GetFilesSoftQuota returns the FilesSoftQuota field value if set, zero value otherwise.
func (o *DfsQuota) GetFilesSoftQuota() int64 {
	if o == nil || IsNil(o.FilesSoftQuota) {
		var ret int64
		return ret
	}
	return *o.FilesSoftQuota
}

// GetFilesSoftQuotaOk returns a tuple with the FilesSoftQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetFilesSoftQuotaOk() (*int64, bool) {
	if o == nil || IsNil(o.FilesSoftQuota) {
		return nil, false
	}
	return o.FilesSoftQuota, true
}

// HasFilesSoftQuota returns a boolean if a field has been set.
func (o *DfsQuota) HasFilesSoftQuota() bool {
	if o != nil && !IsNil(o.FilesSoftQuota) {
		return true
	}

	return false
}

// SetFilesSoftQuota gets a reference to the given int64 and assigns it to the FilesSoftQuota field.
func (o *DfsQuota) SetFilesSoftQuota(v int64) {
	o.FilesSoftQuota = &v
}

// GetFilesSuggestQuota returns the FilesSuggestQuota field value if set, zero value otherwise.
func (o *DfsQuota) GetFilesSuggestQuota() int64 {
	if o == nil || IsNil(o.FilesSuggestQuota) {
		var ret int64
		return ret
	}
	return *o.FilesSuggestQuota
}

// GetFilesSuggestQuotaOk returns a tuple with the FilesSuggestQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetFilesSuggestQuotaOk() (*int64, bool) {
	if o == nil || IsNil(o.FilesSuggestQuota) {
		return nil, false
	}
	return o.FilesSuggestQuota, true
}

// HasFilesSuggestQuota returns a boolean if a field has been set.
func (o *DfsQuota) HasFilesSuggestQuota() bool {
	if o != nil && !IsNil(o.FilesSuggestQuota) {
		return true
	}

	return false
}

// SetFilesSuggestQuota gets a reference to the given int64 and assigns it to the FilesSuggestQuota field.
func (o *DfsQuota) SetFilesSuggestQuota(v int64) {
	o.FilesSuggestQuota = &v
}

// GetFsUser returns the FsUser field value if set, zero value otherwise.
func (o *DfsQuota) GetFsUser() FSUserNestview {
	if o == nil || IsNil(o.FsUser) {
		var ret FSUserNestview
		return ret
	}
	return *o.FsUser
}

// GetFsUserOk returns a tuple with the FsUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetFsUserOk() (*FSUserNestview, bool) {
	if o == nil || IsNil(o.FsUser) {
		return nil, false
	}
	return o.FsUser, true
}

// HasFsUser returns a boolean if a field has been set.
func (o *DfsQuota) HasFsUser() bool {
	if o != nil && !IsNil(o.FsUser) {
		return true
	}

	return false
}

// SetFsUser gets a reference to the given FSUserNestview and assigns it to the FsUser field.
func (o *DfsQuota) SetFsUser(v FSUserNestview) {
	o.FsUser = &v
}

// GetFsUserGroup returns the FsUserGroup field value if set, zero value otherwise.
func (o *DfsQuota) GetFsUserGroup() FSUserGroupNestview {
	if o == nil || IsNil(o.FsUserGroup) {
		var ret FSUserGroupNestview
		return ret
	}
	return *o.FsUserGroup
}

// GetFsUserGroupOk returns a tuple with the FsUserGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetFsUserGroupOk() (*FSUserGroupNestview, bool) {
	if o == nil || IsNil(o.FsUserGroup) {
		return nil, false
	}
	return o.FsUserGroup, true
}

// HasFsUserGroup returns a boolean if a field has been set.
func (o *DfsQuota) HasFsUserGroup() bool {
	if o != nil && !IsNil(o.FsUserGroup) {
		return true
	}

	return false
}

// SetFsUserGroup gets a reference to the given FSUserGroupNestview and assigns it to the FsUserGroup field.
func (o *DfsQuota) SetFsUserGroup(v FSUserGroupNestview) {
	o.FsUserGroup = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsQuota) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsQuota) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsQuota) SetId(v int64) {
	o.Id = &v
}

// GetInode returns the Inode field value if set, zero value otherwise.
func (o *DfsQuota) GetInode() int64 {
	if o == nil || IsNil(o.Inode) {
		var ret int64
		return ret
	}
	return *o.Inode
}

// GetInodeOk returns a tuple with the Inode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetInodeOk() (*int64, bool) {
	if o == nil || IsNil(o.Inode) {
		return nil, false
	}
	return o.Inode, true
}

// HasInode returns a boolean if a field has been set.
func (o *DfsQuota) HasInode() bool {
	if o != nil && !IsNil(o.Inode) {
		return true
	}

	return false
}

// SetInode gets a reference to the given int64 and assigns it to the Inode field.
func (o *DfsQuota) SetInode(v int64) {
	o.Inode = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *DfsQuota) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *DfsQuota) HasShared() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *DfsQuota) SetShared(v bool) {
	o.Shared = &v
}

// GetSizeExceedTime returns the SizeExceedTime field value if set, zero value otherwise.
func (o *DfsQuota) GetSizeExceedTime() time.Time {
	if o == nil || IsNil(o.SizeExceedTime) {
		var ret time.Time
		return ret
	}
	return *o.SizeExceedTime
}

// GetSizeExceedTimeOk returns a tuple with the SizeExceedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetSizeExceedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SizeExceedTime) {
		return nil, false
	}
	return o.SizeExceedTime, true
}

// HasSizeExceedTime returns a boolean if a field has been set.
func (o *DfsQuota) HasSizeExceedTime() bool {
	if o != nil && !IsNil(o.SizeExceedTime) {
		return true
	}

	return false
}

// SetSizeExceedTime gets a reference to the given time.Time and assigns it to the SizeExceedTime field.
func (o *DfsQuota) SetSizeExceedTime(v time.Time) {
	o.SizeExceedTime = &v
}

// GetSizeGraceTime returns the SizeGraceTime field value if set, zero value otherwise.
func (o *DfsQuota) GetSizeGraceTime() int64 {
	if o == nil || IsNil(o.SizeGraceTime) {
		var ret int64
		return ret
	}
	return *o.SizeGraceTime
}

// GetSizeGraceTimeOk returns a tuple with the SizeGraceTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetSizeGraceTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.SizeGraceTime) {
		return nil, false
	}
	return o.SizeGraceTime, true
}

// HasSizeGraceTime returns a boolean if a field has been set.
func (o *DfsQuota) HasSizeGraceTime() bool {
	if o != nil && !IsNil(o.SizeGraceTime) {
		return true
	}

	return false
}

// SetSizeGraceTime gets a reference to the given int64 and assigns it to the SizeGraceTime field.
func (o *DfsQuota) SetSizeGraceTime(v int64) {
	o.SizeGraceTime = &v
}

// GetSizeHardQuota returns the SizeHardQuota field value if set, zero value otherwise.
func (o *DfsQuota) GetSizeHardQuota() int64 {
	if o == nil || IsNil(o.SizeHardQuota) {
		var ret int64
		return ret
	}
	return *o.SizeHardQuota
}

// GetSizeHardQuotaOk returns a tuple with the SizeHardQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetSizeHardQuotaOk() (*int64, bool) {
	if o == nil || IsNil(o.SizeHardQuota) {
		return nil, false
	}
	return o.SizeHardQuota, true
}

// HasSizeHardQuota returns a boolean if a field has been set.
func (o *DfsQuota) HasSizeHardQuota() bool {
	if o != nil && !IsNil(o.SizeHardQuota) {
		return true
	}

	return false
}

// SetSizeHardQuota gets a reference to the given int64 and assigns it to the SizeHardQuota field.
func (o *DfsQuota) SetSizeHardQuota(v int64) {
	o.SizeHardQuota = &v
}

// GetSizeRatio returns the SizeRatio field value if set, zero value otherwise.
func (o *DfsQuota) GetSizeRatio() float32 {
	if o == nil || IsNil(o.SizeRatio) {
		var ret float32
		return ret
	}
	return *o.SizeRatio
}

// GetSizeRatioOk returns a tuple with the SizeRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetSizeRatioOk() (*float32, bool) {
	if o == nil || IsNil(o.SizeRatio) {
		return nil, false
	}
	return o.SizeRatio, true
}

// HasSizeRatio returns a boolean if a field has been set.
func (o *DfsQuota) HasSizeRatio() bool {
	if o != nil && !IsNil(o.SizeRatio) {
		return true
	}

	return false
}

// SetSizeRatio gets a reference to the given float32 and assigns it to the SizeRatio field.
func (o *DfsQuota) SetSizeRatio(v float32) {
	o.SizeRatio = &v
}

// GetSizeSoftQuota returns the SizeSoftQuota field value if set, zero value otherwise.
func (o *DfsQuota) GetSizeSoftQuota() int64 {
	if o == nil || IsNil(o.SizeSoftQuota) {
		var ret int64
		return ret
	}
	return *o.SizeSoftQuota
}

// GetSizeSoftQuotaOk returns a tuple with the SizeSoftQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetSizeSoftQuotaOk() (*int64, bool) {
	if o == nil || IsNil(o.SizeSoftQuota) {
		return nil, false
	}
	return o.SizeSoftQuota, true
}

// HasSizeSoftQuota returns a boolean if a field has been set.
func (o *DfsQuota) HasSizeSoftQuota() bool {
	if o != nil && !IsNil(o.SizeSoftQuota) {
		return true
	}

	return false
}

// SetSizeSoftQuota gets a reference to the given int64 and assigns it to the SizeSoftQuota field.
func (o *DfsQuota) SetSizeSoftQuota(v int64) {
	o.SizeSoftQuota = &v
}

// GetSizeSuggestQuota returns the SizeSuggestQuota field value if set, zero value otherwise.
func (o *DfsQuota) GetSizeSuggestQuota() int64 {
	if o == nil || IsNil(o.SizeSuggestQuota) {
		var ret int64
		return ret
	}
	return *o.SizeSuggestQuota
}

// GetSizeSuggestQuotaOk returns a tuple with the SizeSuggestQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetSizeSuggestQuotaOk() (*int64, bool) {
	if o == nil || IsNil(o.SizeSuggestQuota) {
		return nil, false
	}
	return o.SizeSuggestQuota, true
}

// HasSizeSuggestQuota returns a boolean if a field has been set.
func (o *DfsQuota) HasSizeSuggestQuota() bool {
	if o != nil && !IsNil(o.SizeSuggestQuota) {
		return true
	}

	return false
}

// SetSizeSuggestQuota gets a reference to the given int64 and assigns it to the SizeSuggestQuota field.
func (o *DfsQuota) SetSizeSuggestQuota(v int64) {
	o.SizeSuggestQuota = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *DfsQuota) GetSourceType() string {
	if o == nil || IsNil(o.SourceType) {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceType) {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *DfsQuota) HasSourceType() bool {
	if o != nil && !IsNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *DfsQuota) SetSourceType(v string) {
	o.SourceType = &v
}

// GetSourceUuid returns the SourceUuid field value if set, zero value otherwise.
func (o *DfsQuota) GetSourceUuid() string {
	if o == nil || IsNil(o.SourceUuid) {
		var ret string
		return ret
	}
	return *o.SourceUuid
}

// GetSourceUuidOk returns a tuple with the SourceUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetSourceUuidOk() (*string, bool) {
	if o == nil || IsNil(o.SourceUuid) {
		return nil, false
	}
	return o.SourceUuid, true
}

// HasSourceUuid returns a boolean if a field has been set.
func (o *DfsQuota) HasSourceUuid() bool {
	if o != nil && !IsNil(o.SourceUuid) {
		return true
	}

	return false
}

// SetSourceUuid gets a reference to the given string and assigns it to the SourceUuid field.
func (o *DfsQuota) SetSourceUuid(v string) {
	o.SourceUuid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DfsQuota) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DfsQuota) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DfsQuota) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DfsQuota) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DfsQuota) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DfsQuota) SetType(v string) {
	o.Type = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DfsQuota) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuota) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DfsQuota) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DfsQuota) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o DfsQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionStatus) {
		toSerialize["action_status"] = o.ActionStatus
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.DfsPath) {
		toSerialize["dfs_path"] = o.DfsPath
	}
	if !IsNil(o.DfsRootfs) {
		toSerialize["dfs_rootfs"] = o.DfsRootfs
	}
	if !IsNil(o.DomainUser) {
		toSerialize["domain_user"] = o.DomainUser
	}
	if !IsNil(o.DomainUserGroup) {
		toSerialize["domain_user_group"] = o.DomainUserGroup
	}
	if !IsNil(o.FilesExceedTime) {
		toSerialize["files_exceed_time"] = o.FilesExceedTime
	}
	if !IsNil(o.FilesGraceTime) {
		toSerialize["files_grace_time"] = o.FilesGraceTime
	}
	if !IsNil(o.FilesHardQuota) {
		toSerialize["files_hard_quota"] = o.FilesHardQuota
	}
	if !IsNil(o.FilesRatio) {
		toSerialize["files_ratio"] = o.FilesRatio
	}
	if !IsNil(o.FilesSoftQuota) {
		toSerialize["files_soft_quota"] = o.FilesSoftQuota
	}
	if !IsNil(o.FilesSuggestQuota) {
		toSerialize["files_suggest_quota"] = o.FilesSuggestQuota
	}
	if !IsNil(o.FsUser) {
		toSerialize["fs_user"] = o.FsUser
	}
	if !IsNil(o.FsUserGroup) {
		toSerialize["fs_user_group"] = o.FsUserGroup
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Inode) {
		toSerialize["inode"] = o.Inode
	}
	if !IsNil(o.Shared) {
		toSerialize["shared"] = o.Shared
	}
	if !IsNil(o.SizeExceedTime) {
		toSerialize["size_exceed_time"] = o.SizeExceedTime
	}
	if !IsNil(o.SizeGraceTime) {
		toSerialize["size_grace_time"] = o.SizeGraceTime
	}
	if !IsNil(o.SizeHardQuota) {
		toSerialize["size_hard_quota"] = o.SizeHardQuota
	}
	if !IsNil(o.SizeRatio) {
		toSerialize["size_ratio"] = o.SizeRatio
	}
	if !IsNil(o.SizeSoftQuota) {
		toSerialize["size_soft_quota"] = o.SizeSoftQuota
	}
	if !IsNil(o.SizeSuggestQuota) {
		toSerialize["size_suggest_quota"] = o.SizeSuggestQuota
	}
	if !IsNil(o.SourceType) {
		toSerialize["source_type"] = o.SourceType
	}
	if !IsNil(o.SourceUuid) {
		toSerialize["source_uuid"] = o.SourceUuid
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullableDfsQuota struct {
	value *DfsQuota
	isSet bool
}

func (v NullableDfsQuota) Get() *DfsQuota {
	return v.value
}

func (v *NullableDfsQuota) Set(val *DfsQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsQuota(val *DfsQuota) *NullableDfsQuota {
	return &NullableDfsQuota{value: val, isSet: true}
}

func (v NullableDfsQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


