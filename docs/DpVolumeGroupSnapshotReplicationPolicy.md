# DpVolumeGroupSnapshotReplicationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockVolumeGroupNum** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
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

### NewDpVolumeGroupSnapshotReplicationPolicy

`func NewDpVolumeGroupSnapshotReplicationPolicy() *DpVolumeGroupSnapshotReplicationPolicy`

NewDpVolumeGroupSnapshotReplicationPolicy instantiates a new DpVolumeGroupSnapshotReplicationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpVolumeGroupSnapshotReplicationPolicyWithDefaults

`func NewDpVolumeGroupSnapshotReplicationPolicyWithDefaults() *DpVolumeGroupSnapshotReplicationPolicy`

NewDpVolumeGroupSnapshotReplicationPolicyWithDefaults instantiates a new DpVolumeGroupSnapshotReplicationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockVolumeGroupNum

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetBlockVolumeGroupNum() int64`

GetBlockVolumeGroupNum returns the BlockVolumeGroupNum field if non-nil, zero value otherwise.

### GetBlockVolumeGroupNumOk

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetBlockVolumeGroupNumOk() (*int64, bool)`

GetBlockVolumeGroupNumOk returns a tuple with the BlockVolumeGroupNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeGroupNum

`func (o *DpVolumeGroupSnapshotReplicationPolicy) SetBlockVolumeGroupNum(v int64)`

SetBlockVolumeGroupNum sets BlockVolumeGroupNum field to given value.

### HasBlockVolumeGroupNum

`func (o *DpVolumeGroupSnapshotReplicationPolicy) HasBlockVolumeGroupNum() bool`

HasBlockVolumeGroupNum returns a boolean if a field has been set.

### GetCluster

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpVolumeGroupSnapshotReplicationPolicy) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpVolumeGroupSnapshotReplicationPolicy) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DpVolumeGroupSnapshotReplicationPolicy) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DpVolumeGroupSnapshotReplicationPolicy) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCronExpr

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetCronExpr() string`

GetCronExpr returns the CronExpr field if non-nil, zero value otherwise.

### GetCronExprOk

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetCronExprOk() (*string, bool)`

GetCronExprOk returns a tuple with the CronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpr

`func (o *DpVolumeGroupSnapshotReplicationPolicy) SetCronExpr(v string)`

SetCronExpr sets CronExpr field to given value.

### HasCronExpr

`func (o *DpVolumeGroupSnapshotReplicationPolicy) HasCronExpr() bool`

HasCronExpr returns a boolean if a field has been set.

### GetDpGateway

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetDpGateway() DpGatewayNestview`

GetDpGateway returns the DpGateway field if non-nil, zero value otherwise.

### GetDpGatewayOk

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetDpGatewayOk() (*DpGatewayNestview, bool)`

GetDpGatewayOk returns a tuple with the DpGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpGateway

`func (o *DpVolumeGroupSnapshotReplicationPolicy) SetDpGateway(v DpGatewayNestview)`

SetDpGateway sets DpGateway field to given value.

### HasDpGateway

`func (o *DpVolumeGroupSnapshotReplicationPolicy) HasDpGateway() bool`

HasDpGateway returns a boolean if a field has been set.

### GetDpSite

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetDpSite() DpSiteNestview`

GetDpSite returns the DpSite field if non-nil, zero value otherwise.

### GetDpSiteOk

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetDpSiteOk() (*DpSiteNestview, bool)`

GetDpSiteOk returns a tuple with the DpSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpSite

`func (o *DpVolumeGroupSnapshotReplicationPolicy) SetDpSite(v DpSiteNestview)`

SetDpSite sets DpSite field to given value.

### HasDpSite

`func (o *DpVolumeGroupSnapshotReplicationPolicy) HasDpSite() bool`

HasDpSite returns a boolean if a field has been set.

### GetId

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpVolumeGroupSnapshotReplicationPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpVolumeGroupSnapshotReplicationPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpVolumeGroupSnapshotReplicationPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DpVolumeGroupSnapshotReplicationPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetainedMax

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetRetainedMax() int64`

GetRetainedMax returns the RetainedMax field if non-nil, zero value otherwise.

### GetRetainedMaxOk

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetRetainedMaxOk() (*int64, bool)`

GetRetainedMaxOk returns a tuple with the RetainedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedMax

`func (o *DpVolumeGroupSnapshotReplicationPolicy) SetRetainedMax(v int64)`

SetRetainedMax sets RetainedMax field to given value.

### HasRetainedMax

`func (o *DpVolumeGroupSnapshotReplicationPolicy) HasRetainedMax() bool`

HasRetainedMax returns a boolean if a field has been set.

### GetStatus

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpVolumeGroupSnapshotReplicationPolicy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpVolumeGroupSnapshotReplicationPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DpVolumeGroupSnapshotReplicationPolicy) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DpVolumeGroupSnapshotReplicationPolicy) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DpVolumeGroupSnapshotReplicationPolicy) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


