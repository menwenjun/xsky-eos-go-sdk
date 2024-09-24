# DfsTrashFileJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**ProgressInfo** | Pointer to [**ProgressInfo**](ProgressInfo.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Trash** | Pointer to [**DfsTrashNestview**](DfsTrashNestview.md) |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsTrashFileJob

`func NewDfsTrashFileJob() *DfsTrashFileJob`

NewDfsTrashFileJob instantiates a new DfsTrashFileJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsTrashFileJobWithDefaults

`func NewDfsTrashFileJobWithDefaults() *DfsTrashFileJob`

NewDfsTrashFileJobWithDefaults instantiates a new DfsTrashFileJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DfsTrashFileJob) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DfsTrashFileJob) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DfsTrashFileJob) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DfsTrashFileJob) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCreate

`func (o *DfsTrashFileJob) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsTrashFileJob) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsTrashFileJob) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsTrashFileJob) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *DfsTrashFileJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsTrashFileJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsTrashFileJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsTrashFileJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *DfsTrashFileJob) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsTrashFileJob) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsTrashFileJob) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DfsTrashFileJob) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProgress

`func (o *DfsTrashFileJob) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DfsTrashFileJob) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DfsTrashFileJob) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *DfsTrashFileJob) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProgressInfo

`func (o *DfsTrashFileJob) GetProgressInfo() ProgressInfo`

GetProgressInfo returns the ProgressInfo field if non-nil, zero value otherwise.

### GetProgressInfoOk

`func (o *DfsTrashFileJob) GetProgressInfoOk() (*ProgressInfo, bool)`

GetProgressInfoOk returns a tuple with the ProgressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressInfo

`func (o *DfsTrashFileJob) SetProgressInfo(v ProgressInfo)`

SetProgressInfo sets ProgressInfo field to given value.

### HasProgressInfo

`func (o *DfsTrashFileJob) HasProgressInfo() bool`

HasProgressInfo returns a boolean if a field has been set.

### GetStatus

`func (o *DfsTrashFileJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsTrashFileJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsTrashFileJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsTrashFileJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrash

`func (o *DfsTrashFileJob) GetTrash() DfsTrashNestview`

GetTrash returns the Trash field if non-nil, zero value otherwise.

### GetTrashOk

`func (o *DfsTrashFileJob) GetTrashOk() (*DfsTrashNestview, bool)`

GetTrashOk returns a tuple with the Trash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrash

`func (o *DfsTrashFileJob) SetTrash(v DfsTrashNestview)`

SetTrash sets Trash field to given value.

### HasTrash

`func (o *DfsTrashFileJob) HasTrash() bool`

HasTrash returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsTrashFileJob) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsTrashFileJob) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsTrashFileJob) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsTrashFileJob) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


