# MetadataService

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

## Methods

### NewMetadataService

`func NewMetadataService() *MetadataService`

NewMetadataService instantiates a new MetadataService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataServiceWithDefaults

`func NewMetadataServiceWithDefaults() *MetadataService`

NewMetadataServiceWithDefaults instantiates a new MetadataService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *MetadataService) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *MetadataService) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *MetadataService) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *MetadataService) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *MetadataService) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *MetadataService) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *MetadataService) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *MetadataService) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetComplete

`func (o *MetadataService) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *MetadataService) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *MetadataService) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *MetadataService) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetCreate

`func (o *MetadataService) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *MetadataService) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *MetadataService) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *MetadataService) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDisk

`func (o *MetadataService) GetDisk() DiskNestview`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *MetadataService) GetDiskOk() (*DiskNestview, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *MetadataService) SetDisk(v DiskNestview)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *MetadataService) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetHost

`func (o *MetadataService) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MetadataService) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MetadataService) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *MetadataService) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *MetadataService) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataService) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataService) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetaBytes

`func (o *MetadataService) GetMetaBytes() int64`

GetMetaBytes returns the MetaBytes field if non-nil, zero value otherwise.

### GetMetaBytesOk

`func (o *MetadataService) GetMetaBytesOk() (*int64, bool)`

GetMetaBytesOk returns a tuple with the MetaBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaBytes

`func (o *MetadataService) SetMetaBytes(v int64)`

SetMetaBytes sets MetaBytes field to given value.

### HasMetaBytes

`func (o *MetadataService) HasMetaBytes() bool`

HasMetaBytes returns a boolean if a field has been set.

### GetMetadataCluster

`func (o *MetadataService) GetMetadataCluster() MetadataClusterNestview`

GetMetadataCluster returns the MetadataCluster field if non-nil, zero value otherwise.

### GetMetadataClusterOk

`func (o *MetadataService) GetMetadataClusterOk() (*MetadataClusterNestview, bool)`

GetMetadataClusterOk returns a tuple with the MetadataCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCluster

`func (o *MetadataService) SetMetadataCluster(v MetadataClusterNestview)`

SetMetadataCluster sets MetadataCluster field to given value.

### HasMetadataCluster

`func (o *MetadataService) HasMetadataCluster() bool`

HasMetadataCluster returns a boolean if a field has been set.

### GetName

`func (o *MetadataService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataService) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataService) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacementNode

`func (o *MetadataService) GetPlacementNode() PlacementNodeNestview`

GetPlacementNode returns the PlacementNode field if non-nil, zero value otherwise.

### GetPlacementNodeOk

`func (o *MetadataService) GetPlacementNodeOk() (*PlacementNodeNestview, bool)`

GetPlacementNodeOk returns a tuple with the PlacementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementNode

`func (o *MetadataService) SetPlacementNode(v PlacementNodeNestview)`

SetPlacementNode sets PlacementNode field to given value.

### HasPlacementNode

`func (o *MetadataService) HasPlacementNode() bool`

HasPlacementNode returns a boolean if a field has been set.

### GetStatus

`func (o *MetadataService) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetadataService) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MetadataService) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MetadataService) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUp

`func (o *MetadataService) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *MetadataService) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *MetadataService) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *MetadataService) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUpdate

`func (o *MetadataService) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *MetadataService) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *MetadataService) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *MetadataService) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetXmdsId

`func (o *MetadataService) GetXmdsId() int64`

GetXmdsId returns the XmdsId field if non-nil, zero value otherwise.

### GetXmdsIdOk

`func (o *MetadataService) GetXmdsIdOk() (*int64, bool)`

GetXmdsIdOk returns a tuple with the XmdsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmdsId

`func (o *MetadataService) SetXmdsId(v int64)`

SetXmdsId sets XmdsId field to given value.

### HasXmdsId

`func (o *MetadataService) HasXmdsId() bool`

HasXmdsId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


