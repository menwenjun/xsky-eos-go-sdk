# HostNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**AdminIp** | Pointer to **string** |  | [optional] 
**AdminVip** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OspClusterId** | Pointer to **int64** |  | [optional] 
**OspZoneId** | Pointer to **int64** |  | [optional] 
**Roles** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 

## Methods

### NewHostNode

`func NewHostNode() *HostNode`

NewHostNode instantiates a new HostNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostNodeWithDefaults

`func NewHostNodeWithDefaults() *HostNode`

NewHostNodeWithDefaults instantiates a new HostNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *HostNode) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *HostNode) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *HostNode) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *HostNode) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAdminIp

`func (o *HostNode) GetAdminIp() string`

GetAdminIp returns the AdminIp field if non-nil, zero value otherwise.

### GetAdminIpOk

`func (o *HostNode) GetAdminIpOk() (*string, bool)`

GetAdminIpOk returns a tuple with the AdminIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminIp

`func (o *HostNode) SetAdminIp(v string)`

SetAdminIp sets AdminIp field to given value.

### HasAdminIp

`func (o *HostNode) HasAdminIp() bool`

HasAdminIp returns a boolean if a field has been set.

### GetAdminVip

`func (o *HostNode) GetAdminVip() string`

GetAdminVip returns the AdminVip field if non-nil, zero value otherwise.

### GetAdminVipOk

`func (o *HostNode) GetAdminVipOk() (*string, bool)`

GetAdminVipOk returns a tuple with the AdminVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminVip

`func (o *HostNode) SetAdminVip(v string)`

SetAdminVip sets AdminVip field to given value.

### HasAdminVip

`func (o *HostNode) HasAdminVip() bool`

HasAdminVip returns a boolean if a field has been set.

### GetId

`func (o *HostNode) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostNode) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostNode) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *HostNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HostNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HostNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOspClusterId

`func (o *HostNode) GetOspClusterId() int64`

GetOspClusterId returns the OspClusterId field if non-nil, zero value otherwise.

### GetOspClusterIdOk

`func (o *HostNode) GetOspClusterIdOk() (*int64, bool)`

GetOspClusterIdOk returns a tuple with the OspClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspClusterId

`func (o *HostNode) SetOspClusterId(v int64)`

SetOspClusterId sets OspClusterId field to given value.

### HasOspClusterId

`func (o *HostNode) HasOspClusterId() bool`

HasOspClusterId returns a boolean if a field has been set.

### GetOspZoneId

`func (o *HostNode) GetOspZoneId() int64`

GetOspZoneId returns the OspZoneId field if non-nil, zero value otherwise.

### GetOspZoneIdOk

`func (o *HostNode) GetOspZoneIdOk() (*int64, bool)`

GetOspZoneIdOk returns a tuple with the OspZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspZoneId

`func (o *HostNode) SetOspZoneId(v int64)`

SetOspZoneId sets OspZoneId field to given value.

### HasOspZoneId

`func (o *HostNode) HasOspZoneId() bool`

HasOspZoneId returns a boolean if a field has been set.

### GetRoles

`func (o *HostNode) GetRoles() string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *HostNode) GetRolesOk() (*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *HostNode) SetRoles(v string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *HostNode) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetStatus

`func (o *HostNode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostNode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostNode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HostNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUp

`func (o *HostNode) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *HostNode) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *HostNode) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *HostNode) HasUp() bool`

HasUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


