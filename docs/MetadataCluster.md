# MetadataCluster

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

## Methods

### NewMetadataCluster

`func NewMetadataCluster() *MetadataCluster`

NewMetadataCluster instantiates a new MetadataCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataClusterWithDefaults

`func NewMetadataClusterWithDefaults() *MetadataCluster`

NewMetadataClusterWithDefaults instantiates a new MetadataCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *MetadataCluster) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *MetadataCluster) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *MetadataCluster) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *MetadataCluster) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *MetadataCluster) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *MetadataCluster) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *MetadataCluster) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *MetadataCluster) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterName

`func (o *MetadataCluster) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *MetadataCluster) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *MetadataCluster) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *MetadataCluster) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCreate

`func (o *MetadataCluster) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *MetadataCluster) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *MetadataCluster) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *MetadataCluster) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *MetadataCluster) GetDfsRootfs() DfsRootfsNestview`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *MetadataCluster) GetDfsRootfsOk() (*DfsRootfsNestview, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *MetadataCluster) SetDfsRootfs(v DfsRootfsNestview)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *MetadataCluster) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetId

`func (o *MetadataCluster) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataCluster) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataCluster) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadataServiceNum

`func (o *MetadataCluster) GetMetadataServiceNum() int64`

GetMetadataServiceNum returns the MetadataServiceNum field if non-nil, zero value otherwise.

### GetMetadataServiceNumOk

`func (o *MetadataCluster) GetMetadataServiceNumOk() (*int64, bool)`

GetMetadataServiceNumOk returns a tuple with the MetadataServiceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataServiceNum

`func (o *MetadataCluster) SetMetadataServiceNum(v int64)`

SetMetadataServiceNum sets MetadataServiceNum field to given value.

### HasMetadataServiceNum

`func (o *MetadataCluster) HasMetadataServiceNum() bool`

HasMetadataServiceNum returns a boolean if a field has been set.

### GetName

`func (o *MetadataCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimary

`func (o *MetadataCluster) GetPrimary() MetadataServiceNestview`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *MetadataCluster) GetPrimaryOk() (*MetadataServiceNestview, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *MetadataCluster) SetPrimary(v MetadataServiceNestview)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *MetadataCluster) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetPrimaryPlacementNode

`func (o *MetadataCluster) GetPrimaryPlacementNode() PlacementNodeNestview`

GetPrimaryPlacementNode returns the PrimaryPlacementNode field if non-nil, zero value otherwise.

### GetPrimaryPlacementNodeOk

`func (o *MetadataCluster) GetPrimaryPlacementNodeOk() (*PlacementNodeNestview, bool)`

GetPrimaryPlacementNodeOk returns a tuple with the PrimaryPlacementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPlacementNode

`func (o *MetadataCluster) SetPrimaryPlacementNode(v PlacementNodeNestview)`

SetPrimaryPlacementNode sets PrimaryPlacementNode field to given value.

### HasPrimaryPlacementNode

`func (o *MetadataCluster) HasPrimaryPlacementNode() bool`

HasPrimaryPlacementNode returns a boolean if a field has been set.

### GetStatus

`func (o *MetadataCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetadataCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MetadataCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MetadataCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStretched

`func (o *MetadataCluster) GetStretched() bool`

GetStretched returns the Stretched field if non-nil, zero value otherwise.

### GetStretchedOk

`func (o *MetadataCluster) GetStretchedOk() (*bool, bool)`

GetStretchedOk returns a tuple with the Stretched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStretched

`func (o *MetadataCluster) SetStretched(v bool)`

SetStretched sets Stretched field to given value.

### HasStretched

`func (o *MetadataCluster) HasStretched() bool`

HasStretched returns a boolean if a field has been set.

### GetTransLocator

`func (o *MetadataCluster) GetTransLocator() string`

GetTransLocator returns the TransLocator field if non-nil, zero value otherwise.

### GetTransLocatorOk

`func (o *MetadataCluster) GetTransLocatorOk() (*string, bool)`

GetTransLocatorOk returns a tuple with the TransLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransLocator

`func (o *MetadataCluster) SetTransLocator(v string)`

SetTransLocator sets TransLocator field to given value.

### HasTransLocator

`func (o *MetadataCluster) HasTransLocator() bool`

HasTransLocator returns a boolean if a field has been set.

### GetUpdate

`func (o *MetadataCluster) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *MetadataCluster) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *MetadataCluster) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *MetadataCluster) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


