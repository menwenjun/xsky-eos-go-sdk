# DfsFTPShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DfsGatewayGroup** | Pointer to [**DfsGatewayGroupNestview**](DfsGatewayGroupNestview.md) |  | [optional] 
**DfsPath** | Pointer to [**DfsPathNestview**](DfsPathNestview.md) |  | [optional] 
**DfsRootfs** | Pointer to [**DfsRootfsNestview**](DfsRootfsNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalSnapshotNum** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewDfsFTPShare

`func NewDfsFTPShare() *DfsFTPShare`

NewDfsFTPShare instantiates a new DfsFTPShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsFTPShareWithDefaults

`func NewDfsFTPShareWithDefaults() *DfsFTPShare`

NewDfsFTPShareWithDefaults instantiates a new DfsFTPShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsFTPShare) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsFTPShare) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsFTPShare) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsFTPShare) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *DfsFTPShare) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsFTPShare) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsFTPShare) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsFTPShare) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsFTPShare) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsFTPShare) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsFTPShare) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsFTPShare) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *DfsFTPShare) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsFTPShare) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsFTPShare) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsFTPShare) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsGatewayGroup

`func (o *DfsFTPShare) GetDfsGatewayGroup() DfsGatewayGroupNestview`

GetDfsGatewayGroup returns the DfsGatewayGroup field if non-nil, zero value otherwise.

### GetDfsGatewayGroupOk

`func (o *DfsFTPShare) GetDfsGatewayGroupOk() (*DfsGatewayGroupNestview, bool)`

GetDfsGatewayGroupOk returns a tuple with the DfsGatewayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroup

`func (o *DfsFTPShare) SetDfsGatewayGroup(v DfsGatewayGroupNestview)`

SetDfsGatewayGroup sets DfsGatewayGroup field to given value.

### HasDfsGatewayGroup

`func (o *DfsFTPShare) HasDfsGatewayGroup() bool`

HasDfsGatewayGroup returns a boolean if a field has been set.

### GetDfsPath

`func (o *DfsFTPShare) GetDfsPath() DfsPathNestview`

GetDfsPath returns the DfsPath field if non-nil, zero value otherwise.

### GetDfsPathOk

`func (o *DfsFTPShare) GetDfsPathOk() (*DfsPathNestview, bool)`

GetDfsPathOk returns a tuple with the DfsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPath

`func (o *DfsFTPShare) SetDfsPath(v DfsPathNestview)`

SetDfsPath sets DfsPath field to given value.

### HasDfsPath

`func (o *DfsFTPShare) HasDfsPath() bool`

HasDfsPath returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsFTPShare) GetDfsRootfs() DfsRootfsNestview`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsFTPShare) GetDfsRootfsOk() (*DfsRootfsNestview, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsFTPShare) SetDfsRootfs(v DfsRootfsNestview)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsFTPShare) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetId

`func (o *DfsFTPShare) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsFTPShare) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsFTPShare) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsFTPShare) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DfsFTPShare) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsFTPShare) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsFTPShare) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsFTPShare) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DfsFTPShare) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsFTPShare) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsFTPShare) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsFTPShare) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalSnapshotNum

`func (o *DfsFTPShare) GetTotalSnapshotNum() int64`

GetTotalSnapshotNum returns the TotalSnapshotNum field if non-nil, zero value otherwise.

### GetTotalSnapshotNumOk

`func (o *DfsFTPShare) GetTotalSnapshotNumOk() (*int64, bool)`

GetTotalSnapshotNumOk returns a tuple with the TotalSnapshotNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSnapshotNum

`func (o *DfsFTPShare) SetTotalSnapshotNum(v int64)`

SetTotalSnapshotNum sets TotalSnapshotNum field to given value.

### HasTotalSnapshotNum

`func (o *DfsFTPShare) HasTotalSnapshotNum() bool`

HasTotalSnapshotNum returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsFTPShare) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsFTPShare) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsFTPShare) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsFTPShare) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVersion

`func (o *DfsFTPShare) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DfsFTPShare) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DfsFTPShare) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DfsFTPShare) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


