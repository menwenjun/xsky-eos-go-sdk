/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the Task type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Task{}

// Task Task defines a task will be run on a host +X:model:generate
type Task struct {
	Create *time.Time `json:"create,omitempty"`
	Data *string `json:"data,omitempty"`
	ErrorMsg *string `json:"error_msg,omitempty"`
	Execute *time.Time `json:"execute,omitempty"`
	Finish *time.Time `json:"finish,omitempty"`
	Host *HostNestview `json:"host,omitempty"`
	HostRole *string `json:"host_role,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Priority *int64 `json:"priority,omitempty"`
	Retried *bool `json:"retried,omitempty"`
	Retry *int64 `json:"retry,omitempty"`
	RetryType *string `json:"retry_type,omitempty"`
	Status *string `json:"status,omitempty"`
	Type *string `json:"type,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewTask instantiates a new Task object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTask() *Task {
	this := Task{}
	return &this
}

// NewTaskWithDefaults instantiates a new Task object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWithDefaults() *Task {
	this := Task{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *Task) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *Task) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *Task) SetCreate(v time.Time) {
	o.Create = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Task) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Task) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *Task) SetData(v string) {
	o.Data = &v
}

// GetErrorMsg returns the ErrorMsg field value if set, zero value otherwise.
func (o *Task) GetErrorMsg() string {
	if o == nil || IsNil(o.ErrorMsg) {
		var ret string
		return ret
	}
	return *o.ErrorMsg
}

// GetErrorMsgOk returns a tuple with the ErrorMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetErrorMsgOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMsg) {
		return nil, false
	}
	return o.ErrorMsg, true
}

// HasErrorMsg returns a boolean if a field has been set.
func (o *Task) HasErrorMsg() bool {
	if o != nil && !IsNil(o.ErrorMsg) {
		return true
	}

	return false
}

// SetErrorMsg gets a reference to the given string and assigns it to the ErrorMsg field.
func (o *Task) SetErrorMsg(v string) {
	o.ErrorMsg = &v
}

// GetExecute returns the Execute field value if set, zero value otherwise.
func (o *Task) GetExecute() time.Time {
	if o == nil || IsNil(o.Execute) {
		var ret time.Time
		return ret
	}
	return *o.Execute
}

// GetExecuteOk returns a tuple with the Execute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetExecuteOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Execute) {
		return nil, false
	}
	return o.Execute, true
}

// HasExecute returns a boolean if a field has been set.
func (o *Task) HasExecute() bool {
	if o != nil && !IsNil(o.Execute) {
		return true
	}

	return false
}

// SetExecute gets a reference to the given time.Time and assigns it to the Execute field.
func (o *Task) SetExecute(v time.Time) {
	o.Execute = &v
}

// GetFinish returns the Finish field value if set, zero value otherwise.
func (o *Task) GetFinish() time.Time {
	if o == nil || IsNil(o.Finish) {
		var ret time.Time
		return ret
	}
	return *o.Finish
}

// GetFinishOk returns a tuple with the Finish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetFinishOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Finish) {
		return nil, false
	}
	return o.Finish, true
}

// HasFinish returns a boolean if a field has been set.
func (o *Task) HasFinish() bool {
	if o != nil && !IsNil(o.Finish) {
		return true
	}

	return false
}

