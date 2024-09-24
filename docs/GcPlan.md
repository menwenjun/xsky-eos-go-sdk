# GcPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **[]int64** |  | 
**Policies** | [**[]GcPolicy**](GcPolicy.md) |  | 

## Methods

### NewGcPlan

`func NewGcPlan(days []int64, policies []GcPolicy, ) *GcPlan`

NewGcPlan instantiates a new GcPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcPlanWithDefaults

`func NewGcPlanWithDefaults() *GcPlan`

NewGcPlanWithDefaults instantiates a new GcPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *GcPlan) GetDays() []int64`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *GcPlan) GetDaysOk() (*[]int64, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *GcPlan) SetDays(v []int64)`

SetDays sets Days field to given value.


### GetPolicies

`func (o *GcPlan) GetPolicies() []GcPolicy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *GcPlan) GetPoliciesOk() (*[]GcPolicy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *GcPlan) SetPolicies(v []GcPolicy)`

SetPolicies sets Policies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


