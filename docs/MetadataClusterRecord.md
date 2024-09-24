# MetadataClusterRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**ClusterName** | Pointer to **string** | name in ceph | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsRootfs** | Pointer to [**DfsRootfsNestview**](DfsRootfsNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MetadataServiceNum** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | name for user | [optional] 
**Primary** | Pointer to [**MetadataServiceNestview**](MetadataServiceNestview.md) |  | [optional] 
**PrimaryPlacementNode** | Pointer to [**PlacementNodeNestview**](PlacementNodeNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Stretched** | Pointer to **bool** |  | [optional] 
**TransLocator** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Samples** | Pointer to [**[]MetadataClusterStat**](MetadataClusterStat.md) |  | [optional] 

## Methods

### NewMetadataClusterRecord

`func NewMetadataClusterRecord() *MetadataClusterRecord`

NewMetadataClusterRecord instantiates a new MetadataClusterRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataClusterRecordWithDefaults

`func NewMetadataClusterRecordWithDefaults() *MetadataClusterRecord`

NewMetadataClusterRecordWithDefaults instantiates a new MetadataClusterRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *MetadataClusterRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *MetadataClusterRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *MetadataClusterRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *MetadataClusterRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *MetadataClusterRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *MetadataClusterRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *MetadataClusterRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *MetadataClusterRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterName

`func (o *MetadataClusterRecord) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *MetadataClusterRecord) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *MetadataClusterRecord) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *MetadataClusterRecord) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCreate

`func (o *MetadataClusterRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *MetadataClusterRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *MetadataClusterRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *MetadataClusterRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *MetadataClusterRecord) GetDfsRootfs() DfsRootfsNestview`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *MetadataClusterRecord) GetDfsRootfsOk() (*DfsRootfsNestview, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *MetadataClusterRecord) SetDfsRootfs(v DfsRootfsNestview)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *MetadataClusterRecord) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetId

`func (o *MetadataClusterRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataClusterRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataClusterRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataClusterRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadataServiceNum

`func (o *MetadataClusterRecord) GetMetadataServiceNum() int64`

GetMetadataServiceNum returns the MetadataServiceNum field if non-nil, zero value otherwise.

### GetMetadataServiceNumOk

`func (o *MetadataClusterRecord) GetMetadataServiceNumOk() (*int64, bool)`

GetMetadataServiceNumOk returns a tuple with the MetadataServiceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataServiceNum

`func (o *MetadataClusterRecord) SetMetadataServiceNum(v int64)`

SetMetadataServiceNum sets MetadataServiceNum field to given value.

### HasMetadataServiceNum

`func (o *MetadataClusterRecord) HasMetadataServiceNum() bool`

HasMetadataServiceNum returns a boolean if a field has been set.

### GetName

`func (o *MetadataClusterRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataClusterRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataClusterRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataClusterRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimary

`func (o *MetadataClusterRecord) GetPrimary() MetadataServiceNestview`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *MetadataClusterRecord) GetPrimaryOk() (*MetadataServiceNestview, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *MetadataClusterRecord) SetPrimary(v MetadataServiceNestview)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *MetadataClusterRecord) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetPrimaryPlacementNode

`func (o *MetadataClusterRecord) GetPrimaryPlacementNode() PlacementNodeNestview`

GetPrimaryPlacementNode returns the PrimaryPlacementNode field if non-nil, zero value otherwise.

### GetPrimaryPlacementNodeOk

`func (o *MetadataClusterRecord) GetPrimaryPlacementNodeOk() (*PlacementNodeNestview, bool)`

GetPrimaryPlacementNodeOk returns a tuple with the PrimaryPlacementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPlacementNode

`func (o *MetadataClusterRecord) SetPrimaryPlacementNode(v PlacementNodeNestview)`

SetPrimaryPlacementNode sets PrimaryPlacementNode field to given value.

### HasPrimaryPlacementNode

`func (o *MetadataClusterRecord) HasPrimaryPlacementNode() bool`

HasPrimaryPlacementNode returns a boolean if a field has been set.

### GetStatus

`func (o *MetadataClusterRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetadataClusterRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MetadataClusterRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MetadataClusterRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStretched

`func (o *MetadataClusterRecord) GetStretched() bool`

GetStretched returns the Stretched field if non-nil, zero value otherwise.

### GetStretchedOk

`func (o *MetadataClusterRecord) GetStretchedOk() (*bool, bool)`

GetStretchedOk returns a tuple with the Stretched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStretched

`func (o *MetadataClusterRecord) SetStretched(v bool)`

SetStretched sets Stretched field to given value.

### HasStretched

`func (o *MetadataClusterRecord) HasStretched() bool`

HasStretched returns a boolean if a field has been set.

### GetTransLocator

`func (o *MetadataClusterRecord) GetTransLocator() string`

GetTransLocator returns the TransLocator field if non-nil, zero value otherwise.

### GetTransLocatorOk

`func (o *MetadataClusterRecord) GetTransLocatorOk() (*string, bool)`

GetTransLocatorOk returns a tuple with the TransLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransLocator

`func (o *MetadataClusterRecord) SetTransLocator(v string)`

SetTransLocator sets TransLocator field to given value.

### HasTransLocator

`func (o *MetadataClusterRecord) HasTransLocator() bool`

HasTransLocator returns a boolean if a field has been set.

### GetUpdate

`func (o *MetadataClusterRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *MetadataClusterRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *MetadataClusterRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *MetadataClusterRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *MetadataClusterRecord) GetSamples() []MetadataClusterStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *MetadataClusterRecord) GetSamplesOk() (*[]MetadataClusterStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *MetadataClusterRecord) SetSamples(v []MetadataClusterStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *MetadataClusterRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


