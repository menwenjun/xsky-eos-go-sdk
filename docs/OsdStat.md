# OsdStat

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
**ActualKbyte** | Pointer to **int64** |  | [optional] 
**CompressedByte** | Pointer to **int64** |  | [optional] 
**CompressedOriginByte** | Pointer to **int64** |  | [optional] 
**DataKbyte** | Pointer to **int64** |  | [optional] 
**DegradedPercent** | Pointer to **float64** |  | [optional] 
**HealthyPercent** | Pointer to **float64** |  | [optional] 
**OmapTotalKbyte** | Pointer to **float64** |  | [optional] 
**OmapUsedKbyte** | Pointer to **float64** |  | [optional] 
**OmapUsedPercent** | Pointer to **float64** |  | [optional] 
**Partition** | Pointer to [**PartitionStat**](PartitionStat.md) |  | [optional] 
**PgCreatingNum** | Pointer to **int64** |  | [optional] 
**PgDegradedNum** | Pointer to **int64** |  | [optional] 
**PgHealthyNum** | Pointer to **int64** |  | [optional] 
**PgRecoveryNum** | Pointer to **int64** |  | [optional] 
**PgTotalNum** | Pointer to **int64** |  | [optional] 
**PgUnavailableNum** | Pointer to **int64** |  | [optional] 
**RecoveryPercent** | Pointer to **float64** |  | [optional] 
**UnavailablePercent** | Pointer to **float64** |  | [optional] 
**WaterLevel** | Pointer to **float64** |  | [optional] 

## Methods

### NewOsdStat

`func NewOsdStat() *OsdStat`

NewOsdStat instantiates a new OsdStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsdStatWithDefaults

`func NewOsdStatWithDefaults() *OsdStat`

NewOsdStatWithDefaults instantiates a new OsdStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgQueueLen

`func (o *OsdStat) GetAvgQueueLen() float64`

GetAvgQueueLen returns the AvgQueueLen field if non-nil, zero value otherwise.

### GetAvgQueueLenOk

`func (o *OsdStat) GetAvgQueueLenOk() (*float64, bool)`

GetAvgQueueLenOk returns a tuple with the AvgQueueLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgQueueLen

`func (o *OsdStat) SetAvgQueueLen(v float64)`

SetAvgQueueLen sets AvgQueueLen field to given value.

### HasAvgQueueLen

`func (o *OsdStat) HasAvgQueueLen() bool`

HasAvgQueueLen returns a boolean if a field has been set.

### GetCreate

