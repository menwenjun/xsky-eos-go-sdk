# MatchingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | Pointer to **string** |  | [optional] 
**ClassName** | Pointer to **string** |  | [optional] 
**MatchingCondition** | Pointer to [**MatchingRuleMatchCondition**](MatchingRuleMatchCondition.md) |  | [optional] 
**Prior** | Pointer to **int32** |  | [optional] 

## Methods

### NewMatchingRule

`func NewMatchingRule() *MatchingRule`

NewMatchingRule instantiates a new MatchingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchingRuleWithDefaults

`func NewMatchingRuleWithDefaults() *MatchingRule`

NewMatchingRuleWithDefaults instantiates a new MatchingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MatchingRule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MatchingRule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MatchingRule) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *MatchingRule) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetClassName

`func (o *MatchingRule) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *MatchingRule) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *MatchingRule) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *MatchingRule) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetMatchingCondition

`func (o *MatchingRule) GetMatchingCondition() MatchingRuleMatchCondition`

GetMatchingCondition returns the MatchingCondition field if non-nil, zero value otherwise.

### GetMatchingConditionOk

`func (o *MatchingRule) GetMatchingConditionOk() (*MatchingRuleMatchCondition, bool)`

GetMatchingConditionOk returns a tuple with the MatchingCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingCondition

`func (o *MatchingRule) SetMatchingCondition(v MatchingRuleMatchCondition)`

SetMatchingCondition sets MatchingCondition field to given value.

### HasMatchingCondition

`func (o *MatchingRule) HasMatchingCondition() bool`

HasMatchingCondition returns a boolean if a field has been set.

### GetPrior

`func (o *MatchingRule) GetPrior() int32`

GetPrior returns the Prior field if non-nil, zero value otherwise.

### GetPriorOk

`func (o *MatchingRule) GetPriorOk() (*int32, bool)`

GetPriorOk returns a tuple with the Prior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrior

`func (o *MatchingRule) SetPrior(v int32)`

SetPrior sets Prior field to given value.

### HasPrior

`func (o *MatchingRule) HasPrior() bool`

HasPrior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


