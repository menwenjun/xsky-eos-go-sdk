# ObjectStorageBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**AllUserPermission** | Pointer to **string** |  | [optional] 
**AuthUserPermission** | Pointer to **string** |  | [optional] 
**BucketPolicy** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DeleteArchiveStorageClass** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**ExternalQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**Flag** | Pointer to [**BucketFlag**](BucketFlag.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Lifecycle** | Pointer to [**ObjectStorageLifecycle**](ObjectStorageLifecycle.md) |  | [optional] 
**LocalQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**LocalQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**LogDeliveryPermission** | Pointer to **string** |  | [optional] 
**LoggingEnabled** | Pointer to **bool** |  | [optional] 
**LoggingGrants** | Pointer to [**[]OSBucketLoggingGrant**](OSBucketLoggingGrant.md) |  | [optional] 
**LoggingOwner** | Pointer to [**ObjectStorageUserNestview**](ObjectStorageUserNestview.md) |  | [optional] 
**LoggingPrefix** | Pointer to **string** |  | [optional] 
**LoggingSuspended** | Pointer to **bool** |  | [optional] 
**LoggingTargetBucket** | Pointer to [**ObjectStorageBucketNestview**](ObjectStorageBucketNestview.md) |  | [optional] 
**MetadataSearchEnabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NfsClientNum** | Pointer to **int64** |  | [optional] 
**ObjectCover** | Pointer to [**OSBucketObjectCoverConf**](OSBucketObjectCoverConf.md) |  | [optional] 
**ObjectStorageClass** | Pointer to [**OSBucketObjectStorageClass**](OSBucketObjectStorageClass.md) |  | [optional] 
**OsReplicationPathNum** | Pointer to **int64** |  | [optional] 
**OsReplicationZoneNum** | Pointer to **int64** |  | [optional] 
**OsZone** | Pointer to [**ObjectStorageZoneNestview**](ObjectStorageZoneNestview.md) |  | [optional] 
**OsZoneUuid** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**ObjectStorageUserNestview**](ObjectStorageUserNestview.md) |  | [optional] 
**OwnerPermission** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to [**ObjectStoragePolicyNestview**](ObjectStoragePolicyNestview.md) |  | [optional] 
**PolicyEnabled** | Pointer to **bool** |  | [optional] 
**Qos** | Pointer to [**OSBucketQos**](OSBucketQos.md) |  | [optional] 
**QuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**QuotaMaxSize** | Pointer to **int64** |  | [optional] 
**RemoteCluster** | Pointer to [**RemoteClusterNestview**](RemoteClusterNestview.md) |  | [optional] 
**ReplicationUuid** | Pointer to **string** | NOTE: Use name of bucket as replication uuid for simplicity | [optional] 
**RestoreDays** | Pointer to **int64** |  | [optional] 
**Shards** | Pointer to **int64** |  | [optional] 
**SpecificationObjects** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Virtual** | Pointer to **bool** |  | [optional] 

## Methods

### NewObjectStorageBucket

`func NewObjectStorageBucket() *ObjectStorageBucket`

NewObjectStorageBucket instantiates a new ObjectStorageBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageBucketWithDefaults

`func NewObjectStorageBucketWithDefaults() *ObjectStorageBucket`

NewObjectStorageBucketWithDefaults instantiates a new ObjectStorageBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *ObjectStorageBucket) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *ObjectStorageBucket) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *ObjectStorageBucket) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *ObjectStorageBucket) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAllUserPermission

`func (o *ObjectStorageBucket) GetAllUserPermission() string`

GetAllUserPermission returns the AllUserPermission field if non-nil, zero value otherwise.

### GetAllUserPermissionOk

`func (o *ObjectStorageBucket) GetAllUserPermissionOk() (*string, bool)`

GetAllUserPermissionOk returns a tuple with the AllUserPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUserPermission

`func (o *ObjectStorageBucket) SetAllUserPermission(v string)`

SetAllUserPermission sets AllUserPermission field to given value.

### HasAllUserPermission

`func (o *ObjectStorageBucket) HasAllUserPermission() bool`

HasAllUserPermission returns a boolean if a field has been set.

### GetAuthUserPermission

`func (o *ObjectStorageBucket) GetAuthUserPermission() string`

GetAuthUserPermission returns the AuthUserPermission field if non-nil, zero value otherwise.

