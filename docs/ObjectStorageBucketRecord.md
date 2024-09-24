# ObjectStorageBucketRecord

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
**Samples** | Pointer to [**[]ObjectStorageBucketStat**](ObjectStorageBucketStat.md) |  | [optional] 

## Methods

### NewObjectStorageBucketRecord

`func NewObjectStorageBucketRecord() *ObjectStorageBucketRecord`

NewObjectStorageBucketRecord instantiates a new ObjectStorageBucketRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageBucketRecordWithDefaults

`func NewObjectStorageBucketRecordWithDefaults() *ObjectStorageBucketRecord`

NewObjectStorageBucketRecordWithDefaults instantiates a new ObjectStorageBucketRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *ObjectStorageBucketRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *ObjectStorageBucketRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *ObjectStorageBucketRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *ObjectStorageBucketRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAllUserPermission

`func (o *ObjectStorageBucketRecord) GetAllUserPermission() string`

GetAllUserPermission returns the AllUserPermission field if non-nil, zero value otherwise.

### GetAllUserPermissionOk

`func (o *ObjectStorageBucketRecord) GetAllUserPermissionOk() (*string, bool)`

GetAllUserPermissionOk returns a tuple with the AllUserPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUserPermission

`func (o *ObjectStorageBucketRecord) SetAllUserPermission(v string)`

SetAllUserPermission sets AllUserPermission field to given value.

### HasAllUserPermission

`func (o *ObjectStorageBucketRecord) HasAllUserPermission() bool`

HasAllUserPermission returns a boolean if a field has been set.

### GetAuthUserPermission

`func (o *ObjectStorageBucketRecord) GetAuthUserPermission() string`

GetAuthUserPermission returns the AuthUserPermission field if non-nil, zero value otherwise.

### GetAuthUserPermissionOk

`func (o *ObjectStorageBucketRecord) GetAuthUserPermissionOk() (*string, bool)`

GetAuthUserPermissionOk returns a tuple with the AuthUserPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUserPermission

`func (o *ObjectStorageBucketRecord) SetAuthUserPermission(v string)`

SetAuthUserPermission sets AuthUserPermission field to given value.

### HasAuthUserPermission

`func (o *ObjectStorageBucketRecord) HasAuthUserPermission() bool`

HasAuthUserPermission returns a boolean if a field has been set.

### GetBucketPolicy

`func (o *ObjectStorageBucketRecord) GetBucketPolicy() string`

GetBucketPolicy returns the BucketPolicy field if non-nil, zero value otherwise.

### GetBucketPolicyOk

`func (o *ObjectStorageBucketRecord) GetBucketPolicyOk() (*string, bool)`

GetBucketPolicyOk returns a tuple with the BucketPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPolicy

`func (o *ObjectStorageBucketRecord) SetBucketPolicy(v string)`

SetBucketPolicy sets BucketPolicy field to given value.

### HasBucketPolicy

`func (o *ObjectStorageBucketRecord) HasBucketPolicy() bool`

HasBucketPolicy returns a boolean if a field has been set.

### GetCluster

