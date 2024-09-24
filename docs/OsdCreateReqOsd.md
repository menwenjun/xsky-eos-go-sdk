# OsdCreateReqOsd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindPoolId** | Pointer to **int64** | bind pool id | [optional] 
**DiskId** | Pointer to **int64** | data disk id | [optional] 
**MinAllocSize** | Pointer to **int64** | for xstore_min_alloc_size_hdd/xstore_min_alloc_size_ssd | [optional] 
**OmapByte** | Pointer to **int64** | size of omap partition | [optional] 
**PartitionId** | Pointer to **int64** | cache partition id | [optional] 
**ReadCacheSize** | Pointer to **int64** | read cache size in bytes | [optional] 
**Role** | Pointer to **string** | osd role: \&quot;data\&quot; or \&quot;index\&quot;, default is \&quot;data\&quot; | [optional] 

## Methods

### NewOsdCreateReqOsd

`func NewOsdCreateReqOsd() *OsdCreateReqOsd`

NewOsdCreateReqOsd instantiates a new OsdCreateReqOsd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsdCreateReqOsdWithDefaults

`func NewOsdCreateReqOsdWithDefaults() *OsdCreateReqOsd`

NewOsdCreateReqOsdWithDefaults instantiates a new OsdCreateReqOsd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindPoolId

`func (o *OsdCreateReqOsd) GetBindPoolId() int64`

GetBindPoolId returns the BindPoolId field if non-nil, zero value otherwise.

### GetBindPoolIdOk

`func (o *OsdCreateReqOsd) GetBindPoolIdOk() (*int64, bool)`

GetBindPoolIdOk returns a tuple with the BindPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPoolId

`func (o *OsdCreateReqOsd) SetBindPoolId(v int64)`

SetBindPoolId sets BindPoolId field to given value.

### HasBindPoolId

`func (o *OsdCreateReqOsd) HasBindPoolId() bool`

HasBindPoolId returns a boolean if a field has been set.

### GetDiskId

`func (o *OsdCreateReqOsd) GetDiskId() int64`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *OsdCreateReqOsd) GetDiskIdOk() (*int64, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *OsdCreateReqOsd) SetDiskId(v int64)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *OsdCreateReqOsd) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.

### GetMinAllocSize

`func (o *OsdCreateReqOsd) GetMinAllocSize() int64`

GetMinAllocSize returns the MinAllocSize field if non-nil, zero value otherwise.

### GetMinAllocSizeOk

`func (o *OsdCreateReqOsd) GetMinAllocSizeOk() (*int64, bool)`

GetMinAllocSizeOk returns a tuple with the MinAllocSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAllocSize

`func (o *OsdCreateReqOsd) SetMinAllocSize(v int64)`

SetMinAllocSize sets MinAllocSize field to given value.

### HasMinAllocSize

`func (o *OsdCreateReqOsd) HasMinAllocSize() bool`

HasMinAllocSize returns a boolean if a field has been set.

### GetOmapByte

`func (o *OsdCreateReqOsd) GetOmapByte() int64`

GetOmapByte returns the OmapByte field if non-nil, zero value otherwise.

### GetOmapByteOk

`func (o *OsdCreateReqOsd) GetOmapByteOk() (*int64, bool)`

GetOmapByteOk returns a tuple with the OmapByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmapByte

`func (o *OsdCreateReqOsd) SetOmapByte(v int64)`

SetOmapByte sets OmapByte field to given value.

### HasOmapByte

`func (o *OsdCreateReqOsd) HasOmapByte() bool`

HasOmapByte returns a boolean if a field has been set.

### GetPartitionId

`func (o *OsdCreateReqOsd) GetPartitionId() int64`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *OsdCreateReqOsd) GetPartitionIdOk() (*int64, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *OsdCreateReqOsd) SetPartitionId(v int64)`

SetPartitionId sets PartitionId field to given value.

### HasPartitionId

`func (o *OsdCreateReqOsd) HasPartitionId() bool`

HasPartitionId returns a boolean if a field has been set.

### GetReadCacheSize

`func (o *OsdCreateReqOsd) GetReadCacheSize() int64`

GetReadCacheSize returns the ReadCacheSize field if non-nil, zero value otherwise.

### GetReadCacheSizeOk

`func (o *OsdCreateReqOsd) GetReadCacheSizeOk() (*int64, bool)`

GetReadCacheSizeOk returns a tuple with the ReadCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCacheSize

`func (o *OsdCreateReqOsd) SetReadCacheSize(v int64)`

SetReadCacheSize sets ReadCacheSize field to given value.

### HasReadCacheSize

`func (o *OsdCreateReqOsd) HasReadCacheSize() bool`

HasReadCacheSize returns a boolean if a field has been set.

### GetRole

`func (o *OsdCreateReqOsd) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OsdCreateReqOsd) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OsdCreateReqOsd) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OsdCreateReqOsd) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


