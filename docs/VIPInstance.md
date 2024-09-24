# VIPInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**NetworkAddress** | Pointer to [**NetworkAddressNestview**](NetworkAddressNestview.md) |  | [optional] 
**Priority** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Vip** | Pointer to [**VIPNestview**](VIPNestview.md) |  | [optional] 

## Methods

### NewVIPInstance

`func NewVIPInstance() *VIPInstance`

NewVIPInstance instantiates a new VIPInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVIPInstanceWithDefaults

`func NewVIPInstanceWithDefaults() *VIPInstance`

NewVIPInstanceWithDefaults instantiates a new VIPInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *VIPInstance) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *VIPInstance) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *VIPInstance) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *VIPInstance) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCreate

`func (o *VIPInstance) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *VIPInstance) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *VIPInstance) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *VIPInstance) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *VIPInstance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VIPInstance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VIPInstance) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VIPInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkAddress

`func (o *VIPInstance) GetNetworkAddress() NetworkAddressNestview`

GetNetworkAddress returns the NetworkAddress field if non-nil, zero value otherwise.

### GetNetworkAddressOk

`func (o *VIPInstance) GetNetworkAddressOk() (*NetworkAddressNestview, bool)`

GetNetworkAddressOk returns a tuple with the NetworkAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddress

`func (o *VIPInstance) SetNetworkAddress(v NetworkAddressNestview)`

SetNetworkAddress sets NetworkAddress field to given value.

### HasNetworkAddress

`func (o *VIPInstance) HasNetworkAddress() bool`

HasNetworkAddress returns a boolean if a field has been set.

### GetPriority

`func (o *VIPInstance) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *VIPInstance) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *VIPInstance) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *VIPInstance) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *VIPInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VIPInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VIPInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VIPInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *VIPInstance) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *VIPInstance) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *VIPInstance) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *VIPInstance) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVip

`func (o *VIPInstance) GetVip() VIPNestview`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *VIPInstance) GetVipOk() (*VIPNestview, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *VIPInstance) SetVip(v VIPNestview)`

SetVip sets Vip field to given value.

### HasVip

`func (o *VIPInstance) HasVip() bool`

HasVip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


