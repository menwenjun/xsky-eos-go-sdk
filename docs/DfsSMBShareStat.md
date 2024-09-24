# DfsSMBShareStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**DataCacheHitRate** | Pointer to **float64** |  | [optional] 
**DataDeleteBandwidthKbyte** | Pointer to **int64** |  | [optional] 
**DataDeleteIops** | Pointer to **int64** |  | [optional] 
**DataDeleteLatencyUs** | Pointer to **int64** |  | [optional] 
**DataReadBandwidthKbyte** | Pointer to **int64** |  | [optional] 
**DataReadIops** | Pointer to **int64** |  | [optional] 
**DataReadLatencyUs** | Pointer to **int64** |  | [optional] 
**DataWriteBandwidthKbyte** | Pointer to **int64** |  | [optional] 
**DataWriteIops** | Pointer to **int64** |  | [optional] 
**DataWriteLatencyUs** | Pointer to **int64** |  | [optional] 
**MetaCacheHitRate** | Pointer to **float64** |  | [optional] 
**MetaDeleteLatencyUs** | Pointer to **int64** |  | [optional] 
**MetaDeleteOps** | Pointer to **int64** |  | [optional] 
**MetaListLatencyUs** | Pointer to **int64** |  | [optional] 
**MetaListOps** | Pointer to **int64** |  | [optional] 
**MetaReadLatencyUs** | Pointer to **int64** |  | [optional] 
**MetaReadOps** | Pointer to **int64** |  | [optional] 
**MetaWriteLatencyUs** | Pointer to **int64** |  | [optional] 
**MetaWriteOps** | Pointer to **int64** |  | [optional] 
**ReadAheadBandwidthKbyte** | Pointer to **int64** |  | [optional] 
**ReadAheadIops** | Pointer to **int64** |  | [optional] 

## Methods

### NewDfsSMBShareStat

`func NewDfsSMBShareStat() *DfsSMBShareStat`

NewDfsSMBShareStat instantiates a new DfsSMBShareStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBShareStatWithDefaults

`func NewDfsSMBShareStatWithDefaults() *DfsSMBShareStat`

NewDfsSMBShareStatWithDefaults instantiates a new DfsSMBShareStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *DfsSMBShareStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsSMBShareStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsSMBShareStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsSMBShareStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataCacheHitRate

`func (o *DfsSMBShareStat) GetDataCacheHitRate() float64`

GetDataCacheHitRate returns the DataCacheHitRate field if non-nil, zero value otherwise.

### GetDataCacheHitRateOk

`func (o *DfsSMBShareStat) GetDataCacheHitRateOk() (*float64, bool)`

GetDataCacheHitRateOk returns a tuple with the DataCacheHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCacheHitRate

`func (o *DfsSMBShareStat) SetDataCacheHitRate(v float64)`

SetDataCacheHitRate sets DataCacheHitRate field to given value.

### HasDataCacheHitRate

`func (o *DfsSMBShareStat) HasDataCacheHitRate() bool`

HasDataCacheHitRate returns a boolean if a field has been set.

### GetDataDeleteBandwidthKbyte

`func (o *DfsSMBShareStat) GetDataDeleteBandwidthKbyte() int64`

GetDataDeleteBandwidthKbyte returns the DataDeleteBandwidthKbyte field if non-nil, zero value otherwise.

### GetDataDeleteBandwidthKbyteOk

`func (o *DfsSMBShareStat) GetDataDeleteBandwidthKbyteOk() (*int64, bool)`

GetDataDeleteBandwidthKbyteOk returns a tuple with the DataDeleteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeleteBandwidthKbyte

`func (o *DfsSMBShareStat) SetDataDeleteBandwidthKbyte(v int64)`

SetDataDeleteBandwidthKbyte sets DataDeleteBandwidthKbyte field to given value.

### HasDataDeleteBandwidthKbyte

`func (o *DfsSMBShareStat) HasDataDeleteBandwidthKbyte() bool`

HasDataDeleteBandwidthKbyte returns a boolean if a field has been set.

### GetDataDeleteIops

`func (o *DfsSMBShareStat) GetDataDeleteIops() int64`

GetDataDeleteIops returns the DataDeleteIops field if non-nil, zero value otherwise.

### GetDataDeleteIopsOk

`func (o *DfsSMBShareStat) GetDataDeleteIopsOk() (*int64, bool)`

GetDataDeleteIopsOk returns a tuple with the DataDeleteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeleteIops

`func (o *DfsSMBShareStat) SetDataDeleteIops(v int64)`

SetDataDeleteIops sets DataDeleteIops field to given value.

### HasDataDeleteIops

`func (o *DfsSMBShareStat) HasDataDeleteIops() bool`

HasDataDeleteIops returns a boolean if a field has been set.

