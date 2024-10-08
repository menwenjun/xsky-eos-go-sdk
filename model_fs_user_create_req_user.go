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

// checks if the FSUserCreateReqUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FSUserCreateReqUser{}

// FSUserCreateReqUser struct for FSUserCreateReqUser
type FSUserCreateReqUser struct {
	BucketLimit *int64 `json:"bucket_limit,omitempty"`
	BucketPath *string `json:"bucket_path,omitempty"`
	BucketPermission *string `json:"bucket_permission,omitempty"`
	GatewayGroupId *int64 `json:"gateway_group_id,omitempty"`
	S3Enabled *bool `json:"s3_enabled,omitempty"`
	S3Keys []S3Key `json:"s3_keys,omitempty"`
	// description of file storage user
	Description *string `json:"description,omitempty"`
	// email of file storage user
	Email *string `json:"email,omitempty"`
	// name of file storage user
	Name *string `json:"name,omitempty"`
	// password of file storage user
	Password *string `json:"password,omitempty"`
	// primary group of user
	PrimaryGroupId *int64 `json:"primary_group_id,omitempty"`
	// secondary group of user
	UserGroupIds []int64 `json:"user_group_ids,omitempty"`
	// uid of file storage user
	UserId *int64 `json:"user_id,omitempty"`
}

// NewFSUserCreateReqUser instantiates a new FSUserCreateReqUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFSUserCreateReqUser() *FSUserCreateReqUser {
	this := FSUserCreateReqUser{}
	return &this
}

// NewFSUserCreateReqUserWithDefaults instantiates a new FSUserCreateReqUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFSUserCreateReqUserWithDefaults() *FSUserCreateReqUser {
	this := FSUserCreateReqUser{}
	return &this
}

// GetBucketLimit returns the BucketLimit field value if set, zero value otherwise.
func (o *FSUserCreateReqUser) GetBucketLimit() int64 {
	if o == nil || IsNil(o.BucketLimit) {
		var ret int64
		return ret
	}
	return *o.BucketLimit
}

// GetBucketLimitOk returns a tuple with the BucketLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserCreateReqUser) GetBucketLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.BucketLimit) {
		return nil, false
	}
	return o.BucketLimit, true
}

// HasBucketLimit returns a boolean if a field has been set.
func (o *FSUserCreateReqUser) HasBucketLimit() bool {
	if o != nil && !IsNil(o.BucketLimit) {
		return true
	}

	return false
}

// SetBucketLimit gets a reference to the given int64 and assigns it to the BucketLimit field.
func (o *FSUserCreateReqUser) SetBucketLimit(v int64) {
	o.BucketLimit = &v
}

// GetBucketPath returns the BucketPath field value if set, zero value otherwise.
func (o *FSUserCreateReqUser) GetBucketPath() string {
	if o == nil || IsNil(o.BucketPath) {
		var ret string
		return ret
	}
	return *o.BucketPath
}

// GetBucketPathOk returns a tuple with the BucketPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserCreateReqUser) GetBucketPathOk() (*string, bool) {
	if o == nil || IsNil(o.BucketPath) {
		return nil, false
	}
	return o.BucketPath, true
}

// HasBucketPath returns a boolean if a field has been set.
func (o *FSUserCreateReqUser) HasBucketPath() bool {
	if o != nil && !IsNil(o.BucketPath) {
		return true
	}

	return false
}

// SetBucketPath gets a reference to the given string and assigns it to the BucketPath field.
func (o *FSUserCreateReqUser) SetBucketPath(v string) {
	o.BucketPath = &v
}

// GetBucketPermission returns the BucketPermission field value if set, zero value otherwise.
func (o *FSUserCreateReqUser) GetBucketPermission() string {
	if o == nil || IsNil(o.BucketPermission) {
		var ret string
		return ret
	}
	return *o.BucketPermission
}

