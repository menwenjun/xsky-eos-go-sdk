# AlertRuleGroupDetailUpdateReqGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoResolve** | Pointer to **bool** | auto resolved or not | [optional] 
**Enabled** | Pointer to **bool** | enable the alert rule group or not | [optional] 
**TriggerTimeSeconds** | Pointer to **int64** | trigger time in seconds | [optional] 

## Methods

### NewAlertRuleGroupDetailUpdateReqGroup

`func NewAlertRuleGroupDetailUpdateReqGroup() *AlertRuleGroupDetailUpdateReqGroup`

NewAlertRuleGroupDetailUpdateReqGroup instantiates a new AlertRuleGroupDetailUpdateReqGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleGroupDetailUpdateReqGroupWithDefaults

`func NewAlertRuleGroupDetailUpdateReqGroupWithDefaults() *AlertRuleGroupDetailUpdateReqGroup`

NewAlertRuleGroupDetailUpdateReqGroupWithDefaults instantiates a new AlertRuleGroupDetailUpdateReqGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoResolve

`func (o *AlertRuleGroupDetailUpdateReqGroup) GetAutoResolve() bool`

GetAutoResolve returns the AutoResolve field if non-nil, zero value otherwise.

### GetAutoResolveOk

`func (o *AlertRuleGroupDetailUpdateReqGroup) GetAutoResolveOk() (*bool, bool)`

GetAutoResolveOk returns a tuple with the AutoResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoResolve

`func (o *AlertRuleGroupDetailUpdateReqGroup) SetAutoResolve(v bool)`

SetAutoResolve sets AutoResolve field to given value.

### HasAutoResolve

`func (o *AlertRuleGroupDetailUpdateReqGroup) HasAutoResolve() bool`

HasAutoResolve returns a boolean if a field has been set.

### GetEnabled

`func (o *AlertRuleGroupDetailUpdateReqGroup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertRuleGroupDetailUpdateReqGroup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertRuleGroupDetailUpdateReqGroup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AlertRuleGroupDetailUpdateReqGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTriggerTimeSeconds

`func (o *AlertRuleGroupDetailUpdateReqGroup) GetTriggerTimeSeconds() int64`

GetTriggerTimeSeconds returns the TriggerTimeSeconds field if non-nil, zero value otherwise.

### GetTriggerTimeSecondsOk

`func (o *AlertRuleGroupDetailUpdateReqGroup) GetTriggerTimeSecondsOk() (*int64, bool)`

GetTriggerTimeSecondsOk returns a tuple with the TriggerTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTimeSeconds

`func (o *AlertRuleGroupDetailUpdateReqGroup) SetTriggerTimeSeconds(v int64)`

SetTriggerTimeSeconds sets TriggerTimeSeconds field to given value.

### HasTriggerTimeSeconds

`func (o *AlertRuleGroupDetailUpdateReqGroup) HasTriggerTimeSeconds() bool`

HasTriggerTimeSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