### GetDataDeleteLatencyUs

`func (o *DfsSMBShareStat) GetDataDeleteLatencyUs() int64`

GetDataDeleteLatencyUs returns the DataDeleteLatencyUs field if non-nil, zero value otherwise.

### GetDataDeleteLatencyUsOk

`func (o *DfsSMBShareStat) GetDataDeleteLatencyUsOk() (*int64, bool)`

GetDataDeleteLatencyUsOk returns a tuple with the DataDeleteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeleteLatencyUs

`func (o *DfsSMBShareStat) SetDataDeleteLatencyUs(v int64)`

SetDataDeleteLatencyUs sets DataDeleteLatencyUs field to given value.

### HasDataDeleteLatencyUs

`func (o *DfsSMBShareStat) HasDataDeleteLatencyUs() bool`

HasDataDeleteLatencyUs returns a boolean if a field has been set.

### GetDataReadBandwidthKbyte

`func (o *DfsSMBShareStat) GetDataReadBandwidthKbyte() int64`

GetDataReadBandwidthKbyte returns the DataReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetDataReadBandwidthKbyteOk

`func (o *DfsSMBShareStat) GetDataReadBandwidthKbyteOk() (*int64, bool)`

GetDataReadBandwidthKbyteOk returns a tuple with the DataReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadBandwidthKbyte

`func (o *DfsSMBShareStat) SetDataReadBandwidthKbyte(v int64)`

SetDataReadBandwidthKbyte sets DataReadBandwidthKbyte field to given value.

### HasDataReadBandwidthKbyte

`func (o *DfsSMBShareStat) HasDataReadBandwidthKbyte() bool`

HasDataReadBandwidthKbyte returns a boolean if a field has been set.

### GetDataReadIops

`func (o *DfsSMBShareStat) GetDataReadIops() int64`

GetDataReadIops returns the DataReadIops field if non-nil, zero value otherwise.

### GetDataReadIopsOk

`func (o *DfsSMBShareStat) GetDataReadIopsOk() (*int64, bool)`

GetDataReadIopsOk returns a tuple with the DataReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadIops

`func (o *DfsSMBShareStat) SetDataReadIops(v int64)`

SetDataReadIops sets DataReadIops field to given value.

### HasDataReadIops

`func (o *DfsSMBShareStat) HasDataReadIops() bool`

HasDataReadIops returns a boolean if a field has been set.

### GetDataReadLatencyUs

`func (o *DfsSMBShareStat) GetDataReadLatencyUs() int64`

GetDataReadLatencyUs returns the DataReadLatencyUs field if non-nil, zero value otherwise.

### GetDataReadLatencyUsOk

`func (o *DfsSMBShareStat) GetDataReadLatencyUsOk() (*int64, bool)`

GetDataReadLatencyUsOk returns a tuple with the DataReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadLatencyUs

`func (o *DfsSMBShareStat) SetDataReadLatencyUs(v int64)`

SetDataReadLatencyUs sets DataReadLatencyUs field to given value.

### HasDataReadLatencyUs

`func (o *DfsSMBShareStat) HasDataReadLatencyUs() bool`

HasDataReadLatencyUs returns a boolean if a field has been set.

### GetDataWriteBandwidthKbyte

`func (o *DfsSMBShareStat) GetDataWriteBandwidthKbyte() int64`

GetDataWriteBandwidthKbyte returns the DataWriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetDataWriteBandwidthKbyteOk

`func (o *DfsSMBShareStat) GetDataWriteBandwidthKbyteOk() (*int64, bool)`

GetDataWriteBandwidthKbyteOk returns a tuple with the DataWriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWriteBandwidthKbyte

`func (o *DfsSMBShareStat) SetDataWriteBandwidthKbyte(v int64)`

SetDataWriteBandwidthKbyte sets DataWriteBandwidthKbyte field to given value.

### HasDataWriteBandwidthKbyte

`func (o *DfsSMBShareStat) HasDataWriteBandwidthKbyte() bool`

HasDataWriteBandwidthKbyte returns a boolean if a field has been set.

### GetDataWriteIops

`func (o *DfsSMBShareStat) GetDataWriteIops() int64`

GetDataWriteIops returns the DataWriteIops field if non-nil, zero value otherwise.

### GetDataWriteIopsOk

`func (o *DfsSMBShareStat) GetDataWriteIopsOk() (*int64, bool)`

GetDataWriteIopsOk returns a tuple with the DataWriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWriteIops

`func (o *DfsSMBShareStat) SetDataWriteIops(v int64)`

SetDataWriteIops sets DataWriteIops field to given value.

### HasDataWriteIops

`func (o *DfsSMBShareStat) HasDataWriteIops() bool`

HasDataWriteIops returns a boolean if a field has been set.

### GetDataWriteLatencyUs

