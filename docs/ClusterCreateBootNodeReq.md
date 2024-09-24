# ClusterCreateBootNodeReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayIps** | Pointer to **[]string** | gateway ips for s3 | [optional] 
**GatewayNetwork** | Pointer to **string** | gateway network | [optional] 
**HostId** | **int64** | boot node host id | 
**PrivateIp** | Pointer to **string** | cluster private ip for internal data access | [optional] 
**PrivateNetwork** | **string** | private network | 
**PublicIp** | Pointer to **string** | public ip for outside data access | [optional] 
**PublicNetwork** | **string** | public network | 
**Roles** | Pointer to **[]string** | host roles: admin,monitor,block_storage_gateway,file_storage_gateway,s3_gateway,nfs_gateway | [optional] 

## Methods

### NewClusterCreateBootNodeReq

`func NewClusterCreateBootNodeReq(hostId int64, privateNetwork string, publicNetwork string, ) *ClusterCreateBootNodeReq`

NewClusterCreateBootNodeReq instantiates a new ClusterCreateBootNodeReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateBootNodeReqWithDefaults

`func NewClusterCreateBootNodeReqWithDefaults() *ClusterCreateBootNodeReq`

NewClusterCreateBootNodeReqWithDefaults instantiates a new ClusterCreateBootNodeReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayIps

`func (o *ClusterCreateBootNodeReq) GetGatewayIps() []string`

GetGatewayIps returns the GatewayIps field if non-nil, zero value otherwise.

### GetGatewayIpsOk

`func (o *ClusterCreateBootNodeReq) GetGatewayIpsOk() (*[]string, bool)`

GetGatewayIpsOk returns a tuple with the GatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIps

`func (o *ClusterCreateBootNodeReq) SetGatewayIps(v []string)`

SetGatewayIps sets GatewayIps field to given value.

### HasGatewayIps

`func (o *ClusterCreateBootNodeReq) HasGatewayIps() bool`

HasGatewayIps returns a boolean if a field has been set.

### GetGatewayNetwork

`func (o *ClusterCreateBootNodeReq) GetGatewayNetwork() string`

GetGatewayNetwork returns the GatewayNetwork field if non-nil, zero value otherwise.

### GetGatewayNetworkOk

`func (o *ClusterCreateBootNodeReq) GetGatewayNetworkOk() (*string, bool)`

GetGatewayNetworkOk returns a tuple with the GatewayNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNetwork

`func (o *ClusterCreateBootNodeReq) SetGatewayNetwork(v string)`

SetGatewayNetwork sets GatewayNetwork field to given value.

### HasGatewayNetwork

`func (o *ClusterCreateBootNodeReq) HasGatewayNetwork() bool`

HasGatewayNetwork returns a boolean if a field has been set.

### GetHostId

`func (o *ClusterCreateBootNodeReq) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ClusterCreateBootNodeReq) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ClusterCreateBootNodeReq) SetHostId(v int64)`

SetHostId sets HostId field to given value.


### GetPrivateIp

`func (o *ClusterCreateBootNodeReq) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *ClusterCreateBootNodeReq) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *ClusterCreateBootNodeReq) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *ClusterCreateBootNodeReq) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPrivateNetwork

`func (o *ClusterCreateBootNodeReq) GetPrivateNetwork() string`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *ClusterCreateBootNodeReq) GetPrivateNetworkOk() (*string, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *ClusterCreateBootNodeReq) SetPrivateNetwork(v string)`

SetPrivateNetwork sets PrivateNetwork field to given value.


### GetPublicIp

`func (o *ClusterCreateBootNodeReq) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *ClusterCreateBootNodeReq) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *ClusterCreateBootNodeReq) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *ClusterCreateBootNodeReq) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetPublicNetwork

`func (o *ClusterCreateBootNodeReq) GetPublicNetwork() string`

GetPublicNetwork returns the PublicNetwork field if non-nil, zero value otherwise.

### GetPublicNetworkOk

`func (o *ClusterCreateBootNodeReq) GetPublicNetworkOk() (*string, bool)`

GetPublicNetworkOk returns a tuple with the PublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetwork

`func (o *ClusterCreateBootNodeReq) SetPublicNetwork(v string)`

SetPublicNetwork sets PublicNetwork field to given value.


### GetRoles

`func (o *ClusterCreateBootNodeReq) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ClusterCreateBootNodeReq) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ClusterCreateBootNodeReq) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ClusterCreateBootNodeReq) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


