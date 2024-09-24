# Osd

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

## Methods

### NewOsd

`func NewOsd() *Osd`

NewOsd instantiates a new Osd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsdWithDefaults

`func NewOsdWithDefaults() *Osd`

NewOsdWithDefaults instantiates a new Osd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *Osd) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *Osd) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *Osd) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *Osd) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetBindOsdGroup

`func (o *Osd) GetBindOsdGroup() OsdGroupNestview`

GetBindOsdGroup returns the BindOsdGroup field if non-nil, zero value otherwise.

### GetBindOsdGroupOk

`func (o *Osd) GetBindOsdGroupOk() (*OsdGroupNestview, bool)`

GetBindOsdGroupOk returns a tuple with the BindOsdGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindOsdGroup

`func (o *Osd) SetBindOsdGroup(v OsdGroupNestview)`

SetBindOsdGroup sets BindOsdGroup field to given value.

### HasBindOsdGroup

`func (o *Osd) HasBindOsdGroup() bool`

HasBindOsdGroup returns a boolean if a field has been set.

### GetBindPool

`func (o *Osd) GetBindPool() PoolNestview`

GetBindPool returns the BindPool field if non-nil, zero value otherwise.

### GetBindPoolOk

`func (o *Osd) GetBindPoolOk() (*PoolNestview, bool)`

GetBindPoolOk returns a tuple with the BindPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPool

`func (o *Osd) SetBindPool(v PoolNestview)`

SetBindPool sets BindPool field to given value.

### HasBindPool

`func (o *Osd) HasBindPool() bool`

HasBindPool returns a boolean if a field has been set.

### GetCluster

`func (o *Osd) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Osd) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Osd) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Osd) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *Osd) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Osd) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Osd) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Osd) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataDir

`func (o *Osd) GetDataDir() string`

GetDataDir returns the DataDir field if non-nil, zero value otherwise.

### GetDataDirOk

`func (o *Osd) GetDataDirOk() (*string, bool)`

GetDataDirOk returns a tuple with the DataDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDir

`func (o *Osd) SetDataDir(v string)`

SetDataDir sets DataDir field to given value.

### HasDataDir

`func (o *Osd) HasDataDir() bool`

HasDataDir returns a boolean if a field has been set.

### GetDisk

`func (o *Osd) GetDisk() Disk`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *Osd) GetDiskOk() (*Disk, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *Osd) SetDisk(v Disk)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *Osd) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetEncryptEnabled

`func (o *Osd) GetEncryptEnabled() bool`

GetEncryptEnabled returns the EncryptEnabled field if non-nil, zero value otherwise.

### GetEncryptEnabledOk

`func (o *Osd) GetEncryptEnabledOk() (*bool, bool)`

GetEncryptEnabledOk returns a tuple with the EncryptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptEnabled

`func (o *Osd) SetEncryptEnabled(v bool)`

SetEncryptEnabled sets EncryptEnabled field to given value.

### HasEncryptEnabled

`func (o *Osd) HasEncryptEnabled() bool`

HasEncryptEnabled returns a boolean if a field has been set.

### GetExitCount

`func (o *Osd) GetExitCount() int64`

GetExitCount returns the ExitCount field if non-nil, zero value otherwise.

### GetExitCountOk

`func (o *Osd) GetExitCountOk() (*int64, bool)`

GetExitCountOk returns a tuple with the ExitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCount

`func (o *Osd) SetExitCount(v int64)`

SetExitCount sets ExitCount field to given value.

### HasExitCount

`func (o *Osd) HasExitCount() bool`

HasExitCount returns a boolean if a field has been set.

### GetExitTime

`func (o *Osd) GetExitTime() time.Time`

GetExitTime returns the ExitTime field if non-nil, zero value otherwise.

### GetExitTimeOk

`func (o *Osd) GetExitTimeOk() (*time.Time, bool)`

GetExitTimeOk returns a tuple with the ExitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitTime

`func (o *Osd) SetExitTime(v time.Time)`

SetExitTime sets ExitTime field to given value.

### HasExitTime

`func (o *Osd) HasExitTime() bool`

HasExitTime returns a boolean if a field has been set.

### GetHost

`func (o *Osd) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Osd) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Osd) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *Osd) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *Osd) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Osd) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Osd) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Osd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIn

`func (o *Osd) GetIn() bool`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *Osd) GetInOk() (*bool, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *Osd) SetIn(v bool)`

SetIn sets In field to given value.

### HasIn

`func (o *Osd) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetInitTime

`func (o *Osd) GetInitTime() time.Time`

GetInitTime returns the InitTime field if non-nil, zero value otherwise.

### GetInitTimeOk

`func (o *Osd) GetInitTimeOk() (*time.Time, bool)`

GetInitTimeOk returns a tuple with the InitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitTime

`func (o *Osd) SetInitTime(v time.Time)`

SetInitTime sets InitTime field to given value.

### HasInitTime

`func (o *Osd) HasInitTime() bool`

HasInitTime returns a boolean if a field has been set.

### GetLastScrubTime

`func (o *Osd) GetLastScrubTime() time.Time`

GetLastScrubTime returns the LastScrubTime field if non-nil, zero value otherwise.

### GetLastScrubTimeOk

`func (o *Osd) GetLastScrubTimeOk() (*time.Time, bool)`

GetLastScrubTimeOk returns a tuple with the LastScrubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScrubTime

`func (o *Osd) SetLastScrubTime(v time.Time)`

SetLastScrubTime sets LastScrubTime field to given value.

### HasLastScrubTime

`func (o *Osd) HasLastScrubTime() bool`

HasLastScrubTime returns a boolean if a field has been set.

### GetMetaBytes

`func (o *Osd) GetMetaBytes() int64`

GetMetaBytes returns the MetaBytes field if non-nil, zero value otherwise.

### GetMetaBytesOk

`func (o *Osd) GetMetaBytesOk() (*int64, bool)`

GetMetaBytesOk returns a tuple with the MetaBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaBytes

`func (o *Osd) SetMetaBytes(v int64)`

SetMetaBytes sets MetaBytes field to given value.

### HasMetaBytes

`func (o *Osd) HasMetaBytes() bool`

HasMetaBytes returns a boolean if a field has been set.

### GetMinAllocSize

`func (o *Osd) GetMinAllocSize() int64`

GetMinAllocSize returns the MinAllocSize field if non-nil, zero value otherwise.

### GetMinAllocSizeOk

`func (o *Osd) GetMinAllocSizeOk() (*int64, bool)`

GetMinAllocSizeOk returns a tuple with the MinAllocSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAllocSize

`func (o *Osd) SetMinAllocSize(v int64)`

SetMinAllocSize sets MinAllocSize field to given value.

### HasMinAllocSize

`func (o *Osd) HasMinAllocSize() bool`

HasMinAllocSize returns a boolean if a field has been set.

### GetName

