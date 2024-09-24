# AlertRuleGroupDetailResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertRuleGroup** | [**AlertRuleGroup**](AlertRuleGroup.md) |  | 
**AlertRules** | [**[]AlertRule**](AlertRule.md) | alert rule group members | 

## Methods

### NewAlertRuleGroupDetailResp

`func NewAlertRuleGroupDetailResp(alertRuleGroup AlertRuleGroup, alertRules []AlertRule, ) *AlertRuleGroupDetailResp`

NewAlertRuleGroupDetailResp instantiates a new AlertRuleGroupDetailResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleGroupDetailRespWithDefaults

`func NewAlertRuleGroupDetailRespWithDefaults() *AlertRuleGroupDetailResp`

NewAlertRuleGroupDetailRespWithDefaults instantiates a new AlertRuleGroupDetailResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertRuleGroup

`func (o *AlertRuleGroupDetailResp) GetAlertRuleGroup() AlertRuleGroup`

GetAlertRuleGroup returns the AlertRuleGroup field if non-nil, zero value otherwise.

### GetAlertRuleGroupOk

`func (o *AlertRuleGroupDetailResp) GetAlertRuleGroupOk() (*AlertRuleGroup, bool)`

GetAlertRuleGroupOk returns a tuple with the AlertRuleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRuleGroup

`func (o *AlertRuleGroupDetailResp) SetAlertRuleGroup(v AlertRuleGroup)`

SetAlertRuleGroup sets AlertRuleGroup field to given value.


### GetAlertRules

`func (o *AlertRuleGroupDetailResp) GetAlertRules() []AlertRule`

GetAlertRules returns the AlertRules field if non-nil, zero value otherwise.

### GetAlertRulesOk

`func (o *AlertRuleGroupDetailResp) GetAlertRulesOk() (*[]AlertRule, bool)`

GetAlertRulesOk returns a tuple with the AlertRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRules

`func (o *AlertRuleGroupDetailResp) SetAlertRules(v []AlertRule)`

SetAlertRules sets AlertRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


