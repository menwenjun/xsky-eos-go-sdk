# DfsXDCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsPath** | Pointer to [**DfsPathNestview**](DfsPathNestview.md) |  | [optional] 
**DfsRootfs** | Pointer to [**DfsRootfsNestview**](DfsRootfsNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InheritanceType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsXDCache

`func NewDfsXDCache() *DfsXDCache`

NewDfsXDCache instantiates a new DfsXDCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsXDCacheWithDefaults

`func NewDfsXDCacheWithDefaults() *DfsXDCache`

NewDfsXDCacheWithDefaults instantiates a new DfsXDCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DfsXDCache) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsXDCache) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsXDCache) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsXDCache) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsXDCache) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsXDCache) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsXDCache) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsXDCache) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsPath

`func (o *DfsXDCache) GetDfsPath() DfsPathNestview`

GetDfsPath returns the DfsPath field if non-nil, zero value otherwise.

### GetDfsPathOk

`func (o *DfsXDCache) GetDfsPathOk() (*DfsPathNestview, bool)`

GetDfsPathOk returns a tuple with the DfsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPath

`func (o *DfsXDCache) SetDfsPath(v DfsPathNestview)`

SetDfsPath sets DfsPath field to given value.

### HasDfsPath

`func (o *DfsXDCache) HasDfsPath() bool`

HasDfsPath returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsXDCache) GetDfsRootfs() DfsRootfsNestview`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsXDCache) GetDfsRootfsOk() (*DfsRootfsNestview, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsXDCache) SetDfsRootfs(v DfsRootfsNestview)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsXDCache) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetId

`func (o *DfsXDCache) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsXDCache) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsXDCache) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsXDCache) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceType

`func (o *DfsXDCache) GetInheritanceType() string`

GetInheritanceType returns the InheritanceType field if non-nil, zero value otherwise.

### GetInheritanceTypeOk

`func (o *DfsXDCache) GetInheritanceTypeOk() (*string, bool)`

GetInheritanceTypeOk returns a tuple with the InheritanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceType

`func (o *DfsXDCache) SetInheritanceType(v string)`

SetInheritanceType sets InheritanceType field to given value.

### HasInheritanceType

`func (o *DfsXDCache) HasInheritanceType() bool`

HasInheritanceType returns a boolean if a field has been set.

### GetStatus

`func (o *DfsXDCache) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsXDCache) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsXDCache) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsXDCache) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsXDCache) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsXDCache) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsXDCache) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsXDCache) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


