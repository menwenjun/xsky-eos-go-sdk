# ObjectStorageBucketCreateReqBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllUserPermission** | Pointer to **string** | permission setting of all users | [optional] 
**AuthUserPermission** | Pointer to **string** | permission setting of authenticated users | [optional] 
**DeleteArchiveStorageClass** | Pointer to **string** | delete archive storage class | [optional] 
**Description** | Pointer to **string** | bucket description | [optional] 
**ExternalQuotaMaxObjects** | Pointer to **int64** | max number of external storage objects | [optional] 
**ExternalQuotaMaxSize** | Pointer to **int64** | max size of external storage objects | [optional] 
**Flag** | Pointer to [**BucketFlag**](BucketFlag.md) |  | [optional] 
**LocalQuotaMaxObjects** | Pointer to **int64** | max number of local storage objects | [optional] 
**LocalQuotaMaxSize** | Pointer to **int64** | max size of local storage objects | [optional] 
**LogDeliveryPermission** | Pointer to **string** | permission setting of log delivery group | [optional] 
**Name** | **string** | bucket name | 
**ObjectCover** | Pointer to [**OSBucketObjectCoverConf**](OSBucketObjectCoverConf.md) |  | [optional] 
**ObjectStorageClass** | Pointer to [**OSBucketObjectStorageClass**](OSBucketObjectStorageClass.md) |  | [optional] 
**OwnerId** | **int64** | bucket owner | 
**OwnerPermission** | Pointer to **string** | permission setting of owner | [optional] 
**PolicyId** | **int64** | storage policy | 
**QuotaMaxObjects** | Pointer to **int64** | max number of objects | [optional] 
**QuotaMaxSize** | Pointer to **int64** | max size of all objects | [optional] 
**RestoreDays** | Pointer to **int64** | restore expiration days | [optional] 
**SpecificationObjects** | Pointer to **int64** | specification objects | [optional] 

## Methods

### NewObjectStorageBucketCreateReqBucket

`func NewObjectStorageBucketCreateReqBucket(name string, ownerId int64, policyId int64, ) *ObjectStorageBucketCreateReqBucket`

NewObjectStorageBucketCreateReqBucket instantiates a new ObjectStorageBucketCreateReqBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageBucketCreateReqBucketWithDefaults

`func NewObjectStorageBucketCreateReqBucketWithDefaults() *ObjectStorageBucketCreateReqBucket`

NewObjectStorageBucketCreateReqBucketWithDefaults instantiates a new ObjectStorageBucketCreateReqBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllUserPermission

`func (o *ObjectStorageBucketCreateReqBucket) GetAllUserPermission() string`

GetAllUserPermission returns the AllUserPermission field if non-nil, zero value otherwise.

### GetAllUserPermissionOk

`func (o *ObjectStorageBucketCreateReqBucket) GetAllUserPermissionOk() (*string, bool)`

GetAllUserPermissionOk returns a tuple with the AllUserPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUserPermission

`func (o *ObjectStorageBucketCreateReqBucket) SetAllUserPermission(v string)`

SetAllUserPermission sets AllUserPermission field to given value.

### HasAllUserPermission

`func (o *ObjectStorageBucketCreateReqBucket) HasAllUserPermission() bool`

HasAllUserPermission returns a boolean if a field has been set.

### GetAuthUserPermission

`func (o *ObjectStorageBucketCreateReqBucket) GetAuthUserPermission() string`

GetAuthUserPermission returns the AuthUserPermission field if non-nil, zero value otherwise.

### GetAuthUserPermissionOk

`func (o *ObjectStorageBucketCreateReqBucket) GetAuthUserPermissionOk() (*string, bool)`

GetAuthUserPermissionOk returns a tuple with the AuthUserPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUserPermission

`func (o *ObjectStorageBucketCreateReqBucket) SetAuthUserPermission(v string)`

SetAuthUserPermission sets AuthUserPermission field to given value.

### HasAuthUserPermission

`func (o *ObjectStorageBucketCreateReqBucket) HasAuthUserPermission() bool`

HasAuthUserPermission returns a boolean if a field has been set.

### GetDeleteArchiveStorageClass

`func (o *ObjectStorageBucketCreateReqBucket) GetDeleteArchiveStorageClass() string`

