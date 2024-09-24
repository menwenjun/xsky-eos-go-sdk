# RuleMatchingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchingRules** | Pointer to [**[]MatchingRule**](MatchingRule.md) |  | [optional] 
**MismatchingAction** | Pointer to **string** |  | [optional] 

## Methods

### NewRuleMatchingInfo

`func NewRuleMatchingInfo() *RuleMatchingInfo`

NewRuleMatchingInfo instantiates a new RuleMatchingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleMatchingInfoWithDefaults

`func NewRuleMatchingInfoWithDefaults() *RuleMatchingInfo`

NewRuleMatchingInfoWithDefaults instantiates a new RuleMatchingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchingRules

`func (o *RuleMatchingInfo) GetMatchingRules() []MatchingRule`

GetMatchingRules returns the MatchingRules field if non-nil, zero value otherwise.

### GetMatchingRulesOk

`func (o *RuleMatchingInfo) GetMatchingRulesOk() (*[]MatchingRule, bool)`

GetMatchingRulesOk returns a tuple with the MatchingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingRules

`func (o *RuleMatchingInfo) SetMatchingRules(v []MatchingRule)`

SetMatchingRules sets MatchingRules field to given value.

### HasMatchingRules

`func (o *RuleMatchingInfo) HasMatchingRules() bool`

HasMatchingRules returns a boolean if a field has been set.

### GetMismatchingAction

`func (o *RuleMatchingInfo) GetMismatchingAction() string`

GetMismatchingAction returns the MismatchingAction field if non-nil, zero value otherwise.

### GetMismatchingActionOk

`func (o *RuleMatchingInfo) GetMismatchingActionOk() (*string, bool)`

GetMismatchingActionOk returns a tuple with the MismatchingAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchingAction

`func (o *RuleMatchingInfo) SetMismatchingAction(v string)`

SetMismatchingAction sets MismatchingAction field to given value.

### HasMismatchingAction

`func (o *RuleMatchingInfo) HasMismatchingAction() bool`

HasMismatchingAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