### GetAuthUserPermissionOk

`func (o *ObjectStorageBucket) GetAuthUserPermissionOk() (*string, bool)`

GetAuthUserPermissionOk returns a tuple with the AuthUserPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUserPermission

`func (o *ObjectStorageBucket) SetAuthUserPermission(v string)`

SetAuthUserPermission sets AuthUserPermission field to given value.

### HasAuthUserPermission

`func (o *ObjectStorageBucket) HasAuthUserPermission() bool`

HasAuthUserPermission returns a boolean if a field has been set.

### GetBucketPolicy

`func (o *ObjectStorageBucket) GetBucketPolicy() string`

GetBucketPolicy returns the BucketPolicy field if non-nil, zero value otherwise.

### GetBucketPolicyOk

`func (o *ObjectStorageBucket) GetBucketPolicyOk() (*string, bool)`

GetBucketPolicyOk returns a tuple with the BucketPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPolicy

`func (o *ObjectStorageBucket) SetBucketPolicy(v string)`

SetBucketPolicy sets BucketPolicy field to given value.

### HasBucketPolicy

`func (o *ObjectStorageBucket) HasBucketPolicy() bool`

HasBucketPolicy returns a boolean if a field has been set.

### GetCluster

`func (o *ObjectStorageBucket) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ObjectStorageBucket) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ObjectStorageBucket) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ObjectStorageBucket) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *ObjectStorageBucket) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ObjectStorageBucket) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ObjectStorageBucket) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ObjectStorageBucket) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDeleteArchiveStorageClass

`func (o *ObjectStorageBucket) GetDeleteArchiveStorageClass() string`

GetDeleteArchiveStorageClass returns the DeleteArchiveStorageClass field if non-nil, zero value otherwise.

### GetDeleteArchiveStorageClassOk

`func (o *ObjectStorageBucket) GetDeleteArchiveStorageClassOk() (*string, bool)`

GetDeleteArchiveStorageClassOk returns a tuple with the DeleteArchiveStorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteArchiveStorageClass

`func (o *ObjectStorageBucket) SetDeleteArchiveStorageClass(v string)`

SetDeleteArchiveStorageClass sets DeleteArchiveStorageClass field to given value.

### HasDeleteArchiveStorageClass

`func (o *ObjectStorageBucket) HasDeleteArchiveStorageClass() bool`

HasDeleteArchiveStorageClass returns a boolean if a field has been set.

### GetDescription

`func (o *ObjectStorageBucket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectStorageBucket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectStorageBucket) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectStorageBucket) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalQuotaMaxObjects

`func (o *ObjectStorageBucket) GetExternalQuotaMaxObjects() int64`

GetExternalQuotaMaxObjects returns the ExternalQuotaMaxObjects field if non-nil, zero value otherwise.

### GetExternalQuotaMaxObjectsOk

`func (o *ObjectStorageBucket) GetExternalQuotaMaxObjectsOk() (*int64, bool)`

GetExternalQuotaMaxObjectsOk returns a tuple with the ExternalQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalQuotaMaxObjects

`func (o *ObjectStorageBucket) SetExternalQuotaMaxObjects(v int64)`

SetExternalQuotaMaxObjects sets ExternalQuotaMaxObjects field to given value.

### HasExternalQuotaMaxObjects

`func (o *ObjectStorageBucket) HasExternalQuotaMaxObjects() bool`

HasExternalQuotaMaxObjects returns a boolean if a field has been set.

### GetExternalQuotaMaxSize

`func (o *ObjectStorageBucket) GetExternalQuotaMaxSize() int64`

GetExternalQuotaMaxSize returns the ExternalQuotaMaxSize field if non-nil, zero value otherwise.

### GetExternalQuotaMaxSizeOk

`func (o *ObjectStorageBucket) GetExternalQuotaMaxSizeOk() (*int64, bool)`

GetExternalQuotaMaxSizeOk returns a tuple with the ExternalQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalQuotaMaxSize

`func (o *ObjectStorageBucket) SetExternalQuotaMaxSize(v int64)`

SetExternalQuotaMaxSize sets ExternalQuotaMaxSize field to given value.

### HasExternalQuotaMaxSize

`func (o *ObjectStorageBucket) HasExternalQuotaMaxSize() bool`

HasExternalQuotaMaxSize returns a boolean if a field has been set.

### GetFlag

