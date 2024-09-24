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

// checks if the DfsPolicyFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsPolicyFilter{}

// DfsPolicyFilter DfsPolicyFilter defines filter for DfsStoragePolicy
type DfsPolicyFilter struct {
	AdDomainUserGroups []DomainUserGroup `json:"ad_domain_user_groups,omitempty"`
	AdDomainUsers []DomainUser `json:"ad_domain_users,omitempty"`
	AdUserGroupIds []int64 `json:"ad_user_group_ids,omitempty"`
	AdUserIds []int64 `json:"ad_user_ids,omitempty"`
	DfsStoragePolicy *DfsStoragePolicyNestview `json:"dfs_storage_policy,omitempty"`
	ExcludeFileName *bool `json:"exclude_file_name,omitempty"`
	ExcludeUserName *bool `json:"exclude_user_name,omitempty"`
	FileNames *string `json:"file_names,omitempty"`
	Id *int64 `json:"id,omitempty"`
	LdapDomainUserGroups []DomainUserGroup `json:"ldap_domain_user_groups,omitempty"`
	LdapDomainUsers []DomainUser `json:"ldap_domain_users,omitempty"`
	LdapUserGroupIds []int64 `json:"ldap_user_group_ids,omitempty"`
	LdapUserIds []int64 `json:"ldap_user_ids,omitempty"`
	UserGroupNames []string `json:"user_group_names,omitempty"`
	UserNames []string `json:"user_names,omitempty"`
}

// NewDfsPolicyFilter instantiates a new DfsPolicyFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsPolicyFilter() *DfsPolicyFilter {
	this := DfsPolicyFilter{}
	return &this
}

// NewDfsPolicyFilterWithDefaults instantiates a new DfsPolicyFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsPolicyFilterWithDefaults() *DfsPolicyFilter {
	this := DfsPolicyFilter{}
	return &this
}

// GetAdDomainUserGroups returns the AdDomainUserGroups field value if set, zero value otherwise.
func (o *DfsPolicyFilter) GetAdDomainUserGroups() []DomainUserGroup {
	if o == nil || IsNil(o.AdDomainUserGroups) {
		var ret []DomainUserGroup
		return ret
	}
	return o.AdDomainUserGroups
}

// GetAdDomainUserGroupsOk returns a tuple with the AdDomainUserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPolicyFilter) GetAdDomainUserGroupsOk() ([]DomainUserGroup, bool) {
	if o == nil || IsNil(o.AdDomainUserGroups) {
		return nil, false
	}
	return o.AdDomainUserGroups, true
}

// HasAdDomainUserGroups returns a boolean if a field has been set.
func (o *DfsPolicyFilter) HasAdDomainUserGroups() bool {
	if o != nil && !IsNil(o.AdDomainUserGroups) {
		return true
	}

	return false
}

// SetAdDomainUserGroups gets a reference to the given []DomainUserGroup and assigns it to the AdDomainUserGroups field.
func (o *DfsPolicyFilter) SetAdDomainUserGroups(v []DomainUserGroup) {
	o.AdDomainUserGroups = v
}

// GetAdDomainUsers returns the AdDomainUsers field value if set, zero value otherwise.
func (o *DfsPolicyFilter) GetAdDomainUsers() []DomainUser {
	if o == nil || IsNil(o.AdDomainUsers) {
		var ret []DomainUser
		return ret
	}
	return o.AdDomainUsers
}

// GetAdDomainUsersOk returns a tuple with the AdDomainUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPolicyFilter) GetAdDomainUsersOk() ([]DomainUser, bool) {
	if o == nil || IsNil(o.AdDomainUsers) {
		return nil, false
	}
	return o.AdDomainUsers, true
}

// HasAdDomainUsers returns a boolean if a field has been set.
func (o *DfsPolicyFilter) HasAdDomainUsers() bool {
	if o != nil && !IsNil(o.AdDomainUsers) {
		return true
	}

	return false
}

// SetAdDomainUsers gets a reference to the given []DomainUser and assigns it to the AdDomainUsers field.
func (o *DfsPolicyFilter) SetAdDomainUsers(v []DomainUser) {
	o.AdDomainUsers = v
}

// GetAdUserGroupIds returns the AdUserGroupIds field value if set, zero value otherwise.
func (o *DfsPolicyFilter) GetAdUserGroupIds() []int64 {
	if o == nil || IsNil(o.AdUserGroupIds) {
		var ret []int64
		return ret
	}
	return o.AdUserGroupIds
}

// GetAdUserGroupIdsOk returns a tuple with the AdUserGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPolicyFilter) GetAdUserGroupIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.AdUserGroupIds) {
		return nil, false
	}
	return o.AdUserGroupIds, true
}

