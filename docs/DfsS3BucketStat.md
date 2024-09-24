# DfsS3BucketStat

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
**AllocatedSize** | Pointer to **int64** |  | [optional] 
**Objects** | Pointer to **int64** |  | [optional] 

## Methods

### NewDfsS3BucketStat

`func NewDfsS3BucketStat() *DfsS3BucketStat`

NewDfsS3BucketStat instantiates a new DfsS3BucketStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsS3BucketStatWithDefaults

`func NewDfsS3BucketStatWithDefaults() *DfsS3BucketStat`

NewDfsS3BucketStatWithDefaults instantiates a new DfsS3BucketStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *DfsS3BucketStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsS3BucketStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsS3BucketStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsS3BucketStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataCacheHitRate

`func (o *DfsS3BucketStat) GetDataCacheHitRate() float64`

GetDataCacheHitRate returns the DataCacheHitRate field if non-nil, zero value otherwise.

### GetDataCacheHitRateOk

`func (o *DfsS3BucketStat) GetDataCacheHitRateOk() (*float64, bool)`

GetDataCacheHitRateOk returns a tuple with the DataCacheHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCacheHitRate

`func (o *DfsS3BucketStat) SetDataCacheHitRate(v float64)`

SetDataCacheHitRate sets DataCacheHitRate field to given value.

### HasDataCacheHitRate

`func (o *DfsS3BucketStat) HasDataCacheHitRate() bool`

HasDataCacheHitRate returns a boolean if a field has been set.

### GetDataDeleteBandwidthKbyte

`func (o *DfsS3BucketStat) GetDataDeleteBandwidthKbyte() int64`

GetDataDeleteBandwidthKbyte returns the DataDeleteBandwidthKbyte field if non-nil, zero value otherwise.

### GetDataDeleteBandwidthKbyteOk

`func (o *DfsS3BucketStat) GetDataDeleteBandwidthKbyteOk() (*int64, bool)`

GetDataDeleteBandwidthKbyteOk returns a tuple with the DataDeleteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeleteBandwidthKbyte

`func (o *DfsS3BucketStat) SetDataDeleteBandwidthKbyte(v int64)`

SetDataDeleteBandwidthKbyte sets DataDeleteBandwidthKbyte field to given value.

### HasDataDeleteBandwidthKbyte

`func (o *DfsS3BucketStat) HasDataDeleteBandwidthKbyte() bool`

HasDataDeleteBandwidthKbyte returns a boolean if a field has been set.

### GetDataDeleteIops

`func (o *DfsS3BucketStat) GetDataDeleteIops() int64`

GetDataDeleteIops returns the DataDeleteIops field if non-nil, zero value otherwise.

### GetDataDeleteIopsOk

`func (o *DfsS3BucketStat) GetDataDeleteIopsOk() (*int64, bool)`

GetDataDeleteIopsOk returns a tuple with the DataDeleteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeleteIops

`func (o *DfsS3BucketStat) SetDataDeleteIops(v int64)`

SetDataDeleteIops sets DataDeleteIops field to given value.

### HasDataDeleteIops

`func (o *DfsS3BucketStat) HasDataDeleteIops() bool`

HasDataDeleteIops returns a boolean if a field has been set.

### GetDataDeleteLatencyUs

`func (o *DfsS3BucketStat) GetDataDeleteLatencyUs() int64`

GetDataDeleteLatencyUs returns the DataDeleteLatencyUs field if non-nil, zero value otherwise.

### GetDataDeleteLatencyUsOk

`func (o *DfsS3BucketStat) GetDataDeleteLatencyUsOk() (*int64, bool)`

GetDataDeleteLatencyUsOk returns a tuple with the DataDeleteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeleteLatencyUs

`func (o *DfsS3BucketStat) SetDataDeleteLatencyUs(v int64)`

SetDataDeleteLatencyUs sets DataDeleteLatencyUs field to given value.

### HasDataDeleteLatencyUs

`func (o *DfsS3BucketStat) HasDataDeleteLatencyUs() bool`

HasDataDeleteLatencyUs returns a boolean if a field has been set.

### GetDataReadBandwidthKbyte

`func (o *DfsS3BucketStat) GetDataReadBandwidthKbyte() int64`

GetDataReadBandwidthKbyte returns the DataReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetDataReadBandwidthKbyteOk

`func (o *DfsS3BucketStat) GetDataReadBandwidthKbyteOk() (*int64, bool)`

GetDataReadBandwidthKbyteOk returns a tuple with the DataReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadBandwidthKbyte

`func (o *DfsS3BucketStat) SetDataReadBandwidthKbyte(v int64)`

