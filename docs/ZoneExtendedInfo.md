# ZoneExtendedInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datacenters** | Pointer to [**[]DataCenter**](DataCenter.md) |  | [optional] 
**Fsid** | Pointer to **string** |  | [optional] 
**Gateways** | Pointer to [**[]Gateway**](Gateway.md) |  | [optional] 
**HostNodes** | Pointer to [**[]HostNode**](HostNode.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MetadataActionStatus** | Pointer to **string** |  | [optional] 
**MetadataName** | Pointer to **string** |  | [optional] 
**MetadataStatus** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Routings** | Pointer to [**[]Routing**](Routing.md) |  | [optional] 

## Methods

### NewZoneExtendedInfo

`func NewZoneExtendedInfo() *ZoneExtendedInfo`

NewZoneExtendedInfo instantiates a new ZoneExtendedInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneExtendedInfoWithDefaults

`func NewZoneExtendedInfoWithDefaults() *ZoneExtendedInfo`

NewZoneExtendedInfoWithDefaults instantiates a new ZoneExtendedInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenters

`func (o *ZoneExtendedInfo) GetDatacenters() []DataCenter`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *ZoneExtendedInfo) GetDatacentersOk() (*[]DataCenter, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *ZoneExtendedInfo) SetDatacenters(v []DataCenter)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *ZoneExtendedInfo) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetFsid

`func (o *ZoneExtendedInfo) GetFsid() string`

GetFsid returns the Fsid field if non-nil, zero value otherwise.

### GetFsidOk

`func (o *ZoneExtendedInfo) GetFsidOk() (*string, bool)`

GetFsidOk returns a tuple with the Fsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsid

`func (o *ZoneExtendedInfo) SetFsid(v string)`

SetFsid sets Fsid field to given value.

### HasFsid

`func (o *ZoneExtendedInfo) HasFsid() bool`

HasFsid returns a boolean if a field has been set.

### GetGateways

`func (o *ZoneExtendedInfo) GetGateways() []Gateway`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *ZoneExtendedInfo) GetGatewaysOk() (*[]Gateway, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *ZoneExtendedInfo) SetGateways(v []Gateway)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *ZoneExtendedInfo) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetHostNodes

`func (o *ZoneExtendedInfo) GetHostNodes() []HostNode`

GetHostNodes returns the HostNodes field if non-nil, zero value otherwise.

### GetHostNodesOk

`func (o *ZoneExtendedInfo) GetHostNodesOk() (*[]HostNode, bool)`

GetHostNodesOk returns a tuple with the HostNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNodes

`func (o *ZoneExtendedInfo) SetHostNodes(v []HostNode)`

SetHostNodes sets HostNodes field to given value.

### HasHostNodes

`func (o *ZoneExtendedInfo) HasHostNodes() bool`

HasHostNodes returns a boolean if a field has been set.

### GetId

`func (o *ZoneExtendedInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneExtendedInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneExtendedInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ZoneExtendedInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadataActionStatus

`func (o *ZoneExtendedInfo) GetMetadataActionStatus() string`

GetMetadataActionStatus returns the MetadataActionStatus field if non-nil, zero value otherwise.

### GetMetadataActionStatusOk

`func (o *ZoneExtendedInfo) GetMetadataActionStatusOk() (*string, bool)`

GetMetadataActionStatusOk returns a tuple with the MetadataActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataActionStatus

`func (o *ZoneExtendedInfo) SetMetadataActionStatus(v string)`

SetMetadataActionStatus sets MetadataActionStatus field to given value.

### HasMetadataActionStatus

`func (o *ZoneExtendedInfo) HasMetadataActionStatus() bool`

HasMetadataActionStatus returns a boolean if a field has been set.

### GetMetadataName

`func (o *ZoneExtendedInfo) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *ZoneExtendedInfo) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *ZoneExtendedInfo) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.

### HasMetadataName

`func (o *ZoneExtendedInfo) HasMetadataName() bool`

HasMetadataName returns a boolean if a field has been set.

### GetMetadataStatus

`func (o *ZoneExtendedInfo) GetMetadataStatus() string`

GetMetadataStatus returns the MetadataStatus field if non-nil, zero value otherwise.

### GetMetadataStatusOk

`func (o *ZoneExtendedInfo) GetMetadataStatusOk() (*string, bool)`

GetMetadataStatusOk returns a tuple with the MetadataStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataStatus

`func (o *ZoneExtendedInfo) SetMetadataStatus(v string)`

SetMetadataStatus sets MetadataStatus field to given value.

### HasMetadataStatus

`func (o *ZoneExtendedInfo) HasMetadataStatus() bool`

HasMetadataStatus returns a boolean if a field has been set.

### GetName

`func (o *ZoneExtendedInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneExtendedInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneExtendedInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneExtendedInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoutings

`func (o *ZoneExtendedInfo) GetRoutings() []Routing`

GetRoutings returns the Routings field if non-nil, zero value otherwise.

### GetRoutingsOk

`func (o *ZoneExtendedInfo) GetRoutingsOk() (*[]Routing, bool)`

GetRoutingsOk returns a tuple with the Routings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutings

`func (o *ZoneExtendedInfo) SetRoutings(v []Routing)`

SetRoutings sets Routings field to given value.

### HasRoutings

`func (o *ZoneExtendedInfo) HasRoutings() bool`

HasRoutings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


