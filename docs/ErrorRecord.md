# ErrorRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionLog** | Pointer to [**ActionLogNestview**](ActionLogNestview.md) |  | [optional] 
**Alert** | Pointer to [**AlertNestview**](AlertNestview.md) |  | [optional] 
**AlertInfo** | Pointer to [**AlertInfoNestview**](AlertInfoNestview.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**ExtraData** | Pointer to **map[string]interface{}** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorRecord

`func NewErrorRecord() *ErrorRecord`

NewErrorRecord instantiates a new ErrorRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorRecordWithDefaults

`func NewErrorRecordWithDefaults() *ErrorRecord`

NewErrorRecordWithDefaults instantiates a new ErrorRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionLog

`func (o *ErrorRecord) GetActionLog() ActionLogNestview`

GetActionLog returns the ActionLog field if non-nil, zero value otherwise.

### GetActionLogOk

`func (o *ErrorRecord) GetActionLogOk() (*ActionLogNestview, bool)`

GetActionLogOk returns a tuple with the ActionLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionLog

`func (o *ErrorRecord) SetActionLog(v ActionLogNestview)`

SetActionLog sets ActionLog field to given value.

### HasActionLog

`func (o *ErrorRecord) HasActionLog() bool`

HasActionLog returns a boolean if a field has been set.

### GetAlert

`func (o *ErrorRecord) GetAlert() AlertNestview`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *ErrorRecord) GetAlertOk() (*AlertNestview, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *ErrorRecord) SetAlert(v AlertNestview)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *ErrorRecord) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetAlertInfo

`func (o *ErrorRecord) GetAlertInfo() AlertInfoNestview`

GetAlertInfo returns the AlertInfo field if non-nil, zero value otherwise.

### GetAlertInfoOk

`func (o *ErrorRecord) GetAlertInfoOk() (*AlertInfoNestview, bool)`

GetAlertInfoOk returns a tuple with the AlertInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertInfo

`func (o *ErrorRecord) SetAlertInfo(v AlertInfoNestview)`

SetAlertInfo sets AlertInfo field to given value.

### HasAlertInfo

`func (o *ErrorRecord) HasAlertInfo() bool`

HasAlertInfo returns a boolean if a field has been set.

### GetCode

`func (o *ErrorRecord) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorRecord) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorRecord) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorRecord) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreate

`func (o *ErrorRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ErrorRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ErrorRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ErrorRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDetails

`func (o *ErrorRecord) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorRecord) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorRecord) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ErrorRecord) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetExtraData

`func (o *ErrorRecord) GetExtraData() map[string]interface{}`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ErrorRecord) GetExtraDataOk() (*map[string]interface{}, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ErrorRecord) SetExtraData(v map[string]interface{})`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *ErrorRecord) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetHost

`func (o *ErrorRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ErrorRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ErrorRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *ErrorRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *ErrorRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ErrorRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorRecord) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorRecord) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorRecord) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorRecord) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


