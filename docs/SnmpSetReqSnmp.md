# SnmpSetReqSnmp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Gateways** | Pointer to [**[]SnmpGatewayReq**](SnmpGatewayReq.md) |  | [optional] 
**Receivers** | Pointer to [**[]SnmpTrapReceiverReq**](SnmpTrapReceiverReq.md) |  | [optional] 

## Methods

### NewSnmpSetReqSnmp

`func NewSnmpSetReqSnmp() *SnmpSetReqSnmp`

NewSnmpSetReqSnmp instantiates a new SnmpSetReqSnmp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpSetReqSnmpWithDefaults

`func NewSnmpSetReqSnmpWithDefaults() *SnmpSetReqSnmp`

NewSnmpSetReqSnmpWithDefaults instantiates a new SnmpSetReqSnmp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SnmpSetReqSnmp) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SnmpSetReqSnmp) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SnmpSetReqSnmp) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SnmpSetReqSnmp) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGateways

`func (o *SnmpSetReqSnmp) GetGateways() []SnmpGatewayReq`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *SnmpSetReqSnmp) GetGatewaysOk() (*[]SnmpGatewayReq, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *SnmpSetReqSnmp) SetGateways(v []SnmpGatewayReq)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *SnmpSetReqSnmp) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetReceivers

`func (o *SnmpSetReqSnmp) GetReceivers() []SnmpTrapReceiverReq`

GetReceivers returns the Receivers field if non-nil, zero value otherwise.

### GetReceiversOk

`func (o *SnmpSetReqSnmp) GetReceiversOk() (*[]SnmpTrapReceiverReq, bool)`

GetReceiversOk returns a tuple with the Receivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivers

`func (o *SnmpSetReqSnmp) SetReceivers(v []SnmpTrapReceiverReq)`

SetReceivers sets Receivers field to given value.

### HasReceivers

`func (o *SnmpSetReqSnmp) HasReceivers() bool`

HasReceivers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


