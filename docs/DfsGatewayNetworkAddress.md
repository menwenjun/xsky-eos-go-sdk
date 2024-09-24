# DfsGatewayNetworkAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsGateway** | Pointer to [**DfsGateway**](DfsGateway.md) |  | [optional] 
**DfsGatewayZone** | Pointer to [**DfsGatewayZoneNestview**](DfsGatewayZoneNestview.md) |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**NetworkAddress** | Pointer to [**NetworkAddressNestview**](NetworkAddressNestview.md) |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsGatewayNetworkAddress

`func NewDfsGatewayNetworkAddress() *DfsGatewayNetworkAddress`

NewDfsGatewayNetworkAddress instantiates a new DfsGatewayNetworkAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayNetworkAddressWithDefaults

`func NewDfsGatewayNetworkAddressWithDefaults() *DfsGatewayNetworkAddress`

NewDfsGatewayNetworkAddressWithDefaults instantiates a new DfsGatewayNetworkAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DfsGatewayNetworkAddress) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsGatewayNetworkAddress) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsGatewayNetworkAddress) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsGatewayNetworkAddress) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsGatewayNetworkAddress) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsGatewayNetworkAddress) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsGatewayNetworkAddress) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsGatewayNetworkAddress) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsGateway

`func (o *DfsGatewayNetworkAddress) GetDfsGateway() DfsGateway`

GetDfsGateway returns the DfsGateway field if non-nil, zero value otherwise.

### GetDfsGatewayOk

`func (o *DfsGatewayNetworkAddress) GetDfsGatewayOk() (*DfsGateway, bool)`

GetDfsGatewayOk returns a tuple with the DfsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGateway

`func (o *DfsGatewayNetworkAddress) SetDfsGateway(v DfsGateway)`

SetDfsGateway sets DfsGateway field to given value.

### HasDfsGateway

`func (o *DfsGatewayNetworkAddress) HasDfsGateway() bool`

HasDfsGateway returns a boolean if a field has been set.

### GetDfsGatewayZone

`func (o *DfsGatewayNetworkAddress) GetDfsGatewayZone() DfsGatewayZoneNestview`

GetDfsGatewayZone returns the DfsGatewayZone field if non-nil, zero value otherwise.

### GetDfsGatewayZoneOk

`func (o *DfsGatewayNetworkAddress) GetDfsGatewayZoneOk() (*DfsGatewayZoneNestview, bool)`

GetDfsGatewayZoneOk returns a tuple with the DfsGatewayZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayZone

`func (o *DfsGatewayNetworkAddress) SetDfsGatewayZone(v DfsGatewayZoneNestview)`

SetDfsGatewayZone sets DfsGatewayZone field to given value.

### HasDfsGatewayZone

`func (o *DfsGatewayNetworkAddress) HasDfsGatewayZone() bool`

HasDfsGatewayZone returns a boolean if a field has been set.

### GetHost

`func (o *DfsGatewayNetworkAddress) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DfsGatewayNetworkAddress) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DfsGatewayNetworkAddress) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *DfsGatewayNetworkAddress) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *DfsGatewayNetworkAddress) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsGatewayNetworkAddress) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsGatewayNetworkAddress) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsGatewayNetworkAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkAddress

`func (o *DfsGatewayNetworkAddress) GetNetworkAddress() NetworkAddressNestview`

GetNetworkAddress returns the NetworkAddress field if non-nil, zero value otherwise.

### GetNetworkAddressOk

`func (o *DfsGatewayNetworkAddress) GetNetworkAddressOk() (*NetworkAddressNestview, bool)`

GetNetworkAddressOk returns a tuple with the NetworkAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddress

`func (o *DfsGatewayNetworkAddress) SetNetworkAddress(v NetworkAddressNestview)`

SetNetworkAddress sets NetworkAddress field to given value.

### HasNetworkAddress

`func (o *DfsGatewayNetworkAddress) HasNetworkAddress() bool`

HasNetworkAddress returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsGatewayNetworkAddress) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsGatewayNetworkAddress) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsGatewayNetworkAddress) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsGatewayNetworkAddress) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


