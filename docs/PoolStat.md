# PoolStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualKbyte** | Pointer to **int64** |  | [optional] 
**AvailableActualKbyte** | Pointer to **float64** | actual_kbyte*(1-reserved_data_percent) | [optional] 
**ClientReadBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**ClientReadIoSizeKbyte** | Pointer to **float64** |  | [optional] 
**ClientReadIops** | Pointer to **float64** |  | [optional] 
**ClientReadLatencyUs** | Pointer to **float64** |  | [optional] 
**ClientWriteBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**ClientWriteIoSizeKbyte** | Pointer to **float64** |  | [optional] 
**ClientWriteIops** | Pointer to **float64** |  | [optional] 
**ClientWriteLatencyUs** | Pointer to **float64** |  | [optional] 
**CompressRatio** | Pointer to **float64** |  | [optional] 
**CompressedKbyte** | Pointer to **float64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**CreatingNum** | Pointer to **int64** |  | [optional] 
**DataKbyte** | Pointer to **int64** |  | [optional] 
**DegradedNum** | Pointer to **int64** |  | [optional] 
**DegradedPercent** | Pointer to **float64** |  | [optional] 
**ErrorKbyte** | Pointer to **int64** |  | [optional] 
**FlushReadBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**FlushReadIoSizeKbyte** | Pointer to **float64** |  | [optional] 
**FlushReadIops** | Pointer to **float64** |  | [optional] 
**FlushReadLatencyUs** | Pointer to **float64** |  | [optional] 
**FlushWriteBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**FlushWriteIoSizeKbyte** | Pointer to **float64** |  | [optional] 
**FlushWriteIops** | Pointer to **float64** |  | [optional] 
**FlushWriteLatencyUs** | Pointer to **float64** |  | [optional] 
**GarbageRatio** | Pointer to **float64** |  | [optional] 
**GcReadBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**GcReadIoSizeKbyte** | Pointer to **float64** |  | [optional] 
**GcReadIops** | Pointer to **float64** |  | [optional] 
**GcReadLatencyUs** | Pointer to **float64** |  | [optional] 
**GcRemoveBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**GcRemoveIoSizeKbyte** | Pointer to **float64** |  | [optional] 
**GcRemoveIops** | Pointer to **float64** |  | [optional] 
**GcWriteBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**GcWriteIoSizeKbyte** | Pointer to **float64** |  | [optional] 
**GcWriteIops** | Pointer to **float64** |  | [optional] 
**GcWriteLatencyUs** | Pointer to **float64** |  | [optional] 
**HealthyNum** | Pointer to **int64** |  | [optional] 
**HealthyPercent** | Pointer to **float64** |  | [optional] 
**MaxAvailKbyte** | Pointer to **int64** |  | [optional] 
**MinGarbageKbyte** | Pointer to **int64** | garbage size, for tier pool | [optional] 
**OmapTotalKbyte** | Pointer to **float64** |  | [optional] 
**OmapUsedKbyte** | Pointer to **float64** |  | [optional] 
**OmapUsedPercent** | Pointer to **float64** |  | [optional] 
**OsdCapacityUnbalance** | Pointer to **bool** |  | [optional] 
**PoolCapacityUsage** | Pointer to **int64** |  | [optional] 
**ReadBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**ReadCacheHitRate** | Pointer to **float64** |  | [optional] 
**ReadCacheKbyte** | Pointer to **float64** |  | [optional] 
**ReadCachePercent** | Pointer to **float64** |  | [optional] 
**ReadIoSizeKbyte** | Pointer to **float64** |  | [optional] 
**ReadIops** | Pointer to **float64** |  | [optional] 
**ReadLatencyUs** | Pointer to **float64** |  | [optional] 
**RecoveryBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**RecoveryIops** | Pointer to **float64** |  | [optional] 
**RecoveryNum** | Pointer to **int64** |  | [optional] 
**RecoveryPercent** | Pointer to **float64** |  | [optional] 
**RecoveryRemainSecond** | Pointer to **int64** |  | [optional] 
**RemoveBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**RemoveIops** | Pointer to **float64** |  | [optional] 
**ReservedDataKbyte** | Pointer to **float64** |  | [optional] 
**ReservedDataPercent** | Pointer to **float64** |  | [optional] 
**ReservedUsedKbyte** | Pointer to **float64** |  | [optional] 
**RmwBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**RmwIops** | Pointer to **float64** |  | [optional] 
**SnapKbyte** | Pointer to **int64** |  | [optional] 
**TotalBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**TotalCacheKbyte** | Pointer to **float64** |  | [optional] 
**TotalIops** | Pointer to **float64** |  | [optional] 
**TotalKbyte** | Pointer to **int64** |  | [optional] 
**UnavailableNum** | Pointer to **int64** |  | [optional] 
**UnavailablePercent** | Pointer to **float64** |  | [optional] 
**UnusedDataKbyte** | Pointer to **float64** |  | [optional] 
**UsedKbyte** | Pointer to **int64** |  | [optional] 
**UsedPercent** | Pointer to **float64** |  | [optional] 
**WaterLevel** | Pointer to **float64** | tier pool stat | [optional] 
**WriteBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**WriteCacheKbyte** | Pointer to **float64** |  | [optional] 
**WriteCacheMergeRate** | Pointer to **float64** |  | [optional] 
**WriteCachePercent** | Pointer to **float64** |  | [optional] 
**WriteIoSizeKbyte** | Pointer to **float64** |  | [optional] 
**WriteIops** | Pointer to **float64** |  | [optional] 
**WriteLatencyUs** | Pointer to **float64** |  | [optional] 

## Methods

### NewPoolStat

`func NewPoolStat() *PoolStat`

NewPoolStat instantiates a new PoolStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolStatWithDefaults

`func NewPoolStatWithDefaults() *PoolStat`

NewPoolStatWithDefaults instantiates a new PoolStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualKbyte

`func (o *PoolStat) GetActualKbyte() int64`

GetActualKbyte returns the ActualKbyte field if non-nil, zero value otherwise.

### GetActualKbyteOk

`func (o *PoolStat) GetActualKbyteOk() (*int64, bool)`

GetActualKbyteOk returns a tuple with the ActualKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualKbyte

`func (o *PoolStat) SetActualKbyte(v int64)`

SetActualKbyte sets ActualKbyte field to given value.

### HasActualKbyte

`func (o *PoolStat) HasActualKbyte() bool`

HasActualKbyte returns a boolean if a field has been set.

### GetAvailableActualKbyte

`func (o *PoolStat) GetAvailableActualKbyte() float64`

GetAvailableActualKbyte returns the AvailableActualKbyte field if non-nil, zero value otherwise.

### GetAvailableActualKbyteOk

`func (o *PoolStat) GetAvailableActualKbyteOk() (*float64, bool)`

GetAvailableActualKbyteOk returns a tuple with the AvailableActualKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableActualKbyte

`func (o *PoolStat) SetAvailableActualKbyte(v float64)`

SetAvailableActualKbyte sets AvailableActualKbyte field to given value.

### HasAvailableActualKbyte

`func (o *PoolStat) HasAvailableActualKbyte() bool`

HasAvailableActualKbyte returns a boolean if a field has been set.

### GetClientReadBandwidthKbyte

`func (o *PoolStat) GetClientReadBandwidthKbyte() float64`

GetClientReadBandwidthKbyte returns the ClientReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetClientReadBandwidthKbyteOk

`func (o *PoolStat) GetClientReadBandwidthKbyteOk() (*float64, bool)`

GetClientReadBandwidthKbyteOk returns a tuple with the ClientReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReadBandwidthKbyte

`func (o *PoolStat) SetClientReadBandwidthKbyte(v float64)`

SetClientReadBandwidthKbyte sets ClientReadBandwidthKbyte field to given value.

### HasClientReadBandwidthKbyte

`func (o *PoolStat) HasClientReadBandwidthKbyte() bool`

HasClientReadBandwidthKbyte returns a boolean if a field has been set.

### GetClientReadIoSizeKbyte

`func (o *PoolStat) GetClientReadIoSizeKbyte() float64`

GetClientReadIoSizeKbyte returns the ClientReadIoSizeKbyte field if non-nil, zero value otherwise.

### GetClientReadIoSizeKbyteOk

