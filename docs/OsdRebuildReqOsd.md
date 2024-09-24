# OsdRebuildReqOsd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewDiskId** | Pointer to **int64** | new data disk id | [optional] 
**NewMinAllocSize** | Pointer to **int64** | see OsdCreateReq.Osd.MinAllocSize | [optional] 
**NewOmapByte** | Pointer to **int64** | new size of omap partition | [optional] 
**NewPartitionId** | Pointer to **int64** | new cache partition id | [optional] 
**NewReadCacheSize** | Pointer to **int64** | new read cache size in bytes | [optional] 

## Methods

### NewOsdRebuildReqOsd

`func NewOsdRebuildReqOsd() *OsdRebuildReqOsd`

NewOsdRebuildReqOsd instantiates a new OsdRebuildReqOsd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsdRebuildReqOsdWithDefaults

`func NewOsdRebuildReqOsdWithDefaults() *OsdRebuildReqOsd`

NewOsdRebuildReqOsdWithDefaults instantiates a new OsdRebuildReqOsd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewDiskId

`func (o *OsdRebuildReqOsd) GetNewDiskId() int64`

GetNewDiskId returns the NewDiskId field if non-nil, zero value otherwise.

### GetNewDiskIdOk

`func (o *OsdRebuildReqOsd) GetNewDiskIdOk() (*int64, bool)`

GetNewDiskIdOk returns a tuple with the NewDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDiskId

`func (o *OsdRebuildReqOsd) SetNewDiskId(v int64)`

SetNewDiskId sets NewDiskId field to given value.

### HasNewDiskId

`func (o *OsdRebuildReqOsd) HasNewDiskId() bool`

HasNewDiskId returns a boolean if a field has been set.

### GetNewMinAllocSize

`func (o *OsdRebuildReqOsd) GetNewMinAllocSize() int64`

GetNewMinAllocSize returns the NewMinAllocSize field if non-nil, zero value otherwise.

### GetNewMinAllocSizeOk

`func (o *OsdRebuildReqOsd) GetNewMinAllocSizeOk() (*int64, bool)`

GetNewMinAllocSizeOk returns a tuple with the NewMinAllocSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMinAllocSize

`func (o *OsdRebuildReqOsd) SetNewMinAllocSize(v int64)`

SetNewMinAllocSize sets NewMinAllocSize field to given value.

### HasNewMinAllocSize

`func (o *OsdRebuildReqOsd) HasNewMinAllocSize() bool`

HasNewMinAllocSize returns a boolean if a field has been set.

### GetNewOmapByte

`func (o *OsdRebuildReqOsd) GetNewOmapByte() int64`

GetNewOmapByte returns the NewOmapByte field if non-nil, zero value otherwise.

### GetNewOmapByteOk

`func (o *OsdRebuildReqOsd) GetNewOmapByteOk() (*int64, bool)`

GetNewOmapByteOk returns a tuple with the NewOmapByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOmapByte

`func (o *OsdRebuildReqOsd) SetNewOmapByte(v int64)`

SetNewOmapByte sets NewOmapByte field to given value.

### HasNewOmapByte

`func (o *OsdRebuildReqOsd) HasNewOmapByte() bool`

HasNewOmapByte returns a boolean if a field has been set.

### GetNewPartitionId

`func (o *OsdRebuildReqOsd) GetNewPartitionId() int64`

GetNewPartitionId returns the NewPartitionId field if non-nil, zero value otherwise.

### GetNewPartitionIdOk

`func (o *OsdRebuildReqOsd) GetNewPartitionIdOk() (*int64, bool)`

GetNewPartitionIdOk returns a tuple with the NewPartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPartitionId

`func (o *OsdRebuildReqOsd) SetNewPartitionId(v int64)`

SetNewPartitionId sets NewPartitionId field to given value.

### HasNewPartitionId

`func (o *OsdRebuildReqOsd) HasNewPartitionId() bool`

HasNewPartitionId returns a boolean if a field has been set.

### GetNewReadCacheSize

`func (o *OsdRebuildReqOsd) GetNewReadCacheSize() int64`

GetNewReadCacheSize returns the NewReadCacheSize field if non-nil, zero value otherwise.

### GetNewReadCacheSizeOk

`func (o *OsdRebuildReqOsd) GetNewReadCacheSizeOk() (*int64, bool)`

GetNewReadCacheSizeOk returns a tuple with the NewReadCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewReadCacheSize

`func (o *OsdRebuildReqOsd) SetNewReadCacheSize(v int64)`

SetNewReadCacheSize sets NewReadCacheSize field to given value.

### HasNewReadCacheSize

`func (o *OsdRebuildReqOsd) HasNewReadCacheSize() bool`

HasNewReadCacheSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