// HasAdUserGroupIds returns a boolean if a field has been set.
func (o *DfsPolicyFilter) HasAdUserGroupIds() bool {
	if o != nil && !IsNil(o.AdUserGroupIds) {
		return true
	}

	return false
}

// SetAdUserGroupIds gets a reference to the given []int64 and assigns it to the AdUserGroupIds field.
func (o *DfsPolicyFilter) SetAdUserGroupIds(v []int64) {
	o.AdUserGroupIds = v
}

// GetAdUserIds returns the AdUserIds field value if set, zero value otherwise.
func (o *DfsPolicyFilter) GetAdUserIds() []int64 {
	if o == nil || IsNil(o.AdUserIds) {
		var ret []int64
		return ret
	}
	return o.AdUserIds
}

// GetAdUserIdsOk returns a tuple with the AdUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPolicyFilter) GetAdUserIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.AdUserIds) {
		return nil, false
	}
	return o.AdUserIds, true
}

// HasAdUserIds returns a boolean if a field has been set.
func (o *DfsPolicyFilter) HasAdUserIds() bool {
	if o != nil && !IsNil(o.AdUserIds) {
		return true
	}

	return false
}

// SetAdUserIds gets a reference to the given []int64 and assigns it to the AdUserIds field.
func (o *DfsPolicyFilter) SetAdUserIds(v []int64) {
	o.AdUserIds = v
}

// GetDfsStoragePolicy returns the DfsStoragePolicy field value if set, zero value otherwise.
func (o *DfsPolicyFilter) GetDfsStoragePolicy() DfsStoragePolicyNestview {
	if o == nil || IsNil(o.DfsStoragePolicy) {
		var ret DfsStoragePolicyNestview
		return ret
	}
	return *o.DfsStoragePolicy
}

// GetDfsStoragePolicyOk returns a tuple with the DfsStoragePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPolicyFilter) GetDfsStoragePolicyOk() (*DfsStoragePolicyNestview, bool) {
	if o == nil || IsNil(o.DfsStoragePolicy) {
		return nil, false
	}
	return o.DfsStoragePolicy, true
}

// HasDfsStoragePolicy returns a boolean if a field has been set.
func (o *DfsPolicyFilter) HasDfsStoragePolicy() bool {
	if o != nil && !IsNil(o.DfsStoragePolicy) {
		return true
	}

	return false
}

// SetDfsStoragePolicy gets a reference to the given DfsStoragePolicyNestview and assigns it to the DfsStoragePolicy field.
func (o *DfsPolicyFilter) SetDfsStoragePolicy(v DfsStoragePolicyNestview) {
	o.DfsStoragePolicy = &v
}

// GetExcludeFileName returns the ExcludeFileName field value if set, zero value otherwise.
func (o *DfsPolicyFilter) GetExcludeFileName() bool {
	if o == nil || IsNil(o.ExcludeFileName) {
		var ret bool
		return ret
	}
	return *o.ExcludeFileName
}

// GetExcludeFileNameOk returns a tuple with the ExcludeFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPolicyFilter) GetExcludeFileNameOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeFileName) {
		return nil, false
	}
	return o.ExcludeFileName, true
}

// HasExcludeFileName returns a boolean if a field has been set.
func (o *DfsPolicyFilter) HasExcludeFileName() bool {
	if o != nil && !IsNil(o.ExcludeFileName) {
		return true
	}

	return false
}

// SetExcludeFileName gets a reference to the given bool and assigns it to the ExcludeFileName field.
func (o *DfsPolicyFilter) SetExcludeFileName(v bool) {
	o.ExcludeFileName = &v
}

// GetExcludeUserName returns the ExcludeUserName field value if set, zero value otherwise.
func (o *DfsPolicyFilter) GetExcludeUserName() bool {
	if o == nil || IsNil(o.ExcludeUserName) {
		var ret bool
		return ret
	}
	return *o.ExcludeUserName
}

// GetExcludeUserNameOk returns a tuple with the ExcludeUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPolicyFilter) GetExcludeUserNameOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeUserName) {
		return nil, false
	}
	return o.ExcludeUserName, true
}

// HasExcludeUserName returns a boolean if a field has been set.
func (o *DfsPolicyFilter) HasExcludeUserName() bool {
	if o != nil && !IsNil(o.ExcludeUserName) {
		return true
	}

	return false
}

// SetExcludeUserName gets a reference to the given bool and assigns it to the ExcludeUserName field.
func (o *DfsPolicyFilter) SetExcludeUserName(v bool) {
	o.ExcludeUserName = &v
}

// GetFileNames returns the FileNames field value if set, zero value otherwise.
func (o *DfsPolicyFilter) GetFileNames() string {
	if o == nil || IsNil(o.FileNames) {
		var ret string
		return ret
	}
	return *o.FileNames
}

