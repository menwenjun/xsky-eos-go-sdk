# VIPArgs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** |  | 
**Mask** | Pointer to **int64** |  | [optional] 
**NetworkAddressId** | **int64** |  | 

## Methods

### NewVIPArgs

`func NewVIPArgs(ip string, networkAddressId int64, ) *VIPArgs`

NewVIPArgs instantiates a new VIPArgs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVIPArgsWithDefaults

`func NewVIPArgsWithDefaults() *VIPArgs`

NewVIPArgsWithDefaults instantiates a new VIPArgs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *VIPArgs) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *VIPArgs) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *VIPArgs) SetIp(v string)`

SetIp sets Ip field to given value.


### GetMask

`func (o *VIPArgs) GetMask() int64`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *VIPArgs) GetMaskOk() (*int64, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *VIPArgs) SetMask(v int64)`

SetMask sets Mask field to given value.

### HasMask

`func (o *VIPArgs) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetNetworkAddressId

`func (o *VIPArgs) GetNetworkAddressId() int64`

GetNetworkAddressId returns the NetworkAddressId field if non-nil, zero value otherwise.

### GetNetworkAddressIdOk

`func (o *VIPArgs) GetNetworkAddressIdOk() (*int64, bool)`

GetNetworkAddressIdOk returns a tuple with the NetworkAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddressId

`func (o *VIPArgs) SetNetworkAddressId(v int64)`

SetNetworkAddressId sets NetworkAddressId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


