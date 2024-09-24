# DfsNFSShareACL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllSquash** | Pointer to **bool** |  | [optional] 
**Clients** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsNfsShare** | Pointer to [**DfsNFSShareNestview**](DfsNFSShareNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**RootSquash** | Pointer to **bool** |  | [optional] 
**Sync** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsNFSShareACL

`func NewDfsNFSShareACL() *DfsNFSShareACL`

NewDfsNFSShareACL instantiates a new DfsNFSShareACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsNFSShareACLWithDefaults

`func NewDfsNFSShareACLWithDefaults() *DfsNFSShareACL`

NewDfsNFSShareACLWithDefaults instantiates a new DfsNFSShareACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllSquash

`func (o *DfsNFSShareACL) GetAllSquash() bool`

GetAllSquash returns the AllSquash field if non-nil, zero value otherwise.

### GetAllSquashOk

`func (o *DfsNFSShareACL) GetAllSquashOk() (*bool, bool)`

GetAllSquashOk returns a tuple with the AllSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSquash

`func (o *DfsNFSShareACL) SetAllSquash(v bool)`

SetAllSquash sets AllSquash field to given value.

### HasAllSquash

`func (o *DfsNFSShareACL) HasAllSquash() bool`

HasAllSquash returns a boolean if a field has been set.

### GetClients

`func (o *DfsNFSShareACL) GetClients() string`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *DfsNFSShareACL) GetClientsOk() (*string, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *DfsNFSShareACL) SetClients(v string)`

SetClients sets Clients field to given value.

### HasClients

`func (o *DfsNFSShareACL) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetCluster

`func (o *DfsNFSShareACL) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsNFSShareACL) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsNFSShareACL) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsNFSShareACL) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsNFSShareACL) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsNFSShareACL) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsNFSShareACL) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsNFSShareACL) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsNfsShare

`func (o *DfsNFSShareACL) GetDfsNfsShare() DfsNFSShareNestview`

GetDfsNfsShare returns the DfsNfsShare field if non-nil, zero value otherwise.

### GetDfsNfsShareOk

`func (o *DfsNFSShareACL) GetDfsNfsShareOk() (*DfsNFSShareNestview, bool)`

GetDfsNfsShareOk returns a tuple with the DfsNfsShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsNfsShare

`func (o *DfsNFSShareACL) SetDfsNfsShare(v DfsNFSShareNestview)`

SetDfsNfsShare sets DfsNfsShare field to given value.

### HasDfsNfsShare

`func (o *DfsNFSShareACL) HasDfsNfsShare() bool`

HasDfsNfsShare returns a boolean if a field has been set.

### GetId

`func (o *DfsNFSShareACL) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsNFSShareACL) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsNFSShareACL) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsNFSShareACL) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermission

`func (o *DfsNFSShareACL) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *DfsNFSShareACL) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *DfsNFSShareACL) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *DfsNFSShareACL) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetRootSquash

`func (o *DfsNFSShareACL) GetRootSquash() bool`

GetRootSquash returns the RootSquash field if non-nil, zero value otherwise.

### GetRootSquashOk

`func (o *DfsNFSShareACL) GetRootSquashOk() (*bool, bool)`

GetRootSquashOk returns a tuple with the RootSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSquash

`func (o *DfsNFSShareACL) SetRootSquash(v bool)`

SetRootSquash sets RootSquash field to given value.

### HasRootSquash

`func (o *DfsNFSShareACL) HasRootSquash() bool`

HasRootSquash returns a boolean if a field has been set.

### GetSync

`func (o *DfsNFSShareACL) GetSync() bool`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *DfsNFSShareACL) GetSyncOk() (*bool, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *DfsNFSShareACL) SetSync(v bool)`

SetSync sets Sync field to given value.

### HasSync

`func (o *DfsNFSShareACL) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetType

`func (o *DfsNFSShareACL) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DfsNFSShareACL) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DfsNFSShareACL) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DfsNFSShareACL) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsNFSShareACL) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsNFSShareACL) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsNFSShareACL) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsNFSShareACL) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


