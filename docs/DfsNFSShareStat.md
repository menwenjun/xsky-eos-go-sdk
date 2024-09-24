# DfsNFSShareStat

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

### NewDfsNFSShareStat

`func NewDfsNFSShareStat() *DfsNFSShareStat`

NewDfsNFSShareStat instantiates a new DfsNFSShareStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsNFSShareStatWithDefaults

`func NewDfsNFSShareStatWithDefaults() *DfsNFSShareStat`

NewDfsNFSShareStatWithDefaults instantiates a new DfsNFSShareStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *DfsNFSShareStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsNFSShareStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsNFSShareStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsNFSShareStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataCacheHitRate

`func (o *DfsNFSShareStat) GetDataCacheHitRate() float64`

GetDataCacheHitRate returns the DataCacheHitRate field if non-nil, zero value otherwise.

### GetDataCacheHitRateOk

`func (o *DfsNFSShareStat) GetDataCacheHitRateOk() (*float64, bool)`

GetDataCacheHitRateOk returns a tuple with the DataCacheHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCacheHitRate

`func (o *DfsNFSShareStat) SetDataCacheHitRate(v float64)`

SetDataCacheHitRate sets DataCacheHitRate field to given value.

### HasDataCacheHitRate

`func (o *DfsNFSShareStat) HasDataCacheHitRate() bool`

HasDataCacheHitRate returns a boolean if a field has been set.

### GetDataDeleteBandwidthKbyte

`func (o *DfsNFSShareStat) GetDataDeleteBandwidthKbyte() int64`

GetDataDeleteBandwidthKbyte returns the DataDeleteBandwidthKbyte field if non-nil, zero value otherwise.

### GetDataDeleteBandwidthKbyteOk

`func (o *DfsNFSShareStat) GetDataDeleteBandwidthKbyteOk() (*int64, bool)`

GetDataDeleteBandwidthKbyteOk returns a tuple with the DataDeleteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeleteBandwidthKbyte

`func (o *DfsNFSShareStat) SetDataDeleteBandwidthKbyte(v int64)`

SetDataDeleteBandwidthKbyte sets DataDeleteBandwidthKbyte field to given value.

### HasDataDeleteBandwidthKbyte

`func (o *DfsNFSShareStat) HasDataDeleteBandwidthKbyte() bool`

HasDataDeleteBandwidthKbyte returns a boolean if a field has been set.

### GetDataDeleteIops

`func (o *DfsNFSShareStat) GetDataDeleteIops() int64`

GetDataDeleteIops returns the DataDeleteIops field if non-nil, zero value otherwise.

### GetDataDeleteIopsOk

`func (o *DfsNFSShareStat) GetDataDeleteIopsOk() (*int64, bool)`

GetDataDeleteIopsOk returns a tuple with the DataDeleteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeleteIops

`func (o *DfsNFSShareStat) SetDataDeleteIops(v int64)`

SetDataDeleteIops sets DataDeleteIops field to given value.

### HasDataDeleteIops

`func (o *DfsNFSShareStat) HasDataDeleteIops() bool`

HasDataDeleteIops returns a boolean if a field has been set.

### GetDataDeleteLatencyUs

`func (o *DfsNFSShareStat) GetDataDeleteLatencyUs() int64`

GetDataDeleteLatencyUs returns the DataDeleteLatencyUs field if non-nil, zero value otherwise.

### GetDataDeleteLatencyUsOk

`func (o *DfsNFSShareStat) GetDataDeleteLatencyUsOk() (*int64, bool)`

GetDataDeleteLatencyUsOk returns a tuple with the DataDeleteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeleteLatencyUs

`func (o *DfsNFSShareStat) SetDataDeleteLatencyUs(v int64)`

SetDataDeleteLatencyUs sets DataDeleteLatencyUs field to given value.

### HasDataDeleteLatencyUs

`func (o *DfsNFSShareStat) HasDataDeleteLatencyUs() bool`

HasDataDeleteLatencyUs returns a boolean if a field has been set.

### GetDataReadBandwidthKbyte

`func (o *DfsNFSShareStat) GetDataReadBandwidthKbyte() int64`

GetDataReadBandwidthKbyte returns the DataReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetDataReadBandwidthKbyteOk

`func (o *DfsNFSShareStat) GetDataReadBandwidthKbyteOk() (*int64, bool)`

