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

// checks if the FSUserGroupAddUsersReqUserGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FSUserGroupAddUsersReqUserGroup{}

// FSUserGroupAddUsersReqUserGroup struct for FSUserGroupAddUsersReqUserGroup
type FSUserGroupAddUsersReqUserGroup struct {
	// ids of users, which are required when type is smb or ftp
	FsUserIds []int64 `json:"fs_user_ids,omitempty"`
}

// NewFSUserGroupAddUsersReqUserGroup instantiates a new FSUserGroupAddUsersReqUserGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFSUserGroupAddUsersReqUserGroup() *FSUserGroupAddUsersReqUserGroup {
	this := FSUserGroupAddUsersReqUserGroup{}
	return &this
}

// NewFSUserGroupAddUsersReqUserGroupWithDefaults instantiates a new FSUserGroupAddUsersReqUserGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFSUserGroupAddUsersReqUserGroupWithDefaults() *FSUserGroupAddUsersReqUserGroup {
	this := FSUserGroupAddUsersReqUserGroup{}
	return &this
}

// GetFsUserIds returns the FsUserIds field value if set, zero value otherwise.
func (o *FSUserGroupAddUsersReqUserGroup) GetFsUserIds() []int64 {
	if o == nil || IsNil(o.FsUserIds) {
		var ret []int64
		return ret
	}
	return o.FsUserIds
}

// GetFsUserIdsOk returns a tuple with the FsUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserGroupAddUsersReqUserGroup) GetFsUserIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.FsUserIds) {
		return nil, false
	}
	return o.FsUserIds, true
}

// HasFsUserIds returns a boolean if a field has been set.
func (o *FSUserGroupAddUsersReqUserGroup) HasFsUserIds() bool {
	if o != nil && !IsNil(o.FsUserIds) {
		return true
	}

	return false
}

// SetFsUserIds gets a reference to the given []int64 and assigns it to the FsUserIds field.
func (o *FSUserGroupAddUsersReqUserGroup) SetFsUserIds(v []int64) {
	o.FsUserIds = v
}

func (o FSUserGroupAddUsersReqUserGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FSUserGroupAddUsersReqUserGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FsUserIds) {
		toSerialize["fs_user_ids"] = o.FsUserIds
	}
	return toSerialize, nil
}

type NullableFSUserGroupAddUsersReqUserGroup struct {
	value *FSUserGroupAddUsersReqUserGroup
	isSet bool
}

func (v NullableFSUserGroupAddUsersReqUserGroup) Get() *FSUserGroupAddUsersReqUserGroup {
	return v.value
}

func (v *NullableFSUserGroupAddUsersReqUserGroup) Set(val *FSUserGroupAddUsersReqUserGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableFSUserGroupAddUsersReqUserGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableFSUserGroupAddUsersReqUserGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFSUserGroupAddUsersReqUserGroup(val *FSUserGroupAddUsersReqUserGroup) *NullableFSUserGroupAddUsersReqUserGroup {
	return &NullableFSUserGroupAddUsersReqUserGroup{value: val, isSet: true}
}

func (v NullableFSUserGroupAddUsersReqUserGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFSUserGroupAddUsersReqUserGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