GetDeleteArchiveStorageClass returns the DeleteArchiveStorageClass field if non-nil, zero value otherwise.

### GetDeleteArchiveStorageClassOk

`func (o *ObjectStorageBucketCreateReqBucket) GetDeleteArchiveStorageClassOk() (*string, bool)`

GetDeleteArchiveStorageClassOk returns a tuple with the DeleteArchiveStorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteArchiveStorageClass

`func (o *ObjectStorageBucketCreateReqBucket) SetDeleteArchiveStorageClass(v string)`

SetDeleteArchiveStorageClass sets DeleteArchiveStorageClass field to given value.

### HasDeleteArchiveStorageClass

`func (o *ObjectStorageBucketCreateReqBucket) HasDeleteArchiveStorageClass() bool`

HasDeleteArchiveStorageClass returns a boolean if a field has been set.

### GetDescription

`func (o *ObjectStorageBucketCreateReqBucket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectStorageBucketCreateReqBucket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectStorageBucketCreateReqBucket) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectStorageBucketCreateReqBucket) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalQuotaMaxObjects

`func (o *ObjectStorageBucketCreateReqBucket) GetExternalQuotaMaxObjects() int64`

GetExternalQuotaMaxObjects returns the ExternalQuotaMaxObjects field if non-nil, zero value otherwise.

### GetExternalQuotaMaxObjectsOk

`func (o *ObjectStorageBucketCreateReqBucket) GetExternalQuotaMaxObjectsOk() (*int64, bool)`

GetExternalQuotaMaxObjectsOk returns a tuple with the ExternalQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalQuotaMaxObjects

`func (o *ObjectStorageBucketCreateReqBucket) SetExternalQuotaMaxObjects(v int64)`

SetExternalQuotaMaxObjects sets ExternalQuotaMaxObjects field to given value.

### HasExternalQuotaMaxObjects

`func (o *ObjectStorageBucketCreateReqBucket) HasExternalQuotaMaxObjects() bool`

HasExternalQuotaMaxObjects returns a boolean if a field has been set.

### GetExternalQuotaMaxSize

`func (o *ObjectStorageBucketCreateReqBucket) GetExternalQuotaMaxSize() int64`

GetExternalQuotaMaxSize returns the ExternalQuotaMaxSize field if non-nil, zero value otherwise.

### GetExternalQuotaMaxSizeOk

`func (o *ObjectStorageBucketCreateReqBucket) GetExternalQuotaMaxSizeOk() (*int64, bool)`

GetExternalQuotaMaxSizeOk returns a tuple with the ExternalQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalQuotaMaxSize

`func (o *ObjectStorageBucketCreateReqBucket) SetExternalQuotaMaxSize(v int64)`

SetExternalQuotaMaxSize sets ExternalQuotaMaxSize field to given value.

### HasExternalQuotaMaxSize

`func (o *ObjectStorageBucketCreateReqBucket) HasExternalQuotaMaxSize() bool`

HasExternalQuotaMaxSize returns a boolean if a field has been set.

### GetFlag

`func (o *ObjectStorageBucketCreateReqBucket) GetFlag() BucketFlag`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *ObjectStorageBucketCreateReqBucket) GetFlagOk() (*BucketFlag, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *ObjectStorageBucketCreateReqBucket) SetFlag(v BucketFlag)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *ObjectStorageBucketCreateReqBucket) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetLocalQuotaMaxObjects

`func (o *ObjectStorageBucketCreateReqBucket) GetLocalQuotaMaxObjects() int64`

GetLocalQuotaMaxObjects returns the LocalQuotaMaxObjects field if non-nil, zero value otherwise.

### GetLocalQuotaMaxObjectsOk

`func (o *ObjectStorageBucketCreateReqBucket) GetLocalQuotaMaxObjectsOk() (*int64, bool)`

GetLocalQuotaMaxObjectsOk returns a tuple with the LocalQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalQuotaMaxObjects

`func (o *ObjectStorageBucketCreateReqBucket) SetLocalQuotaMaxObjects(v int64)`

SetLocalQuotaMaxObjects sets LocalQuotaMaxObjects field to given value.

### HasLocalQuotaMaxObjects

`func (o *ObjectStorageBucketCreateReqBucket) HasLocalQuotaMaxObjects() bool`

HasLocalQuotaMaxObjects returns a boolean if a field has been set.

### GetLocalQuotaMaxSize

`func (o *ObjectStorageBucketCreateReqBucket) GetLocalQuotaMaxSize() int64`

GetLocalQuotaMaxSize returns the LocalQuotaMaxSize field if non-nil, zero value otherwise.

### GetLocalQuotaMaxSizeOk

`func (o *ObjectStorageBucketCreateReqBucket) GetLocalQuotaMaxSizeOk() (*int64, bool)`

GetLocalQuotaMaxSizeOk returns a tuple with the LocalQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalQuotaMaxSize

`func (o *ObjectStorageBucketCreateReqBucket) SetLocalQuotaMaxSize(v int64)`

SetLocalQuotaMaxSize sets LocalQuotaMaxSize field to given value.

### HasLocalQuotaMaxSize

`func (o *ObjectStorageBucketCreateReqBucket) HasLocalQuotaMaxSize() bool`

HasLocalQuotaMaxSize returns a boolean if a field has been set.

### GetLogDeliveryPermission

`func (o *ObjectStorageBucketCreateReqBucket) GetLogDeliveryPermission() string`

GetLogDeliveryPermission returns the LogDeliveryPermission field if non-nil, zero value otherwise.

### GetLogDeliveryPermissionOk

`func (o *ObjectStorageBucketCreateReqBucket) GetLogDeliveryPermissionOk() (*string, bool)`

GetLogDeliveryPermissionOk returns a tuple with the LogDeliveryPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDeliveryPermission

`func (o *ObjectStorageBucketCreateReqBucket) SetLogDeliveryPermission(v string)`

SetLogDeliveryPermission sets LogDeliveryPermission field to given value.

### HasLogDeliveryPermission

`func (o *ObjectStorageBucketCreateReqBucket) HasLogDeliveryPermission() bool`

HasLogDeliveryPermission returns a boolean if a field has been set.

### GetName

`func (o *ObjectStorageBucketCreateReqBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStorageBucketCreateReqBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStorageBucketCreateReqBucket) SetName(v string)`

SetName sets Name field to given value.


### GetObjectCover

`func (o *ObjectStorageBucketCreateReqBucket) GetObjectCover() OSBucketObjectCoverConf`

GetObjectCover returns the ObjectCover field if non-nil, zero value otherwise.

### GetObjectCoverOk

`func (o *ObjectStorageBucketCreateReqBucket) GetObjectCoverOk() (*OSBucketObjectCoverConf, bool)`

GetObjectCoverOk returns a tuple with the ObjectCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCover

`func (o *ObjectStorageBucketCreateReqBucket) SetObjectCover(v OSBucketObjectCoverConf)`

SetObjectCover sets ObjectCover field to given value.

### HasObjectCover

`func (o *ObjectStorageBucketCreateReqBucket) HasObjectCover() bool`

HasObjectCover returns a boolean if a field has been set.

### GetObjectStorageClass

`func (o *ObjectStorageBucketCreateReqBucket) GetObjectStorageClass() OSBucketObjectStorageClass`

GetObjectStorageClass returns the ObjectStorageClass field if non-nil, zero value otherwise.

### GetObjectStorageClassOk

`func (o *ObjectStorageBucketCreateReqBucket) GetObjectStorageClassOk() (*OSBucketObjectStorageClass, bool)`

GetObjectStorageClassOk returns a tuple with the ObjectStorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorageClass

`func (o *ObjectStorageBucketCreateReqBucket) SetObjectStorageClass(v OSBucketObjectStorageClass)`

SetObjectStorageClass sets ObjectStorageClass field to given value.

### HasObjectStorageClass

`func (o *ObjectStorageBucketCreateReqBucket) HasObjectStorageClass() bool`

HasObjectStorageClass returns a boolean if a field has been set.

### GetOwnerId

`func (o *ObjectStorageBucketCreateReqBucket) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ObjectStorageBucketCreateReqBucket) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ObjectStorageBucketCreateReqBucket) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.


