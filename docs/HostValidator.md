# HostValidator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminIp** | Pointer to **string** |  | [optional] 
**CheckTypes** | Pointer to **[]string** |  | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**GatewayIps** | Pointer to **[]string** |  | [optional] 
**Host** | Pointer to [**Host**](Host.md) |  | [optional] 
**HostRoles** | Pointer to **[]string** |  | [optional] 
**HostType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**OspClusterIp** | Pointer to **string** |  | [optional] 
**OspGatewayIps** | Pointer to **[]string** |  | [optional] 
**PrivateIp** | Pointer to **string** |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**Report** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewHostValidator

`func NewHostValidator() *HostValidator`

NewHostValidator instantiates a new HostValidator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostValidatorWithDefaults

`func NewHostValidatorWithDefaults() *HostValidator`

NewHostValidatorWithDefaults instantiates a new HostValidator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminIp

`func (o *HostValidator) GetAdminIp() string`

GetAdminIp returns the AdminIp field if non-nil, zero value otherwise.

### GetAdminIpOk

`func (o *HostValidator) GetAdminIpOk() (*string, bool)`

GetAdminIpOk returns a tuple with the AdminIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminIp

`func (o *HostValidator) SetAdminIp(v string)`

SetAdminIp sets AdminIp field to given value.

### HasAdminIp

`func (o *HostValidator) HasAdminIp() bool`

HasAdminIp returns a boolean if a field has been set.

### GetCheckTypes

`func (o *HostValidator) GetCheckTypes() []string`

GetCheckTypes returns the CheckTypes field if non-nil, zero value otherwise.

### GetCheckTypesOk

`func (o *HostValidator) GetCheckTypesOk() (*[]string, bool)`

GetCheckTypesOk returns a tuple with the CheckTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTypes

`func (o *HostValidator) SetCheckTypes(v []string)`

SetCheckTypes sets CheckTypes field to given value.

### HasCheckTypes

`func (o *HostValidator) HasCheckTypes() bool`

HasCheckTypes returns a boolean if a field has been set.

### GetCluster

`func (o *HostValidator) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HostValidator) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HostValidator) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HostValidator) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *HostValidator) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *HostValidator) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *HostValidator) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *HostValidator) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetGatewayIps

`func (o *HostValidator) GetGatewayIps() []string`

GetGatewayIps returns the GatewayIps field if non-nil, zero value otherwise.

### GetGatewayIpsOk

`func (o *HostValidator) GetGatewayIpsOk() (*[]string, bool)`

GetGatewayIpsOk returns a tuple with the GatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIps

`func (o *HostValidator) SetGatewayIps(v []string)`

SetGatewayIps sets GatewayIps field to given value.

### HasGatewayIps

`func (o *HostValidator) HasGatewayIps() bool`

HasGatewayIps returns a boolean if a field has been set.

### GetHost

`func (o *HostValidator) GetHost() Host`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HostValidator) GetHostOk() (*Host, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HostValidator) SetHost(v Host)`

SetHost sets Host field to given value.

### HasHost

`func (o *HostValidator) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHostRoles

`func (o *HostValidator) GetHostRoles() []string`

GetHostRoles returns the HostRoles field if non-nil, zero value otherwise.

### GetHostRolesOk

`func (o *HostValidator) GetHostRolesOk() (*[]string, bool)`

GetHostRolesOk returns a tuple with the HostRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostRoles

`func (o *HostValidator) SetHostRoles(v []string)`

SetHostRoles sets HostRoles field to given value.

### HasHostRoles

`func (o *HostValidator) HasHostRoles() bool`

HasHostRoles returns a boolean if a field has been set.

### GetHostType

`func (o *HostValidator) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *HostValidator) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *HostValidator) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *HostValidator) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### GetId

`func (o *HostValidator) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostValidator) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostValidator) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *HostValidator) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOspClusterIp

`func (o *HostValidator) GetOspClusterIp() string`

GetOspClusterIp returns the OspClusterIp field if non-nil, zero value otherwise.

### GetOspClusterIpOk

`func (o *HostValidator) GetOspClusterIpOk() (*string, bool)`

GetOspClusterIpOk returns a tuple with the OspClusterIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspClusterIp

`func (o *HostValidator) SetOspClusterIp(v string)`

SetOspClusterIp sets OspClusterIp field to given value.

### HasOspClusterIp

`func (o *HostValidator) HasOspClusterIp() bool`

HasOspClusterIp returns a boolean if a field has been set.

### GetOspGatewayIps

`func (o *HostValidator) GetOspGatewayIps() []string`

GetOspGatewayIps returns the OspGatewayIps field if non-nil, zero value otherwise.

### GetOspGatewayIpsOk

`func (o *HostValidator) GetOspGatewayIpsOk() (*[]string, bool)`

GetOspGatewayIpsOk returns a tuple with the OspGatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspGatewayIps

`func (o *HostValidator) SetOspGatewayIps(v []string)`

SetOspGatewayIps sets OspGatewayIps field to given value.

### HasOspGatewayIps

`func (o *HostValidator) HasOspGatewayIps() bool`

HasOspGatewayIps returns a boolean if a field has been set.

### GetPrivateIp

`func (o *HostValidator) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *HostValidator) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *HostValidator) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *HostValidator) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPublicIp

`func (o *HostValidator) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *HostValidator) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *HostValidator) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *HostValidator) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetReport

`func (o *HostValidator) GetReport() map[string]interface{}`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *HostValidator) GetReportOk() (*map[string]interface{}, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *HostValidator) SetReport(v map[string]interface{})`

SetReport sets Report field to given value.

### HasReport

`func (o *HostValidator) HasReport() bool`

HasReport returns a boolean if a field has been set.

### GetStatus

`func (o *HostValidator) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostValidator) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostValidator) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HostValidator) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *HostValidator) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *HostValidator) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *HostValidator) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *HostValidator) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


