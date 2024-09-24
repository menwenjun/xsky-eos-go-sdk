# ClusterCreateNetworkReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayNetwork** | Pointer to **string** | gateway network | [optional] 
**OspClusterNetwork** | Pointer to **string** | osp cluster network | [optional] 
**OspGatewayIsConnected** | Pointer to **string** | \&quot;true\&quot; represents each network can connect to others. v.v. | [optional] 
**OspGatewayNetwork** | Pointer to **string** | osp gateway network | [optional] 
**PrivateNetwork** | **string** | private network | 
**PublicNetwork** | **string** | public network | 

## Methods

### NewClusterCreateNetworkReq

`func NewClusterCreateNetworkReq(privateNetwork string, publicNetwork string, ) *ClusterCreateNetworkReq`

NewClusterCreateNetworkReq instantiates a new ClusterCreateNetworkReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateNetworkReqWithDefaults

`func NewClusterCreateNetworkReqWithDefaults() *ClusterCreateNetworkReq`

NewClusterCreateNetworkReqWithDefaults instantiates a new ClusterCreateNetworkReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayNetwork

`func (o *ClusterCreateNetworkReq) GetGatewayNetwork() string`

GetGatewayNetwork returns the GatewayNetwork field if non-nil, zero value otherwise.

### GetGatewayNetworkOk

`func (o *ClusterCreateNetworkReq) GetGatewayNetworkOk() (*string, bool)`

GetGatewayNetworkOk returns a tuple with the GatewayNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNetwork

`func (o *ClusterCreateNetworkReq) SetGatewayNetwork(v string)`

SetGatewayNetwork sets GatewayNetwork field to given value.

### HasGatewayNetwork

`func (o *ClusterCreateNetworkReq) HasGatewayNetwork() bool`

HasGatewayNetwork returns a boolean if a field has been set.

### GetOspClusterNetwork

`func (o *ClusterCreateNetworkReq) GetOspClusterNetwork() string`

GetOspClusterNetwork returns the OspClusterNetwork field if non-nil, zero value otherwise.

### GetOspClusterNetworkOk

`func (o *ClusterCreateNetworkReq) GetOspClusterNetworkOk() (*string, bool)`

GetOspClusterNetworkOk returns a tuple with the OspClusterNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspClusterNetwork

`func (o *ClusterCreateNetworkReq) SetOspClusterNetwork(v string)`

SetOspClusterNetwork sets OspClusterNetwork field to given value.

### HasOspClusterNetwork

`func (o *ClusterCreateNetworkReq) HasOspClusterNetwork() bool`

HasOspClusterNetwork returns a boolean if a field has been set.

### GetOspGatewayIsConnected

`func (o *ClusterCreateNetworkReq) GetOspGatewayIsConnected() string`

GetOspGatewayIsConnected returns the OspGatewayIsConnected field if non-nil, zero value otherwise.

### GetOspGatewayIsConnectedOk

`func (o *ClusterCreateNetworkReq) GetOspGatewayIsConnectedOk() (*string, bool)`

GetOspGatewayIsConnectedOk returns a tuple with the OspGatewayIsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspGatewayIsConnected

`func (o *ClusterCreateNetworkReq) SetOspGatewayIsConnected(v string)`

SetOspGatewayIsConnected sets OspGatewayIsConnected field to given value.

### HasOspGatewayIsConnected

`func (o *ClusterCreateNetworkReq) HasOspGatewayIsConnected() bool`

HasOspGatewayIsConnected returns a boolean if a field has been set.

### GetOspGatewayNetwork

`func (o *ClusterCreateNetworkReq) GetOspGatewayNetwork() string`

GetOspGatewayNetwork returns the OspGatewayNetwork field if non-nil, zero value otherwise.

### GetOspGatewayNetworkOk

`func (o *ClusterCreateNetworkReq) GetOspGatewayNetworkOk() (*string, bool)`

GetOspGatewayNetworkOk returns a tuple with the OspGatewayNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspGatewayNetwork

`func (o *ClusterCreateNetworkReq) SetOspGatewayNetwork(v string)`

SetOspGatewayNetwork sets OspGatewayNetwork field to given value.

### HasOspGatewayNetwork

`func (o *ClusterCreateNetworkReq) HasOspGatewayNetwork() bool`

HasOspGatewayNetwork returns a boolean if a field has been set.

### GetPrivateNetwork

`func (o *ClusterCreateNetworkReq) GetPrivateNetwork() string`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *ClusterCreateNetworkReq) GetPrivateNetworkOk() (*string, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *ClusterCreateNetworkReq) SetPrivateNetwork(v string)`

SetPrivateNetwork sets PrivateNetwork field to given value.


### GetPublicNetwork

`func (o *ClusterCreateNetworkReq) GetPublicNetwork() string`

GetPublicNetwork returns the PublicNetwork field if non-nil, zero value otherwise.

### GetPublicNetworkOk

`func (o *ClusterCreateNetworkReq) GetPublicNetworkOk() (*string, bool)`

GetPublicNetworkOk returns a tuple with the PublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetwork

`func (o *ClusterCreateNetworkReq) SetPublicNetwork(v string)`

SetPublicNetwork sets PublicNetwork field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