// SetFinish gets a reference to the given time.Time and assigns it to the Finish field.
func (o *Task) SetFinish(v time.Time) {
	o.Finish = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *Task) GetHost() HostNestview {
	if o == nil || IsNil(o.Host) {
		var ret HostNestview
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetHostOk() (*HostNestview, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *Task) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given HostNestview and assigns it to the Host field.
func (o *Task) SetHost(v HostNestview) {
	o.Host = &v
}

// GetHostRole returns the HostRole field value if set, zero value otherwise.
func (o *Task) GetHostRole() string {
	if o == nil || IsNil(o.HostRole) {
		var ret string
		return ret
	}
	return *o.HostRole
}

// GetHostRoleOk returns a tuple with the HostRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetHostRoleOk() (*string, bool) {
	if o == nil || IsNil(o.HostRole) {
		return nil, false
	}
	return o.HostRole, true
}

// HasHostRole returns a boolean if a field has been set.
func (o *Task) HasHostRole() bool {
	if o != nil && !IsNil(o.HostRole) {
		return true
	}

	return false
}

// SetHostRole gets a reference to the given string and assigns it to the HostRole field.
func (o *Task) SetHostRole(v string) {
	o.HostRole = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Task) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Task) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Task) SetId(v int64) {
	o.Id = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *Task) GetPriority() int64 {
	if o == nil || IsNil(o.Priority) {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetPriorityOk() (*int64, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *Task) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *Task) SetPriority(v int64) {
	o.Priority = &v
}

// GetRetried returns the Retried field value if set, zero value otherwise.
func (o *Task) GetRetried() bool {
	if o == nil || IsNil(o.Retried) {
		var ret bool
		return ret
	}
	return *o.Retried
}

// GetRetriedOk returns a tuple with the Retried field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetRetriedOk() (*bool, bool) {
	if o == nil || IsNil(o.Retried) {
		return nil, false
	}
	return o.Retried, true
}

// HasRetried returns a boolean if a field has been set.
func (o *Task) HasRetried() bool {
	if o != nil && !IsNil(o.Retried) {
		return true
	}

	return false
}

// SetRetried gets a reference to the given bool and assigns it to the Retried field.
func (o *Task) SetRetried(v bool) {
	o.Retried = &v
}

// GetRetry returns the Retry field value if set, zero value otherwise.
func (o *Task) GetRetry() int64 {
	if o == nil || IsNil(o.Retry) {
		var ret int64
		return ret
	}
	return *o.Retry
}

// GetRetryOk returns a tuple with the Retry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetRetryOk() (*int64, bool) {
	if o == nil || IsNil(o.Retry) {
		return nil, false
	}
	return o.Retry, true
}

// HasRetry returns a boolean if a field has been set.
func (o *Task) HasRetry() bool {
	if o != nil && !IsNil(o.Retry) {
		return true
	}

	return false
}

// SetRetry gets a reference to the given int64 and assigns it to the Retry field.
func (o *Task) SetRetry(v int64) {
	o.Retry = &v
}

// GetRetryType returns the RetryType field value if set, zero value otherwise.
func (o *Task) GetRetryType() string {
	if o == nil || IsNil(o.RetryType) {
		var ret string
		return ret
	}
	return *o.RetryType
}

// GetRetryTypeOk returns a tuple with the RetryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetRetryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RetryType) {
		return nil, false
	}
	return o.RetryType, true
}

// HasRetryType returns a boolean if a field has been set.
func (o *Task) HasRetryType() bool {
	if o != nil && !IsNil(o.RetryType) {
		return true
	}

	return false
}

// SetRetryType gets a reference to the given string and assigns it to the RetryType field.
func (o *Task) SetRetryType(v string) {
	o.RetryType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Task) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Task) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Task) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Task) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Task) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Task) SetType(v string) {
	o.Type = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *Task) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *Task) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *Task) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o Task) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Task) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.ErrorMsg) {
		toSerialize["error_msg"] = o.ErrorMsg
	}
	if !IsNil(o.Execute) {
		toSerialize["execute"] = o.Execute
	}
	if !IsNil(o.Finish) {
		toSerialize["finish"] = o.Finish
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.HostRole) {
		toSerialize["host_role"] = o.HostRole
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Retried) {
		toSerialize["retried"] = o.Retried
	}
	if !IsNil(o.Retry) {
		toSerialize["retry"] = o.Retry
	}
	if !IsNil(o.RetryType) {
		toSerialize["retry_type"] = o.RetryType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullableTask struct {
	value *Task
	isSet bool
}

func (v NullableTask) Get() *Task {
	return v.value
}

func (v *NullableTask) Set(val *Task) {
	v.value = val
	v.isSet = true
}

func (v NullableTask) IsSet() bool {
	return v.isSet
}

func (v *NullableTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTask(val *Task) *NullableTask {
	return &NullableTask{value: val, isSet: true}
}

func (v NullableTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


