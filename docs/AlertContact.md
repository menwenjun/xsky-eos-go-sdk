# AlertContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**AlertStrategies** | Pointer to [**[]AlertStrategyNestview**](AlertStrategyNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**EmailAddr** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**WechatContactType** | Pointer to **string** |  | [optional] 
**WechatId** | Pointer to **string** |  | [optional] 

## Methods

### NewAlertContact

`func NewAlertContact() *AlertContact`

NewAlertContact instantiates a new AlertContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertContactWithDefaults

`func NewAlertContactWithDefaults() *AlertContact`

NewAlertContactWithDefaults instantiates a new AlertContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *AlertContact) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *AlertContact) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *AlertContact) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *AlertContact) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAlertStrategies

`func (o *AlertContact) GetAlertStrategies() []AlertStrategyNestview`

GetAlertStrategies returns the AlertStrategies field if non-nil, zero value otherwise.

### GetAlertStrategiesOk

`func (o *AlertContact) GetAlertStrategiesOk() (*[]AlertStrategyNestview, bool)`

GetAlertStrategiesOk returns a tuple with the AlertStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertStrategies

`func (o *AlertContact) SetAlertStrategies(v []AlertStrategyNestview)`

SetAlertStrategies sets AlertStrategies field to given value.

### HasAlertStrategies

`func (o *AlertContact) HasAlertStrategies() bool`

HasAlertStrategies returns a boolean if a field has been set.

### GetCreate

`func (o *AlertContact) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *AlertContact) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *AlertContact) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *AlertContact) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetEmailAddr

`func (o *AlertContact) GetEmailAddr() string`

GetEmailAddr returns the EmailAddr field if non-nil, zero value otherwise.

### GetEmailAddrOk

`func (o *AlertContact) GetEmailAddrOk() (*string, bool)`

GetEmailAddrOk returns a tuple with the EmailAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddr

`func (o *AlertContact) SetEmailAddr(v string)`

SetEmailAddr sets EmailAddr field to given value.

### HasEmailAddr

`func (o *AlertContact) HasEmailAddr() bool`

HasEmailAddr returns a boolean if a field has been set.

### GetId

`func (o *AlertContact) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertContact) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertContact) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AlertContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AlertContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertContact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlertContact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *AlertContact) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertContact) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertContact) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertContact) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *AlertContact) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *AlertContact) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *AlertContact) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *AlertContact) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetWechatContactType

`func (o *AlertContact) GetWechatContactType() string`

GetWechatContactType returns the WechatContactType field if non-nil, zero value otherwise.

### GetWechatContactTypeOk

`func (o *AlertContact) GetWechatContactTypeOk() (*string, bool)`

GetWechatContactTypeOk returns a tuple with the WechatContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWechatContactType

`func (o *AlertContact) SetWechatContactType(v string)`

SetWechatContactType sets WechatContactType field to given value.

### HasWechatContactType

`func (o *AlertContact) HasWechatContactType() bool`

HasWechatContactType returns a boolean if a field has been set.

### GetWechatId

`func (o *AlertContact) GetWechatId() string`

GetWechatId returns the WechatId field if non-nil, zero value otherwise.

### GetWechatIdOk

`func (o *AlertContact) GetWechatIdOk() (*string, bool)`

GetWechatIdOk returns a tuple with the WechatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWechatId

`func (o *AlertContact) SetWechatId(v string)`

SetWechatId sets WechatId field to given value.

### HasWechatId

`func (o *AlertContact) HasWechatId() bool`

HasWechatId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