`func (o *DfsSMBShareStat) GetDataWriteLatencyUs() int64`

GetDataWriteLatencyUs returns the DataWriteLatencyUs field if non-nil, zero value otherwise.

### GetDataWriteLatencyUsOk

`func (o *DfsSMBShareStat) GetDataWriteLatencyUsOk() (*int64, bool)`

GetDataWriteLatencyUsOk returns a tuple with the DataWriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWriteLatencyUs

`func (o *DfsSMBShareStat) SetDataWriteLatencyUs(v int64)`

SetDataWriteLatencyUs sets DataWriteLatencyUs field to given value.

### HasDataWriteLatencyUs

`func (o *DfsSMBShareStat) HasDataWriteLatencyUs() bool`

HasDataWriteLatencyUs returns a boolean if a field has been set.

### GetMetaCacheHitRate

`func (o *DfsSMBShareStat) GetMetaCacheHitRate() float64`

GetMetaCacheHitRate returns the MetaCacheHitRate field if non-nil, zero value otherwise.

### GetMetaCacheHitRateOk

`func (o *DfsSMBShareStat) GetMetaCacheHitRateOk() (*float64, bool)`

GetMetaCacheHitRateOk returns a tuple with the MetaCacheHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaCacheHitRate

`func (o *DfsSMBShareStat) SetMetaCacheHitRate(v float64)`

SetMetaCacheHitRate sets MetaCacheHitRate field to given value.

### HasMetaCacheHitRate

`func (o *DfsSMBShareStat) HasMetaCacheHitRate() bool`

HasMetaCacheHitRate returns a boolean if a field has been set.

### GetMetaDeleteLatencyUs

`func (o *DfsSMBShareStat) GetMetaDeleteLatencyUs() int64`

GetMetaDeleteLatencyUs returns the MetaDeleteLatencyUs field if non-nil, zero value otherwise.

### GetMetaDeleteLatencyUsOk

`func (o *DfsSMBShareStat) GetMetaDeleteLatencyUsOk() (*int64, bool)`

GetMetaDeleteLatencyUsOk returns a tuple with the MetaDeleteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDeleteLatencyUs

`func (o *DfsSMBShareStat) SetMetaDeleteLatencyUs(v int64)`

SetMetaDeleteLatencyUs sets MetaDeleteLatencyUs field to given value.

### HasMetaDeleteLatencyUs

`func (o *DfsSMBShareStat) HasMetaDeleteLatencyUs() bool`

HasMetaDeleteLatencyUs returns a boolean if a field has been set.

### GetMetaDeleteOps

`func (o *DfsSMBShareStat) GetMetaDeleteOps() int64`

GetMetaDeleteOps returns the MetaDeleteOps field if non-nil, zero value otherwise.

### GetMetaDeleteOpsOk

`func (o *DfsSMBShareStat) GetMetaDeleteOpsOk() (*int64, bool)`

GetMetaDeleteOpsOk returns a tuple with the MetaDeleteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDeleteOps

`func (o *DfsSMBShareStat) SetMetaDeleteOps(v int64)`

SetMetaDeleteOps sets MetaDeleteOps field to given value.

### HasMetaDeleteOps

`func (o *DfsSMBShareStat) HasMetaDeleteOps() bool`

HasMetaDeleteOps returns a boolean if a field has been set.

### GetMetaListLatencyUs

`func (o *DfsSMBShareStat) GetMetaListLatencyUs() int64`

GetMetaListLatencyUs returns the MetaListLatencyUs field if non-nil, zero value otherwise.

### GetMetaListLatencyUsOk

`func (o *DfsSMBShareStat) GetMetaListLatencyUsOk() (*int64, bool)`

GetMetaListLatencyUsOk returns a tuple with the MetaListLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaListLatencyUs

`func (o *DfsSMBShareStat) SetMetaListLatencyUs(v int64)`

SetMetaListLatencyUs sets MetaListLatencyUs field to given value.

### HasMetaListLatencyUs

`func (o *DfsSMBShareStat) HasMetaListLatencyUs() bool`

HasMetaListLatencyUs returns a boolean if a field has been set.

### GetMetaListOps

`func (o *DfsSMBShareStat) GetMetaListOps() int64`

GetMetaListOps returns the MetaListOps field if non-nil, zero value otherwise.

### GetMetaListOpsOk

`func (o *DfsSMBShareStat) GetMetaListOpsOk() (*int64, bool)`

GetMetaListOpsOk returns a tuple with the MetaListOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaListOps

`func (o *DfsSMBShareStat) SetMetaListOps(v int64)`

SetMetaListOps sets MetaListOps field to given value.

### HasMetaListOps

`func (o *DfsSMBShareStat) HasMetaListOps() bool`

HasMetaListOps returns a boolean if a field has been set.

