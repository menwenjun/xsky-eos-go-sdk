# ObjectStorageBucketUpdateReqBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllUserPermission** | Pointer to **string** |  | [optional] 
**AuthUserPermission** | Pointer to **string** |  | [optional] 
**DeleteArchiveStorageClass** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**ExternalQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**Flag** | Pointer to [**BucketFlagReq**](BucketFlagReq.md) |  | [optional] 
**LocalQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**LocalQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**LogDeliveryPermission** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **int64** |  | [optional] 
**OwnerPermission** | Pointer to **string** |  | [optional] 
**Qos** | Pointer to [**OSBucketQos**](OSBucketQos.md) |  | [optional] 
**QuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**QuotaMaxSize** | Pointer to **int64** |  | [optional] 
**RestoreDays** | Pointer to **int64** |  | [optional] 

## Methods

### NewObjectStorageBucketUpdateReqBucket

`func NewObjectStorageBucketUpdateReqBucket() *ObjectStorageBucketUpdateReqBucket`

NewObjectStorageBucketUpdateReqBucket instantiates a new ObjectStorageBucketUpdateReqBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageBucketUpdateReqBucketWithDefaults

`func NewObjectStorageBucketUpdateReqBucketWithDefaults() *ObjectStorageBucketUpdateReqBucket`

NewObjectStorageBucketUpdateReqBucketWithDefaults instantiates a new ObjectStorageBucketUpdateReqBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllUserPermission

`func (o *ObjectStorageBucketUpdateReqBucket) GetAllUserPermission() string`

GetAllUserPermission returns the AllUserPermission field if non-nil, zero value otherwise.

### GetAllUserPermissionOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetAllUserPermissionOk() (*string, bool)`

GetAllUserPermissionOk returns a tuple with the AllUserPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUserPermission

`func (o *ObjectStorageBucketUpdateReqBucket) SetAllUserPermission(v string)`

SetAllUserPermission sets AllUserPermission field to given value.

### HasAllUserPermission

`func (o *ObjectStorageBucketUpdateReqBucket) HasAllUserPermission() bool`

HasAllUserPermission returns a boolean if a field has been set.

### GetAuthUserPermission

`func (o *ObjectStorageBucketUpdateReqBucket) GetAuthUserPermission() string`

GetAuthUserPermission returns the AuthUserPermission field if non-nil, zero value otherwise.

### GetAuthUserPermissionOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetAuthUserPermissionOk() (*string, bool)`

GetAuthUserPermissionOk returns a tuple with the AuthUserPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUserPermission

`func (o *ObjectStorageBucketUpdateReqBucket) SetAuthUserPermission(v string)`

SetAuthUserPermission sets AuthUserPermission field to given value.

### HasAuthUserPermission

`func (o *ObjectStorageBucketUpdateReqBucket) HasAuthUserPermission() bool`

HasAuthUserPermission returns a boolean if a field has been set.

### GetDeleteArchiveStorageClass

`func (o *ObjectStorageBucketUpdateReqBucket) GetDeleteArchiveStorageClass() string`

GetDeleteArchiveStorageClass returns the DeleteArchiveStorageClass field if non-nil, zero value otherwise.

### GetDeleteArchiveStorageClassOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetDeleteArchiveStorageClassOk() (*string, bool)`

GetDeleteArchiveStorageClassOk returns a tuple with the DeleteArchiveStorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteArchiveStorageClass

`func (o *ObjectStorageBucketUpdateReqBucket) SetDeleteArchiveStorageClass(v string)`

SetDeleteArchiveStorageClass sets DeleteArchiveStorageClass field to given value.

### HasDeleteArchiveStorageClass

`func (o *ObjectStorageBucketUpdateReqBucket) HasDeleteArchiveStorageClass() bool`

HasDeleteArchiveStorageClass returns a boolean if a field has been set.

### GetDescription

`func (o *ObjectStorageBucketUpdateReqBucket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectStorageBucketUpdateReqBucket) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectStorageBucketUpdateReqBucket) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalQuotaMaxObjects

`func (o *ObjectStorageBucketUpdateReqBucket) GetExternalQuotaMaxObjects() int64`

GetExternalQuotaMaxObjects returns the ExternalQuotaMaxObjects field if non-nil, zero value otherwise.

### GetExternalQuotaMaxObjectsOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetExternalQuotaMaxObjectsOk() (*int64, bool)`

