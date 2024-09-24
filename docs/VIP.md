# VIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**Mask** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**VipGroup** | Pointer to [**VIPGroupNestview**](VIPGroupNestview.md) |  | [optional] 
**VirtualRouterId** | Pointer to **int64** |  | [optional] 

## Methods

### NewVIP

`func NewVIP() *VIP`

NewVIP instantiates a new VIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVIPWithDefaults

`func NewVIPWithDefaults() *VIP`

NewVIPWithDefaults instantiates a new VIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *VIP) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *VIP) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *VIP) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *VIP) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCreate

`func (o *VIP) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *VIP) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *VIP) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *VIP) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *VIP) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VIP) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VIP) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VIP) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *VIP) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *VIP) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *VIP) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *VIP) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMacAddress

`func (o *VIP) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VIP) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VIP) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VIP) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMask

`func (o *VIP) GetMask() int64`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *VIP) GetMaskOk() (*int64, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *VIP) SetMask(v int64)`

SetMask sets Mask field to given value.

### HasMask

`func (o *VIP) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetStatus

`func (o *VIP) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VIP) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VIP) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VIP) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *VIP) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *VIP) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *VIP) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *VIP) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVipGroup

`func (o *VIP) GetVipGroup() VIPGroupNestview`

GetVipGroup returns the VipGroup field if non-nil, zero value otherwise.

### GetVipGroupOk

`func (o *VIP) GetVipGroupOk() (*VIPGroupNestview, bool)`

GetVipGroupOk returns a tuple with the VipGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipGroup

`func (o *VIP) SetVipGroup(v VIPGroupNestview)`

SetVipGroup sets VipGroup field to given value.

### HasVipGroup

`func (o *VIP) HasVipGroup() bool`

HasVipGroup returns a boolean if a field has been set.

### GetVirtualRouterId

`func (o *VIP) GetVirtualRouterId() int64`

GetVirtualRouterId returns the VirtualRouterId field if non-nil, zero value otherwise.

### GetVirtualRouterIdOk

`func (o *VIP) GetVirtualRouterIdOk() (*int64, bool)`

GetVirtualRouterIdOk returns a tuple with the VirtualRouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualRouterId

`func (o *VIP) SetVirtualRouterId(v int64)`

SetVirtualRouterId sets VirtualRouterId field to given value.

### HasVirtualRouterId

`func (o *VIP) HasVirtualRouterId() bool`

HasVirtualRouterId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


