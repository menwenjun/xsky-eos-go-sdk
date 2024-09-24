# AlertStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**AlertContacts** | Pointer to [**[]AlertContactNestview**](AlertContactNestview.md) |  | [optional] 
**AlertRules** | Pointer to [**[]AlertRuleNestview**](AlertRuleNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NotifyMethods** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAlertStrategy

`func NewAlertStrategy() *AlertStrategy`

NewAlertStrategy instantiates a new AlertStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertStrategyWithDefaults

`func NewAlertStrategyWithDefaults() *AlertStrategy`

NewAlertStrategyWithDefaults instantiates a new AlertStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *AlertStrategy) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *AlertStrategy) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *AlertStrategy) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *AlertStrategy) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAlertContacts

`func (o *AlertStrategy) GetAlertContacts() []AlertContactNestview`

GetAlertContacts returns the AlertContacts field if non-nil, zero value otherwise.

### GetAlertContactsOk

`func (o *AlertStrategy) GetAlertContactsOk() (*[]AlertContactNestview, bool)`

GetAlertContactsOk returns a tuple with the AlertContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertContacts

`func (o *AlertStrategy) SetAlertContacts(v []AlertContactNestview)`

SetAlertContacts sets AlertContacts field to given value.

### HasAlertContacts

`func (o *AlertStrategy) HasAlertContacts() bool`

HasAlertContacts returns a boolean if a field has been set.

### GetAlertRules

`func (o *AlertStrategy) GetAlertRules() []AlertRuleNestview`

GetAlertRules returns the AlertRules field if non-nil, zero value otherwise.

### GetAlertRulesOk

`func (o *AlertStrategy) GetAlertRulesOk() (*[]AlertRuleNestview, bool)`

GetAlertRulesOk returns a tuple with the AlertRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRules

`func (o *AlertStrategy) SetAlertRules(v []AlertRuleNestview)`

SetAlertRules sets AlertRules field to given value.

### HasAlertRules

`func (o *AlertStrategy) HasAlertRules() bool`

HasAlertRules returns a boolean if a field has been set.

### GetCreate

`func (o *AlertStrategy) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *AlertStrategy) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *AlertStrategy) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *AlertStrategy) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *AlertStrategy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertStrategy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertStrategy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AlertStrategy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnable

`func (o *AlertStrategy) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AlertStrategy) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AlertStrategy) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *AlertStrategy) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetId

`func (o *AlertStrategy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertStrategy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertStrategy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AlertStrategy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AlertStrategy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertStrategy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertStrategy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlertStrategy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyMethods

`func (o *AlertStrategy) GetNotifyMethods() string`

GetNotifyMethods returns the NotifyMethods field if non-nil, zero value otherwise.

### GetNotifyMethodsOk

`func (o *AlertStrategy) GetNotifyMethodsOk() (*string, bool)`

GetNotifyMethodsOk returns a tuple with the NotifyMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMethods

`func (o *AlertStrategy) SetNotifyMethods(v string)`

SetNotifyMethods sets NotifyMethods field to given value.

### HasNotifyMethods

`func (o *AlertStrategy) HasNotifyMethods() bool`

HasNotifyMethods returns a boolean if a field has been set.

### GetStatus

`func (o *AlertStrategy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertStrategy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertStrategy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertStrategy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *AlertStrategy) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *AlertStrategy) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *AlertStrategy) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *AlertStrategy) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