`func (o *OsdStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OsdStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OsdStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OsdStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetIoUtil

`func (o *OsdStat) GetIoUtil() float64`

GetIoUtil returns the IoUtil field if non-nil, zero value otherwise.

### GetIoUtilOk

`func (o *OsdStat) GetIoUtilOk() (*float64, bool)`

GetIoUtilOk returns a tuple with the IoUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoUtil

`func (o *OsdStat) SetIoUtil(v float64)`

SetIoUtil sets IoUtil field to given value.

### HasIoUtil

`func (o *OsdStat) HasIoUtil() bool`

HasIoUtil returns a boolean if a field has been set.

### GetKbytePerIo

`func (o *OsdStat) GetKbytePerIo() float64`

GetKbytePerIo returns the KbytePerIo field if non-nil, zero value otherwise.

### GetKbytePerIoOk

`func (o *OsdStat) GetKbytePerIoOk() (*float64, bool)`

GetKbytePerIoOk returns a tuple with the KbytePerIo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbytePerIo

`func (o *OsdStat) SetKbytePerIo(v float64)`

SetKbytePerIo sets KbytePerIo field to given value.

### HasKbytePerIo

`func (o *OsdStat) HasKbytePerIo() bool`

HasKbytePerIo returns a boolean if a field has been set.

### GetReadBandwidthKbyte

`func (o *OsdStat) GetReadBandwidthKbyte() float64`

GetReadBandwidthKbyte returns the ReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetReadBandwidthKbyteOk

`func (o *OsdStat) GetReadBandwidthKbyteOk() (*float64, bool)`

GetReadBandwidthKbyteOk returns a tuple with the ReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadBandwidthKbyte

`func (o *OsdStat) SetReadBandwidthKbyte(v float64)`

SetReadBandwidthKbyte sets ReadBandwidthKbyte field to given value.

### HasReadBandwidthKbyte

`func (o *OsdStat) HasReadBandwidthKbyte() bool`

HasReadBandwidthKbyte returns a boolean if a field has been set.

### GetReadIops

`func (o *OsdStat) GetReadIops() float64`

GetReadIops returns the ReadIops field if non-nil, zero value otherwise.

### GetReadIopsOk

`func (o *OsdStat) GetReadIopsOk() (*float64, bool)`

GetReadIopsOk returns a tuple with the ReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIops

`func (o *OsdStat) SetReadIops(v float64)`

SetReadIops sets ReadIops field to given value.

### HasReadIops

`func (o *OsdStat) HasReadIops() bool`

HasReadIops returns a boolean if a field has been set.

### GetReadMergedPs

`func (o *OsdStat) GetReadMergedPs() float64`

GetReadMergedPs returns the ReadMergedPs field if non-nil, zero value otherwise.

### GetReadMergedPsOk

`func (o *OsdStat) GetReadMergedPsOk() (*float64, bool)`

GetReadMergedPsOk returns a tuple with the ReadMergedPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadMergedPs

`func (o *OsdStat) SetReadMergedPs(v float64)`

SetReadMergedPs sets ReadMergedPs field to given value.

### HasReadMergedPs

`func (o *OsdStat) HasReadMergedPs() bool`

HasReadMergedPs returns a boolean if a field has been set.

### GetReadWaitUs

`func (o *OsdStat) GetReadWaitUs() float64`

GetReadWaitUs returns the ReadWaitUs field if non-nil, zero value otherwise.

### GetReadWaitUsOk

`func (o *OsdStat) GetReadWaitUsOk() (*float64, bool)`

GetReadWaitUsOk returns a tuple with the ReadWaitUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWaitUs

`func (o *OsdStat) SetReadWaitUs(v float64)`

SetReadWaitUs sets ReadWaitUs field to given value.

### HasReadWaitUs

`func (o *OsdStat) HasReadWaitUs() bool`

HasReadWaitUs returns a boolean if a field has been set.

### GetTotalBandwidthKbyte

`func (o *OsdStat) GetTotalBandwidthKbyte() float64`

GetTotalBandwidthKbyte returns the TotalBandwidthKbyte field if non-nil, zero value otherwise.

### GetTotalBandwidthKbyteOk

`func (o *OsdStat) GetTotalBandwidthKbyteOk() (*float64, bool)`

GetTotalBandwidthKbyteOk returns a tuple with the TotalBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBandwidthKbyte

`func (o *OsdStat) SetTotalBandwidthKbyte(v float64)`

SetTotalBandwidthKbyte sets TotalBandwidthKbyte field to given value.

### HasTotalBandwidthKbyte

`func (o *OsdStat) HasTotalBandwidthKbyte() bool`

HasTotalBandwidthKbyte returns a boolean if a field has been set.

### GetTotalIoWaitUs

`func (o *OsdStat) GetTotalIoWaitUs() float64`

GetTotalIoWaitUs returns the TotalIoWaitUs field if non-nil, zero value otherwise.

### GetTotalIoWaitUsOk

`func (o *OsdStat) GetTotalIoWaitUsOk() (*float64, bool)`

GetTotalIoWaitUsOk returns a tuple with the TotalIoWaitUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIoWaitUs

`func (o *OsdStat) SetTotalIoWaitUs(v float64)`

SetTotalIoWaitUs sets TotalIoWaitUs field to given value.

### HasTotalIoWaitUs

`func (o *OsdStat) HasTotalIoWaitUs() bool`

HasTotalIoWaitUs returns a boolean if a field has been set.

### GetTotalIops

`func (o *OsdStat) GetTotalIops() float64`

GetTotalIops returns the TotalIops field if non-nil, zero value otherwise.

### GetTotalIopsOk

`func (o *OsdStat) GetTotalIopsOk() (*float64, bool)`

GetTotalIopsOk returns a tuple with the TotalIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIops

`func (o *OsdStat) SetTotalIops(v float64)`

SetTotalIops sets TotalIops field to given value.

### HasTotalIops

`func (o *OsdStat) HasTotalIops() bool`

HasTotalIops returns a boolean if a field has been set.

### GetTotalKbyte

`func (o *OsdStat) GetTotalKbyte() int64`

GetTotalKbyte returns the TotalKbyte field if non-nil, zero value otherwise.

### GetTotalKbyteOk

`func (o *OsdStat) GetTotalKbyteOk() (*int64, bool)`

GetTotalKbyteOk returns a tuple with the TotalKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalKbyte

`func (o *OsdStat) SetTotalKbyte(v int64)`

SetTotalKbyte sets TotalKbyte field to given value.

### HasTotalKbyte

`func (o *OsdStat) HasTotalKbyte() bool`

HasTotalKbyte returns a boolean if a field has been set.

### GetUsedKbyte

`func (o *OsdStat) GetUsedKbyte() int64`

GetUsedKbyte returns the UsedKbyte field if non-nil, zero value otherwise.

### GetUsedKbyteOk

`func (o *OsdStat) GetUsedKbyteOk() (*int64, bool)`

GetUsedKbyteOk returns a tuple with the UsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte

`func (o *OsdStat) SetUsedKbyte(v int64)`

SetUsedKbyte sets UsedKbyte field to given value.

### HasUsedKbyte

`func (o *OsdStat) HasUsedKbyte() bool`

HasUsedKbyte returns a boolean if a field has been set.

### GetUsedPercent

`func (o *OsdStat) GetUsedPercent() float64`

GetUsedPercent returns the UsedPercent field if non-nil, zero value otherwise.

### GetUsedPercentOk

`func (o *OsdStat) GetUsedPercentOk() (*float64, bool)`

GetUsedPercentOk returns a tuple with the UsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPercent

`func (o *OsdStat) SetUsedPercent(v float64)`

SetUsedPercent sets UsedPercent field to given value.

### HasUsedPercent

`func (o *OsdStat) HasUsedPercent() bool`

HasUsedPercent returns a boolean if a field has been set.

### GetWriteBandwidthKbyte

`func (o *OsdStat) GetWriteBandwidthKbyte() float64`

GetWriteBandwidthKbyte returns the WriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetWriteBandwidthKbyteOk

`func (o *OsdStat) GetWriteBandwidthKbyteOk() (*float64, bool)`

GetWriteBandwidthKbyteOk returns a tuple with the WriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBandwidthKbyte

`func (o *OsdStat) SetWriteBandwidthKbyte(v float64)`

SetWriteBandwidthKbyte sets WriteBandwidthKbyte field to given value.

### HasWriteBandwidthKbyte

`func (o *OsdStat) HasWriteBandwidthKbyte() bool`

HasWriteBandwidthKbyte returns a boolean if a field has been set.

### GetWriteIops

`func (o *OsdStat) GetWriteIops() float64`

GetWriteIops returns the WriteIops field if non-nil, zero value otherwise.

### GetWriteIopsOk

`func (o *OsdStat) GetWriteIopsOk() (*float64, bool)`

GetWriteIopsOk returns a tuple with the WriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIops

`func (o *OsdStat) SetWriteIops(v float64)`

SetWriteIops sets WriteIops field to given value.

### HasWriteIops

`func (o *OsdStat) HasWriteIops() bool`

HasWriteIops returns a boolean if a field has been set.

### GetWriteMergedPs

`func (o *OsdStat) GetWriteMergedPs() float64`

GetWriteMergedPs returns the WriteMergedPs field if non-nil, zero value otherwise.

### GetWriteMergedPsOk

`func (o *OsdStat) GetWriteMergedPsOk() (*float64, bool)`

GetWriteMergedPsOk returns a tuple with the WriteMergedPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteMergedPs

`func (o *OsdStat) SetWriteMergedPs(v float64)`

SetWriteMergedPs sets WriteMergedPs field to given value.

### HasWriteMergedPs

`func (o *OsdStat) HasWriteMergedPs() bool`

HasWriteMergedPs returns a boolean if a field has been set.

### GetWriteWaitUs

`func (o *OsdStat) GetWriteWaitUs() float64`

GetWriteWaitUs returns the WriteWaitUs field if non-nil, zero value otherwise.

### GetWriteWaitUsOk

`func (o *OsdStat) GetWriteWaitUsOk() (*float64, bool)`

GetWriteWaitUsOk returns a tuple with the WriteWaitUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteWaitUs

`func (o *OsdStat) SetWriteWaitUs(v float64)`

SetWriteWaitUs sets WriteWaitUs field to given value.

### HasWriteWaitUs

`func (o *OsdStat) HasWriteWaitUs() bool`

HasWriteWaitUs returns a boolean if a field has been set.

### GetActualKbyte

`func (o *OsdStat) GetActualKbyte() int64`

GetActualKbyte returns the ActualKbyte field if non-nil, zero value otherwise.

### GetActualKbyteOk

`func (o *OsdStat) GetActualKbyteOk() (*int64, bool)`

GetActualKbyteOk returns a tuple with the ActualKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualKbyte

`func (o *OsdStat) SetActualKbyte(v int64)`

SetActualKbyte sets ActualKbyte field to given value.

### HasActualKbyte

`func (o *OsdStat) HasActualKbyte() bool`

HasActualKbyte returns a boolean if a field has been set.

### GetCompressedByte

`func (o *OsdStat) GetCompressedByte() int64`

GetCompressedByte returns the CompressedByte field if non-nil, zero value otherwise.

### GetCompressedByteOk

`func (o *OsdStat) GetCompressedByteOk() (*int64, bool)`

GetCompressedByteOk returns a tuple with the CompressedByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedByte

`func (o *OsdStat) SetCompressedByte(v int64)`

SetCompressedByte sets CompressedByte field to given value.

### HasCompressedByte

`func (o *OsdStat) HasCompressedByte() bool`

HasCompressedByte returns a boolean if a field has been set.

### GetCompressedOriginByte

`func (o *OsdStat) GetCompressedOriginByte() int64`

GetCompressedOriginByte returns the CompressedOriginByte field if non-nil, zero value otherwise.

### GetCompressedOriginByteOk

`func (o *OsdStat) GetCompressedOriginByteOk() (*int64, bool)`

GetCompressedOriginByteOk returns a tuple with the CompressedOriginByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedOriginByte

`func (o *OsdStat) SetCompressedOriginByte(v int64)`

SetCompressedOriginByte sets CompressedOriginByte field to given value.

### HasCompressedOriginByte

`func (o *OsdStat) HasCompressedOriginByte() bool`

HasCompressedOriginByte returns a boolean if a field has been set.

### GetDataKbyte

`func (o *OsdStat) GetDataKbyte() int64`

GetDataKbyte returns the DataKbyte field if non-nil, zero value otherwise.

### GetDataKbyteOk

`func (o *OsdStat) GetDataKbyteOk() (*int64, bool)`

GetDataKbyteOk returns a tuple with the DataKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataKbyte

`func (o *OsdStat) SetDataKbyte(v int64)`

SetDataKbyte sets DataKbyte field to given value.

### HasDataKbyte

`func (o *OsdStat) HasDataKbyte() bool`

HasDataKbyte returns a boolean if a field has been set.

### GetDegradedPercent

`func (o *OsdStat) GetDegradedPercent() float64`

GetDegradedPercent returns the DegradedPercent field if non-nil, zero value otherwise.

### GetDegradedPercentOk

`func (o *OsdStat) GetDegradedPercentOk() (*float64, bool)`

GetDegradedPercentOk returns a tuple with the DegradedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedPercent

`func (o *OsdStat) SetDegradedPercent(v float64)`

SetDegradedPercent sets DegradedPercent field to given value.

### HasDegradedPercent

`func (o *OsdStat) HasDegradedPercent() bool`

HasDegradedPercent returns a boolean if a field has been set.

### GetHealthyPercent

`func (o *OsdStat) GetHealthyPercent() float64`

GetHealthyPercent returns the HealthyPercent field if non-nil, zero value otherwise.

### GetHealthyPercentOk

`func (o *OsdStat) GetHealthyPercentOk() (*float64, bool)`

GetHealthyPercentOk returns a tuple with the HealthyPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyPercent

`func (o *OsdStat) SetHealthyPercent(v float64)`

SetHealthyPercent sets HealthyPercent field to given value.

### HasHealthyPercent

`func (o *OsdStat) HasHealthyPercent() bool`

HasHealthyPercent returns a boolean if a field has been set.

### GetOmapTotalKbyte

`func (o *OsdStat) GetOmapTotalKbyte() float64`

GetOmapTotalKbyte returns the OmapTotalKbyte field if non-nil, zero value otherwise.

### GetOmapTotalKbyteOk

`func (o *OsdStat) GetOmapTotalKbyteOk() (*float64, bool)`

GetOmapTotalKbyteOk returns a tuple with the OmapTotalKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmapTotalKbyte

`func (o *OsdStat) SetOmapTotalKbyte(v float64)`

SetOmapTotalKbyte sets OmapTotalKbyte field to given value.

### HasOmapTotalKbyte

`func (o *OsdStat) HasOmapTotalKbyte() bool`

HasOmapTotalKbyte returns a boolean if a field has been set.

### GetOmapUsedKbyte

`func (o *OsdStat) GetOmapUsedKbyte() float64`

GetOmapUsedKbyte returns the OmapUsedKbyte field if non-nil, zero value otherwise.

### GetOmapUsedKbyteOk

`func (o *OsdStat) GetOmapUsedKbyteOk() (*float64, bool)`

GetOmapUsedKbyteOk returns a tuple with the OmapUsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmapUsedKbyte

`func (o *OsdStat) SetOmapUsedKbyte(v float64)`

SetOmapUsedKbyte sets OmapUsedKbyte field to given value.

### HasOmapUsedKbyte

`func (o *OsdStat) HasOmapUsedKbyte() bool`

HasOmapUsedKbyte returns a boolean if a field has been set.

### GetOmapUsedPercent

`func (o *OsdStat) GetOmapUsedPercent() float64`

GetOmapUsedPercent returns the OmapUsedPercent field if non-nil, zero value otherwise.

### GetOmapUsedPercentOk

`func (o *OsdStat) GetOmapUsedPercentOk() (*float64, bool)`

GetOmapUsedPercentOk returns a tuple with the OmapUsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmapUsedPercent

`func (o *OsdStat) SetOmapUsedPercent(v float64)`

SetOmapUsedPercent sets OmapUsedPercent field to given value.

### HasOmapUsedPercent

`func (o *OsdStat) HasOmapUsedPercent() bool`

HasOmapUsedPercent returns a boolean if a field has been set.

### GetPartition

`func (o *OsdStat) GetPartition() PartitionStat`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *OsdStat) GetPartitionOk() (*PartitionStat, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *OsdStat) SetPartition(v PartitionStat)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *OsdStat) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPgCreatingNum

`func (o *OsdStat) GetPgCreatingNum() int64`

GetPgCreatingNum returns the PgCreatingNum field if non-nil, zero value otherwise.

### GetPgCreatingNumOk

`func (o *OsdStat) GetPgCreatingNumOk() (*int64, bool)`

GetPgCreatingNumOk returns a tuple with the PgCreatingNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgCreatingNum

`func (o *OsdStat) SetPgCreatingNum(v int64)`

SetPgCreatingNum sets PgCreatingNum field to given value.

### HasPgCreatingNum

`func (o *OsdStat) HasPgCreatingNum() bool`

HasPgCreatingNum returns a boolean if a field has been set.

### GetPgDegradedNum

`func (o *OsdStat) GetPgDegradedNum() int64`

GetPgDegradedNum returns the PgDegradedNum field if non-nil, zero value otherwise.

### GetPgDegradedNumOk

`func (o *OsdStat) GetPgDegradedNumOk() (*int64, bool)`

GetPgDegradedNumOk returns a tuple with the PgDegradedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgDegradedNum

`func (o *OsdStat) SetPgDegradedNum(v int64)`

SetPgDegradedNum sets PgDegradedNum field to given value.

### HasPgDegradedNum

`func (o *OsdStat) HasPgDegradedNum() bool`

HasPgDegradedNum returns a boolean if a field has been set.

### GetPgHealthyNum

`func (o *OsdStat) GetPgHealthyNum() int64`

GetPgHealthyNum returns the PgHealthyNum field if non-nil, zero value otherwise.

### GetPgHealthyNumOk

`func (o *OsdStat) GetPgHealthyNumOk() (*int64, bool)`

GetPgHealthyNumOk returns a tuple with the PgHealthyNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgHealthyNum

`func (o *OsdStat) SetPgHealthyNum(v int64)`

SetPgHealthyNum sets PgHealthyNum field to given value.

### HasPgHealthyNum

`func (o *OsdStat) HasPgHealthyNum() bool`

HasPgHealthyNum returns a boolean if a field has been set.

### GetPgRecoveryNum

`func (o *OsdStat) GetPgRecoveryNum() int64`

GetPgRecoveryNum returns the PgRecoveryNum field if non-nil, zero value otherwise.

### GetPgRecoveryNumOk

`func (o *OsdStat) GetPgRecoveryNumOk() (*int64, bool)`

GetPgRecoveryNumOk returns a tuple with the PgRecoveryNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgRecoveryNum

`func (o *OsdStat) SetPgRecoveryNum(v int64)`

SetPgRecoveryNum sets PgRecoveryNum field to given value.

### HasPgRecoveryNum

`func (o *OsdStat) HasPgRecoveryNum() bool`

HasPgRecoveryNum returns a boolean if a field has been set.

### GetPgTotalNum

`func (o *OsdStat) GetPgTotalNum() int64`

GetPgTotalNum returns the PgTotalNum field if non-nil, zero value otherwise.

### GetPgTotalNumOk

`func (o *OsdStat) GetPgTotalNumOk() (*int64, bool)`

GetPgTotalNumOk returns a tuple with the PgTotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgTotalNum

`func (o *OsdStat) SetPgTotalNum(v int64)`

SetPgTotalNum sets PgTotalNum field to given value.

### HasPgTotalNum

`func (o *OsdStat) HasPgTotalNum() bool`

HasPgTotalNum returns a boolean if a field has been set.

### GetPgUnavailableNum

`func (o *OsdStat) GetPgUnavailableNum() int64`

GetPgUnavailableNum returns the PgUnavailableNum field if non-nil, zero value otherwise.

### GetPgUnavailableNumOk

`func (o *OsdStat) GetPgUnavailableNumOk() (*int64, bool)`

GetPgUnavailableNumOk returns a tuple with the PgUnavailableNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgUnavailableNum

`func (o *OsdStat) SetPgUnavailableNum(v int64)`

SetPgUnavailableNum sets PgUnavailableNum field to given value.

### HasPgUnavailableNum

`func (o *OsdStat) HasPgUnavailableNum() bool`

HasPgUnavailableNum returns a boolean if a field has been set.

### GetRecoveryPercent

`func (o *OsdStat) GetRecoveryPercent() float64`

GetRecoveryPercent returns the RecoveryPercent field if non-nil, zero value otherwise.

### GetRecoveryPercentOk

`func (o *OsdStat) GetRecoveryPercentOk() (*float64, bool)`

GetRecoveryPercentOk returns a tuple with the RecoveryPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPercent

`func (o *OsdStat) SetRecoveryPercent(v float64)`

SetRecoveryPercent sets RecoveryPercent field to given value.

### HasRecoveryPercent

`func (o *OsdStat) HasRecoveryPercent() bool`

HasRecoveryPercent returns a boolean if a field has been set.

### GetUnavailablePercent

`func (o *OsdStat) GetUnavailablePercent() float64`

GetUnavailablePercent returns the UnavailablePercent field if non-nil, zero value otherwise.

### GetUnavailablePercentOk

`func (o *OsdStat) GetUnavailablePercentOk() (*float64, bool)`

GetUnavailablePercentOk returns a tuple with the UnavailablePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailablePercent

`func (o *OsdStat) SetUnavailablePercent(v float64)`

SetUnavailablePercent sets UnavailablePercent field to given value.

### HasUnavailablePercent

`func (o *OsdStat) HasUnavailablePercent() bool`

HasUnavailablePercent returns a boolean if a field has been set.

### GetWaterLevel

`func (o *OsdStat) GetWaterLevel() float64`

GetWaterLevel returns the WaterLevel field if non-nil, zero value otherwise.

### GetWaterLevelOk

`func (o *OsdStat) GetWaterLevelOk() (*float64, bool)`

GetWaterLevelOk returns a tuple with the WaterLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaterLevel

`func (o *OsdStat) SetWaterLevel(v float64)`

SetWaterLevel sets WaterLevel field to given value.

### HasWaterLevel

`func (o *OsdStat) HasWaterLevel() bool`

HasWaterLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