`func (o *ObjectStorageBucketRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ObjectStorageBucketRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ObjectStorageBucketRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ObjectStorageBucketRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *ObjectStorageBucketRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ObjectStorageBucketRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ObjectStorageBucketRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ObjectStorageBucketRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDeleteArchiveStorageClass

`func (o *ObjectStorageBucketRecord) GetDeleteArchiveStorageClass() string`

GetDeleteArchiveStorageClass returns the DeleteArchiveStorageClass field if non-nil, zero value otherwise.

### GetDeleteArchiveStorageClassOk

`func (o *ObjectStorageBucketRecord) GetDeleteArchiveStorageClassOk() (*string, bool)`

GetDeleteArchiveStorageClassOk returns a tuple with the DeleteArchiveStorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteArchiveStorageClass

`func (o *ObjectStorageBucketRecord) SetDeleteArchiveStorageClass(v string)`

SetDeleteArchiveStorageClass sets DeleteArchiveStorageClass field to given value.

### HasDeleteArchiveStorageClass

`func (o *ObjectStorageBucketRecord) HasDeleteArchiveStorageClass() bool`

HasDeleteArchiveStorageClass returns a boolean if a field has been set.

### GetDescription

`func (o *ObjectStorageBucketRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectStorageBucketRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectStorageBucketRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectStorageBucketRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalQuotaMaxObjects

`func (o *ObjectStorageBucketRecord) GetExternalQuotaMaxObjects() int64`

GetExternalQuotaMaxObjects returns the ExternalQuotaMaxObjects field if non-nil, zero value otherwise.

### GetExternalQuotaMaxObjectsOk

`func (o *ObjectStorageBucketRecord) GetExternalQuotaMaxObjectsOk() (*int64, bool)`

GetExternalQuotaMaxObjectsOk returns a tuple with the ExternalQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalQuotaMaxObjects

`func (o *ObjectStorageBucketRecord) SetExternalQuotaMaxObjects(v int64)`

SetExternalQuotaMaxObjects sets ExternalQuotaMaxObjects field to given value.

### HasExternalQuotaMaxObjects

`func (o *ObjectStorageBucketRecord) HasExternalQuotaMaxObjects() bool`

HasExternalQuotaMaxObjects returns a boolean if a field has been set.

### GetExternalQuotaMaxSize

`func (o *ObjectStorageBucketRecord) GetExternalQuotaMaxSize() int64`

GetExternalQuotaMaxSize returns the ExternalQuotaMaxSize field if non-nil, zero value otherwise.

### GetExternalQuotaMaxSizeOk

`func (o *ObjectStorageBucketRecord) GetExternalQuotaMaxSizeOk() (*int64, bool)`

GetExternalQuotaMaxSizeOk returns a tuple with the ExternalQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalQuotaMaxSize

`func (o *ObjectStorageBucketRecord) SetExternalQuotaMaxSize(v int64)`

SetExternalQuotaMaxSize sets ExternalQuotaMaxSize field to given value.

### HasExternalQuotaMaxSize

`func (o *ObjectStorageBucketRecord) HasExternalQuotaMaxSize() bool`

HasExternalQuotaMaxSize returns a boolean if a field has been set.

### GetFlag

`func (o *ObjectStorageBucketRecord) GetFlag() BucketFlag`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *ObjectStorageBucketRecord) GetFlagOk() (*BucketFlag, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *ObjectStorageBucketRecord) SetFlag(v BucketFlag)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *ObjectStorageBucketRecord) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetId

`func (o *ObjectStorageBucketRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStorageBucketRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStorageBucketRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStorageBucketRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLifecycle

`func (o *ObjectStorageBucketRecord) GetLifecycle() ObjectStorageLifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *ObjectStorageBucketRecord) GetLifecycleOk() (*ObjectStorageLifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *ObjectStorageBucketRecord) SetLifecycle(v ObjectStorageLifecycle)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *ObjectStorageBucketRecord) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetLocalQuotaMaxObjects

`func (o *ObjectStorageBucketRecord) GetLocalQuotaMaxObjects() int64`

GetLocalQuotaMaxObjects returns the LocalQuotaMaxObjects field if non-nil, zero value otherwise.

### GetLocalQuotaMaxObjectsOk

`func (o *ObjectStorageBucketRecord) GetLocalQuotaMaxObjectsOk() (*int64, bool)`

GetLocalQuotaMaxObjectsOk returns a tuple with the LocalQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalQuotaMaxObjects

`func (o *ObjectStorageBucketRecord) SetLocalQuotaMaxObjects(v int64)`

SetLocalQuotaMaxObjects sets LocalQuotaMaxObjects field to given value.

### HasLocalQuotaMaxObjects

`func (o *ObjectStorageBucketRecord) HasLocalQuotaMaxObjects() bool`

HasLocalQuotaMaxObjects returns a boolean if a field has been set.

### GetLocalQuotaMaxSize

`func (o *ObjectStorageBucketRecord) GetLocalQuotaMaxSize() int64`

GetLocalQuotaMaxSize returns the LocalQuotaMaxSize field if non-nil, zero value otherwise.

### GetLocalQuotaMaxSizeOk

`func (o *ObjectStorageBucketRecord) GetLocalQuotaMaxSizeOk() (*int64, bool)`

