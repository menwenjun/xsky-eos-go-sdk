# DiskStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgQueueLen** | Pointer to **float64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**IoUtil** | Pointer to **float64** |  | [optional] 
**KbytePerIo** | Pointer to **float64** |  | [optional] 
**ReadBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**ReadIops** | Pointer to **float64** |  | [optional] 
**ReadMergedPs** | Pointer to **float64** |  | [optional] 
**ReadWaitUs** | Pointer to **float64** |  | [optional] 
**TotalBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**TotalIoWaitUs** | Pointer to **float64** |  | [optional] 
**TotalIops** | Pointer to **float64** |  | [optional] 
**TotalKbyte** | Pointer to **int64** |  | [optional] 
**UsedKbyte** | Pointer to **int64** |  | [optional] 
**UsedPercent** | Pointer to **float64** |  | [optional] 
**WriteBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**WriteIops** | Pointer to **float64** |  | [optional] 
**WriteMergedPs** | Pointer to **float64** |  | [optional] 
**WriteWaitUs** | Pointer to **float64** |  | [optional] 

## Methods

### NewDiskStat

`func NewDiskStat() *DiskStat`

NewDiskStat instantiates a new DiskStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskStatWithDefaults

`func NewDiskStatWithDefaults() *DiskStat`

NewDiskStatWithDefaults instantiates a new DiskStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgQueueLen

`func (o *DiskStat) GetAvgQueueLen() float64`

GetAvgQueueLen returns the AvgQueueLen field if non-nil, zero value otherwise.

### GetAvgQueueLenOk

`func (o *DiskStat) GetAvgQueueLenOk() (*float64, bool)`

GetAvgQueueLenOk returns a tuple with the AvgQueueLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgQueueLen

`func (o *DiskStat) SetAvgQueueLen(v float64)`

SetAvgQueueLen sets AvgQueueLen field to given value.

### HasAvgQueueLen

`func (o *DiskStat) HasAvgQueueLen() bool`

HasAvgQueueLen returns a boolean if a field has been set.

### GetCreate

