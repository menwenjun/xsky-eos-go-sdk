# SnapshotCreateReqSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockVolumeId** | **int64** |  | 
**Description** | **string** |  | 
**Name** | **string** |  | 
**Uid** | Pointer to **string** |  | [optional] 

## Methods

### NewSnapshotCreateReqSnapshot

`func NewSnapshotCreateReqSnapshot(blockVolumeId int64, description string, name string, ) *SnapshotCreateReqSnapshot`

NewSnapshotCreateReqSnapshot instantiates a new SnapshotCreateReqSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotCreateReqSnapshotWithDefaults

`func NewSnapshotCreateReqSnapshotWithDefaults() *SnapshotCreateReqSnapshot`

NewSnapshotCreateReqSnapshotWithDefaults instantiates a new SnapshotCreateReqSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockVolumeId

`func (o *SnapshotCreateReqSnapshot) GetBlockVolumeId() int64`

GetBlockVolumeId returns the BlockVolumeId field if non-nil, zero value otherwise.

### GetBlockVolumeIdOk

`func (o *SnapshotCreateReqSnapshot) GetBlockVolumeIdOk() (*int64, bool)`

GetBlockVolumeIdOk returns a tuple with the BlockVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeId

`func (o *SnapshotCreateReqSnapshot) SetBlockVolumeId(v int64)`

SetBlockVolumeId sets BlockVolumeId field to given value.


### GetDescription

`func (o *SnapshotCreateReqSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnapshotCreateReqSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnapshotCreateReqSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *SnapshotCreateReqSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnapshotCreateReqSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnapshotCreateReqSnapshot) SetName(v string)`

SetName sets Name field to given value.


### GetUid

`func (o *SnapshotCreateReqSnapshot) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SnapshotCreateReqSnapshot) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SnapshotCreateReqSnapshot) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SnapshotCreateReqSnapshot) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