// GetFileNamesOk returns a tuple with the FileNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPolicyFilter) GetFileNamesOk() (*string, bool) {
	if o == nil || IsNil(o.FileNames) {
		return nil, false
	}
	return o.FileNames, true
}

// HasFileNames returns a boolean if a field has been set.
func (o *DfsPolicyFilter) HasFileNames() bool {
	if o != nil && !IsNil(o.FileNames) {
		return true
	}

	return false
}

// SetFileNames gets a reference to the given string and assigns it to the FileNames field.
func (o *DfsPolicyFilter) SetFileNames(v string) {
	o.FileNames = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsPolicyFilter) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPolicyFilter) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsPolicyFilter) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsPolicyFilter) SetId(v int64) {
	o.Id = &v
}

// GetLdapDomainUserGroups returns the LdapDomainUserGroups field value if set, zero value otherwise.
func (o *DfsPolicyFilter) GetLdapDomainUserGroups() []DomainUserGroup {
	if o == nil || IsNil(o.LdapDomainUserGroups) {
		var ret []DomainUserGroup
		return ret
	}
	return o.LdapDomainUserGroups
}

// GetLdapDomainUserGroupsOk returns a tuple with the LdapDomainUserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPolicyFilter) GetLdapDomainUserGroupsOk() ([]DomainUserGroup, bool) {
	if o == nil || IsNil(o.LdapDomainUserGroups) {
		return nil, false
	}
	return o.LdapDomainUserGroups, true
}

// HasLdapDomainUserGroups returns a boolean if a field has been set.
func (o *DfsPolicyFilter) HasLdapDomainUserGroups() bool {
	if o != nil && !IsNil(o.LdapDomainUserGroups) {
		return true
	}

	return false
}

// SetLdapDomainUserGroups gets a reference to the given []DomainUserGroup and assigns it to the LdapDomainUserGroups field.
func (o *DfsPolicyFilter) SetLdapDomainUserGroups(v []DomainUserGroup) {
	o.LdapDomainUserGroups = v
}

// GetLdapDomainUsers returns the LdapDomainUsers field value if set, zero value otherwise.
func (o *DfsPolicyFilter) GetLdapDomainUsers() []DomainUser {
	if o == nil || IsNil(o.LdapDomainUsers) {
		var ret []DomainUser
		return ret
	}
	return o.LdapDomainUsers
}

// GetLdapDomainUsersOk returns a tuple with the LdapDomainUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPolicyFilter) GetLdapDomainUsersOk() ([]DomainUser, bool) {
	if o == nil || IsNil(o.LdapDomainUsers) {
		return nil, false
	}
	return o.LdapDomainUsers, true
}

// HasLdapDomainUsers returns a boolean if a field has been set.
func (o *DfsPolicyFilter) HasLdapDomainUsers() bool {
	if o != nil && !IsNil(o.LdapDomainUsers) {
		return true
	}

	return false
}

// SetLdapDomainUsers gets a reference to the given []DomainUser and assigns it to the LdapDomainUsers field.
func (o *DfsPolicyFilter) SetLdapDomainUsers(v []DomainUser) {
	o.LdapDomainUsers = v
}

// GetLdapUserGroupIds returns the LdapUserGroupIds field value if set, zero value otherwise.
func (o *DfsPolicyFilter) GetLdapUserGroupIds() []int64 {
	if o == nil || IsNil(o.LdapUserGroupIds) {
		var ret []int64
		return ret
	}
	return o.LdapUserGroupIds
}

// GetLdapUserGroupIdsOk returns a tuple with the LdapUserGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPolicyFilter) GetLdapUserGroupIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.LdapUserGroupIds) {
		return nil, false
	}
	return o.LdapUserGroupIds, true
}

// HasLdapUserGroupIds returns a boolean if a field has been set.
func (o *DfsPolicyFilter) HasLdapUserGroupIds() bool {
	if o != nil && !IsNil(o.LdapUserGroupIds) {
		return true
	}

	return false
}

// SetLdapUserGroupIds gets a reference to the given []int64 and assigns it to the LdapUserGroupIds field.
func (o *DfsPolicyFilter) SetLdapUserGroupIds(v []int64) {
	o.LdapUserGroupIds = v
}

// GetLdapUserIds returns the LdapUserIds field value if set, zero value otherwise.
func (o *DfsPolicyFilter) GetLdapUserIds() []int64 {
	if o == nil || IsNil(o.LdapUserIds) {
		var ret []int64
		return ret
	}
	return o.LdapUserIds
}

