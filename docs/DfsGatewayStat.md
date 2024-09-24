# DfsGatewayStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketConnNum** | Pointer to **int64** |  | [optional] 
**ConnNum** | Pointer to **int64** |  | [optional] 
**CpuUtil** | Pointer to **float64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DataCacheHitRate** | Pointer to **float64** |  | [optional] 
**FtpConnNum** | Pointer to **int64** |  | [optional] 
**HdfsConnNum** | Pointer to **int64** |  | [optional] 
**MemUsagePercent** | Pointer to **float64** |  | [optional] 
**MetaCacheHitRate** | Pointer to **float64** |  | [optional] 
**MetaDeleteLatencyUs** | Pointer to **float64** |  | [optional] 
**MetaDeleteOps** | Pointer to **float64** |  | [optional] 
**MetaListLatencyUs** | Pointer to **float64** |  | [optional] 
**MetaListOps** | Pointer to **float64** |  | [optional] 
**MetaReadLatencyUs** | Pointer to **float64** |  | [optional] 
**MetaReadOps** | Pointer to **float64** |  | [optional] 
**MetaTruncLatencyUs** | Pointer to **float64** |  | [optional] 
**MetaTruncOps** | Pointer to **float64** |  | [optional] 
**MetaWriteLatencyUs** | Pointer to **float64** |  | [optional] 
**MetaWriteOps** | Pointer to **float64** |  | [optional] 
**NfsConnNum** | Pointer to **int64** |  | [optional] 
**ReadAheadBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**ReadAheadIops** | Pointer to **float64** |  | [optional] 
**ReadBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**ReadIops** | Pointer to **float64** |  | [optional] 
**ReadLatencyUs** | Pointer to **float64** |  | [optional] 
**SmbSessionNum** | Pointer to **int64** |  | [optional] 
**WriteBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**WriteIops** | Pointer to **float64** |  | [optional] 
**WriteLatencyUs** | Pointer to **float64** |  | [optional] 

## Methods

### NewDfsGatewayStat

`func NewDfsGatewayStat() *DfsGatewayStat`

NewDfsGatewayStat instantiates a new DfsGatewayStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayStatWithDefaults

`func NewDfsGatewayStatWithDefaults() *DfsGatewayStat`

NewDfsGatewayStatWithDefaults instantiates a new DfsGatewayStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketConnNum

`func (o *DfsGatewayStat) GetBucketConnNum() int64`

GetBucketConnNum returns the BucketConnNum field if non-nil, zero value otherwise.

### GetBucketConnNumOk

`func (o *DfsGatewayStat) GetBucketConnNumOk() (*int64, bool)`

GetBucketConnNumOk returns a tuple with the BucketConnNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketConnNum

`func (o *DfsGatewayStat) SetBucketConnNum(v int64)`

SetBucketConnNum sets BucketConnNum field to given value.

### HasBucketConnNum

`func (o *DfsGatewayStat) HasBucketConnNum() bool`

HasBucketConnNum returns a boolean if a field has been set.

### GetConnNum

`func (o *DfsGatewayStat) GetConnNum() int64`

GetConnNum returns the ConnNum field if non-nil, zero value otherwise.

### GetConnNumOk

`func (o *DfsGatewayStat) GetConnNumOk() (*int64, bool)`

GetConnNumOk returns a tuple with the ConnNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnNum

`func (o *DfsGatewayStat) SetConnNum(v int64)`

SetConnNum sets ConnNum field to given value.

### HasConnNum

`func (o *DfsGatewayStat) HasConnNum() bool`

HasConnNum returns a boolean if a field has been set.

### GetCpuUtil

`func (o *DfsGatewayStat) GetCpuUtil() float64`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *DfsGatewayStat) GetCpuUtilOk() (*float64, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *DfsGatewayStat) SetCpuUtil(v float64)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *DfsGatewayStat) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCreate

`func (o *DfsGatewayStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsGatewayStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsGatewayStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsGatewayStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataCacheHitRate

`func (o *DfsGatewayStat) GetDataCacheHitRate() float64`

GetDataCacheHitRate returns the DataCacheHitRate field if non-nil, zero value otherwise.

### GetDataCacheHitRateOk

`func (o *DfsGatewayStat) GetDataCacheHitRateOk() (*float64, bool)`

GetDataCacheHitRateOk returns a tuple with the DataCacheHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCacheHitRate

`func (o *DfsGatewayStat) SetDataCacheHitRate(v float64)`

SetDataCacheHitRate sets DataCacheHitRate field to given value.

### HasDataCacheHitRate

`func (o *DfsGatewayStat) HasDataCacheHitRate() bool`

HasDataCacheHitRate returns a boolean if a field has been set.

### GetFtpConnNum

`func (o *DfsGatewayStat) GetFtpConnNum() int64`

GetFtpConnNum returns the FtpConnNum field if non-nil, zero value otherwise.

### GetFtpConnNumOk

`func (o *DfsGatewayStat) GetFtpConnNumOk() (*int64, bool)`

GetFtpConnNumOk returns a tuple with the FtpConnNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpConnNum

`func (o *DfsGatewayStat) SetFtpConnNum(v int64)`

SetFtpConnNum sets FtpConnNum field to given value.

### HasFtpConnNum

`func (o *DfsGatewayStat) HasFtpConnNum() bool`

HasFtpConnNum returns a boolean if a field has been set.

### GetHdfsConnNum

`func (o *DfsGatewayStat) GetHdfsConnNum() int64`

GetHdfsConnNum returns the HdfsConnNum field if non-nil, zero value otherwise.

### GetHdfsConnNumOk

`func (o *DfsGatewayStat) GetHdfsConnNumOk() (*int64, bool)`

GetHdfsConnNumOk returns a tuple with the HdfsConnNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsConnNum

`func (o *DfsGatewayStat) SetHdfsConnNum(v int64)`

SetHdfsConnNum sets HdfsConnNum field to given value.

### HasHdfsConnNum

`func (o *DfsGatewayStat) HasHdfsConnNum() bool`

HasHdfsConnNum returns a boolean if a field has been set.

### GetMemUsagePercent

`func (o *DfsGatewayStat) GetMemUsagePercent() float64`

GetMemUsagePercent returns the MemUsagePercent field if non-nil, zero value otherwise.

### GetMemUsagePercentOk

`func (o *DfsGatewayStat) GetMemUsagePercentOk() (*float64, bool)`

GetMemUsagePercentOk returns a tuple with the MemUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsagePercent

`func (o *DfsGatewayStat) SetMemUsagePercent(v float64)`

SetMemUsagePercent sets MemUsagePercent field to given value.

### HasMemUsagePercent

`func (o *DfsGatewayStat) HasMemUsagePercent() bool`

HasMemUsagePercent returns a boolean if a field has been set.

### GetMetaCacheHitRate

`func (o *DfsGatewayStat) GetMetaCacheHitRate() float64`

GetMetaCacheHitRate returns the MetaCacheHitRate field if non-nil, zero value otherwise.

### GetMetaCacheHitRateOk

`func (o *DfsGatewayStat) GetMetaCacheHitRateOk() (*float64, bool)`

GetMetaCacheHitRateOk returns a tuple with the MetaCacheHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaCacheHitRate

`func (o *DfsGatewayStat) SetMetaCacheHitRate(v float64)`

SetMetaCacheHitRate sets MetaCacheHitRate field to given value.

### HasMetaCacheHitRate

`func (o *DfsGatewayStat) HasMetaCacheHitRate() bool`

HasMetaCacheHitRate returns a boolean if a field has been set.

### GetMetaDeleteLatencyUs

`func (o *DfsGatewayStat) GetMetaDeleteLatencyUs() float64`

GetMetaDeleteLatencyUs returns the MetaDeleteLatencyUs field if non-nil, zero value otherwise.

### GetMetaDeleteLatencyUsOk

`func (o *DfsGatewayStat) GetMetaDeleteLatencyUsOk() (*float64, bool)`

GetMetaDeleteLatencyUsOk returns a tuple with the MetaDeleteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDeleteLatencyUs

`func (o *DfsGatewayStat) SetMetaDeleteLatencyUs(v float64)`

SetMetaDeleteLatencyUs sets MetaDeleteLatencyUs field to given value.

### HasMetaDeleteLatencyUs

`func (o *DfsGatewayStat) HasMetaDeleteLatencyUs() bool`

HasMetaDeleteLatencyUs returns a boolean if a field has been set.

### GetMetaDeleteOps

`func (o *DfsGatewayStat) GetMetaDeleteOps() float64`

GetMetaDeleteOps returns the MetaDeleteOps field if non-nil, zero value otherwise.

### GetMetaDeleteOpsOk

`func (o *DfsGatewayStat) GetMetaDeleteOpsOk() (*float64, bool)`

GetMetaDeleteOpsOk returns a tuple with the MetaDeleteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDeleteOps

`func (o *DfsGatewayStat) SetMetaDeleteOps(v float64)`

SetMetaDeleteOps sets MetaDeleteOps field to given value.

### HasMetaDeleteOps

`func (o *DfsGatewayStat) HasMetaDeleteOps() bool`

HasMetaDeleteOps returns a boolean if a field has been set.

### GetMetaListLatencyUs

`func (o *DfsGatewayStat) GetMetaListLatencyUs() float64`

GetMetaListLatencyUs returns the MetaListLatencyUs field if non-nil, zero value otherwise.

### GetMetaListLatencyUsOk

`func (o *DfsGatewayStat) GetMetaListLatencyUsOk() (*float64, bool)`

GetMetaListLatencyUsOk returns a tuple with the MetaListLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaListLatencyUs

`func (o *DfsGatewayStat) SetMetaListLatencyUs(v float64)`

SetMetaListLatencyUs sets MetaListLatencyUs field to given value.

### HasMetaListLatencyUs

`func (o *DfsGatewayStat) HasMetaListLatencyUs() bool`

HasMetaListLatencyUs returns a boolean if a field has been set.

### GetMetaListOps

`func (o *DfsGatewayStat) GetMetaListOps() float64`

GetMetaListOps returns the MetaListOps field if non-nil, zero value otherwise.

### GetMetaListOpsOk

`func (o *DfsGatewayStat) GetMetaListOpsOk() (*float64, bool)`

GetMetaListOpsOk returns a tuple with the MetaListOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaListOps

`func (o *DfsGatewayStat) SetMetaListOps(v float64)`

SetMetaListOps sets MetaListOps field to given value.

### HasMetaListOps

`func (o *DfsGatewayStat) HasMetaListOps() bool`

HasMetaListOps returns a boolean if a field has been set.

### GetMetaReadLatencyUs

`func (o *DfsGatewayStat) GetMetaReadLatencyUs() float64`

GetMetaReadLatencyUs returns the MetaReadLatencyUs field if non-nil, zero value otherwise.

### GetMetaReadLatencyUsOk

`func (o *DfsGatewayStat) GetMetaReadLatencyUsOk() (*float64, bool)`

GetMetaReadLatencyUsOk returns a tuple with the MetaReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaReadLatencyUs

`func (o *DfsGatewayStat) SetMetaReadLatencyUs(v float64)`

SetMetaReadLatencyUs sets MetaReadLatencyUs field to given value.

### HasMetaReadLatencyUs

`func (o *DfsGatewayStat) HasMetaReadLatencyUs() bool`

HasMetaReadLatencyUs returns a boolean if a field has been set.

### GetMetaReadOps

`func (o *DfsGatewayStat) GetMetaReadOps() float64`

GetMetaReadOps returns the MetaReadOps field if non-nil, zero value otherwise.

### GetMetaReadOpsOk

`func (o *DfsGatewayStat) GetMetaReadOpsOk() (*float64, bool)`

GetMetaReadOpsOk returns a tuple with the MetaReadOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaReadOps

`func (o *DfsGatewayStat) SetMetaReadOps(v float64)`

SetMetaReadOps sets MetaReadOps field to given value.

### HasMetaReadOps

`func (o *DfsGatewayStat) HasMetaReadOps() bool`

HasMetaReadOps returns a boolean if a field has been set.

### GetMetaTruncLatencyUs

`func (o *DfsGatewayStat) GetMetaTruncLatencyUs() float64`

GetMetaTruncLatencyUs returns the MetaTruncLatencyUs field if non-nil, zero value otherwise.

### GetMetaTruncLatencyUsOk

`func (o *DfsGatewayStat) GetMetaTruncLatencyUsOk() (*float64, bool)`

GetMetaTruncLatencyUsOk returns a tuple with the MetaTruncLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTruncLatencyUs

`func (o *DfsGatewayStat) SetMetaTruncLatencyUs(v float64)`

SetMetaTruncLatencyUs sets MetaTruncLatencyUs field to given value.

### HasMetaTruncLatencyUs

`func (o *DfsGatewayStat) HasMetaTruncLatencyUs() bool`

HasMetaTruncLatencyUs returns a boolean if a field has been set.

### GetMetaTruncOps

`func (o *DfsGatewayStat) GetMetaTruncOps() float64`

GetMetaTruncOps returns the MetaTruncOps field if non-nil, zero value otherwise.

### GetMetaTruncOpsOk

`func (o *DfsGatewayStat) GetMetaTruncOpsOk() (*float64, bool)`

GetMetaTruncOpsOk returns a tuple with the MetaTruncOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTruncOps

`func (o *DfsGatewayStat) SetMetaTruncOps(v float64)`

SetMetaTruncOps sets MetaTruncOps field to given value.

### HasMetaTruncOps

`func (o *DfsGatewayStat) HasMetaTruncOps() bool`

HasMetaTruncOps returns a boolean if a field has been set.

### GetMetaWriteLatencyUs

`func (o *DfsGatewayStat) GetMetaWriteLatencyUs() float64`

GetMetaWriteLatencyUs returns the MetaWriteLatencyUs field if non-nil, zero value otherwise.

### GetMetaWriteLatencyUsOk

`func (o *DfsGatewayStat) GetMetaWriteLatencyUsOk() (*float64, bool)`

GetMetaWriteLatencyUsOk returns a tuple with the MetaWriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaWriteLatencyUs

`func (o *DfsGatewayStat) SetMetaWriteLatencyUs(v float64)`

SetMetaWriteLatencyUs sets MetaWriteLatencyUs field to given value.

### HasMetaWriteLatencyUs

`func (o *DfsGatewayStat) HasMetaWriteLatencyUs() bool`

HasMetaWriteLatencyUs returns a boolean if a field has been set.

### GetMetaWriteOps

`func (o *DfsGatewayStat) GetMetaWriteOps() float64`

GetMetaWriteOps returns the MetaWriteOps field if non-nil, zero value otherwise.

### GetMetaWriteOpsOk

`func (o *DfsGatewayStat) GetMetaWriteOpsOk() (*float64, bool)`

GetMetaWriteOpsOk returns a tuple with the MetaWriteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaWriteOps

`func (o *DfsGatewayStat) SetMetaWriteOps(v float64)`

SetMetaWriteOps sets MetaWriteOps field to given value.

### HasMetaWriteOps

`func (o *DfsGatewayStat) HasMetaWriteOps() bool`

HasMetaWriteOps returns a boolean if a field has been set.

### GetNfsConnNum

`func (o *DfsGatewayStat) GetNfsConnNum() int64`

GetNfsConnNum returns the NfsConnNum field if non-nil, zero value otherwise.

### GetNfsConnNumOk

`func (o *DfsGatewayStat) GetNfsConnNumOk() (*int64, bool)`

GetNfsConnNumOk returns a tuple with the NfsConnNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsConnNum

`func (o *DfsGatewayStat) SetNfsConnNum(v int64)`

SetNfsConnNum sets NfsConnNum field to given value.

### HasNfsConnNum

`func (o *DfsGatewayStat) HasNfsConnNum() bool`

HasNfsConnNum returns a boolean if a field has been set.

### GetReadAheadBandwidthKbyte

`func (o *DfsGatewayStat) GetReadAheadBandwidthKbyte() float64`

GetReadAheadBandwidthKbyte returns the ReadAheadBandwidthKbyte field if non-nil, zero value otherwise.

### GetReadAheadBandwidthKbyteOk

`func (o *DfsGatewayStat) GetReadAheadBandwidthKbyteOk() (*float64, bool)`

GetReadAheadBandwidthKbyteOk returns a tuple with the ReadAheadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadBandwidthKbyte

`func (o *DfsGatewayStat) SetReadAheadBandwidthKbyte(v float64)`

SetReadAheadBandwidthKbyte sets ReadAheadBandwidthKbyte field to given value.

### HasReadAheadBandwidthKbyte

`func (o *DfsGatewayStat) HasReadAheadBandwidthKbyte() bool`

HasReadAheadBandwidthKbyte returns a boolean if a field has been set.

### GetReadAheadIops

`func (o *DfsGatewayStat) GetReadAheadIops() float64`

GetReadAheadIops returns the ReadAheadIops field if non-nil, zero value otherwise.

### GetReadAheadIopsOk

`func (o *DfsGatewayStat) GetReadAheadIopsOk() (*float64, bool)`

GetReadAheadIopsOk returns a tuple with the ReadAheadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadIops

`func (o *DfsGatewayStat) SetReadAheadIops(v float64)`

SetReadAheadIops sets ReadAheadIops field to given value.

### HasReadAheadIops

`func (o *DfsGatewayStat) HasReadAheadIops() bool`

HasReadAheadIops returns a boolean if a field has been set.

### GetReadBandwidthKbyte

`func (o *DfsGatewayStat) GetReadBandwidthKbyte() float64`

GetReadBandwidthKbyte returns the ReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetReadBandwidthKbyteOk

`func (o *DfsGatewayStat) GetReadBandwidthKbyteOk() (*float64, bool)`

GetReadBandwidthKbyteOk returns a tuple with the ReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadBandwidthKbyte

`func (o *DfsGatewayStat) SetReadBandwidthKbyte(v float64)`

SetReadBandwidthKbyte sets ReadBandwidthKbyte field to given value.

### HasReadBandwidthKbyte

`func (o *DfsGatewayStat) HasReadBandwidthKbyte() bool`

HasReadBandwidthKbyte returns a boolean if a field has been set.

### GetReadIops

`func (o *DfsGatewayStat) GetReadIops() float64`

GetReadIops returns the ReadIops field if non-nil, zero value otherwise.

### GetReadIopsOk

`func (o *DfsGatewayStat) GetReadIopsOk() (*float64, bool)`

GetReadIopsOk returns a tuple with the ReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIops

`func (o *DfsGatewayStat) SetReadIops(v float64)`

SetReadIops sets ReadIops field to given value.

### HasReadIops

`func (o *DfsGatewayStat) HasReadIops() bool`

HasReadIops returns a boolean if a field has been set.

### GetReadLatencyUs

`func (o *DfsGatewayStat) GetReadLatencyUs() float64`

GetReadLatencyUs returns the ReadLatencyUs field if non-nil, zero value otherwise.

### GetReadLatencyUsOk

`func (o *DfsGatewayStat) GetReadLatencyUsOk() (*float64, bool)`

GetReadLatencyUsOk returns a tuple with the ReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadLatencyUs

`func (o *DfsGatewayStat) SetReadLatencyUs(v float64)`

SetReadLatencyUs sets ReadLatencyUs field to given value.

### HasReadLatencyUs

`func (o *DfsGatewayStat) HasReadLatencyUs() bool`

HasReadLatencyUs returns a boolean if a field has been set.

### GetSmbSessionNum

`func (o *DfsGatewayStat) GetSmbSessionNum() int64`

GetSmbSessionNum returns the SmbSessionNum field if non-nil, zero value otherwise.

### GetSmbSessionNumOk

`func (o *DfsGatewayStat) GetSmbSessionNumOk() (*int64, bool)`

GetSmbSessionNumOk returns a tuple with the SmbSessionNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbSessionNum

`func (o *DfsGatewayStat) SetSmbSessionNum(v int64)`

SetSmbSessionNum sets SmbSessionNum field to given value.

### HasSmbSessionNum

`func (o *DfsGatewayStat) HasSmbSessionNum() bool`

HasSmbSessionNum returns a boolean if a field has been set.

### GetWriteBandwidthKbyte

`func (o *DfsGatewayStat) GetWriteBandwidthKbyte() float64`

GetWriteBandwidthKbyte returns the WriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetWriteBandwidthKbyteOk

`func (o *DfsGatewayStat) GetWriteBandwidthKbyteOk() (*float64, bool)`

GetWriteBandwidthKbyteOk returns a tuple with the WriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBandwidthKbyte

`func (o *DfsGatewayStat) SetWriteBandwidthKbyte(v float64)`

SetWriteBandwidthKbyte sets WriteBandwidthKbyte field to given value.

### HasWriteBandwidthKbyte

`func (o *DfsGatewayStat) HasWriteBandwidthKbyte() bool`

HasWriteBandwidthKbyte returns a boolean if a field has been set.

### GetWriteIops

`func (o *DfsGatewayStat) GetWriteIops() float64`

GetWriteIops returns the WriteIops field if non-nil, zero value otherwise.

### GetWriteIopsOk

`func (o *DfsGatewayStat) GetWriteIopsOk() (*float64, bool)`

GetWriteIopsOk returns a tuple with the WriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIops

`func (o *DfsGatewayStat) SetWriteIops(v float64)`

SetWriteIops sets WriteIops field to given value.

### HasWriteIops

`func (o *DfsGatewayStat) HasWriteIops() bool`

HasWriteIops returns a boolean if a field has been set.

### GetWriteLatencyUs

`func (o *DfsGatewayStat) GetWriteLatencyUs() float64`

GetWriteLatencyUs returns the WriteLatencyUs field if non-nil, zero value otherwise.

### GetWriteLatencyUsOk

`func (o *DfsGatewayStat) GetWriteLatencyUsOk() (*float64, bool)`

GetWriteLatencyUsOk returns a tuple with the WriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLatencyUs

`func (o *DfsGatewayStat) SetWriteLatencyUs(v float64)`

SetWriteLatencyUs sets WriteLatencyUs field to given value.

### HasWriteLatencyUs

`func (o *DfsGatewayStat) HasWriteLatencyUs() bool`

HasWriteLatencyUs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


