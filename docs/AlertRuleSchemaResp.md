# AlertRuleSchemaResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleSchemaMap** | [**map[string]AlertRuleSchema**](AlertRuleSchema.md) | rule schemas | 
**TriggerModes** | **[]string** | trigger modes | 

## Methods

### NewAlertRuleSchemaResp

`func NewAlertRuleSchemaResp(ruleSchemaMap map[string]AlertRuleSchema, triggerModes []string, ) *AlertRuleSchemaResp`

NewAlertRuleSchemaResp instantiates a new AlertRuleSchemaResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleSchemaRespWithDefaults

`func NewAlertRuleSchemaRespWithDefaults() *AlertRuleSchemaResp`

NewAlertRuleSchemaRespWithDefaults instantiates a new AlertRuleSchemaResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleSchemaMap

`func (o *AlertRuleSchemaResp) GetRuleSchemaMap() map[string]AlertRuleSchema`

GetRuleSchemaMap returns the RuleSchemaMap field if non-nil, zero value otherwise.

### GetRuleSchemaMapOk

`func (o *AlertRuleSchemaResp) GetRuleSchemaMapOk() (*map[string]AlertRuleSchema, bool)`

GetRuleSchemaMapOk returns a tuple with the RuleSchemaMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSchemaMap

`func (o *AlertRuleSchemaResp) SetRuleSchemaMap(v map[string]AlertRuleSchema)`

SetRuleSchemaMap sets RuleSchemaMap field to given value.


### GetTriggerModes

`func (o *AlertRuleSchemaResp) GetTriggerModes() []string`

GetTriggerModes returns the TriggerModes field if non-nil, zero value otherwise.

### GetTriggerModesOk

`func (o *AlertRuleSchemaResp) GetTriggerModesOk() (*[]string, bool)`

GetTriggerModesOk returns a tuple with the TriggerModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerModes

`func (o *AlertRuleSchemaResp) SetTriggerModes(v []string)`

SetTriggerModes sets TriggerModes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


