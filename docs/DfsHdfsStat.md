# DfsHdfsStat

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

### NewDfsHdfsStat

`func NewDfsHdfsStat() *DfsHdfsStat`

NewDfsHdfsStat instantiates a new DfsHdfsStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsHdfsStatWithDefaults

`func NewDfsHdfsStatWithDefaults() *DfsHdfsStat`

NewDfsHdfsStatWithDefaults instantiates a new DfsHdfsStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *DfsHdfsStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsHdfsStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsHdfsStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsHdfsStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataCacheHitRate

`func (o *DfsHdfsStat) GetDataCacheHitRate() float64`

GetDataCacheHitRate returns the DataCacheHitRate field if non-nil, zero value otherwise.

### GetDataCacheHitRateOk

`func (o *DfsHdfsStat) GetDataCacheHitRateOk() (*float64, bool)`

GetDataCacheHitRateOk returns a tuple with the DataCacheHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCacheHitRate

`func (o *DfsHdfsStat) SetDataCacheHitRate(v float64)`

SetDataCacheHitRate sets DataCacheHitRate field to given value.

### HasDataCacheHitRate

`func (o *DfsHdfsStat) HasDataCacheHitRate() bool`

HasDataCacheHitRate returns a boolean if a field has been set.

### GetDataDeleteBandwidthKbyte

`func (o *DfsHdfsStat) GetDataDeleteBandwidthKbyte() int64`

GetDataDeleteBandwidthKbyte returns the DataDeleteBandwidthKbyte field if non-nil, zero value otherwise.

### GetDataDeleteBandwidthKbyteOk

`func (o *DfsHdfsStat) GetDataDeleteBandwidthKbyteOk() (*int64, bool)`

GetDataDeleteBandwidthKbyteOk returns a tuple with the DataDeleteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeleteBandwidthKbyte

`func (o *DfsHdfsStat) SetDataDeleteBandwidthKbyte(v int64)`

SetDataDeleteBandwidthKbyte sets DataDeleteBandwidthKbyte field to given value.

### HasDataDeleteBandwidthKbyte

`func (o *DfsHdfsStat) HasDataDeleteBandwidthKbyte() bool`

HasDataDeleteBandwidthKbyte returns a boolean if a field has been set.

### GetDataDeleteIops

`func (o *DfsHdfsStat) GetDataDeleteIops() int64`

GetDataDeleteIops returns the DataDeleteIops field if non-nil, zero value otherwise.

### GetDataDeleteIopsOk

`func (o *DfsHdfsStat) GetDataDeleteIopsOk() (*int64, bool)`

GetDataDeleteIopsOk returns a tuple with the DataDeleteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeleteIops

`func (o *DfsHdfsStat) SetDataDeleteIops(v int64)`

SetDataDeleteIops sets DataDeleteIops field to given value.

### HasDataDeleteIops

`func (o *DfsHdfsStat) HasDataDeleteIops() bool`

HasDataDeleteIops returns a boolean if a field has been set.

### GetDataDeleteLatencyUs

`func (o *DfsHdfsStat) GetDataDeleteLatencyUs() int64`

GetDataDeleteLatencyUs returns the DataDeleteLatencyUs field if non-nil, zero value otherwise.

### GetDataDeleteLatencyUsOk

`func (o *DfsHdfsStat) GetDataDeleteLatencyUsOk() (*int64, bool)`

GetDataDeleteLatencyUsOk returns a tuple with the DataDeleteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeleteLatencyUs

`func (o *DfsHdfsStat) SetDataDeleteLatencyUs(v int64)`

SetDataDeleteLatencyUs sets DataDeleteLatencyUs field to given value.

### HasDataDeleteLatencyUs

`func (o *DfsHdfsStat) HasDataDeleteLatencyUs() bool`

HasDataDeleteLatencyUs returns a boolean if a field has been set.

### GetDataReadBandwidthKbyte

`func (o *DfsHdfsStat) GetDataReadBandwidthKbyte() int64`

GetDataReadBandwidthKbyte returns the DataReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetDataReadBandwidthKbyteOk

`func (o *DfsHdfsStat) GetDataReadBandwidthKbyteOk() (*int64, bool)`

GetDataReadBandwidthKbyteOk returns a tuple with the DataReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadBandwidthKbyte

`func (o *DfsHdfsStat) SetDataReadBandwidthKbyte(v int64)`

SetDataReadBandwidthKbyte sets DataReadBandwidthKbyte field to given value.

### HasDataReadBandwidthKbyte

`func (o *DfsHdfsStat) HasDataReadBandwidthKbyte() bool`

HasDataReadBandwidthKbyte returns a boolean if a field has been set.

### GetDataReadIops

`func (o *DfsHdfsStat) GetDataReadIops() int64`

GetDataReadIops returns the DataReadIops field if non-nil, zero value otherwise.

### GetDataReadIopsOk

`func (o *DfsHdfsStat) GetDataReadIopsOk() (*int64, bool)`

GetDataReadIopsOk returns a tuple with the DataReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadIops

`func (o *DfsHdfsStat) SetDataReadIops(v int64)`

SetDataReadIops sets DataReadIops field to given value.

### HasDataReadIops

`func (o *DfsHdfsStat) HasDataReadIops() bool`

HasDataReadIops returns a boolean if a field has been set.

### GetDataReadLatencyUs

`func (o *DfsHdfsStat) GetDataReadLatencyUs() int64`

GetDataReadLatencyUs returns the DataReadLatencyUs field if non-nil, zero value otherwise.

### GetDataReadLatencyUsOk

`func (o *DfsHdfsStat) GetDataReadLatencyUsOk() (*int64, bool)`

GetDataReadLatencyUsOk returns a tuple with the DataReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadLatencyUs

`func (o *DfsHdfsStat) SetDataReadLatencyUs(v int64)`

SetDataReadLatencyUs sets DataReadLatencyUs field to given value.

### HasDataReadLatencyUs

`func (o *DfsHdfsStat) HasDataReadLatencyUs() bool`

HasDataReadLatencyUs returns a boolean if a field has been set.

### GetDataWriteBandwidthKbyte

`func (o *DfsHdfsStat) GetDataWriteBandwidthKbyte() int64`

GetDataWriteBandwidthKbyte returns the DataWriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetDataWriteBandwidthKbyteOk

`func (o *DfsHdfsStat) GetDataWriteBandwidthKbyteOk() (*int64, bool)`

GetDataWriteBandwidthKbyteOk returns a tuple with the DataWriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWriteBandwidthKbyte

`func (o *DfsHdfsStat) SetDataWriteBandwidthKbyte(v int64)`

SetDataWriteBandwidthKbyte sets DataWriteBandwidthKbyte field to given value.

### HasDataWriteBandwidthKbyte

`func (o *DfsHdfsStat) HasDataWriteBandwidthKbyte() bool`

HasDataWriteBandwidthKbyte returns a boolean if a field has been set.

### GetDataWriteIops

`func (o *DfsHdfsStat) GetDataWriteIops() int64`

GetDataWriteIops returns the DataWriteIops field if non-nil, zero value otherwise.

### GetDataWriteIopsOk

`func (o *DfsHdfsStat) GetDataWriteIopsOk() (*int64, bool)`

GetDataWriteIopsOk returns a tuple with the DataWriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWriteIops

`func (o *DfsHdfsStat) SetDataWriteIops(v int64)`

SetDataWriteIops sets DataWriteIops field to given value.

### HasDataWriteIops

`func (o *DfsHdfsStat) HasDataWriteIops() bool`

HasDataWriteIops returns a boolean if a field has been set.

### GetDataWriteLatencyUs

`func (o *DfsHdfsStat) GetDataWriteLatencyUs() int64`

GetDataWriteLatencyUs returns the DataWriteLatencyUs field if non-nil, zero value otherwise.

### GetDataWriteLatencyUsOk

`func (o *DfsHdfsStat) GetDataWriteLatencyUsOk() (*int64, bool)`

GetDataWriteLatencyUsOk returns a tuple with the DataWriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWriteLatencyUs

`func (o *DfsHdfsStat) SetDataWriteLatencyUs(v int64)`

SetDataWriteLatencyUs sets DataWriteLatencyUs field to given value.

### HasDataWriteLatencyUs

`func (o *DfsHdfsStat) HasDataWriteLatencyUs() bool`

HasDataWriteLatencyUs returns a boolean if a field has been set.

### GetMetaCacheHitRate

