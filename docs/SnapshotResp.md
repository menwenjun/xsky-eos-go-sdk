# SnapshotResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockSnapshot** | [**Snapshot**](Snapshot.md) |  | 

## Methods

### NewSnapshotResp

`func NewSnapshotResp(blockSnapshot Snapshot, ) *SnapshotResp`

NewSnapshotResp instantiates a new SnapshotResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotRespWithDefaults

`func NewSnapshotRespWithDefaults() *SnapshotResp`

NewSnapshotRespWithDefaults instantiates a new SnapshotResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockSnapshot

`func (o *SnapshotResp) GetBlockSnapshot() Snapshot`

GetBlockSnapshot returns the BlockSnapshot field if non-nil, zero value otherwise.

### GetBlockSnapshotOk

`func (o *SnapshotResp) GetBlockSnapshotOk() (*Snapshot, bool)`

GetBlockSnapshotOk returns a tuple with the BlockSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSnapshot

`func (o *SnapshotResp) SetBlockSnapshot(v Snapshot)`

SetBlockSnapshot sets BlockSnapshot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


