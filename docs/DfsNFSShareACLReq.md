# DfsNFSShareACLReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllSquash** | Pointer to **bool** | all squash | [optional] 
**Clients** | Pointer to **string** | ip or network list, separated by comma | [optional] 
**Permission** | Pointer to **string** | readonly or readwrite access | [optional] 
**RootSquash** | Pointer to **bool** | root squash | [optional] 
**Sync** | Pointer to **bool** | write to disk by sync or async | [optional] 
**Type** | Pointer to **string** | type of share acl | [optional] 

## Methods

### NewDfsNFSShareACLReq

`func NewDfsNFSShareACLReq() *DfsNFSShareACLReq`

NewDfsNFSShareACLReq instantiates a new DfsNFSShareACLReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsNFSShareACLReqWithDefaults

`func NewDfsNFSShareACLReqWithDefaults() *DfsNFSShareACLReq`

NewDfsNFSShareACLReqWithDefaults instantiates a new DfsNFSShareACLReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllSquash

`func (o *DfsNFSShareACLReq) GetAllSquash() bool`

GetAllSquash returns the AllSquash field if non-nil, zero value otherwise.

### GetAllSquashOk

`func (o *DfsNFSShareACLReq) GetAllSquashOk() (*bool, bool)`

GetAllSquashOk returns a tuple with the AllSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSquash

`func (o *DfsNFSShareACLReq) SetAllSquash(v bool)`

SetAllSquash sets AllSquash field to given value.

### HasAllSquash

`func (o *DfsNFSShareACLReq) HasAllSquash() bool`

HasAllSquash returns a boolean if a field has been set.

### GetClients

`func (o *DfsNFSShareACLReq) GetClients() string`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *DfsNFSShareACLReq) GetClientsOk() (*string, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *DfsNFSShareACLReq) SetClients(v string)`

SetClients sets Clients field to given value.

### HasClients

`func (o *DfsNFSShareACLReq) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetPermission

`func (o *DfsNFSShareACLReq) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *DfsNFSShareACLReq) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *DfsNFSShareACLReq) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *DfsNFSShareACLReq) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetRootSquash

`func (o *DfsNFSShareACLReq) GetRootSquash() bool`

GetRootSquash returns the RootSquash field if non-nil, zero value otherwise.

### GetRootSquashOk

`func (o *DfsNFSShareACLReq) GetRootSquashOk() (*bool, bool)`

GetRootSquashOk returns a tuple with the RootSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSquash

`func (o *DfsNFSShareACLReq) SetRootSquash(v bool)`

SetRootSquash sets RootSquash field to given value.

### HasRootSquash

`func (o *DfsNFSShareACLReq) HasRootSquash() bool`

HasRootSquash returns a boolean if a field has been set.

### GetSync

`func (o *DfsNFSShareACLReq) GetSync() bool`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *DfsNFSShareACLReq) GetSyncOk() (*bool, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *DfsNFSShareACLReq) SetSync(v bool)`

SetSync sets Sync field to given value.

### HasSync

`func (o *DfsNFSShareACLReq) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetType

`func (o *DfsNFSShareACLReq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DfsNFSShareACLReq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DfsNFSShareACLReq) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DfsNFSShareACLReq) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


