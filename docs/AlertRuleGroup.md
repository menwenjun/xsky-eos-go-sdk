# AlertRuleGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoResolve** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Levels** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**TriggerTimeSeconds** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAlertRuleGroup

`func NewAlertRuleGroup() *AlertRuleGroup`

NewAlertRuleGroup instantiates a new AlertRuleGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleGroupWithDefaults

`func NewAlertRuleGroupWithDefaults() *AlertRuleGroup`

NewAlertRuleGroupWithDefaults instantiates a new AlertRuleGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoResolve

`func (o *AlertRuleGroup) GetAutoResolve() bool`

GetAutoResolve returns the AutoResolve field if non-nil, zero value otherwise.

### GetAutoResolveOk

`func (o *AlertRuleGroup) GetAutoResolveOk() (*bool, bool)`

GetAutoResolveOk returns a tuple with the AutoResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoResolve

`func (o *AlertRuleGroup) SetAutoResolve(v bool)`

SetAutoResolve sets AutoResolve field to given value.

### HasAutoResolve

`func (o *AlertRuleGroup) HasAutoResolve() bool`

HasAutoResolve returns a boolean if a field has been set.

### GetCreate

`func (o *AlertRuleGroup) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *AlertRuleGroup) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *AlertRuleGroup) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *AlertRuleGroup) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetEnabled

`func (o *AlertRuleGroup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertRuleGroup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertRuleGroup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AlertRuleGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *AlertRuleGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertRuleGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertRuleGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AlertRuleGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevels

`func (o *AlertRuleGroup) GetLevels() []string`

GetLevels returns the Levels field if non-nil, zero value otherwise.

### GetLevelsOk

`func (o *AlertRuleGroup) GetLevelsOk() (*[]string, bool)`

GetLevelsOk returns a tuple with the Levels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevels

`func (o *AlertRuleGroup) SetLevels(v []string)`

SetLevels sets Levels field to given value.

### HasLevels

`func (o *AlertRuleGroup) HasLevels() bool`

HasLevels returns a boolean if a field has been set.

### GetName

`func (o *AlertRuleGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertRuleGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertRuleGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlertRuleGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceType

`func (o *AlertRuleGroup) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AlertRuleGroup) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AlertRuleGroup) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AlertRuleGroup) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTriggerTimeSeconds

`func (o *AlertRuleGroup) GetTriggerTimeSeconds() int64`

GetTriggerTimeSeconds returns the TriggerTimeSeconds field if non-nil, zero value otherwise.

### GetTriggerTimeSecondsOk

`func (o *AlertRuleGroup) GetTriggerTimeSecondsOk() (*int64, bool)`

GetTriggerTimeSecondsOk returns a tuple with the TriggerTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTimeSeconds

`func (o *AlertRuleGroup) SetTriggerTimeSeconds(v int64)`

SetTriggerTimeSeconds sets TriggerTimeSeconds field to given value.

### HasTriggerTimeSeconds

`func (o *AlertRuleGroup) HasTriggerTimeSeconds() bool`

HasTriggerTimeSeconds returns a boolean if a field has been set.

### GetUpdate

`func (o *AlertRuleGroup) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *AlertRuleGroup) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *AlertRuleGroup) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *AlertRuleGroup) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


