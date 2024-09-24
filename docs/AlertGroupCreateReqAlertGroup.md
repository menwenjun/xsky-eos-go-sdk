# AlertGroupCreateReqAlertGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertRuleIds** | **[]int64** |  | 
**EmailGroupIds** | **[]int64** |  | 
**Name** | **string** |  | 

## Methods

### NewAlertGroupCreateReqAlertGroup

`func NewAlertGroupCreateReqAlertGroup(alertRuleIds []int64, emailGroupIds []int64, name string, ) *AlertGroupCreateReqAlertGroup`

NewAlertGroupCreateReqAlertGroup instantiates a new AlertGroupCreateReqAlertGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertGroupCreateReqAlertGroupWithDefaults

`func NewAlertGroupCreateReqAlertGroupWithDefaults() *AlertGroupCreateReqAlertGroup`

NewAlertGroupCreateReqAlertGroupWithDefaults instantiates a new AlertGroupCreateReqAlertGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertRuleIds

`func (o *AlertGroupCreateReqAlertGroup) GetAlertRuleIds() []int64`

GetAlertRuleIds returns the AlertRuleIds field if non-nil, zero value otherwise.

### GetAlertRuleIdsOk

`func (o *AlertGroupCreateReqAlertGroup) GetAlertRuleIdsOk() (*[]int64, bool)`

GetAlertRuleIdsOk returns a tuple with the AlertRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRuleIds

`func (o *AlertGroupCreateReqAlertGroup) SetAlertRuleIds(v []int64)`

SetAlertRuleIds sets AlertRuleIds field to given value.


### GetEmailGroupIds

`func (o *AlertGroupCreateReqAlertGroup) GetEmailGroupIds() []int64`

GetEmailGroupIds returns the EmailGroupIds field if non-nil, zero value otherwise.

### GetEmailGroupIdsOk

`func (o *AlertGroupCreateReqAlertGroup) GetEmailGroupIdsOk() (*[]int64, bool)`

GetEmailGroupIdsOk returns a tuple with the EmailGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailGroupIds

`func (o *AlertGroupCreateReqAlertGroup) SetEmailGroupIds(v []int64)`

SetEmailGroupIds sets EmailGroupIds field to given value.


### GetName

`func (o *AlertGroupCreateReqAlertGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertGroupCreateReqAlertGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertGroupCreateReqAlertGroup) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


