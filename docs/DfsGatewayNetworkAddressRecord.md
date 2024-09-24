# DfsGatewayNetworkAddressRecord

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
**Samples** | Pointer to [**[]DfsGatewayStat**](DfsGatewayStat.md) | TODO(liubo): next version maybe change to DfsGatewayNetworkAddressStat | [optional] 

## Methods

### NewDfsGatewayNetworkAddressRecord

`func NewDfsGatewayNetworkAddressRecord() *DfsGatewayNetworkAddressRecord`

NewDfsGatewayNetworkAddressRecord instantiates a new DfsGatewayNetworkAddressRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayNetworkAddressRecordWithDefaults

`func NewDfsGatewayNetworkAddressRecordWithDefaults() *DfsGatewayNetworkAddressRecord`

NewDfsGatewayNetworkAddressRecordWithDefaults instantiates a new DfsGatewayNetworkAddressRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DfsGatewayNetworkAddressRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsGatewayNetworkAddressRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsGatewayNetworkAddressRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsGatewayNetworkAddressRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsGatewayNetworkAddressRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsGatewayNetworkAddressRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsGatewayNetworkAddressRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsGatewayNetworkAddressRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsGateway

`func (o *DfsGatewayNetworkAddressRecord) GetDfsGateway() DfsGateway`

GetDfsGateway returns the DfsGateway field if non-nil, zero value otherwise.

### GetDfsGatewayOk

`func (o *DfsGatewayNetworkAddressRecord) GetDfsGatewayOk() (*DfsGateway, bool)`

GetDfsGatewayOk returns a tuple with the DfsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGateway

`func (o *DfsGatewayNetworkAddressRecord) SetDfsGateway(v DfsGateway)`

SetDfsGateway sets DfsGateway field to given value.

### HasDfsGateway

`func (o *DfsGatewayNetworkAddressRecord) HasDfsGateway() bool`

HasDfsGateway returns a boolean if a field has been set.

### GetDfsGatewayZone

`func (o *DfsGatewayNetworkAddressRecord) GetDfsGatewayZone() DfsGatewayZoneNestview`

GetDfsGatewayZone returns the DfsGatewayZone field if non-nil, zero value otherwise.

### GetDfsGatewayZoneOk

`func (o *DfsGatewayNetworkAddressRecord) GetDfsGatewayZoneOk() (*DfsGatewayZoneNestview, bool)`

GetDfsGatewayZoneOk returns a tuple with the DfsGatewayZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayZone

`func (o *DfsGatewayNetworkAddressRecord) SetDfsGatewayZone(v DfsGatewayZoneNestview)`

SetDfsGatewayZone sets DfsGatewayZone field to given value.

### HasDfsGatewayZone

`func (o *DfsGatewayNetworkAddressRecord) HasDfsGatewayZone() bool`

HasDfsGatewayZone returns a boolean if a field has been set.

### GetHost

`func (o *DfsGatewayNetworkAddressRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DfsGatewayNetworkAddressRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DfsGatewayNetworkAddressRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *DfsGatewayNetworkAddressRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *DfsGatewayNetworkAddressRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsGatewayNetworkAddressRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsGatewayNetworkAddressRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsGatewayNetworkAddressRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkAddress

`func (o *DfsGatewayNetworkAddressRecord) GetNetworkAddress() NetworkAddressNestview`

GetNetworkAddress returns the NetworkAddress field if non-nil, zero value otherwise.

### GetNetworkAddressOk

`func (o *DfsGatewayNetworkAddressRecord) GetNetworkAddressOk() (*NetworkAddressNestview, bool)`

GetNetworkAddressOk returns a tuple with the NetworkAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddress

`func (o *DfsGatewayNetworkAddressRecord) SetNetworkAddress(v NetworkAddressNestview)`

SetNetworkAddress sets NetworkAddress field to given value.

### HasNetworkAddress

`func (o *DfsGatewayNetworkAddressRecord) HasNetworkAddress() bool`

HasNetworkAddress returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsGatewayNetworkAddressRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsGatewayNetworkAddressRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsGatewayNetworkAddressRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsGatewayNetworkAddressRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *DfsGatewayNetworkAddressRecord) GetSamples() []DfsGatewayStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *DfsGatewayNetworkAddressRecord) GetSamplesOk() (*[]DfsGatewayStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *DfsGatewayNetworkAddressRecord) SetSamples(v []DfsGatewayStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *DfsGatewayNetworkAddressRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


