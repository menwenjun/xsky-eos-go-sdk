# DfsVIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsGatewayGroup** | Pointer to [**DfsGatewayGroupNestview**](DfsGatewayGroupNestview.md) |  | [optional] 
**DfsGatewayZone** | Pointer to [**DfsGatewayZoneNestview**](DfsGatewayZoneNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**PrimaryGateway** | Pointer to [**DfsGatewayNestview**](DfsGatewayNestview.md) |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Vip** | Pointer to **string** |  | [optional] 
**VipMask** | Pointer to **int64** |  | [optional] 

## Methods

### NewDfsVIP

`func NewDfsVIP() *DfsVIP`

NewDfsVIP instantiates a new DfsVIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsVIPWithDefaults

`func NewDfsVIPWithDefaults() *DfsVIP`

NewDfsVIPWithDefaults instantiates a new DfsVIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsVIP) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsVIP) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsVIP) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsVIP) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *DfsVIP) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsVIP) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsVIP) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsVIP) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsVIP) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsVIP) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsVIP) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsVIP) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsGatewayGroup

`func (o *DfsVIP) GetDfsGatewayGroup() DfsGatewayGroupNestview`

GetDfsGatewayGroup returns the DfsGatewayGroup field if non-nil, zero value otherwise.

### GetDfsGatewayGroupOk

`func (o *DfsVIP) GetDfsGatewayGroupOk() (*DfsGatewayGroupNestview, bool)`

GetDfsGatewayGroupOk returns a tuple with the DfsGatewayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroup

`func (o *DfsVIP) SetDfsGatewayGroup(v DfsGatewayGroupNestview)`

SetDfsGatewayGroup sets DfsGatewayGroup field to given value.

### HasDfsGatewayGroup

`func (o *DfsVIP) HasDfsGatewayGroup() bool`

HasDfsGatewayGroup returns a boolean if a field has been set.

### GetDfsGatewayZone

`func (o *DfsVIP) GetDfsGatewayZone() DfsGatewayZoneNestview`

GetDfsGatewayZone returns the DfsGatewayZone field if non-nil, zero value otherwise.

### GetDfsGatewayZoneOk

`func (o *DfsVIP) GetDfsGatewayZoneOk() (*DfsGatewayZoneNestview, bool)`

GetDfsGatewayZoneOk returns a tuple with the DfsGatewayZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayZone

`func (o *DfsVIP) SetDfsGatewayZone(v DfsGatewayZoneNestview)`

SetDfsGatewayZone sets DfsGatewayZone field to given value.

### HasDfsGatewayZone

`func (o *DfsVIP) HasDfsGatewayZone() bool`

HasDfsGatewayZone returns a boolean if a field has been set.

### GetId

`func (o *DfsVIP) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsVIP) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsVIP) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsVIP) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrimaryGateway

`func (o *DfsVIP) GetPrimaryGateway() DfsGatewayNestview`

GetPrimaryGateway returns the PrimaryGateway field if non-nil, zero value otherwise.

### GetPrimaryGatewayOk

`func (o *DfsVIP) GetPrimaryGatewayOk() (*DfsGatewayNestview, bool)`

GetPrimaryGatewayOk returns a tuple with the PrimaryGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGateway

`func (o *DfsVIP) SetPrimaryGateway(v DfsGatewayNestview)`

SetPrimaryGateway sets PrimaryGateway field to given value.

### HasPrimaryGateway

`func (o *DfsVIP) HasPrimaryGateway() bool`

HasPrimaryGateway returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsVIP) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsVIP) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsVIP) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsVIP) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVip

`func (o *DfsVIP) GetVip() string`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *DfsVIP) GetVipOk() (*string, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *DfsVIP) SetVip(v string)`

SetVip sets Vip field to given value.

### HasVip

`func (o *DfsVIP) HasVip() bool`

HasVip returns a boolean if a field has been set.

### GetVipMask

`func (o *DfsVIP) GetVipMask() int64`

GetVipMask returns the VipMask field if non-nil, zero value otherwise.

### GetVipMaskOk

`func (o *DfsVIP) GetVipMaskOk() (*int64, bool)`

GetVipMaskOk returns a tuple with the VipMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipMask

`func (o *DfsVIP) SetVipMask(v int64)`

SetVipMask sets VipMask field to given value.

### HasVipMask

`func (o *DfsVIP) HasVipMask() bool`

HasVipMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


