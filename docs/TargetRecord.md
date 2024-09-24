# TargetRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPath** | Pointer to [**AccessPathNestview**](AccessPathNestview.md) |  | [optional] 
**Board** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DcName** | Pointer to **string** |  | [optional] 
**GatewayIps** | Pointer to **string** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Iqn** | Pointer to **string** |  | [optional] 
**Logout** | Pointer to **bool** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**LunsInfo** | Pointer to [**[]LunInfo**](LunInfo.md) |  | [optional] 

## Methods

### NewTargetRecord

`func NewTargetRecord() *TargetRecord`

NewTargetRecord instantiates a new TargetRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetRecordWithDefaults

`func NewTargetRecordWithDefaults() *TargetRecord`

NewTargetRecordWithDefaults instantiates a new TargetRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessPath

`func (o *TargetRecord) GetAccessPath() AccessPathNestview`

GetAccessPath returns the AccessPath field if non-nil, zero value otherwise.

### GetAccessPathOk

`func (o *TargetRecord) GetAccessPathOk() (*AccessPathNestview, bool)`

GetAccessPathOk returns a tuple with the AccessPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPath

`func (o *TargetRecord) SetAccessPath(v AccessPathNestview)`

SetAccessPath sets AccessPath field to given value.

### HasAccessPath

`func (o *TargetRecord) HasAccessPath() bool`

HasAccessPath returns a boolean if a field has been set.

### GetBoard

`func (o *TargetRecord) GetBoard() int64`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *TargetRecord) GetBoardOk() (*int64, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *TargetRecord) SetBoard(v int64)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *TargetRecord) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetCluster

`func (o *TargetRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *TargetRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *TargetRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *TargetRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *TargetRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *TargetRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *TargetRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *TargetRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDcName

`func (o *TargetRecord) GetDcName() string`

GetDcName returns the DcName field if non-nil, zero value otherwise.

### GetDcNameOk

`func (o *TargetRecord) GetDcNameOk() (*string, bool)`

GetDcNameOk returns a tuple with the DcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcName

`func (o *TargetRecord) SetDcName(v string)`

SetDcName sets DcName field to given value.

### HasDcName

`func (o *TargetRecord) HasDcName() bool`

HasDcName returns a boolean if a field has been set.

### GetGatewayIps

`func (o *TargetRecord) GetGatewayIps() string`

GetGatewayIps returns the GatewayIps field if non-nil, zero value otherwise.

### GetGatewayIpsOk

`func (o *TargetRecord) GetGatewayIpsOk() (*string, bool)`

GetGatewayIpsOk returns a tuple with the GatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIps

`func (o *TargetRecord) SetGatewayIps(v string)`

SetGatewayIps sets GatewayIps field to given value.

### HasGatewayIps

`func (o *TargetRecord) HasGatewayIps() bool`

HasGatewayIps returns a boolean if a field has been set.

### GetHost

`func (o *TargetRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TargetRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TargetRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *TargetRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *TargetRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TargetRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TargetRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TargetRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIqn

`func (o *TargetRecord) GetIqn() string`

GetIqn returns the Iqn field if non-nil, zero value otherwise.

### GetIqnOk

`func (o *TargetRecord) GetIqnOk() (*string, bool)`

GetIqnOk returns a tuple with the Iqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqn

`func (o *TargetRecord) SetIqn(v string)`

SetIqn sets Iqn field to given value.

### HasIqn

`func (o *TargetRecord) HasIqn() bool`

HasIqn returns a boolean if a field has been set.

### GetLogout

`func (o *TargetRecord) GetLogout() bool`

GetLogout returns the Logout field if non-nil, zero value otherwise.

### GetLogoutOk

`func (o *TargetRecord) GetLogoutOk() (*bool, bool)`

GetLogoutOk returns a tuple with the Logout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogout

`func (o *TargetRecord) SetLogout(v bool)`

SetLogout sets Logout field to given value.

### HasLogout

`func (o *TargetRecord) HasLogout() bool`

HasLogout returns a boolean if a field has been set.

### GetPort

`func (o *TargetRecord) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TargetRecord) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TargetRecord) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *TargetRecord) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *TargetRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TargetRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TargetRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TargetRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *TargetRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *TargetRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *TargetRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *TargetRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetLunsInfo

`func (o *TargetRecord) GetLunsInfo() []LunInfo`

GetLunsInfo returns the LunsInfo field if non-nil, zero value otherwise.

### GetLunsInfoOk

`func (o *TargetRecord) GetLunsInfoOk() (*[]LunInfo, bool)`

GetLunsInfoOk returns a tuple with the LunsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunsInfo

`func (o *TargetRecord) SetLunsInfo(v []LunInfo)`

SetLunsInfo sets LunsInfo field to given value.

### HasLunsInfo

`func (o *TargetRecord) HasLunsInfo() bool`

HasLunsInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