### GetMetaReadLatencyUs

`func (o *DfsSMBShareStat) GetMetaReadLatencyUs() int64`

GetMetaReadLatencyUs returns the MetaReadLatencyUs field if non-nil, zero value otherwise.

### GetMetaReadLatencyUsOk

`func (o *DfsSMBShareStat) GetMetaReadLatencyUsOk() (*int64, bool)`

GetMetaReadLatencyUsOk returns a tuple with the MetaReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaReadLatencyUs

`func (o *DfsSMBShareStat) SetMetaReadLatencyUs(v int64)`

SetMetaReadLatencyUs sets MetaReadLatencyUs field to given value.

### HasMetaReadLatencyUs

`func (o *DfsSMBShareStat) HasMetaReadLatencyUs() bool`

HasMetaReadLatencyUs returns a boolean if a field has been set.

### GetMetaReadOps

`func (o *DfsSMBShareStat) GetMetaReadOps() int64`

GetMetaReadOps returns the MetaReadOps field if non-nil, zero value otherwise.

### GetMetaReadOpsOk

`func (o *DfsSMBShareStat) GetMetaReadOpsOk() (*int64, bool)`

GetMetaReadOpsOk returns a tuple with the MetaReadOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaReadOps

`func (o *DfsSMBShareStat) SetMetaReadOps(v int64)`

SetMetaReadOps sets MetaReadOps field to given value.

### HasMetaReadOps

`func (o *DfsSMBShareStat) HasMetaReadOps() bool`

HasMetaReadOps returns a boolean if a field has been set.

### GetMetaWriteLatencyUs

`func (o *DfsSMBShareStat) GetMetaWriteLatencyUs() int64`

GetMetaWriteLatencyUs returns the MetaWriteLatencyUs field if non-nil, zero value otherwise.

### GetMetaWriteLatencyUsOk

`func (o *DfsSMBShareStat) GetMetaWriteLatencyUsOk() (*int64, bool)`

GetMetaWriteLatencyUsOk returns a tuple with the MetaWriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaWriteLatencyUs

`func (o *DfsSMBShareStat) SetMetaWriteLatencyUs(v int64)`

SetMetaWriteLatencyUs sets MetaWriteLatencyUs field to given value.

### HasMetaWriteLatencyUs

`func (o *DfsSMBShareStat) HasMetaWriteLatencyUs() bool`

HasMetaWriteLatencyUs returns a boolean if a field has been set.

### GetMetaWriteOps

`func (o *DfsSMBShareStat) GetMetaWriteOps() int64`

GetMetaWriteOps returns the MetaWriteOps field if non-nil, zero value otherwise.

### GetMetaWriteOpsOk

`func (o *DfsSMBShareStat) GetMetaWriteOpsOk() (*int64, bool)`

GetMetaWriteOpsOk returns a tuple with the MetaWriteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaWriteOps

`func (o *DfsSMBShareStat) SetMetaWriteOps(v int64)`

SetMetaWriteOps sets MetaWriteOps field to given value.

### HasMetaWriteOps

`func (o *DfsSMBShareStat) HasMetaWriteOps() bool`

HasMetaWriteOps returns a boolean if a field has been set.

### GetReadAheadBandwidthKbyte

`func (o *DfsSMBShareStat) GetReadAheadBandwidthKbyte() int64`

GetReadAheadBandwidthKbyte returns the ReadAheadBandwidthKbyte field if non-nil, zero value otherwise.

### GetReadAheadBandwidthKbyteOk

`func (o *DfsSMBShareStat) GetReadAheadBandwidthKbyteOk() (*int64, bool)`

GetReadAheadBandwidthKbyteOk returns a tuple with the ReadAheadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadBandwidthKbyte

`func (o *DfsSMBShareStat) SetReadAheadBandwidthKbyte(v int64)`

SetReadAheadBandwidthKbyte sets ReadAheadBandwidthKbyte field to given value.

### HasReadAheadBandwidthKbyte

`func (o *DfsSMBShareStat) HasReadAheadBandwidthKbyte() bool`

HasReadAheadBandwidthKbyte returns a boolean if a field has been set.

### GetReadAheadIops

`func (o *DfsSMBShareStat) GetReadAheadIops() int64`

GetReadAheadIops returns the ReadAheadIops field if non-nil, zero value otherwise.

### GetReadAheadIopsOk

`func (o *DfsSMBShareStat) GetReadAheadIopsOk() (*int64, bool)`

GetReadAheadIopsOk returns a tuple with the ReadAheadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadIops

`func (o *DfsSMBShareStat) SetReadAheadIops(v int64)`

SetReadAheadIops sets ReadAheadIops field to given value.

### HasReadAheadIops

`func (o *DfsSMBShareStat) HasReadAheadIops() bool`

HasReadAheadIops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