`func (o *PoolStat) GetClientReadIoSizeKbyteOk() (*float64, bool)`

GetClientReadIoSizeKbyteOk returns a tuple with the ClientReadIoSizeKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReadIoSizeKbyte

`func (o *PoolStat) SetClientReadIoSizeKbyte(v float64)`

SetClientReadIoSizeKbyte sets ClientReadIoSizeKbyte field to given value.

### HasClientReadIoSizeKbyte

`func (o *PoolStat) HasClientReadIoSizeKbyte() bool`

HasClientReadIoSizeKbyte returns a boolean if a field has been set.

### GetClientReadIops

`func (o *PoolStat) GetClientReadIops() float64`

GetClientReadIops returns the ClientReadIops field if non-nil, zero value otherwise.

### GetClientReadIopsOk

`func (o *PoolStat) GetClientReadIopsOk() (*float64, bool)`

GetClientReadIopsOk returns a tuple with the ClientReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReadIops

`func (o *PoolStat) SetClientReadIops(v float64)`

SetClientReadIops sets ClientReadIops field to given value.

### HasClientReadIops

`func (o *PoolStat) HasClientReadIops() bool`

HasClientReadIops returns a boolean if a field has been set.

### GetClientReadLatencyUs

`func (o *PoolStat) GetClientReadLatencyUs() float64`

GetClientReadLatencyUs returns the ClientReadLatencyUs field if non-nil, zero value otherwise.

### GetClientReadLatencyUsOk

`func (o *PoolStat) GetClientReadLatencyUsOk() (*float64, bool)`

GetClientReadLatencyUsOk returns a tuple with the ClientReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReadLatencyUs

`func (o *PoolStat) SetClientReadLatencyUs(v float64)`

SetClientReadLatencyUs sets ClientReadLatencyUs field to given value.

### HasClientReadLatencyUs

`func (o *PoolStat) HasClientReadLatencyUs() bool`

HasClientReadLatencyUs returns a boolean if a field has been set.

### GetClientWriteBandwidthKbyte

`func (o *PoolStat) GetClientWriteBandwidthKbyte() float64`

GetClientWriteBandwidthKbyte returns the ClientWriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetClientWriteBandwidthKbyteOk

`func (o *PoolStat) GetClientWriteBandwidthKbyteOk() (*float64, bool)`

GetClientWriteBandwidthKbyteOk returns a tuple with the ClientWriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientWriteBandwidthKbyte

`func (o *PoolStat) SetClientWriteBandwidthKbyte(v float64)`

SetClientWriteBandwidthKbyte sets ClientWriteBandwidthKbyte field to given value.

### HasClientWriteBandwidthKbyte

`func (o *PoolStat) HasClientWriteBandwidthKbyte() bool`

HasClientWriteBandwidthKbyte returns a boolean if a field has been set.

### GetClientWriteIoSizeKbyte

`func (o *PoolStat) GetClientWriteIoSizeKbyte() float64`

GetClientWriteIoSizeKbyte returns the ClientWriteIoSizeKbyte field if non-nil, zero value otherwise.

### GetClientWriteIoSizeKbyteOk

`func (o *PoolStat) GetClientWriteIoSizeKbyteOk() (*float64, bool)`

GetClientWriteIoSizeKbyteOk returns a tuple with the ClientWriteIoSizeKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientWriteIoSizeKbyte

`func (o *PoolStat) SetClientWriteIoSizeKbyte(v float64)`

SetClientWriteIoSizeKbyte sets ClientWriteIoSizeKbyte field to given value.

### HasClientWriteIoSizeKbyte

`func (o *PoolStat) HasClientWriteIoSizeKbyte() bool`

HasClientWriteIoSizeKbyte returns a boolean if a field has been set.

### GetClientWriteIops

`func (o *PoolStat) GetClientWriteIops() float64`

GetClientWriteIops returns the ClientWriteIops field if non-nil, zero value otherwise.

### GetClientWriteIopsOk

`func (o *PoolStat) GetClientWriteIopsOk() (*float64, bool)`

GetClientWriteIopsOk returns a tuple with the ClientWriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientWriteIops

`func (o *PoolStat) SetClientWriteIops(v float64)`

SetClientWriteIops sets ClientWriteIops field to given value.

### HasClientWriteIops

`func (o *PoolStat) HasClientWriteIops() bool`

HasClientWriteIops returns a boolean if a field has been set.

### GetClientWriteLatencyUs

`func (o *PoolStat) GetClientWriteLatencyUs() float64`

GetClientWriteLatencyUs returns the ClientWriteLatencyUs field if non-nil, zero value otherwise.

### GetClientWriteLatencyUsOk

`func (o *PoolStat) GetClientWriteLatencyUsOk() (*float64, bool)`

GetClientWriteLatencyUsOk returns a tuple with the ClientWriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientWriteLatencyUs

`func (o *PoolStat) SetClientWriteLatencyUs(v float64)`

SetClientWriteLatencyUs sets ClientWriteLatencyUs field to given value.

### HasClientWriteLatencyUs

`func (o *PoolStat) HasClientWriteLatencyUs() bool`

HasClientWriteLatencyUs returns a boolean if a field has been set.

### GetCompressRatio

`func (o *PoolStat) GetCompressRatio() float64`

GetCompressRatio returns the CompressRatio field if non-nil, zero value otherwise.

### GetCompressRatioOk

`func (o *PoolStat) GetCompressRatioOk() (*float64, bool)`

GetCompressRatioOk returns a tuple with the CompressRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressRatio

`func (o *PoolStat) SetCompressRatio(v float64)`

SetCompressRatio sets CompressRatio field to given value.

### HasCompressRatio

`func (o *PoolStat) HasCompressRatio() bool`

HasCompressRatio returns a boolean if a field has been set.

### GetCompressedKbyte

`func (o *PoolStat) GetCompressedKbyte() float64`

GetCompressedKbyte returns the CompressedKbyte field if non-nil, zero value otherwise.

### GetCompressedKbyteOk

`func (o *PoolStat) GetCompressedKbyteOk() (*float64, bool)`

GetCompressedKbyteOk returns a tuple with the CompressedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedKbyte

`func (o *PoolStat) SetCompressedKbyte(v float64)`

SetCompressedKbyte sets CompressedKbyte field to given value.

### HasCompressedKbyte

`func (o *PoolStat) HasCompressedKbyte() bool`

HasCompressedKbyte returns a boolean if a field has been set.

### GetCreate

