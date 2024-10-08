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

// checks if the DfsSMBShareACLReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsSMBShareACLReq{}

// DfsSMBShareACLReq struct for DfsSMBShareACLReq
type DfsSMBShareACLReq struct {
	// description of user or user group
	Description *string `json:"description,omitempty"`
	// id of user group
	FsUserGroupId *int64 `json:"fs_user_group_id,omitempty"`
	// id of user group
	FsUserGroupName *string `json:"fs_user_group_name,omitempty"`
	// id of user
	FsUserId *int64 `json:"fs_user_id,omitempty"`
	// id of user
	FsUserName *string `json:"fs_user_name,omitempty"`
	// id of user group
	Id *int64 `json:"id,omitempty"`
	// readonly or readwrite access
	Permission *string `json:"permission,omitempty"`
	// security of share acl
	Security *string `json:"security,omitempty"`
	// type of share acl
	Type *string `json:"type,omitempty"`
}

// NewDfsSMBShareACLReq instantiates a new DfsSMBShareACLReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsSMBShareACLReq() *DfsSMBShareACLReq {
	this := DfsSMBShareACLReq{}
	return &this
}

// NewDfsSMBShareACLReqWithDefaults instantiates a new DfsSMBShareACLReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsSMBShareACLReqWithDefaults() *DfsSMBShareACLReq {
	this := DfsSMBShareACLReq{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DfsSMBShareACLReq) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBShareACLReq) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DfsSMBShareACLReq) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DfsSMBShareACLReq) SetDescription(v string) {
	o.Description = &v
}

// GetFsUserGroupId returns the FsUserGroupId field value if set, zero value otherwise.
func (o *DfsSMBShareACLReq) GetFsUserGroupId() int64 {
	if o == nil || IsNil(o.FsUserGroupId) {
		var ret int64
		return ret
	}
	return *o.FsUserGroupId
}

// GetFsUserGroupIdOk returns a tuple with the FsUserGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBShareACLReq) GetFsUserGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.FsUserGroupId) {
		return nil, false
	}
	return o.FsUserGroupId, true
}

// HasFsUserGroupId returns a boolean if a field has been set.
func (o *DfsSMBShareACLReq) HasFsUserGroupId() bool {
	if o != nil && !IsNil(o.FsUserGroupId) {
		return true
	}

	return false
}

// SetFsUserGroupId gets a reference to the given int64 and assigns it to the FsUserGroupId field.
func (o *DfsSMBShareACLReq) SetFsUserGroupId(v int64) {
	o.FsUserGroupId = &v
}

// GetFsUserGroupName returns the FsUserGroupName field value if set, zero value otherwise.
func (o *DfsSMBShareACLReq) GetFsUserGroupName() string {
	if o == nil || IsNil(o.FsUserGroupName) {
		var ret string
		return ret
	}
	return *o.FsUserGroupName
}

// GetFsUserGroupNameOk returns a tuple with the FsUserGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBShareACLReq) GetFsUserGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.FsUserGroupName) {
		return nil, false
	}
	return o.FsUserGroupName, true
}

// HasFsUserGroupName returns a boolean if a field has been set.
func (o *DfsSMBShareACLReq) HasFsUserGroupName() bool {
	if o != nil && !IsNil(o.FsUserGroupName) {
		return true
	}

	return false
}

// SetFsUserGroupName gets a reference to the given string and assigns it to the FsUserGroupName field.
func (o *DfsSMBShareACLReq) SetFsUserGroupName(v string) {
	o.FsUserGroupName = &v
}

// GetFsUserId returns the FsUserId field value if set, zero value otherwise.
func (o *DfsSMBShareACLReq) GetFsUserId() int64 {
	if o == nil || IsNil(o.FsUserId) {
		var ret int64
		return ret
	}
	return *o.FsUserId
}

// GetFsUserIdOk returns a tuple with the FsUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBShareACLReq) GetFsUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.FsUserId) {
		return nil, false
	}
	return o.FsUserId, true
}

// HasFsUserId returns a boolean if a field has been set.
func (o *DfsSMBShareACLReq) HasFsUserId() bool {
	if o != nil && !IsNil(o.FsUserId) {
		return true
	}

	return false
}

// SetFsUserId gets a reference to the given int64 and assigns it to the FsUserId field.
func (o *DfsSMBShareACLReq) SetFsUserId(v int64) {
	o.FsUserId = &v
}

// GetFsUserName returns the FsUserName field value if set, zero value otherwise.
func (o *DfsSMBShareACLReq) GetFsUserName() string {
	if o == nil || IsNil(o.FsUserName) {
		var ret string
		return ret
	}
	return *o.FsUserName
}

// GetFsUserNameOk returns a tuple with the FsUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBShareACLReq) GetFsUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.FsUserName) {
		return nil, false
	}
	return o.FsUserName, true
}

// HasFsUserName returns a boolean if a field has been set.
func (o *DfsSMBShareACLReq) HasFsUserName() bool {
	if o != nil && !IsNil(o.FsUserName) {
		return true
	}

	return false
}

// SetFsUserName gets a reference to the given string and assigns it to the FsUserName field.
func (o *DfsSMBShareACLReq) SetFsUserName(v string) {
	o.FsUserName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsSMBShareACLReq) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBShareACLReq) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsSMBShareACLReq) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsSMBShareACLReq) SetId(v int64) {
	o.Id = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *DfsSMBShareACLReq) GetPermission() string {
	if o == nil || IsNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBShareACLReq) GetPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *DfsSMBShareACLReq) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *DfsSMBShareACLReq) SetPermission(v string) {
	o.Permission = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *DfsSMBShareACLReq) GetSecurity() string {
	if o == nil || IsNil(o.Security) {
		var ret string
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBShareACLReq) GetSecurityOk() (*string, bool) {
	if o == nil || IsNil(o.Security) {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *DfsSMBShareACLReq) HasSecurity() bool {
	if o != nil && !IsNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given string and assigns it to the Security field.
func (o *DfsSMBShareACLReq) SetSecurity(v string) {
	o.Security = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DfsSMBShareACLReq) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBShareACLReq) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DfsSMBShareACLReq) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DfsSMBShareACLReq) SetType(v string) {
	o.Type = &v
}

func (o DfsSMBShareACLReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsSMBShareACLReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FsUserGroupId) {
		toSerialize["fs_user_group_id"] = o.FsUserGroupId
	}
	if !IsNil(o.FsUserGroupName) {
		toSerialize["fs_user_group_name"] = o.FsUserGroupName
	}
	if !IsNil(o.FsUserId) {
		toSerialize["fs_user_id"] = o.FsUserId
	}
	if !IsNil(o.FsUserName) {
		toSerialize["fs_user_name"] = o.FsUserName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}
	if !IsNil(o.Security) {
		toSerialize["security"] = o.Security
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDfsSMBShareACLReq struct {
	value *DfsSMBShareACLReq
	isSet bool
}

func (v NullableDfsSMBShareACLReq) Get() *DfsSMBShareACLReq {
	return v.value
}

func (v *NullableDfsSMBShareACLReq) Set(val *DfsSMBShareACLReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsSMBShareACLReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsSMBShareACLReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsSMBShareACLReq(val *DfsSMBShareACLReq) *NullableDfsSMBShareACLReq {
	return &NullableDfsSMBShareACLReq{value: val, isSet: true}
}

func (v NullableDfsSMBShareACLReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsSMBShareACLReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


