# AdminVIPCreateReqAdminVip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** |  | 
**Ip** | Pointer to **string** |  | [optional] 
**Mask** | Pointer to **int64** |  | [optional] 

## Methods

### NewAdminVIPCreateReqAdminVip

`func NewAdminVIPCreateReqAdminVip(enable bool, ) *AdminVIPCreateReqAdminVip`

NewAdminVIPCreateReqAdminVip instantiates a new AdminVIPCreateReqAdminVip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminVIPCreateReqAdminVipWithDefaults

`func NewAdminVIPCreateReqAdminVipWithDefaults() *AdminVIPCreateReqAdminVip`

NewAdminVIPCreateReqAdminVipWithDefaults instantiates a new AdminVIPCreateReqAdminVip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *AdminVIPCreateReqAdminVip) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AdminVIPCreateReqAdminVip) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AdminVIPCreateReqAdminVip) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetIp

`func (o *AdminVIPCreateReqAdminVip) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AdminVIPCreateReqAdminVip) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AdminVIPCreateReqAdminVip) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AdminVIPCreateReqAdminVip) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMask

`func (o *AdminVIPCreateReqAdminVip) GetMask() int64`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *AdminVIPCreateReqAdminVip) GetMaskOk() (*int64, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *AdminVIPCreateReqAdminVip) SetMask(v int64)`

SetMask sets Mask field to given value.

### HasMask

`func (o *AdminVIPCreateReqAdminVip) HasMask() bool`

HasMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


