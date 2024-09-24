# DfsGatewayZoneCreateReqGatewayZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsGatewayGroupId** | **int64** | id of dfs gateway group | 
**DfsGateways** | [**[]DfsGatewayReq**](DfsGatewayReq.md) | dfs gateways list | 
**DfsVipGateways** | Pointer to **[]string** | vip gateway list | [optional] 
**DfsVips** | **[]string** | vip list | 
**Name** | **string** | gateway zone name | 

## Methods

### NewDfsGatewayZoneCreateReqGatewayZone

`func NewDfsGatewayZoneCreateReqGatewayZone(dfsGatewayGroupId int64, dfsGateways []DfsGatewayReq, dfsVips []string, name string, ) *DfsGatewayZoneCreateReqGatewayZone`

NewDfsGatewayZoneCreateReqGatewayZone instantiates a new DfsGatewayZoneCreateReqGatewayZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayZoneCreateReqGatewayZoneWithDefaults

`func NewDfsGatewayZoneCreateReqGatewayZoneWithDefaults() *DfsGatewayZoneCreateReqGatewayZone`

NewDfsGatewayZoneCreateReqGatewayZoneWithDefaults instantiates a new DfsGatewayZoneCreateReqGatewayZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsGatewayGroupId

`func (o *DfsGatewayZoneCreateReqGatewayZone) GetDfsGatewayGroupId() int64`

GetDfsGatewayGroupId returns the DfsGatewayGroupId field if non-nil, zero value otherwise.

### GetDfsGatewayGroupIdOk

`func (o *DfsGatewayZoneCreateReqGatewayZone) GetDfsGatewayGroupIdOk() (*int64, bool)`

GetDfsGatewayGroupIdOk returns a tuple with the DfsGatewayGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroupId

`func (o *DfsGatewayZoneCreateReqGatewayZone) SetDfsGatewayGroupId(v int64)`

SetDfsGatewayGroupId sets DfsGatewayGroupId field to given value.


### GetDfsGateways

`func (o *DfsGatewayZoneCreateReqGatewayZone) GetDfsGateways() []DfsGatewayReq`

GetDfsGateways returns the DfsGateways field if non-nil, zero value otherwise.

### GetDfsGatewaysOk

`func (o *DfsGatewayZoneCreateReqGatewayZone) GetDfsGatewaysOk() (*[]DfsGatewayReq, bool)`

GetDfsGatewaysOk returns a tuple with the DfsGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGateways

`func (o *DfsGatewayZoneCreateReqGatewayZone) SetDfsGateways(v []DfsGatewayReq)`

SetDfsGateways sets DfsGateways field to given value.


### GetDfsVipGateways

`func (o *DfsGatewayZoneCreateReqGatewayZone) GetDfsVipGateways() []string`

GetDfsVipGateways returns the DfsVipGateways field if non-nil, zero value otherwise.

### GetDfsVipGatewaysOk

`func (o *DfsGatewayZoneCreateReqGatewayZone) GetDfsVipGatewaysOk() (*[]string, bool)`

GetDfsVipGatewaysOk returns a tuple with the DfsVipGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsVipGateways

`func (o *DfsGatewayZoneCreateReqGatewayZone) SetDfsVipGateways(v []string)`

SetDfsVipGateways sets DfsVipGateways field to given value.

### HasDfsVipGateways

`func (o *DfsGatewayZoneCreateReqGatewayZone) HasDfsVipGateways() bool`

HasDfsVipGateways returns a boolean if a field has been set.

### GetDfsVips

`func (o *DfsGatewayZoneCreateReqGatewayZone) GetDfsVips() []string`

GetDfsVips returns the DfsVips field if non-nil, zero value otherwise.

### GetDfsVipsOk

`func (o *DfsGatewayZoneCreateReqGatewayZone) GetDfsVipsOk() (*[]string, bool)`

GetDfsVipsOk returns a tuple with the DfsVips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsVips

`func (o *DfsGatewayZoneCreateReqGatewayZone) SetDfsVips(v []string)`

SetDfsVips sets DfsVips field to given value.


### GetName

`func (o *DfsGatewayZoneCreateReqGatewayZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsGatewayZoneCreateReqGatewayZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsGatewayZoneCreateReqGatewayZone) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


