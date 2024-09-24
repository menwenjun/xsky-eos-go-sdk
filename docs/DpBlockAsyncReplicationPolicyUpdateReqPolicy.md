# DpBlockAsyncReplicationPolicyUpdateReqPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpr** | Pointer to **string** | cron expression | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**RetainedMax** | Pointer to **int64** | retained max number of snapshots | [optional] 

## Methods

### NewDpBlockAsyncReplicationPolicyUpdateReqPolicy

`func NewDpBlockAsyncReplicationPolicyUpdateReqPolicy() *DpBlockAsyncReplicationPolicyUpdateReqPolicy`

NewDpBlockAsyncReplicationPolicyUpdateReqPolicy instantiates a new DpBlockAsyncReplicationPolicyUpdateReqPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockAsyncReplicationPolicyUpdateReqPolicyWithDefaults

`func NewDpBlockAsyncReplicationPolicyUpdateReqPolicyWithDefaults() *DpBlockAsyncReplicationPolicyUpdateReqPolicy`

NewDpBlockAsyncReplicationPolicyUpdateReqPolicyWithDefaults instantiates a new DpBlockAsyncReplicationPolicyUpdateReqPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpr

`func (o *DpBlockAsyncReplicationPolicyUpdateReqPolicy) GetCronExpr() string`

GetCronExpr returns the CronExpr field if non-nil, zero value otherwise.

### GetCronExprOk

`func (o *DpBlockAsyncReplicationPolicyUpdateReqPolicy) GetCronExprOk() (*string, bool)`

GetCronExprOk returns a tuple with the CronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpr

`func (o *DpBlockAsyncReplicationPolicyUpdateReqPolicy) SetCronExpr(v string)`

SetCronExpr sets CronExpr field to given value.

### HasCronExpr

`func (o *DpBlockAsyncReplicationPolicyUpdateReqPolicy) HasCronExpr() bool`

HasCronExpr returns a boolean if a field has been set.

### GetName

`func (o *DpBlockAsyncReplicationPolicyUpdateReqPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpBlockAsyncReplicationPolicyUpdateReqPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpBlockAsyncReplicationPolicyUpdateReqPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DpBlockAsyncReplicationPolicyUpdateReqPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetainedMax

`func (o *DpBlockAsyncReplicationPolicyUpdateReqPolicy) GetRetainedMax() int64`

GetRetainedMax returns the RetainedMax field if non-nil, zero value otherwise.

### GetRetainedMaxOk

`func (o *DpBlockAsyncReplicationPolicyUpdateReqPolicy) GetRetainedMaxOk() (*int64, bool)`

GetRetainedMaxOk returns a tuple with the RetainedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedMax

`func (o *DpBlockAsyncReplicationPolicyUpdateReqPolicy) SetRetainedMax(v int64)`

SetRetainedMax sets RetainedMax field to given value.

### HasRetainedMax

`func (o *DpBlockAsyncReplicationPolicyUpdateReqPolicy) HasRetainedMax() bool`

HasRetainedMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


