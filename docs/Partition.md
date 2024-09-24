# Partition

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

## Methods

### NewPartition

`func NewPartition() *Partition`

NewPartition instantiates a new Partition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionWithDefaults

`func NewPartitionWithDefaults() *Partition`

NewPartitionWithDefaults instantiates a new Partition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *Partition) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *Partition) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *Partition) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *Partition) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *Partition) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Partition) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Partition) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Partition) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *Partition) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Partition) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Partition) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Partition) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDisk

`func (o *Partition) GetDisk() DiskNestview`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *Partition) GetDiskOk() (*DiskNestview, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *Partition) SetDisk(v DiskNestview)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *Partition) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetId

`func (o *Partition) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Partition) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Partition) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Partition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOmapByte

`func (o *Partition) GetOmapByte() int64`

GetOmapByte returns the OmapByte field if non-nil, zero value otherwise.

### GetOmapByteOk

`func (o *Partition) GetOmapByteOk() (*int64, bool)`

GetOmapByteOk returns a tuple with the OmapByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmapByte

`func (o *Partition) SetOmapByte(v int64)`

SetOmapByte sets OmapByte field to given value.

### HasOmapByte

`func (o *Partition) HasOmapByte() bool`

HasOmapByte returns a boolean if a field has been set.

### GetOmapDevicePath

`func (o *Partition) GetOmapDevicePath() string`

GetOmapDevicePath returns the OmapDevicePath field if non-nil, zero value otherwise.

### GetOmapDevicePathOk

`func (o *Partition) GetOmapDevicePathOk() (*string, bool)`

GetOmapDevicePathOk returns a tuple with the OmapDevicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmapDevicePath

`func (o *Partition) SetOmapDevicePath(v string)`

SetOmapDevicePath sets OmapDevicePath field to given value.

### HasOmapDevicePath

`func (o *Partition) HasOmapDevicePath() bool`

HasOmapDevicePath returns a boolean if a field has been set.

### GetOspMetadataCluster

`func (o *Partition) GetOspMetadataCluster() OspMetadataCluster`

GetOspMetadataCluster returns the OspMetadataCluster field if non-nil, zero value otherwise.

### GetOspMetadataClusterOk

`func (o *Partition) GetOspMetadataClusterOk() (*OspMetadataCluster, bool)`

GetOspMetadataClusterOk returns a tuple with the OspMetadataCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspMetadataCluster

`func (o *Partition) SetOspMetadataCluster(v OspMetadataCluster)`

SetOspMetadataCluster sets OspMetadataCluster field to given value.

### HasOspMetadataCluster

`func (o *Partition) HasOspMetadataCluster() bool`

HasOspMetadataCluster returns a boolean if a field has been set.

### GetPath

`func (o *Partition) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Partition) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Partition) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Partition) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSize

`func (o *Partition) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Partition) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Partition) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *Partition) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *Partition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Partition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Partition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Partition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Partition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Partition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Partition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Partition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *Partition) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Partition) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Partition) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Partition) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUsed

`func (o *Partition) GetUsed() bool`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *Partition) GetUsedOk() (*bool, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *Partition) SetUsed(v bool)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *Partition) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetUuid

`func (o *Partition) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Partition) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Partition) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Partition) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *Partition) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Partition) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Partition) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Partition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


