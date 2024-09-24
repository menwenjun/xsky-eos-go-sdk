# DfsSnapshotUpdateReqDfsSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description | [optional] 
**Lock** | Pointer to **bool** | lock snapshot or not | [optional] 
**Name** | Pointer to **string** | snapshot name | [optional] 
**Retention** | Pointer to **string** | retention time | [optional] 

## Methods

### NewDfsSnapshotUpdateReqDfsSnapshot

`func NewDfsSnapshotUpdateReqDfsSnapshot() *DfsSnapshotUpdateReqDfsSnapshot`

NewDfsSnapshotUpdateReqDfsSnapshot instantiates a new DfsSnapshotUpdateReqDfsSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSnapshotUpdateReqDfsSnapshotWithDefaults

`func NewDfsSnapshotUpdateReqDfsSnapshotWithDefaults() *DfsSnapshotUpdateReqDfsSnapshot`

NewDfsSnapshotUpdateReqDfsSnapshotWithDefaults instantiates a new DfsSnapshotUpdateReqDfsSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DfsSnapshotUpdateReqDfsSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsSnapshotUpdateReqDfsSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsSnapshotUpdateReqDfsSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsSnapshotUpdateReqDfsSnapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLock

`func (o *DfsSnapshotUpdateReqDfsSnapshot) GetLock() bool`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *DfsSnapshotUpdateReqDfsSnapshot) GetLockOk() (*bool, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *DfsSnapshotUpdateReqDfsSnapshot) SetLock(v bool)`

SetLock sets Lock field to given value.

### HasLock

`func (o *DfsSnapshotUpdateReqDfsSnapshot) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetName

`func (o *DfsSnapshotUpdateReqDfsSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsSnapshotUpdateReqDfsSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsSnapshotUpdateReqDfsSnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsSnapshotUpdateReqDfsSnapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetention

`func (o *DfsSnapshotUpdateReqDfsSnapshot) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *DfsSnapshotUpdateReqDfsSnapshot) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *DfsSnapshotUpdateReqDfsSnapshot) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *DfsSnapshotUpdateReqDfsSnapshot) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


