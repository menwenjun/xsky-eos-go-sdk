# TasksResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tasks** | [**[]Task**](Task.md) | tasks | 

## Methods

### NewTasksResp

`func NewTasksResp(tasks []Task, ) *TasksResp`

NewTasksResp instantiates a new TasksResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasksRespWithDefaults

`func NewTasksRespWithDefaults() *TasksResp`

NewTasksRespWithDefaults instantiates a new TasksResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTasks

`func (o *TasksResp) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *TasksResp) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *TasksResp) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


