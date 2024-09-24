# BootNodeReqBootNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminNetwork** | Pointer to **string** | admin network: 10.0.0.0/24 | [optional] 
**FsId** | Pointer to **string** | fs id | [optional] 
**GatewayNetwork** | Pointer to **string** | gateway networks, multiple networks splited by comma: 10.0.3.0/24,10.0.4.0/24 | [optional] 
**PrivateNetwork** | **string** | private network : 10.0.2.0/24 | 
**PublicNetwork** | **string** | public network: 10.0.1.0/24 | 
**RegionName** | Pointer to **string** | region name | [optional] 

## Methods

### NewBootNodeReqBootNode

`func NewBootNodeReqBootNode(privateNetwork string, publicNetwork string, ) *BootNodeReqBootNode`

NewBootNodeReqBootNode instantiates a new BootNodeReqBootNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootNodeReqBootNodeWithDefaults

`func NewBootNodeReqBootNodeWithDefaults() *BootNodeReqBootNode`

NewBootNodeReqBootNodeWithDefaults instantiates a new BootNodeReqBootNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminNetwork

`func (o *BootNodeReqBootNode) GetAdminNetwork() string`

GetAdminNetwork returns the AdminNetwork field if non-nil, zero value otherwise.

### GetAdminNetworkOk

`func (o *BootNodeReqBootNode) GetAdminNetworkOk() (*string, bool)`

GetAdminNetworkOk returns a tuple with the AdminNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminNetwork

`func (o *BootNodeReqBootNode) SetAdminNetwork(v string)`

SetAdminNetwork sets AdminNetwork field to given value.

### HasAdminNetwork

`func (o *BootNodeReqBootNode) HasAdminNetwork() bool`

HasAdminNetwork returns a boolean if a field has been set.

### GetFsId

`func (o *BootNodeReqBootNode) GetFsId() string`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *BootNodeReqBootNode) GetFsIdOk() (*string, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *BootNodeReqBootNode) SetFsId(v string)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *BootNodeReqBootNode) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetGatewayNetwork

`func (o *BootNodeReqBootNode) GetGatewayNetwork() string`

GetGatewayNetwork returns the GatewayNetwork field if non-nil, zero value otherwise.

### GetGatewayNetworkOk

`func (o *BootNodeReqBootNode) GetGatewayNetworkOk() (*string, bool)`

GetGatewayNetworkOk returns a tuple with the GatewayNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNetwork

`func (o *BootNodeReqBootNode) SetGatewayNetwork(v string)`

SetGatewayNetwork sets GatewayNetwork field to given value.

### HasGatewayNetwork

`func (o *BootNodeReqBootNode) HasGatewayNetwork() bool`

HasGatewayNetwork returns a boolean if a field has been set.

### GetPrivateNetwork

`func (o *BootNodeReqBootNode) GetPrivateNetwork() string`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *BootNodeReqBootNode) GetPrivateNetworkOk() (*string, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *BootNodeReqBootNode) SetPrivateNetwork(v string)`

SetPrivateNetwork sets PrivateNetwork field to given value.


### GetPublicNetwork

`func (o *BootNodeReqBootNode) GetPublicNetwork() string`

GetPublicNetwork returns the PublicNetwork field if non-nil, zero value otherwise.

### GetPublicNetworkOk

`func (o *BootNodeReqBootNode) GetPublicNetworkOk() (*string, bool)`

GetPublicNetworkOk returns a tuple with the PublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetwork

`func (o *BootNodeReqBootNode) SetPublicNetwork(v string)`

SetPublicNetwork sets PublicNetwork field to given value.


### GetRegionName

`func (o *BootNodeReqBootNode) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *BootNodeReqBootNode) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *BootNodeReqBootNode) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *BootNodeReqBootNode) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


