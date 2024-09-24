# DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpr** | Pointer to **string** | cron expression | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**RetainedMax** | Pointer to **int64** | retained max number of snapshots | [optional] 

## Methods

### NewDpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy

`func NewDpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy() *DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy`

NewDpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy instantiates a new DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicyWithDefaults

`func NewDpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicyWithDefaults() *DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy`

NewDpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicyWithDefaults instantiates a new DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpr

`func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy) GetCronExpr() string`

GetCronExpr returns the CronExpr field if non-nil, zero value otherwise.

### GetCronExprOk

`func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy) GetCronExprOk() (*string, bool)`

GetCronExprOk returns a tuple with the CronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpr

`func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy) SetCronExpr(v string)`

SetCronExpr sets CronExpr field to given value.

### HasCronExpr

`func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy) HasCronExpr() bool`

HasCronExpr returns a boolean if a field has been set.

### GetName

`func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetainedMax

`func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy) GetRetainedMax() int64`

GetRetainedMax returns the RetainedMax field if non-nil, zero value otherwise.

### GetRetainedMaxOk

`func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy) GetRetainedMaxOk() (*int64, bool)`

GetRetainedMaxOk returns a tuple with the RetainedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedMax

`func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy) SetRetainedMax(v int64)`

SetRetainedMax sets RetainedMax field to given value.

### HasRetainedMax

`func (o *DpVolumeGroupSnapshotReplicationPolicyUpdateReqPolicy) HasRetainedMax() bool`

HasRetainedMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


