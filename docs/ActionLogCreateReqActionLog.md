# ActionLogCreateReqActionLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**ClientIp** | **string** |  | 
**Finish** | **time.Time** |  | 
**Id** | **int64** |  | 
**Message** | Pointer to **string** |  | [optional] 
**NewData** | **string** |  | 
**OldData** | **string** |  | 
**OspClusterId** | **int64** |  | 
**Parameter** | Pointer to **string** |  | [optional] 
**ResourceId** | **int64** |  | 
**ResourceType** | **string** |  | 
**Start** | **time.Time** |  | 
**Status** | Pointer to **string** |  | [optional] 
**UserId** | **int64** |  | 

## Methods

### NewActionLogCreateReqActionLog

`func NewActionLogCreateReqActionLog(action string, clientIp string, finish time.Time, id int64, newData string, oldData string, ospClusterId int64, resourceId int64, resourceType string, start time.Time, userId int64, ) *ActionLogCreateReqActionLog`

NewActionLogCreateReqActionLog instantiates a new ActionLogCreateReqActionLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionLogCreateReqActionLogWithDefaults

`func NewActionLogCreateReqActionLogWithDefaults() *ActionLogCreateReqActionLog`

NewActionLogCreateReqActionLogWithDefaults instantiates a new ActionLogCreateReqActionLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ActionLogCreateReqActionLog) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ActionLogCreateReqActionLog) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ActionLogCreateReqActionLog) SetAction(v string)`

SetAction sets Action field to given value.


### GetClientIp

`func (o *ActionLogCreateReqActionLog) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *ActionLogCreateReqActionLog) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *ActionLogCreateReqActionLog) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.


### GetFinish

`func (o *ActionLogCreateReqActionLog) GetFinish() time.Time`

GetFinish returns the Finish field if non-nil, zero value otherwise.

### GetFinishOk

`func (o *ActionLogCreateReqActionLog) GetFinishOk() (*time.Time, bool)`

GetFinishOk returns a tuple with the Finish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinish

`func (o *ActionLogCreateReqActionLog) SetFinish(v time.Time)`

SetFinish sets Finish field to given value.


### GetId

`func (o *ActionLogCreateReqActionLog) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionLogCreateReqActionLog) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionLogCreateReqActionLog) SetId(v int64)`

SetId sets Id field to given value.


### GetMessage

`func (o *ActionLogCreateReqActionLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ActionLogCreateReqActionLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ActionLogCreateReqActionLog) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ActionLogCreateReqActionLog) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNewData

`func (o *ActionLogCreateReqActionLog) GetNewData() string`

GetNewData returns the NewData field if non-nil, zero value otherwise.

### GetNewDataOk

`func (o *ActionLogCreateReqActionLog) GetNewDataOk() (*string, bool)`

GetNewDataOk returns a tuple with the NewData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewData

`func (o *ActionLogCreateReqActionLog) SetNewData(v string)`

SetNewData sets NewData field to given value.


### GetOldData

`func (o *ActionLogCreateReqActionLog) GetOldData() string`

GetOldData returns the OldData field if non-nil, zero value otherwise.

### GetOldDataOk

`func (o *ActionLogCreateReqActionLog) GetOldDataOk() (*string, bool)`

GetOldDataOk returns a tuple with the OldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldData

`func (o *ActionLogCreateReqActionLog) SetOldData(v string)`

SetOldData sets OldData field to given value.


### GetOspClusterId

`func (o *ActionLogCreateReqActionLog) GetOspClusterId() int64`

GetOspClusterId returns the OspClusterId field if non-nil, zero value otherwise.

### GetOspClusterIdOk

`func (o *ActionLogCreateReqActionLog) GetOspClusterIdOk() (*int64, bool)`

GetOspClusterIdOk returns a tuple with the OspClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspClusterId

`func (o *ActionLogCreateReqActionLog) SetOspClusterId(v int64)`

SetOspClusterId sets OspClusterId field to given value.


### GetParameter

`func (o *ActionLogCreateReqActionLog) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ActionLogCreateReqActionLog) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ActionLogCreateReqActionLog) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *ActionLogCreateReqActionLog) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetResourceId

`func (o *ActionLogCreateReqActionLog) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ActionLogCreateReqActionLog) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ActionLogCreateReqActionLog) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *ActionLogCreateReqActionLog) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ActionLogCreateReqActionLog) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ActionLogCreateReqActionLog) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetStart

`func (o *ActionLogCreateReqActionLog) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ActionLogCreateReqActionLog) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ActionLogCreateReqActionLog) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetStatus

`func (o *ActionLogCreateReqActionLog) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionLogCreateReqActionLog) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionLogCreateReqActionLog) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActionLogCreateReqActionLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *ActionLogCreateReqActionLog) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ActionLogCreateReqActionLog) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ActionLogCreateReqActionLog) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