`func (o *ObjectStorageBucket) GetFlag() BucketFlag`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *ObjectStorageBucket) GetFlagOk() (*BucketFlag, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *ObjectStorageBucket) SetFlag(v BucketFlag)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *ObjectStorageBucket) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetId

`func (o *ObjectStorageBucket) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStorageBucket) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStorageBucket) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStorageBucket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLifecycle

`func (o *ObjectStorageBucket) GetLifecycle() ObjectStorageLifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *ObjectStorageBucket) GetLifecycleOk() (*ObjectStorageLifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *ObjectStorageBucket) SetLifecycle(v ObjectStorageLifecycle)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *ObjectStorageBucket) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetLocalQuotaMaxObjects

`func (o *ObjectStorageBucket) GetLocalQuotaMaxObjects() int64`

GetLocalQuotaMaxObjects returns the LocalQuotaMaxObjects field if non-nil, zero value otherwise.

### GetLocalQuotaMaxObjectsOk

`func (o *ObjectStorageBucket) GetLocalQuotaMaxObjectsOk() (*int64, bool)`

GetLocalQuotaMaxObjectsOk returns a tuple with the LocalQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalQuotaMaxObjects

`func (o *ObjectStorageBucket) SetLocalQuotaMaxObjects(v int64)`

SetLocalQuotaMaxObjects sets LocalQuotaMaxObjects field to given value.

### HasLocalQuotaMaxObjects

`func (o *ObjectStorageBucket) HasLocalQuotaMaxObjects() bool`

HasLocalQuotaMaxObjects returns a boolean if a field has been set.

### GetLocalQuotaMaxSize

`func (o *ObjectStorageBucket) GetLocalQuotaMaxSize() int64`

GetLocalQuotaMaxSize returns the LocalQuotaMaxSize field if non-nil, zero value otherwise.

### GetLocalQuotaMaxSizeOk

`func (o *ObjectStorageBucket) GetLocalQuotaMaxSizeOk() (*int64, bool)`

GetLocalQuotaMaxSizeOk returns a tuple with the LocalQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalQuotaMaxSize

`func (o *ObjectStorageBucket) SetLocalQuotaMaxSize(v int64)`

SetLocalQuotaMaxSize sets LocalQuotaMaxSize field to given value.

### HasLocalQuotaMaxSize

`func (o *ObjectStorageBucket) HasLocalQuotaMaxSize() bool`

HasLocalQuotaMaxSize returns a boolean if a field has been set.

### GetLogDeliveryPermission

`func (o *ObjectStorageBucket) GetLogDeliveryPermission() string`

GetLogDeliveryPermission returns the LogDeliveryPermission field if non-nil, zero value otherwise.

### GetLogDeliveryPermissionOk

`func (o *ObjectStorageBucket) GetLogDeliveryPermissionOk() (*string, bool)`

GetLogDeliveryPermissionOk returns a tuple with the LogDeliveryPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDeliveryPermission

`func (o *ObjectStorageBucket) SetLogDeliveryPermission(v string)`

SetLogDeliveryPermission sets LogDeliveryPermission field to given value.

### HasLogDeliveryPermission

`func (o *ObjectStorageBucket) HasLogDeliveryPermission() bool`

HasLogDeliveryPermission returns a boolean if a field has been set.

### GetLoggingEnabled

`func (o *ObjectStorageBucket) GetLoggingEnabled() bool`

GetLoggingEnabled returns the LoggingEnabled field if non-nil, zero value otherwise.

### GetLoggingEnabledOk

`func (o *ObjectStorageBucket) GetLoggingEnabledOk() (*bool, bool)`

GetLoggingEnabledOk returns a tuple with the LoggingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingEnabled

`func (o *ObjectStorageBucket) SetLoggingEnabled(v bool)`

SetLoggingEnabled sets LoggingEnabled field to given value.

### HasLoggingEnabled

`func (o *ObjectStorageBucket) HasLoggingEnabled() bool`

HasLoggingEnabled returns a boolean if a field has been set.

### GetLoggingGrants

`func (o *ObjectStorageBucket) GetLoggingGrants() []OSBucketLoggingGrant`

GetLoggingGrants returns the LoggingGrants field if non-nil, zero value otherwise.

### GetLoggingGrantsOk

`func (o *ObjectStorageBucket) GetLoggingGrantsOk() (*[]OSBucketLoggingGrant, bool)`

