# DfsFTPShareStat

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

### NewDfsFTPShareStat

`func NewDfsFTPShareStat() *DfsFTPShareStat`

NewDfsFTPShareStat instantiates a new DfsFTPShareStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsFTPShareStatWithDefaults

`func NewDfsFTPShareStatWithDefaults() *DfsFTPShareStat`

NewDfsFTPShareStatWithDefaults instantiates a new DfsFTPShareStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *DfsFTPShareStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsFTPShareStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsFTPShareStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsFTPShareStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataCacheHitRate

`func (o *DfsFTPShareStat) GetDataCacheHitRate() float64`

GetDataCacheHitRate returns the DataCacheHitRate field if non-nil, zero value otherwise.

### GetDataCacheHitRateOk

`func (o *DfsFTPShareStat) GetDataCacheHitRateOk() (*float64, bool)`

GetDataCacheHitRateOk returns a tuple with the DataCacheHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCacheHitRate

`func (o *DfsFTPShareStat) SetDataCacheHitRate(v float64)`

SetDataCacheHitRate sets DataCacheHitRate field to given value.

### HasDataCacheHitRate

`func (o *DfsFTPShareStat) HasDataCacheHitRate() bool`

HasDataCacheHitRate returns a boolean if a field has been set.

### GetDataDeleteBandwidthKbyte

`func (o *DfsFTPShareStat) GetDataDeleteBandwidthKbyte() int64`

GetDataDeleteBandwidthKbyte returns the DataDeleteBandwidthKbyte field if non-nil, zero value otherwise.

### GetDataDeleteBandwidthKbyteOk

`func (o *DfsFTPShareStat) GetDataDeleteBandwidthKbyteOk() (*int64, bool)`

GetDataDeleteBandwidthKbyteOk returns a tuple with the DataDeleteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeleteBandwidthKbyte

`func (o *DfsFTPShareStat) SetDataDeleteBandwidthKbyte(v int64)`

SetDataDeleteBandwidthKbyte sets DataDeleteBandwidthKbyte field to given value.

### HasDataDeleteBandwidthKbyte

`func (o *DfsFTPShareStat) HasDataDeleteBandwidthKbyte() bool`

HasDataDeleteBandwidthKbyte returns a boolean if a field has been set.

### GetDataDeleteIops

`func (o *DfsFTPShareStat) GetDataDeleteIops() int64`

GetDataDeleteIops returns the DataDeleteIops field if non-nil, zero value otherwise.

### GetDataDeleteIopsOk

`func (o *DfsFTPShareStat) GetDataDeleteIopsOk() (*int64, bool)`

GetDataDeleteIopsOk returns a tuple with the DataDeleteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeleteIops

`func (o *DfsFTPShareStat) SetDataDeleteIops(v int64)`

SetDataDeleteIops sets DataDeleteIops field to given value.

### HasDataDeleteIops

`func (o *DfsFTPShareStat) HasDataDeleteIops() bool`

HasDataDeleteIops returns a boolean if a field has been set.

### GetDataDeleteLatencyUs

`func (o *DfsFTPShareStat) GetDataDeleteLatencyUs() int64`

GetDataDeleteLatencyUs returns the DataDeleteLatencyUs field if non-nil, zero value otherwise.

### GetDataDeleteLatencyUsOk

`func (o *DfsFTPShareStat) GetDataDeleteLatencyUsOk() (*int64, bool)`

GetDataDeleteLatencyUsOk returns a tuple with the DataDeleteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeleteLatencyUs

`func (o *DfsFTPShareStat) SetDataDeleteLatencyUs(v int64)`

SetDataDeleteLatencyUs sets DataDeleteLatencyUs field to given value.

### HasDataDeleteLatencyUs

`func (o *DfsFTPShareStat) HasDataDeleteLatencyUs() bool`

HasDataDeleteLatencyUs returns a boolean if a field has been set.

### GetDataReadBandwidthKbyte

`func (o *DfsFTPShareStat) GetDataReadBandwidthKbyte() int64`

GetDataReadBandwidthKbyte returns the DataReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetDataReadBandwidthKbyteOk

`func (o *DfsFTPShareStat) GetDataReadBandwidthKbyteOk() (*int64, bool)`

GetDataReadBandwidthKbyteOk returns a tuple with the DataReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadBandwidthKbyte

`func (o *DfsFTPShareStat) SetDataReadBandwidthKbyte(v int64)`

SetDataReadBandwidthKbyte sets DataReadBandwidthKbyte field to given value.

### HasDataReadBandwidthKbyte

`func (o *DfsFTPShareStat) HasDataReadBandwidthKbyte() bool`

HasDataReadBandwidthKbyte returns a boolean if a field has been set.

### GetDataReadIops

`func (o *DfsFTPShareStat) GetDataReadIops() int64`

GetDataReadIops returns the DataReadIops field if non-nil, zero value otherwise.

### GetDataReadIopsOk

`func (o *DfsFTPShareStat) GetDataReadIopsOk() (*int64, bool)`

GetDataReadIopsOk returns a tuple with the DataReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadIops

`func (o *DfsFTPShareStat) SetDataReadIops(v int64)`

SetDataReadIops sets DataReadIops field to given value.

### HasDataReadIops

`func (o *DfsFTPShareStat) HasDataReadIops() bool`

HasDataReadIops returns a boolean if a field has been set.

### GetDataReadLatencyUs

`func (o *DfsFTPShareStat) GetDataReadLatencyUs() int64`

GetDataReadLatencyUs returns the DataReadLatencyUs field if non-nil, zero value otherwise.

### GetDataReadLatencyUsOk

`func (o *DfsFTPShareStat) GetDataReadLatencyUsOk() (*int64, bool)`

GetDataReadLatencyUsOk returns a tuple with the DataReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadLatencyUs

`func (o *DfsFTPShareStat) SetDataReadLatencyUs(v int64)`

SetDataReadLatencyUs sets DataReadLatencyUs field to given value.

### HasDataReadLatencyUs

`func (o *DfsFTPShareStat) HasDataReadLatencyUs() bool`

HasDataReadLatencyUs returns a boolean if a field has been set.

### GetDataWriteBandwidthKbyte

`func (o *DfsFTPShareStat) GetDataWriteBandwidthKbyte() int64`

GetDataWriteBandwidthKbyte returns the DataWriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetDataWriteBandwidthKbyteOk

`func (o *DfsFTPShareStat) GetDataWriteBandwidthKbyteOk() (*int64, bool)`

GetDataWriteBandwidthKbyteOk returns a tuple with the DataWriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWriteBandwidthKbyte

`func (o *DfsFTPShareStat) SetDataWriteBandwidthKbyte(v int64)`

SetDataWriteBandwidthKbyte sets DataWriteBandwidthKbyte field to given value.

### HasDataWriteBandwidthKbyte

`func (o *DfsFTPShareStat) HasDataWriteBandwidthKbyte() bool`

HasDataWriteBandwidthKbyte returns a boolean if a field has been set.

### GetDataWriteIops

`func (o *DfsFTPShareStat) GetDataWriteIops() int64`

GetDataWriteIops returns the DataWriteIops field if non-nil, zero value otherwise.

### GetDataWriteIopsOk

`func (o *DfsFTPShareStat) GetDataWriteIopsOk() (*int64, bool)`

GetDataWriteIopsOk returns a tuple with the DataWriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWriteIops

`func (o *DfsFTPShareStat) SetDataWriteIops(v int64)`

SetDataWriteIops sets DataWriteIops field to given value.

### HasDataWriteIops

`func (o *DfsFTPShareStat) HasDataWriteIops() bool`

HasDataWriteIops returns a boolean if a field has been set.

### GetDataWriteLatencyUs

`func (o *DfsFTPShareStat) GetDataWriteLatencyUs() int64`

GetDataWriteLatencyUs returns the DataWriteLatencyUs field if non-nil, zero value otherwise.

### GetDataWriteLatencyUsOk

`func (o *DfsFTPShareStat) GetDataWriteLatencyUsOk() (*int64, bool)`

GetDataWriteLatencyUsOk returns a tuple with the DataWriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWriteLatencyUs

`func (o *DfsFTPShareStat) SetDataWriteLatencyUs(v int64)`

SetDataWriteLatencyUs sets DataWriteLatencyUs field to given value.

### HasDataWriteLatencyUs

`func (o *DfsFTPShareStat) HasDataWriteLatencyUs() bool`

HasDataWriteLatencyUs returns a boolean if a field has been set.

### GetMetaCacheHitRate

`func (o *DfsFTPShareStat) GetMetaCacheHitRate() float64`