// GetBucketPermissionOk returns a tuple with the BucketPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserCreateReqUser) GetBucketPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.BucketPermission) {
		return nil, false
	}
	return o.BucketPermission, true
}

// HasBucketPermission returns a boolean if a field has been set.
func (o *FSUserCreateReqUser) HasBucketPermission() bool {
	if o != nil && !IsNil(o.BucketPermission) {
		return true
	}

	return false
}

// SetBucketPermission gets a reference to the given string and assigns it to the BucketPermission field.
func (o *FSUserCreateReqUser) SetBucketPermission(v string) {
	o.BucketPermission = &v
}

// GetGatewayGroupId returns the GatewayGroupId field value if set, zero value otherwise.
func (o *FSUserCreateReqUser) GetGatewayGroupId() int64 {
	if o == nil || IsNil(o.GatewayGroupId) {
		var ret int64
		return ret
	}
	return *o.GatewayGroupId
}

// GetGatewayGroupIdOk returns a tuple with the GatewayGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserCreateReqUser) GetGatewayGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayGroupId) {
		return nil, false
	}
	return o.GatewayGroupId, true
}

// HasGatewayGroupId returns a boolean if a field has been set.
func (o *FSUserCreateReqUser) HasGatewayGroupId() bool {
	if o != nil && !IsNil(o.GatewayGroupId) {
		return true
	}

	return false
}

// SetGatewayGroupId gets a reference to the given int64 and assigns it to the GatewayGroupId field.
func (o *FSUserCreateReqUser) SetGatewayGroupId(v int64) {
	o.GatewayGroupId = &v
}

// GetS3Enabled returns the S3Enabled field value if set, zero value otherwise.
func (o *FSUserCreateReqUser) GetS3Enabled() bool {
	if o == nil || IsNil(o.S3Enabled) {
		var ret bool
		return ret
	}
	return *o.S3Enabled
}

// GetS3EnabledOk returns a tuple with the S3Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserCreateReqUser) GetS3EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.S3Enabled) {
		return nil, false
	}
	return o.S3Enabled, true
}

// HasS3Enabled returns a boolean if a field has been set.
func (o *FSUserCreateReqUser) HasS3Enabled() bool {
	if o != nil && !IsNil(o.S3Enabled) {
		return true
	}

	return false
}

// SetS3Enabled gets a reference to the given bool and assigns it to the S3Enabled field.
func (o *FSUserCreateReqUser) SetS3Enabled(v bool) {
	o.S3Enabled = &v
}

// GetS3Keys returns the S3Keys field value if set, zero value otherwise.
func (o *FSUserCreateReqUser) GetS3Keys() []S3Key {
	if o == nil || IsNil(o.S3Keys) {
		var ret []S3Key
		return ret
	}
	return o.S3Keys
}

// GetS3KeysOk returns a tuple with the S3Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserCreateReqUser) GetS3KeysOk() ([]S3Key, bool) {
	if o == nil || IsNil(o.S3Keys) {
		return nil, false
	}
	return o.S3Keys, true
}

// HasS3Keys returns a boolean if a field has been set.
func (o *FSUserCreateReqUser) HasS3Keys() bool {
	if o != nil && !IsNil(o.S3Keys) {
		return true
	}

	return false
}

// SetS3Keys gets a reference to the given []S3Key and assigns it to the S3Keys field.
func (o *FSUserCreateReqUser) SetS3Keys(v []S3Key) {
	o.S3Keys = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FSUserCreateReqUser) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserCreateReqUser) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FSUserCreateReqUser) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FSUserCreateReqUser) SetDescription(v string) {
	o.Description = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *FSUserCreateReqUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserCreateReqUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *FSUserCreateReqUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *FSUserCreateReqUser) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FSUserCreateReqUser) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserCreateReqUser) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FSUserCreateReqUser) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FSUserCreateReqUser) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *FSUserCreateReqUser) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserCreateReqUser) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *FSUserCreateReqUser) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *FSUserCreateReqUser) SetPassword(v string) {
	o.Password = &v
}