GetDataReadBandwidthKbyteOk returns a tuple with the DataReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadBandwidthKbyte

`func (o *DfsNFSShareStat) SetDataReadBandwidthKbyte(v int64)`

SetDataReadBandwidthKbyte sets DataReadBandwidthKbyte field to given value.

### HasDataReadBandwidthKbyte

`func (o *DfsNFSShareStat) HasDataReadBandwidthKbyte() bool`

HasDataReadBandwidthKbyte returns a boolean if a field has been set.

### GetDataReadIops

`func (o *DfsNFSShareStat) GetDataReadIops() int64`

GetDataReadIops returns the DataReadIops field if non-nil, zero value otherwise.

### GetDataReadIopsOk

`func (o *DfsNFSShareStat) GetDataReadIopsOk() (*int64, bool)`

GetDataReadIopsOk returns a tuple with the DataReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadIops

`func (o *DfsNFSShareStat) SetDataReadIops(v int64)`

SetDataReadIops sets DataReadIops field to given value.

### HasDataReadIops

`func (o *DfsNFSShareStat) HasDataReadIops() bool`

HasDataReadIops returns a boolean if a field has been set.

### GetDataReadLatencyUs

`func (o *DfsNFSShareStat) GetDataReadLatencyUs() int64`

GetDataReadLatencyUs returns the DataReadLatencyUs field if non-nil, zero value otherwise.

### GetDataReadLatencyUsOk

`func (o *DfsNFSShareStat) GetDataReadLatencyUsOk() (*int64, bool)`

GetDataReadLatencyUsOk returns a tuple with the DataReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadLatencyUs

`func (o *DfsNFSShareStat) SetDataReadLatencyUs(v int64)`

SetDataReadLatencyUs sets DataReadLatencyUs field to given value.

### HasDataReadLatencyUs

`func (o *DfsNFSShareStat) HasDataReadLatencyUs() bool`

HasDataReadLatencyUs returns a boolean if a field has been set.

### GetDataWriteBandwidthKbyte

`func (o *DfsNFSShareStat) GetDataWriteBandwidthKbyte() int64`

GetDataWriteBandwidthKbyte returns the DataWriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetDataWriteBandwidthKbyteOk

`func (o *DfsNFSShareStat) GetDataWriteBandwidthKbyteOk() (*int64, bool)`

GetDataWriteBandwidthKbyteOk returns a tuple with the DataWriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWriteBandwidthKbyte

`func (o *DfsNFSShareStat) SetDataWriteBandwidthKbyte(v int64)`

SetDataWriteBandwidthKbyte sets DataWriteBandwidthKbyte field to given value.

### HasDataWriteBandwidthKbyte

`func (o *DfsNFSShareStat) HasDataWriteBandwidthKbyte() bool`

HasDataWriteBandwidthKbyte returns a boolean if a field has been set.

### GetDataWriteIops

`func (o *DfsNFSShareStat) GetDataWriteIops() int64`

GetDataWriteIops returns the DataWriteIops field if non-nil, zero value otherwise.

### GetDataWriteIopsOk

`func (o *DfsNFSShareStat) GetDataWriteIopsOk() (*int64, bool)`

GetDataWriteIopsOk returns a tuple with the DataWriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWriteIops

`func (o *DfsNFSShareStat) SetDataWriteIops(v int64)`

SetDataWriteIops sets DataWriteIops field to given value.

### HasDataWriteIops

`func (o *DfsNFSShareStat) HasDataWriteIops() bool`

HasDataWriteIops returns a boolean if a field has been set.

### GetDataWriteLatencyUs

`func (o *DfsNFSShareStat) GetDataWriteLatencyUs() int64`

GetDataWriteLatencyUs returns the DataWriteLatencyUs field if non-nil, zero value otherwise.

### GetDataWriteLatencyUsOk

`func (o *DfsNFSShareStat) GetDataWriteLatencyUsOk() (*int64, bool)`

GetDataWriteLatencyUsOk returns a tuple with the DataWriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWriteLatencyUs

`func (o *DfsNFSShareStat) SetDataWriteLatencyUs(v int64)`

SetDataWriteLatencyUs sets DataWriteLatencyUs field to given value.

### HasDataWriteLatencyUs

`func (o *DfsNFSShareStat) HasDataWriteLatencyUs() bool`

HasDataWriteLatencyUs returns a boolean if a field has been set.

### GetMetaCacheHitRate

