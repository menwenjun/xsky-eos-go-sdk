# DpBlockAsyncReplicationPolicyCreateReqPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpr** | **string** | cron expression | 
**DpGatewayId** | **int64** | dp gateway | 
**DpSiteId** | **int64** | dp site | 
**Name** | **string** | name | 
**RetainedMax** | Pointer to **int64** | retained max number of snapshots | [optional] 
**Timeout** | Pointer to **int64** | timeout | [optional] 

## Methods

### NewDpBlockAsyncReplicationPolicyCreateReqPolicy

`func NewDpBlockAsyncReplicationPolicyCreateReqPolicy(cronExpr string, dpGatewayId int64, dpSiteId int64, name string, ) *DpBlockAsyncReplicationPolicyCreateReqPolicy`

NewDpBlockAsyncReplicationPolicyCreateReqPolicy instantiates a new DpBlockAsyncReplicationPolicyCreateReqPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockAsyncReplicationPolicyCreateReqPolicyWithDefaults

`func NewDpBlockAsyncReplicationPolicyCreateReqPolicyWithDefaults() *DpBlockAsyncReplicationPolicyCreateReqPolicy`

NewDpBlockAsyncReplicationPolicyCreateReqPolicyWithDefaults instantiates a new DpBlockAsyncReplicationPolicyCreateReqPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpr

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) GetCronExpr() string`

GetCronExpr returns the CronExpr field if non-nil, zero value otherwise.

### GetCronExprOk

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) GetCronExprOk() (*string, bool)`

GetCronExprOk returns a tuple with the CronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpr

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) SetCronExpr(v string)`

SetCronExpr sets CronExpr field to given value.


### GetDpGatewayId

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) GetDpGatewayId() int64`

GetDpGatewayId returns the DpGatewayId field if non-nil, zero value otherwise.

### GetDpGatewayIdOk

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) GetDpGatewayIdOk() (*int64, bool)`

GetDpGatewayIdOk returns a tuple with the DpGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpGatewayId

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) SetDpGatewayId(v int64)`

SetDpGatewayId sets DpGatewayId field to given value.


### GetDpSiteId

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) GetDpSiteId() int64`

GetDpSiteId returns the DpSiteId field if non-nil, zero value otherwise.

### GetDpSiteIdOk

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) GetDpSiteIdOk() (*int64, bool)`

GetDpSiteIdOk returns a tuple with the DpSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpSiteId

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) SetDpSiteId(v int64)`

SetDpSiteId sets DpSiteId field to given value.


### GetName

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetRetainedMax

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) GetRetainedMax() int64`

GetRetainedMax returns the RetainedMax field if non-nil, zero value otherwise.

### GetRetainedMaxOk

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) GetRetainedMaxOk() (*int64, bool)`

GetRetainedMaxOk returns a tuple with the RetainedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedMax

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) SetRetainedMax(v int64)`

SetRetainedMax sets RetainedMax field to given value.

### HasRetainedMax

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) HasRetainedMax() bool`

HasRetainedMax returns a boolean if a field has been set.

### GetTimeout

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DpBlockAsyncReplicationPolicyCreateReqPolicy) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