GetLocalQuotaMaxSizeOk returns a tuple with the LocalQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalQuotaMaxSize

`func (o *ObjectStorageBucketRecord) SetLocalQuotaMaxSize(v int64)`

SetLocalQuotaMaxSize sets LocalQuotaMaxSize field to given value.

### HasLocalQuotaMaxSize

`func (o *ObjectStorageBucketRecord) HasLocalQuotaMaxSize() bool`

HasLocalQuotaMaxSize returns a boolean if a field has been set.

### GetLogDeliveryPermission

`func (o *ObjectStorageBucketRecord) GetLogDeliveryPermission() string`

GetLogDeliveryPermission returns the LogDeliveryPermission field if non-nil, zero value otherwise.

### GetLogDeliveryPermissionOk

`func (o *ObjectStorageBucketRecord) GetLogDeliveryPermissionOk() (*string, bool)`

GetLogDeliveryPermissionOk returns a tuple with the LogDeliveryPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDeliveryPermission

`func (o *ObjectStorageBucketRecord) SetLogDeliveryPermission(v string)`

SetLogDeliveryPermission sets LogDeliveryPermission field to given value.

### HasLogDeliveryPermission

`func (o *ObjectStorageBucketRecord) HasLogDeliveryPermission() bool`

HasLogDeliveryPermission returns a boolean if a field has been set.

### GetLoggingEnabled

`func (o *ObjectStorageBucketRecord) GetLoggingEnabled() bool`

GetLoggingEnabled returns the LoggingEnabled field if non-nil, zero value otherwise.

### GetLoggingEnabledOk

`func (o *ObjectStorageBucketRecord) GetLoggingEnabledOk() (*bool, bool)`

GetLoggingEnabledOk returns a tuple with the LoggingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingEnabled

`func (o *ObjectStorageBucketRecord) SetLoggingEnabled(v bool)`

SetLoggingEnabled sets LoggingEnabled field to given value.

### HasLoggingEnabled

`func (o *ObjectStorageBucketRecord) HasLoggingEnabled() bool`

HasLoggingEnabled returns a boolean if a field has been set.

### GetLoggingGrants

`func (o *ObjectStorageBucketRecord) GetLoggingGrants() []OSBucketLoggingGrant`

GetLoggingGrants returns the LoggingGrants field if non-nil, zero value otherwise.

### GetLoggingGrantsOk

`func (o *ObjectStorageBucketRecord) GetLoggingGrantsOk() (*[]OSBucketLoggingGrant, bool)`

GetLoggingGrantsOk returns a tuple with the LoggingGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingGrants

`func (o *ObjectStorageBucketRecord) SetLoggingGrants(v []OSBucketLoggingGrant)`

SetLoggingGrants sets LoggingGrants field to given value.

### HasLoggingGrants

`func (o *ObjectStorageBucketRecord) HasLoggingGrants() bool`

HasLoggingGrants returns a boolean if a field has been set.

### GetLoggingOwner

`func (o *ObjectStorageBucketRecord) GetLoggingOwner() ObjectStorageUserNestview`

GetLoggingOwner returns the LoggingOwner field if non-nil, zero value otherwise.

### GetLoggingOwnerOk

`func (o *ObjectStorageBucketRecord) GetLoggingOwnerOk() (*ObjectStorageUserNestview, bool)`

GetLoggingOwnerOk returns a tuple with the LoggingOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingOwner

`func (o *ObjectStorageBucketRecord) SetLoggingOwner(v ObjectStorageUserNestview)`

SetLoggingOwner sets LoggingOwner field to given value.

### HasLoggingOwner

`func (o *ObjectStorageBucketRecord) HasLoggingOwner() bool`

HasLoggingOwner returns a boolean if a field has been set.

### GetLoggingPrefix

`func (o *ObjectStorageBucketRecord) GetLoggingPrefix() string`

GetLoggingPrefix returns the LoggingPrefix field if non-nil, zero value otherwise.

### GetLoggingPrefixOk

`func (o *ObjectStorageBucketRecord) GetLoggingPrefixOk() (*string, bool)`

GetLoggingPrefixOk returns a tuple with the LoggingPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingPrefix

`func (o *ObjectStorageBucketRecord) SetLoggingPrefix(v string)`

