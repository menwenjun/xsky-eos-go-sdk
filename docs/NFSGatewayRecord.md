# NFSGatewayRecord

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
**Samples** | Pointer to [**[]NFSGatewayStat**](NFSGatewayStat.md) |  | [optional] 

## Methods

### NewNFSGatewayRecord

`func NewNFSGatewayRecord() *NFSGatewayRecord`

NewNFSGatewayRecord instantiates a new NFSGatewayRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFSGatewayRecordWithDefaults

`func NewNFSGatewayRecordWithDefaults() *NFSGatewayRecord`

NewNFSGatewayRecordWithDefaults instantiates a new NFSGatewayRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *NFSGatewayRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *NFSGatewayRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *NFSGatewayRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *NFSGatewayRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetBucketNum

`func (o *NFSGatewayRecord) GetBucketNum() int64`

GetBucketNum returns the BucketNum field if non-nil, zero value otherwise.

### GetBucketNumOk

`func (o *NFSGatewayRecord) GetBucketNumOk() (*int64, bool)`

GetBucketNumOk returns a tuple with the BucketNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketNum

`func (o *NFSGatewayRecord) SetBucketNum(v int64)`

SetBucketNum sets BucketNum field to given value.

### HasBucketNum

`func (o *NFSGatewayRecord) HasBucketNum() bool`

HasBucketNum returns a boolean if a field has been set.

### GetCluster

`func (o *NFSGatewayRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *NFSGatewayRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *NFSGatewayRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *NFSGatewayRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *NFSGatewayRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *NFSGatewayRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *NFSGatewayRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *NFSGatewayRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *NFSGatewayRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NFSGatewayRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NFSGatewayRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NFSGatewayRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGatewayIp

`func (o *NFSGatewayRecord) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *NFSGatewayRecord) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *NFSGatewayRecord) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *NFSGatewayRecord) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetGatewayName

`func (o *NFSGatewayRecord) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *NFSGatewayRecord) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *NFSGatewayRecord) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *NFSGatewayRecord) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### GetHost

`func (o *NFSGatewayRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NFSGatewayRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NFSGatewayRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *NFSGatewayRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *NFSGatewayRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NFSGatewayRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NFSGatewayRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NFSGatewayRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMountClients

`func (o *NFSGatewayRecord) GetMountClients() string`

GetMountClients returns the MountClients field if non-nil, zero value otherwise.

### GetMountClientsOk

`func (o *NFSGatewayRecord) GetMountClientsOk() (*string, bool)`

GetMountClientsOk returns a tuple with the MountClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountClients

`func (o *NFSGatewayRecord) SetMountClients(v string)`

SetMountClients sets MountClients field to given value.

### HasMountClients

`func (o *NFSGatewayRecord) HasMountClients() bool`

HasMountClients returns a boolean if a field has been set.

### GetMountNum

`func (o *NFSGatewayRecord) GetMountNum() int64`

GetMountNum returns the MountNum field if non-nil, zero value otherwise.

### GetMountNumOk

`func (o *NFSGatewayRecord) GetMountNumOk() (*int64, bool)`

GetMountNumOk returns a tuple with the MountNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountNum

`func (o *NFSGatewayRecord) SetMountNum(v int64)`

SetMountNum sets MountNum field to given value.

### HasMountNum

`func (o *NFSGatewayRecord) HasMountNum() bool`

HasMountNum returns a boolean if a field has been set.

### GetName

`func (o *NFSGatewayRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NFSGatewayRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NFSGatewayRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NFSGatewayRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOspZoneId

`func (o *NFSGatewayRecord) GetOspZoneId() int64`

GetOspZoneId returns the OspZoneId field if non-nil, zero value otherwise.

### GetOspZoneIdOk

`func (o *NFSGatewayRecord) GetOspZoneIdOk() (*int64, bool)`

GetOspZoneIdOk returns a tuple with the OspZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspZoneId

`func (o *NFSGatewayRecord) SetOspZoneId(v int64)`

SetOspZoneId sets OspZoneId field to given value.

### HasOspZoneId

`func (o *NFSGatewayRecord) HasOspZoneId() bool`

HasOspZoneId returns a boolean if a field has been set.

### GetPlatformIp

`func (o *NFSGatewayRecord) GetPlatformIp() string`

GetPlatformIp returns the PlatformIp field if non-nil, zero value otherwise.

### GetPlatformIpOk

`func (o *NFSGatewayRecord) GetPlatformIpOk() (*string, bool)`

GetPlatformIpOk returns a tuple with the PlatformIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIp

`func (o *NFSGatewayRecord) SetPlatformIp(v string)`

SetPlatformIp sets PlatformIp field to given value.

### HasPlatformIp

`func (o *NFSGatewayRecord) HasPlatformIp() bool`

HasPlatformIp returns a boolean if a field has been set.

### GetPort

`func (o *NFSGatewayRecord) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NFSGatewayRecord) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NFSGatewayRecord) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *NFSGatewayRecord) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *NFSGatewayRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NFSGatewayRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NFSGatewayRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NFSGatewayRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *NFSGatewayRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *NFSGatewayRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *NFSGatewayRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *NFSGatewayRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *NFSGatewayRecord) GetSamples() []NFSGatewayStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *NFSGatewayRecord) GetSamplesOk() (*[]NFSGatewayStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *NFSGatewayRecord) SetSamples(v []NFSGatewayStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *NFSGatewayRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