`func (o *PoolStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *PoolStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *PoolStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *PoolStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCreatingNum

`func (o *PoolStat) GetCreatingNum() int64`

GetCreatingNum returns the CreatingNum field if non-nil, zero value otherwise.

### GetCreatingNumOk

`func (o *PoolStat) GetCreatingNumOk() (*int64, bool)`

GetCreatingNumOk returns a tuple with the CreatingNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatingNum

`func (o *PoolStat) SetCreatingNum(v int64)`

SetCreatingNum sets CreatingNum field to given value.

### HasCreatingNum

`func (o *PoolStat) HasCreatingNum() bool`

HasCreatingNum returns a boolean if a field has been set.

### GetDataKbyte

`func (o *PoolStat) GetDataKbyte() int64`

GetDataKbyte returns the DataKbyte field if non-nil, zero value otherwise.

### GetDataKbyteOk

`func (o *PoolStat) GetDataKbyteOk() (*int64, bool)`

GetDataKbyteOk returns a tuple with the DataKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataKbyte

`func (o *PoolStat) SetDataKbyte(v int64)`

SetDataKbyte sets DataKbyte field to given value.

### HasDataKbyte

`func (o *PoolStat) HasDataKbyte() bool`

HasDataKbyte returns a boolean if a field has been set.

### GetDegradedNum

`func (o *PoolStat) GetDegradedNum() int64`

GetDegradedNum returns the DegradedNum field if non-nil, zero value otherwise.

### GetDegradedNumOk

`func (o *PoolStat) GetDegradedNumOk() (*int64, bool)`

GetDegradedNumOk returns a tuple with the DegradedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedNum

`func (o *PoolStat) SetDegradedNum(v int64)`

SetDegradedNum sets DegradedNum field to given value.

### HasDegradedNum

`func (o *PoolStat) HasDegradedNum() bool`

HasDegradedNum returns a boolean if a field has been set.

### GetDegradedPercent

`func (o *PoolStat) GetDegradedPercent() float64`

GetDegradedPercent returns the DegradedPercent field if non-nil, zero value otherwise.

### GetDegradedPercentOk

`func (o *PoolStat) GetDegradedPercentOk() (*float64, bool)`

GetDegradedPercentOk returns a tuple with the DegradedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedPercent

`func (o *PoolStat) SetDegradedPercent(v float64)`

SetDegradedPercent sets DegradedPercent field to given value.

### HasDegradedPercent

`func (o *PoolStat) HasDegradedPercent() bool`

HasDegradedPercent returns a boolean if a field has been set.

### GetErrorKbyte

`func (o *PoolStat) GetErrorKbyte() int64`

GetErrorKbyte returns the ErrorKbyte field if non-nil, zero value otherwise.

### GetErrorKbyteOk

`func (o *PoolStat) GetErrorKbyteOk() (*int64, bool)`

GetErrorKbyteOk returns a tuple with the ErrorKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorKbyte

`func (o *PoolStat) SetErrorKbyte(v int64)`

SetErrorKbyte sets ErrorKbyte field to given value.

### HasErrorKbyte

`func (o *PoolStat) HasErrorKbyte() bool`

HasErrorKbyte returns a boolean if a field has been set.

### GetFlushReadBandwidthKbyte

`func (o *PoolStat) GetFlushReadBandwidthKbyte() float64`

GetFlushReadBandwidthKbyte returns the FlushReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetFlushReadBandwidthKbyteOk

`func (o *PoolStat) GetFlushReadBandwidthKbyteOk() (*float64, bool)`

GetFlushReadBandwidthKbyteOk returns a tuple with the FlushReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushReadBandwidthKbyte

`func (o *PoolStat) SetFlushReadBandwidthKbyte(v float64)`

SetFlushReadBandwidthKbyte sets FlushReadBandwidthKbyte field to given value.

### HasFlushReadBandwidthKbyte

`func (o *PoolStat) HasFlushReadBandwidthKbyte() bool`

HasFlushReadBandwidthKbyte returns a boolean if a field has been set.

### GetFlushReadIoSizeKbyte

`func (o *PoolStat) GetFlushReadIoSizeKbyte() float64`

GetFlushReadIoSizeKbyte returns the FlushReadIoSizeKbyte field if non-nil, zero value otherwise.

### GetFlushReadIoSizeKbyteOk

`func (o *PoolStat) GetFlushReadIoSizeKbyteOk() (*float64, bool)`

GetFlushReadIoSizeKbyteOk returns a tuple with the FlushReadIoSizeKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushReadIoSizeKbyte

`func (o *PoolStat) SetFlushReadIoSizeKbyte(v float64)`

SetFlushReadIoSizeKbyte sets FlushReadIoSizeKbyte field to given value.

### HasFlushReadIoSizeKbyte

`func (o *PoolStat) HasFlushReadIoSizeKbyte() bool`

HasFlushReadIoSizeKbyte returns a boolean if a field has been set.

### GetFlushReadIops

`func (o *PoolStat) GetFlushReadIops() float64`

GetFlushReadIops returns the FlushReadIops field if non-nil, zero value otherwise.

### GetFlushReadIopsOk

`func (o *PoolStat) GetFlushReadIopsOk() (*float64, bool)`

GetFlushReadIopsOk returns a tuple with the FlushReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushReadIops

`func (o *PoolStat) SetFlushReadIops(v float64)`

SetFlushReadIops sets FlushReadIops field to given value.

### HasFlushReadIops

`func (o *PoolStat) HasFlushReadIops() bool`

HasFlushReadIops returns a boolean if a field has been set.

### GetFlushReadLatencyUs

`func (o *PoolStat) GetFlushReadLatencyUs() float64`

GetFlushReadLatencyUs returns the FlushReadLatencyUs field if non-nil, zero value otherwise.

### GetFlushReadLatencyUsOk

`func (o *PoolStat) GetFlushReadLatencyUsOk() (*float64, bool)`

GetFlushReadLatencyUsOk returns a tuple with the FlushReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushReadLatencyUs

`func (o *PoolStat) SetFlushReadLatencyUs(v float64)`

SetFlushReadLatencyUs sets FlushReadLatencyUs field to given value.

### HasFlushReadLatencyUs

`func (o *PoolStat) HasFlushReadLatencyUs() bool`

HasFlushReadLatencyUs returns a boolean if a field has been set.

### GetFlushWriteBandwidthKbyte

`func (o *PoolStat) GetFlushWriteBandwidthKbyte() float64`

GetFlushWriteBandwidthKbyte returns the FlushWriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetFlushWriteBandwidthKbyteOk

`func (o *PoolStat) GetFlushWriteBandwidthKbyteOk() (*float64, bool)`

GetFlushWriteBandwidthKbyteOk returns a tuple with the FlushWriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushWriteBandwidthKbyte

`func (o *PoolStat) SetFlushWriteBandwidthKbyte(v float64)`

SetFlushWriteBandwidthKbyte sets FlushWriteBandwidthKbyte field to given value.

### HasFlushWriteBandwidthKbyte

`func (o *PoolStat) HasFlushWriteBandwidthKbyte() bool`

HasFlushWriteBandwidthKbyte returns a boolean if a field has been set.

### GetFlushWriteIoSizeKbyte

`func (o *PoolStat) GetFlushWriteIoSizeKbyte() float64`

GetFlushWriteIoSizeKbyte returns the FlushWriteIoSizeKbyte field if non-nil, zero value otherwise.

### GetFlushWriteIoSizeKbyteOk

`func (o *PoolStat) GetFlushWriteIoSizeKbyteOk() (*float64, bool)`

GetFlushWriteIoSizeKbyteOk returns a tuple with the FlushWriteIoSizeKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushWriteIoSizeKbyte

`func (o *PoolStat) SetFlushWriteIoSizeKbyte(v float64)`

SetFlushWriteIoSizeKbyte sets FlushWriteIoSizeKbyte field to given value.

### HasFlushWriteIoSizeKbyte

`func (o *PoolStat) HasFlushWriteIoSizeKbyte() bool`

HasFlushWriteIoSizeKbyte returns a boolean if a field has been set.

### GetFlushWriteIops

`func (o *PoolStat) GetFlushWriteIops() float64`

GetFlushWriteIops returns the FlushWriteIops field if non-nil, zero value otherwise.

### GetFlushWriteIopsOk

`func (o *PoolStat) GetFlushWriteIopsOk() (*float64, bool)`

GetFlushWriteIopsOk returns a tuple with the FlushWriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushWriteIops

`func (o *PoolStat) SetFlushWriteIops(v float64)`

SetFlushWriteIops sets FlushWriteIops field to given value.

### HasFlushWriteIops

`func (o *PoolStat) HasFlushWriteIops() bool`

HasFlushWriteIops returns a boolean if a field has been set.

### GetFlushWriteLatencyUs

`func (o *PoolStat) GetFlushWriteLatencyUs() float64`

GetFlushWriteLatencyUs returns the FlushWriteLatencyUs field if non-nil, zero value otherwise.

### GetFlushWriteLatencyUsOk

`func (o *PoolStat) GetFlushWriteLatencyUsOk() (*float64, bool)`

GetFlushWriteLatencyUsOk returns a tuple with the FlushWriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushWriteLatencyUs

`func (o *PoolStat) SetFlushWriteLatencyUs(v float64)`

SetFlushWriteLatencyUs sets FlushWriteLatencyUs field to given value.

### HasFlushWriteLatencyUs

`func (o *PoolStat) HasFlushWriteLatencyUs() bool`

HasFlushWriteLatencyUs returns a boolean if a field has been set.

### GetGarbageRatio

`func (o *PoolStat) GetGarbageRatio() float64`

GetGarbageRatio returns the GarbageRatio field if non-nil, zero value otherwise.

### GetGarbageRatioOk

`func (o *PoolStat) GetGarbageRatioOk() (*float64, bool)`

GetGarbageRatioOk returns a tuple with the GarbageRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarbageRatio

`func (o *PoolStat) SetGarbageRatio(v float64)`

SetGarbageRatio sets GarbageRatio field to given value.

### HasGarbageRatio

`func (o *PoolStat) HasGarbageRatio() bool`

HasGarbageRatio returns a boolean if a field has been set.

### GetGcReadBandwidthKbyte

`func (o *PoolStat) GetGcReadBandwidthKbyte() float64`

GetGcReadBandwidthKbyte returns the GcReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetGcReadBandwidthKbyteOk

`func (o *PoolStat) GetGcReadBandwidthKbyteOk() (*float64, bool)`

GetGcReadBandwidthKbyteOk returns a tuple with the GcReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcReadBandwidthKbyte

`func (o *PoolStat) SetGcReadBandwidthKbyte(v float64)`

SetGcReadBandwidthKbyte sets GcReadBandwidthKbyte field to given value.

### HasGcReadBandwidthKbyte

`func (o *PoolStat) HasGcReadBandwidthKbyte() bool`

HasGcReadBandwidthKbyte returns a boolean if a field has been set.

### GetGcReadIoSizeKbyte

`func (o *PoolStat) GetGcReadIoSizeKbyte() float64`

GetGcReadIoSizeKbyte returns the GcReadIoSizeKbyte field if non-nil, zero value otherwise.

### GetGcReadIoSizeKbyteOk

`func (o *PoolStat) GetGcReadIoSizeKbyteOk() (*float64, bool)`

GetGcReadIoSizeKbyteOk returns a tuple with the GcReadIoSizeKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcReadIoSizeKbyte

`func (o *PoolStat) SetGcReadIoSizeKbyte(v float64)`

SetGcReadIoSizeKbyte sets GcReadIoSizeKbyte field to given value.

### HasGcReadIoSizeKbyte

`func (o *PoolStat) HasGcReadIoSizeKbyte() bool`

HasGcReadIoSizeKbyte returns a boolean if a field has been set.

### GetGcReadIops

`func (o *PoolStat) GetGcReadIops() float64`

GetGcReadIops returns the GcReadIops field if non-nil, zero value otherwise.

### GetGcReadIopsOk

`func (o *PoolStat) GetGcReadIopsOk() (*float64, bool)`

GetGcReadIopsOk returns a tuple with the GcReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcReadIops

`func (o *PoolStat) SetGcReadIops(v float64)`

SetGcReadIops sets GcReadIops field to given value.

### HasGcReadIops

`func (o *PoolStat) HasGcReadIops() bool`

HasGcReadIops returns a boolean if a field has been set.

### GetGcReadLatencyUs

`func (o *PoolStat) GetGcReadLatencyUs() float64`

GetGcReadLatencyUs returns the GcReadLatencyUs field if non-nil, zero value otherwise.

### GetGcReadLatencyUsOk

`func (o *PoolStat) GetGcReadLatencyUsOk() (*float64, bool)`

GetGcReadLatencyUsOk returns a tuple with the GcReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcReadLatencyUs

`func (o *PoolStat) SetGcReadLatencyUs(v float64)`

SetGcReadLatencyUs sets GcReadLatencyUs field to given value.

### HasGcReadLatencyUs

`func (o *PoolStat) HasGcReadLatencyUs() bool`

HasGcReadLatencyUs returns a boolean if a field has been set.

### GetGcRemoveBandwidthKbyte

`func (o *PoolStat) GetGcRemoveBandwidthKbyte() float64`

GetGcRemoveBandwidthKbyte returns the GcRemoveBandwidthKbyte field if non-nil, zero value otherwise.

### GetGcRemoveBandwidthKbyteOk

`func (o *PoolStat) GetGcRemoveBandwidthKbyteOk() (*float64, bool)`

GetGcRemoveBandwidthKbyteOk returns a tuple with the GcRemoveBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcRemoveBandwidthKbyte

`func (o *PoolStat) SetGcRemoveBandwidthKbyte(v float64)`

SetGcRemoveBandwidthKbyte sets GcRemoveBandwidthKbyte field to given value.

### HasGcRemoveBandwidthKbyte

`func (o *PoolStat) HasGcRemoveBandwidthKbyte() bool`

HasGcRemoveBandwidthKbyte returns a boolean if a field has been set.

### GetGcRemoveIoSizeKbyte

`func (o *PoolStat) GetGcRemoveIoSizeKbyte() float64`

GetGcRemoveIoSizeKbyte returns the GcRemoveIoSizeKbyte field if non-nil, zero value otherwise.

### GetGcRemoveIoSizeKbyteOk

`func (o *PoolStat) GetGcRemoveIoSizeKbyteOk() (*float64, bool)`

GetGcRemoveIoSizeKbyteOk returns a tuple with the GcRemoveIoSizeKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcRemoveIoSizeKbyte

`func (o *PoolStat) SetGcRemoveIoSizeKbyte(v float64)`

SetGcRemoveIoSizeKbyte sets GcRemoveIoSizeKbyte field to given value.

### HasGcRemoveIoSizeKbyte

`func (o *PoolStat) HasGcRemoveIoSizeKbyte() bool`

HasGcRemoveIoSizeKbyte returns a boolean if a field has been set.

### GetGcRemoveIops

`func (o *PoolStat) GetGcRemoveIops() float64`

GetGcRemoveIops returns the GcRemoveIops field if non-nil, zero value otherwise.

### GetGcRemoveIopsOk

`func (o *PoolStat) GetGcRemoveIopsOk() (*float64, bool)`

GetGcRemoveIopsOk returns a tuple with the GcRemoveIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcRemoveIops

`func (o *PoolStat) SetGcRemoveIops(v float64)`

SetGcRemoveIops sets GcRemoveIops field to given value.

### HasGcRemoveIops

`func (o *PoolStat) HasGcRemoveIops() bool`

HasGcRemoveIops returns a boolean if a field has been set.

### GetGcWriteBandwidthKbyte

`func (o *PoolStat) GetGcWriteBandwidthKbyte() float64`

GetGcWriteBandwidthKbyte returns the GcWriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetGcWriteBandwidthKbyteOk

`func (o *PoolStat) GetGcWriteBandwidthKbyteOk() (*float64, bool)`

GetGcWriteBandwidthKbyteOk returns a tuple with the GcWriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcWriteBandwidthKbyte

`func (o *PoolStat) SetGcWriteBandwidthKbyte(v float64)`

SetGcWriteBandwidthKbyte sets GcWriteBandwidthKbyte field to given value.

### HasGcWriteBandwidthKbyte

`func (o *PoolStat) HasGcWriteBandwidthKbyte() bool`

HasGcWriteBandwidthKbyte returns a boolean if a field has been set.

### GetGcWriteIoSizeKbyte

`func (o *PoolStat) GetGcWriteIoSizeKbyte() float64`

GetGcWriteIoSizeKbyte returns the GcWriteIoSizeKbyte field if non-nil, zero value otherwise.

### GetGcWriteIoSizeKbyteOk

`func (o *PoolStat) GetGcWriteIoSizeKbyteOk() (*float64, bool)`

GetGcWriteIoSizeKbyteOk returns a tuple with the GcWriteIoSizeKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcWriteIoSizeKbyte

`func (o *PoolStat) SetGcWriteIoSizeKbyte(v float64)`

SetGcWriteIoSizeKbyte sets GcWriteIoSizeKbyte field to given value.

### HasGcWriteIoSizeKbyte

`func (o *PoolStat) HasGcWriteIoSizeKbyte() bool`

HasGcWriteIoSizeKbyte returns a boolean if a field has been set.

### GetGcWriteIops

`func (o *PoolStat) GetGcWriteIops() float64`

GetGcWriteIops returns the GcWriteIops field if non-nil, zero value otherwise.

### GetGcWriteIopsOk

`func (o *PoolStat) GetGcWriteIopsOk() (*float64, bool)`

GetGcWriteIopsOk returns a tuple with the GcWriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcWriteIops

`func (o *PoolStat) SetGcWriteIops(v float64)`

SetGcWriteIops sets GcWriteIops field to given value.

### HasGcWriteIops

`func (o *PoolStat) HasGcWriteIops() bool`

HasGcWriteIops returns a boolean if a field has been set.

### GetGcWriteLatencyUs

`func (o *PoolStat) GetGcWriteLatencyUs() float64`

GetGcWriteLatencyUs returns the GcWriteLatencyUs field if non-nil, zero value otherwise.

### GetGcWriteLatencyUsOk

`func (o *PoolStat) GetGcWriteLatencyUsOk() (*float64, bool)`

GetGcWriteLatencyUsOk returns a tuple with the GcWriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcWriteLatencyUs

`func (o *PoolStat) SetGcWriteLatencyUs(v float64)`

SetGcWriteLatencyUs sets GcWriteLatencyUs field to given value.

### HasGcWriteLatencyUs

`func (o *PoolStat) HasGcWriteLatencyUs() bool`

HasGcWriteLatencyUs returns a boolean if a field has been set.

### GetHealthyNum

`func (o *PoolStat) GetHealthyNum() int64`

GetHealthyNum returns the HealthyNum field if non-nil, zero value otherwise.

### GetHealthyNumOk

`func (o *PoolStat) GetHealthyNumOk() (*int64, bool)`

GetHealthyNumOk returns a tuple with the HealthyNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyNum

`func (o *PoolStat) SetHealthyNum(v int64)`

SetHealthyNum sets HealthyNum field to given value.

### HasHealthyNum

`func (o *PoolStat) HasHealthyNum() bool`

HasHealthyNum returns a boolean if a field has been set.

### GetHealthyPercent

`func (o *PoolStat) GetHealthyPercent() float64`

GetHealthyPercent returns the HealthyPercent field if non-nil, zero value otherwise.

### GetHealthyPercentOk

`func (o *PoolStat) GetHealthyPercentOk() (*float64, bool)`

GetHealthyPercentOk returns a tuple with the HealthyPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyPercent

`func (o *PoolStat) SetHealthyPercent(v float64)`

SetHealthyPercent sets HealthyPercent field to given value.

### HasHealthyPercent

`func (o *PoolStat) HasHealthyPercent() bool`

HasHealthyPercent returns a boolean if a field has been set.

### GetMaxAvailKbyte

`func (o *PoolStat) GetMaxAvailKbyte() int64`

GetMaxAvailKbyte returns the MaxAvailKbyte field if non-nil, zero value otherwise.

### GetMaxAvailKbyteOk

`func (o *PoolStat) GetMaxAvailKbyteOk() (*int64, bool)`

GetMaxAvailKbyteOk returns a tuple with the MaxAvailKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAvailKbyte

`func (o *PoolStat) SetMaxAvailKbyte(v int64)`

SetMaxAvailKbyte sets MaxAvailKbyte field to given value.

### HasMaxAvailKbyte

`func (o *PoolStat) HasMaxAvailKbyte() bool`

HasMaxAvailKbyte returns a boolean if a field has been set.

### GetMinGarbageKbyte

`func (o *PoolStat) GetMinGarbageKbyte() int64`

GetMinGarbageKbyte returns the MinGarbageKbyte field if non-nil, zero value otherwise.

### GetMinGarbageKbyteOk

`func (o *PoolStat) GetMinGarbageKbyteOk() (*int64, bool)`

GetMinGarbageKbyteOk returns a tuple with the MinGarbageKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGarbageKbyte

`func (o *PoolStat) SetMinGarbageKbyte(v int64)`

SetMinGarbageKbyte sets MinGarbageKbyte field to given value.

### HasMinGarbageKbyte

`func (o *PoolStat) HasMinGarbageKbyte() bool`

HasMinGarbageKbyte returns a boolean if a field has been set.

### GetOmapTotalKbyte

`func (o *PoolStat) GetOmapTotalKbyte() float64`

GetOmapTotalKbyte returns the OmapTotalKbyte field if non-nil, zero value otherwise.

### GetOmapTotalKbyteOk

`func (o *PoolStat) GetOmapTotalKbyteOk() (*float64, bool)`

GetOmapTotalKbyteOk returns a tuple with the OmapTotalKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmapTotalKbyte

`func (o *PoolStat) SetOmapTotalKbyte(v float64)`

SetOmapTotalKbyte sets OmapTotalKbyte field to given value.

### HasOmapTotalKbyte

`func (o *PoolStat) HasOmapTotalKbyte() bool`

HasOmapTotalKbyte returns a boolean if a field has been set.

### GetOmapUsedKbyte

`func (o *PoolStat) GetOmapUsedKbyte() float64`

GetOmapUsedKbyte returns the OmapUsedKbyte field if non-nil, zero value otherwise.

### GetOmapUsedKbyteOk

`func (o *PoolStat) GetOmapUsedKbyteOk() (*float64, bool)`

GetOmapUsedKbyteOk returns a tuple with the OmapUsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmapUsedKbyte

`func (o *PoolStat) SetOmapUsedKbyte(v float64)`

SetOmapUsedKbyte sets OmapUsedKbyte field to given value.

### HasOmapUsedKbyte

`func (o *PoolStat) HasOmapUsedKbyte() bool`

HasOmapUsedKbyte returns a boolean if a field has been set.

### GetOmapUsedPercent

`func (o *PoolStat) GetOmapUsedPercent() float64`

GetOmapUsedPercent returns the OmapUsedPercent field if non-nil, zero value otherwise.

### GetOmapUsedPercentOk

`func (o *PoolStat) GetOmapUsedPercentOk() (*float64, bool)`

GetOmapUsedPercentOk returns a tuple with the OmapUsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmapUsedPercent

`func (o *PoolStat) SetOmapUsedPercent(v float64)`

SetOmapUsedPercent sets OmapUsedPercent field to given value.

### HasOmapUsedPercent

`func (o *PoolStat) HasOmapUsedPercent() bool`

HasOmapUsedPercent returns a boolean if a field has been set.

### GetOsdCapacityUnbalance

`func (o *PoolStat) GetOsdCapacityUnbalance() bool`

GetOsdCapacityUnbalance returns the OsdCapacityUnbalance field if non-nil, zero value otherwise.

### GetOsdCapacityUnbalanceOk

`func (o *PoolStat) GetOsdCapacityUnbalanceOk() (*bool, bool)`

GetOsdCapacityUnbalanceOk returns a tuple with the OsdCapacityUnbalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdCapacityUnbalance

`func (o *PoolStat) SetOsdCapacityUnbalance(v bool)`

SetOsdCapacityUnbalance sets OsdCapacityUnbalance field to given value.

### HasOsdCapacityUnbalance

`func (o *PoolStat) HasOsdCapacityUnbalance() bool`

HasOsdCapacityUnbalance returns a boolean if a field has been set.

### GetPoolCapacityUsage

`func (o *PoolStat) GetPoolCapacityUsage() int64`

GetPoolCapacityUsage returns the PoolCapacityUsage field if non-nil, zero value otherwise.

### GetPoolCapacityUsageOk

`func (o *PoolStat) GetPoolCapacityUsageOk() (*int64, bool)`

GetPoolCapacityUsageOk returns a tuple with the PoolCapacityUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolCapacityUsage

`func (o *PoolStat) SetPoolCapacityUsage(v int64)`

SetPoolCapacityUsage sets PoolCapacityUsage field to given value.

### HasPoolCapacityUsage

`func (o *PoolStat) HasPoolCapacityUsage() bool`

HasPoolCapacityUsage returns a boolean if a field has been set.

### GetReadBandwidthKbyte

`func (o *PoolStat) GetReadBandwidthKbyte() float64`

GetReadBandwidthKbyte returns the ReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetReadBandwidthKbyteOk

`func (o *PoolStat) GetReadBandwidthKbyteOk() (*float64, bool)`

GetReadBandwidthKbyteOk returns a tuple with the ReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadBandwidthKbyte

`func (o *PoolStat) SetReadBandwidthKbyte(v float64)`

SetReadBandwidthKbyte sets ReadBandwidthKbyte field to given value.

### HasReadBandwidthKbyte

`func (o *PoolStat) HasReadBandwidthKbyte() bool`

HasReadBandwidthKbyte returns a boolean if a field has been set.

### GetReadCacheHitRate

`func (o *PoolStat) GetReadCacheHitRate() float64`

GetReadCacheHitRate returns the ReadCacheHitRate field if non-nil, zero value otherwise.

### GetReadCacheHitRateOk

`func (o *PoolStat) GetReadCacheHitRateOk() (*float64, bool)`

GetReadCacheHitRateOk returns a tuple with the ReadCacheHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCacheHitRate

`func (o *PoolStat) SetReadCacheHitRate(v float64)`

SetReadCacheHitRate sets ReadCacheHitRate field to given value.

### HasReadCacheHitRate

`func (o *PoolStat) HasReadCacheHitRate() bool`

HasReadCacheHitRate returns a boolean if a field has been set.

### GetReadCacheKbyte

`func (o *PoolStat) GetReadCacheKbyte() float64`

GetReadCacheKbyte returns the ReadCacheKbyte field if non-nil, zero value otherwise.

### GetReadCacheKbyteOk

`func (o *PoolStat) GetReadCacheKbyteOk() (*float64, bool)`

GetReadCacheKbyteOk returns a tuple with the ReadCacheKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCacheKbyte

`func (o *PoolStat) SetReadCacheKbyte(v float64)`

SetReadCacheKbyte sets ReadCacheKbyte field to given value.

### HasReadCacheKbyte

`func (o *PoolStat) HasReadCacheKbyte() bool`

HasReadCacheKbyte returns a boolean if a field has been set.

### GetReadCachePercent

`func (o *PoolStat) GetReadCachePercent() float64`

GetReadCachePercent returns the ReadCachePercent field if non-nil, zero value otherwise.

### GetReadCachePercentOk

`func (o *PoolStat) GetReadCachePercentOk() (*float64, bool)`

GetReadCachePercentOk returns a tuple with the ReadCachePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCachePercent

`func (o *PoolStat) SetReadCachePercent(v float64)`

SetReadCachePercent sets ReadCachePercent field to given value.

### HasReadCachePercent

`func (o *PoolStat) HasReadCachePercent() bool`

HasReadCachePercent returns a boolean if a field has been set.

### GetReadIoSizeKbyte

`func (o *PoolStat) GetReadIoSizeKbyte() float64`

GetReadIoSizeKbyte returns the ReadIoSizeKbyte field if non-nil, zero value otherwise.

### GetReadIoSizeKbyteOk

`func (o *PoolStat) GetReadIoSizeKbyteOk() (*float64, bool)`

GetReadIoSizeKbyteOk returns a tuple with the ReadIoSizeKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIoSizeKbyte

`func (o *PoolStat) SetReadIoSizeKbyte(v float64)`

SetReadIoSizeKbyte sets ReadIoSizeKbyte field to given value.

### HasReadIoSizeKbyte

`func (o *PoolStat) HasReadIoSizeKbyte() bool`

HasReadIoSizeKbyte returns a boolean if a field has been set.

### GetReadIops

`func (o *PoolStat) GetReadIops() float64`

GetReadIops returns the ReadIops field if non-nil, zero value otherwise.

### GetReadIopsOk

`func (o *PoolStat) GetReadIopsOk() (*float64, bool)`

GetReadIopsOk returns a tuple with the ReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIops

`func (o *PoolStat) SetReadIops(v float64)`

SetReadIops sets ReadIops field to given value.

### HasReadIops

`func (o *PoolStat) HasReadIops() bool`

HasReadIops returns a boolean if a field has been set.

### GetReadLatencyUs

`func (o *PoolStat) GetReadLatencyUs() float64`

GetReadLatencyUs returns the ReadLatencyUs field if non-nil, zero value otherwise.

### GetReadLatencyUsOk

`func (o *PoolStat) GetReadLatencyUsOk() (*float64, bool)`

GetReadLatencyUsOk returns a tuple with the ReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadLatencyUs

`func (o *PoolStat) SetReadLatencyUs(v float64)`

SetReadLatencyUs sets ReadLatencyUs field to given value.

### HasReadLatencyUs

`func (o *PoolStat) HasReadLatencyUs() bool`

HasReadLatencyUs returns a boolean if a field has been set.

### GetRecoveryBandwidthKbyte

`func (o *PoolStat) GetRecoveryBandwidthKbyte() float64`

GetRecoveryBandwidthKbyte returns the RecoveryBandwidthKbyte field if non-nil, zero value otherwise.

### GetRecoveryBandwidthKbyteOk

`func (o *PoolStat) GetRecoveryBandwidthKbyteOk() (*float64, bool)`

GetRecoveryBandwidthKbyteOk returns a tuple with the RecoveryBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryBandwidthKbyte

`func (o *PoolStat) SetRecoveryBandwidthKbyte(v float64)`

SetRecoveryBandwidthKbyte sets RecoveryBandwidthKbyte field to given value.

### HasRecoveryBandwidthKbyte

`func (o *PoolStat) HasRecoveryBandwidthKbyte() bool`

HasRecoveryBandwidthKbyte returns a boolean if a field has been set.

### GetRecoveryIops

`func (o *PoolStat) GetRecoveryIops() float64`

GetRecoveryIops returns the RecoveryIops field if non-nil, zero value otherwise.

### GetRecoveryIopsOk

`func (o *PoolStat) GetRecoveryIopsOk() (*float64, bool)`

GetRecoveryIopsOk returns a tuple with the RecoveryIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryIops

`func (o *PoolStat) SetRecoveryIops(v float64)`

SetRecoveryIops sets RecoveryIops field to given value.

### HasRecoveryIops

`func (o *PoolStat) HasRecoveryIops() bool`

HasRecoveryIops returns a boolean if a field has been set.

### GetRecoveryNum

`func (o *PoolStat) GetRecoveryNum() int64`

GetRecoveryNum returns the RecoveryNum field if non-nil, zero value otherwise.

### GetRecoveryNumOk

`func (o *PoolStat) GetRecoveryNumOk() (*int64, bool)`

GetRecoveryNumOk returns a tuple with the RecoveryNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryNum

`func (o *PoolStat) SetRecoveryNum(v int64)`

SetRecoveryNum sets RecoveryNum field to given value.

### HasRecoveryNum

`func (o *PoolStat) HasRecoveryNum() bool`

HasRecoveryNum returns a boolean if a field has been set.

### GetRecoveryPercent

`func (o *PoolStat) GetRecoveryPercent() float64`

GetRecoveryPercent returns the RecoveryPercent field if non-nil, zero value otherwise.

### GetRecoveryPercentOk

`func (o *PoolStat) GetRecoveryPercentOk() (*float64, bool)`

GetRecoveryPercentOk returns a tuple with the RecoveryPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPercent

`func (o *PoolStat) SetRecoveryPercent(v float64)`

SetRecoveryPercent sets RecoveryPercent field to given value.

### HasRecoveryPercent

`func (o *PoolStat) HasRecoveryPercent() bool`

HasRecoveryPercent returns a boolean if a field has been set.

### GetRecoveryRemainSecond

`func (o *PoolStat) GetRecoveryRemainSecond() int64`

GetRecoveryRemainSecond returns the RecoveryRemainSecond field if non-nil, zero value otherwise.

### GetRecoveryRemainSecondOk

`func (o *PoolStat) GetRecoveryRemainSecondOk() (*int64, bool)`

GetRecoveryRemainSecondOk returns a tuple with the RecoveryRemainSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryRemainSecond

`func (o *PoolStat) SetRecoveryRemainSecond(v int64)`

SetRecoveryRemainSecond sets RecoveryRemainSecond field to given value.

### HasRecoveryRemainSecond

`func (o *PoolStat) HasRecoveryRemainSecond() bool`

HasRecoveryRemainSecond returns a boolean if a field has been set.

### GetRemoveBandwidthKbyte

`func (o *PoolStat) GetRemoveBandwidthKbyte() float64`

GetRemoveBandwidthKbyte returns the RemoveBandwidthKbyte field if non-nil, zero value otherwise.

### GetRemoveBandwidthKbyteOk

`func (o *PoolStat) GetRemoveBandwidthKbyteOk() (*float64, bool)`

GetRemoveBandwidthKbyteOk returns a tuple with the RemoveBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveBandwidthKbyte

`func (o *PoolStat) SetRemoveBandwidthKbyte(v float64)`

SetRemoveBandwidthKbyte sets RemoveBandwidthKbyte field to given value.

### HasRemoveBandwidthKbyte

`func (o *PoolStat) HasRemoveBandwidthKbyte() bool`

HasRemoveBandwidthKbyte returns a boolean if a field has been set.

### GetRemoveIops

`func (o *PoolStat) GetRemoveIops() float64`

GetRemoveIops returns the RemoveIops field if non-nil, zero value otherwise.

### GetRemoveIopsOk

`func (o *PoolStat) GetRemoveIopsOk() (*float64, bool)`

GetRemoveIopsOk returns a tuple with the RemoveIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveIops

`func (o *PoolStat) SetRemoveIops(v float64)`

SetRemoveIops sets RemoveIops field to given value.

### HasRemoveIops

`func (o *PoolStat) HasRemoveIops() bool`

HasRemoveIops returns a boolean if a field has been set.

### GetReservedDataKbyte

`func (o *PoolStat) GetReservedDataKbyte() float64`

GetReservedDataKbyte returns the ReservedDataKbyte field if non-nil, zero value otherwise.

### GetReservedDataKbyteOk

`func (o *PoolStat) GetReservedDataKbyteOk() (*float64, bool)`

GetReservedDataKbyteOk returns a tuple with the ReservedDataKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedDataKbyte

`func (o *PoolStat) SetReservedDataKbyte(v float64)`

SetReservedDataKbyte sets ReservedDataKbyte field to given value.

### HasReservedDataKbyte

`func (o *PoolStat) HasReservedDataKbyte() bool`

HasReservedDataKbyte returns a boolean if a field has been set.

### GetReservedDataPercent

`func (o *PoolStat) GetReservedDataPercent() float64`

GetReservedDataPercent returns the ReservedDataPercent field if non-nil, zero value otherwise.

### GetReservedDataPercentOk

`func (o *PoolStat) GetReservedDataPercentOk() (*float64, bool)`

GetReservedDataPercentOk returns a tuple with the ReservedDataPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedDataPercent

`func (o *PoolStat) SetReservedDataPercent(v float64)`

SetReservedDataPercent sets ReservedDataPercent field to given value.

### HasReservedDataPercent

`func (o *PoolStat) HasReservedDataPercent() bool`

HasReservedDataPercent returns a boolean if a field has been set.

### GetReservedUsedKbyte

`func (o *PoolStat) GetReservedUsedKbyte() float64`

GetReservedUsedKbyte returns the ReservedUsedKbyte field if non-nil, zero value otherwise.

### GetReservedUsedKbyteOk

`func (o *PoolStat) GetReservedUsedKbyteOk() (*float64, bool)`

GetReservedUsedKbyteOk returns a tuple with the ReservedUsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedUsedKbyte

`func (o *PoolStat) SetReservedUsedKbyte(v float64)`

SetReservedUsedKbyte sets ReservedUsedKbyte field to given value.

### HasReservedUsedKbyte

`func (o *PoolStat) HasReservedUsedKbyte() bool`

HasReservedUsedKbyte returns a boolean if a field has been set.

### GetRmwBandwidthKbyte

`func (o *PoolStat) GetRmwBandwidthKbyte() float64`

GetRmwBandwidthKbyte returns the RmwBandwidthKbyte field if non-nil, zero value otherwise.

### GetRmwBandwidthKbyteOk

`func (o *PoolStat) GetRmwBandwidthKbyteOk() (*float64, bool)`

GetRmwBandwidthKbyteOk returns a tuple with the RmwBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmwBandwidthKbyte

`func (o *PoolStat) SetRmwBandwidthKbyte(v float64)`

SetRmwBandwidthKbyte sets RmwBandwidthKbyte field to given value.

### HasRmwBandwidthKbyte

`func (o *PoolStat) HasRmwBandwidthKbyte() bool`

HasRmwBandwidthKbyte returns a boolean if a field has been set.

### GetRmwIops

`func (o *PoolStat) GetRmwIops() float64`

GetRmwIops returns the RmwIops field if non-nil, zero value otherwise.

### GetRmwIopsOk

`func (o *PoolStat) GetRmwIopsOk() (*float64, bool)`

GetRmwIopsOk returns a tuple with the RmwIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmwIops

`func (o *PoolStat) SetRmwIops(v float64)`

SetRmwIops sets RmwIops field to given value.

### HasRmwIops

`func (o *PoolStat) HasRmwIops() bool`

HasRmwIops returns a boolean if a field has been set.

### GetSnapKbyte

`func (o *PoolStat) GetSnapKbyte() int64`

GetSnapKbyte returns the SnapKbyte field if non-nil, zero value otherwise.

### GetSnapKbyteOk

`func (o *PoolStat) GetSnapKbyteOk() (*int64, bool)`

GetSnapKbyteOk returns a tuple with the SnapKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapKbyte

`func (o *PoolStat) SetSnapKbyte(v int64)`

SetSnapKbyte sets SnapKbyte field to given value.

### HasSnapKbyte

`func (o *PoolStat) HasSnapKbyte() bool`

HasSnapKbyte returns a boolean if a field has been set.

### GetTotalBandwidthKbyte

`func (o *PoolStat) GetTotalBandwidthKbyte() float64`

GetTotalBandwidthKbyte returns the TotalBandwidthKbyte field if non-nil, zero value otherwise.

### GetTotalBandwidthKbyteOk

`func (o *PoolStat) GetTotalBandwidthKbyteOk() (*float64, bool)`

GetTotalBandwidthKbyteOk returns a tuple with the TotalBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBandwidthKbyte

`func (o *PoolStat) SetTotalBandwidthKbyte(v float64)`

SetTotalBandwidthKbyte sets TotalBandwidthKbyte field to given value.

### HasTotalBandwidthKbyte

`func (o *PoolStat) HasTotalBandwidthKbyte() bool`

HasTotalBandwidthKbyte returns a boolean if a field has been set.

### GetTotalCacheKbyte

`func (o *PoolStat) GetTotalCacheKbyte() float64`

GetTotalCacheKbyte returns the TotalCacheKbyte field if non-nil, zero value otherwise.

### GetTotalCacheKbyteOk

`func (o *PoolStat) GetTotalCacheKbyteOk() (*float64, bool)`

GetTotalCacheKbyteOk returns a tuple with the TotalCacheKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCacheKbyte

`func (o *PoolStat) SetTotalCacheKbyte(v float64)`

SetTotalCacheKbyte sets TotalCacheKbyte field to given value.

### HasTotalCacheKbyte

`func (o *PoolStat) HasTotalCacheKbyte() bool`

HasTotalCacheKbyte returns a boolean if a field has been set.

### GetTotalIops

`func (o *PoolStat) GetTotalIops() float64`

GetTotalIops returns the TotalIops field if non-nil, zero value otherwise.

### GetTotalIopsOk

`func (o *PoolStat) GetTotalIopsOk() (*float64, bool)`

GetTotalIopsOk returns a tuple with the TotalIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIops

`func (o *PoolStat) SetTotalIops(v float64)`

SetTotalIops sets TotalIops field to given value.

### HasTotalIops

`func (o *PoolStat) HasTotalIops() bool`

HasTotalIops returns a boolean if a field has been set.

### GetTotalKbyte

`func (o *PoolStat) GetTotalKbyte() int64`

GetTotalKbyte returns the TotalKbyte field if non-nil, zero value otherwise.

### GetTotalKbyteOk

`func (o *PoolStat) GetTotalKbyteOk() (*int64, bool)`

GetTotalKbyteOk returns a tuple with the TotalKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalKbyte

`func (o *PoolStat) SetTotalKbyte(v int64)`

SetTotalKbyte sets TotalKbyte field to given value.

### HasTotalKbyte

`func (o *PoolStat) HasTotalKbyte() bool`

HasTotalKbyte returns a boolean if a field has been set.

### GetUnavailableNum

`func (o *PoolStat) GetUnavailableNum() int64`

GetUnavailableNum returns the UnavailableNum field if non-nil, zero value otherwise.

### GetUnavailableNumOk

`func (o *PoolStat) GetUnavailableNumOk() (*int64, bool)`

GetUnavailableNumOk returns a tuple with the UnavailableNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableNum

`func (o *PoolStat) SetUnavailableNum(v int64)`

SetUnavailableNum sets UnavailableNum field to given value.

### HasUnavailableNum

`func (o *PoolStat) HasUnavailableNum() bool`

HasUnavailableNum returns a boolean if a field has been set.

### GetUnavailablePercent

`func (o *PoolStat) GetUnavailablePercent() float64`

GetUnavailablePercent returns the UnavailablePercent field if non-nil, zero value otherwise.

### GetUnavailablePercentOk

`func (o *PoolStat) GetUnavailablePercentOk() (*float64, bool)`

GetUnavailablePercentOk returns a tuple with the UnavailablePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailablePercent

`func (o *PoolStat) SetUnavailablePercent(v float64)`

SetUnavailablePercent sets UnavailablePercent field to given value.

### HasUnavailablePercent

`func (o *PoolStat) HasUnavailablePercent() bool`

HasUnavailablePercent returns a boolean if a field has been set.

### GetUnusedDataKbyte

`func (o *PoolStat) GetUnusedDataKbyte() float64`

GetUnusedDataKbyte returns the UnusedDataKbyte field if non-nil, zero value otherwise.

### GetUnusedDataKbyteOk

`func (o *PoolStat) GetUnusedDataKbyteOk() (*float64, bool)`

GetUnusedDataKbyteOk returns a tuple with the UnusedDataKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnusedDataKbyte

`func (o *PoolStat) SetUnusedDataKbyte(v float64)`

SetUnusedDataKbyte sets UnusedDataKbyte field to given value.

### HasUnusedDataKbyte

`func (o *PoolStat) HasUnusedDataKbyte() bool`

HasUnusedDataKbyte returns a boolean if a field has been set.

### GetUsedKbyte

`func (o *PoolStat) GetUsedKbyte() int64`

GetUsedKbyte returns the UsedKbyte field if non-nil, zero value otherwise.

### GetUsedKbyteOk

`func (o *PoolStat) GetUsedKbyteOk() (*int64, bool)`

GetUsedKbyteOk returns a tuple with the UsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte

`func (o *PoolStat) SetUsedKbyte(v int64)`

SetUsedKbyte sets UsedKbyte field to given value.

### HasUsedKbyte

`func (o *PoolStat) HasUsedKbyte() bool`

HasUsedKbyte returns a boolean if a field has been set.

### GetUsedPercent

`func (o *PoolStat) GetUsedPercent() float64`

GetUsedPercent returns the UsedPercent field if non-nil, zero value otherwise.

### GetUsedPercentOk

`func (o *PoolStat) GetUsedPercentOk() (*float64, bool)`

GetUsedPercentOk returns a tuple with the UsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPercent

`func (o *PoolStat) SetUsedPercent(v float64)`

SetUsedPercent sets UsedPercent field to given value.

### HasUsedPercent

`func (o *PoolStat) HasUsedPercent() bool`

HasUsedPercent returns a boolean if a field has been set.

### GetWaterLevel

`func (o *PoolStat) GetWaterLevel() float64`

GetWaterLevel returns the WaterLevel field if non-nil, zero value otherwise.

### GetWaterLevelOk

`func (o *PoolStat) GetWaterLevelOk() (*float64, bool)`

GetWaterLevelOk returns a tuple with the WaterLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaterLevel

`func (o *PoolStat) SetWaterLevel(v float64)`

SetWaterLevel sets WaterLevel field to given value.

### HasWaterLevel

`func (o *PoolStat) HasWaterLevel() bool`

HasWaterLevel returns a boolean if a field has been set.

### GetWriteBandwidthKbyte

`func (o *PoolStat) GetWriteBandwidthKbyte() float64`

GetWriteBandwidthKbyte returns the WriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetWriteBandwidthKbyteOk

`func (o *PoolStat) GetWriteBandwidthKbyteOk() (*float64, bool)`

GetWriteBandwidthKbyteOk returns a tuple with the WriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBandwidthKbyte

`func (o *PoolStat) SetWriteBandwidthKbyte(v float64)`

SetWriteBandwidthKbyte sets WriteBandwidthKbyte field to given value.

### HasWriteBandwidthKbyte

`func (o *PoolStat) HasWriteBandwidthKbyte() bool`

HasWriteBandwidthKbyte returns a boolean if a field has been set.

### GetWriteCacheKbyte

`func (o *PoolStat) GetWriteCacheKbyte() float64`

GetWriteCacheKbyte returns the WriteCacheKbyte field if non-nil, zero value otherwise.

### GetWriteCacheKbyteOk

`func (o *PoolStat) GetWriteCacheKbyteOk() (*float64, bool)`

GetWriteCacheKbyteOk returns a tuple with the WriteCacheKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCacheKbyte

`func (o *PoolStat) SetWriteCacheKbyte(v float64)`

SetWriteCacheKbyte sets WriteCacheKbyte field to given value.

### HasWriteCacheKbyte

`func (o *PoolStat) HasWriteCacheKbyte() bool`

HasWriteCacheKbyte returns a boolean if a field has been set.

### GetWriteCacheMergeRate

`func (o *PoolStat) GetWriteCacheMergeRate() float64`

GetWriteCacheMergeRate returns the WriteCacheMergeRate field if non-nil, zero value otherwise.

### GetWriteCacheMergeRateOk

`func (o *PoolStat) GetWriteCacheMergeRateOk() (*float64, bool)`

GetWriteCacheMergeRateOk returns a tuple with the WriteCacheMergeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCacheMergeRate

`func (o *PoolStat) SetWriteCacheMergeRate(v float64)`

SetWriteCacheMergeRate sets WriteCacheMergeRate field to given value.

### HasWriteCacheMergeRate

`func (o *PoolStat) HasWriteCacheMergeRate() bool`

HasWriteCacheMergeRate returns a boolean if a field has been set.

### GetWriteCachePercent

`func (o *PoolStat) GetWriteCachePercent() float64`

GetWriteCachePercent returns the WriteCachePercent field if non-nil, zero value otherwise.

### GetWriteCachePercentOk

`func (o *PoolStat) GetWriteCachePercentOk() (*float64, bool)`

GetWriteCachePercentOk returns a tuple with the WriteCachePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCachePercent

`func (o *PoolStat) SetWriteCachePercent(v float64)`

SetWriteCachePercent sets WriteCachePercent field to given value.

### HasWriteCachePercent

`func (o *PoolStat) HasWriteCachePercent() bool`

HasWriteCachePercent returns a boolean if a field has been set.

### GetWriteIoSizeKbyte

`func (o *PoolStat) GetWriteIoSizeKbyte() float64`

GetWriteIoSizeKbyte returns the WriteIoSizeKbyte field if non-nil, zero value otherwise.

### GetWriteIoSizeKbyteOk

`func (o *PoolStat) GetWriteIoSizeKbyteOk() (*float64, bool)`

GetWriteIoSizeKbyteOk returns a tuple with the WriteIoSizeKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIoSizeKbyte

`func (o *PoolStat) SetWriteIoSizeKbyte(v float64)`

SetWriteIoSizeKbyte sets WriteIoSizeKbyte field to given value.

### HasWriteIoSizeKbyte

`func (o *PoolStat) HasWriteIoSizeKbyte() bool`

HasWriteIoSizeKbyte returns a boolean if a field has been set.

### GetWriteIops

`func (o *PoolStat) GetWriteIops() float64`

GetWriteIops returns the WriteIops field if non-nil, zero value otherwise.

### GetWriteIopsOk

`func (o *PoolStat) GetWriteIopsOk() (*float64, bool)`

GetWriteIopsOk returns a tuple with the WriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIops

`func (o *PoolStat) SetWriteIops(v float64)`

SetWriteIops sets WriteIops field to given value.

### HasWriteIops

`func (o *PoolStat) HasWriteIops() bool`

HasWriteIops returns a boolean if a field has been set.

### GetWriteLatencyUs

`func (o *PoolStat) GetWriteLatencyUs() float64`

GetWriteLatencyUs returns the WriteLatencyUs field if non-nil, zero value otherwise.

### GetWriteLatencyUsOk

`func (o *PoolStat) GetWriteLatencyUsOk() (*float64, bool)`

GetWriteLatencyUsOk returns a tuple with the WriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLatencyUs

`func (o *PoolStat) SetWriteLatencyUs(v float64)`

SetWriteLatencyUs sets WriteLatencyUs field to given value.

### HasWriteLatencyUs

`func (o *PoolStat) HasWriteLatencyUs() bool`

HasWriteLatencyUs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