SetLoggingPrefix sets LoggingPrefix field to given value.

### HasLoggingPrefix

`func (o *ObjectStorageBucketRecord) HasLoggingPrefix() bool`

HasLoggingPrefix returns a boolean if a field has been set.

### GetLoggingSuspended

`func (o *ObjectStorageBucketRecord) GetLoggingSuspended() bool`

GetLoggingSuspended returns the LoggingSuspended field if non-nil, zero value otherwise.

### GetLoggingSuspendedOk

`func (o *ObjectStorageBucketRecord) GetLoggingSuspendedOk() (*bool, bool)`

GetLoggingSuspendedOk returns a tuple with the LoggingSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingSuspended

`func (o *ObjectStorageBucketRecord) SetLoggingSuspended(v bool)`

SetLoggingSuspended sets LoggingSuspended field to given value.

### HasLoggingSuspended

`func (o *ObjectStorageBucketRecord) HasLoggingSuspended() bool`

HasLoggingSuspended returns a boolean if a field has been set.

### GetLoggingTargetBucket

`func (o *ObjectStorageBucketRecord) GetLoggingTargetBucket() ObjectStorageBucketNestview`

GetLoggingTargetBucket returns the LoggingTargetBucket field if non-nil, zero value otherwise.

### GetLoggingTargetBucketOk

`func (o *ObjectStorageBucketRecord) GetLoggingTargetBucketOk() (*ObjectStorageBucketNestview, bool)`

GetLoggingTargetBucketOk returns a tuple with the LoggingTargetBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingTargetBucket

`func (o *ObjectStorageBucketRecord) SetLoggingTargetBucket(v ObjectStorageBucketNestview)`

SetLoggingTargetBucket sets LoggingTargetBucket field to given value.

### HasLoggingTargetBucket

`func (o *ObjectStorageBucketRecord) HasLoggingTargetBucket() bool`

HasLoggingTargetBucket returns a boolean if a field has been set.

### GetMetadataSearchEnabled

`func (o *ObjectStorageBucketRecord) GetMetadataSearchEnabled() bool`

GetMetadataSearchEnabled returns the MetadataSearchEnabled field if non-nil, zero value otherwise.

### GetMetadataSearchEnabledOk

`func (o *ObjectStorageBucketRecord) GetMetadataSearchEnabledOk() (*bool, bool)`

GetMetadataSearchEnabledOk returns a tuple with the MetadataSearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSearchEnabled

`func (o *ObjectStorageBucketRecord) SetMetadataSearchEnabled(v bool)`

SetMetadataSearchEnabled sets MetadataSearchEnabled field to given value.

### HasMetadataSearchEnabled

`func (o *ObjectStorageBucketRecord) HasMetadataSearchEnabled() bool`

HasMetadataSearchEnabled returns a boolean if a field has been set.

### GetName

