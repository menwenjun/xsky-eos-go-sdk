# FCPortErrCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**ErrorFrames** | Pointer to **int64** |  | [optional] 
**ErrorFramesOffset** | Pointer to **int64** |  | [optional] 
**ErrorFramesValue** | Pointer to **int64** |  | [optional] 
**FcPort** | Pointer to [**FCPortNestview**](FCPortNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LinkFailureCount** | Pointer to **int64** |  | [optional] 
**LossOfSignalCount** | Pointer to **int64** |  | [optional] 
**LossOfSyncCount** | Pointer to **int64** |  | [optional] 
**Start** | Pointer to **time.Time** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFCPortErrCode

`func NewFCPortErrCode() *FCPortErrCode`

NewFCPortErrCode instantiates a new FCPortErrCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFCPortErrCodeWithDefaults

`func NewFCPortErrCodeWithDefaults() *FCPortErrCode`

NewFCPortErrCodeWithDefaults instantiates a new FCPortErrCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *FCPortErrCode) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *FCPortErrCode) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *FCPortErrCode) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *FCPortErrCode) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetErrorFrames

`func (o *FCPortErrCode) GetErrorFrames() int64`

GetErrorFrames returns the ErrorFrames field if non-nil, zero value otherwise.

### GetErrorFramesOk

`func (o *FCPortErrCode) GetErrorFramesOk() (*int64, bool)`

GetErrorFramesOk returns a tuple with the ErrorFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFrames

`func (o *FCPortErrCode) SetErrorFrames(v int64)`

SetErrorFrames sets ErrorFrames field to given value.

### HasErrorFrames

`func (o *FCPortErrCode) HasErrorFrames() bool`

HasErrorFrames returns a boolean if a field has been set.

### GetErrorFramesOffset

`func (o *FCPortErrCode) GetErrorFramesOffset() int64`

GetErrorFramesOffset returns the ErrorFramesOffset field if non-nil, zero value otherwise.

### GetErrorFramesOffsetOk

`func (o *FCPortErrCode) GetErrorFramesOffsetOk() (*int64, bool)`

GetErrorFramesOffsetOk returns a tuple with the ErrorFramesOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFramesOffset

`func (o *FCPortErrCode) SetErrorFramesOffset(v int64)`

SetErrorFramesOffset sets ErrorFramesOffset field to given value.

### HasErrorFramesOffset

`func (o *FCPortErrCode) HasErrorFramesOffset() bool`

HasErrorFramesOffset returns a boolean if a field has been set.

### GetErrorFramesValue

`func (o *FCPortErrCode) GetErrorFramesValue() int64`

GetErrorFramesValue returns the ErrorFramesValue field if non-nil, zero value otherwise.

### GetErrorFramesValueOk

`func (o *FCPortErrCode) GetErrorFramesValueOk() (*int64, bool)`

GetErrorFramesValueOk returns a tuple with the ErrorFramesValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFramesValue

`func (o *FCPortErrCode) SetErrorFramesValue(v int64)`

SetErrorFramesValue sets ErrorFramesValue field to given value.

### HasErrorFramesValue

`func (o *FCPortErrCode) HasErrorFramesValue() bool`

HasErrorFramesValue returns a boolean if a field has been set.

### GetFcPort

`func (o *FCPortErrCode) GetFcPort() FCPortNestview`

GetFcPort returns the FcPort field if non-nil, zero value otherwise.

### GetFcPortOk

`func (o *FCPortErrCode) GetFcPortOk() (*FCPortNestview, bool)`

GetFcPortOk returns a tuple with the FcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPort

`func (o *FCPortErrCode) SetFcPort(v FCPortNestview)`

SetFcPort sets FcPort field to given value.

### HasFcPort

`func (o *FCPortErrCode) HasFcPort() bool`

HasFcPort returns a boolean if a field has been set.

### GetId

`func (o *FCPortErrCode) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FCPortErrCode) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FCPortErrCode) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FCPortErrCode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinkFailureCount

`func (o *FCPortErrCode) GetLinkFailureCount() int64`

GetLinkFailureCount returns the LinkFailureCount field if non-nil, zero value otherwise.

### GetLinkFailureCountOk

`func (o *FCPortErrCode) GetLinkFailureCountOk() (*int64, bool)`

GetLinkFailureCountOk returns a tuple with the LinkFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkFailureCount

`func (o *FCPortErrCode) SetLinkFailureCount(v int64)`

SetLinkFailureCount sets LinkFailureCount field to given value.

### HasLinkFailureCount

`func (o *FCPortErrCode) HasLinkFailureCount() bool`

HasLinkFailureCount returns a boolean if a field has been set.

### GetLossOfSignalCount

`func (o *FCPortErrCode) GetLossOfSignalCount() int64`

GetLossOfSignalCount returns the LossOfSignalCount field if non-nil, zero value otherwise.

### GetLossOfSignalCountOk

`func (o *FCPortErrCode) GetLossOfSignalCountOk() (*int64, bool)`

GetLossOfSignalCountOk returns a tuple with the LossOfSignalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossOfSignalCount

`func (o *FCPortErrCode) SetLossOfSignalCount(v int64)`

SetLossOfSignalCount sets LossOfSignalCount field to given value.

### HasLossOfSignalCount

`func (o *FCPortErrCode) HasLossOfSignalCount() bool`

HasLossOfSignalCount returns a boolean if a field has been set.

### GetLossOfSyncCount

`func (o *FCPortErrCode) GetLossOfSyncCount() int64`

GetLossOfSyncCount returns the LossOfSyncCount field if non-nil, zero value otherwise.

### GetLossOfSyncCountOk

`func (o *FCPortErrCode) GetLossOfSyncCountOk() (*int64, bool)`

GetLossOfSyncCountOk returns a tuple with the LossOfSyncCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossOfSyncCount

`func (o *FCPortErrCode) SetLossOfSyncCount(v int64)`

SetLossOfSyncCount sets LossOfSyncCount field to given value.

### HasLossOfSyncCount

`func (o *FCPortErrCode) HasLossOfSyncCount() bool`

HasLossOfSyncCount returns a boolean if a field has been set.

### GetStart

`func (o *FCPortErrCode) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *FCPortErrCode) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *FCPortErrCode) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *FCPortErrCode) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUpdate

`func (o *FCPortErrCode) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *FCPortErrCode) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *FCPortErrCode) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *FCPortErrCode) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


