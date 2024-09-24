# VolumeStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedTask** | Pointer to [**VolumeFailedTask**](VolumeFailedTask.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**IoSize** | Pointer to [**VolumeIOSize**](VolumeIOSize.md) |  | [optional] 
**IoSizeLat** | Pointer to [**VolumeIOLatency**](VolumeIOLatency.md) |  | [optional] 
**MigrateWriteBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**MigrateWriteIops** | Pointer to **float64** |  | [optional] 
**MigrateWriteLatencyUs** | Pointer to **float64** |  | [optional] 
**NonIoTask** | Pointer to [**VolumeNonIOTask**](VolumeNonIOTask.md) |  | [optional] 
**QueueDepth** | Pointer to **float64** |  | [optional] 
**ReadBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**ReadIops** | Pointer to **float64** |  | [optional] 
**ReadLatencyUs** | Pointer to **float64** |  | [optional] 
**TotalBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**TotalIops** | Pointer to **float64** |  | [optional] 
**WriteBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**WriteIops** | Pointer to **float64** |  | [optional] 
**WriteLatencyUs** | Pointer to **float64** |  | [optional] 

## Methods

### NewVolumeStat

`func NewVolumeStat() *VolumeStat`

NewVolumeStat instantiates a new VolumeStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeStatWithDefaults

`func NewVolumeStatWithDefaults() *VolumeStat`

NewVolumeStatWithDefaults instantiates a new VolumeStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedTask

`func (o *VolumeStat) GetFailedTask() VolumeFailedTask`

GetFailedTask returns the FailedTask field if non-nil, zero value otherwise.

### GetFailedTaskOk

`func (o *VolumeStat) GetFailedTaskOk() (*VolumeFailedTask, bool)`

GetFailedTaskOk returns a tuple with the FailedTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedTask

`func (o *VolumeStat) SetFailedTask(v VolumeFailedTask)`

SetFailedTask sets FailedTask field to given value.

### HasFailedTask

`func (o *VolumeStat) HasFailedTask() bool`

HasFailedTask returns a boolean if a field has been set.

### GetCreate