// GetPrimaryGroupId returns the PrimaryGroupId field value if set, zero value otherwise.
func (o *FSUserCreateReqUser) GetPrimaryGroupId() int64 {
	if o == nil || IsNil(o.PrimaryGroupId) {
		var ret int64
		return ret
	}
	return *o.PrimaryGroupId
}

// GetPrimaryGroupIdOk returns a tuple with the PrimaryGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserCreateReqUser) GetPrimaryGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PrimaryGroupId) {
		return nil, false
	}
	return o.PrimaryGroupId, true
}

// HasPrimaryGroupId returns a boolean if a field has been set.
func (o *FSUserCreateReqUser) HasPrimaryGroupId() bool {
	if o != nil && !IsNil(o.PrimaryGroupId) {
		return true
	}

	return false
}

// SetPrimaryGroupId gets a reference to the given int64 and assigns it to the PrimaryGroupId field.
func (o *FSUserCreateReqUser) SetPrimaryGroupId(v int64) {
	o.PrimaryGroupId = &v
}

// GetUserGroupIds returns the UserGroupIds field value if set, zero value otherwise.
func (o *FSUserCreateReqUser) GetUserGroupIds() []int64 {
	if o == nil || IsNil(o.UserGroupIds) {
		var ret []int64
		return ret
	}
	return o.UserGroupIds
}

// GetUserGroupIdsOk returns a tuple with the UserGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserCreateReqUser) GetUserGroupIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.UserGroupIds) {
		return nil, false
	}
	return o.UserGroupIds, true
}

// HasUserGroupIds returns a boolean if a field has been set.
func (o *FSUserCreateReqUser) HasUserGroupIds() bool {
	if o != nil && !IsNil(o.UserGroupIds) {
		return true
	}

	return false
}

// SetUserGroupIds gets a reference to the given []int64 and assigns it to the UserGroupIds field.
func (o *FSUserCreateReqUser) SetUserGroupIds(v []int64) {
	o.UserGroupIds = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *FSUserCreateReqUser) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSUserCreateReqUser) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *FSUserCreateReqUser) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *FSUserCreateReqUser) SetUserId(v int64) {
	o.UserId = &v
}

func (o FSUserCreateReqUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FSUserCreateReqUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BucketLimit) {
		toSerialize["bucket_limit"] = o.BucketLimit
	}
	if !IsNil(o.BucketPath) {
		toSerialize["bucket_path"] = o.BucketPath
	}
	if !IsNil(o.BucketPermission) {
		toSerialize["bucket_permission"] = o.BucketPermission
	}
	if !IsNil(o.GatewayGroupId) {
		toSerialize["gateway_group_id"] = o.GatewayGroupId
	}
	if !IsNil(o.S3Enabled) {
		toSerialize["s3_enabled"] = o.S3Enabled
	}
	if !IsNil(o.S3Keys) {
		toSerialize["s3_keys"] = o.S3Keys
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PrimaryGroupId) {
		toSerialize["primary_group_id"] = o.PrimaryGroupId
	}
	if !IsNil(o.UserGroupIds) {
		toSerialize["user_group_ids"] = o.UserGroupIds
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableFSUserCreateReqUser struct {
	value *FSUserCreateReqUser
	isSet bool
}

func (v NullableFSUserCreateReqUser) Get() *FSUserCreateReqUser {
	return v.value
}

func (v *NullableFSUserCreateReqUser) Set(val *FSUserCreateReqUser) {
	v.value = val
	v.isSet = true
}

func (v NullableFSUserCreateReqUser) IsSet() bool {
	return v.isSet
}

func (v *NullableFSUserCreateReqUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFSUserCreateReqUser(val *FSUserCreateReqUser) *NullableFSUserCreateReqUser {
	return &NullableFSUserCreateReqUser{value: val, isSet: true}
}

func (v NullableFSUserCreateReqUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFSUserCreateReqUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