SetDataReadBandwidthKbyte sets DataReadBandwidthKbyte field to given value.

### HasDataReadBandwidthKbyte

`func (o *DfsS3BucketStat) HasDataReadBandwidthKbyte() bool`

HasDataReadBandwidthKbyte returns a boolean if a field has been set.

### GetDataReadIops

`func (o *DfsS3BucketStat) GetDataReadIops() int64`

GetDataReadIops returns the DataReadIops field if non-nil, zero value otherwise.

### GetDataReadIopsOk

`func (o *DfsS3BucketStat) GetDataReadIopsOk() (*int64, bool)`

GetDataReadIopsOk returns a tuple with the DataReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadIops

`func (o *DfsS3BucketStat) SetDataReadIops(v int64)`

SetDataReadIops sets DataReadIops field to given value.

### HasDataReadIops

`func (o *DfsS3BucketStat) HasDataReadIops() bool`

HasDataReadIops returns a boolean if a field has been set.

### GetDataReadLatencyUs

`func (o *DfsS3BucketStat) GetDataReadLatencyUs() int64`

GetDataReadLatencyUs returns the DataReadLatencyUs field if non-nil, zero value otherwise.

### GetDataReadLatencyUsOk

`func (o *DfsS3BucketStat) GetDataReadLatencyUsOk() (*int64, bool)`

GetDataReadLatencyUsOk returns a tuple with the DataReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadLatencyUs

`func (o *DfsS3BucketStat) SetDataReadLatencyUs(v int64)`

SetDataReadLatencyUs sets DataReadLatencyUs field to given value.

### HasDataReadLatencyUs

`func (o *DfsS3BucketStat) HasDataReadLatencyUs() bool`

HasDataReadLatencyUs returns a boolean if a field has been set.

### GetDataWriteBandwidthKbyte

`func (o *DfsS3BucketStat) GetDataWriteBandwidthKbyte() int64`

GetDataWriteBandwidthKbyte returns the DataWriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetDataWriteBandwidthKbyteOk

`func (o *DfsS3BucketStat) GetDataWriteBandwidthKbyteOk() (*int64, bool)`

GetDataWriteBandwidthKbyteOk returns a tuple with the DataWriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWriteBandwidthKbyte

`func (o *DfsS3BucketStat) SetDataWriteBandwidthKbyte(v int64)`

SetDataWriteBandwidthKbyte sets DataWriteBandwidthKbyte field to given value.

### HasDataWriteBandwidthKbyte

`func (o *DfsS3BucketStat) HasDataWriteBandwidthKbyte() bool`

HasDataWriteBandwidthKbyte returns a boolean if a field has been set.

### GetDataWriteIops

`func (o *DfsS3BucketStat) GetDataWriteIops() int64`

GetDataWriteIops returns the DataWriteIops field if non-nil, zero value otherwise.

### GetDataWriteIopsOk

`func (o *DfsS3BucketStat) GetDataWriteIopsOk() (*int64, bool)`

GetDataWriteIopsOk returns a tuple with the DataWriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWriteIops

`func (o *DfsS3BucketStat) SetDataWriteIops(v int64)`

SetDataWriteIops sets DataWriteIops field to given value.

### HasDataWriteIops

`func (o *DfsS3BucketStat) HasDataWriteIops() bool`

HasDataWriteIops returns a boolean if a field has been set.

### GetDataWriteLatencyUs

`func (o *DfsS3BucketStat) GetDataWriteLatencyUs() int64`

GetDataWriteLatencyUs returns the DataWriteLatencyUs field if non-nil, zero value otherwise.

### GetDataWriteLatencyUsOk

`func (o *DfsS3BucketStat) GetDataWriteLatencyUsOk() (*int64, bool)`

GetDataWriteLatencyUsOk returns a tuple with the DataWriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWriteLatencyUs

`func (o *DfsS3BucketStat) SetDataWriteLatencyUs(v int64)`

SetDataWriteLatencyUs sets DataWriteLatencyUs field to given value.

### HasDataWriteLatencyUs

`func (o *DfsS3BucketStat) HasDataWriteLatencyUs() bool`

HasDataWriteLatencyUs returns a boolean if a field has been set.

### GetMetaCacheHitRate

`func (o *DfsS3BucketStat) GetMetaCacheHitRate() float64`

GetMetaCacheHitRate returns the MetaCacheHitRate field if non-nil, zero value otherwise.

### GetMetaCacheHitRateOk

`func (o *DfsS3BucketStat) GetMetaCacheHitRateOk() (*float64, bool)`

GetMetaCacheHitRateOk returns a tuple with the MetaCacheHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaCacheHitRate

`func (o *DfsS3BucketStat) SetMetaCacheHitRate(v float64)`

SetMetaCacheHitRate sets MetaCacheHitRate field to given value.

### HasMetaCacheHitRate

`func (o *DfsS3BucketStat) HasMetaCacheHitRate() bool`

HasMetaCacheHitRate returns a boolean if a field has been set.

### GetMetaDeleteLatencyUs

`func (o *DfsS3BucketStat) GetMetaDeleteLatencyUs() int64`

GetMetaDeleteLatencyUs returns the MetaDeleteLatencyUs field if non-nil, zero value otherwise.

### GetMetaDeleteLatencyUsOk

`func (o *DfsS3BucketStat) GetMetaDeleteLatencyUsOk() (*int64, bool)`

GetMetaDeleteLatencyUsOk returns a tuple with the MetaDeleteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDeleteLatencyUs

`func (o *DfsS3BucketStat) SetMetaDeleteLatencyUs(v int64)`

SetMetaDeleteLatencyUs sets MetaDeleteLatencyUs field to given value.

### HasMetaDeleteLatencyUs

`func (o *DfsS3BucketStat) HasMetaDeleteLatencyUs() bool`

HasMetaDeleteLatencyUs returns a boolean if a field has been set.

### GetMetaDeleteOps

`func (o *DfsS3BucketStat) GetMetaDeleteOps() int64`

GetMetaDeleteOps returns the MetaDeleteOps field if non-nil, zero value otherwise.

### GetMetaDeleteOpsOk

`func (o *DfsS3BucketStat) GetMetaDeleteOpsOk() (*int64, bool)`

GetMetaDeleteOpsOk returns a tuple with the MetaDeleteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDeleteOps

`func (o *DfsS3BucketStat) SetMetaDeleteOps(v int64)`

SetMetaDeleteOps sets MetaDeleteOps field to given value.

### HasMetaDeleteOps

`func (o *DfsS3BucketStat) HasMetaDeleteOps() bool`

HasMetaDeleteOps returns a boolean if a field has been set.

### GetMetaListLatencyUs

`func (o *DfsS3BucketStat) GetMetaListLatencyUs() int64`

GetMetaListLatencyUs returns the MetaListLatencyUs field if non-nil, zero value otherwise.

### GetMetaListLatencyUsOk

`func (o *DfsS3BucketStat) GetMetaListLatencyUsOk() (*int64, bool)`

GetMetaListLatencyUsOk returns a tuple with the MetaListLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaListLatencyUs

`func (o *DfsS3BucketStat) SetMetaListLatencyUs(v int64)`

SetMetaListLatencyUs sets MetaListLatencyUs field to given value.

### HasMetaListLatencyUs

`func (o *DfsS3BucketStat) HasMetaListLatencyUs() bool`

HasMetaListLatencyUs returns a boolean if a field has been set.

### GetMetaListOps

`func (o *DfsS3BucketStat) GetMetaListOps() int64`

GetMetaListOps returns the MetaListOps field if non-nil, zero value otherwise.

### GetMetaListOpsOk

`func (o *DfsS3BucketStat) GetMetaListOpsOk() (*int64, bool)`

GetMetaListOpsOk returns a tuple with the MetaListOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaListOps

`func (o *DfsS3BucketStat) SetMetaListOps(v int64)`

SetMetaListOps sets MetaListOps field to given value.

### HasMetaListOps

`func (o *DfsS3BucketStat) HasMetaListOps() bool`

HasMetaListOps returns a boolean if a field has been set.

### GetMetaReadLatencyUs

`func (o *DfsS3BucketStat) GetMetaReadLatencyUs() int64`

GetMetaReadLatencyUs returns the MetaReadLatencyUs field if non-nil, zero value otherwise.

### GetMetaReadLatencyUsOk

`func (o *DfsS3BucketStat) GetMetaReadLatencyUsOk() (*int64, bool)`

GetMetaReadLatencyUsOk returns a tuple with the MetaReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaReadLatencyUs

`func (o *DfsS3BucketStat) SetMetaReadLatencyUs(v int64)`

SetMetaReadLatencyUs sets MetaReadLatencyUs field to given value.

### HasMetaReadLatencyUs

`func (o *DfsS3BucketStat) HasMetaReadLatencyUs() bool`

HasMetaReadLatencyUs returns a boolean if a field has been set.

### GetMetaReadOps

`func (o *DfsS3BucketStat) GetMetaReadOps() int64`

GetMetaReadOps returns the MetaReadOps field if non-nil, zero value otherwise.

### GetMetaReadOpsOk

`func (o *DfsS3BucketStat) GetMetaReadOpsOk() (*int64, bool)`

