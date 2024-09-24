# AlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlarmId** | Pointer to **string** |  | [optional] 
**BasicType** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**HasResourceBlacklist** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**ResourceTypeDesc** | Pointer to **string** |  | [optional] 
**TriggerMode** | Pointer to **string** |  | [optional] 
**TriggerPeriod** | Pointer to **int64** |  | [optional] 
**TriggerValue** | Pointer to **string** |  | [optional] 
**TriggerValueDesc** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 

## Methods

### NewAlertRule

`func NewAlertRule() *AlertRule`

NewAlertRule instantiates a new AlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleWithDefaults

`func NewAlertRuleWithDefaults() *AlertRule`

NewAlertRuleWithDefaults instantiates a new AlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmId

`func (o *AlertRule) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *AlertRule) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *AlertRule) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.

### HasAlarmId

`func (o *AlertRule) HasAlarmId() bool`

HasAlarmId returns a boolean if a field has been set.

### GetBasicType

`func (o *AlertRule) GetBasicType() string`

GetBasicType returns the BasicType field if non-nil, zero value otherwise.

### GetBasicTypeOk

`func (o *AlertRule) GetBasicTypeOk() (*string, bool)`

GetBasicTypeOk returns a tuple with the BasicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicType

`func (o *AlertRule) SetBasicType(v string)`

SetBasicType sets BasicType field to given value.

### HasBasicType

`func (o *AlertRule) HasBasicType() bool`

HasBasicType returns a boolean if a field has been set.

### GetCreate

`func (o *AlertRule) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *AlertRule) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *AlertRule) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *AlertRule) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *AlertRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AlertRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AlertRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AlertRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroup

`func (o *AlertRule) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AlertRule) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AlertRule) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AlertRule) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHasResourceBlacklist

`func (o *AlertRule) GetHasResourceBlacklist() bool`

GetHasResourceBlacklist returns the HasResourceBlacklist field if non-nil, zero value otherwise.

### GetHasResourceBlacklistOk

`func (o *AlertRule) GetHasResourceBlacklistOk() (*bool, bool)`

GetHasResourceBlacklistOk returns a tuple with the HasResourceBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasResourceBlacklist

`func (o *AlertRule) SetHasResourceBlacklist(v bool)`

SetHasResourceBlacklist sets HasResourceBlacklist field to given value.

### HasHasResourceBlacklist

`func (o *AlertRule) HasHasResourceBlacklist() bool`

HasHasResourceBlacklist returns a boolean if a field has been set.

### GetId

`func (o *AlertRule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertRule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertRule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AlertRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevel

`func (o *AlertRule) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AlertRule) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AlertRule) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AlertRule) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetResourceType

`func (o *AlertRule) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AlertRule) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AlertRule) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AlertRule) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceTypeDesc

`func (o *AlertRule) GetResourceTypeDesc() string`

GetResourceTypeDesc returns the ResourceTypeDesc field if non-nil, zero value otherwise.

### GetResourceTypeDescOk

`func (o *AlertRule) GetResourceTypeDescOk() (*string, bool)`

GetResourceTypeDescOk returns a tuple with the ResourceTypeDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypeDesc

`func (o *AlertRule) SetResourceTypeDesc(v string)`

SetResourceTypeDesc sets ResourceTypeDesc field to given value.

### HasResourceTypeDesc

`func (o *AlertRule) HasResourceTypeDesc() bool`

HasResourceTypeDesc returns a boolean if a field has been set.

### GetTriggerMode

`func (o *AlertRule) GetTriggerMode() string`

GetTriggerMode returns the TriggerMode field if non-nil, zero value otherwise.

### GetTriggerModeOk

`func (o *AlertRule) GetTriggerModeOk() (*string, bool)`

GetTriggerModeOk returns a tuple with the TriggerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMode

`func (o *AlertRule) SetTriggerMode(v string)`

SetTriggerMode sets TriggerMode field to given value.

### HasTriggerMode

`func (o *AlertRule) HasTriggerMode() bool`

HasTriggerMode returns a boolean if a field has been set.

### GetTriggerPeriod

`func (o *AlertRule) GetTriggerPeriod() int64`

GetTriggerPeriod returns the TriggerPeriod field if non-nil, zero value otherwise.

### GetTriggerPeriodOk

`func (o *AlertRule) GetTriggerPeriodOk() (*int64, bool)`

GetTriggerPeriodOk returns a tuple with the TriggerPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPeriod

`func (o *AlertRule) SetTriggerPeriod(v int64)`

SetTriggerPeriod sets TriggerPeriod field to given value.

### HasTriggerPeriod

`func (o *AlertRule) HasTriggerPeriod() bool`

HasTriggerPeriod returns a boolean if a field has been set.

### GetTriggerValue

`func (o *AlertRule) GetTriggerValue() string`

GetTriggerValue returns the TriggerValue field if non-nil, zero value otherwise.

### GetTriggerValueOk

`func (o *AlertRule) GetTriggerValueOk() (*string, bool)`

GetTriggerValueOk returns a tuple with the TriggerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerValue

`func (o *AlertRule) SetTriggerValue(v string)`

SetTriggerValue sets TriggerValue field to given value.

### HasTriggerValue

`func (o *AlertRule) HasTriggerValue() bool`

HasTriggerValue returns a boolean if a field has been set.

### GetTriggerValueDesc

`func (o *AlertRule) GetTriggerValueDesc() string`

GetTriggerValueDesc returns the TriggerValueDesc field if non-nil, zero value otherwise.

### GetTriggerValueDescOk

`func (o *AlertRule) GetTriggerValueDescOk() (*string, bool)`

GetTriggerValueDescOk returns a tuple with the TriggerValueDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerValueDesc

`func (o *AlertRule) SetTriggerValueDesc(v string)`

SetTriggerValueDesc sets TriggerValueDesc field to given value.

### HasTriggerValueDesc

`func (o *AlertRule) HasTriggerValueDesc() bool`

HasTriggerValueDesc returns a boolean if a field has been set.

### GetType

`func (o *AlertRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AlertRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *AlertRule) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *AlertRule) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *AlertRule) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *AlertRule) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetWeight

`func (o *AlertRule) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *AlertRule) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *AlertRule) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *AlertRule) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


