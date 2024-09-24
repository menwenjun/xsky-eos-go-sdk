# PartitionRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Disk** | Pointer to [**DiskNestview**](DiskNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**OmapByte** | Pointer to **int64** |  | [optional] 
**OmapDevicePath** | Pointer to **string** | omap part path | [optional] 
**OspMetadataCluster** | Pointer to [**OspMetadataCluster**](OspMetadataCluster.md) |  | [optional] 
**Path** | Pointer to **string** | part path, generated from part num and disk device | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Used** | Pointer to **bool** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**Samples** | Pointer to [**[]PartitionStat**](PartitionStat.md) |  | [optional] 

## Methods

### NewPartitionRecord

`func NewPartitionRecord() *PartitionRecord`

NewPartitionRecord instantiates a new PartitionRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionRecordWithDefaults

`func NewPartitionRecordWithDefaults() *PartitionRecord`

NewPartitionRecordWithDefaults instantiates a new PartitionRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *PartitionRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *PartitionRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *PartitionRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *PartitionRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *PartitionRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PartitionRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PartitionRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PartitionRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *PartitionRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *PartitionRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *PartitionRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *PartitionRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDisk

`func (o *PartitionRecord) GetDisk() DiskNestview`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *PartitionRecord) GetDiskOk() (*DiskNestview, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *PartitionRecord) SetDisk(v DiskNestview)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *PartitionRecord) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetId

`func (o *PartitionRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartitionRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartitionRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PartitionRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOmapByte

`func (o *PartitionRecord) GetOmapByte() int64`

GetOmapByte returns the OmapByte field if non-nil, zero value otherwise.

### GetOmapByteOk

`func (o *PartitionRecord) GetOmapByteOk() (*int64, bool)`

GetOmapByteOk returns a tuple with the OmapByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmapByte

`func (o *PartitionRecord) SetOmapByte(v int64)`

SetOmapByte sets OmapByte field to given value.

### HasOmapByte

`func (o *PartitionRecord) HasOmapByte() bool`

HasOmapByte returns a boolean if a field has been set.

### GetOmapDevicePath

`func (o *PartitionRecord) GetOmapDevicePath() string`

GetOmapDevicePath returns the OmapDevicePath field if non-nil, zero value otherwise.

### GetOmapDevicePathOk

`func (o *PartitionRecord) GetOmapDevicePathOk() (*string, bool)`

GetOmapDevicePathOk returns a tuple with the OmapDevicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmapDevicePath

`func (o *PartitionRecord) SetOmapDevicePath(v string)`

SetOmapDevicePath sets OmapDevicePath field to given value.

### HasOmapDevicePath

`func (o *PartitionRecord) HasOmapDevicePath() bool`

HasOmapDevicePath returns a boolean if a field has been set.

### GetOspMetadataCluster

`func (o *PartitionRecord) GetOspMetadataCluster() OspMetadataCluster`

GetOspMetadataCluster returns the OspMetadataCluster field if non-nil, zero value otherwise.

### GetOspMetadataClusterOk

`func (o *PartitionRecord) GetOspMetadataClusterOk() (*OspMetadataCluster, bool)`

GetOspMetadataClusterOk returns a tuple with the OspMetadataCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspMetadataCluster

`func (o *PartitionRecord) SetOspMetadataCluster(v OspMetadataCluster)`

SetOspMetadataCluster sets OspMetadataCluster field to given value.

### HasOspMetadataCluster

`func (o *PartitionRecord) HasOspMetadataCluster() bool`

HasOspMetadataCluster returns a boolean if a field has been set.

### GetPath

`func (o *PartitionRecord) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PartitionRecord) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PartitionRecord) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PartitionRecord) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSize

`func (o *PartitionRecord) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PartitionRecord) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PartitionRecord) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *PartitionRecord) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *PartitionRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PartitionRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PartitionRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PartitionRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *PartitionRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartitionRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartitionRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PartitionRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *PartitionRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *PartitionRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *PartitionRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *PartitionRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUsed

`func (o *PartitionRecord) GetUsed() bool`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *PartitionRecord) GetUsedOk() (*bool, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *PartitionRecord) SetUsed(v bool)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *PartitionRecord) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetUuid

`func (o *PartitionRecord) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PartitionRecord) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PartitionRecord) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PartitionRecord) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *PartitionRecord) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PartitionRecord) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PartitionRecord) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PartitionRecord) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSamples

`func (o *PartitionRecord) GetSamples() []PartitionStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *PartitionRecord) GetSamplesOk() (*[]PartitionStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *PartitionRecord) SetSamples(v []PartitionStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *PartitionRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


