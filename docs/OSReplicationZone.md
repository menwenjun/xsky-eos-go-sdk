# OSReplicationZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Dirty** | Pointer to **bool** |  | [optional] 
**OsBucket** | Pointer to [**ObjectStorageBucketNestview**](ObjectStorageBucketNestview.md) |  | [optional] 
**OsBucketOwnerCluster** | Pointer to [**RemoteClusterNestview**](RemoteClusterNestview.md) |  | [optional] 
**OsBucketOwnerZone** | Pointer to [**ObjectStorageZoneNestview**](ObjectStorageZoneNestview.md) |  | [optional] 
**OsRemotePolicy** | Pointer to [**OSRemotePolicyNestview**](OSRemotePolicyNestview.md) |  | [optional] 
**OsRemotePolicyUuid** | Pointer to **string** |  | [optional] 
**OsReplicationPathNum** | Pointer to **int64** |  | [optional] 
**OsZone** | Pointer to [**ObjectStorageZoneNestview**](ObjectStorageZoneNestview.md) |  | [optional] 
**OsZoneUuid** | Pointer to **string** |  | [optional] 
**Readonly** | Pointer to **bool** |  | [optional] 
**ReplicationUuid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewOSReplicationZone

`func NewOSReplicationZone() *OSReplicationZone`

NewOSReplicationZone instantiates a new OSReplicationZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSReplicationZoneWithDefaults

`func NewOSReplicationZoneWithDefaults() *OSReplicationZone`

NewOSReplicationZoneWithDefaults instantiates a new OSReplicationZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *OSReplicationZone) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSReplicationZone) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSReplicationZone) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSReplicationZone) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSReplicationZone) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSReplicationZone) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSReplicationZone) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSReplicationZone) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDirty

`func (o *OSReplicationZone) GetDirty() bool`

GetDirty returns the Dirty field if non-nil, zero value otherwise.

### GetDirtyOk

`func (o *OSReplicationZone) GetDirtyOk() (*bool, bool)`

GetDirtyOk returns a tuple with the Dirty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirty

`func (o *OSReplicationZone) SetDirty(v bool)`

SetDirty sets Dirty field to given value.

### HasDirty

`func (o *OSReplicationZone) HasDirty() bool`

HasDirty returns a boolean if a field has been set.

### GetOsBucket

`func (o *OSReplicationZone) GetOsBucket() ObjectStorageBucketNestview`

GetOsBucket returns the OsBucket field if non-nil, zero value otherwise.

### GetOsBucketOk

`func (o *OSReplicationZone) GetOsBucketOk() (*ObjectStorageBucketNestview, bool)`

GetOsBucketOk returns a tuple with the OsBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBucket

`func (o *OSReplicationZone) SetOsBucket(v ObjectStorageBucketNestview)`

SetOsBucket sets OsBucket field to given value.

### HasOsBucket

`func (o *OSReplicationZone) HasOsBucket() bool`

HasOsBucket returns a boolean if a field has been set.

### GetOsBucketOwnerCluster

`func (o *OSReplicationZone) GetOsBucketOwnerCluster() RemoteClusterNestview`

GetOsBucketOwnerCluster returns the OsBucketOwnerCluster field if non-nil, zero value otherwise.

### GetOsBucketOwnerClusterOk

`func (o *OSReplicationZone) GetOsBucketOwnerClusterOk() (*RemoteClusterNestview, bool)`

GetOsBucketOwnerClusterOk returns a tuple with the OsBucketOwnerCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBucketOwnerCluster

`func (o *OSReplicationZone) SetOsBucketOwnerCluster(v RemoteClusterNestview)`

SetOsBucketOwnerCluster sets OsBucketOwnerCluster field to given value.

### HasOsBucketOwnerCluster

`func (o *OSReplicationZone) HasOsBucketOwnerCluster() bool`

HasOsBucketOwnerCluster returns a boolean if a field has been set.

### GetOsBucketOwnerZone

`func (o *OSReplicationZone) GetOsBucketOwnerZone() ObjectStorageZoneNestview`

GetOsBucketOwnerZone returns the OsBucketOwnerZone field if non-nil, zero value otherwise.

### GetOsBucketOwnerZoneOk

`func (o *OSReplicationZone) GetOsBucketOwnerZoneOk() (*ObjectStorageZoneNestview, bool)`

GetOsBucketOwnerZoneOk returns a tuple with the OsBucketOwnerZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBucketOwnerZone

`func (o *OSReplicationZone) SetOsBucketOwnerZone(v ObjectStorageZoneNestview)`

SetOsBucketOwnerZone sets OsBucketOwnerZone field to given value.

### HasOsBucketOwnerZone

`func (o *OSReplicationZone) HasOsBucketOwnerZone() bool`

HasOsBucketOwnerZone returns a boolean if a field has been set.

### GetOsRemotePolicy

`func (o *OSReplicationZone) GetOsRemotePolicy() OSRemotePolicyNestview`

