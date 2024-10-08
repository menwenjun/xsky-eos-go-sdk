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

// checks if the ObjectStorageBucketUpdateReqBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStorageBucketUpdateReqBucket{}

// ObjectStorageBucketUpdateReqBucket struct for ObjectStorageBucketUpdateReqBucket
type ObjectStorageBucketUpdateReqBucket struct {
	AllUserPermission *string `json:"all_user_permission,omitempty"`
	AuthUserPermission *string `json:"auth_user_permission,omitempty"`
	DeleteArchiveStorageClass *string `json:"delete_archive_storage_class,omitempty"`
	Description *string `json:"description,omitempty"`
	ExternalQuotaMaxObjects *int64 `json:"external_quota_max_objects,omitempty"`
	ExternalQuotaMaxSize *int64 `json:"external_quota_max_size,omitempty"`
	Flag *BucketFlagReq `json:"flag,omitempty"`
	LocalQuotaMaxObjects *int64 `json:"local_quota_max_objects,omitempty"`
	LocalQuotaMaxSize *int64 `json:"local_quota_max_size,omitempty"`
	LogDeliveryPermission *string `json:"log_delivery_permission,omitempty"`
	OwnerId *int64 `json:"owner_id,omitempty"`
	OwnerPermission *string `json:"owner_permission,omitempty"`
	Qos *OSBucketQos `json:"qos,omitempty"`
	QuotaMaxObjects *int64 `json:"quota_max_objects,omitempty"`
	QuotaMaxSize *int64 `json:"quota_max_size,omitempty"`
	RestoreDays *int64 `json:"restore_days,omitempty"`
}

// NewObjectStorageBucketUpdateReqBucket instantiates a new ObjectStorageBucketUpdateReqBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStorageBucketUpdateReqBucket() *ObjectStorageBucketUpdateReqBucket {
	this := ObjectStorageBucketUpdateReqBucket{}
	return &this
}

// NewObjectStorageBucketUpdateReqBucketWithDefaults instantiates a new ObjectStorageBucketUpdateReqBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStorageBucketUpdateReqBucketWithDefaults() *ObjectStorageBucketUpdateReqBucket {
	this := ObjectStorageBucketUpdateReqBucket{}
	return &this
}

// GetAllUserPermission returns the AllUserPermission field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetAllUserPermission() string {
	if o == nil || IsNil(o.AllUserPermission) {
		var ret string
		return ret
	}
	return *o.AllUserPermission
}

// GetAllUserPermissionOk returns a tuple with the AllUserPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetAllUserPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.AllUserPermission) {
		return nil, false
	}
	return o.AllUserPermission, true
}

// HasAllUserPermission returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasAllUserPermission() bool {
	if o != nil && !IsNil(o.AllUserPermission) {
		return true
	}

	return false
}

// SetAllUserPermission gets a reference to the given string and assigns it to the AllUserPermission field.
func (o *ObjectStorageBucketUpdateReqBucket) SetAllUserPermission(v string) {
	o.AllUserPermission = &v
}

// GetAuthUserPermission returns the AuthUserPermission field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetAuthUserPermission() string {
	if o == nil || IsNil(o.AuthUserPermission) {
		var ret string
		return ret
	}
	return *o.AuthUserPermission
}

// GetAuthUserPermissionOk returns a tuple with the AuthUserPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetAuthUserPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.AuthUserPermission) {
		return nil, false
	}
	return o.AuthUserPermission, true
}

// HasAuthUserPermission returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasAuthUserPermission() bool {
	if o != nil && !IsNil(o.AuthUserPermission) {
		return true
	}

	return false
}

// SetAuthUserPermission gets a reference to the given string and assigns it to the AuthUserPermission field.
func (o *ObjectStorageBucketUpdateReqBucket) SetAuthUserPermission(v string) {
	o.AuthUserPermission = &v
}

// GetDeleteArchiveStorageClass returns the DeleteArchiveStorageClass field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetDeleteArchiveStorageClass() string {
	if o == nil || IsNil(o.DeleteArchiveStorageClass) {
		var ret string
		return ret
	}
	return *o.DeleteArchiveStorageClass
}

// GetDeleteArchiveStorageClassOk returns a tuple with the DeleteArchiveStorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetDeleteArchiveStorageClassOk() (*string, bool) {
	if o == nil || IsNil(o.DeleteArchiveStorageClass) {
		return nil, false
	}
	return o.DeleteArchiveStorageClass, true
}

