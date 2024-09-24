# DfsSnapChangelistTaskCreateReqSnapChangelistTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DstSnapshotId** | **int64** | destination snapshot id | 
**Path** | **string** | snap change list task root path | 
**SrcSnapshotId** | Pointer to **int64** | source snapshot id | [optional] 

## Methods

### NewDfsSnapChangelistTaskCreateReqSnapChangelistTask

`func NewDfsSnapChangelistTaskCreateReqSnapChangelistTask(dstSnapshotId int64, path string, ) *DfsSnapChangelistTaskCreateReqSnapChangelistTask`

NewDfsSnapChangelistTaskCreateReqSnapChangelistTask instantiates a new DfsSnapChangelistTaskCreateReqSnapChangelistTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSnapChangelistTaskCreateReqSnapChangelistTaskWithDefaults

`func NewDfsSnapChangelistTaskCreateReqSnapChangelistTaskWithDefaults() *DfsSnapChangelistTaskCreateReqSnapChangelistTask`

NewDfsSnapChangelistTaskCreateReqSnapChangelistTaskWithDefaults instantiates a new DfsSnapChangelistTaskCreateReqSnapChangelistTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDstSnapshotId

`func (o *DfsSnapChangelistTaskCreateReqSnapChangelistTask) GetDstSnapshotId() int64`

GetDstSnapshotId returns the DstSnapshotId field if non-nil, zero value otherwise.

### GetDstSnapshotIdOk

`func (o *DfsSnapChangelistTaskCreateReqSnapChangelistTask) GetDstSnapshotIdOk() (*int64, bool)`

GetDstSnapshotIdOk returns a tuple with the DstSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstSnapshotId

`func (o *DfsSnapChangelistTaskCreateReqSnapChangelistTask) SetDstSnapshotId(v int64)`

SetDstSnapshotId sets DstSnapshotId field to given value.


### GetPath

`func (o *DfsSnapChangelistTaskCreateReqSnapChangelistTask) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsSnapChangelistTaskCreateReqSnapChangelistTask) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsSnapChangelistTaskCreateReqSnapChangelistTask) SetPath(v string)`

SetPath sets Path field to given value.


### GetSrcSnapshotId

`func (o *DfsSnapChangelistTaskCreateReqSnapChangelistTask) GetSrcSnapshotId() int64`

GetSrcSnapshotId returns the SrcSnapshotId field if non-nil, zero value otherwise.

### GetSrcSnapshotIdOk

`func (o *DfsSnapChangelistTaskCreateReqSnapChangelistTask) GetSrcSnapshotIdOk() (*int64, bool)`

GetSrcSnapshotIdOk returns a tuple with the SrcSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcSnapshotId

`func (o *DfsSnapChangelistTaskCreateReqSnapChangelistTask) SetSrcSnapshotId(v int64)`

SetSrcSnapshotId sets SrcSnapshotId field to given value.

### HasSrcSnapshotId

`func (o *DfsSnapChangelistTaskCreateReqSnapChangelistTask) HasSrcSnapshotId() bool`

HasSrcSnapshotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


