# MatchingRuleMatchCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CondOperator** | Pointer to **string** |  | [optional] 
**Conditions** | Pointer to [**[]Condition**](Condition.md) |  | [optional] 

## Methods

### NewMatchingRuleMatchCondition

`func NewMatchingRuleMatchCondition() *MatchingRuleMatchCondition`

NewMatchingRuleMatchCondition instantiates a new MatchingRuleMatchCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchingRuleMatchConditionWithDefaults

`func NewMatchingRuleMatchConditionWithDefaults() *MatchingRuleMatchCondition`

NewMatchingRuleMatchConditionWithDefaults instantiates a new MatchingRuleMatchCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondOperator

`func (o *MatchingRuleMatchCondition) GetCondOperator() string`

GetCondOperator returns the CondOperator field if non-nil, zero value otherwise.

### GetCondOperatorOk

`func (o *MatchingRuleMatchCondition) GetCondOperatorOk() (*string, bool)`

GetCondOperatorOk returns a tuple with the CondOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondOperator

`func (o *MatchingRuleMatchCondition) SetCondOperator(v string)`

SetCondOperator sets CondOperator field to given value.

### HasCondOperator

`func (o *MatchingRuleMatchCondition) HasCondOperator() bool`

HasCondOperator returns a boolean if a field has been set.

### GetConditions

`func (o *MatchingRuleMatchCondition) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *MatchingRuleMatchCondition) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *MatchingRuleMatchCondition) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *MatchingRuleMatchCondition) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