// HasDeleteArchiveStorageClass returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasDeleteArchiveStorageClass() bool {
	if o != nil && !IsNil(o.DeleteArchiveStorageClass) {
		return true
	}

	return false
}

// SetDeleteArchiveStorageClass gets a reference to the given string and assigns it to the DeleteArchiveStorageClass field.
func (o *ObjectStorageBucketUpdateReqBucket) SetDeleteArchiveStorageClass(v string) {
	o.DeleteArchiveStorageClass = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ObjectStorageBucketUpdateReqBucket) SetDescription(v string) {
	o.Description = &v
}

// GetExternalQuotaMaxObjects returns the ExternalQuotaMaxObjects field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetExternalQuotaMaxObjects() int64 {
	if o == nil || IsNil(o.ExternalQuotaMaxObjects) {
		var ret int64
		return ret
	}
	return *o.ExternalQuotaMaxObjects
}

// GetExternalQuotaMaxObjectsOk returns a tuple with the ExternalQuotaMaxObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetExternalQuotaMaxObjectsOk() (*int64, bool) {
	if o == nil || IsNil(o.ExternalQuotaMaxObjects) {
		return nil, false
	}
	return o.ExternalQuotaMaxObjects, true
}

// HasExternalQuotaMaxObjects returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasExternalQuotaMaxObjects() bool {
	if o != nil && !IsNil(o.ExternalQuotaMaxObjects) {
		return true
	}

	return false
}

// SetExternalQuotaMaxObjects gets a reference to the given int64 and assigns it to the ExternalQuotaMaxObjects field.
func (o *ObjectStorageBucketUpdateReqBucket) SetExternalQuotaMaxObjects(v int64) {
	o.ExternalQuotaMaxObjects = &v
}

// GetExternalQuotaMaxSize returns the ExternalQuotaMaxSize field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetExternalQuotaMaxSize() int64 {
	if o == nil || IsNil(o.ExternalQuotaMaxSize) {
		var ret int64
		return ret
	}
	return *o.ExternalQuotaMaxSize
}

// GetExternalQuotaMaxSizeOk returns a tuple with the ExternalQuotaMaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetExternalQuotaMaxSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.ExternalQuotaMaxSize) {
		return nil, false
	}
	return o.ExternalQuotaMaxSize, true
}

// HasExternalQuotaMaxSize returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasExternalQuotaMaxSize() bool {
	if o != nil && !IsNil(o.ExternalQuotaMaxSize) {
		return true
	}

	return false
}

// SetExternalQuotaMaxSize gets a reference to the given int64 and assigns it to the ExternalQuotaMaxSize field.
func (o *ObjectStorageBucketUpdateReqBucket) SetExternalQuotaMaxSize(v int64) {
	o.ExternalQuotaMaxSize = &v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetFlag() BucketFlagReq {
	if o == nil || IsNil(o.Flag) {
		var ret BucketFlagReq
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetFlagOk() (*BucketFlagReq, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given BucketFlagReq and assigns it to the Flag field.
func (o *ObjectStorageBucketUpdateReqBucket) SetFlag(v BucketFlagReq) {
	o.Flag = &v
}

// GetLocalQuotaMaxObjects returns the LocalQuotaMaxObjects field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetLocalQuotaMaxObjects() int64 {
	if o == nil || IsNil(o.LocalQuotaMaxObjects) {
		var ret int64
		return ret
	}
	return *o.LocalQuotaMaxObjects
}

// GetLocalQuotaMaxObjectsOk returns a tuple with the LocalQuotaMaxObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetLocalQuotaMaxObjectsOk() (*int64, bool) {
	if o == nil || IsNil(o.LocalQuotaMaxObjects) {
		return nil, false
	}
	return o.LocalQuotaMaxObjects, true
}

// HasLocalQuotaMaxObjects returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasLocalQuotaMaxObjects() bool {
	if o != nil && !IsNil(o.LocalQuotaMaxObjects) {
		return true
	}

	return false
}

// SetLocalQuotaMaxObjects gets a reference to the given int64 and assigns it to the LocalQuotaMaxObjects field.
func (o *ObjectStorageBucketUpdateReqBucket) SetLocalQuotaMaxObjects(v int64) {
	o.LocalQuotaMaxObjects = &v
}

// GetLocalQuotaMaxSize returns the LocalQuotaMaxSize field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetLocalQuotaMaxSize() int64 {
	if o == nil || IsNil(o.LocalQuotaMaxSize) {
		var ret int64
		return ret
	}
	return *o.LocalQuotaMaxSize
}