GetLoggingGrantsOk returns a tuple with the LoggingGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingGrants

`func (o *ObjectStorageBucket) SetLoggingGrants(v []OSBucketLoggingGrant)`

SetLoggingGrants sets LoggingGrants field to given value.

### HasLoggingGrants

`func (o *ObjectStorageBucket) HasLoggingGrants() bool`

HasLoggingGrants returns a boolean if a field has been set.

### GetLoggingOwner

`func (o *ObjectStorageBucket) GetLoggingOwner() ObjectStorageUserNestview`

GetLoggingOwner returns the LoggingOwner field if non-nil, zero value otherwise.

### GetLoggingOwnerOk

`func (o *ObjectStorageBucket) GetLoggingOwnerOk() (*ObjectStorageUserNestview, bool)`

GetLoggingOwnerOk returns a tuple with the LoggingOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingOwner

`func (o *ObjectStorageBucket) SetLoggingOwner(v ObjectStorageUserNestview)`

SetLoggingOwner sets LoggingOwner field to given value.

### HasLoggingOwner

`func (o *ObjectStorageBucket) HasLoggingOwner() bool`

HasLoggingOwner returns a boolean if a field has been set.

### GetLoggingPrefix

`func (o *ObjectStorageBucket) GetLoggingPrefix() string`

GetLoggingPrefix returns the LoggingPrefix field if non-nil, zero value otherwise.

### GetLoggingPrefixOk

`func (o *ObjectStorageBucket) GetLoggingPrefixOk() (*string, bool)`

GetLoggingPrefixOk returns a tuple with the LoggingPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingPrefix

`func (o *ObjectStorageBucket) SetLoggingPrefix(v string)`

SetLoggingPrefix sets LoggingPrefix field to given value.

### HasLoggingPrefix

`func (o *ObjectStorageBucket) HasLoggingPrefix() bool`

HasLoggingPrefix returns a boolean if a field has been set.

### GetLoggingSuspended

`func (o *ObjectStorageBucket) GetLoggingSuspended() bool`

GetLoggingSuspended returns the LoggingSuspended field if non-nil, zero value otherwise.

### GetLoggingSuspendedOk

`func (o *ObjectStorageBucket) GetLoggingSuspendedOk() (*bool, bool)`

GetLoggingSuspendedOk returns a tuple with the LoggingSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingSuspended

`func (o *ObjectStorageBucket) SetLoggingSuspended(v bool)`

SetLoggingSuspended sets LoggingSuspended field to given value.

### HasLoggingSuspended

`func (o *ObjectStorageBucket) HasLoggingSuspended() bool`

HasLoggingSuspended returns a boolean if a field has been set.

### GetLoggingTargetBucket

`func (o *ObjectStorageBucket) GetLoggingTargetBucket() ObjectStorageBucketNestview`

GetLoggingTargetBucket returns the LoggingTargetBucket field if non-nil, zero value otherwise.

### GetLoggingTargetBucketOk

`func (o *ObjectStorageBucket) GetLoggingTargetBucketOk() (*ObjectStorageBucketNestview, bool)`

GetLoggingTargetBucketOk returns a tuple with the LoggingTargetBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingTargetBucket

`func (o *ObjectStorageBucket) SetLoggingTargetBucket(v ObjectStorageBucketNestview)`

SetLoggingTargetBucket sets LoggingTargetBucket field to given value.

### HasLoggingTargetBucket

`func (o *ObjectStorageBucket) HasLoggingTargetBucket() bool`

HasLoggingTargetBucket returns a boolean if a field has been set.

### GetMetadataSearchEnabled

`func (o *ObjectStorageBucket) GetMetadataSearchEnabled() bool`

GetMetadataSearchEnabled returns the MetadataSearchEnabled field if non-nil, zero value otherwise.

### GetMetadataSearchEnabledOk

`func (o *ObjectStorageBucket) GetMetadataSearchEnabledOk() (*bool, bool)`

GetMetadataSearchEnabledOk returns a tuple with the MetadataSearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSearchEnabled

`func (o *ObjectStorageBucket) SetMetadataSearchEnabled(v bool)`

SetMetadataSearchEnabled sets MetadataSearchEnabled field to given value.

### HasMetadataSearchEnabled

`func (o *ObjectStorageBucket) HasMetadataSearchEnabled() bool`

