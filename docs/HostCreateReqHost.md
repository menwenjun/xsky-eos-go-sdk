# HostCreateReqHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminIp** | **string** | admin ip | 
**ClusterId** | **int64** | cluster id | 
**Description** | Pointer to **string** | host description | [optional] 
**GatewayIps** | Pointer to **[]string** | gateway ips for s3 | [optional] 
**MetaDevice** | Pointer to **string** | meta device for docker | [optional] 
**OspClusterId** | Pointer to **int64** | osp cluster id | [optional] 
**OspClusterIp** | Pointer to **string** | osp cluster ip | [optional] 
**OspGatewayIps** | Pointer to **[]string** | osp gateway ips | [optional] 
**PrivateIp** | Pointer to **string** | cluster private ip for internal data access | [optional] 
**ProtectionDomainId** | Pointer to **int64** | deprecated | [optional] 
**PublicIp** | Pointer to **string** | public ip for outside data access | [optional] 
**Roles** | Pointer to **[]string** | host roles: admin,monitor,block_storage_gateway,file_storage_gateway,s3_gateway,nfs_gateway | [optional] 
**Type** | Pointer to **string** | storage server or storage client | [optional] 

## Methods

### NewHostCreateReqHost

`func NewHostCreateReqHost(adminIp string, clusterId int64, ) *HostCreateReqHost`

NewHostCreateReqHost instantiates a new HostCreateReqHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostCreateReqHostWithDefaults

`func NewHostCreateReqHostWithDefaults() *HostCreateReqHost`

NewHostCreateReqHostWithDefaults instantiates a new HostCreateReqHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminIp

`func (o *HostCreateReqHost) GetAdminIp() string`

GetAdminIp returns the AdminIp field if non-nil, zero value otherwise.

### GetAdminIpOk

`func (o *HostCreateReqHost) GetAdminIpOk() (*string, bool)`

GetAdminIpOk returns a tuple with the AdminIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminIp

`func (o *HostCreateReqHost) SetAdminIp(v string)`

SetAdminIp sets AdminIp field to given value.


### GetClusterId

`func (o *HostCreateReqHost) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HostCreateReqHost) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HostCreateReqHost) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.


### GetDescription

`func (o *HostCreateReqHost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HostCreateReqHost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HostCreateReqHost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HostCreateReqHost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGatewayIps

`func (o *HostCreateReqHost) GetGatewayIps() []string`

GetGatewayIps returns the GatewayIps field if non-nil, zero value otherwise.

### GetGatewayIpsOk

`func (o *HostCreateReqHost) GetGatewayIpsOk() (*[]string, bool)`

GetGatewayIpsOk returns a tuple with the GatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIps

`func (o *HostCreateReqHost) SetGatewayIps(v []string)`

SetGatewayIps sets GatewayIps field to given value.

### HasGatewayIps

`func (o *HostCreateReqHost) HasGatewayIps() bool`

HasGatewayIps returns a boolean if a field has been set.

### GetMetaDevice

`func (o *HostCreateReqHost) GetMetaDevice() string`

GetMetaDevice returns the MetaDevice field if non-nil, zero value otherwise.

### GetMetaDeviceOk

`func (o *HostCreateReqHost) GetMetaDeviceOk() (*string, bool)`

GetMetaDeviceOk returns a tuple with the MetaDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDevice

`func (o *HostCreateReqHost) SetMetaDevice(v string)`

SetMetaDevice sets MetaDevice field to given value.

### HasMetaDevice

`func (o *HostCreateReqHost) HasMetaDevice() bool`

HasMetaDevice returns a boolean if a field has been set.

### GetOspClusterId

`func (o *HostCreateReqHost) GetOspClusterId() int64`

GetOspClusterId returns the OspClusterId field if non-nil, zero value otherwise.

### GetOspClusterIdOk

`func (o *HostCreateReqHost) GetOspClusterIdOk() (*int64, bool)`

GetOspClusterIdOk returns a tuple with the OspClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspClusterId

`func (o *HostCreateReqHost) SetOspClusterId(v int64)`

SetOspClusterId sets OspClusterId field to given value.

### HasOspClusterId

`func (o *HostCreateReqHost) HasOspClusterId() bool`

HasOspClusterId returns a boolean if a field has been set.

### GetOspClusterIp

`func (o *HostCreateReqHost) GetOspClusterIp() string`

GetOspClusterIp returns the OspClusterIp field if non-nil, zero value otherwise.

### GetOspClusterIpOk

`func (o *HostCreateReqHost) GetOspClusterIpOk() (*string, bool)`

GetOspClusterIpOk returns a tuple with the OspClusterIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspClusterIp

`func (o *HostCreateReqHost) SetOspClusterIp(v string)`

SetOspClusterIp sets OspClusterIp field to given value.

### HasOspClusterIp

`func (o *HostCreateReqHost) HasOspClusterIp() bool`

HasOspClusterIp returns a boolean if a field has been set.

### GetOspGatewayIps

`func (o *HostCreateReqHost) GetOspGatewayIps() []string`

GetOspGatewayIps returns the OspGatewayIps field if non-nil, zero value otherwise.

### GetOspGatewayIpsOk

`func (o *HostCreateReqHost) GetOspGatewayIpsOk() (*[]string, bool)`

GetOspGatewayIpsOk returns a tuple with the OspGatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspGatewayIps

`func (o *HostCreateReqHost) SetOspGatewayIps(v []string)`

SetOspGatewayIps sets OspGatewayIps field to given value.

### HasOspGatewayIps

`func (o *HostCreateReqHost) HasOspGatewayIps() bool`

HasOspGatewayIps returns a boolean if a field has been set.

### GetPrivateIp

`func (o *HostCreateReqHost) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *HostCreateReqHost) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *HostCreateReqHost) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *HostCreateReqHost) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetProtectionDomainId

`func (o *HostCreateReqHost) GetProtectionDomainId() int64`

GetProtectionDomainId returns the ProtectionDomainId field if non-nil, zero value otherwise.

### GetProtectionDomainIdOk

`func (o *HostCreateReqHost) GetProtectionDomainIdOk() (*int64, bool)`

GetProtectionDomainIdOk returns a tuple with the ProtectionDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionDomainId

`func (o *HostCreateReqHost) SetProtectionDomainId(v int64)`

SetProtectionDomainId sets ProtectionDomainId field to given value.

### HasProtectionDomainId

`func (o *HostCreateReqHost) HasProtectionDomainId() bool`

HasProtectionDomainId returns a boolean if a field has been set.

### GetPublicIp

`func (o *HostCreateReqHost) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *HostCreateReqHost) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *HostCreateReqHost) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *HostCreateReqHost) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRoles

`func (o *HostCreateReqHost) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *HostCreateReqHost) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *HostCreateReqHost) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *HostCreateReqHost) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetType

`func (o *HostCreateReqHost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HostCreateReqHost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HostCreateReqHost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HostCreateReqHost) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


