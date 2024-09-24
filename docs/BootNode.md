# BootNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminNetwork** | **string** |  | 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**GatewayNetwork** | **string** |  | 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**PrivateNetwork** | **string** |  | 
**PublicNetwork** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewBootNode

`func NewBootNode(adminNetwork string, gatewayNetwork string, privateNetwork string, publicNetwork string, ) *BootNode`

NewBootNode instantiates a new BootNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootNodeWithDefaults

`func NewBootNodeWithDefaults() *BootNode`

NewBootNodeWithDefaults instantiates a new BootNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminNetwork

`func (o *BootNode) GetAdminNetwork() string`

GetAdminNetwork returns the AdminNetwork field if non-nil, zero value otherwise.

### GetAdminNetworkOk

`func (o *BootNode) GetAdminNetworkOk() (*string, bool)`

GetAdminNetworkOk returns a tuple with the AdminNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminNetwork

`func (o *BootNode) SetAdminNetwork(v string)`

SetAdminNetwork sets AdminNetwork field to given value.


### GetCluster

`func (o *BootNode) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *BootNode) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *BootNode) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *BootNode) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetGatewayNetwork

`func (o *BootNode) GetGatewayNetwork() string`

GetGatewayNetwork returns the GatewayNetwork field if non-nil, zero value otherwise.

### GetGatewayNetworkOk

`func (o *BootNode) GetGatewayNetworkOk() (*string, bool)`

GetGatewayNetworkOk returns a tuple with the GatewayNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNetwork

`func (o *BootNode) SetGatewayNetwork(v string)`

SetGatewayNetwork sets GatewayNetwork field to given value.


### GetHost

`func (o *BootNode) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *BootNode) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *BootNode) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *BootNode) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *BootNode) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BootNode) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BootNode) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BootNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrivateNetwork

`func (o *BootNode) GetPrivateNetwork() string`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *BootNode) GetPrivateNetworkOk() (*string, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *BootNode) SetPrivateNetwork(v string)`

SetPrivateNetwork sets PrivateNetwork field to given value.


### GetPublicNetwork

`func (o *BootNode) GetPublicNetwork() string`

GetPublicNetwork returns the PublicNetwork field if non-nil, zero value otherwise.

### GetPublicNetworkOk

`func (o *BootNode) GetPublicNetworkOk() (*string, bool)`

GetPublicNetworkOk returns a tuple with the PublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetwork

`func (o *BootNode) SetPublicNetwork(v string)`

SetPublicNetwork sets PublicNetwork field to given value.


### GetStatus

`func (o *BootNode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BootNode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BootNode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BootNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


