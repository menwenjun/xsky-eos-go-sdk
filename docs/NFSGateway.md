# NFSGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**BucketNum** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GatewayIp** | Pointer to **string** |  | [optional] 
**GatewayName** | Pointer to **string** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MountClients** | Pointer to **string** |  | [optional] 
**MountNum** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OspZoneId** | Pointer to **int64** |  | [optional] 
**PlatformIp** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNFSGateway

`func NewNFSGateway() *NFSGateway`

NewNFSGateway instantiates a new NFSGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFSGatewayWithDefaults

`func NewNFSGatewayWithDefaults() *NFSGateway`

NewNFSGatewayWithDefaults instantiates a new NFSGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *NFSGateway) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *NFSGateway) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *NFSGateway) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *NFSGateway) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetBucketNum

`func (o *NFSGateway) GetBucketNum() int64`

GetBucketNum returns the BucketNum field if non-nil, zero value otherwise.

### GetBucketNumOk

`func (o *NFSGateway) GetBucketNumOk() (*int64, bool)`

GetBucketNumOk returns a tuple with the BucketNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketNum

`func (o *NFSGateway) SetBucketNum(v int64)`

SetBucketNum sets BucketNum field to given value.

### HasBucketNum

`func (o *NFSGateway) HasBucketNum() bool`

HasBucketNum returns a boolean if a field has been set.

### GetCluster

`func (o *NFSGateway) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *NFSGateway) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *NFSGateway) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *NFSGateway) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *NFSGateway) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *NFSGateway) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *NFSGateway) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *NFSGateway) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *NFSGateway) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NFSGateway) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NFSGateway) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NFSGateway) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGatewayIp

`func (o *NFSGateway) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *NFSGateway) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *NFSGateway) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *NFSGateway) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetGatewayName

`func (o *NFSGateway) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *NFSGateway) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *NFSGateway) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *NFSGateway) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### GetHost

`func (o *NFSGateway) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NFSGateway) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NFSGateway) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *NFSGateway) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *NFSGateway) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NFSGateway) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NFSGateway) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NFSGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMountClients

`func (o *NFSGateway) GetMountClients() string`

GetMountClients returns the MountClients field if non-nil, zero value otherwise.

### GetMountClientsOk

`func (o *NFSGateway) GetMountClientsOk() (*string, bool)`

GetMountClientsOk returns a tuple with the MountClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountClients

`func (o *NFSGateway) SetMountClients(v string)`

SetMountClients sets MountClients field to given value.

### HasMountClients

`func (o *NFSGateway) HasMountClients() bool`

HasMountClients returns a boolean if a field has been set.

### GetMountNum

`func (o *NFSGateway) GetMountNum() int64`

GetMountNum returns the MountNum field if non-nil, zero value otherwise.

### GetMountNumOk

`func (o *NFSGateway) GetMountNumOk() (*int64, bool)`

GetMountNumOk returns a tuple with the MountNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountNum

`func (o *NFSGateway) SetMountNum(v int64)`

SetMountNum sets MountNum field to given value.

### HasMountNum

`func (o *NFSGateway) HasMountNum() bool`

HasMountNum returns a boolean if a field has been set.

### GetName

`func (o *NFSGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NFSGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NFSGateway) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NFSGateway) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOspZoneId

`func (o *NFSGateway) GetOspZoneId() int64`

GetOspZoneId returns the OspZoneId field if non-nil, zero value otherwise.

### GetOspZoneIdOk

`func (o *NFSGateway) GetOspZoneIdOk() (*int64, bool)`

GetOspZoneIdOk returns a tuple with the OspZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspZoneId

`func (o *NFSGateway) SetOspZoneId(v int64)`

SetOspZoneId sets OspZoneId field to given value.

### HasOspZoneId

`func (o *NFSGateway) HasOspZoneId() bool`

HasOspZoneId returns a boolean if a field has been set.

### GetPlatformIp

`func (o *NFSGateway) GetPlatformIp() string`

GetPlatformIp returns the PlatformIp field if non-nil, zero value otherwise.

### GetPlatformIpOk

`func (o *NFSGateway) GetPlatformIpOk() (*string, bool)`

GetPlatformIpOk returns a tuple with the PlatformIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIp

`func (o *NFSGateway) SetPlatformIp(v string)`

SetPlatformIp sets PlatformIp field to given value.

### HasPlatformIp

`func (o *NFSGateway) HasPlatformIp() bool`

HasPlatformIp returns a boolean if a field has been set.

### GetPort

`func (o *NFSGateway) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NFSGateway) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NFSGateway) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *NFSGateway) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *NFSGateway) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NFSGateway) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NFSGateway) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NFSGateway) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *NFSGateway) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *NFSGateway) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *NFSGateway) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *NFSGateway) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


