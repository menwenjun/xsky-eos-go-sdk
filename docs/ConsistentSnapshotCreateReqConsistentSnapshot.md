# ConsistentSnapshotCreateReqConsistentSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockVolumeIds** | **[]int64** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SnapUids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewConsistentSnapshotCreateReqConsistentSnapshot

`func NewConsistentSnapshotCreateReqConsistentSnapshot(blockVolumeIds []int64, ) *ConsistentSnapshotCreateReqConsistentSnapshot`

NewConsistentSnapshotCreateReqConsistentSnapshot instantiates a new ConsistentSnapshotCreateReqConsistentSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsistentSnapshotCreateReqConsistentSnapshotWithDefaults

`func NewConsistentSnapshotCreateReqConsistentSnapshotWithDefaults() *ConsistentSnapshotCreateReqConsistentSnapshot`

NewConsistentSnapshotCreateReqConsistentSnapshotWithDefaults instantiates a new ConsistentSnapshotCreateReqConsistentSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockVolumeIds

`func (o *ConsistentSnapshotCreateReqConsistentSnapshot) GetBlockVolumeIds() []int64`

GetBlockVolumeIds returns the BlockVolumeIds field if non-nil, zero value otherwise.

### GetBlockVolumeIdsOk

`func (o *ConsistentSnapshotCreateReqConsistentSnapshot) GetBlockVolumeIdsOk() (*[]int64, bool)`

GetBlockVolumeIdsOk returns a tuple with the BlockVolumeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeIds

`func (o *ConsistentSnapshotCreateReqConsistentSnapshot) SetBlockVolumeIds(v []int64)`

SetBlockVolumeIds sets BlockVolumeIds field to given value.


### GetDescription

`func (o *ConsistentSnapshotCreateReqConsistentSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsistentSnapshotCreateReqConsistentSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsistentSnapshotCreateReqConsistentSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsistentSnapshotCreateReqConsistentSnapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ConsistentSnapshotCreateReqConsistentSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsistentSnapshotCreateReqConsistentSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsistentSnapshotCreateReqConsistentSnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsistentSnapshotCreateReqConsistentSnapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSnapUids

`func (o *ConsistentSnapshotCreateReqConsistentSnapshot) GetSnapUids() []string`

GetSnapUids returns the SnapUids field if non-nil, zero value otherwise.

### GetSnapUidsOk

`func (o *ConsistentSnapshotCreateReqConsistentSnapshot) GetSnapUidsOk() (*[]string, bool)`

GetSnapUidsOk returns a tuple with the SnapUids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapUids

`func (o *ConsistentSnapshotCreateReqConsistentSnapshot) SetSnapUids(v []string)`

SetSnapUids sets SnapUids field to given value.

### HasSnapUids

`func (o *ConsistentSnapshotCreateReqConsistentSnapshot) HasSnapUids() bool`

HasSnapUids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


