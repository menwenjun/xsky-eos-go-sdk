# ObjectStorageGateway

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

## Methods

### NewObjectStorageGateway

`func NewObjectStorageGateway() *ObjectStorageGateway`

NewObjectStorageGateway instantiates a new ObjectStorageGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageGatewayWithDefaults

`func NewObjectStorageGatewayWithDefaults() *ObjectStorageGateway`

NewObjectStorageGatewayWithDefaults instantiates a new ObjectStorageGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *ObjectStorageGateway) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ObjectStorageGateway) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ObjectStorageGateway) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ObjectStorageGateway) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *ObjectStorageGateway) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ObjectStorageGateway) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ObjectStorageGateway) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ObjectStorageGateway) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *ObjectStorageGateway) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectStorageGateway) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectStorageGateway) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectStorageGateway) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEtag

`func (o *ObjectStorageGateway) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *ObjectStorageGateway) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *ObjectStorageGateway) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *ObjectStorageGateway) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetGatewayIp

`func (o *ObjectStorageGateway) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *ObjectStorageGateway) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *ObjectStorageGateway) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *ObjectStorageGateway) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetGatewayName

`func (o *ObjectStorageGateway) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *ObjectStorageGateway) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *ObjectStorageGateway) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *ObjectStorageGateway) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### GetHost

`func (o *ObjectStorageGateway) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ObjectStorageGateway) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ObjectStorageGateway) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *ObjectStorageGateway) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *ObjectStorageGateway) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStorageGateway) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStorageGateway) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStorageGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ObjectStorageGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStorageGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStorageGateway) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectStorageGateway) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsZone

`func (o *ObjectStorageGateway) GetOsZone() ObjectStorageZoneNestview`

GetOsZone returns the OsZone field if non-nil, zero value otherwise.

### GetOsZoneOk

`func (o *ObjectStorageGateway) GetOsZoneOk() (*ObjectStorageZoneNestview, bool)`

GetOsZoneOk returns a tuple with the OsZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsZone

`func (o *ObjectStorageGateway) SetOsZone(v ObjectStorageZoneNestview)`

SetOsZone sets OsZone field to given value.

### HasOsZone

`func (o *ObjectStorageGateway) HasOsZone() bool`

HasOsZone returns a boolean if a field has been set.

### GetPort

`func (o *ObjectStorageGateway) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ObjectStorageGateway) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ObjectStorageGateway) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ObjectStorageGateway) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRole

`func (o *ObjectStorageGateway) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ObjectStorageGateway) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ObjectStorageGateway) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ObjectStorageGateway) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *ObjectStorageGateway) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStorageGateway) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStorageGateway) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectStorageGateway) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *ObjectStorageGateway) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ObjectStorageGateway) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ObjectStorageGateway) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ObjectStorageGateway) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