GetOsRemotePolicy returns the OsRemotePolicy field if non-nil, zero value otherwise.

### GetOsRemotePolicyOk

`func (o *OSReplicationZone) GetOsRemotePolicyOk() (*OSRemotePolicyNestview, bool)`

GetOsRemotePolicyOk returns a tuple with the OsRemotePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRemotePolicy

`func (o *OSReplicationZone) SetOsRemotePolicy(v OSRemotePolicyNestview)`

SetOsRemotePolicy sets OsRemotePolicy field to given value.

### HasOsRemotePolicy

`func (o *OSReplicationZone) HasOsRemotePolicy() bool`

HasOsRemotePolicy returns a boolean if a field has been set.

### GetOsRemotePolicyUuid

`func (o *OSReplicationZone) GetOsRemotePolicyUuid() string`

GetOsRemotePolicyUuid returns the OsRemotePolicyUuid field if non-nil, zero value otherwise.

### GetOsRemotePolicyUuidOk

`func (o *OSReplicationZone) GetOsRemotePolicyUuidOk() (*string, bool)`

GetOsRemotePolicyUuidOk returns a tuple with the OsRemotePolicyUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRemotePolicyUuid

`func (o *OSReplicationZone) SetOsRemotePolicyUuid(v string)`

SetOsRemotePolicyUuid sets OsRemotePolicyUuid field to given value.

### HasOsRemotePolicyUuid

`func (o *OSReplicationZone) HasOsRemotePolicyUuid() bool`

HasOsRemotePolicyUuid returns a boolean if a field has been set.

### GetOsReplicationPathNum

`func (o *OSReplicationZone) GetOsReplicationPathNum() int64`

GetOsReplicationPathNum returns the OsReplicationPathNum field if non-nil, zero value otherwise.

### GetOsReplicationPathNumOk

`func (o *OSReplicationZone) GetOsReplicationPathNumOk() (*int64, bool)`

GetOsReplicationPathNumOk returns a tuple with the OsReplicationPathNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsReplicationPathNum

`func (o *OSReplicationZone) SetOsReplicationPathNum(v int64)`

SetOsReplicationPathNum sets OsReplicationPathNum field to given value.

### HasOsReplicationPathNum

`func (o *OSReplicationZone) HasOsReplicationPathNum() bool`

HasOsReplicationPathNum returns a boolean if a field has been set.

### GetOsZone

`func (o *OSReplicationZone) GetOsZone() ObjectStorageZoneNestview`

GetOsZone returns the OsZone field if non-nil, zero value otherwise.

### GetOsZoneOk

`func (o *OSReplicationZone) GetOsZoneOk() (*ObjectStorageZoneNestview, bool)`

GetOsZoneOk returns a tuple with the OsZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsZone

`func (o *OSReplicationZone) SetOsZone(v ObjectStorageZoneNestview)`

SetOsZone sets OsZone field to given value.

### HasOsZone

`func (o *OSReplicationZone) HasOsZone() bool`

HasOsZone returns a boolean if a field has been set.

### GetOsZoneUuid

`func (o *OSReplicationZone) GetOsZoneUuid() string`

GetOsZoneUuid returns the OsZoneUuid field if non-nil, zero value otherwise.

### GetOsZoneUuidOk

`func (o *OSReplicationZone) GetOsZoneUuidOk() (*string, bool)`

GetOsZoneUuidOk returns a tuple with the OsZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsZoneUuid

`func (o *OSReplicationZone) SetOsZoneUuid(v string)`

SetOsZoneUuid sets OsZoneUuid field to given value.

### HasOsZoneUuid

`func (o *OSReplicationZone) HasOsZoneUuid() bool`

HasOsZoneUuid returns a boolean if a field has been set.

### GetReadonly

`func (o *OSReplicationZone) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *OSReplicationZone) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *OSReplicationZone) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *OSReplicationZone) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetReplicationUuid

`func (o *OSReplicationZone) GetReplicationUuid() string`

GetReplicationUuid returns the ReplicationUuid field if non-nil, zero value otherwise.

### GetReplicationUuidOk

`func (o *OSReplicationZone) GetReplicationUuidOk() (*string, bool)`

GetReplicationUuidOk returns a tuple with the ReplicationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationUuid

`func (o *OSReplicationZone) SetReplicationUuid(v string)`

SetReplicationUuid sets ReplicationUuid field to given value.

### HasReplicationUuid

`func (o *OSReplicationZone) HasReplicationUuid() bool`

HasReplicationUuid returns a boolean if a field has been set.

### GetStatus

`func (o *OSReplicationZone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSReplicationZone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSReplicationZone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSReplicationZone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *OSReplicationZone) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OSReplicationZone) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OSReplicationZone) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OSReplicationZone) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUuid

`func (o *OSReplicationZone) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OSReplicationZone) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OSReplicationZone) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *OSReplicationZone) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


