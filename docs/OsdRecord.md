# OsdRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**BindOsdGroup** | Pointer to [**OsdGroupNestview**](OsdGroupNestview.md) |  | [optional] 
**BindPool** | Pointer to [**PoolNestview**](PoolNestview.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DataDir** | Pointer to **string** |  | [optional] 
**Disk** | Pointer to [**Disk**](Disk.md) |  | [optional] 
**EncryptEnabled** | Pointer to **bool** |  | [optional] 
**ExitCount** | Pointer to **int64** |  | [optional] 
**ExitTime** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**In** | Pointer to **bool** |  | [optional] 
**InitTime** | Pointer to **time.Time** |  | [optional] 
**LastScrubTime** | Pointer to **time.Time** |  | [optional] 
**MetaBytes** | Pointer to **int64** |  | [optional] 
**MinAllocSize** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumaNode** | Pointer to **int64** |  | [optional] 
**OmapByte** | Pointer to **int64** |  | [optional] 
**OsdGroup** | Pointer to [**OsdGroupNestview**](OsdGroupNestview.md) |  | [optional] 
**OsdId** | Pointer to **int64** |  | [optional] 
**OspMetadataCluster** | Pointer to [**OspMetadataCluster**](OspMetadataCluster.md) |  | [optional] 
**Partition** | Pointer to [**Partition**](Partition.md) |  | [optional] 
**Pool** | Pointer to [**PoolNestview**](PoolNestview.md) |  | [optional] 
**ReadCacheSize** | Pointer to **int64** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusClass** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Samples** | Pointer to [**[]OsdStat**](OsdStat.md) |  | [optional] 

## Methods

### NewOsdRecord

`func NewOsdRecord() *OsdRecord`

NewOsdRecord instantiates a new OsdRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsdRecordWithDefaults

`func NewOsdRecordWithDefaults() *OsdRecord`

NewOsdRecordWithDefaults instantiates a new OsdRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *OsdRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *OsdRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *OsdRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *OsdRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetBindOsdGroup

`func (o *OsdRecord) GetBindOsdGroup() OsdGroupNestview`

GetBindOsdGroup returns the BindOsdGroup field if non-nil, zero value otherwise.

### GetBindOsdGroupOk

`func (o *OsdRecord) GetBindOsdGroupOk() (*OsdGroupNestview, bool)`

GetBindOsdGroupOk returns a tuple with the BindOsdGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindOsdGroup

`func (o *OsdRecord) SetBindOsdGroup(v OsdGroupNestview)`

SetBindOsdGroup sets BindOsdGroup field to given value.

### HasBindOsdGroup

`func (o *OsdRecord) HasBindOsdGroup() bool`

HasBindOsdGroup returns a boolean if a field has been set.

### GetBindPool

`func (o *OsdRecord) GetBindPool() PoolNestview`

GetBindPool returns the BindPool field if non-nil, zero value otherwise.

### GetBindPoolOk

`func (o *OsdRecord) GetBindPoolOk() (*PoolNestview, bool)`

GetBindPoolOk returns a tuple with the BindPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPool

`func (o *OsdRecord) SetBindPool(v PoolNestview)`

SetBindPool sets BindPool field to given value.

### HasBindPool

`func (o *OsdRecord) HasBindPool() bool`

HasBindPool returns a boolean if a field has been set.

### GetCluster

`func (o *OsdRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OsdRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OsdRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OsdRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OsdRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OsdRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OsdRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OsdRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataDir

`func (o *OsdRecord) GetDataDir() string`

GetDataDir returns the DataDir field if non-nil, zero value otherwise.

### GetDataDirOk

`func (o *OsdRecord) GetDataDirOk() (*string, bool)`

GetDataDirOk returns a tuple with the DataDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDir

`func (o *OsdRecord) SetDataDir(v string)`

SetDataDir sets DataDir field to given value.

### HasDataDir

`func (o *OsdRecord) HasDataDir() bool`

HasDataDir returns a boolean if a field has been set.

### GetDisk

`func (o *OsdRecord) GetDisk() Disk`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *OsdRecord) GetDiskOk() (*Disk, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *OsdRecord) SetDisk(v Disk)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *OsdRecord) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetEncryptEnabled

`func (o *OsdRecord) GetEncryptEnabled() bool`

GetEncryptEnabled returns the EncryptEnabled field if non-nil, zero value otherwise.

### GetEncryptEnabledOk

`func (o *OsdRecord) GetEncryptEnabledOk() (*bool, bool)`

GetEncryptEnabledOk returns a tuple with the EncryptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptEnabled

`func (o *OsdRecord) SetEncryptEnabled(v bool)`

SetEncryptEnabled sets EncryptEnabled field to given value.

### HasEncryptEnabled

`func (o *OsdRecord) HasEncryptEnabled() bool`

HasEncryptEnabled returns a boolean if a field has been set.

### GetExitCount

`func (o *OsdRecord) GetExitCount() int64`

GetExitCount returns the ExitCount field if non-nil, zero value otherwise.

### GetExitCountOk

`func (o *OsdRecord) GetExitCountOk() (*int64, bool)`

GetExitCountOk returns a tuple with the ExitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCount

`func (o *OsdRecord) SetExitCount(v int64)`

SetExitCount sets ExitCount field to given value.

### HasExitCount

`func (o *OsdRecord) HasExitCount() bool`

HasExitCount returns a boolean if a field has been set.

### GetExitTime

`func (o *OsdRecord) GetExitTime() time.Time`

GetExitTime returns the ExitTime field if non-nil, zero value otherwise.

### GetExitTimeOk

`func (o *OsdRecord) GetExitTimeOk() (*time.Time, bool)`

GetExitTimeOk returns a tuple with the ExitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitTime

`func (o *OsdRecord) SetExitTime(v time.Time)`

SetExitTime sets ExitTime field to given value.

### HasExitTime

`func (o *OsdRecord) HasExitTime() bool`

HasExitTime returns a boolean if a field has been set.

### GetHost

`func (o *OsdRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *OsdRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *OsdRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *OsdRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *OsdRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OsdRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OsdRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OsdRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIn

`func (o *OsdRecord) GetIn() bool`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *OsdRecord) GetInOk() (*bool, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *OsdRecord) SetIn(v bool)`

SetIn sets In field to given value.

### HasIn

`func (o *OsdRecord) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetInitTime

`func (o *OsdRecord) GetInitTime() time.Time`

GetInitTime returns the InitTime field if non-nil, zero value otherwise.

### GetInitTimeOk

`func (o *OsdRecord) GetInitTimeOk() (*time.Time, bool)`

GetInitTimeOk returns a tuple with the InitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitTime

`func (o *OsdRecord) SetInitTime(v time.Time)`

SetInitTime sets InitTime field to given value.

### HasInitTime

`func (o *OsdRecord) HasInitTime() bool`

HasInitTime returns a boolean if a field has been set.

### GetLastScrubTime

`func (o *OsdRecord) GetLastScrubTime() time.Time`

GetLastScrubTime returns the LastScrubTime field if non-nil, zero value otherwise.

### GetLastScrubTimeOk

`func (o *OsdRecord) GetLastScrubTimeOk() (*time.Time, bool)`

GetLastScrubTimeOk returns a tuple with the LastScrubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScrubTime

`func (o *OsdRecord) SetLastScrubTime(v time.Time)`

SetLastScrubTime sets LastScrubTime field to given value.

### HasLastScrubTime

`func (o *OsdRecord) HasLastScrubTime() bool`

HasLastScrubTime returns a boolean if a field has been set.

### GetMetaBytes

`func (o *OsdRecord) GetMetaBytes() int64`

GetMetaBytes returns the MetaBytes field if non-nil, zero value otherwise.

### GetMetaBytesOk

`func (o *OsdRecord) GetMetaBytesOk() (*int64, bool)`

GetMetaBytesOk returns a tuple with the MetaBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaBytes

`func (o *OsdRecord) SetMetaBytes(v int64)`

SetMetaBytes sets MetaBytes field to given value.

### HasMetaBytes

`func (o *OsdRecord) HasMetaBytes() bool`

HasMetaBytes returns a boolean if a field has been set.

### GetMinAllocSize

`func (o *OsdRecord) GetMinAllocSize() int64`

GetMinAllocSize returns the MinAllocSize field if non-nil, zero value otherwise.

### GetMinAllocSizeOk

`func (o *OsdRecord) GetMinAllocSizeOk() (*int64, bool)`

GetMinAllocSizeOk returns a tuple with the MinAllocSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAllocSize

`func (o *OsdRecord) SetMinAllocSize(v int64)`

SetMinAllocSize sets MinAllocSize field to given value.

### HasMinAllocSize

`func (o *OsdRecord) HasMinAllocSize() bool`

HasMinAllocSize returns a boolean if a field has been set.

### GetName

`func (o *OsdRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsdRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsdRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsdRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumaNode

`func (o *OsdRecord) GetNumaNode() int64`

GetNumaNode returns the NumaNode field if non-nil, zero value otherwise.

### GetNumaNodeOk

`func (o *OsdRecord) GetNumaNodeOk() (*int64, bool)`

GetNumaNodeOk returns a tuple with the NumaNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaNode

`func (o *OsdRecord) SetNumaNode(v int64)`

SetNumaNode sets NumaNode field to given value.

### HasNumaNode

`func (o *OsdRecord) HasNumaNode() bool`

HasNumaNode returns a boolean if a field has been set.

### GetOmapByte

`func (o *OsdRecord) GetOmapByte() int64`

GetOmapByte returns the OmapByte field if non-nil, zero value otherwise.

### GetOmapByteOk

`func (o *OsdRecord) GetOmapByteOk() (*int64, bool)`

GetOmapByteOk returns a tuple with the OmapByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmapByte

`func (o *OsdRecord) SetOmapByte(v int64)`

SetOmapByte sets OmapByte field to given value.

### HasOmapByte

`func (o *OsdRecord) HasOmapByte() bool`

HasOmapByte returns a boolean if a field has been set.

### GetOsdGroup

`func (o *OsdRecord) GetOsdGroup() OsdGroupNestview`

GetOsdGroup returns the OsdGroup field if non-nil, zero value otherwise.

### GetOsdGroupOk

`func (o *OsdRecord) GetOsdGroupOk() (*OsdGroupNestview, bool)`

GetOsdGroupOk returns a tuple with the OsdGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdGroup

`func (o *OsdRecord) SetOsdGroup(v OsdGroupNestview)`

SetOsdGroup sets OsdGroup field to given value.

### HasOsdGroup

`func (o *OsdRecord) HasOsdGroup() bool`

HasOsdGroup returns a boolean if a field has been set.

### GetOsdId

`func (o *OsdRecord) GetOsdId() int64`

GetOsdId returns the OsdId field if non-nil, zero value otherwise.

### GetOsdIdOk

`func (o *OsdRecord) GetOsdIdOk() (*int64, bool)`

GetOsdIdOk returns a tuple with the OsdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdId

`func (o *OsdRecord) SetOsdId(v int64)`

SetOsdId sets OsdId field to given value.

### HasOsdId

`func (o *OsdRecord) HasOsdId() bool`

HasOsdId returns a boolean if a field has been set.

### GetOspMetadataCluster

`func (o *OsdRecord) GetOspMetadataCluster() OspMetadataCluster`

GetOspMetadataCluster returns the OspMetadataCluster field if non-nil, zero value otherwise.

### GetOspMetadataClusterOk

`func (o *OsdRecord) GetOspMetadataClusterOk() (*OspMetadataCluster, bool)`

GetOspMetadataClusterOk returns a tuple with the OspMetadataCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspMetadataCluster

`func (o *OsdRecord) SetOspMetadataCluster(v OspMetadataCluster)`

SetOspMetadataCluster sets OspMetadataCluster field to given value.

### HasOspMetadataCluster

`func (o *OsdRecord) HasOspMetadataCluster() bool`

HasOspMetadataCluster returns a boolean if a field has been set.

### GetPartition

`func (o *OsdRecord) GetPartition() Partition`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *OsdRecord) GetPartitionOk() (*Partition, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *OsdRecord) SetPartition(v Partition)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *OsdRecord) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPool

`func (o *OsdRecord) GetPool() PoolNestview`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *OsdRecord) GetPoolOk() (*PoolNestview, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *OsdRecord) SetPool(v PoolNestview)`

SetPool sets Pool field to given value.

### HasPool

`func (o *OsdRecord) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetReadCacheSize

`func (o *OsdRecord) GetReadCacheSize() int64`

GetReadCacheSize returns the ReadCacheSize field if non-nil, zero value otherwise.

### GetReadCacheSizeOk

`func (o *OsdRecord) GetReadCacheSizeOk() (*int64, bool)`

GetReadCacheSizeOk returns a tuple with the ReadCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCacheSize

`func (o *OsdRecord) SetReadCacheSize(v int64)`

SetReadCacheSize sets ReadCacheSize field to given value.

### HasReadCacheSize

`func (o *OsdRecord) HasReadCacheSize() bool`

HasReadCacheSize returns a boolean if a field has been set.

### GetRole

`func (o *OsdRecord) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OsdRecord) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OsdRecord) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OsdRecord) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *OsdRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OsdRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OsdRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OsdRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusClass

`func (o *OsdRecord) GetStatusClass() string`

GetStatusClass returns the StatusClass field if non-nil, zero value otherwise.

### GetStatusClassOk

`func (o *OsdRecord) GetStatusClassOk() (*string, bool)`

GetStatusClassOk returns a tuple with the StatusClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusClass

`func (o *OsdRecord) SetStatusClass(v string)`

SetStatusClass sets StatusClass field to given value.

### HasStatusClass

`func (o *OsdRecord) HasStatusClass() bool`

HasStatusClass returns a boolean if a field has been set.

### GetType

`func (o *OsdRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OsdRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OsdRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OsdRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUp

`func (o *OsdRecord) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *OsdRecord) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *OsdRecord) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *OsdRecord) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUpdate

`func (o *OsdRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OsdRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OsdRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OsdRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUuid

`func (o *OsdRecord) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OsdRecord) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OsdRecord) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *OsdRecord) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetSamples

`func (o *OsdRecord) GetSamples() []OsdStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *OsdRecord) GetSamplesOk() (*[]OsdStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *OsdRecord) SetSamples(v []OsdStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *OsdRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


