# DpDfsSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsPath** | Pointer to [**DfsPath**](DfsPath.md) |  | [optional] 
**DpDfsSnapshotPolicy** | Pointer to [**DpDfsSnapshotPolicyNestview**](DpDfsSnapshotPolicyNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDpDfsSnapshot

`func NewDpDfsSnapshot() *DpDfsSnapshot`

NewDpDfsSnapshot instantiates a new DpDfsSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpDfsSnapshotWithDefaults

`func NewDpDfsSnapshotWithDefaults() *DpDfsSnapshot`

NewDpDfsSnapshotWithDefaults instantiates a new DpDfsSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DpDfsSnapshot) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpDfsSnapshot) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpDfsSnapshot) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpDfsSnapshot) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DpDfsSnapshot) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DpDfsSnapshot) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DpDfsSnapshot) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DpDfsSnapshot) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsPath

`func (o *DpDfsSnapshot) GetDfsPath() DfsPath`

GetDfsPath returns the DfsPath field if non-nil, zero value otherwise.

### GetDfsPathOk

`func (o *DpDfsSnapshot) GetDfsPathOk() (*DfsPath, bool)`

GetDfsPathOk returns a tuple with the DfsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPath

`func (o *DpDfsSnapshot) SetDfsPath(v DfsPath)`

SetDfsPath sets DfsPath field to given value.

### HasDfsPath

`func (o *DpDfsSnapshot) HasDfsPath() bool`

HasDfsPath returns a boolean if a field has been set.

### GetDpDfsSnapshotPolicy

`func (o *DpDfsSnapshot) GetDpDfsSnapshotPolicy() DpDfsSnapshotPolicyNestview`

GetDpDfsSnapshotPolicy returns the DpDfsSnapshotPolicy field if non-nil, zero value otherwise.

### GetDpDfsSnapshotPolicyOk

`func (o *DpDfsSnapshot) GetDpDfsSnapshotPolicyOk() (*DpDfsSnapshotPolicyNestview, bool)`

GetDpDfsSnapshotPolicyOk returns a tuple with the DpDfsSnapshotPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpDfsSnapshotPolicy

`func (o *DpDfsSnapshot) SetDpDfsSnapshotPolicy(v DpDfsSnapshotPolicyNestview)`

SetDpDfsSnapshotPolicy sets DpDfsSnapshotPolicy field to given value.

### HasDpDfsSnapshotPolicy

`func (o *DpDfsSnapshot) HasDpDfsSnapshotPolicy() bool`

HasDpDfsSnapshotPolicy returns a boolean if a field has been set.

### GetId

`func (o *DpDfsSnapshot) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpDfsSnapshot) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpDfsSnapshot) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpDfsSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *DpDfsSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpDfsSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpDfsSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpDfsSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DpDfsSnapshot) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DpDfsSnapshot) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DpDfsSnapshot) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DpDfsSnapshot) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


