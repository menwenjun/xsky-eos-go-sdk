# DpBlockSnapshotPolicyCreateReqPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpr** | **string** | cron expression | 
**DpGatewayId** | **int64** | dp gateway | 
**Name** | **string** | name | 
**RetainedMax** | **int64** | retained max number of snapshots | 

## Methods

### NewDpBlockSnapshotPolicyCreateReqPolicy

`func NewDpBlockSnapshotPolicyCreateReqPolicy(cronExpr string, dpGatewayId int64, name string, retainedMax int64, ) *DpBlockSnapshotPolicyCreateReqPolicy`

NewDpBlockSnapshotPolicyCreateReqPolicy instantiates a new DpBlockSnapshotPolicyCreateReqPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockSnapshotPolicyCreateReqPolicyWithDefaults

`func NewDpBlockSnapshotPolicyCreateReqPolicyWithDefaults() *DpBlockSnapshotPolicyCreateReqPolicy`

NewDpBlockSnapshotPolicyCreateReqPolicyWithDefaults instantiates a new DpBlockSnapshotPolicyCreateReqPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpr

`func (o *DpBlockSnapshotPolicyCreateReqPolicy) GetCronExpr() string`

GetCronExpr returns the CronExpr field if non-nil, zero value otherwise.

### GetCronExprOk

`func (o *DpBlockSnapshotPolicyCreateReqPolicy) GetCronExprOk() (*string, bool)`

GetCronExprOk returns a tuple with the CronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpr

`func (o *DpBlockSnapshotPolicyCreateReqPolicy) SetCronExpr(v string)`

SetCronExpr sets CronExpr field to given value.


### GetDpGatewayId

`func (o *DpBlockSnapshotPolicyCreateReqPolicy) GetDpGatewayId() int64`

GetDpGatewayId returns the DpGatewayId field if non-nil, zero value otherwise.

### GetDpGatewayIdOk

`func (o *DpBlockSnapshotPolicyCreateReqPolicy) GetDpGatewayIdOk() (*int64, bool)`

GetDpGatewayIdOk returns a tuple with the DpGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpGatewayId

`func (o *DpBlockSnapshotPolicyCreateReqPolicy) SetDpGatewayId(v int64)`

SetDpGatewayId sets DpGatewayId field to given value.


### GetName

`func (o *DpBlockSnapshotPolicyCreateReqPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpBlockSnapshotPolicyCreateReqPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpBlockSnapshotPolicyCreateReqPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetRetainedMax

`func (o *DpBlockSnapshotPolicyCreateReqPolicy) GetRetainedMax() int64`

GetRetainedMax returns the RetainedMax field if non-nil, zero value otherwise.

### GetRetainedMaxOk

`func (o *DpBlockSnapshotPolicyCreateReqPolicy) GetRetainedMaxOk() (*int64, bool)`

GetRetainedMaxOk returns a tuple with the RetainedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedMax

`func (o *DpBlockSnapshotPolicyCreateReqPolicy) SetRetainedMax(v int64)`

SetRetainedMax sets RetainedMax field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


