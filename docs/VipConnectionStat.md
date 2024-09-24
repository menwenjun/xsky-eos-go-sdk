# VipConnectionStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to **int64** |  | [optional] 
**Vip** | Pointer to **string** |  | [optional] 

## Methods

### NewVipConnectionStat

`func NewVipConnectionStat() *VipConnectionStat`

NewVipConnectionStat instantiates a new VipConnectionStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVipConnectionStatWithDefaults

`func NewVipConnectionStatWithDefaults() *VipConnectionStat`

NewVipConnectionStatWithDefaults instantiates a new VipConnectionStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *VipConnectionStat) GetConnections() int64`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *VipConnectionStat) GetConnectionsOk() (*int64, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *VipConnectionStat) SetConnections(v int64)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *VipConnectionStat) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetVip

`func (o *VipConnectionStat) GetVip() string`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *VipConnectionStat) GetVipOk() (*string, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *VipConnectionStat) SetVip(v string)`

SetVip sets Vip field to given value.

### HasVip

`func (o *VipConnectionStat) HasVip() bool`

HasVip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