`func (o *DfsNFSShareStat) GetMetaCacheHitRate() float64`

GetMetaCacheHitRate returns the MetaCacheHitRate field if non-nil, zero value otherwise.

### GetMetaCacheHitRateOk

`func (o *DfsNFSShareStat) GetMetaCacheHitRateOk() (*float64, bool)`

GetMetaCacheHitRateOk returns a tuple with the MetaCacheHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaCacheHitRate

`func (o *DfsNFSShareStat) SetMetaCacheHitRate(v float64)`

SetMetaCacheHitRate sets MetaCacheHitRate field to given value.

### HasMetaCacheHitRate

`func (o *DfsNFSShareStat) HasMetaCacheHitRate() bool`

HasMetaCacheHitRate returns a boolean if a field has been set.

### GetMetaDeleteLatencyUs

`func (o *DfsNFSShareStat) GetMetaDeleteLatencyUs() int64`

GetMetaDeleteLatencyUs returns the MetaDeleteLatencyUs field if non-nil, zero value otherwise.

### GetMetaDeleteLatencyUsOk

`func (o *DfsNFSShareStat) GetMetaDeleteLatencyUsOk() (*int64, bool)`

GetMetaDeleteLatencyUsOk returns a tuple with the MetaDeleteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDeleteLatencyUs

`func (o *DfsNFSShareStat) SetMetaDeleteLatencyUs(v int64)`

SetMetaDeleteLatencyUs sets MetaDeleteLatencyUs field to given value.

### HasMetaDeleteLatencyUs

`func (o *DfsNFSShareStat) HasMetaDeleteLatencyUs() bool`

HasMetaDeleteLatencyUs returns a boolean if a field has been set.

### GetMetaDeleteOps

`func (o *DfsNFSShareStat) GetMetaDeleteOps() int64`

GetMetaDeleteOps returns the MetaDeleteOps field if non-nil, zero value otherwise.

### GetMetaDeleteOpsOk

`func (o *DfsNFSShareStat) GetMetaDeleteOpsOk() (*int64, bool)`

GetMetaDeleteOpsOk returns a tuple with the MetaDeleteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDeleteOps

`func (o *DfsNFSShareStat) SetMetaDeleteOps(v int64)`

SetMetaDeleteOps sets MetaDeleteOps field to given value.

### HasMetaDeleteOps

`func (o *DfsNFSShareStat) HasMetaDeleteOps() bool`

HasMetaDeleteOps returns a boolean if a field has been set.

### GetMetaListLatencyUs

`func (o *DfsNFSShareStat) GetMetaListLatencyUs() int64`

GetMetaListLatencyUs returns the MetaListLatencyUs field if non-nil, zero value otherwise.

### GetMetaListLatencyUsOk

`func (o *DfsNFSShareStat) GetMetaListLatencyUsOk() (*int64, bool)`

GetMetaListLatencyUsOk returns a tuple with the MetaListLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaListLatencyUs

`func (o *DfsNFSShareStat) SetMetaListLatencyUs(v int64)`

SetMetaListLatencyUs sets MetaListLatencyUs field to given value.

### HasMetaListLatencyUs

`func (o *DfsNFSShareStat) HasMetaListLatencyUs() bool`

HasMetaListLatencyUs returns a boolean if a field has been set.

### GetMetaListOps

`func (o *DfsNFSShareStat) GetMetaListOps() int64`

GetMetaListOps returns the MetaListOps field if non-nil, zero value otherwise.

### GetMetaListOpsOk

`func (o *DfsNFSShareStat) GetMetaListOpsOk() (*int64, bool)`

GetMetaListOpsOk returns a tuple with the MetaListOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaListOps

`func (o *DfsNFSShareStat) SetMetaListOps(v int64)`

SetMetaListOps sets MetaListOps field to given value.

### HasMetaListOps

`func (o *DfsNFSShareStat) HasMetaListOps() bool`

HasMetaListOps returns a boolean if a field has been set.

### GetMetaReadLatencyUs

`func (o *DfsNFSShareStat) GetMetaReadLatencyUs() int64`

GetMetaReadLatencyUs returns the MetaReadLatencyUs field if non-nil, zero value otherwise.

### GetMetaReadLatencyUsOk

`func (o *DfsNFSShareStat) GetMetaReadLatencyUsOk() (*int64, bool)`

