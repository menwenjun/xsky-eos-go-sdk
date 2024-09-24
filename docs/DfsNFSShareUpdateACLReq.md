# DfsNFSShareUpdateACLReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllSquash** | Pointer to **bool** | all squash | [optional] 
**Clients** | Pointer to **string** | ip or network list, separated by comma | [optional] 
**Id** | Pointer to **int64** | id of user group | [optional] 
**Permission** | Pointer to **string** | readonly or readwrite access | [optional] 
**RootSquash** | Pointer to **bool** | root squash | [optional] 
**Sync** | Pointer to **bool** | write to disk by sync or async | [optional] 

## Methods

### NewDfsNFSShareUpdateACLReq

`func NewDfsNFSShareUpdateACLReq() *DfsNFSShareUpdateACLReq`

NewDfsNFSShareUpdateACLReq instantiates a new DfsNFSShareUpdateACLReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsNFSShareUpdateACLReqWithDefaults

`func NewDfsNFSShareUpdateACLReqWithDefaults() *DfsNFSShareUpdateACLReq`

NewDfsNFSShareUpdateACLReqWithDefaults instantiates a new DfsNFSShareUpdateACLReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllSquash

`func (o *DfsNFSShareUpdateACLReq) GetAllSquash() bool`

GetAllSquash returns the AllSquash field if non-nil, zero value otherwise.

### GetAllSquashOk

`func (o *DfsNFSShareUpdateACLReq) GetAllSquashOk() (*bool, bool)`

GetAllSquashOk returns a tuple with the AllSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSquash

`func (o *DfsNFSShareUpdateACLReq) SetAllSquash(v bool)`

SetAllSquash sets AllSquash field to given value.

### HasAllSquash

`func (o *DfsNFSShareUpdateACLReq) HasAllSquash() bool`

HasAllSquash returns a boolean if a field has been set.

### GetClients

`func (o *DfsNFSShareUpdateACLReq) GetClients() string`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *DfsNFSShareUpdateACLReq) GetClientsOk() (*string, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *DfsNFSShareUpdateACLReq) SetClients(v string)`

SetClients sets Clients field to given value.

### HasClients

`func (o *DfsNFSShareUpdateACLReq) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetId

`func (o *DfsNFSShareUpdateACLReq) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsNFSShareUpdateACLReq) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsNFSShareUpdateACLReq) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsNFSShareUpdateACLReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermission

`func (o *DfsNFSShareUpdateACLReq) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *DfsNFSShareUpdateACLReq) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *DfsNFSShareUpdateACLReq) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *DfsNFSShareUpdateACLReq) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetRootSquash

`func (o *DfsNFSShareUpdateACLReq) GetRootSquash() bool`

GetRootSquash returns the RootSquash field if non-nil, zero value otherwise.

### GetRootSquashOk

`func (o *DfsNFSShareUpdateACLReq) GetRootSquashOk() (*bool, bool)`

GetRootSquashOk returns a tuple with the RootSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSquash

`func (o *DfsNFSShareUpdateACLReq) SetRootSquash(v bool)`

SetRootSquash sets RootSquash field to given value.

### HasRootSquash

`func (o *DfsNFSShareUpdateACLReq) HasRootSquash() bool`

HasRootSquash returns a boolean if a field has been set.

### GetSync

`func (o *DfsNFSShareUpdateACLReq) GetSync() bool`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *DfsNFSShareUpdateACLReq) GetSyncOk() (*bool, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *DfsNFSShareUpdateACLReq) SetSync(v bool)`

SetSync sets Sync field to given value.

### HasSync

`func (o *DfsNFSShareUpdateACLReq) HasSync() bool`

HasSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


