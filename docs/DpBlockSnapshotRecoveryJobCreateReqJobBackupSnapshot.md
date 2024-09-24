# DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | **time.Time** |  | 
**VolumeSize** | **int64** |  | 

## Methods

### NewDpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot

`func NewDpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot(create time.Time, volumeSize int64, ) *DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot`

NewDpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot instantiates a new DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshotWithDefaults

`func NewDpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshotWithDefaults() *DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot`

NewDpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshotWithDefaults instantiates a new DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot) SetCreate(v time.Time)`

SetCreate sets Create field to given value.


### GetVolumeSize

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot) GetVolumeSize() int64`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot) GetVolumeSizeOk() (*int64, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot) SetVolumeSize(v int64)`

SetVolumeSize sets VolumeSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


