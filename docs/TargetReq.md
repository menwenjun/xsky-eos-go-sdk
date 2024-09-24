# TargetReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayIps** | Pointer to **[]string** |  | [optional] 
**HostId** | Pointer to **int64** |  | [optional] 

## Methods

### NewTargetReq

`func NewTargetReq() *TargetReq`

NewTargetReq instantiates a new TargetReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetReqWithDefaults

`func NewTargetReqWithDefaults() *TargetReq`

NewTargetReqWithDefaults instantiates a new TargetReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayIps

`func (o *TargetReq) GetGatewayIps() []string`

GetGatewayIps returns the GatewayIps field if non-nil, zero value otherwise.

### GetGatewayIpsOk

`func (o *TargetReq) GetGatewayIpsOk() (*[]string, bool)`

GetGatewayIpsOk returns a tuple with the GatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIps

`func (o *TargetReq) SetGatewayIps(v []string)`

SetGatewayIps sets GatewayIps field to given value.

### HasGatewayIps

`func (o *TargetReq) HasGatewayIps() bool`

HasGatewayIps returns a boolean if a field has been set.

### GetHostId

`func (o *TargetReq) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *TargetReq) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *TargetReq) SetHostId(v int64)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *TargetReq) HasHostId() bool`

HasHostId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


