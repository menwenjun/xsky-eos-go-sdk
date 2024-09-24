# VIPRecord

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
**CurrentVipInstance** | Pointer to [**VIPInstanceNestview**](VIPInstanceNestview.md) |  | [optional] 
**DefaultVipInstance** | Pointer to [**VIPInstanceNestview**](VIPInstanceNestview.md) |  | [optional] 

## Methods

### NewVIPRecord

`func NewVIPRecord() *VIPRecord`

NewVIPRecord instantiates a new VIPRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVIPRecordWithDefaults

`func NewVIPRecordWithDefaults() *VIPRecord`

NewVIPRecordWithDefaults instantiates a new VIPRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *VIPRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *VIPRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *VIPRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *VIPRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCreate

`func (o *VIPRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *VIPRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *VIPRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *VIPRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *VIPRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VIPRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VIPRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VIPRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *VIPRecord) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *VIPRecord) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *VIPRecord) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *VIPRecord) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMacAddress

`func (o *VIPRecord) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VIPRecord) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VIPRecord) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VIPRecord) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMask

`func (o *VIPRecord) GetMask() int64`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *VIPRecord) GetMaskOk() (*int64, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *VIPRecord) SetMask(v int64)`

SetMask sets Mask field to given value.

### HasMask

`func (o *VIPRecord) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetStatus

`func (o *VIPRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VIPRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VIPRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VIPRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *VIPRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *VIPRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *VIPRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *VIPRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVipGroup

`func (o *VIPRecord) GetVipGroup() VIPGroupNestview`

GetVipGroup returns the VipGroup field if non-nil, zero value otherwise.

### GetVipGroupOk

`func (o *VIPRecord) GetVipGroupOk() (*VIPGroupNestview, bool)`

GetVipGroupOk returns a tuple with the VipGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipGroup

`func (o *VIPRecord) SetVipGroup(v VIPGroupNestview)`

SetVipGroup sets VipGroup field to given value.

### HasVipGroup

`func (o *VIPRecord) HasVipGroup() bool`

HasVipGroup returns a boolean if a field has been set.

### GetVirtualRouterId

`func (o *VIPRecord) GetVirtualRouterId() int64`

GetVirtualRouterId returns the VirtualRouterId field if non-nil, zero value otherwise.

### GetVirtualRouterIdOk

`func (o *VIPRecord) GetVirtualRouterIdOk() (*int64, bool)`

GetVirtualRouterIdOk returns a tuple with the VirtualRouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualRouterId

`func (o *VIPRecord) SetVirtualRouterId(v int64)`

SetVirtualRouterId sets VirtualRouterId field to given value.

### HasVirtualRouterId

`func (o *VIPRecord) HasVirtualRouterId() bool`

HasVirtualRouterId returns a boolean if a field has been set.

### GetCurrentVipInstance

`func (o *VIPRecord) GetCurrentVipInstance() VIPInstanceNestview`

GetCurrentVipInstance returns the CurrentVipInstance field if non-nil, zero value otherwise.

### GetCurrentVipInstanceOk

`func (o *VIPRecord) GetCurrentVipInstanceOk() (*VIPInstanceNestview, bool)`

GetCurrentVipInstanceOk returns a tuple with the CurrentVipInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVipInstance

`func (o *VIPRecord) SetCurrentVipInstance(v VIPInstanceNestview)`

SetCurrentVipInstance sets CurrentVipInstance field to given value.

### HasCurrentVipInstance

`func (o *VIPRecord) HasCurrentVipInstance() bool`

HasCurrentVipInstance returns a boolean if a field has been set.

### GetDefaultVipInstance

`func (o *VIPRecord) GetDefaultVipInstance() VIPInstanceNestview`

GetDefaultVipInstance returns the DefaultVipInstance field if non-nil, zero value otherwise.

### GetDefaultVipInstanceOk

`func (o *VIPRecord) GetDefaultVipInstanceOk() (*VIPInstanceNestview, bool)`

GetDefaultVipInstanceOk returns a tuple with the DefaultVipInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVipInstance

`func (o *VIPRecord) SetDefaultVipInstance(v VIPInstanceNestview)`

SetDefaultVipInstance sets DefaultVipInstance field to given value.

### HasDefaultVipInstance

`func (o *VIPRecord) HasDefaultVipInstance() bool`

HasDefaultVipInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


