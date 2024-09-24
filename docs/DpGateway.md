# DpGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminPort** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**GatewayPort** | Pointer to **int64** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PolicyNum** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDpGateway

`func NewDpGateway() *DpGateway`

NewDpGateway instantiates a new DpGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpGatewayWithDefaults

`func NewDpGatewayWithDefaults() *DpGateway`

NewDpGatewayWithDefaults instantiates a new DpGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminPort

`func (o *DpGateway) GetAdminPort() int64`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *DpGateway) GetAdminPortOk() (*int64, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *DpGateway) SetAdminPort(v int64)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *DpGateway) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### GetCluster

`func (o *DpGateway) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpGateway) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpGateway) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpGateway) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DpGateway) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DpGateway) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DpGateway) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DpGateway) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetGatewayPort

`func (o *DpGateway) GetGatewayPort() int64`

GetGatewayPort returns the GatewayPort field if non-nil, zero value otherwise.

### GetGatewayPortOk

`func (o *DpGateway) GetGatewayPortOk() (*int64, bool)`

GetGatewayPortOk returns a tuple with the GatewayPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPort

`func (o *DpGateway) SetGatewayPort(v int64)`

SetGatewayPort sets GatewayPort field to given value.

### HasGatewayPort

`func (o *DpGateway) HasGatewayPort() bool`

HasGatewayPort returns a boolean if a field has been set.

### GetHost

`func (o *DpGateway) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DpGateway) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DpGateway) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *DpGateway) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *DpGateway) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpGateway) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpGateway) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DpGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpGateway) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DpGateway) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyNum

`func (o *DpGateway) GetPolicyNum() int64`

GetPolicyNum returns the PolicyNum field if non-nil, zero value otherwise.

### GetPolicyNumOk

`func (o *DpGateway) GetPolicyNumOk() (*int64, bool)`

GetPolicyNumOk returns a tuple with the PolicyNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyNum

`func (o *DpGateway) SetPolicyNum(v int64)`

SetPolicyNum sets PolicyNum field to given value.

### HasPolicyNum

`func (o *DpGateway) HasPolicyNum() bool`

HasPolicyNum returns a boolean if a field has been set.

### GetStatus

`func (o *DpGateway) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpGateway) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpGateway) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpGateway) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DpGateway) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DpGateway) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DpGateway) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DpGateway) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


