# DfsGatewayGroupUpdateVIPsReqGatewayGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsGatewayZoneId** | Pointer to **int64** | gateway zone to which the VIP belongs | [optional] 
**DfsVipGateways** | Pointer to **[]string** | all gateways used in VIP networks | [optional] 
**DfsVips** | Pointer to **[]string** | vip info of dfs gateway group | [optional] 

## Methods

### NewDfsGatewayGroupUpdateVIPsReqGatewayGroup

`func NewDfsGatewayGroupUpdateVIPsReqGatewayGroup() *DfsGatewayGroupUpdateVIPsReqGatewayGroup`

NewDfsGatewayGroupUpdateVIPsReqGatewayGroup instantiates a new DfsGatewayGroupUpdateVIPsReqGatewayGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayGroupUpdateVIPsReqGatewayGroupWithDefaults

`func NewDfsGatewayGroupUpdateVIPsReqGatewayGroupWithDefaults() *DfsGatewayGroupUpdateVIPsReqGatewayGroup`

NewDfsGatewayGroupUpdateVIPsReqGatewayGroupWithDefaults instantiates a new DfsGatewayGroupUpdateVIPsReqGatewayGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsGatewayZoneId

`func (o *DfsGatewayGroupUpdateVIPsReqGatewayGroup) GetDfsGatewayZoneId() int64`

GetDfsGatewayZoneId returns the DfsGatewayZoneId field if non-nil, zero value otherwise.

### GetDfsGatewayZoneIdOk

`func (o *DfsGatewayGroupUpdateVIPsReqGatewayGroup) GetDfsGatewayZoneIdOk() (*int64, bool)`

GetDfsGatewayZoneIdOk returns a tuple with the DfsGatewayZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayZoneId

`func (o *DfsGatewayGroupUpdateVIPsReqGatewayGroup) SetDfsGatewayZoneId(v int64)`

SetDfsGatewayZoneId sets DfsGatewayZoneId field to given value.

### HasDfsGatewayZoneId

`func (o *DfsGatewayGroupUpdateVIPsReqGatewayGroup) HasDfsGatewayZoneId() bool`

HasDfsGatewayZoneId returns a boolean if a field has been set.

### GetDfsVipGateways

`func (o *DfsGatewayGroupUpdateVIPsReqGatewayGroup) GetDfsVipGateways() []string`

GetDfsVipGateways returns the DfsVipGateways field if non-nil, zero value otherwise.

### GetDfsVipGatewaysOk

`func (o *DfsGatewayGroupUpdateVIPsReqGatewayGroup) GetDfsVipGatewaysOk() (*[]string, bool)`

GetDfsVipGatewaysOk returns a tuple with the DfsVipGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsVipGateways

`func (o *DfsGatewayGroupUpdateVIPsReqGatewayGroup) SetDfsVipGateways(v []string)`

SetDfsVipGateways sets DfsVipGateways field to given value.

### HasDfsVipGateways

`func (o *DfsGatewayGroupUpdateVIPsReqGatewayGroup) HasDfsVipGateways() bool`

HasDfsVipGateways returns a boolean if a field has been set.

### GetDfsVips

`func (o *DfsGatewayGroupUpdateVIPsReqGatewayGroup) GetDfsVips() []string`

GetDfsVips returns the DfsVips field if non-nil, zero value otherwise.

### GetDfsVipsOk

`func (o *DfsGatewayGroupUpdateVIPsReqGatewayGroup) GetDfsVipsOk() (*[]string, bool)`

GetDfsVipsOk returns a tuple with the DfsVips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsVips

`func (o *DfsGatewayGroupUpdateVIPsReqGatewayGroup) SetDfsVips(v []string)`

SetDfsVips sets DfsVips field to given value.

### HasDfsVips

`func (o *DfsGatewayGroupUpdateVIPsReqGatewayGroup) HasDfsVips() bool`

HasDfsVips returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


