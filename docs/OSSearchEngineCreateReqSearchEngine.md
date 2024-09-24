# OSSearchEngineCreateReqSearchEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayDataSize** | **int64** |  | 
**GatewayFlavorId** | **int64** |  | 
**GatewayHttpPort** | Pointer to **int64** |  | [optional] 
**GatewayTransportPort** | Pointer to **int64** |  | [optional] 
**OsSearchGateways** | [**[]OSSearchGatewayReq**](OSSearchGatewayReq.md) |  | 

## Methods

### NewOSSearchEngineCreateReqSearchEngine

`func NewOSSearchEngineCreateReqSearchEngine(gatewayDataSize int64, gatewayFlavorId int64, osSearchGateways []OSSearchGatewayReq, ) *OSSearchEngineCreateReqSearchEngine`

NewOSSearchEngineCreateReqSearchEngine instantiates a new OSSearchEngineCreateReqSearchEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSSearchEngineCreateReqSearchEngineWithDefaults

`func NewOSSearchEngineCreateReqSearchEngineWithDefaults() *OSSearchEngineCreateReqSearchEngine`

NewOSSearchEngineCreateReqSearchEngineWithDefaults instantiates a new OSSearchEngineCreateReqSearchEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayDataSize

`func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayDataSize() int64`

GetGatewayDataSize returns the GatewayDataSize field if non-nil, zero value otherwise.

### GetGatewayDataSizeOk

`func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayDataSizeOk() (*int64, bool)`

GetGatewayDataSizeOk returns a tuple with the GatewayDataSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayDataSize

`func (o *OSSearchEngineCreateReqSearchEngine) SetGatewayDataSize(v int64)`

SetGatewayDataSize sets GatewayDataSize field to given value.


### GetGatewayFlavorId

`func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayFlavorId() int64`

GetGatewayFlavorId returns the GatewayFlavorId field if non-nil, zero value otherwise.

### GetGatewayFlavorIdOk

`func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayFlavorIdOk() (*int64, bool)`

GetGatewayFlavorIdOk returns a tuple with the GatewayFlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayFlavorId

`func (o *OSSearchEngineCreateReqSearchEngine) SetGatewayFlavorId(v int64)`

SetGatewayFlavorId sets GatewayFlavorId field to given value.


### GetGatewayHttpPort

`func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayHttpPort() int64`

GetGatewayHttpPort returns the GatewayHttpPort field if non-nil, zero value otherwise.

### GetGatewayHttpPortOk

`func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayHttpPortOk() (*int64, bool)`

GetGatewayHttpPortOk returns a tuple with the GatewayHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayHttpPort

`func (o *OSSearchEngineCreateReqSearchEngine) SetGatewayHttpPort(v int64)`

SetGatewayHttpPort sets GatewayHttpPort field to given value.

### HasGatewayHttpPort

`func (o *OSSearchEngineCreateReqSearchEngine) HasGatewayHttpPort() bool`

HasGatewayHttpPort returns a boolean if a field has been set.

### GetGatewayTransportPort

`func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayTransportPort() int64`

GetGatewayTransportPort returns the GatewayTransportPort field if non-nil, zero value otherwise.

### GetGatewayTransportPortOk

`func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayTransportPortOk() (*int64, bool)`

GetGatewayTransportPortOk returns a tuple with the GatewayTransportPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransportPort

`func (o *OSSearchEngineCreateReqSearchEngine) SetGatewayTransportPort(v int64)`

SetGatewayTransportPort sets GatewayTransportPort field to given value.

### HasGatewayTransportPort

`func (o *OSSearchEngineCreateReqSearchEngine) HasGatewayTransportPort() bool`

HasGatewayTransportPort returns a boolean if a field has been set.

### GetOsSearchGateways

`func (o *OSSearchEngineCreateReqSearchEngine) GetOsSearchGateways() []OSSearchGatewayReq`

GetOsSearchGateways returns the OsSearchGateways field if non-nil, zero value otherwise.

### GetOsSearchGatewaysOk

`func (o *OSSearchEngineCreateReqSearchEngine) GetOsSearchGatewaysOk() (*[]OSSearchGatewayReq, bool)`

GetOsSearchGatewaysOk returns a tuple with the OsSearchGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsSearchGateways

`func (o *OSSearchEngineCreateReqSearchEngine) SetOsSearchGateways(v []OSSearchGatewayReq)`

SetOsSearchGateways sets OsSearchGateways field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


