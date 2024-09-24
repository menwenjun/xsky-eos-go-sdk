# DiskRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Bytes** | Pointer to **int64** | size of disk | [optional] 
**CacheCreate** | Pointer to **time.Time** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**ControllerId** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Device** | Pointer to **string** |  | [optional] 
**DiskType** | Pointer to **string** |  | [optional] 
**DriverType** | Pointer to **string** |  | [optional] 
**EnclosureId** | Pointer to **string** |  | [optional] 
**Host** | Pointer to [**Host**](Host.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsCache** | Pointer to **bool** | used as cache disk, deprecated | [optional] 
**IsRoot** | Pointer to **bool** | used as root disk, deprecated | [optional] 
**LightingStatus** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PartitionNum** | Pointer to **int64** |  | [optional] 
**PartitionTypes** | Pointer to **[]string** |  | [optional] 
**Partitions** | Pointer to [**[]Partition**](Partition.md) |  | [optional] 
**PowerSafe** | Pointer to **bool** |  | [optional] 
**RaidStatus** | Pointer to **string** |  | [optional] 
**RotationRate** | Pointer to **string** |  | [optional] 
**Rotational** | Pointer to **bool** |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**SlotId** | Pointer to **string** |  | [optional] 
**SmartAttrs** | Pointer to [**[]SmartAttr**](SmartAttr.md) |  | [optional] 
**SsdLifeLeft** | Pointer to **int64** |  | [optional] 
**SsdType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Usage** | Pointer to **string** | disk usage | [optional] 
**Used** | Pointer to **bool** | disk is used, deprecated | [optional] 
**Wwid** | Pointer to **string** |  | [optional] 
**Samples** | Pointer to [**[]DiskStat**](DiskStat.md) |  | [optional] 

## Methods

### NewDiskRecord

`func NewDiskRecord() *DiskRecord`

NewDiskRecord instantiates a new DiskRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskRecordWithDefaults

`func NewDiskRecordWithDefaults() *DiskRecord`

NewDiskRecordWithDefaults instantiates a new DiskRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DiskRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DiskRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DiskRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DiskRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetBytes

`func (o *DiskRecord) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *DiskRecord) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *DiskRecord) SetBytes(v int64)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *DiskRecord) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetCacheCreate

`func (o *DiskRecord) GetCacheCreate() time.Time`

GetCacheCreate returns the CacheCreate field if non-nil, zero value otherwise.

### GetCacheCreateOk

`func (o *DiskRecord) GetCacheCreateOk() (*time.Time, bool)`

GetCacheCreateOk returns a tuple with the CacheCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreate

`func (o *DiskRecord) SetCacheCreate(v time.Time)`

SetCacheCreate sets CacheCreate field to given value.

### HasCacheCreate

`func (o *DiskRecord) HasCacheCreate() bool`

HasCacheCreate returns a boolean if a field has been set.

### GetChannelId

`func (o *DiskRecord) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *DiskRecord) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *DiskRecord) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *DiskRecord) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCluster

`func (o *DiskRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DiskRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DiskRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DiskRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetControllerId

`func (o *DiskRecord) GetControllerId() string`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *DiskRecord) GetControllerIdOk() (*string, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *DiskRecord) SetControllerId(v string)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *DiskRecord) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### GetCreate

`func (o *DiskRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DiskRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DiskRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DiskRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDevice

`func (o *DiskRecord) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DiskRecord) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DiskRecord) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *DiskRecord) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDiskType

`func (o *DiskRecord) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *DiskRecord) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *DiskRecord) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *DiskRecord) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetDriverType

`func (o *DiskRecord) GetDriverType() string`

GetDriverType returns the DriverType field if non-nil, zero value otherwise.

### GetDriverTypeOk

`func (o *DiskRecord) GetDriverTypeOk() (*string, bool)`

GetDriverTypeOk returns a tuple with the DriverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverType

`func (o *DiskRecord) SetDriverType(v string)`

SetDriverType sets DriverType field to given value.

### HasDriverType

`func (o *DiskRecord) HasDriverType() bool`

HasDriverType returns a boolean if a field has been set.

### GetEnclosureId

`func (o *DiskRecord) GetEnclosureId() string`

GetEnclosureId returns the EnclosureId field if non-nil, zero value otherwise.

### GetEnclosureIdOk

`func (o *DiskRecord) GetEnclosureIdOk() (*string, bool)`

GetEnclosureIdOk returns a tuple with the EnclosureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureId

`func (o *DiskRecord) SetEnclosureId(v string)`

SetEnclosureId sets EnclosureId field to given value.

### HasEnclosureId

`func (o *DiskRecord) HasEnclosureId() bool`

