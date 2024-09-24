# DpDfsSnapshotPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**CronExpr** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProtectedPathNum** | Pointer to **int64** |  | [optional] 
**RetainTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDpDfsSnapshotPolicy

`func NewDpDfsSnapshotPolicy() *DpDfsSnapshotPolicy`

NewDpDfsSnapshotPolicy instantiates a new DpDfsSnapshotPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpDfsSnapshotPolicyWithDefaults

`func NewDpDfsSnapshotPolicyWithDefaults() *DpDfsSnapshotPolicy`

NewDpDfsSnapshotPolicyWithDefaults instantiates a new DpDfsSnapshotPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DpDfsSnapshotPolicy) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpDfsSnapshotPolicy) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpDfsSnapshotPolicy) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpDfsSnapshotPolicy) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DpDfsSnapshotPolicy) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DpDfsSnapshotPolicy) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DpDfsSnapshotPolicy) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DpDfsSnapshotPolicy) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCronExpr

`func (o *DpDfsSnapshotPolicy) GetCronExpr() string`

GetCronExpr returns the CronExpr field if non-nil, zero value otherwise.

### GetCronExprOk

`func (o *DpDfsSnapshotPolicy) GetCronExprOk() (*string, bool)`

GetCronExprOk returns a tuple with the CronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpr

`func (o *DpDfsSnapshotPolicy) SetCronExpr(v string)`

SetCronExpr sets CronExpr field to given value.

### HasCronExpr

`func (o *DpDfsSnapshotPolicy) HasCronExpr() bool`

HasCronExpr returns a boolean if a field has been set.

### GetDescription

`func (o *DpDfsSnapshotPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DpDfsSnapshotPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DpDfsSnapshotPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DpDfsSnapshotPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DpDfsSnapshotPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpDfsSnapshotPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpDfsSnapshotPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpDfsSnapshotPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DpDfsSnapshotPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpDfsSnapshotPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpDfsSnapshotPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DpDfsSnapshotPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtectedPathNum

`func (o *DpDfsSnapshotPolicy) GetProtectedPathNum() int64`

GetProtectedPathNum returns the ProtectedPathNum field if non-nil, zero value otherwise.

### GetProtectedPathNumOk

`func (o *DpDfsSnapshotPolicy) GetProtectedPathNumOk() (*int64, bool)`

GetProtectedPathNumOk returns a tuple with the ProtectedPathNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedPathNum

`func (o *DpDfsSnapshotPolicy) SetProtectedPathNum(v int64)`

SetProtectedPathNum sets ProtectedPathNum field to given value.

### HasProtectedPathNum

`func (o *DpDfsSnapshotPolicy) HasProtectedPathNum() bool`

HasProtectedPathNum returns a boolean if a field has been set.

### GetRetainTime

`func (o *DpDfsSnapshotPolicy) GetRetainTime() string`

GetRetainTime returns the RetainTime field if non-nil, zero value otherwise.

### GetRetainTimeOk

`func (o *DpDfsSnapshotPolicy) GetRetainTimeOk() (*string, bool)`

GetRetainTimeOk returns a tuple with the RetainTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainTime

`func (o *DpDfsSnapshotPolicy) SetRetainTime(v string)`

SetRetainTime sets RetainTime field to given value.

### HasRetainTime

`func (o *DpDfsSnapshotPolicy) HasRetainTime() bool`

HasRetainTime returns a boolean if a field has been set.

### GetStatus

`func (o *DpDfsSnapshotPolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpDfsSnapshotPolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpDfsSnapshotPolicy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpDfsSnapshotPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DpDfsSnapshotPolicy) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DpDfsSnapshotPolicy) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DpDfsSnapshotPolicy) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DpDfsSnapshotPolicy) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


