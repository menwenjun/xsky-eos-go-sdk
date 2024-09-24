# AlertRuleGroupDetailUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertRuleGroup** | Pointer to [**AlertRuleGroupDetailUpdateReqGroup**](AlertRuleGroupDetailUpdateReqGroup.md) |  | [optional] 
**AlertRules** | Pointer to [**[]AlertRuleGroupDetailUpdateReqRulesElt**](AlertRuleGroupDetailUpdateReqRulesElt.md) | alert rule group members | [optional] 

## Methods

### NewAlertRuleGroupDetailUpdateReq

`func NewAlertRuleGroupDetailUpdateReq() *AlertRuleGroupDetailUpdateReq`

NewAlertRuleGroupDetailUpdateReq instantiates a new AlertRuleGroupDetailUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleGroupDetailUpdateReqWithDefaults

`func NewAlertRuleGroupDetailUpdateReqWithDefaults() *AlertRuleGroupDetailUpdateReq`

NewAlertRuleGroupDetailUpdateReqWithDefaults instantiates a new AlertRuleGroupDetailUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertRuleGroup

`func (o *AlertRuleGroupDetailUpdateReq) GetAlertRuleGroup() AlertRuleGroupDetailUpdateReqGroup`

GetAlertRuleGroup returns the AlertRuleGroup field if non-nil, zero value otherwise.

### GetAlertRuleGroupOk

`func (o *AlertRuleGroupDetailUpdateReq) GetAlertRuleGroupOk() (*AlertRuleGroupDetailUpdateReqGroup, bool)`

GetAlertRuleGroupOk returns a tuple with the AlertRuleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRuleGroup

`func (o *AlertRuleGroupDetailUpdateReq) SetAlertRuleGroup(v AlertRuleGroupDetailUpdateReqGroup)`

SetAlertRuleGroup sets AlertRuleGroup field to given value.

### HasAlertRuleGroup

`func (o *AlertRuleGroupDetailUpdateReq) HasAlertRuleGroup() bool`

HasAlertRuleGroup returns a boolean if a field has been set.

### GetAlertRules

`func (o *AlertRuleGroupDetailUpdateReq) GetAlertRules() []AlertRuleGroupDetailUpdateReqRulesElt`

GetAlertRules returns the AlertRules field if non-nil, zero value otherwise.

### GetAlertRulesOk

`func (o *AlertRuleGroupDetailUpdateReq) GetAlertRulesOk() (*[]AlertRuleGroupDetailUpdateReqRulesElt, bool)`

GetAlertRulesOk returns a tuple with the AlertRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRules

`func (o *AlertRuleGroupDetailUpdateReq) SetAlertRules(v []AlertRuleGroupDetailUpdateReqRulesElt)`

SetAlertRules sets AlertRules field to given value.

### HasAlertRules

`func (o *AlertRuleGroupDetailUpdateReq) HasAlertRules() bool`

HasAlertRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