HasEnclosureId returns a boolean if a field has been set.

### GetHost

`func (o *DiskRecord) GetHost() Host`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DiskRecord) GetHostOk() (*Host, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DiskRecord) SetHost(v Host)`

SetHost sets Host field to given value.

### HasHost

`func (o *DiskRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *DiskRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DiskRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCache

`func (o *DiskRecord) GetIsCache() bool`

GetIsCache returns the IsCache field if non-nil, zero value otherwise.

### GetIsCacheOk

`func (o *DiskRecord) GetIsCacheOk() (*bool, bool)`

GetIsCacheOk returns a tuple with the IsCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCache

`func (o *DiskRecord) SetIsCache(v bool)`

SetIsCache sets IsCache field to given value.

### HasIsCache

`func (o *DiskRecord) HasIsCache() bool`

HasIsCache returns a boolean if a field has been set.

### GetIsRoot

`func (o *DiskRecord) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *DiskRecord) GetIsRootOk() (*bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoot

`func (o *DiskRecord) SetIsRoot(v bool)`

SetIsRoot sets IsRoot field to given value.

### HasIsRoot

`func (o *DiskRecord) HasIsRoot() bool`

HasIsRoot returns a boolean if a field has been set.

### GetLightingStatus

`func (o *DiskRecord) GetLightingStatus() string`

GetLightingStatus returns the LightingStatus field if non-nil, zero value otherwise.

### GetLightingStatusOk

`func (o *DiskRecord) GetLightingStatusOk() (*string, bool)`

GetLightingStatusOk returns a tuple with the LightingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLightingStatus

`func (o *DiskRecord) SetLightingStatus(v string)`

SetLightingStatus sets LightingStatus field to given value.

### HasLightingStatus

`func (o *DiskRecord) HasLightingStatus() bool`

HasLightingStatus returns a boolean if a field has been set.

### GetModel

`func (o *DiskRecord) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DiskRecord) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DiskRecord) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DiskRecord) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *DiskRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiskRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiskRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiskRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartitionNum

`func (o *DiskRecord) GetPartitionNum() int64`

GetPartitionNum returns the PartitionNum field if non-nil, zero value otherwise.

### GetPartitionNumOk

`func (o *DiskRecord) GetPartitionNumOk() (*int64, bool)`

GetPartitionNumOk returns a tuple with the PartitionNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionNum

`func (o *DiskRecord) SetPartitionNum(v int64)`

SetPartitionNum sets PartitionNum field to given value.

### HasPartitionNum

`func (o *DiskRecord) HasPartitionNum() bool`

HasPartitionNum returns a boolean if a field has been set.

### GetPartitionTypes

`func (o *DiskRecord) GetPartitionTypes() []string`

GetPartitionTypes returns the PartitionTypes field if non-nil, zero value otherwise.

### GetPartitionTypesOk

`func (o *DiskRecord) GetPartitionTypesOk() (*[]string, bool)`

GetPartitionTypesOk returns a tuple with the PartitionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionTypes

`func (o *DiskRecord) SetPartitionTypes(v []string)`

SetPartitionTypes sets PartitionTypes field to given value.

### HasPartitionTypes

`func (o *DiskRecord) HasPartitionTypes() bool`

HasPartitionTypes returns a boolean if a field has been set.

### GetPartitions

`func (o *DiskRecord) GetPartitions() []Partition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *DiskRecord) GetPartitionsOk() (*[]Partition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *DiskRecord) SetPartitions(v []Partition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *DiskRecord) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetPowerSafe

`func (o *DiskRecord) GetPowerSafe() bool`

GetPowerSafe returns the PowerSafe field if non-nil, zero value otherwise.

### GetPowerSafeOk

`func (o *DiskRecord) GetPowerSafeOk() (*bool, bool)`

GetPowerSafeOk returns a tuple with the PowerSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSafe

`func (o *DiskRecord) SetPowerSafe(v bool)`

SetPowerSafe sets PowerSafe field to given value.

### HasPowerSafe

`func (o *DiskRecord) HasPowerSafe() bool`

HasPowerSafe returns a boolean if a field has been set.

### GetRaidStatus

`func (o *DiskRecord) GetRaidStatus() string`

GetRaidStatus returns the RaidStatus field if non-nil, zero value otherwise.

### GetRaidStatusOk

`func (o *DiskRecord) GetRaidStatusOk() (*string, bool)`

GetRaidStatusOk returns a tuple with the RaidStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidStatus

`func (o *DiskRecord) SetRaidStatus(v string)`

SetRaidStatus sets RaidStatus field to given value.

### HasRaidStatus

`func (o *DiskRecord) HasRaidStatus() bool`

HasRaidStatus returns a boolean if a field has been set.

### GetRotationRate

`func (o *DiskRecord) GetRotationRate() string`

GetRotationRate returns the RotationRate field if non-nil, zero value otherwise.

### GetRotationRateOk

`func (o *DiskRecord) GetRotationRateOk() (*string, bool)`

GetRotationRateOk returns a tuple with the RotationRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationRate

`func (o *DiskRecord) SetRotationRate(v string)`

SetRotationRate sets RotationRate field to given value.

### HasRotationRate

`func (o *DiskRecord) HasRotationRate() bool`

HasRotationRate returns a boolean if a field has been set.

### GetRotational

`func (o *DiskRecord) GetRotational() bool`

GetRotational returns the Rotational field if non-nil, zero value otherwise.

### GetRotationalOk

`func (o *DiskRecord) GetRotationalOk() (*bool, bool)`

GetRotationalOk returns a tuple with the Rotational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotational

`func (o *DiskRecord) SetRotational(v bool)`

SetRotational sets Rotational field to given value.

### HasRotational

`func (o *DiskRecord) HasRotational() bool`

HasRotational returns a boolean if a field has been set.

### GetSerial

`func (o *DiskRecord) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *DiskRecord) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *DiskRecord) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *DiskRecord) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSlotId

`func (o *DiskRecord) GetSlotId() string`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *DiskRecord) GetSlotIdOk() (*string, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *DiskRecord) SetSlotId(v string)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *DiskRecord) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetSmartAttrs

`func (o *DiskRecord) GetSmartAttrs() []SmartAttr`

GetSmartAttrs returns the SmartAttrs field if non-nil, zero value otherwise.

### GetSmartAttrsOk

`func (o *DiskRecord) GetSmartAttrsOk() (*[]SmartAttr, bool)`

GetSmartAttrsOk returns a tuple with the SmartAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartAttrs

`func (o *DiskRecord) SetSmartAttrs(v []SmartAttr)`

SetSmartAttrs sets SmartAttrs field to given value.

### HasSmartAttrs

`func (o *DiskRecord) HasSmartAttrs() bool`

HasSmartAttrs returns a boolean if a field has been set.

### GetSsdLifeLeft

`func (o *DiskRecord) GetSsdLifeLeft() int64`

GetSsdLifeLeft returns the SsdLifeLeft field if non-nil, zero value otherwise.

### GetSsdLifeLeftOk

`func (o *DiskRecord) GetSsdLifeLeftOk() (*int64, bool)`

GetSsdLifeLeftOk returns a tuple with the SsdLifeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdLifeLeft

`func (o *DiskRecord) SetSsdLifeLeft(v int64)`

SetSsdLifeLeft sets SsdLifeLeft field to given value.

### HasSsdLifeLeft

`func (o *DiskRecord) HasSsdLifeLeft() bool`

HasSsdLifeLeft returns a boolean if a field has been set.

### GetSsdType

`func (o *DiskRecord) GetSsdType() string`

GetSsdType returns the SsdType field if non-nil, zero value otherwise.

### GetSsdTypeOk

`func (o *DiskRecord) GetSsdTypeOk() (*string, bool)`

GetSsdTypeOk returns a tuple with the SsdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdType

`func (o *DiskRecord) SetSsdType(v string)`

SetSsdType sets SsdType field to given value.

### HasSsdType

`func (o *DiskRecord) HasSsdType() bool`

HasSsdType returns a boolean if a field has been set.

### GetStatus

`func (o *DiskRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiskRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiskRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiskRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *DiskRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiskRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiskRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiskRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *DiskRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DiskRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DiskRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DiskRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUsage

`func (o *DiskRecord) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DiskRecord) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DiskRecord) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *DiskRecord) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUsed

`func (o *DiskRecord) GetUsed() bool`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *DiskRecord) GetUsedOk() (*bool, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *DiskRecord) SetUsed(v bool)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *DiskRecord) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetWwid

`func (o *DiskRecord) GetWwid() string`

GetWwid returns the Wwid field if non-nil, zero value otherwise.

### GetWwidOk

`func (o *DiskRecord) GetWwidOk() (*string, bool)`

GetWwidOk returns a tuple with the Wwid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwid

`func (o *DiskRecord) SetWwid(v string)`

SetWwid sets Wwid field to given value.

### HasWwid

`func (o *DiskRecord) HasWwid() bool`

HasWwid returns a boolean if a field has been set.

### GetSamples

`func (o *DiskRecord) GetSamples() []DiskStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *DiskRecord) GetSamplesOk() (*[]DiskStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *DiskRecord) SetSamples(v []DiskStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *DiskRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


