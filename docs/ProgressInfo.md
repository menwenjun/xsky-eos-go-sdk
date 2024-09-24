# ProgressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**CurrentProgress** | Pointer to **float64** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**FailNum** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**SuccessNum** | Pointer to **int64** |  | [optional] 
**TotalNum** | Pointer to **int64** |  | [optional] 
**TotalProgress** | Pointer to **float64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProgressInfo

`func NewProgressInfo() *ProgressInfo`

NewProgressInfo instantiates a new ProgressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgressInfoWithDefaults

`func NewProgressInfoWithDefaults() *ProgressInfo`

NewProgressInfoWithDefaults instantiates a new ProgressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ProgressInfo) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ProgressInfo) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ProgressInfo) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ProgressInfo) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCreate

`func (o *ProgressInfo) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ProgressInfo) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ProgressInfo) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ProgressInfo) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCurrentProgress

`func (o *ProgressInfo) GetCurrentProgress() float64`

GetCurrentProgress returns the CurrentProgress field if non-nil, zero value otherwise.

### GetCurrentProgressOk

`func (o *ProgressInfo) GetCurrentProgressOk() (*float64, bool)`

GetCurrentProgressOk returns a tuple with the CurrentProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentProgress

`func (o *ProgressInfo) SetCurrentProgress(v float64)`

SetCurrentProgress sets CurrentProgress field to given value.

### HasCurrentProgress

`func (o *ProgressInfo) HasCurrentProgress() bool`

HasCurrentProgress returns a boolean if a field has been set.

### GetData

`func (o *ProgressInfo) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProgressInfo) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProgressInfo) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ProgressInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFailNum

`func (o *ProgressInfo) GetFailNum() int64`

GetFailNum returns the FailNum field if non-nil, zero value otherwise.

### GetFailNumOk

`func (o *ProgressInfo) GetFailNumOk() (*int64, bool)`

GetFailNumOk returns a tuple with the FailNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailNum

`func (o *ProgressInfo) SetFailNum(v int64)`

SetFailNum sets FailNum field to given value.

### HasFailNum

`func (o *ProgressInfo) HasFailNum() bool`

HasFailNum returns a boolean if a field has been set.

### GetId

`func (o *ProgressInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProgressInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProgressInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProgressInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSuccessNum

`func (o *ProgressInfo) GetSuccessNum() int64`

GetSuccessNum returns the SuccessNum field if non-nil, zero value otherwise.

### GetSuccessNumOk

`func (o *ProgressInfo) GetSuccessNumOk() (*int64, bool)`

GetSuccessNumOk returns a tuple with the SuccessNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessNum

`func (o *ProgressInfo) SetSuccessNum(v int64)`

SetSuccessNum sets SuccessNum field to given value.

### HasSuccessNum

`func (o *ProgressInfo) HasSuccessNum() bool`

HasSuccessNum returns a boolean if a field has been set.

### GetTotalNum

`func (o *ProgressInfo) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *ProgressInfo) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *ProgressInfo) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.

### HasTotalNum

`func (o *ProgressInfo) HasTotalNum() bool`

HasTotalNum returns a boolean if a field has been set.

### GetTotalProgress

`func (o *ProgressInfo) GetTotalProgress() float64`

GetTotalProgress returns the TotalProgress field if non-nil, zero value otherwise.

### GetTotalProgressOk

`func (o *ProgressInfo) GetTotalProgressOk() (*float64, bool)`

GetTotalProgressOk returns a tuple with the TotalProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProgress

`func (o *ProgressInfo) SetTotalProgress(v float64)`

SetTotalProgress sets TotalProgress field to given value.

### HasTotalProgress

`func (o *ProgressInfo) HasTotalProgress() bool`

HasTotalProgress returns a boolean if a field has been set.

### GetUpdate

`func (o *ProgressInfo) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ProgressInfo) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ProgressInfo) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ProgressInfo) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


