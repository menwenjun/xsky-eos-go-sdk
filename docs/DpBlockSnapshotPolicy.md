# DpBlockSnapshotPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockVolumeNum** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**CronExpr** | Pointer to **string** |  | [optional] 
**DpGateway** | Pointer to [**DpGatewayNestview**](DpGatewayNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RetainedMax** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDpBlockSnapshotPolicy

`func NewDpBlockSnapshotPolicy() *DpBlockSnapshotPolicy`

NewDpBlockSnapshotPolicy instantiates a new DpBlockSnapshotPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockSnapshotPolicyWithDefaults

`func NewDpBlockSnapshotPolicyWithDefaults() *DpBlockSnapshotPolicy`

NewDpBlockSnapshotPolicyWithDefaults instantiates a new DpBlockSnapshotPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockVolumeNum

`func (o *DpBlockSnapshotPolicy) GetBlockVolumeNum() int64`

GetBlockVolumeNum returns the BlockVolumeNum field if non-nil, zero value otherwise.

### GetBlockVolumeNumOk

`func (o *DpBlockSnapshotPolicy) GetBlockVolumeNumOk() (*int64, bool)`

GetBlockVolumeNumOk returns a tuple with the BlockVolumeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeNum

`func (o *DpBlockSnapshotPolicy) SetBlockVolumeNum(v int64)`

SetBlockVolumeNum sets BlockVolumeNum field to given value.

### HasBlockVolumeNum

`func (o *DpBlockSnapshotPolicy) HasBlockVolumeNum() bool`

HasBlockVolumeNum returns a boolean if a field has been set.

### GetCluster

`func (o *DpBlockSnapshotPolicy) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpBlockSnapshotPolicy) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpBlockSnapshotPolicy) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpBlockSnapshotPolicy) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DpBlockSnapshotPolicy) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DpBlockSnapshotPolicy) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DpBlockSnapshotPolicy) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DpBlockSnapshotPolicy) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCronExpr

`func (o *DpBlockSnapshotPolicy) GetCronExpr() string`

GetCronExpr returns the CronExpr field if non-nil, zero value otherwise.

### GetCronExprOk

`func (o *DpBlockSnapshotPolicy) GetCronExprOk() (*string, bool)`

GetCronExprOk returns a tuple with the CronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpr

`func (o *DpBlockSnapshotPolicy) SetCronExpr(v string)`

SetCronExpr sets CronExpr field to given value.

### HasCronExpr

`func (o *DpBlockSnapshotPolicy) HasCronExpr() bool`

HasCronExpr returns a boolean if a field has been set.

### GetDpGateway

`func (o *DpBlockSnapshotPolicy) GetDpGateway() DpGatewayNestview`

GetDpGateway returns the DpGateway field if non-nil, zero value otherwise.

### GetDpGatewayOk

`func (o *DpBlockSnapshotPolicy) GetDpGatewayOk() (*DpGatewayNestview, bool)`

GetDpGatewayOk returns a tuple with the DpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpGateway

`func (o *DpBlockSnapshotPolicy) SetDpGateway(v DpGatewayNestview)`

SetDpGateway sets DpGateway field to given value.

### HasDpGateway

`func (o *DpBlockSnapshotPolicy) HasDpGateway() bool`

HasDpGateway returns a boolean if a field has been set.

### GetId

`func (o *DpBlockSnapshotPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpBlockSnapshotPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpBlockSnapshotPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpBlockSnapshotPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DpBlockSnapshotPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpBlockSnapshotPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpBlockSnapshotPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DpBlockSnapshotPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetainedMax

`func (o *DpBlockSnapshotPolicy) GetRetainedMax() int64`

GetRetainedMax returns the RetainedMax field if non-nil, zero value otherwise.

### GetRetainedMaxOk

`func (o *DpBlockSnapshotPolicy) GetRetainedMaxOk() (*int64, bool)`

GetRetainedMaxOk returns a tuple with the RetainedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedMax

`func (o *DpBlockSnapshotPolicy) SetRetainedMax(v int64)`

SetRetainedMax sets RetainedMax field to given value.

### HasRetainedMax

`func (o *DpBlockSnapshotPolicy) HasRetainedMax() bool`

HasRetainedMax returns a boolean if a field has been set.

### GetStatus

`func (o *DpBlockSnapshotPolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpBlockSnapshotPolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpBlockSnapshotPolicy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpBlockSnapshotPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DpBlockSnapshotPolicy) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DpBlockSnapshotPolicy) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DpBlockSnapshotPolicy) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DpBlockSnapshotPolicy) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