`func (o *DfsHdfsStat) GetMetaCacheHitRate() float64`

GetMetaCacheHitRate returns the MetaCacheHitRate field if non-nil, zero value otherwise.

### GetMetaCacheHitRateOk

`func (o *DfsHdfsStat) GetMetaCacheHitRateOk() (*float64, bool)`

GetMetaCacheHitRateOk returns a tuple with the MetaCacheHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaCacheHitRate

`func (o *DfsHdfsStat) SetMetaCacheHitRate(v float64)`

SetMetaCacheHitRate sets MetaCacheHitRate field to given value.

### HasMetaCacheHitRate

`func (o *DfsHdfsStat) HasMetaCacheHitRate() bool`

HasMetaCacheHitRate returns a boolean if a field has been set.

### GetMetaDeleteLatencyUs

`func (o *DfsHdfsStat) GetMetaDeleteLatencyUs() int64`

GetMetaDeleteLatencyUs returns the MetaDeleteLatencyUs field if non-nil, zero value otherwise.

### GetMetaDeleteLatencyUsOk

`func (o *DfsHdfsStat) GetMetaDeleteLatencyUsOk() (*int64, bool)`

GetMetaDeleteLatencyUsOk returns a tuple with the MetaDeleteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDeleteLatencyUs

`func (o *DfsHdfsStat) SetMetaDeleteLatencyUs(v int64)`

SetMetaDeleteLatencyUs sets MetaDeleteLatencyUs field to given value.

### HasMetaDeleteLatencyUs

`func (o *DfsHdfsStat) HasMetaDeleteLatencyUs() bool`

HasMetaDeleteLatencyUs returns a boolean if a field has been set.

### GetMetaDeleteOps

`func (o *DfsHdfsStat) GetMetaDeleteOps() int64`

GetMetaDeleteOps returns the MetaDeleteOps field if non-nil, zero value otherwise.

### GetMetaDeleteOpsOk

`func (o *DfsHdfsStat) GetMetaDeleteOpsOk() (*int64, bool)`

GetMetaDeleteOpsOk returns a tuple with the MetaDeleteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDeleteOps

`func (o *DfsHdfsStat) SetMetaDeleteOps(v int64)`

SetMetaDeleteOps sets MetaDeleteOps field to given value.

### HasMetaDeleteOps

`func (o *DfsHdfsStat) HasMetaDeleteOps() bool`

HasMetaDeleteOps returns a boolean if a field has been set.

### GetMetaListLatencyUs

`func (o *DfsHdfsStat) GetMetaListLatencyUs() int64`

GetMetaListLatencyUs returns the MetaListLatencyUs field if non-nil, zero value otherwise.

### GetMetaListLatencyUsOk

`func (o *DfsHdfsStat) GetMetaListLatencyUsOk() (*int64, bool)`

GetMetaListLatencyUsOk returns a tuple with the MetaListLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaListLatencyUs

`func (o *DfsHdfsStat) SetMetaListLatencyUs(v int64)`

SetMetaListLatencyUs sets MetaListLatencyUs field to given value.

### HasMetaListLatencyUs

`func (o *DfsHdfsStat) HasMetaListLatencyUs() bool`

HasMetaListLatencyUs returns a boolean if a field has been set.

### GetMetaListOps

`func (o *DfsHdfsStat) GetMetaListOps() int64`

GetMetaListOps returns the MetaListOps field if non-nil, zero value otherwise.

### GetMetaListOpsOk

`func (o *DfsHdfsStat) GetMetaListOpsOk() (*int64, bool)`

GetMetaListOpsOk returns a tuple with the MetaListOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaListOps

`func (o *DfsHdfsStat) SetMetaListOps(v int64)`

SetMetaListOps sets MetaListOps field to given value.

### HasMetaListOps

`func (o *DfsHdfsStat) HasMetaListOps() bool`

HasMetaListOps returns a boolean if a field has been set.

### GetMetaReadLatencyUs

`func (o *DfsHdfsStat) GetMetaReadLatencyUs() int64`

GetMetaReadLatencyUs returns the MetaReadLatencyUs field if non-nil, zero value otherwise.

### GetMetaReadLatencyUsOk

`func (o *DfsHdfsStat) GetMetaReadLatencyUsOk() (*int64, bool)`

GetMetaReadLatencyUsOk returns a tuple with the MetaReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaReadLatencyUs

`func (o *DfsHdfsStat) SetMetaReadLatencyUs(v int64)`

SetMetaReadLatencyUs sets MetaReadLatencyUs field to given value.

### HasMetaReadLatencyUs

`func (o *DfsHdfsStat) HasMetaReadLatencyUs() bool`

HasMetaReadLatencyUs returns a boolean if a field has been set.

### GetMetaReadOps

`func (o *DfsHdfsStat) GetMetaReadOps() int64`

GetMetaReadOps returns the MetaReadOps field if non-nil, zero value otherwise.

### GetMetaReadOpsOk

`func (o *DfsHdfsStat) GetMetaReadOpsOk() (*int64, bool)`

GetMetaReadOpsOk returns a tuple with the MetaReadOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaReadOps

`func (o *DfsHdfsStat) SetMetaReadOps(v int64)`

SetMetaReadOps sets MetaReadOps field to given value.

### HasMetaReadOps

`func (o *DfsHdfsStat) HasMetaReadOps() bool`

HasMetaReadOps returns a boolean if a field has been set.

### GetMetaWriteLatencyUs

`func (o *DfsHdfsStat) GetMetaWriteLatencyUs() int64`

GetMetaWriteLatencyUs returns the MetaWriteLatencyUs field if non-nil, zero value otherwise.

### GetMetaWriteLatencyUsOk

`func (o *DfsHdfsStat) GetMetaWriteLatencyUsOk() (*int64, bool)`

GetMetaWriteLatencyUsOk returns a tuple with the MetaWriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaWriteLatencyUs

`func (o *DfsHdfsStat) SetMetaWriteLatencyUs(v int64)`

SetMetaWriteLatencyUs sets MetaWriteLatencyUs field to given value.

### HasMetaWriteLatencyUs

`func (o *DfsHdfsStat) HasMetaWriteLatencyUs() bool`

HasMetaWriteLatencyUs returns a boolean if a field has been set.

### GetMetaWriteOps

`func (o *DfsHdfsStat) GetMetaWriteOps() int64`

GetMetaWriteOps returns the MetaWriteOps field if non-nil, zero value otherwise.

### GetMetaWriteOpsOk

`func (o *DfsHdfsStat) GetMetaWriteOpsOk() (*int64, bool)`

GetMetaWriteOpsOk returns a tuple with the MetaWriteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaWriteOps

`func (o *DfsHdfsStat) SetMetaWriteOps(v int64)`

SetMetaWriteOps sets MetaWriteOps field to given value.

### HasMetaWriteOps

`func (o *DfsHdfsStat) HasMetaWriteOps() bool`

HasMetaWriteOps returns a boolean if a field has been set.

### GetReadAheadBandwidthKbyte

`func (o *DfsHdfsStat) GetReadAheadBandwidthKbyte() int64`

GetReadAheadBandwidthKbyte returns the ReadAheadBandwidthKbyte field if non-nil, zero value otherwise.

### GetReadAheadBandwidthKbyteOk

`func (o *DfsHdfsStat) GetReadAheadBandwidthKbyteOk() (*int64, bool)`

GetReadAheadBandwidthKbyteOk returns a tuple with the ReadAheadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadBandwidthKbyte

`func (o *DfsHdfsStat) SetReadAheadBandwidthKbyte(v int64)`

SetReadAheadBandwidthKbyte sets ReadAheadBandwidthKbyte field to given value.

### HasReadAheadBandwidthKbyte

`func (o *DfsHdfsStat) HasReadAheadBandwidthKbyte() bool`

HasReadAheadBandwidthKbyte returns a boolean if a field has been set.

### GetReadAheadIops

`func (o *DfsHdfsStat) GetReadAheadIops() int64`

GetReadAheadIops returns the ReadAheadIops field if non-nil, zero value otherwise.

### GetReadAheadIopsOk

`func (o *DfsHdfsStat) GetReadAheadIopsOk() (*int64, bool)`

GetReadAheadIopsOk returns a tuple with the ReadAheadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadIops

`func (o *DfsHdfsStat) SetReadAheadIops(v int64)`

SetReadAheadIops sets ReadAheadIops field to given value.

### HasReadAheadIops

`func (o *DfsHdfsStat) HasReadAheadIops() bool`

HasReadAheadIops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


