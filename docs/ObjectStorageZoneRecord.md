# ObjectStorageZoneRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Endpoints** | Pointer to **[]string** |  | [optional] 
**EpochUuid** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Local** | Pointer to **bool** |  | [optional] 
**Master** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RealmName** | Pointer to **string** |  | [optional] 
**RemoteCluster** | Pointer to [**RemoteClusterNestview**](RemoteClusterNestview.md) |  | [optional] 
**S3LoadBalancerGroup** | Pointer to [**S3LoadBalancerGroupNestview**](S3LoadBalancerGroupNestview.md) |  | [optional] 
**SearchZoneUuid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SwitchStatus** | Pointer to **string** |  | [optional] 
**SystemAccessKey** | Pointer to **string** |  | [optional] 
**SystemSecretKey** | Pointer to **string** |  | [optional] 
**SystemUser** | Pointer to **string** |  | [optional] 
**TierType** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**ZonegroupName** | Pointer to **string** |  | [optional] 
**Samples** | Pointer to [**[]ObjectStorageZoneStat**](ObjectStorageZoneStat.md) |  | [optional] 

## Methods

### NewObjectStorageZoneRecord

`func NewObjectStorageZoneRecord() *ObjectStorageZoneRecord`

NewObjectStorageZoneRecord instantiates a new ObjectStorageZoneRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageZoneRecordWithDefaults

`func NewObjectStorageZoneRecordWithDefaults() *ObjectStorageZoneRecord`

NewObjectStorageZoneRecordWithDefaults instantiates a new ObjectStorageZoneRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ObjectStorageZoneRecord) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ObjectStorageZoneRecord) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ObjectStorageZoneRecord) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ObjectStorageZoneRecord) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCluster

`func (o *ObjectStorageZoneRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ObjectStorageZoneRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ObjectStorageZoneRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ObjectStorageZoneRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnected

`func (o *ObjectStorageZoneRecord) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ObjectStorageZoneRecord) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ObjectStorageZoneRecord) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ObjectStorageZoneRecord) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetCreate

`func (o *ObjectStorageZoneRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ObjectStorageZoneRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ObjectStorageZoneRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ObjectStorageZoneRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetEndpoints

`func (o *ObjectStorageZoneRecord) GetEndpoints() []string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ObjectStorageZoneRecord) GetEndpointsOk() (*[]string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ObjectStorageZoneRecord) SetEndpoints(v []string)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ObjectStorageZoneRecord) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetEpochUuid

`func (o *ObjectStorageZoneRecord) GetEpochUuid() string`

GetEpochUuid returns the EpochUuid field if non-nil, zero value otherwise.

### GetEpochUuidOk

`func (o *ObjectStorageZoneRecord) GetEpochUuidOk() (*string, bool)`

GetEpochUuidOk returns a tuple with the EpochUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochUuid

`func (o *ObjectStorageZoneRecord) SetEpochUuid(v string)`

SetEpochUuid sets EpochUuid field to given value.

### HasEpochUuid

`func (o *ObjectStorageZoneRecord) HasEpochUuid() bool`

HasEpochUuid returns a boolean if a field has been set.

### GetId

`func (o *ObjectStorageZoneRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStorageZoneRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStorageZoneRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStorageZoneRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocal

`func (o *ObjectStorageZoneRecord) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *ObjectStorageZoneRecord) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *ObjectStorageZoneRecord) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *ObjectStorageZoneRecord) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetMaster

`func (o *ObjectStorageZoneRecord) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *ObjectStorageZoneRecord) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *ObjectStorageZoneRecord) SetMaster(v bool)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *ObjectStorageZoneRecord) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetName

`func (o *ObjectStorageZoneRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStorageZoneRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStorageZoneRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectStorageZoneRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRealmName

`func (o *ObjectStorageZoneRecord) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *ObjectStorageZoneRecord) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *ObjectStorageZoneRecord) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.

### HasRealmName

`func (o *ObjectStorageZoneRecord) HasRealmName() bool`

HasRealmName returns a boolean if a field has been set.

### GetRemoteCluster

`func (o *ObjectStorageZoneRecord) GetRemoteCluster() RemoteClusterNestview`

GetRemoteCluster returns the RemoteCluster field if non-nil, zero value otherwise.

### GetRemoteClusterOk

`func (o *ObjectStorageZoneRecord) GetRemoteClusterOk() (*RemoteClusterNestview, bool)`

GetRemoteClusterOk returns a tuple with the RemoteCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCluster

`func (o *ObjectStorageZoneRecord) SetRemoteCluster(v RemoteClusterNestview)`

SetRemoteCluster sets RemoteCluster field to given value.

### HasRemoteCluster

`func (o *ObjectStorageZoneRecord) HasRemoteCluster() bool`

HasRemoteCluster returns a boolean if a field has been set.

### GetS3LoadBalancerGroup

`func (o *ObjectStorageZoneRecord) GetS3LoadBalancerGroup() S3LoadBalancerGroupNestview`

GetS3LoadBalancerGroup returns the S3LoadBalancerGroup field if non-nil, zero value otherwise.

### GetS3LoadBalancerGroupOk

`func (o *ObjectStorageZoneRecord) GetS3LoadBalancerGroupOk() (*S3LoadBalancerGroupNestview, bool)`

GetS3LoadBalancerGroupOk returns a tuple with the S3LoadBalancerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3LoadBalancerGroup

`func (o *ObjectStorageZoneRecord) SetS3LoadBalancerGroup(v S3LoadBalancerGroupNestview)`

SetS3LoadBalancerGroup sets S3LoadBalancerGroup field to given value.

### HasS3LoadBalancerGroup

`func (o *ObjectStorageZoneRecord) HasS3LoadBalancerGroup() bool`

HasS3LoadBalancerGroup returns a boolean if a field has been set.

### GetSearchZoneUuid

`func (o *ObjectStorageZoneRecord) GetSearchZoneUuid() string`

GetSearchZoneUuid returns the SearchZoneUuid field if non-nil, zero value otherwise.

### GetSearchZoneUuidOk

`func (o *ObjectStorageZoneRecord) GetSearchZoneUuidOk() (*string, bool)`

GetSearchZoneUuidOk returns a tuple with the SearchZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchZoneUuid

`func (o *ObjectStorageZoneRecord) SetSearchZoneUuid(v string)`

SetSearchZoneUuid sets SearchZoneUuid field to given value.

### HasSearchZoneUuid

`func (o *ObjectStorageZoneRecord) HasSearchZoneUuid() bool`

HasSearchZoneUuid returns a boolean if a field has been set.

### GetStatus

`func (o *ObjectStorageZoneRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStorageZoneRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStorageZoneRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectStorageZoneRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwitchStatus

`func (o *ObjectStorageZoneRecord) GetSwitchStatus() string`

GetSwitchStatus returns the SwitchStatus field if non-nil, zero value otherwise.

### GetSwitchStatusOk

`func (o *ObjectStorageZoneRecord) GetSwitchStatusOk() (*string, bool)`

GetSwitchStatusOk returns a tuple with the SwitchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchStatus

`func (o *ObjectStorageZoneRecord) SetSwitchStatus(v string)`

SetSwitchStatus sets SwitchStatus field to given value.

### HasSwitchStatus

`func (o *ObjectStorageZoneRecord) HasSwitchStatus() bool`

HasSwitchStatus returns a boolean if a field has been set.

### GetSystemAccessKey

`func (o *ObjectStorageZoneRecord) GetSystemAccessKey() string`

GetSystemAccessKey returns the SystemAccessKey field if non-nil, zero value otherwise.

### GetSystemAccessKeyOk

`func (o *ObjectStorageZoneRecord) GetSystemAccessKeyOk() (*string, bool)`

GetSystemAccessKeyOk returns a tuple with the SystemAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAccessKey

`func (o *ObjectStorageZoneRecord) SetSystemAccessKey(v string)`

SetSystemAccessKey sets SystemAccessKey field to given value.

### HasSystemAccessKey

`func (o *ObjectStorageZoneRecord) HasSystemAccessKey() bool`

HasSystemAccessKey returns a boolean if a field has been set.

### GetSystemSecretKey

`func (o *ObjectStorageZoneRecord) GetSystemSecretKey() string`

GetSystemSecretKey returns the SystemSecretKey field if non-nil, zero value otherwise.

### GetSystemSecretKeyOk

`func (o *ObjectStorageZoneRecord) GetSystemSecretKeyOk() (*string, bool)`

GetSystemSecretKeyOk returns a tuple with the SystemSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemSecretKey

`func (o *ObjectStorageZoneRecord) SetSystemSecretKey(v string)`

SetSystemSecretKey sets SystemSecretKey field to given value.

### HasSystemSecretKey

`func (o *ObjectStorageZoneRecord) HasSystemSecretKey() bool`

HasSystemSecretKey returns a boolean if a field has been set.

### GetSystemUser

`func (o *ObjectStorageZoneRecord) GetSystemUser() string`

GetSystemUser returns the SystemUser field if non-nil, zero value otherwise.

### GetSystemUserOk

`func (o *ObjectStorageZoneRecord) GetSystemUserOk() (*string, bool)`

GetSystemUserOk returns a tuple with the SystemUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUser

`func (o *ObjectStorageZoneRecord) SetSystemUser(v string)`

SetSystemUser sets SystemUser field to given value.

### HasSystemUser

`func (o *ObjectStorageZoneRecord) HasSystemUser() bool`

HasSystemUser returns a boolean if a field has been set.

### GetTierType

`func (o *ObjectStorageZoneRecord) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *ObjectStorageZoneRecord) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *ObjectStorageZoneRecord) SetTierType(v string)`

SetTierType sets TierType field to given value.

### HasTierType

`func (o *ObjectStorageZoneRecord) HasTierType() bool`

HasTierType returns a boolean if a field has been set.

### GetUpdate

`func (o *ObjectStorageZoneRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ObjectStorageZoneRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ObjectStorageZoneRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ObjectStorageZoneRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUuid

`func (o *ObjectStorageZoneRecord) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ObjectStorageZoneRecord) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ObjectStorageZoneRecord) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ObjectStorageZoneRecord) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetZonegroupName

`func (o *ObjectStorageZoneRecord) GetZonegroupName() string`

GetZonegroupName returns the ZonegroupName field if non-nil, zero value otherwise.

### GetZonegroupNameOk

`func (o *ObjectStorageZoneRecord) GetZonegroupNameOk() (*string, bool)`

GetZonegroupNameOk returns a tuple with the ZonegroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonegroupName

`func (o *ObjectStorageZoneRecord) SetZonegroupName(v string)`

SetZonegroupName sets ZonegroupName field to given value.

### HasZonegroupName

`func (o *ObjectStorageZoneRecord) HasZonegroupName() bool`

HasZonegroupName returns a boolean if a field has been set.

### GetSamples

`func (o *ObjectStorageZoneRecord) GetSamples() []ObjectStorageZoneStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ObjectStorageZoneRecord) GetSamplesOk() (*[]ObjectStorageZoneStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ObjectStorageZoneRecord) SetSamples(v []ObjectStorageZoneStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *ObjectStorageZoneRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


