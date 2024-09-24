# OspHostReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | **int64** | host id | 
**OspClusterIp** | **string** | osp cluster ip | 
**OspGatewayIps** | **[]string** | osp gateway ip | 
**Roles** | Pointer to **[]string** | host roles | [optional] 

## Methods

### NewOspHostReq

`func NewOspHostReq(hostId int64, ospClusterIp string, ospGatewayIps []string, ) *OspHostReq`

NewOspHostReq instantiates a new OspHostReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspHostReqWithDefaults

`func NewOspHostReqWithDefaults() *OspHostReq`

NewOspHostReqWithDefaults instantiates a new OspHostReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *OspHostReq) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *OspHostReq) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *OspHostReq) SetHostId(v int64)`

SetHostId sets HostId field to given value.


### GetOspClusterIp

`func (o *OspHostReq) GetOspClusterIp() string`

GetOspClusterIp returns the OspClusterIp field if non-nil, zero value otherwise.

### GetOspClusterIpOk

`func (o *OspHostReq) GetOspClusterIpOk() (*string, bool)`

GetOspClusterIpOk returns a tuple with the OspClusterIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspClusterIp

`func (o *OspHostReq) SetOspClusterIp(v string)`

SetOspClusterIp sets OspClusterIp field to given value.


### GetOspGatewayIps

`func (o *OspHostReq) GetOspGatewayIps() []string`

GetOspGatewayIps returns the OspGatewayIps field if non-nil, zero value otherwise.

### GetOspGatewayIpsOk

`func (o *OspHostReq) GetOspGatewayIpsOk() (*[]string, bool)`

GetOspGatewayIpsOk returns a tuple with the OspGatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspGatewayIps

`func (o *OspHostReq) SetOspGatewayIps(v []string)`

SetOspGatewayIps sets OspGatewayIps field to given value.


### GetRoles

`func (o *OspHostReq) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OspHostReq) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OspHostReq) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OspHostReq) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