GetExternalQuotaMaxObjectsOk returns a tuple with the ExternalQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalQuotaMaxObjects

`func (o *ObjectStorageBucketUpdateReqBucket) SetExternalQuotaMaxObjects(v int64)`

SetExternalQuotaMaxObjects sets ExternalQuotaMaxObjects field to given value.

### HasExternalQuotaMaxObjects

`func (o *ObjectStorageBucketUpdateReqBucket) HasExternalQuotaMaxObjects() bool`

HasExternalQuotaMaxObjects returns a boolean if a field has been set.

### GetExternalQuotaMaxSize

`func (o *ObjectStorageBucketUpdateReqBucket) GetExternalQuotaMaxSize() int64`

GetExternalQuotaMaxSize returns the ExternalQuotaMaxSize field if non-nil, zero value otherwise.

### GetExternalQuotaMaxSizeOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetExternalQuotaMaxSizeOk() (*int64, bool)`

GetExternalQuotaMaxSizeOk returns a tuple with the ExternalQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalQuotaMaxSize

`func (o *ObjectStorageBucketUpdateReqBucket) SetExternalQuotaMaxSize(v int64)`

SetExternalQuotaMaxSize sets ExternalQuotaMaxSize field to given value.

### HasExternalQuotaMaxSize

`func (o *ObjectStorageBucketUpdateReqBucket) HasExternalQuotaMaxSize() bool`

HasExternalQuotaMaxSize returns a boolean if a field has been set.

### GetFlag

`func (o *ObjectStorageBucketUpdateReqBucket) GetFlag() BucketFlagReq`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetFlagOk() (*BucketFlagReq, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *ObjectStorageBucketUpdateReqBucket) SetFlag(v BucketFlagReq)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *ObjectStorageBucketUpdateReqBucket) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetLocalQuotaMaxObjects

`func (o *ObjectStorageBucketUpdateReqBucket) GetLocalQuotaMaxObjects() int64`

GetLocalQuotaMaxObjects returns the LocalQuotaMaxObjects field if non-nil, zero value otherwise.

### GetLocalQuotaMaxObjectsOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetLocalQuotaMaxObjectsOk() (*int64, bool)`

GetLocalQuotaMaxObjectsOk returns a tuple with the LocalQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalQuotaMaxObjects

`func (o *ObjectStorageBucketUpdateReqBucket) SetLocalQuotaMaxObjects(v int64)`

SetLocalQuotaMaxObjects sets LocalQuotaMaxObjects field to given value.

### HasLocalQuotaMaxObjects

`func (o *ObjectStorageBucketUpdateReqBucket) HasLocalQuotaMaxObjects() bool`

HasLocalQuotaMaxObjects returns a boolean if a field has been set.

### GetLocalQuotaMaxSize

`func (o *ObjectStorageBucketUpdateReqBucket) GetLocalQuotaMaxSize() int64`

GetLocalQuotaMaxSize returns the LocalQuotaMaxSize field if non-nil, zero value otherwise.

### GetLocalQuotaMaxSizeOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetLocalQuotaMaxSizeOk() (*int64, bool)`

GetLocalQuotaMaxSizeOk returns a tuple with the LocalQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalQuotaMaxSize

`func (o *ObjectStorageBucketUpdateReqBucket) SetLocalQuotaMaxSize(v int64)`

SetLocalQuotaMaxSize sets LocalQuotaMaxSize field to given value.

### HasLocalQuotaMaxSize

`func (o *ObjectStorageBucketUpdateReqBucket) HasLocalQuotaMaxSize() bool`

HasLocalQuotaMaxSize returns a boolean if a field has been set.

### GetLogDeliveryPermission

`func (o *ObjectStorageBucketUpdateReqBucket) GetLogDeliveryPermission() string`

GetLogDeliveryPermission returns the LogDeliveryPermission field if non-nil, zero value otherwise.

### GetLogDeliveryPermissionOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetLogDeliveryPermissionOk() (*string, bool)`

GetLogDeliveryPermissionOk returns a tuple with the LogDeliveryPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDeliveryPermission

`func (o *ObjectStorageBucketUpdateReqBucket) SetLogDeliveryPermission(v string)`

SetLogDeliveryPermission sets LogDeliveryPermission field to given value.

### HasLogDeliveryPermission

`func (o *ObjectStorageBucketUpdateReqBucket) HasLogDeliveryPermission() bool`

HasLogDeliveryPermission returns a boolean if a field has been set.

### GetOwnerId

`func (o *ObjectStorageBucketUpdateReqBucket) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ObjectStorageBucketUpdateReqBucket) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ObjectStorageBucketUpdateReqBucket) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerPermission

`func (o *ObjectStorageBucketUpdateReqBucket) GetOwnerPermission() string`

GetOwnerPermission returns the OwnerPermission field if non-nil, zero value otherwise.

### GetOwnerPermissionOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetOwnerPermissionOk() (*string, bool)`

GetOwnerPermissionOk returns a tuple with the OwnerPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPermission

`func (o *ObjectStorageBucketUpdateReqBucket) SetOwnerPermission(v string)`

SetOwnerPermission sets OwnerPermission field to given value.

### HasOwnerPermission

`func (o *ObjectStorageBucketUpdateReqBucket) HasOwnerPermission() bool`

HasOwnerPermission returns a boolean if a field has been set.

### GetQos

`func (o *ObjectStorageBucketUpdateReqBucket) GetQos() OSBucketQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetQosOk() (*OSBucketQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *ObjectStorageBucketUpdateReqBucket) SetQos(v OSBucketQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *ObjectStorageBucketUpdateReqBucket) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetQuotaMaxObjects

`func (o *ObjectStorageBucketUpdateReqBucket) GetQuotaMaxObjects() int64`

GetQuotaMaxObjects returns the QuotaMaxObjects field if non-nil, zero value otherwise.

### GetQuotaMaxObjectsOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetQuotaMaxObjectsOk() (*int64, bool)`

GetQuotaMaxObjectsOk returns a tuple with the QuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMaxObjects

`func (o *ObjectStorageBucketUpdateReqBucket) SetQuotaMaxObjects(v int64)`

SetQuotaMaxObjects sets QuotaMaxObjects field to given value.

### HasQuotaMaxObjects

`func (o *ObjectStorageBucketUpdateReqBucket) HasQuotaMaxObjects() bool`

HasQuotaMaxObjects returns a boolean if a field has been set.

### GetQuotaMaxSize

`func (o *ObjectStorageBucketUpdateReqBucket) GetQuotaMaxSize() int64`

GetQuotaMaxSize returns the QuotaMaxSize field if non-nil, zero value otherwise.

### GetQuotaMaxSizeOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetQuotaMaxSizeOk() (*int64, bool)`

GetQuotaMaxSizeOk returns a tuple with the QuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMaxSize

`func (o *ObjectStorageBucketUpdateReqBucket) SetQuotaMaxSize(v int64)`

SetQuotaMaxSize sets QuotaMaxSize field to given value.

### HasQuotaMaxSize

`func (o *ObjectStorageBucketUpdateReqBucket) HasQuotaMaxSize() bool`

HasQuotaMaxSize returns a boolean if a field has been set.

### GetRestoreDays

`func (o *ObjectStorageBucketUpdateReqBucket) GetRestoreDays() int64`

GetRestoreDays returns the RestoreDays field if non-nil, zero value otherwise.

### GetRestoreDaysOk

`func (o *ObjectStorageBucketUpdateReqBucket) GetRestoreDaysOk() (*int64, bool)`

GetRestoreDaysOk returns a tuple with the RestoreDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreDays

`func (o *ObjectStorageBucketUpdateReqBucket) SetRestoreDays(v int64)`

SetRestoreDays sets RestoreDays field to given value.

### HasRestoreDays

`func (o *ObjectStorageBucketUpdateReqBucket) HasRestoreDays() bool`

HasRestoreDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