GetMetaReadOpsOk returns a tuple with the MetaReadOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaReadOps

`func (o *DfsS3BucketStat) SetMetaReadOps(v int64)`

SetMetaReadOps sets MetaReadOps field to given value.

### HasMetaReadOps

`func (o *DfsS3BucketStat) HasMetaReadOps() bool`

HasMetaReadOps returns a boolean if a field has been set.

### GetMetaWriteLatencyUs

`func (o *DfsS3BucketStat) GetMetaWriteLatencyUs() int64`

GetMetaWriteLatencyUs returns the MetaWriteLatencyUs field if non-nil, zero value otherwise.

### GetMetaWriteLatencyUsOk

`func (o *DfsS3BucketStat) GetMetaWriteLatencyUsOk() (*int64, bool)`

GetMetaWriteLatencyUsOk returns a tuple with the MetaWriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaWriteLatencyUs

`func (o *DfsS3BucketStat) SetMetaWriteLatencyUs(v int64)`

SetMetaWriteLatencyUs sets MetaWriteLatencyUs field to given value.

### HasMetaWriteLatencyUs

`func (o *DfsS3BucketStat) HasMetaWriteLatencyUs() bool`

HasMetaWriteLatencyUs returns a boolean if a field has been set.

### GetMetaWriteOps

`func (o *DfsS3BucketStat) GetMetaWriteOps() int64`

GetMetaWriteOps returns the MetaWriteOps field if non-nil, zero value otherwise.

### GetMetaWriteOpsOk

`func (o *DfsS3BucketStat) GetMetaWriteOpsOk() (*int64, bool)`

GetMetaWriteOpsOk returns a tuple with the MetaWriteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaWriteOps

`func (o *DfsS3BucketStat) SetMetaWriteOps(v int64)`

SetMetaWriteOps sets MetaWriteOps field to given value.

### HasMetaWriteOps

`func (o *DfsS3BucketStat) HasMetaWriteOps() bool`

HasMetaWriteOps returns a boolean if a field has been set.

### GetReadAheadBandwidthKbyte

`func (o *DfsS3BucketStat) GetReadAheadBandwidthKbyte() int64`

GetReadAheadBandwidthKbyte returns the ReadAheadBandwidthKbyte field if non-nil, zero value otherwise.

### GetReadAheadBandwidthKbyteOk

`func (o *DfsS3BucketStat) GetReadAheadBandwidthKbyteOk() (*int64, bool)`

GetReadAheadBandwidthKbyteOk returns a tuple with the ReadAheadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadBandwidthKbyte

`func (o *DfsS3BucketStat) SetReadAheadBandwidthKbyte(v int64)`

SetReadAheadBandwidthKbyte sets ReadAheadBandwidthKbyte field to given value.

### HasReadAheadBandwidthKbyte

`func (o *DfsS3BucketStat) HasReadAheadBandwidthKbyte() bool`

HasReadAheadBandwidthKbyte returns a boolean if a field has been set.

### GetReadAheadIops

`func (o *DfsS3BucketStat) GetReadAheadIops() int64`

GetReadAheadIops returns the ReadAheadIops field if non-nil, zero value otherwise.

### GetReadAheadIopsOk

`func (o *DfsS3BucketStat) GetReadAheadIopsOk() (*int64, bool)`

GetReadAheadIopsOk returns a tuple with the ReadAheadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadIops

`func (o *DfsS3BucketStat) SetReadAheadIops(v int64)`

SetReadAheadIops sets ReadAheadIops field to given value.

### HasReadAheadIops

`func (o *DfsS3BucketStat) HasReadAheadIops() bool`

HasReadAheadIops returns a boolean if a field has been set.

### GetAllocatedSize

`func (o *DfsS3BucketStat) GetAllocatedSize() int64`

GetAllocatedSize returns the AllocatedSize field if non-nil, zero value otherwise.

### GetAllocatedSizeOk

`func (o *DfsS3BucketStat) GetAllocatedSizeOk() (*int64, bool)`

GetAllocatedSizeOk returns a tuple with the AllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSize

`func (o *DfsS3BucketStat) SetAllocatedSize(v int64)`

SetAllocatedSize sets AllocatedSize field to given value.

### HasAllocatedSize

`func (o *DfsS3BucketStat) HasAllocatedSize() bool`

HasAllocatedSize returns a boolean if a field has been set.

### GetObjects

`func (o *DfsS3BucketStat) GetObjects() int64`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *DfsS3BucketStat) GetObjectsOk() (*int64, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *DfsS3BucketStat) SetObjects(v int64)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *DfsS3BucketStat) HasObjects() bool`

HasObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