`func (o *ObjectStorageBucketRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStorageBucketRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStorageBucketRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectStorageBucketRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNfsClientNum

`func (o *ObjectStorageBucketRecord) GetNfsClientNum() int64`

GetNfsClientNum returns the NfsClientNum field if non-nil, zero value otherwise.

### GetNfsClientNumOk

`func (o *ObjectStorageBucketRecord) GetNfsClientNumOk() (*int64, bool)`

GetNfsClientNumOk returns a tuple with the NfsClientNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsClientNum

`func (o *ObjectStorageBucketRecord) SetNfsClientNum(v int64)`

SetNfsClientNum sets NfsClientNum field to given value.

### HasNfsClientNum

`func (o *ObjectStorageBucketRecord) HasNfsClientNum() bool`

HasNfsClientNum returns a boolean if a field has been set.

### GetObjectCover

`func (o *ObjectStorageBucketRecord) GetObjectCover() OSBucketObjectCoverConf`

GetObjectCover returns the ObjectCover field if non-nil, zero value otherwise.

### GetObjectCoverOk

`func (o *ObjectStorageBucketRecord) GetObjectCoverOk() (*OSBucketObjectCoverConf, bool)`

GetObjectCoverOk returns a tuple with the ObjectCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCover

`func (o *ObjectStorageBucketRecord) SetObjectCover(v OSBucketObjectCoverConf)`

SetObjectCover sets ObjectCover field to given value.

### HasObjectCover

`func (o *ObjectStorageBucketRecord) HasObjectCover() bool`

HasObjectCover returns a boolean if a field has been set.

### GetObjectStorageClass

`func (o *ObjectStorageBucketRecord) GetObjectStorageClass() OSBucketObjectStorageClass`

GetObjectStorageClass returns the ObjectStorageClass field if non-nil, zero value otherwise.

### GetObjectStorageClassOk

`func (o *ObjectStorageBucketRecord) GetObjectStorageClassOk() (*OSBucketObjectStorageClass, bool)`

GetObjectStorageClassOk returns a tuple with the ObjectStorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorageClass

`func (o *ObjectStorageBucketRecord) SetObjectStorageClass(v OSBucketObjectStorageClass)`

SetObjectStorageClass sets ObjectStorageClass field to given value.

### HasObjectStorageClass

`func (o *ObjectStorageBucketRecord) HasObjectStorageClass() bool`

HasObjectStorageClass returns a boolean if a field has been set.

### GetOsReplicationPathNum

`func (o *ObjectStorageBucketRecord) GetOsReplicationPathNum() int64`

GetOsReplicationPathNum returns the OsReplicationPathNum field if non-nil, zero value otherwise.

### GetOsReplicationPathNumOk

`func (o *ObjectStorageBucketRecord) GetOsReplicationPathNumOk() (*int64, bool)`

GetOsReplicationPathNumOk returns a tuple with the OsReplicationPathNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsReplicationPathNum

`func (o *ObjectStorageBucketRecord) SetOsReplicationPathNum(v int64)`

SetOsReplicationPathNum sets OsReplicationPathNum field to given value.

### HasOsReplicationPathNum

`func (o *ObjectStorageBucketRecord) HasOsReplicationPathNum() bool`

HasOsReplicationPathNum returns a boolean if a field has been set.

### GetOsReplicationZoneNum

`func (o *ObjectStorageBucketRecord) GetOsReplicationZoneNum() int64`

GetOsReplicationZoneNum returns the OsReplicationZoneNum field if non-nil, zero value otherwise.

### GetOsReplicationZoneNumOk

`func (o *ObjectStorageBucketRecord) GetOsReplicationZoneNumOk() (*int64, bool)`

GetOsReplicationZoneNumOk returns a tuple with the OsReplicationZoneNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsReplicationZoneNum

`func (o *ObjectStorageBucketRecord) SetOsReplicationZoneNum(v int64)`

SetOsReplicationZoneNum sets OsReplicationZoneNum field to given value.

### HasOsReplicationZoneNum

`func (o *ObjectStorageBucketRecord) HasOsReplicationZoneNum() bool`

HasOsReplicationZoneNum returns a boolean if a field has been set.

### GetOsZone

`func (o *ObjectStorageBucketRecord) GetOsZone() ObjectStorageZoneNestview`

GetOsZone returns the OsZone field if non-nil, zero value otherwise.

### GetOsZoneOk

`func (o *ObjectStorageBucketRecord) GetOsZoneOk() (*ObjectStorageZoneNestview, bool)`

GetOsZoneOk returns a tuple with the OsZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsZone

`func (o *ObjectStorageBucketRecord) SetOsZone(v ObjectStorageZoneNestview)`

SetOsZone sets OsZone field to given value.

### HasOsZone

`func (o *ObjectStorageBucketRecord) HasOsZone() bool`

HasOsZone returns a boolean if a field has been set.

### GetOsZoneUuid

`func (o *ObjectStorageBucketRecord) GetOsZoneUuid() string`

GetOsZoneUuid returns the OsZoneUuid field if non-nil, zero value otherwise.

### GetOsZoneUuidOk

`func (o *ObjectStorageBucketRecord) GetOsZoneUuidOk() (*string, bool)`

GetOsZoneUuidOk returns a tuple with the OsZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsZoneUuid

`func (o *ObjectStorageBucketRecord) SetOsZoneUuid(v string)`

SetOsZoneUuid sets OsZoneUuid field to given value.

### HasOsZoneUuid

`func (o *ObjectStorageBucketRecord) HasOsZoneUuid() bool`

HasOsZoneUuid returns a boolean if a field has been set.

### GetOwner

`func (o *ObjectStorageBucketRecord) GetOwner() ObjectStorageUserNestview`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ObjectStorageBucketRecord) GetOwnerOk() (*ObjectStorageUserNestview, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ObjectStorageBucketRecord) SetOwner(v ObjectStorageUserNestview)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ObjectStorageBucketRecord) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnerPermission

`func (o *ObjectStorageBucketRecord) GetOwnerPermission() string`

GetOwnerPermission returns the OwnerPermission field if non-nil, zero value otherwise.

### GetOwnerPermissionOk

`func (o *ObjectStorageBucketRecord) GetOwnerPermissionOk() (*string, bool)`

GetOwnerPermissionOk returns a tuple with the OwnerPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPermission

`func (o *ObjectStorageBucketRecord) SetOwnerPermission(v string)`

SetOwnerPermission sets OwnerPermission field to given value.

### HasOwnerPermission

`func (o *ObjectStorageBucketRecord) HasOwnerPermission() bool`

HasOwnerPermission returns a boolean if a field has been set.

### GetPolicy

`func (o *ObjectStorageBucketRecord) GetPolicy() ObjectStoragePolicyNestview`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ObjectStorageBucketRecord) GetPolicyOk() (*ObjectStoragePolicyNestview, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ObjectStorageBucketRecord) SetPolicy(v ObjectStoragePolicyNestview)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ObjectStorageBucketRecord) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetPolicyEnabled

`func (o *ObjectStorageBucketRecord) GetPolicyEnabled() bool`

GetPolicyEnabled returns the PolicyEnabled field if non-nil, zero value otherwise.

### GetPolicyEnabledOk

`func (o *ObjectStorageBucketRecord) GetPolicyEnabledOk() (*bool, bool)`

GetPolicyEnabledOk returns a tuple with the PolicyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEnabled

`func (o *ObjectStorageBucketRecord) SetPolicyEnabled(v bool)`

SetPolicyEnabled sets PolicyEnabled field to given value.

### HasPolicyEnabled

`func (o *ObjectStorageBucketRecord) HasPolicyEnabled() bool`

HasPolicyEnabled returns a boolean if a field has been set.

### GetQos

`func (o *ObjectStorageBucketRecord) GetQos() OSBucketQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *ObjectStorageBucketRecord) GetQosOk() (*OSBucketQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *ObjectStorageBucketRecord) SetQos(v OSBucketQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *ObjectStorageBucketRecord) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetQuotaMaxObjects

`func (o *ObjectStorageBucketRecord) GetQuotaMaxObjects() int64`

GetQuotaMaxObjects returns the QuotaMaxObjects field if non-nil, zero value otherwise.

### GetQuotaMaxObjectsOk

`func (o *ObjectStorageBucketRecord) GetQuotaMaxObjectsOk() (*int64, bool)`

GetQuotaMaxObjectsOk returns a tuple with the QuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMaxObjects

`func (o *ObjectStorageBucketRecord) SetQuotaMaxObjects(v int64)`

SetQuotaMaxObjects sets QuotaMaxObjects field to given value.

### HasQuotaMaxObjects

`func (o *ObjectStorageBucketRecord) HasQuotaMaxObjects() bool`

HasQuotaMaxObjects returns a boolean if a field has been set.

### GetQuotaMaxSize

`func (o *ObjectStorageBucketRecord) GetQuotaMaxSize() int64`

GetQuotaMaxSize returns the QuotaMaxSize field if non-nil, zero value otherwise.

### GetQuotaMaxSizeOk

`func (o *ObjectStorageBucketRecord) GetQuotaMaxSizeOk() (*int64, bool)`

GetQuotaMaxSizeOk returns a tuple with the QuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMaxSize

`func (o *ObjectStorageBucketRecord) SetQuotaMaxSize(v int64)`

SetQuotaMaxSize sets QuotaMaxSize field to given value.

### HasQuotaMaxSize

`func (o *ObjectStorageBucketRecord) HasQuotaMaxSize() bool`

HasQuotaMaxSize returns a boolean if a field has been set.

### GetRemoteCluster

