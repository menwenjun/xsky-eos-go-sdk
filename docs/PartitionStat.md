# PartitionStat

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

### NewPartitionStat

`func NewPartitionStat() *PartitionStat`

NewPartitionStat instantiates a new PartitionStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionStatWithDefaults

`func NewPartitionStatWithDefaults() *PartitionStat`

NewPartitionStatWithDefaults instantiates a new PartitionStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgQueueLen

`func (o *PartitionStat) GetAvgQueueLen() float64`

GetAvgQueueLen returns the AvgQueueLen field if non-nil, zero value otherwise.

### GetAvgQueueLenOk

`func (o *PartitionStat) GetAvgQueueLenOk() (*float64, bool)`

GetAvgQueueLenOk returns a tuple with the AvgQueueLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgQueueLen

`func (o *PartitionStat) SetAvgQueueLen(v float64)`

SetAvgQueueLen sets AvgQueueLen field to given value.

### HasAvgQueueLen

`func (o *PartitionStat) HasAvgQueueLen() bool`

HasAvgQueueLen returns a boolean if a field has been set.

### GetCreate

`func (o *PartitionStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *PartitionStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *PartitionStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *PartitionStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetIoUtil

`func (o *PartitionStat) GetIoUtil() float64`

GetIoUtil returns the IoUtil field if non-nil, zero value otherwise.

### GetIoUtilOk

`func (o *PartitionStat) GetIoUtilOk() (*float64, bool)`

GetIoUtilOk returns a tuple with the IoUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoUtil

`func (o *PartitionStat) SetIoUtil(v float64)`

SetIoUtil sets IoUtil field to given value.

### HasIoUtil

`func (o *PartitionStat) HasIoUtil() bool`

HasIoUtil returns a boolean if a field has been set.

### GetKbytePerIo

`func (o *PartitionStat) GetKbytePerIo() float64`

GetKbytePerIo returns the KbytePerIo field if non-nil, zero value otherwise.

### GetKbytePerIoOk

`func (o *PartitionStat) GetKbytePerIoOk() (*float64, bool)`

GetKbytePerIoOk returns a tuple with the KbytePerIo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbytePerIo

`func (o *PartitionStat) SetKbytePerIo(v float64)`

SetKbytePerIo sets KbytePerIo field to given value.

### HasKbytePerIo

`func (o *PartitionStat) HasKbytePerIo() bool`

HasKbytePerIo returns a boolean if a field has been set.

### GetReadBandwidthKbyte

`func (o *PartitionStat) GetReadBandwidthKbyte() float64`

GetReadBandwidthKbyte returns the ReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetReadBandwidthKbyteOk

`func (o *PartitionStat) GetReadBandwidthKbyteOk() (*float64, bool)`

GetReadBandwidthKbyteOk returns a tuple with the ReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadBandwidthKbyte

`func (o *PartitionStat) SetReadBandwidthKbyte(v float64)`

SetReadBandwidthKbyte sets ReadBandwidthKbyte field to given value.

### HasReadBandwidthKbyte

`func (o *PartitionStat) HasReadBandwidthKbyte() bool`

HasReadBandwidthKbyte returns a boolean if a field has been set.

### GetReadIops

`func (o *PartitionStat) GetReadIops() float64`

GetReadIops returns the ReadIops field if non-nil, zero value otherwise.

### GetReadIopsOk

`func (o *PartitionStat) GetReadIopsOk() (*float64, bool)`

GetReadIopsOk returns a tuple with the ReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIops

`func (o *PartitionStat) SetReadIops(v float64)`

SetReadIops sets ReadIops field to given value.

### HasReadIops

`func (o *PartitionStat) HasReadIops() bool`

HasReadIops returns a boolean if a field has been set.

### GetReadMergedPs

`func (o *PartitionStat) GetReadMergedPs() float64`

GetReadMergedPs returns the ReadMergedPs field if non-nil, zero value otherwise.

### GetReadMergedPsOk

`func (o *PartitionStat) GetReadMergedPsOk() (*float64, bool)`

GetReadMergedPsOk returns a tuple with the ReadMergedPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadMergedPs

`func (o *PartitionStat) SetReadMergedPs(v float64)`

SetReadMergedPs sets ReadMergedPs field to given value.

### HasReadMergedPs

`func (o *PartitionStat) HasReadMergedPs() bool`

HasReadMergedPs returns a boolean if a field has been set.

### GetReadWaitUs

`func (o *PartitionStat) GetReadWaitUs() float64`

GetReadWaitUs returns the ReadWaitUs field if non-nil, zero value otherwise.

### GetReadWaitUsOk

`func (o *PartitionStat) GetReadWaitUsOk() (*float64, bool)`

GetReadWaitUsOk returns a tuple with the ReadWaitUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWaitUs

`func (o *PartitionStat) SetReadWaitUs(v float64)`

SetReadWaitUs sets ReadWaitUs field to given value.

### HasReadWaitUs

`func (o *PartitionStat) HasReadWaitUs() bool`

HasReadWaitUs returns a boolean if a field has been set.

### GetTotalBandwidthKbyte

`func (o *PartitionStat) GetTotalBandwidthKbyte() float64`

GetTotalBandwidthKbyte returns the TotalBandwidthKbyte field if non-nil, zero value otherwise.

### GetTotalBandwidthKbyteOk

`func (o *PartitionStat) GetTotalBandwidthKbyteOk() (*float64, bool)`

GetTotalBandwidthKbyteOk returns a tuple with the TotalBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBandwidthKbyte

`func (o *PartitionStat) SetTotalBandwidthKbyte(v float64)`

SetTotalBandwidthKbyte sets TotalBandwidthKbyte field to given value.

### HasTotalBandwidthKbyte

`func (o *PartitionStat) HasTotalBandwidthKbyte() bool`

HasTotalBandwidthKbyte returns a boolean if a field has been set.

### GetTotalIoWaitUs

`func (o *PartitionStat) GetTotalIoWaitUs() float64`

GetTotalIoWaitUs returns the TotalIoWaitUs field if non-nil, zero value otherwise.

### GetTotalIoWaitUsOk

`func (o *PartitionStat) GetTotalIoWaitUsOk() (*float64, bool)`

GetTotalIoWaitUsOk returns a tuple with the TotalIoWaitUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIoWaitUs

`func (o *PartitionStat) SetTotalIoWaitUs(v float64)`

SetTotalIoWaitUs sets TotalIoWaitUs field to given value.

### HasTotalIoWaitUs

`func (o *PartitionStat) HasTotalIoWaitUs() bool`

HasTotalIoWaitUs returns a boolean if a field has been set.

### GetTotalIops

`func (o *PartitionStat) GetTotalIops() float64`

GetTotalIops returns the TotalIops field if non-nil, zero value otherwise.

### GetTotalIopsOk

`func (o *PartitionStat) GetTotalIopsOk() (*float64, bool)`

GetTotalIopsOk returns a tuple with the TotalIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIops

`func (o *PartitionStat) SetTotalIops(v float64)`

SetTotalIops sets TotalIops field to given value.

### HasTotalIops

`func (o *PartitionStat) HasTotalIops() bool`

HasTotalIops returns a boolean if a field has been set.

### GetTotalKbyte

`func (o *PartitionStat) GetTotalKbyte() int64`

GetTotalKbyte returns the TotalKbyte field if non-nil, zero value otherwise.

### GetTotalKbyteOk

`func (o *PartitionStat) GetTotalKbyteOk() (*int64, bool)`

GetTotalKbyteOk returns a tuple with the TotalKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalKbyte

`func (o *PartitionStat) SetTotalKbyte(v int64)`

SetTotalKbyte sets TotalKbyte field to given value.

### HasTotalKbyte

`func (o *PartitionStat) HasTotalKbyte() bool`

HasTotalKbyte returns a boolean if a field has been set.

### GetUsedKbyte

`func (o *PartitionStat) GetUsedKbyte() int64`

GetUsedKbyte returns the UsedKbyte field if non-nil, zero value otherwise.

### GetUsedKbyteOk

`func (o *PartitionStat) GetUsedKbyteOk() (*int64, bool)`

GetUsedKbyteOk returns a tuple with the UsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte

`func (o *PartitionStat) SetUsedKbyte(v int64)`

SetUsedKbyte sets UsedKbyte field to given value.

### HasUsedKbyte

`func (o *PartitionStat) HasUsedKbyte() bool`

HasUsedKbyte returns a boolean if a field has been set.

### GetUsedPercent

`func (o *PartitionStat) GetUsedPercent() float64`

GetUsedPercent returns the UsedPercent field if non-nil, zero value otherwise.

### GetUsedPercentOk

`func (o *PartitionStat) GetUsedPercentOk() (*float64, bool)`

GetUsedPercentOk returns a tuple with the UsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPercent

`func (o *PartitionStat) SetUsedPercent(v float64)`

SetUsedPercent sets UsedPercent field to given value.

### HasUsedPercent

`func (o *PartitionStat) HasUsedPercent() bool`

HasUsedPercent returns a boolean if a field has been set.

### GetWriteBandwidthKbyte

`func (o *PartitionStat) GetWriteBandwidthKbyte() float64`

GetWriteBandwidthKbyte returns the WriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetWriteBandwidthKbyteOk

`func (o *PartitionStat) GetWriteBandwidthKbyteOk() (*float64, bool)`

GetWriteBandwidthKbyteOk returns a tuple with the WriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBandwidthKbyte

`func (o *PartitionStat) SetWriteBandwidthKbyte(v float64)`

SetWriteBandwidthKbyte sets WriteBandwidthKbyte field to given value.

### HasWriteBandwidthKbyte

`func (o *PartitionStat) HasWriteBandwidthKbyte() bool`

HasWriteBandwidthKbyte returns a boolean if a field has been set.

### GetWriteIops

`func (o *PartitionStat) GetWriteIops() float64`

GetWriteIops returns the WriteIops field if non-nil, zero value otherwise.

### GetWriteIopsOk

`func (o *PartitionStat) GetWriteIopsOk() (*float64, bool)`

GetWriteIopsOk returns a tuple with the WriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIops

`func (o *PartitionStat) SetWriteIops(v float64)`

SetWriteIops sets WriteIops field to given value.

### HasWriteIops

`func (o *PartitionStat) HasWriteIops() bool`

HasWriteIops returns a boolean if a field has been set.

### GetWriteMergedPs

`func (o *PartitionStat) GetWriteMergedPs() float64`

GetWriteMergedPs returns the WriteMergedPs field if non-nil, zero value otherwise.

### GetWriteMergedPsOk

`func (o *PartitionStat) GetWriteMergedPsOk() (*float64, bool)`

GetWriteMergedPsOk returns a tuple with the WriteMergedPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteMergedPs

`func (o *PartitionStat) SetWriteMergedPs(v float64)`

SetWriteMergedPs sets WriteMergedPs field to given value.

### HasWriteMergedPs

`func (o *PartitionStat) HasWriteMergedPs() bool`

HasWriteMergedPs returns a boolean if a field has been set.

### GetWriteWaitUs

`func (o *PartitionStat) GetWriteWaitUs() float64`

GetWriteWaitUs returns the WriteWaitUs field if non-nil, zero value otherwise.

### GetWriteWaitUsOk

`func (o *PartitionStat) GetWriteWaitUsOk() (*float64, bool)`

GetWriteWaitUsOk returns a tuple with the WriteWaitUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteWaitUs

`func (o *PartitionStat) SetWriteWaitUs(v float64)`

SetWriteWaitUs sets WriteWaitUs field to given value.

### HasWriteWaitUs

`func (o *PartitionStat) HasWriteWaitUs() bool`

HasWriteWaitUs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


