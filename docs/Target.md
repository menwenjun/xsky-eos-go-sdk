# Target

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

## Methods

### NewTarget

`func NewTarget() *Target`

NewTarget instantiates a new Target object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetWithDefaults

`func NewTargetWithDefaults() *Target`

NewTargetWithDefaults instantiates a new Target object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessPath

`func (o *Target) GetAccessPath() AccessPathNestview`

GetAccessPath returns the AccessPath field if non-nil, zero value otherwise.

### GetAccessPathOk

`func (o *Target) GetAccessPathOk() (*AccessPathNestview, bool)`

GetAccessPathOk returns a tuple with the AccessPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPath

`func (o *Target) SetAccessPath(v AccessPathNestview)`

SetAccessPath sets AccessPath field to given value.

### HasAccessPath

`func (o *Target) HasAccessPath() bool`

HasAccessPath returns a boolean if a field has been set.

### GetBoard

`func (o *Target) GetBoard() int64`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *Target) GetBoardOk() (*int64, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *Target) SetBoard(v int64)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *Target) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetCluster

`func (o *Target) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Target) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Target) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Target) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *Target) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Target) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Target) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Target) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDcName

`func (o *Target) GetDcName() string`

GetDcName returns the DcName field if non-nil, zero value otherwise.

### GetDcNameOk

`func (o *Target) GetDcNameOk() (*string, bool)`

GetDcNameOk returns a tuple with the DcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcName

`func (o *Target) SetDcName(v string)`

SetDcName sets DcName field to given value.

### HasDcName

`func (o *Target) HasDcName() bool`

HasDcName returns a boolean if a field has been set.

### GetGatewayIps

`func (o *Target) GetGatewayIps() string`

GetGatewayIps returns the GatewayIps field if non-nil, zero value otherwise.

### GetGatewayIpsOk

`func (o *Target) GetGatewayIpsOk() (*string, bool)`

GetGatewayIpsOk returns a tuple with the GatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIps

`func (o *Target) SetGatewayIps(v string)`

SetGatewayIps sets GatewayIps field to given value.

### HasGatewayIps

`func (o *Target) HasGatewayIps() bool`

HasGatewayIps returns a boolean if a field has been set.

### GetHost

`func (o *Target) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Target) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Target) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *Target) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *Target) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Target) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Target) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Target) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIqn

`func (o *Target) GetIqn() string`

GetIqn returns the Iqn field if non-nil, zero value otherwise.

### GetIqnOk

`func (o *Target) GetIqnOk() (*string, bool)`

GetIqnOk returns a tuple with the Iqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqn

`func (o *Target) SetIqn(v string)`

SetIqn sets Iqn field to given value.

### HasIqn

`func (o *Target) HasIqn() bool`

HasIqn returns a boolean if a field has been set.

### GetLogout

`func (o *Target) GetLogout() bool`

GetLogout returns the Logout field if non-nil, zero value otherwise.

### GetLogoutOk

`func (o *Target) GetLogoutOk() (*bool, bool)`

GetLogoutOk returns a tuple with the Logout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogout

`func (o *Target) SetLogout(v bool)`

SetLogout sets Logout field to given value.

### HasLogout

`func (o *Target) HasLogout() bool`

HasLogout returns a boolean if a field has been set.

### GetPort

`func (o *Target) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Target) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Target) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *Target) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *Target) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Target) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Target) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Target) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *Target) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Target) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Target) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Target) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