GetMetaCacheHitRate returns the MetaCacheHitRate field if non-nil, zero value otherwise.

### GetMetaCacheHitRateOk

`func (o *DfsFTPShareStat) GetMetaCacheHitRateOk() (*float64, bool)`

GetMetaCacheHitRateOk returns a tuple with the MetaCacheHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaCacheHitRate

`func (o *DfsFTPShareStat) SetMetaCacheHitRate(v float64)`

SetMetaCacheHitRate sets MetaCacheHitRate field to given value.

### HasMetaCacheHitRate

`func (o *DfsFTPShareStat) HasMetaCacheHitRate() bool`

HasMetaCacheHitRate returns a boolean if a field has been set.

### GetMetaDeleteLatencyUs

`func (o *DfsFTPShareStat) GetMetaDeleteLatencyUs() int64`

GetMetaDeleteLatencyUs returns the MetaDeleteLatencyUs field if non-nil, zero value otherwise.

### GetMetaDeleteLatencyUsOk

`func (o *DfsFTPShareStat) GetMetaDeleteLatencyUsOk() (*int64, bool)`

GetMetaDeleteLatencyUsOk returns a tuple with the MetaDeleteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDeleteLatencyUs

`func (o *DfsFTPShareStat) SetMetaDeleteLatencyUs(v int64)`

SetMetaDeleteLatencyUs sets MetaDeleteLatencyUs field to given value.

### HasMetaDeleteLatencyUs

`func (o *DfsFTPShareStat) HasMetaDeleteLatencyUs() bool`

HasMetaDeleteLatencyUs returns a boolean if a field has been set.

### GetMetaDeleteOps

`func (o *DfsFTPShareStat) GetMetaDeleteOps() int64`

GetMetaDeleteOps returns the MetaDeleteOps field if non-nil, zero value otherwise.

### GetMetaDeleteOpsOk

`func (o *DfsFTPShareStat) GetMetaDeleteOpsOk() (*int64, bool)`

GetMetaDeleteOpsOk returns a tuple with the MetaDeleteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDeleteOps

`func (o *DfsFTPShareStat) SetMetaDeleteOps(v int64)`

SetMetaDeleteOps sets MetaDeleteOps field to given value.

### HasMetaDeleteOps

`func (o *DfsFTPShareStat) HasMetaDeleteOps() bool`

HasMetaDeleteOps returns a boolean if a field has been set.

### GetMetaListLatencyUs

`func (o *DfsFTPShareStat) GetMetaListLatencyUs() int64`

GetMetaListLatencyUs returns the MetaListLatencyUs field if non-nil, zero value otherwise.

### GetMetaListLatencyUsOk

`func (o *DfsFTPShareStat) GetMetaListLatencyUsOk() (*int64, bool)`

GetMetaListLatencyUsOk returns a tuple with the MetaListLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaListLatencyUs

`func (o *DfsFTPShareStat) SetMetaListLatencyUs(v int64)`

SetMetaListLatencyUs sets MetaListLatencyUs field to given value.

### HasMetaListLatencyUs

`func (o *DfsFTPShareStat) HasMetaListLatencyUs() bool`

HasMetaListLatencyUs returns a boolean if a field has been set.

### GetMetaListOps

`func (o *DfsFTPShareStat) GetMetaListOps() int64`

GetMetaListOps returns the MetaListOps field if non-nil, zero value otherwise.

### GetMetaListOpsOk

`func (o *DfsFTPShareStat) GetMetaListOpsOk() (*int64, bool)`

GetMetaListOpsOk returns a tuple with the MetaListOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaListOps

`func (o *DfsFTPShareStat) SetMetaListOps(v int64)`

SetMetaListOps sets MetaListOps field to given value.

### HasMetaListOps

`func (o *DfsFTPShareStat) HasMetaListOps() bool`

HasMetaListOps returns a boolean if a field has been set.

### GetMetaReadLatencyUs

`func (o *DfsFTPShareStat) GetMetaReadLatencyUs() int64`

GetMetaReadLatencyUs returns the MetaReadLatencyUs field if non-nil, zero value otherwise.

### GetMetaReadLatencyUsOk

`func (o *DfsFTPShareStat) GetMetaReadLatencyUsOk() (*int64, bool)`

GetMetaReadLatencyUsOk returns a tuple with the MetaReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaReadLatencyUs

`func (o *DfsFTPShareStat) SetMetaReadLatencyUs(v int64)`

SetMetaReadLatencyUs sets MetaReadLatencyUs field to given value.

### HasMetaReadLatencyUs

`func (o *DfsFTPShareStat) HasMetaReadLatencyUs() bool`

HasMetaReadLatencyUs returns a boolean if a field has been set.

### GetMetaReadOps

`func (o *DfsFTPShareStat) GetMetaReadOps() int64`

GetMetaReadOps returns the MetaReadOps field if non-nil, zero value otherwise.

### GetMetaReadOpsOk

`func (o *DfsFTPShareStat) GetMetaReadOpsOk() (*int64, bool)`

GetMetaReadOpsOk returns a tuple with the MetaReadOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaReadOps

`func (o *DfsFTPShareStat) SetMetaReadOps(v int64)`

SetMetaReadOps sets MetaReadOps field to given value.

### HasMetaReadOps

`func (o *DfsFTPShareStat) HasMetaReadOps() bool`

HasMetaReadOps returns a boolean if a field has been set.

### GetMetaWriteLatencyUs

`func (o *DfsFTPShareStat) GetMetaWriteLatencyUs() int64`

GetMetaWriteLatencyUs returns the MetaWriteLatencyUs field if non-nil, zero value otherwise.

### GetMetaWriteLatencyUsOk

`func (o *DfsFTPShareStat) GetMetaWriteLatencyUsOk() (*int64, bool)`

GetMetaWriteLatencyUsOk returns a tuple with the MetaWriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaWriteLatencyUs

`func (o *DfsFTPShareStat) SetMetaWriteLatencyUs(v int64)`

SetMetaWriteLatencyUs sets MetaWriteLatencyUs field to given value.

### HasMetaWriteLatencyUs

`func (o *DfsFTPShareStat) HasMetaWriteLatencyUs() bool`

HasMetaWriteLatencyUs returns a boolean if a field has been set.

### GetMetaWriteOps

`func (o *DfsFTPShareStat) GetMetaWriteOps() int64`

GetMetaWriteOps returns the MetaWriteOps field if non-nil, zero value otherwise.

### GetMetaWriteOpsOk

`func (o *DfsFTPShareStat) GetMetaWriteOpsOk() (*int64, bool)`

GetMetaWriteOpsOk returns a tuple with the MetaWriteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaWriteOps

`func (o *DfsFTPShareStat) SetMetaWriteOps(v int64)`

SetMetaWriteOps sets MetaWriteOps field to given value.

### HasMetaWriteOps

`func (o *DfsFTPShareStat) HasMetaWriteOps() bool`

HasMetaWriteOps returns a boolean if a field has been set.

### GetReadAheadBandwidthKbyte

`func (o *DfsFTPShareStat) GetReadAheadBandwidthKbyte() int64`

GetReadAheadBandwidthKbyte returns the ReadAheadBandwidthKbyte field if non-nil, zero value otherwise.

### GetReadAheadBandwidthKbyteOk

`func (o *DfsFTPShareStat) GetReadAheadBandwidthKbyteOk() (*int64, bool)`

GetReadAheadBandwidthKbyteOk returns a tuple with the ReadAheadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadBandwidthKbyte

`func (o *DfsFTPShareStat) SetReadAheadBandwidthKbyte(v int64)`

SetReadAheadBandwidthKbyte sets ReadAheadBandwidthKbyte field to given value.

### HasReadAheadBandwidthKbyte

`func (o *DfsFTPShareStat) HasReadAheadBandwidthKbyte() bool`

HasReadAheadBandwidthKbyte returns a boolean if a field has been set.

### GetReadAheadIops

`func (o *DfsFTPShareStat) GetReadAheadIops() int64`

GetReadAheadIops returns the ReadAheadIops field if non-nil, zero value otherwise.

### GetReadAheadIopsOk

`func (o *DfsFTPShareStat) GetReadAheadIopsOk() (*int64, bool)`

GetReadAheadIopsOk returns a tuple with the ReadAheadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadIops

`func (o *DfsFTPShareStat) SetReadAheadIops(v int64)`

SetReadAheadIops sets ReadAheadIops field to given value.

### HasReadAheadIops

`func (o *DfsFTPShareStat) HasReadAheadIops() bool`

HasReadAheadIops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


