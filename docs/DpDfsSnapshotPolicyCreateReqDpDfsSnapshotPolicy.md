# DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpr** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DpGatewayId** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**RetainTime** | **string** |  | 

## Methods

### NewDpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy

`func NewDpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy(cronExpr string, name string, retainTime string, ) *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy`

NewDpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy instantiates a new DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicyWithDefaults

`func NewDpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicyWithDefaults() *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy`

NewDpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicyWithDefaults instantiates a new DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpr

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) GetCronExpr() string`

GetCronExpr returns the CronExpr field if non-nil, zero value otherwise.

### GetCronExprOk

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) GetCronExprOk() (*string, bool)`

GetCronExprOk returns a tuple with the CronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpr

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) SetCronExpr(v string)`

SetCronExpr sets CronExpr field to given value.


### GetDescription

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDpGatewayId

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) GetDpGatewayId() int64`

GetDpGatewayId returns the DpGatewayId field if non-nil, zero value otherwise.

### GetDpGatewayIdOk

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) GetDpGatewayIdOk() (*int64, bool)`

GetDpGatewayIdOk returns a tuple with the DpGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpGatewayId

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) SetDpGatewayId(v int64)`

SetDpGatewayId sets DpGatewayId field to given value.

### HasDpGatewayId

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) HasDpGatewayId() bool`

HasDpGatewayId returns a boolean if a field has been set.

### GetName

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetRetainTime

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) GetRetainTime() string`

GetRetainTime returns the RetainTime field if non-nil, zero value otherwise.

### GetRetainTimeOk

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) GetRetainTimeOk() (*string, bool)`

GetRetainTimeOk returns a tuple with the RetainTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainTime

`func (o *DpDfsSnapshotPolicyCreateReqDpDfsSnapshotPolicy) SetRetainTime(v string)`

SetRetainTime sets RetainTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


