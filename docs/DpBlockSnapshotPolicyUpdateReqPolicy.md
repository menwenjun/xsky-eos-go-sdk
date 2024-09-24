# DpBlockSnapshotPolicyUpdateReqPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpr** | Pointer to **string** | cron expression | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**RetainedMax** | Pointer to **int64** | retained max number of snapshots | [optional] 

## Methods

### NewDpBlockSnapshotPolicyUpdateReqPolicy

`func NewDpBlockSnapshotPolicyUpdateReqPolicy() *DpBlockSnapshotPolicyUpdateReqPolicy`

NewDpBlockSnapshotPolicyUpdateReqPolicy instantiates a new DpBlockSnapshotPolicyUpdateReqPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockSnapshotPolicyUpdateReqPolicyWithDefaults

`func NewDpBlockSnapshotPolicyUpdateReqPolicyWithDefaults() *DpBlockSnapshotPolicyUpdateReqPolicy`

NewDpBlockSnapshotPolicyUpdateReqPolicyWithDefaults instantiates a new DpBlockSnapshotPolicyUpdateReqPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpr

`func (o *DpBlockSnapshotPolicyUpdateReqPolicy) GetCronExpr() string`

GetCronExpr returns the CronExpr field if non-nil, zero value otherwise.

### GetCronExprOk

`func (o *DpBlockSnapshotPolicyUpdateReqPolicy) GetCronExprOk() (*string, bool)`

GetCronExprOk returns a tuple with the CronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpr

`func (o *DpBlockSnapshotPolicyUpdateReqPolicy) SetCronExpr(v string)`

SetCronExpr sets CronExpr field to given value.

### HasCronExpr

`func (o *DpBlockSnapshotPolicyUpdateReqPolicy) HasCronExpr() bool`

HasCronExpr returns a boolean if a field has been set.

### GetName

`func (o *DpBlockSnapshotPolicyUpdateReqPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpBlockSnapshotPolicyUpdateReqPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpBlockSnapshotPolicyUpdateReqPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DpBlockSnapshotPolicyUpdateReqPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetainedMax

`func (o *DpBlockSnapshotPolicyUpdateReqPolicy) GetRetainedMax() int64`

GetRetainedMax returns the RetainedMax field if non-nil, zero value otherwise.

### GetRetainedMaxOk

`func (o *DpBlockSnapshotPolicyUpdateReqPolicy) GetRetainedMaxOk() (*int64, bool)`

GetRetainedMaxOk returns a tuple with the RetainedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedMax

`func (o *DpBlockSnapshotPolicyUpdateReqPolicy) SetRetainedMax(v int64)`

SetRetainedMax sets RetainedMax field to given value.

### HasRetainedMax

`func (o *DpBlockSnapshotPolicyUpdateReqPolicy) HasRetainedMax() bool`

HasRetainedMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


