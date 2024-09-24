# ObjectStorageGatewayRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Etag** | Pointer to **string** |  | [optional] 
**GatewayIp** | Pointer to **string** |  | [optional] 
**GatewayName** | Pointer to **string** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OsZone** | Pointer to [**ObjectStorageZoneNestview**](ObjectStorageZoneNestview.md) |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Samples** | Pointer to [**[]ObjectStorageGatewayStat**](ObjectStorageGatewayStat.md) |  | [optional] 

## Methods

### NewObjectStorageGatewayRecord

`func NewObjectStorageGatewayRecord() *ObjectStorageGatewayRecord`

NewObjectStorageGatewayRecord instantiates a new ObjectStorageGatewayRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageGatewayRecordWithDefaults

`func NewObjectStorageGatewayRecordWithDefaults() *ObjectStorageGatewayRecord`

NewObjectStorageGatewayRecordWithDefaults instantiates a new ObjectStorageGatewayRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *ObjectStorageGatewayRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ObjectStorageGatewayRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ObjectStorageGatewayRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ObjectStorageGatewayRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *ObjectStorageGatewayRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ObjectStorageGatewayRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ObjectStorageGatewayRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ObjectStorageGatewayRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *ObjectStorageGatewayRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectStorageGatewayRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectStorageGatewayRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectStorageGatewayRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEtag

`func (o *ObjectStorageGatewayRecord) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *ObjectStorageGatewayRecord) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *ObjectStorageGatewayRecord) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *ObjectStorageGatewayRecord) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetGatewayIp

`func (o *ObjectStorageGatewayRecord) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *ObjectStorageGatewayRecord) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *ObjectStorageGatewayRecord) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *ObjectStorageGatewayRecord) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetGatewayName

`func (o *ObjectStorageGatewayRecord) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *ObjectStorageGatewayRecord) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *ObjectStorageGatewayRecord) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *ObjectStorageGatewayRecord) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### GetHost

`func (o *ObjectStorageGatewayRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ObjectStorageGatewayRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ObjectStorageGatewayRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *ObjectStorageGatewayRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *ObjectStorageGatewayRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStorageGatewayRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStorageGatewayRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStorageGatewayRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ObjectStorageGatewayRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStorageGatewayRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStorageGatewayRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectStorageGatewayRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsZone

`func (o *ObjectStorageGatewayRecord) GetOsZone() ObjectStorageZoneNestview`

GetOsZone returns the OsZone field if non-nil, zero value otherwise.

### GetOsZoneOk

`func (o *ObjectStorageGatewayRecord) GetOsZoneOk() (*ObjectStorageZoneNestview, bool)`

GetOsZoneOk returns a tuple with the OsZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsZone

`func (o *ObjectStorageGatewayRecord) SetOsZone(v ObjectStorageZoneNestview)`

SetOsZone sets OsZone field to given value.

### HasOsZone

`func (o *ObjectStorageGatewayRecord) HasOsZone() bool`

HasOsZone returns a boolean if a field has been set.

### GetPort

`func (o *ObjectStorageGatewayRecord) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ObjectStorageGatewayRecord) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ObjectStorageGatewayRecord) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ObjectStorageGatewayRecord) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRole

`func (o *ObjectStorageGatewayRecord) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ObjectStorageGatewayRecord) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ObjectStorageGatewayRecord) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ObjectStorageGatewayRecord) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *ObjectStorageGatewayRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStorageGatewayRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStorageGatewayRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectStorageGatewayRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *ObjectStorageGatewayRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ObjectStorageGatewayRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ObjectStorageGatewayRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ObjectStorageGatewayRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *ObjectStorageGatewayRecord) GetSamples() []ObjectStorageGatewayStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ObjectStorageGatewayRecord) GetSamplesOk() (*[]ObjectStorageGatewayStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ObjectStorageGatewayRecord) SetSamples(v []ObjectStorageGatewayStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *ObjectStorageGatewayRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


