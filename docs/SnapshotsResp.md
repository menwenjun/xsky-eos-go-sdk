# SnapshotsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockSnapshots** | [**[]Snapshot**](Snapshot.md) | snapshots | 

## Methods

### NewSnapshotsResp

`func NewSnapshotsResp(blockSnapshots []Snapshot, ) *SnapshotsResp`

NewSnapshotsResp instantiates a new SnapshotsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotsRespWithDefaults

`func NewSnapshotsRespWithDefaults() *SnapshotsResp`

NewSnapshotsRespWithDefaults instantiates a new SnapshotsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockSnapshots

`func (o *SnapshotsResp) GetBlockSnapshots() []Snapshot`

GetBlockSnapshots returns the BlockSnapshots field if non-nil, zero value otherwise.

### GetBlockSnapshotsOk

`func (o *SnapshotsResp) GetBlockSnapshotsOk() (*[]Snapshot, bool)`

GetBlockSnapshotsOk returns a tuple with the BlockSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSnapshots

`func (o *SnapshotsResp) SetBlockSnapshots(v []Snapshot)`

SetBlockSnapshots sets BlockSnapshots field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


