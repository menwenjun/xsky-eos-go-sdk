# SnmpInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | enable snmp or not | 
**Gateways** | [**[]SnmpGateway**](SnmpGateway.md) | snmp gateways | 
**Receivers** | [**[]TrapReceiver**](TrapReceiver.md) | snmp receivers | 

## Methods

### NewSnmpInfo

`func NewSnmpInfo(enabled bool, gateways []SnmpGateway, receivers []TrapReceiver, ) *SnmpInfo`

NewSnmpInfo instantiates a new SnmpInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpInfoWithDefaults

`func NewSnmpInfoWithDefaults() *SnmpInfo`

NewSnmpInfoWithDefaults instantiates a new SnmpInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SnmpInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SnmpInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SnmpInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetGateways

`func (o *SnmpInfo) GetGateways() []SnmpGateway`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *SnmpInfo) GetGatewaysOk() (*[]SnmpGateway, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *SnmpInfo) SetGateways(v []SnmpGateway)`

SetGateways sets Gateways field to given value.


### GetReceivers

`func (o *SnmpInfo) GetReceivers() []TrapReceiver`

GetReceivers returns the Receivers field if non-nil, zero value otherwise.

### GetReceiversOk

`func (o *SnmpInfo) GetReceiversOk() (*[]TrapReceiver, bool)`

GetReceiversOk returns a tuple with the Receivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivers

`func (o *SnmpInfo) SetReceivers(v []TrapReceiver)`

SetReceivers sets Receivers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


