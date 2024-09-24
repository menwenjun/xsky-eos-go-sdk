# DfsGatewayGroupAddGatewaysReqGatewayGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsGatewayZoneId** | Pointer to **int64** | dfs gateway zone id | [optional] 
**DfsGateways** | [**[]DfsGatewayReq**](DfsGatewayReq.md) | dfs gateways list | 

## Methods

### NewDfsGatewayGroupAddGatewaysReqGatewayGroup

`func NewDfsGatewayGroupAddGatewaysReqGatewayGroup(dfsGateways []DfsGatewayReq, ) *DfsGatewayGroupAddGatewaysReqGatewayGroup`

NewDfsGatewayGroupAddGatewaysReqGatewayGroup instantiates a new DfsGatewayGroupAddGatewaysReqGatewayGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayGroupAddGatewaysReqGatewayGroupWithDefaults

`func NewDfsGatewayGroupAddGatewaysReqGatewayGroupWithDefaults() *DfsGatewayGroupAddGatewaysReqGatewayGroup`

NewDfsGatewayGroupAddGatewaysReqGatewayGroupWithDefaults instantiates a new DfsGatewayGroupAddGatewaysReqGatewayGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsGatewayZoneId

`func (o *DfsGatewayGroupAddGatewaysReqGatewayGroup) GetDfsGatewayZoneId() int64`

GetDfsGatewayZoneId returns the DfsGatewayZoneId field if non-nil, zero value otherwise.

### GetDfsGatewayZoneIdOk

`func (o *DfsGatewayGroupAddGatewaysReqGatewayGroup) GetDfsGatewayZoneIdOk() (*int64, bool)`

GetDfsGatewayZoneIdOk returns a tuple with the DfsGatewayZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayZoneId

`func (o *DfsGatewayGroupAddGatewaysReqGatewayGroup) SetDfsGatewayZoneId(v int64)`

SetDfsGatewayZoneId sets DfsGatewayZoneId field to given value.

### HasDfsGatewayZoneId

`func (o *DfsGatewayGroupAddGatewaysReqGatewayGroup) HasDfsGatewayZoneId() bool`

HasDfsGatewayZoneId returns a boolean if a field has been set.

### GetDfsGateways

`func (o *DfsGatewayGroupAddGatewaysReqGatewayGroup) GetDfsGateways() []DfsGatewayReq`

GetDfsGateways returns the DfsGateways field if non-nil, zero value otherwise.

### GetDfsGatewaysOk

`func (o *DfsGatewayGroupAddGatewaysReqGatewayGroup) GetDfsGatewaysOk() (*[]DfsGatewayReq, bool)`

GetDfsGatewaysOk returns a tuple with the DfsGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGateways

`func (o *DfsGatewayGroupAddGatewaysReqGatewayGroup) SetDfsGateways(v []DfsGatewayReq)`

SetDfsGateways sets DfsGateways field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