GetMetaReadLatencyUsOk returns a tuple with the MetaReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaReadLatencyUs

`func (o *DfsNFSShareStat) SetMetaReadLatencyUs(v int64)`

SetMetaReadLatencyUs sets MetaReadLatencyUs field to given value.

### HasMetaReadLatencyUs

`func (o *DfsNFSShareStat) HasMetaReadLatencyUs() bool`

HasMetaReadLatencyUs returns a boolean if a field has been set.

### GetMetaReadOps

`func (o *DfsNFSShareStat) GetMetaReadOps() int64`

GetMetaReadOps returns the MetaReadOps field if non-nil, zero value otherwise.

### GetMetaReadOpsOk

`func (o *DfsNFSShareStat) GetMetaReadOpsOk() (*int64, bool)`

GetMetaReadOpsOk returns a tuple with the MetaReadOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaReadOps

`func (o *DfsNFSShareStat) SetMetaReadOps(v int64)`

SetMetaReadOps sets MetaReadOps field to given value.

### HasMetaReadOps

`func (o *DfsNFSShareStat) HasMetaReadOps() bool`

HasMetaReadOps returns a boolean if a field has been set.

### GetMetaWriteLatencyUs

`func (o *DfsNFSShareStat) GetMetaWriteLatencyUs() int64`

GetMetaWriteLatencyUs returns the MetaWriteLatencyUs field if non-nil, zero value otherwise.

### GetMetaWriteLatencyUsOk

`func (o *DfsNFSShareStat) GetMetaWriteLatencyUsOk() (*int64, bool)`

GetMetaWriteLatencyUsOk returns a tuple with the MetaWriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaWriteLatencyUs

`func (o *DfsNFSShareStat) SetMetaWriteLatencyUs(v int64)`

SetMetaWriteLatencyUs sets MetaWriteLatencyUs field to given value.

### HasMetaWriteLatencyUs

`func (o *DfsNFSShareStat) HasMetaWriteLatencyUs() bool`

HasMetaWriteLatencyUs returns a boolean if a field has been set.

### GetMetaWriteOps

`func (o *DfsNFSShareStat) GetMetaWriteOps() int64`

GetMetaWriteOps returns the MetaWriteOps field if non-nil, zero value otherwise.

### GetMetaWriteOpsOk

`func (o *DfsNFSShareStat) GetMetaWriteOpsOk() (*int64, bool)`

GetMetaWriteOpsOk returns a tuple with the MetaWriteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaWriteOps

`func (o *DfsNFSShareStat) SetMetaWriteOps(v int64)`

SetMetaWriteOps sets MetaWriteOps field to given value.

### HasMetaWriteOps

`func (o *DfsNFSShareStat) HasMetaWriteOps() bool`

HasMetaWriteOps returns a boolean if a field has been set.

### GetReadAheadBandwidthKbyte

`func (o *DfsNFSShareStat) GetReadAheadBandwidthKbyte() int64`

GetReadAheadBandwidthKbyte returns the ReadAheadBandwidthKbyte field if non-nil, zero value otherwise.

### GetReadAheadBandwidthKbyteOk

`func (o *DfsNFSShareStat) GetReadAheadBandwidthKbyteOk() (*int64, bool)`

GetReadAheadBandwidthKbyteOk returns a tuple with the ReadAheadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadBandwidthKbyte

`func (o *DfsNFSShareStat) SetReadAheadBandwidthKbyte(v int64)`

SetReadAheadBandwidthKbyte sets ReadAheadBandwidthKbyte field to given value.

### HasReadAheadBandwidthKbyte

`func (o *DfsNFSShareStat) HasReadAheadBandwidthKbyte() bool`

HasReadAheadBandwidthKbyte returns a boolean if a field has been set.

### GetReadAheadIops

`func (o *DfsNFSShareStat) GetReadAheadIops() int64`

GetReadAheadIops returns the ReadAheadIops field if non-nil, zero value otherwise.

### GetReadAheadIopsOk

`func (o *DfsNFSShareStat) GetReadAheadIopsOk() (*int64, bool)`

GetReadAheadIopsOk returns a tuple with the ReadAheadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadIops

`func (o *DfsNFSShareStat) SetReadAheadIops(v int64)`

SetReadAheadIops sets ReadAheadIops field to given value.

### HasReadAheadIops

`func (o *DfsNFSShareStat) HasReadAheadIops() bool`

HasReadAheadIops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


