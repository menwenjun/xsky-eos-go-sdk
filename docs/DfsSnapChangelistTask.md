# DfsSnapChangelistTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DstName** | Pointer to **string** |  | [optional] 
**DstSnapshot** | Pointer to [**DfsSnapshotNestview**](DfsSnapshotNestview.md) |  | [optional] 
**Gateway** | Pointer to [**DfsGatewayNestview**](DfsGatewayNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **[]string** |  | [optional] 
**SrcName** | Pointer to **string** |  | [optional] 
**SrcSnapshot** | Pointer to [**DfsSnapshotNestview**](DfsSnapshotNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsSnapChangelistTask

`func NewDfsSnapChangelistTask() *DfsSnapChangelistTask`

NewDfsSnapChangelistTask instantiates a new DfsSnapChangelistTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSnapChangelistTaskWithDefaults

`func NewDfsSnapChangelistTaskWithDefaults() *DfsSnapChangelistTask`

NewDfsSnapChangelistTaskWithDefaults instantiates a new DfsSnapChangelistTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DfsSnapChangelistTask) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsSnapChangelistTask) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsSnapChangelistTask) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsSnapChangelistTask) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsSnapChangelistTask) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsSnapChangelistTask) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsSnapChangelistTask) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsSnapChangelistTask) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDstName

`func (o *DfsSnapChangelistTask) GetDstName() string`

GetDstName returns the DstName field if non-nil, zero value otherwise.

### GetDstNameOk

`func (o *DfsSnapChangelistTask) GetDstNameOk() (*string, bool)`

GetDstNameOk returns a tuple with the DstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstName

`func (o *DfsSnapChangelistTask) SetDstName(v string)`

SetDstName sets DstName field to given value.

### HasDstName

`func (o *DfsSnapChangelistTask) HasDstName() bool`

HasDstName returns a boolean if a field has been set.

### GetDstSnapshot

`func (o *DfsSnapChangelistTask) GetDstSnapshot() DfsSnapshotNestview`

GetDstSnapshot returns the DstSnapshot field if non-nil, zero value otherwise.

### GetDstSnapshotOk

`func (o *DfsSnapChangelistTask) GetDstSnapshotOk() (*DfsSnapshotNestview, bool)`

GetDstSnapshotOk returns a tuple with the DstSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstSnapshot

`func (o *DfsSnapChangelistTask) SetDstSnapshot(v DfsSnapshotNestview)`

SetDstSnapshot sets DstSnapshot field to given value.

### HasDstSnapshot

`func (o *DfsSnapChangelistTask) HasDstSnapshot() bool`

HasDstSnapshot returns a boolean if a field has been set.

### GetGateway

`func (o *DfsSnapChangelistTask) GetGateway() DfsGatewayNestview`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *DfsSnapChangelistTask) GetGatewayOk() (*DfsGatewayNestview, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *DfsSnapChangelistTask) SetGateway(v DfsGatewayNestview)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *DfsSnapChangelistTask) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetId

`func (o *DfsSnapChangelistTask) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsSnapChangelistTask) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsSnapChangelistTask) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsSnapChangelistTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *DfsSnapChangelistTask) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsSnapChangelistTask) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsSnapChangelistTask) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DfsSnapChangelistTask) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetResult

`func (o *DfsSnapChangelistTask) GetResult() []string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DfsSnapChangelistTask) GetResultOk() (*[]string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DfsSnapChangelistTask) SetResult(v []string)`

SetResult sets Result field to given value.

### HasResult

`func (o *DfsSnapChangelistTask) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSrcName

`func (o *DfsSnapChangelistTask) GetSrcName() string`

GetSrcName returns the SrcName field if non-nil, zero value otherwise.

### GetSrcNameOk

`func (o *DfsSnapChangelistTask) GetSrcNameOk() (*string, bool)`

GetSrcNameOk returns a tuple with the SrcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcName

`func (o *DfsSnapChangelistTask) SetSrcName(v string)`

SetSrcName sets SrcName field to given value.

### HasSrcName

`func (o *DfsSnapChangelistTask) HasSrcName() bool`

HasSrcName returns a boolean if a field has been set.

### GetSrcSnapshot

`func (o *DfsSnapChangelistTask) GetSrcSnapshot() DfsSnapshotNestview`

GetSrcSnapshot returns the SrcSnapshot field if non-nil, zero value otherwise.

### GetSrcSnapshotOk

`func (o *DfsSnapChangelistTask) GetSrcSnapshotOk() (*DfsSnapshotNestview, bool)`

GetSrcSnapshotOk returns a tuple with the SrcSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcSnapshot

`func (o *DfsSnapChangelistTask) SetSrcSnapshot(v DfsSnapshotNestview)`

SetSrcSnapshot sets SrcSnapshot field to given value.

### HasSrcSnapshot

`func (o *DfsSnapChangelistTask) HasSrcSnapshot() bool`

HasSrcSnapshot returns a boolean if a field has been set.

### GetStatus

`func (o *DfsSnapChangelistTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsSnapChangelistTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsSnapChangelistTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsSnapChangelistTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsSnapChangelistTask) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsSnapChangelistTask) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsSnapChangelistTask) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsSnapChangelistTask) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