// GetLocalQuotaMaxSizeOk returns a tuple with the LocalQuotaMaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetLocalQuotaMaxSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.LocalQuotaMaxSize) {
		return nil, false
	}
	return o.LocalQuotaMaxSize, true
}

// HasLocalQuotaMaxSize returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasLocalQuotaMaxSize() bool {
	if o != nil && !IsNil(o.LocalQuotaMaxSize) {
		return true
	}

	return false
}

// SetLocalQuotaMaxSize gets a reference to the given int64 and assigns it to the LocalQuotaMaxSize field.
func (o *ObjectStorageBucketUpdateReqBucket) SetLocalQuotaMaxSize(v int64) {
	o.LocalQuotaMaxSize = &v
}

// GetLogDeliveryPermission returns the LogDeliveryPermission field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetLogDeliveryPermission() string {
	if o == nil || IsNil(o.LogDeliveryPermission) {
		var ret string
		return ret
	}
	return *o.LogDeliveryPermission
}

// GetLogDeliveryPermissionOk returns a tuple with the LogDeliveryPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetLogDeliveryPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.LogDeliveryPermission) {
		return nil, false
	}
	return o.LogDeliveryPermission, true
}

// HasLogDeliveryPermission returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasLogDeliveryPermission() bool {
	if o != nil && !IsNil(o.LogDeliveryPermission) {
		return true
	}

	return false
}

// SetLogDeliveryPermission gets a reference to the given string and assigns it to the LogDeliveryPermission field.
func (o *ObjectStorageBucketUpdateReqBucket) SetLogDeliveryPermission(v string) {
	o.LogDeliveryPermission = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetOwnerId() int64 {
	if o == nil || IsNil(o.OwnerId) {
		var ret int64
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetOwnerIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given int64 and assigns it to the OwnerId field.
func (o *ObjectStorageBucketUpdateReqBucket) SetOwnerId(v int64) {
	o.OwnerId = &v
}

// GetOwnerPermission returns the OwnerPermission field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetOwnerPermission() string {
	if o == nil || IsNil(o.OwnerPermission) {
		var ret string
		return ret
	}
	return *o.OwnerPermission
}

// GetOwnerPermissionOk returns a tuple with the OwnerPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetOwnerPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerPermission) {
		return nil, false
	}
	return o.OwnerPermission, true
}

// HasOwnerPermission returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasOwnerPermission() bool {
	if o != nil && !IsNil(o.OwnerPermission) {
		return true
	}

	return false
}

// SetOwnerPermission gets a reference to the given string and assigns it to the OwnerPermission field.
func (o *ObjectStorageBucketUpdateReqBucket) SetOwnerPermission(v string) {
	o.OwnerPermission = &v
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetQos() OSBucketQos {
	if o == nil || IsNil(o.Qos) {
		var ret OSBucketQos
		return ret
	}
	return *o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetQosOk() (*OSBucketQos, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given OSBucketQos and assigns it to the Qos field.
func (o *ObjectStorageBucketUpdateReqBucket) SetQos(v OSBucketQos) {
	o.Qos = &v
}

// GetQuotaMaxObjects returns the QuotaMaxObjects field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetQuotaMaxObjects() int64 {
	if o == nil || IsNil(o.QuotaMaxObjects) {
		var ret int64
		return ret
	}
	return *o.QuotaMaxObjects
}

// GetQuotaMaxObjectsOk returns a tuple with the QuotaMaxObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetQuotaMaxObjectsOk() (*int64, bool) {
	if o == nil || IsNil(o.QuotaMaxObjects) {
		return nil, false
	}
	return o.QuotaMaxObjects, true
}

// HasQuotaMaxObjects returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasQuotaMaxObjects() bool {
	if o != nil && !IsNil(o.QuotaMaxObjects) {
		return true
	}

	return false
}

// SetQuotaMaxObjects gets a reference to the given int64 and assigns it to the QuotaMaxObjects field.
func (o *ObjectStorageBucketUpdateReqBucket) SetQuotaMaxObjects(v int64) {
	o.QuotaMaxObjects = &v
}

// GetQuotaMaxSize returns the QuotaMaxSize field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetQuotaMaxSize() int64 {
	if o == nil || IsNil(o.QuotaMaxSize) {
		var ret int64
		return ret
	}
	return *o.QuotaMaxSize
}