`func (o *Osd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Osd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Osd) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Osd) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumaNode

`func (o *Osd) GetNumaNode() int64`

GetNumaNode returns the NumaNode field if non-nil, zero value otherwise.

### GetNumaNodeOk

`func (o *Osd) GetNumaNodeOk() (*int64, bool)`

GetNumaNodeOk returns a tuple with the NumaNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaNode

`func (o *Osd) SetNumaNode(v int64)`

SetNumaNode sets NumaNode field to given value.

### HasNumaNode

`func (o *Osd) HasNumaNode() bool`

HasNumaNode returns a boolean if a field has been set.

### GetOmapByte

`func (o *Osd) GetOmapByte() int64`

GetOmapByte returns the OmapByte field if non-nil, zero value otherwise.

### GetOmapByteOk

`func (o *Osd) GetOmapByteOk() (*int64, bool)`

GetOmapByteOk returns a tuple with the OmapByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmapByte

`func (o *Osd) SetOmapByte(v int64)`

SetOmapByte sets OmapByte field to given value.

### HasOmapByte

`func (o *Osd) HasOmapByte() bool`

HasOmapByte returns a boolean if a field has been set.

### GetOsdGroup

`func (o *Osd) GetOsdGroup() OsdGroupNestview`

GetOsdGroup returns the OsdGroup field if non-nil, zero value otherwise.

### GetOsdGroupOk

`func (o *Osd) GetOsdGroupOk() (*OsdGroupNestview, bool)`

GetOsdGroupOk returns a tuple with the OsdGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdGroup

`func (o *Osd) SetOsdGroup(v OsdGroupNestview)`

SetOsdGroup sets OsdGroup field to given value.

### HasOsdGroup

`func (o *Osd) HasOsdGroup() bool`

HasOsdGroup returns a boolean if a field has been set.

### GetOsdId

`func (o *Osd) GetOsdId() int64`

GetOsdId returns the OsdId field if non-nil, zero value otherwise.

### GetOsdIdOk

`func (o *Osd) GetOsdIdOk() (*int64, bool)`

GetOsdIdOk returns a tuple with the OsdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdId

`func (o *Osd) SetOsdId(v int64)`

SetOsdId sets OsdId field to given value.

### HasOsdId

`func (o *Osd) HasOsdId() bool`

HasOsdId returns a boolean if a field has been set.

### GetOspMetadataCluster

`func (o *Osd) GetOspMetadataCluster() OspMetadataCluster`

GetOspMetadataCluster returns the OspMetadataCluster field if non-nil, zero value otherwise.

### GetOspMetadataClusterOk

`func (o *Osd) GetOspMetadataClusterOk() (*OspMetadataCluster, bool)`

GetOspMetadataClusterOk returns a tuple with the OspMetadataCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspMetadataCluster

`func (o *Osd) SetOspMetadataCluster(v OspMetadataCluster)`

SetOspMetadataCluster sets OspMetadataCluster field to given value.

### HasOspMetadataCluster

`func (o *Osd) HasOspMetadataCluster() bool`

HasOspMetadataCluster returns a boolean if a field has been set.

### GetPartition

`func (o *Osd) GetPartition() Partition`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *Osd) GetPartitionOk() (*Partition, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *Osd) SetPartition(v Partition)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *Osd) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPool

`func (o *Osd) GetPool() PoolNestview`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *Osd) GetPoolOk() (*PoolNestview, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *Osd) SetPool(v PoolNestview)`

SetPool sets Pool field to given value.

### HasPool

`func (o *Osd) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetReadCacheSize

`func (o *Osd) GetReadCacheSize() int64`

GetReadCacheSize returns the ReadCacheSize field if non-nil, zero value otherwise.

### GetReadCacheSizeOk

`func (o *Osd) GetReadCacheSizeOk() (*int64, bool)`

GetReadCacheSizeOk returns a tuple with the ReadCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCacheSize

`func (o *Osd) SetReadCacheSize(v int64)`

SetReadCacheSize sets ReadCacheSize field to given value.

### HasReadCacheSize

`func (o *Osd) HasReadCacheSize() bool`

HasReadCacheSize returns a boolean if a field has been set.

### GetRole

`func (o *Osd) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Osd) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Osd) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Osd) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *Osd) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Osd) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Osd) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Osd) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusClass

`func (o *Osd) GetStatusClass() string`

GetStatusClass returns the StatusClass field if non-nil, zero value otherwise.

### GetStatusClassOk

`func (o *Osd) GetStatusClassOk() (*string, bool)`

GetStatusClassOk returns a tuple with the StatusClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusClass

`func (o *Osd) SetStatusClass(v string)`

SetStatusClass sets StatusClass field to given value.

### HasStatusClass

`func (o *Osd) HasStatusClass() bool`

HasStatusClass returns a boolean if a field has been set.

### GetType

`func (o *Osd) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Osd) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Osd) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Osd) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUp

`func (o *Osd) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *Osd) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *Osd) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *Osd) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUpdate

`func (o *Osd) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Osd) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Osd) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Osd) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUuid

`func (o *Osd) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Osd) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Osd) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Osd) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