`func (o *DiskStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DiskStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DiskStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DiskStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetIoUtil

`func (o *DiskStat) GetIoUtil() float64`

GetIoUtil returns the IoUtil field if non-nil, zero value otherwise.

### GetIoUtilOk

`func (o *DiskStat) GetIoUtilOk() (*float64, bool)`

GetIoUtilOk returns a tuple with the IoUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoUtil

`func (o *DiskStat) SetIoUtil(v float64)`

SetIoUtil sets IoUtil field to given value.

### HasIoUtil

`func (o *DiskStat) HasIoUtil() bool`

HasIoUtil returns a boolean if a field has been set.

### GetKbytePerIo

`func (o *DiskStat) GetKbytePerIo() float64`

GetKbytePerIo returns the KbytePerIo field if non-nil, zero value otherwise.

### GetKbytePerIoOk

`func (o *DiskStat) GetKbytePerIoOk() (*float64, bool)`

GetKbytePerIoOk returns a tuple with the KbytePerIo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbytePerIo

`func (o *DiskStat) SetKbytePerIo(v float64)`

SetKbytePerIo sets KbytePerIo field to given value.

### HasKbytePerIo

`func (o *DiskStat) HasKbytePerIo() bool`

HasKbytePerIo returns a boolean if a field has been set.

### GetReadBandwidthKbyte

`func (o *DiskStat) GetReadBandwidthKbyte() float64`

GetReadBandwidthKbyte returns the ReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetReadBandwidthKbyteOk

`func (o *DiskStat) GetReadBandwidthKbyteOk() (*float64, bool)`

GetReadBandwidthKbyteOk returns a tuple with the ReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadBandwidthKbyte

`func (o *DiskStat) SetReadBandwidthKbyte(v float64)`

SetReadBandwidthKbyte sets ReadBandwidthKbyte field to given value.

### HasReadBandwidthKbyte

`func (o *DiskStat) HasReadBandwidthKbyte() bool`

HasReadBandwidthKbyte returns a boolean if a field has been set.

### GetReadIops

`func (o *DiskStat) GetReadIops() float64`

GetReadIops returns the ReadIops field if non-nil, zero value otherwise.

### GetReadIopsOk

`func (o *DiskStat) GetReadIopsOk() (*float64, bool)`

GetReadIopsOk returns a tuple with the ReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIops

`func (o *DiskStat) SetReadIops(v float64)`

SetReadIops sets ReadIops field to given value.

### HasReadIops

`func (o *DiskStat) HasReadIops() bool`

HasReadIops returns a boolean if a field has been set.

### GetReadMergedPs

`func (o *DiskStat) GetReadMergedPs() float64`

GetReadMergedPs returns the ReadMergedPs field if non-nil, zero value otherwise.

### GetReadMergedPsOk

`func (o *DiskStat) GetReadMergedPsOk() (*float64, bool)`

GetReadMergedPsOk returns a tuple with the ReadMergedPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadMergedPs

`func (o *DiskStat) SetReadMergedPs(v float64)`

SetReadMergedPs sets ReadMergedPs field to given value.

### HasReadMergedPs

`func (o *DiskStat) HasReadMergedPs() bool`

HasReadMergedPs returns a boolean if a field has been set.

### GetReadWaitUs

`func (o *DiskStat) GetReadWaitUs() float64`

GetReadWaitUs returns the ReadWaitUs field if non-nil, zero value otherwise.

### GetReadWaitUsOk

`func (o *DiskStat) GetReadWaitUsOk() (*float64, bool)`

GetReadWaitUsOk returns a tuple with the ReadWaitUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWaitUs

`func (o *DiskStat) SetReadWaitUs(v float64)`

SetReadWaitUs sets ReadWaitUs field to given value.

### HasReadWaitUs

`func (o *DiskStat) HasReadWaitUs() bool`

HasReadWaitUs returns a boolean if a field has been set.

### GetTotalBandwidthKbyte

`func (o *DiskStat) GetTotalBandwidthKbyte() float64`

GetTotalBandwidthKbyte returns the TotalBandwidthKbyte field if non-nil, zero value otherwise.

### GetTotalBandwidthKbyteOk

`func (o *DiskStat) GetTotalBandwidthKbyteOk() (*float64, bool)`

GetTotalBandwidthKbyteOk returns a tuple with the TotalBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBandwidthKbyte

`func (o *DiskStat) SetTotalBandwidthKbyte(v float64)`

SetTotalBandwidthKbyte sets TotalBandwidthKbyte field to given value.

### HasTotalBandwidthKbyte

`func (o *DiskStat) HasTotalBandwidthKbyte() bool`

HasTotalBandwidthKbyte returns a boolean if a field has been set.

### GetTotalIoWaitUs

`func (o *DiskStat) GetTotalIoWaitUs() float64`

GetTotalIoWaitUs returns the TotalIoWaitUs field if non-nil, zero value otherwise.

### GetTotalIoWaitUsOk

`func (o *DiskStat) GetTotalIoWaitUsOk() (*float64, bool)`

GetTotalIoWaitUsOk returns a tuple with the TotalIoWaitUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIoWaitUs

`func (o *DiskStat) SetTotalIoWaitUs(v float64)`

SetTotalIoWaitUs sets TotalIoWaitUs field to given value.

### HasTotalIoWaitUs

`func (o *DiskStat) HasTotalIoWaitUs() bool`

HasTotalIoWaitUs returns a boolean if a field has been set.

### GetTotalIops

`func (o *DiskStat) GetTotalIops() float64`

GetTotalIops returns the TotalIops field if non-nil, zero value otherwise.

### GetTotalIopsOk

`func (o *DiskStat) GetTotalIopsOk() (*float64, bool)`

GetTotalIopsOk returns a tuple with the TotalIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIops

`func (o *DiskStat) SetTotalIops(v float64)`

SetTotalIops sets TotalIops field to given value.

### HasTotalIops

`func (o *DiskStat) HasTotalIops() bool`

HasTotalIops returns a boolean if a field has been set.

### GetTotalKbyte

`func (o *DiskStat) GetTotalKbyte() int64`

GetTotalKbyte returns the TotalKbyte field if non-nil, zero value otherwise.

### GetTotalKbyteOk

`func (o *DiskStat) GetTotalKbyteOk() (*int64, bool)`

GetTotalKbyteOk returns a tuple with the TotalKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalKbyte

`func (o *DiskStat) SetTotalKbyte(v int64)`

SetTotalKbyte sets TotalKbyte field to given value.

### HasTotalKbyte

`func (o *DiskStat) HasTotalKbyte() bool`

HasTotalKbyte returns a boolean if a field has been set.

### GetUsedKbyte

`func (o *DiskStat) GetUsedKbyte() int64`

GetUsedKbyte returns the UsedKbyte field if non-nil, zero value otherwise.

### GetUsedKbyteOk

`func (o *DiskStat) GetUsedKbyteOk() (*int64, bool)`

GetUsedKbyteOk returns a tuple with the UsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte

`func (o *DiskStat) SetUsedKbyte(v int64)`

SetUsedKbyte sets UsedKbyte field to given value.

### HasUsedKbyte

`func (o *DiskStat) HasUsedKbyte() bool`

HasUsedKbyte returns a boolean if a field has been set.

### GetUsedPercent

`func (o *DiskStat) GetUsedPercent() float64`

GetUsedPercent returns the UsedPercent field if non-nil, zero value otherwise.

### GetUsedPercentOk

`func (o *DiskStat) GetUsedPercentOk() (*float64, bool)`

GetUsedPercentOk returns a tuple with the UsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPercent

`func (o *DiskStat) SetUsedPercent(v float64)`

SetUsedPercent sets UsedPercent field to given value.

### HasUsedPercent

`func (o *DiskStat) HasUsedPercent() bool`

HasUsedPercent returns a boolean if a field has been set.

### GetWriteBandwidthKbyte

`func (o *DiskStat) GetWriteBandwidthKbyte() float64`

GetWriteBandwidthKbyte returns the WriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetWriteBandwidthKbyteOk

`func (o *DiskStat) GetWriteBandwidthKbyteOk() (*float64, bool)`

GetWriteBandwidthKbyteOk returns a tuple with the WriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBandwidthKbyte

`func (o *DiskStat) SetWriteBandwidthKbyte(v float64)`

SetWriteBandwidthKbyte sets WriteBandwidthKbyte field to given value.

### HasWriteBandwidthKbyte

`func (o *DiskStat) HasWriteBandwidthKbyte() bool`

HasWriteBandwidthKbyte returns a boolean if a field has been set.

### GetWriteIops

`func (o *DiskStat) GetWriteIops() float64`

GetWriteIops returns the WriteIops field if non-nil, zero value otherwise.

### GetWriteIopsOk

`func (o *DiskStat) GetWriteIopsOk() (*float64, bool)`

GetWriteIopsOk returns a tuple with the WriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIops

`func (o *DiskStat) SetWriteIops(v float64)`

SetWriteIops sets WriteIops field to given value.

### HasWriteIops

`func (o *DiskStat) HasWriteIops() bool`

HasWriteIops returns a boolean if a field has been set.

### GetWriteMergedPs

`func (o *DiskStat) GetWriteMergedPs() float64`

GetWriteMergedPs returns the WriteMergedPs field if non-nil, zero value otherwise.

### GetWriteMergedPsOk

`func (o *DiskStat) GetWriteMergedPsOk() (*float64, bool)`

GetWriteMergedPsOk returns a tuple with the WriteMergedPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteMergedPs

`func (o *DiskStat) SetWriteMergedPs(v float64)`

SetWriteMergedPs sets WriteMergedPs field to given value.

### HasWriteMergedPs

`func (o *DiskStat) HasWriteMergedPs() bool`

HasWriteMergedPs returns a boolean if a field has been set.

### GetWriteWaitUs

`func (o *DiskStat) GetWriteWaitUs() float64`

GetWriteWaitUs returns the WriteWaitUs field if non-nil, zero value otherwise.

### GetWriteWaitUsOk

`func (o *DiskStat) GetWriteWaitUsOk() (*float64, bool)`

GetWriteWaitUsOk returns a tuple with the WriteWaitUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteWaitUs

`func (o *DiskStat) SetWriteWaitUs(v float64)`

SetWriteWaitUs sets WriteWaitUs field to given value.

### HasWriteWaitUs

`func (o *DiskStat) HasWriteWaitUs() bool`

HasWriteWaitUs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