// GetQuotaMaxSizeOk returns a tuple with the QuotaMaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetQuotaMaxSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.QuotaMaxSize) {
		return nil, false
	}
	return o.QuotaMaxSize, true
}

// HasQuotaMaxSize returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasQuotaMaxSize() bool {
	if o != nil && !IsNil(o.QuotaMaxSize) {
		return true
	}

	return false
}

// SetQuotaMaxSize gets a reference to the given int64 and assigns it to the QuotaMaxSize field.
func (o *ObjectStorageBucketUpdateReqBucket) SetQuotaMaxSize(v int64) {
	o.QuotaMaxSize = &v
}

// GetRestoreDays returns the RestoreDays field value if set, zero value otherwise.
func (o *ObjectStorageBucketUpdateReqBucket) GetRestoreDays() int64 {
	if o == nil || IsNil(o.RestoreDays) {
		var ret int64
		return ret
	}
	return *o.RestoreDays
}

// GetRestoreDaysOk returns a tuple with the RestoreDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketUpdateReqBucket) GetRestoreDaysOk() (*int64, bool) {
	if o == nil || IsNil(o.RestoreDays) {
		return nil, false
	}
	return o.RestoreDays, true
}

// HasRestoreDays returns a boolean if a field has been set.
func (o *ObjectStorageBucketUpdateReqBucket) HasRestoreDays() bool {
	if o != nil && !IsNil(o.RestoreDays) {
		return true
	}

	return false
}

// SetRestoreDays gets a reference to the given int64 and assigns it to the RestoreDays field.
func (o *ObjectStorageBucketUpdateReqBucket) SetRestoreDays(v int64) {
	o.RestoreDays = &v
}

func (o ObjectStorageBucketUpdateReqBucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStorageBucketUpdateReqBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllUserPermission) {
		toSerialize["all_user_permission"] = o.AllUserPermission
	}
	if !IsNil(o.AuthUserPermission) {
		toSerialize["auth_user_permission"] = o.AuthUserPermission
	}
	if !IsNil(o.DeleteArchiveStorageClass) {
		toSerialize["delete_archive_storage_class"] = o.DeleteArchiveStorageClass
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalQuotaMaxObjects) {
		toSerialize["external_quota_max_objects"] = o.ExternalQuotaMaxObjects
	}
	if !IsNil(o.ExternalQuotaMaxSize) {
		toSerialize["external_quota_max_size"] = o.ExternalQuotaMaxSize
	}
	if !IsNil(o.Flag) {
		toSerialize["flag"] = o.Flag
	}
	if !IsNil(o.LocalQuotaMaxObjects) {
		toSerialize["local_quota_max_objects"] = o.LocalQuotaMaxObjects
	}
	if !IsNil(o.LocalQuotaMaxSize) {
		toSerialize["local_quota_max_size"] = o.LocalQuotaMaxSize
	}
	if !IsNil(o.LogDeliveryPermission) {
		toSerialize["log_delivery_permission"] = o.LogDeliveryPermission
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.OwnerPermission) {
		toSerialize["owner_permission"] = o.OwnerPermission
	}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	if !IsNil(o.QuotaMaxObjects) {
		toSerialize["quota_max_objects"] = o.QuotaMaxObjects
	}
	if !IsNil(o.QuotaMaxSize) {
		toSerialize["quota_max_size"] = o.QuotaMaxSize
	}
	if !IsNil(o.RestoreDays) {
		toSerialize["restore_days"] = o.RestoreDays
	}
	return toSerialize, nil
}

type NullableObjectStorageBucketUpdateReqBucket struct {
	value *ObjectStorageBucketUpdateReqBucket
	isSet bool
}

func (v NullableObjectStorageBucketUpdateReqBucket) Get() *ObjectStorageBucketUpdateReqBucket {
	return v.value
}

func (v *NullableObjectStorageBucketUpdateReqBucket) Set(val *ObjectStorageBucketUpdateReqBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStorageBucketUpdateReqBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStorageBucketUpdateReqBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStorageBucketUpdateReqBucket(val *ObjectStorageBucketUpdateReqBucket) *NullableObjectStorageBucketUpdateReqBucket {
	return &NullableObjectStorageBucketUpdateReqBucket{value: val, isSet: true}
}

func (v NullableObjectStorageBucketUpdateReqBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStorageBucketUpdateReqBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


