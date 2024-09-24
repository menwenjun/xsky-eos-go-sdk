# CephHostReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayIps** | Pointer to **[]string** | gateway ips for s3 | [optional] 
**HostId** | **int64** | boot node host id | 
**PrivateIp** | Pointer to **string** | cluster private ip for internal data access | [optional] 
**PublicIp** | Pointer to **string** | public ip for outside data access | [optional] 
**Roles** | Pointer to **[]string** | host roles | [optional] 
**Type** | Pointer to **string** | storage server, storage client or storage witness | [optional] 

## Methods

### NewCephHostReq

`func NewCephHostReq(hostId int64, ) *CephHostReq`

NewCephHostReq instantiates a new CephHostReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCephHostReqWithDefaults

`func NewCephHostReqWithDefaults() *CephHostReq`

NewCephHostReqWithDefaults instantiates a new CephHostReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayIps

`func (o *CephHostReq) GetGatewayIps() []string`

GetGatewayIps returns the GatewayIps field if non-nil, zero value otherwise.

### GetGatewayIpsOk

`func (o *CephHostReq) GetGatewayIpsOk() (*[]string, bool)`

GetGatewayIpsOk returns a tuple with the GatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIps

`func (o *CephHostReq) SetGatewayIps(v []string)`

SetGatewayIps sets GatewayIps field to given value.

### HasGatewayIps

`func (o *CephHostReq) HasGatewayIps() bool`

HasGatewayIps returns a boolean if a field has been set.

### GetHostId

`func (o *CephHostReq) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *CephHostReq) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *CephHostReq) SetHostId(v int64)`

SetHostId sets HostId field to given value.


### GetPrivateIp

`func (o *CephHostReq) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *CephHostReq) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *CephHostReq) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *CephHostReq) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPublicIp

`func (o *CephHostReq) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *CephHostReq) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *CephHostReq) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *CephHostReq) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRoles

`func (o *CephHostReq) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CephHostReq) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CephHostReq) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CephHostReq) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetType

`func (o *CephHostReq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CephHostReq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CephHostReq) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CephHostReq) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


