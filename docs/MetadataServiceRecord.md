# MetadataServiceRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Complete** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Disk** | Pointer to [**DiskNestview**](DiskNestview.md) |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MetaBytes** | Pointer to **int64** | journal partition size, 5G | [optional] 
**MetadataCluster** | Pointer to [**MetadataClusterNestview**](MetadataClusterNestview.md) |  | [optional] 
**Name** | Pointer to **string** | xmds.N | [optional] 
**PlacementNode** | Pointer to [**PlacementNodeNestview**](PlacementNodeNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**XmdsId** | Pointer to **int64** |  | [optional] 
**Samples** | Pointer to [**[]MetadataServiceStat**](MetadataServiceStat.md) |  | [optional] 

## Methods

### NewMetadataServiceRecord

`func NewMetadataServiceRecord() *MetadataServiceRecord`

NewMetadataServiceRecord instantiates a new MetadataServiceRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataServiceRecordWithDefaults

`func NewMetadataServiceRecordWithDefaults() *MetadataServiceRecord`

NewMetadataServiceRecordWithDefaults instantiates a new MetadataServiceRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *MetadataServiceRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *MetadataServiceRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *MetadataServiceRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *MetadataServiceRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *MetadataServiceRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *MetadataServiceRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *MetadataServiceRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *MetadataServiceRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetComplete

`func (o *MetadataServiceRecord) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *MetadataServiceRecord) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *MetadataServiceRecord) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *MetadataServiceRecord) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetCreate

`func (o *MetadataServiceRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *MetadataServiceRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *MetadataServiceRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *MetadataServiceRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDisk

`func (o *MetadataServiceRecord) GetDisk() DiskNestview`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *MetadataServiceRecord) GetDiskOk() (*DiskNestview, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *MetadataServiceRecord) SetDisk(v DiskNestview)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *MetadataServiceRecord) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetHost

`func (o *MetadataServiceRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MetadataServiceRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MetadataServiceRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *MetadataServiceRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *MetadataServiceRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataServiceRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataServiceRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataServiceRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetaBytes

`func (o *MetadataServiceRecord) GetMetaBytes() int64`

GetMetaBytes returns the MetaBytes field if non-nil, zero value otherwise.

### GetMetaBytesOk

`func (o *MetadataServiceRecord) GetMetaBytesOk() (*int64, bool)`

GetMetaBytesOk returns a tuple with the MetaBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaBytes

`func (o *MetadataServiceRecord) SetMetaBytes(v int64)`

SetMetaBytes sets MetaBytes field to given value.

### HasMetaBytes

`func (o *MetadataServiceRecord) HasMetaBytes() bool`

HasMetaBytes returns a boolean if a field has been set.

### GetMetadataCluster

`func (o *MetadataServiceRecord) GetMetadataCluster() MetadataClusterNestview`

GetMetadataCluster returns the MetadataCluster field if non-nil, zero value otherwise.

### GetMetadataClusterOk

`func (o *MetadataServiceRecord) GetMetadataClusterOk() (*MetadataClusterNestview, bool)`

GetMetadataClusterOk returns a tuple with the MetadataCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCluster

`func (o *MetadataServiceRecord) SetMetadataCluster(v MetadataClusterNestview)`

SetMetadataCluster sets MetadataCluster field to given value.

### HasMetadataCluster

`func (o *MetadataServiceRecord) HasMetadataCluster() bool`

HasMetadataCluster returns a boolean if a field has been set.

### GetName

`func (o *MetadataServiceRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataServiceRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataServiceRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataServiceRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacementNode

`func (o *MetadataServiceRecord) GetPlacementNode() PlacementNodeNestview`

GetPlacementNode returns the PlacementNode field if non-nil, zero value otherwise.

### GetPlacementNodeOk

`func (o *MetadataServiceRecord) GetPlacementNodeOk() (*PlacementNodeNestview, bool)`

GetPlacementNodeOk returns a tuple with the PlacementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementNode

`func (o *MetadataServiceRecord) SetPlacementNode(v PlacementNodeNestview)`

SetPlacementNode sets PlacementNode field to given value.

### HasPlacementNode

`func (o *MetadataServiceRecord) HasPlacementNode() bool`

HasPlacementNode returns a boolean if a field has been set.

### GetStatus

`func (o *MetadataServiceRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetadataServiceRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MetadataServiceRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MetadataServiceRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUp

`func (o *MetadataServiceRecord) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *MetadataServiceRecord) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *MetadataServiceRecord) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *MetadataServiceRecord) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUpdate

`func (o *MetadataServiceRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *MetadataServiceRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *MetadataServiceRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *MetadataServiceRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetXmdsId

`func (o *MetadataServiceRecord) GetXmdsId() int64`

GetXmdsId returns the XmdsId field if non-nil, zero value otherwise.

### GetXmdsIdOk

`func (o *MetadataServiceRecord) GetXmdsIdOk() (*int64, bool)`

GetXmdsIdOk returns a tuple with the XmdsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmdsId

`func (o *MetadataServiceRecord) SetXmdsId(v int64)`

SetXmdsId sets XmdsId field to given value.

### HasXmdsId

`func (o *MetadataServiceRecord) HasXmdsId() bool`

HasXmdsId returns a boolean if a field has been set.

### GetSamples

`func (o *MetadataServiceRecord) GetSamples() []MetadataServiceStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *MetadataServiceRecord) GetSamplesOk() (*[]MetadataServiceStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *MetadataServiceRecord) SetSamples(v []MetadataServiceStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *MetadataServiceRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