// GetLdapUserIdsOk returns a tuple with the LdapUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPolicyFilter) GetLdapUserIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.LdapUserIds) {
		return nil, false
	}
	return o.LdapUserIds, true
}

// HasLdapUserIds returns a boolean if a field has been set.
func (o *DfsPolicyFilter) HasLdapUserIds() bool {
	if o != nil && !IsNil(o.LdapUserIds) {
		return true
	}

	return false
}

// SetLdapUserIds gets a reference to the given []int64 and assigns it to the LdapUserIds field.
func (o *DfsPolicyFilter) SetLdapUserIds(v []int64) {
	o.LdapUserIds = v
}

// GetUserGroupNames returns the UserGroupNames field value if set, zero value otherwise.
func (o *DfsPolicyFilter) GetUserGroupNames() []string {
	if o == nil || IsNil(o.UserGroupNames) {
		var ret []string
		return ret
	}
	return o.UserGroupNames
}

// GetUserGroupNamesOk returns a tuple with the UserGroupNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPolicyFilter) GetUserGroupNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.UserGroupNames) {
		return nil, false
	}
	return o.UserGroupNames, true
}

// HasUserGroupNames returns a boolean if a field has been set.
func (o *DfsPolicyFilter) HasUserGroupNames() bool {
	if o != nil && !IsNil(o.UserGroupNames) {
		return true
	}

	return false
}

// SetUserGroupNames gets a reference to the given []string and assigns it to the UserGroupNames field.
func (o *DfsPolicyFilter) SetUserGroupNames(v []string) {
	o.UserGroupNames = v
}

// GetUserNames returns the UserNames field value if set, zero value otherwise.
func (o *DfsPolicyFilter) GetUserNames() []string {
	if o == nil || IsNil(o.UserNames) {
		var ret []string
		return ret
	}
	return o.UserNames
}

// GetUserNamesOk returns a tuple with the UserNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPolicyFilter) GetUserNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.UserNames) {
		return nil, false
	}
	return o.UserNames, true
}

// HasUserNames returns a boolean if a field has been set.
func (o *DfsPolicyFilter) HasUserNames() bool {
	if o != nil && !IsNil(o.UserNames) {
		return true
	}

	return false
}

// SetUserNames gets a reference to the given []string and assigns it to the UserNames field.
func (o *DfsPolicyFilter) SetUserNames(v []string) {
	o.UserNames = v
}

func (o DfsPolicyFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsPolicyFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdDomainUserGroups) {
		toSerialize["ad_domain_user_groups"] = o.AdDomainUserGroups
	}
	if !IsNil(o.AdDomainUsers) {
		toSerialize["ad_domain_users"] = o.AdDomainUsers
	}
	if !IsNil(o.AdUserGroupIds) {
		toSerialize["ad_user_group_ids"] = o.AdUserGroupIds
	}
	if !IsNil(o.AdUserIds) {
		toSerialize["ad_user_ids"] = o.AdUserIds
	}
	if !IsNil(o.DfsStoragePolicy) {
		toSerialize["dfs_storage_policy"] = o.DfsStoragePolicy
	}
	if !IsNil(o.ExcludeFileName) {
		toSerialize["exclude_file_name"] = o.ExcludeFileName
	}
	if !IsNil(o.ExcludeUserName) {
		toSerialize["exclude_user_name"] = o.ExcludeUserName
	}
	if !IsNil(o.FileNames) {
		toSerialize["file_names"] = o.FileNames
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LdapDomainUserGroups) {
		toSerialize["ldap_domain_user_groups"] = o.LdapDomainUserGroups
	}
	if !IsNil(o.LdapDomainUsers) {
		toSerialize["ldap_domain_users"] = o.LdapDomainUsers
	}
	if !IsNil(o.LdapUserGroupIds) {
		toSerialize["ldap_user_group_ids"] = o.LdapUserGroupIds
	}
	if !IsNil(o.LdapUserIds) {
		toSerialize["ldap_user_ids"] = o.LdapUserIds
	}
	if !IsNil(o.UserGroupNames) {
		toSerialize["user_group_names"] = o.UserGroupNames
	}
	if !IsNil(o.UserNames) {
		toSerialize["user_names"] = o.UserNames
	}
	return toSerialize, nil
}

type NullableDfsPolicyFilter struct {
	value *DfsPolicyFilter
	isSet bool
}

func (v NullableDfsPolicyFilter) Get() *DfsPolicyFilter {
	return v.value
}

func (v *NullableDfsPolicyFilter) Set(val *DfsPolicyFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsPolicyFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsPolicyFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsPolicyFilter(val *DfsPolicyFilter) *NullableDfsPolicyFilter {
	return &NullableDfsPolicyFilter{value: val, isSet: true}
}

func (v NullableDfsPolicyFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsPolicyFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


