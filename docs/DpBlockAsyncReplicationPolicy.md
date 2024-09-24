# DpBlockAsyncReplicationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockVolumeNum** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**CronExpr** | Pointer to **string** |  | [optional] 
**DpGateway** | Pointer to [**DpGatewayNestview**](DpGatewayNestview.md) |  | [optional] 
**DpSite** | Pointer to [**DpSiteNestview**](DpSiteNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RetainedMax** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDpBlockAsyncReplicationPolicy

`func NewDpBlockAsyncReplicationPolicy() *DpBlockAsyncReplicationPolicy`

NewDpBlockAsyncReplicationPolicy instantiates a new DpBlockAsyncReplicationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockAsyncReplicationPolicyWithDefaults

`func NewDpBlockAsyncReplicationPolicyWithDefaults() *DpBlockAsyncReplicationPolicy`

NewDpBlockAsyncReplicationPolicyWithDefaults instantiates a new DpBlockAsyncReplicationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockVolumeNum

`func (o *DpBlockAsyncReplicationPolicy) GetBlockVolumeNum() int64`

GetBlockVolumeNum returns the BlockVolumeNum field if non-nil, zero value otherwise.

### GetBlockVolumeNumOk

`func (o *DpBlockAsyncReplicationPolicy) GetBlockVolumeNumOk() (*int64, bool)`

GetBlockVolumeNumOk returns a tuple with the BlockVolumeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeNum

`func (o *DpBlockAsyncReplicationPolicy) SetBlockVolumeNum(v int64)`

SetBlockVolumeNum sets BlockVolumeNum field to given value.

### HasBlockVolumeNum

`func (o *DpBlockAsyncReplicationPolicy) HasBlockVolumeNum() bool`

HasBlockVolumeNum returns a boolean if a field has been set.

### GetCluster

`func (o *DpBlockAsyncReplicationPolicy) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpBlockAsyncReplicationPolicy) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpBlockAsyncReplicationPolicy) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpBlockAsyncReplicationPolicy) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DpBlockAsyncReplicationPolicy) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DpBlockAsyncReplicationPolicy) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DpBlockAsyncReplicationPolicy) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DpBlockAsyncReplicationPolicy) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCronExpr

`func (o *DpBlockAsyncReplicationPolicy) GetCronExpr() string`

GetCronExpr returns the CronExpr field if non-nil, zero value otherwise.

### GetCronExprOk

`func (o *DpBlockAsyncReplicationPolicy) GetCronExprOk() (*string, bool)`

GetCronExprOk returns a tuple with the CronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpr

`func (o *DpBlockAsyncReplicationPolicy) SetCronExpr(v string)`

SetCronExpr sets CronExpr field to given value.

### HasCronExpr

`func (o *DpBlockAsyncReplicationPolicy) HasCronExpr() bool`

HasCronExpr returns a boolean if a field has been set.

### GetDpGateway

`func (o *DpBlockAsyncReplicationPolicy) GetDpGateway() DpGatewayNestview`

GetDpGateway returns the DpGateway field if non-nil, zero value otherwise.

### GetDpGatewayOk

`func (o *DpBlockAsyncReplicationPolicy) GetDpGatewayOk() (*DpGatewayNestview, bool)`

GetDpGatewayOk returns a tuple with the DpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpGateway

`func (o *DpBlockAsyncReplicationPolicy) SetDpGateway(v DpGatewayNestview)`

SetDpGateway sets DpGateway field to given value.

### HasDpGateway

`func (o *DpBlockAsyncReplicationPolicy) HasDpGateway() bool`

HasDpGateway returns a boolean if a field has been set.

### GetDpSite

`func (o *DpBlockAsyncReplicationPolicy) GetDpSite() DpSiteNestview`

GetDpSite returns the DpSite field if non-nil, zero value otherwise.

### GetDpSiteOk

`func (o *DpBlockAsyncReplicationPolicy) GetDpSiteOk() (*DpSiteNestview, bool)`

GetDpSiteOk returns a tuple with the DpSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpSite

`func (o *DpBlockAsyncReplicationPolicy) SetDpSite(v DpSiteNestview)`

SetDpSite sets DpSite field to given value.

### HasDpSite

`func (o *DpBlockAsyncReplicationPolicy) HasDpSite() bool`

HasDpSite returns a boolean if a field has been set.

### GetId

`func (o *DpBlockAsyncReplicationPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpBlockAsyncReplicationPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpBlockAsyncReplicationPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpBlockAsyncReplicationPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DpBlockAsyncReplicationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpBlockAsyncReplicationPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpBlockAsyncReplicationPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DpBlockAsyncReplicationPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetainedMax

`func (o *DpBlockAsyncReplicationPolicy) GetRetainedMax() int64`

GetRetainedMax returns the RetainedMax field if non-nil, zero value otherwise.

### GetRetainedMaxOk

`func (o *DpBlockAsyncReplicationPolicy) GetRetainedMaxOk() (*int64, bool)`

GetRetainedMaxOk returns a tuple with the RetainedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedMax

`func (o *DpBlockAsyncReplicationPolicy) SetRetainedMax(v int64)`

SetRetainedMax sets RetainedMax field to given value.

### HasRetainedMax

`func (o *DpBlockAsyncReplicationPolicy) HasRetainedMax() bool`

HasRetainedMax returns a boolean if a field has been set.

### GetStatus

`func (o *DpBlockAsyncReplicationPolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpBlockAsyncReplicationPolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpBlockAsyncReplicationPolicy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpBlockAsyncReplicationPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DpBlockAsyncReplicationPolicy) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DpBlockAsyncReplicationPolicy) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DpBlockAsyncReplicationPolicy) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DpBlockAsyncReplicationPolicy) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