### GetOwnerPermission

`func (o *ObjectStorageBucketCreateReqBucket) GetOwnerPermission() string`

GetOwnerPermission returns the OwnerPermission field if non-nil, zero value otherwise.

### GetOwnerPermissionOk

`func (o *ObjectStorageBucketCreateReqBucket) GetOwnerPermissionOk() (*string, bool)`

GetOwnerPermissionOk returns a tuple with the OwnerPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPermission

`func (o *ObjectStorageBucketCreateReqBucket) SetOwnerPermission(v string)`

SetOwnerPermission sets OwnerPermission field to given value.

### HasOwnerPermission

`func (o *ObjectStorageBucketCreateReqBucket) HasOwnerPermission() bool`

HasOwnerPermission returns a boolean if a field has been set.

### GetPolicyId

`func (o *ObjectStorageBucketCreateReqBucket) GetPolicyId() int64`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ObjectStorageBucketCreateReqBucket) GetPolicyIdOk() (*int64, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ObjectStorageBucketCreateReqBucket) SetPolicyId(v int64)`

SetPolicyId sets PolicyId field to given value.


### GetQuotaMaxObjects

`func (o *ObjectStorageBucketCreateReqBucket) GetQuotaMaxObjects() int64`

GetQuotaMaxObjects returns the QuotaMaxObjects field if non-nil, zero value otherwise.

### GetQuotaMaxObjectsOk

`func (o *ObjectStorageBucketCreateReqBucket) GetQuotaMaxObjectsOk() (*int64, bool)`

GetQuotaMaxObjectsOk returns a tuple with the QuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMaxObjects

`func (o *ObjectStorageBucketCreateReqBucket) SetQuotaMaxObjects(v int64)`

SetQuotaMaxObjects sets QuotaMaxObjects field to given value.

### HasQuotaMaxObjects

`func (o *ObjectStorageBucketCreateReqBucket) HasQuotaMaxObjects() bool`

HasQuotaMaxObjects returns a boolean if a field has been set.

### GetQuotaMaxSize

`func (o *ObjectStorageBucketCreateReqBucket) GetQuotaMaxSize() int64`

GetQuotaMaxSize returns the QuotaMaxSize field if non-nil, zero value otherwise.

### GetQuotaMaxSizeOk

`func (o *ObjectStorageBucketCreateReqBucket) GetQuotaMaxSizeOk() (*int64, bool)`

GetQuotaMaxSizeOk returns a tuple with the QuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMaxSize

`func (o *ObjectStorageBucketCreateReqBucket) SetQuotaMaxSize(v int64)`

SetQuotaMaxSize sets QuotaMaxSize field to given value.

### HasQuotaMaxSize

`func (o *ObjectStorageBucketCreateReqBucket) HasQuotaMaxSize() bool`

HasQuotaMaxSize returns a boolean if a field has been set.

### GetRestoreDays

`func (o *ObjectStorageBucketCreateReqBucket) GetRestoreDays() int64`

GetRestoreDays returns the RestoreDays field if non-nil, zero value otherwise.

### GetRestoreDaysOk

`func (o *ObjectStorageBucketCreateReqBucket) GetRestoreDaysOk() (*int64, bool)`

GetRestoreDaysOk returns a tuple with the RestoreDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreDays

`func (o *ObjectStorageBucketCreateReqBucket) SetRestoreDays(v int64)`

SetRestoreDays sets RestoreDays field to given value.

### HasRestoreDays

`func (o *ObjectStorageBucketCreateReqBucket) HasRestoreDays() bool`

HasRestoreDays returns a boolean if a field has been set.

### GetSpecificationObjects

`func (o *ObjectStorageBucketCreateReqBucket) GetSpecificationObjects() int64`

GetSpecificationObjects returns the SpecificationObjects field if non-nil, zero value otherwise.

### GetSpecificationObjectsOk

`func (o *ObjectStorageBucketCreateReqBucket) GetSpecificationObjectsOk() (*int64, bool)`

GetSpecificationObjectsOk returns a tuple with the SpecificationObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificationObjects

`func (o *ObjectStorageBucketCreateReqBucket) SetSpecificationObjects(v int64)`

SetSpecificationObjects sets SpecificationObjects field to given value.

### HasSpecificationObjects

`func (o *ObjectStorageBucketCreateReqBucket) HasSpecificationObjects() bool`

HasSpecificationObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


