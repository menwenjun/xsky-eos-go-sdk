# DfsNFSShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**CounterDecreased** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DfsGatewayGroup** | Pointer to [**DfsGatewayGroupNestview**](DfsGatewayGroupNestview.md) |  | [optional] 
**DfsPath** | Pointer to [**DfsPathNestview**](DfsPathNestview.md) |  | [optional] 
**DfsRootfs** | Pointer to [**DfsRootfsNestview**](DfsRootfsNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NfsVersions** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalSnapshotNum** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewDfsNFSShare

`func NewDfsNFSShare() *DfsNFSShare`

NewDfsNFSShare instantiates a new DfsNFSShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsNFSShareWithDefaults

`func NewDfsNFSShareWithDefaults() *DfsNFSShare`

NewDfsNFSShareWithDefaults instantiates a new DfsNFSShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsNFSShare) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsNFSShare) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsNFSShare) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsNFSShare) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *DfsNFSShare) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsNFSShare) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsNFSShare) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsNFSShare) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCounterDecreased

`func (o *DfsNFSShare) GetCounterDecreased() bool`

GetCounterDecreased returns the CounterDecreased field if non-nil, zero value otherwise.

### GetCounterDecreasedOk

`func (o *DfsNFSShare) GetCounterDecreasedOk() (*bool, bool)`

GetCounterDecreasedOk returns a tuple with the CounterDecreased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterDecreased

`func (o *DfsNFSShare) SetCounterDecreased(v bool)`

SetCounterDecreased sets CounterDecreased field to given value.

### HasCounterDecreased

`func (o *DfsNFSShare) HasCounterDecreased() bool`

HasCounterDecreased returns a boolean if a field has been set.

### GetCreate

`func (o *DfsNFSShare) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsNFSShare) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsNFSShare) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsNFSShare) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *DfsNFSShare) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsNFSShare) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsNFSShare) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsNFSShare) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsGatewayGroup

`func (o *DfsNFSShare) GetDfsGatewayGroup() DfsGatewayGroupNestview`

GetDfsGatewayGroup returns the DfsGatewayGroup field if non-nil, zero value otherwise.

### GetDfsGatewayGroupOk

`func (o *DfsNFSShare) GetDfsGatewayGroupOk() (*DfsGatewayGroupNestview, bool)`

GetDfsGatewayGroupOk returns a tuple with the DfsGatewayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroup

`func (o *DfsNFSShare) SetDfsGatewayGroup(v DfsGatewayGroupNestview)`

SetDfsGatewayGroup sets DfsGatewayGroup field to given value.

### HasDfsGatewayGroup

`func (o *DfsNFSShare) HasDfsGatewayGroup() bool`

HasDfsGatewayGroup returns a boolean if a field has been set.

### GetDfsPath

`func (o *DfsNFSShare) GetDfsPath() DfsPathNestview`

GetDfsPath returns the DfsPath field if non-nil, zero value otherwise.

### GetDfsPathOk

`func (o *DfsNFSShare) GetDfsPathOk() (*DfsPathNestview, bool)`

GetDfsPathOk returns a tuple with the DfsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPath

`func (o *DfsNFSShare) SetDfsPath(v DfsPathNestview)`

SetDfsPath sets DfsPath field to given value.

### HasDfsPath

`func (o *DfsNFSShare) HasDfsPath() bool`

HasDfsPath returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsNFSShare) GetDfsRootfs() DfsRootfsNestview`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsNFSShare) GetDfsRootfsOk() (*DfsRootfsNestview, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsNFSShare) SetDfsRootfs(v DfsRootfsNestview)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsNFSShare) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetId

`func (o *DfsNFSShare) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsNFSShare) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsNFSShare) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsNFSShare) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DfsNFSShare) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsNFSShare) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsNFSShare) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsNFSShare) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNfsVersions

`func (o *DfsNFSShare) GetNfsVersions() []string`

GetNfsVersions returns the NfsVersions field if non-nil, zero value otherwise.

### GetNfsVersionsOk

`func (o *DfsNFSShare) GetNfsVersionsOk() (*[]string, bool)`

GetNfsVersionsOk returns a tuple with the NfsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsVersions

`func (o *DfsNFSShare) SetNfsVersions(v []string)`

SetNfsVersions sets NfsVersions field to given value.

### HasNfsVersions

`func (o *DfsNFSShare) HasNfsVersions() bool`

HasNfsVersions returns a boolean if a field has been set.

### GetStatus

`func (o *DfsNFSShare) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsNFSShare) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsNFSShare) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsNFSShare) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalSnapshotNum

`func (o *DfsNFSShare) GetTotalSnapshotNum() int64`

GetTotalSnapshotNum returns the TotalSnapshotNum field if non-nil, zero value otherwise.

### GetTotalSnapshotNumOk

`func (o *DfsNFSShare) GetTotalSnapshotNumOk() (*int64, bool)`

GetTotalSnapshotNumOk returns a tuple with the TotalSnapshotNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSnapshotNum

`func (o *DfsNFSShare) SetTotalSnapshotNum(v int64)`

SetTotalSnapshotNum sets TotalSnapshotNum field to given value.

### HasTotalSnapshotNum

`func (o *DfsNFSShare) HasTotalSnapshotNum() bool`

HasTotalSnapshotNum returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsNFSShare) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsNFSShare) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsNFSShare) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsNFSShare) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVersion

`func (o *DfsNFSShare) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DfsNFSShare) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DfsNFSShare) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DfsNFSShare) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


