# HostValidatorCreateReqValidator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminIp** | **string** |  | 
**CheckTypes** | Pointer to **[]string** |  | [optional] 
**ClusterId** | Pointer to **int64** |  | [optional] 
**GatewayIps** | Pointer to **[]string** |  | [optional] 
**HostId** | Pointer to **int64** |  | [optional] 
**HostRoles** | Pointer to **[]string** |  | [optional] 
**HostType** | Pointer to **string** |  | [optional] 
**OspClusterIp** | Pointer to **string** |  | [optional] 
**OspGatewayIps** | Pointer to **[]string** |  | [optional] 
**PrivateIp** | Pointer to **string** |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 

## Methods

### NewHostValidatorCreateReqValidator

`func NewHostValidatorCreateReqValidator(adminIp string, ) *HostValidatorCreateReqValidator`

NewHostValidatorCreateReqValidator instantiates a new HostValidatorCreateReqValidator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostValidatorCreateReqValidatorWithDefaults

`func NewHostValidatorCreateReqValidatorWithDefaults() *HostValidatorCreateReqValidator`

NewHostValidatorCreateReqValidatorWithDefaults instantiates a new HostValidatorCreateReqValidator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminIp

`func (o *HostValidatorCreateReqValidator) GetAdminIp() string`

GetAdminIp returns the AdminIp field if non-nil, zero value otherwise.

### GetAdminIpOk

`func (o *HostValidatorCreateReqValidator) GetAdminIpOk() (*string, bool)`

GetAdminIpOk returns a tuple with the AdminIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminIp

`func (o *HostValidatorCreateReqValidator) SetAdminIp(v string)`

SetAdminIp sets AdminIp field to given value.


### GetCheckTypes

`func (o *HostValidatorCreateReqValidator) GetCheckTypes() []string`

GetCheckTypes returns the CheckTypes field if non-nil, zero value otherwise.

### GetCheckTypesOk

`func (o *HostValidatorCreateReqValidator) GetCheckTypesOk() (*[]string, bool)`

GetCheckTypesOk returns a tuple with the CheckTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTypes

`func (o *HostValidatorCreateReqValidator) SetCheckTypes(v []string)`

SetCheckTypes sets CheckTypes field to given value.

### HasCheckTypes

`func (o *HostValidatorCreateReqValidator) HasCheckTypes() bool`

HasCheckTypes returns a boolean if a field has been set.

### GetClusterId

`func (o *HostValidatorCreateReqValidator) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HostValidatorCreateReqValidator) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HostValidatorCreateReqValidator) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *HostValidatorCreateReqValidator) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetGatewayIps

`func (o *HostValidatorCreateReqValidator) GetGatewayIps() []string`

GetGatewayIps returns the GatewayIps field if non-nil, zero value otherwise.

### GetGatewayIpsOk

`func (o *HostValidatorCreateReqValidator) GetGatewayIpsOk() (*[]string, bool)`

GetGatewayIpsOk returns a tuple with the GatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIps

`func (o *HostValidatorCreateReqValidator) SetGatewayIps(v []string)`

SetGatewayIps sets GatewayIps field to given value.

### HasGatewayIps

`func (o *HostValidatorCreateReqValidator) HasGatewayIps() bool`

HasGatewayIps returns a boolean if a field has been set.

### GetHostId

`func (o *HostValidatorCreateReqValidator) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *HostValidatorCreateReqValidator) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *HostValidatorCreateReqValidator) SetHostId(v int64)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *HostValidatorCreateReqValidator) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetHostRoles

`func (o *HostValidatorCreateReqValidator) GetHostRoles() []string`

GetHostRoles returns the HostRoles field if non-nil, zero value otherwise.

### GetHostRolesOk

`func (o *HostValidatorCreateReqValidator) GetHostRolesOk() (*[]string, bool)`

GetHostRolesOk returns a tuple with the HostRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostRoles

`func (o *HostValidatorCreateReqValidator) SetHostRoles(v []string)`

SetHostRoles sets HostRoles field to given value.

### HasHostRoles

`func (o *HostValidatorCreateReqValidator) HasHostRoles() bool`

HasHostRoles returns a boolean if a field has been set.

### GetHostType

`func (o *HostValidatorCreateReqValidator) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *HostValidatorCreateReqValidator) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *HostValidatorCreateReqValidator) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *HostValidatorCreateReqValidator) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### GetOspClusterIp

`func (o *HostValidatorCreateReqValidator) GetOspClusterIp() string`

GetOspClusterIp returns the OspClusterIp field if non-nil, zero value otherwise.

### GetOspClusterIpOk

`func (o *HostValidatorCreateReqValidator) GetOspClusterIpOk() (*string, bool)`

GetOspClusterIpOk returns a tuple with the OspClusterIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspClusterIp

`func (o *HostValidatorCreateReqValidator) SetOspClusterIp(v string)`

SetOspClusterIp sets OspClusterIp field to given value.

### HasOspClusterIp

`func (o *HostValidatorCreateReqValidator) HasOspClusterIp() bool`

HasOspClusterIp returns a boolean if a field has been set.

### GetOspGatewayIps

`func (o *HostValidatorCreateReqValidator) GetOspGatewayIps() []string`

GetOspGatewayIps returns the OspGatewayIps field if non-nil, zero value otherwise.

### GetOspGatewayIpsOk

`func (o *HostValidatorCreateReqValidator) GetOspGatewayIpsOk() (*[]string, bool)`

GetOspGatewayIpsOk returns a tuple with the OspGatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspGatewayIps

`func (o *HostValidatorCreateReqValidator) SetOspGatewayIps(v []string)`

SetOspGatewayIps sets OspGatewayIps field to given value.

### HasOspGatewayIps

`func (o *HostValidatorCreateReqValidator) HasOspGatewayIps() bool`

HasOspGatewayIps returns a boolean if a field has been set.

### GetPrivateIp

`func (o *HostValidatorCreateReqValidator) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *HostValidatorCreateReqValidator) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *HostValidatorCreateReqValidator) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *HostValidatorCreateReqValidator) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPublicIp

`func (o *HostValidatorCreateReqValidator) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *HostValidatorCreateReqValidator) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *HostValidatorCreateReqValidator) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *HostValidatorCreateReqValidator) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


