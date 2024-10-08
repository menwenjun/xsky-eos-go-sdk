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

// checks if the FSUserGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FSUserGroup{}

// FSUserGroup FSUserGroup defines model of file storage user group
type FSUserGroup struct {
	ActionStatus *string `json:"action_status,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Description *string `json:"description,omitempty"`
	// Not reload default, because it will cause a lot of memory in action log
	FsUsers []FSUserNestview `json:"fs_users,omitempty"`
	GroupId *int64 `json:"group_id,omitempty"`
	HdfsNum *int64 `json:"hdfs_num,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	PolicyFilterNum *int64 `json:"policy_filter_num,omitempty"`
	QuotaNum *int64 `json:"quota_num,omitempty"`
	ShareNums *map[string]int64 `json:"share_nums,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	UserNum *int64 `json:"user_num,omitempty"`
}

// NewFSUserGroup instantiates a new FSUserGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFSUserGroup() *FSUserGroup {
	this := FSUserGroup{}
	return &this
}

// NewFSUserGroupWithDefaults instantiates a new FSUserGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFSUserGroupWithDefaults() *FSUserGroup {
	this := FSUserGroup{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *FSUserGroup) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserGroup) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *FSUserGroup) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *FSUserGroup) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *FSUserGroup) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserGroup) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *FSUserGroup) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *FSUserGroup) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *FSUserGroup) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserGroup) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *FSUserGroup) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *FSUserGroup) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FSUserGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FSUserGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FSUserGroup) SetDescription(v string) {
	o.Description = &v
}

// GetFsUsers returns the FsUsers field value if set, zero value otherwise.
func (o *FSUserGroup) GetFsUsers() []FSUserNestview {
	if o == nil || IsNil(o.FsUsers) {
		var ret []FSUserNestview
		return ret
	}
	return o.FsUsers
}

// GetFsUsersOk returns a tuple with the FsUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserGroup) GetFsUsersOk() ([]FSUserNestview, bool) {
	if o == nil || IsNil(o.FsUsers) {
		return nil, false
	}
	return o.FsUsers, true
}

// HasFsUsers returns a boolean if a field has been set.
func (o *FSUserGroup) HasFsUsers() bool {
	if o != nil && !IsNil(o.FsUsers) {
		return true
	}

	return false
}

// SetFsUsers gets a reference to the given []FSUserNestview and assigns it to the FsUsers field.
func (o *FSUserGroup) SetFsUsers(v []FSUserNestview) {
	o.FsUsers = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *FSUserGroup) GetGroupId() int64 {
	if o == nil || IsNil(o.GroupId) {
		var ret int64
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserGroup) GetGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *FSUserGroup) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given int64 and assigns it to the GroupId field.
func (o *FSUserGroup) SetGroupId(v int64) {
	o.GroupId = &v
}

// GetHdfsNum returns the HdfsNum field value if set, zero value otherwise.
func (o *FSUserGroup) GetHdfsNum() int64 {
	if o == nil || IsNil(o.HdfsNum) {
		var ret int64
		return ret
	}
	return *o.HdfsNum
}

// GetHdfsNumOk returns a tuple with the HdfsNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserGroup) GetHdfsNumOk() (*int64, bool) {
	if o == nil || IsNil(o.HdfsNum) {
		return nil, false
	}
	return o.HdfsNum, true
}

// HasHdfsNum returns a boolean if a field has been set.
func (o *FSUserGroup) HasHdfsNum() bool {
	if o != nil && !IsNil(o.HdfsNum) {
		return true
	}

	return false
}

// SetHdfsNum gets a reference to the given int64 and assigns it to the HdfsNum field.
func (o *FSUserGroup) SetHdfsNum(v int64) {
	o.HdfsNum = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FSUserGroup) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserGroup) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FSUserGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *FSUserGroup) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FSUserGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FSUserGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FSUserGroup) SetName(v string) {
	o.Name = &v
}

// GetPolicyFilterNum returns the PolicyFilterNum field value if set, zero value otherwise.
func (o *FSUserGroup) GetPolicyFilterNum() int64 {
	if o == nil || IsNil(o.PolicyFilterNum) {
		var ret int64
		return ret
	}
	return *o.PolicyFilterNum
}

// GetPolicyFilterNumOk returns a tuple with the PolicyFilterNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserGroup) GetPolicyFilterNumOk() (*int64, bool) {
	if o == nil || IsNil(o.PolicyFilterNum) {
		return nil, false
	}
	return o.PolicyFilterNum, true
}

// HasPolicyFilterNum returns a boolean if a field has been set.
func (o *FSUserGroup) HasPolicyFilterNum() bool {
	if o != nil && !IsNil(o.PolicyFilterNum) {
		return true
	}

	return false
}

// SetPolicyFilterNum gets a reference to the given int64 and assigns it to the PolicyFilterNum field.
func (o *FSUserGroup) SetPolicyFilterNum(v int64) {
	o.PolicyFilterNum = &v
}

// GetQuotaNum returns the QuotaNum field value if set, zero value otherwise.
func (o *FSUserGroup) GetQuotaNum() int64 {
	if o == nil || IsNil(o.QuotaNum) {
		var ret int64
		return ret
	}
	return *o.QuotaNum
}

// GetQuotaNumOk returns a tuple with the QuotaNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserGroup) GetQuotaNumOk() (*int64, bool) {
	if o == nil || IsNil(o.QuotaNum) {
		return nil, false
	}
	return o.QuotaNum, true
}

// HasQuotaNum returns a boolean if a field has been set.
func (o *FSUserGroup) HasQuotaNum() bool {
	if o != nil && !IsNil(o.QuotaNum) {
		return true
	}

	return false
}

// SetQuotaNum gets a reference to the given int64 and assigns it to the QuotaNum field.
func (o *FSUserGroup) SetQuotaNum(v int64) {
	o.QuotaNum = &v
}

// GetShareNums returns the ShareNums field value if set, zero value otherwise.
func (o *FSUserGroup) GetShareNums() map[string]int64 {
	if o == nil || IsNil(o.ShareNums) {
		var ret map[string]int64
		return ret
	}
	return *o.ShareNums
}

// GetShareNumsOk returns a tuple with the ShareNums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserGroup) GetShareNumsOk() (*map[string]int64, bool) {
	if o == nil || IsNil(o.ShareNums) {
		return nil, false
	}
	return o.ShareNums, true
}

// HasShareNums returns a boolean if a field has been set.
func (o *FSUserGroup) HasShareNums() bool {
	if o != nil && !IsNil(o.ShareNums) {
		return true
	}

	return false
}

// SetShareNums gets a reference to the given map[string]int64 and assigns it to the ShareNums field.
func (o *FSUserGroup) SetShareNums(v map[string]int64) {
	o.ShareNums = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *FSUserGroup) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserGroup) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *FSUserGroup) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *FSUserGroup) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetUserNum returns the UserNum field value if set, zero value otherwise.
func (o *FSUserGroup) GetUserNum() int64 {
	if o == nil || IsNil(o.UserNum) {
		var ret int64
		return ret
	}
	return *o.UserNum
}

// GetUserNumOk returns a tuple with the UserNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserGroup) GetUserNumOk() (*int64, bool) {
	if o == nil || IsNil(o.UserNum) {
		return nil, false
	}
	return o.UserNum, true
}

// HasUserNum returns a boolean if a field has been set.
func (o *FSUserGroup) HasUserNum() bool {
	if o != nil && !IsNil(o.UserNum) {
		return true
	}

	return false
}

// SetUserNum gets a reference to the given int64 and assigns it to the UserNum field.
func (o *FSUserGroup) SetUserNum(v int64) {
	o.UserNum = &v
}

func (o FSUserGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FSUserGroup) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FsUsers) {
		toSerialize["fs_users"] = o.FsUsers
	}
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	if !IsNil(o.HdfsNum) {
		toSerialize["hdfs_num"] = o.HdfsNum
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PolicyFilterNum) {
		toSerialize["policy_filter_num"] = o.PolicyFilterNum
	}
	if !IsNil(o.QuotaNum) {
		toSerialize["quota_num"] = o.QuotaNum
	}
	if !IsNil(o.ShareNums) {
		toSerialize["share_nums"] = o.ShareNums
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.UserNum) {
		toSerialize["user_num"] = o.UserNum
	}
	return toSerialize, nil
}

type NullableFSUserGroup struct {
	value *FSUserGroup
	isSet bool
}

func (v NullableFSUserGroup) Get() *FSUserGroup {
	return v.value
}

func (v *NullableFSUserGroup) Set(val *FSUserGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableFSUserGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableFSUserGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFSUserGroup(val *FSUserGroup) *NullableFSUserGroup {
	return &NullableFSUserGroup{value: val, isSet: true}
}

func (v NullableFSUserGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFSUserGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