HasMetadataSearchEnabled returns a boolean if a field has been set.

### GetName

`func (o *ObjectStorageBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStorageBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStorageBucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectStorageBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNfsClientNum

`func (o *ObjectStorageBucket) GetNfsClientNum() int64`

GetNfsClientNum returns the NfsClientNum field if non-nil, zero value otherwise.

### GetNfsClientNumOk

`func (o *ObjectStorageBucket) GetNfsClientNumOk() (*int64, bool)`

GetNfsClientNumOk returns a tuple with the NfsClientNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsClientNum

`func (o *ObjectStorageBucket) SetNfsClientNum(v int64)`

SetNfsClientNum sets NfsClientNum field to given value.

### HasNfsClientNum

`func (o *ObjectStorageBucket) HasNfsClientNum() bool`

HasNfsClientNum returns a boolean if a field has been set.

### GetObjectCover

`func (o *ObjectStorageBucket) GetObjectCover() OSBucketObjectCoverConf`

GetObjectCover returns the ObjectCover field if non-nil, zero value otherwise.

### GetObjectCoverOk

`func (o *ObjectStorageBucket) GetObjectCoverOk() (*OSBucketObjectCoverConf, bool)`

GetObjectCoverOk returns a tuple with the ObjectCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCover

`func (o *ObjectStorageBucket) SetObjectCover(v OSBucketObjectCoverConf)`

SetObjectCover sets ObjectCover field to given value.

### HasObjectCover

`func (o *ObjectStorageBucket) HasObjectCover() bool`

HasObjectCover returns a boolean if a field has been set.

### GetObjectStorageClass

`func (o *ObjectStorageBucket) GetObjectStorageClass() OSBucketObjectStorageClass`

GetObjectStorageClass returns the ObjectStorageClass field if non-nil, zero value otherwise.

### GetObjectStorageClassOk

`func (o *ObjectStorageBucket) GetObjectStorageClassOk() (*OSBucketObjectStorageClass, bool)`

GetObjectStorageClassOk returns a tuple with the ObjectStorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorageClass

`func (o *ObjectStorageBucket) SetObjectStorageClass(v OSBucketObjectStorageClass)`

SetObjectStorageClass sets ObjectStorageClass field to given value.

### HasObjectStorageClass

`func (o *ObjectStorageBucket) HasObjectStorageClass() bool`

HasObjectStorageClass returns a boolean if a field has been set.

### GetOsReplicationPathNum

`func (o *ObjectStorageBucket) GetOsReplicationPathNum() int64`

GetOsReplicationPathNum returns the OsReplicationPathNum field if non-nil, zero value otherwise.

### GetOsReplicationPathNumOk

`func (o *ObjectStorageBucket) GetOsReplicationPathNumOk() (*int64, bool)`

GetOsReplicationPathNumOk returns a tuple with the OsReplicationPathNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsReplicationPathNum

`func (o *ObjectStorageBucket) SetOsReplicationPathNum(v int64)`

SetOsReplicationPathNum sets OsReplicationPathNum field to given value.

### HasOsReplicationPathNum

`func (o *ObjectStorageBucket) HasOsReplicationPathNum() bool`

HasOsReplicationPathNum returns a boolean if a field has been set.

### GetOsReplicationZoneNum

`func (o *ObjectStorageBucket) GetOsReplicationZoneNum() int64`

GetOsReplicationZoneNum returns the OsReplicationZoneNum field if non-nil, zero value otherwise.

### GetOsReplicationZoneNumOk

`func (o *ObjectStorageBucket) GetOsReplicationZoneNumOk() (*int64, bool)`

GetOsReplicationZoneNumOk returns a tuple with the OsReplicationZoneNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsReplicationZoneNum

`func (o *ObjectStorageBucket) SetOsReplicationZoneNum(v int64)`

SetOsReplicationZoneNum sets OsReplicationZoneNum field to given value.

### HasOsReplicationZoneNum

`func (o *ObjectStorageBucket) HasOsReplicationZoneNum() bool`

HasOsReplicationZoneNum returns a boolean if a field has been set.

### GetOsZone

`func (o *ObjectStorageBucket) GetOsZone() ObjectStorageZoneNestview`

GetOsZone returns the OsZone field if non-nil, zero value otherwise.

### GetOsZoneOk

`func (o *ObjectStorageBucket) GetOsZoneOk() (*ObjectStorageZoneNestview, bool)`

GetOsZoneOk returns a tuple with the OsZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsZone

`func (o *ObjectStorageBucket) SetOsZone(v ObjectStorageZoneNestview)`

SetOsZone sets OsZone field to given value.

### HasOsZone

`func (o *ObjectStorageBucket) HasOsZone() bool`

HasOsZone returns a boolean if a field has been set.

### GetOsZoneUuid

`func (o *ObjectStorageBucket) GetOsZoneUuid() string`

GetOsZoneUuid returns the OsZoneUuid field if non-nil, zero value otherwise.

### GetOsZoneUuidOk

`func (o *ObjectStorageBucket) GetOsZoneUuidOk() (*string, bool)`

GetOsZoneUuidOk returns a tuple with the OsZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsZoneUuid

`func (o *ObjectStorageBucket) SetOsZoneUuid(v string)`

SetOsZoneUuid sets OsZoneUuid field to given value.

### HasOsZoneUuid

`func (o *ObjectStorageBucket) HasOsZoneUuid() bool`

HasOsZoneUuid returns a boolean if a field has been set.

### GetOwner

`func (o *ObjectStorageBucket) GetOwner() ObjectStorageUserNestview`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ObjectStorageBucket) GetOwnerOk() (*ObjectStorageUserNestview, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ObjectStorageBucket) SetOwner(v ObjectStorageUserNestview)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ObjectStorageBucket) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnerPermission

`func (o *ObjectStorageBucket) GetOwnerPermission() string`

GetOwnerPermission returns the OwnerPermission field if non-nil, zero value otherwise.

### GetOwnerPermissionOk

`func (o *ObjectStorageBucket) GetOwnerPermissionOk() (*string, bool)`

GetOwnerPermissionOk returns a tuple with the OwnerPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPermission

`func (o *ObjectStorageBucket) SetOwnerPermission(v string)`

SetOwnerPermission sets OwnerPermission field to given value.

### HasOwnerPermission

`func (o *ObjectStorageBucket) HasOwnerPermission() bool`

HasOwnerPermission returns a boolean if a field has been set.

### GetPolicy

`func (o *ObjectStorageBucket) GetPolicy() ObjectStoragePolicyNestview`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ObjectStorageBucket) GetPolicyOk() (*ObjectStoragePolicyNestview, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ObjectStorageBucket) SetPolicy(v ObjectStoragePolicyNestview)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ObjectStorageBucket) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetPolicyEnabled

`func (o *ObjectStorageBucket) GetPolicyEnabled() bool`

GetPolicyEnabled returns the PolicyEnabled field if non-nil, zero value otherwise.

### GetPolicyEnabledOk

`func (o *ObjectStorageBucket) GetPolicyEnabledOk() (*bool, bool)`

GetPolicyEnabledOk returns a tuple with the PolicyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEnabled

`func (o *ObjectStorageBucket) SetPolicyEnabled(v bool)`

SetPolicyEnabled sets PolicyEnabled field to given value.

### HasPolicyEnabled

`func (o *ObjectStorageBucket) HasPolicyEnabled() bool`

HasPolicyEnabled returns a boolean if a field has been set.

### GetQos

`func (o *ObjectStorageBucket) GetQos() OSBucketQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *ObjectStorageBucket) GetQosOk() (*OSBucketQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *ObjectStorageBucket) SetQos(v OSBucketQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *ObjectStorageBucket) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetQuotaMaxObjects

`func (o *ObjectStorageBucket) GetQuotaMaxObjects() int64`

GetQuotaMaxObjects returns the QuotaMaxObjects field if non-nil, zero value otherwise.

### GetQuotaMaxObjectsOk

`func (o *ObjectStorageBucket) GetQuotaMaxObjectsOk() (*int64, bool)`

GetQuotaMaxObjectsOk returns a tuple with the QuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMaxObjects

`func (o *ObjectStorageBucket) SetQuotaMaxObjects(v int64)`

SetQuotaMaxObjects sets QuotaMaxObjects field to given value.

### HasQuotaMaxObjects

`func (o *ObjectStorageBucket) HasQuotaMaxObjects() bool`

HasQuotaMaxObjects returns a boolean if a field has been set.

### GetQuotaMaxSize

`func (o *ObjectStorageBucket) GetQuotaMaxSize() int64`

GetQuotaMaxSize returns the QuotaMaxSize field if non-nil, zero value otherwise.

### GetQuotaMaxSizeOk

`func (o *ObjectStorageBucket) GetQuotaMaxSizeOk() (*int64, bool)`

GetQuotaMaxSizeOk returns a tuple with the QuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMaxSize

`func (o *ObjectStorageBucket) SetQuotaMaxSize(v int64)`

