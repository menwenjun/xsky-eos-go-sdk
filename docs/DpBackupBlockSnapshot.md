# DpBackupBlockSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupType** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**VolumeSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewDpBackupBlockSnapshot

`func NewDpBackupBlockSnapshot() *DpBackupBlockSnapshot`

NewDpBackupBlockSnapshot instantiates a new DpBackupBlockSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBackupBlockSnapshotWithDefaults

`func NewDpBackupBlockSnapshotWithDefaults() *DpBackupBlockSnapshot`

NewDpBackupBlockSnapshotWithDefaults instantiates a new DpBackupBlockSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupType

`func (o *DpBackupBlockSnapshot) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *DpBackupBlockSnapshot) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *DpBackupBlockSnapshot) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *DpBackupBlockSnapshot) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetCreate

`func (o *DpBackupBlockSnapshot) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DpBackupBlockSnapshot) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DpBackupBlockSnapshot) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DpBackupBlockSnapshot) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetSize

`func (o *DpBackupBlockSnapshot) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DpBackupBlockSnapshot) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DpBackupBlockSnapshot) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DpBackupBlockSnapshot) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVolumeSize

`func (o *DpBackupBlockSnapshot) GetVolumeSize() int64`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *DpBackupBlockSnapshot) GetVolumeSizeOk() (*int64, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *DpBackupBlockSnapshot) SetVolumeSize(v int64)`

SetVolumeSize sets VolumeSize field to given value.

### HasVolumeSize

`func (o *DpBackupBlockSnapshot) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


