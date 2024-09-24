# DfsSnapshotCreateReqDfsSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | description | 
**Lock** | Pointer to **bool** | lock snapshot or not | [optional] 
**Name** | **string** | snapshot name | 
**Path** | **string** | path | 
**Retention** | **string** | retention time | 
**RootfsId** | **int64** | id of dfs rootfs | 

## Methods

### NewDfsSnapshotCreateReqDfsSnapshot

`func NewDfsSnapshotCreateReqDfsSnapshot(description string, name string, path string, retention string, rootfsId int64, ) *DfsSnapshotCreateReqDfsSnapshot`

NewDfsSnapshotCreateReqDfsSnapshot instantiates a new DfsSnapshotCreateReqDfsSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSnapshotCreateReqDfsSnapshotWithDefaults

`func NewDfsSnapshotCreateReqDfsSnapshotWithDefaults() *DfsSnapshotCreateReqDfsSnapshot`

NewDfsSnapshotCreateReqDfsSnapshotWithDefaults instantiates a new DfsSnapshotCreateReqDfsSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DfsSnapshotCreateReqDfsSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsSnapshotCreateReqDfsSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsSnapshotCreateReqDfsSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLock

`func (o *DfsSnapshotCreateReqDfsSnapshot) GetLock() bool`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *DfsSnapshotCreateReqDfsSnapshot) GetLockOk() (*bool, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *DfsSnapshotCreateReqDfsSnapshot) SetLock(v bool)`

SetLock sets Lock field to given value.

### HasLock

`func (o *DfsSnapshotCreateReqDfsSnapshot) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetName

`func (o *DfsSnapshotCreateReqDfsSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsSnapshotCreateReqDfsSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsSnapshotCreateReqDfsSnapshot) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *DfsSnapshotCreateReqDfsSnapshot) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsSnapshotCreateReqDfsSnapshot) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsSnapshotCreateReqDfsSnapshot) SetPath(v string)`

SetPath sets Path field to given value.


### GetRetention

`func (o *DfsSnapshotCreateReqDfsSnapshot) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *DfsSnapshotCreateReqDfsSnapshot) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *DfsSnapshotCreateReqDfsSnapshot) SetRetention(v string)`

SetRetention sets Retention field to given value.


### GetRootfsId

`func (o *DfsSnapshotCreateReqDfsSnapshot) GetRootfsId() int64`

GetRootfsId returns the RootfsId field if non-nil, zero value otherwise.

### GetRootfsIdOk

`func (o *DfsSnapshotCreateReqDfsSnapshot) GetRootfsIdOk() (*int64, bool)`

GetRootfsIdOk returns a tuple with the RootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfsId

`func (o *DfsSnapshotCreateReqDfsSnapshot) SetRootfsId(v int64)`

SetRootfsId sets RootfsId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