SetQuotaMaxSize sets QuotaMaxSize field to given value.

### HasQuotaMaxSize

`func (o *ObjectStorageBucket) HasQuotaMaxSize() bool`

HasQuotaMaxSize returns a boolean if a field has been set.

### GetRemoteCluster

`func (o *ObjectStorageBucket) GetRemoteCluster() RemoteClusterNestview`

GetRemoteCluster returns the RemoteCluster field if non-nil, zero value otherwise.

### GetRemoteClusterOk

`func (o *ObjectStorageBucket) GetRemoteClusterOk() (*RemoteClusterNestview, bool)`

GetRemoteClusterOk returns a tuple with the RemoteCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCluster

`func (o *ObjectStorageBucket) SetRemoteCluster(v RemoteClusterNestview)`

SetRemoteCluster sets RemoteCluster field to given value.

### HasRemoteCluster

`func (o *ObjectStorageBucket) HasRemoteCluster() bool`

HasRemoteCluster returns a boolean if a field has been set.

### GetReplicationUuid

`func (o *ObjectStorageBucket) GetReplicationUuid() string`

GetReplicationUuid returns the ReplicationUuid field if non-nil, zero value otherwise.

### GetReplicationUuidOk

`func (o *ObjectStorageBucket) GetReplicationUuidOk() (*string, bool)`

GetReplicationUuidOk returns a tuple with the ReplicationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationUuid

`func (o *ObjectStorageBucket) SetReplicationUuid(v string)`

SetReplicationUuid sets ReplicationUuid field to given value.

### HasReplicationUuid

`func (o *ObjectStorageBucket) HasReplicationUuid() bool`

HasReplicationUuid returns a boolean if a field has been set.

### GetRestoreDays

`func (o *ObjectStorageBucket) GetRestoreDays() int64`

GetRestoreDays returns the RestoreDays field if non-nil, zero value otherwise.

### GetRestoreDaysOk

`func (o *ObjectStorageBucket) GetRestoreDaysOk() (*int64, bool)`

GetRestoreDaysOk returns a tuple with the RestoreDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreDays

`func (o *ObjectStorageBucket) SetRestoreDays(v int64)`

SetRestoreDays sets RestoreDays field to given value.

### HasRestoreDays

`func (o *ObjectStorageBucket) HasRestoreDays() bool`

HasRestoreDays returns a boolean if a field has been set.

### GetShards

`func (o *ObjectStorageBucket) GetShards() int64`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *ObjectStorageBucket) GetShardsOk() (*int64, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *ObjectStorageBucket) SetShards(v int64)`

SetShards sets Shards field to given value.

### HasShards

`func (o *ObjectStorageBucket) HasShards() bool`

HasShards returns a boolean if a field has been set.

### GetSpecificationObjects

`func (o *ObjectStorageBucket) GetSpecificationObjects() int64`

GetSpecificationObjects returns the SpecificationObjects field if non-nil, zero value otherwise.

### GetSpecificationObjectsOk

`func (o *ObjectStorageBucket) GetSpecificationObjectsOk() (*int64, bool)`

GetSpecificationObjectsOk returns a tuple with the SpecificationObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificationObjects

`func (o *ObjectStorageBucket) SetSpecificationObjects(v int64)`

SetSpecificationObjects sets SpecificationObjects field to given value.

### HasSpecificationObjects

`func (o *ObjectStorageBucket) HasSpecificationObjects() bool`

HasSpecificationObjects returns a boolean if a field has been set.

### GetStatus

`func (o *ObjectStorageBucket) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStorageBucket) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStorageBucket) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectStorageBucket) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *ObjectStorageBucket) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ObjectStorageBucket) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ObjectStorageBucket) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ObjectStorageBucket) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVirtual

`func (o *ObjectStorageBucket) GetVirtual() bool`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *ObjectStorageBucket) GetVirtualOk() (*bool, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *ObjectStorageBucket) SetVirtual(v bool)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *ObjectStorageBucket) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


