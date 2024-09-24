# AlertRuleUpdateReqRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | enable the alert rule or not | [optional] 
**TriggerPeriod** | Pointer to **int64** | trigger period of alert rule | [optional] 
**TriggerValue** | Pointer to **string** | trigger value of alert rule | [optional] 

## Methods

### NewAlertRuleUpdateReqRule

`func NewAlertRuleUpdateReqRule() *AlertRuleUpdateReqRule`

NewAlertRuleUpdateReqRule instantiates a new AlertRuleUpdateReqRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleUpdateReqRuleWithDefaults

`func NewAlertRuleUpdateReqRuleWithDefaults() *AlertRuleUpdateReqRule`

NewAlertRuleUpdateReqRuleWithDefaults instantiates a new AlertRuleUpdateReqRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AlertRuleUpdateReqRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertRuleUpdateReqRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertRuleUpdateReqRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AlertRuleUpdateReqRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTriggerPeriod

`func (o *AlertRuleUpdateReqRule) GetTriggerPeriod() int64`

GetTriggerPeriod returns the TriggerPeriod field if non-nil, zero value otherwise.

### GetTriggerPeriodOk

`func (o *AlertRuleUpdateReqRule) GetTriggerPeriodOk() (*int64, bool)`

GetTriggerPeriodOk returns a tuple with the TriggerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPeriod

`func (o *AlertRuleUpdateReqRule) SetTriggerPeriod(v int64)`

SetTriggerPeriod sets TriggerPeriod field to given value.

### HasTriggerPeriod

`func (o *AlertRuleUpdateReqRule) HasTriggerPeriod() bool`

HasTriggerPeriod returns a boolean if a field has been set.

### GetTriggerValue

`func (o *AlertRuleUpdateReqRule) GetTriggerValue() string`

GetTriggerValue returns the TriggerValue field if non-nil, zero value otherwise.

### GetTriggerValueOk

`func (o *AlertRuleUpdateReqRule) GetTriggerValueOk() (*string, bool)`

GetTriggerValueOk returns a tuple with the TriggerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerValue

`func (o *AlertRuleUpdateReqRule) SetTriggerValue(v string)`

SetTriggerValue sets TriggerValue field to given value.

### HasTriggerValue

`func (o *AlertRuleUpdateReqRule) HasTriggerValue() bool`

HasTriggerValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


