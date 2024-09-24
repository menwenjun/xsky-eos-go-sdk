# OSReplicationPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**OsZoneUuids** | Pointer to **[]string** |  | [optional] 
**OsZones** | Pointer to [**[]ObjectStorageZoneNestview**](ObjectStorageZoneNestview.md) |  | [optional] 
**ReplicationUuid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Suspended** | Pointer to **bool** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewOSReplicationPath

`func NewOSReplicationPath() *OSReplicationPath`

NewOSReplicationPath instantiates a new OSReplicationPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSReplicationPathWithDefaults

`func NewOSReplicationPathWithDefaults() *OSReplicationPath`

NewOSReplicationPathWithDefaults instantiates a new OSReplicationPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *OSReplicationPath) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSReplicationPath) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSReplicationPath) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSReplicationPath) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSReplicationPath) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSReplicationPath) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSReplicationPath) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSReplicationPath) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetOsZoneUuids

`func (o *OSReplicationPath) GetOsZoneUuids() []string`

GetOsZoneUuids returns the OsZoneUuids field if non-nil, zero value otherwise.

### GetOsZoneUuidsOk

`func (o *OSReplicationPath) GetOsZoneUuidsOk() (*[]string, bool)`

GetOsZoneUuidsOk returns a tuple with the OsZoneUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsZoneUuids

`func (o *OSReplicationPath) SetOsZoneUuids(v []string)`

SetOsZoneUuids sets OsZoneUuids field to given value.

### HasOsZoneUuids

`func (o *OSReplicationPath) HasOsZoneUuids() bool`

HasOsZoneUuids returns a boolean if a field has been set.

### GetOsZones

`func (o *OSReplicationPath) GetOsZones() []ObjectStorageZoneNestview`

GetOsZones returns the OsZones field if non-nil, zero value otherwise.

### GetOsZonesOk

`func (o *OSReplicationPath) GetOsZonesOk() (*[]ObjectStorageZoneNestview, bool)`

GetOsZonesOk returns a tuple with the OsZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsZones

`func (o *OSReplicationPath) SetOsZones(v []ObjectStorageZoneNestview)`

SetOsZones sets OsZones field to given value.

### HasOsZones

`func (o *OSReplicationPath) HasOsZones() bool`

HasOsZones returns a boolean if a field has been set.

### GetReplicationUuid

`func (o *OSReplicationPath) GetReplicationUuid() string`

GetReplicationUuid returns the ReplicationUuid field if non-nil, zero value otherwise.

### GetReplicationUuidOk

`func (o *OSReplicationPath) GetReplicationUuidOk() (*string, bool)`

GetReplicationUuidOk returns a tuple with the ReplicationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationUuid

`func (o *OSReplicationPath) SetReplicationUuid(v string)`

SetReplicationUuid sets ReplicationUuid field to given value.

### HasReplicationUuid

`func (o *OSReplicationPath) HasReplicationUuid() bool`

HasReplicationUuid returns a boolean if a field has been set.

### GetStatus

`func (o *OSReplicationPath) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSReplicationPath) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSReplicationPath) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSReplicationPath) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuspended

`func (o *OSReplicationPath) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *OSReplicationPath) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *OSReplicationPath) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *OSReplicationPath) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetUpdate

`func (o *OSReplicationPath) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OSReplicationPath) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OSReplicationPath) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OSReplicationPath) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUuid

`func (o *OSReplicationPath) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OSReplicationPath) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OSReplicationPath) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *OSReplicationPath) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