`func (o *VolumeStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *VolumeStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *VolumeStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *VolumeStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetIoSize

`func (o *VolumeStat) GetIoSize() VolumeIOSize`

GetIoSize returns the IoSize field if non-nil, zero value otherwise.

### GetIoSizeOk

`func (o *VolumeStat) GetIoSizeOk() (*VolumeIOSize, bool)`

GetIoSizeOk returns a tuple with the IoSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoSize

`func (o *VolumeStat) SetIoSize(v VolumeIOSize)`

SetIoSize sets IoSize field to given value.

### HasIoSize

`func (o *VolumeStat) HasIoSize() bool`

HasIoSize returns a boolean if a field has been set.

### GetIoSizeLat

`func (o *VolumeStat) GetIoSizeLat() VolumeIOLatency`

GetIoSizeLat returns the IoSizeLat field if non-nil, zero value otherwise.

### GetIoSizeLatOk

`func (o *VolumeStat) GetIoSizeLatOk() (*VolumeIOLatency, bool)`

GetIoSizeLatOk returns a tuple with the IoSizeLat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoSizeLat

`func (o *VolumeStat) SetIoSizeLat(v VolumeIOLatency)`

SetIoSizeLat sets IoSizeLat field to given value.

### HasIoSizeLat

`func (o *VolumeStat) HasIoSizeLat() bool`

HasIoSizeLat returns a boolean if a field has been set.

### GetMigrateWriteBandwidthKbyte

`func (o *VolumeStat) GetMigrateWriteBandwidthKbyte() float64`

GetMigrateWriteBandwidthKbyte returns the MigrateWriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetMigrateWriteBandwidthKbyteOk

`func (o *VolumeStat) GetMigrateWriteBandwidthKbyteOk() (*float64, bool)`

GetMigrateWriteBandwidthKbyteOk returns a tuple with the MigrateWriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateWriteBandwidthKbyte

`func (o *VolumeStat) SetMigrateWriteBandwidthKbyte(v float64)`

SetMigrateWriteBandwidthKbyte sets MigrateWriteBandwidthKbyte field to given value.

### HasMigrateWriteBandwidthKbyte

`func (o *VolumeStat) HasMigrateWriteBandwidthKbyte() bool`

HasMigrateWriteBandwidthKbyte returns a boolean if a field has been set.

### GetMigrateWriteIops

`func (o *VolumeStat) GetMigrateWriteIops() float64`

GetMigrateWriteIops returns the MigrateWriteIops field if non-nil, zero value otherwise.

### GetMigrateWriteIopsOk

`func (o *VolumeStat) GetMigrateWriteIopsOk() (*float64, bool)`

GetMigrateWriteIopsOk returns a tuple with the MigrateWriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateWriteIops

`func (o *VolumeStat) SetMigrateWriteIops(v float64)`

SetMigrateWriteIops sets MigrateWriteIops field to given value.

### HasMigrateWriteIops

`func (o *VolumeStat) HasMigrateWriteIops() bool`

HasMigrateWriteIops returns a boolean if a field has been set.

### GetMigrateWriteLatencyUs

`func (o *VolumeStat) GetMigrateWriteLatencyUs() float64`

GetMigrateWriteLatencyUs returns the MigrateWriteLatencyUs field if non-nil, zero value otherwise.

### GetMigrateWriteLatencyUsOk

`func (o *VolumeStat) GetMigrateWriteLatencyUsOk() (*float64, bool)`

GetMigrateWriteLatencyUsOk returns a tuple with the MigrateWriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateWriteLatencyUs

`func (o *VolumeStat) SetMigrateWriteLatencyUs(v float64)`

SetMigrateWriteLatencyUs sets MigrateWriteLatencyUs field to given value.

### HasMigrateWriteLatencyUs

`func (o *VolumeStat) HasMigrateWriteLatencyUs() bool`

HasMigrateWriteLatencyUs returns a boolean if a field has been set.

### GetNonIoTask

`func (o *VolumeStat) GetNonIoTask() VolumeNonIOTask`

GetNonIoTask returns the NonIoTask field if non-nil, zero value otherwise.

### GetNonIoTaskOk

`func (o *VolumeStat) GetNonIoTaskOk() (*VolumeNonIOTask, bool)`

GetNonIoTaskOk returns a tuple with the NonIoTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonIoTask

`func (o *VolumeStat) SetNonIoTask(v VolumeNonIOTask)`

SetNonIoTask sets NonIoTask field to given value.

### HasNonIoTask

`func (o *VolumeStat) HasNonIoTask() bool`

HasNonIoTask returns a boolean if a field has been set.

### GetQueueDepth

`func (o *VolumeStat) GetQueueDepth() float64`

GetQueueDepth returns the QueueDepth field if non-nil, zero value otherwise.

### GetQueueDepthOk

`func (o *VolumeStat) GetQueueDepthOk() (*float64, bool)`

GetQueueDepthOk returns a tuple with the QueueDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDepth

`func (o *VolumeStat) SetQueueDepth(v float64)`

SetQueueDepth sets QueueDepth field to given value.

### HasQueueDepth

`func (o *VolumeStat) HasQueueDepth() bool`

HasQueueDepth returns a boolean if a field has been set.

### GetReadBandwidthKbyte

`func (o *VolumeStat) GetReadBandwidthKbyte() float64`

GetReadBandwidthKbyte returns the ReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetReadBandwidthKbyteOk

`func (o *VolumeStat) GetReadBandwidthKbyteOk() (*float64, bool)`

GetReadBandwidthKbyteOk returns a tuple with the ReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadBandwidthKbyte

`func (o *VolumeStat) SetReadBandwidthKbyte(v float64)`

SetReadBandwidthKbyte sets ReadBandwidthKbyte field to given value.

### HasReadBandwidthKbyte

`func (o *VolumeStat) HasReadBandwidthKbyte() bool`

HasReadBandwidthKbyte returns a boolean if a field has been set.

### GetReadIops

`func (o *VolumeStat) GetReadIops() float64`

GetReadIops returns the ReadIops field if non-nil, zero value otherwise.

### GetReadIopsOk

`func (o *VolumeStat) GetReadIopsOk() (*float64, bool)`

GetReadIopsOk returns a tuple with the ReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIops

`func (o *VolumeStat) SetReadIops(v float64)`

SetReadIops sets ReadIops field to given value.

### HasReadIops

`func (o *VolumeStat) HasReadIops() bool`

HasReadIops returns a boolean if a field has been set.

### GetReadLatencyUs

`func (o *VolumeStat) GetReadLatencyUs() float64`

GetReadLatencyUs returns the ReadLatencyUs field if non-nil, zero value otherwise.

### GetReadLatencyUsOk

`func (o *VolumeStat) GetReadLatencyUsOk() (*float64, bool)`

GetReadLatencyUsOk returns a tuple with the ReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadLatencyUs

`func (o *VolumeStat) SetReadLatencyUs(v float64)`

SetReadLatencyUs sets ReadLatencyUs field to given value.

### HasReadLatencyUs

`func (o *VolumeStat) HasReadLatencyUs() bool`

HasReadLatencyUs returns a boolean if a field has been set.

### GetTotalBandwidthKbyte

`func (o *VolumeStat) GetTotalBandwidthKbyte() float64`

GetTotalBandwidthKbyte returns the TotalBandwidthKbyte field if non-nil, zero value otherwise.

### GetTotalBandwidthKbyteOk

`func (o *VolumeStat) GetTotalBandwidthKbyteOk() (*float64, bool)`

GetTotalBandwidthKbyteOk returns a tuple with the TotalBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBandwidthKbyte

`func (o *VolumeStat) SetTotalBandwidthKbyte(v float64)`

SetTotalBandwidthKbyte sets TotalBandwidthKbyte field to given value.

### HasTotalBandwidthKbyte

`func (o *VolumeStat) HasTotalBandwidthKbyte() bool`

HasTotalBandwidthKbyte returns a boolean if a field has been set.

### GetTotalIops

`func (o *VolumeStat) GetTotalIops() float64`

GetTotalIops returns the TotalIops field if non-nil, zero value otherwise.

### GetTotalIopsOk

`func (o *VolumeStat) GetTotalIopsOk() (*float64, bool)`

GetTotalIopsOk returns a tuple with the TotalIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIops

`func (o *VolumeStat) SetTotalIops(v float64)`

SetTotalIops sets TotalIops field to given value.

### HasTotalIops

`func (o *VolumeStat) HasTotalIops() bool`

HasTotalIops returns a boolean if a field has been set.

### GetWriteBandwidthKbyte

`func (o *VolumeStat) GetWriteBandwidthKbyte() float64`

GetWriteBandwidthKbyte returns the WriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetWriteBandwidthKbyteOk

`func (o *VolumeStat) GetWriteBandwidthKbyteOk() (*float64, bool)`

GetWriteBandwidthKbyteOk returns a tuple with the WriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBandwidthKbyte

`func (o *VolumeStat) SetWriteBandwidthKbyte(v float64)`

SetWriteBandwidthKbyte sets WriteBandwidthKbyte field to given value.

### HasWriteBandwidthKbyte

`func (o *VolumeStat) HasWriteBandwidthKbyte() bool`

HasWriteBandwidthKbyte returns a boolean if a field has been set.

### GetWriteIops

`func (o *VolumeStat) GetWriteIops() float64`

GetWriteIops returns the WriteIops field if non-nil, zero value otherwise.

### GetWriteIopsOk

`func (o *VolumeStat) GetWriteIopsOk() (*float64, bool)`

GetWriteIopsOk returns a tuple with the WriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIops

`func (o *VolumeStat) SetWriteIops(v float64)`

SetWriteIops sets WriteIops field to given value.

### HasWriteIops

`func (o *VolumeStat) HasWriteIops() bool`

HasWriteIops returns a boolean if a field has been set.

### GetWriteLatencyUs

`func (o *VolumeStat) GetWriteLatencyUs() float64`

GetWriteLatencyUs returns the WriteLatencyUs field if non-nil, zero value otherwise.

### GetWriteLatencyUsOk

`func (o *VolumeStat) GetWriteLatencyUsOk() (*float64, bool)`

GetWriteLatencyUsOk returns a tuple with the WriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLatencyUs

`func (o *VolumeStat) SetWriteLatencyUs(v float64)`

SetWriteLatencyUs sets WriteLatencyUs field to given value.

### HasWriteLatencyUs

`func (o *VolumeStat) HasWriteLatencyUs() bool`

HasWriteLatencyUs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