`func (o *ObjectStorageBucketRecord) GetRemoteCluster() RemoteClusterNestview`

GetRemoteCluster returns the RemoteCluster field if non-nil, zero value otherwise.

### GetRemoteClusterOk

`func (o *ObjectStorageBucketRecord) GetRemoteClusterOk() (*RemoteClusterNestview, bool)`

GetRemoteClusterOk returns a tuple with the RemoteCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCluster

`func (o *ObjectStorageBucketRecord) SetRemoteCluster(v RemoteClusterNestview)`

SetRemoteCluster sets RemoteCluster field to given value.

### HasRemoteCluster

`func (o *ObjectStorageBucketRecord) HasRemoteCluster() bool`

HasRemoteCluster returns a boolean if a field has been set.

### GetReplicationUuid

`func (o *ObjectStorageBucketRecord) GetReplicationUuid() string`

GetReplicationUuid returns the ReplicationUuid field if non-nil, zero value otherwise.

### GetReplicationUuidOk

`func (o *ObjectStorageBucketRecord) GetReplicationUuidOk() (*string, bool)`

GetReplicationUuidOk returns a tuple with the ReplicationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationUuid

`func (o *ObjectStorageBucketRecord) SetReplicationUuid(v string)`

SetReplicationUuid sets ReplicationUuid field to given value.

### HasReplicationUuid

`func (o *ObjectStorageBucketRecord) HasReplicationUuid() bool`

HasReplicationUuid returns a boolean if a field has been set.

### GetRestoreDays

`func (o *ObjectStorageBucketRecord) GetRestoreDays() int64`

GetRestoreDays returns the RestoreDays field if non-nil, zero value otherwise.

### GetRestoreDaysOk

`func (o *ObjectStorageBucketRecord) GetRestoreDaysOk() (*int64, bool)`

GetRestoreDaysOk returns a tuple with the RestoreDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreDays

`func (o *ObjectStorageBucketRecord) SetRestoreDays(v int64)`

SetRestoreDays sets RestoreDays field to given value.

### HasRestoreDays

`func (o *ObjectStorageBucketRecord) HasRestoreDays() bool`

HasRestoreDays returns a boolean if a field has been set.

### GetShards

`func (o *ObjectStorageBucketRecord) GetShards() int64`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *ObjectStorageBucketRecord) GetShardsOk() (*int64, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *ObjectStorageBucketRecord) SetShards(v int64)`

SetShards sets Shards field to given value.

### HasShards

`func (o *ObjectStorageBucketRecord) HasShards() bool`

HasShards returns a boolean if a field has been set.

### GetSpecificationObjects

`func (o *ObjectStorageBucketRecord) GetSpecificationObjects() int64`

GetSpecificationObjects returns the SpecificationObjects field if non-nil, zero value otherwise.

### GetSpecificationObjectsOk

`func (o *ObjectStorageBucketRecord) GetSpecificationObjectsOk() (*int64, bool)`

GetSpecificationObjectsOk returns a tuple with the SpecificationObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificationObjects

`func (o *ObjectStorageBucketRecord) SetSpecificationObjects(v int64)`

SetSpecificationObjects sets SpecificationObjects field to given value.

### HasSpecificationObjects

`func (o *ObjectStorageBucketRecord) HasSpecificationObjects() bool`

HasSpecificationObjects returns a boolean if a field has been set.

### GetStatus

`func (o *ObjectStorageBucketRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStorageBucketRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStorageBucketRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectStorageBucketRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *ObjectStorageBucketRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ObjectStorageBucketRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ObjectStorageBucketRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ObjectStorageBucketRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVirtual

`func (o *ObjectStorageBucketRecord) GetVirtual() bool`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *ObjectStorageBucketRecord) GetVirtualOk() (*bool, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *ObjectStorageBucketRecord) SetVirtual(v bool)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *ObjectStorageBucketRecord) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.

### GetSamples

`func (o *ObjectStorageBucketRecord) GetSamples() []ObjectStorageBucketStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ObjectStorageBucketRecord) GetSamplesOk() (*[]ObjectStorageBucketStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ObjectStorageBucketRecord) SetSamples(v []ObjectStorageBucketStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *ObjectStorageBucketRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


