# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**ErrorMsg** | Pointer to **string** |  | [optional] 
**Execute** | Pointer to **time.Time** |  | [optional] 
**Finish** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**HostRole** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Priority** | Pointer to **int64** |  | [optional] 
**Retried** | Pointer to **bool** |  | [optional] 
**Retry** | Pointer to **int64** |  | [optional] 
**RetryType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *Task) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Task) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Task) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Task) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetData

`func (o *Task) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Task) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Task) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Task) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrorMsg

`func (o *Task) GetErrorMsg() string`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *Task) GetErrorMsgOk() (*string, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *Task) SetErrorMsg(v string)`

SetErrorMsg sets ErrorMsg field to given value.

### HasErrorMsg

`func (o *Task) HasErrorMsg() bool`

HasErrorMsg returns a boolean if a field has been set.

### GetExecute

`func (o *Task) GetExecute() time.Time`

GetExecute returns the Execute field if non-nil, zero value otherwise.

### GetExecuteOk

`func (o *Task) GetExecuteOk() (*time.Time, bool)`

GetExecuteOk returns a tuple with the Execute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecute

`func (o *Task) SetExecute(v time.Time)`

SetExecute sets Execute field to given value.

### HasExecute

`func (o *Task) HasExecute() bool`

HasExecute returns a boolean if a field has been set.

### GetFinish

`func (o *Task) GetFinish() time.Time`

GetFinish returns the Finish field if non-nil, zero value otherwise.

### GetFinishOk

`func (o *Task) GetFinishOk() (*time.Time, bool)`

GetFinishOk returns a tuple with the Finish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinish

`func (o *Task) SetFinish(v time.Time)`

SetFinish sets Finish field to given value.

### HasFinish

`func (o *Task) HasFinish() bool`

HasFinish returns a boolean if a field has been set.

### GetHost

`func (o *Task) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Task) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Task) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *Task) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHostRole

`func (o *Task) GetHostRole() string`

GetHostRole returns the HostRole field if non-nil, zero value otherwise.

### GetHostRoleOk

`func (o *Task) GetHostRoleOk() (*string, bool)`

GetHostRoleOk returns a tuple with the HostRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostRole

`func (o *Task) SetHostRole(v string)`

SetHostRole sets HostRole field to given value.

### HasHostRole

`func (o *Task) HasHostRole() bool`

HasHostRole returns a boolean if a field has been set.

### GetId

`func (o *Task) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Task) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Task) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *Task) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Task) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Task) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Task) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRetried

`func (o *Task) GetRetried() bool`

GetRetried returns the Retried field if non-nil, zero value otherwise.

### GetRetriedOk

`func (o *Task) GetRetriedOk() (*bool, bool)`

GetRetriedOk returns a tuple with the Retried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetried

`func (o *Task) SetRetried(v bool)`

SetRetried sets Retried field to given value.

### HasRetried

`func (o *Task) HasRetried() bool`

HasRetried returns a boolean if a field has been set.

### GetRetry

`func (o *Task) GetRetry() int64`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *Task) GetRetryOk() (*int64, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *Task) SetRetry(v int64)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *Task) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetRetryType

`func (o *Task) GetRetryType() string`

GetRetryType returns the RetryType field if non-nil, zero value otherwise.

### GetRetryTypeOk

`func (o *Task) GetRetryTypeOk() (*string, bool)`

GetRetryTypeOk returns a tuple with the RetryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryType

`func (o *Task) SetRetryType(v string)`

SetRetryType sets RetryType field to given value.

### HasRetryType

`func (o *Task) HasRetryType() bool`

HasRetryType returns a boolean if a field has been set.

### GetStatus

`func (o *Task) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Task) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Task) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Task) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Task) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Task) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Task) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Task) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *Task) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Task) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Task) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Task) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


