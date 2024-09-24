# ObjectStorageZone

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

## Methods

### NewObjectStorageZone

`func NewObjectStorageZone() *ObjectStorageZone`

NewObjectStorageZone instantiates a new ObjectStorageZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageZoneWithDefaults

`func NewObjectStorageZoneWithDefaults() *ObjectStorageZone`

NewObjectStorageZoneWithDefaults instantiates a new ObjectStorageZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ObjectStorageZone) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ObjectStorageZone) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ObjectStorageZone) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ObjectStorageZone) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCluster

`func (o *ObjectStorageZone) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ObjectStorageZone) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ObjectStorageZone) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ObjectStorageZone) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnected

`func (o *ObjectStorageZone) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ObjectStorageZone) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ObjectStorageZone) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ObjectStorageZone) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetCreate

`func (o *ObjectStorageZone) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ObjectStorageZone) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ObjectStorageZone) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ObjectStorageZone) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetEndpoints

`func (o *ObjectStorageZone) GetEndpoints() []string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ObjectStorageZone) GetEndpointsOk() (*[]string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ObjectStorageZone) SetEndpoints(v []string)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ObjectStorageZone) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetEpochUuid

`func (o *ObjectStorageZone) GetEpochUuid() string`

GetEpochUuid returns the EpochUuid field if non-nil, zero value otherwise.

### GetEpochUuidOk

`func (o *ObjectStorageZone) GetEpochUuidOk() (*string, bool)`

GetEpochUuidOk returns a tuple with the EpochUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochUuid

`func (o *ObjectStorageZone) SetEpochUuid(v string)`

SetEpochUuid sets EpochUuid field to given value.

### HasEpochUuid

`func (o *ObjectStorageZone) HasEpochUuid() bool`

HasEpochUuid returns a boolean if a field has been set.

### GetId

`func (o *ObjectStorageZone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStorageZone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStorageZone) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStorageZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocal

`func (o *ObjectStorageZone) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *ObjectStorageZone) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *ObjectStorageZone) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *ObjectStorageZone) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetMaster

`func (o *ObjectStorageZone) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *ObjectStorageZone) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *ObjectStorageZone) SetMaster(v bool)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *ObjectStorageZone) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetName

`func (o *ObjectStorageZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStorageZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStorageZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectStorageZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRealmName

`func (o *ObjectStorageZone) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *ObjectStorageZone) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *ObjectStorageZone) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.

### HasRealmName

`func (o *ObjectStorageZone) HasRealmName() bool`

HasRealmName returns a boolean if a field has been set.

### GetRemoteCluster

`func (o *ObjectStorageZone) GetRemoteCluster() RemoteClusterNestview`

GetRemoteCluster returns the RemoteCluster field if non-nil, zero value otherwise.

### GetRemoteClusterOk

`func (o *ObjectStorageZone) GetRemoteClusterOk() (*RemoteClusterNestview, bool)`

GetRemoteClusterOk returns a tuple with the RemoteCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCluster

`func (o *ObjectStorageZone) SetRemoteCluster(v RemoteClusterNestview)`

SetRemoteCluster sets RemoteCluster field to given value.

### HasRemoteCluster

`func (o *ObjectStorageZone) HasRemoteCluster() bool`

HasRemoteCluster returns a boolean if a field has been set.

### GetS3LoadBalancerGroup

`func (o *ObjectStorageZone) GetS3LoadBalancerGroup() S3LoadBalancerGroupNestview`

GetS3LoadBalancerGroup returns the S3LoadBalancerGroup field if non-nil, zero value otherwise.

### GetS3LoadBalancerGroupOk

`func (o *ObjectStorageZone) GetS3LoadBalancerGroupOk() (*S3LoadBalancerGroupNestview, bool)`

GetS3LoadBalancerGroupOk returns a tuple with the S3LoadBalancerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3LoadBalancerGroup

`func (o *ObjectStorageZone) SetS3LoadBalancerGroup(v S3LoadBalancerGroupNestview)`

SetS3LoadBalancerGroup sets S3LoadBalancerGroup field to given value.

### HasS3LoadBalancerGroup

`func (o *ObjectStorageZone) HasS3LoadBalancerGroup() bool`

HasS3LoadBalancerGroup returns a boolean if a field has been set.

### GetSearchZoneUuid

`func (o *ObjectStorageZone) GetSearchZoneUuid() string`

GetSearchZoneUuid returns the SearchZoneUuid field if non-nil, zero value otherwise.

### GetSearchZoneUuidOk

`func (o *ObjectStorageZone) GetSearchZoneUuidOk() (*string, bool)`

GetSearchZoneUuidOk returns a tuple with the SearchZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchZoneUuid

`func (o *ObjectStorageZone) SetSearchZoneUuid(v string)`

SetSearchZoneUuid sets SearchZoneUuid field to given value.

### HasSearchZoneUuid

`func (o *ObjectStorageZone) HasSearchZoneUuid() bool`

HasSearchZoneUuid returns a boolean if a field has been set.

### GetStatus

`func (o *ObjectStorageZone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStorageZone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStorageZone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectStorageZone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwitchStatus

`func (o *ObjectStorageZone) GetSwitchStatus() string`

GetSwitchStatus returns the SwitchStatus field if non-nil, zero value otherwise.

### GetSwitchStatusOk

`func (o *ObjectStorageZone) GetSwitchStatusOk() (*string, bool)`

GetSwitchStatusOk returns a tuple with the SwitchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchStatus

`func (o *ObjectStorageZone) SetSwitchStatus(v string)`

SetSwitchStatus sets SwitchStatus field to given value.

### HasSwitchStatus

`func (o *ObjectStorageZone) HasSwitchStatus() bool`

HasSwitchStatus returns a boolean if a field has been set.

### GetSystemAccessKey

`func (o *ObjectStorageZone) GetSystemAccessKey() string`

GetSystemAccessKey returns the SystemAccessKey field if non-nil, zero value otherwise.

### GetSystemAccessKeyOk

`func (o *ObjectStorageZone) GetSystemAccessKeyOk() (*string, bool)`

GetSystemAccessKeyOk returns a tuple with the SystemAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAccessKey

`func (o *ObjectStorageZone) SetSystemAccessKey(v string)`

SetSystemAccessKey sets SystemAccessKey field to given value.

### HasSystemAccessKey

`func (o *ObjectStorageZone) HasSystemAccessKey() bool`

HasSystemAccessKey returns a boolean if a field has been set.

### GetSystemSecretKey

`func (o *ObjectStorageZone) GetSystemSecretKey() string`

GetSystemSecretKey returns the SystemSecretKey field if non-nil, zero value otherwise.

### GetSystemSecretKeyOk

`func (o *ObjectStorageZone) GetSystemSecretKeyOk() (*string, bool)`

GetSystemSecretKeyOk returns a tuple with the SystemSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemSecretKey

`func (o *ObjectStorageZone) SetSystemSecretKey(v string)`

SetSystemSecretKey sets SystemSecretKey field to given value.

### HasSystemSecretKey

`func (o *ObjectStorageZone) HasSystemSecretKey() bool`

HasSystemSecretKey returns a boolean if a field has been set.

### GetSystemUser

`func (o *ObjectStorageZone) GetSystemUser() string`

GetSystemUser returns the SystemUser field if non-nil, zero value otherwise.

### GetSystemUserOk

`func (o *ObjectStorageZone) GetSystemUserOk() (*string, bool)`

GetSystemUserOk returns a tuple with the SystemUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUser

`func (o *ObjectStorageZone) SetSystemUser(v string)`

SetSystemUser sets SystemUser field to given value.

### HasSystemUser

`func (o *ObjectStorageZone) HasSystemUser() bool`

HasSystemUser returns a boolean if a field has been set.

### GetTierType

`func (o *ObjectStorageZone) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *ObjectStorageZone) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *ObjectStorageZone) SetTierType(v string)`

SetTierType sets TierType field to given value.

### HasTierType

`func (o *ObjectStorageZone) HasTierType() bool`

HasTierType returns a boolean if a field has been set.

### GetUpdate

`func (o *ObjectStorageZone) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ObjectStorageZone) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ObjectStorageZone) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ObjectStorageZone) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUuid

`func (o *ObjectStorageZone) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ObjectStorageZone) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ObjectStorageZone) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ObjectStorageZone) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetZonegroupName

`func (o *ObjectStorageZone) GetZonegroupName() string`

GetZonegroupName returns the ZonegroupName field if non-nil, zero value otherwise.

### GetZonegroupNameOk

`func (o *ObjectStorageZone) GetZonegroupNameOk() (*string, bool)`

GetZonegroupNameOk returns a tuple with the ZonegroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonegroupName

`func (o *ObjectStorageZone) SetZonegroupName(v string)`

SetZonegroupName sets ZonegroupName field to given value.

### HasZonegroupName

`func (o *ObjectStorageZone) HasZonegroupName() bool`

HasZonegroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


